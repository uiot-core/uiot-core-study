package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"strconv"
	"./uiot"
	"time"
)

var uiotClient *uiot.UIotClient

const productSN = "tr7xjdvgm5pirv1w"
const deviceSN = "12345678"
const region = "cn-sh2"
const zone = "cn-sh2-02"
const projectid = "org-vk3oxc"

//const projectid = "org-vk3oxc"

func main() {
	cfg, credential := loadConfig()
	cfg.BaseUrl = "https://api-cn-sh2.iot.ucloud.cn"
	uiotClient = uiot.NewClient(cfg, credential)

	/* Send downlink message */
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			PublishUIoTCoreMQTTMessage()
		}
	}

	//UpdateUIoTCoreDeviceShadow()
	//SendUIoTCoreDeviceCommand()
	//CreateUIoTCoreDevice()
	//BatchCreateUIoTCoreDevice()
	//BatchCreateUIoTCoreDeviceWithSN()

}

func SetUIoTCoreDeviceProperty() {
	req := uiot.SetUIoTCoreDevicePropertyRequest{}
	req.ProductSN = ucloud.String(productSN)
	req.DeviceSN = ucloud.String(deviceSN)
	req.Zone = ucloud.String(zone)
	req.Property = ucloud.String(ToBase64(`{"temperature":15}`))
	req.Desired = true
	resp, err := uiotClient.SetUIoTCoreDeviceProperty(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		strresp, _ := json.Marshal(resp)
		fmt.Println(string(strresp))
	}
}

var i int = 0
func PublishUIoTCoreMQTTMessage() {
	req := uiot.PublishUIoTCoreMQTTMessageRequest{}
	req.ProductSN = ucloud.String(productSN)
	req.DeviceSN = ucloud.String(deviceSN)

	/*
	// binary
	req.TopicFullName = ucloud.String(`/` + productSN + `/` + deviceSN + `/sub`)
	bs := []byte{0x0, 0x22, 0x23, 0x87, 0xF3, 0xA1}
	req.MessageContent = ucloud.String(base64.StdEncoding.EncodeToString(bs))
	req.Qos = 0
	resp, err := uiotClient.PublishUIoTCoreMQTTMessage(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		strresp, _ := json.Marshal(resp)
		fmt.Println(string(strresp))
	}
	*/

	// json message
	i++
	req.TopicFullName = ucloud.String(`/` + productSN + `/` + deviceSN + `/set`)
	req.MessageContent = ucloud.String(ToBase64(`{"downlink-test":` + strconv.Itoa(i) + `}`))
	resp, err := uiotClient.PublishUIoTCoreMQTTMessage(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		strresp, _ := json.Marshal(resp)
		fmt.Println(string(strresp))
	}
}

func UpdateUIoTCoreDeviceShadow() {
	req := uiot.UpdateUIoTCoreDeviceShadowRequest{}
	req.ProductSN = ucloud.String(productSN)
	req.DeviceSN = ucloud.String(deviceSN)
	req.Zone = ucloud.String(zone)
	req.Desired = ucloud.String(ToBase64(`{"temperature":17}`))
	req.ShadowVersion = ucloud.String("1")
	resp, err := uiotClient.UpdateUIoTCoreDeviceShadow(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		strresp, _ := json.Marshal(resp)
		fmt.Println(string(strresp))
	}
}

func SendUIoTCoreDeviceCommand() {
	req := uiot.SendUIoTCoreDeviceCommandRequest{}
	req.ProductSN = ucloud.String(productSN)
	req.DeviceSN = ucloud.String(deviceSN)
	req.Zone = ucloud.String(zone)
	req.Identifier = ucloud.String("downtemperature")
	req.Input = ucloud.String(ToBase64(`{"downvalue":10}`))
	req.Method = ucloud.String("async")
	req.Timeout = ucloud.Int(15)
	resp, err := uiotClient.SendUIoTCoreDeviceCommand(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		strresp, _ := json.Marshal(resp)
		fmt.Println(string(strresp))
	}
}

func CreateUIoTCoreDevice() {
	req := uiot.CreateUIoTCoreDeviceRequest{}
	req.ProductSN = ucloud.String(productSN)
	//req.DeviceSN = ucloud.String(`78:982:00020202`)
	req.Zone = ucloud.String(zone)
	resp, err := uiotClient.CreateUIoTCoreDevice(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		strresp, _ := json.Marshal(resp)
		fmt.Println(string(strresp))
	}
}

func BatchCreateUIoTCoreDevice() {
	req := uiot.BatchCreateUIoTCoreDeviceRequest{}
	req.ProductSN = ucloud.String(productSN)
	//req.DeviceSN = ucloud.String(`78:982:00020202`)
	req.DeviceCount = 10
	req.Zone = ucloud.String(zone)
	resp, err := uiotClient.BatchCreateUIoTCoreDevice(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		strresp, _ := json.Marshal(resp)
		fmt.Println(string(strresp))
	}
}

func BatchCreateUIoTCoreDeviceWithSN() {
	req := uiot.BatchCreateUIoTCoreDeviceWithSNRequest{}
	req.ProductSN = ucloud.String(productSN)
	//req.DeviceSN = ucloud.String(`78:982:00020202`)
	req.DeviceSN = []string{"1234567", "1236786", "876544"}
	req.Zone = ucloud.String(zone)
	resp, err := uiotClient.BatchCreateUIoTCoreDeviceWithSN(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		strresp, _ := json.Marshal(resp)
		fmt.Println(string(strresp))
	}
}

func QueryUIoTCoreDeviceFileList() string {
	req := uiot.QueryUIoTCoreDeviceFileListRequest{}
	req.ProductSN = ucloud.String(productSN)
	req.DeviceSN = ucloud.String(deviceSN)
	req.Zone = ucloud.String(zone)

	resp, err := uiotClient.QueryUIoTCoreDeviceFileList(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		/*
		strresp,_ := json.Marshal(resp)
		fmt.Println(string(strresp))
		*/
		if resp.TotalCount > 0 {
			return resp.FileList[0].URL
		}
	}

	return "error"
}

func GetUIoTCoreDeviceFileURL(FileName string) string {
	req := uiot.GetUIoTCoreDeviceFileURLRequest{}
	req.ProductSN = ucloud.String(productSN)
	req.DeviceSN = ucloud.String(deviceSN)
	req.Zone = ucloud.String(zone)
	req.FileName = ucloud.String(FileName)
	resp, err := uiotClient.GetUIoTCoreDeviceFileURL(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		/*
		strresp,_ := json.Marshal(resp)
		fmt.Println(string(strresp))
		*/
		return resp.URL
	}
	return "error"
}

const privateKey = "your_privateKey"
const publicKey = "your_publicKey"

func loadConfig() (*ucloud.Config, *auth.Credential) {
	cfg := ucloud.NewConfig()
	cfg.LogLevel = log.DebugLevel
	cfg.LogLevel = log.ErrorLevel
	cfg.Region = region
	cfg.ProjectId = projectid
	cfg.Zone = zone

	credential := auth.NewCredential()
	credential.PrivateKey = privateKey
	credential.PublicKey = publicKey

	log.Info("setup clients ...")

	return &cfg, &credential
}

func ToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
