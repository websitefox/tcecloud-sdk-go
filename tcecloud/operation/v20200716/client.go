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

package v20200716

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-07-16"

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

func NewGetHostAreaRequest() (request *GetHostAreaRequest) {
	request = &GetHostAreaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "GetHostArea")
	return
}

func NewGetHostAreaResponse() (response *GetHostAreaResponse) {
	response = &GetHostAreaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询机房地域
func (c *Client) GetHostArea(request *GetHostAreaRequest) (response *GetHostAreaResponse, err error) {
	if request == nil {
		request = NewGetHostAreaRequest()
	}
	response = NewGetHostAreaResponse()
	err = c.Send(request, response)
	return
}

func NewGetAnnounceClassRequest() (request *GetAnnounceClassRequest) {
	request = &GetAnnounceClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "GetAnnounceClass")
	return
}

func NewGetAnnounceClassResponse() (response *GetAnnounceClassResponse) {
	response = &GetAnnounceClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询公告分类信息
func (c *Client) GetAnnounceClass(request *GetAnnounceClassRequest) (response *GetAnnounceClassResponse, err error) {
	if request == nil {
		request = NewGetAnnounceClassRequest()
	}
	response = NewGetAnnounceClassResponse()
	err = c.Send(request, response)
	return
}

func NewGetAnnounceContentRequest() (request *GetAnnounceContentRequest) {
	request = &GetAnnounceContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "GetAnnounceContent")
	return
}

func NewGetAnnounceContentResponse() (response *GetAnnounceContentResponse) {
	response = &GetAnnounceContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询公告内容
func (c *Client) GetAnnounceContent(request *GetAnnounceContentRequest) (response *GetAnnounceContentResponse, err error) {
	if request == nil {
		request = NewGetAnnounceContentRequest()
	}
	response = NewGetAnnounceContentResponse()
	err = c.Send(request, response)
	return
}

func NewGetAnnounceListRequest() (request *GetAnnounceListRequest) {
	request = &GetAnnounceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "GetAnnounceList")
	return
}

func NewGetAnnounceListResponse() (response *GetAnnounceListResponse) {
	response = &GetAnnounceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询公告列表
func (c *Client) GetAnnounceList(request *GetAnnounceListRequest) (response *GetAnnounceListResponse, err error) {
	if request == nil {
		request = NewGetAnnounceListRequest()
	}
	response = NewGetAnnounceListResponse()
	err = c.Send(request, response)
	return
}
