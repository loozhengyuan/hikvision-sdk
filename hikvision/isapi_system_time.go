package hikvision

import (
	"encoding/xml"
	"net/url"
)

// Time represents the XML_Time and JSON_Time resource.
type Time struct {
	XMLName           xml.Name `xml:"Time,omitempty"`
	XMLVersion        string   `xml:"version,attr"`
	XMLNamespace      string   `xml:"xmlns,attr"`
	TimeMode          string   `xml:"timeMode,omitempty" json:"timeMode,omitempty"`
	LocalTime         string   `xml:"localTime,omitempty" json:"localTime,omitempty"`
	TimeZone          string   `xml:"timeZone,omitempty" json:"timeZone,omitempty"`
	SatelliteInterval string   `xml:"satelliteInterval,omitempty" json:"satelliteInterval,omitempty"`
}

// GetTime executes a HTTP GET request to the
// /ISAPI/System/time endpoint.
func (c *Client) GetTime() (resp *Time, err error) {
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
func (c *Client) PutTime(data *Time) (resp *ResponseStatus, err error) {
	path := "/ISAPI/System/time"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	body, err := c.PutXML(u, data)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
