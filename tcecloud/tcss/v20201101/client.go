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

package v20201101

import (
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-11-01"

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

func NewDeleteSearchTemplateRequest() (request *DeleteSearchTemplateRequest) {
	request = &DeleteSearchTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteSearchTemplate")
	return
}

func NewDeleteSearchTemplateResponse() (response *DeleteSearchTemplateResponse) {
	response = &DeleteSearchTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除检索模板
func (c *Client) DeleteSearchTemplate(request *DeleteSearchTemplateRequest) (response *DeleteSearchTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteSearchTemplateRequest()
	}
	response = NewDeleteSearchTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulScanInfoRequest() (request *DescribeVulScanInfoRequest) {
	request = &DescribeVulScanInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulScanInfo")
	return
}

func NewDescribeVulScanInfoResponse() (response *DescribeVulScanInfoResponse) {
	response = &DescribeVulScanInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞扫描任务信息
func (c *Client) DescribeVulScanInfo(request *DescribeVulScanInfoRequest) (response *DescribeVulScanInfoResponse, err error) {
	if request == nil {
		request = NewDescribeVulScanInfoRequest()
	}
	response = NewDescribeVulScanInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateReverseShellRegexpWhiteListExportJobRequest() (request *CreateReverseShellRegexpWhiteListExportJobRequest) {
	request = &CreateReverseShellRegexpWhiteListExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateReverseShellRegexpWhiteListExportJob")
	return
}

func NewCreateReverseShellRegexpWhiteListExportJobResponse() (response *CreateReverseShellRegexpWhiteListExportJobResponse) {
	response = &CreateReverseShellRegexpWhiteListExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建反弹shell正则白名单导出任务
func (c *Client) CreateReverseShellRegexpWhiteListExportJob(request *CreateReverseShellRegexpWhiteListExportJobRequest) (response *CreateReverseShellRegexpWhiteListExportJobResponse, err error) {
	if request == nil {
		request = NewCreateReverseShellRegexpWhiteListExportJobRequest()
	}
	response = NewCreateReverseShellRegexpWhiteListExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceAssetPolicyItemListRequest() (request *DescribeComplianceAssetPolicyItemListRequest) {
	request = &DescribeComplianceAssetPolicyItemListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceAssetPolicyItemList")
	return
}

func NewDescribeComplianceAssetPolicyItemListResponse() (response *DescribeComplianceAssetPolicyItemListResponse) {
	response = &DescribeComplianceAssetPolicyItemListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某资产下的检测项列表
func (c *Client) DescribeComplianceAssetPolicyItemList(request *DescribeComplianceAssetPolicyItemListRequest) (response *DescribeComplianceAssetPolicyItemListResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceAssetPolicyItemListRequest()
	}
	response = NewDescribeComplianceAssetPolicyItemListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulScanTaskRequest() (request *CreateVulScanTaskRequest) {
	request = &CreateVulScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateVulScanTask")
	return
}

func NewCreateVulScanTaskResponse() (response *CreateVulScanTaskResponse) {
	response = &CreateVulScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建漏洞扫描任务
func (c *Client) CreateVulScanTask(request *CreateVulScanTaskRequest) (response *CreateVulScanTaskResponse, err error) {
	if request == nil {
		request = NewCreateVulScanTaskRequest()
	}
	response = NewCreateVulScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeSafeStateRequest() (request *DescribeEscapeSafeStateRequest) {
	request = &DescribeEscapeSafeStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeSafeState")
	return
}

func NewDescribeEscapeSafeStateResponse() (response *DescribeEscapeSafeStateResponse) {
	response = &DescribeEscapeSafeStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeEscapeSafeState 查询容器逃逸安全状态
func (c *Client) DescribeEscapeSafeState(request *DescribeEscapeSafeStateRequest) (response *DescribeEscapeSafeStateResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeSafeStateRequest()
	}
	response = NewDescribeEscapeSafeStateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyK8sApiAbnormalRuleInfoRequest() (request *ModifyK8sApiAbnormalRuleInfoRequest) {
	request = &ModifyK8sApiAbnormalRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyK8sApiAbnormalRuleInfo")
	return
}

func NewModifyK8sApiAbnormalRuleInfoResponse() (response *ModifyK8sApiAbnormalRuleInfoResponse) {
	response = &ModifyK8sApiAbnormalRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改k8sapi异常规则信息
func (c *Client) ModifyK8sApiAbnormalRuleInfo(request *ModifyK8sApiAbnormalRuleInfoRequest) (response *ModifyK8sApiAbnormalRuleInfoResponse, err error) {
	if request == nil {
		request = NewModifyK8sApiAbnormalRuleInfoRequest()
	}
	response = NewModifyK8sApiAbnormalRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateK8sApiAbnormalRuleExportJobRequest() (request *CreateK8sApiAbnormalRuleExportJobRequest) {
	request = &CreateK8sApiAbnormalRuleExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateK8sApiAbnormalRuleExportJob")
	return
}

func NewCreateK8sApiAbnormalRuleExportJobResponse() (response *CreateK8sApiAbnormalRuleExportJobResponse) {
	response = &CreateK8sApiAbnormalRuleExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建k8sApi异常规则导出任务
func (c *Client) CreateK8sApiAbnormalRuleExportJob(request *CreateK8sApiAbnormalRuleExportJobRequest) (response *CreateK8sApiAbnormalRuleExportJobResponse, err error) {
	if request == nil {
		request = NewCreateK8sApiAbnormalRuleExportJobRequest()
	}
	response = NewCreateK8sApiAbnormalRuleExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRiskDnsEventStatusRequest() (request *ModifyRiskDnsEventStatusRequest) {
	request = &ModifyRiskDnsEventStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyRiskDnsEventStatus")
	return
}

func NewModifyRiskDnsEventStatusResponse() (response *ModifyRiskDnsEventStatusResponse) {
	response = &ModifyRiskDnsEventStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑恶意请求事件状态
func (c *Client) ModifyRiskDnsEventStatus(request *ModifyRiskDnsEventStatusRequest) (response *ModifyRiskDnsEventStatusResponse, err error) {
	if request == nil {
		request = NewModifyRiskDnsEventStatusRequest()
	}
	response = NewModifyRiskDnsEventStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSupportRegionsRequest() (request *DescribeSupportRegionsRequest) {
	request = &DescribeSupportRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSupportRegions")
	return
}

func NewDescribeSupportRegionsResponse() (response *DescribeSupportRegionsResponse) {
	response = &DescribeSupportRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持的地域
func (c *Client) DescribeSupportRegions(request *DescribeSupportRegionsRequest) (response *DescribeSupportRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeSupportRegionsRequest()
	}
	response = NewDescribeSupportRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeEventInfoRequest() (request *DescribeEscapeEventInfoRequest) {
	request = &DescribeEscapeEventInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventInfo")
	return
}

func NewDescribeEscapeEventInfoResponse() (response *DescribeEscapeEventInfoResponse) {
	response = &DescribeEscapeEventInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeEscapeEventInfo 查询容器逃逸事件列表
func (c *Client) DescribeEscapeEventInfo(request *DescribeEscapeEventInfoRequest) (response *DescribeEscapeEventInfoResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeEventInfoRequest()
	}
	response = NewDescribeEscapeEventInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalRuleScopeListRequest() (request *DescribeK8sApiAbnormalRuleScopeListRequest) {
	request = &DescribeK8sApiAbnormalRuleScopeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalRuleScopeList")
	return
}

func NewDescribeK8sApiAbnormalRuleScopeListResponse() (response *DescribeK8sApiAbnormalRuleScopeListResponse) {
	response = &DescribeK8sApiAbnormalRuleScopeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8sapi 异常规则中范围列表
func (c *Client) DescribeK8sApiAbnormalRuleScopeList(request *DescribeK8sApiAbnormalRuleScopeListRequest) (response *DescribeK8sApiAbnormalRuleScopeListResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalRuleScopeListRequest()
	}
	response = NewDescribeK8sApiAbnormalRuleScopeListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyContainerNetStatusRequest() (request *ModifyContainerNetStatusRequest) {
	request = &ModifyContainerNetStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyContainerNetStatus")
	return
}

func NewModifyContainerNetStatusResponse() (response *ModifyContainerNetStatusResponse) {
	response = &ModifyContainerNetStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 隔离容器网络状态
func (c *Client) ModifyContainerNetStatus(request *ModifyContainerNetStatusRequest) (response *ModifyContainerNetStatusResponse, err error) {
	if request == nil {
		request = NewModifyContainerNetStatusRequest()
	}
	response = NewModifyContainerNetStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallPolicyYamlDetailRequest() (request *DescribeNetworkFirewallPolicyYamlDetailRequest) {
	request = &DescribeNetworkFirewallPolicyYamlDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPolicyYamlDetail")
	return
}

func NewDescribeNetworkFirewallPolicyYamlDetailResponse() (response *DescribeNetworkFirewallPolicyYamlDetailResponse) {
	response = &DescribeNetworkFirewallPolicyYamlDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络集群查看Yaml网络策略详情
func (c *Client) DescribeNetworkFirewallPolicyYamlDetail(request *DescribeNetworkFirewallPolicyYamlDetailRequest) (response *DescribeNetworkFirewallPolicyYamlDetailResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallPolicyYamlDetailRequest()
	}
	response = NewDescribeNetworkFirewallPolicyYamlDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetHostListRequest() (request *DescribeAssetHostListRequest) {
	request = &DescribeAssetHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetHostList")
	return
}

func NewDescribeAssetHostListResponse() (response *DescribeAssetHostListResponse) {
	response = &DescribeAssetHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全搜索查询主机列表
func (c *Client) DescribeAssetHostList(request *DescribeAssetHostListRequest) (response *DescribeAssetHostListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetHostListRequest()
	}
	response = NewDescribeAssetHostListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkTopologyFlowSwitchListRequest() (request *DescribeNetworkTopologyFlowSwitchListRequest) {
	request = &DescribeNetworkTopologyFlowSwitchListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkTopologyFlowSwitchList")
	return
}

func NewDescribeNetworkTopologyFlowSwitchListResponse() (response *DescribeNetworkTopologyFlowSwitchListResponse) {
	response = &DescribeNetworkTopologyFlowSwitchListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络拓扑集群流量开关列表
func (c *Client) DescribeNetworkTopologyFlowSwitchList(request *DescribeNetworkTopologyFlowSwitchListRequest) (response *DescribeNetworkTopologyFlowSwitchListResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkTopologyFlowSwitchListRequest()
	}
	response = NewDescribeNetworkTopologyFlowSwitchListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteK8sApiAbnormalRuleRequest() (request *DeleteK8sApiAbnormalRuleRequest) {
	request = &DeleteK8sApiAbnormalRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteK8sApiAbnormalRule")
	return
}

func NewDeleteK8sApiAbnormalRuleResponse() (response *DeleteK8sApiAbnormalRuleResponse) {
	response = &DeleteK8sApiAbnormalRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除k8sapi异常事件规则
func (c *Client) DeleteK8sApiAbnormalRule(request *DeleteK8sApiAbnormalRuleRequest) (response *DeleteK8sApiAbnormalRuleResponse, err error) {
	if request == nil {
		request = NewDeleteK8sApiAbnormalRuleRequest()
	}
	response = NewDeleteK8sApiAbnormalRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellEventsExportRequest() (request *DescribeReverseShellEventsExportRequest) {
	request = &DescribeReverseShellEventsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellEventsExport")
	return
}

func NewDescribeReverseShellEventsExportResponse() (response *DescribeReverseShellEventsExportResponse) {
	response = &DescribeReverseShellEventsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时反弹shell事件列表信息导出
func (c *Client) DescribeReverseShellEventsExport(request *DescribeReverseShellEventsExportRequest) (response *DescribeReverseShellEventsExportResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellEventsExportRequest()
	}
	response = NewDescribeReverseShellEventsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageScanStatusRequest() (request *DescribeAssetImageScanStatusRequest) {
	request = &DescribeAssetImageScanStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageScanStatus")
	return
}

func NewDescribeAssetImageScanStatusResponse() (response *DescribeAssetImageScanStatusResponse) {
	response = &DescribeAssetImageScanStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全查询镜像扫描状态
func (c *Client) DescribeAssetImageScanStatus(request *DescribeAssetImageScanStatusRequest) (response *DescribeAssetImageScanStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageScanStatusRequest()
	}
	response = NewDescribeAssetImageScanStatusResponse()
	err = c.Send(request, response)
	return
}

func NewAddAndPublishNetworkFirewallPolicyYamlDetailRequest() (request *AddAndPublishNetworkFirewallPolicyYamlDetailRequest) {
	request = &AddAndPublishNetworkFirewallPolicyYamlDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddAndPublishNetworkFirewallPolicyYamlDetail")
	return
}

func NewAddAndPublishNetworkFirewallPolicyYamlDetailResponse() (response *AddAndPublishNetworkFirewallPolicyYamlDetailResponse) {
	response = &AddAndPublishNetworkFirewallPolicyYamlDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建Yaml网络策略并发布任务
func (c *Client) AddAndPublishNetworkFirewallPolicyYamlDetail(request *AddAndPublishNetworkFirewallPolicyYamlDetailRequest) (response *AddAndPublishNetworkFirewallPolicyYamlDetailResponse, err error) {
	if request == nil {
		request = NewAddAndPublishNetworkFirewallPolicyYamlDetailRequest()
	}
	response = NewAddAndPublishNetworkFirewallPolicyYamlDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryListExportRequest() (request *DescribeAssetImageRegistryListExportRequest) {
	request = &DescribeAssetImageRegistryListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryListExport")
	return
}

func NewDescribeAssetImageRegistryListExportResponse() (response *DescribeAssetImageRegistryListExportResponse) {
	response = &DescribeAssetImageRegistryListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库镜像列表导出
func (c *Client) DescribeAssetImageRegistryListExport(request *DescribeAssetImageRegistryListExportRequest) (response *DescribeAssetImageRegistryListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryListExportRequest()
	}
	response = NewDescribeAssetImageRegistryListExportResponse()
	err = c.Send(request, response)
	return
}

func NewModifyK8sApiAbnormalEventStatusRequest() (request *ModifyK8sApiAbnormalEventStatusRequest) {
	request = &ModifyK8sApiAbnormalEventStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyK8sApiAbnormalEventStatus")
	return
}

func NewModifyK8sApiAbnormalEventStatusResponse() (response *ModifyK8sApiAbnormalEventStatusResponse) {
	response = &ModifyK8sApiAbnormalEventStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改k8sapi异常事件状态
func (c *Client) ModifyK8sApiAbnormalEventStatus(request *ModifyK8sApiAbnormalEventStatusRequest) (response *ModifyK8sApiAbnormalEventStatusResponse, err error) {
	if request == nil {
		request = NewModifyK8sApiAbnormalEventStatusRequest()
	}
	response = NewModifyK8sApiAbnormalEventStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePurchaseStateInfoRequest() (request *DescribePurchaseStateInfoRequest) {
	request = &DescribePurchaseStateInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribePurchaseStateInfo")
	return
}

func NewDescribePurchaseStateInfoResponse() (response *DescribePurchaseStateInfoResponse) {
	response = &DescribePurchaseStateInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribePurchaseStateInfo 查询容器安全服务已购买信息
func (c *Client) DescribePurchaseStateInfo(request *DescribePurchaseStateInfoRequest) (response *DescribePurchaseStateInfoResponse, err error) {
	if request == nil {
		request = NewDescribePurchaseStateInfoRequest()
	}
	response = NewDescribePurchaseStateInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogAlertMsgRequest() (request *DescribeSecLogAlertMsgRequest) {
	request = &DescribeSecLogAlertMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogAlertMsg")
	return
}

func NewDescribeSecLogAlertMsgResponse() (response *DescribeSecLogAlertMsgResponse) {
	response = &DescribeSecLogAlertMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志告警信息
func (c *Client) DescribeSecLogAlertMsg(request *DescribeSecLogAlertMsgRequest) (response *DescribeSecLogAlertMsgResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogAlertMsgRequest()
	}
	response = NewDescribeSecLogAlertMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyEventListRequest() (request *DescribeImageDenyEventListRequest) {
	request = &DescribeImageDenyEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageDenyEventList")
	return
}

func NewDescribeImageDenyEventListResponse() (response *DescribeImageDenyEventListResponse) {
	response = &DescribeImageDenyEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截事件列表
func (c *Client) DescribeImageDenyEventList(request *DescribeImageDenyEventListRequest) (response *DescribeImageDenyEventListResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyEventListRequest()
	}
	response = NewDescribeImageDenyEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSuperNodeInstallTaskListRequest() (request *DescribeSuperNodeInstallTaskListRequest) {
	request = &DescribeSuperNodeInstallTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSuperNodeInstallTaskList")
	return
}

func NewDescribeSuperNodeInstallTaskListResponse() (response *DescribeSuperNodeInstallTaskListResponse) {
	response = &DescribeSuperNodeInstallTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询超级节点安装任务列表
func (c *Client) DescribeSuperNodeInstallTaskList(request *DescribeSuperNodeInstallTaskListRequest) (response *DescribeSuperNodeInstallTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeSuperNodeInstallTaskListRequest()
	}
	response = NewDescribeSuperNodeInstallTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyRuleDetailRequest() (request *DescribeImageDenyRuleDetailRequest) {
	request = &DescribeImageDenyRuleDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageDenyRuleDetail")
	return
}

func NewDescribeImageDenyRuleDetailResponse() (response *DescribeImageDenyRuleDetailResponse) {
	response = &DescribeImageDenyRuleDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截规则详情
func (c *Client) DescribeImageDenyRuleDetail(request *DescribeImageDenyRuleDetailRequest) (response *DescribeImageDenyRuleDetailResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyRuleDetailRequest()
	}
	response = NewDescribeImageDenyRuleDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndexListRequest() (request *DescribeIndexListRequest) {
	request = &DescribeIndexListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeIndexList")
	return
}

func NewDescribeIndexListResponse() (response *DescribeIndexListResponse) {
	response = &DescribeIndexListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取索引列表
func (c *Client) DescribeIndexList(request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
	if request == nil {
		request = NewDescribeIndexListRequest()
	}
	response = NewDescribeIndexListResponse()
	err = c.Send(request, response)
	return
}

func NewAddOrModifyImageDenyRuleRequest() (request *AddOrModifyImageDenyRuleRequest) {
	request = &AddOrModifyImageDenyRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddOrModifyImageDenyRule")
	return
}

func NewAddOrModifyImageDenyRuleResponse() (response *AddOrModifyImageDenyRuleResponse) {
	response = &AddOrModifyImageDenyRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或编辑镜像拦截规则
func (c *Client) AddOrModifyImageDenyRule(request *AddOrModifyImageDenyRuleRequest) (response *AddOrModifyImageDenyRuleResponse, err error) {
	if request == nil {
		request = NewAddOrModifyImageDenyRuleRequest()
	}
	response = NewAddOrModifyImageDenyRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaliciousConnectionWhiteListRequest() (request *DescribeMaliciousConnectionWhiteListRequest) {
	request = &DescribeMaliciousConnectionWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeMaliciousConnectionWhiteList")
	return
}

func NewDescribeMaliciousConnectionWhiteListResponse() (response *DescribeMaliciousConnectionWhiteListResponse) {
	response = &DescribeMaliciousConnectionWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意外连白名单
func (c *Client) DescribeMaliciousConnectionWhiteList(request *DescribeMaliciousConnectionWhiteListRequest) (response *DescribeMaliciousConnectionWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeMaliciousConnectionWhiteListRequest()
	}
	response = NewDescribeMaliciousConnectionWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveLocalStorageItemRequest() (request *RemoveLocalStorageItemRequest) {
	request = &RemoveLocalStorageItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "RemoveLocalStorageItem")
	return
}

func NewRemoveLocalStorageItemResponse() (response *RemoveLocalStorageItemResponse) {
	response = &RemoveLocalStorageItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除本地存储数据
func (c *Client) RemoveLocalStorageItem(request *RemoveLocalStorageItemRequest) (response *RemoveLocalStorageItemResponse, err error) {
	if request == nil {
		request = NewRemoveLocalStorageItemRequest()
	}
	response = NewRemoveLocalStorageItemResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
	request = &DescribeClusterDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterDetail")
	return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
	response = &DescribeClusterDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单个集群的详细信息
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
	if request == nil {
		request = NewDescribeClusterDetailRequest()
	}
	response = NewDescribeClusterDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateReverseShellWhiteListsExportJobRequest() (request *CreateReverseShellWhiteListsExportJobRequest) {
	request = &CreateReverseShellWhiteListsExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateReverseShellWhiteListsExportJob")
	return
}

func NewCreateReverseShellWhiteListsExportJobResponse() (response *CreateReverseShellWhiteListsExportJobResponse) {
	response = &CreateReverseShellWhiteListsExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建反弹shell白名单列表导出任务
func (c *Client) CreateReverseShellWhiteListsExportJob(request *CreateReverseShellWhiteListsExportJobRequest) (response *CreateReverseShellWhiteListsExportJobResponse, err error) {
	if request == nil {
		request = NewCreateReverseShellWhiteListsExportJobRequest()
	}
	response = NewCreateReverseShellWhiteListsExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCompliancePolicyItemFromWhitelistRequest() (request *DeleteCompliancePolicyItemFromWhitelistRequest) {
	request = &DeleteCompliancePolicyItemFromWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteCompliancePolicyItemFromWhitelist")
	return
}

func NewDeleteCompliancePolicyItemFromWhitelistResponse() (response *DeleteCompliancePolicyItemFromWhitelistResponse) {
	response = &DeleteCompliancePolicyItemFromWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从白名单中删除将指定的检测项。
func (c *Client) DeleteCompliancePolicyItemFromWhitelist(request *DeleteCompliancePolicyItemFromWhitelistRequest) (response *DeleteCompliancePolicyItemFromWhitelistResponse, err error) {
	if request == nil {
		request = NewDeleteCompliancePolicyItemFromWhitelistRequest()
	}
	response = NewDeleteCompliancePolicyItemFromWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryVulListRequest() (request *DescribeAssetImageRegistryVulListRequest) {
	request = &DescribeAssetImageRegistryVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVulList")
	return
}

func NewDescribeAssetImageRegistryVulListResponse() (response *DescribeAssetImageRegistryVulListResponse) {
	response = &DescribeAssetImageRegistryVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询镜像漏洞列表
func (c *Client) DescribeAssetImageRegistryVulList(request *DescribeAssetImageRegistryVulListRequest) (response *DescribeAssetImageRegistryVulListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryVulListRequest()
	}
	response = NewDescribeAssetImageRegistryVulListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkTopologyRealFlowListRequest() (request *DescribeNetworkTopologyRealFlowListRequest) {
	request = &DescribeNetworkTopologyRealFlowListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkTopologyRealFlowList")
	return
}

func NewDescribeNetworkTopologyRealFlowListResponse() (response *DescribeNetworkTopologyRealFlowListResponse) {
	response = &DescribeNetworkTopologyRealFlowListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络查查询网络拓扑实时流量列表
func (c *Client) DescribeNetworkTopologyRealFlowList(request *DescribeNetworkTopologyRealFlowListRequest) (response *DescribeNetworkTopologyRealFlowListResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkTopologyRealFlowListRequest()
	}
	response = NewDescribeNetworkTopologyRealFlowListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReverseShellInnerIPShowStatusRequest() (request *ModifyReverseShellInnerIPShowStatusRequest) {
	request = &ModifyReverseShellInnerIPShowStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyReverseShellInnerIPShowStatus")
	return
}

func NewModifyReverseShellInnerIPShowStatusResponse() (response *ModifyReverseShellInnerIPShowStatusResponse) {
	response = &ModifyReverseShellInnerIPShowStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改反弹shelll内网ip状态
func (c *Client) ModifyReverseShellInnerIPShowStatus(request *ModifyReverseShellInnerIPShowStatusRequest) (response *ModifyReverseShellInnerIPShowStatusResponse, err error) {
	if request == nil {
		request = NewModifyReverseShellInnerIPShowStatusRequest()
	}
	response = NewModifyReverseShellInnerIPShowStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCompliancePolicyItemIgnoredAssetListRequest() (request *DescribeCompliancePolicyItemIgnoredAssetListRequest) {
	request = &DescribeCompliancePolicyItemIgnoredAssetListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeCompliancePolicyItemIgnoredAssetList")
	return
}

func NewDescribeCompliancePolicyItemIgnoredAssetListResponse() (response *DescribeCompliancePolicyItemIgnoredAssetListResponse) {
	response = &DescribeCompliancePolicyItemIgnoredAssetListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照 检测项 → 资产 的两级层次展开的第二层级：资产层级。
func (c *Client) DescribeCompliancePolicyItemIgnoredAssetList(request *DescribeCompliancePolicyItemIgnoredAssetListRequest) (response *DescribeCompliancePolicyItemIgnoredAssetListResponse, err error) {
	if request == nil {
		request = NewDescribeCompliancePolicyItemIgnoredAssetListRequest()
	}
	response = NewDescribeCompliancePolicyItemIgnoredAssetListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCheckComponentRequest() (request *CreateCheckComponentRequest) {
	request = &CreateCheckComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateCheckComponent")
	return
}

func NewCreateCheckComponentResponse() (response *CreateCheckComponentResponse) {
	response = &CreateCheckComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安装检查组件，创建防护容器
func (c *Client) CreateCheckComponent(request *CreateCheckComponentRequest) (response *CreateCheckComponentResponse, err error) {
	if request == nil {
		request = NewCreateCheckComponentRequest()
	}
	response = NewCreateCheckComponentResponse()
	err = c.Send(request, response)
	return
}

func NewExportVirusListRequest() (request *ExportVirusListRequest) {
	request = &ExportVirusListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ExportVirusList")
	return
}

func NewExportVirusListResponse() (response *ExportVirusListResponse) {
	response = &ExportVirusListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时文件查杀事件列表导出
func (c *Client) ExportVirusList(request *ExportVirusListRequest) (response *ExportVirusListResponse, err error) {
	if request == nil {
		request = NewExportVirusListRequest()
	}
	response = NewExportVirusListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateK8sApiAbnormalEventExportJobRequest() (request *CreateK8sApiAbnormalEventExportJobRequest) {
	request = &CreateK8sApiAbnormalEventExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateK8sApiAbnormalEventExportJob")
	return
}

func NewCreateK8sApiAbnormalEventExportJobResponse() (response *CreateK8sApiAbnormalEventExportJobResponse) {
	response = &CreateK8sApiAbnormalEventExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建k8s api异常事件导出任务
func (c *Client) CreateK8sApiAbnormalEventExportJob(request *CreateK8sApiAbnormalEventExportJobRequest) (response *CreateK8sApiAbnormalEventExportJobResponse, err error) {
	if request == nil {
		request = NewCreateK8sApiAbnormalEventExportJobRequest()
	}
	response = NewCreateK8sApiAbnormalEventExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageAutoAuthorizedLogListRequest() (request *DescribeImageAutoAuthorizedLogListRequest) {
	request = &DescribeImageAutoAuthorizedLogListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageAutoAuthorizedLogList")
	return
}

func NewDescribeImageAutoAuthorizedLogListResponse() (response *DescribeImageAutoAuthorizedLogListResponse) {
	response = &DescribeImageAutoAuthorizedLogListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像自动授权结果列表
func (c *Client) DescribeImageAutoAuthorizedLogList(request *DescribeImageAutoAuthorizedLogListRequest) (response *DescribeImageAutoAuthorizedLogListResponse, err error) {
	if request == nil {
		request = NewDescribeImageAutoAuthorizedLogListRequest()
	}
	response = NewDescribeImageAutoAuthorizedLogListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryVirusListExportRequest() (request *DescribeAssetImageRegistryVirusListExportRequest) {
	request = &DescribeAssetImageRegistryVirusListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVirusListExport")
	return
}

func NewDescribeAssetImageRegistryVirusListExportResponse() (response *DescribeAssetImageRegistryVirusListExportResponse) {
	response = &DescribeAssetImageRegistryVirusListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库木马信息列表导出
func (c *Client) DescribeAssetImageRegistryVirusListExport(request *DescribeAssetImageRegistryVirusListExportRequest) (response *DescribeAssetImageRegistryVirusListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryVirusListExportRequest()
	}
	response = NewDescribeAssetImageRegistryVirusListExportResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkFirewallClusterRefreshRequest() (request *CreateNetworkFirewallClusterRefreshRequest) {
	request = &CreateNetworkFirewallClusterRefreshRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateNetworkFirewallClusterRefresh")
	return
}

func NewCreateNetworkFirewallClusterRefreshResponse() (response *CreateNetworkFirewallClusterRefreshResponse) {
	response = &CreateNetworkFirewallClusterRefreshResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络集群下发刷新任务
func (c *Client) CreateNetworkFirewallClusterRefresh(request *CreateNetworkFirewallClusterRefreshRequest) (response *CreateNetworkFirewallClusterRefreshResponse, err error) {
	if request == nil {
		request = NewCreateNetworkFirewallClusterRefreshRequest()
	}
	response = NewCreateNetworkFirewallClusterRefreshResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusSampleDownloadUrlRequest() (request *DescribeVirusSampleDownloadUrlRequest) {
	request = &DescribeVirusSampleDownloadUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusSampleDownloadUrl")
	return
}

func NewDescribeVirusSampleDownloadUrlResponse() (response *DescribeVirusSampleDownloadUrlResponse) {
	response = &DescribeVirusSampleDownloadUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马样本下载url
func (c *Client) DescribeVirusSampleDownloadUrl(request *DescribeVirusSampleDownloadUrlRequest) (response *DescribeVirusSampleDownloadUrlResponse, err error) {
	if request == nil {
		request = NewDescribeVirusSampleDownloadUrlRequest()
	}
	response = NewDescribeVirusSampleDownloadUrlResponse()
	err = c.Send(request, response)
	return
}

func NewCreateManualUpgradeVersionRequest() (request *CreateManualUpgradeVersionRequest) {
	request = &CreateManualUpgradeVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateManualUpgradeVersion")
	return
}

func NewCreateManualUpgradeVersionResponse() (response *CreateManualUpgradeVersionResponse) {
	response = &CreateManualUpgradeVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 手动创建终端版本升级任务
func (c *Client) CreateManualUpgradeVersion(request *CreateManualUpgradeVersionRequest) (response *CreateManualUpgradeVersionResponse, err error) {
	if request == nil {
		request = NewCreateManualUpgradeVersionRequest()
	}
	response = NewCreateManualUpgradeVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusScanConfigRequest() (request *DescribeVirusScanConfigRequest) {
	request = &DescribeVirusScanConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusScanConfig")
	return
}

func NewDescribeVirusScanConfigResponse() (response *DescribeVirusScanConfigResponse) {
	response = &DescribeVirusScanConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询文件查杀新设置
func (c *Client) DescribeVirusScanConfig(request *DescribeVirusScanConfigRequest) (response *DescribeVirusScanConfigResponse, err error) {
	if request == nil {
		request = NewDescribeVirusScanConfigRequest()
	}
	response = NewDescribeVirusScanConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkTopologyFlowSwitchStatusRequest() (request *DescribeNetworkTopologyFlowSwitchStatusRequest) {
	request = &DescribeNetworkTopologyFlowSwitchStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkTopologyFlowSwitchStatus")
	return
}

func NewDescribeNetworkTopologyFlowSwitchStatusResponse() (response *DescribeNetworkTopologyFlowSwitchStatusResponse) {
	response = &DescribeNetworkTopologyFlowSwitchStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络查询网络拓扑集群流量开关任务执行状态
func (c *Client) DescribeNetworkTopologyFlowSwitchStatus(request *DescribeNetworkTopologyFlowSwitchStatusRequest) (response *DescribeNetworkTopologyFlowSwitchStatusResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkTopologyFlowSwitchStatusRequest()
	}
	response = NewDescribeNetworkTopologyFlowSwitchStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalEventInfoRequest() (request *DescribeK8sApiAbnormalEventInfoRequest) {
	request = &DescribeK8sApiAbnormalEventInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalEventInfo")
	return
}

func NewDescribeK8sApiAbnormalEventInfoResponse() (response *DescribeK8sApiAbnormalEventInfoResponse) {
	response = &DescribeK8sApiAbnormalEventInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8s api 异常事件详情
func (c *Client) DescribeK8sApiAbnormalEventInfo(request *DescribeK8sApiAbnormalEventInfoRequest) (response *DescribeK8sApiAbnormalEventInfoResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalEventInfoRequest()
	}
	response = NewDescribeK8sApiAbnormalEventInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetClusterListRequest() (request *DescribeAssetClusterListRequest) {
	request = &DescribeAssetClusterListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetClusterList")
	return
}

func NewDescribeAssetClusterListResponse() (response *DescribeAssetClusterListResponse) {
	response = &DescribeAssetClusterListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群列表
func (c *Client) DescribeAssetClusterList(request *DescribeAssetClusterListRequest) (response *DescribeAssetClusterListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetClusterListRequest()
	}
	response = NewDescribeAssetClusterListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceEventDetailRequest() (request *DescribeVulDefenceEventDetailRequest) {
	request = &DescribeVulDefenceEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefenceEventDetail")
	return
}

func NewDescribeVulDefenceEventDetailResponse() (response *DescribeVulDefenceEventDetailResponse) {
	response = &DescribeVulDefenceEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞防御事件详情
func (c *Client) DescribeVulDefenceEventDetail(request *DescribeVulDefenceEventDetailRequest) (response *DescribeVulDefenceEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceEventDetailRequest()
	}
	response = NewDescribeVulDefenceEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewStopVulScanTaskRequest() (request *StopVulScanTaskRequest) {
	request = &StopVulScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "StopVulScanTask")
	return
}

func NewStopVulScanTaskResponse() (response *StopVulScanTaskResponse) {
	response = &StopVulScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止漏洞扫描任务
func (c *Client) StopVulScanTask(request *StopVulScanTaskRequest) (response *StopVulScanTaskResponse, err error) {
	if request == nil {
		request = NewStopVulScanTaskRequest()
	}
	response = NewStopVulScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessEventTendencyRequest() (request *DescribeAbnormalProcessEventTendencyRequest) {
	request = &DescribeAbnormalProcessEventTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessEventTendency")
	return
}

func NewDescribeAbnormalProcessEventTendencyResponse() (response *DescribeAbnormalProcessEventTendencyResponse) {
	response = &DescribeAbnormalProcessEventTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询待处理异常进程事件趋势
func (c *Client) DescribeAbnormalProcessEventTendency(request *DescribeAbnormalProcessEventTendencyRequest) (response *DescribeAbnormalProcessEventTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessEventTendencyRequest()
	}
	response = NewDescribeAbnormalProcessEventTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulDefenceEventExportJobRequest() (request *CreateVulDefenceEventExportJobRequest) {
	request = &CreateVulDefenceEventExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateVulDefenceEventExportJob")
	return
}

func NewCreateVulDefenceEventExportJobResponse() (response *CreateVulDefenceEventExportJobResponse) {
	response = &CreateVulDefenceEventExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建漏洞防御导出任务
func (c *Client) CreateVulDefenceEventExportJob(request *CreateVulDefenceEventExportJobRequest) (response *CreateVulDefenceEventExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVulDefenceEventExportJobRequest()
	}
	response = NewCreateVulDefenceEventExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterIngressListRequest() (request *DescribeClusterIngressListRequest) {
	request = &DescribeClusterIngressListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterIngressList")
	return
}

func NewDescribeClusterIngressListResponse() (response *DescribeClusterIngressListResponse) {
	response = &DescribeClusterIngressListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群的Ingress列表
func (c *Client) DescribeClusterIngressList(request *DescribeClusterIngressListRequest) (response *DescribeClusterIngressListResponse, err error) {
	if request == nil {
		request = NewDescribeClusterIngressListRequest()
	}
	response = NewDescribeClusterIngressListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetImageScanTaskRequest() (request *CreateAssetImageScanTaskRequest) {
	request = &CreateAssetImageScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageScanTask")
	return
}

func NewCreateAssetImageScanTaskResponse() (response *CreateAssetImageScanTaskResponse) {
	response = &CreateAssetImageScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全创建镜像扫描任务
func (c *Client) CreateAssetImageScanTask(request *CreateAssetImageScanTaskRequest) (response *CreateAssetImageScanTaskResponse, err error) {
	if request == nil {
		request = NewCreateAssetImageScanTaskRequest()
	}
	response = NewCreateAssetImageScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogJoinObjectListRequest() (request *DescribeSecLogJoinObjectListRequest) {
	request = &DescribeSecLogJoinObjectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogJoinObjectList")
	return
}

func NewDescribeSecLogJoinObjectListResponse() (response *DescribeSecLogJoinObjectListResponse) {
	response = &DescribeSecLogJoinObjectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志接入对象列表
func (c *Client) DescribeSecLogJoinObjectList(request *DescribeSecLogJoinObjectListRequest) (response *DescribeSecLogJoinObjectListResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogJoinObjectListRequest()
	}
	response = NewDescribeSecLogJoinObjectListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageDenyEventExportJobRequest() (request *CreateImageDenyEventExportJobRequest) {
	request = &CreateImageDenyEventExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateImageDenyEventExportJob")
	return
}

func NewCreateImageDenyEventExportJobResponse() (response *CreateImageDenyEventExportJobResponse) {
	response = &CreateImageDenyEventExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像拦截事件导出任务
func (c *Client) CreateImageDenyEventExportJob(request *CreateImageDenyEventExportJobRequest) (response *CreateImageDenyEventExportJobResponse, err error) {
	if request == nil {
		request = NewCreateImageDenyEventExportJobRequest()
	}
	response = NewCreateImageDenyEventExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetImageVirusExportJobRequest() (request *CreateAssetImageVirusExportJobRequest) {
	request = &CreateAssetImageVirusExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageVirusExportJob")
	return
}

func NewCreateAssetImageVirusExportJobResponse() (response *CreateAssetImageVirusExportJobResponse) {
	response = &CreateAssetImageVirusExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建本地镜像木马列表导出任务
func (c *Client) CreateAssetImageVirusExportJob(request *CreateAssetImageVirusExportJobRequest) (response *CreateAssetImageVirusExportJobResponse, err error) {
	if request == nil {
		request = NewCreateAssetImageVirusExportJobRequest()
	}
	response = NewCreateAssetImageVirusExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetComponentListRequest() (request *DescribeAssetComponentListRequest) {
	request = &DescribeAssetComponentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetComponentList")
	return
}

func NewDescribeAssetComponentListResponse() (response *DescribeAssetComponentListResponse) {
	response = &DescribeAssetComponentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全搜索查询容器组件列表
func (c *Client) DescribeAssetComponentList(request *DescribeAssetComponentListRequest) (response *DescribeAssetComponentListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetComponentListRequest()
	}
	response = NewDescribeAssetComponentListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceHostRequest() (request *DescribeVulDefenceHostRequest) {
	request = &DescribeVulDefenceHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefenceHost")
	return
}

func NewDescribeVulDefenceHostResponse() (response *DescribeVulDefenceHostResponse) {
	response = &DescribeVulDefenceHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞防御的主机列表
func (c *Client) DescribeVulDefenceHost(request *DescribeVulDefenceHostRequest) (response *DescribeVulDefenceHostResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceHostRequest()
	}
	response = NewDescribeVulDefenceHostResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaliciousConnectionWhiteListRequest() (request *DeleteMaliciousConnectionWhiteListRequest) {
	request = &DeleteMaliciousConnectionWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteMaliciousConnectionWhiteList")
	return
}

func NewDeleteMaliciousConnectionWhiteListResponse() (response *DeleteMaliciousConnectionWhiteListResponse) {
	response = &DeleteMaliciousConnectionWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除恶意请求外连白名单
func (c *Client) DeleteMaliciousConnectionWhiteList(request *DeleteMaliciousConnectionWhiteListRequest) (response *DeleteMaliciousConnectionWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteMaliciousConnectionWhiteListRequest()
	}
	response = NewDeleteMaliciousConnectionWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyK8sApiAbnormalRuleStatusRequest() (request *ModifyK8sApiAbnormalRuleStatusRequest) {
	request = &ModifyK8sApiAbnormalRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyK8sApiAbnormalRuleStatus")
	return
}

func NewModifyK8sApiAbnormalRuleStatusResponse() (response *ModifyK8sApiAbnormalRuleStatusResponse) {
	response = &ModifyK8sApiAbnormalRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改k8sapi异常事件规则状态
func (c *Client) ModifyK8sApiAbnormalRuleStatus(request *ModifyK8sApiAbnormalRuleStatusRequest) (response *ModifyK8sApiAbnormalRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyK8sApiAbnormalRuleStatusRequest()
	}
	response = NewModifyK8sApiAbnormalRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewAddOrModifyMaliciousConnectionBlackListRequest() (request *AddOrModifyMaliciousConnectionBlackListRequest) {
	request = &AddOrModifyMaliciousConnectionBlackListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddOrModifyMaliciousConnectionBlackList")
	return
}

func NewAddOrModifyMaliciousConnectionBlackListResponse() (response *AddOrModifyMaliciousConnectionBlackListResponse) {
	response = &AddOrModifyMaliciousConnectionBlackListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加或修改恶意外连的黑名单
func (c *Client) AddOrModifyMaliciousConnectionBlackList(request *AddOrModifyMaliciousConnectionBlackListRequest) (response *AddOrModifyMaliciousConnectionBlackListResponse, err error) {
	if request == nil {
		request = NewAddOrModifyMaliciousConnectionBlackListRequest()
	}
	response = NewAddOrModifyMaliciousConnectionBlackListResponse()
	err = c.Send(request, response)
	return
}

func NewAddClusterNodeToWhiteListRequest() (request *AddClusterNodeToWhiteListRequest) {
	request = &AddClusterNodeToWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddClusterNodeToWhiteList")
	return
}

func NewAddClusterNodeToWhiteListResponse() (response *AddClusterNodeToWhiteListResponse) {
	response = &AddClusterNodeToWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加集群节点资产到白名单，只支持单个添加
func (c *Client) AddClusterNodeToWhiteList(request *AddClusterNodeToWhiteListRequest) (response *AddClusterNodeToWhiteListResponse, err error) {
	if request == nil {
		request = NewAddClusterNodeToWhiteListRequest()
	}
	response = NewAddClusterNodeToWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellWhiteListDetailRequest() (request *DescribeReverseShellWhiteListDetailRequest) {
	request = &DescribeReverseShellWhiteListDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellWhiteListDetail")
	return
}

func NewDescribeReverseShellWhiteListDetailResponse() (response *DescribeReverseShellWhiteListDetailResponse) {
	response = &DescribeReverseShellWhiteListDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时反弹shell白名单详细信息
func (c *Client) DescribeReverseShellWhiteListDetail(request *DescribeReverseShellWhiteListDetailRequest) (response *DescribeReverseShellWhiteListDetailResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellWhiteListDetailRequest()
	}
	response = NewDescribeReverseShellWhiteListDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageScanTaskRequest() (request *DescribeAssetImageScanTaskRequest) {
	request = &DescribeAssetImageScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageScanTask")
	return
}

func NewDescribeAssetImageScanTaskResponse() (response *DescribeAssetImageScanTaskResponse) {
	response = &DescribeAssetImageScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询正在一键扫描的镜像扫描taskid
func (c *Client) DescribeAssetImageScanTask(request *DescribeAssetImageScanTaskRequest) (response *DescribeAssetImageScanTaskResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageScanTaskRequest()
	}
	response = NewDescribeAssetImageScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageBindRuleInfoRequest() (request *DescribeAssetImageBindRuleInfoRequest) {
	request = &DescribeAssetImageBindRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageBindRuleInfo")
	return
}

func NewDescribeAssetImageBindRuleInfoResponse() (response *DescribeAssetImageBindRuleInfoResponse) {
	response = &DescribeAssetImageBindRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像绑定规则列表信息，包含运行时访问控制和异常进程公用
func (c *Client) DescribeAssetImageBindRuleInfo(request *DescribeAssetImageBindRuleInfoRequest) (response *DescribeAssetImageBindRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageBindRuleInfoRequest()
	}
	response = NewDescribeAssetImageBindRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsListRequest() (request *DescribeRiskDnsListRequest) {
	request = &DescribeRiskDnsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskDnsList")
	return
}

func NewDescribeRiskDnsListResponse() (response *DescribeRiskDnsListResponse) {
	response = &DescribeRiskDnsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意请求事件列表
func (c *Client) DescribeRiskDnsList(request *DescribeRiskDnsListRequest) (response *DescribeRiskDnsListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsListRequest()
	}
	response = NewDescribeRiskDnsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeESAggregationsRequest() (request *DescribeESAggregationsRequest) {
	request = &DescribeESAggregationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeESAggregations")
	return
}

func NewDescribeESAggregationsResponse() (response *DescribeESAggregationsResponse) {
	response = &DescribeESAggregationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ES字段聚合结果
func (c *Client) DescribeESAggregations(request *DescribeESAggregationsRequest) (response *DescribeESAggregationsResponse, err error) {
	if request == nil {
		request = NewDescribeESAggregationsRequest()
	}
	response = NewDescribeESAggregationsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulDefenceEventStatusRequest() (request *ModifyVulDefenceEventStatusRequest) {
	request = &ModifyVulDefenceEventStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVulDefenceEventStatus")
	return
}

func NewModifyVulDefenceEventStatusResponse() (response *ModifyVulDefenceEventStatusResponse) {
	response = &ModifyVulDefenceEventStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改漏洞防御事件状态
func (c *Client) ModifyVulDefenceEventStatus(request *ModifyVulDefenceEventStatusRequest) (response *ModifyVulDefenceEventStatusResponse, err error) {
	if request == nil {
		request = NewModifyVulDefenceEventStatusRequest()
	}
	response = NewModifyVulDefenceEventStatusResponse()
	err = c.Send(request, response)
	return
}

func NewAddCompliancePolicyItemToWhitelistRequest() (request *AddCompliancePolicyItemToWhitelistRequest) {
	request = &AddCompliancePolicyItemToWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddCompliancePolicyItemToWhitelist")
	return
}

func NewAddCompliancePolicyItemToWhitelistResponse() (response *AddCompliancePolicyItemToWhitelistResponse) {
	response = &AddCompliancePolicyItemToWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将指定的检测项添加到白名单中，不显示未通过结果。
func (c *Client) AddCompliancePolicyItemToWhitelist(request *AddCompliancePolicyItemToWhitelistRequest) (response *AddCompliancePolicyItemToWhitelistResponse, err error) {
	if request == nil {
		request = NewAddCompliancePolicyItemToWhitelistRequest()
	}
	response = NewAddCompliancePolicyItemToWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageSimpleListRequest() (request *DescribeAssetImageSimpleListRequest) {
	request = &DescribeAssetImageSimpleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageSimpleList")
	return
}

func NewDescribeAssetImageSimpleListResponse() (response *DescribeAssetImageSimpleListResponse) {
	response = &DescribeAssetImageSimpleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全搜索查询镜像简略信息列表
func (c *Client) DescribeAssetImageSimpleList(request *DescribeAssetImageSimpleListRequest) (response *DescribeAssetImageSimpleListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageSimpleListRequest()
	}
	response = NewDescribeAssetImageSimpleListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEscapeWhiteListRequest() (request *DeleteEscapeWhiteListRequest) {
	request = &DeleteEscapeWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteEscapeWhiteList")
	return
}

func NewDeleteEscapeWhiteListResponse() (response *DeleteEscapeWhiteListResponse) {
	response = &DeleteEscapeWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除逃逸白名单
func (c *Client) DeleteEscapeWhiteList(request *DeleteEscapeWhiteListRequest) (response *DeleteEscapeWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteEscapeWhiteListRequest()
	}
	response = NewDeleteEscapeWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventEscapeImageListRequest() (request *DescribeEventEscapeImageListRequest) {
	request = &DescribeEventEscapeImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEventEscapeImageList")
	return
}

func NewDescribeEventEscapeImageListResponse() (response *DescribeEventEscapeImageListResponse) {
	response = &DescribeEventEscapeImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeRiskContainerImageList查询风险容器镜像列表
func (c *Client) DescribeEventEscapeImageList(request *DescribeEventEscapeImageListRequest) (response *DescribeEventEscapeImageListResponse, err error) {
	if request == nil {
		request = NewDescribeEventEscapeImageListRequest()
	}
	response = NewDescribeEventEscapeImageListResponse()
	err = c.Send(request, response)
	return
}

func NewGetLocalStorageItemRequest() (request *GetLocalStorageItemRequest) {
	request = &GetLocalStorageItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "GetLocalStorageItem")
	return
}

func NewGetLocalStorageItemResponse() (response *GetLocalStorageItemResponse) {
	response = &GetLocalStorageItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取本地存储数据
func (c *Client) GetLocalStorageItem(request *GetLocalStorageItemRequest) (response *GetLocalStorageItemResponse, err error) {
	if request == nil {
		request = NewGetLocalStorageItemRequest()
	}
	response = NewGetLocalStorageItemResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceAssetListRequest() (request *DescribeComplianceAssetListRequest) {
	request = &DescribeComplianceAssetListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceAssetList")
	return
}

func NewDescribeComplianceAssetListResponse() (response *DescribeComplianceAssetListResponse) {
	response = &DescribeComplianceAssetListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某类资产的列表
func (c *Client) DescribeComplianceAssetList(request *DescribeComplianceAssetListRequest) (response *DescribeComplianceAssetListResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceAssetListRequest()
	}
	response = NewDescribeComplianceAssetListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeRegexpWhiteListRequest() (request *DescribeEscapeRegexpWhiteListRequest) {
	request = &DescribeEscapeRegexpWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeRegexpWhiteList")
	return
}

func NewDescribeEscapeRegexpWhiteListResponse() (response *DescribeEscapeRegexpWhiteListResponse) {
	response = &DescribeEscapeRegexpWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器逃逸正则白名单列表
func (c *Client) DescribeEscapeRegexpWhiteList(request *DescribeEscapeRegexpWhiteListRequest) (response *DescribeEscapeRegexpWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeRegexpWhiteListRequest()
	}
	response = NewDescribeEscapeRegexpWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNodeWhiteListRequest() (request *DescribeClusterNodeWhiteListRequest) {
	request = &DescribeClusterNodeWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterNodeWhiteList")
	return
}

func NewDescribeClusterNodeWhiteListResponse() (response *DescribeClusterNodeWhiteListResponse) {
	response = &DescribeClusterNodeWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取忽略的节点信息白名单
func (c *Client) DescribeClusterNodeWhiteList(request *DescribeClusterNodeWhiteListRequest) (response *DescribeClusterNodeWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNodeWhiteListRequest()
	}
	response = NewDescribeClusterNodeWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterUninstallCmdRequest() (request *DescribeClusterUninstallCmdRequest) {
	request = &DescribeClusterUninstallCmdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterUninstallCmd")
	return
}

func NewDescribeClusterUninstallCmdResponse() (response *DescribeClusterUninstallCmdResponse) {
	response = &DescribeClusterUninstallCmdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群卸载命令
func (c *Client) DescribeClusterUninstallCmd(request *DescribeClusterUninstallCmdRequest) (response *DescribeClusterUninstallCmdResponse, err error) {
	if request == nil {
		request = NewDescribeClusterUninstallCmdRequest()
	}
	response = NewDescribeClusterUninstallCmdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageHostListRequest() (request *DescribeAssetImageHostListRequest) {
	request = &DescribeAssetImageHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageHostList")
	return
}

func NewDescribeAssetImageHostListResponse() (response *DescribeAssetImageHostListResponse) {
	response = &DescribeAssetImageHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全查询镜像关联主机
func (c *Client) DescribeAssetImageHostList(request *DescribeAssetImageHostListRequest) (response *DescribeAssetImageHostListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageHostListRequest()
	}
	response = NewDescribeAssetImageHostListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalSummaryRequest() (request *DescribeK8sApiAbnormalSummaryRequest) {
	request = &DescribeK8sApiAbnormalSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalSummary")
	return
}

func NewDescribeK8sApiAbnormalSummaryResponse() (response *DescribeK8sApiAbnormalSummaryResponse) {
	response = &DescribeK8sApiAbnormalSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8sapi异常事件统计
func (c *Client) DescribeK8sApiAbnormalSummary(request *DescribeK8sApiAbnormalSummaryRequest) (response *DescribeK8sApiAbnormalSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalSummaryRequest()
	}
	response = NewDescribeK8sApiAbnormalSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceSettingRequest() (request *DescribeVulDefenceSettingRequest) {
	request = &DescribeVulDefenceSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefenceSetting")
	return
}

func NewDescribeVulDefenceSettingResponse() (response *DescribeVulDefenceSettingResponse) {
	response = &DescribeVulDefenceSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞防御设置信息
func (c *Client) DescribeVulDefenceSetting(request *DescribeVulDefenceSettingRequest) (response *DescribeVulDefenceSettingResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceSettingRequest()
	}
	response = NewDescribeVulDefenceSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSuperNodeInfoRequest() (request *DescribeAssetSuperNodeInfoRequest) {
	request = &DescribeAssetSuperNodeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetSuperNodeInfo")
	return
}

func NewDescribeAssetSuperNodeInfoResponse() (response *DescribeAssetSuperNodeInfoResponse) {
	response = &DescribeAssetSuperNodeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询超级节点详情
func (c *Client) DescribeAssetSuperNodeInfo(request *DescribeAssetSuperNodeInfoRequest) (response *DescribeAssetSuperNodeInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSuperNodeInfoRequest()
	}
	response = NewDescribeAssetSuperNodeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkTopologyClusterChartRequest() (request *DescribeNetworkTopologyClusterChartRequest) {
	request = &DescribeNetworkTopologyClusterChartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkTopologyClusterChart")
	return
}

func NewDescribeNetworkTopologyClusterChartResponse() (response *DescribeNetworkTopologyClusterChartResponse) {
	response = &DescribeNetworkTopologyClusterChartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络拓扑集群图
func (c *Client) DescribeNetworkTopologyClusterChart(request *DescribeNetworkTopologyClusterChartRequest) (response *DescribeNetworkTopologyClusterChartResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkTopologyClusterChartRequest()
	}
	response = NewDescribeNetworkTopologyClusterChartResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHybridClustersRequest() (request *DescribeHybridClustersRequest) {
	request = &DescribeHybridClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeHybridClusters")
	return
}

func NewDescribeHybridClustersResponse() (response *DescribeHybridClustersResponse) {
	response = &DescribeHybridClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 枚举所有未导入的混合云集群
func (c *Client) DescribeHybridClusters(request *DescribeHybridClustersRequest) (response *DescribeHybridClustersResponse, err error) {
	if request == nil {
		request = NewDescribeHybridClustersRequest()
	}
	response = NewDescribeHybridClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellWhiteListsRequest() (request *DescribeReverseShellWhiteListsRequest) {
	request = &DescribeReverseShellWhiteListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellWhiteLists")
	return
}

func NewDescribeReverseShellWhiteListsResponse() (response *DescribeReverseShellWhiteListsResponse) {
	response = &DescribeReverseShellWhiteListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时运行时反弹shell白名单列表信息
func (c *Client) DescribeReverseShellWhiteLists(request *DescribeReverseShellWhiteListsRequest) (response *DescribeReverseShellWhiteListsResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellWhiteListsRequest()
	}
	response = NewDescribeReverseShellWhiteListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallPolicyDetailRequest() (request *DescribeNetworkFirewallPolicyDetailRequest) {
	request = &DescribeNetworkFirewallPolicyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPolicyDetail")
	return
}

func NewDescribeNetworkFirewallPolicyDetailResponse() (response *DescribeNetworkFirewallPolicyDetailResponse) {
	response = &DescribeNetworkFirewallPolicyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络集群查看策略详情
func (c *Client) DescribeNetworkFirewallPolicyDetail(request *DescribeNetworkFirewallPolicyDetailRequest) (response *DescribeNetworkFirewallPolicyDetailResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallPolicyDetailRequest()
	}
	response = NewDescribeNetworkFirewallPolicyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchImageAutoAuthorizedRuleRequest() (request *SwitchImageAutoAuthorizedRuleRequest) {
	request = &SwitchImageAutoAuthorizedRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "SwitchImageAutoAuthorizedRule")
	return
}

func NewSwitchImageAutoAuthorizedRuleResponse() (response *SwitchImageAutoAuthorizedRuleResponse) {
	response = &SwitchImageAutoAuthorizedRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑本地镜像自动授权开关
func (c *Client) SwitchImageAutoAuthorizedRule(request *SwitchImageAutoAuthorizedRuleRequest) (response *SwitchImageAutoAuthorizedRuleResponse, err error) {
	if request == nil {
		request = NewSwitchImageAutoAuthorizedRuleRequest()
	}
	response = NewSwitchImageAutoAuthorizedRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserPublicClusterRequest() (request *CreateUserPublicClusterRequest) {
	request = &CreateUserPublicClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateUserPublicCluster")
	return
}

func NewCreateUserPublicClusterResponse() (response *CreateUserPublicClusterResponse) {
	response = &CreateUserPublicClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建用户自建集群
func (c *Client) CreateUserPublicCluster(request *CreateUserPublicClusterRequest) (response *CreateUserPublicClusterResponse, err error) {
	if request == nil {
		request = NewCreateUserPublicClusterRequest()
	}
	response = NewCreateUserPublicClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetAppServiceListRequest() (request *DescribeAssetAppServiceListRequest) {
	request = &DescribeAssetAppServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetAppServiceList")
	return
}

func NewDescribeAssetAppServiceListResponse() (response *DescribeAssetAppServiceListResponse) {
	response = &DescribeAssetAppServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全查询app服务列表
func (c *Client) DescribeAssetAppServiceList(request *DescribeAssetAppServiceListRequest) (response *DescribeAssetAppServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetAppServiceListRequest()
	}
	response = NewDescribeAssetAppServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewAddClusterWorkloadToWhiteListRequest() (request *AddClusterWorkloadToWhiteListRequest) {
	request = &AddClusterWorkloadToWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddClusterWorkloadToWhiteList")
	return
}

func NewAddClusterWorkloadToWhiteListResponse() (response *AddClusterWorkloadToWhiteListResponse) {
	response = &AddClusterWorkloadToWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加工作负载名称到白名单
func (c *Client) AddClusterWorkloadToWhiteList(request *AddClusterWorkloadToWhiteListRequest) (response *AddClusterWorkloadToWhiteListResponse, err error) {
	if request == nil {
		request = NewAddClusterWorkloadToWhiteListRequest()
	}
	response = NewAddClusterWorkloadToWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallEventsExportRequest() (request *DescribeRiskSyscallEventsExportRequest) {
	request = &DescribeRiskSyscallEventsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallEventsExport")
	return
}

func NewDescribeRiskSyscallEventsExportResponse() (response *DescribeRiskSyscallEventsExportResponse) {
	response = &DescribeRiskSyscallEventsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时高危系统调用列表导出
func (c *Client) DescribeRiskSyscallEventsExport(request *DescribeRiskSyscallEventsExportRequest) (response *DescribeRiskSyscallEventsExportResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallEventsExportRequest()
	}
	response = NewDescribeRiskSyscallEventsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryRiskListExportRequest() (request *DescribeAssetImageRegistryRiskListExportRequest) {
	request = &DescribeAssetImageRegistryRiskListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRiskListExport")
	return
}

func NewDescribeAssetImageRegistryRiskListExportResponse() (response *DescribeAssetImageRegistryRiskListExportResponse) {
	response = &DescribeAssetImageRegistryRiskListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库敏感信息列表导出
func (c *Client) DescribeAssetImageRegistryRiskListExport(request *DescribeAssetImageRegistryRiskListExportRequest) (response *DescribeAssetImageRegistryRiskListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryRiskListExportRequest()
	}
	response = NewDescribeAssetImageRegistryRiskListExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusScanTimeoutSettingRequest() (request *DescribeVirusScanTimeoutSettingRequest) {
	request = &DescribeVirusScanTimeoutSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusScanTimeoutSetting")
	return
}

func NewDescribeVirusScanTimeoutSettingResponse() (response *DescribeVirusScanTimeoutSettingResponse) {
	response = &DescribeVirusScanTimeoutSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时文件扫描超时设置查询
func (c *Client) DescribeVirusScanTimeoutSetting(request *DescribeVirusScanTimeoutSettingRequest) (response *DescribeVirusScanTimeoutSettingResponse, err error) {
	if request == nil {
		request = NewDescribeVirusScanTimeoutSettingRequest()
	}
	response = NewDescribeVirusScanTimeoutSettingResponse()
	err = c.Send(request, response)
	return
}

func NewRenewImageAuthorizeStateRequest() (request *RenewImageAuthorizeStateRequest) {
	request = &RenewImageAuthorizeStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "RenewImageAuthorizeState")
	return
}

func NewRenewImageAuthorizeStateResponse() (response *RenewImageAuthorizeStateResponse) {
	response = &RenewImageAuthorizeStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// RenewImageAuthorizeState   授权镜像扫描
func (c *Client) RenewImageAuthorizeState(request *RenewImageAuthorizeStateRequest) (response *RenewImageAuthorizeStateResponse, err error) {
	if request == nil {
		request = NewRenewImageAuthorizeStateRequest()
	}
	response = NewRenewImageAuthorizeStateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessControlsRuleExportJobRequest() (request *CreateAccessControlsRuleExportJobRequest) {
	request = &CreateAccessControlsRuleExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateAccessControlsRuleExportJob")
	return
}

func NewCreateAccessControlsRuleExportJobResponse() (response *CreateAccessControlsRuleExportJobResponse) {
	response = &CreateAccessControlsRuleExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建文件篡改规则导出任务
func (c *Client) CreateAccessControlsRuleExportJob(request *CreateAccessControlsRuleExportJobRequest) (response *CreateAccessControlsRuleExportJobResponse, err error) {
	if request == nil {
		request = NewCreateAccessControlsRuleExportJobRequest()
	}
	response = NewCreateAccessControlsRuleExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAbnormalProcessRulesRequest() (request *DeleteAbnormalProcessRulesRequest) {
	request = &DeleteAbnormalProcessRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteAbnormalProcessRules")
	return
}

func NewDeleteAbnormalProcessRulesResponse() (response *DeleteAbnormalProcessRulesResponse) {
	response = &DeleteAbnormalProcessRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运行异常进程策略
func (c *Client) DeleteAbnormalProcessRules(request *DeleteAbnormalProcessRulesRequest) (response *DeleteAbnormalProcessRulesResponse, err error) {
	if request == nil {
		request = NewDeleteAbnormalProcessRulesRequest()
	}
	response = NewDeleteAbnormalProcessRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogCleanSettingInfoRequest() (request *DescribeSecLogCleanSettingInfoRequest) {
	request = &DescribeSecLogCleanSettingInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogCleanSettingInfo")
	return
}

func NewDescribeSecLogCleanSettingInfoResponse() (response *DescribeSecLogCleanSettingInfoResponse) {
	response = &DescribeSecLogCleanSettingInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志清理设置详情
func (c *Client) DescribeSecLogCleanSettingInfo(request *DescribeSecLogCleanSettingInfoRequest) (response *DescribeSecLogCleanSettingInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogCleanSettingInfoRequest()
	}
	response = NewDescribeSecLogCleanSettingInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetContainerNetworkInfoRequest() (request *DescribeAssetContainerNetworkInfoRequest) {
	request = &DescribeAssetContainerNetworkInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetContainerNetworkInfo")
	return
}

func NewDescribeAssetContainerNetworkInfoResponse() (response *DescribeAssetContainerNetworkInfoResponse) {
	response = &DescribeAssetContainerNetworkInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器网络信息
func (c *Client) DescribeAssetContainerNetworkInfo(request *DescribeAssetContainerNetworkInfoRequest) (response *DescribeAssetContainerNetworkInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetContainerNetworkInfoRequest()
	}
	response = NewDescribeAssetContainerNetworkInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetContainerExportJobRequest() (request *CreateAssetContainerExportJobRequest) {
	request = &CreateAssetContainerExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetContainerExportJob")
	return
}

func NewCreateAssetContainerExportJobResponse() (response *CreateAssetContainerExportJobResponse) {
	response = &CreateAssetContainerExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建资产容器列表导出任务
func (c *Client) CreateAssetContainerExportJob(request *CreateAssetContainerExportJobRequest) (response *CreateAssetContainerExportJobResponse, err error) {
	if request == nil {
		request = NewCreateAssetContainerExportJobRequest()
	}
	response = NewCreateAssetContainerExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetImageRegistryScanStopOneKeyRequest() (request *ModifyAssetImageRegistryScanStopOneKeyRequest) {
	request = &ModifyAssetImageRegistryScanStopOneKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyAssetImageRegistryScanStopOneKey")
	return
}

func NewModifyAssetImageRegistryScanStopOneKeyResponse() (response *ModifyAssetImageRegistryScanStopOneKeyResponse) {
	response = &ModifyAssetImageRegistryScanStopOneKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库停止镜像一键扫描任务
func (c *Client) ModifyAssetImageRegistryScanStopOneKey(request *ModifyAssetImageRegistryScanStopOneKeyRequest) (response *ModifyAssetImageRegistryScanStopOneKeyResponse, err error) {
	if request == nil {
		request = NewModifyAssetImageRegistryScanStopOneKeyRequest()
	}
	response = NewModifyAssetImageRegistryScanStopOneKeyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVirusScanAgainRequest() (request *CreateVirusScanAgainRequest) {
	request = &CreateVirusScanAgainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateVirusScanAgain")
	return
}

func NewCreateVirusScanAgainResponse() (response *CreateVirusScanAgainResponse) {
	response = &CreateVirusScanAgainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时文件查杀重新检测
func (c *Client) CreateVirusScanAgain(request *CreateVirusScanAgainRequest) (response *CreateVirusScanAgainResponse, err error) {
	if request == nil {
		request = NewCreateVirusScanAgainRequest()
	}
	response = NewCreateVirusScanAgainResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEscapeEventsExportJobRequest() (request *CreateEscapeEventsExportJobRequest) {
	request = &CreateEscapeEventsExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateEscapeEventsExportJob")
	return
}

func NewCreateEscapeEventsExportJobResponse() (response *CreateEscapeEventsExportJobResponse) {
	response = &CreateEscapeEventsExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建逃逸事件导出异步任务
func (c *Client) CreateEscapeEventsExportJob(request *CreateEscapeEventsExportJobRequest) (response *CreateEscapeEventsExportJobResponse, err error) {
	if request == nil {
		request = NewCreateEscapeEventsExportJobRequest()
	}
	response = NewCreateEscapeEventsExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSuperNodeMachineRequest() (request *DeleteSuperNodeMachineRequest) {
	request = &DeleteSuperNodeMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteSuperNodeMachine")
	return
}

func NewDeleteSuperNodeMachineResponse() (response *DeleteSuperNodeMachineResponse) {
	response = &DeleteSuperNodeMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 卸载超级节点Agent客户端
func (c *Client) DeleteSuperNodeMachine(request *DeleteSuperNodeMachineRequest) (response *DeleteSuperNodeMachineResponse, err error) {
	if request == nil {
		request = NewDeleteSuperNodeMachineRequest()
	}
	response = NewDeleteSuperNodeMachineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublicProxyInstallCommandRequest() (request *DescribePublicProxyInstallCommandRequest) {
	request = &DescribePublicProxyInstallCommandRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribePublicProxyInstallCommand")
	return
}

func NewDescribePublicProxyInstallCommandResponse() (response *DescribePublicProxyInstallCommandResponse) {
	response = &DescribePublicProxyInstallCommandResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取公共代理安装命令
func (c *Client) DescribePublicProxyInstallCommand(request *DescribePublicProxyInstallCommandRequest) (response *DescribePublicProxyInstallCommandResponse, err error) {
	if request == nil {
		request = NewDescribePublicProxyInstallCommandRequest()
	}
	response = NewDescribePublicProxyInstallCommandResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetImageScanStopRequest() (request *ModifyAssetImageScanStopRequest) {
	request = &ModifyAssetImageScanStopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyAssetImageScanStop")
	return
}

func NewModifyAssetImageScanStopResponse() (response *ModifyAssetImageScanStopResponse) {
	response = &ModifyAssetImageScanStopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全停止镜像扫描
func (c *Client) ModifyAssetImageScanStop(request *ModifyAssetImageScanStopRequest) (response *ModifyAssetImageScanStopResponse, err error) {
	if request == nil {
		request = NewModifyAssetImageScanStopRequest()
	}
	response = NewModifyAssetImageScanStopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRiskListExportRequest() (request *DescribeAssetImageRiskListExportRequest) {
	request = &DescribeAssetImageRiskListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRiskListExport")
	return
}

func NewDescribeAssetImageRiskListExportResponse() (response *DescribeAssetImageRiskListExportResponse) {
	response = &DescribeAssetImageRiskListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全搜索查询镜像风险列表导出
func (c *Client) DescribeAssetImageRiskListExport(request *DescribeAssetImageRiskListExportRequest) (response *DescribeAssetImageRiskListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRiskListExportRequest()
	}
	response = NewDescribeAssetImageRiskListExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryScanStatusOneKeyRequest() (request *DescribeAssetImageRegistryScanStatusOneKeyRequest) {
	request = &DescribeAssetImageRegistryScanStatusOneKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryScanStatusOneKey")
	return
}

func NewDescribeAssetImageRegistryScanStatusOneKeyResponse() (response *DescribeAssetImageRegistryScanStatusOneKeyResponse) {
	response = &DescribeAssetImageRegistryScanStatusOneKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询一键镜像扫描状态
func (c *Client) DescribeAssetImageRegistryScanStatusOneKey(request *DescribeAssetImageRegistryScanStatusOneKeyRequest) (response *DescribeAssetImageRegistryScanStatusOneKeyResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryScanStatusOneKeyRequest()
	}
	response = NewDescribeAssetImageRegistryScanStatusOneKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRegistryAutoAuthorizedRepoListRequest() (request *DescribeImageRegistryAutoAuthorizedRepoListRequest) {
	request = &DescribeImageRegistryAutoAuthorizedRepoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRegistryAutoAuthorizedRepoList")
	return
}

func NewDescribeImageRegistryAutoAuthorizedRepoListResponse() (response *DescribeImageRegistryAutoAuthorizedRepoListResponse) {
	response = &DescribeImageRegistryAutoAuthorizedRepoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像仓库自动授权配置待选项目空间和镜像列表
func (c *Client) DescribeImageRegistryAutoAuthorizedRepoList(request *DescribeImageRegistryAutoAuthorizedRepoListRequest) (response *DescribeImageRegistryAutoAuthorizedRepoListResponse, err error) {
	if request == nil {
		request = NewDescribeImageRegistryAutoAuthorizedRepoListRequest()
	}
	response = NewDescribeImageRegistryAutoAuthorizedRepoListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogDeliveryKafkaOptionsRequest() (request *DescribeSecLogDeliveryKafkaOptionsRequest) {
	request = &DescribeSecLogDeliveryKafkaOptionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogDeliveryKafkaOptions")
	return
}

func NewDescribeSecLogDeliveryKafkaOptionsResponse() (response *DescribeSecLogDeliveryKafkaOptionsResponse) {
	response = &DescribeSecLogDeliveryKafkaOptionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志投递kafka可选项
func (c *Client) DescribeSecLogDeliveryKafkaOptions(request *DescribeSecLogDeliveryKafkaOptionsRequest) (response *DescribeSecLogDeliveryKafkaOptionsResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogDeliveryKafkaOptionsRequest()
	}
	response = NewDescribeSecLogDeliveryKafkaOptionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirusMonitorConfigRequest() (request *ModifyVirusMonitorConfigRequest) {
	request = &ModifyVirusMonitorConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusMonitorConfig")
	return
}

func NewModifyVirusMonitorConfigResponse() (response *ModifyVirusMonitorConfigResponse) {
	response = &ModifyVirusMonitorConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时更新文件查杀实时监控设置信息
func (c *Client) ModifyVirusMonitorConfig(request *ModifyVirusMonitorConfigRequest) (response *ModifyVirusMonitorConfigResponse, err error) {
	if request == nil {
		request = NewModifyVirusMonitorConfigRequest()
	}
	response = NewModifyVirusMonitorConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogDeliveryClsOptionsRequest() (request *DescribeSecLogDeliveryClsOptionsRequest) {
	request = &DescribeSecLogDeliveryClsOptionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogDeliveryClsOptions")
	return
}

func NewDescribeSecLogDeliveryClsOptionsResponse() (response *DescribeSecLogDeliveryClsOptionsResponse) {
	response = &DescribeSecLogDeliveryClsOptionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志投递cls可选项
func (c *Client) DescribeSecLogDeliveryClsOptions(request *DescribeSecLogDeliveryClsOptionsRequest) (response *DescribeSecLogDeliveryClsOptionsResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogDeliveryClsOptionsRequest()
	}
	response = NewDescribeSecLogDeliveryClsOptionsResponse()
	err = c.Send(request, response)
	return
}

func NewAddEscapeWhiteListRequest() (request *AddEscapeWhiteListRequest) {
	request = &AddEscapeWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddEscapeWhiteList")
	return
}

func NewAddEscapeWhiteListResponse() (response *AddEscapeWhiteListResponse) {
	response = &AddEscapeWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增逃逸白名单
func (c *Client) AddEscapeWhiteList(request *AddEscapeWhiteListRequest) (response *AddEscapeWhiteListResponse, err error) {
	if request == nil {
		request = NewAddEscapeWhiteListRequest()
	}
	response = NewAddEscapeWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogJoinSuperObjectListRequest() (request *DescribeSecLogJoinSuperObjectListRequest) {
	request = &DescribeSecLogJoinSuperObjectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogJoinSuperObjectList")
	return
}

func NewDescribeSecLogJoinSuperObjectListResponse() (response *DescribeSecLogJoinSuperObjectListResponse) {
	response = &DescribeSecLogJoinSuperObjectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志接入超级对象列表
func (c *Client) DescribeSecLogJoinSuperObjectList(request *DescribeSecLogJoinSuperObjectListRequest) (response *DescribeSecLogJoinSuperObjectListResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogJoinSuperObjectListRequest()
	}
	response = NewDescribeSecLogJoinSuperObjectListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryRegistryDetailRequest() (request *DescribeAssetImageRegistryRegistryDetailRequest) {
	request = &DescribeAssetImageRegistryRegistryDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRegistryDetail")
	return
}

func NewDescribeAssetImageRegistryRegistryDetailResponse() (response *DescribeAssetImageRegistryRegistryDetailResponse) {
	response = &DescribeAssetImageRegistryRegistryDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看单个镜像仓库详细信息
func (c *Client) DescribeAssetImageRegistryRegistryDetail(request *DescribeAssetImageRegistryRegistryDetailRequest) (response *DescribeAssetImageRegistryRegistryDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryRegistryDetailRequest()
	}
	response = NewDescribeAssetImageRegistryRegistryDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNewPurchaseInfoRequest() (request *DescribeNewPurchaseInfoRequest) {
	request = &DescribeNewPurchaseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNewPurchaseInfo")
	return
}

func NewDescribeNewPurchaseInfoResponse() (response *DescribeNewPurchaseInfoResponse) {
	response = &DescribeNewPurchaseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询新购信息
func (c *Client) DescribeNewPurchaseInfo(request *DescribeNewPurchaseInfoRequest) (response *DescribeNewPurchaseInfoResponse, err error) {
	if request == nil {
		request = NewDescribeNewPurchaseInfoRequest()
	}
	response = NewDescribeNewPurchaseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceTaskAssetSummaryRequest() (request *DescribeComplianceTaskAssetSummaryRequest) {
	request = &DescribeComplianceTaskAssetSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceTaskAssetSummary")
	return
}

func NewDescribeComplianceTaskAssetSummaryResponse() (response *DescribeComplianceTaskAssetSummaryResponse) {
	response = &DescribeComplianceTaskAssetSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询上次任务的资产通过率汇总信息
func (c *Client) DescribeComplianceTaskAssetSummary(request *DescribeComplianceTaskAssetSummaryRequest) (response *DescribeComplianceTaskAssetSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceTaskAssetSummaryRequest()
	}
	response = NewDescribeComplianceTaskAssetSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageDetailRequest() (request *DescribeAssetImageDetailRequest) {
	request = &DescribeAssetImageDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageDetail")
	return
}

func NewDescribeAssetImageDetailResponse() (response *DescribeAssetImageDetailResponse) {
	response = &DescribeAssetImageDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像详细信息
func (c *Client) DescribeAssetImageDetail(request *DescribeAssetImageDetailRequest) (response *DescribeAssetImageDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageDetailRequest()
	}
	response = NewDescribeAssetImageDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersNamespacesRequest() (request *DescribeClustersNamespacesRequest) {
	request = &DescribeClustersNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClustersNamespaces")
	return
}

func NewDescribeClustersNamespacesResponse() (response *DescribeClustersNamespacesResponse) {
	response = &DescribeClustersNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群的命名空间
func (c *Client) DescribeClustersNamespaces(request *DescribeClustersNamespacesRequest) (response *DescribeClustersNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeClustersNamespacesRequest()
	}
	response = NewDescribeClustersNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSuperNodePodListRequest() (request *DescribeSuperNodePodListRequest) {
	request = &DescribeSuperNodePodListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSuperNodePodList")
	return
}

func NewDescribeSuperNodePodListResponse() (response *DescribeSuperNodePodListResponse) {
	response = &DescribeSuperNodePodListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询超级节点pod列表
func (c *Client) DescribeSuperNodePodList(request *DescribeSuperNodePodListRequest) (response *DescribeSuperNodePodListResponse, err error) {
	if request == nil {
		request = NewDescribeSuperNodePodListRequest()
	}
	response = NewDescribeSuperNodePodListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellSystemPolicyConfigRequest() (request *DescribeReverseShellSystemPolicyConfigRequest) {
	request = &DescribeReverseShellSystemPolicyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellSystemPolicyConfig")
	return
}

func NewDescribeReverseShellSystemPolicyConfigResponse() (response *DescribeReverseShellSystemPolicyConfigResponse) {
	response = &DescribeReverseShellSystemPolicyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询反弹shell系统策略配置
func (c *Client) DescribeReverseShellSystemPolicyConfig(request *DescribeReverseShellSystemPolicyConfigRequest) (response *DescribeReverseShellSystemPolicyConfigResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellSystemPolicyConfigRequest()
	}
	response = NewDescribeReverseShellSystemPolicyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNewestVulRequest() (request *DescribeNewestVulRequest) {
	request = &DescribeNewestVulRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNewestVul")
	return
}

func NewDescribeNewestVulResponse() (response *DescribeNewestVulResponse) {
	response = &DescribeNewestVulResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询最新披露漏洞列表
func (c *Client) DescribeNewestVul(request *DescribeNewestVulRequest) (response *DescribeNewestVulResponse, err error) {
	if request == nil {
		request = NewDescribeNewestVulRequest()
	}
	response = NewDescribeNewestVulResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirusScanTimeoutSettingRequest() (request *ModifyVirusScanTimeoutSettingRequest) {
	request = &ModifyVirusScanTimeoutSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusScanTimeoutSetting")
	return
}

func NewModifyVirusScanTimeoutSettingResponse() (response *ModifyVirusScanTimeoutSettingResponse) {
	response = &ModifyVirusScanTimeoutSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时文件扫描超时设置
func (c *Client) ModifyVirusScanTimeoutSetting(request *ModifyVirusScanTimeoutSettingRequest) (response *ModifyVirusScanTimeoutSettingResponse, err error) {
	if request == nil {
		request = NewModifyVirusScanTimeoutSettingRequest()
	}
	response = NewModifyVirusScanTimeoutSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulScanAuthorizedImageSummaryRequest() (request *DescribeVulScanAuthorizedImageSummaryRequest) {
	request = &DescribeVulScanAuthorizedImageSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulScanAuthorizedImageSummary")
	return
}

func NewDescribeVulScanAuthorizedImageSummaryResponse() (response *DescribeVulScanAuthorizedImageSummaryResponse) {
	response = &DescribeVulScanAuthorizedImageSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统计漏洞扫描页已授权和未扫描镜像数
func (c *Client) DescribeVulScanAuthorizedImageSummary(request *DescribeVulScanAuthorizedImageSummaryRequest) (response *DescribeVulScanAuthorizedImageSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeVulScanAuthorizedImageSummaryRequest()
	}
	response = NewDescribeVulScanAuthorizedImageSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSuperNodeAgentInstallJobRequest() (request *CreateSuperNodeAgentInstallJobRequest) {
	request = &CreateSuperNodeAgentInstallJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateSuperNodeAgentInstallJob")
	return
}

func NewCreateSuperNodeAgentInstallJobResponse() (response *CreateSuperNodeAgentInstallJobResponse) {
	response = &CreateSuperNodeAgentInstallJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建超级节点agent安装任务
func (c *Client) CreateSuperNodeAgentInstallJob(request *CreateSuperNodeAgentInstallJobRequest) (response *CreateSuperNodeAgentInstallJobResponse, err error) {
	if request == nil {
		request = NewCreateSuperNodeAgentInstallJobRequest()
	}
	response = NewCreateSuperNodeAgentInstallJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterAssetSummaryRequest() (request *DescribeClusterAssetSummaryRequest) {
	request = &DescribeClusterAssetSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterAssetSummary")
	return
}

func NewDescribeClusterAssetSummaryResponse() (response *DescribeClusterAssetSummaryResponse) {
	response = &DescribeClusterAssetSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户集群资产的总览信息
func (c *Client) DescribeClusterAssetSummary(request *DescribeClusterAssetSummaryRequest) (response *DescribeClusterAssetSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeClusterAssetSummaryRequest()
	}
	response = NewDescribeClusterAssetSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirusFileStatusRequest() (request *ModifyVirusFileStatusRequest) {
	request = &ModifyVirusFileStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusFileStatus")
	return
}

func NewModifyVirusFileStatusResponse() (response *ModifyVirusFileStatusResponse) {
	response = &ModifyVirusFileStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时更新木马文件事件状态
func (c *Client) ModifyVirusFileStatus(request *ModifyVirusFileStatusRequest) (response *ModifyVirusFileStatusResponse, err error) {
	if request == nil {
		request = NewModifyVirusFileStatusRequest()
	}
	response = NewModifyVirusFileStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCheckNetworkFirewallPolicyYamlRequest() (request *CheckNetworkFirewallPolicyYamlRequest) {
	request = &CheckNetworkFirewallPolicyYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CheckNetworkFirewallPolicyYaml")
	return
}

func NewCheckNetworkFirewallPolicyYamlResponse() (response *CheckNetworkFirewallPolicyYamlResponse) {
	response = &CheckNetworkFirewallPolicyYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建检查Yaml网络策略任务
func (c *Client) CheckNetworkFirewallPolicyYaml(request *CheckNetworkFirewallPolicyYamlRequest) (response *CheckNetworkFirewallPolicyYamlResponse, err error) {
	if request == nil {
		request = NewCheckNetworkFirewallPolicyYamlRequest()
	}
	response = NewCheckNetworkFirewallPolicyYamlResponse()
	err = c.Send(request, response)
	return
}

func NewAddAndPublishNetworkFirewallPolicyDetailRequest() (request *AddAndPublishNetworkFirewallPolicyDetailRequest) {
	request = &AddAndPublishNetworkFirewallPolicyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddAndPublishNetworkFirewallPolicyDetail")
	return
}

func NewAddAndPublishNetworkFirewallPolicyDetailResponse() (response *AddAndPublishNetworkFirewallPolicyDetailResponse) {
	response = &AddAndPublishNetworkFirewallPolicyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建网络策略添加并发布任务
func (c *Client) AddAndPublishNetworkFirewallPolicyDetail(request *AddAndPublishNetworkFirewallPolicyDetailRequest) (response *AddAndPublishNetworkFirewallPolicyDetailResponse, err error) {
	if request == nil {
		request = NewAddAndPublishNetworkFirewallPolicyDetailRequest()
	}
	response = NewAddAndPublishNetworkFirewallPolicyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewAddModReverseShellRegexpWhiteListRequest() (request *AddModReverseShellRegexpWhiteListRequest) {
	request = &AddModReverseShellRegexpWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddModReverseShellRegexpWhiteList")
	return
}

func NewAddModReverseShellRegexpWhiteListResponse() (response *AddModReverseShellRegexpWhiteListResponse) {
	response = &AddModReverseShellRegexpWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加修改反弹shell正则白名单
func (c *Client) AddModReverseShellRegexpWhiteList(request *AddModReverseShellRegexpWhiteListRequest) (response *AddModReverseShellRegexpWhiteListResponse, err error) {
	if request == nil {
		request = NewAddModReverseShellRegexpWhiteListRequest()
	}
	response = NewAddModReverseShellRegexpWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusSummaryRequest() (request *DescribeVirusSummaryRequest) {
	request = &DescribeVirusSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusSummary")
	return
}

func NewDescribeVirusSummaryResponse() (response *DescribeVirusSummaryResponse) {
	response = &DescribeVirusSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询木马概览信息
func (c *Client) DescribeVirusSummary(request *DescribeVirusSummaryRequest) (response *DescribeVirusSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeVirusSummaryRequest()
	}
	response = NewDescribeVirusSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReverseShellEventsRequest() (request *DeleteReverseShellEventsRequest) {
	request = &DeleteReverseShellEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteReverseShellEvents")
	return
}

func NewDeleteReverseShellEventsResponse() (response *DeleteReverseShellEventsResponse) {
	response = &DeleteReverseShellEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运行时反弹shell事件
func (c *Client) DeleteReverseShellEvents(request *DeleteReverseShellEventsRequest) (response *DeleteReverseShellEventsResponse, err error) {
	if request == nil {
		request = NewDeleteReverseShellEventsRequest()
	}
	response = NewDeleteReverseShellEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebVulListRequest() (request *DescribeWebVulListRequest) {
	request = &DescribeWebVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeWebVulList")
	return
}

func NewDescribeWebVulListResponse() (response *DescribeWebVulListResponse) {
	response = &DescribeWebVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询web应用漏洞列表
func (c *Client) DescribeWebVulList(request *DescribeWebVulListRequest) (response *DescribeWebVulListResponse, err error) {
	if request == nil {
		request = NewDescribeWebVulListRequest()
	}
	response = NewDescribeWebVulListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallDetailRequest() (request *DescribeRiskSyscallDetailRequest) {
	request = &DescribeRiskSyscallDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallDetail")
	return
}

func NewDescribeRiskSyscallDetailResponse() (response *DescribeRiskSyscallDetailResponse) {
	response = &DescribeRiskSyscallDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询高危系统调用事件详细信息
func (c *Client) DescribeRiskSyscallDetail(request *DescribeRiskSyscallDetailRequest) (response *DescribeRiskSyscallDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallDetailRequest()
	}
	response = NewDescribeRiskSyscallDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirusMonitorSettingRequest() (request *ModifyVirusMonitorSettingRequest) {
	request = &ModifyVirusMonitorSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusMonitorSetting")
	return
}

func NewModifyVirusMonitorSettingResponse() (response *ModifyVirusMonitorSettingResponse) {
	response = &ModifyVirusMonitorSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时更新文件查杀实时监控设置
func (c *Client) ModifyVirusMonitorSetting(request *ModifyVirusMonitorSettingRequest) (response *ModifyVirusMonitorSettingResponse, err error) {
	if request == nil {
		request = NewModifyVirusMonitorSettingRequest()
	}
	response = NewModifyVirusMonitorSettingResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterDetailExportJobRequest() (request *CreateClusterDetailExportJobRequest) {
	request = &CreateClusterDetailExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateClusterDetailExportJob")
	return
}

func NewCreateClusterDetailExportJobResponse() (response *CreateClusterDetailExportJobResponse) {
	response = &CreateClusterDetailExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群详情导出任务,只支持单个集群的详情导出
func (c *Client) CreateClusterDetailExportJob(request *CreateClusterDetailExportJobRequest) (response *CreateClusterDetailExportJobResponse, err error) {
	if request == nil {
		request = NewCreateClusterDetailExportJobRequest()
	}
	response = NewCreateClusterDetailExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRiskListExportJobRequest() (request *CreateRiskListExportJobRequest) {
	request = &CreateRiskListExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateRiskListExportJob")
	return
}

func NewCreateRiskListExportJobResponse() (response *CreateRiskListExportJobResponse) {
	response = &CreateRiskListExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建风险项导出任务，可指定1个或多个风险导出
func (c *Client) CreateRiskListExportJob(request *CreateRiskListExportJobRequest) (response *CreateRiskListExportJobResponse, err error) {
	if request == nil {
		request = NewCreateRiskListExportJobRequest()
	}
	response = NewCreateRiskListExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserClusterRequest() (request *DeleteUserClusterRequest) {
	request = &DeleteUserClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteUserCluster")
	return
}

func NewDeleteUserClusterResponse() (response *DeleteUserClusterResponse) {
	response = &DeleteUserClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户自建集群
func (c *Client) DeleteUserCluster(request *DeleteUserClusterRequest) (response *DeleteUserClusterResponse, err error) {
	if request == nil {
		request = NewDeleteUserClusterRequest()
	}
	response = NewDeleteUserClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaliciousConnectionBlackListRequest() (request *DeleteMaliciousConnectionBlackListRequest) {
	request = &DeleteMaliciousConnectionBlackListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteMaliciousConnectionBlackList")
	return
}

func NewDeleteMaliciousConnectionBlackListResponse() (response *DeleteMaliciousConnectionBlackListResponse) {
	response = &DeleteMaliciousConnectionBlackListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除恶意请求外连黑名单
func (c *Client) DeleteMaliciousConnectionBlackList(request *DeleteMaliciousConnectionBlackListRequest) (response *DeleteMaliciousConnectionBlackListResponse, err error) {
	if request == nil {
		request = NewDeleteMaliciousConnectionBlackListRequest()
	}
	response = NewDeleteMaliciousConnectionBlackListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDefenceVulExportJobRequest() (request *CreateDefenceVulExportJobRequest) {
	request = &CreateDefenceVulExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateDefenceVulExportJob")
	return
}

func NewCreateDefenceVulExportJobResponse() (response *CreateDefenceVulExportJobResponse) {
	response = &CreateDefenceVulExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建支持防御的漏洞导出任务
func (c *Client) CreateDefenceVulExportJob(request *CreateDefenceVulExportJobRequest) (response *CreateDefenceVulExportJobResponse, err error) {
	if request == nil {
		request = NewCreateDefenceVulExportJobRequest()
	}
	response = NewCreateDefenceVulExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRegistryTimingScanTaskRequest() (request *DescribeImageRegistryTimingScanTaskRequest) {
	request = &DescribeImageRegistryTimingScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRegistryTimingScanTask")
	return
}

func NewDescribeImageRegistryTimingScanTaskResponse() (response *DescribeImageRegistryTimingScanTaskResponse) {
	response = &DescribeImageRegistryTimingScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查看定时任务
func (c *Client) DescribeImageRegistryTimingScanTask(request *DescribeImageRegistryTimingScanTaskRequest) (response *DescribeImageRegistryTimingScanTaskResponse, err error) {
	if request == nil {
		request = NewDescribeImageRegistryTimingScanTaskRequest()
	}
	response = NewDescribeImageRegistryTimingScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAffectedWorkloadListRequest() (request *DescribeAffectedWorkloadListRequest) {
	request = &DescribeAffectedWorkloadListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAffectedWorkloadList")
	return
}

func NewDescribeAffectedWorkloadListResponse() (response *DescribeAffectedWorkloadListResponse) {
	response = &DescribeAffectedWorkloadListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询workload类型的影响范围，返回workload列表
func (c *Client) DescribeAffectedWorkloadList(request *DescribeAffectedWorkloadListRequest) (response *DescribeAffectedWorkloadListResponse, err error) {
	if request == nil {
		request = NewDescribeAffectedWorkloadListRequest()
	}
	response = NewDescribeAffectedWorkloadListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAbnormalProcessRuleStatusRequest() (request *ModifyAbnormalProcessRuleStatusRequest) {
	request = &ModifyAbnormalProcessRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyAbnormalProcessRuleStatus")
	return
}

func NewModifyAbnormalProcessRuleStatusResponse() (response *ModifyAbnormalProcessRuleStatusResponse) {
	response = &ModifyAbnormalProcessRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改运行时异常进程策略的开启关闭状态
func (c *Client) ModifyAbnormalProcessRuleStatus(request *ModifyAbnormalProcessRuleStatusRequest) (response *ModifyAbnormalProcessRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyAbnormalProcessRuleStatusRequest()
	}
	response = NewModifyAbnormalProcessRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulImageListRequest() (request *DescribeVulImageListRequest) {
	request = &DescribeVulImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulImageList")
	return
}

func NewDescribeVulImageListResponse() (response *DescribeVulImageListResponse) {
	response = &DescribeVulImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞影响的镜像列表
func (c *Client) DescribeVulImageList(request *DescribeVulImageListRequest) (response *DescribeVulImageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulImageListRequest()
	}
	response = NewDescribeVulImageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalRuleInfoRequest() (request *DescribeK8sApiAbnormalRuleInfoRequest) {
	request = &DescribeK8sApiAbnormalRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalRuleInfo")
	return
}

func NewDescribeK8sApiAbnormalRuleInfoResponse() (response *DescribeK8sApiAbnormalRuleInfoResponse) {
	response = &DescribeK8sApiAbnormalRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8sapi异常请求规则详情
func (c *Client) DescribeK8sApiAbnormalRuleInfo(request *DescribeK8sApiAbnormalRuleInfoRequest) (response *DescribeK8sApiAbnormalRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalRuleInfoRequest()
	}
	response = NewDescribeK8sApiAbnormalRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlEventsRequest() (request *DescribeAccessControlEventsRequest) {
	request = &DescribeAccessControlEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlEvents")
	return
}

func NewDescribeAccessControlEventsResponse() (response *DescribeAccessControlEventsResponse) {
	response = &DescribeAccessControlEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时访问控制事件列表
func (c *Client) DescribeAccessControlEvents(request *DescribeAccessControlEventsRequest) (response *DescribeAccessControlEventsResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlEventsRequest()
	}
	response = NewDescribeAccessControlEventsResponse()
	err = c.Send(request, response)
	return
}

func NewStopVirusScanTaskRequest() (request *StopVirusScanTaskRequest) {
	request = &StopVirusScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "StopVirusScanTask")
	return
}

func NewStopVirusScanTaskResponse() (response *StopVirusScanTaskResponse) {
	response = &StopVirusScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时停止木马查杀任务
func (c *Client) StopVirusScanTask(request *StopVirusScanTaskRequest) (response *StopVirusScanTaskResponse, err error) {
	if request == nil {
		request = NewStopVirusScanTaskRequest()
	}
	response = NewStopVirusScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallClusterContainerSecurityRequest() (request *UninstallClusterContainerSecurityRequest) {
	request = &UninstallClusterContainerSecurityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "UninstallClusterContainerSecurity")
	return
}

func NewUninstallClusterContainerSecurityResponse() (response *UninstallClusterContainerSecurityResponse) {
	response = &UninstallClusterContainerSecurityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 卸载集群容器安全
func (c *Client) UninstallClusterContainerSecurity(request *UninstallClusterContainerSecurityRequest) (response *UninstallClusterContainerSecurityResponse, err error) {
	if request == nil {
		request = NewUninstallClusterContainerSecurityRequest()
	}
	response = NewUninstallClusterContainerSecurityResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEscapeRegexpRuleStatusRequest() (request *ModifyEscapeRegexpRuleStatusRequest) {
	request = &ModifyEscapeRegexpRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyEscapeRegexpRuleStatus")
	return
}

func NewModifyEscapeRegexpRuleStatusResponse() (response *ModifyEscapeRegexpRuleStatusResponse) {
	response = &ModifyEscapeRegexpRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改容器逃逸正则规则状态
func (c *Client) ModifyEscapeRegexpRuleStatus(request *ModifyEscapeRegexpRuleStatusRequest) (response *ModifyEscapeRegexpRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyEscapeRegexpRuleStatusRequest()
	}
	response = NewModifyEscapeRegexpRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirusScanSettingRequest() (request *ModifyVirusScanSettingRequest) {
	request = &ModifyVirusScanSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusScanSetting")
	return
}

func NewModifyVirusScanSettingResponse() (response *ModifyVirusScanSettingResponse) {
	response = &ModifyVirusScanSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时更新文件查杀设置
func (c *Client) ModifyVirusScanSetting(request *ModifyVirusScanSettingRequest) (response *ModifyVirusScanSettingResponse, err error) {
	if request == nil {
		request = NewModifyVirusScanSettingRequest()
	}
	response = NewModifyVirusScanSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusManualScanEstimateTimeoutRequest() (request *DescribeVirusManualScanEstimateTimeoutRequest) {
	request = &DescribeVirusManualScanEstimateTimeoutRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusManualScanEstimateTimeout")
	return
}

func NewDescribeVirusManualScanEstimateTimeoutResponse() (response *DescribeVirusManualScanEstimateTimeoutResponse) {
	response = &DescribeVirusManualScanEstimateTimeoutResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马一键检测预估超时时间
func (c *Client) DescribeVirusManualScanEstimateTimeout(request *DescribeVirusManualScanEstimateTimeoutRequest) (response *DescribeVirusManualScanEstimateTimeoutResponse, err error) {
	if request == nil {
		request = NewDescribeVirusManualScanEstimateTimeoutRequest()
	}
	response = NewDescribeVirusManualScanEstimateTimeoutResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkFirewallPublishRequest() (request *CreateNetworkFirewallPublishRequest) {
	request = &CreateNetworkFirewallPublishRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateNetworkFirewallPublish")
	return
}

func NewCreateNetworkFirewallPublishResponse() (response *CreateNetworkFirewallPublishResponse) {
	response = &CreateNetworkFirewallPublishResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建网络策略发布任务
func (c *Client) CreateNetworkFirewallPublish(request *CreateNetworkFirewallPublishRequest) (response *CreateNetworkFirewallPublishResponse, err error) {
	if request == nil {
		request = NewCreateNetworkFirewallPublishRequest()
	}
	response = NewCreateNetworkFirewallPublishResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusTaskListRequest() (request *DescribeVirusTaskListRequest) {
	request = &DescribeVirusTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusTaskList")
	return
}

func NewDescribeVirusTaskListResponse() (response *DescribeVirusTaskListResponse) {
	response = &DescribeVirusTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询文件查杀任务列表
func (c *Client) DescribeVirusTaskList(request *DescribeVirusTaskListRequest) (response *DescribeVirusTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeVirusTaskListRequest()
	}
	response = NewDescribeVirusTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebServiceListRequest() (request *DescribeAssetWebServiceListRequest) {
	request = &DescribeAssetWebServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetWebServiceList")
	return
}

func NewDescribeAssetWebServiceListResponse() (response *DescribeAssetWebServiceListResponse) {
	response = &DescribeAssetWebServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全查询web服务列表
func (c *Client) DescribeAssetWebServiceList(request *DescribeAssetWebServiceListRequest) (response *DescribeAssetWebServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebServiceListRequest()
	}
	response = NewDescribeAssetWebServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNameListRequest() (request *DescribeClusterNameListRequest) {
	request = &DescribeClusterNameListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterNameList")
	return
}

func NewDescribeClusterNameListResponse() (response *DescribeClusterNameListResponse) {
	response = &DescribeClusterNameListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户所有集群名字
func (c *Client) DescribeClusterNameList(request *DescribeClusterNameListRequest) (response *DescribeClusterNameListResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNameListRequest()
	}
	response = NewDescribeClusterNameListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlRulesRequest() (request *DescribeAccessControlRulesRequest) {
	request = &DescribeAccessControlRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlRules")
	return
}

func NewDescribeAccessControlRulesResponse() (response *DescribeAccessControlRulesResponse) {
	response = &DescribeAccessControlRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行访问控制策略列表信息
func (c *Client) DescribeAccessControlRules(request *DescribeAccessControlRulesRequest) (response *DescribeAccessControlRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlRulesRequest()
	}
	response = NewDescribeAccessControlRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportJobResultRequest() (request *DescribeExportJobResultRequest) {
	request = &DescribeExportJobResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeExportJobResult")
	return
}

func NewDescribeExportJobResultResponse() (response *DescribeExportJobResultResponse) {
	response = &DescribeExportJobResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询导出任务的结果
func (c *Client) DescribeExportJobResult(request *DescribeExportJobResultRequest) (response *DescribeExportJobResultResponse, err error) {
	if request == nil {
		request = NewDescribeExportJobResultRequest()
	}
	response = NewDescribeExportJobResultResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEmergencyVulExportJobRequest() (request *CreateEmergencyVulExportJobRequest) {
	request = &CreateEmergencyVulExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateEmergencyVulExportJob")
	return
}

func NewCreateEmergencyVulExportJobResponse() (response *CreateEmergencyVulExportJobResponse) {
	response = &CreateEmergencyVulExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建应急漏洞导出任务
func (c *Client) CreateEmergencyVulExportJob(request *CreateEmergencyVulExportJobRequest) (response *CreateEmergencyVulExportJobResponse, err error) {
	if request == nil {
		request = NewCreateEmergencyVulExportJobRequest()
	}
	response = NewCreateEmergencyVulExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteComplianceAssetPolicySetFromWhitelistRequest() (request *DeleteComplianceAssetPolicySetFromWhitelistRequest) {
	request = &DeleteComplianceAssetPolicySetFromWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteComplianceAssetPolicySetFromWhitelist")
	return
}

func NewDeleteComplianceAssetPolicySetFromWhitelistResponse() (response *DeleteComplianceAssetPolicySetFromWhitelistResponse) {
	response = &DeleteComplianceAssetPolicySetFromWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 移除安全合规忽略(资产+检测项)列表，不显示指定的检查项包含的资产内容
// 参考的AddCompliancePolicyAssetSetToWhitelist，除输入字段外，其它应该是一致的，如果有不同可能是定义的不对
func (c *Client) DeleteComplianceAssetPolicySetFromWhitelist(request *DeleteComplianceAssetPolicySetFromWhitelistRequest) (response *DeleteComplianceAssetPolicySetFromWhitelistResponse, err error) {
	if request == nil {
		request = NewDeleteComplianceAssetPolicySetFromWhitelistRequest()
	}
	response = NewDeleteComplianceAssetPolicySetFromWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulDefenceSettingRequest() (request *ModifyVulDefenceSettingRequest) {
	request = &ModifyVulDefenceSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVulDefenceSetting")
	return
}

func NewModifyVulDefenceSettingResponse() (response *ModifyVulDefenceSettingResponse) {
	response = &ModifyVulDefenceSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑漏洞防御设置
func (c *Client) ModifyVulDefenceSetting(request *ModifyVulDefenceSettingRequest) (response *ModifyVulDefenceSettingResponse, err error) {
	if request == nil {
		request = NewModifyVulDefenceSettingRequest()
	}
	response = NewModifyVulDefenceSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeESHitsRequest() (request *DescribeESHitsRequest) {
	request = &DescribeESHitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeESHits")
	return
}

func NewDescribeESHitsResponse() (response *DescribeESHitsResponse) {
	response = &DescribeESHitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ES查询文档列表
func (c *Client) DescribeESHits(request *DescribeESHitsRequest) (response *DescribeESHitsResponse, err error) {
	if request == nil {
		request = NewDescribeESHitsRequest()
	}
	response = NewDescribeESHitsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTcrRegistryDetailRequest() (request *UpdateTcrRegistryDetailRequest) {
	request = &UpdateTcrRegistryDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "UpdateTcrRegistryDetail")
	return
}

func NewUpdateTcrRegistryDetailResponse() (response *UpdateTcrRegistryDetailResponse) {
	response = &UpdateTcrRegistryDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新TCR镜像仓库配置
func (c *Client) UpdateTcrRegistryDetail(request *UpdateTcrRegistryDetailRequest) (response *UpdateTcrRegistryDetailResponse, err error) {
	if request == nil {
		request = NewUpdateTcrRegistryDetailRequest()
	}
	response = NewUpdateTcrRegistryDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessRulesRequest() (request *DescribeAbnormalProcessRulesRequest) {
	request = &DescribeAbnormalProcessRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessRules")
	return
}

func NewDescribeAbnormalProcessRulesResponse() (response *DescribeAbnormalProcessRulesResponse) {
	response = &DescribeAbnormalProcessRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时异常进程策略列表信息
func (c *Client) DescribeAbnormalProcessRules(request *DescribeAbnormalProcessRulesRequest) (response *DescribeAbnormalProcessRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessRulesRequest()
	}
	response = NewDescribeAbnormalProcessRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellInnerIPShowStatusRequest() (request *DescribeReverseShellInnerIPShowStatusRequest) {
	request = &DescribeReverseShellInnerIPShowStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellInnerIPShowStatus")
	return
}

func NewDescribeReverseShellInnerIPShowStatusResponse() (response *DescribeReverseShellInnerIPShowStatusResponse) {
	response = &DescribeReverseShellInnerIPShowStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询反弹shell内网IP展示状态
func (c *Client) DescribeReverseShellInnerIPShowStatus(request *DescribeReverseShellInnerIPShowStatusRequest) (response *DescribeReverseShellInnerIPShowStatusResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellInnerIPShowStatusRequest()
	}
	response = NewDescribeReverseShellInnerIPShowStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirusAutoIsolateSettingRequest() (request *ModifyVirusAutoIsolateSettingRequest) {
	request = &ModifyVirusAutoIsolateSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusAutoIsolateSetting")
	return
}

func NewModifyVirusAutoIsolateSettingResponse() (response *ModifyVirusAutoIsolateSettingResponse) {
	response = &ModifyVirusAutoIsolateSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改木马自动隔离设置
func (c *Client) ModifyVirusAutoIsolateSetting(request *ModifyVirusAutoIsolateSettingRequest) (response *ModifyVirusAutoIsolateSettingResponse, err error) {
	if request == nil {
		request = NewModifyVirusAutoIsolateSettingRequest()
	}
	response = NewModifyVirusAutoIsolateSettingResponse()
	err = c.Send(request, response)
	return
}

func NewAddCompliancePolicyAssetSetToWhitelistRequest() (request *AddCompliancePolicyAssetSetToWhitelistRequest) {
	request = &AddCompliancePolicyAssetSetToWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddCompliancePolicyAssetSetToWhitelist")
	return
}

func NewAddCompliancePolicyAssetSetToWhitelistResponse() (response *AddCompliancePolicyAssetSetToWhitelistResponse) {
	response = &AddCompliancePolicyAssetSetToWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增安全合规忽略(检测项+资产)列表，不显示指定的检查项包含的资产内容
// 参考的AddCompliancePolicyItemToWhitelist，除输入字段外，其它应该是一致的，如果有不同可能是定义的不对
func (c *Client) AddCompliancePolicyAssetSetToWhitelist(request *AddCompliancePolicyAssetSetToWhitelistRequest) (response *AddCompliancePolicyAssetSetToWhitelistResponse, err error) {
	if request == nil {
		request = NewAddCompliancePolicyAssetSetToWhitelistRequest()
	}
	response = NewAddCompliancePolicyAssetSetToWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewResetSecLogTopicConfigRequest() (request *ResetSecLogTopicConfigRequest) {
	request = &ResetSecLogTopicConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ResetSecLogTopicConfig")
	return
}

func NewResetSecLogTopicConfigResponse() (response *ResetSecLogTopicConfigResponse) {
	response = &ResetSecLogTopicConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置安全日志主题设置
func (c *Client) ResetSecLogTopicConfig(request *ResetSecLogTopicConfigRequest) (response *ResetSecLogTopicConfigResponse, err error) {
	if request == nil {
		request = NewResetSecLogTopicConfigRequest()
	}
	response = NewResetSecLogTopicConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCheckItemListRequest() (request *DescribeCheckItemListRequest) {
	request = &DescribeCheckItemListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeCheckItemList")
	return
}

func NewDescribeCheckItemListResponse() (response *DescribeCheckItemListResponse) {
	response = &DescribeCheckItemListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有检查项接口，返回总数和检查项列表
func (c *Client) DescribeCheckItemList(request *DescribeCheckItemListRequest) (response *DescribeCheckItemListResponse, err error) {
	if request == nil {
		request = NewDescribeCheckItemListRequest()
	}
	response = NewDescribeCheckItemListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInspectionReportRequest() (request *DescribeInspectionReportRequest) {
	request = &DescribeInspectionReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeInspectionReport")
	return
}

func NewDescribeInspectionReportResponse() (response *DescribeInspectionReportResponse) {
	response = &DescribeInspectionReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询检查报告
func (c *Client) DescribeInspectionReport(request *DescribeInspectionReportRequest) (response *DescribeInspectionReportResponse, err error) {
	if request == nil {
		request = NewDescribeInspectionReportRequest()
	}
	response = NewDescribeInspectionReportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageComponentListRequest() (request *DescribeImageComponentListRequest) {
	request = &DescribeImageComponentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageComponentList")
	return
}

func NewDescribeImageComponentListResponse() (response *DescribeImageComponentListResponse) {
	response = &DescribeImageComponentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地镜像组件列表
func (c *Client) DescribeImageComponentList(request *DescribeImageComponentListRequest) (response *DescribeImageComponentListResponse, err error) {
	if request == nil {
		request = NewDescribeImageComponentListRequest()
	}
	response = NewDescribeImageComponentListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRegistryNamespaceListRequest() (request *DescribeImageRegistryNamespaceListRequest) {
	request = &DescribeImageRegistryNamespaceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRegistryNamespaceList")
	return
}

func NewDescribeImageRegistryNamespaceListResponse() (response *DescribeImageRegistryNamespaceListResponse) {
	response = &DescribeImageRegistryNamespaceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户镜像仓库下的命令空间列表
func (c *Client) DescribeImageRegistryNamespaceList(request *DescribeImageRegistryNamespaceListRequest) (response *DescribeImageRegistryNamespaceListResponse, err error) {
	if request == nil {
		request = NewDescribeImageRegistryNamespaceListRequest()
	}
	response = NewDescribeImageRegistryNamespaceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnfinishRefreshTaskRequest() (request *DescribeUnfinishRefreshTaskRequest) {
	request = &DescribeUnfinishRefreshTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeUnfinishRefreshTask")
	return
}

func NewDescribeUnfinishRefreshTaskResponse() (response *DescribeUnfinishRefreshTaskResponse) {
	response = &DescribeUnfinishRefreshTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询未完成的刷新资产任务信息
func (c *Client) DescribeUnfinishRefreshTask(request *DescribeUnfinishRefreshTaskRequest) (response *DescribeUnfinishRefreshTaskResponse, err error) {
	if request == nil {
		request = NewDescribeUnfinishRefreshTaskRequest()
	}
	response = NewDescribeUnfinishRefreshTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirusScanConfigRequest() (request *ModifyVirusScanConfigRequest) {
	request = &ModifyVirusScanConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusScanConfig")
	return
}

func NewModifyVirusScanConfigResponse() (response *ModifyVirusScanConfigResponse) {
	response = &ModifyVirusScanConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时更新文件查杀新设置
func (c *Client) ModifyVirusScanConfig(request *ModifyVirusScanConfigRequest) (response *ModifyVirusScanConfigResponse, err error) {
	if request == nil {
		request = NewModifyVirusScanConfigRequest()
	}
	response = NewModifyVirusScanConfigResponse()
	err = c.Send(request, response)
	return
}

func NewScanCompliancePolicyItemsRequest() (request *ScanCompliancePolicyItemsRequest) {
	request = &ScanCompliancePolicyItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ScanCompliancePolicyItems")
	return
}

func NewScanCompliancePolicyItemsResponse() (response *ScanCompliancePolicyItemsResponse) {
	response = &ScanCompliancePolicyItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新检测选的检测项下的所有资产，返回创建的合规检查任务的ID。
func (c *Client) ScanCompliancePolicyItems(request *ScanCompliancePolicyItemsRequest) (response *ScanCompliancePolicyItemsResponse, err error) {
	if request == nil {
		request = NewScanCompliancePolicyItemsRequest()
	}
	response = NewScanCompliancePolicyItemsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePodContainersRequest() (request *DescribePodContainersRequest) {
	request = &DescribePodContainersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribePodContainers")
	return
}

func NewDescribePodContainersResponse() (response *DescribePodContainersResponse) {
	response = &DescribePodContainersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询pod对应的容器信息
func (c *Client) DescribePodContainers(request *DescribePodContainersRequest) (response *DescribePodContainersResponse, err error) {
	if request == nil {
		request = NewDescribePodContainersRequest()
	}
	response = NewDescribePodContainersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceWhitelistItemListRequest() (request *DescribeComplianceWhitelistItemListRequest) {
	request = &DescribeComplianceWhitelistItemListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceWhitelistItemList")
	return
}

func NewDescribeComplianceWhitelistItemListResponse() (response *DescribeComplianceWhitelistItemListResponse) {
	response = &DescribeComplianceWhitelistItemListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询白名单列表
func (c *Client) DescribeComplianceWhitelistItemList(request *DescribeComplianceWhitelistItemListRequest) (response *DescribeComplianceWhitelistItemListResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceWhitelistItemListRequest()
	}
	response = NewDescribeComplianceWhitelistItemListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRefreshTaskRequest() (request *CreateRefreshTaskRequest) {
	request = &CreateRefreshTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateRefreshTask")
	return
}

func NewCreateRefreshTaskResponse() (response *CreateRefreshTaskResponse) {
	response = &CreateRefreshTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下发刷新任务，会刷新资产信息
func (c *Client) CreateRefreshTask(request *CreateRefreshTaskRequest) (response *CreateRefreshTaskResponse, err error) {
	if request == nil {
		request = NewCreateRefreshTaskRequest()
	}
	response = NewCreateRefreshTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetImageRegistryScanStopRequest() (request *ModifyAssetImageRegistryScanStopRequest) {
	request = &ModifyAssetImageRegistryScanStopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyAssetImageRegistryScanStop")
	return
}

func NewModifyAssetImageRegistryScanStopResponse() (response *ModifyAssetImageRegistryScanStopResponse) {
	response = &ModifyAssetImageRegistryScanStopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库停止镜像扫描任务
func (c *Client) ModifyAssetImageRegistryScanStop(request *ModifyAssetImageRegistryScanStopRequest) (response *ModifyAssetImageRegistryScanStopResponse, err error) {
	if request == nil {
		request = NewModifyAssetImageRegistryScanStopRequest()
	}
	response = NewModifyAssetImageRegistryScanStopResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetworkFirewallPolicyDetailRequest() (request *DeleteNetworkFirewallPolicyDetailRequest) {
	request = &DeleteNetworkFirewallPolicyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteNetworkFirewallPolicyDetail")
	return
}

func NewDeleteNetworkFirewallPolicyDetailResponse() (response *DeleteNetworkFirewallPolicyDetailResponse) {
	response = &DeleteNetworkFirewallPolicyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建网络策略删除任务
func (c *Client) DeleteNetworkFirewallPolicyDetail(request *DeleteNetworkFirewallPolicyDetailRequest) (response *DeleteNetworkFirewallPolicyDetailResponse, err error) {
	if request == nil {
		request = NewDeleteNetworkFirewallPolicyDetailRequest()
	}
	response = NewDeleteNetworkFirewallPolicyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetContainerPortListRequest() (request *DescribeAssetContainerPortListRequest) {
	request = &DescribeAssetContainerPortListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetContainerPortList")
	return
}

func NewDescribeAssetContainerPortListResponse() (response *DescribeAssetContainerPortListResponse) {
	response = &DescribeAssetContainerPortListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器端口列表
func (c *Client) DescribeAssetContainerPortList(request *DescribeAssetContainerPortListRequest) (response *DescribeAssetContainerPortListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetContainerPortListRequest()
	}
	response = NewDescribeAssetContainerPortListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkTopologyWorkLoadDetailRequest() (request *DescribeNetworkTopologyWorkLoadDetailRequest) {
	request = &DescribeNetworkTopologyWorkLoadDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkTopologyWorkLoadDetail")
	return
}

func NewDescribeNetworkTopologyWorkLoadDetailResponse() (response *DescribeNetworkTopologyWorkLoadDetailResponse) {
	response = &DescribeNetworkTopologyWorkLoadDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络查询网络拓扑负载详情
func (c *Client) DescribeNetworkTopologyWorkLoadDetail(request *DescribeNetworkTopologyWorkLoadDetailRequest) (response *DescribeNetworkTopologyWorkLoadDetailResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkTopologyWorkLoadDetailRequest()
	}
	response = NewDescribeNetworkTopologyWorkLoadDetailResponse()
	err = c.Send(request, response)
	return
}

func NewAddEditRiskSyscallWhiteListRequest() (request *AddEditRiskSyscallWhiteListRequest) {
	request = &AddEditRiskSyscallWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddEditRiskSyscallWhiteList")
	return
}

func NewAddEditRiskSyscallWhiteListResponse() (response *AddEditRiskSyscallWhiteListResponse) {
	response = &AddEditRiskSyscallWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加编辑运行时高危系统调用白名单
func (c *Client) AddEditRiskSyscallWhiteList(request *AddEditRiskSyscallWhiteListRequest) (response *AddEditRiskSyscallWhiteListResponse, err error) {
	if request == nil {
		request = NewAddEditRiskSyscallWhiteListRequest()
	}
	response = NewAddEditRiskSyscallWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalTendencyRequest() (request *DescribeK8sApiAbnormalTendencyRequest) {
	request = &DescribeK8sApiAbnormalTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalTendency")
	return
}

func NewDescribeK8sApiAbnormalTendencyResponse() (response *DescribeK8sApiAbnormalTendencyResponse) {
	response = &DescribeK8sApiAbnormalTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8sapi异常事件趋势
func (c *Client) DescribeK8sApiAbnormalTendency(request *DescribeK8sApiAbnormalTendencyRequest) (response *DescribeK8sApiAbnormalTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalTendencyRequest()
	}
	response = NewDescribeK8sApiAbnormalTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallPodLabelsListRequest() (request *DescribeNetworkFirewallPodLabelsListRequest) {
	request = &DescribeNetworkFirewallPodLabelsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPodLabelsList")
	return
}

func NewDescribeNetworkFirewallPodLabelsListResponse() (response *DescribeNetworkFirewallPodLabelsListResponse) {
	response = &DescribeNetworkFirewallPodLabelsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群网络pod标签
func (c *Client) DescribeNetworkFirewallPodLabelsList(request *DescribeNetworkFirewallPodLabelsListRequest) (response *DescribeNetworkFirewallPodLabelsListResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallPodLabelsListRequest()
	}
	response = NewDescribeNetworkFirewallPodLabelsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageAutoAuthorizedRuleRequest() (request *DescribeImageAutoAuthorizedRuleRequest) {
	request = &DescribeImageAutoAuthorizedRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageAutoAuthorizedRule")
	return
}

func NewDescribeImageAutoAuthorizedRuleResponse() (response *DescribeImageAutoAuthorizedRuleResponse) {
	response = &DescribeImageAutoAuthorizedRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地镜像自动授权规则
func (c *Client) DescribeImageAutoAuthorizedRule(request *DescribeImageAutoAuthorizedRuleRequest) (response *DescribeImageAutoAuthorizedRuleResponse, err error) {
	if request == nil {
		request = NewDescribeImageAutoAuthorizedRuleRequest()
	}
	response = NewDescribeImageAutoAuthorizedRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryImageLayerDetailRequest() (request *DescribeAssetImageRegistryImageLayerDetailRequest) {
	request = &DescribeAssetImageRegistryImageLayerDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryImageLayerDetail")
	return
}

func NewDescribeAssetImageRegistryImageLayerDetailResponse() (response *DescribeAssetImageRegistryImageLayerDetailResponse) {
	response = &DescribeAssetImageRegistryImageLayerDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像层信息接口
func (c *Client) DescribeAssetImageRegistryImageLayerDetail(request *DescribeAssetImageRegistryImageLayerDetailRequest) (response *DescribeAssetImageRegistryImageLayerDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryImageLayerDetailRequest()
	}
	response = NewDescribeAssetImageRegistryImageLayerDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSummaryRequest() (request *DescribeAssetSummaryRequest) {
	request = &DescribeAssetSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetSummary")
	return
}

func NewDescribeAssetSummaryResponse() (response *DescribeAssetSummaryResponse) {
	response = &DescribeAssetSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账户容器、镜像等统计信息
func (c *Client) DescribeAssetSummary(request *DescribeAssetSummaryRequest) (response *DescribeAssetSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSummaryRequest()
	}
	response = NewDescribeAssetSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellDetailRequest() (request *DescribeReverseShellDetailRequest) {
	request = &DescribeReverseShellDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellDetail")
	return
}

func NewDescribeReverseShellDetailResponse() (response *DescribeReverseShellDetailResponse) {
	response = &DescribeReverseShellDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时反弹shell事件详细信息
func (c *Client) DescribeReverseShellDetail(request *DescribeReverseShellDetailRequest) (response *DescribeReverseShellDetailResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellDetailRequest()
	}
	response = NewDescribeReverseShellDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkTopologyFlowListRequest() (request *DescribeNetworkTopologyFlowListRequest) {
	request = &DescribeNetworkTopologyFlowListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkTopologyFlowList")
	return
}

func NewDescribeNetworkTopologyFlowListResponse() (response *DescribeNetworkTopologyFlowListResponse) {
	response = &DescribeNetworkTopologyFlowListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络查查询网络拓扑流量列表
func (c *Client) DescribeNetworkTopologyFlowList(request *DescribeNetworkTopologyFlowListRequest) (response *DescribeNetworkTopologyFlowListResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkTopologyFlowListRequest()
	}
	response = NewDescribeNetworkTopologyFlowListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceAssetDetailInfoRequest() (request *DescribeComplianceAssetDetailInfoRequest) {
	request = &DescribeComplianceAssetDetailInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceAssetDetailInfo")
	return
}

func NewDescribeComplianceAssetDetailInfoResponse() (response *DescribeComplianceAssetDetailInfoResponse) {
	response = &DescribeComplianceAssetDetailInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个资产的详情
func (c *Client) DescribeComplianceAssetDetailInfo(request *DescribeComplianceAssetDetailInfoRequest) (response *DescribeComplianceAssetDetailInfoResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceAssetDetailInfoRequest()
	}
	response = NewDescribeComplianceAssetDetailInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSuperNodeListRequest() (request *DescribeAssetSuperNodeListRequest) {
	request = &DescribeAssetSuperNodeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetSuperNodeList")
	return
}

func NewDescribeAssetSuperNodeListResponse() (response *DescribeAssetSuperNodeListResponse) {
	response = &DescribeAssetSuperNodeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询超级节点列表
func (c *Client) DescribeAssetSuperNodeList(request *DescribeAssetSuperNodeListRequest) (response *DescribeAssetSuperNodeListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSuperNodeListRequest()
	}
	response = NewDescribeAssetSuperNodeListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSuperNodeExportJobRequest() (request *CreateSuperNodeExportJobRequest) {
	request = &CreateSuperNodeExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateSuperNodeExportJob")
	return
}

func NewCreateSuperNodeExportJobResponse() (response *CreateSuperNodeExportJobResponse) {
	response = &CreateSuperNodeExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建超级节点列表导出
func (c *Client) CreateSuperNodeExportJob(request *CreateSuperNodeExportJobRequest) (response *CreateSuperNodeExportJobResponse, err error) {
	if request == nil {
		request = NewCreateSuperNodeExportJobRequest()
	}
	response = NewCreateSuperNodeExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentInstallCommandRequest() (request *DescribeAgentInstallCommandRequest) {
	request = &DescribeAgentInstallCommandRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAgentInstallCommand")
	return
}

func NewDescribeAgentInstallCommandResponse() (response *DescribeAgentInstallCommandResponse) {
	response = &DescribeAgentInstallCommandResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询agent安装命令
func (c *Client) DescribeAgentInstallCommand(request *DescribeAgentInstallCommandRequest) (response *DescribeAgentInstallCommandResponse, err error) {
	if request == nil {
		request = NewDescribeAgentInstallCommandRequest()
	}
	response = NewDescribeAgentInstallCommandResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveAssetImageRegistryRegistryDetailRequest() (request *RemoveAssetImageRegistryRegistryDetailRequest) {
	request = &RemoveAssetImageRegistryRegistryDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "RemoveAssetImageRegistryRegistryDetail")
	return
}

func NewRemoveAssetImageRegistryRegistryDetailResponse() (response *RemoveAssetImageRegistryRegistryDetailResponse) {
	response = &RemoveAssetImageRegistryRegistryDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除单个镜像仓库详细信息
func (c *Client) RemoveAssetImageRegistryRegistryDetail(request *RemoveAssetImageRegistryRegistryDetailRequest) (response *RemoveAssetImageRegistryRegistryDetailResponse, err error) {
	if request == nil {
		request = NewRemoveAssetImageRegistryRegistryDetailRequest()
	}
	response = NewRemoveAssetImageRegistryRegistryDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogKafkaUINRequest() (request *DescribeSecLogKafkaUINRequest) {
	request = &DescribeSecLogKafkaUINRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogKafkaUIN")
	return
}

func NewDescribeSecLogKafkaUINResponse() (response *DescribeSecLogKafkaUINResponse) {
	response = &DescribeSecLogKafkaUINResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志KafkaUIN
func (c *Client) DescribeSecLogKafkaUIN(request *DescribeSecLogKafkaUINRequest) (response *DescribeSecLogKafkaUINResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogKafkaUINRequest()
	}
	response = NewDescribeSecLogKafkaUINResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventEscapeImageDetailRequest() (request *DescribeEventEscapeImageDetailRequest) {
	request = &DescribeEventEscapeImageDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEventEscapeImageDetail")
	return
}

func NewDescribeEventEscapeImageDetailResponse() (response *DescribeEventEscapeImageDetailResponse) {
	response = &DescribeEventEscapeImageDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询风险容器镜像详情
func (c *Client) DescribeEventEscapeImageDetail(request *DescribeEventEscapeImageDetailRequest) (response *DescribeEventEscapeImageDetailResponse, err error) {
	if request == nil {
		request = NewDescribeEventEscapeImageDetailRequest()
	}
	response = NewDescribeEventEscapeImageDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallNamespaceLabelListRequest() (request *DescribeNetworkFirewallNamespaceLabelListRequest) {
	request = &DescribeNetworkFirewallNamespaceLabelListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallNamespaceLabelList")
	return
}

func NewDescribeNetworkFirewallNamespaceLabelListResponse() (response *DescribeNetworkFirewallNamespaceLabelListResponse) {
	response = &DescribeNetworkFirewallNamespaceLabelListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群网络空间标签列表
func (c *Client) DescribeNetworkFirewallNamespaceLabelList(request *DescribeNetworkFirewallNamespaceLabelListRequest) (response *DescribeNetworkFirewallNamespaceLabelListResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallNamespaceLabelListRequest()
	}
	response = NewDescribeNetworkFirewallNamespaceLabelListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulTopRankingRequest() (request *DescribeVulTopRankingRequest) {
	request = &DescribeVulTopRankingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulTopRanking")
	return
}

func NewDescribeVulTopRankingResponse() (response *DescribeVulTopRankingResponse) {
	response = &DescribeVulTopRankingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞Top排名列表
func (c *Client) DescribeVulTopRanking(request *DescribeVulTopRankingRequest) (response *DescribeVulTopRankingResponse, err error) {
	if request == nil {
		request = NewDescribeVulTopRankingRequest()
	}
	response = NewDescribeVulTopRankingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterRiskRequest() (request *DescribeClusterRiskRequest) {
	request = &DescribeClusterRiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterRisk")
	return
}

func NewDescribeClusterRiskResponse() (response *DescribeClusterRiskResponse) {
	response = &DescribeClusterRiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按集群类别查询风险数量
func (c *Client) DescribeClusterRisk(request *DescribeClusterRiskRequest) (response *DescribeClusterRiskResponse, err error) {
	if request == nil {
		request = NewDescribeClusterRiskRequest()
	}
	response = NewDescribeClusterRiskResponse()
	err = c.Send(request, response)
	return
}

func NewAddEditReverseShellWhiteListRequest() (request *AddEditReverseShellWhiteListRequest) {
	request = &AddEditReverseShellWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddEditReverseShellWhiteList")
	return
}

func NewAddEditReverseShellWhiteListResponse() (response *AddEditReverseShellWhiteListResponse) {
	response = &AddEditReverseShellWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加编辑运行时反弹shell白名单
func (c *Client) AddEditReverseShellWhiteList(request *AddEditReverseShellWhiteListRequest) (response *AddEditReverseShellWhiteListResponse, err error) {
	if request == nil {
		request = NewAddEditReverseShellWhiteListRequest()
	}
	response = NewAddEditReverseShellWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalRuleListRequest() (request *DescribeK8sApiAbnormalRuleListRequest) {
	request = &DescribeK8sApiAbnormalRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalRuleList")
	return
}

func NewDescribeK8sApiAbnormalRuleListResponse() (response *DescribeK8sApiAbnormalRuleListResponse) {
	response = &DescribeK8sApiAbnormalRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8sapi异常请求规则列表
func (c *Client) DescribeK8sApiAbnormalRuleList(request *DescribeK8sApiAbnormalRuleListRequest) (response *DescribeK8sApiAbnormalRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalRuleListRequest()
	}
	response = NewDescribeK8sApiAbnormalRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceEventRequest() (request *DescribeVulDefenceEventRequest) {
	request = &DescribeVulDefenceEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefenceEvent")
	return
}

func NewDescribeVulDefenceEventResponse() (response *DescribeVulDefenceEventResponse) {
	response = &DescribeVulDefenceEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞防御事件列表
func (c *Client) DescribeVulDefenceEvent(request *DescribeVulDefenceEventRequest) (response *DescribeVulDefenceEventResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceEventRequest()
	}
	response = NewDescribeVulDefenceEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchLogsRequest() (request *DescribeSearchLogsRequest) {
	request = &DescribeSearchLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSearchLogs")
	return
}

func NewDescribeSearchLogsResponse() (response *DescribeSearchLogsResponse) {
	response = &DescribeSearchLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取历史搜索记录
func (c *Client) DescribeSearchLogs(request *DescribeSearchLogsRequest) (response *DescribeSearchLogsResponse, err error) {
	if request == nil {
		request = NewDescribeSearchLogsRequest()
	}
	response = NewDescribeSearchLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchExportListRequest() (request *DescribeSearchExportListRequest) {
	request = &DescribeSearchExportListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSearchExportList")
	return
}

func NewDescribeSearchExportListResponse() (response *DescribeSearchExportListResponse) {
	response = &DescribeSearchExportListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出ES查询文档列表
func (c *Client) DescribeSearchExportList(request *DescribeSearchExportListRequest) (response *DescribeSearchExportListResponse, err error) {
	if request == nil {
		request = NewDescribeSearchExportListRequest()
	}
	response = NewDescribeSearchExportListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyImageRegistryAutoAuthorizedStatusRequest() (request *ModifyImageRegistryAutoAuthorizedStatusRequest) {
	request = &ModifyImageRegistryAutoAuthorizedStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyImageRegistryAutoAuthorizedStatus")
	return
}

func NewModifyImageRegistryAutoAuthorizedStatusResponse() (response *ModifyImageRegistryAutoAuthorizedStatusResponse) {
	response = &ModifyImageRegistryAutoAuthorizedStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库自动授权开关
func (c *Client) ModifyImageRegistryAutoAuthorizedStatus(request *ModifyImageRegistryAutoAuthorizedStatusRequest) (response *ModifyImageRegistryAutoAuthorizedStatusResponse, err error) {
	if request == nil {
		request = NewModifyImageRegistryAutoAuthorizedStatusRequest()
	}
	response = NewModifyImageRegistryAutoAuthorizedStatusResponse()
	err = c.Send(request, response)
	return
}

func NewSetComplianceHostRangeRequest() (request *SetComplianceHostRangeRequest) {
	request = &SetComplianceHostRangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "SetComplianceHostRange")
	return
}

func NewSetComplianceHostRangeResponse() (response *SetComplianceHostRangeResponse) {
	response = &SetComplianceHostRangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置合规安全扫描的主机范围
func (c *Client) SetComplianceHostRange(request *SetComplianceHostRangeRequest) (response *SetComplianceHostRangeResponse, err error) {
	if request == nil {
		request = NewSetComplianceHostRangeRequest()
	}
	response = NewSetComplianceHostRangeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostPayDetailRequest() (request *DescribePostPayDetailRequest) {
	request = &DescribePostPayDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribePostPayDetail")
	return
}

func NewDescribePostPayDetailResponse() (response *DescribePostPayDetailResponse) {
	response = &DescribePostPayDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribePostPayDetail  查询后付费详情
func (c *Client) DescribePostPayDetail(request *DescribePostPayDetailRequest) (response *DescribePostPayDetailResponse, err error) {
	if request == nil {
		request = NewDescribePostPayDetailRequest()
	}
	response = NewDescribePostPayDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulScanLocalImageListRequest() (request *DescribeVulScanLocalImageListRequest) {
	request = &DescribeVulScanLocalImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulScanLocalImageList")
	return
}

func NewDescribeVulScanLocalImageListResponse() (response *DescribeVulScanLocalImageListResponse) {
	response = &DescribeVulScanLocalImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞扫描任务的本地镜像列表
func (c *Client) DescribeVulScanLocalImageList(request *DescribeVulScanLocalImageListRequest) (response *DescribeVulScanLocalImageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulScanLocalImageListRequest()
	}
	response = NewDescribeVulScanLocalImageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegistryVulScanAuthorizedImageSummaryRequest() (request *DescribeRegistryVulScanAuthorizedImageSummaryRequest) {
	request = &DescribeRegistryVulScanAuthorizedImageSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRegistryVulScanAuthorizedImageSummary")
	return
}

func NewDescribeRegistryVulScanAuthorizedImageSummaryResponse() (response *DescribeRegistryVulScanAuthorizedImageSummaryResponse) {
	response = &DescribeRegistryVulScanAuthorizedImageSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统计仓库漏洞扫描页已授权和未扫描镜像数
func (c *Client) DescribeRegistryVulScanAuthorizedImageSummary(request *DescribeRegistryVulScanAuthorizedImageSummaryRequest) (response *DescribeRegistryVulScanAuthorizedImageSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeRegistryVulScanAuthorizedImageSummaryRequest()
	}
	response = NewDescribeRegistryVulScanAuthorizedImageSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterIngressYamlRequest() (request *DescribeClusterIngressYamlRequest) {
	request = &DescribeClusterIngressYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterIngressYaml")
	return
}

func NewDescribeClusterIngressYamlResponse() (response *DescribeClusterIngressYamlResponse) {
	response = &DescribeClusterIngressYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群Ingress的yaml文件
func (c *Client) DescribeClusterIngressYaml(request *DescribeClusterIngressYamlRequest) (response *DescribeClusterIngressYamlResponse, err error) {
	if request == nil {
		request = NewDescribeClusterIngressYamlRequest()
	}
	response = NewDescribeClusterIngressYamlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulImageSummaryRequest() (request *DescribeVulImageSummaryRequest) {
	request = &DescribeVulImageSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulImageSummary")
	return
}

func NewDescribeVulImageSummaryResponse() (response *DescribeVulImageSummaryResponse) {
	response = &DescribeVulImageSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞镜像统计
func (c *Client) DescribeVulImageSummary(request *DescribeVulImageSummaryRequest) (response *DescribeVulImageSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeVulImageSummaryRequest()
	}
	response = NewDescribeVulImageSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogVasInfoRequest() (request *DescribeSecLogVasInfoRequest) {
	request = &DescribeSecLogVasInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogVasInfo")
	return
}

func NewDescribeSecLogVasInfoResponse() (response *DescribeSecLogVasInfoResponse) {
	response = &DescribeSecLogVasInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志商品信息
func (c *Client) DescribeSecLogVasInfo(request *DescribeSecLogVasInfoRequest) (response *DescribeSecLogVasInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogVasInfoRequest()
	}
	response = NewDescribeSecLogVasInfoResponse()
	err = c.Send(request, response)
	return
}

func NewInitializeUserComplianceEnvironmentRequest() (request *InitializeUserComplianceEnvironmentRequest) {
	request = &InitializeUserComplianceEnvironmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "InitializeUserComplianceEnvironment")
	return
}

func NewInitializeUserComplianceEnvironmentResponse() (response *InitializeUserComplianceEnvironmentResponse) {
	response = &InitializeUserComplianceEnvironmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为客户初始化合规基线的使用环境，创建必要的数据和选项。
func (c *Client) InitializeUserComplianceEnvironment(request *InitializeUserComplianceEnvironmentRequest) (response *InitializeUserComplianceEnvironmentResponse, err error) {
	if request == nil {
		request = NewInitializeUserComplianceEnvironmentRequest()
	}
	response = NewInitializeUserComplianceEnvironmentResponse()
	err = c.Send(request, response)
	return
}

func NewAddOrModifyMaliciousConnectionWhiteListRequest() (request *AddOrModifyMaliciousConnectionWhiteListRequest) {
	request = &AddOrModifyMaliciousConnectionWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddOrModifyMaliciousConnectionWhiteList")
	return
}

func NewAddOrModifyMaliciousConnectionWhiteListResponse() (response *AddOrModifyMaliciousConnectionWhiteListResponse) {
	response = &AddOrModifyMaliciousConnectionWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加或修改恶意外连的白名单
func (c *Client) AddOrModifyMaliciousConnectionWhiteList(request *AddOrModifyMaliciousConnectionWhiteListRequest) (response *AddOrModifyMaliciousConnectionWhiteListResponse, err error) {
	if request == nil {
		request = NewAddOrModifyMaliciousConnectionWhiteListRequest()
	}
	response = NewAddOrModifyMaliciousConnectionWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableNodesCountRequest() (request *DescribeAvailableNodesCountRequest) {
	request = &DescribeAvailableNodesCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAvailableNodesCount")
	return
}

func NewDescribeAvailableNodesCountResponse() (response *DescribeAvailableNodesCountResponse) {
	response = &DescribeAvailableNodesCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询开启的节点数信息
func (c *Client) DescribeAvailableNodesCount(request *DescribeAvailableNodesCountRequest) (response *DescribeAvailableNodesCountResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableNodesCountRequest()
	}
	response = NewDescribeAvailableNodesCountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetImageRegistryScanTaskRequest() (request *CreateAssetImageRegistryScanTaskRequest) {
	request = &CreateAssetImageRegistryScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageRegistryScanTask")
	return
}

func NewCreateAssetImageRegistryScanTaskResponse() (response *CreateAssetImageRegistryScanTaskResponse) {
	response = &CreateAssetImageRegistryScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库创建镜像扫描任务
func (c *Client) CreateAssetImageRegistryScanTask(request *CreateAssetImageRegistryScanTaskRequest) (response *CreateAssetImageRegistryScanTaskResponse, err error) {
	if request == nil {
		request = NewCreateAssetImageRegistryScanTaskRequest()
	}
	response = NewCreateAssetImageRegistryScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSetCheckModeRequest() (request *SetCheckModeRequest) {
	request = &SetCheckModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "SetCheckMode")
	return
}

func NewSetCheckModeResponse() (response *SetCheckModeResponse) {
	response = &SetCheckModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置检测模式和自动检查
func (c *Client) SetCheckMode(request *SetCheckModeRequest) (response *SetCheckModeResponse, err error) {
	if request == nil {
		request = NewSetCheckModeRequest()
	}
	response = NewSetCheckModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeEventTendencyRequest() (request *DescribeEscapeEventTendencyRequest) {
	request = &DescribeEscapeEventTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventTendency")
	return
}

func NewDescribeEscapeEventTendencyResponse() (response *DescribeEscapeEventTendencyResponse) {
	response = &DescribeEscapeEventTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询待处理逃逸事件趋势
func (c *Client) DescribeEscapeEventTendency(request *DescribeEscapeEventTendencyRequest) (response *DescribeEscapeEventTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeEventTendencyRequest()
	}
	response = NewDescribeEscapeEventTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewSetLocalStorageExpireRequest() (request *SetLocalStorageExpireRequest) {
	request = &SetLocalStorageExpireRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "SetLocalStorageExpire")
	return
}

func NewSetLocalStorageExpireResponse() (response *SetLocalStorageExpireResponse) {
	response = &SetLocalStorageExpireResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置本地存储过期时间
func (c *Client) SetLocalStorageExpire(request *SetLocalStorageExpireRequest) (response *SetLocalStorageExpireResponse, err error) {
	if request == nil {
		request = NewSetLocalStorageExpireRequest()
	}
	response = NewSetLocalStorageExpireResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalEventListRequest() (request *DescribeK8sApiAbnormalEventListRequest) {
	request = &DescribeK8sApiAbnormalEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalEventList")
	return
}

func NewDescribeK8sApiAbnormalEventListResponse() (response *DescribeK8sApiAbnormalEventListResponse) {
	response = &DescribeK8sApiAbnormalEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8s api异常事件列表
func (c *Client) DescribeK8sApiAbnormalEventList(request *DescribeK8sApiAbnormalEventListRequest) (response *DescribeK8sApiAbnormalEventListResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalEventListRequest()
	}
	response = NewDescribeK8sApiAbnormalEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeEventTypeSummaryRequest() (request *DescribeEscapeEventTypeSummaryRequest) {
	request = &DescribeEscapeEventTypeSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventTypeSummary")
	return
}

func NewDescribeEscapeEventTypeSummaryResponse() (response *DescribeEscapeEventTypeSummaryResponse) {
	response = &DescribeEscapeEventTypeSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统计容器逃逸各事件类型和待处理事件数
func (c *Client) DescribeEscapeEventTypeSummary(request *DescribeEscapeEventTypeSummaryRequest) (response *DescribeEscapeEventTypeSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeEventTypeSummaryRequest()
	}
	response = NewDescribeEscapeEventTypeSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespaceListRequest() (request *DescribeNamespaceListRequest) {
	request = &DescribeNamespaceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNamespaceList")
	return
}

func NewDescribeNamespaceListResponse() (response *DescribeNamespaceListResponse) {
	response = &DescribeNamespaceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取命名空间列表
func (c *Client) DescribeNamespaceList(request *DescribeNamespaceListRequest) (response *DescribeNamespaceListResponse, err error) {
	if request == nil {
		request = NewDescribeNamespaceListRequest()
	}
	response = NewDescribeNamespaceListResponse()
	err = c.Send(request, response)
	return
}

func NewAddIgnoreVulRequest() (request *AddIgnoreVulRequest) {
	request = &AddIgnoreVulRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddIgnoreVul")
	return
}

func NewAddIgnoreVulResponse() (response *AddIgnoreVulResponse) {
	response = &AddIgnoreVulResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增漏洞扫描忽略漏洞
func (c *Client) AddIgnoreVul(request *AddIgnoreVulRequest) (response *AddIgnoreVulResponse, err error) {
	if request == nil {
		request = NewAddIgnoreVulRequest()
	}
	response = NewAddIgnoreVulResponse()
	err = c.Send(request, response)
	return
}

func NewAddAssetImageRegistryRegistryDetailRequest() (request *AddAssetImageRegistryRegistryDetailRequest) {
	request = &AddAssetImageRegistryRegistryDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddAssetImageRegistryRegistryDetail")
	return
}

func NewAddAssetImageRegistryRegistryDetailResponse() (response *AddAssetImageRegistryRegistryDetailResponse) {
	response = &AddAssetImageRegistryRegistryDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增单个镜像仓库详细信息
func (c *Client) AddAssetImageRegistryRegistryDetail(request *AddAssetImageRegistryRegistryDetailRequest) (response *AddAssetImageRegistryRegistryDetailResponse, err error) {
	if request == nil {
		request = NewAddAssetImageRegistryRegistryDetailRequest()
	}
	response = NewAddAssetImageRegistryRegistryDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterWorkloadFromWhiteListRequest() (request *DeleteClusterWorkloadFromWhiteListRequest) {
	request = &DeleteClusterWorkloadFromWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteClusterWorkloadFromWhiteList")
	return
}

func NewDeleteClusterWorkloadFromWhiteListResponse() (response *DeleteClusterWorkloadFromWhiteListResponse) {
	response = &DeleteClusterWorkloadFromWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从白名单中删除工作负载
func (c *Client) DeleteClusterWorkloadFromWhiteList(request *DeleteClusterWorkloadFromWhiteListRequest) (response *DeleteClusterWorkloadFromWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteClusterWorkloadFromWhiteListRequest()
	}
	response = NewDeleteClusterWorkloadFromWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewSetLocalStorageItemRequest() (request *SetLocalStorageItemRequest) {
	request = &SetLocalStorageItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "SetLocalStorageItem")
	return
}

func NewSetLocalStorageItemResponse() (response *SetLocalStorageItemResponse) {
	response = &SetLocalStorageItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置本地存储数据
func (c *Client) SetLocalStorageItem(request *SetLocalStorageItemRequest) (response *SetLocalStorageItemResponse, err error) {
	if request == nil {
		request = NewSetLocalStorageItemRequest()
	}
	response = NewSetLocalStorageItemResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetImageScanSettingRequest() (request *CreateAssetImageScanSettingRequest) {
	request = &CreateAssetImageScanSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageScanSetting")
	return
}

func NewCreateAssetImageScanSettingResponse() (response *CreateAssetImageScanSettingResponse) {
	response = &CreateAssetImageScanSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加容器安全镜像扫描设置
func (c *Client) CreateAssetImageScanSetting(request *CreateAssetImageScanSettingRequest) (response *CreateAssetImageScanSettingResponse, err error) {
	if request == nil {
		request = NewCreateAssetImageScanSettingRequest()
	}
	response = NewCreateAssetImageScanSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanIgnoreVulListRequest() (request *DescribeScanIgnoreVulListRequest) {
	request = &DescribeScanIgnoreVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeScanIgnoreVulList")
	return
}

func NewDescribeScanIgnoreVulListResponse() (response *DescribeScanIgnoreVulListResponse) {
	response = &DescribeScanIgnoreVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询扫描忽略的漏洞列表
func (c *Client) DescribeScanIgnoreVulList(request *DescribeScanIgnoreVulListRequest) (response *DescribeScanIgnoreVulListResponse, err error) {
	if request == nil {
		request = NewDescribeScanIgnoreVulListRequest()
	}
	response = NewDescribeScanIgnoreVulListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallNamespaceListRequest() (request *DescribeNetworkFirewallNamespaceListRequest) {
	request = &DescribeNetworkFirewallNamespaceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallNamespaceList")
	return
}

func NewDescribeNetworkFirewallNamespaceListResponse() (response *DescribeNetworkFirewallNamespaceListResponse) {
	response = &DescribeNetworkFirewallNamespaceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群网络空间列表
func (c *Client) DescribeNetworkFirewallNamespaceList(request *DescribeNetworkFirewallNamespaceListRequest) (response *DescribeNetworkFirewallNamespaceListResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallNamespaceListRequest()
	}
	response = NewDescribeNetworkFirewallNamespaceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulLevelImageSummaryRequest() (request *DescribeVulLevelImageSummaryRequest) {
	request = &DescribeVulLevelImageSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulLevelImageSummary")
	return
}

func NewDescribeVulLevelImageSummaryResponse() (response *DescribeVulLevelImageSummaryResponse) {
	response = &DescribeVulLevelImageSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应急漏洞各威胁等级统计镜像数
func (c *Client) DescribeVulLevelImageSummary(request *DescribeVulLevelImageSummaryRequest) (response *DescribeVulLevelImageSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeVulLevelImageSummaryRequest()
	}
	response = NewDescribeVulLevelImageSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEscapeWhiteListRequest() (request *ModifyEscapeWhiteListRequest) {
	request = &ModifyEscapeWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyEscapeWhiteList")
	return
}

func NewModifyEscapeWhiteListResponse() (response *ModifyEscapeWhiteListResponse) {
	response = &ModifyEscapeWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改逃逸白名单
func (c *Client) ModifyEscapeWhiteList(request *ModifyEscapeWhiteListRequest) (response *ModifyEscapeWhiteListResponse, err error) {
	if request == nil {
		request = NewModifyEscapeWhiteListRequest()
	}
	response = NewModifyEscapeWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePromotionActivityRequest() (request *DescribePromotionActivityRequest) {
	request = &DescribePromotionActivityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribePromotionActivity")
	return
}

func NewDescribePromotionActivityResponse() (response *DescribePromotionActivityResponse) {
	response = &DescribePromotionActivityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询促销活动
func (c *Client) DescribePromotionActivity(request *DescribePromotionActivityRequest) (response *DescribePromotionActivityResponse, err error) {
	if request == nil {
		request = NewDescribePromotionActivityRequest()
	}
	response = NewDescribePromotionActivityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIngressForwardConfigRequest() (request *DescribeIngressForwardConfigRequest) {
	request = &DescribeIngressForwardConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeIngressForwardConfig")
	return
}

func NewDescribeIngressForwardConfigResponse() (response *DescribeIngressForwardConfigResponse) {
	response = &DescribeIngressForwardConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Ingress的转发配置信息列表
func (c *Client) DescribeIngressForwardConfig(request *DescribeIngressForwardConfigRequest) (response *DescribeIngressForwardConfigResponse, err error) {
	if request == nil {
		request = NewDescribeIngressForwardConfigRequest()
	}
	response = NewDescribeIngressForwardConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulImageExportJobRequest() (request *CreateVulImageExportJobRequest) {
	request = &CreateVulImageExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateVulImageExportJob")
	return
}

func NewCreateVulImageExportJobResponse() (response *CreateVulImageExportJobResponse) {
	response = &CreateVulImageExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建受漏洞影响的镜像导出任务
func (c *Client) CreateVulImageExportJob(request *CreateVulImageExportJobRequest) (response *CreateVulImageExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVulImageExportJobRequest()
	}
	response = NewCreateVulImageExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIACCheckTaskRequest() (request *CreateIACCheckTaskRequest) {
	request = &CreateIACCheckTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateIACCheckTask")
	return
}

func NewCreateIACCheckTaskResponse() (response *CreateIACCheckTaskResponse) {
	response = &CreateIACCheckTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建IAC检测任务
func (c *Client) CreateIACCheckTask(request *CreateIACCheckTaskRequest) (response *CreateIACCheckTaskResponse, err error) {
	if request == nil {
		request = NewCreateIACCheckTaskRequest()
	}
	response = NewCreateIACCheckTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountStatusRequest() (request *DescribeAccountStatusRequest) {
	request = &DescribeAccountStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccountStatus")
	return
}

func NewDescribeAccountStatusResponse() (response *DescribeAccountStatusResponse) {
	response = &DescribeAccountStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账号状态
func (c *Client) DescribeAccountStatus(request *DescribeAccountStatusRequest) (response *DescribeAccountStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAccountStatusRequest()
	}
	response = NewDescribeAccountStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkTopologyFlowSwitchTaskRequest() (request *CreateNetworkTopologyFlowSwitchTaskRequest) {
	request = &CreateNetworkTopologyFlowSwitchTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateNetworkTopologyFlowSwitchTask")
	return
}

func NewCreateNetworkTopologyFlowSwitchTaskResponse() (response *CreateNetworkTopologyFlowSwitchTaskResponse) {
	response = &CreateNetworkTopologyFlowSwitchTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建网络拓扑集群流量开关任务
func (c *Client) CreateNetworkTopologyFlowSwitchTask(request *CreateNetworkTopologyFlowSwitchTaskRequest) (response *CreateNetworkTopologyFlowSwitchTaskResponse, err error) {
	if request == nil {
		request = NewCreateNetworkTopologyFlowSwitchTaskRequest()
	}
	response = NewCreateNetworkTopologyFlowSwitchTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetPortListRequest() (request *DescribeAssetPortListRequest) {
	request = &DescribeAssetPortListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetPortList")
	return
}

func NewDescribeAssetPortListResponse() (response *DescribeAssetPortListResponse) {
	response = &DescribeAssetPortListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全搜索查询端口占用列表
func (c *Client) DescribeAssetPortList(request *DescribeAssetPortListRequest) (response *DescribeAssetPortListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetPortListRequest()
	}
	response = NewDescribeAssetPortListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistrySummaryRequest() (request *DescribeAssetImageRegistrySummaryRequest) {
	request = &DescribeAssetImageRegistrySummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistrySummary")
	return
}

func NewDescribeAssetImageRegistrySummaryResponse() (response *DescribeAssetImageRegistrySummaryResponse) {
	response = &DescribeAssetImageRegistrySummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询镜像统计信息
func (c *Client) DescribeAssetImageRegistrySummary(request *DescribeAssetImageRegistrySummaryRequest) (response *DescribeAssetImageRegistrySummaryResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistrySummaryRequest()
	}
	response = NewDescribeAssetImageRegistrySummaryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirusAutoIsolateExampleSwitchRequest() (request *ModifyVirusAutoIsolateExampleSwitchRequest) {
	request = &ModifyVirusAutoIsolateExampleSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusAutoIsolateExampleSwitch")
	return
}

func NewModifyVirusAutoIsolateExampleSwitchResponse() (response *ModifyVirusAutoIsolateExampleSwitchResponse) {
	response = &ModifyVirusAutoIsolateExampleSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改木马自动隔离样本开关
func (c *Client) ModifyVirusAutoIsolateExampleSwitch(request *ModifyVirusAutoIsolateExampleSwitchRequest) (response *ModifyVirusAutoIsolateExampleSwitchResponse, err error) {
	if request == nil {
		request = NewModifyVirusAutoIsolateExampleSwitchRequest()
	}
	response = NewModifyVirusAutoIsolateExampleSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallWhiteListsRequest() (request *DescribeRiskSyscallWhiteListsRequest) {
	request = &DescribeRiskSyscallWhiteListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallWhiteLists")
	return
}

func NewDescribeRiskSyscallWhiteListsResponse() (response *DescribeRiskSyscallWhiteListsResponse) {
	response = &DescribeRiskSyscallWhiteListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时高危系统调用白名单列表信息
func (c *Client) DescribeRiskSyscallWhiteLists(request *DescribeRiskSyscallWhiteListsRequest) (response *DescribeRiskSyscallWhiteListsResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallWhiteListsRequest()
	}
	response = NewDescribeRiskSyscallWhiteListsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAssetImageRegistryRegistryDetailRequest() (request *UpdateAssetImageRegistryRegistryDetailRequest) {
	request = &UpdateAssetImageRegistryRegistryDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "UpdateAssetImageRegistryRegistryDetail")
	return
}

func NewUpdateAssetImageRegistryRegistryDetailResponse() (response *UpdateAssetImageRegistryRegistryDetailResponse) {
	response = &UpdateAssetImageRegistryRegistryDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新单个镜像仓库详细信息
func (c *Client) UpdateAssetImageRegistryRegistryDetail(request *UpdateAssetImageRegistryRegistryDetailRequest) (response *UpdateAssetImageRegistryRegistryDetailResponse, err error) {
	if request == nil {
		request = NewUpdateAssetImageRegistryRegistryDetailRequest()
	}
	response = NewUpdateAssetImageRegistryRegistryDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExportComplianceStatusListJobRequest() (request *CreateExportComplianceStatusListJobRequest) {
	request = &CreateExportComplianceStatusListJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateExportComplianceStatusListJob")
	return
}

func NewCreateExportComplianceStatusListJobResponse() (response *CreateExportComplianceStatusListJobResponse) {
	response = &CreateExportComplianceStatusListJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建一个导出安全合规信息的任务
func (c *Client) CreateExportComplianceStatusListJob(request *CreateExportComplianceStatusListJobRequest) (response *CreateExportComplianceStatusListJobResponse, err error) {
	if request == nil {
		request = NewCreateExportComplianceStatusListJobRequest()
	}
	response = NewCreateExportComplianceStatusListJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRiskTendencyRequest() (request *DescribeImageRiskTendencyRequest) {
	request = &DescribeImageRiskTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRiskTendency")
	return
}

func NewDescribeImageRiskTendencyResponse() (response *DescribeImageRiskTendencyResponse) {
	response = &DescribeImageRiskTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器安全本地镜像风险趋势
func (c *Client) DescribeImageRiskTendency(request *DescribeImageRiskTendencyRequest) (response *DescribeImageRiskTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeImageRiskTendencyRequest()
	}
	response = NewDescribeImageRiskTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEventEscapeImageStatusRequest() (request *ModifyEventEscapeImageStatusRequest) {
	request = &ModifyEventEscapeImageStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyEventEscapeImageStatus")
	return
}

func NewModifyEventEscapeImageStatusResponse() (response *ModifyEventEscapeImageStatusResponse) {
	response = &ModifyEventEscapeImageStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改风险容器镜像状态
func (c *Client) ModifyEventEscapeImageStatus(request *ModifyEventEscapeImageStatusRequest) (response *ModifyEventEscapeImageStatusResponse, err error) {
	if request == nil {
		request = NewModifyEventEscapeImageStatusRequest()
	}
	response = NewModifyEventEscapeImageStatusResponse()
	err = c.Send(request, response)
	return
}

func NewScanComplianceAssetsRequest() (request *ScanComplianceAssetsRequest) {
	request = &ScanComplianceAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ScanComplianceAssets")
	return
}

func NewScanComplianceAssetsResponse() (response *ScanComplianceAssetsResponse) {
	response = &ScanComplianceAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新检测选定的资产
func (c *Client) ScanComplianceAssets(request *ScanComplianceAssetsRequest) (response *ScanComplianceAssetsResponse, err error) {
	if request == nil {
		request = NewScanComplianceAssetsRequest()
	}
	response = NewScanComplianceAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulRegistryImageListRequest() (request *DescribeVulRegistryImageListRequest) {
	request = &DescribeVulRegistryImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulRegistryImageList")
	return
}

func NewDescribeVulRegistryImageListResponse() (response *DescribeVulRegistryImageListResponse) {
	response = &DescribeVulRegistryImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞影响的仓库镜像列表
func (c *Client) DescribeVulRegistryImageList(request *DescribeVulRegistryImageListRequest) (response *DescribeVulRegistryImageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulRegistryImageListRequest()
	}
	response = NewDescribeVulRegistryImageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulIgnoreLocalImageListRequest() (request *DescribeVulIgnoreLocalImageListRequest) {
	request = &DescribeVulIgnoreLocalImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulIgnoreLocalImageList")
	return
}

func NewDescribeVulIgnoreLocalImageListResponse() (response *DescribeVulIgnoreLocalImageListResponse) {
	response = &DescribeVulIgnoreLocalImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞扫描忽略的本地镜像列表
func (c *Client) DescribeVulIgnoreLocalImageList(request *DescribeVulIgnoreLocalImageListRequest) (response *DescribeVulIgnoreLocalImageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulIgnoreLocalImageListRequest()
	}
	response = NewDescribeVulIgnoreLocalImageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkTopologyWorkLoadChartRequest() (request *DescribeNetworkTopologyWorkLoadChartRequest) {
	request = &DescribeNetworkTopologyWorkLoadChartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkTopologyWorkLoadChart")
	return
}

func NewDescribeNetworkTopologyWorkLoadChartResponse() (response *DescribeNetworkTopologyWorkLoadChartResponse) {
	response = &DescribeNetworkTopologyWorkLoadChartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络查询网络拓扑负载图
func (c *Client) DescribeNetworkTopologyWorkLoadChart(request *DescribeNetworkTopologyWorkLoadChartRequest) (response *DescribeNetworkTopologyWorkLoadChartResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkTopologyWorkLoadChartRequest()
	}
	response = NewDescribeNetworkTopologyWorkLoadChartResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerSecEventSummaryRequest() (request *DescribeContainerSecEventSummaryRequest) {
	request = &DescribeContainerSecEventSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeContainerSecEventSummary")
	return
}

func NewDescribeContainerSecEventSummaryResponse() (response *DescribeContainerSecEventSummaryResponse) {
	response = &DescribeContainerSecEventSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器安全未处理事件信息
func (c *Client) DescribeContainerSecEventSummary(request *DescribeContainerSecEventSummaryRequest) (response *DescribeContainerSecEventSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeContainerSecEventSummaryRequest()
	}
	response = NewDescribeContainerSecEventSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewAddEditAccessControlRuleRequest() (request *AddEditAccessControlRuleRequest) {
	request = &AddEditAccessControlRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddEditAccessControlRule")
	return
}

func NewAddEditAccessControlRuleResponse() (response *AddEditAccessControlRuleResponse) {
	response = &AddEditAccessControlRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加编辑运行时访问控制策略
func (c *Client) AddEditAccessControlRule(request *AddEditAccessControlRuleRequest) (response *AddEditAccessControlRuleResponse, err error) {
	if request == nil {
		request = NewAddEditAccessControlRuleRequest()
	}
	response = NewAddEditAccessControlRuleResponse()
	err = c.Send(request, response)
	return
}

func NewKeysLocalStorageRequest() (request *KeysLocalStorageRequest) {
	request = &KeysLocalStorageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "KeysLocalStorage")
	return
}

func NewKeysLocalStorageResponse() (response *KeysLocalStorageResponse) {
	response = &KeysLocalStorageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取本地存储键值列表
func (c *Client) KeysLocalStorage(request *KeysLocalStorageRequest) (response *KeysLocalStorageResponse, err error) {
	if request == nil {
		request = NewKeysLocalStorageRequest()
	}
	response = NewKeysLocalStorageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemVulListRequest() (request *DescribeSystemVulListRequest) {
	request = &DescribeSystemVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSystemVulList")
	return
}

func NewDescribeSystemVulListResponse() (response *DescribeSystemVulListResponse) {
	response = &DescribeSystemVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询系统漏洞列表
func (c *Client) DescribeSystemVulList(request *DescribeSystemVulListRequest) (response *DescribeSystemVulListResponse, err error) {
	if request == nil {
		request = NewDescribeSystemVulListRequest()
	}
	response = NewDescribeSystemVulListResponse()
	err = c.Send(request, response)
	return
}

func NewImageRegistryConnDetectRequest() (request *ImageRegistryConnDetectRequest) {
	request = &ImageRegistryConnDetectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ImageRegistryConnDetect")
	return
}

func NewImageRegistryConnDetectResponse() (response *ImageRegistryConnDetectResponse) {
	response = &ImageRegistryConnDetectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库联通性检测
func (c *Client) ImageRegistryConnDetect(request *ImageRegistryConnDetectRequest) (response *ImageRegistryConnDetectResponse, err error) {
	if request == nil {
		request = NewImageRegistryConnDetectRequest()
	}
	response = NewImageRegistryConnDetectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAffectedClusterCountRequest() (request *DescribeAffectedClusterCountRequest) {
	request = &DescribeAffectedClusterCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAffectedClusterCount")
	return
}

func NewDescribeAffectedClusterCountResponse() (response *DescribeAffectedClusterCountResponse) {
	response = &DescribeAffectedClusterCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取受影响的集群数量，返回数量
func (c *Client) DescribeAffectedClusterCount(request *DescribeAffectedClusterCountRequest) (response *DescribeAffectedClusterCountResponse, err error) {
	if request == nil {
		request = NewDescribeAffectedClusterCountRequest()
	}
	response = NewDescribeAffectedClusterCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageVirusListExportRequest() (request *DescribeAssetImageVirusListExportRequest) {
	request = &DescribeAssetImageVirusListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVirusListExport")
	return
}

func NewDescribeAssetImageVirusListExportResponse() (response *DescribeAssetImageVirusListExportResponse) {
	response = &DescribeAssetImageVirusListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全搜索查询镜像木马列表导出
func (c *Client) DescribeAssetImageVirusListExport(request *DescribeAssetImageVirusListExportRequest) (response *DescribeAssetImageVirusListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageVirusListExportRequest()
	}
	response = NewDescribeAssetImageVirusListExportResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReverseShellStatusRequest() (request *ModifyReverseShellStatusRequest) {
	request = &ModifyReverseShellStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyReverseShellStatus")
	return
}

func NewModifyReverseShellStatusResponse() (response *ModifyReverseShellStatusResponse) {
	response = &ModifyReverseShellStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改反弹shell事件的状态信息
func (c *Client) ModifyReverseShellStatus(request *ModifyReverseShellStatusRequest) (response *ModifyReverseShellStatusResponse, err error) {
	if request == nil {
		request = NewModifyReverseShellStatusRequest()
	}
	response = NewModifyReverseShellStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRefreshTaskRequest() (request *DescribeRefreshTaskRequest) {
	request = &DescribeRefreshTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRefreshTask")
	return
}

func NewDescribeRefreshTaskResponse() (response *DescribeRefreshTaskResponse) {
	response = &DescribeRefreshTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询刷新任务
func (c *Client) DescribeRefreshTask(request *DescribeRefreshTaskRequest) (response *DescribeRefreshTaskResponse, err error) {
	if request == nil {
		request = NewDescribeRefreshTaskRequest()
	}
	response = NewDescribeRefreshTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateComponentExportJobRequest() (request *CreateComponentExportJobRequest) {
	request = &CreateComponentExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateComponentExportJob")
	return
}

func NewCreateComponentExportJobResponse() (response *CreateComponentExportJobResponse) {
	response = &CreateComponentExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地镜像组件列表导出
func (c *Client) CreateComponentExportJob(request *CreateComponentExportJobRequest) (response *CreateComponentExportJobResponse, err error) {
	if request == nil {
		request = NewCreateComponentExportJobRequest()
	}
	response = NewCreateComponentExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterServiceYamlRequest() (request *DescribeClusterServiceYamlRequest) {
	request = &DescribeClusterServiceYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterServiceYaml")
	return
}

func NewDescribeClusterServiceYamlResponse() (response *DescribeClusterServiceYamlResponse) {
	response = &DescribeClusterServiceYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群service的yaml文件
func (c *Client) DescribeClusterServiceYaml(request *DescribeClusterServiceYamlRequest) (response *DescribeClusterServiceYamlResponse, err error) {
	if request == nil {
		request = NewDescribeClusterServiceYamlRequest()
	}
	response = NewDescribeClusterServiceYamlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyRuleListRequest() (request *DescribeImageDenyRuleListRequest) {
	request = &DescribeImageDenyRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageDenyRuleList")
	return
}

func NewDescribeImageDenyRuleListResponse() (response *DescribeImageDenyRuleListResponse) {
	response = &DescribeImageDenyRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截规则列表
func (c *Client) DescribeImageDenyRuleList(request *DescribeImageDenyRuleListRequest) (response *DescribeImageDenyRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyRuleListRequest()
	}
	response = NewDescribeImageDenyRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterWorkloadsRequest() (request *DescribeClusterWorkloadsRequest) {
	request = &DescribeClusterWorkloadsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterWorkloads")
	return
}

func NewDescribeClusterWorkloadsResponse() (response *DescribeClusterWorkloadsResponse) {
	response = &DescribeClusterWorkloadsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群工作负载信息
func (c *Client) DescribeClusterWorkloads(request *DescribeClusterWorkloadsRequest) (response *DescribeClusterWorkloadsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterWorkloadsRequest()
	}
	response = NewDescribeClusterWorkloadsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogStorageStatisticRequest() (request *DescribeLogStorageStatisticRequest) {
	request = &DescribeLogStorageStatisticRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeLogStorageStatistic")
	return
}

func NewDescribeLogStorageStatisticResponse() (response *DescribeLogStorageStatisticResponse) {
	response = &DescribeLogStorageStatisticResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志检索容量使用统计
func (c *Client) DescribeLogStorageStatistic(request *DescribeLogStorageStatisticRequest) (response *DescribeLogStorageStatisticResponse, err error) {
	if request == nil {
		request = NewDescribeLogStorageStatisticRequest()
	}
	response = NewDescribeLogStorageStatisticResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusAutoIsolateSampleListRequest() (request *DescribeVirusAutoIsolateSampleListRequest) {
	request = &DescribeVirusAutoIsolateSampleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusAutoIsolateSampleList")
	return
}

func NewDescribeVirusAutoIsolateSampleListResponse() (response *DescribeVirusAutoIsolateSampleListResponse) {
	response = &DescribeVirusAutoIsolateSampleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马自动隔离样本列表
func (c *Client) DescribeVirusAutoIsolateSampleList(request *DescribeVirusAutoIsolateSampleListRequest) (response *DescribeVirusAutoIsolateSampleListResponse, err error) {
	if request == nil {
		request = NewDescribeVirusAutoIsolateSampleListRequest()
	}
	response = NewDescribeVirusAutoIsolateSampleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlRuleDetailRequest() (request *DescribeAccessControlRuleDetailRequest) {
	request = &DescribeAccessControlRuleDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlRuleDetail")
	return
}

func NewDescribeAccessControlRuleDetailResponse() (response *DescribeAccessControlRuleDetailResponse) {
	response = &DescribeAccessControlRuleDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时访问控制策略详细信息
func (c *Client) DescribeAccessControlRuleDetail(request *DescribeAccessControlRuleDetailRequest) (response *DescribeAccessControlRuleDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlRuleDetailRequest()
	}
	response = NewDescribeAccessControlRuleDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallClusterRefreshStatusRequest() (request *DescribeNetworkFirewallClusterRefreshStatusRequest) {
	request = &DescribeNetworkFirewallClusterRefreshStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallClusterRefreshStatus")
	return
}

func NewDescribeNetworkFirewallClusterRefreshStatusResponse() (response *DescribeNetworkFirewallClusterRefreshStatusResponse) {
	response = &DescribeNetworkFirewallClusterRefreshStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络查询资产任务进度
func (c *Client) DescribeNetworkFirewallClusterRefreshStatus(request *DescribeNetworkFirewallClusterRefreshStatusRequest) (response *DescribeNetworkFirewallClusterRefreshStatusResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallClusterRefreshStatusRequest()
	}
	response = NewDescribeNetworkFirewallClusterRefreshStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRiskListRequest() (request *DescribeAssetImageRiskListRequest) {
	request = &DescribeAssetImageRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRiskList")
	return
}

func NewDescribeAssetImageRiskListResponse() (response *DescribeAssetImageRiskListResponse) {
	response = &DescribeAssetImageRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全查询镜像风险列表
func (c *Client) DescribeAssetImageRiskList(request *DescribeAssetImageRiskListRequest) (response *DescribeAssetImageRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRiskListRequest()
	}
	response = NewDescribeAssetImageRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDetailRequest() (request *DescribeVulDetailRequest) {
	request = &DescribeVulDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDetail")
	return
}

func NewDescribeVulDetailResponse() (response *DescribeVulDetailResponse) {
	response = &DescribeVulDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞详情
func (c *Client) DescribeVulDetail(request *DescribeVulDetailRequest) (response *DescribeVulDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVulDetailRequest()
	}
	response = NewDescribeVulDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlRulesExportRequest() (request *DescribeAccessControlRulesExportRequest) {
	request = &DescribeAccessControlRulesExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlRulesExport")
	return
}

func NewDescribeAccessControlRulesExportResponse() (response *DescribeAccessControlRulesExportResponse) {
	response = &DescribeAccessControlRulesExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时访问控制策略列表导出
func (c *Client) DescribeAccessControlRulesExport(request *DescribeAccessControlRulesExportRequest) (response *DescribeAccessControlRulesExportResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlRulesExportRequest()
	}
	response = NewDescribeAccessControlRulesExportResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCompliancePolicyAssetSetFromWhitelistRequest() (request *DeleteCompliancePolicyAssetSetFromWhitelistRequest) {
	request = &DeleteCompliancePolicyAssetSetFromWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteCompliancePolicyAssetSetFromWhitelist")
	return
}

func NewDeleteCompliancePolicyAssetSetFromWhitelistResponse() (response *DeleteCompliancePolicyAssetSetFromWhitelistResponse) {
	response = &DeleteCompliancePolicyAssetSetFromWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增安全合规忽略(检测项+资产)列表，不显示指定的检查项包含的资产内容
func (c *Client) DeleteCompliancePolicyAssetSetFromWhitelist(request *DeleteCompliancePolicyAssetSetFromWhitelistRequest) (response *DeleteCompliancePolicyAssetSetFromWhitelistResponse, err error) {
	if request == nil {
		request = NewDeleteCompliancePolicyAssetSetFromWhitelistRequest()
	}
	response = NewDeleteCompliancePolicyAssetSetFromWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHostExportJobRequest() (request *CreateHostExportJobRequest) {
	request = &CreateHostExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateHostExportJob")
	return
}

func NewCreateHostExportJobResponse() (response *CreateHostExportJobResponse) {
	response = &CreateHostExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建主机列表导出任务
func (c *Client) CreateHostExportJob(request *CreateHostExportJobRequest) (response *CreateHostExportJobResponse, err error) {
	if request == nil {
		request = NewCreateHostExportJobRequest()
	}
	response = NewCreateHostExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallWhiteListDetailRequest() (request *DescribeRiskSyscallWhiteListDetailRequest) {
	request = &DescribeRiskSyscallWhiteListDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallWhiteListDetail")
	return
}

func NewDescribeRiskSyscallWhiteListDetailResponse() (response *DescribeRiskSyscallWhiteListDetailResponse) {
	response = &DescribeRiskSyscallWhiteListDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时高危系统调用白名单详细信息
func (c *Client) DescribeRiskSyscallWhiteListDetail(request *DescribeRiskSyscallWhiteListDetailRequest) (response *DescribeRiskSyscallWhiteListDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallWhiteListDetailRequest()
	}
	response = NewDescribeRiskSyscallWhiteListDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryVulListExportRequest() (request *DescribeAssetImageRegistryVulListExportRequest) {
	request = &DescribeAssetImageRegistryVulListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVulListExport")
	return
}

func NewDescribeAssetImageRegistryVulListExportResponse() (response *DescribeAssetImageRegistryVulListExportResponse) {
	response = &DescribeAssetImageRegistryVulListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库漏洞列表导出
func (c *Client) DescribeAssetImageRegistryVulListExport(request *DescribeAssetImageRegistryVulListExportRequest) (response *DescribeAssetImageRegistryVulListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryVulListExportRequest()
	}
	response = NewDescribeAssetImageRegistryVulListExportResponse()
	err = c.Send(request, response)
	return
}

func NewAddEditAbnormalProcessRuleRequest() (request *AddEditAbnormalProcessRuleRequest) {
	request = &AddEditAbnormalProcessRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddEditAbnormalProcessRule")
	return
}

func NewAddEditAbnormalProcessRuleResponse() (response *AddEditAbnormalProcessRuleResponse) {
	response = &AddEditAbnormalProcessRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加编辑运行时异常进程策略
func (c *Client) AddEditAbnormalProcessRule(request *AddEditAbnormalProcessRuleRequest) (response *AddEditAbnormalProcessRuleResponse, err error) {
	if request == nil {
		request = NewAddEditAbnormalProcessRuleRequest()
	}
	response = NewAddEditAbnormalProcessRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDBServiceListRequest() (request *DescribeAssetDBServiceListRequest) {
	request = &DescribeAssetDBServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetDBServiceList")
	return
}

func NewDescribeAssetDBServiceListResponse() (response *DescribeAssetDBServiceListResponse) {
	response = &DescribeAssetDBServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全查询db服务列表
func (c *Client) DescribeAssetDBServiceList(request *DescribeAssetDBServiceListRequest) (response *DescribeAssetDBServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDBServiceListRequest()
	}
	response = NewDescribeAssetDBServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessDetailRequest() (request *DescribeAbnormalProcessDetailRequest) {
	request = &DescribeAbnormalProcessDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessDetail")
	return
}

func NewDescribeAbnormalProcessDetailResponse() (response *DescribeAbnormalProcessDetailResponse) {
	response = &DescribeAbnormalProcessDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时异常进程事件详细信息
func (c *Client) DescribeAbnormalProcessDetail(request *DescribeAbnormalProcessDetailRequest) (response *DescribeAbnormalProcessDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessDetailRequest()
	}
	response = NewDescribeAbnormalProcessDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEscapeRuleRequest() (request *ModifyEscapeRuleRequest) {
	request = &ModifyEscapeRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyEscapeRule")
	return
}

func NewModifyEscapeRuleResponse() (response *ModifyEscapeRuleResponse) {
	response = &ModifyEscapeRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyEscapeRule  修改容器逃逸扫描规则信息
func (c *Client) ModifyEscapeRule(request *ModifyEscapeRuleRequest) (response *ModifyEscapeRuleResponse, err error) {
	if request == nil {
		request = NewModifyEscapeRuleRequest()
	}
	response = NewModifyEscapeRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulRegistryImageExportJobRequest() (request *CreateVulRegistryImageExportJobRequest) {
	request = &CreateVulRegistryImageExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateVulRegistryImageExportJob")
	return
}

func NewCreateVulRegistryImageExportJobResponse() (response *CreateVulRegistryImageExportJobResponse) {
	response = &CreateVulRegistryImageExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建受漏洞影响的仓库镜像导出任务
func (c *Client) CreateVulRegistryImageExportJob(request *CreateVulRegistryImageExportJobRequest) (response *CreateVulRegistryImageExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVulRegistryImageExportJobRequest()
	}
	response = NewCreateVulRegistryImageExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEscapeEventStatusRequest() (request *ModifyEscapeEventStatusRequest) {
	request = &ModifyEscapeEventStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyEscapeEventStatus")
	return
}

func NewModifyEscapeEventStatusResponse() (response *ModifyEscapeEventStatusResponse) {
	response = &ModifyEscapeEventStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyEscapeEventStatus  修改容器逃逸扫描事件状态
func (c *Client) ModifyEscapeEventStatus(request *ModifyEscapeEventStatusRequest) (response *ModifyEscapeEventStatusResponse, err error) {
	if request == nil {
		request = NewModifyEscapeEventStatusRequest()
	}
	response = NewModifyEscapeEventStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetContainerListRequest() (request *DescribeAssetContainerListRequest) {
	request = &DescribeAssetContainerListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetContainerList")
	return
}

func NewDescribeAssetContainerListResponse() (response *DescribeAssetContainerListResponse) {
	response = &DescribeAssetContainerListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索查询容器列表
func (c *Client) DescribeAssetContainerList(request *DescribeAssetContainerListRequest) (response *DescribeAssetContainerListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetContainerListRequest()
	}
	response = NewDescribeAssetContainerListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAbnormalProcessStatusRequest() (request *ModifyAbnormalProcessStatusRequest) {
	request = &ModifyAbnormalProcessStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyAbnormalProcessStatus")
	return
}

func NewModifyAbnormalProcessStatusResponse() (response *ModifyAbnormalProcessStatusResponse) {
	response = &ModifyAbnormalProcessStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改异常进程事件的状态信息
func (c *Client) ModifyAbnormalProcessStatus(request *ModifyAbnormalProcessStatusRequest) (response *ModifyAbnormalProcessStatusResponse, err error) {
	if request == nil {
		request = NewModifyAbnormalProcessStatusRequest()
	}
	response = NewModifyAbnormalProcessStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyImageRegistryAutoAuthorizedRuleRequest() (request *ModifyImageRegistryAutoAuthorizedRuleRequest) {
	request = &ModifyImageRegistryAutoAuthorizedRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyImageRegistryAutoAuthorizedRule")
	return
}

func NewModifyImageRegistryAutoAuthorizedRuleResponse() (response *ModifyImageRegistryAutoAuthorizedRuleResponse) {
	response = &ModifyImageRegistryAutoAuthorizedRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新镜像仓库自动授权配置
func (c *Client) ModifyImageRegistryAutoAuthorizedRule(request *ModifyImageRegistryAutoAuthorizedRuleRequest) (response *ModifyImageRegistryAutoAuthorizedRuleResponse, err error) {
	if request == nil {
		request = NewModifyImageRegistryAutoAuthorizedRuleRequest()
	}
	response = NewModifyImageRegistryAutoAuthorizedRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAddNetworkFirewallPolicyDetailRequest() (request *AddNetworkFirewallPolicyDetailRequest) {
	request = &AddNetworkFirewallPolicyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddNetworkFirewallPolicyDetail")
	return
}

func NewAddNetworkFirewallPolicyDetailResponse() (response *AddNetworkFirewallPolicyDetailResponse) {
	response = &AddNetworkFirewallPolicyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建网络策略添加任务
func (c *Client) AddNetworkFirewallPolicyDetail(request *AddNetworkFirewallPolicyDetailRequest) (response *AddNetworkFirewallPolicyDetailResponse, err error) {
	if request == nil {
		request = NewAddNetworkFirewallPolicyDetailRequest()
	}
	response = NewAddNetworkFirewallPolicyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulnerabilityFocusRuleRequest() (request *ModifyVulnerabilityFocusRuleRequest) {
	request = &ModifyVulnerabilityFocusRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyVulnerabilityFocusRule")
	return
}

func NewModifyVulnerabilityFocusRuleResponse() (response *ModifyVulnerabilityFocusRuleResponse) {
	response = &ModifyVulnerabilityFocusRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改漏洞重点关注规则
func (c *Client) ModifyVulnerabilityFocusRule(request *ModifyVulnerabilityFocusRuleRequest) (response *ModifyVulnerabilityFocusRuleResponse, err error) {
	if request == nil {
		request = NewModifyVulnerabilityFocusRuleRequest()
	}
	response = NewModifyVulnerabilityFocusRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAbnormalProcessRulesExportJobRequest() (request *CreateAbnormalProcessRulesExportJobRequest) {
	request = &CreateAbnormalProcessRulesExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateAbnormalProcessRulesExportJob")
	return
}

func NewCreateAbnormalProcessRulesExportJobResponse() (response *CreateAbnormalProcessRulesExportJobResponse) {
	response = &CreateAbnormalProcessRulesExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建异常进程规则导出任务
func (c *Client) CreateAbnormalProcessRulesExportJob(request *CreateAbnormalProcessRulesExportJobRequest) (response *CreateAbnormalProcessRulesExportJobResponse, err error) {
	if request == nil {
		request = NewCreateAbnormalProcessRulesExportJobRequest()
	}
	response = NewCreateAbnormalProcessRulesExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAffectedNodeListRequest() (request *DescribeAffectedNodeListRequest) {
	request = &DescribeAffectedNodeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAffectedNodeList")
	return
}

func NewDescribeAffectedNodeListResponse() (response *DescribeAffectedNodeListResponse) {
	response = &DescribeAffectedNodeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点类型的影响范围，返回节点列表
func (c *Client) DescribeAffectedNodeList(request *DescribeAffectedNodeListRequest) (response *DescribeAffectedNodeListResponse, err error) {
	if request == nil {
		request = NewDescribeAffectedNodeListRequest()
	}
	response = NewDescribeAffectedNodeListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWebVulExportJobRequest() (request *CreateWebVulExportJobRequest) {
	request = &CreateWebVulExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateWebVulExportJob")
	return
}

func NewCreateWebVulExportJobResponse() (response *CreateWebVulExportJobResponse) {
	response = &CreateWebVulExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建web漏洞导出任务
func (c *Client) CreateWebVulExportJob(request *CreateWebVulExportJobRequest) (response *CreateWebVulExportJobResponse, err error) {
	if request == nil {
		request = NewCreateWebVulExportJobRequest()
	}
	response = NewCreateWebVulExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskListRequest() (request *DescribeRiskListRequest) {
	request = &DescribeRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskList")
	return
}

func NewDescribeRiskListResponse() (response *DescribeRiskListResponse) {
	response = &DescribeRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询最近一次任务发现的风险项的信息列表，支持根据特殊字段进行过滤
func (c *Client) DescribeRiskList(request *DescribeRiskListRequest) (response *DescribeRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskListRequest()
	}
	response = NewDescribeRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlEventsExportRequest() (request *DescribeAccessControlEventsExportRequest) {
	request = &DescribeAccessControlEventsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlEventsExport")
	return
}

func NewDescribeAccessControlEventsExportResponse() (response *DescribeAccessControlEventsExportResponse) {
	response = &DescribeAccessControlEventsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时访问控制事件列表导出
func (c *Client) DescribeAccessControlEventsExport(request *DescribeAccessControlEventsExportRequest) (response *DescribeAccessControlEventsExportResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlEventsExportRequest()
	}
	response = NewDescribeAccessControlEventsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceHostAssetsSyncStatusRequest() (request *DescribeComplianceHostAssetsSyncStatusRequest) {
	request = &DescribeComplianceHostAssetsSyncStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceHostAssetsSyncStatus")
	return
}

func NewDescribeComplianceHostAssetsSyncStatusResponse() (response *DescribeComplianceHostAssetsSyncStatusResponse) {
	response = &DescribeComplianceHostAssetsSyncStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取合规安全主机列表的同步状态
func (c *Client) DescribeComplianceHostAssetsSyncStatus(request *DescribeComplianceHostAssetsSyncStatusRequest) (response *DescribeComplianceHostAssetsSyncStatusResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceHostAssetsSyncStatusRequest()
	}
	response = NewDescribeComplianceHostAssetsSyncStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddImageDenyRuleInfoRequest() (request *DescribeAddImageDenyRuleInfoRequest) {
	request = &DescribeAddImageDenyRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAddImageDenyRuleInfo")
	return
}

func NewDescribeAddImageDenyRuleInfoResponse() (response *DescribeAddImageDenyRuleInfoResponse) {
	response = &DescribeAddImageDenyRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户新增镜像拦截规则信息
func (c *Client) DescribeAddImageDenyRuleInfo(request *DescribeAddImageDenyRuleInfoRequest) (response *DescribeAddImageDenyRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAddImageDenyRuleInfoRequest()
	}
	response = NewDescribeAddImageDenyRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEscapeWhiteListExportJobRequest() (request *CreateEscapeWhiteListExportJobRequest) {
	request = &CreateEscapeWhiteListExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateEscapeWhiteListExportJob")
	return
}

func NewCreateEscapeWhiteListExportJobResponse() (response *CreateEscapeWhiteListExportJobResponse) {
	response = &CreateEscapeWhiteListExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建逃逸白名单导出任务
func (c *Client) CreateEscapeWhiteListExportJob(request *CreateEscapeWhiteListExportJobRequest) (response *CreateEscapeWhiteListExportJobResponse, err error) {
	if request == nil {
		request = NewCreateEscapeWhiteListExportJobRequest()
	}
	response = NewCreateEscapeWhiteListExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateComplianceTaskRequest() (request *CreateComplianceTaskRequest) {
	request = &CreateComplianceTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateComplianceTask")
	return
}

func NewCreateComplianceTaskResponse() (response *CreateComplianceTaskResponse) {
	response = &CreateComplianceTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建合规检查任务，在资产级别触发重新检测时使用。
func (c *Client) CreateComplianceTask(request *CreateComplianceTaskRequest) (response *CreateComplianceTaskResponse, err error) {
	if request == nil {
		request = NewCreateComplianceTaskRequest()
	}
	response = NewCreateComplianceTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetContainerDetailRequest() (request *DescribeAssetContainerDetailRequest) {
	request = &DescribeAssetContainerDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetContainerDetail")
	return
}

func NewDescribeAssetContainerDetailResponse() (response *DescribeAssetContainerDetailResponse) {
	response = &DescribeAssetContainerDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器详细信息
func (c *Client) DescribeAssetContainerDetail(request *DescribeAssetContainerDetailRequest) (response *DescribeAssetContainerDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetContainerDetailRequest()
	}
	response = NewDescribeAssetContainerDetailResponse()
	err = c.Send(request, response)
	return
}

func NewSyncAssetImageRegistryAssetRequest() (request *SyncAssetImageRegistryAssetRequest) {
	request = &SyncAssetImageRegistryAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "SyncAssetImageRegistryAsset")
	return
}

func NewSyncAssetImageRegistryAssetResponse() (response *SyncAssetImageRegistryAssetResponse) {
	response = &SyncAssetImageRegistryAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库资产刷新
func (c *Client) SyncAssetImageRegistryAsset(request *SyncAssetImageRegistryAssetRequest) (response *SyncAssetImageRegistryAssetResponse, err error) {
	if request == nil {
		request = NewSyncAssetImageRegistryAssetRequest()
	}
	response = NewSyncAssetImageRegistryAssetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyEventDetailRequest() (request *DescribeImageDenyEventDetailRequest) {
	request = &DescribeImageDenyEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageDenyEventDetail")
	return
}

func NewDescribeImageDenyEventDetailResponse() (response *DescribeImageDenyEventDetailResponse) {
	response = &DescribeImageDenyEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截事件详情
func (c *Client) DescribeImageDenyEventDetail(request *DescribeImageDenyEventDetailRequest) (response *DescribeImageDenyEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyEventDetailRequest()
	}
	response = NewDescribeImageDenyEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageExportJobRequest() (request *CreateImageExportJobRequest) {
	request = &CreateImageExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateImageExportJob")
	return
}

func NewCreateImageExportJobResponse() (response *CreateImageExportJobResponse) {
	response = &CreateImageExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像导出任务
func (c *Client) CreateImageExportJob(request *CreateImageExportJobRequest) (response *CreateImageExportJobResponse, err error) {
	if request == nil {
		request = NewCreateImageExportJobRequest()
	}
	response = NewCreateImageExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoAuthorizedRuleHostRequest() (request *DescribeAutoAuthorizedRuleHostRequest) {
	request = &DescribeAutoAuthorizedRuleHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAutoAuthorizedRuleHost")
	return
}

func NewDescribeAutoAuthorizedRuleHostResponse() (response *DescribeAutoAuthorizedRuleHostResponse) {
	response = &DescribeAutoAuthorizedRuleHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询自动授权规则授权范围主机信息
func (c *Client) DescribeAutoAuthorizedRuleHost(request *DescribeAutoAuthorizedRuleHostRequest) (response *DescribeAutoAuthorizedRuleHostResponse, err error) {
	if request == nil {
		request = NewDescribeAutoAuthorizedRuleHostRequest()
	}
	response = NewDescribeAutoAuthorizedRuleHostResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulContainerListRequest() (request *DescribeVulContainerListRequest) {
	request = &DescribeVulContainerListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulContainerList")
	return
}

func NewDescribeVulContainerListResponse() (response *DescribeVulContainerListResponse) {
	response = &DescribeVulContainerListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询受漏洞的容器列表
func (c *Client) DescribeVulContainerList(request *DescribeVulContainerListRequest) (response *DescribeVulContainerListResponse, err error) {
	if request == nil {
		request = NewDescribeVulContainerListRequest()
	}
	response = NewDescribeVulContainerListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceHostListRequest() (request *DescribeComplianceHostListRequest) {
	request = &DescribeComplianceHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceHostList")
	return
}

func NewDescribeComplianceHostListResponse() (response *DescribeComplianceHostListResponse) {
	response = &DescribeComplianceHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机列表
func (c *Client) DescribeComplianceHostList(request *DescribeComplianceHostListRequest) (response *DescribeComplianceHostListResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceHostListRequest()
	}
	response = NewDescribeComplianceHostListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProVersionInfoRequest() (request *DescribeProVersionInfoRequest) {
	request = &DescribeProVersionInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeProVersionInfo")
	return
}

func NewDescribeProVersionInfoResponse() (response *DescribeProVersionInfoResponse) {
	response = &DescribeProVersionInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeProVersionInfo  查询专业版需购买信息
func (c *Client) DescribeProVersionInfo(request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
	if request == nil {
		request = NewDescribeProVersionInfoRequest()
	}
	response = NewDescribeProVersionInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSuperNodeInstallTaskDetailRequest() (request *DescribeSuperNodeInstallTaskDetailRequest) {
	request = &DescribeSuperNodeInstallTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSuperNodeInstallTaskDetail")
	return
}

func NewDescribeSuperNodeInstallTaskDetailResponse() (response *DescribeSuperNodeInstallTaskDetailResponse) {
	response = &DescribeSuperNodeInstallTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询超级节点安装任务详情
func (c *Client) DescribeSuperNodeInstallTaskDetail(request *DescribeSuperNodeInstallTaskDetailRequest) (response *DescribeSuperNodeInstallTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeSuperNodeInstallTaskDetailRequest()
	}
	response = NewDescribeSuperNodeInstallTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageVulListRequest() (request *DescribeAssetImageVulListRequest) {
	request = &DescribeAssetImageVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVulList")
	return
}

func NewDescribeAssetImageVulListResponse() (response *DescribeAssetImageVulListResponse) {
	response = &DescribeAssetImageVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全查询镜像漏洞列表
func (c *Client) DescribeAssetImageVulList(request *DescribeAssetImageVulListRequest) (response *DescribeAssetImageVulListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageVulListRequest()
	}
	response = NewDescribeAssetImageVulListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEscapeRegexpWhiteListExportJobRequest() (request *CreateEscapeRegexpWhiteListExportJobRequest) {
	request = &CreateEscapeRegexpWhiteListExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateEscapeRegexpWhiteListExportJob")
	return
}

func NewCreateEscapeRegexpWhiteListExportJobResponse() (response *CreateEscapeRegexpWhiteListExportJobResponse) {
	response = &CreateEscapeRegexpWhiteListExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建容器逃逸正则白名单导出任务
func (c *Client) CreateEscapeRegexpWhiteListExportJob(request *CreateEscapeRegexpWhiteListExportJobRequest) (response *CreateEscapeRegexpWhiteListExportJobResponse, err error) {
	if request == nil {
		request = NewCreateEscapeRegexpWhiteListExportJobRequest()
	}
	response = NewCreateEscapeRegexpWhiteListExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulLevelSummaryRequest() (request *DescribeVulLevelSummaryRequest) {
	request = &DescribeVulLevelSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulLevelSummary")
	return
}

func NewDescribeVulLevelSummaryResponse() (response *DescribeVulLevelSummaryResponse) {
	response = &DescribeVulLevelSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞各威胁等级统计数
func (c *Client) DescribeVulLevelSummary(request *DescribeVulLevelSummaryRequest) (response *DescribeVulLevelSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeVulLevelSummaryRequest()
	}
	response = NewDescribeVulLevelSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRegistryAutoAuthorizedRuleRequest() (request *DescribeImageRegistryAutoAuthorizedRuleRequest) {
	request = &DescribeImageRegistryAutoAuthorizedRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRegistryAutoAuthorizedRule")
	return
}

func NewDescribeImageRegistryAutoAuthorizedRuleResponse() (response *DescribeImageRegistryAutoAuthorizedRuleResponse) {
	response = &DescribeImageRegistryAutoAuthorizedRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像仓库自动授权配置
func (c *Client) DescribeImageRegistryAutoAuthorizedRule(request *DescribeImageRegistryAutoAuthorizedRuleRequest) (response *DescribeImageRegistryAutoAuthorizedRuleResponse, err error) {
	if request == nil {
		request = NewDescribeImageRegistryAutoAuthorizedRuleRequest()
	}
	response = NewDescribeImageRegistryAutoAuthorizedRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallPolicyStatusRequest() (request *DescribeNetworkFirewallPolicyStatusRequest) {
	request = &DescribeNetworkFirewallPolicyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPolicyStatus")
	return
}

func NewDescribeNetworkFirewallPolicyStatusResponse() (response *DescribeNetworkFirewallPolicyStatusResponse) {
	response = &DescribeNetworkFirewallPolicyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络查询网络策略策略执行状态
func (c *Client) DescribeNetworkFirewallPolicyStatus(request *DescribeNetworkFirewallPolicyStatusRequest) (response *DescribeNetworkFirewallPolicyStatusResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallPolicyStatusRequest()
	}
	response = NewDescribeNetworkFirewallPolicyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewAddComplianceAssetPolicySetToWhitelistRequest() (request *AddComplianceAssetPolicySetToWhitelistRequest) {
	request = &AddComplianceAssetPolicySetToWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddComplianceAssetPolicySetToWhitelist")
	return
}

func NewAddComplianceAssetPolicySetToWhitelistResponse() (response *AddComplianceAssetPolicySetToWhitelistResponse) {
	response = &AddComplianceAssetPolicySetToWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增安全合规忽略(资产+检测项列表)列表，不显示指定的检查项包含的资产内容
// 参考的AddCompliancePolicyItemToWhitelist，除输入字段外，其它应该是一致的，如果有不同可能是定义的不对
func (c *Client) AddComplianceAssetPolicySetToWhitelist(request *AddComplianceAssetPolicySetToWhitelistRequest) (response *AddComplianceAssetPolicySetToWhitelistResponse, err error) {
	if request == nil {
		request = NewAddComplianceAssetPolicySetToWhitelistRequest()
	}
	response = NewAddComplianceAssetPolicySetToWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageDenyEventRequest() (request *DeleteImageDenyEventRequest) {
	request = &DeleteImageDenyEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteImageDenyEvent")
	return
}

func NewDeleteImageDenyEventResponse() (response *DeleteImageDenyEventResponse) {
	response = &DeleteImageDenyEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除镜像拦截事件
func (c *Client) DeleteImageDenyEvent(request *DeleteImageDenyEventRequest) (response *DeleteImageDenyEventResponse, err error) {
	if request == nil {
		request = NewDeleteImageDenyEventRequest()
	}
	response = NewDeleteImageDenyEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefencePluginRequest() (request *DescribeVulDefencePluginRequest) {
	request = &DescribeVulDefencePluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefencePlugin")
	return
}

func NewDescribeVulDefencePluginResponse() (response *DescribeVulDefencePluginResponse) {
	response = &DescribeVulDefencePluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞防御插件列表
func (c *Client) DescribeVulDefencePlugin(request *DescribeVulDefencePluginRequest) (response *DescribeVulDefencePluginResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefencePluginRequest()
	}
	response = NewDescribeVulDefencePluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallEventsRequest() (request *DescribeRiskSyscallEventsRequest) {
	request = &DescribeRiskSyscallEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallEvents")
	return
}

func NewDescribeRiskSyscallEventsResponse() (response *DescribeRiskSyscallEventsResponse) {
	response = &DescribeRiskSyscallEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时运行时高危系统调用列表信息
func (c *Client) DescribeRiskSyscallEvents(request *DescribeRiskSyscallEventsRequest) (response *DescribeRiskSyscallEventsResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallEventsRequest()
	}
	response = NewDescribeRiskSyscallEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellEventsRequest() (request *DescribeReverseShellEventsRequest) {
	request = &DescribeReverseShellEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellEvents")
	return
}

func NewDescribeReverseShellEventsResponse() (response *DescribeReverseShellEventsResponse) {
	response = &DescribeReverseShellEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时反弹shell事件列表信息
func (c *Client) DescribeReverseShellEvents(request *DescribeReverseShellEventsRequest) (response *DescribeReverseShellEventsResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellEventsRequest()
	}
	response = NewDescribeReverseShellEventsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulContainerExportJobRequest() (request *CreateVulContainerExportJobRequest) {
	request = &CreateVulContainerExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateVulContainerExportJob")
	return
}

func NewCreateVulContainerExportJobResponse() (response *CreateVulContainerExportJobResponse) {
	response = &CreateVulContainerExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建受漏洞影响的容器导出任务
func (c *Client) CreateVulContainerExportJob(request *CreateVulContainerExportJobRequest) (response *CreateVulContainerExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVulContainerExportJobRequest()
	}
	response = NewCreateVulContainerExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewModifyImageDenyRuleStatusRequest() (request *ModifyImageDenyRuleStatusRequest) {
	request = &ModifyImageDenyRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyImageDenyRuleStatus")
	return
}

func NewModifyImageDenyRuleStatusResponse() (response *ModifyImageDenyRuleStatusResponse) {
	response = &ModifyImageDenyRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑镜像拦截规则开关状态
func (c *Client) ModifyImageDenyRuleStatus(request *ModifyImageDenyRuleStatusRequest) (response *ModifyImageDenyRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyImageDenyRuleStatusRequest()
	}
	response = NewModifyImageDenyRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterNodeFromWhiteListRequest() (request *DeleteClusterNodeFromWhiteListRequest) {
	request = &DeleteClusterNodeFromWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteClusterNodeFromWhiteList")
	return
}

func NewDeleteClusterNodeFromWhiteListResponse() (response *DeleteClusterNodeFromWhiteListResponse) {
	response = &DeleteClusterNodeFromWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从白名单中删除节点资产信息
func (c *Client) DeleteClusterNodeFromWhiteList(request *DeleteClusterNodeFromWhiteListRequest) (response *DeleteClusterNodeFromWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteClusterNodeFromWhiteListRequest()
	}
	response = NewDeleteClusterNodeFromWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessEventsExportRequest() (request *DescribeAbnormalProcessEventsExportRequest) {
	request = &DescribeAbnormalProcessEventsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessEventsExport")
	return
}

func NewDescribeAbnormalProcessEventsExportResponse() (response *DescribeAbnormalProcessEventsExportResponse) {
	response = &DescribeAbnormalProcessEventsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时异常进程事件列表信息导出
func (c *Client) DescribeAbnormalProcessEventsExport(request *DescribeAbnormalProcessEventsExportRequest) (response *DescribeAbnormalProcessEventsExportResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessEventsExportRequest()
	}
	response = NewDescribeAbnormalProcessEventsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusAutoIsolateSettingRequest() (request *DescribeVirusAutoIsolateSettingRequest) {
	request = &DescribeVirusAutoIsolateSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusAutoIsolateSetting")
	return
}

func NewDescribeVirusAutoIsolateSettingResponse() (response *DescribeVirusAutoIsolateSettingResponse) {
	response = &DescribeVirusAutoIsolateSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马自动隔离设置
func (c *Client) DescribeVirusAutoIsolateSetting(request *DescribeVirusAutoIsolateSettingRequest) (response *DescribeVirusAutoIsolateSettingResponse, err error) {
	if request == nil {
		request = NewDescribeVirusAutoIsolateSettingRequest()
	}
	response = NewDescribeVirusAutoIsolateSettingResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateNetworkFirewallPolicyYamlDetailRequest() (request *UpdateNetworkFirewallPolicyYamlDetailRequest) {
	request = &UpdateNetworkFirewallPolicyYamlDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "UpdateNetworkFirewallPolicyYamlDetail")
	return
}

func NewUpdateNetworkFirewallPolicyYamlDetailResponse() (response *UpdateNetworkFirewallPolicyYamlDetailResponse) {
	response = &UpdateNetworkFirewallPolicyYamlDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络更新Yaml网络策略任务
func (c *Client) UpdateNetworkFirewallPolicyYamlDetail(request *UpdateNetworkFirewallPolicyYamlDetailRequest) (response *UpdateNetworkFirewallPolicyYamlDetailResponse, err error) {
	if request == nil {
		request = NewUpdateNetworkFirewallPolicyYamlDetailRequest()
	}
	response = NewUpdateNetworkFirewallPolicyYamlDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkTopologyFlowDetailRequest() (request *DescribeNetworkTopologyFlowDetailRequest) {
	request = &DescribeNetworkTopologyFlowDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkTopologyFlowDetail")
	return
}

func NewDescribeNetworkTopologyFlowDetailResponse() (response *DescribeNetworkTopologyFlowDetailResponse) {
	response = &DescribeNetworkTopologyFlowDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络查看网络拓扑流量详情
func (c *Client) DescribeNetworkTopologyFlowDetail(request *DescribeNetworkTopologyFlowDetailRequest) (response *DescribeNetworkTopologyFlowDetailResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkTopologyFlowDetailRequest()
	}
	response = NewDescribeNetworkTopologyFlowDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIACCheckTaskStatusRequest() (request *DescribeIACCheckTaskStatusRequest) {
	request = &DescribeIACCheckTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeIACCheckTaskStatus")
	return
}

func NewDescribeIACCheckTaskStatusResponse() (response *DescribeIACCheckTaskStatusResponse) {
	response = &DescribeIACCheckTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看IAC检测任务状态
func (c *Client) DescribeIACCheckTaskStatus(request *DescribeIACCheckTaskStatusRequest) (response *DescribeIACCheckTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeIACCheckTaskStatusRequest()
	}
	response = NewDescribeIACCheckTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyRuleSummaryRequest() (request *DescribeImageDenyRuleSummaryRequest) {
	request = &DescribeImageDenyRuleSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageDenyRuleSummary")
	return
}

func NewDescribeImageDenyRuleSummaryResponse() (response *DescribeImageDenyRuleSummaryResponse) {
	response = &DescribeImageDenyRuleSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截规则统计
func (c *Client) DescribeImageDenyRuleSummary(request *DescribeImageDenyRuleSummaryRequest) (response *DescribeImageDenyRuleSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyRuleSummaryRequest()
	}
	response = NewDescribeImageDenyRuleSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsEventDetailRequest() (request *DescribeRiskDnsEventDetailRequest) {
	request = &DescribeRiskDnsEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskDnsEventDetail")
	return
}

func NewDescribeRiskDnsEventDetailResponse() (response *DescribeRiskDnsEventDetailResponse) {
	response = &DescribeRiskDnsEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意请求事件详情
func (c *Client) DescribeRiskDnsEventDetail(request *DescribeRiskDnsEventDetailRequest) (response *DescribeRiskDnsEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsEventDetailRequest()
	}
	response = NewDescribeRiskDnsEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentDaemonSetCmdRequest() (request *DescribeAgentDaemonSetCmdRequest) {
	request = &DescribeAgentDaemonSetCmdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAgentDaemonSetCmd")
	return
}

func NewDescribeAgentDaemonSetCmdResponse() (response *DescribeAgentDaemonSetCmdResponse) {
	response = &DescribeAgentDaemonSetCmdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询平行容器安装命令
func (c *Client) DescribeAgentDaemonSetCmd(request *DescribeAgentDaemonSetCmdRequest) (response *DescribeAgentDaemonSetCmdResponse, err error) {
	if request == nil {
		request = NewDescribeAgentDaemonSetCmdRequest()
	}
	response = NewDescribeAgentDaemonSetCmdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSyncLastTimeRequest() (request *DescribeAssetSyncLastTimeRequest) {
	request = &DescribeAssetSyncLastTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetSyncLastTime")
	return
}

func NewDescribeAssetSyncLastTimeResponse() (response *DescribeAssetSyncLastTimeResponse) {
	response = &DescribeAssetSyncLastTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产同步最近时间
func (c *Client) DescribeAssetSyncLastTime(request *DescribeAssetSyncLastTimeRequest) (response *DescribeAssetSyncLastTimeResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSyncLastTimeRequest()
	}
	response = NewDescribeAssetSyncLastTimeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkFirewallUndoPublishRequest() (request *CreateNetworkFirewallUndoPublishRequest) {
	request = &CreateNetworkFirewallUndoPublishRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateNetworkFirewallUndoPublish")
	return
}

func NewCreateNetworkFirewallUndoPublishResponse() (response *CreateNetworkFirewallUndoPublishResponse) {
	response = &CreateNetworkFirewallUndoPublishResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建网络策略撤销任务
func (c *Client) CreateNetworkFirewallUndoPublish(request *CreateNetworkFirewallUndoPublishRequest) (response *CreateNetworkFirewallUndoPublishResponse, err error) {
	if request == nil {
		request = NewCreateNetworkFirewallUndoPublishRequest()
	}
	response = NewCreateNetworkFirewallUndoPublishResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpgradeVersionInfoRequest() (request *DescribeUpgradeVersionInfoRequest) {
	request = &DescribeUpgradeVersionInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeUpgradeVersionInfo")
	return
}

func NewDescribeUpgradeVersionInfoResponse() (response *DescribeUpgradeVersionInfoResponse) {
	response = &DescribeUpgradeVersionInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器终端版本的版本升级信息
func (c *Client) DescribeUpgradeVersionInfo(request *DescribeUpgradeVersionInfoRequest) (response *DescribeUpgradeVersionInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUpgradeVersionInfoRequest()
	}
	response = NewDescribeUpgradeVersionInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReverseShellRegexpRuleStatusRequest() (request *ModifyReverseShellRegexpRuleStatusRequest) {
	request = &ModifyReverseShellRegexpRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyReverseShellRegexpRuleStatus")
	return
}

func NewModifyReverseShellRegexpRuleStatusResponse() (response *ModifyReverseShellRegexpRuleStatusResponse) {
	response = &ModifyReverseShellRegexpRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改反弹shell正则规则状态
func (c *Client) ModifyReverseShellRegexpRuleStatus(request *ModifyReverseShellRegexpRuleStatusRequest) (response *ModifyReverseShellRegexpRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyReverseShellRegexpRuleStatusRequest()
	}
	response = NewModifyReverseShellRegexpRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecLogKafkaUINRequest() (request *ModifySecLogKafkaUINRequest) {
	request = &ModifySecLogKafkaUINRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogKafkaUIN")
	return
}

func NewModifySecLogKafkaUINResponse() (response *ModifySecLogKafkaUINResponse) {
	response = &ModifySecLogKafkaUINResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改安全日志kafkaUIN
func (c *Client) ModifySecLogKafkaUIN(request *ModifySecLogKafkaUINRequest) (response *ModifySecLogKafkaUINResponse, err error) {
	if request == nil {
		request = NewModifySecLogKafkaUINRequest()
	}
	response = NewModifySecLogKafkaUINResponse()
	err = c.Send(request, response)
	return
}

func NewOpenTcssTrialRequest() (request *OpenTcssTrialRequest) {
	request = &OpenTcssTrialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "OpenTcssTrial")
	return
}

func NewOpenTcssTrialResponse() (response *OpenTcssTrialResponse) {
	response = &OpenTcssTrialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开通容器安全服务试用
func (c *Client) OpenTcssTrial(request *OpenTcssTrialRequest) (response *OpenTcssTrialResponse, err error) {
	if request == nil {
		request = NewOpenTcssTrialRequest()
	}
	response = NewOpenTcssTrialResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceTaskPolicyItemSummaryListRequest() (request *DescribeComplianceTaskPolicyItemSummaryListRequest) {
	request = &DescribeComplianceTaskPolicyItemSummaryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceTaskPolicyItemSummaryList")
	return
}

func NewDescribeComplianceTaskPolicyItemSummaryListResponse() (response *DescribeComplianceTaskPolicyItemSummaryListResponse) {
	response = &DescribeComplianceTaskPolicyItemSummaryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询最近一次任务发现的检测项的汇总信息列表，按照 检测项 → 资产 的两级层次展开。
func (c *Client) DescribeComplianceTaskPolicyItemSummaryList(request *DescribeComplianceTaskPolicyItemSummaryListRequest) (response *DescribeComplianceTaskPolicyItemSummaryListResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceTaskPolicyItemSummaryListRequest()
	}
	response = NewDescribeComplianceTaskPolicyItemSummaryListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIgnoreVulRequest() (request *DeleteIgnoreVulRequest) {
	request = &DeleteIgnoreVulRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteIgnoreVul")
	return
}

func NewDeleteIgnoreVulResponse() (response *DeleteIgnoreVulResponse) {
	response = &DeleteIgnoreVulResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消漏洞扫描忽略漏洞
func (c *Client) DeleteIgnoreVul(request *DeleteIgnoreVulRequest) (response *DeleteIgnoreVulResponse, err error) {
	if request == nil {
		request = NewDeleteIgnoreVulRequest()
	}
	response = NewDeleteIgnoreVulResponse()
	err = c.Send(request, response)
	return
}

func NewModifyImageAuthorizedRequest() (request *ModifyImageAuthorizedRequest) {
	request = &ModifyImageAuthorizedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyImageAuthorized")
	return
}

func NewModifyImageAuthorizedResponse() (response *ModifyImageAuthorizedResponse) {
	response = &ModifyImageAuthorizedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量授权镜像扫描V2.0
func (c *Client) ModifyImageAuthorized(request *ModifyImageAuthorizedRequest) (response *ModifyImageAuthorizedResponse, err error) {
	if request == nil {
		request = NewModifyImageAuthorizedRequest()
	}
	response = NewModifyImageAuthorizedResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublicKeyRequest() (request *DescribePublicKeyRequest) {
	request = &DescribePublicKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribePublicKey")
	return
}

func NewDescribePublicKeyResponse() (response *DescribePublicKeyResponse) {
	response = &DescribePublicKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取公钥
func (c *Client) DescribePublicKey(request *DescribePublicKeyRequest) (response *DescribePublicKeyResponse, err error) {
	if request == nil {
		request = NewDescribePublicKeyRequest()
	}
	response = NewDescribePublicKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecEventsTendencyRequest() (request *DescribeSecEventsTendencyRequest) {
	request = &DescribeSecEventsTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecEventsTendency")
	return
}

func NewDescribeSecEventsTendencyResponse() (response *DescribeSecEventsTendencyResponse) {
	response = &DescribeSecEventsTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器运行时安全事件趋势
func (c *Client) DescribeSecEventsTendency(request *DescribeSecEventsTendencyRequest) (response *DescribeSecEventsTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeSecEventsTendencyRequest()
	}
	response = NewDescribeSecEventsTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessControlRuleStatusRequest() (request *ModifyAccessControlRuleStatusRequest) {
	request = &ModifyAccessControlRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyAccessControlRuleStatus")
	return
}

func NewModifyAccessControlRuleStatusResponse() (response *ModifyAccessControlRuleStatusResponse) {
	response = &ModifyAccessControlRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改运行时访问控制策略的状态，启用或者禁用
func (c *Client) ModifyAccessControlRuleStatus(request *ModifyAccessControlRuleStatusRequest) (response *ModifyAccessControlRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyAccessControlRuleStatusRequest()
	}
	response = NewModifyAccessControlRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeValueAddedSrvInfoRequest() (request *DescribeValueAddedSrvInfoRequest) {
	request = &DescribeValueAddedSrvInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeValueAddedSrvInfo")
	return
}

func NewDescribeValueAddedSrvInfoResponse() (response *DescribeValueAddedSrvInfoResponse) {
	response = &DescribeValueAddedSrvInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeValueAddedSrvInfo查询增值服务需购买信息
func (c *Client) DescribeValueAddedSrvInfo(request *DescribeValueAddedSrvInfoRequest) (response *DescribeValueAddedSrvInfoResponse, err error) {
	if request == nil {
		request = NewDescribeValueAddedSrvInfoRequest()
	}
	response = NewDescribeValueAddedSrvInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeEventDetailRequest() (request *DescribeEscapeEventDetailRequest) {
	request = &DescribeEscapeEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventDetail")
	return
}

func NewDescribeEscapeEventDetailResponse() (response *DescribeEscapeEventDetailResponse) {
	response = &DescribeEscapeEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeEscapeEventDetail  查询容器逃逸事件详情
func (c *Client) DescribeEscapeEventDetail(request *DescribeEscapeEventDetailRequest) (response *DescribeEscapeEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeEventDetailRequest()
	}
	response = NewDescribeEscapeEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusAutoIsolateSampleDetailRequest() (request *DescribeVirusAutoIsolateSampleDetailRequest) {
	request = &DescribeVirusAutoIsolateSampleDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusAutoIsolateSampleDetail")
	return
}

func NewDescribeVirusAutoIsolateSampleDetailResponse() (response *DescribeVirusAutoIsolateSampleDetailResponse) {
	response = &DescribeVirusAutoIsolateSampleDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马自动隔离样本详情
func (c *Client) DescribeVirusAutoIsolateSampleDetail(request *DescribeVirusAutoIsolateSampleDetailRequest) (response *DescribeVirusAutoIsolateSampleDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVirusAutoIsolateSampleDetailRequest()
	}
	response = NewDescribeVirusAutoIsolateSampleDetailResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAndPublishNetworkFirewallPolicyDetailRequest() (request *UpdateAndPublishNetworkFirewallPolicyDetailRequest) {
	request = &UpdateAndPublishNetworkFirewallPolicyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "UpdateAndPublishNetworkFirewallPolicyDetail")
	return
}

func NewUpdateAndPublishNetworkFirewallPolicyDetailResponse() (response *UpdateAndPublishNetworkFirewallPolicyDetailResponse) {
	response = &UpdateAndPublishNetworkFirewallPolicyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建网络策略更新并发布任务
func (c *Client) UpdateAndPublishNetworkFirewallPolicyDetail(request *UpdateAndPublishNetworkFirewallPolicyDetailRequest) (response *UpdateAndPublishNetworkFirewallPolicyDetailResponse, err error) {
	if request == nil {
		request = NewUpdateAndPublishNetworkFirewallPolicyDetailRequest()
	}
	response = NewUpdateAndPublishNetworkFirewallPolicyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusMonitorConfigRequest() (request *DescribeVirusMonitorConfigRequest) {
	request = &DescribeVirusMonitorConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusMonitorConfig")
	return
}

func NewDescribeVirusMonitorConfigResponse() (response *DescribeVirusMonitorConfigResponse) {
	response = &DescribeVirusMonitorConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询文件查杀实时监控设置信息
func (c *Client) DescribeVirusMonitorConfig(request *DescribeVirusMonitorConfigRequest) (response *DescribeVirusMonitorConfigResponse, err error) {
	if request == nil {
		request = NewDescribeVirusMonitorConfigRequest()
	}
	response = NewDescribeVirusMonitorConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportJobManageListRequest() (request *DescribeExportJobManageListRequest) {
	request = &DescribeExportJobManageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeExportJobManageList")
	return
}

func NewDescribeExportJobManageListResponse() (response *DescribeExportJobManageListResponse) {
	response = &DescribeExportJobManageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询导出任务管理列表
func (c *Client) DescribeExportJobManageList(request *DescribeExportJobManageListRequest) (response *DescribeExportJobManageListResponse, err error) {
	if request == nil {
		request = NewDescribeExportJobManageListRequest()
	}
	response = NewDescribeExportJobManageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceScanFailedAssetListRequest() (request *DescribeComplianceScanFailedAssetListRequest) {
	request = &DescribeComplianceScanFailedAssetListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceScanFailedAssetList")
	return
}

func NewDescribeComplianceScanFailedAssetListResponse() (response *DescribeComplianceScanFailedAssetListResponse) {
	response = &DescribeComplianceScanFailedAssetListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照 资产 → 检测项 二层结构展示的信息。这里查询第一层 资产的通过率汇总信息。
func (c *Client) DescribeComplianceScanFailedAssetList(request *DescribeComplianceScanFailedAssetListRequest) (response *DescribeComplianceScanFailedAssetListResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceScanFailedAssetListRequest()
	}
	response = NewDescribeComplianceScanFailedAssetListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetImageRiskExportJobRequest() (request *CreateAssetImageRiskExportJobRequest) {
	request = &CreateAssetImageRiskExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageRiskExportJob")
	return
}

func NewCreateAssetImageRiskExportJobResponse() (response *CreateAssetImageRiskExportJobResponse) {
	response = &CreateAssetImageRiskExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建本地镜像敏感信息列表导出任务
func (c *Client) CreateAssetImageRiskExportJob(request *CreateAssetImageRiskExportJobRequest) (response *CreateAssetImageRiskExportJobResponse, err error) {
	if request == nil {
		request = NewCreateAssetImageRiskExportJobRequest()
	}
	response = NewCreateAssetImageRiskExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRiskSyscallStatusRequest() (request *ModifyRiskSyscallStatusRequest) {
	request = &ModifyRiskSyscallStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyRiskSyscallStatus")
	return
}

func NewModifyRiskSyscallStatusResponse() (response *ModifyRiskSyscallStatusResponse) {
	response = &ModifyRiskSyscallStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改高危系统调用事件的状态信息
func (c *Client) ModifyRiskSyscallStatus(request *ModifyRiskSyscallStatusRequest) (response *ModifyRiskSyscallStatusResponse, err error) {
	if request == nil {
		request = NewModifyRiskSyscallStatusRequest()
	}
	response = NewModifyRiskSyscallStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecLogJoinStateRequest() (request *ModifySecLogJoinStateRequest) {
	request = &ModifySecLogJoinStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogJoinState")
	return
}

func NewModifySecLogJoinStateResponse() (response *ModifySecLogJoinStateResponse) {
	response = &ModifySecLogJoinStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改安全日志接入状态
func (c *Client) ModifySecLogJoinState(request *ModifySecLogJoinStateRequest) (response *ModifySecLogJoinStateResponse, err error) {
	if request == nil {
		request = NewModifySecLogJoinStateRequest()
	}
	response = NewModifySecLogJoinStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaliciousConnectionRuleSummaryRequest() (request *DescribeMaliciousConnectionRuleSummaryRequest) {
	request = &DescribeMaliciousConnectionRuleSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeMaliciousConnectionRuleSummary")
	return
}

func NewDescribeMaliciousConnectionRuleSummaryResponse() (response *DescribeMaliciousConnectionRuleSummaryResponse) {
	response = &DescribeMaliciousConnectionRuleSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恶意外连黑白名单统计，主要展示黑白名单的数据统计
func (c *Client) DescribeMaliciousConnectionRuleSummary(request *DescribeMaliciousConnectionRuleSummaryRequest) (response *DescribeMaliciousConnectionRuleSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeMaliciousConnectionRuleSummaryRequest()
	}
	response = NewDescribeMaliciousConnectionRuleSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMachineRequest() (request *DeleteMachineRequest) {
	request = &DeleteMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteMachine")
	return
}

func NewDeleteMachineResponse() (response *DeleteMachineResponse) {
	response = &DeleteMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 卸载Agent客户端
func (c *Client) DeleteMachine(request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
	if request == nil {
		request = NewDeleteMachineRequest()
	}
	response = NewDeleteMachineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRiskDnsEventExportJobRequest() (request *CreateRiskDnsEventExportJobRequest) {
	request = &CreateRiskDnsEventExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateRiskDnsEventExportJob")
	return
}

func NewCreateRiskDnsEventExportJobResponse() (response *CreateRiskDnsEventExportJobResponse) {
	response = &CreateRiskDnsEventExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建恶意请求事件导出任务
func (c *Client) CreateRiskDnsEventExportJob(request *CreateRiskDnsEventExportJobRequest) (response *CreateRiskDnsEventExportJobResponse, err error) {
	if request == nil {
		request = NewCreateRiskDnsEventExportJobRequest()
	}
	response = NewCreateRiskDnsEventExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSuperNodePodExportJobRequest() (request *CreateSuperNodePodExportJobRequest) {
	request = &CreateSuperNodePodExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateSuperNodePodExportJob")
	return
}

func NewCreateSuperNodePodExportJobResponse() (response *CreateSuperNodePodExportJobResponse) {
	response = &CreateSuperNodePodExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建超级节点pod列表导出
func (c *Client) CreateSuperNodePodExportJob(request *CreateSuperNodePodExportJobRequest) (response *CreateSuperNodePodExportJobResponse, err error) {
	if request == nil {
		request = NewCreateSuperNodePodExportJobRequest()
	}
	response = NewCreateSuperNodePodExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCompliancePolicyItemAffectedAssetListRequest() (request *DescribeCompliancePolicyItemAffectedAssetListRequest) {
	request = &DescribeCompliancePolicyItemAffectedAssetListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeCompliancePolicyItemAffectedAssetList")
	return
}

func NewDescribeCompliancePolicyItemAffectedAssetListResponse() (response *DescribeCompliancePolicyItemAffectedAssetListResponse) {
	response = &DescribeCompliancePolicyItemAffectedAssetListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照 检测项 → 资产 的两级层次展开的第二层级：资产层级。
func (c *Client) DescribeCompliancePolicyItemAffectedAssetList(request *DescribeCompliancePolicyItemAffectedAssetListRequest) (response *DescribeCompliancePolicyItemAffectedAssetListResponse, err error) {
	if request == nil {
		request = NewDescribeCompliancePolicyItemAffectedAssetListRequest()
	}
	response = NewDescribeCompliancePolicyItemAffectedAssetListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulIgnoreRegistryImageListRequest() (request *DescribeVulIgnoreRegistryImageListRequest) {
	request = &DescribeVulIgnoreRegistryImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulIgnoreRegistryImageList")
	return
}

func NewDescribeVulIgnoreRegistryImageListResponse() (response *DescribeVulIgnoreRegistryImageListResponse) {
	response = &DescribeVulIgnoreRegistryImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞扫描忽略的仓库镜像列表
func (c *Client) DescribeVulIgnoreRegistryImageList(request *DescribeVulIgnoreRegistryImageListRequest) (response *DescribeVulIgnoreRegistryImageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulIgnoreRegistryImageListRequest()
	}
	response = NewDescribeVulIgnoreRegistryImageListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSuperNodeInstallTaskExportJobRequest() (request *CreateSuperNodeInstallTaskExportJobRequest) {
	request = &CreateSuperNodeInstallTaskExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateSuperNodeInstallTaskExportJob")
	return
}

func NewCreateSuperNodeInstallTaskExportJobResponse() (response *CreateSuperNodeInstallTaskExportJobResponse) {
	response = &CreateSuperNodeInstallTaskExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建超级节点安装任务列表导出任务
func (c *Client) CreateSuperNodeInstallTaskExportJob(request *CreateSuperNodeInstallTaskExportJobRequest) (response *CreateSuperNodeInstallTaskExportJobResponse, err error) {
	if request == nil {
		request = NewCreateSuperNodeInstallTaskExportJobRequest()
	}
	response = NewCreateSuperNodeInstallTaskExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSupportDefenceVulRequest() (request *DescribeSupportDefenceVulRequest) {
	request = &DescribeSupportDefenceVulRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSupportDefenceVul")
	return
}

func NewDescribeSupportDefenceVulResponse() (response *DescribeSupportDefenceVulResponse) {
	response = &DescribeSupportDefenceVulResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询支持防御的漏洞列表
func (c *Client) DescribeSupportDefenceVul(request *DescribeSupportDefenceVulRequest) (response *DescribeSupportDefenceVulResponse, err error) {
	if request == nil {
		request = NewDescribeSupportDefenceVulRequest()
	}
	response = NewDescribeSupportDefenceVulResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallPolicyDiscoverRequest() (request *DescribeNetworkFirewallPolicyDiscoverRequest) {
	request = &DescribeNetworkFirewallPolicyDiscoverRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPolicyDiscover")
	return
}

func NewDescribeNetworkFirewallPolicyDiscoverResponse() (response *DescribeNetworkFirewallPolicyDiscoverResponse) {
	response = &DescribeNetworkFirewallPolicyDiscoverResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络查询网络策略自动发现任务进度
func (c *Client) DescribeNetworkFirewallPolicyDiscover(request *DescribeNetworkFirewallPolicyDiscoverRequest) (response *DescribeNetworkFirewallPolicyDiscoverResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallPolicyDiscoverRequest()
	}
	response = NewDescribeNetworkFirewallPolicyDiscoverResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaliciousConnectionBlackListRequest() (request *DescribeMaliciousConnectionBlackListRequest) {
	request = &DescribeMaliciousConnectionBlackListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeMaliciousConnectionBlackList")
	return
}

func NewDescribeMaliciousConnectionBlackListResponse() (response *DescribeMaliciousConnectionBlackListResponse) {
	response = &DescribeMaliciousConnectionBlackListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意外连黑名单
func (c *Client) DescribeMaliciousConnectionBlackList(request *DescribeMaliciousConnectionBlackListRequest) (response *DescribeMaliciousConnectionBlackListResponse, err error) {
	if request == nil {
		request = NewDescribeMaliciousConnectionBlackListRequest()
	}
	response = NewDescribeMaliciousConnectionBlackListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccessControlRulesRequest() (request *DeleteAccessControlRulesRequest) {
	request = &DeleteAccessControlRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteAccessControlRules")
	return
}

func NewDeleteAccessControlRulesResponse() (response *DeleteAccessControlRulesResponse) {
	response = &DeleteAccessControlRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运行访问控制策略
func (c *Client) DeleteAccessControlRules(request *DeleteAccessControlRulesRequest) (response *DeleteAccessControlRulesResponse, err error) {
	if request == nil {
		request = NewDeleteAccessControlRulesRequest()
	}
	response = NewDeleteAccessControlRulesResponse()
	err = c.Send(request, response)
	return
}

func NewSetClusterCheckIdWhiteListStatusRequest() (request *SetClusterCheckIdWhiteListStatusRequest) {
	request = &SetClusterCheckIdWhiteListStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "SetClusterCheckIdWhiteListStatus")
	return
}

func NewSetClusterCheckIdWhiteListStatusResponse() (response *SetClusterCheckIdWhiteListStatusResponse) {
	response = &SetClusterCheckIdWhiteListStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置检查项ID白名单状态
func (c *Client) SetClusterCheckIdWhiteListStatus(request *SetClusterCheckIdWhiteListStatusRequest) (response *SetClusterCheckIdWhiteListStatusResponse, err error) {
	if request == nil {
		request = NewSetClusterCheckIdWhiteListStatusRequest()
	}
	response = NewSetClusterCheckIdWhiteListStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessControlStatusRequest() (request *ModifyAccessControlStatusRequest) {
	request = &ModifyAccessControlStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyAccessControlStatus")
	return
}

func NewModifyAccessControlStatusResponse() (response *ModifyAccessControlStatusResponse) {
	response = &ModifyAccessControlStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改运行时访问控制事件状态信息
func (c *Client) ModifyAccessControlStatus(request *ModifyAccessControlStatusRequest) (response *ModifyAccessControlStatusResponse, err error) {
	if request == nil {
		request = NewModifyAccessControlStatusRequest()
	}
	response = NewModifyAccessControlStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulScanRegistryImageListRequest() (request *DescribeVulScanRegistryImageListRequest) {
	request = &DescribeVulScanRegistryImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulScanRegistryImageList")
	return
}

func NewDescribeVulScanRegistryImageListResponse() (response *DescribeVulScanRegistryImageListResponse) {
	response = &DescribeVulScanRegistryImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞扫描任务的仓库镜像列表
func (c *Client) DescribeVulScanRegistryImageList(request *DescribeVulScanRegistryImageListRequest) (response *DescribeVulScanRegistryImageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulScanRegistryImageListRequest()
	}
	response = NewDescribeVulScanRegistryImageListResponse()
	err = c.Send(request, response)
	return
}

func NewAddNetworkFirewallPolicyYamlDetailRequest() (request *AddNetworkFirewallPolicyYamlDetailRequest) {
	request = &AddNetworkFirewallPolicyYamlDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddNetworkFirewallPolicyYamlDetail")
	return
}

func NewAddNetworkFirewallPolicyYamlDetailResponse() (response *AddNetworkFirewallPolicyYamlDetailResponse) {
	response = &AddNetworkFirewallPolicyYamlDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建Yaml网络策略添加任务
func (c *Client) AddNetworkFirewallPolicyYamlDetail(request *AddNetworkFirewallPolicyYamlDetailRequest) (response *AddNetworkFirewallPolicyYamlDetailResponse, err error) {
	if request == nil {
		request = NewAddNetworkFirewallPolicyYamlDetailRequest()
	}
	response = NewAddNetworkFirewallPolicyYamlDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoUpgradeSwitchRequest() (request *ModifyAutoUpgradeSwitchRequest) {
	request = &ModifyAutoUpgradeSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyAutoUpgradeSwitch")
	return
}

func NewModifyAutoUpgradeSwitchResponse() (response *ModifyAutoUpgradeSwitchResponse) {
	response = &ModifyAutoUpgradeSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改终端自动升级开关
func (c *Client) ModifyAutoUpgradeSwitch(request *ModifyAutoUpgradeSwitchRequest) (response *ModifyAutoUpgradeSwitchResponse, err error) {
	if request == nil {
		request = NewModifyAutoUpgradeSwitchRequest()
	}
	response = NewModifyAutoUpgradeSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEmergencyVulListRequest() (request *DescribeEmergencyVulListRequest) {
	request = &DescribeEmergencyVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEmergencyVulList")
	return
}

func NewDescribeEmergencyVulListResponse() (response *DescribeEmergencyVulListResponse) {
	response = &DescribeEmergencyVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应急漏洞列表
func (c *Client) DescribeEmergencyVulList(request *DescribeEmergencyVulListRequest) (response *DescribeEmergencyVulListResponse, err error) {
	if request == nil {
		request = NewDescribeEmergencyVulListRequest()
	}
	response = NewDescribeEmergencyVulListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterCheckTimerSettingsRequest() (request *ModifyClusterCheckTimerSettingsRequest) {
	request = &ModifyClusterCheckTimerSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyClusterCheckTimerSettings")
	return
}

func NewModifyClusterCheckTimerSettingsResponse() (response *ModifyClusterCheckTimerSettingsResponse) {
	response = &ModifyClusterCheckTimerSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群检查定时设置
func (c *Client) ModifyClusterCheckTimerSettings(request *ModifyClusterCheckTimerSettingsRequest) (response *ModifyClusterCheckTimerSettingsResponse, err error) {
	if request == nil {
		request = NewModifyClusterCheckTimerSettingsRequest()
	}
	response = NewModifyClusterCheckTimerSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogDeliveryClsSettingRequest() (request *DescribeSecLogDeliveryClsSettingRequest) {
	request = &DescribeSecLogDeliveryClsSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogDeliveryClsSetting")
	return
}

func NewDescribeSecLogDeliveryClsSettingResponse() (response *DescribeSecLogDeliveryClsSettingResponse) {
	response = &DescribeSecLogDeliveryClsSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志投递Cls配置
func (c *Client) DescribeSecLogDeliveryClsSetting(request *DescribeSecLogDeliveryClsSettingRequest) (response *DescribeSecLogDeliveryClsSettingResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogDeliveryClsSettingRequest()
	}
	response = NewDescribeSecLogDeliveryClsSettingResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecLogDeliveryClsSettingRequest() (request *ModifySecLogDeliveryClsSettingRequest) {
	request = &ModifySecLogDeliveryClsSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogDeliveryClsSetting")
	return
}

func NewModifySecLogDeliveryClsSettingResponse() (response *ModifySecLogDeliveryClsSettingResponse) {
	response = &ModifySecLogDeliveryClsSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新安全日志-日志投递cls配置
func (c *Client) ModifySecLogDeliveryClsSetting(request *ModifySecLogDeliveryClsSettingRequest) (response *ModifySecLogDeliveryClsSettingResponse, err error) {
	if request == nil {
		request = NewModifySecLogDeliveryClsSettingRequest()
	}
	response = NewModifySecLogDeliveryClsSettingResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterAssetExportJobRequest() (request *CreateClusterAssetExportJobRequest) {
	request = &CreateClusterAssetExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateClusterAssetExportJob")
	return
}

func NewCreateClusterAssetExportJobResponse() (response *CreateClusterAssetExportJobResponse) {
	response = &CreateClusterAssetExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群详情导出任务,只支持单个集群的资产导出
func (c *Client) CreateClusterAssetExportJob(request *CreateClusterAssetExportJobRequest) (response *CreateClusterAssetExportJobResponse, err error) {
	if request == nil {
		request = NewCreateClusterAssetExportJobRequest()
	}
	response = NewCreateClusterAssetExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReverseShellSystemPolicyConfigRequest() (request *ModifyReverseShellSystemPolicyConfigRequest) {
	request = &ModifyReverseShellSystemPolicyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyReverseShellSystemPolicyConfig")
	return
}

func NewModifyReverseShellSystemPolicyConfigResponse() (response *ModifyReverseShellSystemPolicyConfigResponse) {
	response = &ModifyReverseShellSystemPolicyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改反弹shell系统策略配置
func (c *Client) ModifyReverseShellSystemPolicyConfig(request *ModifyReverseShellSystemPolicyConfigRequest) (response *ModifyReverseShellSystemPolicyConfigResponse, err error) {
	if request == nil {
		request = NewModifyReverseShellSystemPolicyConfigRequest()
	}
	response = NewModifyReverseShellSystemPolicyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCompliancePeriodTaskRequest() (request *ModifyCompliancePeriodTaskRequest) {
	request = &ModifyCompliancePeriodTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyCompliancePeriodTask")
	return
}

func NewModifyCompliancePeriodTaskResponse() (response *ModifyCompliancePeriodTaskResponse) {
	response = &ModifyCompliancePeriodTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改定时任务的设置，包括检测周期、开启/禁用合规基准。
func (c *Client) ModifyCompliancePeriodTask(request *ModifyCompliancePeriodTaskRequest) (response *ModifyCompliancePeriodTaskResponse, err error) {
	if request == nil {
		request = NewModifyCompliancePeriodTaskRequest()
	}
	response = NewModifyCompliancePeriodTaskResponse()
	err = c.Send(request, response)
	return
}

func NewAddEditImageAutoAuthorizedRuleRequest() (request *AddEditImageAutoAuthorizedRuleRequest) {
	request = &AddEditImageAutoAuthorizedRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddEditImageAutoAuthorizedRule")
	return
}

func NewAddEditImageAutoAuthorizedRuleResponse() (response *AddEditImageAutoAuthorizedRuleResponse) {
	response = &AddEditImageAutoAuthorizedRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或编辑本地镜像自动授权规则
func (c *Client) AddEditImageAutoAuthorizedRule(request *AddEditImageAutoAuthorizedRuleRequest) (response *AddEditImageAutoAuthorizedRuleResponse, err error) {
	if request == nil {
		request = NewAddEditImageAutoAuthorizedRuleRequest()
	}
	response = NewAddEditImageAutoAuthorizedRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageAuthorizedInfoRequest() (request *DescribeImageAuthorizedInfoRequest) {
	request = &DescribeImageAuthorizedInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageAuthorizedInfo")
	return
}

func NewDescribeImageAuthorizedInfoResponse() (response *DescribeImageAuthorizedInfoResponse) {
	response = &DescribeImageAuthorizedInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeImageAuthorizedInfo  查询镜像授权信息
func (c *Client) DescribeImageAuthorizedInfo(request *DescribeImageAuthorizedInfoRequest) (response *DescribeImageAuthorizedInfoResponse, err error) {
	if request == nil {
		request = NewDescribeImageAuthorizedInfoRequest()
	}
	response = NewDescribeImageAuthorizedInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogJoinTypeListRequest() (request *DescribeSecLogJoinTypeListRequest) {
	request = &DescribeSecLogJoinTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogJoinTypeList")
	return
}

func NewDescribeSecLogJoinTypeListResponse() (response *DescribeSecLogJoinTypeListResponse) {
	response = &DescribeSecLogJoinTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志接入列表
func (c *Client) DescribeSecLogJoinTypeList(request *DescribeSecLogJoinTypeListRequest) (response *DescribeSecLogJoinTypeListResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogJoinTypeListRequest()
	}
	response = NewDescribeSecLogJoinTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceEventTendencyRequest() (request *DescribeVulDefenceEventTendencyRequest) {
	request = &DescribeVulDefenceEventTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefenceEventTendency")
	return
}

func NewDescribeVulDefenceEventTendencyResponse() (response *DescribeVulDefenceEventTendencyResponse) {
	response = &DescribeVulDefenceEventTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞防御攻击事件趋势
func (c *Client) DescribeVulDefenceEventTendency(request *DescribeVulDefenceEventTendencyRequest) (response *DescribeVulDefenceEventTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceEventTendencyRequest()
	}
	response = NewDescribeVulDefenceEventTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessEventsRequest() (request *DescribeAbnormalProcessEventsRequest) {
	request = &DescribeAbnormalProcessEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessEvents")
	return
}

func NewDescribeAbnormalProcessEventsResponse() (response *DescribeAbnormalProcessEventsResponse) {
	response = &DescribeAbnormalProcessEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时异常进程事件列表信息
func (c *Client) DescribeAbnormalProcessEvents(request *DescribeAbnormalProcessEventsRequest) (response *DescribeAbnormalProcessEventsResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessEventsRequest()
	}
	response = NewDescribeAbnormalProcessEventsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateNetworkFirewallPolicyDetailRequest() (request *UpdateNetworkFirewallPolicyDetailRequest) {
	request = &UpdateNetworkFirewallPolicyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "UpdateNetworkFirewallPolicyDetail")
	return
}

func NewUpdateNetworkFirewallPolicyDetailResponse() (response *UpdateNetworkFirewallPolicyDetailResponse) {
	response = &UpdateNetworkFirewallPolicyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建网络策略更新任务
func (c *Client) UpdateNetworkFirewallPolicyDetail(request *UpdateNetworkFirewallPolicyDetailRequest) (response *UpdateNetworkFirewallPolicyDetailResponse, err error) {
	if request == nil {
		request = NewUpdateNetworkFirewallPolicyDetailRequest()
	}
	response = NewUpdateNetworkFirewallPolicyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecLogCleanSettingInfoRequest() (request *ModifySecLogCleanSettingInfoRequest) {
	request = &ModifySecLogCleanSettingInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogCleanSettingInfo")
	return
}

func NewModifySecLogCleanSettingInfoResponse() (response *ModifySecLogCleanSettingInfoResponse) {
	response = &ModifySecLogCleanSettingInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改安全日志清理设置信息
func (c *Client) ModifySecLogCleanSettingInfo(request *ModifySecLogCleanSettingInfoRequest) (response *ModifySecLogCleanSettingInfoResponse, err error) {
	if request == nil {
		request = NewModifySecLogCleanSettingInfoRequest()
	}
	response = NewModifySecLogCleanSettingInfoResponse()
	err = c.Send(request, response)
	return
}

func NewApplyTrialRequest() (request *ApplyTrialRequest) {
	request = &ApplyTrialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ApplyTrial")
	return
}

func NewApplyTrialResponse() (response *ApplyTrialResponse) {
	response = &ApplyTrialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键申请试用
func (c *Client) ApplyTrial(request *ApplyTrialRequest) (response *ApplyTrialResponse, err error) {
	if request == nil {
		request = NewApplyTrialRequest()
	}
	response = NewApplyTrialResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetHostDetailRequest() (request *DescribeAssetHostDetailRequest) {
	request = &DescribeAssetHostDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetHostDetail")
	return
}

func NewDescribeAssetHostDetailResponse() (response *DescribeAssetHostDetailResponse) {
	response = &DescribeAssetHostDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机详细信息
func (c *Client) DescribeAssetHostDetail(request *DescribeAssetHostDetailRequest) (response *DescribeAssetHostDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetHostDetailRequest()
	}
	response = NewDescribeAssetHostDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRiskSummaryRequest() (request *DescribeImageRiskSummaryRequest) {
	request = &DescribeImageRiskSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRiskSummary")
	return
}

func NewDescribeImageRiskSummaryResponse() (response *DescribeImageRiskSummaryResponse) {
	response = &DescribeImageRiskSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地镜像风险概览
func (c *Client) DescribeImageRiskSummary(request *DescribeImageRiskSummaryRequest) (response *DescribeImageRiskSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeImageRiskSummaryRequest()
	}
	response = NewDescribeImageRiskSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeWhiteListRequest() (request *DescribeEscapeWhiteListRequest) {
	request = &DescribeEscapeWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeWhiteList")
	return
}

func NewDescribeEscapeWhiteListResponse() (response *DescribeEscapeWhiteListResponse) {
	response = &DescribeEscapeWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询逃逸白名单
func (c *Client) DescribeEscapeWhiteList(request *DescribeEscapeWhiteListRequest) (response *DescribeEscapeWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeWhiteListRequest()
	}
	response = NewDescribeEscapeWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageDenyRuleExportJobRequest() (request *CreateImageDenyRuleExportJobRequest) {
	request = &CreateImageDenyRuleExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateImageDenyRuleExportJob")
	return
}

func NewCreateImageDenyRuleExportJobResponse() (response *CreateImageDenyRuleExportJobResponse) {
	response = &CreateImageDenyRuleExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像拦截规则导出任务
func (c *Client) CreateImageDenyRuleExportJob(request *CreateImageDenyRuleExportJobRequest) (response *CreateImageDenyRuleExportJobResponse, err error) {
	if request == nil {
		request = NewCreateImageDenyRuleExportJobRequest()
	}
	response = NewCreateImageDenyRuleExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskTargetCountRequest() (request *DescribeRiskTargetCountRequest) {
	request = &DescribeRiskTargetCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskTargetCount")
	return
}

func NewDescribeRiskTargetCountResponse() (response *DescribeRiskTargetCountResponse) {
	response = &DescribeRiskTargetCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按检查对象展示风险数量
func (c *Client) DescribeRiskTargetCount(request *DescribeRiskTargetCountRequest) (response *DescribeRiskTargetCountResponse, err error) {
	if request == nil {
		request = NewDescribeRiskTargetCountRequest()
	}
	response = NewDescribeRiskTargetCountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulExportJobRequest() (request *CreateVulExportJobRequest) {
	request = &CreateVulExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateVulExportJob")
	return
}

func NewCreateVulExportJobResponse() (response *CreateVulExportJobResponse) {
	response = &CreateVulExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地镜像漏洞列表导出
func (c *Client) CreateVulExportJob(request *CreateVulExportJobRequest) (response *CreateVulExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVulExportJobRequest()
	}
	response = NewCreateVulExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlDetailRequest() (request *DescribeAccessControlDetailRequest) {
	request = &DescribeAccessControlDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlDetail")
	return
}

func NewDescribeAccessControlDetailResponse() (response *DescribeAccessControlDetailResponse) {
	response = &DescribeAccessControlDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时访问控制事件的详细信息
func (c *Client) DescribeAccessControlDetail(request *DescribeAccessControlDetailRequest) (response *DescribeAccessControlDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlDetailRequest()
	}
	response = NewDescribeAccessControlDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryListRequest() (request *DescribeAssetImageRegistryListRequest) {
	request = &DescribeAssetImageRegistryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryList")
	return
}

func NewDescribeAssetImageRegistryListResponse() (response *DescribeAssetImageRegistryListResponse) {
	response = &DescribeAssetImageRegistryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库镜像仓库列表
func (c *Client) DescribeAssetImageRegistryList(request *DescribeAssetImageRegistryListRequest) (response *DescribeAssetImageRegistryListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryListRequest()
	}
	response = NewDescribeAssetImageRegistryListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEscapeRegexpWhiteListRequest() (request *DeleteEscapeRegexpWhiteListRequest) {
	request = &DeleteEscapeRegexpWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteEscapeRegexpWhiteList")
	return
}

func NewDeleteEscapeRegexpWhiteListResponse() (response *DeleteEscapeRegexpWhiteListResponse) {
	response = &DeleteEscapeRegexpWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除容器逃逸正则白名单
func (c *Client) DeleteEscapeRegexpWhiteList(request *DeleteEscapeRegexpWhiteListRequest) (response *DeleteEscapeRegexpWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteEscapeRegexpWhiteListRequest()
	}
	response = NewDeleteEscapeRegexpWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateImageRegistryTimingScanTaskRequest() (request *UpdateImageRegistryTimingScanTaskRequest) {
	request = &UpdateImageRegistryTimingScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "UpdateImageRegistryTimingScanTask")
	return
}

func NewUpdateImageRegistryTimingScanTaskResponse() (response *UpdateImageRegistryTimingScanTaskResponse) {
	response = &UpdateImageRegistryTimingScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库更新定时任务
func (c *Client) UpdateImageRegistryTimingScanTask(request *UpdateImageRegistryTimingScanTaskRequest) (response *UpdateImageRegistryTimingScanTaskResponse, err error) {
	if request == nil {
		request = NewUpdateImageRegistryTimingScanTaskRequest()
	}
	response = NewUpdateImageRegistryTimingScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulDefenceHostExportJobRequest() (request *CreateVulDefenceHostExportJobRequest) {
	request = &CreateVulDefenceHostExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateVulDefenceHostExportJob")
	return
}

func NewCreateVulDefenceHostExportJobResponse() (response *CreateVulDefenceHostExportJobResponse) {
	response = &CreateVulDefenceHostExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建漏洞防御主机导出任务
func (c *Client) CreateVulDefenceHostExportJob(request *CreateVulDefenceHostExportJobRequest) (response *CreateVulDefenceHostExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVulDefenceHostExportJobRequest()
	}
	response = NewCreateVulDefenceHostExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVirusScanTaskRequest() (request *CreateVirusScanTaskRequest) {
	request = &CreateVirusScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateVirusScanTask")
	return
}

func NewCreateVirusScanTaskResponse() (response *CreateVirusScanTaskResponse) {
	response = &CreateVirusScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时文件查杀一键扫描
func (c *Client) CreateVirusScanTask(request *CreateVirusScanTaskRequest) (response *CreateVirusScanTaskResponse, err error) {
	if request == nil {
		request = NewCreateVirusScanTaskRequest()
	}
	response = NewCreateVirusScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterCheckTaskRequest() (request *CreateClusterCheckTaskRequest) {
	request = &CreateClusterCheckTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateClusterCheckTask")
	return
}

func NewCreateClusterCheckTaskResponse() (response *CreateClusterCheckTaskResponse) {
	response = &CreateClusterCheckTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群检查任务，用户检查用户的集群相关风险项
func (c *Client) CreateClusterCheckTask(request *CreateClusterCheckTaskRequest) (response *CreateClusterCheckTaskResponse, err error) {
	if request == nil {
		request = NewCreateClusterCheckTaskRequest()
	}
	response = NewCreateClusterCheckTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProcessEventsExportJobRequest() (request *CreateProcessEventsExportJobRequest) {
	request = &CreateProcessEventsExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateProcessEventsExportJob")
	return
}

func NewCreateProcessEventsExportJobResponse() (response *CreateProcessEventsExportJobResponse) {
	response = &CreateProcessEventsExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建异常进程事件导出异步任务
func (c *Client) CreateProcessEventsExportJob(request *CreateProcessEventsExportJobRequest) (response *CreateProcessEventsExportJobResponse, err error) {
	if request == nil {
		request = NewCreateProcessEventsExportJobRequest()
	}
	response = NewCreateProcessEventsExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewConfirmNetworkFirewallPolicyRequest() (request *ConfirmNetworkFirewallPolicyRequest) {
	request = &ConfirmNetworkFirewallPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ConfirmNetworkFirewallPolicy")
	return
}

func NewConfirmNetworkFirewallPolicyResponse() (response *ConfirmNetworkFirewallPolicyResponse) {
	response = &ConfirmNetworkFirewallPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络创建网络策略确认任务
func (c *Client) ConfirmNetworkFirewallPolicy(request *ConfirmNetworkFirewallPolicyRequest) (response *ConfirmNetworkFirewallPolicyResponse, err error) {
	if request == nil {
		request = NewConfirmNetworkFirewallPolicyRequest()
	}
	response = NewConfirmNetworkFirewallPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusAutoIsolateSampleDownloadURLRequest() (request *DescribeVirusAutoIsolateSampleDownloadURLRequest) {
	request = &DescribeVirusAutoIsolateSampleDownloadURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusAutoIsolateSampleDownloadURL")
	return
}

func NewDescribeVirusAutoIsolateSampleDownloadURLResponse() (response *DescribeVirusAutoIsolateSampleDownloadURLResponse) {
	response = &DescribeVirusAutoIsolateSampleDownloadURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马自动隔离样本下载链接
func (c *Client) DescribeVirusAutoIsolateSampleDownloadURL(request *DescribeVirusAutoIsolateSampleDownloadURLRequest) (response *DescribeVirusAutoIsolateSampleDownloadURLResponse, err error) {
	if request == nil {
		request = NewDescribeVirusAutoIsolateSampleDownloadURLRequest()
	}
	response = NewDescribeVirusAutoIsolateSampleDownloadURLResponse()
	err = c.Send(request, response)
	return
}

func NewCheckRepeatAssetImageRegistryRequest() (request *CheckRepeatAssetImageRegistryRequest) {
	request = &CheckRepeatAssetImageRegistryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CheckRepeatAssetImageRegistry")
	return
}

func NewCheckRepeatAssetImageRegistryResponse() (response *CheckRepeatAssetImageRegistryResponse) {
	response = &CheckRepeatAssetImageRegistryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查单个镜像仓库名是否重复
func (c *Client) CheckRepeatAssetImageRegistry(request *CheckRepeatAssetImageRegistryRequest) (response *CheckRepeatAssetImageRegistryResponse, err error) {
	if request == nil {
		request = NewCheckRepeatAssetImageRegistryRequest()
	}
	response = NewCheckRepeatAssetImageRegistryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessRulesExportRequest() (request *DescribeAbnormalProcessRulesExportRequest) {
	request = &DescribeAbnormalProcessRulesExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessRulesExport")
	return
}

func NewDescribeAbnormalProcessRulesExportResponse() (response *DescribeAbnormalProcessRulesExportResponse) {
	response = &DescribeAbnormalProcessRulesExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时异常进程策略列表信息导出
func (c *Client) DescribeAbnormalProcessRulesExport(request *DescribeAbnormalProcessRulesExportRequest) (response *DescribeAbnormalProcessRulesExportResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessRulesExportRequest()
	}
	response = NewDescribeAbnormalProcessRulesExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallPolicyListRequest() (request *DescribeNetworkFirewallPolicyListRequest) {
	request = &DescribeNetworkFirewallPolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPolicyList")
	return
}

func NewDescribeNetworkFirewallPolicyListResponse() (response *DescribeNetworkFirewallPolicyListResponse) {
	response = &DescribeNetworkFirewallPolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群网络策略列表
func (c *Client) DescribeNetworkFirewallPolicyList(request *DescribeNetworkFirewallPolicyListRequest) (response *DescribeNetworkFirewallPolicyListResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallPolicyListRequest()
	}
	response = NewDescribeNetworkFirewallPolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageAutoAuthorizedTaskListRequest() (request *DescribeImageAutoAuthorizedTaskListRequest) {
	request = &DescribeImageAutoAuthorizedTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageAutoAuthorizedTaskList")
	return
}

func NewDescribeImageAutoAuthorizedTaskListResponse() (response *DescribeImageAutoAuthorizedTaskListResponse) {
	response = &DescribeImageAutoAuthorizedTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像自动授权任务列表
func (c *Client) DescribeImageAutoAuthorizedTaskList(request *DescribeImageAutoAuthorizedTaskListRequest) (response *DescribeImageAutoAuthorizedTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeImageAutoAuthorizedTaskListRequest()
	}
	response = NewDescribeImageAutoAuthorizedTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryAssetStatusRequest() (request *DescribeAssetImageRegistryAssetStatusRequest) {
	request = &DescribeAssetImageRegistryAssetStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryAssetStatus")
	return
}

func NewDescribeAssetImageRegistryAssetStatusResponse() (response *DescribeAssetImageRegistryAssetStatusResponse) {
	response = &DescribeAssetImageRegistryAssetStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看镜像仓库资产更新进度状态
func (c *Client) DescribeAssetImageRegistryAssetStatus(request *DescribeAssetImageRegistryAssetStatusRequest) (response *DescribeAssetImageRegistryAssetStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryAssetStatusRequest()
	}
	response = NewDescribeAssetImageRegistryAssetStatusResponse()
	err = c.Send(request, response)
	return
}

func NewClearLocalStorageRequest() (request *ClearLocalStorageRequest) {
	request = &ClearLocalStorageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ClearLocalStorage")
	return
}

func NewClearLocalStorageResponse() (response *ClearLocalStorageResponse) {
	response = &ClearLocalStorageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清理本地存储数据
func (c *Client) ClearLocalStorage(request *ClearLocalStorageRequest) (response *ClearLocalStorageResponse, err error) {
	if request == nil {
		request = NewClearLocalStorageRequest()
	}
	response = NewClearLocalStorageResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterAccessRequest() (request *CreateClusterAccessRequest) {
	request = &CreateClusterAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateClusterAccess")
	return
}

func NewCreateClusterAccessResponse() (response *CreateClusterAccessResponse) {
	response = &CreateClusterAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群接入
func (c *Client) CreateClusterAccess(request *CreateClusterAccessRequest) (response *CreateClusterAccessResponse, err error) {
	if request == nil {
		request = NewCreateClusterAccessRequest()
	}
	response = NewCreateClusterAccessResponse()
	err = c.Send(request, response)
	return
}

func NewCreateK8sApiAbnormalRuleInfoRequest() (request *CreateK8sApiAbnormalRuleInfoRequest) {
	request = &CreateK8sApiAbnormalRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateK8sApiAbnormalRuleInfo")
	return
}

func NewCreateK8sApiAbnormalRuleInfoResponse() (response *CreateK8sApiAbnormalRuleInfoResponse) {
	response = &CreateK8sApiAbnormalRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建k8sapi异常事件规则
func (c *Client) CreateK8sApiAbnormalRuleInfo(request *CreateK8sApiAbnormalRuleInfoRequest) (response *CreateK8sApiAbnormalRuleInfoResponse, err error) {
	if request == nil {
		request = NewCreateK8sApiAbnormalRuleInfoRequest()
	}
	response = NewCreateK8sApiAbnormalRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryRiskInfoListRequest() (request *DescribeAssetImageRegistryRiskInfoListRequest) {
	request = &DescribeAssetImageRegistryRiskInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRiskInfoList")
	return
}

func NewDescribeAssetImageRegistryRiskInfoListResponse() (response *DescribeAssetImageRegistryRiskInfoListResponse) {
	response = &DescribeAssetImageRegistryRiskInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询镜像高危行为列表
func (c *Client) DescribeAssetImageRegistryRiskInfoList(request *DescribeAssetImageRegistryRiskInfoListRequest) (response *DescribeAssetImageRegistryRiskInfoListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryRiskInfoListRequest()
	}
	response = NewDescribeAssetImageRegistryRiskInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSearchTemplateRequest() (request *CreateSearchTemplateRequest) {
	request = &CreateSearchTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateSearchTemplate")
	return
}

func NewCreateSearchTemplateResponse() (response *CreateSearchTemplateResponse) {
	response = &CreateSearchTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加检索模板
func (c *Client) CreateSearchTemplate(request *CreateSearchTemplateRequest) (response *CreateSearchTemplateResponse, err error) {
	if request == nil {
		request = NewCreateSearchTemplateRequest()
	}
	response = NewCreateSearchTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusDetailRequest() (request *DescribeVirusDetailRequest) {
	request = &DescribeVirusDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusDetail")
	return
}

func NewDescribeVirusDetailResponse() (response *DescribeVirusDetailResponse) {
	response = &DescribeVirusDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询木马文件信息
func (c *Client) DescribeVirusDetail(request *DescribeVirusDetailRequest) (response *DescribeVirusDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVirusDetailRequest()
	}
	response = NewDescribeVirusDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEventEscapeImageExportJobRequest() (request *CreateEventEscapeImageExportJobRequest) {
	request = &CreateEventEscapeImageExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateEventEscapeImageExportJob")
	return
}

func NewCreateEventEscapeImageExportJobResponse() (response *CreateEventEscapeImageExportJobResponse) {
	response = &CreateEventEscapeImageExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建风险容器镜像导出任务
func (c *Client) CreateEventEscapeImageExportJob(request *CreateEventEscapeImageExportJobRequest) (response *CreateEventEscapeImageExportJobResponse, err error) {
	if request == nil {
		request = NewCreateEventEscapeImageExportJobRequest()
	}
	response = NewCreateEventEscapeImageExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryRegistryListRequest() (request *DescribeAssetImageRegistryRegistryListRequest) {
	request = &DescribeAssetImageRegistryRegistryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRegistryList")
	return
}

func NewDescribeAssetImageRegistryRegistryListResponse() (response *DescribeAssetImageRegistryRegistryListResponse) {
	response = &DescribeAssetImageRegistryRegistryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库仓库列表
func (c *Client) DescribeAssetImageRegistryRegistryList(request *DescribeAssetImageRegistryRegistryListRequest) (response *DescribeAssetImageRegistryRegistryListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryRegistryListRequest()
	}
	response = NewDescribeAssetImageRegistryRegistryListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCompliancePeriodTaskListRequest() (request *DescribeCompliancePeriodTaskListRequest) {
	request = &DescribeCompliancePeriodTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeCompliancePeriodTaskList")
	return
}

func NewDescribeCompliancePeriodTaskListResponse() (response *DescribeCompliancePeriodTaskListResponse) {
	response = &DescribeCompliancePeriodTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询合规检测的定时任务列表
func (c *Client) DescribeCompliancePeriodTaskList(request *DescribeCompliancePeriodTaskListRequest) (response *DescribeCompliancePeriodTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeCompliancePeriodTaskListRequest()
	}
	response = NewDescribeCompliancePeriodTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAssetImageAuthInfoRequest() (request *UpdateAssetImageAuthInfoRequest) {
	request = &UpdateAssetImageAuthInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "UpdateAssetImageAuthInfo")
	return
}

func NewUpdateAssetImageAuthInfoResponse() (response *UpdateAssetImageAuthInfoResponse) {
	response = &UpdateAssetImageAuthInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新仓库镜像授权信息
func (c *Client) UpdateAssetImageAuthInfo(request *UpdateAssetImageAuthInfoRequest) (response *UpdateAssetImageAuthInfoResponse, err error) {
	if request == nil {
		request = NewUpdateAssetImageAuthInfoRequest()
	}
	response = NewUpdateAssetImageAuthInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulnerabilityFocusRuleDetailRequest() (request *DescribeVulnerabilityFocusRuleDetailRequest) {
	request = &DescribeVulnerabilityFocusRuleDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulnerabilityFocusRuleDetail")
	return
}

func NewDescribeVulnerabilityFocusRuleDetailResponse() (response *DescribeVulnerabilityFocusRuleDetailResponse) {
	response = &DescribeVulnerabilityFocusRuleDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞重点关注规则
func (c *Client) DescribeVulnerabilityFocusRuleDetail(request *DescribeVulnerabilityFocusRuleDetailRequest) (response *DescribeVulnerabilityFocusRuleDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVulnerabilityFocusRuleDetailRequest()
	}
	response = NewDescribeVulnerabilityFocusRuleDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeRegexpWhiteListInfoRequest() (request *DescribeEscapeRegexpWhiteListInfoRequest) {
	request = &DescribeEscapeRegexpWhiteListInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeRegexpWhiteListInfo")
	return
}

func NewDescribeEscapeRegexpWhiteListInfoResponse() (response *DescribeEscapeRegexpWhiteListInfoResponse) {
	response = &DescribeEscapeRegexpWhiteListInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器逃逸白名单详情
func (c *Client) DescribeEscapeRegexpWhiteListInfo(request *DescribeEscapeRegexpWhiteListInfoRequest) (response *DescribeEscapeRegexpWhiteListInfoResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeRegexpWhiteListInfoRequest()
	}
	response = NewDescribeEscapeRegexpWhiteListInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserClusterRequest() (request *DescribeUserClusterRequest) {
	request = &DescribeUserClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeUserCluster")
	return
}

func NewDescribeUserClusterResponse() (response *DescribeUserClusterResponse) {
	response = &DescribeUserClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全概览和集群安全页进入调用该接口，查询用户集群相关信息。
func (c *Client) DescribeUserCluster(request *DescribeUserClusterRequest) (response *DescribeUserClusterResponse, err error) {
	if request == nil {
		request = NewDescribeUserClusterRequest()
	}
	response = NewDescribeUserClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessLevelSummaryRequest() (request *DescribeAbnormalProcessLevelSummaryRequest) {
	request = &DescribeAbnormalProcessLevelSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessLevelSummary")
	return
}

func NewDescribeAbnormalProcessLevelSummaryResponse() (response *DescribeAbnormalProcessLevelSummaryResponse) {
	response = &DescribeAbnormalProcessLevelSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统计异常进程各威胁等级待处理事件数
func (c *Client) DescribeAbnormalProcessLevelSummary(request *DescribeAbnormalProcessLevelSummaryRequest) (response *DescribeAbnormalProcessLevelSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessLevelSummaryRequest()
	}
	response = NewDescribeAbnormalProcessLevelSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellRegexpWhiteListInfoRequest() (request *DescribeReverseShellRegexpWhiteListInfoRequest) {
	request = &DescribeReverseShellRegexpWhiteListInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellRegexpWhiteListInfo")
	return
}

func NewDescribeReverseShellRegexpWhiteListInfoResponse() (response *DescribeReverseShellRegexpWhiteListInfoResponse) {
	response = &DescribeReverseShellRegexpWhiteListInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询反弹shell正则白名单详情
func (c *Client) DescribeReverseShellRegexpWhiteListInfo(request *DescribeReverseShellRegexpWhiteListInfoRequest) (response *DescribeReverseShellRegexpWhiteListInfoResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellRegexpWhiteListInfoRequest()
	}
	response = NewDescribeReverseShellRegexpWhiteListInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecLogDeliveryKafkaSettingRequest() (request *ModifySecLogDeliveryKafkaSettingRequest) {
	request = &ModifySecLogDeliveryKafkaSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogDeliveryKafkaSetting")
	return
}

func NewModifySecLogDeliveryKafkaSettingResponse() (response *ModifySecLogDeliveryKafkaSettingResponse) {
	response = &ModifySecLogDeliveryKafkaSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新安全日志投递kafka设置
func (c *Client) ModifySecLogDeliveryKafkaSetting(request *ModifySecLogDeliveryKafkaSettingRequest) (response *ModifySecLogDeliveryKafkaSettingResponse, err error) {
	if request == nil {
		request = NewModifySecLogDeliveryKafkaSettingRequest()
	}
	response = NewModifySecLogDeliveryKafkaSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulSummaryRequest() (request *DescribeVulSummaryRequest) {
	request = &DescribeVulSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulSummary")
	return
}

func NewDescribeVulSummaryResponse() (response *DescribeVulSummaryResponse) {
	response = &DescribeVulSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞风险统计概览
func (c *Client) DescribeVulSummary(request *DescribeVulSummaryRequest) (response *DescribeVulSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeVulSummaryRequest()
	}
	response = NewDescribeVulSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrialStatusRequest() (request *DescribeTrialStatusRequest) {
	request = &DescribeTrialStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeTrialStatus")
	return
}

func NewDescribeTrialStatusResponse() (response *DescribeTrialStatusResponse) {
	response = &DescribeTrialStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询是否可试用
func (c *Client) DescribeTrialStatus(request *DescribeTrialStatusRequest) (response *DescribeTrialStatusResponse, err error) {
	if request == nil {
		request = NewDescribeTrialStatusRequest()
	}
	response = NewDescribeTrialStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserWorkloadKindsRequest() (request *DescribeUserWorkloadKindsRequest) {
	request = &DescribeUserWorkloadKindsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeUserWorkloadKinds")
	return
}

func NewDescribeUserWorkloadKindsResponse() (response *DescribeUserWorkloadKindsResponse) {
	response = &DescribeUserWorkloadKindsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 枚举用户所有的工作负载类型
func (c *Client) DescribeUserWorkloadKinds(request *DescribeUserWorkloadKindsRequest) (response *DescribeUserWorkloadKindsResponse, err error) {
	if request == nil {
		request = NewDescribeUserWorkloadKindsRequest()
	}
	response = NewDescribeUserWorkloadKindsResponse()
	err = c.Send(request, response)
	return
}

func NewSyncComplianceHostAssetsRequest() (request *SyncComplianceHostAssetsRequest) {
	request = &SyncComplianceHostAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "SyncComplianceHostAssets")
	return
}

func NewSyncComplianceHostAssetsResponse() (response *SyncComplianceHostAssetsResponse) {
	response = &SyncComplianceHostAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步合规安全扫描的主机列表
func (c *Client) SyncComplianceHostAssets(request *SyncComplianceHostAssetsRequest) (response *SyncComplianceHostAssetsResponse, err error) {
	if request == nil {
		request = NewSyncComplianceHostAssetsRequest()
	}
	response = NewSyncComplianceHostAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReverseShellRegexpWhiteListRequest() (request *DeleteReverseShellRegexpWhiteListRequest) {
	request = &DeleteReverseShellRegexpWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteReverseShellRegexpWhiteList")
	return
}

func NewDeleteReverseShellRegexpWhiteListResponse() (response *DeleteReverseShellRegexpWhiteListResponse) {
	response = &DeleteReverseShellRegexpWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除反弹shell正则白名单
func (c *Client) DeleteReverseShellRegexpWhiteList(request *DeleteReverseShellRegexpWhiteListRequest) (response *DeleteReverseShellRegexpWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteReverseShellRegexpWhiteListRequest()
	}
	response = NewDeleteReverseShellRegexpWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterServiceListRequest() (request *DescribeClusterServiceListRequest) {
	request = &DescribeClusterServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterServiceList")
	return
}

func NewDescribeClusterServiceListResponse() (response *DescribeClusterServiceListResponse) {
	response = &DescribeClusterServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群service列表
func (c *Client) DescribeClusterServiceList(request *DescribeClusterServiceListRequest) (response *DescribeClusterServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeClusterServiceListRequest()
	}
	response = NewDescribeClusterServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRegistryAutoAuthorizedRuleImageRequest() (request *DescribeImageRegistryAutoAuthorizedRuleImageRequest) {
	request = &DescribeImageRegistryAutoAuthorizedRuleImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRegistryAutoAuthorizedRuleImage")
	return
}

func NewDescribeImageRegistryAutoAuthorizedRuleImageResponse() (response *DescribeImageRegistryAutoAuthorizedRuleImageResponse) {
	response = &DescribeImageRegistryAutoAuthorizedRuleImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像仓库自动授权的范围详情
func (c *Client) DescribeImageRegistryAutoAuthorizedRuleImage(request *DescribeImageRegistryAutoAuthorizedRuleImageRequest) (response *DescribeImageRegistryAutoAuthorizedRuleImageResponse, err error) {
	if request == nil {
		request = NewDescribeImageRegistryAutoAuthorizedRuleImageRequest()
	}
	response = NewDescribeImageRegistryAutoAuthorizedRuleImageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterSummaryRequest() (request *DescribeClusterSummaryRequest) {
	request = &DescribeClusterSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterSummary")
	return
}

func NewDescribeClusterSummaryResponse() (response *DescribeClusterSummaryResponse) {
	response = &DescribeClusterSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户集群资产总览
func (c *Client) DescribeClusterSummary(request *DescribeClusterSummaryRequest) (response *DescribeClusterSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeClusterSummaryRequest()
	}
	response = NewDescribeClusterSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusScanTaskStatusRequest() (request *DescribeVirusScanTaskStatusRequest) {
	request = &DescribeVirusScanTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusScanTaskStatus")
	return
}

func NewDescribeVirusScanTaskStatusResponse() (response *DescribeVirusScanTaskStatusResponse) {
	response = &DescribeVirusScanTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询文件查杀任务状态
func (c *Client) DescribeVirusScanTaskStatus(request *DescribeVirusScanTaskStatusRequest) (response *DescribeVirusScanTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVirusScanTaskStatusRequest()
	}
	response = NewDescribeVirusScanTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIACCheckResultListRequest() (request *DescribeIACCheckResultListRequest) {
	request = &DescribeIACCheckResultListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeIACCheckResultList")
	return
}

func NewDescribeIACCheckResultListResponse() (response *DescribeIACCheckResultListResponse) {
	response = &DescribeIACCheckResultListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询IAC检测任务结果列表
func (c *Client) DescribeIACCheckResultList(request *DescribeIACCheckResultListRequest) (response *DescribeIACCheckResultListResponse, err error) {
	if request == nil {
		request = NewDescribeIACCheckResultListRequest()
	}
	response = NewDescribeIACCheckResultListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAndPublishNetworkFirewallPolicyYamlDetailRequest() (request *UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest) {
	request = &UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "UpdateAndPublishNetworkFirewallPolicyYamlDetail")
	return
}

func NewUpdateAndPublishNetworkFirewallPolicyYamlDetailResponse() (response *UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse) {
	response = &UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络更新Yaml网络策略并发布任务
func (c *Client) UpdateAndPublishNetworkFirewallPolicyYamlDetail(request *UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest) (response *UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse, err error) {
	if request == nil {
		request = NewUpdateAndPublishNetworkFirewallPolicyYamlDetailRequest()
	}
	response = NewUpdateAndPublishNetworkFirewallPolicyYamlDetailResponse()
	err = c.Send(request, response)
	return
}

func NewAddModEscapeRegexpWhiteListRequest() (request *AddModEscapeRegexpWhiteListRequest) {
	request = &AddModEscapeRegexpWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddModEscapeRegexpWhiteList")
	return
}

func NewAddModEscapeRegexpWhiteListResponse() (response *AddModEscapeRegexpWhiteListResponse) {
	response = &AddModEscapeRegexpWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加修改容器逃逸正则白名单
func (c *Client) AddModEscapeRegexpWhiteList(request *AddModEscapeRegexpWhiteListRequest) (response *AddModEscapeRegexpWhiteListResponse, err error) {
	if request == nil {
		request = NewAddModEscapeRegexpWhiteListRequest()
	}
	response = NewAddModEscapeRegexpWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNodesRequest() (request *DescribeClusterNodesRequest) {
	request = &DescribeClusterNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterNodes")
	return
}

func NewDescribeClusterNodesResponse() (response *DescribeClusterNodesResponse) {
	response = &DescribeClusterNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群节点信息
func (c *Client) DescribeClusterNodes(request *DescribeClusterNodesRequest) (response *DescribeClusterNodesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNodesRequest()
	}
	response = NewDescribeClusterNodesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserPodListRequest() (request *DescribeUserPodListRequest) {
	request = &DescribeUserPodListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeUserPodList")
	return
}

func NewDescribeUserPodListResponse() (response *DescribeUserPodListResponse) {
	response = &DescribeUserPodListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的pod列表
func (c *Client) DescribeUserPodList(request *DescribeUserPodListRequest) (response *DescribeUserPodListResponse, err error) {
	if request == nil {
		request = NewDescribeUserPodListRequest()
	}
	response = NewDescribeUserPodListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyEventTendencyRequest() (request *DescribeImageDenyEventTendencyRequest) {
	request = &DescribeImageDenyEventTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageDenyEventTendency")
	return
}

func NewDescribeImageDenyEventTendencyResponse() (response *DescribeImageDenyEventTendencyResponse) {
	response = &DescribeImageDenyEventTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截事件趋势
func (c *Client) DescribeImageDenyEventTendency(request *DescribeImageDenyEventTendencyRequest) (response *DescribeImageDenyEventTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyEventTendencyRequest()
	}
	response = NewDescribeImageDenyEventTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryDetailRequest() (request *DescribeAssetImageRegistryDetailRequest) {
	request = &DescribeAssetImageRegistryDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryDetail")
	return
}

func NewDescribeAssetImageRegistryDetailResponse() (response *DescribeAssetImageRegistryDetailResponse) {
	response = &DescribeAssetImageRegistryDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库镜像仓库列表详情
func (c *Client) DescribeAssetImageRegistryDetail(request *DescribeAssetImageRegistryDetailRequest) (response *DescribeAssetImageRegistryDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryDetailRequest()
	}
	response = NewDescribeAssetImageRegistryDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkFirewallPolicyDiscoverRequest() (request *CreateNetworkFirewallPolicyDiscoverRequest) {
	request = &CreateNetworkFirewallPolicyDiscoverRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateNetworkFirewallPolicyDiscover")
	return
}

func NewCreateNetworkFirewallPolicyDiscoverResponse() (response *CreateNetworkFirewallPolicyDiscoverResponse) {
	response = &CreateNetworkFirewallPolicyDiscoverResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器网络集群网络策略创建自动发现任务
func (c *Client) CreateNetworkFirewallPolicyDiscover(request *CreateNetworkFirewallPolicyDiscoverRequest) (response *CreateNetworkFirewallPolicyDiscoverResponse, err error) {
	if request == nil {
		request = NewCreateNetworkFirewallPolicyDiscoverRequest()
	}
	response = NewCreateNetworkFirewallPolicyDiscoverResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchTemplatesRequest() (request *DescribeSearchTemplatesRequest) {
	request = &DescribeSearchTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSearchTemplates")
	return
}

func NewDescribeSearchTemplatesResponse() (response *DescribeSearchTemplatesResponse) {
	response = &DescribeSearchTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取快速检索列表
func (c *Client) DescribeSearchTemplates(request *DescribeSearchTemplatesRequest) (response *DescribeSearchTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeSearchTemplatesRequest()
	}
	response = NewDescribeSearchTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTcssSummaryRequest() (request *DescribeTcssSummaryRequest) {
	request = &DescribeTcssSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeTcssSummary")
	return
}

func NewDescribeTcssSummaryResponse() (response *DescribeTcssSummaryResponse) {
	response = &DescribeTcssSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器安全概览信息
func (c *Client) DescribeTcssSummary(request *DescribeTcssSummaryRequest) (response *DescribeTcssSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeTcssSummaryRequest()
	}
	response = NewDescribeTcssSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogDeliveryKafkaSettingRequest() (request *DescribeSecLogDeliveryKafkaSettingRequest) {
	request = &DescribeSecLogDeliveryKafkaSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogDeliveryKafkaSetting")
	return
}

func NewDescribeSecLogDeliveryKafkaSettingResponse() (response *DescribeSecLogDeliveryKafkaSettingResponse) {
	response = &DescribeSecLogDeliveryKafkaSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志投递kafka配置
func (c *Client) DescribeSecLogDeliveryKafkaSetting(request *DescribeSecLogDeliveryKafkaSettingRequest) (response *DescribeSecLogDeliveryKafkaSettingResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogDeliveryKafkaSettingRequest()
	}
	response = NewDescribeSecLogDeliveryKafkaSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellRegexpWhiteListRequest() (request *DescribeReverseShellRegexpWhiteListRequest) {
	request = &DescribeReverseShellRegexpWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellRegexpWhiteList")
	return
}

func NewDescribeReverseShellRegexpWhiteListResponse() (response *DescribeReverseShellRegexpWhiteListResponse) {
	response = &DescribeReverseShellRegexpWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询反弹shell正则白名单列表
func (c *Client) DescribeReverseShellRegexpWhiteList(request *DescribeReverseShellRegexpWhiteListRequest) (response *DescribeReverseShellRegexpWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellRegexpWhiteListRequest()
	}
	response = NewDescribeReverseShellRegexpWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusEstimateScanTimeRequest() (request *DescribeVirusEstimateScanTimeRequest) {
	request = &DescribeVirusEstimateScanTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusEstimateScanTime")
	return
}

func NewDescribeVirusEstimateScanTimeResponse() (response *DescribeVirusEstimateScanTimeResponse) {
	response = &DescribeVirusEstimateScanTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马一键检测预估扫描时间
func (c *Client) DescribeVirusEstimateScanTime(request *DescribeVirusEstimateScanTimeRequest) (response *DescribeVirusEstimateScanTimeResponse, err error) {
	if request == nil {
		request = NewDescribeVirusEstimateScanTimeRequest()
	}
	response = NewDescribeVirusEstimateScanTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageVulListExportRequest() (request *DescribeAssetImageVulListExportRequest) {
	request = &DescribeAssetImageVulListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVulListExport")
	return
}

func NewDescribeAssetImageVulListExportResponse() (response *DescribeAssetImageVulListExportResponse) {
	response = &DescribeAssetImageVulListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全搜索查询镜像漏洞列表导出
func (c *Client) DescribeAssetImageVulListExport(request *DescribeAssetImageVulListExportRequest) (response *DescribeAssetImageVulListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageVulListExportRequest()
	}
	response = NewDescribeAssetImageVulListExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWarningRulesRequest() (request *DescribeWarningRulesRequest) {
	request = &DescribeWarningRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeWarningRules")
	return
}

func NewDescribeWarningRulesResponse() (response *DescribeWarningRulesResponse) {
	response = &DescribeWarningRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警策略列表
func (c *Client) DescribeWarningRules(request *DescribeWarningRulesRequest) (response *DescribeWarningRulesResponse, err error) {
	if request == nil {
		request = NewDescribeWarningRulesRequest()
	}
	response = NewDescribeWarningRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallNamesRequest() (request *DescribeRiskSyscallNamesRequest) {
	request = &DescribeRiskSyscallNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallNames")
	return
}

func NewDescribeRiskSyscallNamesResponse() (response *DescribeRiskSyscallNamesResponse) {
	response = &DescribeRiskSyscallNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时高危系统调用系统名称列表
func (c *Client) DescribeRiskSyscallNames(request *DescribeRiskSyscallNamesRequest) (response *DescribeRiskSyscallNamesResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallNamesRequest()
	}
	response = NewDescribeRiskSyscallNamesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageListRequest() (request *DescribeAssetImageListRequest) {
	request = &DescribeAssetImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageList")
	return
}

func NewDescribeAssetImageListResponse() (response *DescribeAssetImageListResponse) {
	response = &DescribeAssetImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全搜索查询镜像列表
func (c *Client) DescribeAssetImageList(request *DescribeAssetImageListRequest) (response *DescribeAssetImageListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageListRequest()
	}
	response = NewDescribeAssetImageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskResultSummaryRequest() (request *DescribeTaskResultSummaryRequest) {
	request = &DescribeTaskResultSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeTaskResultSummary")
	return
}

func NewDescribeTaskResultSummaryResponse() (response *DescribeTaskResultSummaryResponse) {
	response = &DescribeTaskResultSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询检查结果总览，返回受影响的节点数量，返回7天的数据，总共7个
func (c *Client) DescribeTaskResultSummary(request *DescribeTaskResultSummaryRequest) (response *DescribeTaskResultSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeTaskResultSummaryRequest()
	}
	response = NewDescribeTaskResultSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageScanSettingRequest() (request *DescribeAssetImageScanSettingRequest) {
	request = &DescribeAssetImageScanSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageScanSetting")
	return
}

func NewDescribeAssetImageScanSettingResponse() (response *DescribeAssetImageScanSettingResponse) {
	response = &DescribeAssetImageScanSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像扫描设置信息
func (c *Client) DescribeAssetImageScanSetting(request *DescribeAssetImageScanSettingRequest) (response *DescribeAssetImageScanSettingResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageScanSettingRequest()
	}
	response = NewDescribeAssetImageScanSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageListExportRequest() (request *DescribeAssetImageListExportRequest) {
	request = &DescribeAssetImageListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageListExport")
	return
}

func NewDescribeAssetImageListExportResponse() (response *DescribeAssetImageListExportResponse) {
	response = &DescribeAssetImageListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全搜索查询镜像列表导出
func (c *Client) DescribeAssetImageListExport(request *DescribeAssetImageListExportRequest) (response *DescribeAssetImageListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageListExportRequest()
	}
	response = NewDescribeAssetImageListExportResponse()
	err = c.Send(request, response)
	return
}

func NewCleanTrialRequest() (request *CleanTrialRequest) {
	request = &CleanTrialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CleanTrial")
	return
}

func NewCleanTrialResponse() (response *CleanTrialResponse) {
	response = &CleanTrialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键重置试用
func (c *Client) CleanTrial(request *CleanTrialRequest) (response *CleanTrialResponse, err error) {
	if request == nil {
		request = NewCleanTrialRequest()
	}
	response = NewCleanTrialResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterWorkloadWhiteListRequest() (request *DescribeClusterWorkloadWhiteListRequest) {
	request = &DescribeClusterWorkloadWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterWorkloadWhiteList")
	return
}

func NewDescribeClusterWorkloadWhiteListResponse() (response *DescribeClusterWorkloadWhiteListResponse) {
	response = &DescribeClusterWorkloadWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群工作负载白名单
func (c *Client) DescribeClusterWorkloadWhiteList(request *DescribeClusterWorkloadWhiteListRequest) (response *DescribeClusterWorkloadWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeClusterWorkloadWhiteListRequest()
	}
	response = NewDescribeClusterWorkloadWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallClusterListRequest() (request *DescribeNetworkFirewallClusterListRequest) {
	request = &DescribeNetworkFirewallClusterListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallClusterList")
	return
}

func NewDescribeNetworkFirewallClusterListResponse() (response *DescribeNetworkFirewallClusterListResponse) {
	response = &DescribeNetworkFirewallClusterListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群策略列表
func (c *Client) DescribeNetworkFirewallClusterList(request *DescribeNetworkFirewallClusterListRequest) (response *DescribeNetworkFirewallClusterListResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallClusterListRequest()
	}
	response = NewDescribeNetworkFirewallClusterListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReverseShellWhiteListsRequest() (request *DeleteReverseShellWhiteListsRequest) {
	request = &DeleteReverseShellWhiteListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteReverseShellWhiteLists")
	return
}

func NewDeleteReverseShellWhiteListsResponse() (response *DeleteReverseShellWhiteListsResponse) {
	response = &DeleteReverseShellWhiteListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运行时反弹shell白名单
func (c *Client) DeleteReverseShellWhiteLists(request *DeleteReverseShellWhiteListsRequest) (response *DeleteReverseShellWhiteListsResponse, err error) {
	if request == nil {
		request = NewDeleteReverseShellWhiteListsRequest()
	}
	response = NewDeleteReverseShellWhiteListsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserClusterExportJobRequest() (request *CreateUserClusterExportJobRequest) {
	request = &CreateUserClusterExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateUserClusterExportJob")
	return
}

func NewCreateUserClusterExportJobResponse() (response *CreateUserClusterExportJobResponse) {
	response = &CreateUserClusterExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群信息导出任务，可以指定一个或者多个导出
func (c *Client) CreateUserClusterExportJob(request *CreateUserClusterExportJobRequest) (response *CreateUserClusterExportJobResponse, err error) {
	if request == nil {
		request = NewCreateUserClusterExportJobRequest()
	}
	response = NewCreateUserClusterExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRiskSyscallWhiteListsRequest() (request *DeleteRiskSyscallWhiteListsRequest) {
	request = &DeleteRiskSyscallWhiteListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteRiskSyscallWhiteLists")
	return
}

func NewDeleteRiskSyscallWhiteListsResponse() (response *DeleteRiskSyscallWhiteListsResponse) {
	response = &DeleteRiskSyscallWhiteListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运行时高危系统调用白名单
func (c *Client) DeleteRiskSyscallWhiteLists(request *DeleteRiskSyscallWhiteListsRequest) (response *DeleteRiskSyscallWhiteListsResponse, err error) {
	if request == nil {
		request = NewDeleteRiskSyscallWhiteListsRequest()
	}
	response = NewDeleteRiskSyscallWhiteListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryVirusListRequest() (request *DescribeAssetImageRegistryVirusListRequest) {
	request = &DescribeAssetImageRegistryVirusListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVirusList")
	return
}

func NewDescribeAssetImageRegistryVirusListResponse() (response *DescribeAssetImageRegistryVirusListResponse) {
	response = &DescribeAssetImageRegistryVirusListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询木马病毒列表
func (c *Client) DescribeAssetImageRegistryVirusList(request *DescribeAssetImageRegistryVirusListRequest) (response *DescribeAssetImageRegistryVirusListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryVirusListRequest()
	}
	response = NewDescribeAssetImageRegistryVirusListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkFirewallAuditRecordRequest() (request *DescribeNetworkFirewallAuditRecordRequest) {
	request = &DescribeNetworkFirewallAuditRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallAuditRecord")
	return
}

func NewDescribeNetworkFirewallAuditRecordResponse() (response *DescribeNetworkFirewallAuditRecordResponse) {
	response = &DescribeNetworkFirewallAuditRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群策略审计列表
func (c *Client) DescribeNetworkFirewallAuditRecord(request *DescribeNetworkFirewallAuditRecordRequest) (response *DescribeNetworkFirewallAuditRecordResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkFirewallAuditRecordRequest()
	}
	response = NewDescribeNetworkFirewallAuditRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusScanSettingRequest() (request *DescribeVirusScanSettingRequest) {
	request = &DescribeVirusScanSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusScanSetting")
	return
}

func NewDescribeVirusScanSettingResponse() (response *DescribeVirusScanSettingResponse) {
	response = &DescribeVirusScanSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询文件查杀设置
func (c *Client) DescribeVirusScanSetting(request *DescribeVirusScanSettingRequest) (response *DescribeVirusScanSettingResponse, err error) {
	if request == nil {
		request = NewDescribeVirusScanSettingRequest()
	}
	response = NewDescribeVirusScanSettingResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecLogJoinObjectsRequest() (request *ModifySecLogJoinObjectsRequest) {
	request = &ModifySecLogJoinObjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogJoinObjects")
	return
}

func NewModifySecLogJoinObjectsResponse() (response *ModifySecLogJoinObjectsResponse) {
	response = &ModifySecLogJoinObjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改安全日志接入对象
func (c *Client) ModifySecLogJoinObjects(request *ModifySecLogJoinObjectsRequest) (response *ModifySecLogJoinObjectsResponse, err error) {
	if request == nil {
		request = NewModifySecLogJoinObjectsRequest()
	}
	response = NewModifySecLogJoinObjectsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnauthorizedCoresTendencyRequest() (request *DescribeUnauthorizedCoresTendencyRequest) {
	request = &DescribeUnauthorizedCoresTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeUnauthorizedCoresTendency")
	return
}

func NewDescribeUnauthorizedCoresTendencyResponse() (response *DescribeUnauthorizedCoresTendencyResponse) {
	response = &DescribeUnauthorizedCoresTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当天未授权核数趋势
func (c *Client) DescribeUnauthorizedCoresTendency(request *DescribeUnauthorizedCoresTendencyRequest) (response *DescribeUnauthorizedCoresTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeUnauthorizedCoresTendencyRequest()
	}
	response = NewDescribeUnauthorizedCoresTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetRequest() (request *ModifyAssetRequest) {
	request = &ModifyAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ModifyAsset")
	return
}

func NewModifyAssetResponse() (response *ModifyAssetResponse) {
	response = &ModifyAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全主机资产刷新
func (c *Client) ModifyAsset(request *ModifyAssetRequest) (response *ModifyAssetResponse, err error) {
	if request == nil {
		request = NewModifyAssetRequest()
	}
	response = NewModifyAssetResponse()
	err = c.Send(request, response)
	return
}

func NewAddEditWarningRulesRequest() (request *AddEditWarningRulesRequest) {
	request = &AddEditWarningRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "AddEditWarningRules")
	return
}

func NewAddEditWarningRulesResponse() (response *AddEditWarningRulesResponse) {
	response = &AddEditWarningRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加编辑告警策略
func (c *Client) AddEditWarningRules(request *AddEditWarningRulesRequest) (response *AddEditWarningRulesResponse, err error) {
	if request == nil {
		request = NewAddEditWarningRulesRequest()
	}
	response = NewAddEditWarningRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageVirusListRequest() (request *DescribeAssetImageVirusListRequest) {
	request = &DescribeAssetImageVirusListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVirusList")
	return
}

func NewDescribeAssetImageVirusListResponse() (response *DescribeAssetImageVirusListResponse) {
	response = &DescribeAssetImageVirusListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全查询镜像病毒列表
func (c *Client) DescribeAssetImageVirusList(request *DescribeAssetImageVirusListRequest) (response *DescribeAssetImageVirusListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageVirusListRequest()
	}
	response = NewDescribeAssetImageVirusListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeRuleInfoRequest() (request *DescribeEscapeRuleInfoRequest) {
	request = &DescribeEscapeRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeRuleInfo")
	return
}

func NewDescribeEscapeRuleInfoResponse() (response *DescribeEscapeRuleInfoResponse) {
	response = &DescribeEscapeRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeEscapeRuleInfo 查询容器逃逸扫描规则信息
func (c *Client) DescribeEscapeRuleInfo(request *DescribeEscapeRuleInfoRequest) (response *DescribeEscapeRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeRuleInfoRequest()
	}
	response = NewDescribeEscapeRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeABTestConfigRequest() (request *DescribeABTestConfigRequest) {
	request = &DescribeABTestConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeABTestConfig")
	return
}

func NewDescribeABTestConfigResponse() (response *DescribeABTestConfigResponse) {
	response = &DescribeABTestConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户当前灰度配置
func (c *Client) DescribeABTestConfig(request *DescribeABTestConfigRequest) (response *DescribeABTestConfigResponse, err error) {
	if request == nil {
		request = NewDescribeABTestConfigRequest()
	}
	response = NewDescribeABTestConfigResponse()
	err = c.Send(request, response)
	return
}

func NewScanComplianceAssetsByPolicyItemRequest() (request *ScanComplianceAssetsByPolicyItemRequest) {
	request = &ScanComplianceAssetsByPolicyItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ScanComplianceAssetsByPolicyItem")
	return
}

func NewScanComplianceAssetsByPolicyItemResponse() (response *ScanComplianceAssetsByPolicyItemResponse) {
	response = &ScanComplianceAssetsByPolicyItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用指定的检测项重新检测选定的资产，返回创建的合规检查任务的ID。
func (c *Client) ScanComplianceAssetsByPolicyItem(request *ScanComplianceAssetsByPolicyItemRequest) (response *ScanComplianceAssetsByPolicyItemResponse, err error) {
	if request == nil {
		request = NewScanComplianceAssetsByPolicyItemRequest()
	}
	response = NewScanComplianceAssetsByPolicyItemResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetImageRegistryScanTaskOneKeyRequest() (request *CreateAssetImageRegistryScanTaskOneKeyRequest) {
	request = &CreateAssetImageRegistryScanTaskOneKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageRegistryScanTaskOneKey")
	return
}

func NewCreateAssetImageRegistryScanTaskOneKeyResponse() (response *CreateAssetImageRegistryScanTaskOneKeyResponse) {
	response = &CreateAssetImageRegistryScanTaskOneKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库创建镜像一键扫描任务
func (c *Client) CreateAssetImageRegistryScanTaskOneKey(request *CreateAssetImageRegistryScanTaskOneKeyRequest) (response *CreateAssetImageRegistryScanTaskOneKeyResponse, err error) {
	if request == nil {
		request = NewCreateAssetImageRegistryScanTaskOneKeyRequest()
	}
	response = NewCreateAssetImageRegistryScanTaskOneKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerAssetSummaryRequest() (request *DescribeContainerAssetSummaryRequest) {
	request = &DescribeContainerAssetSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeContainerAssetSummary")
	return
}

func NewDescribeContainerAssetSummaryResponse() (response *DescribeContainerAssetSummaryResponse) {
	response = &DescribeContainerAssetSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器资产概览信息
func (c *Client) DescribeContainerAssetSummary(request *DescribeContainerAssetSummaryRequest) (response *DescribeContainerAssetSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeContainerAssetSummaryRequest()
	}
	response = NewDescribeContainerAssetSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRiskSyscallEventsRequest() (request *DeleteRiskSyscallEventsRequest) {
	request = &DeleteRiskSyscallEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteRiskSyscallEvents")
	return
}

func NewDeleteRiskSyscallEventsResponse() (response *DeleteRiskSyscallEventsResponse) {
	response = &DeleteRiskSyscallEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运行时高危系统调用事件
func (c *Client) DeleteRiskSyscallEvents(request *DeleteRiskSyscallEventsRequest) (response *DeleteRiskSyscallEventsResponse, err error) {
	if request == nil {
		request = NewDeleteRiskSyscallEventsRequest()
	}
	response = NewDeleteRiskSyscallEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageSimpleListRequest() (request *DescribeImageSimpleListRequest) {
	request = &DescribeImageSimpleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageSimpleList")
	return
}

func NewDescribeImageSimpleListResponse() (response *DescribeImageSimpleListResponse) {
	response = &DescribeImageSimpleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeImageSimpleList 查询全部镜像列表
func (c *Client) DescribeImageSimpleList(request *DescribeImageSimpleListRequest) (response *DescribeImageSimpleListResponse, err error) {
	if request == nil {
		request = NewDescribeImageSimpleListRequest()
	}
	response = NewDescribeImageSimpleListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVirusScanNewTaskRequest() (request *CreateVirusScanNewTaskRequest) {
	request = &CreateVirusScanNewTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateVirusScanNewTask")
	return
}

func NewCreateVirusScanNewTaskResponse() (response *CreateVirusScanNewTaskResponse) {
	response = &CreateVirusScanNewTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运行时文件查杀一键扫描
func (c *Client) CreateVirusScanNewTask(request *CreateVirusScanNewTaskRequest) (response *CreateVirusScanNewTaskResponse, err error) {
	if request == nil {
		request = NewCreateVirusScanNewTaskRequest()
	}
	response = NewCreateVirusScanNewTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulTendencyRequest() (request *DescribeVulTendencyRequest) {
	request = &DescribeVulTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulTendency")
	return
}

func NewDescribeVulTendencyResponse() (response *DescribeVulTendencyResponse) {
	response = &DescribeVulTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地镜像、仓库镜像中严重&高危的漏洞趋势
func (c *Client) DescribeVulTendency(request *DescribeVulTendencyRequest) (response *DescribeVulTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeVulTendencyRequest()
	}
	response = NewDescribeVulTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetProcessListRequest() (request *DescribeAssetProcessListRequest) {
	request = &DescribeAssetProcessListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetProcessList")
	return
}

func NewDescribeAssetProcessListResponse() (response *DescribeAssetProcessListResponse) {
	response = &DescribeAssetProcessListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器安全搜索查询进程列表
func (c *Client) DescribeAssetProcessList(request *DescribeAssetProcessListRequest) (response *DescribeAssetProcessListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetProcessListRequest()
	}
	response = NewDescribeAssetProcessListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSystemVulExportJobRequest() (request *CreateSystemVulExportJobRequest) {
	request = &CreateSystemVulExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateSystemVulExportJob")
	return
}

func NewCreateSystemVulExportJobResponse() (response *CreateSystemVulExportJobResponse) {
	response = &CreateSystemVulExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建系统漏洞导出任务
func (c *Client) CreateSystemVulExportJob(request *CreateSystemVulExportJobRequest) (response *CreateSystemVulExportJobResponse, err error) {
	if request == nil {
		request = NewCreateSystemVulExportJobRequest()
	}
	response = NewCreateSystemVulExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusListRequest() (request *DescribeVirusListRequest) {
	request = &DescribeVirusListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusList")
	return
}

func NewDescribeVirusListResponse() (response *DescribeVirusListResponse) {
	response = &DescribeVirusListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时文件查杀事件列表
func (c *Client) DescribeVirusList(request *DescribeVirusListRequest) (response *DescribeVirusListResponse, err error) {
	if request == nil {
		request = NewDescribeVirusListRequest()
	}
	response = NewDescribeVirusListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeEventsExportRequest() (request *DescribeEscapeEventsExportRequest) {
	request = &DescribeEscapeEventsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventsExport")
	return
}

func NewDescribeEscapeEventsExportResponse() (response *DescribeEscapeEventsExportResponse) {
	response = &DescribeEscapeEventsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeEscapeEventsExport  查询容器逃逸事件列表导出
func (c *Client) DescribeEscapeEventsExport(request *DescribeEscapeEventsExportRequest) (response *DescribeEscapeEventsExportResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeEventsExportRequest()
	}
	response = NewDescribeEscapeEventsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageDenyRuleRequest() (request *DeleteImageDenyRuleRequest) {
	request = &DeleteImageDenyRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DeleteImageDenyRule")
	return
}

func NewDeleteImageDenyRuleResponse() (response *DeleteImageDenyRuleResponse) {
	response = &DeleteImageDenyRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除镜像拦截规则
func (c *Client) DeleteImageDenyRule(request *DeleteImageDenyRuleRequest) (response *DeleteImageDenyRuleResponse, err error) {
	if request == nil {
		request = NewDeleteImageDenyRuleRequest()
	}
	response = NewDeleteImageDenyRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusMonitorSettingRequest() (request *DescribeVirusMonitorSettingRequest) {
	request = &DescribeVirusMonitorSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusMonitorSetting")
	return
}

func NewDescribeVirusMonitorSettingResponse() (response *DescribeVirusMonitorSettingResponse) {
	response = &DescribeVirusMonitorSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询文件查杀实时监控设置
func (c *Client) DescribeVirusMonitorSetting(request *DescribeVirusMonitorSettingRequest) (response *DescribeVirusMonitorSettingResponse, err error) {
	if request == nil {
		request = NewDescribeVirusMonitorSettingRequest()
	}
	response = NewDescribeVirusMonitorSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCompliancePolicyItemAffectedSummaryRequest() (request *DescribeCompliancePolicyItemAffectedSummaryRequest) {
	request = &DescribeCompliancePolicyItemAffectedSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeCompliancePolicyItemAffectedSummary")
	return
}

func NewDescribeCompliancePolicyItemAffectedSummaryResponse() (response *DescribeCompliancePolicyItemAffectedSummaryResponse) {
	response = &DescribeCompliancePolicyItemAffectedSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照 检测项 → 资产 的两级层次展开的第一层级：检测项层级。
func (c *Client) DescribeCompliancePolicyItemAffectedSummary(request *DescribeCompliancePolicyItemAffectedSummaryRequest) (response *DescribeCompliancePolicyItemAffectedSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeCompliancePolicyItemAffectedSummaryRequest()
	}
	response = NewDescribeCompliancePolicyItemAffectedSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportJobDownloadURLRequest() (request *DescribeExportJobDownloadURLRequest) {
	request = &DescribeExportJobDownloadURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeExportJobDownloadURL")
	return
}

func NewDescribeExportJobDownloadURLResponse() (response *DescribeExportJobDownloadURLResponse) {
	response = &DescribeExportJobDownloadURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询导出任务下载URL
func (c *Client) DescribeExportJobDownloadURL(request *DescribeExportJobDownloadURLRequest) (response *DescribeExportJobDownloadURLResponse, err error) {
	if request == nil {
		request = NewDescribeExportJobDownloadURLRequest()
	}
	response = NewDescribeExportJobDownloadURLResponse()
	err = c.Send(request, response)
	return
}

func NewScanComplianceScanFailedAssetsRequest() (request *ScanComplianceScanFailedAssetsRequest) {
	request = &ScanComplianceScanFailedAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "ScanComplianceScanFailedAssets")
	return
}

func NewScanComplianceScanFailedAssetsResponse() (response *ScanComplianceScanFailedAssetsResponse) {
	response = &ScanComplianceScanFailedAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新检测选定的检测失败的资产下的所有失败的检测项，返回创建的合规检查任务的ID。
func (c *Client) ScanComplianceScanFailedAssets(request *ScanComplianceScanFailedAssetsRequest) (response *ScanComplianceScanFailedAssetsResponse, err error) {
	if request == nil {
		request = NewScanComplianceScanFailedAssetsRequest()
	}
	response = NewScanComplianceScanFailedAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterCheckTimerSettingRequest() (request *DescribeClusterCheckTimerSettingRequest) {
	request = &DescribeClusterCheckTimerSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterCheckTimerSetting")
	return
}

func NewDescribeClusterCheckTimerSettingResponse() (response *DescribeClusterCheckTimerSettingResponse) {
	response = &DescribeClusterCheckTimerSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群定时检查设置
func (c *Client) DescribeClusterCheckTimerSetting(request *DescribeClusterCheckTimerSettingRequest) (response *DescribeClusterCheckTimerSettingResponse, err error) {
	if request == nil {
		request = NewDescribeClusterCheckTimerSettingRequest()
	}
	response = NewDescribeClusterCheckTimerSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpgradeAgentListRequest() (request *DescribeUpgradeAgentListRequest) {
	request = &DescribeUpgradeAgentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeUpgradeAgentList")
	return
}

func NewDescribeUpgradeAgentListResponse() (response *DescribeUpgradeAgentListResponse) {
	response = &DescribeUpgradeAgentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询升级终端列表详情
func (c *Client) DescribeUpgradeAgentList(request *DescribeUpgradeAgentListRequest) (response *DescribeUpgradeAgentListResponse, err error) {
	if request == nil {
		request = NewDescribeUpgradeAgentListRequest()
	}
	response = NewDescribeUpgradeAgentListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessRuleDetailRequest() (request *DescribeAbnormalProcessRuleDetailRequest) {
	request = &DescribeAbnormalProcessRuleDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessRuleDetail")
	return
}

func NewDescribeAbnormalProcessRuleDetailResponse() (response *DescribeAbnormalProcessRuleDetailResponse) {
	response = &DescribeAbnormalProcessRuleDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时异常策略详细信息
func (c *Client) DescribeAbnormalProcessRuleDetail(request *DescribeAbnormalProcessRuleDetailRequest) (response *DescribeAbnormalProcessRuleDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessRuleDetailRequest()
	}
	response = NewDescribeAbnormalProcessRuleDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusEventTendencyRequest() (request *DescribeVirusEventTendencyRequest) {
	request = &DescribeVirusEventTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusEventTendency")
	return
}

func NewDescribeVirusEventTendencyResponse() (response *DescribeVirusEventTendencyResponse) {
	response = &DescribeVirusEventTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马事件趋势
func (c *Client) DescribeVirusEventTendency(request *DescribeVirusEventTendencyRequest) (response *DescribeVirusEventTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeVirusEventTendencyRequest()
	}
	response = NewDescribeVirusEventTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsSummaryRequest() (request *DescribeRiskDnsSummaryRequest) {
	request = &DescribeRiskDnsSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskDnsSummary")
	return
}

func NewDescribeRiskDnsSummaryResponse() (response *DescribeRiskDnsSummaryResponse) {
	response = &DescribeRiskDnsSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意请求事件统计
func (c *Client) DescribeRiskDnsSummary(request *DescribeRiskDnsSummaryRequest) (response *DescribeRiskDnsSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsSummaryRequest()
	}
	response = NewDescribeRiskDnsSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrModifyPostPayCoresRequest() (request *CreateOrModifyPostPayCoresRequest) {
	request = &CreateOrModifyPostPayCoresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcss", APIVersion, "CreateOrModifyPostPayCores")
	return
}

func NewCreateOrModifyPostPayCoresResponse() (response *CreateOrModifyPostPayCoresResponse) {
	response = &CreateOrModifyPostPayCoresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateOrModifyPostPayCores  创建或者编辑弹性计费上限
func (c *Client) CreateOrModifyPostPayCores(request *CreateOrModifyPostPayCoresRequest) (response *CreateOrModifyPostPayCoresResponse, err error) {
	if request == nil {
		request = NewCreateOrModifyPostPayCoresRequest()
	}
	response = NewCreateOrModifyPostPayCoresResponse()
	err = c.Send(request, response)
	return
}
