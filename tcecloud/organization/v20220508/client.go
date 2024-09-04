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

package v20220508

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2022-05-08"

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

func NewDeleteOrganizationNodeMembersRequest() (request *DeleteOrganizationNodeMembersRequest) {
	request = &DeleteOrganizationNodeMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationNodeMembers")
	return
}

func NewDeleteOrganizationNodeMembersResponse() (response *DeleteOrganizationNodeMembersResponse) {
	response = &DeleteOrganizationNodeMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除节点成员
func (c *Client) DeleteOrganizationNodeMembers(request *DeleteOrganizationNodeMembersRequest) (response *DeleteOrganizationNodeMembersResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationNodeMembersRequest()
	}
	response = NewDeleteOrganizationNodeMembersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodesByParentRequest() (request *DescribeOrganizationNodesByParentRequest) {
	request = &DescribeOrganizationNodesByParentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodesByParent")
	return
}

func NewDescribeOrganizationNodesByParentResponse() (response *DescribeOrganizationNodesByParentResponse) {
	response = &DescribeOrganizationNodesByParentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据父节点查询子节点
func (c *Client) DescribeOrganizationNodesByParent(request *DescribeOrganizationNodesByParentRequest) (response *DescribeOrganizationNodesByParentResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodesByParentRequest()
	}
	response = NewDescribeOrganizationNodesByParentResponse()
	err = c.Send(request, response)
	return
}

func NewCancelOrganizationMemberAuthAccountRequest() (request *CancelOrganizationMemberAuthAccountRequest) {
	request = &CancelOrganizationMemberAuthAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationMemberAuthAccount")
	return
}

func NewCancelOrganizationMemberAuthAccountResponse() (response *CancelOrganizationMemberAuthAccountResponse) {
	response = &CancelOrganizationMemberAuthAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消成员和子账号的授权关系
func (c *Client) CancelOrganizationMemberAuthAccount(request *CancelOrganizationMemberAuthAccountRequest) (response *CancelOrganizationMemberAuthAccountResponse, err error) {
	if request == nil {
		request = NewCancelOrganizationMemberAuthAccountRequest()
	}
	response = NewCancelOrganizationMemberAuthAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationNodesRequest() (request *DeleteOrganizationNodesRequest) {
	request = &DeleteOrganizationNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationNodes")
	return
}

func NewDeleteOrganizationNodesResponse() (response *DeleteOrganizationNodesResponse) {
	response = &DeleteOrganizationNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除组织节点
func (c *Client) DeleteOrganizationNodes(request *DeleteOrganizationNodesRequest) (response *DeleteOrganizationNodesResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationNodesRequest()
	}
	response = NewDeleteOrganizationNodesResponse()
	err = c.Send(request, response)
	return
}

func NewAddOrganizationNodeRequest() (request *AddOrganizationNodeRequest) {
	request = &AddOrganizationNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "AddOrganizationNode")
	return
}

func NewAddOrganizationNodeResponse() (response *AddOrganizationNodeResponse) {
	response = &AddOrganizationNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加集团组织节点
func (c *Client) AddOrganizationNode(request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
	if request == nil {
		request = NewAddOrganizationNodeRequest()
	}
	response = NewAddOrganizationNodeResponse()
	err = c.Send(request, response)
	return
}

func NewMoveOrganizationNodeMembersRequest() (request *MoveOrganizationNodeMembersRequest) {
	request = &MoveOrganizationNodeMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "MoveOrganizationNodeMembers")
	return
}

func NewMoveOrganizationNodeMembersResponse() (response *MoveOrganizationNodeMembersResponse) {
	response = &MoveOrganizationNodeMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 移动节点成员
func (c *Client) MoveOrganizationNodeMembers(request *MoveOrganizationNodeMembersRequest) (response *MoveOrganizationNodeMembersResponse, err error) {
	if request == nil {
		request = NewMoveOrganizationNodeMembersRequest()
	}
	response = NewMoveOrganizationNodeMembersResponse()
	err = c.Send(request, response)
	return
}

func NewCheckMemberExistRequest() (request *CheckMemberExistRequest) {
	request = &CheckMemberExistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CheckMemberExist")
	return
}

func NewCheckMemberExistResponse() (response *CheckMemberExistResponse) {
	response = &CheckMemberExistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量检查成员账号是否存在，返回存在的账号列表
func (c *Client) CheckMemberExist(request *CheckMemberExistRequest) (response *CheckMemberExistResponse, err error) {
	if request == nil {
		request = NewCheckMemberExistRequest()
	}
	response = NewCheckMemberExistResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOrganizationNodeRequest() (request *UpdateOrganizationNodeRequest) {
	request = &UpdateOrganizationNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationNode")
	return
}

func NewUpdateOrganizationNodeResponse() (response *UpdateOrganizationNodeResponse) {
	response = &UpdateOrganizationNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新组织节点
func (c *Client) UpdateOrganizationNode(request *UpdateOrganizationNodeRequest) (response *UpdateOrganizationNodeResponse, err error) {
	if request == nil {
		request = NewUpdateOrganizationNodeRequest()
	}
	response = NewUpdateOrganizationNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMembersRequest() (request *DescribeOrganizationMembersRequest) {
	request = &DescribeOrganizationMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMembers")
	return
}

func NewDescribeOrganizationMembersResponse() (response *DescribeOrganizationMembersResponse) {
	response = &DescribeOrganizationMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 成员账号列表
func (c *Client) DescribeOrganizationMembers(request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMembersRequest()
	}
	response = NewDescribeOrganizationMembersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodeRequest() (request *DescribeOrganizationNodeRequest) {
	request = &DescribeOrganizationNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNode")
	return
}

func NewDescribeOrganizationNodeResponse() (response *DescribeOrganizationNodeResponse) {
	response = &DescribeOrganizationNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单节点信息
func (c *Client) DescribeOrganizationNode(request *DescribeOrganizationNodeRequest) (response *DescribeOrganizationNodeResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodeRequest()
	}
	response = NewDescribeOrganizationNodeResponse()
	err = c.Send(request, response)
	return
}

func NewBindOrganizationMemberAuthAccountRequest() (request *BindOrganizationMemberAuthAccountRequest) {
	request = &BindOrganizationMemberAuthAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "BindOrganizationMemberAuthAccount")
	return
}

func NewBindOrganizationMemberAuthAccountResponse() (response *BindOrganizationMemberAuthAccountResponse) {
	response = &BindOrganizationMemberAuthAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定成员和子账号授权关系
func (c *Client) BindOrganizationMemberAuthAccount(request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
	if request == nil {
		request = NewBindOrganizationMemberAuthAccountRequest()
	}
	response = NewBindOrganizationMemberAuthAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationRequest() (request *DescribeOrganizationRequest) {
	request = &DescribeOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganization")
	return
}

func NewDescribeOrganizationResponse() (response *DescribeOrganizationResponse) {
	response = &DescribeOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所属集团组织信息
func (c *Client) DescribeOrganization(request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationRequest()
	}
	response = NewDescribeOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodeMembersRequest() (request *DescribeOrganizationNodeMembersRequest) {
	request = &DescribeOrganizationNodeMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodeMembers")
	return
}

func NewDescribeOrganizationNodeMembersResponse() (response *DescribeOrganizationNodeMembersResponse) {
	response = &DescribeOrganizationNodeMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点成员列表
func (c *Client) DescribeOrganizationNodeMembers(request *DescribeOrganizationNodeMembersRequest) (response *DescribeOrganizationNodeMembersResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodeMembersRequest()
	}
	response = NewDescribeOrganizationNodeMembersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrganizationMemberPolicyRequest() (request *CreateOrganizationMemberPolicyRequest) {
	request = &CreateOrganizationMemberPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMemberPolicy")
	return
}

func NewCreateOrganizationMemberPolicyResponse() (response *CreateOrganizationMemberPolicyResponse) {
	response = &CreateOrganizationMemberPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加成员授权策略
func (c *Client) CreateOrganizationMemberPolicy(request *CreateOrganizationMemberPolicyRequest) (response *CreateOrganizationMemberPolicyResponse, err error) {
	if request == nil {
		request = NewCreateOrganizationMemberPolicyRequest()
	}
	response = NewCreateOrganizationMemberPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMemberRequest() (request *DescribeOrganizationMemberRequest) {
	request = &DescribeOrganizationMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMember")
	return
}

func NewDescribeOrganizationMemberResponse() (response *DescribeOrganizationMemberResponse) {
	response = &DescribeOrganizationMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织成员详情
func (c *Client) DescribeOrganizationMember(request *DescribeOrganizationMemberRequest) (response *DescribeOrganizationMemberResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMemberRequest()
	}
	response = NewDescribeOrganizationMemberResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodesRequest() (request *DescribeOrganizationNodesRequest) {
	request = &DescribeOrganizationNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodes")
	return
}

func NewDescribeOrganizationNodesResponse() (response *DescribeOrganizationNodesResponse) {
	response = &DescribeOrganizationNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组织节点列表
func (c *Client) DescribeOrganizationNodes(request *DescribeOrganizationNodesRequest) (response *DescribeOrganizationNodesResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodesRequest()
	}
	response = NewDescribeOrganizationNodesResponse()
	err = c.Send(request, response)
	return
}

func NewCancelOrganizationMemberAuthAccountForDeletionSubAccountRequest() (request *CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest) {
	request = &CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationMemberAuthAccountForDeletionSubAccount")
	return
}

func NewCancelOrganizationMemberAuthAccountForDeletionSubAccountResponse() (response *CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse) {
	response = &CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消成员和子账号的授权关系(删除子账号消息调用)
func (c *Client) CancelOrganizationMemberAuthAccountForDeletionSubAccount(request *CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest) (response *CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse, err error) {
	if request == nil {
		request = NewCancelOrganizationMemberAuthAccountForDeletionSubAccountRequest()
	}
	response = NewCancelOrganizationMemberAuthAccountForDeletionSubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewMoveOrganizationNodeRequest() (request *MoveOrganizationNodeRequest) {
	request = &MoveOrganizationNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "MoveOrganizationNode")
	return
}

func NewMoveOrganizationNodeResponse() (response *MoveOrganizationNodeResponse) {
	response = &MoveOrganizationNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 移动组织节点
func (c *Client) MoveOrganizationNode(request *MoveOrganizationNodeRequest) (response *MoveOrganizationNodeResponse, err error) {
	if request == nil {
		request = NewMoveOrganizationNodeRequest()
	}
	response = NewMoveOrganizationNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMemberAuthAccountsRequest() (request *DescribeOrganizationMemberAuthAccountsRequest) {
	request = &DescribeOrganizationMemberAuthAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberAuthAccounts")
	return
}

func NewDescribeOrganizationMemberAuthAccountsResponse() (response *DescribeOrganizationMemberAuthAccountsResponse) {
	response = &DescribeOrganizationMemberAuthAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询成员授权子账号列表
func (c *Client) DescribeOrganizationMemberAuthAccounts(request *DescribeOrganizationMemberAuthAccountsRequest) (response *DescribeOrganizationMemberAuthAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMemberAuthAccountsRequest()
	}
	response = NewDescribeOrganizationMemberAuthAccountsResponse()
	err = c.Send(request, response)
	return
}
