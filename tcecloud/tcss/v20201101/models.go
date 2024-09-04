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
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeNetworkFirewallPolicyDiscoverResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态，可能为：Task_Running,Task_Succ,Task_Error,Task_NoExist

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallPolicyDiscoverResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPolicyDiscoverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEditAbnormalProcessRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddEditAbnormalProcessRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditAbnormalProcessRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePurchaseStateInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：可申请试用可购买；1：只可购买(含试用审核不通过和试用过期)；2：试用生效中；3：专业版生效中；4：专业版过期

		State *int64 `json:"State,omitempty" name:"State"`
		// 总核数

		CoresCnt *uint64 `json:"CoresCnt,omitempty" name:"CoresCnt"`
		// 已购买核数

		AuthorizedCoresCnt *uint64 `json:"AuthorizedCoresCnt,omitempty" name:"AuthorizedCoresCnt"`
		// 镜像数

		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
		// 已授权镜像数

		AuthorizedImageCnt *uint64 `json:"AuthorizedImageCnt,omitempty" name:"AuthorizedImageCnt"`
		// 已购买镜像授权数

		PurchasedAuthorizedCnt *uint64 `json:"PurchasedAuthorizedCnt,omitempty" name:"PurchasedAuthorizedCnt"`
		// 过期时间

		ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`
		// 0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费(用户设置)

		AutomaticRenewal *int64 `json:"AutomaticRenewal,omitempty" name:"AutomaticRenewal"`
		// 试用期间赠送镜像授权数，可能会过期

		GivenAuthorizedCnt *uint64 `json:"GivenAuthorizedCnt,omitempty" name:"GivenAuthorizedCnt"`
		// 起始时间

		BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
		// 子状态(具体意义依据State字段而定)
		// State为4时，有效值为: ISOLATE(隔离) DESTROED(已销毁)

		SubState *string `json:"SubState,omitempty" name:"SubState"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePurchaseStateInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePurchaseStateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVirusListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryVirusListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVirusListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlEventsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAccessControlEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageDenyRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageDenyRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageDenyRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAndPublishNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAndPublishNetworkFirewallPolicyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAndPublishNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterCheckTimerSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClusterCheckTimerSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterCheckTimerSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulIgnoreLocalImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像列表

		List []*VulIgnoreLocalImage `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulIgnoreLocalImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulIgnoreLocalImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 入站规则
	//
	// 全部允许：1
	//
	// 全部拒绝 ：2
	//
	// 自定义：3

	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`
	// 出站规则
	//
	// 全部允许：1
	//
	// 全部拒绝 ：2
	//
	// 自定义：3

	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`
	// pod选择器

	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 自定义规则

	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

func (r *AddNetworkFirewallPolicyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckNetworkFirewallPolicyYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckNetworkFirewallPolicyYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckNetworkFirewallPolicyYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSimpleListRequest struct {
	*tchttp.BaseRequest

	// IsAuthorized 是否已经授权, 0:否 1:是 无:全部

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeImageSimpleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageSimpleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallEventsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Excel下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 任务Id

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallEventsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetListRequest struct {
	*tchttp.BaseRequest

	// 资产类型列表。

	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的数据量，默认为10，最大为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询过滤器

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceAssetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterDetailExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要导出的集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateClusterDetailExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterDetailExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHybridClustersRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeHybridClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHybridClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRiskSyscallEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRiskSyscallEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRiskSyscallEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefenceEvent struct {

	// 漏洞CVEID

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 入侵状态

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 攻击源IP

	SourceIP *string `json:"SourceIP,omitempty" name:"SourceIP"`
	// 攻击源ip地址所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 容器ID

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 处理状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件ID

	EventID *int64 `json:"EventID,omitempty" name:"EventID"`
	// 首次发现时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 隔离状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 最近发现时间

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// 	"NODE_DESTROYED"      //节点已销毁
	// 	"CONTAINER_EXITED"    //容器已退出
	// 	"CONTAINER_DESTROYED" //容器已销毁
	// 	"SHARED_HOST"         // 容器与主机共享网络
	// 	"RESOURCE_LIMIT"      //隔离操作资源超限
	// 	"UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 主机QUUID/超级节点ID

	QUUID *string `json:"QUUID,omitempty" name:"QUUID"`
	// 主机内网IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机名称/超级节点名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 节点类型[NORMAL:普通节点|SUPER:超级节点]

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 超级节点唯一ID

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 超级节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type DescribeAssetImageRegistryVulListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryVulListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVulListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageComponent struct {

	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 组件路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 组件类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 组件漏洞数量

	VulCount *uint64 `json:"VulCount,omitempty" name:"VulCount"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

type ImagesVul struct {

	// 漏洞id

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 分类2

	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
	// 风险等级

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 描述

	Des *string `json:"Des,omitempty" name:"Des"`
	// 解决方案

	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`
	// 引用

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 防御方案

	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`
	// 提交时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// CVSS V3分数

	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`
	// CVSS V3描述

	CVSSV3Desc *string `json:"CVSSV3Desc,omitempty" name:"CVSSV3Desc"`
	// 是否是重点关注：true：是，false：不是

	IsSuggest *bool `json:"IsSuggest,omitempty" name:"IsSuggest"`
	// 修复版本号

	FixedVersions *string `json:"FixedVersions,omitempty" name:"FixedVersions"`
	// 漏洞标签:"CanBeFixed","DynamicLevelPoc","DynamicLevelExp"

	Tag []*string `json:"Tag,omitempty" name:"Tag"`
	// 攻击热度

	AttackLevel *int64 `json:"AttackLevel,omitempty" name:"AttackLevel"`
}

type RiskSyscallEventDescription struct {

	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 系统调用名称

	SyscallName *string `json:"SyscallName,omitempty" name:"SyscallName"`
	// 事件最后一次处理的时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type CreateVulExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式desc ，asc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateVulExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscapeEventDescription struct {

	// 事件规则

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 事件最后一次处理的时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type DescribeAssetImageRegistryVirusListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像信息

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVirusListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVirusListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousConnectionRuleSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 黑名单内容数量

		BlackListCount *uint64 `json:"BlackListCount,omitempty" name:"BlackListCount"`
		// 黑名单域名数量

		BlackListDomainCount *uint64 `json:"BlackListDomainCount,omitempty" name:"BlackListDomainCount"`
		// 黑名单IP数量

		BlackListIPCount *uint64 `json:"BlackListIPCount,omitempty" name:"BlackListIPCount"`
		// 白名单内容数量

		WhiteListCount *uint64 `json:"WhiteListCount,omitempty" name:"WhiteListCount"`
		// 白单域名数量

		WhiteListDomainCount *uint64 `json:"WhiteListDomainCount,omitempty" name:"WhiteListDomainCount"`
		// 白名单IP数量

		WhiteListIPCount *uint64 `json:"WhiteListIPCount,omitempty" name:"WhiteListIPCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaliciousConnectionRuleSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionRuleSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetImageRegistryScanStopOneKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetImageRegistryScanStopOneKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetImageRegistryScanStopOneKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表

		List []*K8sApiAbnormalEventListItem `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogDeliveryClsOptionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// cls可选日志集合列表(仅当入参ClsRegion不为空时返回)

		LogSetList []*ClsLogsetInfo `json:"LogSetList,omitempty" name:"LogSetList"`
		// 可选地域列表(仅当入参ClsRegion为空时返回)

		RegionList []*RegionInfo `json:"RegionList,omitempty" name:"RegionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogDeliveryClsOptionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogDeliveryClsOptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDBServiceListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetDBServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDBServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SuperNodePodListItem struct {

	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// podIP

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// cpu需求核数

	CpuRequest *int64 `json:"CpuRequest,omitempty" name:"CpuRequest"`
	// cpu限制核数

	CpuLimit *int64 `json:"CpuLimit,omitempty" name:"CpuLimit"`
	// 内存需求大小

	MemRequest *int64 `json:"MemRequest,omitempty" name:"MemRequest"`
	// 内存限制大小

	MemLimit *int64 `json:"MemLimit,omitempty" name:"MemLimit"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 工作负载名称

	DeploymentName *string `json:"DeploymentName,omitempty" name:"DeploymentName"`
	// 工作负载id

	DeploymentID *string `json:"DeploymentID,omitempty" name:"DeploymentID"`
	// 启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 关联容器个数

	RelateContainerCount *uint64 `json:"RelateContainerCount,omitempty" name:"RelateContainerCount"`
	// 运行时间

	RunningTime *string `json:"RunningTime,omitempty" name:"RunningTime"`
	// PodUid

	PodUid *string `json:"PodUid,omitempty" name:"PodUid"`
	// 计费核数

	ChargeCoresCnt *uint64 `json:"ChargeCoresCnt,omitempty" name:"ChargeCoresCnt"`
	// 防护状态

	DefendStatus *string `json:"DefendStatus,omitempty" name:"DefendStatus"`
}

type ResetSecLogTopicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置类型(ckafka/cls)

	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
	// 日志类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

func (r *ResetSecLogTopicConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetSecLogTopicConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNetworkFirewallPolicyYamlDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNetworkFirewallPolicyYamlDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNetworkFirewallPolicyYamlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyRuleDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 规则名称

		RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
		// 规则类型 RULE_RISK：风险， RULE_PRIVILEGE：特权

		RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
		// 生效的镜像数量

		EffectImageCount *int64 `json:"EffectImageCount,omitempty" name:"EffectImageCount"`
		// 是否对全部扫描镜像生效。0:全选镜像，1:自选镜像

		IsEffectAllImage *int64 `json:"IsEffectAllImage,omitempty" name:"IsEffectAllImage"`
		// 规则开始生效时间

		EffectTime *string `json:"EffectTime,omitempty" name:"EffectTime"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 操作用户

		OperationUin *string `json:"OperationUin,omitempty" name:"OperationUin"`
		// 生效状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中

		EffectStatus *string `json:"EffectStatus,omitempty" name:"EffectStatus"`
		// 规则描述

		RuleDescription *string `json:"RuleDescription,omitempty" name:"RuleDescription"`
		// 启用状态 0:开启，1:关闭

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 漏洞，0:未选中，1:选中

		Vul *int64 `json:"Vul,omitempty" name:"Vul"`
		// cve编号

		CVEIDSet []*string `json:"CVEIDSet,omitempty" name:"CVEIDSet"`
		// 组件编号

		ComponentSet []*string `json:"ComponentSet,omitempty" name:"ComponentSet"`
		// 漏洞分类

		VulClassSet []*string `json:"VulClassSet,omitempty" name:"VulClassSet"`
		// 漏洞等级

		VulLevelSet []*string `json:"VulLevelSet,omitempty" name:"VulLevelSet"`
		// 漏洞标签

		VulLabelSet []*string `json:"VulLabelSet,omitempty" name:"VulLabelSet"`
		// 木马，0:未选中，1:选中

		Virus *int64 `json:"Virus,omitempty" name:"Virus"`
		// 木马md5列表

		VirusMD5Set []*string `json:"VirusMD5Set,omitempty" name:"VirusMD5Set"`
		// 木马等级

		VirusLevelSet []*string `json:"VirusLevelSet,omitempty" name:"VirusLevelSet"`
		// 病毒名

		VirusName []*string `json:"VirusName,omitempty" name:"VirusName"`
		// 敏感信息，0:未选中，1:选中

		Risk *int64 `json:"Risk,omitempty" name:"Risk"`
		// 敏感等级

		RiskLevelSet []*string `json:"RiskLevelSet,omitempty" name:"RiskLevelSet"`
		// 敏感信息分类

		RiskType []*string `json:"RiskType,omitempty" name:"RiskType"`
		// 特权启动 0:不允许，1:允许

		PrivilegeRun *int64 `json:"PrivilegeRun,omitempty" name:"PrivilegeRun"`
		// 特权类型,

		PrivilegeDetail []*string `json:"PrivilegeDetail,omitempty" name:"PrivilegeDetail"`
		// 镜像ID列表

		EffectImageSet []*string `json:"EffectImageSet,omitempty" name:"EffectImageSet"`
		// 多少天后生效

		EffectDay *uint64 `json:"EffectDay,omitempty" name:"EffectDay"`
		// 规则RuelD

		RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyRuleDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRuleDetailRequest struct {
	*tchttp.BaseRequest

	// 策略唯一id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 镜像id, 仅仅在事件加白的时候使用

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAccessControlRuleDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRuleDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEditRiskSyscallWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddEditRiskSyscallWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditRiskSyscallWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryImageLayerDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像层列表

		List []*ImageLayer `json:"List,omitempty" name:"List"`
		// 个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryImageLayerDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryImageLayerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrModifyMaliciousConnectionBlackListRequest struct {
	*tchttp.BaseRequest

	// 枚举
	// IP: IP
	// 域名：DOMAIN

	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`
	// 黑名单域名

	BlackDomainList []*string `json:"BlackDomainList,omitempty" name:"BlackDomainList"`
	// 黑名单IP

	BlackIPList []*string `json:"BlackIPList,omitempty" name:"BlackIPList"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 黑名单记录id，只有修改时需要

	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *AddOrModifyMaliciousConnectionBlackListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrModifyMaliciousConnectionBlackListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyRuleSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像拦截规则总数(含关闭的和开启的)

		RuleTotalCount *uint64 `json:"RuleTotalCount,omitempty" name:"RuleTotalCount"`
		// 开启的镜像拦截规则数

		EnabledRuleCount *uint64 `json:"EnabledRuleCount,omitempty" name:"EnabledRuleCount"`
		// 观察期中的镜像拦截规则数

		ObservedRuleCount *uint64 `json:"ObservedRuleCount,omitempty" name:"ObservedRuleCount"`
		// 已生效的镜像拦截规则数

		EffectiveRuleCount *uint64 `json:"EffectiveRuleCount,omitempty" name:"EffectiveRuleCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyRuleSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageRegistryAutoAuthorizedStatusRequest struct {
	*tchttp.BaseRequest

	// 0关闭 1打开

	IsEnabled *uint64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 规则ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ModifyImageRegistryAutoAuthorizedStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageRegistryAutoAuthorizedStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPodListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Pod列表详细信息

		PodList []*ClusterPodInfo `json:"PodList,omitempty" name:"PodList"`
		// Pod列表总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserPodListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPodListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryScanStatusOneKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像个数

		ImageTotal *uint64 `json:"ImageTotal,omitempty" name:"ImageTotal"`
		// 扫描镜像个数

		ImageScanCnt *uint64 `json:"ImageScanCnt,omitempty" name:"ImageScanCnt"`
		// 扫描进度列表

		ImageStatus []*ImageProgress `json:"ImageStatus,omitempty" name:"ImageStatus"`
		// 安全个数

		SuccessCount *uint64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 风险个数

		RiskCount *uint64 `json:"RiskCount,omitempty" name:"RiskCount"`
		// 总的扫描进度

		Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`
		// 总的扫描状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 扫描剩余时间

		ScanRemainTime *uint64 `json:"ScanRemainTime,omitempty" name:"ScanRemainTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryScanStatusOneKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryScanStatusOneKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClusterSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEscapeRuleRequest struct {
	*tchttp.BaseRequest

	// 需要修改的数组

	RuleSet []*EscapeRuleEnabled `json:"RuleSet,omitempty" name:"RuleSet"`
}

func (r *ModifyEscapeRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEscapeRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 模板列表

		List []*SearchTemplate `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSearchTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FocusVulnerabilityRuleItem struct {

	// CVSS V3评分

	CvssScore *float64 `json:"CvssScore,omitempty" name:"CvssScore"`
	// 漏洞标签

	Tags []*VulnerabilityFocusRuleTags `json:"Tags,omitempty" name:"Tags"`
}

type DescribeVulImageSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulImageSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulImageSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessControlsRuleExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By []*string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateAccessControlsRuleExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlsRuleExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddClusterWorkloadToWhiteListRequest struct {
	*tchttp.BaseRequest

	// 受影响的Workload信息

	AffectedWorkloadInfo *AffectedWorkloadItem `json:"AffectedWorkloadInfo,omitempty" name:"AffectedWorkloadInfo"`
	// 检测项ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
}

func (r *AddClusterWorkloadToWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterWorkloadToWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSearchTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSearchTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSearchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSystemVulExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSystemVulExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSystemVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEscapeWhiteListExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEscapeWhiteListExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEscapeWhiteListExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImagesBindRuleSimpleInfo struct {

	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 绑定规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type DescribeK8sApiAbnormalEventListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>MatchRules - string  - 是否必填: 否 -命中规则筛选</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>
	// <li>Status - string  - 是否必填: 否 -事件状态筛选</li>
	// <li>MatchRuleType - string  - 是否必填: 否 -命中规则类型筛选</li>
	// <li>ClusterRunningStatus - string  - 是否必填: 否 -集群运行状态</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段
	// LatestFoundTime: 最近生成时间
	// AlarmCount: 告警数量

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeK8sApiAbnormalEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportDefenceVulRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：披露时间：SubmitTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeSupportDefenceVulRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSupportDefenceVulRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVirusListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>FileName - String - 是否必填：否 - 文件名称</li>
	// <li>FilePath - String - 是否必填：否 - 文件路径</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名称</li>
	// <li>ContainerName- String - 是否必填：是 - 容器名称</li>
	// <li>ContainerId- string - 是否必填：否 - 容器id</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称</li>
	// <li>ImageId- string - 是否必填：否 - 镜像id</li>
	// <li>IsRealTime- int - 是否必填：否 - 过滤是否实时监控数据</li>
	// <li>TaskId- string - 是否必填：否 - 任务ID</li>
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>ContainerNetStatus - String -是否必填: 否 -  容器网络状态筛选 NORMAL ISOLATED ISOLATING RESTORING RESTORE_FAILED</li>
	// <li>ContainerStatus - string -是否必填: 否 - 容器状态 RUNNING PAUSED STOPPED CREATED DESTROYED RESTARTING REMOVING</li>
	// <li>AutoIsolateMode - string -是否必填: 否 - 隔离方式 MANUAL AUTO</li>
	// <li>MD5 - string -是否必填: 否 - md5 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *ExportVirusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVirusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像漏洞列表

		List []*ImagesVul `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceScanFailedAssetListRequest struct {
	*tchttp.BaseRequest

	// 资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产

	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的数据量，默认为10，最大为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询过滤器

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceScanFailedAssetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceScanFailedAssetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpgradeVersionInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 终端最新版本

		LatestVersion *string `json:"LatestVersion,omitempty" name:"LatestVersion"`
		// 低版本终端数量

		LowVersionCnt *uint64 `json:"LowVersionCnt,omitempty" name:"LowVersionCnt"`
		// 低版本列表

		LowVersionInfo *string `json:"LowVersionInfo,omitempty" name:"LowVersionInfo"`
		// 自动升级开关

		AutoSwitch *bool `json:"AutoSwitch,omitempty" name:"AutoSwitch"`
		// 最近一次升级终端数量

		LastUpgradeCnt *uint64 `json:"LastUpgradeCnt,omitempty" name:"LastUpgradeCnt"`
		// 最近一次升级时间

		LastUpgradeTime *string `json:"LastUpgradeTime,omitempty" name:"LastUpgradeTime"`
		// 升级状态：正在升级：UPGRADING， 完成：FINISHED

		Status *string `json:"Status,omitempty" name:"Status"`
		// 最新md5

		Md5 *string `json:"Md5,omitempty" name:"Md5"`
		// cos临时ID

		TmpSecretID *string `json:"TmpSecretID,omitempty" name:"TmpSecretID"`
		// cos临时key

		TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
		// cos临时token

		SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
		// 开始时间

		StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
		// 到期时间

		ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// 桶名

		BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
		// 地域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpgradeVersionInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpgradeVersionInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEditAccessControlRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddEditAccessControlRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitializeUserComplianceEnvironmentRequest struct {
	*tchttp.BaseRequest
}

func (r *InitializeUserComplianceEnvironmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitializeUserComplianceEnvironmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousConnectionWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求白名单总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 恶意请求白名单列表

		List []*MaliciousConnectionRuleInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaliciousConnectionWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoAuthorizedRuleHostInfo struct {

	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip即内网ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 镜像个数

	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
	// 容器个数

	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 主机实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// docker 版本

	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`
	// agent运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type K8sApiAbnormalTendencyItem struct {

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 异常UA请求事件数

	ExceptionUARequestCount *uint64 `json:"ExceptionUARequestCount,omitempty" name:"ExceptionUARequestCount"`
	// 匿名用户权限事件数

	AnonymousUserRightCount *uint64 `json:"AnonymousUserRightCount,omitempty" name:"AnonymousUserRightCount"`
	// 凭据信息获取事件数

	CredentialInformationObtainCount *uint64 `json:"CredentialInformationObtainCount,omitempty" name:"CredentialInformationObtainCount"`
	// 敏感数据挂载事件数

	SensitiveDataMountCount *uint64 `json:"SensitiveDataMountCount,omitempty" name:"SensitiveDataMountCount"`
	// 命令执行事件数

	CmdExecCount *uint64 `json:"CmdExecCount,omitempty" name:"CmdExecCount"`
	// 异常定时任务事件数

	AbnormalScheduledTaskCount *uint64 `json:"AbnormalScheduledTaskCount,omitempty" name:"AbnormalScheduledTaskCount"`
	// 静态Pod创建数

	StaticsPodCreateCount *uint64 `json:"StaticsPodCreateCount,omitempty" name:"StaticsPodCreateCount"`
	// 可疑容器创建数

	DoubtfulContainerCreateCount *uint64 `json:"DoubtfulContainerCreateCount,omitempty" name:"DoubtfulContainerCreateCount"`
	// 自定义规则事件数

	UserDefinedRuleCount *uint64 `json:"UserDefinedRuleCount,omitempty" name:"UserDefinedRuleCount"`
	// 匿名访问事件数

	AnonymousAccessCount *uint64 `json:"AnonymousAccessCount,omitempty" name:"AnonymousAccessCount"`
	// 特权容器事件数

	PrivilegeContainerCount *uint64 `json:"PrivilegeContainerCount,omitempty" name:"PrivilegeContainerCount"`
}

type ModifyReverseShellSystemPolicyConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReverseShellSystemPolicyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReverseShellSystemPolicyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTcssSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTcssSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTcssSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulTendencyRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 枚举类型：
	// LATEST：最新版本
	// CONTAINER: 运行容器

	SphereOfInfluence *string `json:"SphereOfInfluence,omitempty" name:"SphereOfInfluence"`
}

func (r *DescribeVulTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNetworkFirewallPolicyYamlDetailRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// base64编码的networkpolicy yaml字符串

	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateNetworkFirewallPolicyYamlDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNetworkFirewallPolicyYamlDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AffectedNodeItem struct {

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 内网ip地址

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 节点的角色，Master、Work等

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// k8s版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 运行时组件,docker或者containerd

	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 检查结果的验证信息

	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type DescribeABTestConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 灰度项目配置

		Config []*ABTestConfig `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeABTestConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeABTestConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncComplianceHostAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机同步状态， NORMAL，RUNNING

		SyncStatus *string `json:"SyncStatus,omitempty" name:"SyncStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncComplianceHostAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncComplianceHostAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkFirewallPublishResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkFirewallPublishResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkFirewallPublishResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefencePlugin struct {

	// java进程pid

	PID *int64 `json:"PID,omitempty" name:"PID"`
	// 进程主类名

	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`
	// 插件运行状态：注入中:INJECTING，注入成功：SUCCESS，注入失败：FAIL，插件超时：TIMEOUT，插件退出：QUIT

	Status *string `json:"Status,omitempty" name:"Status"`
	// 错误日志

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
}

type DescribeComplianceWhitelistItemListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白名单项的列表。

		WhitelistItemSet []*ComplianceWhitelistItem `json:"WhitelistItemSet,omitempty" name:"WhitelistItemSet"`
		// 白名单项的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceWhitelistItemListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceWhitelistItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIgnoreVulResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddIgnoreVulResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddIgnoreVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProVersionInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专业版开始时间，补充购买时才不为空

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 专业版结束时间，补充购买时才不为空

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 需购买的机器核数

		CoresCnt *uint64 `json:"CoresCnt,omitempty" name:"CoresCnt"`
		// 弹性计费上限

		MaxPostPayCoresCnt *uint64 `json:"MaxPostPayCoresCnt,omitempty" name:"MaxPostPayCoresCnt"`
		// 资源ID

		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
		// 购买状态
		// 待购: Pending
		// 已购: Normal
		// 隔离: Isolate

		BuyStatus *string `json:"BuyStatus,omitempty" name:"BuyStatus"`
		// 是否曾经购买过(false:未曾 true:曾经购买过)

		IsPurchased *bool `json:"IsPurchased,omitempty" name:"IsPurchased"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProVersionInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProVersionInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEscapeRegexpRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 启用状态

	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyEscapeRegexpRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEscapeRegexpRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanIgnoreVul struct {

	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
	// 漏洞CVEID

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 忽略的仓库镜像数

	RegistryImageCount *int64 `json:"RegistryImageCount,omitempty" name:"RegistryImageCount"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否忽略所有镜像：0：否/1：是

	IsIgnoreAll *int64 `json:"IsIgnoreAll,omitempty" name:"IsIgnoreAll"`
	// 忽略的本地镜像数

	LocalImageCount *int64 `json:"LocalImageCount,omitempty" name:"LocalImageCount"`
}

type ModifyVirusScanTimeoutSettingRequest struct {
	*tchttp.BaseRequest

	// 超时时长单位小时(5~24h)

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 设置类型0一键检测，1定时检测

	ScanType *uint64 `json:"ScanType,omitempty" name:"ScanType"`
}

func (r *ModifyVirusScanTimeoutSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusScanTimeoutSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyTrialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 申请结果，0成功，其他失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 试用失败原因，成功可为空；eg：请勿重复申请

		Reason *string `json:"Reason,omitempty" name:"Reason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyTrialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyTrialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessChildRuleInfo struct {

	// 子策略id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 策略模式，   RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截

	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 威胁等级，HIGH:高，MIDDLE:中，LOW:低

	RuleLevel *string `json:"RuleLevel,omitempty" name:"RuleLevel"`
}

type ContainerNetwork struct {

	// endpoint id

	EndpointID *string `json:"EndpointID,omitempty" name:"EndpointID"`
	// 模式:bridge

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 网络名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网络ID

	NetworkID *string `json:"NetworkID,omitempty" name:"NetworkID"`
	// 网关

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
	// IPV4地址

	Ipv4 *string `json:"Ipv4,omitempty" name:"Ipv4"`
	// IPV6地址

	Ipv6 *string `json:"Ipv6,omitempty" name:"Ipv6"`
	// MAC 地址

	MAC *string `json:"MAC,omitempty" name:"MAC"`
}

type ModifyAbnormalProcessRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAbnormalProcessRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAbnormalProcessRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSuperNodePodListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*SuperNodePodListItem `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSuperNodePodListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSuperNodePodListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryTimingScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 定时扫描开关

		Enable *bool `json:"Enable,omitempty" name:"Enable"`
		// 定时任务扫描时间

		ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
		// 定时扫描间隔

		ScanPeriod *uint64 `json:"ScanPeriod,omitempty" name:"ScanPeriod"`
		// 扫描类型数组

		ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`
		// 扫描全部镜像

		All *bool `json:"All,omitempty" name:"All"`
		// 自定义扫描镜像

		Images []*ImageInfo `json:"Images,omitempty" name:"Images"`
		// 自动以扫描镜像Id

		Id []*uint64 `json:"Id,omitempty" name:"Id"`
		// 是否扫描最新版本镜像

		Latest *bool `json:"Latest,omitempty" name:"Latest"`
		// 扫描结束时间

		ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`
		// 仓库类型 tcr,ccr,harbor

		RegistryType []*string `json:"RegistryType,omitempty" name:"RegistryType"`
		// 是否存在运行中的容器

		ContainerRunning *bool `json:"ContainerRunning,omitempty" name:"ContainerRunning"`
		// 扫描范围 0全部镜像，1自选镜像，2推荐扫描镜像

		ScanScope *uint64 `json:"ScanScope,omitempty" name:"ScanScope"`
		// 命名空间

		Namespace []*string `json:"Namespace,omitempty" name:"Namespace"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRegistryTimingScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryTimingScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulDefenceHostExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，正常：SUCCESS，异常：FAIL， NO_DEFENCE:未防御</li>
	// <li>KeyWords- string - 是否必填：否 - 主机名称/IP。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，最大值为100000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：更新时间：ModifyTime/首次开启时间：CreateTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateVulDefenceHostExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulDefenceHostExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkFirewallClusterRefreshRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateNetworkFirewallClusterRefreshRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkFirewallClusterRefreshRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogJoinObjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecLogJoinObjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogJoinObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRegistryTypeCount struct {

	// 仓库类型

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 镜像数量

	ImageCount *int64 `json:"ImageCount,omitempty" name:"ImageCount"`
}

type DescribeEscapeEventTendencyRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeEscapeEventTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageScanStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像个数

		ImageTotal *uint64 `json:"ImageTotal,omitempty" name:"ImageTotal"`
		// 扫描镜像个数

		ImageScanCnt *uint64 `json:"ImageScanCnt,omitempty" name:"ImageScanCnt"`
		// 扫描状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 扫描进度 ImageScanCnt/ImageTotal *100

		Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`
		// 安全个数

		SuccessCount *uint64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 风险个数

		RiskCount *uint64 `json:"RiskCount,omitempty" name:"RiskCount"`
		// 剩余扫描时间

		LeftSeconds *uint64 `json:"LeftSeconds,omitempty" name:"LeftSeconds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageScanStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageScanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskTargetCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRiskTargetCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskTargetCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVirusScanNewTaskRequest struct {
	*tchttp.BaseRequest

	// 自选扫描范围的容器id或者节点id

	ScanIDs []*ScanRangeInfo `json:"ScanIDs,omitempty" name:"ScanIDs"`
	// SCAN_NODE:扫描节点 SCAN_CONTAINER:扫描容器

	ScanRangeType *string `json:"ScanRangeType,omitempty" name:"ScanRangeType"`
	// 超时时长，单位小时

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 自选排除或扫描的地址

	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
	// 扫描路径模式：
	// SCAN_PATH_ALL：全部路径
	// SCAN_PATH_DEFAULT：默认路径
	// SCAN_PATH_USER_DEFINE：用户自定义路径
	//

	ScanPathMode *string `json:"ScanPathMode,omitempty" name:"ScanPathMode"`
	// true:包含路径 false:排除路径

	IsIncludePath *bool `json:"IsIncludePath,omitempty" name:"IsIncludePath"`
}

func (r *CreateVirusScanNewTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVirusScanNewTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSuperNodeMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSuperNodeMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSuperNodeMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 木马列表

		List []*VirusInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启定期扫描

		EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`
		// 检测周期每隔多少天

		Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`
		// 扫描开始时间

		BeginScanAt *string `json:"BeginScanAt,omitempty" name:"BeginScanAt"`
		// 扫描全部路径

		ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`
		// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径

		ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`
		// 超时时长，单位小时

		Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
		// 扫描范围0容器1主机节点

		ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`
		// true 全选，false 自选

		ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`
		// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定

		ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`
		// 自选排除或扫描的地址

		ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
		// 一键检测的超时设置

		ClickTimeout *uint64 `json:"ClickTimeout,omitempty" name:"ClickTimeout"`
		// 扫描路径模式：
		// SCAN_PATH_ALL：全部路径
		// SCAN_PATH_DEFAULT：默认路径
		// SCAN_PATH_USER_DEFINE：用户自定义路径
		//

		ScanPathMode *string `json:"ScanPathMode,omitempty" name:"ScanPathMode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusScanSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机的列表。

		HostList []*ComplianceHostInfo `json:"HostList,omitempty" name:"HostList"`
		// 主机的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 是否为全部主机节点

		AllChecked *bool `json:"AllChecked,omitempty" name:"AllChecked"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本地镜像重新漏洞扫描时的任务ID

		LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`
		// 仓库镜像重新漏洞扫描时的任务ID

		RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIngressForwardConfigRequest struct {
	*tchttp.BaseRequest

	// Ingress名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIngressForwardConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIngressForwardConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 访问控制策略信息列表

		RuleSet []*RuleBaseInfo `json:"RuleSet,omitempty" name:"RuleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostInfo struct {

	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip即内网ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 业务组

	Group *string `json:"Group,omitempty" name:"Group"`
	// docker 版本

	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`
	// docker 文件系统类型

	DockerFileSystemDriver *string `json:"DockerFileSystemDriver,omitempty" name:"DockerFileSystemDriver"`
	// 镜像个数

	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
	// 容器个数

	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// agent运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 是否是Containerd

	IsContainerd *bool `json:"IsContainerd,omitempty" name:"IsContainerd"`
	// 主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 主机实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 地域ID

	RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`
	// 所属项目

	Project *ProjectInfo `json:"Project,omitempty" name:"Project"`
	// 标签

	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
	// 集群id

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群接入状态

	ClusterAccessedStatus *string `json:"ClusterAccessedStatus,omitempty" name:"ClusterAccessedStatus"`
	// 计费核数

	ChargeCoresCnt *uint64 `json:"ChargeCoresCnt,omitempty" name:"ChargeCoresCnt"`
	// 防护状态:
	// 已防护: Defended
	// 未防护: UnDefended

	DefendStatus *string `json:"DefendStatus,omitempty" name:"DefendStatus"`
}

type DescribeVirusAutoIsolateSampleDownloadURLRequest struct {
	*tchttp.BaseRequest

	// 样本Md5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
}

func (r *DescribeVirusAutoIsolateSampleDownloadURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleDownloadURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryScanStatusOneKeyRequest struct {
	*tchttp.BaseRequest

	// 需要获取进度的镜像列表

	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`
	// 是否获取全部镜像

	All *bool `json:"All,omitempty" name:"All"`
	// 需要获取进度的镜像列表Id

	Id []*uint64 `json:"Id,omitempty" name:"Id"`
	// 获取进度的任务ID

	TaskID *uint64 `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *DescribeAssetImageRegistryScanStatusOneKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryScanStatusOneKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDetailRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
}

func (r *DescribeVulDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellWhiteListsRequest struct {
	*tchttp.BaseRequest

	// 白名单ids

	WhiteListIdSet []*string `json:"WhiteListIdSet,omitempty" name:"WhiteListIdSet"`
}

func (r *DeleteReverseShellWhiteListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellWhiteListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSuperNodeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 超级节点列表

		List []*SuperNodeListItem `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSuperNodeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSuperNodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveLocalStorageItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveLocalStorageItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveLocalStorageItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsEventDetailRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	EventID *uint64 `json:"EventID,omitempty" name:"EventID"`
}

func (r *DescribeRiskDnsEventDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsEventDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulContainerExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulContainerExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulContainerExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReverseShellStatusRequest struct {
	*tchttp.BaseRequest

	// 处理事件ids

	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
	// 标记事件的状态，
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyReverseShellStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReverseShellStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageVirusInfo struct {

	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 风险等级

	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 修护建议

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 首次发现时间

	FirstScanTime *string `json:"FirstScanTime,omitempty" name:"FirstScanTime"`
	// 最近扫描时间

	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
	// 文件md5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 检测平台
	// 1: 云查杀引擎
	// 2: tav
	// 3: binaryAi
	// 4: 异常行为
	// 5: 威胁情报

	CheckPlatform []*string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
}

type DescribeReverseShellRegexpWhiteListInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		RuleInfo *RegexpRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellRegexpWhiteListInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellRegexpWhiteListInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// base64编码的networkpolicy yaml字符串

	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageDenyEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageDenyEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageDenyEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterUninstallCmdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命令

		Command *string `json:"Command,omitempty" name:"Command"`
		// 文件下载地址

		URL *string `json:"URL,omitempty" name:"URL"`
		// 文件内容

		FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterUninstallCmdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterUninstallCmdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventEscapeImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险容器镜像列表

		List []*EventEscapeImageInfo `json:"List,omitempty" name:"List"`
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventEscapeImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventEscapeImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetImageRegistryAssetRequest struct {
	*tchttp.BaseRequest

	// 是否同步所有镜像仓库

	All *bool `json:"All,omitempty" name:"All"`
	// 需要同步的部分镜像仓库

	RegistryIds []*uint64 `json:"RegistryIds,omitempty" name:"RegistryIds"`
}

func (r *SyncAssetImageRegistryAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetImageRegistryAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventEscapeImageListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数:
	// EventType: 事件类型(MOUNT_SENSITIVE_PTAH:敏感路径挂载 PRIVILEGE_CONTAINER_START:特权容器)
	// Status: 事件状态(EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略)
	// ImageID: 镜像id
	// ImageName:镜像名称

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEventEscapeImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventEscapeImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetImageRegistryScanStopOneKeyRequest struct {
	*tchttp.BaseRequest

	// 是否扫描全部镜像

	All *bool `json:"All,omitempty" name:"All"`
	// 扫描的镜像列表

	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`
	// 扫描的镜像列表Id

	Id []*uint64 `json:"Id,omitempty" name:"Id"`
	// 停止的任务ID

	TaskID *uint64 `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *ModifyAssetImageRegistryScanStopOneKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetImageRegistryScanStopOneKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNewestVulRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNewestVulRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewestVulRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventType struct {

	// 事件类型, 字符串枚举，例如：ESCAPE_SYSCALL

	Type *string `json:"Type,omitempty" name:"Type"`
	// 事件名称，例如：敏感路径挂载

	EventName *string `json:"EventName,omitempty" name:"EventName"`
}

type CreateUserClusterExportJobRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表，为空表示导出全部

	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`
}

func (r *CreateUserClusterExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserClusterExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellInnerIPShowStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 展示内网IP

		InnerIPShow *bool `json:"InnerIPShow,omitempty" name:"InnerIPShow"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellInnerIPShowStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellInnerIPShowStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查杀容器个数

		ContainerTotal *uint64 `json:"ContainerTotal,omitempty" name:"ContainerTotal"`
		// 风险容器个数

		RiskContainerCnt *uint64 `json:"RiskContainerCnt,omitempty" name:"RiskContainerCnt"`
		// 扫描状态 任务状态:
		// SCAN_NONE:无，
		// SCAN_SCANNING:正在扫描中，
		// SCAN_FINISH：扫描完成，
		// SCAN_TIMEOUT：扫描超时
		// SCAN_CANCELING: 取消中
		// SCAN_CANCELED:已取消

		Status *string `json:"Status,omitempty" name:"Status"`
		// 扫描进度 I

		Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`
		// 已经扫描了的容器个数

		ContainerScanCnt *uint64 `json:"ContainerScanCnt,omitempty" name:"ContainerScanCnt"`
		// 风险个数

		RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
		// 剩余扫描时间

		LeftSeconds *uint64 `json:"LeftSeconds,omitempty" name:"LeftSeconds"`
		// 扫描开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 扫描结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 扫描类型，"CYCLE"：周期扫描， "MANUAL"：手动扫描

		ScanType *string `json:"ScanType,omitempty" name:"ScanType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusScanTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetClusterCheckIdWhiteListStatusRequest struct {
	*tchttp.BaseRequest

	// 检测项ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 是否忽略该项

	IsIgnored *bool `json:"IsIgnored,omitempty" name:"IsIgnored"`
}

func (r *SetClusterCheckIdWhiteListStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetClusterCheckIdWhiteListStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopVulScanTaskRequest struct {
	*tchttp.BaseRequest

	// 本地镜像漏洞扫描任务ID

	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`
	// 本地镜像ID，无则全部

	LocalImageIDs []*string `json:"LocalImageIDs,omitempty" name:"LocalImageIDs"`
	// 仓库镜像ID，无则全部

	RegistryImageIDs []*uint64 `json:"RegistryImageIDs,omitempty" name:"RegistryImageIDs"`
	// 仓库镜像漏洞扫描任务ID

	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
}

func (r *StopVulScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopVulScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAssetSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户的pod总量

		PodTotalCount *uint64 `json:"PodTotalCount,omitempty" name:"PodTotalCount"`
		// 用户运行状态的Pod数量

		RunningPodCount *uint64 `json:"RunningPodCount,omitempty" name:"RunningPodCount"`
		// Pending状态的Pod数量

		PendingPodCount *uint64 `json:"PendingPodCount,omitempty" name:"PendingPodCount"`
		// 用户的集群Service总量

		ServiceTotalCount *uint64 `json:"ServiceTotalCount,omitempty" name:"ServiceTotalCount"`
		// ClusterIp类型的Service数量

		ClusterIpCount *uint64 `json:"ClusterIpCount,omitempty" name:"ClusterIpCount"`
		// NodePort类型的Service数量

		NodePortCount *uint64 `json:"NodePortCount,omitempty" name:"NodePortCount"`
		// 用户的集群Ingress总量

		IngressTotalCount *uint64 `json:"IngressTotalCount,omitempty" name:"IngressTotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterAssetSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAssetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePeriodTaskListRequest struct {
	*tchttp.BaseRequest

	// 资产的类型，取值为：
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCompliancePeriodTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePeriodTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogCleanSettingInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecLogCleanSettingInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogCleanSettingInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRiskDnsEventExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRiskDnsEventExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskDnsEventExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待处理逃逸事件趋势

		List []*EscapeEventTendencyInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeEventTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageBindRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像绑定规则列表

		ImageBindRuleSet []*ImagesBindRuleInfo `json:"ImageBindRuleSet,omitempty" name:"ImageBindRuleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageBindRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageBindRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageAutoAuthorizedLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 自动授权结果镜像列表

		List []*AutoAuthorizedImageInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageAutoAuthorizedLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageAutoAuthorizedLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReverseShellStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReverseShellStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReverseShellStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLocalStorageExpireResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLocalStorageExpireResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLocalStorageExpireResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSuperNodeInstallTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	//
	// TaskID - String - 是否必填：否 - 任务ID

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeSuperNodeInstallTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSuperNodeInstallTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyWorkLoadDetailRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 负载名称

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 负载类型

	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤器

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkTopologyWorkLoadDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyWorkLoadDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReverseShellRegexpWhiteListExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateReverseShellRegexpWhiteListExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateReverseShellRegexpWhiteListExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanAuthorizedImageSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulScanAuthorizedImageSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanAuthorizedImageSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulnerabilityFocusRuleRequest struct {
	*tchttp.BaseRequest

	// CVSS评分

	CvssScore *float64 `json:"CvssScore,omitempty" name:"CvssScore"`
	// 漏洞标签

	Tags []*VulnerabilityFocusRuleTags `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyVulnerabilityFocusRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulnerabilityFocusRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNewPurchaseInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNewPurchaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewPurchaseInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalEventInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件详情

		Info *K8sApiAbnormalEventInfo `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalEventInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeESAggregationsRequest struct {
	*tchttp.BaseRequest

	// ES聚合条件JSON

	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeESAggregationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeESAggregationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanRegistryImageListRequest struct {
	*tchttp.BaseRequest

	// 漏洞扫描任务ID

	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件 OnlyShowLatest- string “true”代表是，“false”代表否 - 是否必填：否 - 仅展示影响最新版本镜像的漏洞
	// 	ImageID- string - 是否必填：否 - 镜像ID
	// 	ImageName- String -是否必填: 否 - 镜像名称
	// 	ScanStatus- string -是否必填: 否 - 检测状态。
	// 	NOT_SCAN：待检测，SCANNING：检查中，SCANNED：检查完成，SCAN_ERR：检查失败，CANCELED：检测停止

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulScanRegistryImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanRegistryImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVirusScanSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteComplianceAssetPolicySetFromWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteComplianceAssetPolicySetFromWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteComplianceAssetPolicySetFromWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallPolicyYamlDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群Id

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 策略名

		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
		// base64编码的yaml字符串

		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
		// 策略描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 策略创建时间

		PolicyCreateTime *string `json:"PolicyCreateTime,omitempty" name:"PolicyCreateTime"`
		// 策略源类型，分为System和Manual，分别代表手动和系统添加

		PolicySourceType *string `json:"PolicySourceType,omitempty" name:"PolicySourceType"`
		// 网络策略对应的网络插件

		NetworkPolicyPlugin *string `json:"NetworkPolicyPlugin,omitempty" name:"NetworkPolicyPlugin"`
		// 网络策略状态

		PublishStatus *string `json:"PublishStatus,omitempty" name:"PublishStatus"`
		// 网络发布结果

		PublishResult *string `json:"PublishResult,omitempty" name:"PublishResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallPolicyYamlDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPolicyYamlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户状态

		UserStatus *string `json:"UserStatus,omitempty" name:"UserStatus"`
		// 版本

		Version *string `json:"Version,omitempty" name:"Version"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryTimingScanTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageRegistryTimingScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryTimingScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogDeliveryClsSettingRequest struct {
	*tchttp.BaseRequest

	// 日志信息

	List []*SecLogDeliveryClsSettingInfo `json:"List,omitempty" name:"List"`
}

func (r *ModifySecLogDeliveryClsSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogDeliveryClsSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AffectedWorkloadItem struct {

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 工作负载名称

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 工作负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 检测结果的验证信息

	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
}

type DescribeVulScanInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本次扫描的本地镜像总数

		LocalImageScanCount *int64 `json:"LocalImageScanCount,omitempty" name:"LocalImageScanCount"`
		// 忽略的漏洞数量

		IgnoreVulCount *int64 `json:"IgnoreVulCount,omitempty" name:"IgnoreVulCount"`
		// 漏洞扫描的开始时间

		ScanStartTime *string `json:"ScanStartTime,omitempty" name:"ScanStartTime"`
		// 漏洞扫描的结束时间

		ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`
		// 发现风险镜像数量

		FoundRiskImageCount *int64 `json:"FoundRiskImageCount,omitempty" name:"FoundRiskImageCount"`
		// 本地发现漏洞数量

		FoundVulCount *int64 `json:"FoundVulCount,omitempty" name:"FoundVulCount"`
		// 扫描进度

		ScanProgress *float64 `json:"ScanProgress,omitempty" name:"ScanProgress"`
		// 本次扫描的仓库镜像总数

		RegistryImageScanCount *int64 `json:"RegistryImageScanCount,omitempty" name:"RegistryImageScanCount"`
		// 本地镜像最近一次的漏洞任务扫描ID

		LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`
		// 扫描状态:NOT_SCAN:未扫描，SCANNING:扫描中,SCANNED:完成，CANCELED：扫描停止

		Status *string `json:"Status,omitempty" name:"Status"`
		// 剩余时间，秒

		RemainingTime *float64 `json:"RemainingTime,omitempty" name:"RemainingTime"`
		// 仓库镜像最近一次的漏洞任务扫描ID

		RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
		// 仓库发现漏洞数量

		RegistryFoundVulCount *int64 `json:"RegistryFoundVulCount,omitempty" name:"RegistryFoundVulCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulScanInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallWhiteListDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白名单基本信息

		WhiteListDetailInfo *RiskSyscallWhiteListInfo `json:"WhiteListDetailInfo,omitempty" name:"WhiteListDetailInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallWhiteListDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallWhiteListDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSimpleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像列表

		ImageList []*ImageSimpleInfo `json:"ImageList,omitempty" name:"ImageList"`
		// 镜像数

		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageSimpleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageSimpleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulRegistryImageListRequest struct {
	*tchttp.BaseRequest

	// 漏洞ID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// OnlyAffectedNewestImage bool 是否影响最新镜像
	// ImageDigest 镜像Digest，支持模糊查询
	// ImageId 镜像ID，支持模糊查询
	// Namespace 命名空间，支持模糊查询
	// ImageTag 镜像版本，支持模糊查询
	// InstanceName 实例名称，支持模糊查询
	// ImageName 镜像名，支持模糊查询
	// ImageRepoAddress 镜像地址，支持模糊查询

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulRegistryImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulRegistryImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCheckItemListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检查项详情数组

		ClusterCheckItems []*ClusterCheckItem `json:"ClusterCheckItems,omitempty" name:"ClusterCheckItems"`
		// 检查项总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCheckItemListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCheckItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyRealFlowListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络流量列表
		//

		FlowList []*NetworkTopologyFlow `json:"FlowList,omitempty" name:"FlowList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkTopologyRealFlowListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyRealFlowListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageRegistryAutoAuthorizedStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageRegistryAutoAuthorizedStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageRegistryAutoAuthorizedStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCompliancePolicyItemFromWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCompliancePolicyItemFromWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCompliancePolicyItemFromWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`
		// 进程信息

		ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`
		// 事件描述

		EventDetail *EscapeEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`
		// 父进程信息

		ParentProcessInfo *ProcessBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`
		// 祖先进程信息

		AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeEventDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunTimeRiskInfo struct {

	// 数量

	Cnt *uint64 `json:"Cnt,omitempty" name:"Cnt"`
	// 风险等级：
	// CRITICAL: 严重
	// HIGH: 高
	// MEDIUM：中
	// LOW: 低

	Level *string `json:"Level,omitempty" name:"Level"`
}

type RenewImageAuthorizeStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewImageAuthorizeStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewImageAuthorizeStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusMonitorConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVirusMonitorConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusMonitorConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessLevelSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异常进程高危待处理事件数

		HighLevelEventCount *int64 `json:"HighLevelEventCount,omitempty" name:"HighLevelEventCount"`
		// 异常进程中危待处理事件数

		MediumLevelEventCount *int64 `json:"MediumLevelEventCount,omitempty" name:"MediumLevelEventCount"`
		// 异常进程低危待处理事件数

		LowLevelEventCount *int64 `json:"LowLevelEventCount,omitempty" name:"LowLevelEventCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessLevelSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessLevelSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIgnoreVulResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIgnoreVulResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIgnoreVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageRiskExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为10000

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *CreateAssetImageRiskExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageRiskExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenTcssTrialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 试用开通结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 试用开通开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenTcssTrialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenTcssTrialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIndexListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetSecLogTopicConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetSecLogTopicConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetSecLogTopicConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlEventDescription struct {

	// 事件规则

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 命中规则详细信息

	MatchRule *AccessControlChildRuleInfo `json:"MatchRule,omitempty" name:"MatchRule"`
	// 命中规则名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 命中规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 事件最后一次处理的时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type AddAndPublishNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 入站规则
	//
	// 全部允许：1
	//
	// 全部拒绝 ：2
	//
	// 自定义：3

	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`
	// 出站规则
	//
	// 全部允许：1
	//
	// 全部拒绝 ：2
	//
	// 自定义：3

	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`
	// pod选择器

	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 自定义规则

	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

func (r *AddAndPublishNetworkFirewallPolicyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAndPublishNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReverseShellInnerIPShowStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReverseShellInnerIPShowStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReverseShellInnerIPShowStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCheckComponentRequest struct {
	*tchttp.BaseRequest

	// 要安装的集群列表信息

	ClusterInfoList []*ClusterCreateComponentItem `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`
}

func (r *CreateCheckComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCheckComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAssetImageRegistryRegistryDetailRequest struct {
	*tchttp.BaseRequest

	// 仓库名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 仓库url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 仓库类型，列表：harbor

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 仓库版本

	RegistryVersion *string `json:"RegistryVersion,omitempty" name:"RegistryVersion"`
	// 网络类型，列表：public（公网）

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 区域，列表：default（默认）

	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`
	// 限速

	SpeedLimit *int64 `json:"SpeedLimit,omitempty" name:"SpeedLimit"`
	// 安全模式（证书校验）：0（默认） 非安全模式（跳过证书校验）：1

	Insecure *uint64 `json:"Insecure,omitempty" name:"Insecure"`
	// 联通性检测的配置

	ConnDetectConfig []*ConnDetectConfig `json:"ConnDetectConfig,omitempty" name:"ConnDetectConfig"`
}

func (r *UpdateAssetImageRegistryRegistryDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAssetImageRegistryRegistryDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogJoinObjectsRequest struct {
	*tchttp.BaseRequest

	// 绑定主机quuid列表

	BindList []*string `json:"BindList,omitempty" name:"BindList"`
	// 待解绑主机quuid列表

	UnBindList []*string `json:"UnBindList,omitempty" name:"UnBindList"`
	// 日志类型
	// bash日志: container_bash
	// 容器启动: container_launch
	// k8sApi: k8s_api

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 节点类型:
	// NORMAL: 普通节点(默认值)
	// SUPER: 超级节点
	//

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}

func (r *ModifySecLogJoinObjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogJoinObjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogVasInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 购买状态
		// 待购: Pending
		// 已购: Normal
		// 隔离: Isolate

		BuyStatus *string `json:"BuyStatus,omitempty" name:"BuyStatus"`
		// 存储时长(月)

		LogSaveMonth *int64 `json:"LogSaveMonth,omitempty" name:"LogSaveMonth"`
		// 起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 截止时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 存储容量(GB)

		LogCapacity *uint64 `json:"LogCapacity,omitempty" name:"LogCapacity"`
		// 资源ID

		ResourceID *string `json:"ResourceID,omitempty" name:"ResourceID"`
		// 是否曾经购买过(false:未曾 true:曾经购买过)

		IsPurchased *bool `json:"IsPurchased,omitempty" name:"IsPurchased"`
		// 试用存储容量(GB)

		TrialCapacity *uint64 `json:"TrialCapacity,omitempty" name:"TrialCapacity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogVasInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogVasInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageVul struct {

	// 漏洞id

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 观点验证程序id

	POCID *string `json:"POCID,omitempty" name:"POCID"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 涉及组件信息

	Components []*ComponentsInfo `json:"Components,omitempty" name:"Components"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 分类2

	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 描述

	Des *string `json:"Des,omitempty" name:"Des"`
	// 解决方案

	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`
	// 引用

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 防御方案

	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`
	// 提交时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// Cvss分数

	CvssScore *string `json:"CvssScore,omitempty" name:"CvssScore"`
	// Cvss信息

	CvssVector *string `json:"CvssVector,omitempty" name:"CvssVector"`
	// 是否建议修复

	IsSuggest *string `json:"IsSuggest,omitempty" name:"IsSuggest"`
	// 修复版本号

	FixedVersions *string `json:"FixedVersions,omitempty" name:"FixedVersions"`
	// 漏洞标签:"CanBeFixed","DynamicLevelPoc","DynamicLevelExp"

	Tag []*string `json:"Tag,omitempty" name:"Tag"`
	// 组件名

	Component *string `json:"Component,omitempty" name:"Component"`
	// 组件版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 攻击热度 0-3

	AttackLevel *int64 `json:"AttackLevel,omitempty" name:"AttackLevel"`
	// 镜像层信息列表

	LayerInfos []*ImageVulLayerInfo `json:"LayerInfos,omitempty" name:"LayerInfos"`
}

type DescribeAffectedNodeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受影响的节点总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 受影响的节点列表

		AffectedNodeList []*AffectedNodeItem `json:"AffectedNodeList,omitempty" name:"AffectedNodeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAffectedNodeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedNodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自动隔离开关(true:开 false:关)

		AutoIsolateSwitch *bool `json:"AutoIsolateSwitch,omitempty" name:"AutoIsolateSwitch"`
		// 是否中断隔离文件关联的进程(true:是 false:否)

		IsKillProgress *bool `json:"IsKillProgress,omitempty" name:"IsKillProgress"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSuperNodePodListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>NodeUniqueID- String - 是否必填：否 - 节点唯一id </li>
	// <li>PodName- String - 是否必填：否 - Pod示例名称 </li>
	// <li>PodIP- String - 是否必填：否 - POD IP </li>
	// <li>Namespace- String - 是否必填：否 - 命名空间 </li>
	// <li>Deployment- String - 是否必填：否 - 所属工作负载 </li>
	// <li>Status- String - 是否必填：否 - 状态 </li>
	//

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeSuperNodePodListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSuperNodePodListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkPolicyInfoItem struct {

	// 网络策略名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网络策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 发布状态：
	//
	// 开启待确认：PublishedNoConfirm
	//
	// 开启已确认：PublishedConfirmed
	//
	// 关闭中：unPublishing
	//
	// 开启中：Publishing
	//
	// 待开启：unPublishEdit

	PublishStatus *string `json:"PublishStatus,omitempty" name:"PublishStatus"`
	// 策略类型：
	//
	// 自动发现：System
	//
	// 手动添加：Manual

	PolicySourceType *string `json:"PolicySourceType,omitempty" name:"PolicySourceType"`
	// 策略空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 策略创建日期

	PolicyCreateTime *string `json:"PolicyCreateTime,omitempty" name:"PolicyCreateTime"`
	// 策略类型
	//
	// kube-router：KubeRouter
	//
	// cilium：Cilium

	NetworkPolicyPlugin *string `json:"NetworkPolicyPlugin,omitempty" name:"NetworkPolicyPlugin"`
	// 策略发布结果

	PublishResult *string `json:"PublishResult,omitempty" name:"PublishResult"`
	// 入站规则
	//
	// 全部允许：1
	//
	// 全部拒绝 ：2
	//
	// 自定义：3

	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`
	// 入站规则
	//
	// 全部允许：1
	//
	// 全部拒绝 ：2
	//
	// 自定义：3

	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`
	// 作用对象

	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
	// 网络策略Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeAssetImageRegistryDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像Digest

		ImageDigest *string `json:"ImageDigest,omitempty" name:"ImageDigest"`
		// 镜像地址

		ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`
		// 镜像类型

		RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
		// 仓库名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 镜像版本

		ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
		// 扫描时间

		ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
		// 扫描状态

		ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
		// 安全漏洞数

		VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`
		// 木马病毒数

		VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`
		// 风险行为数

		RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
		// 敏感信息数

		SentiveInfoCnt *uint64 `json:"SentiveInfoCnt,omitempty" name:"SentiveInfoCnt"`
		// 是否可信镜像

		IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`
		// 镜像系统

		OsName *string `json:"OsName,omitempty" name:"OsName"`
		// 木马扫描错误

		ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`
		// 漏洞扫描错误

		ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`
		// 层文件信息

		LayerInfo *string `json:"LayerInfo,omitempty" name:"LayerInfo"`
		// 实例id

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例名称

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 命名空间

		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
		// 高危扫描错误

		ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`
		// 木马信息扫描进度

		ScanVirusProgress *uint64 `json:"ScanVirusProgress,omitempty" name:"ScanVirusProgress"`
		// 漏洞扫描进度

		ScanVulProgress *uint64 `json:"ScanVulProgress,omitempty" name:"ScanVulProgress"`
		// 敏感扫描进度

		ScanRiskProgress *uint64 `json:"ScanRiskProgress,omitempty" name:"ScanRiskProgress"`
		// 剩余扫描时间秒

		ScanRemainTime *uint64 `json:"ScanRemainTime,omitempty" name:"ScanRemainTime"`
		// cve扫描状态

		CveStatus *string `json:"CveStatus,omitempty" name:"CveStatus"`
		// 高危扫描状态

		RiskStatus *string `json:"RiskStatus,omitempty" name:"RiskStatus"`
		// 木马扫描状态

		VirusStatus *string `json:"VirusStatus,omitempty" name:"VirusStatus"`
		// 总进度

		Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
		// 授权状态

		IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
		// 镜像大小

		ImageSize *uint64 `json:"ImageSize,omitempty" name:"ImageSize"`
		// 镜像Id

		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
		// 镜像区域

		RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`
		// 镜像创建的时间

		ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`
		// 敏感信息数

		SensitiveInfoCnt *uint64 `json:"SensitiveInfoCnt,omitempty" name:"SensitiveInfoCnt"`
		// Id

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSearchTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSearchTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSearchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterDetailExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务Id

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterDetailExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterDetailExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateK8sApiAbnormalRuleExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By []*string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateK8sApiAbnormalRuleExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateK8sApiAbnormalRuleExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ABTestConfig struct {

	// 灰度项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// true：正在灰度，false：不在灰度

	Status *bool `json:"Status,omitempty" name:"Status"`
}

type CreateSystemVulExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞子类型</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateSystemVulExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSystemVulExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessEventTendencyInfo struct {

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 待处理代理软件事件数

	ProxyToolEventCount *int64 `json:"ProxyToolEventCount,omitempty" name:"ProxyToolEventCount"`
	// 待处理横向参透事件数

	TransferControlEventCount *int64 `json:"TransferControlEventCount,omitempty" name:"TransferControlEventCount"`
	// 待处理恶意命令事件数

	AttackCmdEventCount *int64 `json:"AttackCmdEventCount,omitempty" name:"AttackCmdEventCount"`
	// 待处理反弹shell事件数

	ReverseShellEventCount *int64 `json:"ReverseShellEventCount,omitempty" name:"ReverseShellEventCount"`
	// 待处理无文件程序执行事件数

	FilelessEventCount *int64 `json:"FilelessEventCount,omitempty" name:"FilelessEventCount"`
	// 待处理高危命令事件数

	RiskCmdEventCount *int64 `json:"RiskCmdEventCount,omitempty" name:"RiskCmdEventCount"`
	// 待处理敏感服务异常子进程启动事件数

	AbnormalChildProcessEventCount *int64 `json:"AbnormalChildProcessEventCount,omitempty" name:"AbnormalChildProcessEventCount"`
	// 待处理自定义规则事件数

	UserDefinedRuleEventCount *int64 `json:"UserDefinedRuleEventCount,omitempty" name:"UserDefinedRuleEventCount"`
}

type AddEditImageAutoAuthorizedRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddEditImageAutoAuthorizedRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditImageAutoAuthorizedRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启定期扫描

		EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`
		// 检测周期每隔多少天

		Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`
		// 扫描开始时间

		BeginScanAt *string `json:"BeginScanAt,omitempty" name:"BeginScanAt"`
		// 超时时长，单位小时

		Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
		// SCAN_NODE:扫描节点
		// SCAN_CONTAINER:扫描容器

		ScanRangeType *string `json:"ScanRangeType,omitempty" name:"ScanRangeType"`
		// 自选扫描范围的容器id或者节点id

		ScanIDs []*ScanRangeInfo `json:"ScanIDs,omitempty" name:"ScanIDs"`
		// 自选排除或扫描的地址

		ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
		// 扫描路径模式：
		// SCAN_PATH_ALL：全部路径
		// SCAN_PATH_DEFAULT：默认路径
		// SCAN_PATH_USER_DEFINE：用户自定义路径
		//

		ScanPathMode *string `json:"ScanPathMode,omitempty" name:"ScanPathMode"`
		// true:包含路径
		// false:排除路径

		IsIncludePath *bool `json:"IsIncludePath,omitempty" name:"IsIncludePath"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusScanConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallAuditRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群审计总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群的审计详细信息

		AuditList []*NetworkAuditRecord `json:"AuditList,omitempty" name:"AuditList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallAuditRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallAuditRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCompliancePeriodTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCompliancePeriodTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCompliancePeriodTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirusAutoIsolateSampleInfo struct {

	// 文件MD5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 病毒名

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 最近编辑时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 自动隔离开关(true:开 false:关)

	AutoIsolateSwitch *bool `json:"AutoIsolateSwitch,omitempty" name:"AutoIsolateSwitch"`
}

type DescribeAssetImageRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像病毒列表

		List []*ImageRiskInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddCompliancePolicyAssetSetToWhitelistRequest struct {
	*tchttp.BaseRequest

	// 检查项ID

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 需要忽略指定检查项内的资产ID列表

	CustomerAssetItemIdSet []*uint64 `json:"CustomerAssetItemIdSet,omitempty" name:"CustomerAssetItemIdSet"`
}

func (r *AddCompliancePolicyAssetSetToWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCompliancePolicyAssetSetToWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetFilters struct {

	// 过滤键的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 是否模糊查询

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type SoftQuotaDayInfo struct {

	// 扣费时间

	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`
	// 计费核数(已废弃)

	CoresCnt *uint64 `json:"CoresCnt,omitempty" name:"CoresCnt"`
}

type ModifySecLogDeliveryKafkaSettingRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 用户名

	User *string `json:"User,omitempty" name:"User"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 日志类型队列

	LogTypeList []*SecLogDeliveryKafkaSettingInfo `json:"LogTypeList,omitempty" name:"LogTypeList"`
	// 接入类型

	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`
	// kafka版本号

	KafkaVersion *string `json:"KafkaVersion,omitempty" name:"KafkaVersion"`
	// 地域ID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
}

func (r *ModifySecLogDeliveryKafkaSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogDeliveryKafkaSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRiskSyscallWhiteListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRiskSyscallWhiteListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRiskSyscallWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogDeliveryClsSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecLogDeliveryClsSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogDeliveryClsSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecLogJoinInfo struct {

	// 已接入普通主机数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 已接入超级节点数量

	SuperNodeCount *uint64 `json:"SuperNodeCount,omitempty" name:"SuperNodeCount"`
	// 是否已接入(true:已接入 false:未接入)

	IsJoined *bool `json:"IsJoined,omitempty" name:"IsJoined"`
	// 日志类型(
	// 容器bash:  "container_bash"
	// 容器启动: "container_launch"
	// k8sApi: "k8s_api"
	// )

	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

type DeleteUserClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 地域

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

func (r *DeleteUserClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallEventsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeRiskSyscallEventsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallEventsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表

		List []*HostInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulAffectedComponentInfo struct {

	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件版本

	Version []*string `json:"Version,omitempty" name:"Version"`
	// 组件修复版本

	FixedVersion []*string `json:"FixedVersion,omitempty" name:"FixedVersion"`
}

type DescribeVirusDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像ID

		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
		// 镜像名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 木马文件大小

		Size *uint64 `json:"Size,omitempty" name:"Size"`
		// 木马文件路径

		FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
		// 最近生成时间

		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
		// 病毒名称

		VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
		// 风险等级 RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。

		RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
		// 容器名称

		ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
		// 容器id

		ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
		// 主机名称

		HostName *string `json:"HostName,omitempty" name:"HostName"`
		// 主机id

		HostId *string `json:"HostId,omitempty" name:"HostId"`
		// 进程名称

		ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
		// 进程路径

		ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
		// 进程md5

		ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
		// 进程id

		ProcessId *uint64 `json:"ProcessId,omitempty" name:"ProcessId"`
		// 进程参数

		ProcessArgv *string `json:"ProcessArgv,omitempty" name:"ProcessArgv"`
		// 进程链

		ProcessChan *string `json:"ProcessChan,omitempty" name:"ProcessChan"`
		// 进程组

		ProcessAccountGroup *string `json:"ProcessAccountGroup,omitempty" name:"ProcessAccountGroup"`
		// 进程启动用户

		ProcessStartAccount *string `json:"ProcessStartAccount,omitempty" name:"ProcessStartAccount"`
		// 进程文件权限

		ProcessFileAuthority *string `json:"ProcessFileAuthority,omitempty" name:"ProcessFileAuthority"`
		// 来源：0：一键扫描， 1：定时扫描 2：实时监控

		SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
		// 标签

		Tags []*string `json:"Tags,omitempty" name:"Tags"`
		// 事件描述

		HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
		// 建议方案

		SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
		// 备注

		Mark *string `json:"Mark,omitempty" name:"Mark"`
		// 风险文件名称

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// 文件MD5

		FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`
		// 事件类型

		EventType *string `json:"EventType,omitempty" name:"EventType"`
		// 集群名称

		PodName *string `json:"PodName,omitempty" name:"PodName"`
		// DEAL_NONE:文件待处理
		// DEAL_IGNORE:已经忽略
		// DEAL_ADD_WHITELIST:加白
		// DEAL_DEL:文件已经删除
		// DEAL_ISOLATE:已经隔离
		// DEAL_ISOLATING:隔离中
		// DEAL_ISOLATE_FAILED:隔离失败
		// DEAL_RECOVERING:恢复中
		// DEAL_RECOVER_FAILED: 恢复失败

		Status *string `json:"Status,omitempty" name:"Status"`
		// 失败子状态:
		// FILE_NOT_FOUND:文件不存在
		// FILE_ABNORMAL:文件异常
		// FILE_ABNORMAL_DEAL_RECOVER:恢复文件时，文件异常
		// BACKUP_FILE_NOT_FOUND:备份文件不存在
		// CONTAINER_NOT_FOUND_DEAL_ISOLATE:隔离时，容器不存在
		// CONTAINER_NOT_FOUND_DEAL_RECOVER:恢复时，容器不存在

		SubStatus *string `json:"SubStatus,omitempty" name:"SubStatus"`
		// 内网ip

		HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
		// 外网ip

		ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`
		// 父进程启动用户

		PProcessStartUser *string `json:"PProcessStartUser,omitempty" name:"PProcessStartUser"`
		// 父进程用户组

		PProcessUserGroup *string `json:"PProcessUserGroup,omitempty" name:"PProcessUserGroup"`
		// 父进程路径

		PProcessPath *string `json:"PProcessPath,omitempty" name:"PProcessPath"`
		// 父进程命令行参数

		PProcessParam *string `json:"PProcessParam,omitempty" name:"PProcessParam"`
		// 祖先进程启动用户

		AncestorProcessStartUser *string `json:"AncestorProcessStartUser,omitempty" name:"AncestorProcessStartUser"`
		// 祖先进程用户组

		AncestorProcessUserGroup *string `json:"AncestorProcessUserGroup,omitempty" name:"AncestorProcessUserGroup"`
		// 祖先进程路径

		AncestorProcessPath *string `json:"AncestorProcessPath,omitempty" name:"AncestorProcessPath"`
		// 祖先进程命令行参数

		AncestorProcessParam *string `json:"AncestorProcessParam,omitempty" name:"AncestorProcessParam"`
		// 事件最后一次处理的时间

		OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
		// 容器隔离状态

		ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
		// 容器隔离子状态

		ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
		// 容器隔离操作来源

		ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
		// 检测平台
		// 1: 云查杀引擎
		// 2: tav
		// 3: binaryAi
		// 4: 异常行为
		// 5: 威胁情报

		CheckPlatform []*string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
		// 文件访问时间

		FileAccessTime *string `json:"FileAccessTime,omitempty" name:"FileAccessTime"`
		// 文件修改时间

		FileModifyTime *string `json:"FileModifyTime,omitempty" name:"FileModifyTime"`
		// 节点子网ID

		NodeSubNetID *string `json:"NodeSubNetID,omitempty" name:"NodeSubNetID"`
		// 节点子网名称

		NodeSubNetName *string `json:"NodeSubNetName,omitempty" name:"NodeSubNetName"`
		// 节点子网网段

		NodeSubNetCIDR *string `json:"NodeSubNetCIDR,omitempty" name:"NodeSubNetCIDR"`
		// 集群id

		ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
		// pod ip

		PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
		// pod状态

		PodStatus *string `json:"PodStatus,omitempty" name:"PodStatus"`
		// 节点唯一ID

		NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
		// 节点类型：NORMAL普通节点、SUPER超级节点

		NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
		// 节点ID

		NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// Namespace

		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
		// 工作负载类型

		WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncComplianceHostAssetsRequest struct {
	*tchttp.BaseRequest
}

func (r *SyncComplianceHostAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncComplianceHostAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefencePluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞防御插件列表

		List []*VulDefencePlugin `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefencePluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefencePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReverseShellRegexpRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReverseShellRegexpRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReverseShellRegexpRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetComplianceHostRangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 如果需要立即执行，返回执行任务id

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetComplianceHostRangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetComplianceHostRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceWhitelistItemListRequest struct {
	*tchttp.BaseRequest

	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 要获取的数量，默认为10，最大为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 资产类型列表。

	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`
	// 查询过滤器

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 desc asc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeComplianceWhitelistItemListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceWhitelistItemListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostDetailRequest struct {
	*tchttp.BaseRequest

	// 主机id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *DescribeAssetHostDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserClusterRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeUserClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallNamespaceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群的详细信息

		ClusterNamespaceList []*NetworkClusterNamespaceInfo `json:"ClusterNamespaceList,omitempty" name:"ClusterNamespaceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallNamespaceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallNamespaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCompliancePolicyItemFromWhitelistRequest struct {
	*tchttp.BaseRequest

	// 指定的白名单项的ID的列表

	WhitelistIdSet []*uint64 `json:"WhitelistIdSet,omitempty" name:"WhitelistIdSet"`
}

func (r *DeleteCompliancePolicyItemFromWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCompliancePolicyItemFromWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞防御事件列表

		List []*VulDefenceEvent `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefenceEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 同步任务发送结果

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerNetworkInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// endpoint id

		EndpointID *string `json:"EndpointID,omitempty" name:"EndpointID"`
		// 模式:bridge

		Mode *string `json:"Mode,omitempty" name:"Mode"`
		// 网络名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 网络ID

		NetworkID *string `json:"NetworkID,omitempty" name:"NetworkID"`
		// 网关

		Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
		// IPV4地址

		Ipv4 *string `json:"Ipv4,omitempty" name:"Ipv4"`
		// IPV6地址

		Ipv6 *string `json:"Ipv6,omitempty" name:"Ipv6"`
		// MAC 地址

		MAC *string `json:"MAC,omitempty" name:"MAC"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetContainerNetworkInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerNetworkInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIACCheckTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeIACCheckTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIACCheckTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyFlowSwitchListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 节点总数

		NetworkClusterSwitchList []*NetworkClusterSwitchListItem `json:"NetworkClusterSwitchList,omitempty" name:"NetworkClusterSwitchList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkTopologyFlowSwitchListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyFlowSwitchListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusEventTendencyRequest struct {
	*tchttp.BaseRequest

	// 趋势周期(默认为7天)

	TendencyPeriod *uint64 `json:"TendencyPeriod,omitempty" name:"TendencyPeriod"`
}

func (r *DescribeVirusEventTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusEventTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaliciousConnectionBlackListRequest struct {
	*tchttp.BaseRequest

	// 黑名单记录ID数组

	IDSet []*int64 `json:"IDSet,omitempty" name:"IDSet"`
}

func (r *DeleteMaliciousConnectionBlackListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaliciousConnectionBlackListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WhiteListNodeInfo struct {

	// 加入白名单的时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 白名单节点Id

	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 节点的角色，Master、Work等

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// 内网ip地址

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type AddAndPublishNetworkFirewallPolicyYamlDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAndPublishNetworkFirewallPolicyYamlDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAndPublishNetworkFirewallPolicyYamlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PodContainerInfo struct {

	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// CPU使用率

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 内存使用 KB

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 重启次数

	RestartCount *uint64 `json:"RestartCount,omitempty" name:"RestartCount"`
}

type CreateImageExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateImageExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterIngressListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群Ingress列表

		ClusterIngressList []*ClusterIngressInfo `json:"ClusterIngressList,omitempty" name:"ClusterIngressList"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterIngressListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterIngressListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SuperNodeInstallTaskListItem struct {

	// 任务id

	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
	// 安装起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 安装超级节点个数

	SuperNodeCount *uint64 `json:"SuperNodeCount,omitempty" name:"SuperNodeCount"`
	// 安装来源:
	//   TCSS:容器安全服务
	//   EKS:容器服务

	Source *string `json:"Source,omitempty" name:"Source"`
	// 状态	INSTALLING:安装中;
	// 	INSTALLATIONFAILURE:安装失败;
	// 	INSTALLATIONSUCCESS:安装成功
	// 	PARTIALINSTALLATIONFAILURE:部分安装失败

	Status *string `json:"Status,omitempty" name:"Status"`
	// 成功个数

	SuccessCount *uint64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
	// 失败个数

	FailCount *uint64 `json:"FailCount,omitempty" name:"FailCount"`
}

type DescribeIngressForwardConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Ingress的转发配置信息列表

		ConfigList []*IngressForwardConfig `json:"ConfigList,omitempty" name:"ConfigList"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIngressForwardConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIngressForwardConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusMonitorSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVirusMonitorSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusMonitorSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyEventDetailRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	EventID *int64 `json:"EventID,omitempty" name:"EventID"`
}

func (r *DescribeImageDenyEventDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellWhiteListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 白名单信息列表

		WhiteListSet []*ReverseShellWhiteListBaseInfo `json:"WhiteListSet,omitempty" name:"WhiteListSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellWhiteListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmergencyVulInfo struct {

	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 漏洞标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// CVSS V3分数

	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// CVE编号

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// 漏洞披露时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// 最近发现时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 应急漏洞风险情况：NOT_SCAN：未扫描，SCANNING：扫描中，SCANNED_NOT_RISK：已扫描，暂未风险 ，SCANNED_RISK：已扫描存在风险

	Status *string `json:"Status,omitempty" name:"Status"`
	// 漏洞ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 防御状态，NO_DEFENDED:未防御，DEFENDED:已防御

	DefenceStatus *string `json:"DefenceStatus,omitempty" name:"DefenceStatus"`
	// 漏洞防御主机范围: MANUAL:自选主机节点，ALL:全部

	DefenceScope *string `json:"DefenceScope,omitempty" name:"DefenceScope"`
	// 漏洞防御主机数量

	DefenceHostCount *int64 `json:"DefenceHostCount,omitempty" name:"DefenceHostCount"`
	// 已防御攻击次数

	DefendedCount *int64 `json:"DefendedCount,omitempty" name:"DefendedCount"`
}

type DescribeAssetImageRegistryVirusListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像信息

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVirusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVirusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRiskDnsEventStatusRequest struct {
	*tchttp.BaseRequest

	// 恶意请求事件ID数组。加白时必需，否则Filters和EventIDSet二者选其一。

	EventIDSet []*uint64 `json:"EventIDSet,omitempty" name:"EventIDSet"`
	// 标记事件的状态：
	// EVENT_UNDEAL:未处理（取消忽略），
	// EVENT_DEALED:已处理，
	// EVENT_IGNORE:忽略，
	// EVENT_DELETE：已删除
	// EVENT_ADD_WHITE：加白
	// EVENT_ISOLATE_CONTAINER：隔离容器
	// EVENT_RESOTRE_CONTAINER：恢复容器

	EventStatus *string `json:"EventStatus,omitempty" name:"EventStatus"`
	// 白名单域名/IP

	Address *string `json:"Address,omitempty" name:"Address"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 相同的请求域名/IP事件加白处理

	AllSameEventAddWhite *bool `json:"AllSameEventAddWhite,omitempty" name:"AllSameEventAddWhite"`
	// 加白的事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP

	AddWhiteEventType *string `json:"AddWhiteEventType,omitempty" name:"AddWhiteEventType"`
}

func (r *ModifyRiskDnsEventStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskDnsEventStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 逃逸白名单列表

		List []*EscapeWhiteListInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEscapeRegexpWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEscapeRegexpWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEscapeRegexpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRefreshTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 刷新任务状态，可能为：Task_New,Task_Running,Task_Finish,Task_Error,Task_NoExist

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRefreshTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRefreshTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLocalStorageItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 值

		Value *string `json:"Value,omitempty" name:"Value"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLocalStorageItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLocalStorageItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProcessEventsExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProcessEventsExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProcessEventsExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetProcessListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>RunAs - String - 是否必填：否 - 运行用户筛选，</li>
	// <li>ContainerID - String - 是否必填：否 - 容器id</li>
	// <li>HostID- String - 是否必填：是 - 主机id</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>ProcessName- string - 是否必填：否 - 进程名搜索</li>
	// <li>Pid- string - 是否必填：否 - 进程id搜索(关联进程)</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetProcessListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddClusterNodeToWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加检测项白名单的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddClusterNodeToWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterNodeToWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanComplianceScanFailedAssetsRequest struct {
	*tchttp.BaseRequest

	// 要重新扫描的客户资产项ID的列表。

	CustomerAssetIdSet []*uint64 `json:"CustomerAssetIdSet,omitempty" name:"CustomerAssetIdSet"`
}

func (r *ScanComplianceScanFailedAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanComplianceScanFailedAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaliciousConnectionBlackListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaliciousConnectionBlackListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaliciousConnectionBlackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnauthorizedCoresTendency struct {

	// 日期

	DateTime *string `json:"DateTime,omitempty" name:"DateTime"`
	// 未授权的核数

	CoresCount *int64 `json:"CoresCount,omitempty" name:"CoresCount"`
}

type DescribeNetworkFirewallPolicyDiscoverRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeNetworkFirewallPolicyDiscoverRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPolicyDiscoverRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterIngressListRequest struct {
	*tchttp.BaseRequest

	// 集群Id，不填表示查询用户所有ingress

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：Name, Address等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClusterIngressListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterIngressListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWebVulExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞子类型</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateWebVulExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWebVulExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageSimpleInfo struct {

	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 类型

	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
	// 关联容器数

	ContainerCnt *int64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
}

type WhiteListRegexpExpressionInfo struct {

	// 逻辑符号
	// 与 (AND)
	// 或 (OR)
	// 非 (NOT)

	LogicSymbol *string `json:"LogicSymbol,omitempty" name:"LogicSymbol"`
	// 匹配字段

	MatchField *string `json:"MatchField,omitempty" name:"MatchField"`
	// 匹配内容

	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
}

type DeleteK8sApiAbnormalRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteK8sApiAbnormalRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteK8sApiAbnormalRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLocalStorageItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLocalStorageItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLocalStorageItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetImageRegistryScanStopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetImageRegistryScanStopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetImageRegistryScanStopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSuperNodeInstallTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 超级节点安装任务详情列表

		List []*SuperNodeInstallTaskDetailInfo `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSuperNodeInstallTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSuperNodeInstallTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellRegexpWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReverseShellRegexpWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellRegexpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusMonitorConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVirusMonitorConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusMonitorConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulRegistryImageExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulRegistryImageExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulRegistryImageExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkPorts struct {

	// 网络策略协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 网络策略策略端口

	Port *string `json:"Port,omitempty" name:"Port"`
}

type ModifyAutoUpgradeSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoUpgradeSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoUpgradeSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEscapeRegexpWhiteListExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleName - string  - 是否必填: 否 -规则名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateEscapeRegexpWhiteListExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEscapeRegexpWhiteListExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkTopologyFlowSwitchTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建确认任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkTopologyFlowSwitchTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkTopologyFlowSwitchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRiskListExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一的任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRiskListExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskListExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命名空间列表

		ClusterNamespaceList []*ClusterNamespaceInfo `json:"ClusterNamespaceList,omitempty" name:"ClusterNamespaceList"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceK8SDetailInfo struct {

	// K8S集群的名称。

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// K8S集群的版本。

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
}

type DescribeIndexListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ES 索引信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndexListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebServiceListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>
	// <li>Type- String - 是否必填：否 - 主机运行状态筛选，"Apache"
	// "Jboss"
	// "lighttpd"
	// "Nginx"
	// "Tomcat"</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetWebServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启:0: 关闭 1:开启

		IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
		// 漏洞防御主机范围: 0:自选主机节点，1:全部

		Scope *int64 `json:"Scope,omitempty" name:"Scope"`
		// 漏洞防御主机数量

		HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
		// 开启漏洞防御异常主机数量

		ExceptionHostCount *int64 `json:"ExceptionHostCount,omitempty" name:"ExceptionHostCount"`
		// 自选漏洞防御主机

		HostIDs []*string `json:"HostIDs,omitempty" name:"HostIDs"`
		// 开通容器安全的主机总数

		HostTotalCount *int64 `json:"HostTotalCount,omitempty" name:"HostTotalCount"`
		// 支持防御的漏洞数

		SupportDefenseVulCount *int64 `json:"SupportDefenseVulCount,omitempty" name:"SupportDefenseVulCount"`
		// 普通节点个数

		HostNodeCount *int64 `json:"HostNodeCount,omitempty" name:"HostNodeCount"`
		// 超级节点范围

		SuperScope *int64 `json:"SuperScope,omitempty" name:"SuperScope"`
		// 超级节点个数

		SuperNodeCount *int64 `json:"SuperNodeCount,omitempty" name:"SuperNodeCount"`
		// 超级节点Id列表

		SuperNodeIds []*string `json:"SuperNodeIds,omitempty" name:"SuperNodeIds"`
		// 开通容器安全的超级结点总数

		NodeTotalCount *int64 `json:"NodeTotalCount,omitempty" name:"NodeTotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefenceSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulLevelImageSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulLevelImageSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulLevelImageSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirusTendencyInfo struct {

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 待处理事件总数

	PendingEventCount *uint64 `json:"PendingEventCount,omitempty" name:"PendingEventCount"`
	// 风险容器总数

	RiskContainerCount *uint64 `json:"RiskContainerCount,omitempty" name:"RiskContainerCount"`
	// 事件总数

	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
	// 隔离事件总数

	IsolateEventCount *uint64 `json:"IsolateEventCount,omitempty" name:"IsolateEventCount"`
}

type ComplianceAffectedAsset struct {

	// 为客户分配的唯一的资产项的ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 资产项的名称。

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产项的类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 检测状态
	//
	// CHECK_INIT, 待检测
	//
	// CHECK_RUNNING, 检测中
	//
	// CHECK_FINISHED, 检测完成
	//
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 节点名称。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 上次检测的时间，格式为”YYYY-MM-DD HH:m::SS“。
	//
	// 如果没有检测过，此处为”0000-00-00 00:00:00“。

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 检测结果。取值为：
	//
	// RESULT_FAILED: 未通过
	//
	// RESULT_PASSED: 通过

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 主机IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 镜像的tag

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 检查项验证信息

	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
	// 主机实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeAccessControlDetailRequest struct {
	*tchttp.BaseRequest

	// 事件唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeAccessControlDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest

	// 策略Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeNetworkFirewallPolicyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReverseShellInnerIPShowStatusRequest struct {
	*tchttp.BaseRequest

	// 内网ip展示

	InnerIPShow *bool `json:"InnerIPShow,omitempty" name:"InnerIPShow"`
}

func (r *ModifyReverseShellInnerIPShowStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReverseShellInnerIPShowStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkCustomPolicy struct {

	// 网络策略方向，分为FROM和TO

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 网络策略策略端口

	Ports []*NetworkPorts `json:"Ports,omitempty" name:"Ports"`
	// 网络策略策略对象
	//
	// 开启待确认：PublishedNoConfirm
	//
	// 开启已确认：PublishedConfirmed
	//
	// 关闭中：unPublishing
	//
	// 开启中：Publishing
	//
	// 待开启：unPublishEdit

	Peer []*NetworkPeer `json:"Peer,omitempty" name:"Peer"`
}

type ClsTopicInfo struct {

	// 主题ID

	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ModifyImageAuthorizedRequest struct {
	*tchttp.BaseRequest

	// 根据满足条件的本地镜像授权，AllLocalImages为false且LocalImageIds为空和UpdatedLocalImageCnt大于0时，需要。

	LocalImageFilter []*AssetFilters `json:"LocalImageFilter,omitempty" name:"LocalImageFilter"`
	// 根据满足条件的仓库镜像授权，AllRegistryImages为false且RegistryImageIds为空和UpdatedRegistryImageCnt大于0时，需要。

	RegistryImageFilter []*AssetFilters `json:"RegistryImageFilter,omitempty" name:"RegistryImageFilter"`
	// 根据满足条件的镜像授权,同时排除的本地镜像。

	ExcludeLocalImageIds []*string `json:"ExcludeLocalImageIds,omitempty" name:"ExcludeLocalImageIds"`
	// 根据满足条件的镜像授权,同时排除的仓库镜像。

	ExcludeRegistryImageIds []*string `json:"ExcludeRegistryImageIds,omitempty" name:"ExcludeRegistryImageIds"`
	// 根据本地镜像ids授权，优先权高于根据满足条件的镜像授权。AllLocalImages为false且LocalImageFilter为空和UpdatedLocalImageCnt大于0时，需要。

	LocalImageIds []*string `json:"LocalImageIds,omitempty" name:"LocalImageIds"`
	// 根据仓库镜像Ids授权，优先权高于根据满足条件的镜像授。AllRegistryImages为false且RegistryImageFilter为空和UpdatedRegistryImageCnt大于0时，需要。

	RegistryImageIds []*string `json:"RegistryImageIds,omitempty" name:"RegistryImageIds"`
	// 本地镜像是否全部授权的标识，优先权高于根据本地镜像ids授权。等于true时需UpdatedLocalImageCnt大于0。

	AllLocalImages *bool `json:"AllLocalImages,omitempty" name:"AllLocalImages"`
	// 仓库镜像是否全部授权的标识，优先权高于根据镜像ids授权。等于true时需UpdatedRegistryImageCnt大于0。

	AllRegistryImages *bool `json:"AllRegistryImages,omitempty" name:"AllRegistryImages"`
	// 指定操作授权的本地镜像数量，判断优先权最高，实际多出的镜像随机忽略，实际不足的部分也忽略。

	UpdatedLocalImageCnt *uint64 `json:"UpdatedLocalImageCnt,omitempty" name:"UpdatedLocalImageCnt"`
	// 指定操作授权的仓库镜像数量，判断优先权最高，实际多出的镜像随机忽略，实际不足的部分也忽略；

	UpdatedRegistryImageCnt *uint64 `json:"UpdatedRegistryImageCnt,omitempty" name:"UpdatedRegistryImageCnt"`
	// 是否仅最新的镜像；RegistryImageFilter不为空且UpdatedRegistryImageCnt大于0时仓库镜像需要。

	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
	// 根据满足条件的本地镜像授权,本地镜像来源；ASSETIMAGE:本地镜像列表；IMAGEALL:同步本地镜像；AllLocalImages为false且LocalImageIds为空和UpdatedLocalImageCnt大于0时，需要

	ImageSourceType *string `json:"ImageSourceType,omitempty" name:"ImageSourceType"`
}

func (r *ModifyImageAuthorizedRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageAuthorizedRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageScanStatusRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *DescribeAssetImageScanStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageScanStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogJoinObjectListRequest struct {
	*tchttp.BaseRequest

	// 日志类型
	// bash: "container_bash",
	// 启动: "container_launch",
	// "k8s": "k8s_api"

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 主机状态 </li>
	// <li>HostIP- String - 是否必填：否 - 主机内网IP </li>
	// <li>PublicIP- String - 是否必填：否 - 主机外网IP </li>
	// <li>HostName- String - 是否必填：否 - 主机名称 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeSecLogJoinObjectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogJoinObjectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanAuthorizedImageSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全部已授权的本地镜像数

		AllAuthorizedImageCount *int64 `json:"AllAuthorizedImageCount,omitempty" name:"AllAuthorizedImageCount"`
		// 已授权未扫描的本地镜像数

		UnScanAuthorizedImageCount *int64 `json:"UnScanAuthorizedImageCount,omitempty" name:"UnScanAuthorizedImageCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulScanAuthorizedImageSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanAuthorizedImageSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogVasInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecLogVasInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogVasInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulDefenceEventStatusRequest struct {
	*tchttp.BaseRequest

	// 事件IDs数组

	EventIDs []*int64 `json:"EventIDs,omitempty" name:"EventIDs"`
	// 操作状态：
	// EVENT_DEALED：已处理，EVENT_IGNORE：忽略，EVENT_ISOLATE_CONTAINER：隔离容器，EVENT_DEL：删除

	Status *string `json:"Status,omitempty" name:"Status"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyVulDefenceEventStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceEventStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServiceYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// yaml格式字符串,base64编码

		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterServiceYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServiceYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulIgnoreLocalImage struct {

	// 记录ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像大小

	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
}

type DescribeReverseShellEventsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeReverseShellEventsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateK8sApiAbnormalRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 规则详情

	RuleInfo *K8sApiAbnormalRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`
	// 拷贝规则ID(适用于复制规则场景)

	CopySrcRuleID *string `json:"CopySrcRuleID,omitempty" name:"CopySrcRuleID"`
	// 事件ID(适用于事件加白场景)

	EventID *uint64 `json:"EventID,omitempty" name:"EventID"`
}

func (r *CreateK8sApiAbnormalRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateK8sApiAbnormalRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessEventsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessEventsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceInfo struct {

	// 服务id

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 服务名 例如nginx/redis

	Type *string `json:"Type,omitempty" name:"Type"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 账号

	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
	// 监听端口

	Listen []*string `json:"Listen,omitempty" name:"Listen"`
	// 配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// 关联进程数

	ProcessCnt *uint64 `json:"ProcessCnt,omitempty" name:"ProcessCnt"`
	// 访问日志

	AccessLog *string `json:"AccessLog,omitempty" name:"AccessLog"`
	// 错误日志

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
	// 数据目录

	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`
	// web目录

	WebRoot *string `json:"WebRoot,omitempty" name:"WebRoot"`
	// 关联的进程id

	Pids []*uint64 `json:"Pids,omitempty" name:"Pids"`
	// 服务类型 app,web,db

	MainType *string `json:"MainType,omitempty" name:"MainType"`
	// 执行文件

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 服务命令行参数

	Parameter *string `json:"Parameter,omitempty" name:"Parameter"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// podip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 超级节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
}

type DescribeImageAuthorizedInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageAuthorizedInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageAuthorizedInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterRiskTarget struct {

	// 检查对象名称

	RiskTarget *string `json:"RiskTarget,omitempty" name:"RiskTarget"`
	// 严重的风险数量

	SeriousRiskCount *uint64 `json:"SeriousRiskCount,omitempty" name:"SeriousRiskCount"`
	// 高危的风险数量

	HighRiskCount *uint64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
	// 中危的风险数量

	MiddleRiskCount *uint64 `json:"MiddleRiskCount,omitempty" name:"MiddleRiskCount"`
	// 提示的风险数量

	HintRiskCount *uint64 `json:"HintRiskCount,omitempty" name:"HintRiskCount"`
}

type ModifyContainerNetStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContainerNetStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyContainerNetStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateManualUpgradeVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateManualUpgradeVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateManualUpgradeVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePolicyItemAffectedSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回各检测项影响的资产的汇总信息。

		PolicyItemSummary *CompliancePolicyItemSummary `json:"PolicyItemSummary,omitempty" name:"PolicyItemSummary"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCompliancePolicyItemAffectedSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePolicyItemAffectedSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`
		// 进程信息

		ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`
		// 被篡改信息

		TamperedFileInfo *FileAttributeInfo `json:"TamperedFileInfo,omitempty" name:"TamperedFileInfo"`
		// 事件描述

		EventDetail *AccessControlEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`
		// 父进程信息

		ParentProcessInfo *ProcessBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`
		// 祖先进程信息

		AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRulesExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessRulesExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRulesExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeRuleInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeEscapeRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportRegionsRequest struct {
	*tchttp.BaseRequest

	// 1表示获取用户资产相关的地域列表。
	// 2表示获取所有支持的地域列表。

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSupportRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSupportRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusManualScanEstimateTimeoutRequest struct {
	*tchttp.BaseRequest

	// 扫描范围0容器1主机节点

	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`
	// true 全选，false 自选

	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`
	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定

	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`
}

func (r *DescribeVirusManualScanEstimateTimeoutRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusManualScanEstimateTimeoutRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulDefenceSettingRequest struct {
	*tchttp.BaseRequest

	// 是否开启:0: 关闭 1:开启

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 漏洞防御主机范围:0：自选 1: 全部主机

	Scope *int64 `json:"Scope,omitempty" name:"Scope"`
	// 自选漏洞防御主机

	HostIDs []*string `json:"HostIDs,omitempty" name:"HostIDs"`
	// 漏洞防御超级节点范围:0：自选 1: 全部

	SuperScope *int64 `json:"SuperScope,omitempty" name:"SuperScope"`
	// 超级节点Id列表

	NodeIds []*string `json:"NodeIds,omitempty" name:"NodeIds"`
}

func (r *ModifyVulDefenceSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageComponentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像组件列表

		List []*ImageComponent `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageComponentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageComponentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageAutoAuthorizedRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则是否生效，0:不生效，1:已生效

		IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
		// 授权范围类别，MANUAL:自选主机节点，ALL:全部镜像

		RangeType *string `json:"RangeType,omitempty" name:"RangeType"`
		// 授权范围是自选主机时的主机数量

		HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
		// 每天最大的镜像授权数限制, 0表示无限制

		MaxDailyCount *int64 `json:"MaxDailyCount,omitempty" name:"MaxDailyCount"`
		// 规则id，用未设置时为0

		RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
		// 自动扫描开关，0：关闭，1：开启

		AutoScanEnabled *int64 `json:"AutoScanEnabled,omitempty" name:"AutoScanEnabled"`
		// 自动扫描范围

		ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageAutoAuthorizedRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageAutoAuthorizedRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageDenyRuleExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType- String - 是否必填：否 -规则类型 RULE_RISK：风险， RULE_PRIVILEGE：特权</li>
	// <li>EffectStatus- String - 是否必填：否 - 生效状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中。</li>
	// <li>RuleName- string - 是否必填：否 - 规则名称。</li>
	// <li>Status- string - 是否必填：否 - 开启状态 0：开启，1：关闭。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：生效时间：EffectTime，更新时间：UpdateTime，Status：启动状态

	By *string `json:"By,omitempty" name:"By"`
	// 置顶已开启规则 true：是 ，否：false

	TopTurnOn *bool `json:"TopTurnOn,omitempty" name:"TopTurnOn"`
}

func (r *CreateImageDenyRuleExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageDenyRuleExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

func (r *DescribeK8sApiAbnormalRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalRuleScopeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 列表

		List []*K8sApiAbnormalRuleScopeInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalRuleScopeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleScopeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageBindRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"EventType","Values":[""]}]
	// EventType取值：
	// "FILE_ABNORMAL_READ" 访问控制
	// "MALICE_PROCESS_START" 恶意进程启动

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAssetImageBindRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageBindRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryNamespaceListRequest struct {
	*tchttp.BaseRequest

	// 本次查询的起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本次查询的数据量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询的过滤条件。Name字段可取值"Namespace"。

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeImageRegistryNamespaceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryNamespaceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellSystemPolicyConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeReverseShellSystemPolicyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellSystemPolicyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAssetImageAuthInfoRequest struct {
	*tchttp.BaseRequest

	// 镜像ids

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 是否全部未授权镜像

	All *bool `json:"All,omitempty" name:"All"`
	// 1标记已授权 0取消授权

	IsAuthed *uint64 `json:"IsAuthed,omitempty" name:"IsAuthed"`
}

func (r *UpdateAssetImageAuthInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAssetImageAuthInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRiskSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全漏洞

		VulnerabilityCnt []*RunTimeRiskInfo `json:"VulnerabilityCnt,omitempty" name:"VulnerabilityCnt"`
		// 木马病毒

		MalwareVirusCnt []*RunTimeRiskInfo `json:"MalwareVirusCnt,omitempty" name:"MalwareVirusCnt"`
		// 敏感信息

		RiskCnt []*RunTimeRiskInfo `json:"RiskCnt,omitempty" name:"RiskCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRiskSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRiskSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportComplianceStatusListJobRequest struct {
	*tchttp.BaseRequest

	// 要导出信息的资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 按照检测项导出，还是按照资产导出。true: 按照资产导出；false: 按照检测项导出。

	ExportByAsset *bool `json:"ExportByAsset,omitempty" name:"ExportByAsset"`
	// 要导出的资产ID列表或检测项ID列表，由ExportByAsset的取值决定。

	IdList []*uint64 `json:"IdList,omitempty" name:"IdList"`
	// true, 全部导出；false, 根据IdList来导出数据。

	ExportAll *bool `json:"ExportAll,omitempty" name:"ExportAll"`
}

func (r *CreateExportComplianceStatusListJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportComplianceStatusListJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalRuleScopeInfo struct {

	// 范围
	// 系统事件:
	// ANONYMOUS_ACCESS: 匿名访问
	// ABNORMAL_UA_REQ: 异常UA请求
	// ANONYMOUS_ABNORMAL_PERMISSION: 匿名用户权限异动
	// GET_CREDENTIALS: 凭据信息获取
	// MOUNT_SENSITIVE_PATH: 敏感路径挂载
	// COMMAND_RUN: 命令执行
	// PRIVILEGE_CONTAINER: 特权容器
	// EXCEPTION_CRONTAB_TASK: 异常定时任务
	// STATICS_POD: 静态pod创建
	// ABNORMAL_CREATE_POD: 异常pod创建
	// USER_DEFINED: 用户自定义

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 动作(RULE_MODE_ALERT: 告警 RULE_MODE_RELEASE:放行)

	Action *string `json:"Action,omitempty" name:"Action"`
	// 威胁等级 HIGH:高级 MIDDLE: 中级 LOW:低级 NOTICE:提示

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 开关状态(true:开 false:关) 适用于系统规则

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 是否被删除 适用于自定义规则入参

	IsDelete *bool `json:"IsDelete,omitempty" name:"IsDelete"`
}

type DescribeSecLogDeliveryClsSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志类型列表

		LogTypeList []*SecLogDeliveryClsSettingInfo `json:"LogTypeList,omitempty" name:"LogTypeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogDeliveryClsSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogDeliveryClsSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEscapeWhiteListExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>EventType- String - 是否必填：否 - 加白事件类型，ESCAPE_CGROUPS：利用cgroup机制逃逸，ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸， ESCAPE_DOCKER_API：访问Docker API接口逃逸，ESCAPE_VUL_OCCURRED：逃逸漏洞利用，MOUNT_SENSITIVE_PTAH：敏感路径挂载，PRIVILEGE_CONTAINER_START：特权容器， PRIVILEGE：程序提权逃逸</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：主机数量：HostCount，容器数量：ContainerCount，更新时间：UpdateTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateEscapeWhiteListExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEscapeWhiteListExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterIngressInfo struct {

	// Ingress名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 预留字段，当前不生效.

	Type *string `json:"Type,omitempty" name:"Type"`
	// VIP地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// Ingress对应的后端服务，多个以,隔开

	BackendInfo *string `json:"BackendInfo,omitempty" name:"BackendInfo"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 后端服务数量

	BackendCount *uint64 `json:"BackendCount,omitempty" name:"BackendCount"`
	// 集群状态

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
}

type RunTimeFilters struct {

	// 过滤键的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 是否模糊查询

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type VulAffectedImageInfo struct {

	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 关联的主机数

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 关联的容器数

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 组件列表

	ComponentList []*VulAffectedImageComponentInfo `json:"ComponentList,omitempty" name:"ComponentList"`
}

type CreateReverseShellWhiteListsExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateReverseShellWhiteListsExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateReverseShellWhiteListsExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskListRequest struct {
	*tchttp.BaseRequest

	// 要查询的集群ID，如果不指定，则查询用户所有的风险项

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：RiskLevel风险等级, RiskTarget检查对象，风险对象,RiskType风险类别,RiskAttribute检测项所属的风险类型,Name

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyRuleListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType- String - 是否必填：否 -规则类型 RULE_RISK：风险， RULE_PRIVILEGE：特权</li>
	// <li>EffectStatus- String - 是否必填：否 - 生效状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中。</li>
	// <li>RuleName- string - 是否必填：否 - 规则名称。</li>
	// <li>Status- string - 是否必填：否 - 开启状态 0：开启，1：关闭。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：生效时间：EffectTime，更新时间：UpdateTime

	By *string `json:"By,omitempty" name:"By"`
	// 置顶已开启规则 true：是 ，否：false

	TopTurnOn *bool `json:"TopTurnOn,omitempty" name:"TopTurnOn"`
}

func (r *DescribeImageDenyRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusScanConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVirusScanConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusScanConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryAutoAuthorizedRepoListRequest struct {
	*tchttp.BaseRequest

	// NAMESPACE/REPO

	QueryObject *string `json:"QueryObject,omitempty" name:"QueryObject"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤器

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc，desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeImageRegistryAutoAuthorizedRepoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryAutoAuthorizedRepoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchImageAutoAuthorizedRuleRequest struct {
	*tchttp.BaseRequest

	// 规则是否生效，0:不生效，1:已生效

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 规则id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *SwitchImageAutoAuthorizedRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchImageAutoAuthorizedRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Status - String - 是否必填：否 - agent状态筛选，"ALL":"全部"(或不传该字段),"UNINSTALL"："未安装","OFFLINE"："离线", "ONLINE"："防护中"</li>
	// <li>HostName - String - 是否必填：否 - 主机名筛选</li>
	// <li>Group- String - 是否必填：否 - 主机群组搜索</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>HostID- string - 是否必填：否 - 主机id搜索</li>
	// <li>DockerVersion- string - 是否必填：否 - docker版本搜索</li>
	// <li>MachineType- string - 是否必填：否 - 主机来源MachineType搜索，"ALL":"全部"(或不传该字段),主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；</li>
	// <li>DockerStatus- string - 是否必填：否 - docker安装状态，"ALL":"全部"(或不传该字段),"INSTALL":"已安装","UNINSTALL":"未安装"</li>
	// <li>ProjectID- string - 是否必填：否 - 所属项目id搜索</li>
	// <li>Tag:xxx(tag:key)- string- 是否必填：否 - 标签键值搜索 示例Filters":[{"Name":"tag:tke-kind","Values":["service"]}]</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像拦截列表

		List []*ImageDenyEvent `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunTimeTendencyInfo struct {

	// 当天时间

	CurTime *string `json:"CurTime,omitempty" name:"CurTime"`
	// 当前数量

	Cnt *uint64 `json:"Cnt,omitempty" name:"Cnt"`
}

type ComplianceAssetDetailInfo struct {

	// 客户资产的ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 资产类别。

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 资产的名称。

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产所属的节点的名称。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 资产所在的主机的名称。

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 资产所在的主机的IP。

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 检测状态
	// CHECK_INIT, 待检测
	// CHECK_RUNNING, 检测中
	// CHECK_FINISHED, 检测完成
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 此类资产通过的检测项的数目。

	PassedPolicyItemCount *uint64 `json:"PassedPolicyItemCount,omitempty" name:"PassedPolicyItemCount"`
	// 此类资产未通过的检测的数目。

	FailedPolicyItemCount *uint64 `json:"FailedPolicyItemCount,omitempty" name:"FailedPolicyItemCount"`
	// 上次检测的时间。

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 检测结果：
	// RESULT_FAILED: 未通过。
	// RESULT_PASSED: 通过。

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 资产的运行状态。

	AssetStatus *string `json:"AssetStatus,omitempty" name:"AssetStatus"`
	// 创建资产的时间。
	// ASSET_NORMAL: 正常运行，
	// ASSET_PAUSED: 暂停运行，
	// ASSET_STOPPED: 停止运行，
	// ASSET_ABNORMAL: 异常

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
}

type DescribeReverseShellWhiteListDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		WhiteListDetailInfo *ReverseShellWhiteListInfo `json:"WhiteListDetailInfo,omitempty" name:"WhiteListDetailInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellWhiteListDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellWhiteListDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellInnerIPShowStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeReverseShellInnerIPShowStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellInnerIPShowStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefenceEventTendency struct {

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
}

type DescribeClusterRiskRequest struct {
	*tchttp.BaseRequest

	// 漏洞风险CVERisk,配置风险ConfigRisk

	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
}

func (r *DescribeClusterRiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanCompliancePolicyItemsRequest struct {
	*tchttp.BaseRequest

	// 要重新扫描的客户检测项的列表。

	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

func (r *ScanCompliancePolicyItemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanCompliancePolicyItemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableNodesCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 普通节点计数信息

		NormalNodesCount *uint64 `json:"NormalNodesCount,omitempty" name:"NormalNodesCount"`
		// 超级节点计数信息

		SuperNodesCount *uint64 `json:"SuperNodesCount,omitempty" name:"SuperNodesCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableNodesCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableNodesCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskTargetCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检查对象风险数量

		ClusterRiskTargets []*ClusterRiskTarget `json:"ClusterRiskTargets,omitempty" name:"ClusterRiskTargets"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskTargetCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskTargetCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeEventsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeAgentInfo struct {

	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 内部ip

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 升级前版本

	CurrentVersion *string `json:"CurrentVersion,omitempty" name:"CurrentVersion"`
	// 升级后的版本

	UpgradeVersion *string `json:"UpgradeVersion,omitempty" name:"UpgradeVersion"`
	// 升级状态，                                                成功：SUCCESSED,
	// 失败：FAILED,
	// 升级中：UPGRADING

	Status *string `json:"Status,omitempty" name:"Status"`
	// 唯一id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
}

type DescribeNamespaceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命名空间列表

		Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespaceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegistryVulScanAuthorizedImageSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全部已授权的仓库镜像数

		AllAuthorizedImageCount *int64 `json:"AllAuthorizedImageCount,omitempty" name:"AllAuthorizedImageCount"`
		// 已授权未扫描的仓库镜像数

		UnScanAuthorizedImageCount *int64 `json:"UnScanAuthorizedImageCount,omitempty" name:"UnScanAuthorizedImageCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegistryVulScanAuthorizedImageSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegistryVulScanAuthorizedImageSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetComponentListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 容器id

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
}

func (r *DescribeAssetComponentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetComponentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyWorkLoadChartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 边总数

		EdgeListTotalCount *uint64 `json:"EdgeListTotalCount,omitempty" name:"EdgeListTotalCount"`
		// 节点总数

		NodeListTotalCount *uint64 `json:"NodeListTotalCount,omitempty" name:"NodeListTotalCount"`
		// 模糊状态：precise：准确；fuzzy：模糊；

		Fuzzy *string `json:"Fuzzy,omitempty" name:"Fuzzy"`
		// 边信息

		EdgeList *string `json:"EdgeList,omitempty" name:"EdgeList"`
		// 节点信息

		NodeList *string `json:"NodeList,omitempty" name:"NodeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkTopologyWorkLoadChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyWorkLoadChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrialStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否可以申请试用，true可以申请

		CanApplyTrial *bool `json:"CanApplyTrial,omitempty" name:"CanApplyTrial"`
		// 无法试用原因，可试用为空；eg：不是企业用户，不符合试用条件

		NotApplyReason *string `json:"NotApplyReason,omitempty" name:"NotApplyReason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrialStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrialStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusAutoIsolateSettingRequest struct {
	*tchttp.BaseRequest

	// 自动隔离开关(true:开 false:关)

	AutoIsolateSwitch *bool `json:"AutoIsolateSwitch,omitempty" name:"AutoIsolateSwitch"`
	// 是否中断隔离文件关联的进程(true:是 false:否)

	IsKillProgress *bool `json:"IsKillProgress,omitempty" name:"IsKillProgress"`
}

func (r *ModifyVirusAutoIsolateSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusAutoIsolateSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageRegistryScanTaskOneKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 扫描任务id

		TaskID *uint64 `json:"TaskID,omitempty" name:"TaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetImageRegistryScanTaskOneKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageRegistryScanTaskOneKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessEventDescription struct {

	// 事件规则

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 命中规则详细信息

	MatchRule *AbnormalProcessChildRuleInfo `json:"MatchRule,omitempty" name:"MatchRule"`
	// 命中规则名称，PROXY_TOOL：代理软件，TRANSFER_CONTROL：横向渗透，ATTACK_CMD：恶意命令，REVERSE_SHELL：反弹shell，FILELESS：无文件程序执行，RISK_CMD：高危命令，ABNORMAL_CHILD_PROC：敏感服务异常子进程启动，USER_DEFINED_RULE：用户自定义规则

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 命中规则的id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 事件最后一次处理的时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
	// 命中策略名称：SYSTEM_DEFINED_RULE （系统策略）或  用户自定义的策略名字

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type MaliciousConnectionRuleInfo struct {

	// 枚举：
	// IP: 表示ipv4或者ipv6
	// DOMAIN: 表示域名

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 自定义黑白名单的域名/IP

	Address *string `json:"Address,omitempty" name:"Address"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 规则ID

	RuleID *uint64 `json:"RuleID,omitempty" name:"RuleID"`
}

type DeleteMaliciousConnectionWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单记录ID数组

	IDSet []*int64 `json:"IDSet,omitempty" name:"IDSet"`
}

func (r *DeleteMaliciousConnectionWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaliciousConnectionWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallPodLabelsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群pod总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群pod详细信息

		PodList []*NetworkClusterPodInfo `json:"PodList,omitempty" name:"PodList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallPodLabelsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPodLabelsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAbnormalProcessRulesExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateAbnormalProcessRulesExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAbnormalProcessRulesExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SupportDefenceVul struct {

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 漏洞标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 漏洞CVSS

	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`
	// 漏洞威胁等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 漏洞CVEID

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞披露时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
}

type DescribeRefreshTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 新任务ID

	NewTaskID *string `json:"NewTaskID,omitempty" name:"NewTaskID"`
}

func (r *DescribeRefreshTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRefreshTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodeWhiteListRequest struct {
	*tchttp.BaseRequest

	// 检测项ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：InstanceId实例id,InstanceRole节点的角色，Master、Work等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClusterNodeWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodeWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PortInfo struct {

	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 对外ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 主机端口

	PublicPort *uint64 `json:"PublicPort,omitempty" name:"PublicPort"`
	// 容器端口

	ContainerPort *uint64 `json:"ContainerPort,omitempty" name:"ContainerPort"`
	// 容器Pid

	ContainerPID *uint64 `json:"ContainerPID,omitempty" name:"ContainerPID"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 容器内监听地址

	ListenContainer *string `json:"ListenContainer,omitempty" name:"ListenContainer"`
	// 容器外监听地址

	ListenHost *string `json:"ListenHost,omitempty" name:"ListenHost"`
	// 运行账号

	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// podip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 超级节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
}

type WhiteListWorkload struct {

	// 加入白名单的时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 白名单的workloadId

	WorkloadId *uint64 `json:"WorkloadId,omitempty" name:"WorkloadId"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 工作负载名称

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 工作负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeValueAddedSrvInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeValueAddedSrvInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeValueAddedSrvInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTcrRegistryDetailRequest struct {
	*tchttp.BaseRequest

	// 仓库ID

	RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`
	// 网络类型，列表：public（公网） vpc（内网）

	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

func (r *UpdateTcrRegistryDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTcrRegistryDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetClusterCheckIdWhiteListStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加检测项白名单的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetClusterCheckIdWhiteListStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetClusterCheckIdWhiteListStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrModifyPostPayCoresRequest struct {
	*tchttp.BaseRequest

	// 弹性计费上限，最小值500

	CoresCnt *uint64 `json:"CoresCnt,omitempty" name:"CoresCnt"`
}

func (r *CreateOrModifyPostPayCoresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrModifyPostPayCoresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalRuleInfo struct {

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 规则信息列表

	RuleInfoList []*K8sApiAbnormalRuleScopeInfo `json:"RuleInfoList,omitempty" name:"RuleInfoList"`
	// 生效集群IDSet

	EffectClusterIDSet []*string `json:"EffectClusterIDSet,omitempty" name:"EffectClusterIDSet"`
	// 规则类型
	// RT_SYSTEM 系统规则
	// RT_USER 用户自定义

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 是否所有集群生效

	EffectAllCluster *bool `json:"EffectAllCluster,omitempty" name:"EffectAllCluster"`
}

type DescribeNewestVulResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞PocID

		PocID *string `json:"PocID,omitempty" name:"PocID"`
		// 漏洞名称

		VulName *string `json:"VulName,omitempty" name:"VulName"`
		// 披露时间

		SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
		// 应急漏洞风险情况：NOT_SCAN：未扫描，SCANNING：扫描中，SCANNED：已扫描

		Status *string `json:"Status,omitempty" name:"Status"`
		// 漏洞CVEID

		CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNewestVulResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewestVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogKafkaUINResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 目标UIN

		DstUIN *string `json:"DstUIN,omitempty" name:"DstUIN"`
		// 授权状态

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogKafkaUINResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogKafkaUINResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskSyscallWhiteListInfo struct {

	// 系统调用名称，通过DescribeRiskSyscallNames接口获取枚举列表

	SyscallNames []*string `json:"SyscallNames,omitempty" name:"SyscallNames"`
	// 目标进程

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 镜像id数组，为空代表全部

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 白名单id，如果新建则id为空

	Id *string `json:"Id,omitempty" name:"Id"`
}

type CreateClusterAssetExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要导出的集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateClusterAssetExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterAssetExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteComplianceAssetPolicySetFromWhitelistRequest struct {
	*tchttp.BaseRequest

	// 资产ID

	AssetItemId *uint64 `json:"AssetItemId,omitempty" name:"AssetItemId"`
	// 需要忽略指定资产内的检查项ID列表

	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

func (r *DeleteComplianceAssetPolicySetFromWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteComplianceAssetPolicySetFromWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventTypeSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容器逃逸事件数

		ContainerEscapeEventCount *int64 `json:"ContainerEscapeEventCount,omitempty" name:"ContainerEscapeEventCount"`
		// 程序提权事件数

		ProcessPrivilegeEventCount *int64 `json:"ProcessPrivilegeEventCount,omitempty" name:"ProcessPrivilegeEventCount"`
		// 风险容器事件数

		RiskContainerEventCount *int64 `json:"RiskContainerEventCount,omitempty" name:"RiskContainerEventCount"`
		// 逃逸事件待处理数

		PendingEscapeEventCount *int64 `json:"PendingEscapeEventCount,omitempty" name:"PendingEscapeEventCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeEventTypeSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventTypeSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecLogAlertMsgInfo struct {

	// 告警类型

	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`
	// 告警值

	MsgValue *string `json:"MsgValue,omitempty" name:"MsgValue"`
	// 状态(0:关闭 1:开启)

	State *bool `json:"State,omitempty" name:"State"`
}

type DescribeVulLevelSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 高危漏洞数

		HighLevelVulCount *int64 `json:"HighLevelVulCount,omitempty" name:"HighLevelVulCount"`
		// 中危漏洞数

		MediumLevelVulCount *int64 `json:"MediumLevelVulCount,omitempty" name:"MediumLevelVulCount"`
		// 低危漏洞数

		LowLevelVulCount *int64 `json:"LowLevelVulCount,omitempty" name:"LowLevelVulCount"`
		// 严重漏洞数

		CriticalLevelVulCount *int64 `json:"CriticalLevelVulCount,omitempty" name:"CriticalLevelVulCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulLevelSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulLevelSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 连接错误信息

		HealthCheckErr *string `json:"HealthCheckErr,omitempty" name:"HealthCheckErr"`
		// 名称错误信息

		NameRepeatErr *string `json:"NameRepeatErr,omitempty" name:"NameRepeatErr"`
		// 仓库唯一id

		RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAssetImageRegistryRegistryDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAssetImageRegistryRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulImageExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulImageExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulImageExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyEventListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>EventType- String - 是否必填：否 -事件类型 EVENT_RISK:风险事件类型，EVENT_PRIVILEGE:特权。</li>
	// <li>DealBehavior- String - 是否必填：否 - 执行动作,BEHAVIOR_ALERT:告警，BEHAVIOR_HOLDUP_SUCCESSED:拦截。</li>
	// <li>RuleName- string - 是否必填：否 - 规则名称。</li>
	// <li>NodeName- string - 是否必填：否 - 节点名称。</li>
	// <li>NodeIP- string - 是否必填：否 - 内外IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：告警数量：EventCount，最近生成时间：LatestFoundTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeImageDenyEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalTendencyRequest struct {
	*tchttp.BaseRequest

	// 趋势周期(默认为7天)

	TendencyPeriod *uint64 `json:"TendencyPeriod,omitempty" name:"TendencyPeriod"`
}

func (r *DescribeK8sApiAbnormalTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallPolicyListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeNetworkFirewallPolicyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPolicyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterCheckTimerSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 集群ID列表

		ClusterIDs []*string `json:"ClusterIDs,omitempty" name:"ClusterIDs"`
		// 周期天数 若是1 则表示每天一扫 若是3 则表示每三天一扫

		CycleDay *uint64 `json:"CycleDay,omitempty" name:"CycleDay"`
		// 扫描开始时间

		ScanTimeBegin *string `json:"ScanTimeBegin,omitempty" name:"ScanTimeBegin"`
		// 扫描截止时间

		ScanTimeEnd *string `json:"ScanTimeEnd,omitempty" name:"ScanTimeEnd"`
		// 扫描范围 0:全部集群 1:部分集群

		ScanScope *uint64 `json:"ScanScope,omitempty" name:"ScanScope"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterCheckTimerSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterCheckTimerSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecLogJoinSuperObjectInfo struct {

	// 超级节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 超级节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 所属集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 所属集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 节点状态:Running,Ready,Notready,Initializing,Failed,Error

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网ID

	SubNetID *string `json:"SubNetID,omitempty" name:"SubNetID"`
	// 子网名称

	SubNetName *string `json:"SubNetName,omitempty" name:"SubNetName"`
	// 子网网段

	SubNetCidr *string `json:"SubNetCidr,omitempty" name:"SubNetCidr"`
	// 关联pod数

	RelatePodCount *uint64 `json:"RelatePodCount,omitempty" name:"RelatePodCount"`
	// 关联容器数

	RelateContainerCount *uint64 `json:"RelateContainerCount,omitempty" name:"RelateContainerCount"`
	// 节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 接入状态(true:已接入 false:未接入)

	JoinState *bool `json:"JoinState,omitempty" name:"JoinState"`
}

type EscapeRuleEnabled struct {

	// 规则类型
	//    ESCAPE_HOST_ACESS_FILE:宿主机文件访问逃逸
	//    ESCAPE_MOUNT_NAMESPACE:MountNamespace逃逸
	//    ESCAPE_PRIVILEDGE:程序提权逃逸
	//    ESCAPE_PRIVILEDGE_CONTAINER_START:特权容器启动逃逸
	//    ESCAPE_MOUNT_SENSITIVE_PTAH:敏感路径挂载
	//    ESCAPE_SYSCALL:Syscall逃逸

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否打开：false否 ，true是

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
}

type CreateEventEscapeImageExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEventEscapeImageExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEventEscapeImageExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetSimpleImageInfo struct {

	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 关联容器个数

	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 最后扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 镜像大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type ModifyVirusScanTimeoutSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVirusScanTimeoutSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusScanTimeoutSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVirusScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVirusScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVirusScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryNamespaceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可返回的命令空间的总量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的命令空间列表

		NamespaceList []*string `json:"NamespaceList,omitempty" name:"NamespaceList"`
		// 返回的命令空间详细信息列表

		NamespaceDetail []*NamespaceInfo `json:"NamespaceDetail,omitempty" name:"NamespaceDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRegistryNamespaceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryNamespaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddClusterNodeToWhiteListRequest struct {
	*tchttp.BaseRequest

	// 受影响的节点信息

	AffectedNodeInfo *AffectedNodeItem `json:"AffectedNodeInfo,omitempty" name:"AffectedNodeInfo"`
	// 检测项ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
}

func (r *AddClusterNodeToWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterNodeToWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerAssetSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeContainerAssetSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerAssetSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserPublicClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群类型:Openshift或者Kubernetes

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 是否自动安装防御容器

	InstallDefender *bool `json:"InstallDefender,omitempty" name:"InstallDefender"`
	// 是否自动检查

	ClusterAutoCheck *bool `json:"ClusterAutoCheck,omitempty" name:"ClusterAutoCheck"`
	// 配置文件

	KubeConfig *string `json:"KubeConfig,omitempty" name:"KubeConfig"`
	// 地域

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
	// 不填或者为空表示公网接入,非空表示VPC接入

	VPCId *string `json:"VPCId,omitempty" name:"VPCId"`
	// VPCId为非空时,该项必填,为CVM或者LB

	ApiServerType *string `json:"ApiServerType,omitempty" name:"ApiServerType"`
	// 混合云传入集群Id，其它类型的留空

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateUserPublicClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserPublicClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefenceHost struct {

	// 主机名称/超级节点名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机ip即内网ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机QUUID/超级节点ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 插件状态，正常：SUCCESS，异常：FAIL， NO_DEFENDED:未防御

	Status *string `json:"Status,omitempty" name:"Status"`
	// 外网ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 首次开启时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 节点类型[NORMAL:普通节点|SUPER:超级节点]

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 超级节点子网名称

	NodeSubNetName *string `json:"NodeSubNetName,omitempty" name:"NodeSubNetName"`
	// 超级节点子网网段

	NodeSubNetCIDR *string `json:"NodeSubNetCIDR,omitempty" name:"NodeSubNetCIDR"`
	// 超级节点子网ID

	NodeSubNetID *string `json:"NodeSubNetID,omitempty" name:"NodeSubNetID"`
	// 超级节点唯一ID

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 超级节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// Pod Ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// Pod 名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
}

type DescribeReverseShellRegexpWhiteListInfoRequest struct {
	*tchttp.BaseRequest

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

func (r *DescribeReverseShellRegexpWhiteListInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellRegexpWhiteListInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRegistryConnDetectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 联通性检测结果

		ConnDetectResult []*RegistryConnDetectResult `json:"ConnDetectResult,omitempty" name:"ConnDetectResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageRegistryConnDetectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImageRegistryConnDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddImageDenyRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否存在特权拦截规则

		ExistPrivilegeDenyRule *bool `json:"ExistPrivilegeDenyRule,omitempty" name:"ExistPrivilegeDenyRule"`
		// 已扫描的镜像数量

		ScannedImageCount *int64 `json:"ScannedImageCount,omitempty" name:"ScannedImageCount"`
		// 特权拦截规则RuleID

		PrivilegeDenyRuleID *string `json:"PrivilegeDenyRuleID,omitempty" name:"PrivilegeDenyRuleID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddImageDenyRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddImageDenyRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEscapeEventsExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEscapeEventsExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEscapeEventsExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceImageDetailInfo struct {

	// 镜像在主机上的ID。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像的名称。

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像的Tag。

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 镜像所在远程仓库的路径。

	Repository *string `json:"Repository,omitempty" name:"Repository"`
}

type DescribeRiskSyscallEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 高危系统调用数组

		EventSet []*RiskSyscallEventInfo `json:"EventSet,omitempty" name:"EventSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIACCheckTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态，可能为：Task_Running,Task_Succ,Task_Error,Task_NoExist

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 任务结果详情

		TaskResult []*string `json:"TaskResult,omitempty" name:"TaskResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIACCheckTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIACCheckTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CleanTrialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 申请结果，0成功，其他失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 试用失败原因，成功可为空；eg：请勿重复申请

		Reason *string `json:"Reason,omitempty" name:"Reason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CleanTrialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CleanTrialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVulListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 漏洞名称名称筛选，</li>
	// <li>Level - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *DescribeAssetImageVulListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVulListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessControlRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccessControlRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessControlRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceIgnoredAsset struct {

	// 资产项的ID，如容器id。

	CustomerAssetId *string `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 资产项的名称。

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产项的类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 为客户分配的唯一的资产项的ID。

	AssetId *uint64 `json:"AssetId,omitempty" name:"AssetId"`
	// 资产状态，
	// CONTAINER_UNKNOWN  //未知状态，处于三种状态之外
	// 	CONTAINER_EXIT     //容器处于退出状态
	// 	CONTAINER_RUNNING  //容器处于正常运行状态
	// 	CONTAINER_PAUSE    //容器处于暂停状态
	// 	CONTAINER_ERROR   //查询容器状态出错

	AssetStatus *string `json:"AssetStatus,omitempty" name:"AssetStatus"`
	// 节点名称。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 主机IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 镜像的tag

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
}

type AddEditAccessControlRuleRequest struct {
	*tchttp.BaseRequest

	// 增加策略信息，策略id为空，编辑策略是id不能为空

	RuleInfo *AccessControlRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`
	// 仅在白名单时候使用

	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *AddEditAccessControlRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditAccessControlRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmNetworkFirewallPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建确认任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmNetworkFirewallPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmNetworkFirewallPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallNamespaceListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeNetworkFirewallNamespaceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallNamespaceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReverseShellSystemPolicyConfigRequest struct {
	*tchttp.BaseRequest

	// 内网告警展示

	InnerNetAlarmShow *bool `json:"InnerNetAlarmShow,omitempty" name:"InnerNetAlarmShow"`
}

func (r *ModifyReverseShellSystemPolicyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReverseShellSystemPolicyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscapeEventInfo struct {

	// 事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 状态，EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件记录的唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// pod(实例)的名字

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 事件名字，
	// 宿主机文件访问逃逸、
	// Syscall逃逸、
	// MountNamespace逃逸、
	// 程序提权逃逸、
	// 特权容器启动逃逸、
	// 敏感路径挂载

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 镜像id，用于跳转

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器id，用于跳转

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 事件解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 节点IP

	NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`
	// 主机IP

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 节点所属集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 节点类型：NORMAL普通节点、SUPER超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// pod ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 节点公网ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 节点内网ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type DescribeIACCheckResultListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeIACCheckResultListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIACCheckResultListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallWhiteListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 白名单信息列表

		WhiteListSet []*RiskSyscallWhiteListBaseInfo `json:"WhiteListSet,omitempty" name:"WhiteListSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallWhiteListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulTopRankingInfo struct {

	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
	// 威胁等级,CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低

	Level *string `json:"Level,omitempty" name:"Level"`
	// 影响的镜像数

	AffectedImageCount *int64 `json:"AffectedImageCount,omitempty" name:"AffectedImageCount"`
	// 影响的容器数

	AffectedContainerCount *int64 `json:"AffectedContainerCount,omitempty" name:"AffectedContainerCount"`
	// 漏洞ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
}

type DescribeVirusDetailRequest struct {
	*tchttp.BaseRequest

	// 木马文件id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeVirusDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeK8sApiAbnormalSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageAutoAuthorizedTask struct {

	// 任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 授权方式，AUTO:自动授权，MANUAL:手动授权

	Type *string `json:"Type,omitempty" name:"Type"`
	// 任务日期

	AuthorizedDate *string `json:"AuthorizedDate,omitempty" name:"AuthorizedDate"`
	// 镜像来源，LOCAL:本地镜像，REGISTRY:仓库镜像

	Source *string `json:"Source,omitempty" name:"Source"`
	// 最近授权时间

	LastAuthorizedTime *string `json:"LastAuthorizedTime,omitempty" name:"LastAuthorizedTime"`
	// 自动授权成功数

	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
	// 自动授权失败数

	FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
	// 最近任务失败码，REACH_LIMIT:达到授权上限，LICENSE_INSUFFICIENT:授权数不足

	LatestFailCode *string `json:"LatestFailCode,omitempty" name:"LatestFailCode"`
}

type VulAffectedImageComponentInfo struct {

	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 组件修复版本

	FixedVersion *string `json:"FixedVersion,omitempty" name:"FixedVersion"`
	// 组件路径

	Path *string `json:"Path,omitempty" name:"Path"`
}

type DescribeVirusSampleDownloadUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 样本下载地址

		FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusSampleDownloadUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusSampleDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetImageScanStopRequest struct {
	*tchttp.BaseRequest

	// 任务id；任务id，镜像id和根据过滤条件筛选三选一。

	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
	// 镜像id；任务id，镜像id和根据过滤条件筛选三选一。

	Images []*string `json:"Images,omitempty" name:"Images"`
	// 根据过滤条件筛选出镜像；任务id，镜像id和根据过滤条件筛选三选一。

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 根据过滤条件筛选出镜像，再排除个别镜像

	ExcludeImageIds *string `json:"ExcludeImageIds,omitempty" name:"ExcludeImageIds"`
}

func (r *ModifyAssetImageScanStopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetImageScanStopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyFlowSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeNetworkTopologyFlowSwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyFlowSwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogDeliveryKafkaOptionsRequest struct {
	*tchttp.BaseRequest

	// 地域，若为空则返回所有可选地域

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
}

func (r *DescribeSecLogDeliveryKafkaOptionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogDeliveryKafkaOptionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetProcessListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口列表

		List []*ProcessInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetProcessListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEventEscapeImageExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数:
	// EventType: 事件类型(MOUNT_SENSITIVE_PTAH:敏感路径挂载 PRIVILEGE_CONTAINER_START:特权容器)
	// ImageID: 镜像id
	// ImageName:镜像名称

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：latest_found_time

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateEventEscapeImageExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEventEscapeImageExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用个数

		AppCnt *uint64 `json:"AppCnt,omitempty" name:"AppCnt"`
		// 容器个数

		ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
		// 暂停的容器个数

		ContainerPause *uint64 `json:"ContainerPause,omitempty" name:"ContainerPause"`
		// 运行的容器个数

		ContainerRunning *uint64 `json:"ContainerRunning,omitempty" name:"ContainerRunning"`
		// 停止运行的容器个数

		ContainerStop *uint64 `json:"ContainerStop,omitempty" name:"ContainerStop"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 数据库个数

		DbCnt *uint64 `json:"DbCnt,omitempty" name:"DbCnt"`
		// 镜像个数

		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
		// 主机在线个数

		HostOnline *uint64 `json:"HostOnline,omitempty" name:"HostOnline"`
		// 主机个数

		HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`
		// 有风险的镜像个数

		ImageHasRiskInfoCnt *uint64 `json:"ImageHasRiskInfoCnt,omitempty" name:"ImageHasRiskInfoCnt"`
		// 有病毒的镜像个数

		ImageHasVirusCnt *uint64 `json:"ImageHasVirusCnt,omitempty" name:"ImageHasVirusCnt"`
		// 有漏洞的镜像个数

		ImageHasVulsCnt *uint64 `json:"ImageHasVulsCnt,omitempty" name:"ImageHasVulsCnt"`
		// 不受信任的镜像个数

		ImageUntrustCnt *uint64 `json:"ImageUntrustCnt,omitempty" name:"ImageUntrustCnt"`
		// 监听的端口个数

		ListenPortCnt *uint64 `json:"ListenPortCnt,omitempty" name:"ListenPortCnt"`
		// 进程个数

		ProcessCnt *uint64 `json:"ProcessCnt,omitempty" name:"ProcessCnt"`
		// web服务个数

		WebServiceCnt *uint64 `json:"WebServiceCnt,omitempty" name:"WebServiceCnt"`
		// 最近镜像扫描时间

		LatestImageScanTime *string `json:"LatestImageScanTime,omitempty" name:"LatestImageScanTime"`
		// 风险镜像个数

		ImageUnsafeCnt *uint64 `json:"ImageUnsafeCnt,omitempty" name:"ImageUnsafeCnt"`
		// 主机未安装agent数量

		HostUnInstallCnt *uint64 `json:"HostUnInstallCnt,omitempty" name:"HostUnInstallCnt"`
		// 超级节点个数

		SuperNodeCnt *uint64 `json:"SuperNodeCnt,omitempty" name:"SuperNodeCnt"`
		// 超级节点运行个数

		SuperNodeRunningCnt *uint64 `json:"SuperNodeRunningCnt,omitempty" name:"SuperNodeRunningCnt"`
		// 今日新增镜像个数

		TodayNewImageCnt *uint64 `json:"TodayNewImageCnt,omitempty" name:"TodayNewImageCnt"`
		// 今日新增风险镜像个数

		TodayUnsafeImageCnt *uint64 `json:"TodayUnsafeImageCnt,omitempty" name:"TodayUnsafeImageCnt"`
		// 推荐处置镜像个数

		RecommendedFixImageCnt *uint64 `json:"RecommendedFixImageCnt,omitempty" name:"RecommendedFixImageCnt"`
		// 已扫描镜像个数

		ScannedImageCnt *uint64 `json:"ScannedImageCnt,omitempty" name:"ScannedImageCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImagesInfo struct {

	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 镜像大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 主机个数

	HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`
	// 容器个数

	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 漏洞个数

	VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`
	// 病毒个数

	VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`
	// 敏感信息个数

	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
	// 是否信任镜像

	IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`
	// 镜像系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// agent镜像扫描错误

	AgentError *string `json:"AgentError,omitempty" name:"AgentError"`
	// 后端镜像扫描错误

	ScanError *string `json:"ScanError,omitempty" name:"ScanError"`
	// 扫描状态

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 木马扫描错误信息

	ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`
	// 漏洞扫描错误信息

	ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`
	// 风险扫描错误信息

	ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`
	// 是否是重点关注镜像，为0不是，非0是

	IsSuggest *uint64 `json:"IsSuggest,omitempty" name:"IsSuggest"`
	// 是否授权，1是0否

	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
	// 组件个数

	ComponentCnt *uint64 `json:"ComponentCnt,omitempty" name:"ComponentCnt"`
	// 严重漏洞数

	CriticalLevelVulCnt *uint64 `json:"CriticalLevelVulCnt,omitempty" name:"CriticalLevelVulCnt"`
	// 高危漏洞数

	HighLevelVulCnt *uint64 `json:"HighLevelVulCnt,omitempty" name:"HighLevelVulCnt"`
	// 中危漏洞数

	MediumLevelVulCnt *uint64 `json:"MediumLevelVulCnt,omitempty" name:"MediumLevelVulCnt"`
	// 低危漏洞数

	LowLevelVulCnt *uint64 `json:"LowLevelVulCnt,omitempty" name:"LowLevelVulCnt"`
	// 是否最新版本镜像

	IsLatestImage *bool `json:"IsLatestImage,omitempty" name:"IsLatestImage"`
	// 是否推荐处置

	RecommendedFix *bool `json:"RecommendedFix,omitempty" name:"RecommendedFix"`
}

type DescribeVulContainerListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容器列表

		List []*VulAffectedContainerInfo `json:"List,omitempty" name:"List"`
		// 容器总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulContainerListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulContainerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageAutoAuthorizedLogListRequest struct {
	*tchttp.BaseRequest

	// 自动授权任务Id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// Status授权结果，SUCCESS:成功，REACH_LIMIT:达到授权上限，LICENSE_INSUFFICIENT:授权数不足

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段：AuthorizedTime

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式，asc，desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeImageAutoAuthorizedLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageAutoAuthorizedLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceScanFailedAsset struct {

	// 客户资产的ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 资产类别。

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 检测状态
	// CHECK_INIT, 待检测
	// CHECK_RUNNING, 检测中
	// CHECK_FINISHED, 检测完成
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 资产的名称。

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产检测失败的原因。

	FailureReason *string `json:"FailureReason,omitempty" name:"FailureReason"`
	// 检测失败的处理建议。

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 检测的时间。

	CheckTime *string `json:"CheckTime,omitempty" name:"CheckTime"`
}

type VulIgnoreRegistryImage struct {

	// 记录ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 仓库名称

	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`
	// 镜像版本

	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
	// 仓库地址

	RegistryPath *string `json:"RegistryPath,omitempty" name:"RegistryPath"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
}

type ModifyVirusAutoIsolateExampleSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVirusAutoIsolateExampleSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusAutoIsolateExampleSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallDetailRequest struct {
	*tchttp.BaseRequest

	// 事件唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeRiskSyscallDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusScanSettingRequest struct {
	*tchttp.BaseRequest

	// 是否开启定期扫描

	EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`
	// 检测周期每隔多少天(1|3|7)

	Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`
	// 扫描开始时间

	BeginScanAt *string `json:"BeginScanAt,omitempty" name:"BeginScanAt"`
	// 扫描全部路径(true:全选,false:自选)

	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`
	// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径

	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`
	// 超时时长(5~24h)

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 扫描范围0容器1主机节点

	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`
	// true 全选，false 自选

	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`
	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定

	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`
	// 扫描路径

	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
	// 扫描路径模式：
	// SCAN_PATH_ALL：全部路径
	// SCAN_PATH_DEFAULT：默认路径
	// SCAN_PATH_USER_DEFINE：用户自定义路径
	//

	ScanPathMode *string `json:"ScanPathMode,omitempty" name:"ScanPathMode"`
}

func (r *ModifyVirusScanSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusScanSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterNodeInfo struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 内网ip地址

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 节点的角色，Master、Work等

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// 实例的状态（running 运行中，initializing 初始化中，failed 异常）

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// agent安装状态

	AgentStatus *string `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// 公网ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机类型(普通节点情况)

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 节点类型(
	// NORMAL: 普通节点
	// SUPER:超级节点
	// )

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 计费核数

	ChargeCoresCnt *uint64 `json:"ChargeCoresCnt,omitempty" name:"ChargeCoresCnt"`
	// 防护状态:
	// 已防护: Defended
	// 未防护: UnDefended

	DefendStatus *string `json:"DefendStatus,omitempty" name:"DefendStatus"`
}

type ModifyVirusFileStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVirusFileStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusFileStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddImageDenyRuleInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAddImageDenyRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddImageDenyRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddModReverseShellRegexpWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddModReverseShellRegexpWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModReverseShellRegexpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 入站规则
	//
	// 全部允许：1
	//
	// 全部拒绝 ：2
	//
	// 自定义：3

	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`
	// 出站规则
	//
	// 全部允许：1
	//
	// 全部拒绝 ：2
	//
	// 自定义：3

	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`
	// pod选择器

	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 自定义规则

	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

func (r *UpdateNetworkFirewallPolicyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEditWarningRulesRequest struct {
	*tchttp.BaseRequest

	// 告警开关策略

	WarningRules []*WarningRule `json:"WarningRules,omitempty" name:"WarningRules"`
}

func (r *AddEditWarningRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditWarningRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceAssetInfo struct {

	// 客户资产的ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 资产类别。

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 资产的名称。

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 当资产为镜像时，这个字段为镜像Tag。

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 资产所在的主机IP。

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 资产所属的节点的名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 检测状态
	//
	// CHECK_INIT, 待检测
	//
	// CHECK_RUNNING, 检测中
	//
	// CHECK_FINISHED, 检测完成
	//
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 此类资产通过的检测项的数目。

	PassedPolicyItemCount *uint64 `json:"PassedPolicyItemCount,omitempty" name:"PassedPolicyItemCount"`
	// 此类资产未通过的检测的数目。

	FailedPolicyItemCount *uint64 `json:"FailedPolicyItemCount,omitempty" name:"FailedPolicyItemCount"`
	// 上次检测的时间。

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 检测结果：
	// RESULT_FAILED: 未通过。
	// RESULT_PASSED: 通过。

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 主机节点的实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type SuperNodeListItem struct {

	// 超级节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 超级节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 所属集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 所属集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 节点状态:Running,Ready,Notready,Initializing,Failed,Error

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网ID

	SubNetID *string `json:"SubNetID,omitempty" name:"SubNetID"`
	// 子网名称

	SubNetName *string `json:"SubNetName,omitempty" name:"SubNetName"`
	// 子网网段

	SubNetCidr *string `json:"SubNetCidr,omitempty" name:"SubNetCidr"`
	// 可用区ID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 关联pod数

	RelatePodCount *uint64 `json:"RelatePodCount,omitempty" name:"RelatePodCount"`
	// 关联容器数

	RelateContainerCount *uint64 `json:"RelateContainerCount,omitempty" name:"RelateContainerCount"`
	// agent安装状态UNINSTALL:未安装;INSTALLED:已安装;INSTALLING:安装中;

	AgentStatus *string `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// 节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 集群接入状态

	ClusterAccessedStatus *string `json:"ClusterAccessedStatus,omitempty" name:"ClusterAccessedStatus"`
	// 计费核数

	ChargeCoresCnt *uint64 `json:"ChargeCoresCnt,omitempty" name:"ChargeCoresCnt"`
	// 防护状态:
	// 已防护: Defended
	// 未防护: UnDefended

	DefendStatus *string `json:"DefendStatus,omitempty" name:"DefendStatus"`
}

type DescribeVirusAutoIsolateSampleDownloadURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 样本下载链接

		FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSampleDownloadURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkTopologyPodList struct {

	// pod名称
	//

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// pod状态
	//

	PodStatus *string `json:"PodStatus,omitempty" name:"PodStatus"`
	// pod标签
	//

	PodLabels *string `json:"PodLabels,omitempty" name:"PodLabels"`
	// pod Ip
	//

	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`
	// Node ip
	//

	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`
	// 创建时间
	//

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type RegistryConnDetectParam struct {

	// 镜像仓库名称

	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 检测的用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 检测的密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 检测地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// 检测主机的quuid 或者 backend

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 检测主机的uuid 或者 backend

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 仓库类型

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// api版本

	ApiVersion *string `json:"ApiVersion,omitempty" name:"ApiVersion"`
	// retry:重试
	// first:首次调用联通性检测
	// query:查询结果

	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`
	// tcr情况下instance_id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type RiskDnsEventInfo struct {

	// 事件ID

	EventID *uint64 `json:"EventID,omitempty" name:"EventID"`
	// 事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 恶意请求域名/IP

	Address *string `json:"Address,omitempty" name:"Address"`
	// 容器ID

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 隔离状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 首次发现时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 事件状态
	// EVENT_UNDEAL： 待处理
	// EVENT_DEALED：已处理
	// EVENT_IGNORE： 已忽略
	// EVENT_ADD_WHITE：已加白

	EventStatus *string `json:"EventStatus,omitempty" name:"EventStatus"`
	// 恶意请求次数

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 事件描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 恶意IP所属城市

	City *string `json:"City,omitempty" name:"City"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 内网IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点类型：NORMAL普通节点、SUPER超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// pod ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// pod 名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type AddModEscapeRegexpWhiteListRequest struct {
	*tchttp.BaseRequest

	// 规则信息

	RuleInfo *RegexpRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`
	// 加白名单事件类型 ESCAPE_CGROUPS：利用cgroup机制逃逸 ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸 ESCAPE_DOCKER_API：访问Docker API接口逃逸 ESCAPE_VUL_OCCURRED：逃逸漏洞利用 MOUNT_SENSITIVE_PTAH：敏感路径挂载 PRIVILEGE_CONTAINER_START：特权容器 PRIVILEGE：程序提权逃逸

	EventType []*string `json:"EventType,omitempty" name:"EventType"`
}

func (r *AddModEscapeRegexpWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModEscapeRegexpWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentInstallCommandRequest struct {
	*tchttp.BaseRequest

	// 是否是腾讯云

	IsCloud *bool `json:"IsCloud,omitempty" name:"IsCloud"`
	// 网络类型：basic-基础网络，private-VPC, public-公网，direct-专线

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 地域标示, NetType=direct时必填

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// VpcId, NetType=direct时必填

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 命令有效期，非腾讯云时必填

	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
	// 标签ID列表，IsCloud=false时才会生效

	TagIds []*uint64 `json:"TagIds,omitempty" name:"TagIds"`
}

func (r *DescribeAgentInstallCommandRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentInstallCommandRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyEventTendencyRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeImageDenyEventTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受影响的镜像列表

		List []*VulAffectedImageInfo `json:"List,omitempty" name:"List"`
		// 镜像总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAbnormalProcessRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 策略的ids

	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`
	// 策略开关，true开启，false关闭

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
}

func (r *ModifyAbnormalProcessRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAbnormalProcessRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageDenyEvent struct {

	// 事件类型 EVENT_RISK:风险事件类型，EVENT_PRIVILEGE:特权

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则RuleID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则类型

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 规则启用状态 0:开启，1:关闭

	RuleStatus *int64 `json:"RuleStatus,omitempty" name:"RuleStatus"`
	// 规则策略状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中

	RuleEffectStatus *string `json:"RuleEffectStatus,omitempty" name:"RuleEffectStatus"`
	// 规则内容

	RuleInfo []*string `json:"RuleInfo,omitempty" name:"RuleInfo"`
	// 规则描述

	RuleDescription *string `json:"RuleDescription,omitempty" name:"RuleDescription"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 内网IP

	NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`
	// 主机Quuid

	QUUID *string `json:"QUUID,omitempty" name:"QUUID"`
	// 首次生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 执行动作:
	// BEHAVIOR_ALERT:告警，
	// BEHAVIOR_HOLDUP_SUCCESSED:拦截

	DealBehavior *string `json:"DealBehavior,omitempty" name:"DealBehavior"`
	// 事件ID

	EventID *int64 `json:"EventID,omitempty" name:"EventID"`
	// 外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 超级节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// pod ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// pod name

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type DescribeAssetImageDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像ID

		ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
		// 镜像名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 镜像摘要

		ImageDigest *string `json:"ImageDigest,omitempty" name:"ImageDigest"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 镜像大小

		Size *uint64 `json:"Size,omitempty" name:"Size"`
		// 关联主机个数

		HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`
		// 关联容器个数

		ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
		// 最近扫描时间

		ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
		// 漏洞个数

		VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`
		// 风险行为数

		RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
		// 敏感信息数

		SensitiveInfoCnt *uint64 `json:"SensitiveInfoCnt,omitempty" name:"SensitiveInfoCnt"`
		// 是否信任镜像

		IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`
		// 镜像系统

		OsName *string `json:"OsName,omitempty" name:"OsName"`
		// agent镜像扫描错误

		AgentError *string `json:"AgentError,omitempty" name:"AgentError"`
		// 后端镜像扫描错误

		ScanError *string `json:"ScanError,omitempty" name:"ScanError"`
		// 系统架构

		Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
		// 作者

		Author *string `json:"Author,omitempty" name:"Author"`
		// 构建历史

		BuildHistory *string `json:"BuildHistory,omitempty" name:"BuildHistory"`
		// 木马扫描进度

		ScanVirusProgress *uint64 `json:"ScanVirusProgress,omitempty" name:"ScanVirusProgress"`
		// 漏洞扫进度

		ScanVulProgress *uint64 `json:"ScanVulProgress,omitempty" name:"ScanVulProgress"`
		// 敏感信息扫描进度

		ScanRiskProgress *uint64 `json:"ScanRiskProgress,omitempty" name:"ScanRiskProgress"`
		// 木马扫描错误

		ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`
		// 漏洞扫描错误

		ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`
		// 敏感信息错误

		ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`
		// 镜像扫描状态

		ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
		// 木马病毒数

		VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`
		// 镜像扫描状态

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 剩余扫描时间

		RemainScanTime *uint64 `json:"RemainScanTime,omitempty" name:"RemainScanTime"`
		// 授权为：1，未授权为：0

		IsAuthorized *int64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallPolicyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群的详细信息

		NetPolicy []*NetworkPolicyInfoItem `json:"NetPolicy,omitempty" name:"NetPolicy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallPolicyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchLogsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSearchLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserWorkloadKindsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserWorkloadKindsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserWorkloadKindsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogStorageStatisticResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总容量（单位：B）

		TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`
		// 已使用容量（单位：B）

		UsedSize *uint64 `json:"UsedSize,omitempty" name:"UsedSize"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogStorageStatisticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogStorageStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVirusSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkClusterSwitchListItem struct {

	// 集群Id
	//

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 区域
	//

	Region *string `json:"Region,omitempty" name:"Region"`
	// 集群名称
	//

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群类型
	//

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群版本
	//

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 集群OS信息
	//

	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`
	// 集群状态
	//

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 集群流量开关 ON:打开，OFF关闭
	//

	FlowSwitch *string `json:"FlowSwitch,omitempty" name:"FlowSwitch"`
	// 集群拥有节点个数

	ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitempty" name:"ClusterNodeNum"`
	// 集群流量支持Enable支持，Disable不支持

	FlowSupport *string `json:"FlowSupport,omitempty" name:"FlowSupport"`
	// 集群流量不支持原因

	FlowUnsupportReason *string `json:"FlowUnsupportReason,omitempty" name:"FlowUnsupportReason"`
}

type DescribeVirusAutoIsolateSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVirusAutoIsolateSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRuleDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运行时策略详细信息

		RuleDetail *AccessControlRuleInfo `json:"RuleDetail,omitempty" name:"RuleDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlRuleDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeESHitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ES查询结果JSON

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeESHitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeESHitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServiceListRequest struct {
	*tchttp.BaseRequest

	// Pod的名字,不填表示查询所有

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ServiceName, ExternalIp等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 命名空间,不填表示查询所有

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群Id,不填表示查询所有

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusSampleDownloadUrlRequest struct {
	*tchttp.BaseRequest

	// 木马id

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeVirusSampleDownloadUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusSampleDownloadUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险详情数组

		ClusterRiskItems []*ClusterRiskItem `json:"ClusterRiskItems,omitempty" name:"ClusterRiskItems"`
		// 风险项的总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回资产的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回各类资产的列表。

		AssetInfoList []*ComplianceAssetInfo `json:"AssetInfoList,omitempty" name:"AssetInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceAssetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoAuthorizedImageInfo struct {

	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 授权时间

	AuthorizedTime *string `json:"AuthorizedTime,omitempty" name:"AuthorizedTime"`
	// 授权结果，SUCCESS:成功，REACH_LIMIT:达到授权上限，LICENSE_INSUFFICIENT:授权数不足'

	Status *string `json:"Status,omitempty" name:"Status"`
	// 是否授权，1：是，0：否

	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
}

type ComplianceWhitelistItem struct {

	// 白名单项的ID。

	WhitelistItemId *uint64 `json:"WhitelistItemId,omitempty" name:"WhitelistItemId"`
	// 客户检测项的ID。

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 检测项的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 合规标准的名称。

	StandardName *string `json:"StandardName,omitempty" name:"StandardName"`
	// 合规标准的ID。

	StandardId *uint64 `json:"StandardId,omitempty" name:"StandardId"`
	// 检测项影响的资产的数目。

	AffectedAssetCount *uint64 `json:"AffectedAssetCount,omitempty" name:"AffectedAssetCount"`
	// 最后更新的时间

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 加入到白名单的时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
}

type DescribeVirusAutoIsolateSampleListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>MD5- String - 是否必填：否 - md5 </li>
	// <li>AutoIsolateSwitch- String - 是否必填：否 - 自动隔离开关 </li>
	// <li>VirusName- String - 是否必填：否 - 病毒名 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVirusAutoIsolateSampleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageAutoAuthorizedTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自动授权任务列表

		List []*ImageAutoAuthorizedTask `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageAutoAuthorizedTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageAutoAuthorizedTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLocalStorageItemRequest struct {
	*tchttp.BaseRequest

	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *GetLocalStorageItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLocalStorageItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanRangeInfo struct {

	// true:选择全部；
	// false:部分选择

	IsAll *bool `json:"IsAll,omitempty" name:"IsAll"`
	// SCAN_NORMAL:普通节点；
	// SCAN_SUPER:超级节点
	// SCAN_CONTAINER:容器

	RangeType *string `json:"RangeType,omitempty" name:"RangeType"`
	// 选择的ID

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
}

type DescribePurchaseStateInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePurchaseStateInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePurchaseStateInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateK8sApiAbnormalEventExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>MatchRules - string  - 是否必填: 否 -命中规则筛选</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>
	// <li>Status - string  - 是否必填: 否 -事件状态筛选</li>
	// <li>MatchRuleType - string  - 是否必填: 否 -命中规则类型筛选</li>
	// <li>ClusterRunningStatus - string  - 是否必填: 否 -集群运行状态</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateK8sApiAbnormalEventExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateK8sApiAbnormalEventExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeRegexpWhiteListInfoRequest struct {
	*tchttp.BaseRequest

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

func (r *DescribeEscapeRegexpWhiteListInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeRegexpWhiteListInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceBenchmarkStandardEnable struct {

	// 合规标准的ID。

	StandardId *uint64 `json:"StandardId,omitempty" name:"StandardId"`
	// 是否启用合规标准

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type ClusterCheckTaskItem struct {

	// 指定要扫描的集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群所属地域

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
	// 指定要扫描的节点IP

	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`
	// 按照要扫描的workload名字

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
}

type ReverseShellWhiteListInfo struct {

	// 目标IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 目标端口

	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`
	// 目标进程

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 镜像id数组，为空代表全部

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 白名单id，如果新建则id为空

	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeVulSummaryRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- string- 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>OnlyAffectedContainer-string- 是否必填：否 - 仅展示影响容器的漏洞,true,false</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞分类: SYSTEM:系统漏洞 WEB:web应用漏洞 ALL:全部漏洞</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 反弹shell数组

		EventSet []*ReverseShellEventInfo `json:"EventSet,omitempty" name:"EventSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRiskDnsEventStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRiskDnsEventStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskDnsEventStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanIgnoreVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>VulName- string - 是否必填：否 - 漏洞名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式:DESC,ACS

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段 UpdateTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeScanIgnoreVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanIgnoreVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像漏洞列表

		List []*ImageVul `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenTcssTrialRequest struct {
	*tchttp.BaseRequest
}

func (r *OpenTcssTrialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenTcssTrialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetClusterListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群列表

		List []*AssetClusterListItem `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetClusterListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalEventInfoRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeK8sApiAbnormalEventInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalEventInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserPublicClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserPublicClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserPublicClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulIgnoreLocalImageListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 排序方式:DESC,ACS

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段 ImageSize

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulIgnoreLocalImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulIgnoreLocalImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群id

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 集群名字

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 当前集群扫描任务的进度，100表示扫描完成.

		ScanTaskProgress *int64 `json:"ScanTaskProgress,omitempty" name:"ScanTaskProgress"`
		// 集群版本

		ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
		// 运行时组件

		ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
		// 集群节点数

		ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitempty" name:"ClusterNodeNum"`
		// 集群状态 (Running 运行中 Creating 创建中 Abnormal 异常 )

		ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
		// 集群类型：为托管集群MANAGED_CLUSTER、独立集群INDEPENDENT_CLUSTER

		ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
		// 集群区域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 严重风险检查项的数量

		SeriousRiskCount *uint64 `json:"SeriousRiskCount,omitempty" name:"SeriousRiskCount"`
		// 高风险检查项的数量

		HighRiskCount *uint64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
		// 中风险检查项的数量

		MiddleRiskCount *uint64 `json:"MiddleRiskCount,omitempty" name:"MiddleRiskCount"`
		// 提示风险检查项的数量

		HintRiskCount *uint64 `json:"HintRiskCount,omitempty" name:"HintRiskCount"`
		// 检查任务的状态

		CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
		// 防御容器状态

		DefenderStatus *string `json:"DefenderStatus,omitempty" name:"DefenderStatus"`
		// 扫描任务创建时间

		TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`
		// 网络类型.PublicNetwork为公网类型,VPCNetwork为VPC网络

		NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
		// API Server地址

		ApiServerAddress *string `json:"ApiServerAddress,omitempty" name:"ApiServerAddress"`
		// 节点数

		NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
		// 命名空间数

		NamespaceCount *uint64 `json:"NamespaceCount,omitempty" name:"NamespaceCount"`
		// 工作负载数

		WorkloadCount *uint64 `json:"WorkloadCount,omitempty" name:"WorkloadCount"`
		// Pod数量

		PodCount *uint64 `json:"PodCount,omitempty" name:"PodCount"`
		// Service数量

		ServiceCount *uint64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
		// Ingress数量

		IngressCount *uint64 `json:"IngressCount,omitempty" name:"IngressCount"`
		// 主节点的ip列表

		MasterIps *string `json:"MasterIps,omitempty" name:"MasterIps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerNetworkInfoRequest struct {
	*tchttp.BaseRequest

	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
}

func (r *DescribeAssetContainerNetworkInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerNetworkInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略Id数组

	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteNetworkFirewallPolicyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceEventTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞防御事件趋势

		DefendedList []*VulDefenceEventTendency `json:"DefendedList,omitempty" name:"DefendedList"`
		// 漏洞攻击事件趋势

		AttackList []*VulDefenceEventTendency `json:"AttackList,omitempty" name:"AttackList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefenceEventTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessControlStatusRequest struct {
	*tchttp.BaseRequest

	// 处理事件ids

	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
	// 标记事件的状态，
	// EVENT_DEALED:事件已经处理
	//      EVENT_INGNORE：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 备注事件信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAccessControlStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulDefenceSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserClusterExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务唯一ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserClusterExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserClusterExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryRiskListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryRiskListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRiskListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceEventTendencyRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeVulDefenceEventTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceEventTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRulesExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAccessControlRulesExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskSyscallEventInfo struct {

	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 事件解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件详细描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 系统调用名称

	SyscallName *string `json:"SyscallName,omitempty" name:"SyscallName"`
	// 状态，EVENT_UNDEAL:事件未处理
	//     EVENT_DEALED:事件已经处理
	//     EVENT_INGNORE：事件已经忽略
	//     EVENT_ADD_WHITE：时间已经加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// pod(实例)的名字

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 系统监控名称是否存在

	RuleExist *bool `json:"RuleExist,omitempty" name:"RuleExist"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 节点类型：NORMAL普通节点、SUPER超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 集群I'D

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// pod ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 节点公网ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// uuid

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 节点内网ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type CreateAssetContainerExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetContainerExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetContainerExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulDefenceEventExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulDefenceEventExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulDefenceEventExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageAutoAuthorizedRuleRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageAutoAuthorizedRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageAutoAuthorizedRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterServiceInfo struct {

	// service名称.

	Name *string `json:"Name,omitempty" name:"Name"`
	// service类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Selector信息

	Selectors *string `json:"Selectors,omitempty" name:"Selectors"`
	// Selector数量

	SelectorCount *uint64 `json:"SelectorCount,omitempty" name:"SelectorCount"`
	// 负载均衡Ip

	ExternalIp *string `json:"ExternalIp,omitempty" name:"ExternalIp"`
	// 服务Ip

	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`
	// string类型,以,分隔

	Port *string `json:"Port,omitempty" name:"Port"`
	// 所属集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 所属集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 关联Pod名字

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 关联pod数量

	PodCount *uint64 `json:"PodCount,omitempty" name:"PodCount"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 集群状态

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// service唯一标识符

	ServiceUniqueID *string `json:"ServiceUniqueID,omitempty" name:"ServiceUniqueID"`
}

type ImageRiskInfo struct {

	// 行为

	Behavior *uint64 `json:"Behavior,omitempty" name:"Behavior"`
	// 类型

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 级别

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 详情

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 解决方案

	InstructionContent *string `json:"InstructionContent,omitempty" name:"InstructionContent"`
}

type DescribePromotionActivityRequest struct {
	*tchttp.BaseRequest

	// 活动ID

	ActiveID *uint64 `json:"ActiveID,omitempty" name:"ActiveID"`
}

func (r *DescribePromotionActivityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePromotionActivityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServiceYamlRequest struct {
	*tchttp.BaseRequest

	// Service名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// Service的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群地域

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

func (r *DescribeClusterServiceYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServiceYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetContainerExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	//
	// ContainerName - String - 是否必填：否 - 容器名称模糊搜索
	// Status - String - 是否必填：否 - 容器运行状态筛选，0："created",1："running", 2："paused", 3："restarting", 4："removing", 5："exited", 6："dead"
	// Runas - String - 是否必填：否 - 运行用户筛选
	// ImageName- String - 是否必填：否 - 镜像名称搜索
	// HostIP- string - 是否必填：否 - 主机ip搜索
	// OrderBy - String 是否必填：否 -排序字段，支持：cpu_usage, mem_usage的动态排序 ["cpu_usage","+"] '+'升序、'-'降序
	// NetStatus - String -是否必填: 否 - 容器网络状态筛选 normal isolated isolating isolate_failed restoring restore_failed
	// PodID - String -是否必填: 否 - PodID筛选
	// NodeUniqueID - String -是否必填: 否 - SuperNode筛选
	// PodUid - String -是否必填: 否 - Pod筛选
	// PodIP - String -是否必填: 否 - PodIP筛选
	// NodeType - String -是否必填: 否 - 节点类型筛选:NORMAL:普通节点;SUPER:超级节点

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 需要返回的数量，默认为10，最大值为10000

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateAssetContainerExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetContainerExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateComponentExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式desc ，asc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateComponentExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateComponentExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscapeWhiteListInfo struct {

	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 白名单记录ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 关联主机数量

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 关联容器数量

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 加白事件类型

	EventType []*string `json:"EventType,omitempty" name:"EventType"`
	// 创建时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 镜像大小

	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`
}

type DescribeIACCheckResultListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IAC检测结果总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// IAC检测结果列表

		IACList []*IACItem `json:"IACList,omitempty" name:"IACList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIACCheckResultListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIACCheckResultListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeVirusTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件查杀列表

		List []*VirusTaskInfo `json:"List,omitempty" name:"List"`
		// 总数量(容器任务数量)

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVirusScanAgainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVirusScanAgainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVirusScanAgainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallPolicyStatusRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeNetworkFirewallPolicyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPolicyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAbnormalProcessRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAbnormalProcessRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAbnormalProcessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *DescribeVirusScanTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistrySummaryRequest struct {
	*tchttp.BaseRequest

	// 过滤字段

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetImageRegistrySummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistrySummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyFlowSwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态，可能为：Task_Running,Task_Succ,Task_Error,Task_NoExist

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 任务结果

		TaskResult []*string `json:"TaskResult,omitempty" name:"TaskResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkTopologyFlowSwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyFlowSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryDetailRequest struct {
	*tchttp.BaseRequest

	// 仓库列表id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeAssetImageRegistryDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerMount struct {

	// 挂载类型 bind

	Type *string `json:"Type,omitempty" name:"Type"`
	// 宿主机路径

	Source *string `json:"Source,omitempty" name:"Source"`
	// 容器内路径

	Destination *string `json:"Destination,omitempty" name:"Destination"`
	// 模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 读写权限

	RW *bool `json:"RW,omitempty" name:"RW"`
	// 传播类型

	Propagation *string `json:"Propagation,omitempty" name:"Propagation"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 驱动

	Driver *string `json:"Driver,omitempty" name:"Driver"`
}

type DescribeVulDefenceHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞防御的主机列表

		List []*VulDefenceHost `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefenceHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceEventDetailRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	EventID *int64 `json:"EventID,omitempty" name:"EventID"`
}

func (r *DescribeVulDefenceEventDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceEventDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogJoinStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecLogJoinStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogJoinStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkFirewallUndoPublishResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkFirewallUndoPublishResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkFirewallUndoPublishResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRulesRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAbnormalProcessRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerSecEventSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 未处理逃逸事件

		UnhandledEscapeCnt *uint64 `json:"UnhandledEscapeCnt,omitempty" name:"UnhandledEscapeCnt"`
		// 未处理反弹shell事件

		UnhandledReverseShellCnt *uint64 `json:"UnhandledReverseShellCnt,omitempty" name:"UnhandledReverseShellCnt"`
		// 未处理高危系统调用

		UnhandledRiskSyscallCnt *uint64 `json:"UnhandledRiskSyscallCnt,omitempty" name:"UnhandledRiskSyscallCnt"`
		// 未处理异常进程

		UnhandledAbnormalProcessCnt *uint64 `json:"UnhandledAbnormalProcessCnt,omitempty" name:"UnhandledAbnormalProcessCnt"`
		// 未处理文件篡改

		UnhandledFileCnt *uint64 `json:"UnhandledFileCnt,omitempty" name:"UnhandledFileCnt"`
		// 未处理木马事件

		UnhandledVirusEventCnt *uint64 `json:"UnhandledVirusEventCnt,omitempty" name:"UnhandledVirusEventCnt"`
		// 未处理恶意外连事件

		UnhandledMaliciousConnectionEventCnt *uint64 `json:"UnhandledMaliciousConnectionEventCnt,omitempty" name:"UnhandledMaliciousConnectionEventCnt"`
		// 未处理k8sApi事件

		UnhandledK8sApiEventCnt *uint64 `json:"UnhandledK8sApiEventCnt,omitempty" name:"UnhandledK8sApiEventCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerSecEventSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerSecEventSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryAutoAuthorizedRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仓库类型

		RangeType *string `json:"RangeType,omitempty" name:"RangeType"`
		// 每天最大授权数

		MaxDailyCount *int64 `json:"MaxDailyCount,omitempty" name:"MaxDailyCount"`
		// 开关

		IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
		// 插入时间

		InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
		// 最近更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 授权范围数量

		NamespaceCount *uint64 `json:"NamespaceCount,omitempty" name:"NamespaceCount"`
		// 授权范围数量

		ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`
		// 规则ID

		RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
		// 自动扫描开关

		AutoScanEnabled *int64 `json:"AutoScanEnabled,omitempty" name:"AutoScanEnabled"`
		// 自动扫描范围

		ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRegistryAutoAuthorizedRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryAutoAuthorizedRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoUpgradeSwitchRequest struct {
	*tchttp.BaseRequest

	// 开关

	Switch *bool `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyAutoUpgradeSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoUpgradeSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkClusterNamespaceInfo struct {

	// 网络空间标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 网络空间名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

type UpdateNetworkFirewallPolicyYamlDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNetworkFirewallPolicyYamlDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNetworkFirewallPolicyYamlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageNewVul struct {

	// 漏洞id

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 观点验证程序id

	POCID *string `json:"POCID,omitempty" name:"POCID"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 涉及组件信息

	Components []*ComponentsInfo `json:"Components,omitempty" name:"Components"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 分类2

	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 描述

	Des *string `json:"Des,omitempty" name:"Des"`
	// 解决方案

	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`
	// 引用

	Reference []*string `json:"Reference,omitempty" name:"Reference"`
	// 防御方案

	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`
	// 提交时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// Cvss分数

	CvssScore *string `json:"CvssScore,omitempty" name:"CvssScore"`
	// Cvss信息

	CvssVector *string `json:"CvssVector,omitempty" name:"CvssVector"`
	// 是否建议修复

	IsSuggest *string `json:"IsSuggest,omitempty" name:"IsSuggest"`
	// 修复版本号

	FixedVersions *string `json:"FixedVersions,omitempty" name:"FixedVersions"`
	// 漏洞标签:"CanBeFixed","DynamicLevelPoc","DynamicLevelExp"

	Tag []*string `json:"Tag,omitempty" name:"Tag"`
	// 组件名

	Component *string `json:"Component,omitempty" name:"Component"`
	// 组件版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 攻击热度 0-3

	AttackLevel *int64 `json:"AttackLevel,omitempty" name:"AttackLevel"`
	// 镜像层指令列表

	LayerCmds []*string `json:"LayerCmds,omitempty" name:"LayerCmds"`
	// 镜像层id列表

	LayerIds []*string `json:"LayerIds,omitempty" name:"LayerIds"`
}

type DescribeAbnormalProcessEventsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAbnormalProcessEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInspectionReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 报告名称

		ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
		// 下载链接

		ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInspectionReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInspectionReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddModEscapeRegexpWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddModEscapeRegexpWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModEscapeRegexpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersNamespacesRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤字段

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClustersNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEventEscapeImageStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEventEscapeImageStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEventEscapeImageStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAssetSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClusterAssetSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAssetSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterPodInfo struct {

	// Pod名称.

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// Pod状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// Pod IP

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 节点内网Ip

	NodeLanIP *string `json:"NodeLanIP,omitempty" name:"NodeLanIP"`
	// 所属的工作负载名字

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 所属工作负载类型

	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`
	// 所属集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 所属集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 所属命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 运行时间

	Age *string `json:"Age,omitempty" name:"Age"`
	// 创建时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 重启次数

	Restarts *uint64 `json:"Restarts,omitempty" name:"Restarts"`
	// 关联的service名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 关联的service数量

	ServiceCount *uint64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
	// 关联的容器名字

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 关联的容器数量

	ContainerCount *uint64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// CPU占用率

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 内存占用量

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// Pod标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 集群状态

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 工作负载标签

	WorkloadLabels *string `json:"WorkloadLabels,omitempty" name:"WorkloadLabels"`
	// 容器Id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机Id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// abc

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// NORMAL：普通节点 SUPER：超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 计费核数

	ChargeCoresCnt *uint64 `json:"ChargeCoresCnt,omitempty" name:"ChargeCoresCnt"`
}

type ImagesBindRuleInfo struct {

	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 关联容器数量

	ContainerCnt *int64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 绑定规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 镜像大小

	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`
	// 最近扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
}

type CreateRefreshTaskRequest struct {
	*tchttp.BaseRequest

	// 指定集群列表,若为空则标识同步所有集群

	ClusterIDs []*string `json:"ClusterIDs,omitempty" name:"ClusterIDs"`
	// 是否只同步列表

	IsSyncListOnly *bool `json:"IsSyncListOnly,omitempty" name:"IsSyncListOnly"`
}

func (r *CreateRefreshTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRefreshTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRiskSyscallEventsRequest struct {
	*tchttp.BaseRequest

	// 事件ids

	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
}

func (r *DeleteRiskSyscallEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRiskSyscallEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanRegistryImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像列表

		List []*VulScanRegistryImageInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulScanRegistryImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanRegistryImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanLocalImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像列表

		List []*VulScanImageInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulScanLocalImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanLocalImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageDenyEventExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageDenyEventExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageDenyEventExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerSecEventSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeContainerSecEventSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerSecEventSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeABTestConfigRequest struct {
	*tchttp.BaseRequest

	// 灰度项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *DescribeABTestConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeABTestConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCheckComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// "InstallSucc"表示安装成功，"InstallFailed"表示安装失败

		InstallResult *string `json:"InstallResult,omitempty" name:"InstallResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCheckComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCheckComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRiskTendencyInfo struct {

	// 趋势列表

	ImageRiskSet []*RunTimeTendencyInfo `json:"ImageRiskSet,omitempty" name:"ImageRiskSet"`
	// 风险类型：
	// IRT_VULNERABILITY : 安全漏洞
	// IRT_MALWARE_VIRUS: 木马病毒
	// IRT_RISK:敏感信息

	ImageRiskType *string `json:"ImageRiskType,omitempty" name:"ImageRiskType"`
}

type DeleteClusterWorkloadFromWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除检测项白名单的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterWorkloadFromWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterWorkloadFromWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulImageSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受严重或高危漏洞影响的镜像数

		SeriousVulImageCount *int64 `json:"SeriousVulImageCount,omitempty" name:"SeriousVulImageCount"`
		// 已扫描的镜像数

		ScannedImageCount *int64 `json:"ScannedImageCount,omitempty" name:"ScannedImageCount"`
		// 漏洞总数量

		VulTotalCount *int64 `json:"VulTotalCount,omitempty" name:"VulTotalCount"`
		// 系统漏洞数

		SysTemVulCount *int64 `json:"SysTemVulCount,omitempty" name:"SysTemVulCount"`
		// web应用漏洞数

		WebVulCount *int64 `json:"WebVulCount,omitempty" name:"WebVulCount"`
		// 已授权镜像数

		AllAuthorizedImageCount *int64 `json:"AllAuthorizedImageCount,omitempty" name:"AllAuthorizedImageCount"`
		// 应急漏洞数

		EmergencyVulCount *int64 `json:"EmergencyVulCount,omitempty" name:"EmergencyVulCount"`
		// 支持扫描的漏洞总数量

		SupportVulTotalCount *int64 `json:"SupportVulTotalCount,omitempty" name:"SupportVulTotalCount"`
		// 漏洞库更新时间

		VulLibraryUpdateTime *string `json:"VulLibraryUpdateTime,omitempty" name:"VulLibraryUpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulImageSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulImageSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulContainerExportJobRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID</li>
	// <li>ContainerName- String -是否必填: 否 - 容器名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *CreateVulContainerExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulContainerExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterWorkloadsRequest struct {
	*tchttp.BaseRequest

	// 集群Id,不输入表示查询所有

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：WorkloadName工作负载名字, WorkloadType工作负载类型等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClusterWorkloadsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterWorkloadsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchExportListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSearchExportListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchExportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNetworkFirewallPolicyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRiskListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRiskListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRiskListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousConnectionRuleSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMaliciousConnectionRuleSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionRuleSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageScanSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetImageScanSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEmergencyVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞列表

		List []*EmergencyVulInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEmergencyVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEmergencyVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCompliancePolicyAssetSetFromWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCompliancePolicyAssetSetFromWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCompliancePolicyAssetSetFromWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageScanTaskRequest struct {
	*tchttp.BaseRequest

	// 是否扫描全部镜像；全部镜像，镜像列表和根据过滤条件筛选三选一。

	All *bool `json:"All,omitempty" name:"All"`
	// 需要扫描的镜像列表；全部镜像，镜像列表和根据过滤条件筛选三选一。

	Images []*string `json:"Images,omitempty" name:"Images"`
	// 扫描漏洞；漏洞，木马和风险需选其一

	ScanVul *bool `json:"ScanVul,omitempty" name:"ScanVul"`
	// 扫描木马；漏洞，木马和风险需选其一

	ScanVirus *bool `json:"ScanVirus,omitempty" name:"ScanVirus"`
	// 扫描风险；漏洞，木马和风险需选其一

	ScanRisk *bool `json:"ScanRisk,omitempty" name:"ScanRisk"`
	// 根据过滤条件筛选出镜像；全部镜像，镜像列表和根据过滤条件筛选三选一。

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 根据过滤条件筛选出镜像，再排除个别镜像

	ExcludeImageIds []*string `json:"ExcludeImageIds,omitempty" name:"ExcludeImageIds"`
	// 镜像是否存在运行中的容器

	ContainerRunning *bool `json:"ContainerRunning,omitempty" name:"ContainerRunning"`
	// 扫描范围 0 全部授权镜像，1自选镜像，2 推荐扫描

	ScanScope *uint64 `json:"ScanScope,omitempty" name:"ScanScope"`
	// 任务超时时长单位秒，默认1小时

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
}

func (r *CreateAssetImageScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallAuditRecordRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - Action
	// Name 可取值：publish，unpublish，confirm，add，update，delete

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkFirewallAuditRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallAuditRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则信息

		RuleSet []*EscapeRule `json:"RuleSet,omitempty" name:"RuleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanComplianceAssetsByPolicyItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回重新检测任务的ID。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanComplianceAssetsByPolicyItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanComplianceAssetsByPolicyItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVulListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageVulListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVulListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAbnormalProcessStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAbnormalProcessStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAbnormalProcessStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterRegionItem struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域缩写

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 地域Id

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域中文名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域英文名

	RegionNameEn *string `json:"RegionNameEn,omitempty" name:"RegionNameEn"`
}

type ComponentInfo struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type CreateComponentExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateComponentExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateComponentExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateManualUpgradeVersionRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateManualUpgradeVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateManualUpgradeVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportJobInfo struct {

	// 任务ID

	JobID *string `json:"JobID,omitempty" name:"JobID"`
	// 任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 导出状态

	ExportStatus *string `json:"ExportStatus,omitempty" name:"ExportStatus"`
	// 导出进度

	ExportProgress *int64 `json:"ExportProgress,omitempty" name:"ExportProgress"`
	// 失败原因

	FailureMsg *string `json:"FailureMsg,omitempty" name:"FailureMsg"`
	// 超时时间

	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`
	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
}

type RaspInfo struct {

	// rasp名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// rasp  描述

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateSuperNodePodExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSuperNodePodExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSuperNodePodExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulnerabilityFocusRuleDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统规则

		System *FocusVulnerabilityRuleItem `json:"System,omitempty" name:"System"`
		// 用户规则

		User *FocusVulnerabilityRuleItem `json:"User,omitempty" name:"User"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulnerabilityFocusRuleDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulnerabilityFocusRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeEscapeEventsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrModifyMaliciousConnectionWhiteListRequest struct {
	*tchttp.BaseRequest

	// 枚举
	// IP: IP
	// 域名：DOMAIN

	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`
	// 白名单域名

	WhiteDomainList []*string `json:"WhiteDomainList,omitempty" name:"WhiteDomainList"`
	// 白名单IP

	WhiteIPList []*string `json:"WhiteIPList,omitempty" name:"WhiteIPList"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 白名单记录id，只有修改时需要

	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *AddOrModifyMaliciousConnectionWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrModifyMaliciousConnectionWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		List []*ImageDenyRule `json:"List,omitempty" name:"List"`
		// 规则总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageLayer struct {

	// 序号

	LayerNum *int64 `json:"LayerNum,omitempty" name:"LayerNum"`
	// 层id

	LayerId *string `json:"LayerId,omitempty" name:"LayerId"`
	// 层指令

	LayerCmd *string `json:"LayerCmd,omitempty" name:"LayerCmd"`
	// 层大小

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// 漏洞严重个数

	CriticalLevelVulCnt *int64 `json:"CriticalLevelVulCnt,omitempty" name:"CriticalLevelVulCnt"`
	// 漏洞高危个数

	HighLevelVulCnt *int64 `json:"HighLevelVulCnt,omitempty" name:"HighLevelVulCnt"`
	// 漏洞中危个数

	MediumLevelVulCnt *int64 `json:"MediumLevelVulCnt,omitempty" name:"MediumLevelVulCnt"`
	// 漏洞低危个数

	LowLevelVulCnt *int64 `json:"LowLevelVulCnt,omitempty" name:"LowLevelVulCnt"`
	// 木马个数

	VirusCnt *int64 `json:"VirusCnt,omitempty" name:"VirusCnt"`
	// 敏感信息个数

	RiskCnt *int64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
	// 漏洞列表

	VulList []*ImageNewVul `json:"VulList,omitempty" name:"VulList"`
	// 木马列表

	VirusList []*ImageVirus `json:"VirusList,omitempty" name:"VirusList"`
	// 敏感信息列表

	SensitiveList []*ImageRisk `json:"SensitiveList,omitempty" name:"SensitiveList"`
}

type ProjectInfo struct {

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 项目ID

	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`
}

type DescribeWarningRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略列表

		WarningRules []*WarningRule `json:"WarningRules,omitempty" name:"WarningRules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWarningRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarningRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyWorkLoadChartRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// TimeRange代表时间范围：string类型，前2个值代表开始时间和结束时间
	//

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeNetworkTopologyWorkLoadChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyWorkLoadChartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveAssetImageRegistryRegistryDetailRequest struct {
	*tchttp.BaseRequest

	// 仓库唯一id

	RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *RemoveAssetImageRegistryRegistryDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveAssetImageRegistryRegistryDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRiskTendencyRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeImageRiskTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRiskTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellRegexpWhiteListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	//
	// RuleName- String - 是否必填：否 - 规则名称

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeReverseShellRegexpWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellRegexpWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanComplianceAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回重新检测任务的ID。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanComplianceAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanComplianceAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVirusScanNewTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVirusScanNewTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVirusScanNewTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞趋势列表

		VulTendencySet []*VulTendencyInfo `json:"VulTendencySet,omitempty" name:"VulTendencySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSuperNodeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 超级节点详情

		Info *SuperNodeInfo `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSuperNodeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSuperNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CKafkaTopicInfo struct {

	// 主题ID

	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type DescribeSystemVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞列表

		List []*VulInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyRealFlowListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeNetworkTopologyRealFlowListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyRealFlowListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceAssetSummary struct {

	// 资产类别。

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 是否为客户的首次检测。与CheckStatus配合使用。

	IsCustomerFirstCheck *bool `json:"IsCustomerFirstCheck,omitempty" name:"IsCustomerFirstCheck"`
	// 检测状态
	//
	// CHECK_UNINIT, 用户未启用此功能
	//
	// CHECK_INIT, 待检测
	//
	// CHECK_RUNNING, 检测中
	//
	// CHECK_FINISHED, 检测完成
	//
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 此类别的检测进度，为 0~100 的数。若未在检测中，无此字段。

	CheckProgress *float64 `json:"CheckProgress,omitempty" name:"CheckProgress"`
	// 此类资产通过的检测项的数目。

	PassedPolicyItemCount *uint64 `json:"PassedPolicyItemCount,omitempty" name:"PassedPolicyItemCount"`
	// 此类资产未通过的检测的数目。

	FailedPolicyItemCount *uint64 `json:"FailedPolicyItemCount,omitempty" name:"FailedPolicyItemCount"`
	// 此类资产下未通过的严重级别的检测项的数目。

	FailedCriticalPolicyItemCount *uint64 `json:"FailedCriticalPolicyItemCount,omitempty" name:"FailedCriticalPolicyItemCount"`
	// 此类资产下未通过的高危检测项的数目。

	FailedHighRiskPolicyItemCount *uint64 `json:"FailedHighRiskPolicyItemCount,omitempty" name:"FailedHighRiskPolicyItemCount"`
	// 此类资产下未通过的中危检测项的数目。

	FailedMediumRiskPolicyItemCount *uint64 `json:"FailedMediumRiskPolicyItemCount,omitempty" name:"FailedMediumRiskPolicyItemCount"`
	// 此类资产下未通过的低危检测项的数目。

	FailedLowRiskPolicyItemCount *uint64 `json:"FailedLowRiskPolicyItemCount,omitempty" name:"FailedLowRiskPolicyItemCount"`
	// 此类资产下提示级别的检测项的数目。

	NoticePolicyItemCount *uint64 `json:"NoticePolicyItemCount,omitempty" name:"NoticePolicyItemCount"`
	// 通过检测的资产的数目。

	PassedAssetCount *uint64 `json:"PassedAssetCount,omitempty" name:"PassedAssetCount"`
	// 未通过检测的资产的数目。

	FailedAssetCount *uint64 `json:"FailedAssetCount,omitempty" name:"FailedAssetCount"`
	// 此类资产的合规率，0~100的数。

	AssetPassedRate *float64 `json:"AssetPassedRate,omitempty" name:"AssetPassedRate"`
	// 检测失败的资产的数目。

	ScanFailedAssetCount *uint64 `json:"ScanFailedAssetCount,omitempty" name:"ScanFailedAssetCount"`
	// 上次检测的耗时，单位为秒。

	CheckCostTime *float64 `json:"CheckCostTime,omitempty" name:"CheckCostTime"`
	// 上次检测的时间。

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 定时检测规则。

	PeriodRule *CompliancePeriodTaskRule `json:"PeriodRule,omitempty" name:"PeriodRule"`
	// 已开启的检查项总数

	OpenPolicyItemCount *uint64 `json:"OpenPolicyItemCount,omitempty" name:"OpenPolicyItemCount"`
	// 已忽略的检查项总数

	IgnoredPolicyItemCount *uint64 `json:"IgnoredPolicyItemCount,omitempty" name:"IgnoredPolicyItemCount"`
}

type RuleBaseInfo struct {

	// true: 默认策略，false:自定义策略

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
	// 策略生效镜像数量

	EffectImageCount *uint64 `json:"EffectImageCount,omitempty" name:"EffectImageCount"`
	// 策略Id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 策略更新时间, 存在为空的情况

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 策略名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 编辑用户名称

	EditUserName *string `json:"EditUserName,omitempty" name:"EditUserName"`
	// true: 策略启用，false：策略禁用

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
}

type DescribeAccessControlRulesRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAccessControlRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetRequest struct {
	*tchttp.BaseRequest

	// 全部同步，俩参数必选一个 All优先

	All *bool `json:"All,omitempty" name:"All"`
	// 要同步的主机列表uuid ，俩参数必选一个 All优先

	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
}

func (r *ModifyAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 严重风险影响的节点数量,返回7天数据

		SeriousRiskNodeCount []*uint64 `json:"SeriousRiskNodeCount,omitempty" name:"SeriousRiskNodeCount"`
		// 高风险影响的节点的数量,返回7天数据

		HighRiskNodeCount []*uint64 `json:"HighRiskNodeCount,omitempty" name:"HighRiskNodeCount"`
		// 中风险检查项的节点数量,返回7天数据

		MiddleRiskNodeCount []*uint64 `json:"MiddleRiskNodeCount,omitempty" name:"MiddleRiskNodeCount"`
		// 提示风险检查项的节点数量,返回7天数据

		HintRiskNodeCount []*uint64 `json:"HintRiskNodeCount,omitempty" name:"HintRiskNodeCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResultSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResultSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellWhiteListsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeReverseShellWhiteListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellWhiteListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkClusterNamespaceLabelInfo struct {

	// 网络空间标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 网络空间名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeAssetImageRegistryRegistryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像仓库列表

		List []*ImageRepoRegistryInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryRegistryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRegistryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulnerabilityFocusRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulnerabilityFocusRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulnerabilityFocusRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CKafkaInstanceInfo struct {

	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 主题列表

	TopicList []*CKafkaTopicInfo `json:"TopicList,omitempty" name:"TopicList"`
	// 路由列表

	RouteList []*CkafkaRouteInfo `json:"RouteList,omitempty" name:"RouteList"`
	// kafka版本号

	KafkaVersion *string `json:"KafkaVersion,omitempty" name:"KafkaVersion"`
}

type SuperNodeInstallTaskDetailInfo struct {

	// 超级节点 ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 超级节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 超级节点状态	 Running //运行中
	// 	 Ready //准备
	// 	 Notready //未准备好
	// 	 Initializing //初始化
	// 	 Failed //失败
	// 	 Error //错误

	NodeStatus *string `json:"NodeStatus,omitempty" name:"NodeStatus"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 任务状态	INSTALLING:安装中;
	// 	INSTALL_FAILURE:安装失败;
	// 	INSTALL_SUCCESS:安装成功

	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// 安装任务信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 超级节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 容器数

	ContainerCount *uint64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// pod数

	PodCount *uint64 `json:"PodCount,omitempty" name:"PodCount"`
}

type DescribePublicKeyRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePublicKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallPolicyYamlDetailRequest struct {
	*tchttp.BaseRequest

	// 策略Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeNetworkFirewallPolicyYamlDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPolicyYamlDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirusDetailInfo struct {

	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 木马文件大小

	Size []*uint64 `json:"Size,omitempty" name:"Size"`
	// 木马文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 最近生成时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 风险等级 RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 进程md5

	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
	// 进程id

	ProcessId []*uint64 `json:"ProcessId,omitempty" name:"ProcessId"`
	// 进程参数

	ProcessArgv *string `json:"ProcessArgv,omitempty" name:"ProcessArgv"`
	// 进程链

	ProcessChan *string `json:"ProcessChan,omitempty" name:"ProcessChan"`
	// 进程组

	ProcessAccountGroup *string `json:"ProcessAccountGroup,omitempty" name:"ProcessAccountGroup"`
	// 进程启动用户

	ProcessStartAccount *string `json:"ProcessStartAccount,omitempty" name:"ProcessStartAccount"`
	// 进程文件权限

	ProcessFileAuthority *string `json:"ProcessFileAuthority,omitempty" name:"ProcessFileAuthority"`
	// 来源：0：一键扫描， 1：定时扫描 2：实时监控

	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
	// 集群名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 事件描述

	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
	// 建议方案

	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
	// 备注

	Mark *string `json:"Mark,omitempty" name:"Mark"`
	// 风险文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件MD5

	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`
	// 事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
}

type DescribePodContainersRequest struct {
	*tchttp.BaseRequest

	// Pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Name 可取值：ClusterId集群id,Namespace命名空间等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribePodContainersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePodContainersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterNodeFromWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加检测项白名单的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterNodeFromWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterNodeFromWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像信息

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopVirusScanTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 需要停止的容器id 为空默认停止整个任务

	ContainerIds []*string `json:"ContainerIds,omitempty" name:"ContainerIds"`
}

func (r *StopVirusScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopVirusScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRiskSyscallStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRiskSyscallStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskSyscallStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRulesExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAbnormalProcessRulesExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRulesExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetDetailInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 某资产的详情。

		AssetDetailInfo *ComplianceAssetDetailInfo `json:"AssetDetailInfo,omitempty" name:"AssetDetailInfo"`
		// 当资产为容器时，返回此字段。

		ContainerDetailInfo *ComplianceContainerDetailInfo `json:"ContainerDetailInfo,omitempty" name:"ContainerDetailInfo"`
		// 当资产为镜像时，返回此字段。

		ImageDetailInfo *ComplianceImageDetailInfo `json:"ImageDetailInfo,omitempty" name:"ImageDetailInfo"`
		// 当资产为主机时，返回此字段。

		HostDetailInfo *ComplianceHostDetailInfo `json:"HostDetailInfo,omitempty" name:"HostDetailInfo"`
		// 当资产为K8S时，返回此字段。

		K8SDetailInfo *ComplianceK8SDetailInfo `json:"K8SDetailInfo,omitempty" name:"K8SDetailInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceAssetDetailInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		List []*K8sApiAbnormalRuleListItem `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEscapeRegexpWhiteListRequest struct {
	*tchttp.BaseRequest

	// 规则ID数组

	RuleIDs []*string `json:"RuleIDs,omitempty" name:"RuleIDs"`
}

func (r *DeleteEscapeRegexpWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEscapeRegexpWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeWebVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAbnormalProcessRulesRequest struct {
	*tchttp.BaseRequest

	// 策略的ids

	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`
}

func (r *DeleteAbnormalProcessRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAbnormalProcessRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalRuleListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段。
	// <li>UpdateTime - string  - 是否必填: 否 -最后更新时间</li>
	// <li>EffectClusterCount - string  - 是否必填: 否 -影响集群数</li>

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeK8sApiAbnormalRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellSystemPolicyConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内网告警展示

		InnerNetAlarmShow *bool `json:"InnerNetAlarmShow,omitempty" name:"InnerNetAlarmShow"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellSystemPolicyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellSystemPolicyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAffectedWorkloadListRequest struct {
	*tchttp.BaseRequest

	// 唯一的检测项的ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：WorkloadType,ClusterId

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAffectedWorkloadListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedWorkloadListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellWhiteListDetailRequest struct {
	*tchttp.BaseRequest

	// 白名单id

	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
}

func (r *DescribeReverseShellWhiteListDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellWhiteListDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SuperNodeInfo struct {

	// 超级节点ID

	SuperNodeID *string `json:"SuperNodeID,omitempty" name:"SuperNodeID"`
	// 超级节点名称

	SuperNodeName *string `json:"SuperNodeName,omitempty" name:"SuperNodeName"`
	// 节点状态:Running,Ready,Notready,Initializing,Failed,Error

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网名称

	SubNetName *string `json:"SubNetName,omitempty" name:"SubNetName"`
	// 子网ID

	SubNetID *string `json:"SubNetID,omitempty" name:"SubNetID"`
	// 子网网段

	SubNetCIDR *string `json:"SubNetCIDR,omitempty" name:"SubNetCIDR"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区ID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 核心数

	CpuOnline *int64 `json:"CpuOnline,omitempty" name:"CpuOnline"`
	// 总内存

	MemTotal *int64 `json:"MemTotal,omitempty" name:"MemTotal"`
	// 集群信息

	ClusterInfo *AssetClusterListItem `json:"ClusterInfo,omitempty" name:"ClusterInfo"`
	// 关联pod数量

	RelatePodCount *uint64 `json:"RelatePodCount,omitempty" name:"RelatePodCount"`
	// 关联容器数量

	RelateContainerCount *uint64 `json:"RelateContainerCount,omitempty" name:"RelateContainerCount"`
	// agent安装状态:
	// 	UNINSTALL:未安装;
	// 	INSTALLED:已安装;
	// 	INSTALLING:安装中;

	AgentStatus *string `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// 超级节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Kubelet版本

	KubeletVersion *string `json:"KubeletVersion,omitempty" name:"KubeletVersion"`
}

type ConnDetectConfig struct {

	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetImageRegistryListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段
	// IsAuthorized是否授权，取值全部all，未授权0，已授权1

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式，asc，desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 是否仅展示各repository最新的镜像, 默认为false

	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
	// 是否仅展示运行中容器镜像

	IsRunning *bool `json:"IsRunning,omitempty" name:"IsRunning"`
}

func (r *DescribeAssetImageRegistryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEscapeWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单记录ID数组

	IDSet []*int64 `json:"IDSet,omitempty" name:"IDSet"`
}

func (r *DeleteEscapeWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEscapeWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEditAbnormalProcessRuleRequest struct {
	*tchttp.BaseRequest

	// 增加策略信息，策略id为空，编辑策略是id不能为空

	RuleInfo *AbnormalProcessRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`
	// 仅在加白的时候带上

	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *AddEditAbnormalProcessRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditAbnormalProcessRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVulListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像信息

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVulListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVulListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePodContainersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// pod的容器列表

		PodContainerList []*PodContainerInfo `json:"PodContainerList,omitempty" name:"PodContainerList"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePodContainersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePodContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearLocalStorageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearLocalStorageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearLocalStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryAssetStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 更新进度状态,doing更新中，success更新成功，failed失败

		Status *string `json:"Status,omitempty" name:"Status"`
		// 错误信息

		Err *string `json:"Err,omitempty" name:"Err"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryAssetStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryAssetStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEventEscapeImageStatusRequest struct {
	*tchttp.BaseRequest

	// 标记事件的状态:
	// EVENT_DEALED:事件处理
	// EVENT_IGNORE"：事件忽略
	// EVENT_DEL:事件删除
	// EVENT_ADD_WHITE:事件加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 唯一值

	UniqueKeys []*string `json:"UniqueKeys,omitempty" name:"UniqueKeys"`
	// 加白镜像id

	ImageIdSet []*string `json:"ImageIdSet,omitempty" name:"ImageIdSet"`
	// 加白事件类型

	EventType []*string `json:"EventType,omitempty" name:"EventType"`
}

func (r *ModifyEventEscapeImageStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEventEscapeImageStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddModReverseShellRegexpWhiteListRequest struct {
	*tchttp.BaseRequest

	// 规则信息

	RuleInfo *RegexpRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`
}

func (r *AddModReverseShellRegexpWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModReverseShellRegexpWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportJobDownloadURLRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	JobID *string `json:"JobID,omitempty" name:"JobID"`
}

func (r *DescribeExportJobDownloadURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportJobDownloadURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogDeliveryKafkaSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecLogDeliveryKafkaSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogDeliveryKafkaSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryAutoAuthorizedRepoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		List []*AutoAuthImageInfo `json:"List,omitempty" name:"List"`
		// 数据量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRegistryAutoAuthorizedRepoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryAutoAuthorizedRepoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeRegexpWhiteListInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		RuleInfo *RegexpRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`
		// 加白名单事件类型 ESCAPE_CGROUPS：利用cgroup机制逃逸 ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸 ESCAPE_DOCKER_API：访问Docker API接口逃逸 ESCAPE_VUL_OCCURRED：逃逸漏洞利用 MOUNT_SENSITIVE_PTAH：敏感路径挂载 PRIVILEGE_CONTAINER_START：特权容器 PRIVILEGE：程序提权逃逸

		EventType []*string `json:"EventType,omitempty" name:"EventType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeRegexpWhiteListInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeRegexpWhiteListInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceTaskAssetSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回用户的状态，
		//
		// USER_UNINIT: 用户未初始化。
		// USER_INITIALIZING，表示用户正在初始化环境。
		// USER_NORMAL: 正常状态。

		Status *string `json:"Status,omitempty" name:"Status"`
		// 返回各类资产的汇总信息的列表。

		AssetSummaryList []*ComplianceAssetSummary `json:"AssetSummaryList,omitempty" name:"AssetSummaryList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceTaskAssetSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceTaskAssetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerPortListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口列表

		List []*ContainerPortInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetContainerPortListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerPortListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousConnectionBlackListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>RequestType- string - 是否必填：否 - 请求类型，全部请求类型：ALL；域名：DOMAIN；IP: IP</li>
	// <li>BlackDomain- string - 是否必填：否 - 自定义黑域名</li>
	// <li>BlackIP- string - 是否必填：否 - 自定义黑IP</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeMaliciousConnectionBlackListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionBlackListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageDenyRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageDenyRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageDenyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousConnectionWhiteListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>RequestType- string - 是否必填：是 - 请求类型，全部请求类型：ALL；域名：DOMAIN；IP: IP</li>
	// <li>WhiteDomain- string - 是否必填：否 - 自定义白域名</li>
	// <li>WhiteIP- string - 是否必填：否 - 自定义白名单IP</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeMaliciousConnectionWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallWhiteListDetailRequest struct {
	*tchttp.BaseRequest

	// 白名单id

	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
}

func (r *DescribeRiskSyscallWhiteListDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallWhiteListDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClsLogsetInfo struct {

	// 日志集ID

	LogsetID *string `json:"LogsetID,omitempty" name:"LogsetID"`
	// 日志集名称

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// cls主题列表

	TopicList []*ClsTopicInfo `json:"TopicList,omitempty" name:"TopicList"`
}

type RiskSyscallWhiteListBaseInfo struct {

	// 白名单id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 镜像数量

	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`
	// 连接进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 系统调用名称列表

	SyscallNames []*string `json:"SyscallNames,omitempty" name:"SyscallNames"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否是全局白名单，true全局

	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 镜像id数组

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

type DescribeAssetImageScanSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetImageScanSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageScanSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterAccessRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterIDs []*string `json:"ClusterIDs,omitempty" name:"ClusterIDs"`
}

func (r *CreateClusterAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanInfoRequest struct {
	*tchttp.BaseRequest

	// 本地镜像漏洞扫描任务ID，无则返回最近一次的漏洞任务扫描

	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`
	// 仓库镜像漏洞扫描任务ID，无则返回最近一次的漏洞任务扫描

	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
}

func (r *DescribeVulScanInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVirusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVirusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetImageScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlRuleInfo struct {

	// 策略id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 开关,true:开启，false:禁用

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 生效镜像id，空数组代表全部镜像

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 用户策略的子策略数组

	ChildRules []*AccessControlChildRuleInfo `json:"ChildRules,omitempty" name:"ChildRules"`
	// 策略名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 系统策略的子策略数组

	SystemChildRules []*AccessControlSystemChildRuleInfo `json:"SystemChildRules,omitempty" name:"SystemChildRules"`
	// 是否是系统默认策略

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type VulDefenceEventDetail struct {

	// 漏洞CVEID

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 入侵状态

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 攻击源IP

	SourceIP *string `json:"SourceIP,omitempty" name:"SourceIP"`
	// 攻击源ip地址所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 容器ID

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 处理状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 攻击源端口

	SourcePort []*string `json:"SourcePort,omitempty" name:"SourcePort"`
	// 事件ID

	EventID *int64 `json:"EventID,omitempty" name:"EventID"`
	// 主机名称/超级节点名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机内网IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// Pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 危害描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 修复建议

	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`
	// 攻击包

	NetworkPayload *string `json:"NetworkPayload,omitempty" name:"NetworkPayload"`
	// 进程PID

	PID *int64 `json:"PID,omitempty" name:"PID"`
	// 进程主类名

	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`
	// 堆栈信息

	StackTrace *string `json:"StackTrace,omitempty" name:"StackTrace"`
	// 监听账号

	ServerAccount *string `json:"ServerAccount,omitempty" name:"ServerAccount"`
	// 监听端口

	ServerPort *string `json:"ServerPort,omitempty" name:"ServerPort"`
	// 进程路径

	ServerExe *string `json:"ServerExe,omitempty" name:"ServerExe"`
	// 进程命令行参数

	ServerArg *string `json:"ServerArg,omitempty" name:"ServerArg"`
	// 主机QUUID/超级节点ID

	QUUID *string `json:"QUUID,omitempty" name:"QUUID"`
	// 隔离状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// 	"NODE_DESTROYED"      //节点已销毁
	// 	"CONTAINER_EXITED"    //容器已退出
	// 	"CONTAINER_DESTROYED" //容器已销毁
	// 	"SHARED_HOST"         // 容器与主机共享网络
	// 	"RESOURCE_LIMIT"      //隔离操作资源超限
	// 	"UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 接口Url

	JNDIUrl *string `json:"JNDIUrl,omitempty" name:"JNDIUrl"`
	// rasp detail

	RaspDetail []*RaspInfo `json:"RaspDetail,omitempty" name:"RaspDetail"`
	// 超级节点子网名称

	NodeSubNetName *string `json:"NodeSubNetName,omitempty" name:"NodeSubNetName"`
	// 超级节点子网网段

	NodeSubNetCIDR *string `json:"NodeSubNetCIDR,omitempty" name:"NodeSubNetCIDR"`
	// pod ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 节点类型[NORMAL:普通节点|SUPER:超级节点]

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 超级节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 超级节点唯一ID

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 超级节点子网ID

	NodeSubNetID *string `json:"NodeSubNetID,omitempty" name:"NodeSubNetID"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// Namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 工作负载

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
}

type DescribeExportJobManageListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>ExportStatus- string -是否必填: 否 - 导出状态 RUNNING: 导出中 SUCCESS:导出完成 FAILURE:失败
	// </li>
	// <li>ExportSource- string -是否必填: 否 - 导出来源 LocalImage: 本地镜像 RegistryImage: 仓库镜像
	// </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段
	// InsertTime: 创建时间

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeExportJobManageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportJobManageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势列表

		List []*K8sApiAbnormalTendencyItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCheckModeRequest struct {
	*tchttp.BaseRequest

	// 要设置的集群ID列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 集群检查模式(正常模式"Cluster_Normal"、主动模式"Cluster_Actived"、不设置"Cluster_Unset")

	ClusterCheckMode *string `json:"ClusterCheckMode,omitempty" name:"ClusterCheckMode"`
	// 0不设置 1打开 2关闭

	ClusterAutoCheck *uint64 `json:"ClusterAutoCheck,omitempty" name:"ClusterAutoCheck"`
}

func (r *SetCheckModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCheckModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceEventDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞防御事件详细

		EventDetail *VulDefenceEventDetail `json:"EventDetail,omitempty" name:"EventDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefenceEventDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数
	// InnerNetAlarmShow- int - 是否必填：1 - 内网告警展示 0 - 不展示

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeReverseShellEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateImageRegistryTimingScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateImageRegistryTimingScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateImageRegistryTimingScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAssetImageAuthInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAssetImageAuthInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAssetImageAuthInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanTimeoutSettingRequest struct {
	*tchttp.BaseRequest

	// 设置类型0一键检测，1定时检测

	ScanType *uint64 `json:"ScanType,omitempty" name:"ScanType"`
}

func (r *DescribeVirusScanTimeoutSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanTimeoutSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceHostInfo struct {

	// 主机唯一id

	HostId *uint64 `json:"HostId,omitempty" name:"HostId"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机状态 ASSET_OFFLINE 回收  ASSET_ABNORMAL 离线
	// ASSET_NORMAL 正常

	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`
	// 镜像数量

	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
	// 容器数量

	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 是否为已选项

	Checked *bool `json:"Checked,omitempty" name:"Checked"`
	// 主机实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeAssetImageRiskListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 风险级别 1,2,3,4，</li>
	// <li>Behavior - String - 是否必填：否 - 风险行为 1,2,3,4</li>
	// <li>Type - String - 是否必填：否 - 风险类型  1,2,</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellRegexpWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 列表

		List []*RegexpRuleListItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellRegexpWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellRegexpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageDenyEventRequest struct {
	*tchttp.BaseRequest

	// 事件记录ID数组。Filters和EventIDSet二者选其一

	EventIDSet []*int64 `json:"EventIDSet,omitempty" name:"EventIDSet"`
	// 过滤条件。Filters和EventIDSet二者选其一
	// <li>EventType- String - 是否必填：否 -事件类型 EVENT_RISK:风险事件类型，EVENT_PRIVILEGE:特权。</li>
	// <li>DealBehavior- String - 是否必填：否 - 执行动作,BEHAVIOR_ALERT:告警，BEHAVIOR_HOLDUP_SUCCESSED:拦截。</li>
	// <li>RuleName- string - 是否必填：否 - 规则名称。</li>
	// <li>NodeName- string - 是否必填：否 - 节点名称。</li>
	// <li>NodeIP- string - 是否必填：否 - 内外IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 过滤条件中不删除的事件ID数组

	ExcludeEventIDSet []*int64 `json:"ExcludeEventIDSet,omitempty" name:"ExcludeEventIDSet"`
}

func (r *DeleteImageDenyEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageDenyEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryAutoAuthorizedRuleImageRequest struct {
	*tchttp.BaseRequest

	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 规则ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeImageRegistryAutoAuthorizedRuleImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryAutoAuthorizedRuleImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRisk struct {

	// 高危行为

	Behavior *uint64 `json:"Behavior,omitempty" name:"Behavior"`
	// 种类

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 解决方案

	InstructionContent *string `json:"InstructionContent,omitempty" name:"InstructionContent"`
}

type DescribeVulImageListRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>HostIP- string -是否必填: 否 - 内网IP</li>
	// <li>PublicIP- string -是否必填: 否 - 外网IP</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>HostName- string -是否必填: 否 - 主机名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEditWarningRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddEditWarningRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditWarningRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkFirewallPolicyDiscoverRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateNetworkFirewallPolicyDiscoverRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkFirewallPolicyDiscoverRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRiskListExportRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 风险级别 1,2,3,4，</li>
	// <li>Behavior - String - 是否必填：否 - 风险行为 1,2,3,4</li>
	// <li>Type - String - 是否必填：否 - 风险类型  1,2,</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *DescribeAssetImageRiskListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRiskListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageRegistryAutoAuthorizedRuleRequest struct {
	*tchttp.BaseRequest

	// 0关闭 1打开

	IsEnabled *uint64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 全部：ALL  自选项目空间：NAMESPACE 自选镜像仓库：REPO

	RuleRange *string `json:"RuleRange,omitempty" name:"RuleRange"`
	// 每天最大授权数

	MaxDailyCnt *int64 `json:"MaxDailyCnt,omitempty" name:"MaxDailyCnt"`
	// RuleRange=NAMESPACE/REPO需提供

	AutoAuthImageInfos []*AutoAuthImageInfo `json:"AutoAuthImageInfos,omitempty" name:"AutoAuthImageInfos"`
	// 过滤器。支持值：RegistryType/Namespace，只支持全匹配

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排除的条目。结合Filters使用，用于剔除应用过滤器Filters查询之后的结果。

	ExcludeSet []*AutoAuthImageInfo `json:"ExcludeSet,omitempty" name:"ExcludeSet"`
	// 规则ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 自动扫描开关

	AutoScanEnabled *uint64 `json:"AutoScanEnabled,omitempty" name:"AutoScanEnabled"`
	// 自动扫描范围

	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`
}

func (r *ModifyImageRegistryAutoAuthorizedRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageRegistryAutoAuthorizedRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRegistryConnDetectRequest struct {
	*tchttp.BaseRequest

	// 联通性检测的内容

	ConnDetectContent []*RegistryConnDetectParam `json:"ConnDetectContent,omitempty" name:"ConnDetectContent"`
	// 检测来源，list表示列表页面  modfiy表示编辑页面

	From *string `json:"From,omitempty" name:"From"`
}

func (r *ImageRegistryConnDetectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImageRegistryConnDetectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckRepeatAssetImageRegistryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否重复

		IsRepeat *bool `json:"IsRepeat,omitempty" name:"IsRepeat"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckRepeatAssetImageRegistryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRepeatAssetImageRegistryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageAuthorizedInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总共有效的镜像授权数

		TotalAuthorizedCnt *uint64 `json:"TotalAuthorizedCnt,omitempty" name:"TotalAuthorizedCnt"`
		// 已使用镜像授权数

		UsedAuthorizedCnt *uint64 `json:"UsedAuthorizedCnt,omitempty" name:"UsedAuthorizedCnt"`
		// 已开启扫描镜像数

		ScannedImageCnt *uint64 `json:"ScannedImageCnt,omitempty" name:"ScannedImageCnt"`
		// 未开启扫描镜像数

		NotScannedImageCnt *uint64 `json:"NotScannedImageCnt,omitempty" name:"NotScannedImageCnt"`
		// 本地未开启扫描镜像数

		NotScannedLocalImageCnt *uint64 `json:"NotScannedLocalImageCnt,omitempty" name:"NotScannedLocalImageCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageAuthorizedInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageAuthorizedInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessRuleInfo struct {

	// 策略id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// true:策略启用，false:策略禁用

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 生效镜像id，空数组代表全部镜像

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 用户策略的子策略数组

	ChildRules []*AbnormalProcessChildRuleInfo `json:"ChildRules,omitempty" name:"ChildRules"`
	// 策略名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 系统策略的子策略数组

	SystemChildRules []*AbnormalProcessSystemChildRuleInfo `json:"SystemChildRules,omitempty" name:"SystemChildRules"`
	// 是否是系统默认策略

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type DescribeAssetImageRegistryListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusScanConfigRequest struct {
	*tchttp.BaseRequest

	// 是否开启定期扫描

	EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`
	// 检测周期每隔多少天(1|3|7)

	Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`
	// 扫描开始时间

	BeginScanAt *string `json:"BeginScanAt,omitempty" name:"BeginScanAt"`
	// 超时时长(5~24h)

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// SCAN_NODE:扫描节点 SCAN_CONTAINER:扫描容器

	ScanRangeType *string `json:"ScanRangeType,omitempty" name:"ScanRangeType"`
	// 自选扫描范围的容器id或者节点id

	ScanIDs []*ScanRangeInfo `json:"ScanIDs,omitempty" name:"ScanIDs"`
	// 扫描路径

	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
	// 扫描路径模式：
	// SCAN_PATH_ALL：全部路径
	// SCAN_PATH_DEFAULT：默认路径
	// SCAN_PATH_USER_DEFINE：用户自定义路径
	//

	ScanPathMode *string `json:"ScanPathMode,omitempty" name:"ScanPathMode"`
	// true:包含路径 false:排除路径

	IsIncludePath *bool `json:"IsIncludePath,omitempty" name:"IsIncludePath"`
}

func (r *ModifyVirusScanConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusScanConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePeriodTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 定时任务的总量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 定时任务信息的列表

		PeriodTaskSet []*CompliancePeriodTask `json:"PeriodTaskSet,omitempty" name:"PeriodTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCompliancePeriodTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePeriodTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessSystemChildRuleInfo struct {

	// 子策略Id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 子策略状态，true为开启，false为关闭

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 策略模式,  RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截

	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`
	// 子策略检测的行为类型
	// PROXY_TOOL： 代理软件
	// TRANSFER_CONTROL：横向渗透
	// ATTACK_CMD： 恶意命令
	// REVERSE_SHELL：反弹shell
	// FILELESS：无文件程序执行
	// RISK_CMD：高危命令
	// ABNORMAL_CHILD_PROC: 敏感服务异常子进程启动

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 威胁等级，HIGH:高，MIDDLE:中，LOW:低

	RuleLevel *string `json:"RuleLevel,omitempty" name:"RuleLevel"`
}

type ReverseShellEventDescription struct {

	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 目标地址

	DstAddress *string `json:"DstAddress,omitempty" name:"DstAddress"`
	// 事件最后一次处理的时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type DescribeVulDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞详情信息

		VulInfo *VulDetailInfo `json:"VulInfo,omitempty" name:"VulInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecLogJoinObjectInfo struct {

	// 主机ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机状态

	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 接入状态(true:已接入  false:未接入)

	JoinState *bool `json:"JoinState,omitempty" name:"JoinState"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 集群主节点地址

	ClusterMainAddress *string `json:"ClusterMainAddress,omitempty" name:"ClusterMainAddress"`
}

type DescribeNetworkTopologyFlowSwitchListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkTopologyFlowSwitchListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyFlowSwitchListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkTopologyFlow struct {

	// 流量id
	//

	Id *string `json:"Id,omitempty" name:"Id"`
	// 流量时间
	//

	FlowTime *string `json:"FlowTime,omitempty" name:"FlowTime"`
	// 源ip
	//

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 源pod
	//

	SrcPod *string `json:"SrcPod,omitempty" name:"SrcPod"`
	// 源负载
	//

	SrcWorkLoad *string `json:"SrcWorkLoad,omitempty" name:"SrcWorkLoad"`
	// 源关联策略
	//

	SrcPolicy *uint64 `json:"SrcPolicy,omitempty" name:"SrcPolicy"`
	// 源标签
	//

	SrcLabel *string `json:"SrcLabel,omitempty" name:"SrcLabel"`
	// 源负载类型

	SrcWorkLoadKind *string `json:"SrcWorkLoadKind,omitempty" name:"SrcWorkLoadKind"`
	// 源命名空间

	SrcNamespace *string `json:"SrcNamespace,omitempty" name:"SrcNamespace"`
	// 目标ip
	//

	DestIp *string `json:"DestIp,omitempty" name:"DestIp"`
	// 目标pod
	//

	DestPod *string `json:"DestPod,omitempty" name:"DestPod"`
	// 目标负载
	//

	DestWorkLoad *string `json:"DestWorkLoad,omitempty" name:"DestWorkLoad"`
	// 目标负载类型
	//

	DestWorkLoadKind *string `json:"DestWorkLoadKind,omitempty" name:"DestWorkLoadKind"`
	// 目标命名空间

	DestNamespace *string `json:"DestNamespace,omitempty" name:"DestNamespace"`
	// 目标关联策略
	//

	DestPolicy *uint64 `json:"DestPolicy,omitempty" name:"DestPolicy"`
	// 目标标签
	//

	DestLabel *string `json:"DestLabel,omitempty" name:"DestLabel"`
	// 协议
	//

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 端口
	//

	Port *string `json:"Port,omitempty" name:"Port"`
	// 动作
	//

	Action *string `json:"Action,omitempty" name:"Action"`
}

type ImageRepoRegistryInfo struct {

	// 仓库id

	RegistryId *uint64 `json:"RegistryId,omitempty" name:"RegistryId"`
	// 仓库名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 仓库类型，列表：harbor、tcr

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 仓库url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 网络类型，列表：public

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 区域，列表：default

	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`
	// 仓库版本

	RegistryVersion *string `json:"RegistryVersion,omitempty" name:"RegistryVersion"`
	// 仓库连接状态连接状态，成功：success 失败：failed，待废弃，请使用ConnDetectStatus

	ConnectStatus *string `json:"ConnectStatus,omitempty" name:"ConnectStatus"`
	// 仓库连接错误信息，待废弃，请使用ConnDetectException

	ConnectMsg *string `json:"ConnectMsg,omitempty" name:"ConnectMsg"`
	// 联通性检测方式

	ConnDetectType *string `json:"ConnDetectType,omitempty" name:"ConnDetectType"`
	// 联通性检测主机数

	ConnDetectHostCount *uint64 `json:"ConnDetectHostCount,omitempty" name:"ConnDetectHostCount"`
	// 联通性检测详情

	ConnDetectDetail []*RegistryConnDetectResult `json:"ConnDetectDetail,omitempty" name:"ConnDetectDetail"`
	// tcr情况下的instance_id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 最近同步成功时间

	LatestSyncTime *string `json:"LatestSyncTime,omitempty" name:"LatestSyncTime"`
	// 同步状态

	SyncStatus *string `json:"SyncStatus,omitempty" name:"SyncStatus"`
	// 同步失败原因

	SyncFailReason *string `json:"SyncFailReason,omitempty" name:"SyncFailReason"`
	// 同步失败解决方案

	SyncSolution *string `json:"SyncSolution,omitempty" name:"SyncSolution"`
	// 同步失败信息

	SyncMessage *string `json:"SyncMessage,omitempty" name:"SyncMessage"`
}

type DeleteMaliciousConnectionWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaliciousConnectionWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaliciousConnectionWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`
		// 进程信息

		ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`
		// 父进程信息

		ParentProcessInfo *ProcessDetailBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`
		// 事件描述

		EventDetail *RiskSyscallEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`
		// 祖先进程信息

		AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 节点列表

		ClusterNodeList []*ClusterNodeInfo `json:"ClusterNodeList,omitempty" name:"ClusterNodeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeSystemVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待处理的事件总数

		UnhandledEventCount *int64 `json:"UnhandledEventCount,omitempty" name:"UnhandledEventCount"`
		// 待处理的恶意域名请求数

		RiskDnsEventCount *int64 `json:"RiskDnsEventCount,omitempty" name:"RiskDnsEventCount"`
		// 待处理的恶意IP请求数

		RiskIPEventCount *int64 `json:"RiskIPEventCount,omitempty" name:"RiskIPEventCount"`
		// 影响的容器总数

		EffectContainerCount *int64 `json:"EffectContainerCount,omitempty" name:"EffectContainerCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryAutoAuthorizedRuleImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		List []*AutoAuthImageInfo `json:"List,omitempty" name:"List"`
		// 数据量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRegistryAutoAuthorizedRuleImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryAutoAuthorizedRuleImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterWorkloadWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检测项名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// workload的白名单列表

		WhiteListWorkloads []*WhiteListWorkload `json:"WhiteListWorkloads,omitempty" name:"WhiteListWorkloads"`
		// workload类型白名单总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterWorkloadWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterWorkloadWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRiskTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本地镜像新增风险趋势信息列表

		ImageRiskTendencySet []*ImageRiskTendencyInfo `json:"ImageRiskTendencySet,omitempty" name:"ImageRiskTendencySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRiskTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRiskTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmNetworkFirewallPolicyRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略Id数组

	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ConfirmNetworkFirewallPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmNetworkFirewallPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpgradeAgentListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeUpgradeAgentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpgradeAgentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRuleDetailRequest struct {
	*tchttp.BaseRequest

	// 策略唯一id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 镜像id, 在添加白名单的时候使用

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAbnormalProcessRuleDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRuleDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallClusterContainerSecurityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallClusterContainerSecurityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallClusterContainerSecurityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyWorkLoadDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// pod个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// pod列表

		PodList []*NetworkTopologyPodList `json:"PodList,omitempty" name:"PodList"`
		// 命名空间

		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
		// workload名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 出站类型

		WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`
		// services

		Service *string `json:"Service,omitempty" name:"Service"`
		// workload标签

		Labels *string `json:"Labels,omitempty" name:"Labels"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkTopologyWorkLoadDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyWorkLoadDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSearchTemplateRequest struct {
	*tchttp.BaseRequest

	// 搜索模板

	SearchTemplate *SearchTemplate `json:"SearchTemplate,omitempty" name:"SearchTemplate"`
}

func (r *CreateSearchTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSearchTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageVirusExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为10000

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *CreateAssetImageVirusExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageVirusExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusAutoIsolateSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVirusAutoIsolateSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusAutoIsolateSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageAuthorizedResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageAuthorizedResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageAuthorizedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventEscapeImageDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一值

		UniqueKey *string `json:"UniqueKey,omitempty" name:"UniqueKey"`
		// 事件类型 MOUNT_SENSITIVE_PTAH : 敏感路径挂载
		// PRIVILEGE_CONTAINER_START:特权容器

		EventType *string `json:"EventType,omitempty" name:"EventType"`
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 操作时间

		OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
		// 镜像ID

		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
		// 镜像名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 影响容器数量

		ContainerCount *uint64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
		// 告警数量

		EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
		// 首次生成时间

		FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
		// 最近生成时间

		LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
		// 状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 解决方案

		Solution *string `json:"Solution,omitempty" name:"Solution"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventEscapeImageDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventEscapeImageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTaskResultSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResultSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddClusterWorkloadToWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加检测项白名单的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddClusterWorkloadToWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterWorkloadToWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterIngressYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Ingress格式字符串,base64编码

		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterIngressYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterIngressYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeESAggregationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ES聚合结果JSON

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeESAggregationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeESAggregationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>FileName - String - 是否必填：否 - 文件名称</li>
	// <li>FilePath - String - 是否必填：否 - 文件路径</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名称</li>
	// <li>ContainerName- String - 是否必填：是 - 容器名称</li>
	// <li>ContainerId- string - 是否必填：否 - 容器id</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称</li>
	// <li>ImageId- string - 是否必填：否 - 镜像id</li>
	// <li>IsRealTime- int - 是否必填：否 - 过滤是否实时监控数据</li>
	// <li>TaskId- string - 是否必填：否 - 任务ID</li>
	// <li>ContainerNetStatus - String -是否必填: 否 -  容器网络状态筛选 NORMAL ISOLATED ISOLATING RESTORING RESTORE_FAILED</li>
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>ContainerStatus - string -是否必填: 否 - 容器状态 RUNNING PAUSED STOPPED CREATED DESTROYED RESTARTING REMOVING</li>
	// <li>AutoIsolateMode - string -是否必填: 否 - 隔离方式 MANUAL AUTO</li>
	// <li>MD5 - string -是否必填: 否 - md5 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVirusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageHost struct {

	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
}

type ModifyClusterCheckTimerSettingsRequest struct {
	*tchttp.BaseRequest

	// 状态(true:开启 false:关闭)

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 集群ID列表 为空则标识全部

	ClusterIDs []*string `json:"ClusterIDs,omitempty" name:"ClusterIDs"`
	// 扫描周期天数 若是1 则表示每天 若是3 则表示每三天

	CycleDay *uint64 `json:"CycleDay,omitempty" name:"CycleDay"`
	// 扫描开始时间

	ScanTimeBegin *string `json:"ScanTimeBegin,omitempty" name:"ScanTimeBegin"`
	// 扫描截止时间

	ScanTimeEnd *string `json:"ScanTimeEnd,omitempty" name:"ScanTimeEnd"`
	// 扫描范围 0:全部集群 1:部分集群

	ScanScope *uint64 `json:"ScanScope,omitempty" name:"ScanScope"`
}

func (r *ModifyClusterCheckTimerSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterCheckTimerSettingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 系统调用名称列表

		SyscallNames []*string `json:"SyscallNames,omitempty" name:"SyscallNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSuperNodeExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSuperNodeExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSuperNodeExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSampleDetailRequest struct {
	*tchttp.BaseRequest

	// 文件MD5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
}

func (r *DescribeVirusAutoIsolateSampleDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReverseShellEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportJobResultRequest struct {
	*tchttp.BaseRequest

	// CreateExportComplianceStatusListJob返回的JobId字段的值

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeExportJobResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportJobResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInspectionReportRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInspectionReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInspectionReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulInfo struct {

	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 漏洞标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// CVSS V3分数

	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// CVE编号

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞子类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// 首次发现时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 最近发现时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 漏洞ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 影响本地镜像数

	LocalImageCount *int64 `json:"LocalImageCount,omitempty" name:"LocalImageCount"`
	// 影响容器数

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 影响仓库镜像数

	RegistryImageCount *int64 `json:"RegistryImageCount,omitempty" name:"RegistryImageCount"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 防御状态，NO_DEFENDED:未防御，DEFENDED:已防御

	DefenceStatus *string `json:"DefenceStatus,omitempty" name:"DefenceStatus"`
	// 漏洞防御主机范围: MANUAL:自选主机节点，ALL:全部

	DefenceScope *string `json:"DefenceScope,omitempty" name:"DefenceScope"`
	// 漏洞防御主机数量

	DefenceHostCount *int64 `json:"DefenceHostCount,omitempty" name:"DefenceHostCount"`
	// 已防御攻击次数

	DefendedCount *int64 `json:"DefendedCount,omitempty" name:"DefendedCount"`
}

type DescribeAccessControlRulesExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlRulesExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEditRiskSyscallWhiteListRequest struct {
	*tchttp.BaseRequest

	// 仅在添加事件白名单时候使用

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 增加或编辑白名单信。新增白名单时WhiteListInfo.id为空，编辑白名单WhiteListInfo.id不能为空.

	WhiteListInfo *RiskSyscallWhiteListInfo `json:"WhiteListInfo,omitempty" name:"WhiteListInfo"`
}

func (r *AddEditRiskSyscallWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditRiskSyscallWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrModifyMaliciousConnectionWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOrModifyMaliciousConnectionWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrModifyMaliciousConnectionWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPortListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口列表

		List []*PortInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetPortListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// db服务列表

		List []*ServiceInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetAppServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogDeliveryKafkaSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecLogDeliveryKafkaSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogDeliveryKafkaSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceTaskAssetSummaryRequest struct {
	*tchttp.BaseRequest

	// 资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产

	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`
}

func (r *DescribeComplianceTaskAssetSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceTaskAssetSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusMonitorConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启实时监控

		EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`
		// true:包含路径 false:排除路径

		IsIncludePath *bool `json:"IsIncludePath,omitempty" name:"IsIncludePath"`
		// 自选排除或扫描的地址

		ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
		// 扫描路径模式：
		// SCAN_PATH_ALL：全部路径
		// SCAN_PATH_DEFAULT：默认路径
		// SCAN_PATH_USER_DEFINE：用户自定义路径
		//

		ScanPathMode *string `json:"ScanPathMode,omitempty" name:"ScanPathMode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusMonitorConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusMonitorConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterIngressYamlRequest struct {
	*tchttp.BaseRequest

	// ingress名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群地域

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

func (r *DescribeClusterIngressYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterIngressYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulDefenceSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulDefenceSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEmergencyVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEmergencyVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEmergencyVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoAuthorizedRuleHostRequest struct {
	*tchttp.BaseRequest

	// 规则id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 需要返回的数量，默认为全部；

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序方式，asc，desc

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAutoAuthorizedRuleHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoAuthorizedRuleHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalEventInfo struct {

	// 命中规则名称

	MatchRuleName *string `json:"MatchRuleName,omitempty" name:"MatchRuleName"`
	// 命中规则类型

	MatchRuleType *string `json:"MatchRuleType,omitempty" name:"MatchRuleType"`
	// 告警等级

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群运行状态

	ClusterRunningStatus *string `json:"ClusterRunningStatus,omitempty" name:"ClusterRunningStatus"`
	// 初次生成时间

	FirstCreateTime *string `json:"FirstCreateTime,omitempty" name:"FirstCreateTime"`
	// 最近一次生成时间

	LastCreateTime *string `json:"LastCreateTime,omitempty" name:"LastCreateTime"`
	// 告警数量

	AlarmCount *uint64 `json:"AlarmCount,omitempty" name:"AlarmCount"`
	// 状态
	// "EVENT_UNDEAL":未处理
	// "EVENT_DEALED": 已处理
	// "EVENT_IGNORE": 忽略
	// "EVENT_DEL": 删除
	// "EVENT_ADD_WHITE": 加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 集群masterIP

	ClusterMasterIP *string `json:"ClusterMasterIP,omitempty" name:"ClusterMasterIP"`
	// k8s版本

	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
	// 运行时组件

	RunningComponent []*string `json:"RunningComponent,omitempty" name:"RunningComponent"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 建议

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 请求信息

	Info *string `json:"Info,omitempty" name:"Info"`
	// 规则ID

	MatchRuleID *string `json:"MatchRuleID,omitempty" name:"MatchRuleID"`
	// 高亮字段数组

	HighLightFields []*string `json:"HighLightFields,omitempty" name:"HighLightFields"`
	// 命中规则

	MatchRule *K8sApiAbnormalRuleScopeInfo `json:"MatchRule,omitempty" name:"MatchRule"`
}

type ModifyRiskSyscallStatusRequest struct {
	*tchttp.BaseRequest

	// 处理事件ids

	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
	// 标记事件的状态，
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyRiskSyscallStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskSyscallStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventTypeSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeEscapeEventTypeSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventTypeSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogKafkaUINRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecLogKafkaUINRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogKafkaUINRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceBenchmarkStandard struct {

	// 合规标准的ID

	StandardId *uint64 `json:"StandardId,omitempty" name:"StandardId"`
	// 合规标准的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 合规标准包含的数目

	PolicyItemCount *uint64 `json:"PolicyItemCount,omitempty" name:"PolicyItemCount"`
	// 是否启用此标准

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 标准的描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateClusterAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulTopRankingRequest struct {
	*tchttp.BaseRequest

	// 漏洞分类: SYSTEM:系统漏洞 WEB:web应用漏洞 EMERGENCY:应急漏洞

	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
}

func (r *DescribeVulTopRankingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulTopRankingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerInfo struct {

	// 容器id

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 容器运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 运行用户

	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
	// 命令行

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// CPU使用率 *1000

	CPUUsage *uint64 `json:"CPUUsage,omitempty" name:"CPUUsage"`
	// 内存使用 kb

	RamUsage *uint64 `json:"RamUsage,omitempty" name:"RamUsage"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像id

	POD *string `json:"POD,omitempty" name:"POD"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`
	// 网络子状态

	NetSubStatus *string `json:"NetSubStatus,omitempty" name:"NetSubStatus"`
	// 隔离来源

	IsolateSource *string `json:"IsolateSource,omitempty" name:"IsolateSource"`
	// 隔离时间

	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
	// 超级节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// podip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 节点类型:节点类型：NORMAL普通节点、SUPER超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 超级节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 所属Pod的CPU

	PodCpu *int64 `json:"PodCpu,omitempty" name:"PodCpu"`
	// 所属Pod的内存

	PodMem *int64 `json:"PodMem,omitempty" name:"PodMem"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// pod uid

	PodUid *string `json:"PodUid,omitempty" name:"PodUid"`
}

type DescribeCompliancePolicyItemAffectedAssetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回各检测项所影响的资产的列表。

		AffectedAssetList []*ComplianceAffectedAsset `json:"AffectedAssetList,omitempty" name:"AffectedAssetList"`
		// 检测项影响的资产的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCompliancePolicyItemAffectedAssetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePolicyItemAffectedAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 访问控制事件数组

		EventSet []*AccessControlEventInfo `json:"EventSet,omitempty" name:"EventSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageRegistryScanTaskRequest struct {
	*tchttp.BaseRequest

	// 是否扫描全部镜像

	All *bool `json:"All,omitempty" name:"All"`
	// 扫描的镜像列表

	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`
	// 扫描类型数组

	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`
	// 扫描的镜像列表

	Id []*uint64 `json:"Id,omitempty" name:"Id"`
	// 过滤条件

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 不需要扫描的镜像列表, 与Filters配合使用

	ExcludeImageList []*uint64 `json:"ExcludeImageList,omitempty" name:"ExcludeImageList"`
	// 是否仅扫描各repository最新版的镜像, 与Filters配合使用

	OnlyScanLatest *bool `json:"OnlyScanLatest,omitempty" name:"OnlyScanLatest"`
}

func (r *CreateAssetImageRegistryScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageRegistryScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetDetailInfoRequest struct {
	*tchttp.BaseRequest

	// 客户资产ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
}

func (r *DescribeComplianceAssetDetailInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetDetailInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageDenyEventTendency struct {

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
}

type AddCompliancePolicyItemToWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCompliancePolicyItemToWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCompliancePolicyItemToWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessControlRulesRequest struct {
	*tchttp.BaseRequest

	// 策略的ids

	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`
}

func (r *DeleteAccessControlRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessControlRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HybridClusterItem struct {

	// 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群主Ip列表,以;隔开

	MasterIpList *string `json:"MasterIpList,omitempty" name:"MasterIpList"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type AddComplianceAssetPolicySetToWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddComplianceAssetPolicySetToWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddComplianceAssetPolicySetToWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像列表

		List []*ImageHost `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVirusScanConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrModifyImageDenyRuleRequest struct {
	*tchttp.BaseRequest

	// 规则RuleID，编辑时必需

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则类型 RULE_RISK：风险， RULE_PRIVILEGE：特权

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 是否对全部扫描镜像生效。0:全选镜像，1:自选镜像

	IsEffectAllImage *int64 `json:"IsEffectAllImage,omitempty" name:"IsEffectAllImage"`
	// 启用状态 0:开启，1:关闭

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 规则描述

	RuleDescription *string `json:"RuleDescription,omitempty" name:"RuleDescription"`
	// 漏洞，0:未选中，1:选中。RuleType=RULE_RISK时Vul、Virus和Risk至少三者选其一

	Vul *int64 `json:"Vul,omitempty" name:"Vul"`
	// cve编号，Vul=1时CVEIDSet、ComponentSet和RiskLevelSet三者选其一

	CVEIDSet []*string `json:"CVEIDSet,omitempty" name:"CVEIDSet"`
	// 组件编号，Vul=1时CVEIDSet、ComponentSet、VulClassSet和RiskLevelSet四者选其一

	ComponentSet []*string `json:"ComponentSet,omitempty" name:"ComponentSet"`
	// 漏洞分类，Vul=1时CVEIDSet、ComponentSet、VulClassSet和RiskLevelSet四者选其一

	VulClassSet []*string `json:"VulClassSet,omitempty" name:"VulClassSet"`
	// 漏洞等级，Vul=1时CVEIDSet、ComponentSet、VulClassSet和RiskLevelSet四者选其一

	VulLevelSet []*string `json:"VulLevelSet,omitempty" name:"VulLevelSet"`
	// 漏洞标签，Virus=1时VirusMD5Set、ComponentSet和RiskLevelSet三者选其一

	VulLabelSet []*string `json:"VulLabelSet,omitempty" name:"VulLabelSet"`
	// 木马，0:未选中，1:选中。RuleType=RULE_RISK时Vul、Virus和Risk至少三者选其一

	Virus *int64 `json:"Virus,omitempty" name:"Virus"`
	// 木马md5列表，Virus=1时VirusMD5Set、VirusLevelSet和RiskLevelSet三者选其一

	VirusMD5Set []*string `json:"VirusMD5Set,omitempty" name:"VirusMD5Set"`
	// 木马等级，Virus=1时VirusMD5Set、VirusLevelSet和RiskLevelSet三者选其一

	VirusLevelSet []*string `json:"VirusLevelSet,omitempty" name:"VirusLevelSet"`
	// 病毒名，Virus=1时VirusMD5Set和RiskLevelSet二者选其一

	VirusName []*string `json:"VirusName,omitempty" name:"VirusName"`
	// 敏感信息，0:未选中，1:选中。RuleType=RULE_RISK时Vul、Virus和Risk至少三者选其一

	Risk *int64 `json:"Risk,omitempty" name:"Risk"`
	// 敏感等级，Risk=1时RiskType和RiskLevelSet二者选其一

	RiskLevelSet []*string `json:"RiskLevelSet,omitempty" name:"RiskLevelSet"`
	// 敏感信息分类，Risk=1时RiskType和RiskLevelSet二者选其一

	RiskType []*string `json:"RiskType,omitempty" name:"RiskType"`
	// 特权启动 0:不允许，1:允许

	PrivilegeRun *int64 `json:"PrivilegeRun,omitempty" name:"PrivilegeRun"`
	// 特权类型，RuleType=RULE_PRIVILEGE时特权容器拦截时必需

	PrivilegeDetail []*string `json:"PrivilegeDetail,omitempty" name:"PrivilegeDetail"`
	// 自选镜像数组，非全部镜像时必需

	EffectImageSet []*string `json:"EffectImageSet,omitempty" name:"EffectImageSet"`
	// 多少天后生效，0即立即生效

	EffectDay *uint64 `json:"EffectDay,omitempty" name:"EffectDay"`
}

func (r *AddOrModifyImageDenyRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrModifyImageDenyRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAbnormalProcessStatusRequest struct {
	*tchttp.BaseRequest

	// 处理事件ids

	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
	// 标记事件的状态，
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAbnormalProcessStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAbnormalProcessStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateComplianceTaskRequest struct {
	*tchttp.BaseRequest

	// 指定要扫描的资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产
	// AssetTypeSet, PolicySetId, PeriodTaskId三个参数，必须要给其中一个参数填写有效的值。

	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`
	// 按照策略集ID指定的策略执行合规检查。

	PolicySetId *uint64 `json:"PolicySetId,omitempty" name:"PolicySetId"`
	// 按照定时任务ID指定的策略执行合规检查。

	PeriodTaskId *uint64 `json:"PeriodTaskId,omitempty" name:"PeriodTaskId"`
}

func (r *CreateComplianceTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateComplianceTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogDeliveryKafkaOptionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceList []*CKafkaInstanceInfo `json:"InstanceList,omitempty" name:"InstanceList"`
		// 地域列表

		RegionList []*RegionInfo `json:"RegionList,omitempty" name:"RegionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogDeliveryKafkaOptionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogDeliveryKafkaOptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 逃逸事件数组

		EventSet []*EscapeEventInfo `json:"EventSet,omitempty" name:"EventSet"`
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeEventInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetClusterListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>Status - string  - 是否必填: 否 -集群状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段。

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAssetClusterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetClusterListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceEventRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_DEFENDED：已防御</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  入侵状态，防御成功：EVENT_DEFENDED，尝试攻击：EVENT_ATTACK </li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>VulName- string - 是否必填：否 - 漏洞名称。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号。</li>
	// <li>SourceIP- string - 是否必填：否 - 攻击源IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：事件数量：EventCount

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulDefenceEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusManualScanEstimateTimeoutResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预估超时时间(h)

		Timeout *float64 `json:"Timeout,omitempty" name:"Timeout"`
		// 单台主机并行扫描容器数

		ContainerScanConcurrencyCount *uint64 `json:"ContainerScanConcurrencyCount,omitempty" name:"ContainerScanConcurrencyCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusManualScanEstimateTimeoutResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusManualScanEstimateTimeoutResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusEstimateScanTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预估超时时间(h)

		Timeout *float64 `json:"Timeout,omitempty" name:"Timeout"`
		// 单节点并行扫描容器数

		ContainerScanConcurrencyCount *uint64 `json:"ContainerScanConcurrencyCount,omitempty" name:"ContainerScanConcurrencyCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusEstimateScanTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusEstimateScanTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ContainerName - String - 是否必填：否 - 容器名称模糊搜索</li>
	// <li>Status - String - 是否必填：否 - 容器运行状态筛选，0："created",1："running", 2："paused", 3："restarting", 4："removing", 5："exited", 6："dead" </li>
	// <li>Runas - String - 是否必填：否 - 运行用户筛选</li>
	// <li>ImageName- String - 是否必填：否 - 镜像名称搜索</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>OrderBy - String 是否必填：否 -排序字段，支持：cpu_usage, mem_usage的动态排序 ["cpu_usage","+"]  '+'升序、'-'降序</li>
	// <li>NetStatus - String -是否必填: 否 -  容器网络状态筛选 normal isolated isolating isolate_failed restoring restore_failed</li>
	// <li>PodID - String -是否必填: 否 - PodID筛选</li>
	// <li>NodeUniqueID - String -是否必填: 否 - SuperNode筛选</li>
	// <li>PodUid - String -是否必填: 否 - Pod筛选</li>
	// <li>PodIP - String -是否必填: 否 - PodIP筛选</li>
	// <li>NodeType - String -是否必填: 否 - 节点类型筛选:NORMAL:普通节点;SUPER:超级节点</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetContainerListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetPolicyItemListRequest struct {
	*tchttp.BaseRequest

	// 客户资产的ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 要获取的数据量，默认为10，最大为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤器列表。Name字段支持
	// RiskLevel

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceAssetPolicyItemListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetPolicyItemListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalEventListItem struct {

	// 事件ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 命中规则类型

	MatchRuleType *string `json:"MatchRuleType,omitempty" name:"MatchRuleType"`
	// 威胁等级

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群运行状态

	ClusterRunningStatus *string `json:"ClusterRunningStatus,omitempty" name:"ClusterRunningStatus"`
	// 初次生成时间

	FirstCreateTime *string `json:"FirstCreateTime,omitempty" name:"FirstCreateTime"`
	// 最近一次生成时间

	LastCreateTime *string `json:"LastCreateTime,omitempty" name:"LastCreateTime"`
	// 告警数量

	AlarmCount *uint64 `json:"AlarmCount,omitempty" name:"AlarmCount"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 规则类型

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 描述信息

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 解决方案

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 命中规则

	MatchRule *K8sApiAbnormalRuleScopeInfo `json:"MatchRule,omitempty" name:"MatchRule"`
}

type CreateVulDefenceEventExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_DEFENDED：已防御</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  入侵状态，防御成功：EVENT_DEFENDED，尝试攻击：EVENT_ATTACK </li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>VulName- string - 是否必填：否 - 漏洞名称。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号。</li>
	// <li>SourceIP- string - 是否必填：否 - 攻击源IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，最大值为100000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：事件数量：EventCount

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateVulDefenceEventExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulDefenceEventExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreVul struct {

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 忽略的镜像ID，空表示全部

	ImageIDs []*string `json:"ImageIDs,omitempty" name:"ImageIDs"`
	// 当有镜像时
	// 镜像类型: LOCAL 本地镜像 REGISTRY 仓库镜像

	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
}

type ModifyVirusFileStatusRequest struct {
	*tchttp.BaseRequest

	// 处理事件id

	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
	// 标记事件的状态，
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//     EVENT_DEL:事件删除
	//     EVENT_ADD_WHITE:事件加白
	//     EVENT_PENDING: 事件待处理
	// 	EVENT_ISOLATE_CONTAINER: 隔离容器
	// 	EVENT_RESOTRE_CONTAINER: 恢复容器

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否后续自动隔离相同MD5文件

	AutoIsolate *bool `json:"AutoIsolate,omitempty" name:"AutoIsolate"`
}

func (r *ModifyVirusFileStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusFileStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTcrRegistryDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 更新错误提示

		UpdateErr *string `json:"UpdateErr,omitempty" name:"UpdateErr"`
		// 仓库ID

		RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTcrRegistryDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTcrRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryImageLayerDetailRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 风险级别筛选，</li>
	// <li>RiskType - String - 是否必填：否 - 风险类型名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryImageLayerDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryImageLayerDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterCheckItem struct {

	// 唯一的检测项的ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 风险项的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 检测项详细描述。

	ItemDetail *string `json:"ItemDetail,omitempty" name:"ItemDetail"`
	// 威胁等级。严重Serious,高危High,中危Middle,提示Hint

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 检查对象、风险对象.Runc,Kubelet,Containerd,Pods

	RiskTarget *string `json:"RiskTarget,omitempty" name:"RiskTarget"`
	// 风险类别,漏洞风险CVERisk,配置风险ConfigRisk

	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
	// 检测项所属的风险类型,权限提升:PrivilegePromotion,拒绝服务:RefuseService,目录穿越:DirectoryEscape,未授权访问:UnauthorizedAccess,权限许可和访问控制问题:PrivilegeAndAccessControl,敏感信息泄露:SensitiveInfoLeak

	RiskAttribute *string `json:"RiskAttribute,omitempty" name:"RiskAttribute"`
	// 风险特征,Tag.存在EXP:ExistEXP,存在POD:ExistPOC,无需重启:NoNeedReboot, 服务重启:ServerRestart,远程信息泄露:RemoteInfoLeak,远程拒绝服务:RemoteRefuseService,远程利用:RemoteExploit,远程执行:RemoteExecute

	RiskProperty *string `json:"RiskProperty,omitempty" name:"RiskProperty"`
	// CVE编号

	CVENumber *string `json:"CVENumber,omitempty" name:"CVENumber"`
	// 披露时间

	DiscoverTime *string `json:"DiscoverTime,omitempty" name:"DiscoverTime"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// CVSS信息,用于画图

	CVSS *string `json:"CVSS,omitempty" name:"CVSS"`
	// CVSS分数

	CVSSScore *string `json:"CVSSScore,omitempty" name:"CVSSScore"`
	// 参考连接

	RelateLink *string `json:"RelateLink,omitempty" name:"RelateLink"`
	// 影响类型，为Node或者Workload

	AffectedType *string `json:"AffectedType,omitempty" name:"AffectedType"`
	// 受影响的版本信息

	AffectedVersion *string `json:"AffectedVersion,omitempty" name:"AffectedVersion"`
	// 忽略的资产数量

	IgnoredAssetNum *int64 `json:"IgnoredAssetNum,omitempty" name:"IgnoredAssetNum"`
	// 是否忽略该检测项

	IsIgnored *bool `json:"IsIgnored,omitempty" name:"IsIgnored"`
	// 受影响评估

	RiskAssessment *string `json:"RiskAssessment,omitempty" name:"RiskAssessment"`
}

type ImageInfo struct {

	// 镜像id

	ImageDigest *string `json:"ImageDigest,omitempty" name:"ImageDigest"`
	// 仓库类型

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 镜像仓库地址

	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像tag

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 强制扫描

	Force *string `json:"Force,omitempty" name:"Force"`
}

type DescribeNetworkTopologyFlowListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络流量列表
		//

		FlowList []*NetworkTopologyFlow `json:"FlowList,omitempty" name:"FlowList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkTopologyFlowListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyFlowListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageProgress struct {

	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 仓库类型

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 镜像仓库地址

	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 仓库名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像tag

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 镜像扫描状态

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 镜像cve扫描进度

	CveProgress *uint64 `json:"CveProgress,omitempty" name:"CveProgress"`
	// 镜像敏感扫描进度

	RiskProgress *uint64 `json:"RiskProgress,omitempty" name:"RiskProgress"`
	// 镜像木马扫描进度

	VirusProgress *uint64 `json:"VirusProgress,omitempty" name:"VirusProgress"`
}

type DescribeExportJobManageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务列表

		List []*ExportJobInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportJobManageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportJobManageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSuperNodeInstallTaskExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSuperNodeInstallTaskExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSuperNodeInstallTaskExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileAttributeInfo struct {

	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件类型

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// 文件大小(字节)

	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`
	// 文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 文件创建时间

	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`
	// 最近被篡改文件创建时间

	LatestTamperedFileMTime *string `json:"LatestTamperedFileMTime,omitempty" name:"LatestTamperedFileMTime"`
	// 新文件内容

	NewFile *string `json:"NewFile,omitempty" name:"NewFile"`
	// 新旧文件的差异

	FileDiff *string `json:"FileDiff,omitempty" name:"FileDiff"`
}

type DescribeEscapeEventInfoRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,
	// Status：状态(EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略)
	// EventType: 事件类型(MOUNT_SENSITIVE_PTAH:敏感路径挂载 PRIVILEGE_CONTAINER_START:特权容器 PRIVILEGE:提权事件
	//     ESCAPE_VUL_OCCURRED:逃逸漏洞利用 ESCAPE_DOCKER_API:访问Docker API接口逃逸 ESCAPE_TAMPER_SENSITIVE_FILE:篡改敏感文件逃逸 ESCAPE_CGROUPS:利用cgroup机制逃逸)
	// ContainerNetStatus: 容器隔离状态 (NORMAL:正常 ISOLATED:已隔离 ISOLATE_FAILED:隔离失败 ISOLATE_FAILED:解除隔离失败 RESTORING:解除隔离中 ISOLATING:隔离中)
	// ContainerStatus: 容器状态(CREATED:已创建 RUNNING:正常运行 PAUSED:暂停运行 STOPPED:停止运行 RESTARTING:重启中 REMOVING:迁移中 DEAD:DEAD UNKNOWN：未知 DESTROYED:已销毁)
	// ForeignUniqueKey:镜像ID及事件类型唯一值

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEscapeEventInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNewPurchaseInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否可购买
		// (标识是否可以新购，如果已购买则不允许走新购逻辑)

		CanBuy *bool `json:"CanBuy,omitempty" name:"CanBuy"`
		// 需要购买的核心数

		CoresCnt *uint64 `json:"CoresCnt,omitempty" name:"CoresCnt"`
		// 需要购买的主机数

		HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`
		// 弹性计费上限

		MaxPostPayCoresCnt *uint64 `json:"MaxPostPayCoresCnt,omitempty" name:"MaxPostPayCoresCnt"`
		// 待扫描本地镜像数

		UnScannedLocalImageCnt *uint64 `json:"UnScannedLocalImageCnt,omitempty" name:"UnScannedLocalImageCnt"`
		// 待扫描仓库镜像数

		UnScannedRepositoryImageCnt *uint64 `json:"UnScannedRepositoryImageCnt,omitempty" name:"UnScannedRepositoryImageCnt"`
		// 待使用镜像授权数

		UnUsedAuthorizedImageCnt *uint64 `json:"UnUsedAuthorizedImageCnt,omitempty" name:"UnUsedAuthorizedImageCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNewPurchaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewPurchaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEditReverseShellWhiteListRequest struct {
	*tchttp.BaseRequest

	// 增加或编辑白名单信息。新增白名单时WhiteListInfo.id为空，编辑白名单WhiteListInfo.id不能为空。

	WhiteListInfo *ReverseShellWhiteListInfo `json:"WhiteListInfo,omitempty" name:"WhiteListInfo"`
	// 仅在添加事件白名单时候使用

	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *AddEditReverseShellWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditReverseShellWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求事件列表

		List []*RiskDnsEventInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmergencyVulExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEmergencyVulExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEmergencyVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecTendencyEventInfo struct {

	// 趋势列表

	EventSet []*RunTimeTendencyInfo `json:"EventSet,omitempty" name:"EventSet"`
	// 事件类型：
	// ET_ESCAPE : 容器逃逸
	// ET_REVERSE_SHELL: 反弹shell
	// ET_RISK_SYSCALL:高危系统调用
	// ET_ABNORMAL_PROCESS: 异常进程
	// ET_ACCESS_CONTROL 文件篡改
	// ET_VIRUS 木马事件
	// ET_MALICIOUS_CONNECTION 恶意外连事件

	EventType *string `json:"EventType,omitempty" name:"EventType"`
}

type RemoveLocalStorageItemRequest struct {
	*tchttp.BaseRequest

	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *RemoveLocalStorageItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveLocalStorageItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogJoinTypeListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecLogJoinTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogJoinTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompliancePeriodTaskRule struct {

	// 执行的频率（几天一次），取值为：1,3,7。

	Frequency *uint64 `json:"Frequency,omitempty" name:"Frequency"`
	// 在这天的什么时间执行，格式为：HH:mm:SS。

	ExecutionTime *string `json:"ExecutionTime,omitempty" name:"ExecutionTime"`
	// 是否开启

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type CleanTrialRequest struct {
	*tchttp.BaseRequest
}

func (r *CleanTrialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CleanTrialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveAssetImageRegistryRegistryDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveAssetImageRegistryRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateK8sApiAbnormalRuleExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateK8sApiAbnormalRuleExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateK8sApiAbnormalRuleExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryRegistryListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段
	// IsAuthorized是否授权，取值全部all，未授权0，已授权1

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式，asc，desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageRegistryRegistryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRegistryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSuperNodeInstallTaskListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// Status- String - 是否必填：否 - 状态 INSTALLING:安装中; INSTALLATIONFAILURE:安装失败; INSTALLATIONSUCCESS:安装成功 PARTIALINSTALLATIONFAILURE:部分安装失败
	// Source- String - 是否必填：否 - 安装来源: TCSS:容器安全服务 EKS:容器服务
	// TimeRange- String - 是否必填：否 - 安装时间范围,传一个时间默认为开始时间,两个为开始时间和结束时间

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeSuperNodeInstallTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSuperNodeInstallTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRiskSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageRiskSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRiskSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEscapeWhiteListRequest struct {
	*tchttp.BaseRequest

	// 加白名单事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸

	EventType []*string `json:"EventType,omitempty" name:"EventType"`
	// 白名单记录ID

	IDSet []*int64 `json:"IDSet,omitempty" name:"IDSet"`
}

func (r *ModifyEscapeWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEscapeWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNetworkFirewallPolicyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSuperNodeAgentInstallJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSuperNodeAgentInstallJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSuperNodeAgentInstallJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReverseShellWhiteListsExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateReverseShellWhiteListsExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateReverseShellWhiteListsExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageScanSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关

		Enable *bool `json:"Enable,omitempty" name:"Enable"`
		// 扫描时刻(完整时间;后端按0时区解析时分秒)

		ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
		// 扫描间隔

		ScanPeriod *uint64 `json:"ScanPeriod,omitempty" name:"ScanPeriod"`
		// 扫描木马

		ScanVirus *bool `json:"ScanVirus,omitempty" name:"ScanVirus"`
		// 扫描敏感信息

		ScanRisk *bool `json:"ScanRisk,omitempty" name:"ScanRisk"`
		// 扫描漏洞

		ScanVul *bool `json:"ScanVul,omitempty" name:"ScanVul"`
		// 扫描全部镜像

		All *bool `json:"All,omitempty" name:"All"`
		// 自定义扫描镜像

		Images []*string `json:"Images,omitempty" name:"Images"`
		// 镜像是否存在运行中的容器

		ContainerRunning *bool `json:"ContainerRunning,omitempty" name:"ContainerRunning"`
		// 扫描范围 0 全部授权镜像，1自选镜像，2 推荐扫描

		ScanScope *uint64 `json:"ScanScope,omitempty" name:"ScanScope"`
		// 扫描结束时间 02:00 时分

		ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageScanSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVirusListExportRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>RiskLevel - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 列表支持字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *DescribeAssetImageVirusListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVirusListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 异常进程策略信息列表

		RuleSet []*RuleBaseInfo `json:"RuleSet,omitempty" name:"RuleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterWorkloadsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群workload信息

		ClusterWorkloadList []*ClusterWorkloadInfo `json:"ClusterWorkloadList,omitempty" name:"ClusterWorkloadList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterWorkloadsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterWorkloadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVirusScanAgainRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 需要扫描的容器id集合

	ContainerIds []*string `json:"ContainerIds,omitempty" name:"ContainerIds"`
	// 是否是扫描全部超时的

	TimeoutAll *bool `json:"TimeoutAll,omitempty" name:"TimeoutAll"`
	// 重新设置的超时时长

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
}

func (r *CreateVirusScanAgainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVirusScanAgainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePolicyItemIgnoredAssetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回各检测项所影响的资产的列表。

		IgnoredAssetList []*ComplianceIgnoredAsset `json:"IgnoredAssetList,omitempty" name:"IgnoredAssetList"`
		// 检测项影响的资产的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCompliancePolicyItemIgnoredAssetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePolicyItemIgnoredAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeysLocalStorageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 键列表

		Keys []*string `json:"Keys,omitempty" name:"Keys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *KeysLocalStorageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KeysLocalStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessControlRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessControlRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalRuleListItem struct {

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则类型
	// RT_SYSTEM 系统规则
	// RT_USER 用户自定义

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 受影响集群总数

	EffectClusterCount *uint64 `json:"EffectClusterCount,omitempty" name:"EffectClusterCount"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 编辑账号

	OprUin *string `json:"OprUin,omitempty" name:"OprUin"`
	// 状态

	Status *bool `json:"Status,omitempty" name:"Status"`
}

type ModifyEscapeEventStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEscapeEventStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEscapeEventStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmergencyVulExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateEmergencyVulExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEmergencyVulExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 连接错误信息

		HealthCheckErr *string `json:"HealthCheckErr,omitempty" name:"HealthCheckErr"`
		// 名称错误信息

		NameRepeatErr *string `json:"NameRepeatErr,omitempty" name:"NameRepeatErr"`
		// 仓库唯一id

		RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAssetImageRegistryRegistryDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAssetImageRegistryRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallWhiteListsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeRiskSyscallWhiteListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallWhiteListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanComplianceScanFailedAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回重新检测任务的ID。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanComplianceScanFailedAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanComplianceScanFailedAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePolicyItemAffectedAssetListRequest struct {
	*tchttp.BaseRequest

	// DescribeComplianceTaskPolicyItemSummaryList返回的CustomerPolicyItemId，表示检测项的ID。

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// Name - String
	// Name 可取值：NodeName, CheckResult

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCompliancePolicyItemAffectedAssetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePolicyItemAffectedAssetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞总数量

		VulTotalCount *int64 `json:"VulTotalCount,omitempty" name:"VulTotalCount"`
		// 严重及高危漏洞数量

		SeriousVulCount *int64 `json:"SeriousVulCount,omitempty" name:"SeriousVulCount"`
		// 重点关注漏洞数量

		SuggestVulCount *int64 `json:"SuggestVulCount,omitempty" name:"SuggestVulCount"`
		// 有Poc或者Exp的漏洞数量

		PocExpLevelVulCount *int64 `json:"PocExpLevelVulCount,omitempty" name:"PocExpLevelVulCount"`
		// 有远程Exp的漏洞数量

		RemoteExpLevelVulCount *int64 `json:"RemoteExpLevelVulCount,omitempty" name:"RemoteExpLevelVulCount"`
		// 受严重或高危漏洞影响的最新版本镜像数

		SeriousVulNewestImageCount *int64 `json:"SeriousVulNewestImageCount,omitempty" name:"SeriousVulNewestImageCount"`
		// 系统漏洞重点关注数

		SystemVulnerabilityFocusCount *int64 `json:"SystemVulnerabilityFocusCount,omitempty" name:"SystemVulnerabilityFocusCount"`
		// web漏洞重点关注数

		WebVulnerabilityFocusCount *int64 `json:"WebVulnerabilityFocusCount,omitempty" name:"WebVulnerabilityFocusCount"`
		// 受影响本地镜像数

		SeriousVulnerabilityLocalImageCount *int64 `json:"SeriousVulnerabilityLocalImageCount,omitempty" name:"SeriousVulnerabilityLocalImageCount"`
		// 受影响仓库镜像数

		SeriousVulnerabilityRegistryImageCount *int64 `json:"SeriousVulnerabilityRegistryImageCount,omitempty" name:"SeriousVulnerabilityRegistryImageCount"`
		// 应急漏洞数量

		EmergencyVulnerabilityCount *int64 `json:"EmergencyVulnerabilityCount,omitempty" name:"EmergencyVulnerabilityCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallClusterListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群的详细信息

		ClusterInfoList []*NetworkClusterInfoItem `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallClusterListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogKafkaUINRequest struct {
	*tchttp.BaseRequest

	// 目标UIN

	DstUIN *string `json:"DstUIN,omitempty" name:"DstUIN"`
}

func (r *ModifySecLogKafkaUINRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogKafkaUINRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkFirewallUndoPublishRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略Id数组

	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *CreateNetworkFirewallUndoPublishRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkFirewallUndoPublishRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEscapeRegexpRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEscapeRegexpRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEscapeRegexpRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodesRequest struct {
	*tchttp.BaseRequest

	// 集群Id,不输入表示查询所有

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name 可取值：
	// DefendStatus（防护状态）:
	// 	Defended 已防护
	// 	UnDefended 未防护
	// AgentStatus (容器agent状态):
	//  	OFFLINE 离线
	//  	ONLINE 在线
	//  	UNINSTALL 未安装
	// InstanceState (节点状态):
	//   	Running 运行中
	//   	Ready 准备
	//   	Notready 未准备好
	//   	Initializing 初始化
	//   	Failed 失败
	//   	Error 错误
	// InstanceRole (节点角色)
	//     WORKER 工作节点
	//     MASTER_ETCD 主节点
	//     SUPER 超级节点

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClusterNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousConnectionBlackListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求白名单总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 恶意请求白名单列表

		List []*MaliciousConnectionRuleInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaliciousConnectionBlackListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionBlackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpgradeAgentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 升级终端列表

		List []*UpgradeAgentInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpgradeAgentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpgradeAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSyncLastTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产最近同步时间

		AssetSyncLastTime *string `json:"AssetSyncLastTime,omitempty" name:"AssetSyncLastTime"`
		// 任务状态
		// PENDING:待处理
		// PROCESSING:处理中
		// PROCESSED:已完成

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 任务进度(百分比)

		TaskProcess *int64 `json:"TaskProcess,omitempty" name:"TaskProcess"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSyncLastTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncLastTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkAuditRecord struct {

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 操作人

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 策略名

	NetworkPolicyName *string `json:"NetworkPolicyName,omitempty" name:"NetworkPolicyName"`
	// 操作时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
	// 操作人appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 操作人uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 策略id

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type NetworkClusterInfoItem struct {

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 集群操作系统

	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 集群网络插件

	NetworkPolicyPlugin *string `json:"NetworkPolicyPlugin,omitempty" name:"NetworkPolicyPlugin"`
	// 集群状态

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 总策略数量

	TotalRuleCount *int64 `json:"TotalRuleCount,omitempty" name:"TotalRuleCount"`
	// 已开启策略数量

	EnableRuleCount *int64 `json:"EnableRuleCount,omitempty" name:"EnableRuleCount"`
	// 集群网络插件状态，正常：Running 不正常：Error

	NetworkPolicyPluginStatus *string `json:"NetworkPolicyPluginStatus,omitempty" name:"NetworkPolicyPluginStatus"`
	// 集群网络插件错误信息

	NetworkPolicyPluginError *string `json:"NetworkPolicyPluginError,omitempty" name:"NetworkPolicyPluginError"`
	// 容器网络插件

	ClusterNetworkSettings *string `json:"ClusterNetworkSettings,omitempty" name:"ClusterNetworkSettings"`
}

type CreateVulRegistryImageExportJobRequest struct {
	*tchttp.BaseRequest

	// 漏洞ID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// OnlyAffectedNewestImage bool 是否影响最新镜像
	// ImageDigest 镜像Digest，支持模糊查询
	// ImageId 镜像ID，支持模糊查询
	// Namespace 命名空间，支持模糊查询
	// ImageTag 镜像版本，支持模糊查询
	// InstanceName 实例名称，支持模糊查询
	// ImageName 镜像名，支持模糊查询
	// ImageRepoAddress 镜像地址，支持模糊查询

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateVulRegistryImageExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulRegistryImageExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterCheckTaskRequest struct {
	*tchttp.BaseRequest

	// 指定要扫描的集群信息

	ClusterCheckTaskList []*ClusterCheckTaskItem `json:"ClusterCheckTaskList,omitempty" name:"ClusterCheckTaskList"`
}

func (r *CreateClusterCheckTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterCheckTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIgnoreVulRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID 信息列表

	List []*ModifyIgnoreVul `json:"List,omitempty" name:"List"`
}

func (r *DeleteIgnoreVulRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIgnoreVulRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecLogDeliveryKafkaSettingInfo struct {

	// 日志类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 主题ID

	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 投递状态(false:关 true:开)

	State *bool `json:"State,omitempty" name:"State"`
}

type DeleteEscapeWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEscapeWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEscapeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageVirus struct {

	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 风险等级

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 文件类型

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// 文件路径

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件md5

	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`
	// 大小

	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`
	// 首次发现时间

	FirstScanTime *string `json:"FirstScanTime,omitempty" name:"FirstScanTime"`
	// 最近扫描时间

	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
}

type DescribeNetworkTopologyFlowDetailRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 流量id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeNetworkTopologyFlowDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyFlowDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulScanRegistryImageInfo struct {

	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 扫描状态。NOT_SCAN:未扫描，SCANNING:扫描中,SCANNED:完成，CANCELED：扫描停止;
	// SCAN_ERR：检查失败

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 扫描时长

	ScanDuration *int64 `json:"ScanDuration,omitempty" name:"ScanDuration"`
	// 高危漏洞数

	HighLevelVulCount *int64 `json:"HighLevelVulCount,omitempty" name:"HighLevelVulCount"`
	// 中危漏洞数

	MediumLevelVulCount *int64 `json:"MediumLevelVulCount,omitempty" name:"MediumLevelVulCount"`
	// 低危漏洞数

	LowLevelVulCount *int64 `json:"LowLevelVulCount,omitempty" name:"LowLevelVulCount"`
	// 严重漏洞数

	CriticalLevelVulCount *int64 `json:"CriticalLevelVulCount,omitempty" name:"CriticalLevelVulCount"`
	// 镜像版本

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 镜像地址

	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`
	// 漏洞扫描的开始时间

	ScanStartTime *string `json:"ScanStartTime,omitempty" name:"ScanStartTime"`
	// 漏洞扫描的结束时间

	ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`
	// 扫描任务消息
	// "MessageRegistryDownload":     "仓库镜像层下载错误",
	// "MessageRegistryNet":          "仓库镜像层信息获取错误，可能原因是仓库未开启外网访问",
	// "MessageRegistryTimeout":      "仓库镜像任务超时",
	// "MessageRegistryScan":         "仓库镜像扫描错误",
	// "MessageRegistryInner":        "仓库内部拉取错误,请检查仓库配置",
	// "MessageRegistryInnerTimeout": "仓库内部拉取超时,可能原因是仓库未开启外网访问"

	Message *string `json:"Message,omitempty" name:"Message"`
	// 是否是最新镜像

	IsLatestImage *bool `json:"IsLatestImage,omitempty" name:"IsLatestImage"`
	// 仓库内部镜像ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type ComplianceAssetPolicyItem struct {

	// 为客户分配的唯一的检测项的ID。

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 检测项的原始ID

	BasePolicyItemId *uint64 `json:"BasePolicyItemId,omitempty" name:"BasePolicyItemId"`
	// 检测项的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 检测项所属的类型的名称

	Category *string `json:"Category,omitempty" name:"Category"`
	// 所属的合规标准的ID

	BenchmarkStandardId *uint64 `json:"BenchmarkStandardId,omitempty" name:"BenchmarkStandardId"`
	// 所属的合规标准的名称

	BenchmarkStandardName *string `json:"BenchmarkStandardName,omitempty" name:"BenchmarkStandardName"`
	// 威胁等级

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 检测状态
	// CHECK_INIT, 待检测
	// CHECK_RUNNING, 检测中
	// CHECK_FINISHED, 检测完成
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 检测结果
	// RESULT_PASSED: 通过
	// RESULT_FAILED: 未通过

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 检测项对应的白名单项的ID。如果存在且非0，表示检测项被用户忽略。

	WhitelistId *uint64 `json:"WhitelistId,omitempty" name:"WhitelistId"`
	// 处理建议。

	FixSuggestion *string `json:"FixSuggestion,omitempty" name:"FixSuggestion"`
	// 最近检测的时间。

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 验证信息

	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
}

type DescribeK8sApiAbnormalRuleScopeListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 过滤条件。
	// <li>Action - string -是否必填: 否 - 执行动作</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeK8sApiAbnormalRuleScopeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleScopeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceHostDetailInfo struct {

	// 主机上的Docker版本。

	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`
	// 主机上的K8S的版本。

	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`
	// 主机上Containerd版本

	ContainerdVersion *string `json:"ContainerdVersion,omitempty" name:"ContainerdVersion"`
}

type DescribeImageDenyEventDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件ID

		EventID *int64 `json:"EventID,omitempty" name:"EventID"`
		// 事件类型 EVENT_RISK:风险事件类型，EVENT_PRIVILEGE:特权

		EventType *string `json:"EventType,omitempty" name:"EventType"`
		// 规则名称

		RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
		// 规则RuleID

		RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
		// 规则类型

		RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
		// 规则启用状态 0:开启，1:关闭

		RuleStatus *int64 `json:"RuleStatus,omitempty" name:"RuleStatus"`
		// 规则策略状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中

		RuleEffectStatus *string `json:"RuleEffectStatus,omitempty" name:"RuleEffectStatus"`
		// 规则内容

		RuleInfo []*string `json:"RuleInfo,omitempty" name:"RuleInfo"`
		// 规则描述

		RuleDescription *string `json:"RuleDescription,omitempty" name:"RuleDescription"`
		// 镜像ID

		ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
		// 镜像名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 节点名称

		NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
		// 内网IP

		NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`
		// 外网IP

		PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
		// 主机Quuid

		QUUID *string `json:"QUUID,omitempty" name:"QUUID"`
		// 首次生成时间

		FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
		// 最近生成时间

		LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
		// 事件数量

		EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
		// 执行动作:
		// BEHAVIOR_ALERT:告警，
		// BEHAVIOR_HOLDUP_SUCCESSED:拦截

		DealBehavior *string `json:"DealBehavior,omitempty" name:"DealBehavior"`
		// Pod名称

		PodName *string `json:"PodName,omitempty" name:"PodName"`
		// 规则开始拦截时间

		RuleEffectTime *string `json:"RuleEffectTime,omitempty" name:"RuleEffectTime"`
		// 事件描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 镜像启动参数

		StartParam *string `json:"StartParam,omitempty" name:"StartParam"`
		// 解决方案

		Solution *string `json:"Solution,omitempty" name:"Solution"`
		// pod ip

		PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
		//  pod状态

		PodStatus *string `json:"PodStatus,omitempty" name:"PodStatus"`
		// 集群id

		ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
		// 节点类型

		NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
		// 节点id

		NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
		// 节点唯一id

		NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
		// 节点子网id

		NodeSubNetID *string `json:"NodeSubNetID,omitempty" name:"NodeSubNetID"`
		// 节点子网名称

		NodeSubNetName *string `json:"NodeSubNetName,omitempty" name:"NodeSubNetName"`
		// 节点子网cidr

		NodeSubNetCIDR *string `json:"NodeSubNetCIDR,omitempty" name:"NodeSubNetCIDR"`
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyEventDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReverseShellRegexpWhiteListExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleName - string  - 是否必填: 否 -规则名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateReverseShellRegexpWhiteListExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateReverseShellRegexpWhiteListExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWebVulExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWebVulExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWebVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableNodesCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableNodesCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableNodesCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanComplianceAssetsByPolicyItemRequest struct {
	*tchttp.BaseRequest

	// 指定的检测项的ID

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 要重新扫描的客户资产项ID的列表。

	CustomerAssetIdSet []*uint64 `json:"CustomerAssetIdSet,omitempty" name:"CustomerAssetIdSet"`
}

func (r *ScanComplianceAssetsByPolicyItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanComplianceAssetsByPolicyItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTcssSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像总数

		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
		// 已扫描镜像数

		ScannedImageCnt *uint64 `json:"ScannedImageCnt,omitempty" name:"ScannedImageCnt"`
		// 待扫描镜像个数

		UnScannedImageCnt *uint64 `json:"UnScannedImageCnt,omitempty" name:"UnScannedImageCnt"`
		// 本地镜像个数

		LocalImageCnt *uint64 `json:"LocalImageCnt,omitempty" name:"LocalImageCnt"`
		// 仓库镜像个数

		RepositoryImageCnt *uint64 `json:"RepositoryImageCnt,omitempty" name:"RepositoryImageCnt"`
		// 风险本地镜像个数

		RiskLocalImageCnt *uint64 `json:"RiskLocalImageCnt,omitempty" name:"RiskLocalImageCnt"`
		// 风险仓库镜像个数

		RiskRepositoryImageCnt *uint64 `json:"RiskRepositoryImageCnt,omitempty" name:"RiskRepositoryImageCnt"`
		// 容器个数

		ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
		// 风险容器个数

		RiskContainerCnt *uint64 `json:"RiskContainerCnt,omitempty" name:"RiskContainerCnt"`
		// 集群个数

		ClusterCnt *uint64 `json:"ClusterCnt,omitempty" name:"ClusterCnt"`
		// 风险集群个数

		RiskClusterCnt *uint64 `json:"RiskClusterCnt,omitempty" name:"RiskClusterCnt"`
		// 待扫描漏洞个数

		UnScannedVulCnt *uint64 `json:"UnScannedVulCnt,omitempty" name:"UnScannedVulCnt"`
		// 风险漏洞个数

		RiskVulCnt *uint64 `json:"RiskVulCnt,omitempty" name:"RiskVulCnt"`
		// 安全基线待扫描项个数

		UnScannedBaseLineCnt *uint64 `json:"UnScannedBaseLineCnt,omitempty" name:"UnScannedBaseLineCnt"`
		// 安全基线风险个数

		RiskBaseLineCnt *uint64 `json:"RiskBaseLineCnt,omitempty" name:"RiskBaseLineCnt"`
		// 运行时(高危)待处理事件个数

		RuntimeUnhandleEventCnt *uint64 `json:"RuntimeUnhandleEventCnt,omitempty" name:"RuntimeUnhandleEventCnt"`
		// 待扫描集群个数

		UnScannedClusterCnt *uint64 `json:"UnScannedClusterCnt,omitempty" name:"UnScannedClusterCnt"`
		// 是否已扫描镜像

		ScanImageStatus *bool `json:"ScanImageStatus,omitempty" name:"ScanImageStatus"`
		// 是否已扫描集群

		ScanClusterStatus *bool `json:"ScanClusterStatus,omitempty" name:"ScanClusterStatus"`
		// 是否已扫描基线

		ScanBaseLineStatus *bool `json:"ScanBaseLineStatus,omitempty" name:"ScanBaseLineStatus"`
		// 是否已扫描漏洞

		ScanVulStatus *bool `json:"ScanVulStatus,omitempty" name:"ScanVulStatus"`
		// 漏洞影响镜像数

		VulRiskImageCnt *uint64 `json:"VulRiskImageCnt,omitempty" name:"VulRiskImageCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTcssSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTcssSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表

		List []*ServiceInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecEventsTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运行时安全事件趋势信息列表

		EventTendencySet []*SecTendencyEventInfo `json:"EventTendencySet,omitempty" name:"EventTendencySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecEventsTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecEventsTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallClusterContainerSecurityRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterIDs []*string `json:"ClusterIDs,omitempty" name:"ClusterIDs"`
}

func (r *UninstallClusterContainerSecurityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallClusterContainerSecurityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnauthorizedCoresTendencyRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUnauthorizedCoresTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnauthorizedCoresTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellWhiteListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReverseShellWhiteListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNameListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群名称列表

		ClusterNames []*string `json:"ClusterNames,omitempty" name:"ClusterNames"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterNameListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogDeliveryKafkaSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消息队列实例ID

		InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
		// 消息队列实例名称

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 域名

		Domain *string `json:"Domain,omitempty" name:"Domain"`
		// 日志类型队列

		LogTypeList []*SecLogDeliveryKafkaSettingInfo `json:"LogTypeList,omitempty" name:"LogTypeList"`
		// 用户名

		User *string `json:"User,omitempty" name:"User"`
		// 地域ID

		RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogDeliveryKafkaSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogDeliveryKafkaSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryAutoAuthorizedRuleRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageRegistryAutoAuthorizedRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryAutoAuthorizedRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePolicyItemIgnoredAssetListRequest struct {
	*tchttp.BaseRequest

	// 检测项的ID

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// Name - String
	// Name 可取值：NodeName, CheckResult

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCompliancePolicyItemIgnoredAssetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePolicyItemIgnoredAssetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetImageRegistryAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncAssetImageRegistryAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetImageRegistryAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearLocalStorageRequest struct {
	*tchttp.BaseRequest
}

func (r *ClearLocalStorageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearLocalStorageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetComponentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件列表

		List []*ComponentInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetComponentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetComponentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 异常进程数组

		EventSet []*AbnormalProcessEventInfo `json:"EventSet,omitempty" name:"EventSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IACItem struct {

	// 规则名

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 严重程度

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 种类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 是否通过;Pass:通过;NoPass:不通过

	Passed *string `json:"Passed,omitempty" name:"Passed"`
}

type CreateAssetImageRiskExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetImageRiskExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageRiskExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAndPublishNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAndPublishNetworkFirewallPolicyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAndPublishNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompliancePolicyAssetSetItem struct {

	// 检查项ID

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 需要忽略指定检查项内的资产ID列表，为空表示所有

	CustomerAssetItemIdSet []*uint64 `json:"CustomerAssetItemIdSet,omitempty" name:"CustomerAssetItemIdSet"`
}

type DeleteImageDenyRuleRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。Filters和RuleIDSet二者选其一
	// <li>RuleType- String - 是否必填：否 -规则类型 RULE_RISK：风险， RULE_PRIVILEGE：特权</li>
	// <li>EffectStatus- String - 是否必填：否 - 生效状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中。</li>
	// <li>RuleName- string - 是否必填：否 - 规则名称。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 指定删除的规则RuleID数组。Filters和RuleIDSet二者选其一

	RuleIDSet []*string `json:"RuleIDSet,omitempty" name:"RuleIDSet"`
	// 过滤条件中不删除的规则RuleID数组

	ExcludeRuleIDSet []*string `json:"ExcludeRuleIDSet,omitempty" name:"ExcludeRuleIDSet"`
}

func (r *DeleteImageDenyRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageDenyRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefencePluginRequest struct {
	*tchttp.BaseRequest

	// 主机HostID或超级节点UniqueId

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>
	// Status- String - 是否必填：否 -插件运行状态：注入中:INJECTING，注入成功：SUCCESS，注入失败：FAIL，插件超时：TIMEOUT，插件退出：QUIT
	// </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulDefencePluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefencePluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetImageRegistryScanStopRequest struct {
	*tchttp.BaseRequest

	// 是否扫描全部镜像

	All *bool `json:"All,omitempty" name:"All"`
	// 扫描的镜像列表

	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`
	// 扫描的镜像列表

	Id []*uint64 `json:"Id,omitempty" name:"Id"`
	// 过滤条件

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 不要扫描的镜像列表，与Filters配合使用

	ExcludeImageList []*uint64 `json:"ExcludeImageList,omitempty" name:"ExcludeImageList"`
	// 是否仅扫描各repository最新版本的镜像

	OnlyScanLatest *bool `json:"OnlyScanLatest,omitempty" name:"OnlyScanLatest"`
	// 停止的任务ID

	TaskID *uint64 `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *ModifyAssetImageRegistryScanStopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetImageRegistryScanStopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulAffectedContainerInfo struct {

	// 内网IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 容器ID

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// Pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// PodIP值

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 节点类型[NORMAL:普通节点|SUPER:超级节点]

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 超级节点唯一ID

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 超级节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 超级节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type DescribeK8sApiAbnormalSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待处理事件个数

		UnhandleEventCount *uint64 `json:"UnhandleEventCount,omitempty" name:"UnhandleEventCount"`
		// 待处理高危事件个数

		UnhandleHighLevelEventCount *uint64 `json:"UnhandleHighLevelEventCount,omitempty" name:"UnhandleHighLevelEventCount"`
		// 待处理中危事件个数

		UnhandleMediumLevelEventCount *uint64 `json:"UnhandleMediumLevelEventCount,omitempty" name:"UnhandleMediumLevelEventCount"`
		// 待处理低危事件个数

		UnhandleLowLevelEventCount *uint64 `json:"UnhandleLowLevelEventCount,omitempty" name:"UnhandleLowLevelEventCount"`
		// 待处理提示级别事件个数

		UnhandleNoticeLevelEventCount *uint64 `json:"UnhandleNoticeLevelEventCount,omitempty" name:"UnhandleNoticeLevelEventCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSuperNodeInstallTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务列表

		List []*SuperNodeInstallTaskListItem `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSuperNodeInstallTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSuperNodeInstallTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsEventDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件ID

		EventID *uint64 `json:"EventID,omitempty" name:"EventID"`
		// 事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP

		EventType *string `json:"EventType,omitempty" name:"EventType"`
		// 恶意请求次数

		EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
		// 首次发现时间

		FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
		// 最近生成时间

		LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
		// 容器ID

		ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
		// 容器名称

		ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
		// 隔离状态
		// 未隔离  	NORMAL
		// 已隔离		ISOLATED
		// 隔离中		ISOLATING
		// 隔离失败	ISOLATE_FAILED
		// 解除隔离中  RESTORING
		// 解除隔离失败 RESTORE_FAILED

		ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
		// 容器状态
		// 正在运行: RUNNING
		// 暂停: PAUSED
		// 停止: STOPPED
		// 已经创建: CREATED
		// 已经销毁: DESTROYED
		// 正在重启中: RESTARTING
		// 迁移中: REMOVING

		ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
		// 容器子状态
		// "AGENT_OFFLINE"       //Agent离线
		// "NODE_DESTROYED"      //节点已销毁
		// "CONTAINER_EXITED"    //容器已退出
		// "CONTAINER_DESTROYED" //容器已销毁
		// "SHARED_HOST"         // 容器与主机共享网络
		// "RESOURCE_LIMIT"      //隔离操作资源超限
		// "UNKNOW"              // 原因未知

		ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
		// 容器隔离操作来源

		ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
		// 镜像ID

		ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
		// 镜像名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 主机名称

		HostName *string `json:"HostName,omitempty" name:"HostName"`
		// 内网IP

		HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
		// 外网IP

		PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
		// 节点名称

		PodName *string `json:"PodName,omitempty" name:"PodName"`
		// 事件描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 解决方案

		Solution *string `json:"Solution,omitempty" name:"Solution"`
		// 参考链接

		Reference []*string `json:"Reference,omitempty" name:"Reference"`
		// 恶意域名或IP

		Address *string `json:"Address,omitempty" name:"Address"`
		// 恶意IP所属城市

		City *string `json:"City,omitempty" name:"City"`
		// 命中规则类型
		// SYSTEM：系统规则
		//  USER：用户自定义

		MatchRuleType *string `json:"MatchRuleType,omitempty" name:"MatchRuleType"`
		// 标签特征

		FeatureLabel *string `json:"FeatureLabel,omitempty" name:"FeatureLabel"`
		// 进程权限

		ProcessAuthority *string `json:"ProcessAuthority,omitempty" name:"ProcessAuthority"`
		// 进程md5

		ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
		// 进程启动用户

		ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`
		// 进程用户组

		ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`
		// 进程路径

		ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
		// 进程树

		ProcessTree *string `json:"ProcessTree,omitempty" name:"ProcessTree"`
		// 进程命令行参数

		ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
		// 父进程启动用户

		ParentProcessStartUser *string `json:"ParentProcessStartUser,omitempty" name:"ParentProcessStartUser"`
		// 父进程用户组

		ParentProcessUserGroup *string `json:"ParentProcessUserGroup,omitempty" name:"ParentProcessUserGroup"`
		// 父进程路径

		ParentProcessPath *string `json:"ParentProcessPath,omitempty" name:"ParentProcessPath"`
		// 父进程命令行参数

		ParentProcessParam *string `json:"ParentProcessParam,omitempty" name:"ParentProcessParam"`
		// 祖先进程启动用户

		AncestorProcessStartUser *string `json:"AncestorProcessStartUser,omitempty" name:"AncestorProcessStartUser"`
		// 祖先进程用户组

		AncestorProcessUserGroup *string `json:"AncestorProcessUserGroup,omitempty" name:"AncestorProcessUserGroup"`
		// 祖先进程路径

		AncestorProcessPath *string `json:"AncestorProcessPath,omitempty" name:"AncestorProcessPath"`
		// 祖先进程命令行参数

		AncestorProcessParam *string `json:"AncestorProcessParam,omitempty" name:"AncestorProcessParam"`
		// 主机ID

		HostID *string `json:"HostID,omitempty" name:"HostID"`
		// 事件状态
		// EVENT_UNDEAL： 待处理
		// EVENT_DEALED：已处理
		// EVENT_IGNORE： 已忽略
		// EVENT_ADD_WHITE：已加白

		EventStatus *string `json:"EventStatus,omitempty" name:"EventStatus"`
		// 操作时间

		OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 节点类型

		NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
		// 节点名称

		NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
		// 节点子网ID

		NodeSubNetID *string `json:"NodeSubNetID,omitempty" name:"NodeSubNetID"`
		// 节点子网名称

		NodeSubNetName *string `json:"NodeSubNetName,omitempty" name:"NodeSubNetName"`
		// 节点子网网段

		NodeSubNetCIDR *string `json:"NodeSubNetCIDR,omitempty" name:"NodeSubNetCIDR"`
		// 集群ID

		ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
		// podip

		PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
		// pod状态

		PodStatus *string `json:"PodStatus,omitempty" name:"PodStatus"`
		// 节点唯一id

		NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
		// 节点ID名称

		NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// Namespace

		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
		// 工作负载类型

		WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsEventDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVirusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像病毒列表

		List []*ImageVirusInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 病毒扫描状态
		// 0:未扫描
		// 1:扫描中
		// 2:扫描完成
		// 3:扫描出错
		// 4:扫描取消

		VirusScanStatus *uint64 `json:"VirusScanStatus,omitempty" name:"VirusScanStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageVirusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeWhiteListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>EventType- String - 是否必填：否 - 加白事件类型，ESCAPE_CGROUPS：利用cgroup机制逃逸，ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸， ESCAPE_DOCKER_API：访问Docker API接口逃逸，ESCAPE_VUL_OCCURRED：逃逸漏洞利用，MOUNT_SENSITIVE_PTAH：敏感路径挂载，PRIVILEGE_CONTAINER_START：特权容器， PRIVILEGE：程序提权逃逸</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：主机数量：HostCount，容器数量：ContainerCount，更新时间：UpdateTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEscapeWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoAuthImageInfo struct {

	// 仓库类型

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 仓库实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 项目空间（命名空间）

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 唯一Key

	UniqueKey *string `json:"UniqueKey,omitempty" name:"UniqueKey"`
	// 镜像仓数量

	RepoCount *uint64 `json:"RepoCount,omitempty" name:"RepoCount"`
}

type DescribeAssetImageRegistryVirusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像漏洞列表

		List []*ImageVirus `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryVirusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEscapeRegexpWhiteListExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEscapeRegexpWhiteListExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEscapeRegexpWhiteListExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryRegistryDetailRequest struct {
	*tchttp.BaseRequest

	// 仓库唯一id

	RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeAssetImageRegistryRegistryDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRegistryDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopVulScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopVulScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopVulScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSuperNodeListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>NodeID- String - 是否必填：否 - ID </li>
	// <li>NodeName- String - 是否必填：否 - 超级节点名称 </li>
	// <li>SubnetName- String - 是否必填：否 - VPC子网 </li>
	// <li>AgentStatus- String - 是否必填：否 - 安装状态UNINSTALL:未安装;INSTALLED:已安装;INSTALLING:安装中; </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetSuperNodeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSuperNodeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryAssetStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetImageRegistryAssetStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryAssetStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessDetailInfo struct {

	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程权限

	ProcessAuthority *string `json:"ProcessAuthority,omitempty" name:"ProcessAuthority"`
	// 进程pid

	ProcessId *uint64 `json:"ProcessId,omitempty" name:"ProcessId"`
	// 进程启动用户

	ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`
	// 进程用户组

	ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 进程树

	ProcessTree *string `json:"ProcessTree,omitempty" name:"ProcessTree"`
	// 进程md5

	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
	// 进程命令行参数

	ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
}

type DescribeClusterServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群Service列表信息

		ClusterServiceList []*ClusterServiceInfo `json:"ClusterServiceList,omitempty" name:"ClusterServiceList"`
		// 集群Service列表总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSampleDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件Md5值

		MD5 *string `json:"MD5,omitempty" name:"MD5"`
		// 文件大小(B)

		Size *uint64 `json:"Size,omitempty" name:"Size"`
		// 病毒名

		VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
		// 风险等级 RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。

		RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
		// 查杀引擎

		KillEngine []*string `json:"KillEngine,omitempty" name:"KillEngine"`
		// 标签

		Tags []*string `json:"Tags,omitempty" name:"Tags"`
		// 事件描述

		HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
		// 建议方案

		SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
		// 参考链接

		ReferenceLink *string `json:"ReferenceLink,omitempty" name:"ReferenceLink"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSampleDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPortListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>All - String - 是否必填：否 - 模糊查询可选字段</li>
	// <li>RunAs - String - 是否必填：否 - 运行用户筛选，</li>
	// <li>ContainerID - String - 是否必填：否 - 容器id</li>
	// <li>HostID- String - 是否必填：是 - 主机id</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>ProcessName- string - 是否必填：否 - 进程名搜索</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetPortListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRiskListExportJobRequest struct {
	*tchttp.BaseRequest

	// 风险项Id列表，为空则导出所有风险项

	CheckItemIdList []*uint64 `json:"CheckItemIdList,omitempty" name:"CheckItemIdList"`
}

func (r *CreateRiskListExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskListExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyK8sApiAbnormalEventStatusRequest struct {
	*tchttp.BaseRequest

	// 事件ID集合

	EventIDSet []*uint64 `json:"EventIDSet,omitempty" name:"EventIDSet"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyK8sApiAbnormalEventStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyK8sApiAbnormalEventStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAndPublishNetworkFirewallPolicyYamlDetailRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// base64编码的networkpolicy yaml字符串

	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *AddAndPublishNetworkFirewallPolicyYamlDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAndPublishNetworkFirewallPolicyYamlDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusScanSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVirusScanSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterCreateComponentItem struct {

	// 要安装组件的集群ID。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 该集群对应的地域

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

type ReverseShellWhiteListBaseInfo struct {

	// 白名单id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 镜像数量

	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`
	// 连接进程名字

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 目标地址ip

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 目标端口

	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`
	// 是否是全局白名单，true全局

	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 镜像id数组，为空代表全部

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

type DescribeSecEventsTendencyRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeSecEventsTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecEventsTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulTopRankingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞Top排名信息列表

		List []*VulTopRankingInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulTopRankingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulTopRankingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogCleanSettingInfoRequest struct {
	*tchttp.BaseRequest

	// 触发清理的储量底线(50-99)

	ReservesLimit *uint64 `json:"ReservesLimit,omitempty" name:"ReservesLimit"`
	// 清理停止时的储量截至线(>0,小于ReservesLimit)

	ReservesDeadline *uint64 `json:"ReservesDeadline,omitempty" name:"ReservesDeadline"`
	// 触发清理的存储天数(>=1)

	DayLimit *uint64 `json:"DayLimit,omitempty" name:"DayLimit"`
}

func (r *ModifySecLogCleanSettingInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogCleanSettingInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageSimpleListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 镜像名、镜像id 称筛选，</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageSimpleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageSimpleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEscapeEventsExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,Status：EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：latest_found_time

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateEscapeEventsExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEscapeEventsExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkTopologyFlowSwitchTaskRequest struct {
	*tchttp.BaseRequest

	// 集群Id数组

	ClusterId []*string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 流量开关 ON 开启 OFF关闭
	//

	FlowSwitch *string `json:"FlowSwitch,omitempty" name:"FlowSwitch"`
}

func (r *CreateNetworkTopologyFlowSwitchTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkTopologyFlowSwitchTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallClusterListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkFirewallClusterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallClusterListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAbnormalProcessRulesExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAbnormalProcessRulesExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAbnormalProcessRulesExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyK8sApiAbnormalRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 规则详情

	RuleInfo *K8sApiAbnormalRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`
}

func (r *ModifyK8sApiAbnormalRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyK8sApiAbnormalRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulDefenceHostExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulDefenceHostExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulDefenceHostExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlEventInfo struct {

	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 命中规则名称

	MatchRuleName *string `json:"MatchRuleName,omitempty" name:"MatchRuleName"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 动作执行结果，   BEHAVIOR_NONE: 无
	//     BEHAVIOR_ALERT: 告警
	//     BEHAVIOR_RELEASE：放行
	//     BEHAVIOR_HOLDUP_FAILED:拦截失败
	//     BEHAVIOR_HOLDUP_SUCCESSED：拦截失败

	Behavior *string `json:"Behavior,omitempty" name:"Behavior"`
	// 状态0:未处理  “EVENT_UNDEAL”:事件未处理
	//     "EVENT_DEALED":事件已经处理
	//     "EVENT_INGNORE"：事件已经忽略

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件记录的唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 事件类型， FILE_ABNORMAL_READ:文件异常读取

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 镜像id, 用于跳转

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器id, 用于跳转

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 事件解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件详细描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 命中策略id

	MatchRuleId *string `json:"MatchRuleId,omitempty" name:"MatchRuleId"`
	// 命中规则行为：
	// RULE_MODE_RELEASE 放行
	// RULE_MODE_ALERT  告警
	// RULE_MODE_HOLDUP 拦截

	MatchAction *string `json:"MatchAction,omitempty" name:"MatchAction"`
	// 命中规则进程信息

	MatchProcessPath *string `json:"MatchProcessPath,omitempty" name:"MatchProcessPath"`
	// 命中规则文件信息

	MatchFilePath *string `json:"MatchFilePath,omitempty" name:"MatchFilePath"`
	// 文件路径，包含名字

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 规则是否存在

	RuleExist *bool `json:"RuleExist,omitempty" name:"RuleExist"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 规则组id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 节点名称：如果是超级节点，展示的实质上是它的node_id

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// pod ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 节点类型：NORMAL普通节点、SUPER超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 集群id

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 节点的唯一id，主要是超级节点使用

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 节点公网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// uuid

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 节点内网ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type DescribeSecLogCleanSettingInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecLogCleanSettingInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogCleanSettingInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageRegistryScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的任务ID

		TaskID *uint64 `json:"TaskID,omitempty" name:"TaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetImageRegistryScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageRegistryScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogAlertMsgRequest struct {
	*tchttp.BaseRequest

	// 告警类型
	// 日志储量告警: log_reserve_full
	// 日志存储时间告警: log_save_day_limit
	// kafka实例/公网域名不可用: kafka_instance_domain_unavailable
	// kafka 用户名密码错误: kafka_user_passwd_wrong
	// kafka后台报错字段: kafka_field_wrong

	Type []*string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSecLogAlertMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogAlertMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		Info *K8sApiAbnormalRuleInfo `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrialStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTrialStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrialStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEscapeWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEscapeWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEscapeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最近的一次扫描任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 木马影响容器个数

		RiskContainerCnt *uint64 `json:"RiskContainerCnt,omitempty" name:"RiskContainerCnt"`
		// 待处理风险个数

		RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
		// 病毒库更新时间

		VirusDataBaseModifyTime *string `json:"VirusDataBaseModifyTime,omitempty" name:"VirusDataBaseModifyTime"`
		// 木马影响容器个数较昨日增长

		RiskContainerIncrease *int64 `json:"RiskContainerIncrease,omitempty" name:"RiskContainerIncrease"`
		// 待处理风险个数较昨日增长

		RiskIncrease *int64 `json:"RiskIncrease,omitempty" name:"RiskIncrease"`
		// 隔离事件个数较昨日新增

		IsolateIncrease *int64 `json:"IsolateIncrease,omitempty" name:"IsolateIncrease"`
		// 隔离事件总数

		IsolateCnt *int64 `json:"IsolateCnt,omitempty" name:"IsolateCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportDefenceVulResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持防御的漏洞列表

		List []*SupportDefenceVul `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSupportDefenceVulResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSupportDefenceVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogKafkaUINResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecLogKafkaUINResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogKafkaUINResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessBaseInfo struct {

	// 进程启动用户

	ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`
	// 进程用户组

	ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 进程命令行参数

	ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
}

type ModifyVulDefenceEventStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulDefenceEventStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceEventStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageDenyEventExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>EventType- String - 是否必填：否 -事件类型 EVENT_RISK:风险事件类型，EVENT_PRIVILEGE:特权。</li>
	// <li>DealBehavior- String - 是否必填：否 - 执行动作,BEHAVIOR_ALERT:告警，BEHAVIOR_HOLDUP_SUCCESSED:拦截。</li>
	// <li>RuleName- string - 是否必填：否 - 规则名称。</li>
	// <li>NodeName- string - 是否必填：否 - 节点名称。</li>
	// <li>NodeIP- string - 是否必填：否 - 内外IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：告警数量：EventCount，最近生成时间：LatestFoundTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateImageDenyEventExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageDenyEventExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogJoinSuperObjectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 接入对象列表

		List []*SecLogJoinSuperObjectInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogJoinSuperObjectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogJoinSuperObjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterNodeFromWhiteListRequest struct {
	*tchttp.BaseRequest

	// 要删除的节点白名单Id列表

	NodeIdList []*uint64 `json:"NodeIdList,omitempty" name:"NodeIdList"`
}

func (r *DeleteClusterNodeFromWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterNodeFromWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDetailInfo struct {

	// CVE编号

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 漏洞标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 漏洞类型

	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
	// 漏洞威胁等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 漏洞披露时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// 漏洞描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// CVSS V3描述

	CVSSV3Desc *string `json:"CVSSV3Desc,omitempty" name:"CVSSV3Desc"`
	// 漏洞修复建议

	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`
	// 缓解措施

	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`
	// 参考链接

	Reference []*string `json:"Reference,omitempty" name:"Reference"`
	// CVSS V3分数

	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`
	// 受漏洞影响的组件列表

	ComponentList []*VulAffectedComponentInfo `json:"ComponentList,omitempty" name:"ComponentList"`
	// 影响本地镜像数

	LocalImageCount *int64 `json:"LocalImageCount,omitempty" name:"LocalImageCount"`
	// 影响容器数

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 影响仓库镜像数

	RegistryImageCount *int64 `json:"RegistryImageCount,omitempty" name:"RegistryImageCount"`
	// 漏洞子类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// 影响最新本地镜像数

	LocalNewestImageCount *int64 `json:"LocalNewestImageCount,omitempty" name:"LocalNewestImageCount"`
	// 影响最新仓库镜像数

	RegistryNewestImageCount *int64 `json:"RegistryNewestImageCount,omitempty" name:"RegistryNewestImageCount"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 防御状态，NO_DEFENDED:未防御，DEFENDED:已防御

	DefenceStatus *string `json:"DefenceStatus,omitempty" name:"DefenceStatus"`
	// 漏洞防御主机范围: MANUAL:自选主机节点，ALL:全部

	DefenceScope *string `json:"DefenceScope,omitempty" name:"DefenceScope"`
	// 漏洞防御主机数量

	DefenceHostCount *int64 `json:"DefenceHostCount,omitempty" name:"DefenceHostCount"`
	// 已防御攻击次数

	DefendedCount *int64 `json:"DefendedCount,omitempty" name:"DefendedCount"`
	// 是否已扫描，NOT_SCAN:未扫描,SCANNED:已扫描

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
}

type DescribePublicProxyInstallCommandRequest struct {
	*tchttp.BaseRequest

	// ip

	IP *string `json:"IP,omitempty" name:"IP"`
}

func (r *DescribePublicProxyInstallCommandRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicProxyInstallCommandRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仓库名

		Name *string `json:"Name,omitempty" name:"Name"`
		// 用户名

		Username *string `json:"Username,omitempty" name:"Username"`
		// 密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 仓库url

		Url *string `json:"Url,omitempty" name:"Url"`
		// 仓库类型，列表：harbor

		RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
		// 仓库版本

		RegistryVersion *string `json:"RegistryVersion,omitempty" name:"RegistryVersion"`
		// 网络类型，列表：public（公网）,private（私网）

		NetType *string `json:"NetType,omitempty" name:"NetType"`
		// 区域，列表:default（默认）

		RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`
		// 限速

		SpeedLimit *uint64 `json:"SpeedLimit,omitempty" name:"SpeedLimit"`
		// 安全模式（证书校验）：0（默认） 非安全模式（跳过证书校验）：1

		Insecure *uint64 `json:"Insecure,omitempty" name:"Insecure"`
		// 联通性检测结果详情

		ConnDetectDetail []*RegistryConnDetectResult `json:"ConnDetectDetail,omitempty" name:"ConnDetectDetail"`
		// tcr情况下instance_id

		InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryRegistryDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodeWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检测项名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 忽略的节点信息列表

		WhiteListNodes []*WhiteListNodeInfo `json:"WhiteListNodes,omitempty" name:"WhiteListNodes"`
		// 节点类型白名单总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterNodeWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportJobDownloadURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadURL *string `json:"DownloadURL,omitempty" name:"DownloadURL"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportJobDownloadURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportJobDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulLevelImageSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 高危漏洞最新本地镜像占比

		HighLevelVulLocalImagePercent *float64 `json:"HighLevelVulLocalImagePercent,omitempty" name:"HighLevelVulLocalImagePercent"`
		// 中危漏洞最新本地镜像占比

		MediumLevelVulLocalImagePercent *float64 `json:"MediumLevelVulLocalImagePercent,omitempty" name:"MediumLevelVulLocalImagePercent"`
		// 低危漏洞最新本地镜像占比

		LowLevelVulLocalImagePercent *float64 `json:"LowLevelVulLocalImagePercent,omitempty" name:"LowLevelVulLocalImagePercent"`
		// 严重漏洞最新本地镜像占比

		CriticalLevelVulLocalImagePercent *float64 `json:"CriticalLevelVulLocalImagePercent,omitempty" name:"CriticalLevelVulLocalImagePercent"`
		// 影响的最新版本本地镜像数

		LocalNewestImageCount *int64 `json:"LocalNewestImageCount,omitempty" name:"LocalNewestImageCount"`
		// 影响的最新版本仓库镜像数

		RegistryNewestImageCount *int64 `json:"RegistryNewestImageCount,omitempty" name:"RegistryNewestImageCount"`
		// 高危漏洞最新仓库镜像占比

		HighLevelVulRegistryImagePercent *float64 `json:"HighLevelVulRegistryImagePercent,omitempty" name:"HighLevelVulRegistryImagePercent"`
		// 中危漏洞最新仓库镜像占比

		MediumLevelVulRegistryImagePercent *float64 `json:"MediumLevelVulRegistryImagePercent,omitempty" name:"MediumLevelVulRegistryImagePercent"`
		// 低危漏洞最新仓库镜像占比

		LowLevelVulRegistryImagePercent *float64 `json:"LowLevelVulRegistryImagePercent,omitempty" name:"LowLevelVulRegistryImagePercent"`
		// 严重漏洞最新仓库镜像占比

		CriticalLevelVulRegistryImagePercent *float64 `json:"CriticalLevelVulRegistryImagePercent,omitempty" name:"CriticalLevelVulRegistryImagePercent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulLevelImageSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulLevelImageSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulAffectedRegistryImageInfo struct {

	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像版本

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 镜像命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 镜像地址

	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`
	// 组件列表

	ComponentList []*VulAffectedImageComponentInfo `json:"ComponentList,omitempty" name:"ComponentList"`
	// 是否为镜像的最新版本

	IsLatestImage *bool `json:"IsLatestImage,omitempty" name:"IsLatestImage"`
	// 内部镜像资产ID

	ImageAssetId *int64 `json:"ImageAssetId,omitempty" name:"ImageAssetId"`
}

type DescribeAssetImageVirusListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>RiskLevel - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序 asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAssetImageVirusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVirusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserWorkloadKindsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户工作负载类型列表

		WorkloadKinds []*string `json:"WorkloadKinds,omitempty" name:"WorkloadKinds"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserWorkloadKindsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserWorkloadKindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerDetailRequest struct {
	*tchttp.BaseRequest

	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
}

func (r *DescribeAssetContainerDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageRegistryScanTaskOneKeyRequest struct {
	*tchttp.BaseRequest

	// 是否扫描全部镜像(废弃)

	All *bool `json:"All,omitempty" name:"All"`
	// 扫描的镜像列表

	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`
	// 扫描类型数组

	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`
	// 扫描的镜像列表Id

	Id []*uint64 `json:"Id,omitempty" name:"Id"`
	// 是否最新镜像

	IsLatest *bool `json:"IsLatest,omitempty" name:"IsLatest"`
	// 扫描范围 0全部镜像，1自选镜像，2推荐扫描镜像

	ScanScope *uint64 `json:"ScanScope,omitempty" name:"ScanScope"`
	// 仓库类型

	RegistryType []*string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 命名空间

	Namespace []*string `json:"Namespace,omitempty" name:"Namespace"`
	// 是否存在运行中的容器

	ContainerRunning *bool `json:"ContainerRunning,omitempty" name:"ContainerRunning"`
	// 任务超时时长单位s

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
}

func (r *CreateAssetImageRegistryScanTaskOneKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageRegistryScanTaskOneKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterAssetExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务Id

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterAssetExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterAssetExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessEventTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待处理异常进程事件趋势

		List []*AbnormalProcessEventTendencyInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessEventTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterWorkloadWhiteListRequest struct {
	*tchttp.BaseRequest

	// 检测项ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterId集群Id,ClusterName集群名字,WorkloadName工作负载名称

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClusterWorkloadWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterWorkloadWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEscapeWhiteListRequest struct {
	*tchttp.BaseRequest

	// 加白名单事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸

	EventType []*string `json:"EventType,omitempty" name:"EventType"`
	// 加白名单镜像ID数组

	ImageIDs []*string `json:"ImageIDs,omitempty" name:"ImageIDs"`
}

func (r *AddEscapeWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEscapeWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportComplianceStatusListJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的导出任务的ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportComplianceStatusListJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportComplianceStatusListJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageDetailRequest struct {
	*tchttp.BaseRequest

	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *DescribeAssetImageDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoAuthorizedRuleHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像自动授权规则授权范围主机列表

		List []*AutoAuthorizedRuleHostInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoAuthorizedRuleHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoAuthorizedRuleHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegexpRuleListItem struct {

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 生效表达式个数

	EffectiveExpression *uint64 `json:"EffectiveExpression,omitempty" name:"EffectiveExpression"`
	// 最新编辑时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 最近编辑账号

	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`
	// 启用状态

	Status *bool `json:"Status,omitempty" name:"Status"`
}

type RegionInfo struct {

	// 地域标识

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type DescribeImageComponentListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式desc ，asc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeImageComponentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageComponentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessDetailRequest struct {
	*tchttp.BaseRequest

	// 事件唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeAbnormalProcessDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像仓库列表

		List []*ImageRepoInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像列表

		List []*ImagesInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPodListRequest struct {
	*tchttp.BaseRequest

	// 集群Id,不填表示获取用户所有pod

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name 可取值：ClusterId集群id,Namespace命名空间等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// Service名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeUserPodListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPodListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageDenyRule struct {

	// 规则RuleID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则类型 RULE_RISK：风险， RULE_PRIVILEGE：特权

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 生效的镜像数量

	EffectImageCount *int64 `json:"EffectImageCount,omitempty" name:"EffectImageCount"`
	// 是否对全部扫描镜像生效。0:全选镜像，1:自选镜像

	IsEffectAllImage *int64 `json:"IsEffectAllImage,omitempty" name:"IsEffectAllImage"`
	// 规则开始生效时间

	EffectTime *string `json:"EffectTime,omitempty" name:"EffectTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 操作用户

	OperationUin *string `json:"OperationUin,omitempty" name:"OperationUin"`
	// 启用状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 生效状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中

	EffectStatus *string `json:"EffectStatus,omitempty" name:"EffectStatus"`
	// 规则ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type SetComplianceHostRangeRequest struct {
	*tchttp.BaseRequest

	// 添加主机id

	AddHostIds []*uint64 `json:"AddHostIds,omitempty" name:"AddHostIds"`
	// 是否需要立即执行扫描

	PerformScan *bool `json:"PerformScan,omitempty" name:"PerformScan"`
	// 添加所有主机

	AddAll *bool `json:"AddAll,omitempty" name:"AddAll"`
}

func (r *SetComplianceHostRangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetComplianceHostRangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulRegistryImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 仓库镜像列表

		List []*VulAffectedRegistryImageInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulRegistryImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulRegistryImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建删除任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetworkFirewallPolicyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessEventsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAbnormalProcessEventsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusMonitorConfigRequest struct {
	*tchttp.BaseRequest

	// 是否开启实时监控

	EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`
	// true:包含路径 false:排除路径

	IsIncludePath *bool `json:"IsIncludePath,omitempty" name:"IsIncludePath"`
	// 自选排除或扫描的地址

	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
	// 扫描路径模式：
	// SCAN_PATH_ALL：全部路径
	// SCAN_PATH_DEFAULT：默认路径
	// SCAN_PATH_USER_DEFINE：用户自定义路径
	//

	ScanPathMode *string `json:"ScanPathMode,omitempty" name:"ScanPathMode"`
}

func (r *ModifyVirusMonitorConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusMonitorConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchExportListRequest struct {
	*tchttp.BaseRequest

	// ES查询条件JSON

	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeSearchExportListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchExportListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEditImageAutoAuthorizedRuleRequest struct {
	*tchttp.BaseRequest

	// 授权范围类别，MANUAL:自选主机节点，ALL:全部镜像

	RangeType *string `json:"RangeType,omitempty" name:"RangeType"`
	// 自选主机id，当授权范围为MANUAL:自选主机时且HostIdFilters为空时，必填

	HostIdSet []*string `json:"HostIdSet,omitempty" name:"HostIdSet"`
	// 每天最大的镜像授权数限制, 0表示无限制

	MaxDailyCount *int64 `json:"MaxDailyCount,omitempty" name:"MaxDailyCount"`
	// 规则id，在编辑时，必填

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 根据条件过滤，当授权范围为MANUAL:自选主机时且HostIdSet为空时，必填

	HostIdFilters []*AssetFilters `json:"HostIdFilters,omitempty" name:"HostIdFilters"`
	// 根据条件过滤而且排除指定主机id

	ExcludeHostIdSet []*string `json:"ExcludeHostIdSet,omitempty" name:"ExcludeHostIdSet"`
	// 规则是否生效，0:不生效，1:已生效

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 自动扫描开关

	AutoScanEnabled *int64 `json:"AutoScanEnabled,omitempty" name:"AutoScanEnabled"`
	// 自动扫描范围

	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`
}

func (r *AddEditImageAutoAuthorizedRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditImageAutoAuthorizedRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceAssetPolicySetItem struct {

	// 资产ID

	CustomerAssetItemId *uint64 `json:"CustomerAssetItemId,omitempty" name:"CustomerAssetItemId"`
	// 需要忽略指定资产内的检查项ID列表，为空表示所有

	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

type DescribeAffectedClusterCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAffectedClusterCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedClusterCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerPortListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>ContainerID - string  - 是否必填: 必填 -协议</li>
	// <li>ProtocolType - string  - 是否必填: 否 -协议</li>
	// <li>ContainerPort - uin64  - 是否必填: 否 -容器端口</li>
	// <li>HostPort - uin64  - 是否必填: 否 -宿主机端口</li>
	// <li>HostIP - string  - 是否必填: 否 -宿主机IP</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段
	// LatestFoundTime: 最近生成时间
	// AlarmCount: 告警数量

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAssetContainerPortListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerPortListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportJobResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出的状态。取值为, SUCCESS:成功、FAILURE:失败，RUNNING: 进行中。

		ExportStatus *string `json:"ExportStatus,omitempty" name:"ExportStatus"`
		// 返回下载URL

		DownloadURL *string `json:"DownloadURL,omitempty" name:"DownloadURL"`
		// 当ExportStatus为RUNNING时，返回导出进度。0~100范围的浮点数。

		ExportProgress *float64 `json:"ExportProgress,omitempty" name:"ExportProgress"`
		// 失败原因

		FailureMsg *string `json:"FailureMsg,omitempty" name:"FailureMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportJobResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportJobResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogJoinSuperObjectListRequest struct {
	*tchttp.BaseRequest

	// 日志类型 bash: "container_bash", 启动: "container_launch", "k8s": "k8s_api"

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	//
	// Status- String - 是否必填：否 - 主机状态
	// HostIP- String - 是否必填：否 - 主机内网IP
	// PublicIP- String - 是否必填：否 - 主机外网IP
	// HostName- String - 是否必填：否 - 主机名称

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 	排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeSecLogJoinSuperObjectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogJoinSuperObjectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlSystemChildRuleInfo struct {

	// 子策略Id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 策略模式,  RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截

	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`
	// 子策略状态，true为开启，false为关闭

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 子策略检测的入侵行为类型
	// CHANGE_CRONTAB：篡改计划任务
	// CHANGE_SYS_BIN：篡改系统程序
	// CHANGE_USRCFG：篡改用户配置

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

type ComplianceContainerDetailInfo struct {

	// 容器在主机上的ID。

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 容器所属的Pod的名称。

	PodName *string `json:"PodName,omitempty" name:"PodName"`
}

type VirusInfo struct {

	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// DEAL_NONE:文件待处理
	// DEAL_IGNORE:已经忽略
	// DEAL_ADD_WHITELIST:加白
	// DEAL_DEL:文件已经删除
	// DEAL_ISOLATE:已经隔离
	// DEAL_ISOLATING:隔离中
	// DEAL_ISOLATE_FAILED:隔离失败
	// DEAL_RECOVERING:恢复中
	// DEAL_RECOVER_FAILED: 恢复失败

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 事件描述

	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
	// 建议方案

	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
	// 失败子状态:
	// FILE_NOT_FOUND:文件不存在
	// FILE_ABNORMAL:文件异常
	// FILE_ABNORMAL_DEAL_RECOVER:恢复文件时，文件异常
	// BACKUP_FILE_NOT_FOUND:备份文件不存在
	// CONTAINER_NOT_FOUND_DEAL_ISOLATE:隔离时，容器不存在
	// CONTAINER_NOT_FOUND_DEAL_RECOVER:恢复时，容器不存在
	// TIMEOUT: 超时
	// TOO_MANY: 任务过多
	// OFFLINE: 离线
	// INTERNAL: 服务内部错误
	// VALIDATION: 参数非法

	SubStatus *string `json:"SubStatus,omitempty" name:"SubStatus"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// 	"NODE_DESTROYED"      //节点已销毁
	// 	"CONTAINER_EXITED"    //容器已退出
	// 	"CONTAINER_DESTROYED" //容器已销毁
	// 	"SHARED_HOST"         // 容器与主机共享网络
	// 	"RESOURCE_LIMIT"      //隔离操作资源超限
	// 	"UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// md5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 风险等级 RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 检测平台
	// 1: 云查杀引擎
	// 2: tav
	// 3: binaryAi
	// 4: 异常行为
	// 5: 威胁情报

	CheckPlatform []*string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
	// 节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// pod ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// pod(实例)的名字

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 节点所属集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 节点类型：NORMAL普通节点、SUPER超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点内网IP

	InnerIP *string `json:"InnerIP,omitempty" name:"InnerIP"`
	// 节点唯一ID

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 普通节点ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type DescribeVirusAutoIsolateSampleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 木马自动隔离样本列表

		List []*VirusAutoIsolateSampleInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSampleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePostPayDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹性计费扣费详情

		SoftQuotaDayDetail []*SoftQuotaDayInfo `json:"SoftQuotaDayDetail,omitempty" name:"SoftQuotaDayDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePostPayDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePostPayDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceTaskPolicyItemSummaryListRequest struct {
	*tchttp.BaseRequest

	// 资产类型。仅查询与指定资产类型相关的检测项。
	//
	// ASSET_CONTAINER, 容器
	//
	// ASSET_IMAGE, 镜像
	//
	// ASSET_HOST, 主机
	//
	// ASSET_K8S, K8S资产

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// Name - String
	// Name 可取值：ItemType, StandardId,  RiskLevel。
	// 当为K8S资产时，还可取ClusterName。

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceTaskPolicyItemSummaryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceTaskPolicyItemSummaryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageVirusExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetImageVirusExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageVirusExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyFlowDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量时间

		FlowTime *string `json:"FlowTime,omitempty" name:"FlowTime"`
		// 源ip

		SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
		// 源pod

		SrcPod *string `json:"SrcPod,omitempty" name:"SrcPod"`
		// 源负载

		SrcWorkLoad *string `json:"SrcWorkLoad,omitempty" name:"SrcWorkLoad"`
		// 源负载类型

		SrcWorkLoadKind *string `json:"SrcWorkLoadKind,omitempty" name:"SrcWorkLoadKind"`
		// 源命名空间

		SrcNamespace *string `json:"SrcNamespace,omitempty" name:"SrcNamespace"`
		// 源关联策略

		SrcPolicy *uint64 `json:"SrcPolicy,omitempty" name:"SrcPolicy"`
		// 源节点IP

		SrcNodeIp *string `json:"SrcNodeIp,omitempty" name:"SrcNodeIp"`
		// 源节点名称

		SrcNodeName *string `json:"SrcNodeName,omitempty" name:"SrcNodeName"`
		// 源标签

		SrcLabel *string `json:"SrcLabel,omitempty" name:"SrcLabel"`
		// 目标ip

		DestIp *string `json:"DestIp,omitempty" name:"DestIp"`
		// 目标pod

		DestPod *string `json:"DestPod,omitempty" name:"DestPod"`
		// 目标负载

		DestWorkLoad *string `json:"DestWorkLoad,omitempty" name:"DestWorkLoad"`
		// 目标关联策略

		DestPolicy *uint64 `json:"DestPolicy,omitempty" name:"DestPolicy"`
		// 目标标签

		DestLabel *string `json:"DestLabel,omitempty" name:"DestLabel"`
		// 目标负载类型

		DestWorkLoadKind *string `json:"DestWorkLoadKind,omitempty" name:"DestWorkLoadKind"`
		// 目标命名空间

		DestNamespace *string `json:"DestNamespace,omitempty" name:"DestNamespace"`
		// 目标节点ip

		DestNodeIp *string `json:"DestNodeIp,omitempty" name:"DestNodeIp"`
		// 目标节点名称

		DestNodeName *string `json:"DestNodeName,omitempty" name:"DestNodeName"`
		// 端口，和网络策略保持一致，字符串类型

		Port *string `json:"Port,omitempty" name:"Port"`
		// 动作
		//

		Action *string `json:"Action,omitempty" name:"Action"`
		// 协议

		Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
		// 标签类型

		Flag *string `json:"Flag,omitempty" name:"Flag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkTopologyFlowDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyFlowDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogDeliveryClsSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecLogDeliveryClsSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogDeliveryClsSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEditReverseShellWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddEditReverseShellWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEditReverseShellWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSearchTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteSearchTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSearchTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIACCheckTaskRequest struct {
	*tchttp.BaseRequest

	// IAC检测类型

	IACType *string `json:"IACType,omitempty" name:"IACType"`
	// IAC检测版本

	IACVersion *string `json:"IACVersion,omitempty" name:"IACVersion"`
	// IAC文件类型

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// IAC检测内容，需base64编码

	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *CreateIACCheckTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIACCheckTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyFlowListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeNetworkTopologyFlowListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyFlowListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompliancePeriodTask struct {

	// 周期任务的ID

	PeriodTaskId *uint64 `json:"PeriodTaskId,omitempty" name:"PeriodTaskId"`
	// 资产类型。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 最近一次触发的时间

	LastTriggerTime *string `json:"LastTriggerTime,omitempty" name:"LastTriggerTime"`
	// 总的检查项数目

	TotalPolicyItemCount *uint64 `json:"TotalPolicyItemCount,omitempty" name:"TotalPolicyItemCount"`
	// 周期设置

	PeriodRule *CompliancePeriodTaskRule `json:"PeriodRule,omitempty" name:"PeriodRule"`
	// 合规标准列表

	BenchmarkStandardSet []*ComplianceBenchmarkStandard `json:"BenchmarkStandardSet,omitempty" name:"BenchmarkStandardSet"`
}

type CreateSuperNodeExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>NodeID- String - 是否必填：否 - ID </li>
	// <li>NodeName- String - 是否必填：否 - 超级节点名称 </li>
	// <li>SubnetName- String - 是否必填：否 - VPC子网 </li>
	// <li>AgentStatus- String - 是否必填：否 - 安装状态NOTINSTALLED:未安装;INSTALLED:已安装;INSTALLING:安装中; </li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 需要返回的数量，默认为10，最大值为10000

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateSuperNodeExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSuperNodeExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusMonitorSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVirusMonitorSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusMonitorSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogJoinTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接入日志列表

		List []*SecLogJoinInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogJoinTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogJoinTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSuperNodeInstallTaskExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。 Status- String - 是否必填：否 - 状态 INSTALLING:安装中; INSTALLATIONFAILURE:安装失败; INSTALLATIONSUCCESS:安装成功 PARTIALINSTALLATIONFAILURE:部分安装失败 Source- String - 是否必填：否 - 安装来源: TCSS:容器安全服务 EKS:容器服务 TimeRange- String - 是否必填：否 - 安装时间范围,传一个时间默认为开始时间,两个为开始时间和结束时间

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 需要返回的数量，默认为10，最大值为10000

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateSuperNodeInstallTaskExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSuperNodeInstallTaskExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckRepeatAssetImageRegistryRequest struct {
	*tchttp.BaseRequest

	// 仓库名

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CheckRepeatAssetImageRegistryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRepeatAssetImageRegistryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`
		// 进程信息

		ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`
		// 父进程信息

		ParentProcessInfo *ProcessDetailBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`
		// 事件描述

		EventDetail *AbnormalProcessEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`
		// 祖先进程信息

		AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSyncLastTimeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetSyncLastTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncLastTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventDetailRequest struct {
	*tchttp.BaseRequest

	// 事件唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
}

func (r *DescribeEscapeEventDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceFilters struct {

	// 过滤键的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 是否模糊查询。默认为是。

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type CreateComplianceTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的合规检查任务的ID。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateComplianceTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateComplianceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscapeEventTendencyInfo struct {

	// 待处理风险容器事件总数

	RiskContainerEventCount *int64 `json:"RiskContainerEventCount,omitempty" name:"RiskContainerEventCount"`
	// 待处理程序特权事件总数

	ProcessPrivilegeEventCount *int64 `json:"ProcessPrivilegeEventCount,omitempty" name:"ProcessPrivilegeEventCount"`
	// 待处理容器逃逸事件总数

	ContainerEscapeEventCount *int64 `json:"ContainerEscapeEventCount,omitempty" name:"ContainerEscapeEventCount"`
	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
}

type DeleteMachineRequest struct {
	*tchttp.BaseRequest

	// 客户端Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DeleteMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddCompliancePolicyItemToWhitelistRequest struct {
	*tchttp.BaseRequest

	// 要忽略的检测项的ID的列表

	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

func (r *AddCompliancePolicyItemToWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCompliancePolicyItemToWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessControlRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 策略的ids

	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`
	// 策略开关，true:代表开启， false代表关闭

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
}

func (r *ModifyAccessControlRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVirusListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageVirusListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVirusListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlEventsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAccessControlEventsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlEventsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanIgnoreVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总是

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞列表

		List []*ScanIgnoreVul `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanIgnoreVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanIgnoreVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCompliancePolicyAssetSetFromWhitelistRequest struct {
	*tchttp.BaseRequest

	// （检查项ID+资产ID列表）的列表

	PolicyAssetSetList []*CompliancePolicyAssetSetItem `json:"PolicyAssetSetList,omitempty" name:"PolicyAssetSetList"`
}

func (r *DeleteCompliancePolicyAssetSetFromWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCompliancePolicyAssetSetFromWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkFirewallPublishRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略Id数组

	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *CreateNetworkFirewallPublishRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkFirewallPublishRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CkafkaRouteInfo struct {

	// 路由ID

	RouteID *int64 `json:"RouteID,omitempty" name:"RouteID"`
	// 域名名称

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名端口

	DomainPort *uint64 `json:"DomainPort,omitempty" name:"DomainPort"`
	// 虚拟ip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 虚拟ip类型

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
	// 接入类型
	// // 0：PLAINTEXT (明文方式，没有带用户信息老版本及社区版本都支持)
	// 	// 1：SASL_PLAINTEXT（明文方式，不过在数据开始时，会通过SASL方式登录鉴权，仅社区版本支持）
	// 	// 2：SSL（SSL加密通信，没有带用户信息，老版本及社区版本都支持）
	// 	// 3：SASL_SSL（SSL加密通信，在数据开始时，会通过SASL方式登录鉴权，仅社区版本支持）

	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`
}

type DescribeNetworkFirewallNamespaceLabelListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群空间标签详细信息

		ClusterNamespaceLabelList []*NetworkClusterNamespaceLabelInfo `json:"ClusterNamespaceLabelList,omitempty" name:"ClusterNamespaceLabelList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallNamespaceLabelListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallNamespaceLabelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchImageAutoAuthorizedRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchImageAutoAuthorizedRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchImageAutoAuthorizedRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulLevelSummaryRequest struct {
	*tchttp.BaseRequest

	// 漏洞分类: SYSTEM:系统漏洞 WEB:web应用漏洞 EMERGENCY:应急漏洞

	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
}

func (r *DescribeVulLevelSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulLevelSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageAutoAuthorizedTaskListRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤字段
	// Status授权结果，全部授权成功：ALLSUCCSESS，部分授权失败：PARTIALFAIL,全部授权失败：ALLFAIL
	// Type授权方式，AUTO:自动授权，MANUAL:手动授权
	// Source 镜像来源，LOCAL:本地镜像，REGISTRY:仓库镜像

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeImageAutoAuthorizedTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageAutoAuthorizedTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyK8sApiAbnormalRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyK8sApiAbnormalRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyK8sApiAbnormalRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRepoInfo struct {

	// 镜像Digest

	ImageDigest *string `json:"ImageDigest,omitempty" name:"ImageDigest"`
	// 镜像仓库地址

	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`
	// 仓库类型

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像版本

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 镜像大小

	ImageSize *uint64 `json:"ImageSize,omitempty" name:"ImageSize"`
	// 最近扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 扫描状态

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 安全漏洞数

	VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`
	// 木马病毒数

	VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`
	// 风险行为数

	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
	// 敏感信息数

	SentiveInfoCnt *uint64 `json:"SentiveInfoCnt,omitempty" name:"SentiveInfoCnt"`
	// 是否可信镜像

	IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`
	// 镜像系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 木马扫描错误

	ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`
	// 漏洞扫描错误

	ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 高危扫描错误

	ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`
	// 敏感信息扫描进度

	ScanVirusProgress *uint64 `json:"ScanVirusProgress,omitempty" name:"ScanVirusProgress"`
	// 木马扫描进度

	ScanVulProgress *uint64 `json:"ScanVulProgress,omitempty" name:"ScanVulProgress"`
	// 漏洞扫描进度

	ScanRiskProgress *uint64 `json:"ScanRiskProgress,omitempty" name:"ScanRiskProgress"`
	// 剩余扫描时间秒

	ScanRemainTime *uint64 `json:"ScanRemainTime,omitempty" name:"ScanRemainTime"`
	// cve扫描状态

	CveStatus *string `json:"CveStatus,omitempty" name:"CveStatus"`
	// 高危扫描状态

	RiskStatus *string `json:"RiskStatus,omitempty" name:"RiskStatus"`
	// 木马扫描状态

	VirusStatus *string `json:"VirusStatus,omitempty" name:"VirusStatus"`
	// 总进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 授权状态

	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
	// 仓库区域

	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`
	// 列表id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 镜像Id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像创建的时间

	ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`
	// 是否为镜像的最新版本

	IsLatestImage *bool `json:"IsLatestImage,omitempty" name:"IsLatestImage"`
	// low级别漏洞个数

	LowLevelVulCnt *uint64 `json:"LowLevelVulCnt,omitempty" name:"LowLevelVulCnt"`
	// medium级别漏洞个数

	MediumLevelVulCnt *uint64 `json:"MediumLevelVulCnt,omitempty" name:"MediumLevelVulCnt"`
	// high级别漏洞个数

	HighLevelVulCnt *uint64 `json:"HighLevelVulCnt,omitempty" name:"HighLevelVulCnt"`
	// critical级别漏洞个数

	CriticalLevelVulCnt *uint64 `json:"CriticalLevelVulCnt,omitempty" name:"CriticalLevelVulCnt"`
	// 关联容器数

	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 组件数

	ComponentCnt *uint64 `json:"ComponentCnt,omitempty" name:"ComponentCnt"`
	// 是否运行中

	IsRunning *bool `json:"IsRunning,omitempty" name:"IsRunning"`
	// 是否存在必修漏洞

	HasNeedFixVul *bool `json:"HasNeedFixVul,omitempty" name:"HasNeedFixVul"`
	// 敏感信息

	SensitiveInfoCnt *uint64 `json:"SensitiveInfoCnt,omitempty" name:"SensitiveInfoCnt"`
	// 是否推荐处置

	RecommendedFix *bool `json:"RecommendedFix,omitempty" name:"RecommendedFix"`
}

type DescribeValueAddedSrvInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仓库镜像未授权数量

		RegistryImageCnt *uint64 `json:"RegistryImageCnt,omitempty" name:"RegistryImageCnt"`
		// 本地镜像未授权数量

		LocalImageCnt *uint64 `json:"LocalImageCnt,omitempty" name:"LocalImageCnt"`
		// 未使用的镜像安全扫描授权数

		UnusedAuthorizedCnt *uint64 `json:"UnusedAuthorizedCnt,omitempty" name:"UnusedAuthorizedCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeValueAddedSrvInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeValueAddedSrvInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePolicyItemAffectedSummaryRequest struct {
	*tchttp.BaseRequest

	// DescribeComplianceTaskPolicyItemSummaryList返回的CustomerPolicyItemId，表示检测项的ID。

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
}

func (r *DescribeCompliancePolicyItemAffectedSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePolicyItemAffectedSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentDaemonSetCmdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安装命令

		Command *string `json:"Command,omitempty" name:"Command"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentDaemonSetCmdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentDaemonSetCmdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellRegexpWhiteListRequest struct {
	*tchttp.BaseRequest

	// 规则ID数组

	RuleIDs []*string `json:"RuleIDs,omitempty" name:"RuleIDs"`
}

func (r *DeleteReverseShellRegexpWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellRegexpWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusEstimateScanTimeRequest struct {
	*tchttp.BaseRequest

	// SCAN_NODE:扫描节点 SCAN_CONTAINER:扫描容器

	ScanRangeType *string `json:"ScanRangeType,omitempty" name:"ScanRangeType"`
	// 自选扫描范围的容器id或者节点id

	ScanIDs []*ScanRangeInfo `json:"ScanIDs,omitempty" name:"ScanIDs"`
}

func (r *DescribeVirusEstimateScanTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusEstimateScanTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrModifyImageDenyRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOrModifyImageDenyRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrModifyImageDenyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusMonitorSettingRequest struct {
	*tchttp.BaseRequest

	// 是否开启定期扫描

	EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`
	// 扫描全部路径

	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`
	// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径(扫描范围只能小于等于1)

	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`
	// 自选排除或扫描的地址

	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
	// 扫描路径模式：
	// SCAN_PATH_ALL：全部路径
	// SCAN_PATH_DEFAULT：默认路径
	// SCAN_PATH_USER_DEFINE：用户自定义路径
	//

	ScanPathMode *string `json:"ScanPathMode,omitempty" name:"ScanPathMode"`
}

func (r *ModifyVirusMonitorSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusMonitorSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageScanTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetImageScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRiskSyscallNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppServiceListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetAppServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompliancePolicyItemSummary struct {

	// 为客户分配的唯一的检测项的ID。

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 检测项的原始ID。

	BasePolicyItemId *uint64 `json:"BasePolicyItemId,omitempty" name:"BasePolicyItemId"`
	// 检测项的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 检测项所属的类型，枚举字符串。

	Category *string `json:"Category,omitempty" name:"Category"`
	// 所属的合规标准

	BenchmarkStandardName *string `json:"BenchmarkStandardName,omitempty" name:"BenchmarkStandardName"`
	// 威胁等级。RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 检测项所属的资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 最近检测的时间

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 检测状态
	//
	// CHECK_INIT, 待检测
	//
	// CHECK_RUNNING, 检测中
	//
	// CHECK_FINISHED, 检测完成
	//
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 检测结果。RESULT_PASSED: 通过
	//
	// RESULT_FAILED: 未通过

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 通过检测的资产的数目

	PassedAssetCount *uint64 `json:"PassedAssetCount,omitempty" name:"PassedAssetCount"`
	// 未通过检测的资产的数目

	FailedAssetCount *uint64 `json:"FailedAssetCount,omitempty" name:"FailedAssetCount"`
	// 检测项对应的白名单项的ID。如果存在且非0，表示检测项被用户忽略。

	WhitelistId *uint64 `json:"WhitelistId,omitempty" name:"WhitelistId"`
	// 处理建议。

	FixSuggestion *string `json:"FixSuggestion,omitempty" name:"FixSuggestion"`
	// 所属的合规标准的ID

	BenchmarkStandardId *uint64 `json:"BenchmarkStandardId,omitempty" name:"BenchmarkStandardId"`
	// 检测项适用的版本

	ApplicableVersion *string `json:"ApplicableVersion,omitempty" name:"ApplicableVersion"`
	// 检查项描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 检查项审计方法

	AuditProcedure *string `json:"AuditProcedure,omitempty" name:"AuditProcedure"`
}

type DescribeNamespaceListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：Namespace

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNamespaceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespaceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageRegistryAutoAuthorizedRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageRegistryAutoAuthorizedRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageRegistryAutoAuthorizedRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeSafeStateRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeEscapeSafeStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeSafeStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCheckItemListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name 可取值：
	// Name: 检查项名称
	// RiskType: 风险类别
	// RiskLevel: 风险等级
	// RiskTarget: 检查对象
	// RiskAttribute: 检测项所属分型线类型
	// Enable: 检查项是否开启(0:关闭 1:开启)

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCheckItemListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCheckItemListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReverseShellRegexpRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 启用状态

	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyReverseShellRegexpRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReverseShellRegexpRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterWorkloadFromWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单工作负载ID列表信息

	WorkloadIdList []*uint64 `json:"WorkloadIdList,omitempty" name:"WorkloadIdList"`
}

func (r *DeleteClusterWorkloadFromWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterWorkloadFromWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterCheckTimerSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterCheckTimerSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterCheckTimerSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEscapeEventStatusRequest struct {
	*tchttp.BaseRequest

	// 处理事件ids

	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
	// 标记事件的状态：
	// EVENT_UNDEAL:未处理（取消忽略），
	// EVENT_DEALED:已处理，
	// EVENT_IGNORE:忽略，
	// EVENT_DELETE：已删除
	// EVENT_ADD_WHITE：加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 加白镜像ID数组

	ImageIDs []*string `json:"ImageIDs,omitempty" name:"ImageIDs"`
	// 加白事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸

	EventType []*string `json:"EventType,omitempty" name:"EventType"`
}

func (r *ModifyEscapeEventStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEscapeEventStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallEventsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeRiskSyscallEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessControlStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessControlStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetImageScanStopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 停止状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetImageScanStopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetImageScanStopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryRiskInfoListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像id

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryRiskInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRiskInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerPortInfo struct {

	// 宿主机IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 宿主机端口

	HostPort *uint64 `json:"HostPort,omitempty" name:"HostPort"`
	// 容器端口

	ContainerPort *uint64 `json:"ContainerPort,omitempty" name:"ContainerPort"`
	// 协议类型

	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
}

type DescribeVirusScanTimeoutSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 超时时长单位小时

		Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusScanTimeoutSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanTimeoutSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProcessEventsExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,Status：EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：latest_found_time

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateProcessEventsExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProcessEventsExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallClusterRefreshStatusRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeNetworkFirewallClusterRefreshStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallClusterRefreshStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentInstallCommandResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// linux系统安装命令

		LinuxCommand *string `json:"LinuxCommand,omitempty" name:"LinuxCommand"`
		// windows系统安装命令（windows2008及以上）

		WindowsCommand *string `json:"WindowsCommand,omitempty" name:"WindowsCommand"`
		// windows系统安装命令第一步（windows2003）

		WindowsStepOne *string `json:"WindowsStepOne,omitempty" name:"WindowsStepOne"`
		// windows系统安装命令第二步（windows2003）

		WindowsStepTwo *string `json:"WindowsStepTwo,omitempty" name:"WindowsStepTwo"`
		// windows版agent下载链接

		WindowsDownloadUrl *string `json:"WindowsDownloadUrl,omitempty" name:"WindowsDownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentInstallCommandResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentInstallCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventEscapeImageInfo struct {

	// 镜像id，用于跳转

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 唯一值

	UniqueKey *string `json:"UniqueKey,omitempty" name:"UniqueKey"`
	// 事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 原始事件类型

	OriginEventType *string `json:"OriginEventType,omitempty" name:"OriginEventType"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 容器数量

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 状态，EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略

	Status *string `json:"Status,omitempty" name:"Status"`
	// 风险描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
}

type VulTendencyInfo struct {

	// 漏洞趋势列表

	VulSet []*RunTimeTendencyInfo `json:"VulSet,omitempty" name:"VulSet"`
	// 漏洞影响的镜像类型：
	// LOCAL：本地镜像
	// REGISTRY: 仓库镜像

	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
}

type DescribeAssetImageListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAssetImageListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterCustomParameters struct {

	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ComponentsInfo struct {

	// 组件名称

	Component *string `json:"Component,omitempty" name:"Component"`
	// 组件版本信息

	Version *string `json:"Version,omitempty" name:"Version"`
	// 可修复版本

	FixedVersion *string `json:"FixedVersion,omitempty" name:"FixedVersion"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type NetworkClusterPodInfo struct {

	// pod名字

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// pod空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// pod标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// pod类型

	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`
}

type CreateRiskDnsEventExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>EventStatus- String - 是否必填：否 - 事件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_ADD_WHITE：已加白</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>RiskDns- string - 是否必填：否 - 恶意域名。</li>
	// <li>RiskIP- string - 是否必填：否 - 恶意IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，最大值为100000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：事件数量：EventCount

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateRiskDnsEventExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskDnsEventExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanLocalImageListRequest struct {
	*tchttp.BaseRequest

	// 漏洞扫描任务ID

	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ScanStatus- string -是否必填: 否 - 检测状态。WAIT_SCAN：待检测，SCANNING：检查中，SCANNED：检查完成，SCAN_ERR：检查失败，CANCELED：检测停止</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulScanLocalImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanLocalImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVirusScanTaskRequest struct {
	*tchttp.BaseRequest

	// 是否扫描所有路径

	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`
	// 当ScanPathAll为false生效 0扫描以下路径 1、扫描除以下路径

	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`
	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定

	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`
	// 扫描范围0容器1主机节点

	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`
	// true 全选，false 自选

	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`
	// 超时时长，单位小时

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 自选排除或扫描的地址

	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
	// 扫描路径模式：
	// SCAN_PATH_ALL：全部路径
	// SCAN_PATH_DEFAULT：默认路径
	// SCAN_PATH_USER_DEFINE：用户自定义路径
	//

	ScanPathMode *string `json:"ScanPathMode,omitempty" name:"ScanPathMode"`
}

func (r *CreateVirusScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVirusScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyK8sApiAbnormalEventStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyK8sApiAbnormalEventStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyK8sApiAbnormalEventStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceScanFailedAssetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回检测失败的资产的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回各类检测失败的资产的汇总信息的列表。

		ScanFailedAssetList []*ComplianceScanFailedAsset `json:"ScanFailedAssetList,omitempty" name:"ScanFailedAssetList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceScanFailedAssetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceScanFailedAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulImageExportJobRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ClientIP- string -是否必填: 否 - 内网IP</li>
	// <li>PublicIP- string -是否必填: 否 - 外网IP</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>HostName- string -是否必填: 否 - 主机名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateVulImageExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulImageExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkFirewallClusterRefreshResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的集群检查任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建检查任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkFirewallClusterRefreshResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkFirewallClusterRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表

		RegionList []*ClusterRegionItem `json:"RegionList,omitempty" name:"RegionList"`
		// 数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSupportRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSupportRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetPolicyItemListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回检测项的总数。如果用户未启用基线检查，此处返回0。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回某个资产下的检测项的列表。

		AssetPolicyItemList []*ComplianceAssetPolicyItem `json:"AssetPolicyItemList,omitempty" name:"AssetPolicyItemList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceAssetPolicyItemListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetPolicyItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageScanSettingRequest struct {
	*tchttp.BaseRequest

	// 开关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 扫描开始时间
	// 01:00 时分

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 扫描周期

	ScanPeriod *uint64 `json:"ScanPeriod,omitempty" name:"ScanPeriod"`
	// 扫描木马

	ScanVirus *bool `json:"ScanVirus,omitempty" name:"ScanVirus"`
	// 扫描敏感信息

	ScanRisk *bool `json:"ScanRisk,omitempty" name:"ScanRisk"`
	// 扫描漏洞

	ScanVul *bool `json:"ScanVul,omitempty" name:"ScanVul"`
	// 全部镜像

	All *bool `json:"All,omitempty" name:"All"`
	// 自定义镜像

	Images []*string `json:"Images,omitempty" name:"Images"`
	// 镜像是否存在运行中的容器

	ContainerRunning *bool `json:"ContainerRunning,omitempty" name:"ContainerRunning"`
	// 扫描范围 0 全部授权镜像，1自选镜像，2 推荐扫描

	ScanScope *uint64 `json:"ScanScope,omitempty" name:"ScanScope"`
	// 扫描结束时间
	// 02:00 时分

	ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`
}

func (r *CreateAssetImageScanSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageScanSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogJoinObjectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 接入对象列表

		List []*SecLogJoinObjectInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogJoinObjectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogJoinObjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyEventTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像拦截成功事件趋势

		DenyList []*ImageDenyEventTendency `json:"DenyList,omitempty" name:"DenyList"`
		// 镜像拦截告警事件趋势

		AlarmList []*ImageDenyEventTendency `json:"AlarmList,omitempty" name:"AlarmList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyEventTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlEventsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 任务id

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlEventsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnfinishRefreshTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUnfinishRefreshTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnfinishRefreshTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewImageAuthorizeStateRequest struct {
	*tchttp.BaseRequest

	// 镜像ids

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 是否全部未授权镜像

	AllImages *bool `json:"AllImages,omitempty" name:"AllImages"`
	// 是否授权后自动扫描

	NeedScan *bool `json:"NeedScan,omitempty" name:"NeedScan"`
	// 扫描类型

	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`
}

func (r *RenewImageAuthorizeStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewImageAuthorizeStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirusTaskInfo struct {

	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像Id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 节点名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 节点内网ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 扫描状态：
	// WAIT: 等待扫描
	// FAILED: 失败
	// SCANNING: 扫描中
	// FINISHED: 结束
	// CANCELING: 取消中
	// CANCELED: 已取消
	// CANCEL_FAILED： 取消失败

	Status *string `json:"Status,omitempty" name:"Status"`
	// 检测开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 检测结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 风险个数

	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
	// 事件id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 错误原因:
	// SEND_SUCCESSED: 下发成功
	// SCAN_WAIT: agent排队扫描等待中
	// OFFLINE: 离线
	// SEND_FAILED:下发失败
	// TIMEOUT: 超时
	// LOW_AGENT_VERSION: 客户端版本过低
	// AGENT_NOT_FOUND: 镜像所属客户端版不存在
	// TOO_MANY: 任务过多
	// VALIDATION: 参数非法
	// INTERNAL: 服务内部错误
	// MISC: 其他错误
	// UNAUTH: 所在镜像未授权
	// SEND_CANCEL_SUCCESSED:下发成功

	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
	// 节点类型：NORMAL普通节点、SUPER超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
}

type AddEscapeWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddEscapeWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEscapeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessEventInfo struct {

	// 进程目录

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 事件类型，MALICE_PROCESS_START:恶意进程启动

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 命中规则名称，PROXY_TOOL：代理软件，TRANSFER_CONTROL：横向渗透，ATTACK_CMD：恶意命令，REVERSE_SHELL：反弹shell，FILELESS：无文件程序执行，RISK_CMD：高危命令，ABNORMAL_CHILD_PROC：敏感服务异常子进程启动，USER_DEFINED_RULE：用户自定义规则

	MatchRuleName *string `json:"MatchRuleName,omitempty" name:"MatchRuleName"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 动作执行结果，    BEHAVIOR_NONE: 无
	//     BEHAVIOR_ALERT: 告警
	//     BEHAVIOR_RELEASE：放行
	//     BEHAVIOR_HOLDUP_FAILED:拦截失败
	//     BEHAVIOR_HOLDUP_SUCCESSED：拦截失败

	Behavior *string `json:"Behavior,omitempty" name:"Behavior"`
	// 状态，EVENT_UNDEAL:事件未处理
	//     EVENT_DEALED:事件已经处理
	//     EVENT_INGNORE：事件已经忽略

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件记录的唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 镜像id，用于跳转

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器id，用于跳转

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 事件解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件详细描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 命中策略id

	MatchRuleId *string `json:"MatchRuleId,omitempty" name:"MatchRuleId"`
	// 命中规则行为：
	// RULE_MODE_RELEASE 放行
	// RULE_MODE_ALERT  告警
	// RULE_MODE_HOLDUP 拦截

	MatchAction *string `json:"MatchAction,omitempty" name:"MatchAction"`
	// 命中规则进程信息

	MatchProcessPath *string `json:"MatchProcessPath,omitempty" name:"MatchProcessPath"`
	// 规则是否存在

	RuleExist *bool `json:"RuleExist,omitempty" name:"RuleExist"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 规则组Id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 命中策略名称：SYSTEM_DEFINED_RULE （系统策略）或  用户自定义的策略名字

	MatchGroupName *string `json:"MatchGroupName,omitempty" name:"MatchGroupName"`
	// 命中规则等级，HIGH：高危，MIDDLE：中危，LOW：低危。

	MatchRuleLevel *string `json:"MatchRuleLevel,omitempty" name:"MatchRuleLevel"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 节点类型：NORMAL普通节点、SUPER超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// pod 名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// pod ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 集群id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 节点公网ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// uuid

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 节点内网ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type RegistryConnDetectResult struct {

	// 联通性检测的主机quuid 或者 backend

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 联通性检测的主机uuid 或者 backend

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 检测结果状态

	ConnDetectStatus *string `json:"ConnDetectStatus,omitempty" name:"ConnDetectStatus"`
	// 检测结果信息

	ConnDetectMessage *string `json:"ConnDetectMessage,omitempty" name:"ConnDetectMessage"`
	// 失败的解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 失败原因

	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
}

type SearchTemplate struct {

	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 检索名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 检索索引类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 检索语句

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 时间范围

	TimeRange *string `json:"TimeRange,omitempty" name:"TimeRange"`
	// 转换的检索语句内容

	Query *string `json:"Query,omitempty" name:"Query"`
	// 检索方式。输入框检索：standard,过滤，检索：simple

	Flag *string `json:"Flag,omitempty" name:"Flag"`
	// 展示数据

	DisplayData *string `json:"DisplayData,omitempty" name:"DisplayData"`
}

type DeleteK8sApiAbnormalRuleRequest struct {
	*tchttp.BaseRequest

	// 规则ID集合

	RuleIDSet []*string `json:"RuleIDSet,omitempty" name:"RuleIDSet"`
}

func (r *DeleteK8sApiAbnormalRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteK8sApiAbnormalRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulnerabilityFocusRuleDetailRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulnerabilityFocusRuleDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulnerabilityFocusRuleDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSuperNodePodExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>NodeUniqueID- String - 是否必填：否 - 节点唯一id </li>
	// <li>PodName- String - 是否必填：否 - Pod示例名称 </li>
	// <li>PodIP- String - 是否必填：否 - POD IP </li>
	// <li>NameSpace- String - 是否必填：否 - 命名空间 </li>
	// <li>Deployment- String - 是否必填：否 - 所属工作负载 </li>
	// <li>Status- String - 是否必填：否 - 状态 </li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 需要返回的数量，默认为10，最大值为10000

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateSuperNodePodExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSuperNodePodExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogAlertMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警消息队列

		List []*SecLogAlertMsgInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogAlertMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogAlertMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusTaskListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ContainerName - String - 是否必填：否 - 容器名称</li>
	// <li>ContainerId - String - 是否必填：否 - 容器id</li>
	// <li>Hostname - String - 是否必填：否 - 主机名称</li>
	// <li>HostIp- String - 是否必填：否 - 主机IP</li>
	// <li>ImageId- String - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String - 是否必填：否 - 镜像名称</li>
	// <li>Status- String - 是否必填：否 - 状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVirusTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAffectedClusterCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 严重风险的集群数量

		SeriousRiskClusterCount *uint64 `json:"SeriousRiskClusterCount,omitempty" name:"SeriousRiskClusterCount"`
		// 高危风险的集群数量

		HighRiskClusterCount *uint64 `json:"HighRiskClusterCount,omitempty" name:"HighRiskClusterCount"`
		// 中危风险的集群数量

		MiddleRiskClusterCount *uint64 `json:"MiddleRiskClusterCount,omitempty" name:"MiddleRiskClusterCount"`
		// 低危风险的集群数量

		HintRiskClusterCount *uint64 `json:"HintRiskClusterCount,omitempty" name:"HintRiskClusterCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAffectedClusterCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedClusterCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公钥

		PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`
		// 进程信息

		ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`
		// 父进程信息

		ParentProcessInfo *ProcessDetailBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`
		// 事件描述

		EventDetail *ReverseShellEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`
		// 祖先进程信息

		AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyK8sApiAbnormalRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyK8sApiAbnormalRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyK8sApiAbnormalRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogCleanSettingInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 触发清理的储量底线

		ReservesLimit *uint64 `json:"ReservesLimit,omitempty" name:"ReservesLimit"`
		// 清理停止时的储量截至线

		ReservesDeadline *uint64 `json:"ReservesDeadline,omitempty" name:"ReservesDeadline"`
		// 触发清理的存储天数

		DayLimit *uint64 `json:"DayLimit,omitempty" name:"DayLimit"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogCleanSettingInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogCleanSettingInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchTemplatesRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSearchTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 严重的风险数量

		SeriousRiskCount *uint64 `json:"SeriousRiskCount,omitempty" name:"SeriousRiskCount"`
		// 高危的风险数量

		HighRiskCount *uint64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
		// 中危的风险数量

		MiddleRiskCount *uint64 `json:"MiddleRiskCount,omitempty" name:"MiddleRiskCount"`
		// 提示的风险数量

		HintRiskCount *uint64 `json:"HintRiskCount,omitempty" name:"HintRiskCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterRiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusEventTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势列表

		List []*VirusTendencyInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusEventTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogStorageStatisticRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogStorageStatisticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogStorageStatisticRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInfoItem struct {

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 集群操作系统

	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群节点数

	ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitempty" name:"ClusterNodeNum"`
	// 集群区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 监控组件的状态，为Defender_Uninstall、Defender_Normal、Defender_Error、Defender_Installing

	DefenderStatus *string `json:"DefenderStatus,omitempty" name:"DefenderStatus"`
	// 集群状态

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 集群的检测模式，为Cluster_Normal或者Cluster_Actived.

	ClusterCheckMode *string `json:"ClusterCheckMode,omitempty" name:"ClusterCheckMode"`
	// 是否自动定期检测

	ClusterAutoCheck *bool `json:"ClusterAutoCheck,omitempty" name:"ClusterAutoCheck"`
	// 防护容器部署失败原因，为UserDaemonSetNotReady时,和UnreadyNodeNum转成"N个节点防御容器为就绪"，其他错误直接展示

	DefenderErrorReason *string `json:"DefenderErrorReason,omitempty" name:"DefenderErrorReason"`
	// 防御容器没有ready状态的节点数量

	UnreadyNodeNum *uint64 `json:"UnreadyNodeNum,omitempty" name:"UnreadyNodeNum"`
	// 严重风险检查项的数量

	SeriousRiskCount *int64 `json:"SeriousRiskCount,omitempty" name:"SeriousRiskCount"`
	// 高风险检查项的数量

	HighRiskCount *int64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
	// 中风险检查项的数量

	MiddleRiskCount *int64 `json:"MiddleRiskCount,omitempty" name:"MiddleRiskCount"`
	// 提示风险检查项的数量

	HintRiskCount *int64 `json:"HintRiskCount,omitempty" name:"HintRiskCount"`
	// 检查失败原因

	CheckFailReason *string `json:"CheckFailReason,omitempty" name:"CheckFailReason"`
	// 检查状态,为Task_Running, NoRisk, HasRisk, Uncheck, Task_Error

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 任务创建时间,检查时间

	TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`
	// 接入状态:
	// 未接入: AccessedNone
	// 已防护: AccessedDefended
	// 未防护: AccessedInstalled
	// 部分防护: AccessedPartialDefence
	// 接入异常: AccessedException
	// 卸载异常: AccessedUninstallException
	// 接入中: AccessedInstalling
	// 卸载中: AccessedUninstalling

	AccessedStatus *string `json:"AccessedStatus,omitempty" name:"AccessedStatus"`
	// 接入失败原因

	AccessedSubStatus *string `json:"AccessedSubStatus,omitempty" name:"AccessedSubStatus"`
	// 节点总数

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 离线节点数

	OffLineNodeCount *uint64 `json:"OffLineNodeCount,omitempty" name:"OffLineNodeCount"`
	// 未安装agent节点数

	UnInstallAgentNodeCount *uint64 `json:"UnInstallAgentNodeCount,omitempty" name:"UnInstallAgentNodeCount"`
	// 计费核数

	ChargeCoresCnt *uint64 `json:"ChargeCoresCnt,omitempty" name:"ChargeCoresCnt"`
}

type CreateDefenceVulExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDefenceVulExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDefenceVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterNamespaceInfo struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type CreateHostExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - agent状态筛选，"ALL":"全部"(或不传该字段),"UNINSTALL"："未安装","OFFLINE"："离线", "ONLINE"："防护中"</li>
	// <li>HostName - String - 是否必填：否 - 主机名筛选</li>
	// <li>Group- String - 是否必填：否 - 主机群组搜索</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>HostID- string - 是否必填：否 - 主机id搜索</li>
	// <li>DockerVersion- string - 是否必填：否 - docker版本搜索</li>
	// <li>MachineType- string - 是否必填：否 - 主机来源MachineType搜索，"ALL":"全部"(或不传该字段),主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；</li>
	// <li>DockerStatus- string - 是否必填：否 - docker安装状态，"ALL":"全部"(或不传该字段),"INSTALL":"已安装","UNINSTALL":"未安装"</li>
	// <li>ProjectID- string - 是否必填：否 - 所属项目id搜索</li>
	// <li>Tag:xxx(tag:key)- string- 是否必填：否 - 标签键值搜索 示例Filters":[{"Name":"tag:tke-kind","Values":["service"]}]</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 需要返回的数量，默认为10，最大值为10000

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateHostExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryRiskListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像信息

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryRiskListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRiskListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecLogDeliveryClsSettingInfo struct {

	// 日志类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 投递状态(true:开启 false:关闭)

	State *bool `json:"State,omitempty" name:"State"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 日志集

	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`
	// 主题ID

	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
	// 日志集名称

	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type CreateHostExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHostExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulIgnoreRegistryImageListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
}

func (r *DescribeVulIgnoreRegistryImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulIgnoreRegistryImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式，asc，desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 是否仅展示repository版本最新的镜像，默认为false

	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
}

func (r *DescribeAssetImageRegistryListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRiskDnsSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePostPayDetailRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePostPayDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePostPayDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetClusterListItem struct {

	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群状态
	// CSR_RUNNING: 运行中
	// CSR_EXCEPTION:异常
	// CSR_DEL:已经删除

	Status *string `json:"Status,omitempty" name:"Status"`
	// 绑定规则名称

	BindRuleName *string `json:"BindRuleName,omitempty" name:"BindRuleName"`
	// 集群类型:
	// CT_TKE:TKE集群;
	// CT_USER_CREATE:用户自建集群;
	// CT_TKE_SERVERLESS:TKE Serverless集群;

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 内存量

	MemLimit *int64 `json:"MemLimit,omitempty" name:"MemLimit"`
	// cpu

	CpuLimit *int64 `json:"CpuLimit,omitempty" name:"CpuLimit"`
}

type ModifyVirusAutoIsolateExampleSwitchRequest struct {
	*tchttp.BaseRequest

	// 文件Md5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 开关(开:true 关: false)

	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyVirusAutoIsolateExampleSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirusAutoIsolateExampleSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkPeer struct {

	// 对象类型：
	//
	// 命名空间：NamespaceSelector，代表NamespaceSelector有值
	//
	// pod类型：PodSelector，代表NamespaceSelector和PodSelector都有值
	//
	// ip类型：IPBlock，代表只有IPBlock有值

	PeerType *string `json:"PeerType,omitempty" name:"PeerType"`
	// 空间选择器

	NamespaceSelector *string `json:"NamespaceSelector,omitempty" name:"NamespaceSelector"`
	// pod选择器

	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
	// Ip选择器

	IPBlock *string `json:"IPBlock,omitempty" name:"IPBlock"`
}

type WarningRule struct {

	// 告警事件类型：
	// 镜像仓库安全-木马：IMG_REG_VIRUS
	// 镜像仓库安全-漏洞：IMG_REG_VUL
	// 镜像仓库安全-敏感信息：IMG_REG_RISK
	// 镜像安全-木马：IMG_VIRUS
	// 镜像安全-漏洞：IMG_VUL
	// 镜像安全-敏感信息：IMG_RISK
	// 镜像安全-镜像拦截：IMG_INTERCEPT
	// 运行时安全-容器逃逸：RUNTIME_ESCAPE
	// 运行时安全-异常进程：RUNTIME_FILE
	// 运行时安全-异常文件访问：RUNTIME_PROCESS
	// 运行时安全-高危系统调用：RUNTIME_SYSCALL
	// 运行时安全-反弹Shell：RUNTIME_REVERSE_SHELL
	// 运行时安全-木马：RUNTIME_VIRUS

	Type *string `json:"Type,omitempty" name:"Type"`
	// 开关状态：
	// 打开：ON
	// 关闭：OFF

	Switch *string `json:"Switch,omitempty" name:"Switch"`
	// 告警开始时间，格式: HH:mm

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 告警结束时间，格式: HH:mm

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 告警等级策略控制，二进制位每位代表一个含义，值以字符串类型传递
	// 控制开关分为高、中、低，则二进制位分别为：第1位:低，第2位:中，第3位:高，0表示关闭、1表示打开。
	// 如：高危和中危打开告警，低危关闭告警，则二进制值为：110
	// 告警类型不区分等级控制，则传1。

	ControlBits *string `json:"ControlBits,omitempty" name:"ControlBits"`
}

type DescribePublicProxyInstallCommandResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 保活命令

		KeepAliveCommand *string `json:"KeepAliveCommand,omitempty" name:"KeepAliveCommand"`
		// nginx命令

		NginxCommand *string `json:"NginxCommand,omitempty" name:"NginxCommand"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicProxyInstallCommandResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicProxyInstallCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRefreshTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的集群检查任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建检查任务的结果，"Succ"为成功，"Failed"为失败

		CreateResult *string `json:"CreateResult,omitempty" name:"CreateResult"`
		// 返回创建的新集群检查任务ID

		NewTaskID *string `json:"NewTaskID,omitempty" name:"NewTaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRefreshTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRefreshTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlChildRuleInfo struct {

	// 子策略id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 策略模式,  RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截

	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 被访问文件路径，仅仅在访问控制生效

	TargetFilePath *string `json:"TargetFilePath,omitempty" name:"TargetFilePath"`
}

type CreateDefenceVulExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateDefenceVulExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDefenceVulExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAffectedWorkloadListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受影响的workload列表数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 受影响的workload列表

		AffectedWorkloadList []*AffectedWorkloadItem `json:"AffectedWorkloadList,omitempty" name:"AffectedWorkloadList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAffectedWorkloadListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedWorkloadListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallClusterRefreshStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态，可能为：Task_Running,Task_Succ,Task_Error,Task_NoExist

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallClusterRefreshStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallClusterRefreshStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceHostAssetsSyncStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机同步状态， NORMAL，RUNNING

		SyncStatus *string `json:"SyncStatus,omitempty" name:"SyncStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceHostAssetsSyncStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceHostAssetsSyncStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 有风险的集群数量

		RiskClusterCount *uint64 `json:"RiskClusterCount,omitempty" name:"RiskClusterCount"`
		// 未检查的集群数量

		UncheckClusterCount *uint64 `json:"UncheckClusterCount,omitempty" name:"UncheckClusterCount"`
		// 托管集群数量

		ManagedClusterCount *uint64 `json:"ManagedClusterCount,omitempty" name:"ManagedClusterCount"`
		// 独立集群数量

		IndependentClusterCount *uint64 `json:"IndependentClusterCount,omitempty" name:"IndependentClusterCount"`
		// 无风险的集群数量

		NoRiskClusterCount *uint64 `json:"NoRiskClusterCount,omitempty" name:"NoRiskClusterCount"`
		// 已经检查集群数

		CheckedClusterCount *uint64 `json:"CheckedClusterCount,omitempty" name:"CheckedClusterCount"`
		// 自动检查集群数

		AutoCheckClusterCount *uint64 `json:"AutoCheckClusterCount,omitempty" name:"AutoCheckClusterCount"`
		// 手动检查集群数

		ManualCheckClusterCount *uint64 `json:"ManualCheckClusterCount,omitempty" name:"ManualCheckClusterCount"`
		// 检查失败集群数

		FailedClusterCount *uint64 `json:"FailedClusterCount,omitempty" name:"FailedClusterCount"`
		// 未导入的集群数量

		NotImportedClusterCount *uint64 `json:"NotImportedClusterCount,omitempty" name:"NotImportedClusterCount"`
		// eks集群数量

		ServerlessClusterCount *uint64 `json:"ServerlessClusterCount,omitempty" name:"ServerlessClusterCount"`
		// TKE集群数量

		TkeClusterCount *uint64 `json:"TkeClusterCount,omitempty" name:"TkeClusterCount"`
		// 用户自建腾讯云集群数量

		UserCreateTencentClusterCount *uint64 `json:"UserCreateTencentClusterCount,omitempty" name:"UserCreateTencentClusterCount"`
		// 用户自建集群混合云数量

		UserCreateHybridClusterCount *uint64 `json:"UserCreateHybridClusterCount,omitempty" name:"UserCreateHybridClusterCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulIgnoreRegistryImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像列表

		List []*VulIgnoreRegistryImage `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulIgnoreRegistryImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulIgnoreRegistryImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIgnoreVulRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID信息列表

	List []*ModifyIgnoreVul `json:"List,omitempty" name:"List"`
}

func (r *AddIgnoreVulRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddIgnoreVulRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>EventStatus- String - 是否必填：否 - 事件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_ADD_WHITE：已加白</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>RiskDns- string - 是否必填：否 - 恶意域名。</li>
	// <li>RiskIP- string - 是否必填：否 - 恶意IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：告警数量：EventCount，最近生成时间：LatestFoundTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeRiskDnsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterCheckTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的集群检查任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建检查任务的结果，"Succ"为成功，其他的为失败原因

		CreateResult *string `json:"CreateResult,omitempty" name:"CreateResult"`
		// 返回创建的集群新任务ID

		NewTaskID *string `json:"NewTaskID,omitempty" name:"NewTaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterCheckTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterCheckTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDBServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// db服务列表

		List []*ServiceInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDBServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDBServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterUninstallCmdRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *DescribeClusterUninstallCmdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterUninstallCmdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressForwardConfig struct {

	// 协议.

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听端口

	ListeningPort *uint64 `json:"ListeningPort,omitempty" name:"ListeningPort"`
	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// URL路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 后端服务

	BackendServiceName *string `json:"BackendServiceName,omitempty" name:"BackendServiceName"`
	// 服务端口

	ServicePort *uint64 `json:"ServicePort,omitempty" name:"ServicePort"`
	// 重定向目标地址

	RewriteAddress *string `json:"RewriteAddress,omitempty" name:"RewriteAddress"`
	// 1 ForwardConfig:转发配置,2 RewriteForwardConfig:重定向转发配置

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type ScanComplianceAssetsRequest struct {
	*tchttp.BaseRequest

	// 要重新扫描的客户资产项ID的列表。

	CustomerAssetIdSet []*uint64 `json:"CustomerAssetIdSet,omitempty" name:"CustomerAssetIdSet"`
}

func (r *ScanComplianceAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanComplianceAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceHostListRequest struct {
	*tchttp.BaseRequest

	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 要获取的数量，默认为10，最大为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询过滤器

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 desc asc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 已选择列表 和 未选择的列表

	Checked *bool `json:"Checked,omitempty" name:"Checked"`
}

func (r *DescribeComplianceHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegistryVulScanAuthorizedImageSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegistryVulScanAuthorizedImageSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegistryVulScanAuthorizedImageSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckNetworkFirewallPolicyYamlRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// base64编码的networkpolicy yaml字符串

	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CheckNetworkFirewallPolicyYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckNetworkFirewallPolicyYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群的详细信息

		ClusterInfoList []*ClusterInfoItem `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessControlsRuleExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessControlsRuleExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlsRuleExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessLevelSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAbnormalProcessLevelSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessLevelSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容器安全uuid

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 主机名

		HostName *string `json:"HostName,omitempty" name:"HostName"`
		// 主机分组

		Group *string `json:"Group,omitempty" name:"Group"`
		// 主机IP

		HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
		// 操作系统

		OsName *string `json:"OsName,omitempty" name:"OsName"`
		// agent版本

		AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
		// 内核版本

		KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`
		// docker版本

		DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`
		// docker api版本

		DockerAPIVersion *string `json:"DockerAPIVersion,omitempty" name:"DockerAPIVersion"`
		// docker go 版本

		DockerGoVersion *string `json:"DockerGoVersion,omitempty" name:"DockerGoVersion"`
		// docker 文件系统类型

		DockerFileSystemDriver *string `json:"DockerFileSystemDriver,omitempty" name:"DockerFileSystemDriver"`
		// docker root 目录

		DockerRootDir *string `json:"DockerRootDir,omitempty" name:"DockerRootDir"`
		// 镜像数

		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
		// 容器数

		ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
		// k8s IP

		K8sMasterIP *string `json:"K8sMasterIP,omitempty" name:"K8sMasterIP"`
		// k8s version

		K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
		// kube proxy

		KubeProxyVersion *string `json:"KubeProxyVersion,omitempty" name:"KubeProxyVersion"`
		// "UNINSTALL"："未安装","OFFLINE"："离线", "ONLINE"："防护中

		Status *string `json:"Status,omitempty" name:"Status"`
		// 是否Containerd

		IsContainerd *bool `json:"IsContainerd,omitempty" name:"IsContainerd"`
		// 主机来源;"TENCENTCLOUD":"腾讯云服务器","OTHERCLOUD":"非腾讯云服务器"

		MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
		// 外网ip

		PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
		// 主机实例ID

		InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
		// 地域ID

		RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`
		// 所属项目

		Project *ProjectInfo `json:"Project,omitempty" name:"Project"`
		// 标签

		Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
		// 集群ID

		ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 集群接入状态

		ClusterAccessedStatus *string `json:"ClusterAccessedStatus,omitempty" name:"ClusterAccessedStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetHostDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunTimeEventBaseInfo struct {

	// 事件唯一ID

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 事件发现时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 状态， “EVENT_UNDEAL”:事件未处理
	//     "EVENT_DEALED":事件已经处理
	//     "EVENT_INGNORE"：事件已经忽略

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件名称：
	// 宿主机文件访问逃逸、
	// Syscall逃逸、
	// MountNamespace逃逸、
	// 程序提权逃逸、
	// 特权容器启动逃逸、
	// 敏感路径挂载
	// 恶意进程启动
	// 文件篡改

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 事件类型
	//    ESCAPE_HOST_ACESS_FILE:宿主机文件访问逃逸
	//    ESCAPE_MOUNT_NAMESPACE:MountNamespace逃逸
	//    ESCAPE_PRIVILEDGE:程序提权逃逸
	//    ESCAPE_PRIVILEDGE_CONTAINER_START:特权容器启动逃逸
	//    ESCAPE_MOUNT_SENSITIVE_PTAH:敏感路径挂载
	//    ESCAPE_SYSCALL:Syscall逃逸

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 内网ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 外网ip

	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 节点类型:NORMAL:普通节点;SUPER:超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点子网ID

	NodeSubNetID *string `json:"NodeSubNetID,omitempty" name:"NodeSubNetID"`
	// 节点子网名称

	NodeSubNetName *string `json:"NodeSubNetName,omitempty" name:"NodeSubNetName"`
	// 节点子网网段

	NodeSubNetCIDR *string `json:"NodeSubNetCIDR,omitempty" name:"NodeSubNetCIDR"`
	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// podIP

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// pod状态

	PodStatus *string `json:"PodStatus,omitempty" name:"PodStatus"`
	// 集群id

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// uuid

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// Namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// WorkloadType

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
}

type PromotionActivityContent struct {

	// 月数

	MonthNum *uint64 `json:"MonthNum,omitempty" name:"MonthNum"`
	// 核数最低限量

	CoresCountLimit *uint64 `json:"CoresCountLimit,omitempty" name:"CoresCountLimit"`
	// 专业版折扣

	ProfessionalDiscount *uint64 `json:"ProfessionalDiscount,omitempty" name:"ProfessionalDiscount"`
	// 附赠镜像数

	ImageAuthorizationNum *uint64 `json:"ImageAuthorizationNum,omitempty" name:"ImageAuthorizationNum"`
}

type DescribeAssetContainerListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容器列表

		List []*ContainerInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetContainerListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePromotionActivityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 促销活动内容

		List []*PromotionActivityContent `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePromotionActivityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePromotionActivityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkFirewallPolicyDiscoverResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的集群检查任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建检查任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkFirewallPolicyDiscoverResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkFirewallPolicyDiscoverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageDenyRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 规则ID

	RuleIDSet []*string `json:"RuleIDSet,omitempty" name:"RuleIDSet"`
	// 启用状态 0:开启，1:关闭

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyImageDenyRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageDenyRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群Id

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 策略名

		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
		// 命名空间

		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
		// 入站类型

		FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`
		// 出站类型

		ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`
		// 自定义规则

		CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
		// pod选择器

		PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
		// 策略描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 策略创建时间

		PolicyCreateTime *string `json:"PolicyCreateTime,omitempty" name:"PolicyCreateTime"`
		// 策略源类型，分为System和Manual，分别代表手动和系统添加

		PolicySourceType *string `json:"PolicySourceType,omitempty" name:"PolicySourceType"`
		// 网络策略对应的网络插件

		NetworkPolicyPlugin *string `json:"NetworkPolicyPlugin,omitempty" name:"NetworkPolicyPlugin"`
		// 网络策略状态

		PublishStatus *string `json:"PublishStatus,omitempty" name:"PublishStatus"`
		// 网络发布结果

		PublishResult *string `json:"PublishResult,omitempty" name:"PublishResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallPolicyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NamespaceInfo struct {

	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 包含仓库数

	RegistryCnt *int64 `json:"RegistryCnt,omitempty" name:"RegistryCnt"`
	// 包含镜像数

	ImageCnt *int64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
	// 包含风险镜像数

	RiskImageCnt *int64 `json:"RiskImageCnt,omitempty" name:"RiskImageCnt"`
}

type ModifyCompliancePeriodTaskRequest struct {
	*tchttp.BaseRequest

	// 要修改的定时任务的ID，由DescribeCompliancePeriodTaskList接口返回。

	PeriodTaskId *uint64 `json:"PeriodTaskId,omitempty" name:"PeriodTaskId"`
	// 定时任务的周期规则。不填时，不修改。

	PeriodRule *CompliancePeriodTaskRule `json:"PeriodRule,omitempty" name:"PeriodRule"`
	// 设置合规标准。不填时，不修改。

	StandardSettings []*ComplianceBenchmarkStandardEnable `json:"StandardSettings,omitempty" name:"StandardSettings"`
}

func (r *ModifyCompliancePeriodTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCompliancePeriodTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 漏洞名称名称筛选，</li>
	// <li>Level - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReverseShellEventInfo struct {

	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 事件解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件详细描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 状态，EVENT_UNDEAL:事件未处理
	//     EVENT_DEALED:事件已经处理
	//     EVENT_INGNORE：事件已经忽略
	//     EVENT_ADD_WHITE：时间已经加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 父进程名

	PProcessName *string `json:"PProcessName,omitempty" name:"PProcessName"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 目标地址

	DstAddress *string `json:"DstAddress,omitempty" name:"DstAddress"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// 	"NODE_DESTROYED"      //节点已销毁
	// 	"CONTAINER_EXITED"    //容器已退出
	// 	"CONTAINER_DESTROYED" //容器已销毁
	// 	"SHARED_HOST"         // 容器与主机共享网络
	// 	"RESOURCE_LIMIT"      //隔离操作资源超限
	// 	"UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
}

type DescribeUpgradeVersionInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUpgradeVersionInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpgradeVersionInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNetworkFirewallPolicyYamlDetailRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// base64编码的networkpolicy yaml字符串

	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *AddNetworkFirewallPolicyYamlDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNetworkFirewallPolicyYamlDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyTrialRequest struct {
	*tchttp.BaseRequest
}

func (r *ApplyTrialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyTrialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// 事件ids

	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
}

func (r *DeleteReverseShellEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyRuleDetailRequest struct {
	*tchttp.BaseRequest

	// 规则RuleID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

func (r *DescribeImageDenyRuleDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessEventTendencyRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAbnormalProcessEventTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitializeUserComplianceEnvironmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitializeUserComplianceEnvironmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitializeUserComplianceEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageVulLayerInfo struct {

	// 层id

	LayerId *string `json:"LayerId,omitempty" name:"LayerId"`
	// 层cmd

	LayerCmd *string `json:"LayerCmd,omitempty" name:"LayerCmd"`
}

type DescribeSearchLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 历史搜索记录 保留最新的10条

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSearchLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIACCheckTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的任务的ID，为0表示创建失败。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 创建确认任务的结果，"Succ"为成功，"Failed"为失败

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIACCheckTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIACCheckTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnauthorizedCoresTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 未授权核数趋势

		List []*UnauthorizedCoresTendency `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnauthorizedCoresTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnauthorizedCoresTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateImageRegistryTimingScanTaskRequest struct {
	*tchttp.BaseRequest

	// 定时扫描周期

	ScanPeriod *uint64 `json:"ScanPeriod,omitempty" name:"ScanPeriod"`
	// 扫描木马类型数组

	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`
	// 扫描镜像

	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`
	// 是否扫描所有

	All *bool `json:"All,omitempty" name:"All"`
	// 定时扫描开关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 定时扫描的时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 扫描镜像Id

	Id []*uint64 `json:"Id,omitempty" name:"Id"`
	// 是否扫描最新版本

	Latest *bool `json:"Latest,omitempty" name:"Latest"`
	// 是否存在运行中的容器

	ContainerRunning *bool `json:"ContainerRunning,omitempty" name:"ContainerRunning"`
	// 扫描结束时间

	ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`
	// 扫描范围 0全部镜像，1自选镜像，2推荐扫描镜像

	ScanScope *uint64 `json:"ScanScope,omitempty" name:"ScanScope"`
	// 仓库类型 tcr,ccr,harbor

	RegistryType []*string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 命名空间

	Namespace []*string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *UpdateImageRegistryTimingScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateImageRegistryTimingScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNameListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClusterNameListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNameListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulScanTaskRequest struct {
	*tchttp.BaseRequest

	// 本地镜像扫描范围类型。ALL:全部本地镜像，NOT_SCAN：全部已授权未扫描本地镜像，IMAGEIDS:自选本地镜像ID

	LocalImageScanType *string `json:"LocalImageScanType,omitempty" name:"LocalImageScanType"`
	// 根据已授权的本地镜像IDs扫描，优先权高于根据满足条件的已授权的本地镜像。

	LocalImageIDs []*string `json:"LocalImageIDs,omitempty" name:"LocalImageIDs"`
	// 仓库镜像扫描范围类型。ALL:全部仓库镜像，NOT_SCAN：全部已授权未扫描仓库镜像，IMAGEIDS:自选仓库镜像ID

	RegistryImageScanType *string `json:"RegistryImageScanType,omitempty" name:"RegistryImageScanType"`
	// 根据已授权的仓库镜像IDs扫描，优先权高于根据满足条件的已授权的仓库镜像。

	RegistryImageIDs []*uint64 `json:"RegistryImageIDs,omitempty" name:"RegistryImageIDs"`
	// 本地镜像重新漏洞扫描时的任务ID

	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`
	// 仓库镜像重新漏洞扫描时的任务ID

	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
}

func (r *CreateVulScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceHostRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，正常：SUCCESS，异常：FAIL， NO_DEFENCE:未防御</li>
	// <li>HostName- String - 是否必填：否 - 主机名称/超级节点名称</li>
	// <li>HostIP- String - 是否必填：否 - 主机IP</li>
	// <li>NodeType- String - 是否必填：否 - 节点类型</li>
	// <li>HostName- String - 是否必填：否 - 超级节点名称</li>
	// <li>NodeSubNetCIDR- String - 是否必填：否 - 超级节点CIDR</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：更新时间：ModifyTime/首次开启时间：CreateTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulDefenceHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallPolicyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态，可能为：Task_Running,Task_Succ,Task_Error,Task_NoExist

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// NameRepeat,K8sRuleIngressPortError等

		TaskResult []*string `json:"TaskResult,omitempty" name:"TaskResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkFirewallPolicyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallNamespaceLabelListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeNetworkFirewallNamespaceLabelListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallNamespaceLabelListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWarningRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWarningRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarningRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyRuleSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageDenyRuleSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLocalStorageItemRequest struct {
	*tchttp.BaseRequest

	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 失效时间(秒)

	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`
}

func (r *SetLocalStorageItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLocalStorageItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRuleDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异常进程策略详细信息

		RuleDetail *AbnormalProcessRuleInfo `json:"RuleDetail,omitempty" name:"RuleDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessRuleDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyClusterChartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 边总数

		EdgeListTotalCount *uint64 `json:"EdgeListTotalCount,omitempty" name:"EdgeListTotalCount"`
		// 节点总数

		NodeListTotalCount *uint64 `json:"NodeListTotalCount,omitempty" name:"NodeListTotalCount"`
		// 模糊状态：precise：准确；fuzzy：模糊；

		Fuzzy *string `json:"Fuzzy,omitempty" name:"Fuzzy"`
		// 边信息

		EdgeList *string `json:"EdgeList,omitempty" name:"EdgeList"`
		// 节点信息

		NodeList *string `json:"NodeList,omitempty" name:"NodeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkTopologyClusterChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyClusterChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkFirewallPodLabelsListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeNetworkFirewallPodLabelsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkFirewallPodLabelsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddComplianceAssetPolicySetToWhitelistRequest struct {
	*tchttp.BaseRequest

	// 资产ID+检查项IDs. 列表

	AssetPolicySetList []*ComplianceAssetPolicySetItem `json:"AssetPolicySetList,omitempty" name:"AssetPolicySetList"`
}

func (r *AddComplianceAssetPolicySetToWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddComplianceAssetPolicySetToWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistrySummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像个数

		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
		// 有风险的镜像个数

		ImageHasRiskInfoCnt *uint64 `json:"ImageHasRiskInfoCnt,omitempty" name:"ImageHasRiskInfoCnt"`
		// 有病毒的镜像个数

		ImageHasVirusCnt *uint64 `json:"ImageHasVirusCnt,omitempty" name:"ImageHasVirusCnt"`
		// 有漏洞的镜像个数

		ImageHasVulsCnt *uint64 `json:"ImageHasVulsCnt,omitempty" name:"ImageHasVulsCnt"`
		// 不受信任的镜像个数

		ImageUntrustCnt *uint64 `json:"ImageUntrustCnt,omitempty" name:"ImageUntrustCnt"`
		// 最近镜像扫描时间

		LatestImageScanTime *string `json:"LatestImageScanTime,omitempty" name:"LatestImageScanTime"`
		// 风险镜像个数

		ImageUnsafeCnt *uint64 `json:"ImageUnsafeCnt,omitempty" name:"ImageUnsafeCnt"`
		// 仓库类型数量分布

		RegistryTypeStat []*ImageRegistryTypeCount `json:"RegistryTypeStat,omitempty" name:"RegistryTypeStat"`
		// 已扫描镜像个数

		ScannedImageCnt *uint64 `json:"ScannedImageCnt,omitempty" name:"ScannedImageCnt"`
		// 推荐处置镜像个数

		RecommendedFixImageCnt *uint64 `json:"RecommendedFixImageCnt,omitempty" name:"RecommendedFixImageCnt"`
		// 今日新增风险镜像个数

		TodayUnsafeImageCnt *uint64 `json:"TodayUnsafeImageCnt,omitempty" name:"TodayUnsafeImageCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistrySummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistrySummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerAssetSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容器总数

		ContainerTotalCnt *uint64 `json:"ContainerTotalCnt,omitempty" name:"ContainerTotalCnt"`
		// 正在运行容器数量

		ContainerRunningCnt *uint64 `json:"ContainerRunningCnt,omitempty" name:"ContainerRunningCnt"`
		// 暂停运行容器数量

		ContainerPauseCnt *uint64 `json:"ContainerPauseCnt,omitempty" name:"ContainerPauseCnt"`
		// 停止运行容器数量

		ContainerStopped *uint64 `json:"ContainerStopped,omitempty" name:"ContainerStopped"`
		// 本地镜像数量

		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
		// 主机节点数量

		HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`
		// 主机正在运行节点数量

		HostRunningCnt *uint64 `json:"HostRunningCnt,omitempty" name:"HostRunningCnt"`
		// 主机离线节点数量

		HostOfflineCnt *uint64 `json:"HostOfflineCnt,omitempty" name:"HostOfflineCnt"`
		// 镜像仓库数量

		ImageRegistryCnt *uint64 `json:"ImageRegistryCnt,omitempty" name:"ImageRegistryCnt"`
		// 镜像总数

		ImageTotalCnt *uint64 `json:"ImageTotalCnt,omitempty" name:"ImageTotalCnt"`
		// 主机未安装agent数量

		HostUnInstallCnt *uint64 `json:"HostUnInstallCnt,omitempty" name:"HostUnInstallCnt"`
		// 超级节点个数

		HostSuperNodeCnt *uint64 `json:"HostSuperNodeCnt,omitempty" name:"HostSuperNodeCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerAssetSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerAssetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulScanImageInfo struct {

	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像大小

	Size *float64 `json:"Size,omitempty" name:"Size"`
	// 任务状态:SCANNING:扫描中 FAILED:失败 FINISHED:完成 CANCELED:取消

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 扫描时长

	ScanDuration *float64 `json:"ScanDuration,omitempty" name:"ScanDuration"`
	// 高危漏洞数

	HighLevelVulCount *int64 `json:"HighLevelVulCount,omitempty" name:"HighLevelVulCount"`
	// 中危漏洞数

	MediumLevelVulCount *int64 `json:"MediumLevelVulCount,omitempty" name:"MediumLevelVulCount"`
	// 低危漏洞数

	LowLevelVulCount *int64 `json:"LowLevelVulCount,omitempty" name:"LowLevelVulCount"`
	// 严重漏洞数

	CriticalLevelVulCount *int64 `json:"CriticalLevelVulCount,omitempty" name:"CriticalLevelVulCount"`
	// 本地镜像漏洞扫描任务ID

	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
	// 漏洞扫描的开始时间

	ScanStartTime *string `json:"ScanStartTime,omitempty" name:"ScanStartTime"`
	// 漏洞扫描的结束时间

	ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`
	// 失败原因:TIMEOUT:超时 TOO_MANY:任务过多 OFFLINE:离线

	ErrorStatus *string `json:"ErrorStatus,omitempty" name:"ErrorStatus"`
}

type AddOrModifyMaliciousConnectionBlackListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOrModifyMaliciousConnectionBlackListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrModifyMaliciousConnectionBlackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeysLocalStorageRequest struct {
	*tchttp.BaseRequest
}

func (r *KeysLocalStorageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KeysLocalStorageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopVirusScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopVirusScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopVirusScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机id

		HostID *string `json:"HostID,omitempty" name:"HostID"`
		// 主机ip

		HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
		// 容器名称

		ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
		// 运行状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 运行账户

		RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
		// 命令行

		Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
		// CPU使用率 * 1000

		CPUUsage *uint64 `json:"CPUUsage,omitempty" name:"CPUUsage"`
		// 内存使用 KB

		RamUsage *uint64 `json:"RamUsage,omitempty" name:"RamUsage"`
		// 镜像名

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 镜像ID

		ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
		// 归属POD

		POD *string `json:"POD,omitempty" name:"POD"`
		// k8s 主节点

		K8sMaster *string `json:"K8sMaster,omitempty" name:"K8sMaster"`
		// 容器内进程数

		ProcessCnt *uint64 `json:"ProcessCnt,omitempty" name:"ProcessCnt"`
		// 容器内端口数

		PortCnt *uint64 `json:"PortCnt,omitempty" name:"PortCnt"`
		// 组件数

		ComponentCnt *uint64 `json:"ComponentCnt,omitempty" name:"ComponentCnt"`
		// app数

		AppCnt *uint64 `json:"AppCnt,omitempty" name:"AppCnt"`
		// websvc数

		WebServiceCnt *uint64 `json:"WebServiceCnt,omitempty" name:"WebServiceCnt"`
		// 挂载

		Mounts []*ContainerMount `json:"Mounts,omitempty" name:"Mounts"`
		// 容器网络信息

		Network *ContainerNetwork `json:"Network,omitempty" name:"Network"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 镜像创建时间

		ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`
		// 镜像大小

		ImageSize *uint64 `json:"ImageSize,omitempty" name:"ImageSize"`
		// 主机状态 offline,online,pause

		HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`
		// 网络状态
		// 未隔离  	NORMAL
		// 已隔离		ISOLATED
		// 隔离中		ISOLATING
		// 隔离失败	ISOLATE_FAILED
		// 解除隔离中  RESTORING
		// 解除隔离失败 RESTORE_FAILED

		NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`
		// 网络子状态

		NetSubStatus *string `json:"NetSubStatus,omitempty" name:"NetSubStatus"`
		// 隔离来源

		IsolateSource *string `json:"IsolateSource,omitempty" name:"IsolateSource"`
		// 隔离时间

		IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
		// 节点ID

		NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
		// 节点名称

		NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
		// 节点子网ID

		NodeSubNetID *string `json:"NodeSubNetID,omitempty" name:"NodeSubNetID"`
		// 节点子网名称

		NodeSubNetName *string `json:"NodeSubNetName,omitempty" name:"NodeSubNetName"`
		// 节点子网网段

		NodeSubNetCIDR *string `json:"NodeSubNetCIDR,omitempty" name:"NodeSubNetCIDR"`
		// pod名称

		PodName *string `json:"PodName,omitempty" name:"PodName"`
		// pod ip

		PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
		// pod状态

		PodStatus *string `json:"PodStatus,omitempty" name:"PodStatus"`
		// 集群ID

		ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 节点类型:NORMAL: 普通节点(默认值) SUPER: 超级节点

		NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
		// 超级节点唯一id

		NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
		// 外网ip

		PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetContainerDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAndPublishNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 策略Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 入站规则
	//
	// 全部允许：1
	//
	// 全部拒绝 ：2
	//
	// 自定义：3

	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`
	// 出站规则
	//
	// 全部允许：1
	//
	// 全部拒绝 ：2
	//
	// 自定义：3

	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`
	// pod选择器

	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 自定义规则

	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

func (r *UpdateAndPublishNetworkFirewallPolicyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAndPublishNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞列表

		List []*VulInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAssetImageRegistryRegistryDetailRequest struct {
	*tchttp.BaseRequest

	// 仓库名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 仓库url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 仓库类型，列表：harbor

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 仓库版本

	RegistryVersion *string `json:"RegistryVersion,omitempty" name:"RegistryVersion"`
	// 网络类型，列表：public（公网）

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 区域，列表：default（默认）

	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`
	// 限速

	SpeedLimit *int64 `json:"SpeedLimit,omitempty" name:"SpeedLimit"`
	// 安全模式（证书校验）：0（默认） 非安全模式（跳过证书校验）：1

	Insecure *uint64 `json:"Insecure,omitempty" name:"Insecure"`
	// 联通性检测的记录ID

	ConnDetectConfig []*ConnDetectConfig `json:"ConnDetectConfig,omitempty" name:"ConnDetectConfig"`
	// ”授权&扫描"开关

	NeedScan *bool `json:"NeedScan,omitempty" name:"NeedScan"`
}

func (r *AddAssetImageRegistryRegistryDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAssetImageRegistryRegistryDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageHostListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件 支持ImageID,HostID

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetImageHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceHostAssetsSyncStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeComplianceHostAssetsSyncStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceHostAssetsSyncStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusMonitorSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启实时监控

		EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`
		// 扫描全部路径

		ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`
		// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径

		ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`
		// 自选排除或扫描的地址

		ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
		// 扫描路径模式：
		// SCAN_PATH_ALL：全部路径
		// SCAN_PATH_DEFAULT：默认路径
		// SCAN_PATH_USER_DEFINE：用户自定义路径
		//

		ScanPathMode *string `json:"ScanPathMode,omitempty" name:"ScanPathMode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusMonitorSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusMonitorSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateK8sApiAbnormalRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则ID

		RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateK8sApiAbnormalRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateK8sApiAbnormalRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterRiskItem struct {

	// 检测项相关信息

	CheckItem *ClusterCheckItem `json:"CheckItem,omitempty" name:"CheckItem"`
	// 验证信息

	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
	// 事件描述,检查的错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	// 受影响的集群数量

	AffectedClusterCount *uint64 `json:"AffectedClusterCount,omitempty" name:"AffectedClusterCount"`
	// 受影响的节点数量

	AffectedNodeCount *uint64 `json:"AffectedNodeCount,omitempty" name:"AffectedNodeCount"`
}

type VulnerabilityFocusRuleTags struct {

	// 漏洞重点关注标签数组

	VulnerabilityFocusRuleTagValues []*string `json:"VulnerabilityFocusRuleTagValues,omitempty" name:"VulnerabilityFocusRuleTagValues"`
}

type DescribeAssetSuperNodeInfoRequest struct {
	*tchttp.BaseRequest

	// 超级节点唯一ID

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
}

func (r *DescribeAssetSuperNodeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSuperNodeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegexpRuleInfo struct {

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 启用状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 正则表达式列表

	ExpressionList []*WhiteListRegexpExpressionInfo `json:"ExpressionList,omitempty" name:"ExpressionList"`
	// 最近更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 最近操作账号

	OperatorUIN *string `json:"OperatorUIN,omitempty" name:"OperatorUIN"`
}

type CreateSuperNodeAgentInstallJobRequest struct {
	*tchttp.BaseRequest

	// 超级节点唯一id

	NodeUniqueIdSet []*string `json:"NodeUniqueIdSet,omitempty" name:"NodeUniqueIdSet"`
}

func (r *CreateSuperNodeAgentInstallJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSuperNodeAgentInstallJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSuperNodeMachineRequest struct {
	*tchttp.BaseRequest

	// 客户端唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
}

func (r *DeleteSuperNodeMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSuperNodeMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnfinishRefreshTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回最近的一次任务ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 任务状态，为Task_New,Task_Running,Task_Finish,Task_Error,Task_NoExist.Task_New,Task_Running表示有任务存在，不允许新下发

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 新任务ID

		NewTaskID *string `json:"NewTaskID,omitempty" name:"NewTaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnfinishRefreshTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnfinishRefreshTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageSimpleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像列表

		List []*AssetSimpleImageInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageSimpleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageSimpleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAffectedNodeListRequest struct {
	*tchttp.BaseRequest

	// 唯一的检测项的ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName, ClusterId,InstanceId,PrivateIpAddresses

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAffectedNodeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedNodeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellDetailRequest struct {
	*tchttp.BaseRequest

	// 事件唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeReverseShellDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHybridClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群的信息列表

		HybridClusterList []*HybridClusterItem `json:"HybridClusterList,omitempty" name:"HybridClusterList"`
		// 数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHybridClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHybridClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellEventsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulContainerListRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID</li>
	// <li>ContainerName- String -是否必填: 否 - 容器名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulContainerListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulContainerListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeRegexpWhiteListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	//
	// RuleName- String - 是否必填：否 - 规则名称

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeEscapeRegexpWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeRegexpWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEscapeRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEscapeRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEscapeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageDenyRuleExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageDenyRuleExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageDenyRuleExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscapeRule struct {

	// 规则类型
	// ESCAPE_HOST_ACESS_FILE:宿主机文件访问逃逸
	//    ESCAPE_MOUNT_NAMESPACE:MountNamespace逃逸
	//    ESCAPE_PRIVILEDGE:程序提权逃逸
	//    ESCAPE_PRIVILEDGE_CONTAINER_START:特权容器启动逃逸
	//    ESCAPE_MOUNT_SENSITIVE_PTAH:敏感路径挂载
	// ESCAPE_SYSCALL:Syscall逃逸

	Type *string `json:"Type,omitempty" name:"Type"`
	// 规则名称
	// 宿主机文件访问逃逸、
	// Syscall逃逸、
	// MountNamespace逃逸、
	// 程序提权逃逸、
	// 特权容器启动逃逸、
	// 敏感路径挂载

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否打开：false否 ，true是

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 规则组别。RISK_CONTAINER：风险容器，PROCESS_PRIVILEGE：程序特权，CONTAINER_ESCAPE：容器逃逸

	Group *string `json:"Group,omitempty" name:"Group"`
}

type ClusterWorkloadInfo struct {

	// 工作负载名称

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 工作负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 工作负载标签

	WorkloadLabels *string `json:"WorkloadLabels,omitempty" name:"WorkloadLabels"`
	// 关联的Pod数量

	PodCount *uint64 `json:"PodCount,omitempty" name:"PodCount"`
	// 工作负载Selector

	WorkloadSelectors *string `json:"WorkloadSelectors,omitempty" name:"WorkloadSelectors"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type ScanCompliancePolicyItemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回重新检测任务的ID。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanCompliancePolicyItemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanCompliancePolicyItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLocalStorageExpireRequest struct {
	*tchttp.BaseRequest

	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 过期时间

	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`
}

func (r *SetLocalStorageExpireRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLocalStorageExpireRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentDaemonSetCmdRequest struct {
	*tchttp.BaseRequest

	// 是否是腾讯云

	IsCloud *bool `json:"IsCloud,omitempty" name:"IsCloud"`
	// 网络类型：basic-基础网络，private-VPC, public-公网，direct-专线

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 地域标示, NetType=direct时必填

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// VpcId, NetType=direct时必填

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 命令有效期，非腾讯云时必填

	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
	// 集群自定义参数

	ClusterCustomParameters []*ClusterCustomParameters `json:"ClusterCustomParameters,omitempty" name:"ClusterCustomParameters"`
}

func (r *DescribeAgentDaemonSetCmdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentDaemonSetCmdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkTopologyClusterChartRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// TimeRange代表时间范围：string类型，前2个值代表开始时间和结束时间
	//

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkTopologyClusterChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkTopologyClusterChartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecLogJoinStateRequest struct {
	*tchttp.BaseRequest

	// 日志类型
	// bash日志: container_bash
	// 容器启动: container_launch
	// k8sApi: k8s_api

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 状态(true:开 false:关)

	State *bool `json:"State,omitempty" name:"State"`
}

func (r *ModifySecLogJoinStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecLogJoinStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyK8sApiAbnormalRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 状态(true:开 false:关)

	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyK8sApiAbnormalRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyK8sApiAbnormalRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateK8sApiAbnormalEventExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateK8sApiAbnormalEventExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateK8sApiAbnormalEventExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProVersionInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProVersionInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProVersionInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventEscapeImageDetailRequest struct {
	*tchttp.BaseRequest

	// 唯一值

	UniqueKey *string `json:"UniqueKey,omitempty" name:"UniqueKey"`
}

func (r *DescribeEventEscapeImageDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventEscapeImageDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessDetailBaseInfo struct {

	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程pid

	ProcessId *uint64 `json:"ProcessId,omitempty" name:"ProcessId"`
	// 进程启动用户

	ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`
	// 进程用户组

	ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 进程命令行参数

	ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
}

type DescribeComplianceTaskPolicyItemSummaryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回最近一次合规检查任务的ID。这个任务为本次所展示数据的来源。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 返回检测项的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回各检测项对应的汇总信息的列表。

		PolicyItemSummaryList []*CompliancePolicyItemSummary `json:"PolicyItemSummaryList,omitempty" name:"PolicyItemSummaryList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceTaskPolicyItemSummaryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceTaskPolicyItemSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContainerNetStatusRequest struct {
	*tchttp.BaseRequest

	// 容器ID

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
	// 状态(
	// 隔离容器: EVENT_ISOLATE_CONTAINER
	// 恢复容器: EVENT_RESOTRE_CONTAINER
	// )

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyContainerNetStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyContainerNetStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeRegexpWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 列表

		List []*RegexpRuleListItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeRegexpWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeRegexpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrModifyPostPayCoresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrModifyPostPayCoresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrModifyPostPayCoresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeSafeStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unsafe：存在风险，Safe：暂无风险,UnKnown:未知风险

		IsSafe *string `json:"IsSafe,omitempty" name:"IsSafe"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeSafeStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeSafeStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeESHitsRequest struct {
	*tchttp.BaseRequest

	// ES查询条件JSON

	Query *string `json:"Query,omitempty" name:"Query"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeESHitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeESHitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCheckModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// "Succ"表示设置成功，"Failed"表示设置失败

		SetCheckResult *string `json:"SetCheckResult,omitempty" name:"SetCheckResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCheckModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCheckModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddCompliancePolicyAssetSetToWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCompliancePolicyAssetSetToWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCompliancePolicyAssetSetToWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogDeliveryClsOptionsRequest struct {
	*tchttp.BaseRequest

	// 地域

	ClsRegion *string `json:"ClsRegion,omitempty" name:"ClsRegion"`
}

func (r *DescribeSecLogDeliveryClsOptionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogDeliveryClsOptionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessInfo struct {

	// 进程启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 运行用户

	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
	// 命令行参数

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// Exe路径

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 主机PID

	PID *uint64 `json:"PID,omitempty" name:"PID"`
	// 容器内pid

	ContainerPID *uint64 `json:"ContainerPID,omitempty" name:"ContainerPID"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// podip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 超级节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
}

type DeleteRiskSyscallWhiteListsRequest struct {
	*tchttp.BaseRequest

	// 白名单ids

	WhiteListIdSet []*string `json:"WhiteListIdSet,omitempty" name:"WhiteListIdSet"`
}

func (r *DeleteRiskSyscallWhiteListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRiskSyscallWhiteListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryRiskInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像漏洞列表

		List []*ImageRisk `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryRiskInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRiskInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
