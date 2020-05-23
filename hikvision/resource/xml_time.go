package resource

import (
	"encoding/xml"
)

// XMLTime represents to XML_Time resource.
type XMLTime struct {
	XMLName           xml.Name `xml:"Time,omitempty"`
	XMLVersion        string   `xml:"version,attr"`
	XMLNamespace      string   `xml:"xmlns,attr"`
	TimeMode          string   `xml:"timeMode,omitempty"`
	LocalTime         string   `xml:"localTime,omitempty"`
	TimeZone          string   `xml:"timeZone,omitempty"`
	SatelliteInterval string   `xml:"satelliteInterval,omitempty"`
}

// Kind method outputs the resource kind.
func (r *XMLTime) Kind() string {
	return "xml"
}

// String method outputs the resource in string format.
func (r *XMLTime) String() (string, error) {
	s, err := xml.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// StringIndent method outputs the resource in indented string format.
func (r *XMLTime) StringIndent() (string, error) {
	s, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// NewXMLTime is a constructor for the XMLTime resource.
func NewXMLTime() *XMLTime {
	return &XMLTime{
		XMLVersion:   "2.0",
		XMLNamespace: "http://www.isapi.org/ver20/XMLSchema",
	}
}
