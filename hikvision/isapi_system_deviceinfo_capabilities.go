package hikvision

import (
	"encoding/xml"
	"net/url"
)

// DeviceInfoCap represents the XML_Cap_DeviceInfo and JSON_Cap_DeviceInfo resource.
type DeviceInfoCap struct {
	XMLName              xml.Name `xml:"DeviceInfo,omitempty"`
	XMLVersion           string   `xml:"version,attr"`
	XMLNamespace         string   `xml:"xmlns,attr"`
	DeviceName           string   `xml:"deviceName,omitempty" json:"deviceName,omitempty"`
	LanguageType         string   `xml:"languageType,omitempty" json:"languageType,omitempty"`
	DeviceID             string   `xml:"deviceID,omitempty" json:"deviceID,omitempty"`
	DeviceDescription    string   `xml:"deviceDescription,omitempty" json:"deviceDescription,omitempty"`
	DeviceLocation       string   `xml:"deviceLocation,omitempty" json:"deviceLocation,omitempty"`
	SystemContact        string   `xml:"systemContact,omitempty" json:"systemContact,omitempty"`
	Model                string   `xml:"model,omitempty" json:"model,omitempty"`
	SerialNumber         string   `xml:"serialNumber,omitempty" json:"serialNumber,omitempty"`
	MacAddress           string   `xml:"macAddress,omitempty" json:"macAddress,omitempty"`
	FirmwareVersion      string   `xml:"firmwareVersion,omitempty" json:"firmwareVersion,omitempty"`
	FirmwareReleasedDate string   `xml:"firmwareReleasedDate,omitempty" json:"firmwareReleasedDate,omitempty"`
	BootVersion          string   `xml:"bootVersion,omitempty" json:"bootVersion,omitempty"`
	BootReleasedDate     string   `xml:"bootReleasedDate,omitempty" json:"bootReleasedDate,omitempty"`
	HardwareVersion      string   `xml:"hardwareVersion,omitempty" json:"hardwareVersion,omitempty"`
	EncoderVersion       string   `xml:"encoderVersion,omitempty" json:"encoderVersion,omitempty"`
	EncoderReleasedDate  string   `xml:"encoderReleasedDate,omitempty" json:"encoderReleasedDate,omitempty"`
	DecoderVersion       string   `xml:"decoderVersion,omitempty" json:"decoderVersion,omitempty"`
	DecoderReleasedDate  string   `xml:"decoderReleasedDate,omitempty" json:"decoderReleasedDate,omitempty"`
	DeviceType           string   `xml:"deviceType,omitempty" json:"deviceType,omitempty"`
	TelecontrolID        int      `xml:"telecontrolID,omitempty" json:"telecontrolID,omitempty"`
	SupportBeep          bool     `xml:"supportBeep,omitempty" json:"supportBeep,omitempty"`
	FirmwareVersionInfo  string   `xml:"firmwareVersionInfo,omitempty" json:"firmwareVersionInfo,omitempty"`
	SubChannelEnabled    bool     `xml:"subChannelEnabled,omitempty" json:"subChannelEnabled,omitempty"`
	ThrChannelEnabled    bool     `xml:"thrChannelEnabled,omitempty" json:"thrChannelEnabled,omitempty"`
	ActualFloorNum       int      `xml:"actualFloorNum,omitempty" json:"actualFloorNum,omitempty"`
	RadarVersion         string   `xml:"radarVersion,omitempty" json:"radarVersion,omitempty"`
	PowerOnMode          string   `xml:"powerOnMode,omitempty" json:"powerOnMode,omitempty"`
}

// GetDeviceInfoCapabilities executes a HTTP GET request to the
// /ISAPI/System/deviceInfo/capabilities endpoint.
func (c *Client) GetDeviceInfoCapabilities() (resp *DeviceInfoCap, err error) {
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
