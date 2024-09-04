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

package v20190719

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-07-19"

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

func NewDescribeMountTargetsWithRegionRequest() (request *DescribeMountTargetsWithRegionRequest) {
	request = &DescribeMountTargetsWithRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeMountTargetsWithRegion")
	return
}

func NewDescribeMountTargetsWithRegionResponse() (response *DescribeMountTargetsWithRegionResponse) {
	response = &DescribeMountTargetsWithRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询区域挂载点情况，只用于控制台。
func (c *Client) DescribeMountTargetsWithRegion(request *DescribeMountTargetsWithRegionRequest) (response *DescribeMountTargetsWithRegionResponse, err error) {
	if request == nil {
		request = NewDescribeMountTargetsWithRegionRequest()
	}
	response = NewDescribeMountTargetsWithRegionResponse()
	err = c.Send(request, response)
	return
}

func NewSignUpCfsServiceRequest() (request *SignUpCfsServiceRequest) {
	request = &SignUpCfsServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "SignUpCfsService")
	return
}

func NewSignUpCfsServiceResponse() (response *SignUpCfsServiceResponse) {
	response = &SignUpCfsServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SignUpCfsService）用于开通CFS服务。
func (c *Client) SignUpCfsService(request *SignUpCfsServiceRequest) (response *SignUpCfsServiceResponse, err error) {
	if request == nil {
		request = NewSignUpCfsServiceRequest()
	}
	response = NewSignUpCfsServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfsTagsRequest() (request *DescribeCfsTagsRequest) {
	request = &DescribeCfsTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeCfsTags")
	return
}

func NewDescribeCfsTagsResponse() (response *DescribeCfsTagsResponse) {
	response = &DescribeCfsTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cfs资源池列表
func (c *Client) DescribeCfsTags(request *DescribeCfsTagsRequest) (response *DescribeCfsTagsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsTagsRequest()
	}
	response = NewDescribeCfsTagsResponse()
	err = c.Send(request, response)
	return
}

func NewScaleUpFileSystemRequest() (request *ScaleUpFileSystemRequest) {
	request = &ScaleUpFileSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "ScaleUpFileSystem")
	return
}

func NewScaleUpFileSystemResponse() (response *ScaleUpFileSystemResponse) {
	response = &ScaleUpFileSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 扩容文件系统
func (c *Client) ScaleUpFileSystem(request *ScaleUpFileSystemRequest) (response *ScaleUpFileSystemResponse, err error) {
	if request == nil {
		request = NewScaleUpFileSystemRequest()
	}
	response = NewScaleUpFileSystemResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStatisticsRequest() (request *DescribeStatisticsRequest) {
	request = &DescribeStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeStatistics")
	return
}

func NewDescribeStatisticsResponse() (response *DescribeStatisticsResponse) {
	response = &DescribeStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 不对外，仅供控制台使用
func (c *Client) DescribeStatistics(request *DescribeStatisticsRequest) (response *DescribeStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeStatisticsRequest()
	}
	response = NewDescribeStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCfsFileSystemRequest() (request *DeleteCfsFileSystemRequest) {
	request = &DeleteCfsFileSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DeleteCfsFileSystem")
	return
}

func NewDeleteCfsFileSystemResponse() (response *DeleteCfsFileSystemResponse) {
	response = &DeleteCfsFileSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于删除文件系统
func (c *Client) DeleteCfsFileSystem(request *DeleteCfsFileSystemRequest) (response *DeleteCfsFileSystemResponse, err error) {
	if request == nil {
		request = NewDeleteCfsFileSystemRequest()
	}
	response = NewDeleteCfsFileSystemResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCfsFileSystemNameRequest() (request *UpdateCfsFileSystemNameRequest) {
	request = &UpdateCfsFileSystemNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "UpdateCfsFileSystemName")
	return
}

func NewUpdateCfsFileSystemNameResponse() (response *UpdateCfsFileSystemNameResponse) {
	response = &UpdateCfsFileSystemNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateCfsFileSystemName）用于更新文件系统名
func (c *Client) UpdateCfsFileSystemName(request *UpdateCfsFileSystemNameRequest) (response *UpdateCfsFileSystemNameResponse, err error) {
	if request == nil {
		request = NewUpdateCfsFileSystemNameRequest()
	}
	response = NewUpdateCfsFileSystemNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMountTargetsRequest() (request *DescribeMountTargetsRequest) {
	request = &DescribeMountTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeMountTargets")
	return
}

func NewDescribeMountTargetsResponse() (response *DescribeMountTargetsResponse) {
	response = &DescribeMountTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeMountTargets）用于查询文件系统挂载点信息
func (c *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) (response *DescribeMountTargetsResponse, err error) {
	if request == nil {
		request = NewDescribeMountTargetsRequest()
	}
	response = NewDescribeMountTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfsFileSystemsRequest() (request *DescribeCfsFileSystemsRequest) {
	request = &DescribeCfsFileSystemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeCfsFileSystems")
	return
}

func NewDescribeCfsFileSystemsResponse() (response *DescribeCfsFileSystemsResponse) {
	response = &DescribeCfsFileSystemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCfsFileSystems）用于查询文件系统
func (c *Client) DescribeCfsFileSystems(request *DescribeCfsFileSystemsRequest) (response *DescribeCfsFileSystemsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsFileSystemsRequest()
	}
	response = NewDescribeCfsFileSystemsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCfsFileSystemRequest() (request *CreateCfsFileSystemRequest) {
	request = &CreateCfsFileSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "CreateCfsFileSystem")
	return
}

func NewCreateCfsFileSystemResponse() (response *CreateCfsFileSystemResponse) {
	response = &CreateCfsFileSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于添加新文件系统
func (c *Client) CreateCfsFileSystem(request *CreateCfsFileSystemRequest) (response *CreateCfsFileSystemResponse, err error) {
	if request == nil {
		request = NewCreateCfsFileSystemRequest()
	}
	response = NewCreateCfsFileSystemResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfsServiceStatusRequest() (request *DescribeCfsServiceStatusRequest) {
	request = &DescribeCfsServiceStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeCfsServiceStatus")
	return
}

func NewDescribeCfsServiceStatusResponse() (response *DescribeCfsServiceStatusResponse) {
	response = &DescribeCfsServiceStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCfsServiceStatus）用于查询用户使用CFS的服务状态。
func (c *Client) DescribeCfsServiceStatus(request *DescribeCfsServiceStatusRequest) (response *DescribeCfsServiceStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCfsServiceStatusRequest()
	}
	response = NewDescribeCfsServiceStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfsFileSystemClientsRequest() (request *DescribeCfsFileSystemClientsRequest) {
	request = &DescribeCfsFileSystemClientsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeCfsFileSystemClients")
	return
}

func NewDescribeCfsFileSystemClientsResponse() (response *DescribeCfsFileSystemClientsResponse) {
	response = &DescribeCfsFileSystemClientsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询挂载该文件系统的客户端。此功能需要客户端安装CFS监控插件。
func (c *Client) DescribeCfsFileSystemClients(request *DescribeCfsFileSystemClientsRequest) (response *DescribeCfsFileSystemClientsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsFileSystemClientsRequest()
	}
	response = NewDescribeCfsFileSystemClientsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCfsFileSystemPGroupRequest() (request *UpdateCfsFileSystemPGroupRequest) {
	request = &UpdateCfsFileSystemPGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "UpdateCfsFileSystemPGroup")
	return
}

func NewUpdateCfsFileSystemPGroupResponse() (response *UpdateCfsFileSystemPGroupResponse) {
	response = &UpdateCfsFileSystemPGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateCfsFileSystemPGroup）用于更新文件系统所使用的权限组
func (c *Client) UpdateCfsFileSystemPGroup(request *UpdateCfsFileSystemPGroupRequest) (response *UpdateCfsFileSystemPGroupResponse, err error) {
	if request == nil {
		request = NewUpdateCfsFileSystemPGroupRequest()
	}
	response = NewUpdateCfsFileSystemPGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVipStatusRequest() (request *DescribeVipStatusRequest) {
	request = &DescribeVipStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeVipStatus")
	return
}

func NewDescribeVipStatusResponse() (response *DescribeVipStatusResponse) {
	response = &DescribeVipStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 内部接口，只供控制台使用
func (c *Client) DescribeVipStatus(request *DescribeVipStatusRequest) (response *DescribeVipStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVipStatusRequest()
	}
	response = NewDescribeVipStatusResponse()
	err = c.Send(request, response)
	return
}
