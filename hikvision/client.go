package hikvision

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/loozhengyuan/hikvision-sdk/hikvision/resource"
)

const (
	contentTypeXML  = `application/xml; charset="UTF-8"`
	contentTypeJSON = `application/json; charset="UTF-8"`
)

var (
	// ErrResponseNotOk is returned by Client.Get calls when the
	// the API returns a response with a non-200 HTTP status code.
	ErrResponseNotOk = errors.New("hikvision: response not ok")

	// ErrParseErrorMessageFailure is returned by Client.Get calls
	// when the API call is not successful but there the error
	// message could not be successfully parsed.
	ErrParseErrorMessageFailure = errors.New("hikvision: error parsing error message")
)

// XMLResponseStatus represents to XML_ResponseStatus resource.
type XMLResponseStatus struct {
	XMLName       xml.Name                        `xml:"ResponseStatus,omitempty"`
	XMLVersion    string                          `xml:"version,attr"`
	XMLNamespace  string                          `xml:"xmlns,attr"`
	RequestURL    string                          `xml:"requestURL,omitempty"`
	StatusCode    int                             `xml:"statusCode,omitempty"`
	StatusString  string                          `xml:"statusString,omitempty"`
	ID            int                             `xml:"id,omitempty"`
	SubStatusCode string                          `xml:"subStatusCode,omitempty"`
	ErrorCode     int                             `xml:"errorCode,omitempty"`
	ErrorMsg      string                          `xml:"errorMsg,omitempty"`
	AdditionalErr *XMLResponseStatusAdditionalErr `xml:"AdditionalErr,omitempty"`
}

// XMLResponseStatusAdditionalErr comment
type XMLResponseStatusAdditionalErr struct {
	StatusList []XMLResponseStatusAdditionalErrStatus `xml:"StatusList,omitempty"`
}

// XMLResponseStatusAdditionalErrStatus comment
type XMLResponseStatusAdditionalErrStatus struct {
	ID            string `xml:"id,omitempty"`
	StatusCode    int    `xml:"statusCode,omitempty"`
	StatusString  string `xml:"statusString,omitempty"`
	SubStatusCode string `xml:"subStatusCode,omitempty"`
}

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
func (c *Client) Do(r *http.Request) ([]byte, error) {
	// Send request
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Handle non-success HTTP responses
	// TODO: Check and handle JSON responses
	if resp.StatusCode != http.StatusOK {
		var e resource.XMLResponseStatus
		if err := xml.Unmarshal(body, &e); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrParseErrorMessageFailure, string(body))
		}
		return nil, fmt.Errorf("%w: %v", ErrResponseNotOk, e)
	}

	return body, nil
}

// Get executes a HTTP GET request.
func (c *Client) Get(u *url.URL) ([]byte, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// Put executes a HTTP PUT request.
func (c *Client) Put(u *url.URL, contentType string, data []byte) ([]byte, error) {
	b := bytes.NewBuffer(data)
	req, err := http.NewRequest("GET", u.String(), b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}

// PutXML executes a HTTP PUT request with `application/xml` content type.
func (c *Client) PutXML(u *url.URL, data interface{}) ([]byte, error) {
	b, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return c.Put(u, contentTypeXML, b)
}

// PutJSON executes a HTTP PUT request with `application/json` content type.
func (c *Client) PutJSON(u *url.URL, contentType string, data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return c.Put(u, contentTypeJSON, b)
}
