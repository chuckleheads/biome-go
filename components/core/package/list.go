package pkg

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver"

	"github.com/biome-sh/biome-go/components/core/ident"
)

// InstallTmpPrefix is a path for pre-installed bio packages
const InstallTmpPrefix = ".bio-pkg-install"

// packageListForIdent returns a slice of ident structs built from the contents of
// the given directory, using the given ident to restrict the
// search.
//
// The search is restricted by assuming the package directory
// structure is:
//
//    /base/ORIGIN/NAME/VERSION/RELEASE/
//
func packageListForIdent(pkg ident.Ident, fsRoot string) []ident.Ident {
	var idents []ident.Ident

	entrypoint := filepath.Join(fsRoot, "bio", "pkgs", pkg.String())

	// origin/name
	if pkg.Version == nil {
		walkVersions(pkg.Origin, pkg.Name, &idents, entrypoint)
	} else if pkg.Release == "" {
		// origin/name/version
		walkReleases(pkg.Origin, pkg.Name, pkg.Version, &idents, entrypoint)
	} else {
		if _, err := os.Stat(entrypoint); !os.IsNotExist(err) {
			idents = append(idents, pkg)
		}
	}

	return idents
}

func walkVersions(origin string, name string, idents *[]ident.Ident, dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			version, err := semver.NewVersion(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			walkReleases(origin, name, version, idents, filepath.Join(dir, file.Name()))
		}
	}
}

func walkReleases(origin string, name string, version *semver.Version, idents *[]ident.Ident, dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			*idents = append(*idents, ident.Ident{
				Origin:  origin,
				Name:    name,
				Version: version,
				Release: file.Name(),
			})
		}
	}
}
