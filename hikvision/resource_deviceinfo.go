package hikvision

import (
	"encoding/xml"
	"log"

	"github.com/loozhengyuan/hikvision-sdk/hikvision/resource"
)

// GetDeviceInfo executes a HTTP GET request to the
// /ISAPI/System/deviceInfo endpoint.
// TODO: Returns XML_Response when error
func (c *Client) GetDeviceInfo() (resource.Resource, error) {
	// Set variables
	method := "GET"
	path := "/ISAPI/System/deviceInfo"
	log.Println(method, path)

	// Execute request
	body, err := c.Get(path)
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
