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

package v20180416

import (
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-04-16"

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

func NewCreateCfsRuleRequest() (request *CreateCfsRuleRequest) {
	request = &CreateCfsRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "CreateCfsRule")
	return
}

func NewCreateCfsRuleResponse() (response *CreateCfsRuleResponse) {
	response = &CreateCfsRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCfsRule）用于创建权限组规则。
func (c *Client) CreateCfsRule(request *CreateCfsRuleRequest) (response *CreateCfsRuleResponse, err error) {
	if request == nil {
		request = NewCreateCfsRuleRequest()
	}
	response = NewCreateCfsRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCfsPGroupRequest() (request *CreateCfsPGroupRequest) {
	request = &CreateCfsPGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "CreateCfsPGroup")
	return
}

func NewCreateCfsPGroupResponse() (response *CreateCfsPGroupResponse) {
	response = &CreateCfsPGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCfsPGroup）用于创建权限组
func (c *Client) CreateCfsPGroup(request *CreateCfsPGroupRequest) (response *CreateCfsPGroupResponse, err error) {
	if request == nil {
		request = NewCreateCfsPGroupRequest()
	}
	response = NewCreateCfsPGroupResponse()
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
