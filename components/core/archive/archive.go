package archive

import (
	"archive/tar"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/biome-sh/biome-go/components/core/ident"
	"github.com/ulikunitz/xz"
)

type PackageArchive struct {
	path string
}

func metafileRegex() map[Metafile]*regexp.Regexp {
	// TODO: when hab packages are biome packages, rename this?
	return map[Metafile]*regexp.Regexp{
		CFlags:           regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", CFlags)),
		Deps:             regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", Deps)),
		TDeps:            regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", TDeps)),
		Exposes:          regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", Exposes)),
		Ident:            regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", Ident)),
		LdRunPath:        regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", LdRunPath)),
		LdFlags:          regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", LdFlags)),
		SvcUser:          regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", SvcUser)),
		Services:         regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", Services)),
		ResolvedServices: regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", ResolvedServices)),
		Manifest:         regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", Manifest)),
		Path:             regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", Path)),
		Target:           regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", Target)),
		Type:             regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", Type)),
		Config:           regexp.MustCompile(fmt.Sprintf("^/?hab/pkgs/([^/]+)/([^/]+)/([^/]+)/([^/]+)/%s$", Config)),
	}
}

// New takes a path to a hart file and returns a PackageArchive struct
func New(path string) PackageArchive {
	return PackageArchive{path}
}

// GetMetadata returns the metadata files for a given archive
// This means that each Metafile is a key to the stringified contents of the files
func (archive *PackageArchive) GetMetadata() map[Metafile]string {
	metadata := make(map[Metafile]string)
	file, err := os.Open(archive.path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	// Strip first 5 lines
	for i := 0; i < 5; i++ {
		_, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
	}

	xzf, err := xz.NewReader(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tarReader := tar.NewReader(xzf)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		name := header.Name
		switch header.Typeflag {
		case tar.TypeDir:
			break
		case tar.TypeReg:
			for metafile, reg := range metafileRegex() {
				if reg.MatchString(name) {
					data := make([]byte, header.Size-1)
					_, err := tarReader.Read(data)
					if err != nil {
						panic("Error reading file!!! PANIC!!!!!!")
					}
					metadata[metafile] = string(data[:header.Size-1])
					break
				}
			}
		}
		if len(metadata) == len(metafileRegex()) {
			break
		}
	}
	return metadata
}

// Unpack takes a package archive and unpacks it into an fsroot (defaults to /)
func (archive *PackageArchive) Unpack(fsRoot string) error {
	if fsRoot == "" {
		fsRoot = "/"
	}
	file, err := os.Open(archive.path)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	// Strip first 5 lines
	for i := 0; i < 5; i++ {
		_, err = reader.ReadString('\n')
		if err != nil {
			return err
		}
	}

	xzf, err := xz.NewReader(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tarReader := tar.NewReader(xzf)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		target := filepath.Join(fsRoot, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg, tar.TypeLink, tar.TypeSymlink:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(f, tarReader); err != nil {
				return err
			}

			f.Close()
		}
	}
	return nil
}

// Ident returns a package identifier for an archive
func (archive *PackageArchive) Ident() (ident.Ident, error) {
	return ident.FromString(archive.GetMetadata()[Ident])
}

// Deps returns a list of dependencies for a given archive
func (archive *PackageArchive) Deps() ([]ident.Ident, error) {
	return archive.readIdents(Deps)
}

// TDeps returns a list of dependencies for a given archive
func (archive *PackageArchive) TDeps() ([]ident.Ident, error) {
	return archive.readIdents(TDeps)
}

// PkgServices returns a list of services for a given archive
func (archive *PackageArchive) PkgServices() ([]ident.Ident, error) {
	return archive.readIdents(Services)
}

// ResolvedServices returns a list of resolved services for a given archive
func (archive *PackageArchive) ResolvedServices() ([]ident.Ident, error) {
	return archive.readIdents(ResolvedServices)
}

// readDeps reads the package dependencies for a given Metafile type
func (archive *PackageArchive) readIdents(file Metafile) ([]ident.Ident, error) {
	idents := []ident.Ident{}
	for _, dep := range strings.Split(archive.GetMetadata()[file], "\n") {
		depIdent, err := ident.FromString(dep)
		if err != nil {
			return idents, err
		}
		idents = append(idents, depIdent)
	}
	return idents, nil
}

type ArchiveHeader struct {
	FormatVersion string
	KeyName       string
	HashType      string
	SignatureRaw  string
}
