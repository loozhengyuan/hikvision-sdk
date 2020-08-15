package hikvision

import (
	"encoding/xml"
	"net/url"

	"github.com/loozhengyuan/hikvision-sdk/hikvision/resource"
)

// GetDeviceInfoCapabilities executes a HTTP GET request to the
// /ISAPI/System/deviceInfo/capabilities endpoint.
func (c *Client) GetDeviceInfoCapabilities() (resp *resource.XMLCapDeviceInfo, err error) {
	path := "/ISAPI/System/deviceInfo/capabilities"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	body, err := c.Get(u)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
