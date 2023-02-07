// Package pulse is implements a Doppler API client.
package pulse

import (
	"net/url"
	"strings"
	"sync"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

const (
	defaultBaseURL = "https://api.doppler.com/"
	apiPath        = "v3/"
)

type AuthType int

const (
	BasicAuth AuthType = iota
	PersonalToken
	ServiceToken
	SCIMToken
	AuditToken
)

type Client struct {
	client *retryablehttp.Client

	baseUrl *url.URL

	token string

	authType string

	tokenl sync.RWMutex
}

func (c *Client) setBaseURL(uri string) error {
	if !strings.HasSuffix(uri, "/") {
		uri += "/"
	}

	baseUrl, err := url.Parse(uri)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseUrl.Path, apiPath) {
		baseUrl.Path += apiPath
	}

	c.baseUrl = baseUrl
	return nil
}
