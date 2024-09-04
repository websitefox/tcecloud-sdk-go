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

package v20211001

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-10-01"

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

func NewDescribeOrganizationsRequest() (request *DescribeOrganizationsRequest) {
	request = &DescribeOrganizationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "DescribeOrganizations")
	return
}

func NewDescribeOrganizationsResponse() (response *DescribeOrganizationsResponse) {
	response = &DescribeOrganizationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组织
func (c *Client) DescribeOrganizations(request *DescribeOrganizationsRequest) (response *DescribeOrganizationsResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationsRequest()
	}
	response = NewDescribeOrganizationsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationRequest() (request *DeleteOrganizationRequest) {
	request = &DeleteOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "DeleteOrganization")
	return
}

func NewDeleteOrganizationResponse() (response *DeleteOrganizationResponse) {
	response = &DeleteOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除组织，要求当前组织和其子组织没有关联项目，否则不允许删除
func (c *Client) DeleteOrganization(request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationRequest()
	}
	response = NewDeleteOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationMembersRequest() (request *DeleteOrganizationMembersRequest) {
	request = &DeleteOrganizationMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "DeleteOrganizationMembers")
	return
}

func NewDeleteOrganizationMembersResponse() (response *DeleteOrganizationMembersResponse) {
	response = &DeleteOrganizationMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除组织成员
func (c *Client) DeleteOrganizationMembers(request *DeleteOrganizationMembersRequest) (response *DeleteOrganizationMembersResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationMembersRequest()
	}
	response = NewDeleteOrganizationMembersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationProjectsRequest() (request *DescribeOrganizationProjectsRequest) {
	request = &DescribeOrganizationProjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "DescribeOrganizationProjects")
	return
}

func NewDescribeOrganizationProjectsResponse() (response *DescribeOrganizationProjectsResponse) {
	response = &DescribeOrganizationProjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组织下的项目列表
func (c *Client) DescribeOrganizationProjects(request *DescribeOrganizationProjectsRequest) (response *DescribeOrganizationProjectsResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationProjectsRequest()
	}
	response = NewDescribeOrganizationProjectsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNonMembersRequest() (request *DescribeOrganizationNonMembersRequest) {
	request = &DescribeOrganizationNonMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "DescribeOrganizationNonMembers")
	return
}

func NewDescribeOrganizationNonMembersResponse() (response *DescribeOrganizationNonMembersResponse) {
	response = &DescribeOrganizationNonMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有非组织成员
func (c *Client) DescribeOrganizationNonMembers(request *DescribeOrganizationNonMembersRequest) (response *DescribeOrganizationNonMembersResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNonMembersRequest()
	}
	response = NewDescribeOrganizationNonMembersResponse()
	err = c.Send(request, response)
	return
}

func NewAddOrganizationRequest() (request *AddOrganizationRequest) {
	request = &AddOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "AddOrganization")
	return
}

