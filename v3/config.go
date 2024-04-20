package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"jlink-restful-golang-demo/config"
	"jlink-restful-golang-demo/utils"
	"log"
	"strings"
)

// 获取设备配置
func (jDevice *JLinkDevice) DeviceGetconfig(name string) (data interface{}, err error) {
	res, err := jDevice.GetConfigHttp(name)
	if err != nil {
		return nil, err
	}
	if name == utils.ConfigAVEncEncode {
		var resp *config.AVEncEncodeResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AVEncEncode, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAVEncEncode {
		var resp *config.AVEncEncodeResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AVEncEncode, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAVEncVideoWidget {
		var resp *config.AVEncVideoWidgetResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AVEncVideoWidget, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAVEncVideoColor {
		var resp *config.AVEncVideoColorResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AVEncVideoColor, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAVEncSmartH264 {
		var resp *config.AVEncSmartH264Resp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AVEncSmartH264, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigfVideoVolume {
		var resp *config.FVideoVolumeResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.FVideoVolume, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigfVideoOSDInfo {
		var resp *config.FVideoOSDInfoResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.FVideoOSDInfo, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigRecord {
		var resp *config.RecordResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.Record, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigStorageSnapshot {
		var resp *config.StorageSnapshotResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.StorageSnapshot, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigDetectBlindDetect {
		var resp *config.DetectBlindDetectResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.DetectBlindDetect, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigDetectMotionDetect {
		var resp *config.DetectMotionDetectResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.DetectMotionDetect, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigDetectLossDetect {
		var resp *config.DetectLossDetectResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.DetectLossDetect, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigChannelTitle {
		var resp *config.ChannelTitleResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.ChannelTitle, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigCamera {
		var resp *config.CameraResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.Camera, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigDetectHumanDetection {
		var resp *config.DetectHumanDetectionResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.DetectHumanDetection, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigDetectHumanDetection0 {
		var resp *config.DetectHumanDetection0Resp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.DetectHumanDetection0, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAbilityVoiceTipType {
		var resp *config.AbilityVoiceTipTypeResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AbilityVoiceTipType, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigUartComm {
		var resp *config.UartCommResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.UartComm, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigUartPTZ {
		var resp *config.UartPTZResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.UartPTZ, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigUartRS485 {
		var resp *config.UartRS485Resp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.UartRS485, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetCommon {
		var resp *config.NetWorkNetCommonResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetCommon, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetDHCP {
		var resp *config.NetWorkNetDHCPResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetDHCP, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetEmail {
		var resp *config.NetWorkNetEmailResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetEmail, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetNTP {
		var resp *config.NetWorkNetNTPResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetNTP, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetDNS {
		var resp *config.NetWorkNetDNSResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetDNS, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetFTP {
		var resp *config.NetWorkNetFTPResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetFTP, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAlarmNetIPConflict {
		var resp *config.AlarmNetIPConflictResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AlarmNetIPConflict, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAlarmNetAbort {
		var resp *config.AlarmNetAbortResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AlarmNetAbort, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkUpnp {
		var resp *config.NetWorkUpnpResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkUpnp, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetIPFilter {
		var resp *config.NetWorkNetIPFilterResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetIPFilter, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigOPUTCTimeSetting {
		var resp *config.OPUTCTimeSettingResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.OPUTCTimeSetting, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetOrder {
		var resp *config.NetWorkNetOrderResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetOrder, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkDAS {
		var resp *config.NetWorkDASResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkDAS, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkPMS {
		var resp *config.NetWorkPMSResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkPMS, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNat {
		var resp *config.NetWorkNatResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNat, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkWifi {
		var resp *config.NetWorkWifiResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkWifi, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigFVideoGUISet {
		var resp *config.FVideoGUISetResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.FVideoGUISet, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigStorageStoragePosition {
		var resp *config.StorageStoragePositionResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.StorageStoragePosition, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigStorageStorageNotExist {
		var resp *config.StorageStorageNotExistResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.StorageStorageNotExist, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigStorageStorageLowSpace {
		var resp *config.StorageStorageLowSpaceResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.StorageStorageLowSpace, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigStorageStorageFailure {
		var resp *config.StorageStorageFailureResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.StorageStorageFailure, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigGeneralGeneral {
		var resp *config.GeneralGeneralResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.GeneralGeneral, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigGeneralLocation {
		var resp *config.GeneralLocationResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.GeneralLocation, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigCameraParam {
		var resp *config.CameraParamResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.CameraParam, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigCameraParamEx {
		var resp *config.CameraParamExResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.CameraParamEx, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigSimplifyEncode {
		var resp *config.SimplifyEncodeResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.SimplifyEncode, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAVEncMultiChannelEncode {
		var resp *config.AVEncMultiChannelEncodeResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AVEncMultiChannelEncode, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigfVideoTVAdjust {
		var resp *config.FVideoTVAdjustResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.FVideoTVAdjust, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigfVideoVideoOut {
		var resp *config.FVideoVideoOutResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.FVideoVideoOut, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigfVideoPlay {
		var resp *config.FVideoPlayResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.FVideoPlay, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigfVideoAudioInFormat {
		var resp *config.FVideoAudioInFormatResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.FVideoAudioInFormat, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigfVideoTour {
		var resp *config.FVideoTourResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.FVideoTour, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigfVideoVideoOutPriority {
		var resp *config.FVideoVideoOutPriorityResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.FVideoVideoOutPriority, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAnalyze {
		var resp *config.AnalyzeResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.Analyze, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigVideoDec {
		var resp *config.VideoDecResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.VideoDec, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigVideoChannel {
		var resp *config.VideoChannelResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.VideoChannel, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAbilityEncodePower {
		var resp *config.AbilityEncodePowerResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AbilityEncodePower, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigGeneralSystemState {
		var resp *config.GeneralSystemStateResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.GeneralSystemState, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigMediaDecodeParam {
		var resp *config.MediaDecodeParamResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.MediaDecodeParam, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAlarmLocalAlarm {
		var resp *config.AlarmLocalAlarmResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AlarmLocalAlarm, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAlarmAlarmOut {
		var resp *config.AlarmAlarmOutResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AlarmAlarmOut, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigUartPTZPreset {
		var resp *config.UartPTZPresetResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.UartPTZPreset, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigUartPTZTour {
		var resp *config.UartPTZTourResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.UartPTZTour, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkRemoteDeviceV3 {
		var resp *config.NetWorkRemoteDeviceV3Resp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkRemoteDeviceV3, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetDDNS {
		var resp *config.NetWorkNetDDNSResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetDDNS, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetPPPOE {
		var resp *config.NetWorkNetPPPOEResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetPPPOE, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetARSP {
		var resp *config.NetWorkNetARSPResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetARSP, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkRemoteDevice {
		var resp *config.NetWorkRemoteDeviceResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkRemoteDevice, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNet3G {
		var resp *config.NetWorkNet3GResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNet3G, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkNetMobile {
		var resp *config.NetWorkNetMobileResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkNetMobile, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkDigTimeSyn {
		var resp *config.NetWorkDigTimeSynResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkDigTimeSyn, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkShortMsg {
		var resp *config.NetWorkShortMsgResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkShortMsg, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigNetWorkMultimediaMsg {
		var resp *config.NetWorkMultimediaMsgResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetWorkMultimediaMsg, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigAlarmIPCAlarm {
		var resp *config.AlarmIPCAlarmResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AlarmIPCAlarm, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigGeneralAutoMaintain {
		var resp *config.GeneralAutoMaintainResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.GeneralAutoMaintain, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigUsers {
		var resp *config.UsersResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.Users, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigGroups {
		var resp *config.GroupsResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.Groups, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == utils.ConfigSystemExUserMap {
		var resp *config.SystemExUserMapResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.SystemExUserMap, nil
		}
		return nil, errors.New(resp.RetMsg)
	}

	return
}

// Obtain device configuration Take the DAS configuration of device active registration service as an example
func (jDevice *JLinkDevice) GetConfigHttp(name string) (res []byte, err error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DGetconfigUrl, jDevice.Jdtoken)
	payload := strings.NewReader(`{"Name":"` + name + `"}`)
	resbody, err := jDevice.HttpPost(url, payload)

	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return
	}
	// fmt.Println(string(resbody))
	resp := &utils.Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Unmarshal err:" + err.Error())
		return
	}

	if resp.Code != 2000 {
		log.Println("GetConfig err：" + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	resByre, err := json.Marshal(resp.Data)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return
	}

	return resByre, nil
}

// Setting configuration Take the device actively registering the DAS service as an example
func (jDevice *JLinkDevice) SetConfig(gc interface{}) (bool, error) {
	parm := make(map[string]interface{})
	switch v := gc.(type) {
	case config.AVEncEncodeResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AVEncEncode
	case config.AVEncVideoWidgetResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AVEncVideoWidget
	case config.AVEncVideoColorResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AVEncVideoColor
	case config.AVEncSmartH264Resp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AVEncSmartH264
	case config.FVideoVolumeResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.FVideoVolume
	case config.FVideoOSDInfoResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.FVideoOSDInfo
	case config.RecordResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.Record
	case config.StorageSnapshotResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.StorageSnapshot
	case config.DetectBlindDetectResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.DetectBlindDetect
	case config.DetectMotionDetectResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.DetectMotionDetect
	case config.DetectLossDetectResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.DetectLossDetect
	case config.ChannelTitleResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.ChannelTitle
	case config.CameraResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.Camera
	case config.DetectHumanDetectionResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.DetectHumanDetection
	case config.DetectHumanDetection0Resp:
		parm["Name"] = v.Name
		parm[v.Name] = v.DetectHumanDetection0
	case config.AbilityVoiceTipTypeResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AbilityVoiceTipType
	case config.UartCommResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.UartComm
	case config.UartPTZResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.UartPTZ
	case config.UartRS485Resp:
		parm["Name"] = v.Name
		parm[v.Name] = v.UartRS485
	case config.NetWorkNetCommonResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetCommon
	case config.NetWorkNetDHCPResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetDHCP
	case config.NetWorkNetEmailResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetEmail
	case config.NetWorkNetNTPResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetNTP
	case config.NetWorkNetDNSResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetDNS
	case config.NetWorkNetFTPResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetFTP
	case config.AlarmNetIPConflictResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AlarmNetIPConflict
	case config.AlarmNetAbortResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AlarmNetAbort
	case config.NetWorkUpnpResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkUpnp
	case config.NetWorkNetIPFilterResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetIPFilter
	case config.OPUTCTimeSettingResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.OPUTCTimeSetting
	case config.NetWorkNetOrderResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetOrder
	case config.NetWorkDASResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkDAS
	case config.NetWorkPMSResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkPMS
	case config.NetWorkNatResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNat
	case config.NetWorkWifiResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkWifi
	case config.FVideoGUISetResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.FVideoGUISet
	case config.StorageStoragePositionResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.StorageStoragePosition
	case config.StorageStorageNotExistResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.StorageStorageNotExist
	case config.StorageStorageLowSpaceResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.StorageStorageLowSpace
	case config.StorageStorageFailureResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.StorageStorageFailure
	case config.GeneralGeneralResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.GeneralGeneral
	case config.GeneralLocationResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.GeneralLocation
	case config.CameraParamResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.CameraParam
	case config.CameraParamExResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.CameraParamEx
	case config.SimplifyEncodeResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.SimplifyEncode
	case config.AVEncMultiChannelEncodeResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AVEncMultiChannelEncode
	case config.FVideoTVAdjustResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.FVideoTVAdjust
	case config.FVideoVideoOutResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.FVideoVideoOut
	case config.FVideoPlayResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.FVideoPlay
	case config.FVideoAudioInFormatResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.FVideoAudioInFormat
	case config.FVideoTourResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.FVideoTour
	case config.FVideoVideoOutPriorityResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.FVideoVideoOutPriority
	case config.AnalyzeResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.Analyze
	case config.VideoDecResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.VideoDec
	case config.VideoChannelResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.VideoChannel
	case config.AbilityEncodePowerResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AbilityEncodePower
	case config.GeneralSystemStateResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.GeneralSystemState
	case config.MediaDecodeParamResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.MediaDecodeParam
	case config.AlarmLocalAlarmResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AlarmLocalAlarm
	case config.AlarmAlarmOutResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AlarmAlarmOut
	case config.UartPTZPresetResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.UartPTZPreset
	case config.UartPTZTourResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.UartPTZTour
	case config.NetWorkRemoteDeviceV3Resp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkRemoteDeviceV3
	case config.NetWorkNetDDNSResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetDDNS
	case config.NetWorkNetPPPOEResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetPPPOE
	case config.NetWorkNetARSPResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetARSP
	case config.NetWorkRemoteDeviceResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkRemoteDevice
	case config.NetWorkNet3GResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNet3G
	case config.NetWorkNetMobileResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkNetMobile
	case config.NetWorkDigTimeSynResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkDigTimeSyn
	case config.NetWorkShortMsgResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkShortMsg
	case config.NetWorkMultimediaMsgResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.NetWorkMultimediaMsg
	case config.AlarmIPCAlarmResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.AlarmIPCAlarm
	case config.GeneralAutoMaintainResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.GeneralAutoMaintain
	case config.UsersResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.Users
	case config.GroupsResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.Groups
	case config.SystemExUserMapResp:
		parm["Name"] = v.Name
		parm[v.Name] = v.SystemExUserMap
	default:
		log.Println("以上类型都不是")
	}

	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return false, err
	}
	fmt.Println(string(postData))
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DSetconfigUrl, jDevice.Jdtoken)
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return false, err
	}

	fmt.Println(string(resbody))
	resp := &utils.XYResponse{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if resp.Code == 2000 && resp.Data.Ret == 100 {
		return true, nil
	}
	log.Println("SetConfig err：" + resp.Msg)
	return false, errors.New(resp.Msg)

}
