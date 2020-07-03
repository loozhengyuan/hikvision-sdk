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
	Client  *http.Client
	BaseURL string
}

// NewClient is a constructor for the Client object.
func NewClient(host, username, password string) (*Client, error) {
	u, err := url.Parse("http://" + host)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: &http.Client{
			Transport: NewAuthTransport(username, password),
		},
		BaseURL: u.String(),
	}, nil
}

// Do executes a HTTP request.
func (c *Client) Do(method string, u *url.URL) ([]byte, error) {
	// Create request
	req, err := http.NewRequest(method, u.String(), nil)
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
func (c *Client) DoWithBody(method string, u *url.URL, data resource.Resource) ([]byte, error) {
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
	req, err := http.NewRequest(method, u.String(), b)
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
func (c *Client) Get(u *url.URL) ([]byte, error) {
	return c.Do("GET", u)
}

// Put executes a HTTP PUT request.
func (c *Client) Put(u *url.URL, data *resource.Resource) ([]byte, error) {
	return c.DoWithBody("PUT", u, *data)
}
