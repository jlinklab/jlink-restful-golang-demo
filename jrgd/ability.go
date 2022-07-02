package jrgd

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/ability"
	"jlink-restful-golang-demo/http"
	"log"
)

func (jDevice *JLinkDevice) DeviceAbility(name string) (da interface{}, err error) {
	res, err := jDevice.DeviceAbilityHttp(name)
	if err != nil {
		return
	}

	if name == AbilitySystemFunction {
		var resp *ability.SystemFunctionResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.SystemFunction, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityBlindCapability {
		var resp *ability.BlindCapabilityResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.BlindCapability, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityEncodeCapability {
		var resp *ability.EncodeCapabilityResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.EncodeCapability, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityMotionArea {
		var resp *ability.MotionAreaResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.MotionArea, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityDDNSService {
		var resp *ability.DDNSServiceResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.DDNSService, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityComProtocol {
		var resp *ability.ComProtocolResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.ComProtocol, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityCamera {
		var resp *ability.CameraResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.Camera, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityTalkAudioFormat {
		var resp *ability.TalkAudioFormatResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.TalkAudioFormat, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityMultiLanguage {
		var resp *ability.MultiLanguageResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.MultiLanguage, nil
		}
	} else if name == AbilityUartProtocol {
		var resp *ability.UartProtocolResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.UartProtocol, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityIntelligent {
		var resp *ability.IntelligentResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.Intelligent, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityCameraParamEx {
		var resp *ability.CameraParamExResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.CameraParamEx, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityAbilitySupportVstd {
		var resp *ability.AbilitySupportVstdResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AbilitySupportVstd, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityAbilityDeviceDesc {
		var resp *ability.AbilityDeviceDescResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AbilityDeviceDesc, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityAbilityVGAResolution {
		var resp *ability.AbilityVGAResolutionResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AbilityVGAResolution, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityAbilityGUIThemeList {
		var resp *ability.AbilityGUIThemeListResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AbilityGUIThemeList, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityAbilityDigitalReal {
		var resp *ability.AbilityDigitalRealResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.AbilityDigitalReal, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityVideoAbility {
		var resp *ability.VideoAbilityResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.VideoAbility, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityNetAbility {
		var resp *ability.NetAbilityResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.NetAbility, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityEncStaticParam {
		var resp *ability.EncStaticParamResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.EncStaticParam, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityMaxPreRecord {
		var resp *ability.MaxPreRecordResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.MaxPreRecord, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityMultiChannel {
		var resp *ability.MultiChannelResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.MultiChannel, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityChannelSystemFunctionSupportPeaInHumanPed {
		var resp *ability.ChannelSystemFunctionSupportPeaInHumanPedResp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.Body, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityChannelSystemFunctionSupportFaceDetectV2 {
		var resp *ability.ChannelSystemFunctionSupportFaceDetectV2Resp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.ChannelSystemFunctionSupportFaceDetectV2, nil
		}
		return nil, errors.New(resp.RetMsg)
	} else if name == AbilityChannelSystemFunctionSmartH264 {
		var resp *ability.ChannelSystemFunctionSmartH264Resp
		err = json.Unmarshal(res, &resp)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Ret == 100 {
			return resp.Body, nil
		}
		return nil, errors.New(resp.RetMsg)
	}
	return

}

// Get Device Capability Set
func (jDevice *JLinkDevice) DeviceAbilityHttp(name string) (res []byte, err error) {
	parm := make(map[string]interface{})
	parm["Name"] = name
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("err:" + err.Error())
		return
	}

	resbody, err := http.HttpPost(dAbilityUrl, jDevice.Jdtoken, postData)
	if err != nil {
		log.Println("httpPosterr:" + err.Error())
		return
	}

	resp := &Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println(err)
		return
	}
	if resp.Code != 2000 {
		return res, errors.New(resp.Msg)
	}

	resByre, resByteErr := json.Marshal(resp.Data)
	if resByteErr != nil {
		log.Println(resByteErr)
		return
	}
	return resByre, nil

}
