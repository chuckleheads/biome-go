package pkg

// feels_bad.jpg ^^

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/biome-sh/biome-go/components/builder-depot-client/client"
	"github.com/biome-sh/biome-go/components/core/archive"
	pkgIdent "github.com/biome-sh/biome-go/components/core/ident"
	corePkg "github.com/biome-sh/biome-go/components/core/package"
	"github.com/biome-sh/biome-go/components/bio/pkg/ui"
)

func FromIdent(ident pkgIdent.Ident, channel string) {
	latest := determineLatestFromIdent(ident, channel)
	isInstalled, _ := corePkg.Load(*latest, viper.GetString("fs_root"))

	if isInstalled != nil {
		ui.Status(ui.Using, latest.String())
		ui.End(fmt.Sprintf("Install of %s complete with %d new packages installed.", latest, 0))
		return
	}

	archive, err := getCachedArtifact(*latest)
	if err != nil {
		ui.Fatal(err)
		os.Exit(1)
	}

	FromArchive(archive)
}

func FromArchive(localArchive archive.PackageArchive) {
	var toInstall []archive.PackageArchive
	tdeps, err := localArchive.TDeps()
	if err != nil {
		ui.Fatal(err)
		os.Exit(1)
	}
	for _, dep := range tdeps {
		depArchive, err := getCachedArtifact(dep)
		if err != nil {
			ui.Fatal(err)
			os.Exit(1)
		}
		isInstalled, _ := corePkg.Load(dep, viper.GetString("fs_root"))

		if isInstalled != nil {
			ui.Status(ui.Using, dep.String())
		} else {
			toInstall = append(toInstall, depArchive)
		}
	}
	toInstall = append(toInstall, localArchive)

	for _, installable := range toInstall {
		if err = installable.Unpack(viper.GetString("fs_root")); err != nil {
			fmt.Println(err)
		}
	}
	// TED need to check for binlink stuff here
	if err != nil {
		ui.Fatal(err)
		os.Exit(1)
	}
	localIdent, err := localArchive.Ident()
	if err != nil {
		ui.Fatal(err)
		os.Exit(1)
	}
	ui.Status(ui.Installed, localIdent.String())
}

// getCachedArtifact verifies that an artifact is in the cache.
// If there isn't an artifact that satisfies the constraints
// It will be fetched from Builder
func getCachedArtifact(ident pkgIdent.Ident) (archive.PackageArchive, error) {
	if isArtifactCached(ident) {
		ui.Status(ui.Found, fmt.Sprint("Artifact found in cache for ident:", ident))
	} else {
		ui.Status(ui.Downloading, ident.String())
		cli := client.New(viper.GetString("bldr_url"), viper.GetString("auth_token"))
		_ = cli.DownloadPackage(ident, viper.GetString("artifact_cache"))
	}
	path, err := cachedArtifactPath(ident)
	if err != nil {
		return archive.PackageArchive{}, err
	}
	return archive.New(path), nil
}

func isArtifactCached(ident pkgIdent.Ident) bool {
	path, err := cachedArtifactPath(ident)
	if err != nil {
		return false
	}
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}

func cachedArtifactPath(ident pkgIdent.Ident) (string, error) {
	name, err := ident.ArchiveName()
	if err != nil {
		return "", err
	}
	return filepath.Join(viper.GetString("artifact_cache"), name), nil
}

func determineLatestFromIdent(ident pkgIdent.Ident, channel string) *pkgIdent.Ident {
	// If we have a fully qualified package identifier, then our work is done--there can
	// only be *one* package that satisfies a fully qualified identifier.
	if ident.FullyQualified() {
		return &ident
	}
	cli := client.New(viper.GetString("bldr_url"), viper.GetString("auth_token"))

	ui.Status(ui.Determining, fmt.Sprintf("latest version of %s in the '%s' channel", ident, channel))
	fqi := cli.ShowPackage(ident, channel)

	latestInstalled, err := corePkg.Load(ident, viper.GetString("fs_root"))

	isLess, err := fqi.Compare(latestInstalled.Ident)
	if err != nil {
		return &fqi
	}
	if isLess < 0 {
		ui.Status(ui.Found, fmt.Sprintf("newer installed version (%s) than remote version (%s)", latestInstalled.Ident, fqi))
		return &latestInstalled.Ident
	}

	return &fqi
}
