package client

import (
	"fmt"

	"github.com/go-resty/resty"
)

type Client struct {
	req *resty.Request
}

const DEFAULT_API_PATH = "/v1"

func New(baseURL string, token string) Client {
	// resty.SetDebug(true)
	// There is probably a better way to do this
	resty.SetHostURL(fmt.Sprintf("%s/%s", baseURL, DEFAULT_API_PATH))
	resty.SetHeader("Accept", "application/json")
	resty.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "Habicats are GO!",
	})
	return Client{
		req: authzMeMaybe(resty.R(), token),
	}
}

func authzMeMaybe(req *resty.Request, token string) *resty.Request {
	if token != "" {
		return req.SetAuthToken(token)
	}
	return req
}
