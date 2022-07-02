package jrgd

import (
	"errors"
	"jlink-restful-golang-demo/opedev"
)

//Device control PTZ control as an example
func (jDevice *JLinkDevice) DeviceOperate(pdcd interface{}) (data interface{}, err error) {
	switch v := pdcd.(type) {
	case *opedev.PtzDirectionControlDTO:
		data, err = opedev.OpeDevOPPTZControl(v, jDevice.Jdtoken)
	case *opedev.OPMachineActionReq:
		data, err = opedev.OPMachineControl(v, jDevice.Jdtoken)
	case *opedev.OPDefaultConfigReq:
		data, err = opedev.OpeDevOPDefaultConfig(v, jDevice.Jdtoken)
	case *opedev.StorageInfoReq:
		data, err = opedev.OpeDevStorageInfo(v, jDevice.Jdtoken)
	case *opedev.OPStorageManagerReq:
		data, err = opedev.OpeDevPartitionSize(v, jDevice.Jdtoken)
	case *opedev.OPFileQueryReq:
		data, err = opedev.OpeDevOPFileQuery(v, jDevice.Jdtoken)
	case *opedev.OPTimeQueryReq:
		data, err = opedev.OpeDevOPTimeQuery(v, jDevice.Jdtoken)
	case *opedev.OPSystemConfigReq:
		data, err = opedev.OpeDevOPSystemConfig(v, jDevice.Jdtoken)
	case *opedev.OPTimeSettingReq:
		data, err = opedev.OpeDevOPTimeSetting(v, jDevice.Jdtoken)
	case *opedev.OPLogQueryReq:
		data, err = opedev.OpeDevOPLogQuery(v, jDevice.Jdtoken)
	default:
		return nil, errors.New("No operation of this type")
	}
	return
}
