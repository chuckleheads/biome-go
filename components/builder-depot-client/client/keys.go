package client

import (
	"errors"
)

func (c *Client) FetchOriginSecretKey(origin string) ([]byte, string, error) {
	res, err := c.req.SetPathParams(map[string]string{
		"origin": origin,
	}).
		Get("depot/origins/{origin}/secret_keys/latest")

	if err != nil {
		return nil, "", err
	}
	if res.StatusCode() != 200 {
		return []byte{}, "", errors.New(res.Status())
	}

	return res.Body(), res.Header().Get("X-Filename"), nil
}
