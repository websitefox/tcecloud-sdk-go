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

package v20180420

import (
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-04-20"

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

func NewCreateVpcPrivateLinkRequest() (request *CreateVpcPrivateLinkRequest) {
	request = &CreateVpcPrivateLinkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateVpcPrivateLink")
	return
}

func NewCreateVpcPrivateLinkResponse() (response *CreateVpcPrivateLinkResponse) {
	response = &CreateVpcPrivateLinkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加Vpc通道
func (c *Client) CreateVpcPrivateLink(request *CreateVpcPrivateLinkRequest) (response *CreateVpcPrivateLinkResponse, err error) {
	if request == nil {
		request = NewCreateVpcPrivateLinkRequest()
	}
	response = NewCreateVpcPrivateLinkResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWhiteListRequest() (request *DeleteWhiteListRequest) {
	request = &DeleteWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DeleteWhiteList")
	return
}

func NewDeleteWhiteListResponse() (response *DeleteWhiteListResponse) {
	response = &DeleteWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除白名单
func (c *Client) DeleteWhiteList(request *DeleteWhiteListRequest) (response *DeleteWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteWhiteListRequest()
	}
	response = NewDeleteWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmEnterpriseWechatConfigRequest() (request *DescribeAlarmEnterpriseWechatConfigRequest) {
	request = &DescribeAlarmEnterpriseWechatConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAlarmEnterpriseWechatConfig")
	return
}

func NewDescribeAlarmEnterpriseWechatConfigResponse() (response *DescribeAlarmEnterpriseWechatConfigResponse) {
	response = &DescribeAlarmEnterpriseWechatConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取企业微信配置
func (c *Client) DescribeAlarmEnterpriseWechatConfig(request *DescribeAlarmEnterpriseWechatConfigRequest) (response *DescribeAlarmEnterpriseWechatConfigResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmEnterpriseWechatConfigRequest()
	}
	response = NewDescribeAlarmEnterpriseWechatConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleInfoRequest() (request *DescribeRuleInfoRequest) {
	request = &DescribeRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeRuleInfo")
	return
}

func NewDescribeRuleInfoResponse() (response *DescribeRuleInfoResponse) {
	response = &DescribeRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取对应 ruleID 的数据
func (c *Client) DescribeRuleInfo(request *DescribeRuleInfoRequest) (response *DescribeRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRuleInfoRequest()
	}
	response = NewDescribeRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTencentDbRequest() (request *DescribeTencentDbRequest) {
	request = &DescribeTencentDbRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeTencentDb")
	return
}

func NewDescribeTencentDbResponse() (response *DescribeTencentDbResponse) {
	response = &DescribeTencentDbResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 Tencent db 拉取 cdb 信息 并自动导入到资产 db
func (c *Client) DescribeTencentDb(request *DescribeTencentDbRequest) (response *DescribeTencentDbResponse, err error) {
	if request == nil {
		request = NewDescribeTencentDbRequest()
	}
	response = NewDescribeTencentDbResponse()
	err = c.Send(request, response)
	return
}

func NewModifycdsAuditSaasResourceRequest() (request *ModifycdsAuditSaasResourceRequest) {
	request = &ModifycdsAuditSaasResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifycdsAuditSaasResource")
	return
}

func NewModifycdsAuditSaasResourceResponse() (response *ModifycdsAuditSaasResourceResponse) {
	response = &ModifycdsAuditSaasResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifycdsAuditSaasResource
func (c *Client) ModifycdsAuditSaasResource(request *ModifycdsAuditSaasResourceRequest) (response *ModifycdsAuditSaasResourceResponse, err error) {
	if request == nil {
		request = NewModifycdsAuditSaasResourceRequest()
	}
	response = NewModifycdsAuditSaasResourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperateRuleSaveRequest() (request *ModifyOperateRuleSaveRequest) {
	request = &ModifyOperateRuleSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyOperateRuleSave")
	return
}

func NewModifyOperateRuleSaveResponse() (response *ModifyOperateRuleSaveResponse) {
	response = &ModifyOperateRuleSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改行为分类配置
func (c *Client) ModifyOperateRuleSave(request *ModifyOperateRuleSaveRequest) (response *ModifyOperateRuleSaveResponse, err error) {
	if request == nil {
		request = NewModifyOperateRuleSaveRequest()
	}
	response = NewModifyOperateRuleSaveResponse()
	err = c.Send(request, response)
	return
}

func NewSendAlarmSmsTestRequest() (request *SendAlarmSmsTestRequest) {
	request = &SendAlarmSmsTestRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "SendAlarmSmsTest")
	return
}

func NewSendAlarmSmsTestResponse() (response *SendAlarmSmsTestResponse) {
	response = &SendAlarmSmsTestResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试短信配置
func (c *Client) SendAlarmSmsTest(request *SendAlarmSmsTestRequest) (response *SendAlarmSmsTestResponse, err error) {
	if request == nil {
		request = NewSendAlarmSmsTestRequest()
	}
	response = NewSendAlarmSmsTestResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAgentSaveRequest() (request *ModifyAgentSaveRequest) {
	request = &ModifyAgentSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyAgentSave")
	return
}

func NewModifyAgentSaveResponse() (response *ModifyAgentSaveResponse) {
	response = &ModifyAgentSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑Agent
func (c *Client) ModifyAgentSave(request *ModifyAgentSaveRequest) (response *ModifyAgentSaveResponse, err error) {
	if request == nil {
		request = NewModifyAgentSaveRequest()
	}
	response = NewModifyAgentSaveResponse()
	err = c.Send(request, response)
	return
}

func NewInquireCreatecdsAuditSaasResourceRequest() (request *InquireCreatecdsAuditSaasResourceRequest) {
	request = &InquireCreatecdsAuditSaasResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "InquireCreatecdsAuditSaasResource")
	return
}

func NewInquireCreatecdsAuditSaasResourceResponse() (response *InquireCreatecdsAuditSaasResourceResponse) {
	response = &InquireCreatecdsAuditSaasResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// InquireCreatecdsAuditSaasResource
func (c *Client) InquireCreatecdsAuditSaasResource(request *InquireCreatecdsAuditSaasResourceRequest) (response *InquireCreatecdsAuditSaasResourceResponse, err error) {
	if request == nil {
		request = NewInquireCreatecdsAuditSaasResourceRequest()
	}
	response = NewInquireCreatecdsAuditSaasResourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetsSaveRequest() (request *ModifyAssetsSaveRequest) {
	request = &ModifyAssetsSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyAssetsSave")
	return
}

func NewModifyAssetsSaveResponse() (response *ModifyAssetsSaveResponse) {
	response = &ModifyAssetsSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改数据资产
func (c *Client) ModifyAssetsSave(request *ModifyAssetsSaveRequest) (response *ModifyAssetsSaveResponse, err error) {
	if request == nil {
		request = NewModifyAssetsSaveRequest()
	}
	response = NewModifyAssetsSaveResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetsPermissionRequest() (request *ModifyAssetsPermissionRequest) {
	request = &ModifyAssetsPermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyAssetsPermission")
	return
}

func NewModifyAssetsPermissionResponse() (response *ModifyAssetsPermissionResponse) {
	response = &ModifyAssetsPermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改审计权限
func (c *Client) ModifyAssetsPermission(request *ModifyAssetsPermissionRequest) (response *ModifyAssetsPermissionResponse, err error) {
	if request == nil {
		request = NewModifyAssetsPermissionRequest()
	}
	response = NewModifyAssetsPermissionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReportRequest() (request *DescribeReportRequest) {
	request = &DescribeReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeReport")
	return
}

func NewDescribeReportResponse() (response *DescribeReportResponse) {
	response = &DescribeReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看报表
func (c *Client) DescribeReport(request *DescribeReportRequest) (response *DescribeReportResponse, err error) {
	if request == nil {
		request = NewDescribeReportRequest()
	}
	response = NewDescribeReportResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTimerReportRequest() (request *CreateTimerReportRequest) {
	request = &CreateTimerReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateTimerReport")
	return
}

func NewCreateTimerReportResponse() (response *CreateTimerReportResponse) {
	response = &CreateTimerReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建报表任务
func (c *Client) CreateTimerReport(request *CreateTimerReportRequest) (response *CreateTimerReportResponse, err error) {
	if request == nil {
		request = NewCreateTimerReportRequest()
	}
	response = NewCreateTimerReportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentDeployRequest() (request *DescribeAgentDeployRequest) {
	request = &DescribeAgentDeployRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAgentDeploy")
	return
}

func NewDescribeAgentDeployResponse() (response *DescribeAgentDeployResponse) {
	response = &DescribeAgentDeployResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Agent部署列表
func (c *Client) DescribeAgentDeploy(request *DescribeAgentDeployRequest) (response *DescribeAgentDeployResponse, err error) {
	if request == nil {
		request = NewDescribeAgentDeployRequest()
	}
	response = NewDescribeAgentDeployResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExternalAgentRequest() (request *DescribeExternalAgentRequest) {
	request = &DescribeExternalAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeExternalAgent")
	return
}

func NewDescribeExternalAgentResponse() (response *DescribeExternalAgentResponse) {
	response = &DescribeExternalAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询腾讯云外 Agent 名单列表
func (c *Client) DescribeExternalAgent(request *DescribeExternalAgentRequest) (response *DescribeExternalAgentResponse, err error) {
	if request == nil {
		request = NewDescribeExternalAgentRequest()
	}
	response = NewDescribeExternalAgentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRuleSwitchRequest() (request *ModifyRuleSwitchRequest) {
	request = &ModifyRuleSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyRuleSwitch")
	return
}

func NewModifyRuleSwitchResponse() (response *ModifyRuleSwitchResponse) {
	response = &ModifyRuleSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 规则启用
func (c *Client) ModifyRuleSwitch(request *ModifyRuleSwitchRequest) (response *ModifyRuleSwitchResponse, err error) {
	if request == nil {
		request = NewModifyRuleSwitchRequest()
	}
	response = NewModifyRuleSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDashBoardLogRequest() (request *DescribeDashBoardLogRequest) {
	request = &DescribeDashBoardLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeDashBoardLog")
	return
}

func NewDescribeDashBoardLogResponse() (response *DescribeDashBoardLogResponse) {
	response = &DescribeDashBoardLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 概览页-统计日志
func (c *Client) DescribeDashBoardLog(request *DescribeDashBoardLogRequest) (response *DescribeDashBoardLogResponse, err error) {
	if request == nil {
		request = NewDescribeDashBoardLogRequest()
	}
	response = NewDescribeDashBoardLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMailConfigSaveRequest() (request *ModifyMailConfigSaveRequest) {
	request = &ModifyMailConfigSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyMailConfigSave")
	return
}

func NewModifyMailConfigSaveResponse() (response *ModifyMailConfigSaveResponse) {
	response = &ModifyMailConfigSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 邮件告警设置保存及测试
func (c *Client) ModifyMailConfigSave(request *ModifyMailConfigSaveRequest) (response *ModifyMailConfigSaveResponse, err error) {
	if request == nil {
		request = NewModifyMailConfigSaveRequest()
	}
	response = NewModifyMailConfigSaveResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNewProduceReportRequest() (request *CreateNewProduceReportRequest) {
	request = &CreateNewProduceReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateNewProduceReport")
	return
}

func NewCreateNewProduceReportResponse() (response *CreateNewProduceReportResponse) {
	response = &CreateNewProduceReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新的新建即时报表
func (c *Client) CreateNewProduceReport(request *CreateNewProduceReportRequest) (response *CreateNewProduceReportResponse, err error) {
	if request == nil {
		request = NewCreateNewProduceReportRequest()
	}
	response = NewCreateNewProduceReportResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogDownloadRequest() (request *ModifyLogDownloadRequest) {
	request = &ModifyLogDownloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyLogDownload")
	return
}

func NewModifyLogDownloadResponse() (response *ModifyLogDownloadResponse) {
	response = &ModifyLogDownloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载导出好的日志文件
func (c *Client) ModifyLogDownload(request *ModifyLogDownloadRequest) (response *ModifyLogDownloadResponse, err error) {
	if request == nil {
		request = NewModifyLogDownloadRequest()
	}
	response = NewModifyLogDownloadResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmSmsConfigRequest() (request *DescribeAlarmSmsConfigRequest) {
	request = &DescribeAlarmSmsConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAlarmSmsConfig")
	return
}

func NewDescribeAlarmSmsConfigResponse() (response *DescribeAlarmSmsConfigResponse) {
	response = &DescribeAlarmSmsConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取短信配置
func (c *Client) DescribeAlarmSmsConfig(request *DescribeAlarmSmsConfigRequest) (response *DescribeAlarmSmsConfigResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmSmsConfigRequest()
	}
	response = NewDescribeAlarmSmsConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcAccessRequest() (request *DescribeVpcAccessRequest) {
	request = &DescribeVpcAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeVpcAccess")
	return
}

func NewDescribeVpcAccessResponse() (response *DescribeVpcAccessResponse) {
	response = &DescribeVpcAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VPC通道列表
func (c *Client) DescribeVpcAccess(request *DescribeVpcAccessRequest) (response *DescribeVpcAccessResponse, err error) {
	if request == nil {
		request = NewDescribeVpcAccessRequest()
	}
	response = NewDescribeVpcAccessResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateLogListRequest() (request *DescribeOperateLogListRequest) {
	request = &DescribeOperateLogListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeOperateLogList")
	return
}

func NewDescribeOperateLogListResponse() (response *DescribeOperateLogListResponse) {
	response = &DescribeOperateLogListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询操作日志列表
func (c *Client) DescribeOperateLogList(request *DescribeOperateLogListRequest) (response *DescribeOperateLogListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateLogListRequest()
	}
	response = NewDescribeOperateLogListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRulesRequest() (request *DeleteRulesRequest) {
	request = &DeleteRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DeleteRules")
	return
}

func NewDeleteRulesResponse() (response *DeleteRulesResponse) {
	response = &DeleteRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 规则删除
func (c *Client) DeleteRules(request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
	if request == nil {
		request = NewDeleteRulesRequest()
	}
	response = NewDeleteRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogListRequest() (request *DescribeLogListRequest) {
	request = &DescribeLogListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeLogList")
	return
}

func NewDescribeLogListResponse() (response *DescribeLogListResponse) {
	response = &DescribeLogListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志列表信息
func (c *Client) DescribeLogList(request *DescribeLogListRequest) (response *DescribeLogListResponse, err error) {
	if request == nil {
		request = NewDescribeLogListRequest()
	}
	response = NewDescribeLogListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetsInstRequest() (request *DescribeAssetsInstRequest) {
	request = &DescribeAssetsInstRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAssetsInst")
	return
}

func NewDescribeAssetsInstResponse() (response *DescribeAssetsInstResponse) {
	response = &DescribeAssetsInstResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 概览页-查询资产审计列表
//
func (c *Client) DescribeAssetsInst(request *DescribeAssetsInstRequest) (response *DescribeAssetsInstResponse, err error) {
	if request == nil {
		request = NewDescribeAssetsInstRequest()
	}
	response = NewDescribeAssetsInstResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRulesListRequest() (request *DescribeRulesListRequest) {
	request = &DescribeRulesListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeRulesList")
	return
}

func NewDescribeRulesListResponse() (response *DescribeRulesListResponse) {
	response = &DescribeRulesListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 规则列表查询
//
func (c *Client) DescribeRulesList(request *DescribeRulesListRequest) (response *DescribeRulesListResponse, err error) {
	if request == nil {
		request = NewDescribeRulesListRequest()
	}
	response = NewDescribeRulesListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetsSaveRequest() (request *CreateAssetsSaveRequest) {
	request = &CreateAssetsSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateAssetsSave")
	return
}

func NewCreateAssetsSaveResponse() (response *CreateAssetsSaveResponse) {
	response = &CreateAssetsSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增数据资产（需要支持腾讯云外数据库）
func (c *Client) CreateAssetsSave(request *CreateAssetsSaveRequest) (response *CreateAssetsSaveResponse, err error) {
	if request == nil {
		request = NewCreateAssetsSaveRequest()
	}
	response = NewCreateAssetsSaveResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWhiteListRequest() (request *CreateWhiteListRequest) {
	request = &CreateWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateWhiteList")
	return
}

func NewCreateWhiteListResponse() (response *CreateWhiteListResponse) {
	response = &CreateWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加白名单
func (c *Client) CreateWhiteList(request *CreateWhiteListRequest) (response *CreateWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateWhiteListRequest()
	}
	response = NewCreateWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemVersionRequest() (request *DescribeSystemVersionRequest) {
	request = &DescribeSystemVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeSystemVersion")
	return
}

func NewDescribeSystemVersionResponse() (response *DescribeSystemVersionResponse) {
	response = &DescribeSystemVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户系统信息
func (c *Client) DescribeSystemVersion(request *DescribeSystemVersionRequest) (response *DescribeSystemVersionResponse, err error) {
	if request == nil {
		request = NewDescribeSystemVersionRequest()
	}
	response = NewDescribeSystemVersionResponse()
	err = c.Send(request, response)
	return
}

func NewSendAlarmEnterpriseWechatTestRequest() (request *SendAlarmEnterpriseWechatTestRequest) {
	request = &SendAlarmEnterpriseWechatTestRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "SendAlarmEnterpriseWechatTest")
	return
}

func NewSendAlarmEnterpriseWechatTestResponse() (response *SendAlarmEnterpriseWechatTestResponse) {
	response = &SendAlarmEnterpriseWechatTestResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试企业微信
func (c *Client) SendAlarmEnterpriseWechatTest(request *SendAlarmEnterpriseWechatTestRequest) (response *SendAlarmEnterpriseWechatTestResponse, err error) {
	if request == nil {
		request = NewSendAlarmEnterpriseWechatTestRequest()
	}
	response = NewSendAlarmEnterpriseWechatTestResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAgentBatchDeployRequest() (request *CreateAgentBatchDeployRequest) {
	request = &CreateAgentBatchDeployRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateAgentBatchDeploy")
	return
}

func NewCreateAgentBatchDeployResponse() (response *CreateAgentBatchDeployResponse) {
	response = &CreateAgentBatchDeployResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Agent内网在线部署
func (c *Client) CreateAgentBatchDeploy(request *CreateAgentBatchDeployRequest) (response *CreateAgentBatchDeployResponse, err error) {
	if request == nil {
		request = NewCreateAgentBatchDeployRequest()
	}
	response = NewCreateAgentBatchDeployResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCkafkaStopRequest() (request *ModifyCkafkaStopRequest) {
	request = &ModifyCkafkaStopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyCkafkaStop")
	return
}

func NewModifyCkafkaStopResponse() (response *ModifyCkafkaStopResponse) {
	response = &ModifyCkafkaStopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志类型投递关闭
func (c *Client) ModifyCkafkaStop(request *ModifyCkafkaStopRequest) (response *ModifyCkafkaStopResponse, err error) {
	if request == nil {
		request = NewModifyCkafkaStopRequest()
	}
	response = NewModifyCkafkaStopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllRulesRequest() (request *DescribeAllRulesRequest) {
	request = &DescribeAllRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAllRules")
	return
}

func NewDescribeAllRulesResponse() (response *DescribeAllRulesResponse) {
	response = &DescribeAllRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有规则
func (c *Client) DescribeAllRules(request *DescribeAllRulesRequest) (response *DescribeAllRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAllRulesRequest()
	}
	response = NewDescribeAllRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmSmsConfigSaveRequest() (request *ModifyAlarmSmsConfigSaveRequest) {
	request = &ModifyAlarmSmsConfigSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyAlarmSmsConfigSave")
	return
}

func NewModifyAlarmSmsConfigSaveResponse() (response *ModifyAlarmSmsConfigSaveResponse) {
	response = &ModifyAlarmSmsConfigSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存短信配置
func (c *Client) ModifyAlarmSmsConfigSave(request *ModifyAlarmSmsConfigSaveRequest) (response *ModifyAlarmSmsConfigSaveResponse, err error) {
	if request == nil {
		request = NewModifyAlarmSmsConfigSaveRequest()
	}
	response = NewModifyAlarmSmsConfigSaveResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCvmListRequest() (request *DescribeCvmListRequest) {
	request = &DescribeCvmListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeCvmList")
	return
}

func NewDescribeCvmListResponse() (response *DescribeCvmListResponse) {
	response = &DescribeCvmListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询CVM列表
func (c *Client) DescribeCvmList(request *DescribeCvmListRequest) (response *DescribeCvmListResponse, err error) {
	if request == nil {
		request = NewDescribeCvmListRequest()
	}
	response = NewDescribeCvmListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQpsAndStoreStatsRequest() (request *DescribeQpsAndStoreStatsRequest) {
	request = &DescribeQpsAndStoreStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeQpsAndStoreStats")
	return
}

func NewDescribeQpsAndStoreStatsResponse() (response *DescribeQpsAndStoreStatsResponse) {
	response = &DescribeQpsAndStoreStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Qos与存储状态，实现告警
func (c *Client) DescribeQpsAndStoreStats(request *DescribeQpsAndStoreStatsRequest) (response *DescribeQpsAndStoreStatsResponse, err error) {
	if request == nil {
		request = NewDescribeQpsAndStoreStatsRequest()
	}
	response = NewDescribeQpsAndStoreStatsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAgentUninstallRequest() (request *ModifyAgentUninstallRequest) {
	request = &ModifyAgentUninstallRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyAgentUninstall")
	return
}

func NewModifyAgentUninstallResponse() (response *ModifyAgentUninstallResponse) {
	response = &ModifyAgentUninstallResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 卸载 agent
func (c *Client) ModifyAgentUninstall(request *ModifyAgentUninstallRequest) (response *ModifyAgentUninstallResponse, err error) {
	if request == nil {
		request = NewModifyAgentUninstallRequest()
	}
	response = NewModifyAgentUninstallResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateRuleListRequest() (request *DescribeOperateRuleListRequest) {
	request = &DescribeOperateRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeOperateRuleList")
	return
}

func NewDescribeOperateRuleListResponse() (response *DescribeOperateRuleListResponse) {
	response = &DescribeOperateRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询行为规则配置列表
func (c *Client) DescribeOperateRuleList(request *DescribeOperateRuleListRequest) (response *DescribeOperateRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateRuleListRequest()
	}
	response = NewDescribeOperateRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleSwitchListRequest() (request *DescribeRuleSwitchListRequest) {
	request = &DescribeRuleSwitchListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeRuleSwitchList")
	return
}

func NewDescribeRuleSwitchListResponse() (response *DescribeRuleSwitchListResponse) {
	response = &DescribeRuleSwitchListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 规则启用列表查询
//
func (c *Client) DescribeRuleSwitchList(request *DescribeRuleSwitchListRequest) (response *DescribeRuleSwitchListResponse, err error) {
	if request == nil {
		request = NewDescribeRuleSwitchListRequest()
	}
	response = NewDescribeRuleSwitchListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReportMissionRequest() (request *DeleteReportMissionRequest) {
	request = &DeleteReportMissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DeleteReportMission")
	return
}

func NewDeleteReportMissionResponse() (response *DeleteReportMissionResponse) {
	response = &DeleteReportMissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除报表任务
func (c *Client) DeleteReportMission(request *DeleteReportMissionRequest) (response *DeleteReportMissionResponse, err error) {
	if request == nil {
		request = NewDeleteReportMissionRequest()
	}
	response = NewDeleteReportMissionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReportTplRequest() (request *DescribeReportTplRequest) {
	request = &DescribeReportTplRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeReportTpl")
	return
}

func NewDescribeReportTplResponse() (response *DescribeReportTplResponse) {
	response = &DescribeReportTplResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询上一次的报表配置
func (c *Client) DescribeReportTpl(request *DescribeReportTplRequest) (response *DescribeReportTplResponse, err error) {
	if request == nil {
		request = NewDescribeReportTplRequest()
	}
	response = NewDescribeReportTplResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmSettingRequest() (request *DescribeAlarmSettingRequest) {
	request = &DescribeAlarmSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAlarmSetting")
	return
}

func NewDescribeAlarmSettingResponse() (response *DescribeAlarmSettingResponse) {
	response = &DescribeAlarmSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询各类告警开关的设置
func (c *Client) DescribeAlarmSetting(request *DescribeAlarmSettingRequest) (response *DescribeAlarmSettingResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmSettingRequest()
	}
	response = NewDescribeAlarmSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupLogListRequest() (request *DescribeBackupLogListRequest) {
	request = &DescribeBackupLogListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeBackupLogList")
	return
}

func NewDescribeBackupLogListResponse() (response *DescribeBackupLogListResponse) {
	response = &DescribeBackupLogListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询备份日志列表
func (c *Client) DescribeBackupLogList(request *DescribeBackupLogListRequest) (response *DescribeBackupLogListResponse, err error) {
	if request == nil {
		request = NewDescribeBackupLogListRequest()
	}
	response = NewDescribeBackupLogListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReportMissionListRequest() (request *DescribeReportMissionListRequest) {
	request = &DescribeReportMissionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeReportMissionList")
	return
}

func NewDescribeReportMissionListResponse() (response *DescribeReportMissionListResponse) {
	response = &DescribeReportMissionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询报表任务列表
func (c *Client) DescribeReportMissionList(request *DescribeReportMissionListRequest) (response *DescribeReportMissionListResponse, err error) {
	if request == nil {
		request = NewDescribeReportMissionListRequest()
	}
	response = NewDescribeReportMissionListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRuleSaveRequest() (request *CreateRuleSaveRequest) {
	request = &CreateRuleSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateRuleSave")
	return
}

func NewCreateRuleSaveResponse() (response *CreateRuleSaveResponse) {
	response = &CreateRuleSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加规则
func (c *Client) CreateRuleSave(request *CreateRuleSaveRequest) (response *CreateRuleSaveResponse, err error) {
	if request == nil {
		request = NewCreateRuleSaveRequest()
	}
	response = NewCreateRuleSaveResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAgentStopRequest() (request *ModifyAgentStopRequest) {
	request = &ModifyAgentStopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyAgentStop")
	return
}

func NewModifyAgentStopResponse() (response *ModifyAgentStopResponse) {
	response = &ModifyAgentStopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止 agent
func (c *Client) ModifyAgentStop(request *ModifyAgentStopRequest) (response *ModifyAgentStopResponse, err error) {
	if request == nil {
		request = NewModifyAgentStopRequest()
	}
	response = NewModifyAgentStopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmRulesRequest() (request *DescribeAlarmRulesRequest) {
	request = &DescribeAlarmRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAlarmRules")
	return
}

func NewDescribeAlarmRulesResponse() (response *DescribeAlarmRulesResponse) {
	response = &DescribeAlarmRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前用户告警规则配置
func (c *Client) DescribeAlarmRules(request *DescribeAlarmRulesRequest) (response *DescribeAlarmRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmRulesRequest()
	}
	response = NewDescribeAlarmRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupSettingRequest() (request *DescribeBackupSettingRequest) {
	request = &DescribeBackupSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeBackupSetting")
	return
}

func NewDescribeBackupSettingResponse() (response *DescribeBackupSettingResponse) {
	response = &DescribeBackupSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志备份配置
func (c *Client) DescribeBackupSetting(request *DescribeBackupSettingRequest) (response *DescribeBackupSettingResponse, err error) {
	if request == nil {
		request = NewDescribeBackupSettingRequest()
	}
	response = NewDescribeBackupSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublickeyRequest() (request *DescribePublickeyRequest) {
	request = &DescribePublickeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribePublickey")
	return
}

func NewDescribePublickeyResponse() (response *DescribePublickeyResponse) {
	response = &DescribePublickeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取公钥
func (c *Client) DescribePublickey(request *DescribePublickeyRequest) (response *DescribePublickeyResponse, err error) {
	if request == nil {
		request = NewDescribePublickeyRequest()
	}
	response = NewDescribePublickeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStoreTopLimitRequest() (request *DescribeStoreTopLimitRequest) {
	request = &DescribeStoreTopLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeStoreTopLimit")
	return
}

func NewDescribeStoreTopLimitResponse() (response *DescribeStoreTopLimitResponse) {
	response = &DescribeStoreTopLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志预计可存储的天数
func (c *Client) DescribeStoreTopLimit(request *DescribeStoreTopLimitRequest) (response *DescribeStoreTopLimitResponse, err error) {
	if request == nil {
		request = NewDescribeStoreTopLimitRequest()
	}
	response = NewDescribeStoreTopLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCkafkaRouteListRequest() (request *DescribeCkafkaRouteListRequest) {
	request = &DescribeCkafkaRouteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeCkafkaRouteList")
	return
}

func NewDescribeCkafkaRouteListResponse() (response *DescribeCkafkaRouteListResponse) {
	response = &DescribeCkafkaRouteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Ckafka实例的路由信息
func (c *Client) DescribeCkafkaRouteList(request *DescribeCkafkaRouteListRequest) (response *DescribeCkafkaRouteListResponse, err error) {
	if request == nil {
		request = NewDescribeCkafkaRouteListRequest()
	}
	response = NewDescribeCkafkaRouteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCkafkaTopicListRequest() (request *DescribeCkafkaTopicListRequest) {
	request = &DescribeCkafkaTopicListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeCkafkaTopicList")
	return
}

func NewDescribeCkafkaTopicListResponse() (response *DescribeCkafkaTopicListResponse) {
	response = &DescribeCkafkaTopicListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例的主题列表
func (c *Client) DescribeCkafkaTopicList(request *DescribeCkafkaTopicListRequest) (response *DescribeCkafkaTopicListResponse, err error) {
	if request == nil {
		request = NewDescribeCkafkaTopicListRequest()
	}
	response = NewDescribeCkafkaTopicListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBackupLogListRequest() (request *DeleteBackupLogListRequest) {
	request = &DeleteBackupLogListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DeleteBackupLogList")
	return
}

func NewDeleteBackupLogListResponse() (response *DeleteBackupLogListResponse) {
	response = &DeleteBackupLogListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除备份日志
func (c *Client) DeleteBackupLogList(request *DeleteBackupLogListRequest) (response *DeleteBackupLogListResponse, err error) {
	if request == nil {
		request = NewDeleteBackupLogListRequest()
	}
	response = NewDeleteBackupLogListResponse()
	err = c.Send(request, response)
	return
}

func NewDestroycdsAuditSaasResourceRequest() (request *DestroycdsAuditSaasResourceRequest) {
	request = &DestroycdsAuditSaasResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DestroycdsAuditSaasResource")
	return
}

func NewDestroycdsAuditSaasResourceResponse() (response *DestroycdsAuditSaasResourceResponse) {
	response = &DestroycdsAuditSaasResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DestroycdsAuditSaasResource
func (c *Client) DestroycdsAuditSaasResource(request *DestroycdsAuditSaasResourceRequest) (response *DestroycdsAuditSaasResourceResponse, err error) {
	if request == nil {
		request = NewDestroycdsAuditSaasResourceRequest()
	}
	response = NewDestroycdsAuditSaasResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBehaviourRequest() (request *DescribeBehaviourRequest) {
	request = &DescribeBehaviourRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeBehaviour")
	return
}

func NewDescribeBehaviourResponse() (response *DescribeBehaviourResponse) {
	response = &DescribeBehaviourResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询操作行为分类列表
func (c *Client) DescribeBehaviour(request *DescribeBehaviourRequest) (response *DescribeBehaviourResponse, err error) {
	if request == nil {
		request = NewDescribeBehaviourRequest()
	}
	response = NewDescribeBehaviourResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogDeliveryTypeRequest() (request *DescribeLogDeliveryTypeRequest) {
	request = &DescribeLogDeliveryTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeLogDeliveryType")
	return
}

func NewDescribeLogDeliveryTypeResponse() (response *DescribeLogDeliveryTypeResponse) {
	response = &DescribeLogDeliveryTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志投递的日志类型
func (c *Client) DescribeLogDeliveryType(request *DescribeLogDeliveryTypeRequest) (response *DescribeLogDeliveryTypeResponse, err error) {
	if request == nil {
		request = NewDescribeLogDeliveryTypeRequest()
	}
	response = NewDescribeLogDeliveryTypeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCvmAssetsSaveRequest() (request *CreateCvmAssetsSaveRequest) {
	request = &CreateCvmAssetsSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateCvmAssetsSave")
	return
}

func NewCreateCvmAssetsSaveResponse() (response *CreateCvmAssetsSaveResponse) {
	response = &CreateCvmAssetsSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增 Cvm 实例
func (c *Client) CreateCvmAssetsSave(request *CreateCvmAssetsSaveRequest) (response *CreateCvmAssetsSaveResponse, err error) {
	if request == nil {
		request = NewCreateCvmAssetsSaveRequest()
	}
	response = NewCreateCvmAssetsSaveResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCkafkaStartRequest() (request *ModifyCkafkaStartRequest) {
	request = &ModifyCkafkaStartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyCkafkaStart")
	return
}

func NewModifyCkafkaStartResponse() (response *ModifyCkafkaStartResponse) {
	response = &ModifyCkafkaStartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志投递开启
func (c *Client) ModifyCkafkaStart(request *ModifyCkafkaStartRequest) (response *ModifyCkafkaStartResponse, err error) {
	if request == nil {
		request = NewModifyCkafkaStartRequest()
	}
	response = NewModifyCkafkaStartResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMailConfigRequest() (request *DescribeMailConfigRequest) {
	request = &DescribeMailConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeMailConfig")
	return
}

func NewDescribeMailConfigResponse() (response *DescribeMailConfigResponse) {
	response = &DescribeMailConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取邮件告警信息
func (c *Client) DescribeMailConfig(request *DescribeMailConfigRequest) (response *DescribeMailConfigResponse, err error) {
	if request == nil {
		request = NewDescribeMailConfigRequest()
	}
	response = NewDescribeMailConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCkafkaConfigRequest() (request *DeleteCkafkaConfigRequest) {
	request = &DeleteCkafkaConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DeleteCkafkaConfig")
	return
}

func NewDeleteCkafkaConfigResponse() (response *DeleteCkafkaConfigResponse) {
	response = &DeleteCkafkaConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消日志投递配置
func (c *Client) DeleteCkafkaConfig(request *DeleteCkafkaConfigRequest) (response *DeleteCkafkaConfigResponse, err error) {
	if request == nil {
		request = NewDeleteCkafkaConfigRequest()
	}
	response = NewDeleteCkafkaConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReportListRequest() (request *DescribeReportListRequest) {
	request = &DescribeReportListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeReportList")
	return
}

func NewDescribeReportListResponse() (response *DescribeReportListResponse) {
	response = &DescribeReportListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询报表列表
func (c *Client) DescribeReportList(request *DescribeReportListRequest) (response *DescribeReportListResponse, err error) {
	if request == nil {
		request = NewDescribeReportListRequest()
	}
	response = NewDescribeReportListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReportMissionRequest() (request *ModifyReportMissionRequest) {
	request = &ModifyReportMissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyReportMission")
	return
}

func NewModifyReportMissionResponse() (response *ModifyReportMissionResponse) {
	response = &ModifyReportMissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改报表任务
func (c *Client) ModifyReportMission(request *ModifyReportMissionRequest) (response *ModifyReportMissionResponse, err error) {
	if request == nil {
		request = NewModifyReportMissionRequest()
	}
	response = NewModifyReportMissionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRuleQuickConfigRequest() (request *ModifyRuleQuickConfigRequest) {
	request = &ModifyRuleQuickConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyRuleQuickConfig")
	return
}

func NewModifyRuleQuickConfigResponse() (response *ModifyRuleQuickConfigResponse) {
	response = &ModifyRuleQuickConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 快捷配置规则
func (c *Client) ModifyRuleQuickConfig(request *ModifyRuleQuickConfigRequest) (response *ModifyRuleQuickConfigResponse, err error) {
	if request == nil {
		request = NewModifyRuleQuickConfigRequest()
	}
	response = NewModifyRuleQuickConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRuleSaveRequest() (request *ModifyRuleSaveRequest) {
	request = &ModifyRuleSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyRuleSave")
	return
}

func NewModifyRuleSaveResponse() (response *ModifyRuleSaveResponse) {
	response = &ModifyRuleSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改规则
func (c *Client) ModifyRuleSave(request *ModifyRuleSaveRequest) (response *ModifyRuleSaveResponse, err error) {
	if request == nil {
		request = NewModifyRuleSaveRequest()
	}
	response = NewModifyRuleSaveResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOldAlarmStatusRequest() (request *DescribeOldAlarmStatusRequest) {
	request = &DescribeOldAlarmStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeOldAlarmStatus")
	return
}

func NewDescribeOldAlarmStatusResponse() (response *DescribeOldAlarmStatusResponse) {
	response = &DescribeOldAlarmStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询老的告警开启状态
func (c *Client) DescribeOldAlarmStatus(request *DescribeOldAlarmStatusRequest) (response *DescribeOldAlarmStatusResponse, err error) {
	if request == nil {
		request = NewDescribeOldAlarmStatusRequest()
	}
	response = NewDescribeOldAlarmStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetsTypeRequest() (request *DescribeAssetsTypeRequest) {
	request = &DescribeAssetsTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAssetsType")
	return
}

func NewDescribeAssetsTypeResponse() (response *DescribeAssetsTypeResponse) {
	response = &DescribeAssetsTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有数据资产类型
func (c *Client) DescribeAssetsType(request *DescribeAssetsTypeRequest) (response *DescribeAssetsTypeResponse, err error) {
	if request == nil {
		request = NewDescribeAssetsTypeRequest()
	}
	response = NewDescribeAssetsTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogDeliveryTypeRequest() (request *ModifyLogDeliveryTypeRequest) {
	request = &ModifyLogDeliveryTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyLogDeliveryType")
	return
}

func NewModifyLogDeliveryTypeResponse() (response *ModifyLogDeliveryTypeResponse) {
	response = &ModifyLogDeliveryTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改日志投递配置信息
func (c *Client) ModifyLogDeliveryType(request *ModifyLogDeliveryTypeRequest) (response *ModifyLogDeliveryTypeResponse, err error) {
	if request == nil {
		request = NewModifyLogDeliveryTypeRequest()
	}
	response = NewModifyLogDeliveryTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmRuleRequest() (request *ModifyAlarmRuleRequest) {
	request = &ModifyAlarmRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyAlarmRule")
	return
}

func NewModifyAlarmRuleResponse() (response *ModifyAlarmRuleResponse) {
	response = &ModifyAlarmRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改当前用户告警规则配置
func (c *Client) ModifyAlarmRule(request *ModifyAlarmRuleRequest) (response *ModifyAlarmRuleResponse, err error) {
	if request == nil {
		request = NewModifyAlarmRuleRequest()
	}
	response = NewModifyAlarmRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAnalysisQpsRequest() (request *DescribeAnalysisQpsRequest) {
	request = &DescribeAnalysisQpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAnalysisQps")
	return
}

func NewDescribeAnalysisQpsResponse() (response *DescribeAnalysisQpsResponse) {
	response = &DescribeAnalysisQpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 概览页-流量走势实时统计
func (c *Client) DescribeAnalysisQps(request *DescribeAnalysisQpsRequest) (response *DescribeAnalysisQpsResponse, err error) {
	if request == nil {
		request = NewDescribeAnalysisQpsRequest()
	}
	response = NewDescribeAnalysisQpsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogStoreTrendRequest() (request *DescribeLogStoreTrendRequest) {
	request = &DescribeLogStoreTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeLogStoreTrend")
	return
}

func NewDescribeLogStoreTrendResponse() (response *DescribeLogStoreTrendResponse) {
	response = &DescribeLogStoreTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 概览页-日志存储走势
func (c *Client) DescribeLogStoreTrend(request *DescribeLogStoreTrendRequest) (response *DescribeLogStoreTrendResponse, err error) {
	if request == nil {
		request = NewDescribeLogStoreTrendRequest()
	}
	response = NewDescribeLogStoreTrendResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllDbCountRequest() (request *DescribeAllDbCountRequest) {
	request = &DescribeAllDbCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAllDbCount")
	return
}

func NewDescribeAllDbCountResponse() (response *DescribeAllDbCountResponse) {
	response = &DescribeAllDbCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各类型数据库个数
func (c *Client) DescribeAllDbCount(request *DescribeAllDbCountRequest) (response *DescribeAllDbCountResponse, err error) {
	if request == nil {
		request = NewDescribeAllDbCountRequest()
	}
	response = NewDescribeAllDbCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDashBoardAnalysisRequest() (request *DescribeDashBoardAnalysisRequest) {
	request = &DescribeDashBoardAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeDashBoardAnalysis")
	return
}

func NewDescribeDashBoardAnalysisResponse() (response *DescribeDashBoardAnalysisResponse) {
	response = &DescribeDashBoardAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 概览页-态势分析
func (c *Client) DescribeDashBoardAnalysis(request *DescribeDashBoardAnalysisRequest) (response *DescribeDashBoardAnalysisResponse, err error) {
	if request == nil {
		request = NewDescribeDashBoardAnalysisRequest()
	}
	response = NewDescribeDashBoardAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRestoreLogTaskRequest() (request *ModifyRestoreLogTaskRequest) {
	request = &ModifyRestoreLogTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyRestoreLogTask")
	return
}

func NewModifyRestoreLogTaskResponse() (response *ModifyRestoreLogTaskResponse) {
	response = &ModifyRestoreLogTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恢复备份日志
func (c *Client) ModifyRestoreLogTask(request *ModifyRestoreLogTaskRequest) (response *ModifyRestoreLogTaskResponse, err error) {
	if request == nil {
		request = NewModifyRestoreLogTaskRequest()
	}
	response = NewModifyRestoreLogTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAgentDownloadRequest() (request *CreateAgentDownloadRequest) {
	request = &CreateAgentDownloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateAgentDownload")
	return
}

func NewCreateAgentDownloadResponse() (response *CreateAgentDownloadResponse) {
	response = &CreateAgentDownloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Agent部署-下载Agent
func (c *Client) CreateAgentDownload(request *CreateAgentDownloadRequest) (response *CreateAgentDownloadResponse, err error) {
	if request == nil {
		request = NewCreateAgentDownloadRequest()
	}
	response = NewCreateAgentDownloadResponse()
	err = c.Send(request, response)
	return
}

func NewSendCkafkaTestRequest() (request *SendCkafkaTestRequest) {
	request = &SendCkafkaTestRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "SendCkafkaTest")
	return
}

func NewSendCkafkaTestResponse() (response *SendCkafkaTestResponse) {
	response = &SendCkafkaTestResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户Ckafka联通性测试
func (c *Client) SendCkafkaTest(request *SendCkafkaTestRequest) (response *SendCkafkaTestResponse, err error) {
	if request == nil {
		request = NewSendCkafkaTestRequest()
	}
	response = NewSendCkafkaTestResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReportRequest() (request *DeleteReportRequest) {
	request = &DeleteReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DeleteReport")
	return
}

func NewDeleteReportResponse() (response *DeleteReportResponse) {
	response = &DeleteReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除报表
func (c *Client) DeleteReport(request *DeleteReportRequest) (response *DeleteReportResponse, err error) {
	if request == nil {
		request = NewDeleteReportRequest()
	}
	response = NewDeleteReportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionListsRequest() (request *DescribeRegionListsRequest) {
	request = &DescribeRegionListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeRegionLists")
	return
}

func NewDescribeRegionListsResponse() (response *DescribeRegionListsResponse) {
	response = &DescribeRegionListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过接口获取不同地域支持的地域子集
func (c *Client) DescribeRegionLists(request *DescribeRegionListsRequest) (response *DescribeRegionListsResponse, err error) {
	if request == nil {
		request = NewDescribeRegionListsRequest()
	}
	response = NewDescribeRegionListsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAgentDeployRequest() (request *CreateAgentDeployRequest) {
	request = &CreateAgentDeployRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateAgentDeploy")
	return
}

func NewCreateAgentDeployResponse() (response *CreateAgentDeployResponse) {
	response = &CreateAgentDeployResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Agent部署-Linux 在线部署
func (c *Client) CreateAgentDeploy(request *CreateAgentDeployRequest) (response *CreateAgentDeployResponse, err error) {
	if request == nil {
		request = NewCreateAgentDeployRequest()
	}
	response = NewCreateAgentDeployResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFieldConfigRequest() (request *DescribeFieldConfigRequest) {
	request = &DescribeFieldConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeFieldConfig")
	return
}

func NewDescribeFieldConfigResponse() (response *DescribeFieldConfigResponse) {
	response = &DescribeFieldConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取规则字段及类型
func (c *Client) DescribeFieldConfig(request *DescribeFieldConfigRequest) (response *DescribeFieldConfigResponse, err error) {
	if request == nil {
		request = NewDescribeFieldConfigRequest()
	}
	response = NewDescribeFieldConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAssetsRequest() (request *DeleteAssetsRequest) {
	request = &DeleteAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DeleteAssets")
	return
}

func NewDeleteAssetsResponse() (response *DeleteAssetsResponse) {
	response = &DeleteAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除数据资产
func (c *Client) DeleteAssets(request *DeleteAssetsRequest) (response *DeleteAssetsResponse, err error) {
	if request == nil {
		request = NewDeleteAssetsRequest()
	}
	response = NewDeleteAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSessionListRequest() (request *DescribeSessionListRequest) {
	request = &DescribeSessionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeSessionList")
	return
}

func NewDescribeSessionListResponse() (response *DescribeSessionListResponse) {
	response = &DescribeSessionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询审计会话列表信息
func (c *Client) DescribeSessionList(request *DescribeSessionListRequest) (response *DescribeSessionListResponse, err error) {
	if request == nil {
		request = NewDescribeSessionListRequest()
	}
	response = NewDescribeSessionListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceStateRequest() (request *DescribeServiceStateRequest) {
	request = &DescribeServiceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeServiceState")
	return
}

func NewDescribeServiceStateResponse() (response *DescribeServiceStateResponse) {
	response = &DescribeServiceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户购买信息
func (c *Client) DescribeServiceState(request *DescribeServiceStateRequest) (response *DescribeServiceStateResponse, err error) {
	if request == nil {
		request = NewDescribeServiceStateRequest()
	}
	response = NewDescribeServiceStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomerAdditionRequest() (request *DescribeCustomerAdditionRequest) {
	request = &DescribeCustomerAdditionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeCustomerAddition")
	return
}

func NewDescribeCustomerAdditionResponse() (response *DescribeCustomerAdditionResponse) {
	response = &DescribeCustomerAdditionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户扩展信息
func (c *Client) DescribeCustomerAddition(request *DescribeCustomerAdditionRequest) (response *DescribeCustomerAdditionResponse, err error) {
	if request == nil {
		request = NewDescribeCustomerAdditionRequest()
	}
	response = NewDescribeCustomerAdditionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogTypeConfigListRequest() (request *DescribeLogTypeConfigListRequest) {
	request = &DescribeLogTypeConfigListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeLogTypeConfigList")
	return
}

func NewDescribeLogTypeConfigListResponse() (response *DescribeLogTypeConfigListResponse) {
	response = &DescribeLogTypeConfigListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

//  查询租户日志投递配置
func (c *Client) DescribeLogTypeConfigList(request *DescribeLogTypeConfigListRequest) (response *DescribeLogTypeConfigListResponse, err error) {
	if request == nil {
		request = NewDescribeLogTypeConfigListRequest()
	}
	response = NewDescribeLogTypeConfigListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmSettingSaveRequest() (request *ModifyAlarmSettingSaveRequest) {
	request = &ModifyAlarmSettingSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyAlarmSettingSave")
	return
}

func NewModifyAlarmSettingSaveResponse() (response *ModifyAlarmSettingSaveResponse) {
	response = &ModifyAlarmSettingSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存告警开关设置
func (c *Client) ModifyAlarmSettingSave(request *ModifyAlarmSettingSaveRequest) (response *ModifyAlarmSettingSaveResponse, err error) {
	if request == nil {
		request = NewModifyAlarmSettingSaveRequest()
	}
	response = NewModifyAlarmSettingSaveResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCkafkaSaveRequest() (request *ModifyCkafkaSaveRequest) {
	request = &ModifyCkafkaSaveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyCkafkaSave")
	return
}

func NewModifyCkafkaSaveResponse() (response *ModifyCkafkaSaveResponse) {
	response = &ModifyCkafkaSaveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户Ckafka配置保存
func (c *Client) ModifyCkafkaSave(request *ModifyCkafkaSaveRequest) (response *ModifyCkafkaSaveResponse, err error) {
	if request == nil {
		request = NewModifyCkafkaSaveRequest()
	}
	response = NewModifyCkafkaSaveResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmEnterpriseWechatConfigRequest() (request *ModifyAlarmEnterpriseWechatConfigRequest) {
	request = &ModifyAlarmEnterpriseWechatConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyAlarmEnterpriseWechatConfig")
	return
}

func NewModifyAlarmEnterpriseWechatConfigResponse() (response *ModifyAlarmEnterpriseWechatConfigResponse) {
	response = &ModifyAlarmEnterpriseWechatConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存企业微信
func (c *Client) ModifyAlarmEnterpriseWechatConfig(request *ModifyAlarmEnterpriseWechatConfigRequest) (response *ModifyAlarmEnterpriseWechatConfigResponse, err error) {
	if request == nil {
		request = NewModifyAlarmEnterpriseWechatConfigRequest()
	}
	response = NewModifyAlarmEnterpriseWechatConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAgentStartRequest() (request *ModifyAgentStartRequest) {
	request = &ModifyAgentStartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyAgentStart")
	return
}

func NewModifyAgentStartResponse() (response *ModifyAgentStartResponse) {
	response = &ModifyAgentStartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动 agent
func (c *Client) ModifyAgentStart(request *ModifyAgentStartRequest) (response *ModifyAgentStartResponse, err error) {
	if request == nil {
		request = NewModifyAgentStartRequest()
	}
	response = NewModifyAgentStartResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcPrivateLinkRequest() (request *DeleteVpcPrivateLinkRequest) {
	request = &DeleteVpcPrivateLinkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DeleteVpcPrivateLink")
	return
}

func NewDeleteVpcPrivateLinkResponse() (response *DeleteVpcPrivateLinkResponse) {
	response = &DeleteVpcPrivateLinkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Vpc通道
func (c *Client) DeleteVpcPrivateLink(request *DeleteVpcPrivateLinkRequest) (response *DeleteVpcPrivateLinkResponse, err error) {
	if request == nil {
		request = NewDeleteVpcPrivateLinkRequest()
	}
	response = NewDeleteVpcPrivateLinkResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentListRequest() (request *DescribeAgentListRequest) {
	request = &DescribeAgentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAgentList")
	return
}

func NewDescribeAgentListResponse() (response *DescribeAgentListResponse) {
	response = &DescribeAgentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Agent列表
func (c *Client) DescribeAgentList(request *DescribeAgentListRequest) (response *DescribeAgentListResponse, err error) {
	if request == nil {
		request = NewDescribeAgentListRequest()
	}
	response = NewDescribeAgentListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetsVersionsRequest() (request *DescribeAssetsVersionsRequest) {
	request = &DescribeAssetsVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAssetsVersions")
	return
}

func NewDescribeAssetsVersionsResponse() (response *DescribeAssetsVersionsResponse) {
	response = &DescribeAssetsVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有数据资产类型对应的版本列表
func (c *Client) DescribeAssetsVersions(request *DescribeAssetsVersionsRequest) (response *DescribeAssetsVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeAssetsVersionsRequest()
	}
	response = NewDescribeAssetsVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomerGuideOpenRequest() (request *ModifyCustomerGuideOpenRequest) {
	request = &ModifyCustomerGuideOpenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyCustomerGuideOpen")
	return
}

func NewModifyCustomerGuideOpenResponse() (response *ModifyCustomerGuideOpenResponse) {
	response = &ModifyCustomerGuideOpenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户引导状态
func (c *Client) ModifyCustomerGuideOpen(request *ModifyCustomerGuideOpenRequest) (response *ModifyCustomerGuideOpenResponse, err error) {
	if request == nil {
		request = NewModifyCustomerGuideOpenRequest()
	}
	response = NewModifyCustomerGuideOpenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDashBoardRiskRequest() (request *DescribeDashBoardRiskRequest) {
	request = &DescribeDashBoardRiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeDashBoardRisk")
	return
}

func NewDescribeDashBoardRiskResponse() (response *DescribeDashBoardRiskResponse) {
	response = &DescribeDashBoardRiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 概览页-风险占比统计
func (c *Client) DescribeDashBoardRisk(request *DescribeDashBoardRiskRequest) (response *DescribeDashBoardRiskResponse, err error) {
	if request == nil {
		request = NewDescribeDashBoardRiskRequest()
	}
	response = NewDescribeDashBoardRiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNewReportTplRequest() (request *DescribeNewReportTplRequest) {
	request = &DescribeNewReportTplRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeNewReportTpl")
	return
}

func NewDescribeNewReportTplResponse() (response *DescribeNewReportTplResponse) {
	response = &DescribeNewReportTplResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新的查询上一次的报表配置
//
func (c *Client) DescribeNewReportTpl(request *DescribeNewReportTplRequest) (response *DescribeNewReportTplResponse, err error) {
	if request == nil {
		request = NewDescribeNewReportTplRequest()
	}
	response = NewDescribeNewReportTplResponse()
	err = c.Send(request, response)
	return
}

func NewCreatcdsAuditSaasResourceRequest() (request *CreatcdsAuditSaasResourceRequest) {
	request = &CreatcdsAuditSaasResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreatcdsAuditSaasResource")
	return
}

func NewCreatcdsAuditSaasResourceResponse() (response *CreatcdsAuditSaasResourceResponse) {
	response = &CreatcdsAuditSaasResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreatcdsAuditSaasResource
func (c *Client) CreatcdsAuditSaasResource(request *CreatcdsAuditSaasResourceRequest) (response *CreatcdsAuditSaasResourceResponse, err error) {
	if request == nil {
		request = NewCreatcdsAuditSaasResourceRequest()
	}
	response = NewCreatcdsAuditSaasResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserCkafkaInstanceListRequest() (request *DescribeUserCkafkaInstanceListRequest) {
	request = &DescribeUserCkafkaInstanceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeUserCkafkaInstanceList")
	return
}

func NewDescribeUserCkafkaInstanceListResponse() (response *DescribeUserCkafkaInstanceListResponse) {
	response = &DescribeUserCkafkaInstanceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户ckafka实例列表
func (c *Client) DescribeUserCkafkaInstanceList(request *DescribeUserCkafkaInstanceListRequest) (response *DescribeUserCkafkaInstanceListResponse, err error) {
	if request == nil {
		request = NewDescribeUserCkafkaInstanceListRequest()
	}
	response = NewDescribeUserCkafkaInstanceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetsListRequest() (request *DescribeAssetsListRequest) {
	request = &DescribeAssetsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeAssetsList")
	return
}

func NewDescribeAssetsListResponse() (response *DescribeAssetsListResponse) {
	response = &DescribeAssetsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产列表
func (c *Client) DescribeAssetsList(request *DescribeAssetsListRequest) (response *DescribeAssetsListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetsListRequest()
	}
	response = NewDescribeAssetsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskListRequest() (request *DescribeRiskListRequest) {
	request = &DescribeRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DescribeRiskList")
	return
}

func NewDescribeRiskListResponse() (response *DescribeRiskListResponse) {
	response = &DescribeRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询审计风险列表信息
func (c *Client) DescribeRiskList(request *DescribeRiskListRequest) (response *DescribeRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskListRequest()
	}
	response = NewDescribeRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewInquireModifycdsAuditSaasResourceRequest() (request *InquireModifycdsAuditSaasResourceRequest) {
	request = &InquireModifycdsAuditSaasResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "InquireModifycdsAuditSaasResource")
	return
}

func NewInquireModifycdsAuditSaasResourceResponse() (response *InquireModifycdsAuditSaasResourceResponse) {
	response = &InquireModifycdsAuditSaasResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// InquireModifycdsAuditSaasResource
func (c *Client) InquireModifycdsAuditSaasResource(request *InquireModifycdsAuditSaasResourceRequest) (response *InquireModifycdsAuditSaasResourceResponse, err error) {
	if request == nil {
		request = NewInquireModifycdsAuditSaasResourceRequest()
	}
	response = NewInquireModifycdsAuditSaasResourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBackupSettingRequest() (request *ModifyBackupSettingRequest) {
	request = &ModifyBackupSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "ModifyBackupSetting")
	return
}

func NewModifyBackupSettingResponse() (response *ModifyBackupSettingResponse) {
	response = &ModifyBackupSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改日志备份设置
func (c *Client) ModifyBackupSetting(request *ModifyBackupSettingRequest) (response *ModifyBackupSettingResponse, err error) {
	if request == nil {
		request = NewModifyBackupSettingRequest()
	}
	response = NewModifyBackupSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAgentRequest() (request *DeleteAgentRequest) {
	request = &DeleteAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "DeleteAgent")
	return
}

func NewDeleteAgentResponse() (response *DeleteAgentResponse) {
	response = &DeleteAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除 agent
func (c *Client) DeleteAgent(request *DeleteAgentRequest) (response *DeleteAgentResponse, err error) {
	if request == nil {
		request = NewDeleteAgentRequest()
	}
	response = NewDeleteAgentResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProduceReportRequest() (request *CreateProduceReportRequest) {
	request = &CreateProduceReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateProduceReport")
	return
}

func NewCreateProduceReportResponse() (response *CreateProduceReportResponse) {
	response = &CreateProduceReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建即时报表
func (c *Client) CreateProduceReport(request *CreateProduceReportRequest) (response *CreateProduceReportResponse, err error) {
	if request == nil {
		request = NewCreateProduceReportRequest()
	}
	response = NewCreateProduceReportResponse()
	err = c.Send(request, response)
	return
}

func NewCreateReportPdfRequest() (request *CreateReportPdfRequest) {
	request = &CreateReportPdfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cds", APIVersion, "CreateReportPdf")
	return
}

func NewCreateReportPdfResponse() (response *CreateReportPdfResponse) {
	response = &CreateReportPdfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载报表PDF
func (c *Client) CreateReportPdf(request *CreateReportPdfRequest) (response *CreateReportPdfResponse, err error) {
	if request == nil {
		request = NewCreateReportPdfRequest()
	}
	response = NewCreateReportPdfResponse()
	err = c.Send(request, response)
	return
}
