package resource

import (
	"encoding/xml"
)

// XMLCapDeviceInfo represents to XML_Cap_DeviceInfo resource.
type XMLCapDeviceInfo struct {
	XMLName              xml.Name `xml:"DeviceInfo,omitempty"`
	XMLVersion           string   `xml:"version,attr"`
	XMLNamespace         string   `xml:"xmlns,attr"`
	DeviceName           string   `xml:"deviceName,omitempty"`
	LanguageType         string   `xml:"languageType,omitempty"`
	DeviceID             string   `xml:"deviceID,omitempty"`
	DeviceDescription    string   `xml:"deviceDescription,omitempty"`
	DeviceLocation       string   `xml:"deviceLocation,omitempty"`
	SystemContact        string   `xml:"systemContact,omitempty"`
	Model                string   `xml:"model,omitempty"`
	SerialNumber         string   `xml:"serialNumber,omitempty"`
	MacAddress           string   `xml:"macAddress,omitempty"`
	FirmwareVersion      string   `xml:"firmwareVersion,omitempty"`
	FirmwareReleasedDate string   `xml:"firmwareReleasedDate,omitempty"`
	BootVersion          string   `xml:"bootVersion,omitempty"`
	BootReleasedDate     string   `xml:"bootReleasedDate,omitempty"`
	HardwareVersion      string   `xml:"hardwareVersion,omitempty"`
	EncoderVersion       string   `xml:"encoderVersion,omitempty"`
	EncoderReleasedDate  string   `xml:"encoderReleasedDate,omitempty"`
	DecoderVersion       string   `xml:"decoderVersion,omitempty"`
	DecoderReleasedDate  string   `xml:"decoderReleasedDate,omitempty"`
	DeviceType           string   `xml:"deviceType,omitempty"`
	TelecontrolID        int      `xml:"telecontrolID,omitempty"`
	SupportBeep          bool     `xml:"supportBeep,omitempty"`
	FirmwareVersionInfo  string   `xml:"firmwareVersionInfo,omitempty"`
	SubChannelEnabled    bool     `xml:"subChannelEnabled,omitempty"`
	ThrChannelEnabled    bool     `xml:"thrChannelEnabled,omitempty"`
	ActualFloorNum       int      `xml:"actualFloorNum,omitempty"`
	RadarVersion         string   `xml:"radarVersion,omitempty"`
	PowerOnMode          string   `xml:"powerOnMode,omitempty"`
}

// Kind method outputs the resource kind.
func (r *XMLCapDeviceInfo) Kind() string {
	return "xml"
}

// String method outputs the resource in string format.
func (r *XMLCapDeviceInfo) String() (string, error) {
	s, err := xml.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// StringIndent method outputs the resource in indented string format.
func (r *XMLCapDeviceInfo) StringIndent() (string, error) {
	s, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// NewXMLCapDeviceInfo is a constructor for the XMLCapDeviceInfo resource.
func NewXMLCapDeviceInfo() *XMLCapDeviceInfo {
	return &XMLCapDeviceInfo{
		XMLVersion:   "2.0",
		XMLNamespace: "http://www.isapi.org/ver20/XMLSchema",
	}
}
