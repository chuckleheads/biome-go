package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/biome-sh/biome-go/components/builder-depot-client/types"
	pkg "github.com/biome-sh/biome-go/components/core/ident"
)

// ScheduleJob schedules a new job with builder for a given package
func (c *Client) ScheduleJob(ident pkg.Ident, group bool) error {
	res, err := c.req.SetPathParams(map[string]string{
		"origin": ident.Origin,
		"name":   ident.Name,
	}).
		SetQueryParam("package_only", fmt.Sprintf("%t", group)).
		Post("depot/pkgs/schedule/{origin}/{name}")

	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		return errors.New(res.Status())
	}
	return nil
}

// FetchRdeps fetches reverse dependencies for a given package ident
func (c *Client) FetchRdeps(ident pkg.Ident) ([]pkg.Ident, error) {
	test, err := c.req.SetPathParams(map[string]string{
		"ident": ident.String(),
	}).Get("/rdeps/{ident}")

	if err != nil {
		return []pkg.Ident{}, err
	}
	var res types.RDepsResp
	err = json.Unmarshal(test.Body(), &res)
	if err != nil {
		return []pkg.Ident{}, err
	}
	var rdeps []pkg.Ident
	for _, rdepString := range res.RDeps {
		rdep, err := pkg.FromString(rdepString)
		if err != nil {
			return []pkg.Ident{}, err
		}
		rdeps = append(rdeps, rdep)
	}

	return rdeps, nil
}
