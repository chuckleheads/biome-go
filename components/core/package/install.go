package pkg

import (
	"fmt"
	"sort"

	"github.com/google/go-cmp/cmp"

	"github.com/biome-sh/biome-go/components/core/ident"
)

type PackageInstall struct {
	InstalledPath string
	Ident         ident.Ident
}

func Load(pkg ident.Ident, fsRoot string) (*PackageInstall, error) {
	return resolvePackageInstall(pkg, fsRoot)
}

func resolvePackageInstall(pkg ident.Ident, fsRoot string) (*PackageInstall, error) {
	installedIdents := packageListForIdent(pkg, fsRoot)

	if len(installedIdents) == 0 {
		return nil, fmt.Errorf("No installed packages found")
	}

	if pkg.FullyQualified() {
		for _, idt := range installedIdents {
			if cmp.Equal(idt, pkg) {
				return &PackageInstall{
					InstalledPath: fmt.Sprintf("%s/%s", fsRoot, pkg.String()),
					Ident:         pkg,
				}, nil
			}
		}
		return nil, fmt.Errorf("No installed package found")
	}

	sort.Slice(installedIdents, func(i, j int) bool {
		isLess, err := installedIdents[i].Compare(installedIdents[j])
		if err != nil {
			return false
		}
		if isLess == -1 {
			return true
		}
		return false
	})

	latestInstalledIdent := installedIdents[len(installedIdents)-1]

	return &PackageInstall{
		InstalledPath: fmt.Sprintf("%s/%s", fsRoot, latestInstalledIdent.String()),
		Ident:         latestInstalledIdent,
	}, nil
}
