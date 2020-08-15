package hikvision

import (
	"encoding/xml"
	"net/url"

	"github.com/loozhengyuan/hikvision-sdk/hikvision/resource"
)

// GetTime executes a HTTP GET request to the
// /ISAPI/System/time endpoint.
// TODO: Returns XML_ResponseStatus when error
func (c *Client) GetTime() (resource.Resource, error) {
	// Parse URL
	path := "/ISAPI/System/time"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}

	// Execute request
	body, err := c.Get(u)
	if err != nil {
		return nil, err
	}

	// Unmarshall data
	var output resource.Resource = resource.NewXMLTime()
	err = xml.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// PutTime executes a HTTP PUT request to the
// /ISAPI/System/time endpoint.
func (c *Client) PutTime(data *resource.XMLTime) (resource.Resource, error) {
	// Parse URL
	path := "/ISAPI/System/time"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}

	// Execute request
	var d resource.Resource = data
	body, err := c.Put(u, &d)
	if err != nil {
		return nil, err
	}

	// Unmarshall data
	var output resource.Resource = resource.NewXMLResponseStatus()
	err = xml.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
