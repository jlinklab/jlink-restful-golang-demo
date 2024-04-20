package opedev

import (
	"errors"
	v3 "jlink-restful-golang-demo/v3"
)

// Device control PTZ control as an example
func DeviceOperate(jDevice *v3.JLinkDevice, pdcd interface{}) (data interface{}, err error) {
	switch v := pdcd.(type) {
	case *PtzDirectionControlDTO:
		data, err = OpeDevOPPTZControl(jDevice, v)
	case *OPMachineActionReq:
		data, err = OPMachineControl(jDevice, v)
	case *OPDefaultConfigReq:
		data, err = OpeDevOPDefaultConfig(jDevice, v)
	case *StorageInfoReq:
		data, err = OpeDevStorageInfo(jDevice, v)
	case *OPStorageManagerReq:
		data, err = OpeDevPartitionSize(jDevice, v)
	case *OPFileQueryReq:
		data, err = OpeDevOPFileQuery(jDevice, v)
	case *OPTimeQueryReq:
		data, err = OpeDevOPTimeQuery(jDevice, v)
	case *OPSystemConfigReq:
		data, err = OpeDevOPSystemConfig(jDevice, v)
	case *OPTimeSettingReq:
		data, err = OpeDevOPTimeSetting(jDevice, v)
	case *OPLogQueryReq:
		data, err = OpeDevOPLogQuery(jDevice, v)
	default:
		return nil, errors.New("No operation of this type")
	}
	return
}
