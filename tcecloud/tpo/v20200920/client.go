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

package v20200920

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-09-20"

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

func NewDescribeProjectResourcesRequest() (request *DescribeProjectResourcesRequest) {
	request = &DescribeProjectResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeProjectResources")
	return
}

func NewDescribeProjectResourcesResponse() (response *DescribeProjectResourcesResponse) {
	response = &DescribeProjectResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询项目资源列表
func (c *Client) DescribeProjectResources(request *DescribeProjectResourcesRequest) (response *DescribeProjectResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeProjectResourcesRequest()
	}
	response = NewDescribeProjectResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectMembersRequest() (request *DescribeProjectMembersRequest) {
	request = &DescribeProjectMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeProjectMembers")
	return
}

func NewDescribeProjectMembersResponse() (response *DescribeProjectMembersResponse) {
	response = &DescribeProjectMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询项目成员
func (c *Client) DescribeProjectMembers(request *DescribeProjectMembersRequest) (response *DescribeProjectMembersResponse, err error) {
	if request == nil {
		request = NewDescribeProjectMembersRequest()
	}
	response = NewDescribeProjectMembersResponse()
	err = c.Send(request, response)
	return
}

func NewCheckProjectQuotasRequest() (request *CheckProjectQuotasRequest) {
	request = &CheckProjectQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "CheckProjectQuotas")
	return
}

func NewCheckProjectQuotasResponse() (response *CheckProjectQuotasResponse) {
	response = &CheckProjectQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查配额合法性，转换配额名称到编码
func (c *Client) CheckProjectQuotas(request *CheckProjectQuotasRequest) (response *CheckProjectQuotasResponse, err error) {
	if request == nil {
		request = NewCheckProjectQuotasRequest()
	}
	response = NewCheckProjectQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
	request = &CreateProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "CreateProject")
	return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
	response = &CreateProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建项目
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
	if request == nil {
		request = NewCreateProjectRequest()
	}
	response = NewCreateProjectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectMemberPoliciesRequest() (request *DescribeProjectMemberPoliciesRequest) {
	request = &DescribeProjectMemberPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeProjectMemberPolicies")
	return
}

func NewDescribeProjectMemberPoliciesResponse() (response *DescribeProjectMemberPoliciesResponse) {
	response = &DescribeProjectMemberPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户的项目策略列表
func (c *Client) DescribeProjectMemberPolicies(request *DescribeProjectMemberPoliciesRequest) (response *DescribeProjectMemberPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeProjectMemberPoliciesRequest()
	}
	response = NewDescribeProjectMemberPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectPoliciesRequest() (request *DescribeProjectPoliciesRequest) {
	request = &DescribeProjectPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeProjectPolicies")
	return
}

func NewDescribeProjectPoliciesResponse() (response *DescribeProjectPoliciesResponse) {
	response = &DescribeProjectPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询项目所有策略列表
func (c *Client) DescribeProjectPolicies(request *DescribeProjectPoliciesRequest) (response *DescribeProjectPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeProjectPoliciesRequest()
	}
	response = NewDescribeProjectPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewAddProjectResourceRequest() (request *AddProjectResourceRequest) {
	request = &AddProjectResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "AddProjectResource")
	return
}

func NewAddProjectResourceResponse() (response *AddProjectResourceResponse) {
	response = &AddProjectResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 向项目里添加资源
func (c *Client) AddProjectResource(request *AddProjectResourceRequest) (response *AddProjectResourceResponse, err error) {
	if request == nil {
		request = NewAddProjectResourceRequest()
	}
	response = NewAddProjectResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductTreeRequest() (request *DescribeProductTreeRequest) {
	request = &DescribeProductTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeProductTree")
	return
}

func NewDescribeProductTreeResponse() (response *DescribeProductTreeResponse) {
	response = &DescribeProductTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询各级产品
func (c *Client) DescribeProductTree(request *DescribeProductTreeRequest) (response *DescribeProductTreeResponse, err error) {
	if request == nil {
		request = NewDescribeProductTreeRequest()
	}
	response = NewDescribeProductTreeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
	request = &DescribeProjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeProjects")
	return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
	response = &DescribeProjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询项目列表
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
	if request == nil {
		request = NewDescribeProjectsRequest()
	}
	response = NewDescribeProjectsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
	request = &DescribeResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeResources")
	return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
	response = &DescribeResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源列表
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeResourcesRequest()
	}
	response = NewDescribeResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceAdminProjectsRequest() (request *DescribeResourceAdminProjectsRequest) {
	request = &DescribeResourceAdminProjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeResourceAdminProjects")
	return
}

func NewDescribeResourceAdminProjectsResponse() (response *DescribeResourceAdminProjectsResponse) {
	response = &DescribeResourceAdminProjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询具有转入转出资源权限的项目列表
func (c *Client) DescribeResourceAdminProjects(request *DescribeResourceAdminProjectsRequest) (response *DescribeResourceAdminProjectsResponse, err error) {
	if request == nil {
		request = NewDescribeResourceAdminProjectsRequest()
	}
	response = NewDescribeResourceAdminProjectsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProjectResourceRequest() (request *DeleteProjectResourceRequest) {
	request = &DeleteProjectResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DeleteProjectResource")
	return
}

func NewDeleteProjectResourceResponse() (response *DeleteProjectResourceResponse) {
	response = &DeleteProjectResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将资源移出项目
func (c *Client) DeleteProjectResource(request *DeleteProjectResourceRequest) (response *DeleteProjectResourceResponse, err error) {
	if request == nil {
		request = NewDeleteProjectResourceRequest()
	}
	response = NewDeleteProjectResourceResponse()
	err = c.Send(request, response)
	return
}

func NewProjectNameExistsRequest() (request *ProjectNameExistsRequest) {
	request = &ProjectNameExistsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "ProjectNameExists")
	return
}

func NewProjectNameExistsResponse() (response *ProjectNameExistsResponse) {
	response = &ProjectNameExistsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 项目名是否唯一
func (c *Client) ProjectNameExists(request *ProjectNameExistsRequest) (response *ProjectNameExistsResponse, err error) {
	if request == nil {
		request = NewProjectNameExistsRequest()
	}
	response = NewProjectNameExistsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductsRequest() (request *DescribeProductsRequest) {
	request = &DescribeProductsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeProducts")
	return
}

func NewDescribeProductsResponse() (response *DescribeProductsResponse) {
	response = &DescribeProductsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询接入产品列表
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
	if request == nil {
		request = NewDescribeProductsRequest()
	}
	response = NewDescribeProductsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProjectQuotaRequest() (request *DeleteProjectQuotaRequest) {
	request = &DeleteProjectQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DeleteProjectQuota")
	return
}

func NewDeleteProjectQuotaResponse() (response *DeleteProjectQuotaResponse) {
	response = &DeleteProjectQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除项目配额
func (c *Client) DeleteProjectQuota(request *DeleteProjectQuotaRequest) (response *DeleteProjectQuotaResponse, err error) {
	if request == nil {
		request = NewDeleteProjectQuotaRequest()
	}
	response = NewDeleteProjectQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProjectMemberPolicyRequest() (request *ModifyProjectMemberPolicyRequest) {
	request = &ModifyProjectMemberPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "ModifyProjectMemberPolicy")
	return
}

func NewModifyProjectMemberPolicyResponse() (response *ModifyProjectMemberPolicyResponse) {
	response = &ModifyProjectMemberPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改单个用户的授权策略
func (c *Client) ModifyProjectMemberPolicy(request *ModifyProjectMemberPolicyRequest) (response *ModifyProjectMemberPolicyResponse, err error) {
	if request == nil {
		request = NewModifyProjectMemberPolicyRequest()
	}
	response = NewModifyProjectMemberPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewAddProjectQuotaRequest() (request *AddProjectQuotaRequest) {
	request = &AddProjectQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "AddProjectQuota")
	return
}

func NewAddProjectQuotaResponse() (response *AddProjectQuotaResponse) {
	response = &AddProjectQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加项目配额
func (c *Client) AddProjectQuota(request *AddProjectQuotaRequest) (response *AddProjectQuotaResponse, err error) {
	if request == nil {
		request = NewAddProjectQuotaRequest()
	}
	response = NewAddProjectQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectQuotasRequest() (request *DescribeProjectQuotasRequest) {
	request = &DescribeProjectQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeProjectQuotas")
	return
}

func NewDescribeProjectQuotasResponse() (response *DescribeProjectQuotasResponse) {
	response = &DescribeProjectQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询项目配额
func (c *Client) DescribeProjectQuotas(request *DescribeProjectQuotasRequest) (response *DescribeProjectQuotasResponse, err error) {
	if request == nil {
		request = NewDescribeProjectQuotasRequest()
	}
	response = NewDescribeProjectQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceRegionsRequest() (request *DescribeResourceRegionsRequest) {
	request = &DescribeResourceRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeResourceRegions")
	return
}

func NewDescribeResourceRegionsResponse() (response *DescribeResourceRegionsResponse) {
	response = &DescribeResourceRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有资源区域
func (c *Client) DescribeResourceRegions(request *DescribeResourceRegionsRequest) (response *DescribeResourceRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeResourceRegionsRequest()
	}
	response = NewDescribeResourceRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveProjectMemberRequest() (request *RemoveProjectMemberRequest) {
	request = &RemoveProjectMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "RemoveProjectMember")
	return
}

func NewRemoveProjectMemberResponse() (response *RemoveProjectMemberResponse) {
	response = &RemoveProjectMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从项目移除成员
func (c *Client) RemoveProjectMember(request *RemoveProjectMemberRequest) (response *RemoveProjectMemberResponse, err error) {
	if request == nil {
		request = NewRemoveProjectMemberRequest()
	}
	response = NewRemoveProjectMemberResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectNonMembersRequest() (request *DescribeProjectNonMembersRequest) {
	request = &DescribeProjectNonMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeProjectNonMembers")
	return
}

func NewDescribeProjectNonMembersResponse() (response *DescribeProjectNonMembersResponse) {
	response = &DescribeProjectNonMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前租户下所有非项目成员
func (c *Client) DescribeProjectNonMembers(request *DescribeProjectNonMembersRequest) (response *DescribeProjectNonMembersResponse, err error) {
	if request == nil {
		request = NewDescribeProjectNonMembersRequest()
	}
	response = NewDescribeProjectNonMembersResponse()
	err = c.Send(request, response)
	return
}

func NewBatchAddProjectQuotaRequest() (request *BatchAddProjectQuotaRequest) {
	request = &BatchAddProjectQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "BatchAddProjectQuota")
	return
}

func NewBatchAddProjectQuotaResponse() (response *BatchAddProjectQuotaResponse) {
	response = &BatchAddProjectQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入配额
func (c *Client) BatchAddProjectQuota(request *BatchAddProjectQuotaRequest) (response *BatchAddProjectQuotaResponse, err error) {
	if request == nil {
		request = NewBatchAddProjectQuotaRequest()
	}
	response = NewBatchAddProjectQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewMoveProjectResourceRequest() (request *MoveProjectResourceRequest) {
	request = &MoveProjectResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "MoveProjectResource")
	return
}

func NewMoveProjectResourceResponse() (response *MoveProjectResourceResponse) {
	response = &MoveProjectResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) MoveProjectResource(request *MoveProjectResourceRequest) (response *MoveProjectResourceResponse, err error) {
	if request == nil {
		request = NewMoveProjectResourceRequest()
	}
	response = NewMoveProjectResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceTypesRequest() (request *DescribeResourceTypesRequest) {
	request = &DescribeResourceTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DescribeResourceTypes")
	return
}

func NewDescribeResourceTypesResponse() (response *DescribeResourceTypesResponse) {
	response = &DescribeResourceTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源的类型，如instance, image
func (c *Client) DescribeResourceTypes(request *DescribeResourceTypesRequest) (response *DescribeResourceTypesResponse, err error) {
	if request == nil {
		request = NewDescribeResourceTypesRequest()
	}
	response = NewDescribeResourceTypesResponse()
	err = c.Send(request, response)
	return
}

func NewTransferProjectResourceRequest() (request *TransferProjectResourceRequest) {
	request = &TransferProjectResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "TransferProjectResource")
	return
}

func NewTransferProjectResourceResponse() (response *TransferProjectResourceResponse) {
	response = &TransferProjectResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 转入转出资源
func (c *Client) TransferProjectResource(request *TransferProjectResourceRequest) (response *TransferProjectResourceResponse, err error) {
	if request == nil {
		request = NewTransferProjectResourceRequest()
	}
	response = NewTransferProjectResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
	request = &DeleteProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "DeleteProject")
	return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
	response = &DeleteProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除项目
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
	if request == nil {
		request = NewDeleteProjectRequest()
	}
	response = NewDeleteProjectResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProjectNameRequest() (request *ModifyProjectNameRequest) {
	request = &ModifyProjectNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "ModifyProjectName")
	return
}

func NewModifyProjectNameResponse() (response *ModifyProjectNameResponse) {
	response = &ModifyProjectNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改项目名称
func (c *Client) ModifyProjectName(request *ModifyProjectNameRequest) (response *ModifyProjectNameResponse, err error) {
	if request == nil {
		request = NewModifyProjectNameRequest()
	}
	response = NewModifyProjectNameResponse()
	err = c.Send(request, response)
	return
}

func NewAddProjectMemberPolicyRequest() (request *AddProjectMemberPolicyRequest) {
	request = &AddProjectMemberPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "AddProjectMemberPolicy")
	return
}

func NewAddProjectMemberPolicyResponse() (response *AddProjectMemberPolicyResponse) {
	response = &AddProjectMemberPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 项目添加用户及授权
func (c *Client) AddProjectMemberPolicy(request *AddProjectMemberPolicyRequest) (response *AddProjectMemberPolicyResponse, err error) {
	if request == nil {
		request = NewAddProjectMemberPolicyRequest()
	}
	response = NewAddProjectMemberPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProjectQuotaRequest() (request *ModifyProjectQuotaRequest) {
	request = &ModifyProjectQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tpo", APIVersion, "ModifyProjectQuota")
	return
}

func NewModifyProjectQuotaResponse() (response *ModifyProjectQuotaResponse) {
	response = &ModifyProjectQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改项目配额
func (c *Client) ModifyProjectQuota(request *ModifyProjectQuotaRequest) (response *ModifyProjectQuotaResponse, err error) {
	if request == nil {
		request = NewModifyProjectQuotaRequest()
	}
	response = NewModifyProjectQuotaResponse()
	err = c.Send(request, response)
	return
}
