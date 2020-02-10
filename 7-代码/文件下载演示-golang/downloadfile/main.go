package main

import (
	"./uiot"
	"./utils"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"io"
	"net/http"
	"os"
	"time"
)

var uiotClient *uiot.UIotClient

var clientget = &http.Client{Transport: &http.Transport{
	TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	MaxConnsPerHost:     100,
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
},
	Timeout: time.Second * 1000,
}

const productSN = "ozuz63kum2i4djb3"
const deviceSN = "afnyhnizq9l4l9ev"
const region = "cn-sh2"
const zone = "cn-sh2-02"
const projectid = "org-z44lmf"

func main() {
	// get the url
	cfg, credential := loadConfig()
	cfg.BaseUrl = "https://api-cn-sh2.iot.ucloud.cn"
	uiotClient = uiot.NewClient(cfg, credential)

	url := QueryUIoTCoreDeviceFileList()
	if url == "error" {
		fmt.Println("error filename")
		return
	}

	// get the file
	resp, err := clientget.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	filename := utils.GetRandomString(16)
	utils.WriteFile("./files/"+filename+".download",resp.Body)

	fileupload, err := os.Open("./files/"+filename+".download")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fileupload.Close()

	h := md5.New()
	if _, err := io.Copy(h, fileupload); err != nil {
		fmt.Println(err)
		return
	}
	md5string := hex.EncodeToString(h.Sum(nil))
	fmt.Println("md5sum:"+md5string)

	//io.Copy(ioutil.Discard, resp.Body)
	time.Sleep(500 * time.Millisecond)
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

func PublishUIoTCoreMQTTMessage() {
		req := uiot.PublishUIoTCoreMQTTMessageRequest{}
		req.ProductSN = ucloud.String(productSN)
		req.DeviceSN = ucloud.String(deviceSN)
		req.TopicFullName = ucloud.String(`/` + productSN + `/` + deviceSN + `/sub`)
		bs := []byte{0x0, 0x22, 0x23, 0x87, 0xF3, 0xA1}
		req.MessageContent = ucloud.String(base64.StdEncoding.EncodeToString(bs))
		req.Qos = 1
		resp, err := uiotClient.PublishUIoTCoreMQTTMessage(&req)

		if err != nil {
			fmt.Println(err)
		} else {
			strresp, _ := json.Marshal(resp)
			_ = strresp
			//fmt.Println(string(strresp))
		}

		// json message
		req.TopicFullName = ucloud.String(`/` + productSN + `/` + deviceSN + `/downlink`)
		req.MessageContent = ucloud.String(ToBase64(`{"test":1`))
		resp, err = uiotClient.PublishUIoTCoreMQTTMessage(&req)

		if err != nil {
			fmt.Println(err)
		} else {
			strresp, _ := json.Marshal(resp)
			_ = strresp
			//fmt.Println(string(strresp))
		}

		//fmt.Println(g_count,g_maxtime,g_totaltime/float64(g_count))
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
	cfg.BaseUrl = "https://api-cn-sh2.iot.ucloud.cn"

	credential := auth.NewCredential()
	credential.PrivateKey = privateKey
	credential.PublicKey = publicKey

	log.Info("setup clients ...")

	return &cfg, &credential
}

func ToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

