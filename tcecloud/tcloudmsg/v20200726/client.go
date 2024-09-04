// All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20200726

import (
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-07-26"

var _ = tchttp.POST

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewDeleteSiteMsgRequest() (request *DeleteSiteMsgRequest) {
	request = &DeleteSiteMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "DeleteSiteMsg")
	return
}

func NewDeleteSiteMsgResponse() (response *DeleteSiteMsgResponse) {
	response = &DeleteSiteMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除站内信
func (c *Client) DeleteSiteMsg(request *DeleteSiteMsgRequest) (response *DeleteSiteMsgResponse, err error) {
	if request == nil {
		request = NewDeleteSiteMsgRequest()
	}
	response = NewDeleteSiteMsgResponse()
	err = c.Send(request, response)
	return
}

func NewSetAllSiteMsgReadRequest() (request *SetAllSiteMsgReadRequest) {
	request = &SetAllSiteMsgReadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "SetAllSiteMsgRead")
	return
}

func NewSetAllSiteMsgReadResponse() (response *SetAllSiteMsgReadResponse) {
	response = &SetAllSiteMsgReadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标记全部站内信已读
func (c *Client) SetAllSiteMsgRead(request *SetAllSiteMsgReadRequest) (response *SetAllSiteMsgReadResponse, err error) {
	if request == nil {
		request = NewSetAllSiteMsgReadRequest()
	}
	response = NewSetAllSiteMsgReadResponse()
	err = c.Send(request, response)
	return
}

func NewGetReceiversOnAllTypeRequest() (request *GetReceiversOnAllTypeRequest) {
	request = &GetReceiversOnAllTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "GetReceiversOnAllType")
	return
}

func NewGetReceiversOnAllTypeResponse() (response *GetReceiversOnAllTypeResponse) {
	response = &GetReceiversOnAllTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有的消息订阅
func (c *Client) GetReceiversOnAllType(request *GetReceiversOnAllTypeRequest) (response *GetReceiversOnAllTypeResponse, err error) {
	if request == nil {
		request = NewGetReceiversOnAllTypeRequest()
	}
	response = NewGetReceiversOnAllTypeResponse()
	err = c.Send(request, response)
	return
}

func NewGetSiteMsgDetailRequest() (request *GetSiteMsgDetailRequest) {
	request = &GetSiteMsgDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "GetSiteMsgDetail")
	return
}

func NewGetSiteMsgDetailResponse() (response *GetSiteMsgDetailResponse) {
	response = &GetSiteMsgDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取站内信详情
func (c *Client) GetSiteMsgDetail(request *GetSiteMsgDetailRequest) (response *GetSiteMsgDetailResponse, err error) {
	if request == nil {
		request = NewGetSiteMsgDetailRequest()
	}
	response = NewGetSiteMsgDetailResponse()
	err = c.Send(request, response)
	return
}

func NewListFirstMsgTypeRequest() (request *ListFirstMsgTypeRequest) {
	request = &ListFirstMsgTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "ListFirstMsgType")
	return
}

func NewListFirstMsgTypeResponse() (response *ListFirstMsgTypeResponse) {
	response = &ListFirstMsgTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有的一级消息分类
func (c *Client) ListFirstMsgType(request *ListFirstMsgTypeRequest) (response *ListFirstMsgTypeResponse, err error) {
	if request == nil {
		request = NewListFirstMsgTypeRequest()
	}
	response = NewListFirstMsgTypeResponse()
	err = c.Send(request, response)
	return
}

func NewGetMsgSummaryRequest() (request *GetMsgSummaryRequest) {
	request = &GetMsgSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "GetMsgSummary")
	return
}

