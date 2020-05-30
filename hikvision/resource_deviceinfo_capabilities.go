package hikvision

import (
	"encoding/xml"
	"log"

	"github.com/loozhengyuan/hikvision-sdk/hikvision/resource"
)

// GetDeviceInfoCapabilities executes a HTTP GET request to the
// /ISAPI/System/deviceInfo/capabilities endpoint.
// TODO: Returns XML_Response when error
func (c *Client) GetDeviceInfoCapabilities() (resource.Resource, error) {
	// Set variables
	method := "GET"
	path := "/ISAPI/System/deviceInfo/capabilities"
	log.Println(method, path)

	// Execute request
	body, err := c.Get(path)
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
