package pulse

// ClientOptionFunc can be used to set the client options for a Doppler client.
type ClientOptionFunc func(*Client) error

// WithBaseURL can be used to set the base URL to make calls to the Doppler API.
func WithBaseURL(uri string) ClientOptionFunc {
	return func(c *Client) error {
		return c.setBaseURL(uri)
	}
}
