package resource

import (
	"encoding/xml"
)

// XMLDeviceInfo represents to XML_DeviceInfo resource.
type XMLDeviceInfo struct {
	XMLName              xml.Name                           `xml:"DeviceInfo,omitempty"`
	XMLVersion           string                             `xml:"version,attr"`
	XMLNamespace         string                             `xml:"xmlns,attr"`
	DeviceName           string                             `xml:"deviceName,omitempty"`
	DeviceID             string                             `xml:"deviceID,omitempty"`
	DeviceDescription    string                             `xml:"deviceDescription,omitempty"`
	DeviceLocation       string                             `xml:"deviceLocation,omitempty"`
	DeviceStatus         string                             `xml:"deviceStatus,omitempty"`
	DetailAbnormalStatus *XMLDeviceInfoDetailAbnormalStatus `xml:"DetailAbnormalStatus,omitempty"`
	SystemContact        string                             `xml:"systemContact,omitempty"`
	Model                string                             `xml:"model,omitempty"`
	SerialNumber         string                             `xml:"serialNumber,omitempty"`
	MacAddress           string                             `xml:"macAddress,omitempty"`
	FirmwareVersion      string                             `xml:"firmwareVersion,omitempty"`
	FirmwareReleasedDate string                             `xml:"firmwareReleasedDate,omitempty"`
	BootVersion          string                             `xml:"bootVersion,omitempty"`
	BootReleasedDate     string                             `xml:"bootReleasedDate,omitempty"`
	HardwareVersion      string                             `xml:"hardwareVersion,omitempty"`
	EncoderVersion       string                             `xml:"encoderVersion,omitempty"`
	EncoderReleasedDate  string                             `xml:"encoderReleasedDate,omitempty"`
	DecoderVersion       string                             `xml:"decoderVersion,omitempty"`
	DecoderReleasedDate  string                             `xml:"decoderReleasedDate,omitempty"`
	SoftwareVersion      string                             `xml:"softwareVersion,omitempty"`
	Capacity             int                                `xml:"capacity,omitempty"`
	UsedCapacity         int                                `xml:"usedCapacity,omitempty"`
	DeviceType           string                             `xml:"deviceType,omitempty"`
	TelecontrolID        int                                `xml:"telecontrolID,omitempty"`
	SupportBeep          bool                               `xml:"supportBeep,omitempty"`
	ActualFloorNum       int                                `xml:"actualFloorNum,omitempty"`
	SubChannelEnabled    bool                               `xml:"subChannelEnabled,omitempty"`
	ThrChannelEnabled    bool                               `xml:"thrChannelEnabled,omitempty"`
	RadarVersion         string                             `xml:"radarVersion,omitempty"`
	LocalZoneNum         int                                `xml:"localZoneNum,omitempty"`
	AlarmOutNum          int                                `xml:"alarmOutNum,omitempty"`
	DistanceResolution   float32                            `xml:"distanceResolution,omitempty"`
	AngleResolution      float32                            `xml:"angleResolution,omitempty"`
	SpeedResolution      float32                            `xml:"speedResolution,omitempty"`
	DetectDistance       float32                            `xml:"detectDistance,omitempty"`
	LanguageType         string                             `xml:"languageType,omitempty"`
	RelayNum             int                                `xml:"relayNum,omitempty"`
	ElectroLockNum       int                                `xml:"electroLockNum,omitempty"`
	RS485Num             int                                `xml:"RS485Num,omitempty"`
	PowerOnMode          string                             `xml:"powerOnMode,omitempty"`
}

// XMLDeviceInfoDetailAbnormalStatus represents the device abnormal status,
// which is used in the XML_DeviceInfo resource.
type XMLDeviceInfoDetailAbnormalStatus struct {
	HardDiskFull         bool `xml:"hardDiskFull,omitempty"`
	HardDiskError        bool `xml:"hardDiskError,omitempty"`
	EthernetBroken       bool `xml:"ethernetBroken,omitempty"`
	IPAddrConflict       bool `xml:"ipaddrConflict,omitempty"`
	IllegalAccess        bool `xml:"illegalAccess,omitempty"`
	RecordError          bool `xml:"recordError,omitempty"`
	RaidLogicDiskError   bool `xml:"raidLogicDiskError,omitempty"`
	SpareWorkDeviceError bool `xml:"spareWorkDeviceError,omitempty"`
}

// Kind method outputs the resource kind.
func (r *XMLDeviceInfo) Kind() string {
	return "xml"
}

// String method outputs the resource in string format.
func (r *XMLDeviceInfo) String() (string, error) {
	s, err := xml.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// StringIndent method outputs the resource in indented string format.
func (r *XMLDeviceInfo) StringIndent() (string, error) {
	s, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// NewXMLDeviceInfo is a constructor for the XMLDeviceInfo resource.
func NewXMLDeviceInfo() *XMLDeviceInfo {
	return &XMLDeviceInfo{
		XMLVersion:   "2.0",
		XMLNamespace: "http://www.isapi.org/ver20/XMLSchema",
	}
}
