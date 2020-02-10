package uiot

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)



//  SetUIoTCoreDeviceProperty
type SetUIoTCoreDevicePropertyRequest struct {
	request.CommonBase
	ProductSN *string `required:"true"`
	DeviceSN *string `required:"ture"`
	Property *string `required:"ture"`
	Desired bool
}
type SetUIoTCoreDevicePropertyResponse struct {
	response.CommonBase
	Payload  interface{}
}

func (c *UIotClient) SetUIoTCoreDeviceProperty(req *SetUIoTCoreDevicePropertyRequest) (*SetUIoTCoreDevicePropertyResponse, error){
	c.Client.SetupRequest(req)
	req.SetRetryable(false)
	resp:= SetUIoTCoreDevicePropertyResponse{}
	err := c.Client.InvokeAction("SetUIoTCoreDeviceProperty", req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}

// PublishUIoTCoreMQTTMessage
type PublishUIoTCoreMQTTMessageRequest struct {
	request.CommonBase
	ProductSN *string `required:"true"`
	DeviceSN *string `required:"ture"`
	TopicFullName *string
	MessageContent *string
	Qos int
}
type PublishUIoTCoreMQTTMessageResponse struct {
	response.CommonBase
	Payload  interface{}
}

func (c *UIotClient) PublishUIoTCoreMQTTMessage(req *PublishUIoTCoreMQTTMessageRequest) (*PublishUIoTCoreMQTTMessageResponse, error){
	c.Client.SetupRequest(req)
	req.SetRetryable(false)
	resp:= PublishUIoTCoreMQTTMessageResponse{}
	err := c.Client.InvokeAction("PublishUIoTCoreMQTTMessage", req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}

// UpdateUIoTCoreDeviceShadow
type UpdateUIoTCoreDeviceShadowRequest struct {
	request.CommonBase
	ProductSN *string `required:"true"`
	DeviceSN *string `required:"ture"`
	Desired *string
	ShadowVersion *string
}

type UpdateUIoTCoreDeviceShadowResponse struct {
	response.CommonBase
	Payload  interface{}
	Version interface{}
	Timestamp interface{}
}

func (c *UIotClient) UpdateUIoTCoreDeviceShadow(req *UpdateUIoTCoreDeviceShadowRequest) (*UpdateUIoTCoreDeviceShadowResponse, error){
	c.Client.SetupRequest(req)
	req.SetRetryable(false)
	resp:= UpdateUIoTCoreDeviceShadowResponse{}
	err := c.Client.InvokeAction("UpdateUIoTCoreDeviceShadow", req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}


// SendUIoTCoreDeviceCommand
type SendUIoTCoreDeviceCommandRequest struct {
	request.CommonBase
	ProductSN *string `required:"true"`
	DeviceSN *string `required:"ture"`
	Identifier *string
	Input *string
	Method *string
	Timeout *int
}

type SendUIoTCoreDeviceCommandResponse struct {
	response.CommonBase
	Payload  interface{}
	RequestID interface{}
	Output interface{}
}

func (c *UIotClient) SendUIoTCoreDeviceCommand(req *SendUIoTCoreDeviceCommandRequest) (*SendUIoTCoreDeviceCommandResponse, error){
	c.Client.SetupRequest(req)
	req.SetRetryable(false)
	resp:= SendUIoTCoreDeviceCommandResponse{}
	err := c.Client.InvokeAction("SendUIoTCoreDeviceCommand", req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}

// CreateUIoTCoreDevice
type CreateUIoTCoreDeviceRequest struct {
	request.CommonBase
	ProductSN *string `required:"true"`
	DeviceSN *string `required:"ture"`
}

type CreateUIoTCoreDeviceResponse struct {
	response.CommonBase
	Payload  interface{}
}

func (c *UIotClient) CreateUIoTCoreDevice(req *CreateUIoTCoreDeviceRequest) (*CreateUIoTCoreDeviceResponse, error){
	c.Client.SetupRequest(req)
	req.SetRetryable(false)
	resp:= CreateUIoTCoreDeviceResponse{}
	err := c.Client.InvokeAction("CreateUIoTCoreDevice", req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}

// BatchCreateUIoTCoreDevice
type BatchCreateUIoTCoreDeviceRequest struct {
	request.CommonBase
	ProductSN *string `required:"true"`
	DeviceSN *string `required:"ture"`
	DeviceCount int
}

type BatchCreateUIoTCoreDeviceResponse struct {
	response.CommonBase
	Payload  interface{}
}

func (c *UIotClient) BatchCreateUIoTCoreDevice(req *BatchCreateUIoTCoreDeviceRequest) (*BatchCreateUIoTCoreDeviceResponse, error){
	c.Client.SetupRequest(req)
	req.SetRetryable(false)
	resp:= BatchCreateUIoTCoreDeviceResponse{}
	err := c.Client.InvokeAction("BatchCreateUIoTCoreDevice", req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}


// BatchCreateUIoTCoreDeviceWithSN
type BatchCreateUIoTCoreDeviceWithSNRequest struct {
	request.CommonBase
	ProductSN *string `required:"true"`
	DeviceSN []string
}

type BatchCreateUIoTCoreDeviceWithSNResponse struct {
	response.CommonBase
	Payload  interface{}
}

func (c *UIotClient) BatchCreateUIoTCoreDeviceWithSN(req *BatchCreateUIoTCoreDeviceWithSNRequest) (*BatchCreateUIoTCoreDeviceWithSNResponse, error){
	c.Client.SetupRequest(req)
	req.SetRetryable(false)
	resp:= BatchCreateUIoTCoreDeviceWithSNResponse{}
	err := c.Client.InvokeAction("BatchCreateUIoTCoreDeviceWithSN", req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}

// QueryUIoTCoreDeviceFileList
type FileInfo struct {
	FileName string
	FileSize int
	CreateTime int
	URL string
}

type QueryUIoTCoreDeviceFileListRequest struct {
	request.CommonBase
	ProductSN *string `required:"true"`
	DeviceSN *string
	FileName *string
	Offset int
	Limit int
	URLExpire int
}

type QueryUIoTCoreDeviceFileListResponse struct {
	response.CommonBase
	FileList []FileInfo
	TotalCount int
}

func (c *UIotClient) QueryUIoTCoreDeviceFileList(req *QueryUIoTCoreDeviceFileListRequest) (*QueryUIoTCoreDeviceFileListResponse, error){
	c.Client.SetupRequest(req)
	req.SetRetryable(false)
	resp:= QueryUIoTCoreDeviceFileListResponse{}
	err := c.Client.InvokeAction("QueryUIoTCoreDeviceFileList", req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}


// GetUIoTCoreDeviceFileURL
type GetUIoTCoreDeviceFileURLRequest struct {
	request.CommonBase
	ProductSN *string `required:"true"`
	DeviceSN *string
	FileName *string
	URLExpire int
}

type GetUIoTCoreDeviceFileURLResponse struct {
	response.CommonBase
	URL string
}

func (c *UIotClient) GetUIoTCoreDeviceFileURL(req *GetUIoTCoreDeviceFileURLRequest) (*GetUIoTCoreDeviceFileURLResponse, error){
	c.Client.SetupRequest(req)
	req.SetRetryable(false)
	resp:= GetUIoTCoreDeviceFileURLResponse{}
	err := c.Client.InvokeAction("GetUIoTCoreDeviceFileURL", req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}

// GetUIoTCoreDeviceList
type GetUIoTCoreDeviceListRequest struct {
	request.CommonBase
	ProductSN *string `required:"true"`
	DeviceSN *string
	FileName *string
	URLExpire int
}

type GetUIoTCoreDeviceListResponse struct {
	response.CommonBase
	URL string
}

func (c *UIotClient) GetUIoTCoreDeviceList(req *GetUIoTCoreDeviceListRequest) (*GetUIoTCoreDeviceListResponse, error){
	c.Client.SetupRequest(req)
	req.SetRetryable(false)
	resp:= GetUIoTCoreDeviceListResponse{}
	err := c.Client.InvokeAction("GetUIoTCoreDeviceList", req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}