func NewGetMsgSummaryResponse() (response *GetMsgSummaryResponse) {
	response = &GetMsgSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 站内信统计汇总
func (c *Client) GetMsgSummary(request *GetMsgSummaryRequest) (response *GetMsgSummaryResponse, err error) {
	if request == nil {
		request = NewGetMsgSummaryRequest()
	}
	response = NewGetMsgSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewSetSiteMsgReadRequest() (request *SetSiteMsgReadRequest) {
	request = &SetSiteMsgReadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "SetSiteMsgRead")
	return
}

func NewSetSiteMsgReadResponse() (response *SetSiteMsgReadResponse) {
	response = &SetSiteMsgReadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量设置站内信已读
func (c *Client) SetSiteMsgRead(request *SetSiteMsgReadRequest) (response *SetSiteMsgReadResponse, err error) {
	if request == nil {
		request = NewSetSiteMsgReadRequest()
	}
	response = NewSetSiteMsgReadResponse()
	err = c.Send(request, response)
	return
}

func NewGetMsgListRequest() (request *GetMsgListRequest) {
	request = &GetMsgListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "GetMsgList")
	return
}

func NewGetMsgListResponse() (response *GetMsgListResponse) {
	response = &GetMsgListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看未读的站内信
func (c *Client) GetMsgList(request *GetMsgListRequest) (response *GetMsgListResponse, err error) {
	if request == nil {
		request = NewGetMsgListRequest()
	}
	response = NewGetMsgListResponse()
	err = c.Send(request, response)
	return
}

func NewAddReceiverOnMsgTypeRequest() (request *AddReceiverOnMsgTypeRequest) {
	request = &AddReceiverOnMsgTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "AddReceiverOnMsgType")
	return
}

func NewAddReceiverOnMsgTypeResponse() (response *AddReceiverOnMsgTypeResponse) {
	response = &AddReceiverOnMsgTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量消息订阅增加接收人
func (c *Client) AddReceiverOnMsgType(request *AddReceiverOnMsgTypeRequest) (response *AddReceiverOnMsgTypeResponse, err error) {
	if request == nil {
		request = NewAddReceiverOnMsgTypeRequest()
	}
	response = NewAddReceiverOnMsgTypeResponse()
	err = c.Send(request, response)
	return
}

func NewListSecondaryMsgTypeRequest() (request *ListSecondaryMsgTypeRequest) {
	request = &ListSecondaryMsgTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "ListSecondaryMsgType")
	return
}

func NewListSecondaryMsgTypeResponse() (response *ListSecondaryMsgTypeResponse) {
	response = &ListSecondaryMsgTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取二级消息类型
func (c *Client) ListSecondaryMsgType(request *ListSecondaryMsgTypeRequest) (response *ListSecondaryMsgTypeResponse, err error) {
	if request == nil {
		request = NewListSecondaryMsgTypeRequest()
	}
	response = NewListSecondaryMsgTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReceiverOnMsgTypeRequest() (request *ModifyReceiverOnMsgTypeRequest) {
	request = &ModifyReceiverOnMsgTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "ModifyReceiverOnMsgType")
	return
}

func NewModifyReceiverOnMsgTypeResponse() (response *ModifyReceiverOnMsgTypeResponse) {
	response = &ModifyReceiverOnMsgTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消息订阅修改订阅类型和接收人
func (c *Client) ModifyReceiverOnMsgType(request *ModifyReceiverOnMsgTypeRequest) (response *ModifyReceiverOnMsgTypeResponse, err error) {
	if request == nil {
		request = NewModifyReceiverOnMsgTypeRequest()
	}
	response = NewModifyReceiverOnMsgTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDelReceiverOnMsgTypeRequest() (request *DelReceiverOnMsgTypeRequest) {
	request = &DelReceiverOnMsgTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "DelReceiverOnMsgType")
	return
}

func NewDelReceiverOnMsgTypeResponse() (response *DelReceiverOnMsgTypeResponse) {
	response = &DelReceiverOnMsgTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量消息订阅删除接收人
func (c *Client) DelReceiverOnMsgType(request *DelReceiverOnMsgTypeRequest) (response *DelReceiverOnMsgTypeResponse, err error) {
	if request == nil {
		request = NewDelReceiverOnMsgTypeRequest()
	}
	response = NewDelReceiverOnMsgTypeResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSiteMsgRequest() (request *SearchSiteMsgRequest) {
	request = &SearchSiteMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "SearchSiteMsg")
	return
}

func NewSearchSiteMsgResponse() (response *SearchSiteMsgResponse) {
	response = &SearchSiteMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表过滤查询站内信
func (c *Client) SearchSiteMsg(request *SearchSiteMsgRequest) (response *SearchSiteMsgResponse, err error) {
	if request == nil {
		request = NewSearchSiteMsgRequest()
	}
	response = NewSearchSiteMsgResponse()
	err = c.Send(request, response)
	return
}
