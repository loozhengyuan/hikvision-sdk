package hikvision

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/bobziuchkovski/digest"
)

// Client comment
type Client struct {
	Client *http.Client
	URL    *url.URL
}

// NewClient is a constructor for the Client object.
func NewClient(host, scheme, username, password string) (*Client, error) {
	// Create client
	t := digest.NewTransport(username, password)
	c, err := t.Client()
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: c,
		URL: &url.URL{
			Scheme: scheme,
			Host:   host,
		},
	}, nil
}

// Do is executes a HTTP request.
func (c *Client) Do(method, path string) ([]byte, error) {
	// Get URL
	c.URL.Path = path
	url := c.URL.String()

	// Create request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// Send request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
