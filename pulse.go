// Package pulse is implements a Doppler API client.
package pulse

import (
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-cleanhttp"
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

func NewClient(token string, options ...ClientOptionFunc) (*Client, error) {
	c, err := newClient(options...)
	if err != nil {
		return nil, err
	}
	c.token = token
	return c, nil
}

func (c *Client) BaseURL() *url.URL {
	u := *c.baseUrl
	return &u
}

func newClient(options ...ClientOptionFunc) (*Client, error) {
	c := &Client{}
	c.client = &retryablehttp.Client{
		ErrorHandler: retryablehttp.PassthroughErrorHandler,
		HTTPClient:   cleanhttp.DefaultPooledClient(),
		RetryWaitMin: 100 * time.Millisecond,
		RetryWaitMax: 400 * time.Millisecond,
		RetryMax:     5,
	}

	c.setBaseURL(defaultBaseURL)

	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}
