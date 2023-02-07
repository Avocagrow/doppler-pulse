package pulse

// ClientOptionFunc can be used to set the client options for a Doppler client.
type ClientOptionFunc func(*Client) error

func WithBaseURL(uri string) ClientOptionFunc {
	return func(c *Client) error {
		return c.baseUrl.setBaseURL(uri)
	}
}
