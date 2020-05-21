package types

import (
	"encoding/xml"
)

// Response comment
type Response struct {
	XMLName       xml.Name       `xml:"ResponseStatus,omitempty"`
	RequestURL    string         `xml:"requestURL,omitempty"`
	StatusCode    int            `xml:"statusCode,omitempty"`
	StatusString  string         `xml:"statusString,omitempty"`
	ID            int            `xml:"id,omitempty"`
	SubStatusCode string         `xml:"subStatusCode,omitempty"`
	ErrorCode     int            `xml:"errorCode,omitempty"`
	ErrorMsg      string         `xml:"errorMsg,omitempty"`
	AdditionalErr *AdditionalErr `xml:"AdditionalErr,omitempty"`
}

// AdditionalErr comment
type AdditionalErr struct {
	StatusList []Status `xml:"StatusList,omitempty"`
}

// Status comment
type Status struct {
	ID            string `xml:"id,omitempty"`
	StatusCode    int    `xml:"statusCode,omitempty"`
	StatusString  string `xml:"statusString,omitempty"`
	SubStatusCode string `xml:"subStatusCode,omitempty"`
}
