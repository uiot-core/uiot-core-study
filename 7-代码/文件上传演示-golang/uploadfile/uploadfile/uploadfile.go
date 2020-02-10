package uploadfile

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"../utils"
)


type UploadFile struct {
	Client *http.Client
	ProductSN string
	DeviceSN string
	DeviceSecret string
	FilePath string
	ContentType  string
	Region string
}

type UrlBody struct {
	ProductSN    string
	DeviceSN     string
	FileName     string
	Filesize     int64
	Md5          string
	ContentType  string `json:"Content-Type"`
}

type UrlRet struct {
	RetCode       int
	Authorization string
	URL           string
}

type PutReq struct {
	URL           string
	Authorization string
	ContentType   string `json:"Content-Type"`
	ContentMD5    string `json:"Content-MD5"`
}


func CreateFile(sizem float64) (string, error) {
	var sizek int
	sizek = int(sizem * 1024)
	// check if the dir exit
	_, err := os.Stat("./files")
	if os.IsNotExist(err) {
		err = os.Mkdir("./files", os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	// create the filepath
	file_path := "./files/" + utils.GetRandomString(16) + strconv.FormatFloat(sizem, 'f', -1, 64) + "M.file"

	f, err := os.OpenFile(file_path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 0; i < sizek; i++ {
		_, err := f.WriteAt([]byte(strings.Repeat(utils.GetRandomString(1024), 1)), int64(i)*1024)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}
	return file_path, nil
}

func (u *UploadFile) Uploadfile() error {

	// open file
	fileupload, err := os.Open(u.FilePath)
	if err != nil {
		return  err
	}

	defer fileupload.Close()

	fileinfo, err := fileupload.Stat()
	filesize := fileinfo.Size()

	// caculate the md5
	h := md5.New()
	if _, err := io.Copy(h, fileupload); err != nil {
		return err
	}

	md5string := hex.EncodeToString(h.Sum(nil))

	fmt.Println("md5sum:"+md5string)

	// get the upload url
	urlbody := UrlBody{
		ProductSN:    u.ProductSN,
		DeviceSN:     u.DeviceSN,
		FileName:     filepath.Base(u.FilePath),
		Filesize:     filesize,
		Md5:          md5string,
		ContentType:  u.ContentType}

	bodydata, err := json.Marshal(urlbody)

	// Authorization
	mac := hmac.New(sha256.New, []byte(u.DeviceSecret))
	mac.Write(bodydata)
	authString := hex.EncodeToString(mac.Sum(nil))

	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", "https://file-"+u.Region+".iot.ucloud.cn/api/v1/url", bytes.NewReader(bodydata))

	req.Header.Add("Content-Type", u.ContentType)
	req.Header.Add("Authorization", authString)

	resp, err := u.Client.Do(req)

	if err != nil {
		return  err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return  err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	// put file to ufile
	urlret := UrlRet{}
	if err := json.Unmarshal(body, &urlret); err != nil {
		return  err
	}

	if urlret.RetCode != 0 {
		return errors.New(string(body))
	}

	_, err = fileupload.Seek(0, 0)
	if err != nil {
		return err
	}

	putreq := &PutReq{
		URL:           urlret.URL,
		Authorization: urlret.Authorization,
		ContentType:   u.ContentType,
		ContentMD5:    md5string}

	req, _ = http.NewRequest("PUT", putreq.URL, fileupload)

	req.Header.Add("Content-Type", putreq.ContentType)
	req.Header.Add("Authorization", putreq.Authorization)
	req.Header.Add("Content-MD5", putreq.ContentMD5)

	resp, err = u.Client.Do(req)
	if err != nil {
		return err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return  err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(string(body))
		return errors.New(string(body))
	}

	return nil
}


