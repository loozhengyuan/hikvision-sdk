package hikvision

import (
	"encoding/xml"
	"net/url"

	"github.com/loozhengyuan/hikvision-sdk/hikvision/resource"
)

// GetDeviceInfo executes a HTTP GET request to the
// /ISAPI/System/deviceInfo endpoint.
// TODO: Returns XML_Response when error
func (c *Client) GetDeviceInfo() (resource.Resource, error) {
	// Parse URL
	path := "/ISAPI/System/deviceInfo"
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
	var output resource.Resource = resource.NewXMLDeviceInfo()
	err = xml.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
