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

package v20210801

import (
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-08-01"

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

func NewQueryDataSouceDataRequest() (request *QueryDataSouceDataRequest) {
	request = &QueryDataSouceDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryDataSouceData")
	return
}

func NewQueryDataSouceDataResponse() (response *QueryDataSouceDataResponse) {
	response = &QueryDataSouceDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取独立数据源接口数据
func (c *Client) QueryDataSouceData(request *QueryDataSouceDataRequest) (response *QueryDataSouceDataResponse, err error) {
	if request == nil {
		request = NewQueryDataSouceDataRequest()
	}
	response = NewQueryDataSouceDataResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDataSourceRequest() (request *QueryDataSourceRequest) {
	request = &QueryDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryDataSource")
	return
}

func NewQueryDataSourceResponse() (response *QueryDataSourceResponse) {
	response = &QueryDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询数据源
func (c *Client) QueryDataSource(request *QueryDataSourceRequest) (response *QueryDataSourceResponse, err error) {
	if request == nil {
		request = NewQueryDataSourceRequest()
	}
	response = NewQueryDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPagesRequest() (request *QueryPagesRequest) {
	request = &QueryPagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryPages")
	return
}

func NewQueryPagesResponse() (response *QueryPagesResponse) {
	response = &QueryPagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取页面列表
func (c *Client) QueryPages(request *QueryPagesRequest) (response *QueryPagesResponse, err error) {
	if request == nil {
		request = NewQueryPagesRequest()
	}
	response = NewQueryPagesResponse()
	err = c.Send(request, response)
	return
}
