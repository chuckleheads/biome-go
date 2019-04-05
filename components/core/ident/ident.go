package ident

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
)

// Ident represents a package identifier
type Ident struct {
	Origin  string          `mapstructure:"origin"`
	Name    string          `mapstructure:"name"`
	Version *semver.Version `mapstructure:"version"`
	Release string          `mapstructure:"release"`
}

// New creates a new instance of an Ident from its parts
func New(origin string, name string, version string, release string) (Ident, error) {
	newIdent := Ident{}
	if origin == "" || name == "" {
		errors.Errorf("Origin and Name are required")
	}
	newIdent.Origin = origin
	newIdent.Name = name
	if version != "" {
		newVersion, err := semver.NewVersion(version)
		if err != nil {
			return Ident{}, err
		}
		newIdent.Version = newVersion
	}
	if release != "" {
		newIdent.Release = release
	}
	return newIdent, nil
}

// FromString takes a string in the form of ORIGIN/NAME/VERSION/RELEASE
// and resturns a HabPkg struct
func FromString(rawIdent string) (Ident, error) {
	pkg := Ident{}
	return pkg.FromString(rawIdent)
}

// FromString takes a string in the form of ORIGIN/NAME/VERSION/RELEASE
// and resturns a Ident struct
func (ident Ident) FromString(rawIdent string) (Ident, error) {
	parts := strings.Split(rawIdent, "/")
	if len(parts) < 2 || len(parts) > 4 {
		return Ident{}, errors.Errorf("invalid package identifier '%s'", rawIdent)
	}

	for i, part := range parts {
		switch i {
		case 0:
			ident.Origin = part
		case 1:
			ident.Name = part
		case 2:
			newVersion, err := semver.NewVersion(part)
			if err != nil {
				return Ident{}, err
			}
			ident.Version = newVersion
		case 3:
			ident.Release = part
		}
	}

	return ident, nil
}

// String makes Ident satisfy the Stringer interface.
func (ident Ident) String() string {
	formattedIdent := filepath.Join(ident.Origin, ident.Name)
	if ident.Version != nil {
		formattedIdent = filepath.Join(formattedIdent, ident.Version.Original())
	}
	if ident.Release != "" {
		formattedIdent = filepath.Join(formattedIdent, ident.Release)
	}
	return formattedIdent
}

// ArchiveName converts an ident struct to a PackageArchive filename
func (ident Ident) ArchiveName() (string, error) {
	// TODO: Set a sane default
	return ident.archiveName("x86_64-linux")
}

// TODO: Implement PackageTarget
func (ident Ident) archiveName(target string) (string, error) {
	if ident.FullyQualified() {
		return fmt.Sprintf("%s-%s-%s-%s-%s.hart", ident.Origin, ident.Name, ident.Version.Original(), ident.Release, target), nil
	}
	return "", fmt.Errorf("Fully-qualified package identifier was expected, but found: %s", ident)
}

// FullyQualified determines if a given ident is a fully qualified package identifier
func (ident Ident) FullyQualified() bool {
	if ident.Origin != "" && ident.Name != "" && ident.Version != nil && ident.Release != "" {
		return true
	}
	return false
}

// Compare two packages according to the following:
//
// * origin is ignored in the comparison - my redis and
//   your redis compare the same.
// * If the names are not equal, they cannot be compared.
// * If the versions are greater/lesser,
//   return -1 if self is lesser, return +1 if the other is lesser
// * If the versions are equal, return the greater/lesser
//   return -1 if self is lesser, return +1 if the other is lesser.
func (ident Ident) Compare(other Ident) (int, error) {
	// Names are different - bail
	if ident.Name != other.Name {
		return -2, fmt.Errorf("Different package names can't be compared")
	}

	if ident.Version == nil && other.Version != nil {
		return 1, nil
	} else if ident.Version != nil && other.Version == nil {
		return -1, nil
	} else if ident.Version == nil && other.Version == nil {
		return 0, nil
	}

	// We know we have a version now for both idents, lets compare them
	// Leverage semver package to do the heavy lifting for version comparison.
	// If they are the same version, continue
	if verCmp := ident.Version.Compare(other.Version); verCmp != 0 {
		return verCmp, nil
	}

	if ident.Release == "" && other.Release != "" {
		return 1, nil
	} else if ident.Release != "" && other.Release == "" {
		return -1, nil
	} else if ident.Release == "" && other.Release == "" {
		return 0, nil
	}

	myRelease, err := strconv.Atoi(ident.Release)
	if err != nil {
		return -2, fmt.Errorf("failed to parse my release version for comparison")
	}
	otherRelease, err := strconv.Atoi(other.Release)
	if err != nil {
		return -2, fmt.Errorf("failed to parse other release version for comparison")
	}

	if myRelease < otherRelease {
		return -1, nil
	} else if myRelease > otherRelease {
		return 1, nil
	}
	return 0, nil
}
