package client

import (
	"encoding/json"
	"errors"

	"github.com/biome-sh/biome-go/components/builder-depot-client/types"
)

func (c *Client) CreateChannel(origin string, channel string) error {
	res, err := c.req.SetPathParams(map[string]string{
		"origin":  origin,
		"channel": channel,
	}).Post("depot/channels/{origin}/{channel}")

	if err != nil {
		return err
	}
	if res.StatusCode() != 201 {
		return errors.New(res.Status())
	}
	return nil
}

func (c *Client) DeleteChannel(origin string, channel string) error {
	res, err := c.req.SetPathParams(map[string]string{
		"origin":  origin,
		"channel": channel,
	}).Delete("depot/channels/{origin}/{channel}")

	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		return errors.New(res.Status())
	}
	return nil
}

func (c *Client) ListChannels(origin string) ([]types.ChannelName, error) {
	res, err := c.req.SetPathParams(map[string]string{
		"origin": origin,
	}).Get("depot/channels/{origin}")

	if err != nil {
		return nil, err
	}

	var channels []types.ChannelName
	err = json.Unmarshal(res.Body(), &res)
	if err != nil {
		return nil, err
	}
	return channels, nil
}
