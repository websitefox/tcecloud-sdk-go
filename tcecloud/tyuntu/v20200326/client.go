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

package v20200326

import (
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-03-26"

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

func NewRunActionRequest() (request *RunActionRequest) {
	request = &RunActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tyuntu", APIVersion, "RunAction")
	return
}

func NewRunActionResponse() (response *RunActionResponse) {
	response = &RunActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api explorer 入口
func (c *Client) RunAction(request *RunActionRequest) (response *RunActionResponse, err error) {
	if request == nil {
		request = NewRunActionRequest()
	}
	response = NewRunActionResponse()
	err = c.Send(request, response)
	return
}

func NewGetRegionsRequest() (request *GetRegionsRequest) {
	request = &GetRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tyuntu", APIVersion, "GetRegions")
	return
}

func NewGetRegionsResponse() (response *GetRegionsResponse) {
	response = &GetRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetRegions
func (c *Client) GetRegions(request *GetRegionsRequest) (response *GetRegionsResponse, err error) {
	if request == nil {
		request = NewGetRegionsRequest()
	}
	response = NewGetRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeActionRequest() (request *DescribeActionRequest) {
	request = &DescribeActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tyuntu", APIVersion, "DescribeAction")
	return
}

func NewDescribeActionResponse() (response *DescribeActionResponse) {
	response = &DescribeActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询接口信息
func (c *Client) DescribeAction(request *DescribeActionRequest) (response *DescribeActionResponse, err error) {
	if request == nil {
		request = NewDescribeActionRequest()
	}
	response = NewDescribeActionResponse()
	err = c.Send(request, response)
	return
}

func NewQueryModuleFilterRequest() (request *QueryModuleFilterRequest) {
	request = &QueryModuleFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tyuntu", APIVersion, "QueryModuleFilter")
	return
}

func NewQueryModuleFilterResponse() (response *QueryModuleFilterResponse) {
	response = &QueryModuleFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页、搜索、列表模块
func (c *Client) QueryModuleFilter(request *QueryModuleFilterRequest) (response *QueryModuleFilterResponse, err error) {
	if request == nil {
		request = NewQueryModuleFilterRequest()
	}
	response = NewQueryModuleFilterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSdkDemosRequest() (request *DescribeSdkDemosRequest) {
	request = &DescribeSdkDemosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tyuntu", APIVersion, "DescribeSdkDemos")
	return
}

func NewDescribeSdkDemosResponse() (response *DescribeSdkDemosResponse) {
	response = &DescribeSdkDemosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Action的代码模板
func (c *Client) DescribeSdkDemos(request *DescribeSdkDemosRequest) (response *DescribeSdkDemosResponse, err error) {
	if request == nil {
		request = NewDescribeSdkDemosRequest()
	}
	response = NewDescribeSdkDemosResponse()
	err = c.Send(request, response)
	return
}

func NewGetApiInfoRequest() (request *GetApiInfoRequest) {
	request = &GetApiInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tyuntu", APIVersion, "GetApiInfo")
	return
}

func NewGetApiInfoResponse() (response *GetApiInfoResponse) {
	response = &GetApiInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取API详情
func (c *Client) GetApiInfo(request *GetApiInfoRequest) (response *GetApiInfoResponse, err error) {
	if request == nil {
		request = NewGetApiInfoRequest()
	}
	response = NewGetApiInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetDataTypesRequest() (request *GetDataTypesRequest) {
	request = &GetDataTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tyuntu", APIVersion, "GetDataTypes")
	return
}

func NewGetDataTypesResponse() (response *GetDataTypesResponse) {
	response = &GetDataTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询数据类型
func (c *Client) GetDataTypes(request *GetDataTypesRequest) (response *GetDataTypesResponse, err error) {
	if request == nil {
		request = NewGetDataTypesRequest()
	}
	response = NewGetDataTypesResponse()
	err = c.Send(request, response)
	return
}

func NewQueryApiInfoFilterRequest() (request *QueryApiInfoFilterRequest) {
	request = &QueryApiInfoFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tyuntu", APIVersion, "QueryApiInfoFilter")
	return
}

func NewQueryApiInfoFilterResponse() (response *QueryApiInfoFilterResponse) {
	response = &QueryApiInfoFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页、搜索列表API接口
func (c *Client) QueryApiInfoFilter(request *QueryApiInfoFilterRequest) (response *QueryApiInfoFilterResponse, err error) {
	if request == nil {
		request = NewQueryApiInfoFilterRequest()
	}
	response = NewQueryApiInfoFilterResponse()
	err = c.Send(request, response)
	return
}
