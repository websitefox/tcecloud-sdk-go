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

package v20201207

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-12-07"

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

func NewUpdateDefineSmsChannelRequest() (request *UpdateDefineSmsChannelRequest) {
	request = &UpdateDefineSmsChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "UpdateDefineSmsChannel")
	return
}

func NewUpdateDefineSmsChannelResponse() (response *UpdateDefineSmsChannelResponse) {
	response = &UpdateDefineSmsChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置产品自定义短信通道号
func (c *Client) UpdateDefineSmsChannel(request *UpdateDefineSmsChannelRequest) (response *UpdateDefineSmsChannelResponse, err error) {
	if request == nil {
		request = NewUpdateDefineSmsChannelRequest()
	}
	response = NewUpdateDefineSmsChannelResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDefineSmsChannelRequest() (request *QueryDefineSmsChannelRequest) {
	request = &QueryDefineSmsChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "QueryDefineSmsChannel")
	return
}

func NewQueryDefineSmsChannelResponse() (response *QueryDefineSmsChannelResponse) {
	response = &QueryDefineSmsChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表产品自定义通道号
func (c *Client) QueryDefineSmsChannel(request *QueryDefineSmsChannelRequest) (response *QueryDefineSmsChannelResponse, err error) {
	if request == nil {
		request = NewQueryDefineSmsChannelRequest()
	}
	response = NewQueryDefineSmsChannelResponse()
	err = c.Send(request, response)
	return
}
