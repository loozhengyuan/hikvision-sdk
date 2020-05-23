package hikvision

import (
	"encoding/xml"
	"log"

	"github.com/loozhengyuan/hikvision-sdk/hikvision/resource"
)

// GetTime executes a HTTP GET request to the
// /ISAPI/System/time endpoint.
// TODO: Returns XML_ResponseStatus when error
func (c *Client) GetTime() (resource.Resource, error) {
	// Set variables
	method := "GET"
	path := "/ISAPI/System/time"
	log.Println(method, path)

	// Execute request
	body, err := c.Get(path)
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
	// Set variables
	method := "PUT"
	path := "/ISAPI/System/time"
	log.Println(method, path)

	// Execute request
	var d resource.Resource = data
	body, err := c.Put(path, &d)
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
