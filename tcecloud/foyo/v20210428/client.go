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

package v20210428

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-04-28"

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

func NewQuerySwitchRequest() (request *QuerySwitchRequest) {
	request = &QuerySwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("foyo", APIVersion, "QuerySwitch")
	return
}

func NewQuerySwitchResponse() (response *QuerySwitchResponse) {
	response = &QuerySwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询开关接口
func (c *Client) QuerySwitch(request *QuerySwitchRequest) (response *QuerySwitchResponse, err error) {
	if request == nil {
		request = NewQuerySwitchRequest()
	}
	response = NewQuerySwitchResponse()
	err = c.Send(request, response)
	return
}

func NewAction1Request() (request *Action1Request) {
	request = &Action1Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("foyo", APIVersion, "Action1")
	return
}

func NewAction1Response() (response *Action1Response) {
	response = &Action1Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 时擦出
func (c *Client) Action1(request *Action1Request) (response *Action1Response, err error) {
	if request == nil {
		request = NewAction1Request()
	}
	response = NewAction1Response()
	err = c.Send(request, response)
	return
}
