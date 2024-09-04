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

package v20180813

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-08-13"

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

func NewGetConsoleSessionTokenRequest() (request *GetConsoleSessionTokenRequest) {
	request = &GetConsoleSessionTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sts", APIVersion, "GetConsoleSessionToken")
	return
}

func NewGetConsoleSessionTokenResponse() (response *GetConsoleSessionTokenResponse) {
	response = &GetConsoleSessionTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取控制台临时证书
func (c *Client) GetConsoleSessionToken(request *GetConsoleSessionTokenRequest) (response *GetConsoleSessionTokenResponse, err error) {
	if request == nil {
		request = NewGetConsoleSessionTokenRequest()
	}
	response = NewGetConsoleSessionTokenResponse()
	err = c.Send(request, response)
	return
}

func NewGetFederationTokenRequest() (request *GetFederationTokenRequest) {
	request = &GetFederationTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sts", APIVersion, "GetFederationToken")
	return
}

func NewGetFederationTokenResponse() (response *GetFederationTokenResponse) {
	response = &GetFederationTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取联合身份临时访问凭证
func (c *Client) GetFederationToken(request *GetFederationTokenRequest) (response *GetFederationTokenResponse, err error) {
	if request == nil {
		request = NewGetFederationTokenRequest()
	}
	response = NewGetFederationTokenResponse()
	err = c.Send(request, response)
	return
}

func NewAssumeRoleRequest() (request *AssumeRoleRequest) {
	request = &AssumeRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sts", APIVersion, "AssumeRole")
	return
}

func NewAssumeRoleResponse() (response *AssumeRoleResponse) {
	response = &AssumeRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 申请扮演角色
func (c *Client) AssumeRole(request *AssumeRoleRequest) (response *AssumeRoleResponse, err error) {
	if request == nil {
		request = NewAssumeRoleRequest()
	}
	response = NewAssumeRoleResponse()
	err = c.Send(request, response)
	return
}

func NewGetThirdPartyFederationTokenRequest() (request *GetThirdPartyFederationTokenRequest) {
	request = &GetThirdPartyFederationTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sts", APIVersion, "GetThirdPartyFederationToken")
	return
}

func NewGetThirdPartyFederationTokenResponse() (response *GetThirdPartyFederationTokenResponse) {
	response = &GetThirdPartyFederationTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户临时证书（第三方开发商）
func (c *Client) GetThirdPartyFederationToken(request *GetThirdPartyFederationTokenRequest) (response *GetThirdPartyFederationTokenResponse, err error) {
	if request == nil {
		request = NewGetThirdPartyFederationTokenRequest()
	}
	response = NewGetThirdPartyFederationTokenResponse()
	err = c.Send(request, response)
	return
}
