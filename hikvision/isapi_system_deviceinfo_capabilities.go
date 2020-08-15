package hikvision

import (
	"encoding/xml"
	"net/url"

	"github.com/loozhengyuan/hikvision-sdk/hikvision/resource"
)

// GetDeviceInfoCapabilities executes a HTTP GET request to the
// /ISAPI/System/deviceInfo/capabilities endpoint.
// TODO: Returns XML_Response when error
func (c *Client) GetDeviceInfoCapabilities() (resource.Resource, error) {
	// Parse URL
	path := "/ISAPI/System/deviceInfo/capabilities"
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
	var output resource.Resource = resource.NewXMLCapDeviceInfo()
	err = xml.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
