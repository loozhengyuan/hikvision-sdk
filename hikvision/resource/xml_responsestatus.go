package resource

import (
	"encoding/xml"
)

// XMLResponseStatus represents to XML_ResponseStatus resource.
type XMLResponseStatus struct {
	XMLName       xml.Name                        `xml:"ResponseStatus,omitempty"`
	XMLVersion    string                          `xml:"version,attr"`
	XMLNamespace  string                          `xml:"xmlns,attr"`
	RequestURL    string                          `xml:"requestURL,omitempty"`
	StatusCode    int                             `xml:"statusCode,omitempty"`
	StatusString  string                          `xml:"statusString,omitempty"`
	ID            int                             `xml:"id,omitempty"`
	SubStatusCode string                          `xml:"subStatusCode,omitempty"`
	ErrorCode     int                             `xml:"errorCode,omitempty"`
	ErrorMsg      string                          `xml:"errorMsg,omitempty"`
	AdditionalErr *XMLResponseStatusAdditionalErr `xml:"AdditionalErr,omitempty"`
}

// XMLResponseStatusAdditionalErr comment
type XMLResponseStatusAdditionalErr struct {
	StatusList []XMLResponseStatusAdditionalErrStatus `xml:"StatusList,omitempty"`
}

// XMLResponseStatusAdditionalErrStatus comment
type XMLResponseStatusAdditionalErrStatus struct {
	ID            string `xml:"id,omitempty"`
	StatusCode    int    `xml:"statusCode,omitempty"`
	StatusString  string `xml:"statusString,omitempty"`
	SubStatusCode string `xml:"subStatusCode,omitempty"`
}

// Kind method outputs the resource kind.
func (r *XMLResponseStatus) Kind() string {
	return "xml"
}

// String method outputs the resource in string format.
func (r *XMLResponseStatus) String() (string, error) {
	s, err := xml.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// StringIndent method outputs the resource in indented string format.
func (r *XMLResponseStatus) StringIndent() (string, error) {
	s, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// NewXMLResponseStatus is a constructor for the Time resource.
func NewXMLResponseStatus() *XMLResponseStatus {
	return &XMLResponseStatus{
		XMLVersion:   "2.0",
		XMLNamespace: "http://www.std-cgi.org/ver20/XMLSchema",
	}
}
