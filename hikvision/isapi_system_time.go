package hikvision

import (
	"encoding/xml"
	"net/url"

	"github.com/loozhengyuan/hikvision-sdk/hikvision/resource"
)

// GetTime executes a HTTP GET request to the
// /ISAPI/System/time endpoint.
func (c *Client) GetTime() (resp *resource.XMLTime, err error) {
	path := "/ISAPI/System/time"
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

// PutTime executes a HTTP PUT request to the
// /ISAPI/System/time endpoint.
func (c *Client) PutTime(data *resource.XMLTime) (resp *resource.XMLResponseStatus, err error) {
	path := "/ISAPI/System/time"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	var d resource.Resource = data
	body, err := c.Put(u, &d)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
