package hikvision

import (
	"encoding/xml"
	"net/url"
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

// GetDeviceInfoCapabilities executes a HTTP GET request to the
// /ISAPI/System/deviceInfo/capabilities endpoint.
func (c *Client) GetDeviceInfoCapabilities() (resp *XMLCapDeviceInfo, err error) {
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
