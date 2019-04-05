package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"

	"github.com/biome-sh/biome-go/components/builder-depot-client/types"
	"github.com/biome-sh/biome-go/components/core/archive"
	pkg "github.com/biome-sh/biome-go/components/core/ident"
)

func (c *Client) SearchPackage(searchTerm string) []pkg.Ident {
	encodedTerm := url.QueryEscape(searchTerm)
	test, err := c.req.SetPathParams(map[string]string{
		"term": encodedTerm,
	}).Get("depot/pkgs/search/{term}")

	if err != nil {
		fmt.Println(err)
	}
	var res types.PagedSearchResp
	err = json.Unmarshal(test.Body(), &res)
	if err != nil {
		fmt.Println(err)
	}
	return res.Data
}

func (c *Client) DownloadPackage(ident pkg.Ident, artifactCachePath string) archive.PackageArchive {
	tmpFile, err := ioutil.TempFile("/tmp", "tmpHab")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(tmpFile.Name())
	res, err := c.req.SetOutput(tmpFile.Name()).SetPathParams(map[string]string{
		"package": ident.String(),
	}).Get("depot/pkgs/{package}/download")
	if err != nil {
		fmt.Println(err)
	}
	file := res.Header().Get("X-Filename")
	newPath := filepath.Join(artifactCachePath, file)
	// Does this work in windows? Not sure and I don't have a windows test machine
	os.Link(tmpFile.Name(), newPath)
	return archive.New(newPath)
}

// ShowPackage takes an ident and returns the package metadata. Additionally,
// if the ident is not fully qualified, it will return the latest package for
// the provided channel
// TODO - this should actually return a package type but we haven't invented that yet
func (c *Client) ShowPackage(ident pkg.Ident, channel string) pkg.Ident {
	params := map[string]string{
		"origin":  ident.Origin,
		"name":    ident.Name,
		"channel": channel,
	}

	url := "depot/channels/{origin}/{channel}/pkgs/{name}"

	if ident.Version != nil {
		params["version"] = ident.Version.Original()
		url = fmt.Sprintf("%s/{version}", url)
		if ident.Release != "" {
			params["release"] = ident.Release
			url = fmt.Sprintf("%s/{release}", url)
		}
	}

	if !ident.FullyQualified() {
		url = fmt.Sprintf("%s/latest", url)
	}

	test, err := c.req.SetPathParams(params).Get(url)
	if err != nil {
		fmt.Println(err)
	}

	var res types.ShowPackageResp
	err = json.Unmarshal(test.Body(), &res)
	if err != nil {
		fmt.Println(err)
	}
	return res.Ident
}
