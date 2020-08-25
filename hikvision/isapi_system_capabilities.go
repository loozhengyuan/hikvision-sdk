package hikvision

import (
	"encoding/xml"
	"net/url"
)

// DeviceCap represents the XML_DeviceCap and JSON_DeviceCap resource.
type DeviceCap struct {
	XMLName           xml.Name         `xml:"DeviceCap,omitempty"`
	XMLVersion        string           `xml:"version,attr"`
	XMLNamespace      string           `xml:"xmlns,attr"`
	SysCap            *DeviceCapSysCap `xml:"SysCap,omitempty"`
	VoiceTalkNums     int              `xml:"voicetalkNums,omitempty" json:"voicetalkNums,omitempty"`
	IsSupportSnapshot bool             `xml:"isSupportSnapshot,omitempty" json:"isSupportSnapshot,omitempty"`
	// TODO: SecurityCap
	// TODO: EventCap
	// TODO: ITCCap
	// TODO: ImageCap
	// TODO: RacmCap
	// TODO: PTZCtrlCap <--
	// TODO: SmartCap
	IsSupportEhome            bool `xml:"isSupportEhome,omitempty" json:"isSupportEhome,omitempty"`
	IsSupportStreamingEncrypt bool `xml:"isSupportStreamingEncrypt,omitempty" json:"isSupportStreamingEncrypt,omitempty"`
	// TODO: TestCap <--
	// TODO: ThermalCap
	// TODO: WLAlarmCap
	// TODO: SecurityCPCapabilities
	IsSupportGIS                    bool `xml:"isSupportGIS,omitempty" json:"isSupportGIS,omitempty"`
	IsSupportCompass                bool `xml:"isSupportCompass,omitempty" json:"isSupportCompass,omitempty"`
	IsSupportRoadInfoOverlays       bool `xml:"isSupportRoadInfoOverlays,omitempty" json:"isSupportRoadInfoOverlays,omitempty"`
	IsSupportFaceCaptureStatistics  bool `xml:"isSupportFaceCaptureStatistics,omitempty" json:"isSupportFaceCaptureStatistics,omitempty"`
	IsSupportExternalDevice         bool `xml:"isSupportExternalDevice,omitempty" json:"isSupportExternalDevice,omitempty"`
	IsSupportElectronicsEnlarge     bool `xml:"isSupportElectronicsEnlarge,omitempty" json:"isSupportElectronicsEnlarge,omitempty"`
	IsSupportRemoveStorage          bool `xml:"isSupportRemoveStorage,omitempty" json:"isSupportRemoveStorage,omitempty"`
	IsSupportCloud                  bool `xml:"isSupportCloud,omitempty" json:"isSupportCloud,omitempty"`
	IsSupportRecordHost             bool `xml:"isSupportRecordHost,omitempty" json:"isSupportRecordHost,omitempty"`
	IsSupportEagleEye               bool `xml:"isSupportEagleEye,omitempty" json:"isSupportEagleEye,omitempty"`
	IsSupportPanorama               bool `xml:"isSupportPanorama,omitempty" json:"isSupportPanorama,omitempty"`
	IsSupportFirmwareVersionInfo    bool `xml:"isSupportFirmwareVersionInfo,omitempty" json:"isSupportFirmwareVersionInfo,omitempty"`
	IsSupportExternalWirelessServer bool `xml:"isSupportExternalWirelessServer,omitempty" json:"isSupportExternalWirelessServer,omitempty"`
	IsSupportSetupCalibration       bool `xml:"isSupportSetupCalibration,omitempty" json:"isSupportSetupCalibration,omitempty"`
	IsSupportGetMutexFuncErrMsg     bool `xml:"isSupportGetmutexFuncErrMsg,omitempty" json:"isSupportGetmutexFuncErrMsg,omitempty"`
	IsSupportTokenAuthenticate      bool `xml:"isSupportTokenAuthenticate,omitempty" json:"isSupportTokenAuthenticate,omitempty"`
	IsSupportStreamDualVCA          bool `xml:"isSupportStreamDualVCA,omitempty" json:"isSupportStreamDualVCA,omitempty"`
	IsSupportLaserSpotManual        bool `xml:"isSupportlaserSpotManual,omitempty" json:"isSupportlaserSpotManual,omitempty"`
	IsSupportRTMP                   bool `xml:"isSupportRTMP,omitempty" json:"isSupportRTMP,omitempty"`
	IsSupportTraffic                bool `xml:"isSupportTraffic,omitempty" json:"isSupportTraffic,omitempty"`
}

