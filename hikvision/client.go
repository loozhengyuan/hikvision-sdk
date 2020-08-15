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

var (
	// ErrResponseNotOk is returned by Client.Get calls when the
	// the API returns a response with a non-200 HTTP status code.
	ErrResponseNotOk = errors.New("hikvision: response not ok")

	// ErrParseErrorMessageFailure is returned by Client.Get calls
	// when the API call is not successful but there the error
	// message could not be successfully parsed.
	ErrParseErrorMessageFailure = errors.New("hikvision: error parsing error message")

	// ErrUnhandledContentType is returned when the Content-Type
	// header is not unhandled by the client.
	ErrUnhandledContentType = errors.New("hikvision: unhandled content type")
)

// ResponseStatus represents the XML_ResponseStatus and JSON_ResponseStatus resource.
type ResponseStatus struct {
	XMLName       xml.Name                     `xml:"ResponseStatus,omitempty"`
	XMLVersion    string                       `xml:"version,attr"`
	XMLNamespace  string                       `xml:"xmlns,attr"`
	ID            int                          `xml:"id,omitempty" json:"id,omitempty"`
	RequestURL    string                       `xml:"requestURL,omitempty" json:"requestURL,omitempty"`
	StatusCode    int                          `xml:"statusCode,omitempty" json:"statusCode,omitempty"`
	StatusString  string                       `xml:"statusString,omitempty" json:"statusString,omitempty"`
	SubStatusCode string                       `xml:"subStatusCode,omitempty" json:"subStatusCode,omitempty"`
	ErrorCode     int                          `xml:"errorCode,omitempty" json:"errorCode,omitempty"`
	ErrorMsg      string                       `xml:"errorMsg,omitempty" json:"errorMsg,omitempty"`
	AdditionalErr *ResponseStatusAdditionalErr `xml:"AdditionalErr,omitempty" json:"AdditionalErr,omitempty"`
}

// ResponseStatusAdditionalErr represents the additional error status, which is
// valid when StatusCode is set to 9.
type ResponseStatusAdditionalErr struct {
	StatusList []ResponseStatusAdditionalErrStatus `xml:"StatusList,omitempty" json:"StatusList,omitempty"`
}

// ResponseStatusAdditionalErrStatus represents a single status information.
type ResponseStatusAdditionalErrStatus struct {
	Status string `xml:"Status,omitempty" json:"Status,omitempty"`
}

// ResponseStatusAdditionalErrStatusInfo represents information of status.
type ResponseStatusAdditionalErrStatusInfo struct {
	ID            string `xml:"id,omitempty" json:"id,omitempty"`
	StatusCode    int    `xml:"statusCode,omitempty" json:"statusCode,omitempty"`
	StatusString  string `xml:"statusString,omitempty" json:"statusString,omitempty"`
	SubStatusCode string `xml:"subStatusCode,omitempty" json:"subStatusCode,omitempty"`
	ErrorCode     int    `xml:"errorCode,omitempty" json:"errorCode,omitempty"`
	ErrorMsg      string `xml:"errorMsg,omitempty" json:"errorMsg,omitempty"`
}

func (r *ResponseStatus) String() string {
	return fmt.Sprintf("Status %v: %s", r.StatusCode, r.StatusString)
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

	// Handle non-success HTTP responses
	if resp.StatusCode != http.StatusOK {
		e := ResponseStatus{}
		switch ct := resp.Header.Get("Content-Type"); ct {
		case `application/xml; charset="UTF-8"`:
			if err := xml.Unmarshal(body, &e); err != nil {
				return nil, fmt.Errorf("%w: %v", ErrParseErrorMessageFailure, string(body))
			}
		case `application/json; charset="UTF-8"`:
			if err := json.Unmarshal(body, &e); err != nil {
				return nil, fmt.Errorf("%w: %v", ErrParseErrorMessageFailure, string(body))
			}
		default:
			return nil, fmt.Errorf("%w: %v", ErrUnhandledContentType, ct)
		}
		return nil, fmt.Errorf("%w: %v", ErrResponseNotOk, e.String())
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
