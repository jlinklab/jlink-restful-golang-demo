# restful_api_demo

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

## Get Device Token

```
JDToken,err := jDevice.GetDeviceToken()
```

## Get Device Status

```
JDStatus,err := jDevice.DeviceStatus()
```

## Low-power Device Wakeup

```
wakeupFlag,err := jDevice.WakeUp()
```

## Device Login

> 1.Login with devUsername and decPassword

```
jDLogin,err := jDevice.Login()
```

> 2.Log in to the device with the Token obtained from device sharing

```
jDLogin,err = jDevice.DeviceLoginByToken(loginToken)
```

## Get Device Information

> Take obtaining system information as an example

```
devInfo,err := jDevice.DeviceInfo(jrgd.SystemInfo)
```

## Device Capability Set Acquisition

> Obtain the device capability set after the device is successfully logged in. Take obtaining the system capability set as an example

```
jDAbility,err := jDevice.DeviceAbility(jrgd.SystemFunction)
```

## Get Device Configuration

> Take obtaining the device active registration service DAS configuration as an example

```
jDConfig,err := jDevice.GetConfig(jrgd.NetWorkDas)
```

## Set Device Configuration

> Take the device actively registering the DAS service as an example enable: enable serverAddr: server address Port: server port

```
	if op, ok := jDConfig.(config.NetWorkDAS); ok {
		op.Enable = false
		nw := config.NetWorkDASResp{Name: jrgd.ConfigNetWorkDAS, NetWorkDAS: op}
		setConfigFlag, err := jDevice.SetConfig(nw)
		fmt.Println(setConfigFlag, err)
	}

```

## Device Keepalive

```
keepalive,err := jDevice.Keepalive()
```

## Device Operate

> Take the PTZ direction control as an example

```
ptzDirectionControlDTO := new(jrgd.PtzDirectionControlDTO)
ptzDirectionControlDTO.Command = jrgd.DirectionLeft
ptzDirectionControlDTO.Parameter.Channel = 0
ptzDirectionControlDTO.Parameter.Preset = 65535
ptzDirectionControlDTO.Parameter.Step = 6
jDevice.DeviceOperate(ptzDirectionControlDTO)
```

## Get the liveStream

> Take obtaining the live address of the main stream of the device in hls format as an example;
> rs Account represents the user registered by the open platform through the RS interface, 
> and rs Pass represents the user password registered by the open platform through the RS interface

```
jUser := jrgd.NewJLinkUser(jClient, userName, password)
usertoken,err := jUser.GetUserToken()
fmt.Println(usertoken)
STREAM_EXTRA := "1"
MEDIATYPE_HLS := "hls"
CHANNEL := "0"
PROTOCOL_TS := "ts"
liveStream,err := jDevice.DeviceLivePlayUrl(STREAM_EXTRA, MEDIATYPE_HLS, PROTOCOL_TS, CHANNEL, jUser)
```

## Device Snapshot

>  Capture a picture of a channel, The channel number defaults to 0

```
channel := 0
captureUrl ,err:= jDevice.Capture(channel)
```