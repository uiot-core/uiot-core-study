package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"./uploadfile"
	"time"
)

const productSN = "ozuz63kum2i4djb3"
const deviceSN = "afnyhnizq9l4l9ev"
const deviceSecret = "3ksk8dbg8ny3z3cf"
const region = "cn-sh2"

var client = &http.Client{Transport: &http.Transport{
	TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	MaxConnsPerHost:     100,
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
},
	Timeout: time.Second * 1000,
}

func main() {

	file_path, err := uploadfile.CreateFile(1)
	if err != nil {
		panic(err)
	}

	ufile := &uploadfile.UploadFile{
		Client:       client,
		ProductSN:    productSN,
		DeviceSN:     deviceSN,
		DeviceSecret: deviceSecret,
		FilePath:     file_path,
		ContentType:  "application/octet-stream",
		Region:       region,
	}

	err = ufile.Uploadfile()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(500 * time.Millisecond)
}
