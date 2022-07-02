package jrgd

import (
	"encoding/json"
	"errors"
	"fmt"
	"jlink-restful-golang-demo/config"
	"jlink-restful-golang-demo/http"
	"log"
)

type XYResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data Data   `json:"data"`
}

type Data struct {
	Name string `json:"name"`
	Ret  int    `json:"ret"`
}

func (jDevice *JLinkDevice) GetConfig(name string) (da interface{}, err error) {
	res, err := jDevice.GetConfigHttp(name)
	if err != nil {
		log.Println("GetConfig err：" + err.Error())
		return
	}

	if name == ConfigAVEncEncode {
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
	} else if name == ConfigAVEncEncode {
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
	} else if name == ConfigAVEncVideoWidget {
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
	} else if name == ConfigAVEncVideoColor {
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
	} else if name == ConfigAVEncSmartH264 {
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
	} else if name == ConfigfVideoVolume {
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
	} else if name == ConfigfVideoOSDInfo {
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
	} else if name == ConfigRecord {
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
	} else if name == ConfigStorageSnapshot {
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
	} else if name == ConfigDetectBlindDetect {
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
	} else if name == ConfigDetectMotionDetect {
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
	} else if name == ConfigDetectLossDetect {
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
	} else if name == ConfigChannelTitle {
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
	} else if name == ConfigCamera {
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
	} else if name == ConfigDetectHumanDetection {
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
	} else if name == ConfigDetectHumanDetection0 {
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
	} else if name == ConfigAbilityVoiceTipType {
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
	} else if name == ConfigUartComm {
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
	} else if name == ConfigUartPTZ {
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
	} else if name == ConfigUartRS485 {
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
	} else if name == ConfigNetWorkNetCommon {
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
	} else if name == ConfigNetWorkNetDHCP {
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
	} else if name == ConfigNetWorkNetEmail {
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
	} else if name == ConfigNetWorkNetNTP {
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
	} else if name == ConfigNetWorkNetDNS {
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
	} else if name == ConfigNetWorkNetFTP {
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
	} else if name == ConfigAlarmNetIPConflict {
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
	} else if name == ConfigAlarmNetAbort {
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
	} else if name == ConfigNetWorkUpnp {
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
	} else if name == ConfigNetWorkNetIPFilter {
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
	} else if name == ConfigOPUTCTimeSetting {
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
	} else if name == ConfigNetWorkNetOrder {
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
	} else if name == ConfigNetWorkDAS {
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
	} else if name == ConfigNetWorkPMS {
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
	} else if name == ConfigNetWorkNat {
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
	} else if name == ConfigNetWorkWifi {
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
	} else if name == ConfigFVideoGUISet {
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
	} else if name == ConfigStorageStoragePosition {
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
	} else if name == ConfigStorageStorageNotExist {
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
	} else if name == ConfigStorageStorageLowSpace {
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
	} else if name == ConfigStorageStorageFailure {
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
	} else if name == ConfigGeneralGeneral {
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
	} else if name == ConfigGeneralLocation {
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
	} else if name == ConfigCameraParam {
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
	} else if name == ConfigCameraParamEx {
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
	} else if name == ConfigSimplifyEncode {
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
	} else if name == ConfigAVEncMultiChannelEncode {
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
	} else if name == ConfigfVideoTVAdjust {
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
	} else if name == ConfigfVideoVideoOut {
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
	} else if name == ConfigfVideoPlay {
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
	} else if name == ConfigfVideoAudioInFormat {
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
	} else if name == ConfigfVideoTour {
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
	} else if name == ConfigfVideoVideoOutPriority {
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
	} else if name == ConfigAnalyze {
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
	} else if name == ConfigVideoDec {
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
	} else if name == ConfigVideoChannel {
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
	} else if name == ConfigAbilityEncodePower {
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
	} else if name == ConfigGeneralSystemState {
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
	} else if name == ConfigMediaDecodeParam {
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
	} else if name == ConfigAlarmLocalAlarm {
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
	} else if name == ConfigAlarmAlarmOut {
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
	} else if name == ConfigUartPTZPreset {
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
	} else if name == ConfigUartPTZTour {
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
	} else if name == ConfigNetWorkRemoteDeviceV3 {
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
	} else if name == ConfigNetWorkNetDDNS {
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
	} else if name == ConfigNetWorkNetPPPOE {
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
	} else if name == ConfigNetWorkNetARSP {
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
	} else if name == ConfigNetWorkRemoteDevice {
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
	} else if name == ConfigNetWorkNet3G {
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
	} else if name == ConfigNetWorkNetMobile {
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
	} else if name == ConfigNetWorkDigTimeSyn {
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
	} else if name == ConfigNetWorkShortMsg {
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
	} else if name == ConfigNetWorkMultimediaMsg {
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
	} else if name == ConfigAlarmIPCAlarm {
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
	} else if name == ConfigGeneralAutoMaintain {
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
	} else if name == ConfigUsers {
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
	} else if name == ConfigGroups {
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
	} else if name == ConfigSystemExUserMap {
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
	parm := make(map[string]interface{})
	parm["Name"] = name
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return
	}
	resbody, err := http.HttpPost(dGetconfigUrl, jDevice.Jdtoken, postData)
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return
	}
	// fmt.Println(string(resbody))
	resp := &Response{}
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
	fmt.Println(gc)
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
	resbody, err := http.HttpPost(dSetconfigUrl, jDevice.Jdtoken, postData)
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return false, err
	}
	fmt.Println(string(resbody))
	resp := &XYResponse{}
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