// DeviceCapSysCap describes the capabilities of the system.
type DeviceCapSysCap struct {
	IsSupportDST bool `xml:"isSupportDst,omitempty" json:"isSupportDst,omitempty"`
	// TODO: NetworkCap
	// TODO: IOCap
	// TODO: SerialCap
	// TODO: VideoCap
	// TODO: AudioCap
	// NOTE: Not sure if spelling error is deliberate or not.
	IsSupportHolidy bool `xml:"isSupportHolidy,omitempty" json:"isSupportHolidy,omitempty"`
	// TODO: RebootConfigurationCap
	IsSupportExternalDevice         bool `xml:"isSupportExternalDevice,omitempty" json:"isSupportExternalDevice,omitempty"`
	IsSupportChangedUpload          bool `xml:"isSupportChangedUpload,omitempty" json:"isSupportChangedUpload,omitempty"`
	IsSupportWorkingStatus          bool `xml:"isSupportWorkingStatus,omitempty" json:"isSupportWorkingStatus,omitempty"`
	IsSupportChannelInfoByCondition bool `xml:"isSupportChannelInfoByCondition,omitempty" json:"isSupportChannelInfoByCondition,omitempty"`
	IsSupportDiagnosedDataParameter bool `xml:"isSupportDiagnosedDataParameter,omitempty" json:"isSupportDiagnosedDataParameter,omitempty"`
	IsSupportSimpleDevStatus        bool `xml:"isSupportSimpleDevStatus,omitempty" json:"isSupportSimpleDevStatus,omitempty"`
	IsSupportFlexible               bool `xml:"isSupportFlexible,omitempty" json:"isSupportFlexible,omitempty"`
	IsSupportPTZChannels            bool `xml:"isSupportPTZChannels,omitempty" json:"isSupportPTZChannels,omitempty"`
	IsSupportSubscribeEvent         bool `xml:"isSupportSubscribeEvent,omitempty" json:"isSupportSubscribeEvent,omitempty"`
	IsSupportDiagnosedData          bool `xml:"isSupportDiagnosedData,omitempty" json:"isSupportDiagnosedData,omitempty"`
	IsSupportTimeCap                bool `xml:"isSupportTimeCap,omitempty" json:"isSupportTimeCap,omitempty"`
	IsSupportThermalStreamData      bool `xml:"isSupportThermalStreamData,omitempty" json:"isSupportThermalStreamData,omitempty"`
	IsSupportPostUpdateFirmware     bool `xml:"isSupportPostUpdateFirmware,omitempty" json:"isSupportPostUpdateFirmware,omitempty"`
	IsSupportPostConfigData         bool `xml:"isSupportPostConfigData,omitempty" json:"isSupportPostConfigData,omitempty"`
	IsSupportUserLock               bool `xml:"isSupportUserLock,omitempty" json:"isSupportUserLock,omitempty"`
	IsSupportModuleLock             bool `xml:"isSupportModuleLock,omitempty" json:"isSupportModuleLock,omitempty"`
	IsSupportSoundCfg               bool `xml:"isSupportSoundCfg,omitempty" json:"isSupportSoundCfg,omitempty"`
	IsSupportMetadata               bool `xml:"isSupportMetadata,omitempty" json:"isSupportMetadata,omitempty"`
	IsSupportShutdown               bool `xml:"isSupportShutdown,omitempty" json:"isSupportShutdown,omitempty"`
	// TODO: supportSmartOverlapChannles
}

// GetSystemCapabilities executes a HTTP GET request to the
// /ISAPI/System/capabilties endpoint.
func (c *Client) GetSystemCapabilities() (resp *DeviceCap, err error) {
	path := "/ISAPI/System/capabilties"
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
