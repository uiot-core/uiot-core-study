package uiot

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)


// View Shadow Request
type ViewShadowRequest struct {
	request.CommonBase


	ProductSN *string `required:"true"`
	DeviceSN *string `required:"ture"`

}

type ShadowState struct {
	Reported map[string]interface{}
	Desired map[string]interface{}
}

type ShadowMetadata struct {
	Reported map[string]map[string]int64
	Desired map[string]map[string]int64
}

type ShadowDocument struct {
	State ShadowState
	Metadata ShadowMetadata
}

// View Shadow Response
type ViewShadowResponse struct {
	response.CommonBase

	Payload ShadowDocument
	Version uint32
	Timestamp int64

}

func (c *UIotClient) ViewShadowRequest() *ViewShadowRequest {
	req := &ViewShadowRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

func (c *UIotClient) ViewShadow(req *ViewShadowRequest) (*ViewShadowResponse, error) {
	var err error
	var res ViewShadowResponse
	var reqImmutable = *req
	//reqImmutable.Password = ucloud.String(base64.StdEncoding.EncodeToString([]byte(ucloud.StringValue(req.Password))))

	err = c.Client.InvokeAction("GetUIoTCoreDeviceShadow", &reqImmutable, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}


type UrlBody struct {
	ProductSN    string
	DeviceSN     string
	DeviceSecret string
	FileName     string
	Filesize     int64
	MD5          string
	ContentType  string `json:"Content-Type"`
}

type UrlRet struct {
	RetCode       int
	Authorization string
	URL           string
}