func NewAddOrganizationResponse() (response *AddOrganizationResponse) {
	response = &AddOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加组织
func (c *Client) AddOrganization(request *AddOrganizationRequest) (response *AddOrganizationResponse, err error) {
	if request == nil {
		request = NewAddOrganizationRequest()
	}
	response = NewAddOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationResourcesRequest() (request *DescribeOrganizationResourcesRequest) {
	request = &DescribeOrganizationResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "DescribeOrganizationResources")
	return
}

func NewDescribeOrganizationResourcesResponse() (response *DescribeOrganizationResourcesResponse) {
	response = &DescribeOrganizationResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组织下项目资源列表
func (c *Client) DescribeOrganizationResources(request *DescribeOrganizationResourcesRequest) (response *DescribeOrganizationResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationResourcesRequest()
	}
	response = NewDescribeOrganizationResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewAddOrganizationMemberPolicyRequest() (request *AddOrganizationMemberPolicyRequest) {
	request = &AddOrganizationMemberPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "AddOrganizationMemberPolicy")
	return
}

func NewAddOrganizationMemberPolicyResponse() (response *AddOrganizationMemberPolicyResponse) {
	response = &AddOrganizationMemberPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加用户到组织并授权
func (c *Client) AddOrganizationMemberPolicy(request *AddOrganizationMemberPolicyRequest) (response *AddOrganizationMemberPolicyResponse, err error) {
	if request == nil {
		request = NewAddOrganizationMemberPolicyRequest()
	}
	response = NewAddOrganizationMemberPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOrganizationMemberPolicyRequest() (request *ModifyOrganizationMemberPolicyRequest) {
	request = &ModifyOrganizationMemberPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "ModifyOrganizationMemberPolicy")
	return
}

func NewModifyOrganizationMemberPolicyResponse() (response *ModifyOrganizationMemberPolicyResponse) {
	response = &ModifyOrganizationMemberPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改组织成员策略
func (c *Client) ModifyOrganizationMemberPolicy(request *ModifyOrganizationMemberPolicyRequest) (response *ModifyOrganizationMemberPolicyResponse, err error) {
	if request == nil {
		request = NewModifyOrganizationMemberPolicyRequest()
	}
	response = NewModifyOrganizationMemberPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationQuotasRequest() (request *DescribeOrganizationQuotasRequest) {
	request = &DescribeOrganizationQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "DescribeOrganizationQuotas")
	return
}

func NewDescribeOrganizationQuotasResponse() (response *DescribeOrganizationQuotasResponse) {
	response = &DescribeOrganizationQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组织下的项目配额
func (c *Client) DescribeOrganizationQuotas(request *DescribeOrganizationQuotasRequest) (response *DescribeOrganizationQuotasResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationQuotasRequest()
	}
	response = NewDescribeOrganizationQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMembersRequest() (request *DescribeOrganizationMembersRequest) {
	request = &DescribeOrganizationMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "DescribeOrganizationMembers")
	return
}

func NewDescribeOrganizationMembersResponse() (response *DescribeOrganizationMembersResponse) {
	response = &DescribeOrganizationMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看组织成员
func (c *Client) DescribeOrganizationMembers(request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMembersRequest()
	}
	response = NewDescribeOrganizationMembersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOrganizationRequest() (request *ModifyOrganizationRequest) {
	request = &ModifyOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "ModifyOrganization")
	return
}

func NewModifyOrganizationResponse() (response *ModifyOrganizationResponse) {
	response = &ModifyOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改组织
func (c *Client) ModifyOrganization(request *ModifyOrganizationRequest) (response *ModifyOrganizationResponse, err error) {
	if request == nil {
		request = NewModifyOrganizationRequest()
	}
	response = NewModifyOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationPoliciesTemplateRequest() (request *DescribeOrganizationPoliciesTemplateRequest) {
	request = &DescribeOrganizationPoliciesTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "DescribeOrganizationPoliciesTemplate")
	return
}

func NewDescribeOrganizationPoliciesTemplateResponse() (response *DescribeOrganizationPoliciesTemplateResponse) {
	response = &DescribeOrganizationPoliciesTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组织通用策略模板
func (c *Client) DescribeOrganizationPoliciesTemplate(request *DescribeOrganizationPoliciesTemplateRequest) (response *DescribeOrganizationPoliciesTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationPoliciesTemplateRequest()
	}
	response = NewDescribeOrganizationPoliciesTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOrganizationProjectsRequest() (request *ModifyOrganizationProjectsRequest) {
	request = &ModifyOrganizationProjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("org", APIVersion, "ModifyOrganizationProjects")
	return
}

func NewModifyOrganizationProjectsResponse() (response *ModifyOrganizationProjectsResponse) {
	response = &ModifyOrganizationProjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加/移除组织项目
func (c *Client) ModifyOrganizationProjects(request *ModifyOrganizationProjectsRequest) (response *ModifyOrganizationProjectsResponse, err error) {
	if request == nil {
		request = NewModifyOrganizationProjectsRequest()
	}
	response = NewModifyOrganizationProjectsResponse()
	err = c.Send(request, response)
	return
}
