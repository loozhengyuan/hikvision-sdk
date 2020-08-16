package hikvision

import (
	"encoding/xml"
	"net/url"
)

// DeviceInfo represents the XML_DeviceInfo and JSON_DeviceInfo resource.
type DeviceInfo struct {
	XMLName              xml.Name                        `xml:"DeviceInfo,omitempty"`
	XMLVersion           string                          `xml:"version,attr"`
	XMLNamespace         string                          `xml:"xmlns,attr"`
	DeviceName           string                          `xml:"deviceName,omitempty" json:"deviceName,omitempty"`
	DeviceID             string                          `xml:"deviceID,omitempty" json:"deviceID,omitempty"`
	DeviceDescription    string                          `xml:"deviceDescription,omitempty" json:"deviceDescription,omitempty"`
	DeviceLocation       string                          `xml:"deviceLocation,omitempty" json:"deviceLocation,omitempty"`
	DeviceStatus         string                          `xml:"deviceStatus,omitempty" json:"deviceStatus,omitempty"`
	DetailAbnormalStatus *DeviceInfoDetailAbnormalStatus `xml:"DetailAbnormalStatus,omitempty" json:"DetailAbnormalStatus,omitempty"`
	SystemContact        string                          `xml:"systemContact,omitempty" json:"systemContact,omitempty"`
	Model                string                          `xml:"model,omitempty" json:"model,omitempty"`
	SerialNumber         string                          `xml:"serialNumber,omitempty" json:"serialNumber,omitempty"`
	MacAddress           string                          `xml:"macAddress,omitempty" json:"macAddress,omitempty"`
	FirmwareVersion      string                          `xml:"firmwareVersion,omitempty" json:"firmwareVersion,omitempty"`
	FirmwareReleasedDate string                          `xml:"firmwareReleasedDate,omitempty" json:"firmwareReleasedDate,omitempty"`
	BootVersion          string                          `xml:"bootVersion,omitempty" json:"bootVersion,omitempty"`
	BootReleasedDate     string                          `xml:"bootReleasedDate,omitempty" json:"bootReleasedDate,omitempty"`
	HardwareVersion      string                          `xml:"hardwareVersion,omitempty" json:"hardwareVersion,omitempty"`
	EncoderVersion       string                          `xml:"encoderVersion,omitempty" json:"encoderVersion,omitempty"`
	EncoderReleasedDate  string                          `xml:"encoderReleasedDate,omitempty" json:"encoderReleasedDate,omitempty"`
	DecoderVersion       string                          `xml:"decoderVersion,omitempty" json:"decoderVersion,omitempty"`
	DecoderReleasedDate  string                          `xml:"decoderReleasedDate,omitempty" json:"decoderReleasedDate,omitempty"`
	SoftwareVersion      string                          `xml:"softwareVersion,omitempty" json:"softwareVersion,omitempty"`
	Capacity             int                             `xml:"capacity,omitempty" json:"capacity,omitempty"`
	UsedCapacity         int                             `xml:"usedCapacity,omitempty" json:"usedCapacity,omitempty"`
	DeviceType           string                          `xml:"deviceType,omitempty" json:"deviceType,omitempty"`
	TelecontrolID        int                             `xml:"telecontrolID,omitempty" json:"telecontrolID,omitempty"`
	SupportBeep          bool                            `xml:"supportBeep,omitempty" json:"supportBeep,omitempty"`
	ActualFloorNum       int                             `xml:"actualFloorNum,omitempty" json:"actualFloorNum,omitempty"`
	SubChannelEnabled    bool                            `xml:"subChannelEnabled,omitempty" json:"subChannelEnabled,omitempty"`
	ThrChannelEnabled    bool                            `xml:"thrChannelEnabled,omitempty" json:"thrChannelEnabled,omitempty"`
	RadarVersion         string                          `xml:"radarVersion,omitempty" json:"radarVersion,omitempty"`
	LocalZoneNum         int                             `xml:"localZoneNum,omitempty" json:"localZoneNum,omitempty"`
	AlarmOutNum          int                             `xml:"alarmOutNum,omitempty" json:"alarmOutNum,omitempty"`
	DistanceResolution   float32                         `xml:"distanceResolution,omitempty" json:"distanceResolution,omitempty"`
	AngleResolution      float32                         `xml:"angleResolution,omitempty" json:"angleResolution,omitempty"`
	SpeedResolution      float32                         `xml:"speedResolution,omitempty" json:"speedResolution,omitempty"`
	DetectDistance       float32                         `xml:"detectDistance,omitempty" json:"detectDistance,omitempty"`
	LanguageType         string                          `xml:"languageType,omitempty" json:"languageType,omitempty"`
	RelayNum             int                             `xml:"relayNum,omitempty" json:"relayNum,omitempty"`
	ElectroLockNum       int                             `xml:"electroLockNum,omitempty" json:"electroLockNum,omitempty"`
	RS485Num             int                             `xml:"RS485Num,omitempty" json:"RS485Num,omitempty"`
	PowerOnMode          string                          `xml:"powerOnMode,omitempty" json:"powerOnMode,omitempty"`
}

// DeviceInfoDetailAbnormalStatus represents the device abnormal status,
// which is used in the XML_DeviceInfo resource.
type DeviceInfoDetailAbnormalStatus struct {
	HardDiskFull         bool `xml:"hardDiskFull,omitempty" json:"hardDiskFull,omitempty"`
	HardDiskError        bool `xml:"hardDiskError,omitempty" json:"hardDiskError,omitempty"`
	EthernetBroken       bool `xml:"ethernetBroken,omitempty" json:"ethernetBroken,omitempty"`
	IPAddrConflict       bool `xml:"ipaddrConflict,omitempty" json:"ipaddrConflict,omitempty"`
	IllegalAccess        bool `xml:"illegalAccess,omitempty" json:"illegalAccess,omitempty"`
	RecordError          bool `xml:"recordError,omitempty" json:"recordError,omitempty"`
	RaidLogicDiskError   bool `xml:"raidLogicDiskError,omitempty" json:"raidLogicDiskError,omitempty"`
	SpareWorkDeviceError bool `xml:"spareWorkDeviceError,omitempty" json:"spareWorkDeviceError,omitempty"`
}

// GetDeviceInfo executes a HTTP GET request to the
// /ISAPI/System/deviceInfo endpoint.
func (c *Client) GetDeviceInfo() (resp *DeviceInfo, err error) {
	path := "/ISAPI/System/deviceInfo"
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
