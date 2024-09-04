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

package v20190304

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-03-04"

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

func NewCreatExporJobRequest() (request *CreatExporJobRequest) {
	request = &CreatExporJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "CreatExporJob")
	return
}

func NewCreatExporJobResponse() (response *CreatExporJobResponse) {
	response = &CreatExporJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建导出任务
func (c *Client) CreatExporJob(request *CreatExporJobRequest) (response *CreatExporJobResponse, err error) {
	if request == nil {
		request = NewCreatExporJobRequest()
	}
	response = NewCreatExporJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBucketRequest() (request *CreateBucketRequest) {
	request = &CreateBucketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "CreateBucket")
	return
}

func NewCreateBucketResponse() (response *CreateBucketResponse) {
	response = &CreateBucketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建存储桶
func (c *Client) CreateBucket(request *CreateBucketRequest) (response *CreateBucketResponse, err error) {
	if request == nil {
		request = NewCreateBucketRequest()
	}
	response = NewCreateBucketResponse()
	err = c.Send(request, response)
	return
}

func NewGetExporLoggingListRequest() (request *GetExporLoggingListRequest) {
	request = &GetExporLoggingListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "GetExporLoggingList")
	return
}

func NewGetExporLoggingListResponse() (response *GetExporLoggingListResponse) {
	response = &GetExporLoggingListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务列表
func (c *Client) GetExporLoggingList(request *GetExporLoggingListRequest) (response *GetExporLoggingListResponse, err error) {
	if request == nil {
		request = NewGetExporLoggingListRequest()
	}
	response = NewGetExporLoggingListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteExporJobRequest() (request *DeleteExporJobRequest) {
	request = &DeleteExporJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "DeleteExporJob")
	return
}

func NewDeleteExporJobResponse() (response *DeleteExporJobResponse) {
	response = &DeleteExporJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除任务
func (c *Client) DeleteExporJob(request *DeleteExporJobRequest) (response *DeleteExporJobResponse, err error) {
	if request == nil {
		request = NewDeleteExporJobRequest()
	}
	response = NewDeleteExporJobResponse()
	err = c.Send(request, response)
	return
}

func NewLookupEventsRequest() (request *LookupEventsRequest) {
	request = &LookupEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "LookupEvents")
	return
}

func NewLookupEventsResponse() (response *LookupEventsResponse) {
	response = &LookupEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检索日志接口
func (c *Client) LookupEvents(request *LookupEventsRequest) (response *LookupEventsResponse, err error) {
	if request == nil {
		request = NewLookupEventsRequest()
	}
	response = NewLookupEventsResponse()
	err = c.Send(request, response)
	return
}
