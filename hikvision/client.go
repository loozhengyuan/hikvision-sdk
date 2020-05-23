package hikvision

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/loozhengyuan/hikvision-sdk/hikvision/resource"
)

// Client is a http.Client wrapper that handles authentication.
type Client struct {
	Client *http.Client
	URL    *url.URL
}

// NewClient is a constructor for the Client object.
func NewClient(host, username, password string) (*Client, error) {
	return &Client{
		Client: &http.Client{
			Transport: NewAuthTransport(username, password),
		},
		URL: &url.URL{
			Scheme: "http",
			Host:   host,
		},
	}, nil
}

// Do executes a HTTP request.
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

// DoWithBody executes a HTTP request containing a request body.
func (c *Client) DoWithBody(method, path string, data resource.Resource) ([]byte, error) {
	// Get URL
	c.URL.Path = path
	url := c.URL.String()

	// Handle data
	var kind string
	if data != nil {
		kind = data.Kind()
	}

	// Handle data kind
	b := new(bytes.Buffer)
	headers := map[string]string{}
	switch kind {
	// case "":
	// 	b = nil
	case "xml":
		xml.NewEncoder(b).Encode(data)
		headers["Content-Type"] = `application/xml; charset="UTF-8"`
	case "json":
		json.NewEncoder(b).Encode(data)
		headers["Content-Type"] = `application/json`
		// default:
		// 	return nil, ErrInvalidResourceKind
	}

	// Create request
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, err
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
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

// Get executes a HTTP GET request.
func (c *Client) Get(path string) ([]byte, error) {
	return c.Do("GET", path)
}

// Put executes a HTTP PUT request.
func (c *Client) Put(path string, data *resource.Resource) ([]byte, error) {
	return c.DoWithBody("PUT", path, *data)
}
