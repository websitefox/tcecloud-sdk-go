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

package v20220518

import (
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2022-05-18"

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

func NewGetFlowDetailRequest() (request *GetFlowDetailRequest) {
	request = &GetFlowDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "GetFlowDetail")
	return
}

func NewGetFlowDetailResponse() (response *GetFlowDetailResponse) {
	response = &GetFlowDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户端审批流详情
func (c *Client) GetFlowDetail(request *GetFlowDetailRequest) (response *GetFlowDetailResponse, err error) {
	if request == nil {
		request = NewGetFlowDetailRequest()
	}
	response = NewGetFlowDetailResponse()
	err = c.Send(request, response)
	return
}

func NewQueryApprovalFlowStatusRequest() (request *QueryApprovalFlowStatusRequest) {
	request = &QueryApprovalFlowStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "QueryApprovalFlowStatus")
	return
}

func NewQueryApprovalFlowStatusResponse() (response *QueryApprovalFlowStatusResponse) {
	response = &QueryApprovalFlowStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户关联的审批流状态
func (c *Client) QueryApprovalFlowStatus(request *QueryApprovalFlowStatusRequest) (response *QueryApprovalFlowStatusResponse, err error) {
	if request == nil {
		request = NewQueryApprovalFlowStatusRequest()
	}
	response = NewQueryApprovalFlowStatusResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCustomerApprovalDetailRequest() (request *QueryCustomerApprovalDetailRequest) {
	request = &QueryCustomerApprovalDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "QueryCustomerApprovalDetail")
	return
}

func NewQueryCustomerApprovalDetailResponse() (response *QueryCustomerApprovalDetailResponse) {
	response = &QueryCustomerApprovalDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户端审批单详情
func (c *Client) QueryCustomerApprovalDetail(request *QueryCustomerApprovalDetailRequest) (response *QueryCustomerApprovalDetailResponse, err error) {
	if request == nil {
		request = NewQueryCustomerApprovalDetailRequest()
	}
	response = NewQueryCustomerApprovalDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApprovalFlowRequest() (request *DeleteApprovalFlowRequest) {
	request = &DeleteApprovalFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "DeleteApprovalFlow")
	return
}

func NewDeleteApprovalFlowResponse() (response *DeleteApprovalFlowResponse) {
	response = &DeleteApprovalFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除审批流
func (c *Client) DeleteApprovalFlow(request *DeleteApprovalFlowRequest) (response *DeleteApprovalFlowResponse, err error) {
	if request == nil {
		request = NewDeleteApprovalFlowRequest()
	}
	response = NewDeleteApprovalFlowResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCustomerFlowRequest() (request *CreateCustomerFlowRequest) {
	request = &CreateCustomerFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "CreateCustomerFlow")
	return
}

func NewCreateCustomerFlowResponse() (response *CreateCustomerFlowResponse) {
	response = &CreateCustomerFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端创建审批流
func (c *Client) CreateCustomerFlow(request *CreateCustomerFlowRequest) (response *CreateCustomerFlowResponse, err error) {
	if request == nil {
		request = NewCreateCustomerFlowRequest()
	}
	response = NewCreateCustomerFlowResponse()
	err = c.Send(request, response)
	return
}

func NewWithdrawApplicationRequest() (request *WithdrawApplicationRequest) {
	request = &WithdrawApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "WithdrawApplication")
	return
}

func NewWithdrawApplicationResponse() (response *WithdrawApplicationResponse) {
	response = &WithdrawApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户撤销已提交的审批单， 会通知租户撤销结果， 当前审批人无需处理此审批单
func (c *Client) WithdrawApplication(request *WithdrawApplicationRequest) (response *WithdrawApplicationResponse, err error) {
	if request == nil {
		request = NewWithdrawApplicationRequest()
	}
	response = NewWithdrawApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewGetApprovedSetRequest() (request *GetApprovedSetRequest) {
	request = &GetApprovedSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "GetApprovedSet")
	return
}

func NewGetApprovedSetResponse() (response *GetApprovedSetResponse) {
	response = &GetApprovedSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户端用户当前已审批流程单
func (c *Client) GetApprovedSet(request *GetApprovedSetRequest) (response *GetApprovedSetResponse, err error) {
	if request == nil {
		request = NewGetApprovedSetRequest()
	}
	response = NewGetApprovedSetResponse()
	err = c.Send(request, response)
	return
}

func NewQueryActionSetRequest() (request *QueryActionSetRequest) {
	request = &QueryActionSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "QueryActionSet")
	return
}

func NewQueryActionSetResponse() (response *QueryActionSetResponse) {
	response = &QueryActionSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户端已接入审批流服务接口
func (c *Client) QueryActionSet(request *QueryActionSetRequest) (response *QueryActionSetResponse, err error) {
	if request == nil {
		request = NewQueryActionSetRequest()
	}
	response = NewQueryActionSetResponse()
	err = c.Send(request, response)
	return
}

func NewQueryApplicationFormRequest() (request *QueryApplicationFormRequest) {
	request = &QueryApplicationFormRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "QueryApplicationForm")
	return
}

func NewQueryApplicationFormResponse() (response *QueryApplicationFormResponse) {
	response = &QueryApplicationFormResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户端提交的申请单列表
func (c *Client) QueryApplicationForm(request *QueryApplicationFormRequest) (response *QueryApplicationFormResponse, err error) {
	if request == nil {
		request = NewQueryApplicationFormRequest()
	}
	response = NewQueryApplicationFormResponse()
	err = c.Send(request, response)
	return
}

func NewPerformApprovalRequest() (request *PerformApprovalRequest) {
	request = &PerformApprovalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "PerformApproval")
	return
}

func NewPerformApprovalResponse() (response *PerformApprovalResponse) {
	response = &PerformApprovalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户审批流程单接口
func (c *Client) PerformApproval(request *PerformApprovalRequest) (response *PerformApprovalResponse, err error) {
	if request == nil {
		request = NewPerformApprovalRequest()
	}
	response = NewPerformApprovalResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCustomerActionScopeRequest() (request *QueryCustomerActionScopeRequest) {
	request = &QueryCustomerActionScopeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "QueryCustomerActionScope")
	return
}

func NewQueryCustomerActionScopeResponse() (response *QueryCustomerActionScopeResponse) {
	response = &QueryCustomerActionScopeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户端接入审批流action 已关联的用户
func (c *Client) QueryCustomerActionScope(request *QueryCustomerActionScopeRequest) (response *QueryCustomerActionScopeResponse, err error) {
	if request == nil {
		request = NewQueryCustomerActionScopeRequest()
	}
	response = NewQueryCustomerActionScopeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCustomerApprovalDocRequest() (request *QueryCustomerApprovalDocRequest) {
	request = &QueryCustomerApprovalDocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "QueryCustomerApprovalDoc")
	return
}

func NewQueryCustomerApprovalDocResponse() (response *QueryCustomerApprovalDocResponse) {
	response = &QueryCustomerApprovalDocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户端审批单列表
func (c *Client) QueryCustomerApprovalDoc(request *QueryCustomerApprovalDocRequest) (response *QueryCustomerApprovalDocResponse, err error) {
	if request == nil {
		request = NewQueryCustomerApprovalDocRequest()
	}
	response = NewQueryCustomerApprovalDocResponse()
	err = c.Send(request, response)
	return
}

func NewSubmitApprovalDocRequest() (request *SubmitApprovalDocRequest) {
	request = &SubmitApprovalDocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "SubmitApprovalDoc")
	return
}

func NewSubmitApprovalDocResponse() (response *SubmitApprovalDocResponse) {
	response = &SubmitApprovalDocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户提交申请接口
func (c *Client) SubmitApprovalDoc(request *SubmitApprovalDocRequest) (response *SubmitApprovalDocResponse, err error) {
	if request == nil {
		request = NewSubmitApprovalDocRequest()
	}
	response = NewSubmitApprovalDocResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApprovalFlowRequest() (request *ModifyApprovalFlowRequest) {
	request = &ModifyApprovalFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "ModifyApprovalFlow")
	return
}

func NewModifyApprovalFlowResponse() (response *ModifyApprovalFlowResponse) {
	response = &ModifyApprovalFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改审批流
func (c *Client) ModifyApprovalFlow(request *ModifyApprovalFlowRequest) (response *ModifyApprovalFlowResponse, err error) {
	if request == nil {
		request = NewModifyApprovalFlowRequest()
	}
	response = NewModifyApprovalFlowResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCurrApprovalDocRequest() (request *QueryCurrApprovalDocRequest) {
	request = &QueryCurrApprovalDocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "QueryCurrApprovalDoc")
	return
}

func NewQueryCurrApprovalDocResponse() (response *QueryCurrApprovalDocResponse) {
	response = &QueryCurrApprovalDocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户端用户当前待审批流程单
func (c *Client) QueryCurrApprovalDoc(request *QueryCurrApprovalDocRequest) (response *QueryCurrApprovalDocResponse, err error) {
	if request == nil {
		request = NewQueryCurrApprovalDocRequest()
	}
	response = NewQueryCurrApprovalDocResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowSetRequest() (request *DescribeFlowSetRequest) {
	request = &DescribeFlowSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "DescribeFlowSet")
	return
}

func NewDescribeFlowSetResponse() (response *DescribeFlowSetResponse) {
	response = &DescribeFlowSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户端审批流列表
func (c *Client) DescribeFlowSet(request *DescribeFlowSetRequest) (response *DescribeFlowSetResponse, err error) {
	if request == nil {
		request = NewDescribeFlowSetRequest()
	}
	response = NewDescribeFlowSetResponse()
	err = c.Send(request, response)
	return
}

func NewOperateFlowStatusRequest() (request *OperateFlowStatusRequest) {
	request = &OperateFlowStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "OperateFlowStatus")
	return
}

func NewOperateFlowStatusResponse() (response *OperateFlowStatusResponse) {
	response = &OperateFlowStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 操作审批流状态
func (c *Client) OperateFlowStatus(request *OperateFlowStatusRequest) (response *OperateFlowStatusResponse, err error) {
	if request == nil {
		request = NewOperateFlowStatusRequest()
	}
	response = NewOperateFlowStatusResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCurrApprovalDetailRequest() (request *QueryCurrApprovalDetailRequest) {
	request = &QueryCurrApprovalDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tapproval", APIVersion, "QueryCurrApprovalDetail")
	return
}

func NewQueryCurrApprovalDetailResponse() (response *QueryCurrApprovalDetailResponse) {
	response = &QueryCurrApprovalDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 我的审批-“我的审批/我的申请”：查看审批详情接口
func (c *Client) QueryCurrApprovalDetail(request *QueryCurrApprovalDetailRequest) (response *QueryCurrApprovalDetailResponse, err error) {
	if request == nil {
		request = NewQueryCurrApprovalDetailRequest()
	}
	response = NewQueryCurrApprovalDetailResponse()
	err = c.Send(request, response)
	return
}
