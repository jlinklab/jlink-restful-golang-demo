# OpenApi SDK V3 Demo

## Initialization JLinkClient

> kUuid,kAppKey,KAppSecret,kMovecard Obtained from (open.jftech.com).

```
jClient := jrgd.NewJLinkClient(kUuid, kAppKey, KAppSecret, kMovecard)
```

## Initialization JLinkDevice

> kDeviceId is the device serial number
> devUsername indicates the device login username
> devPassword indicates the device login password

```
jDevice := jrgd.NewJLinkDevice(jClient, kDeviceId, devUsername, devPassword)
```
## Device Bind(The v3 version requires the device to be bound to the user's account in order to operate)

```
DeviceBind,err := jDevice.DeviceBind()
```

## Device UnBind

```
DeviceUnBind,err := jDevice.DeviceUnBind()
```

## Get Device List

```
DeviceList,err := jDevice.DeviceList(1, 10)
```


## Get Device Token

```
DeviceToken,err := jDevice.GetDeviceTokenV3()
```

## Get Device LoginToken

```
err := jDevice.LoginToken()
```

## Get Device Status

```
JDStatus,err := jDevice.DeviceStatus()
```


## Device Login

> 1.Login with devUsername and decPassword

```
jDLogin,err := jDevice.DeviceLoginPassWord()
```

> 2.Log in to the device with the Token obtained from device sharing

```
jDLogin,err = jDevice.DeviceLoginToken()
```

## Get Device Information

> Take obtaining system information as an example

```
devInfo,err := jDevice.DeviceGetinfo(utils.SystemInfo)
```

## Device Capability Set Acquisition

> Obtain the device capability set after the device is successfully logged in. Take obtaining the system capability set as an example

```
jDAbility,err := jDevice.DeviceAbility(utils.SystemFunction)
```

## Get Device Configuration

> Take obtaining the device active registration service DAS configuration as an example

```
jDConfig, err := jDevice.DeviceGetconfig(utils.ConfigNetWorkDAS)
```

## Set Device Configuration

> Take the device actively registering the DAS service as an example enable: enable serverAddr: server address Port: server port

```
	if op, ok := jDConfig.(config.NetWorkDAS); ok {
		op.Enable = false
		nw := config.NetWorkDASResp{Name: utils.ConfigNetWorkDAS, NetWorkDAS: op}
		setConfigFlag, err := jDevice.SetConfig(nw)
		fmt.Println(setConfigFlag, err)
	}

```

## Device Keepalive

```
keepalive,err := jDevice.DeviceKeepalive()
```

## Device Operate

> Take the PTZ direction control as an example

```
ptzDirectionControlDTO := new(opedev.PtzDirectionControlDTO)
ptzDirectionControlDTO.Command = utils.DirectionLeft
ptzDirectionControlDTO.Parameter.Channel = 0
ptzDirectionControlDTO.Parameter.Preset = 65535
ptzDirectionControlDTO.Parameter.Step = 6
opedev.DeviceOperate(jDevice,ptzDirectionControlDTO)
```

## Get the liveStream

> Take obtaining the live address of the main stream of the device in hls format as an example;
> rs Account represents the user registered by the open platform through the RS interface, 
> and rs Pass represents the user password registered by the open platform through the RS interface

```
STREAM_EXTRA := "1"
MEDIATYPE_HLS := "hls"
CHANNEL := "0"
PROTOCOL_TS := "ts"
Livestream, err := jDevice.Livestream(MEDIATYPE_HLS, CHANNEL, STREAM_EXTRA, PROTOCOL_TS)
```

## Device Snapshot

>  Capture a picture of a channel, The channel number defaults to 0

```
captureUrl ,err:= jDevice.Capture(&jrgd.CaptureRequest{Channel: 0, PicType: 0})
```

## Low-power Device Wakeup

```
wakeupFlag,err := jDevice.DeviceWakeup("1")
```

## IOT DoorLock Device SetTempPassword

>  Set temporary password for door lock equipment

```
doorLock.DoorLockSetTempPassword(jDevice, []doorLock.TempPassword{ 
		{
			EndTime:   "2024-04-19 18:45:56",
			Password:  "123456",
			StartTime: "2024-04-19 10:15:56",
			ValidNum:  1,
		}})
```

## IOT DoorLock Device RemoteUnlock

>  Remote unlocking of door lock equipment

```
doorLock.DoorLockRemoteUnlock(jDevice, "password")
```