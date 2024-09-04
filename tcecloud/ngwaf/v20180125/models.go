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

package v20180125

import (
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type RefreshAccessCheckResultRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *RefreshAccessCheckResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshAccessCheckResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddGroupProtectionRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 设置的带宽（非实时）

	BindWidth *int64 `json:"BindWidth,omitempty" name:"BindWidth"`
	// 是否迁到tke

	ToTke *int64 `json:"ToTke,omitempty" name:"ToTke"`
	// 迁移、限速、回滚等功能设置

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
}

func (r *AddGroupProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddGroupProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHostInnerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增防护域名ID

		DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHostInnerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostInnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHybridClusterRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 新的Waf开关状态，如果和已有状态相同认为修改成功

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 该动作类型， Status修改开关，InstanceId绑定实例

	OpType *string `json:"OpType,omitempty" name:"OpType"`
}

func (r *ModifyHybridClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHybridClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotListStat struct {

	// 平均速度

	AvgSpeed *float64 `json:"AvgSpeed,omitempty" name:"AvgSpeed"`
}

type PartDetail struct {

	// 计费细项

	Name *string `json:"Name,omitempty" name:"Name"`
	// 价格明细项

	Value *PartDetailItem `json:"Value,omitempty" name:"Value"`
}

type DescribeCertificateVerifyResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 错误详情

		Detail []*string `json:"Detail,omitempty" name:"Detail"`
		// 过期时间

		NotAfter *string `json:"NotAfter,omitempty" name:"NotAfter"`
		// 证书是否改变:1有改变，0没有改变

		Changed *int64 `json:"Changed,omitempty" name:"Changed"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificateVerifyResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCertificateVerifyResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserEngineTypeRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeUserEngineTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserEngineTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBusinessRiskRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBusinessRiskRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBusinessRiskRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePocTestTaskListItem struct {

	// 动作类型

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务完成进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 任务完成时间

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
	// 测试结果下载地址

	Link *string `json:"Link,omitempty" name:"Link"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type HighPrecisionPrice struct {

	// 价格

	PriceHigh *string `json:"PriceHigh,omitempty" name:"PriceHigh"`
	// 使用

	TotalCostHigh *string `json:"TotalCostHigh,omitempty" name:"TotalCostHigh"`
	// 总价

	RealTotalCostHigh *string `json:"RealTotalCostHigh,omitempty" name:"RealTotalCostHigh"`
}

type InstanceInfo struct {

	// 实例唯一ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例对应资源ID，计费使用

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 实例所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 付费模式

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 自动续费标识。
	// 0：关闭
	// 1：开启

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 弹性计费开关。
	// 0：关闭
	// 1：开启

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 实例套餐版本。
	// 101：小微版
	// 102：超轻版
	// 2：高级版
	// 3：企业版
	// 4：旗舰版
	// 6：独享版

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 实例过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 实例开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 已配置域名个数

	DomainCount *uint64 `json:"DomainCount,omitempty" name:"DomainCount"`
	// 域名数量上限

	SubDomainLimit *uint64 `json:"SubDomainLimit,omitempty" name:"SubDomainLimit"`
	// 已配置主域名个数

	MainDomainCount *uint64 `json:"MainDomainCount,omitempty" name:"MainDomainCount"`
	// 主域名数量上限

	MainDomainLimit *uint64 `json:"MainDomainLimit,omitempty" name:"MainDomainLimit"`
	// 实例30天内QPS峰值

	MaxQPS *uint64 `json:"MaxQPS,omitempty" name:"MaxQPS"`
	// qps扩展包信息

	QPS *QPSPackageNew `json:"QPS,omitempty" name:"QPS"`
	// 域名扩展包信息

	DomainPkg *DomainPackageNew `json:"DomainPkg,omitempty" name:"DomainPkg"`
	// 用户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// clb或saas

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 业务安全包

	FraudPkg *FraudPkg `json:"FraudPkg,omitempty" name:"FraudPkg"`
	// Bot资源包

	BotPkg *BotPkg `json:"BotPkg,omitempty" name:"BotPkg"`
	// bot的qps详情

	BotQPS *BotQPS `json:"BotQPS,omitempty" name:"BotQPS"`
	// qps弹性计费上限

	ElasticBilling *uint64 `json:"ElasticBilling,omitempty" name:"ElasticBilling"`
	// 攻击日志投递开关

	AttackLogPost *int64 `json:"AttackLogPost,omitempty" name:"AttackLogPost"`
	// 带宽峰值，单位为B/s(字节每秒)

	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// api安全是否购买

	APISecurity *uint64 `json:"APISecurity,omitempty" name:"APISecurity"`
	// 购买的qps规格

	QpsStandard *uint64 `json:"QpsStandard,omitempty" name:"QpsStandard"`
	// 购买的带宽规格

	BandwidthStandard *uint64 `json:"BandwidthStandard,omitempty" name:"BandwidthStandard"`
	// 实例状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 实例沙箱qps值

	SandboxQps *uint64 `json:"SandboxQps,omitempty" name:"SandboxQps"`
	// 是否api 安全试用

	IsAPISecurityTrial *uint64 `json:"IsAPISecurityTrial,omitempty" name:"IsAPISecurityTrial"`
	// 重保包

	MajorEventsPkg *MajorEventsPkg `json:"MajorEventsPkg,omitempty" name:"MajorEventsPkg"`
	// 混合云子节点包

	HybridPkg *HybridPkg `json:"HybridPkg,omitempty" name:"HybridPkg"`
	// API安全资源包

	ApiPkg *ApiPkg `json:"ApiPkg,omitempty" name:"ApiPkg"`
	// 小程序安全加速包

	MiniPkg *MiniPkg `json:"MiniPkg,omitempty" name:"MiniPkg"`
	// 小程序qps规格

	MiniQpsStandard *uint64 `json:"MiniQpsStandard,omitempty" name:"MiniQpsStandard"`
	// 小程序qps峰值

	MiniMaxQPS *uint64 `json:"MiniMaxQPS,omitempty" name:"MiniMaxQPS"`
}

type UserInfoCls struct {

	// 计数

	Count *string `json:"Count,omitempty" name:"Count"`
	// 过期时间

	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 续费标志

	Renew *string `json:"Renew,omitempty" name:"Renew"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeIpHitItemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Data *IpHitItemsData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpHitItemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpHitItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpartaProtectionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果总数

		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`
		// 结果列表

		ProtectionList []*Protection `json:"ProtectionList,omitempty" name:"ProtectionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpartaProtectionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpartaProtectionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotFirstPurchaseRequest struct {
	*tchttp.BaseRequest

	// 1表示尾部；2表示腰部，3表示头部

	UserLevel *uint64 `json:"UserLevel,omitempty" name:"UserLevel"`
}

func (r *DescribeBotFirstPurchaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotFirstPurchaseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBatchCustomRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBatchCustomRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBatchCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRewriteCountRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyRewriteCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRewriteCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotIdConfig struct {

	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	BotId *string `json:"BotId,omitempty" name:"BotId"`
	// 规则开关

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 动作配置

	Action *string `json:"Action,omitempty" name:"Action"`
}

type CreateAccessDownloadRecordRsp struct {

	// 是否成功，正常情况为0

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 下载记录编号

	Flow *string `json:"Flow,omitempty" name:"Flow"`
}

type AddBatchCustomRuleRequest struct {
	*tchttp.BaseRequest

	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 有效期
	//

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 优先级

	SortId *int64 `json:"SortId,omitempty" name:"SortId"`
	// 执行动作

	ActionType *int64 `json:"ActionType,omitempty" name:"ActionType"`
	// 重定向地址

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 加白模块

	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 事件Id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 域名列表

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
	// 策略详情列表

	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`
}

func (r *AddBatchCustomRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBatchCustomRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionedIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Data *ActionedIpData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeActionedIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeActionedIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryConstantEnumRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryConstantEnumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryConstantEnumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增防护域名ID

		DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserLevelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 300:正常 400:严格

		Level *uint64 `json:"Level,omitempty" name:"Level"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserLevelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppletAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAppletAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppletAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotLevelRequest struct {
	*tchttp.BaseRequest

	// 等级

	Level *string `json:"Level,omitempty" name:"Level"`
}

func (r *DescribeBotLevelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotLevelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotWhiteStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotWhiteStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotWhiteStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAttackWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAttackWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAttackWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAccessInfo struct {

	// 日志导出任务ID

	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
	// 日志导出查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 日志导出文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 日志文件大小

	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`
	// 日志导出时间排序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 日志导出格式

	Format *string `json:"Format,omitempty" name:"Format"`
	// 日志导出数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 日志下载状态。Processing:导出正在进行中，Complete:导出完成，Failed:导出失败，Expired:日志导出已过期（三天有效期）

	Status *string `json:"Status,omitempty" name:"Status"`
	// 日志导出起始时间

	From *int64 `json:"From,omitempty" name:"From"`
	// 日志导出结束时间

	To *int64 `json:"To,omitempty" name:"To"`
	// 日志导出路径

	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`
	// 日志导出创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type BotModuleStatusEntry struct {

	// bot -ai模块开关状态

	AiStatus *bool `json:"AiStatus,omitempty" name:"AiStatus"`
	// bot -威胁情报开关状态

	TiStatus *bool `json:"TiStatus,omitempty" name:"TiStatus"`
	// bot- 前端对抗模块开关

	JsInjectStatus *bool `json:"JsInjectStatus,omitempty" name:"JsInjectStatus"`
	// bot- 智能统计模块开关

	AutoStatisticStatus *bool `json:"AutoStatisticStatus,omitempty" name:"AutoStatisticStatus"`
}

type DescribeCustomRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则条数

		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`
		// 规则详情

		RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitempty" name:"RuleList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoticeInstanceChangeRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *NoticeInstanceChangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NoticeInstanceChangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HybridPkg struct {

	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 地域

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 申请数量

	InquireNum *int64 `json:"InquireNum,omitempty" name:"InquireNum"`
	// 使用数量

	UsedNum *int64 `json:"UsedNum,omitempty" name:"UsedNum"`
	// 续费标志

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type UpsertBotUCBFeatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正常情况下为null

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpsertBotUCBFeatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertBotUCBFeatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCLSPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已使用cls

		Used *float64 `json:"Used,omitempty" name:"Used"`
		// 资源id

		ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
		// 过期时间

		ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
		// 自动续费开关

		RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
		// cls总量

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 地区

		Region *string `json:"Region,omitempty" name:"Region"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCLSPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCLSPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpartaPackageRenewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true or false

		AutoRenew *string `json:"AutoRenew,omitempty" name:"AutoRenew"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySpartaPackageRenewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpartaPackageRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyBlockPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyBlockPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyBlockPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotSceneUAStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DescribeBotSceneUAStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSceneUAStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotGlobalUAStrategyRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 配置项列表

	Data []*UAStrategy `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyBotGlobalUAStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotGlobalUAStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopStraceIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 柱状图

		Histogram []*string `json:"Histogram,omitempty" name:"Histogram"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopStraceIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopStraceIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotUCBFeatureRuleRsp struct {

	// 规则数目

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 详细的规则

	Res []*string `json:"Res,omitempty" name:"Res"`
}

type ModifyUserLevelRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 防护规则等级 300=standard，400=extended

	Level *uint64 `json:"Level,omitempty" name:"Level"`
}

func (r *ModifyUserLevelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserLevelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotDetailCountTopNResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*CountTopNItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotDetailCountTopNResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotDetailCountTopNResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePocTestTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务列表

		TaskList []*DescribePocTestTaskListItem `json:"TaskList,omitempty" name:"TaskList"`
		// 任务条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePocTestTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePocTestTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopAttackDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Web攻击域名列表

		Web []*KVInt `json:"Web,omitempty" name:"Web"`
		// CC攻击域名列表

		CC []*KVInt `json:"CC,omitempty" name:"CC"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopAttackDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopAttackDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePostCLSFlowsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客户的投递流列表

		PostCLSFlows []*PostCLSFlowInfo `json:"PostCLSFlows,omitempty" name:"PostCLSFlows"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePostCLSFlowsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePostCLSFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotIdRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotIdRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotIdRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotWhiteStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotWhiteStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotWhiteStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ThreatIntelligenceData struct {

	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// IP数量

	IpNum *uint64 `json:"IpNum,omitempty" name:"IpNum"`
}

type QueryBypassAllStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryBypassAllStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBypassAllStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotSceneActionRule struct {

	// 动作策略ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 动作策略名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 策略优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 策略生效状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 分数范围

	Score []*BotScoreRuleEntry `json:"Score,omitempty" name:"Score"`
	// 100-宽松、200-中等、300-严格、0-自定义

	Level *string `json:"Level,omitempty" name:"Level"`
	// 生效范围，为空表示全部范围

	Scope []*BotActionScopeRuleEntry `json:"Scope,omitempty" name:"Scope"`
	// default：默认创建 custom：自定义创建

	Type *string `json:"Type,omitempty" name:"Type"`
	// 匹配范围类型：全局匹配 or 自定义匹配范围

	ScopeType *string `json:"ScopeType,omitempty" name:"ScopeType"`
	// 匹配条件间的与或关系

	ActionMatchType *string `json:"ActionMatchType,omitempty" name:"ActionMatchType"`
}

type AddCustomRuleRequest struct {
	*tchttp.BaseRequest

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 放行的详情

	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
	// 添加规则的来源，默认为空

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 优先级

	SortId *string `json:"SortId,omitempty" name:"SortId"`
	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 过期时间，单位为秒级时间戳，例如1677254399表示过期时间为2023-02-24 23:59:59. 0表示永不过期

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 策略详情

	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`
	// 需要添加策略的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 动作类型，1代表阻断，2代表人机识别，3代表观察，4代表重定向

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *AddCustomRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCustomRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessSpartaProtectionRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domains *string `json:"Domains,omitempty" name:"Domains"`
	// 实例

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 用户id

	AppIdInner *int64 `json:"AppIdInner,omitempty" name:"AppIdInner"`
	// 用户id

	UinInner *string `json:"UinInner,omitempty" name:"UinInner"`
}

func (r *AccessSpartaProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessSpartaProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCopyAreaBanRequest struct {
	*tchttp.BaseRequest

	// 规则字符串

	Rules *string `json:"Rules,omitempty" name:"Rules"`
	// 目标域名字符串，以;分隔

	Target *string `json:"Target,omitempty" name:"Target"`
	// 来源域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *CreateCopyAreaBanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCopyAreaBanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDomainListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名列表

		Data *DescribeInstanceDomainListData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceDomainListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppletAccessV3Request struct {
	*tchttp.BaseRequest

	// 支持的过滤器：AppletID: 小程序ID

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 上限

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAppletAccessV3Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppletAccessV3Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreaBanAreasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回包内容

		Data *DescribeAreaBanAreasRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAreaBanAreasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreaBanAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainWhiteRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		RuleList []*RuleList `json:"RuleList,omitempty" name:"RuleList"`
		// 规则的数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainWhiteRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainWhiteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttackDownloadRecordRequest struct {
	*tchttp.BaseRequest

	// 下载任务记录唯一标记

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteAttackDownloadRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttackDownloadRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainPackage struct {

	// 资源ID

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 是否自动续费，1：自动续费，0：不自动续费

	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
	// 套餐购买个数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 套餐购买地域，clb-waf暂时没有用到

	Region *string `json:"Region,omitempty" name:"Region"`
}

type CreateCachePathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCachePathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCachePathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHybridClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息

		Clusters []*HybridCluster `json:"Clusters,omitempty" name:"Clusters"`
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

type EditOpRuleUpdateLogStatusRequest struct {
	*tchttp.BaseRequest

	// 要编辑的日志id

	UpdateLogIDs []*uint64 `json:"UpdateLogIDs,omitempty" name:"UpdateLogIDs"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *EditOpRuleUpdateLogStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditOpRuleUpdateLogStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccessDownloadRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载访问日志记录数组

		Records []*DownloadAccessRecordInfo `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAccessDownloadRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccessDownloadRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpsertCCRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一般为null

		Data *string `json:"Data,omitempty" name:"Data"`
		// 操作的RuleId

		RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpsertCCRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertCCRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FreshAntiFakeUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果成功与否

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FreshAntiFakeUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FreshAntiFakeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainBotStatusData struct {

	// 数据

	Res []*DomainBotStatus `json:"Res,omitempty" name:"Res"`
}

type AddDomainWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 需要添加的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 需要添加的规则

	Rules []*uint64 `json:"Rules,omitempty" name:"Rules"`
	// 需要添加的规则url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 规则的方法

	Function *string `json:"Function,omitempty" name:"Function"`
	// 规则的开关，0表示规则关闭，1表示规则打开

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *AddDomainWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDomainWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRolesRequest struct {
	*tchttp.BaseRequest

	// 目前支持两种方式，分别是WAF_QCSLinkedRoleInCLS 和 WAF_QCSLinkedRoleInCKafkackafka

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DescribeRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDefaultEngineRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserDefaultEngineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDefaultEngineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRateLimitsRequest struct {
	*tchttp.BaseRequest

	// 规则ID列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DeleteRateLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRateLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功的状态码，需要JSON解码后再使用，返回的格式是{"域名":"状态"}，成功的状态码为Success，其它的为失败的状态码（yunapi定义的错误码）

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CCRuleItems struct {

	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 模式

	Advance *uint64 `json:"Advance,omitempty" name:"Advance"`
	// 限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 范围

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 网址

	Url *string `json:"Url,omitempty" name:"Url"`
	// 匹配类型

	MatchFunc *uint64 `json:"MatchFunc,omitempty" name:"MatchFunc"`
	// 动作

	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`
	// 优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
	// 有效时间

	ValidTime *uint64 `json:"ValidTime,omitempty" name:"ValidTime"`
	// 版本

	TsVersion *uint64 `json:"TsVersion,omitempty" name:"TsVersion"`
	// 规则详情

	Options *string `json:"Options,omitempty" name:"Options"`
	// 规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 关联的Session规则

	SessionApplied []*int64 `json:"SessionApplied,omitempty" name:"SessionApplied"`
}

type CreateAppletAccessRequest struct {
	*tchttp.BaseRequest

	// 小程序接入名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 小程序类型传入 1， 预留扩展用

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 小程序源站域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 服务协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 微信小程序ID

	AppletID *string `json:"AppletID,omitempty" name:"AppletID"`
	// 详情

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateAppletAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppletAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQPSRequest struct {
	*tchttp.BaseRequest

	// 客户的版本，clb-waf,sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeUserQPSRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserQPSRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalSceneInfo struct {

	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 场景名称

	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`
	// 场景优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 场景更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ModifyPocTestAuthorizationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPocTestAuthorizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPocTestAuthorizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Protection struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名Id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// cname地址

	Cname *string `json:"Cname,omitempty" name:"Cname"`
	// waf开关 0表示关闭, 1表示开启

	Status *string `json:"Status,omitempty" name:"Status"`
	// 防护状况：0 正常防护   1 未检测到流量  2 即将到期  3 已到期

	State *string `json:"State,omitempty" name:"State"`
	// 域名创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 防御模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// AI防御模式

	Engine *string `json:"Engine,omitempty" name:"Engine"`
	// 接入IP列表

	Vip []*string `json:"Vip,omitempty" name:"Vip"`
	// 灰度标识，1灰度0不灰度

	IsGray *string `json:"IsGray,omitempty" name:"IsGray"`
	// 灰度区域，不灰传no

	GrayAreas []*string `json:"GrayAreas,omitempty" name:"GrayAreas"`
	// 证书类型，1表示自有证书，2表示腾讯云托管证书

	CertType *string `json:"CertType,omitempty" name:"CertType"`
	// 证书内容

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// 证书私钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 腾讯云托管证书id，当certType=2时有效

	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`
	// 1有cdn，0无cdn

	IsCdn *string `json:"IsCdn,omitempty" name:"IsCdn"`
	// 1开启http2，0不开启

	IsHttp2 *string `json:"IsHttp2,omitempty" name:"IsHttp2"`
	// 1开启websocket，0不开启

	IsWebsocket *string `json:"IsWebsocket,omitempty" name:"IsWebsocket"`
	// 是否http强制跳转https，"1" 是、"0" 否

	HttpsRewrite *string `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`
	// https回源端口

	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`
	// 回源方式，"0"为ip回源，"1"为域名回源

	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// 回源域名

	UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`
	// 回源协议

	UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`
	// 日志包

	Cls *string `json:"Cls,omitempty" name:"Cls"`
	// CC列表

	CCList []*string `json:"CCList,omitempty" name:"CCList"`
	// 负载均衡算法，0表示轮询，1表示ip hash，默认为0

	LoadBalance *string `json:"LoadBalance,omitempty" name:"LoadBalance"`
	// 服务端口配置

	Ports []*PortItem `json:"Ports,omitempty" name:"Ports"`
	// 回源ip

	RsList []*string `json:"RsList,omitempty" name:"RsList"`
	// ip列表

	SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`
	// 是否支持长连接

	IsKeepAlive *string `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`
}

type AddBypassAllRuleRequest struct {
	*tchttp.BaseRequest
}

func (r *AddBypassAllRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBypassAllRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotTokenRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafAutoDenyRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重保护网域名状态

		HWState *int64 `json:"HWState,omitempty" name:"HWState"`
		// 攻击次数阈值

		AttackThreshold *int64 `json:"AttackThreshold,omitempty" name:"AttackThreshold"`
		// 攻击时间阈值

		TimeThreshold *int64 `json:"TimeThreshold,omitempty" name:"TimeThreshold"`
		// 自动封禁时间

		DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitempty" name:"DenyTimeThreshold"`
		// 自动封禁状态

		DefenseStatus *int64 `json:"DefenseStatus,omitempty" name:"DefenseStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafAutoDenyRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafAutoDenyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UAStrategyDetail struct {

	// UA策略项对应的具体类别，UAName:UATyoe = 1：n

	UAType []*string `json:"UAType,omitempty" name:"UAType"`
	// 中文名称

	NameCN *string `json:"NameCN,omitempty" name:"NameCN"`
	// 英文名称

	NameEN *string `json:"NameEN,omitempty" name:"NameEN"`
	// 中文描述

	DescriptionCN *string `json:"DescriptionCN,omitempty" name:"DescriptionCN"`
	// 英文描述

	DescriptionEN *string `json:"DescriptionEN,omitempty" name:"DescriptionEN"`
	// 开关状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 最近修改时间

	TimeStamp *uint64 `json:"TimeStamp,omitempty" name:"TimeStamp"`
}

type AiListFilter struct {

	// 筛选字段

	Key *string `json:"Key,omitempty" name:"Key"`
	// 筛选值：|分隔，or关系

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GlobalWhiteRule struct {

	// 序号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则开关

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 特征序号

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 详细规则

	Rules []*GlobalWhiteCond `json:"Rules,omitempty" name:"Rules"`
}

type SearchLogsSimpleRsp struct {

	// 是否成功执行，0表示成功执行

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 日志的数目

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 日志的内容，可以用作游标

	Context *string `json:"Context,omitempty" name:"Context"`
	// 具体的日志内容

	Data []*string `json:"Data,omitempty" name:"Data"`
}

type ModifyBotSessionKeyRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// session位置

	Category *string `json:"Category,omitempty" name:"Category"`
	// session key

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *ModifyBotSessionKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSessionKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDnsRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// waf的实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 取值1或者2;  1将域名切回WAF  2 切到源站,切到源站是Value不能为空

	Option *uint64 `json:"Option,omitempty" name:"Option"`
	// Value 为Ipv4格式的源站 ip地址的列表,独享版本支持5个源站,填写必须与WAF中添加的源站地址一致

	Value []*string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyDnsRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDnsRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpRuleUpdateLogRequest struct {
	*tchttp.BaseRequest

	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 语言cn/en

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeOpRuleUpdateLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpRuleUpdateLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpRuleUpdateLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则日志总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则日志列表

		List []*RuleUpdateLog `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpRuleUpdateLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpRuleUpdateLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTsRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTsRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTsRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpAccessControlData struct {

	// ip黑白名单

	Res []*IpAccessControlItem `json:"Res,omitempty" name:"Res"`
	// 计数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type PostCKafkaFlowInfo struct {

	// 投递流唯一ID

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
	// 1-访问日志 2-攻击日志

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 最后一次修改时间

	ModifyTime *int64 `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 状态 0-为关闭 1-为启用

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CKafka所在区域

	CKafkaRegion *string `json:"CKafkaRegion,omitempty" name:"CKafkaRegion"`
	// CKafka实例ID

	CKafkaID *string `json:"CKafkaID,omitempty" name:"CKafkaID"`
	// ckafka地址信息

	Brokers *string `json:"Brokers,omitempty" name:"Brokers"`
	// ckafka版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// 主题名称

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 压缩算法，支持gzip 和 lz4

	Compression *string `json:"Compression,omitempty" name:"Compression"`
	// 是否支持SASL,0-关闭，1-开启

	SASLEnable *int64 `json:"SASLEnable,omitempty" name:"SASLEnable"`
	// SASL用户名

	SASLUser *string `json:"SASLUser,omitempty" name:"SASLUser"`
	// SALS密码

	SASLPassword *string `json:"SASLPassword,omitempty" name:"SASLPassword"`
	// 0-异步投递，1-同步投递，目前仅支持同步投递

	Sync *int64 `json:"Sync,omitempty" name:"Sync"`
	// 描述信息

	Content *string `json:"Content,omitempty" name:"Content"`
	// 1-外网TGW，2-支撑环境，默认为支撑环境

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
	// 投递的域名或实例个数

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DescribeDNSDetectDomainListRequest struct {
	*tchttp.BaseRequest

	// 查询的参数，只支持域名

	Search *string `json:"Search,omitempty" name:"Search"`
	// 分页参数

	PageInfo *PageInfoForInt `json:"PageInfo,omitempty" name:"PageInfo"`
}

func (r *DescribeDNSDetectDomainListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSDetectDomainListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpMainClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主类的详细信息

		MainClass []*MainClassItem `json:"MainClass,omitempty" name:"MainClass"`
		// 主类的总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpMainClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpMainClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnData struct {

	// "returnData":{"data":{"AutoRenew":"false"}}

	Data *Data `json:"Data,omitempty" name:"Data"`
}

type DescribeAttackIndexRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAttackIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRulesInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRulesInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableUserSignaturePolicyRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 地域信息

	Areas []*string `json:"Areas,omitempty" name:"Areas"`
}

func (r *EnableUserSignaturePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableUserSignaturePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainConfig struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 大字段配置

	WriteConfig *FieldWriteConfig `json:"WriteConfig,omitempty" name:"WriteConfig"`
}

type DescribeRequestCountRequest struct {
	*tchttp.BaseRequest

	// 查询结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 客户的Appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 被查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// WAF实例ID，不传则不过滤

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 查询开始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
}

func (r *DescribeRequestCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRequestCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppletAccessV3Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 小程序接入列表

		Applets []*AppletAccessV3 `json:"Applets,omitempty" name:"Applets"`
		// 满足过滤条件的接入小程序的总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppletAccessV3Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppletAccessV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackWorldMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 每个城市和攻击数量的数组

		Map []*MapItem `json:"Map,omitempty" name:"Map"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackWorldMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackWorldMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBusinessRiskRequest struct {
	*tchttp.BaseRequest

	// 拦截Url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 账户类型

	AccountTypes *string `json:"AccountTypes,omitempty" name:"AccountTypes"`
	// 账户获取方式

	Options *string `json:"Options,omitempty" name:"Options"`
	// 条件匹配方式

	Extend *string `json:"Extend,omitempty" name:"Extend"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 域分类字段，业务安全默认为：BUSINESS_RISK

	DomainCode *string `json:"DomainCode,omitempty" name:"DomainCode"`
	// 域名，实际Host

	Host *string `json:"Host,omitempty" name:"Host"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 风险等级，天御风险类型，包含：reject（高风险），review（人工审核）pass（无风险）

	Level *string `json:"Level,omitempty" name:"Level"`
	// 风险类型：
	// "1", "账号信用低"
	// "11", "低活跃账号"
	// "2", "垃圾账号"
	// "21", "疑似小号"
	// "22", "违规账号"
	// "3", "无效账号"
	// "4", "黑名单"
	// "5", "白名单"
	// "201", "环境异常"
	// "2011", "非常用IP"
	// "2012", "IP异常"
	// "205", "非公网有效ip"
	// "206", "设备异常"
	// "2061", "非常用设备"
	// "2063", "群控设备"

	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
	// 重定向URL

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 规则动作，50（观察）51（验证码）52（拦截）53（重定向）优先级：观察>验证码>重定向>拦截

	RuleAction *int64 `json:"RuleAction,omitempty" name:"RuleAction"`
}

func (r *CreateBusinessRiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBusinessRiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOpRuleUpdateLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOpRuleUpdateLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOpRuleUpdateLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailedInfo struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 失败信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type ModifyHostStatusInnerRequest struct {
	*tchttp.BaseRequest

	// 域名状态列表

	HostsStatus []*HostStatus `json:"HostsStatus,omitempty" name:"HostsStatus"`
}

func (r *ModifyHostStatusInnerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostStatusInnerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNewOldConfigNew struct {

	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 二级产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费策略id

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// waf实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// QPS数量

	ElasticQps *int64 `json:"ElasticQps,omitempty" name:"ElasticQps"`
	// 弹性账单

	FlexBill *int64 `json:"FlexBill,omitempty" name:"FlexBill"`
	// 1:自动续费，0:不自动续费

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// waf购买的实际地域信息

	RealRegion *int64 `json:"RealRegion,omitempty" name:"RealRegion"`
	// Waf实例对应的二级产品码

	Type *string `json:"Type,omitempty" name:"Type"`
	// 计费细项标签数组

	LabelTypes []*string `json:"LabelTypes,omitempty" name:"LabelTypes"`
	// 计费细项标签数量，一般和SvLabelType一一对应

	LabelCounts []*int64 `json:"LabelCounts,omitempty" name:"LabelCounts"`
	// 变配使用，实例到期时间

	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
	// 对存在的实例购买bot 或api 安全

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// sv_wsm_waf_package_micro计费项

	WafPackageMicroClb *int64 `json:"WafPackageMicroClb,omitempty" name:"WafPackageMicroClb"`
	// sv_wsm_waf_lightpackage_lightpackage

	WafPackageClb *int64 `json:"WafPackageClb,omitempty" name:"WafPackageClb"`
	// sv_wsm_waf_package_premium_clb

	WafPackagePremiumClb *int64 `json:"WafPackagePremiumClb,omitempty" name:"WafPackagePremiumClb"`
	// sv_wsm_waf_package_enterprise_clb

	WafPackageEnterpriseClb *int64 `json:"WafPackageEnterpriseClb,omitempty" name:"WafPackageEnterpriseClb"`
	// sv_wsm_waf_package_premium

	WafPackagePremium *int64 `json:"WafPackagePremium,omitempty" name:"WafPackagePremium"`
	// sv_wsm_waf_package_enterprise

	WafPackageEnterprise *int64 `json:"WafPackageEnterprise,omitempty" name:"WafPackageEnterprise"`
	// sv_wsm_waf_package_ultimate

	WafPackageUltimate *int64 `json:"WafPackageUltimate,omitempty" name:"WafPackageUltimate"`
	// sv_wsm_waf_scene_bot_protection

	WafSceneBotProtection *int64 `json:"WafSceneBotProtection,omitempty" name:"WafSceneBotProtection"`
	// sv_wsm_waf_scene_bot_protection_clb

	WafSceneBotProtectionClb *int64 `json:"WafSceneBotProtectionClb,omitempty" name:"WafSceneBotProtectionClb"`
	// sv_wsm_waf_domain

	WafDomain *int64 `json:"WafDomain,omitempty" name:"WafDomain"`
	// sv_wsm_waf_qps_ep

	WafQps *int64 `json:"WafQps,omitempty" name:"WafQps"`
	// sv_wsm_waf_scls

	WafCls *int64 `json:"WafCls,omitempty" name:"WafCls"`
	// sv_wsm_waf_package_ultimate_clb

	WafPackageUltimateClb *string `json:"WafPackageUltimateClb,omitempty" name:"WafPackageUltimateClb"`
}

type ModifyInstanceElasticModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceElasticModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceElasticModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UCBEntryValue struct {

	// string类型值

	BasicValue *string `json:"BasicValue,omitempty" name:"BasicValue"`
	// 布尔类型值

	LogicValue *bool `json:"LogicValue,omitempty" name:"LogicValue"`
	// string数组类型值

	BelongValue []*string `json:"BelongValue,omitempty" name:"BelongValue"`
	// 指示有效的字段

	ValidKey *string `json:"ValidKey,omitempty" name:"ValidKey"`
	// string数组类型值

	MultiValue []*string `json:"MultiValue,omitempty" name:"MultiValue"`
}

type DescribeBotStatisticPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// bot数目

		PointsBot []*uint64 `json:"PointsBot,omitempty" name:"PointsBot"`
		// 总数目

		PointsTotal []*uint64 `json:"PointsTotal,omitempty" name:"PointsTotal"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotStatisticPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotStatisticPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 业务安全风险类型

		RiskInfo *RiskInfo `json:"RiskInfo,omitempty" name:"RiskInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchApiRulesRequest struct {
	*tchttp.BaseRequest

	// 启动或停用api规则的列表

	ApiList []*SwitchApiRules `json:"ApiList,omitempty" name:"ApiList"`
	// 启用或停用用户api规则的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 启用或停用api安全的状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SwitchApiRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchApiRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiInfoLeakRulesRuleItem struct {

	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 规则动作类型

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 规则创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 详细的规则

	Strategies []*DescribeAntiInfoLeakRulesStrategyItem `json:"Strategies,omitempty" name:"Strategies"`
}

type AddAntiInfoLeakRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则ID

		RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAntiInfoLeakRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAntiInfoLeakRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSceneInfoRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeUserSceneInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserSceneInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchIpAccessControlRequest struct {
	*tchttp.BaseRequest

	// 偏移

	OffSet *uint64 `json:"OffSet,omitempty" name:"OffSet"`
	// 限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 筛选条件，支持 ActionType 40/42，Ip：ip地址，Domain：域名三类

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBatchIpAccessControlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBatchIpAccessControlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PostAttackDownloadTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务task id

		Flow *string `json:"Flow,omitempty" name:"Flow"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PostAttackDownloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PostAttackDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTsResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTsResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTsResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchAreaBanListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段名称

	By *string `json:"By,omitempty" name:"By"`
	// 排序类型

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询字段信息

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBatchAreaBanListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBatchAreaBanListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSpartaProtectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSpartaProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHWDomainStateData struct {

	// HWState（1-未购买、2-购买未激活、3-购买已激活、4-已过期）
	//

	HWState *int64 `json:"HWState,omitempty" name:"HWState"`
}

type IpAccessControlItem struct {

	// mongo表自增Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 动作

	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 更新时间戳

	TsVersion *uint64 `json:"TsVersion,omitempty" name:"TsVersion"`
	// 有效截止时间戳

	ValidTs *uint64 `json:"ValidTs,omitempty" name:"ValidTs"`
	// 生效状态

	ValidStatus *int64 `json:"ValidStatus,omitempty" name:"ValidStatus"`
}

type BotCookieAnalyze struct {

	// 出现最多的cookie

	CookieMax *string `json:"CookieMax,omitempty" name:"CookieMax"`
	// cookie重复比

	CookieRepeat *float64 `json:"CookieRepeat,omitempty" name:"CookieRepeat"`
	// cookie种类

	CookieKind *uint64 `json:"CookieKind,omitempty" name:"CookieKind"`
	// cookie是否滥用

	CookieAbuse *bool `json:"CookieAbuse,omitempty" name:"CookieAbuse"`
	// cookie存在性

	CookieExist *bool `json:"CookieExist,omitempty" name:"CookieExist"`
	// cookie存在比

	CookieValidRate *float64 `json:"CookieValidRate,omitempty" name:"CookieValidRate"`
}

type ModifyAccessPeriodResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessPeriodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotAccessItemTwo struct {

	// 访问ip

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 访问会话ID

	BotToken *string `json:"BotToken,omitempty" name:"BotToken"`
	// 访问域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 访问url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 请求方式

	Method *string `json:"Method,omitempty" name:"Method"`
	// Bot标签

	BotLabel *string `json:"BotLabel,omitempty" name:"BotLabel"`
	// Bot动作

	BotAction *string `json:"BotAction,omitempty" name:"BotAction"`
	// 命中模块

	BotModule *string `json:"BotModule,omitempty" name:"BotModule"`
	// bot得分

	BotScore *int64 `json:"BotScore,omitempty" name:"BotScore"`
	// ip所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 命中时间

	BotTime *int64 `json:"BotTime,omitempty" name:"BotTime"`
	// 是否命中威胁情报

	BotTi *int64 `json:"BotTi,omitempty" name:"BotTi"`
	// 是否命中Ai评估

	BotAi *int64 `json:"BotAi,omitempty" name:"BotAi"`
	// 是否命中智能统计

	BotStat *int64 `json:"BotStat,omitempty" name:"BotStat"`
	// 命中威胁情报的IDC信息

	BotTiIdc *string `json:"BotTiIdc,omitempty" name:"BotTiIdc"`
	// 威胁情报库分类

	BotTiCategory *string `json:"BotTiCategory,omitempty" name:"BotTiCategory"`
	// 威胁情报UserAgent分类

	BotTiUa *string `json:"BotTiUa,omitempty" name:"BotTiUa"`
	// 日志条数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 是否命中智能统计平均速度异常

	BotStatAvgSpeed *int64 `json:"BotStatAvgSpeed,omitempty" name:"BotStatAvgSpeed"`
	// 是否命中智能统计的数量异常

	BotStatNums *int64 `json:"BotStatNums,omitempty" name:"BotStatNums"`
	// 是否命中智能统计的时间会话异常

	BotStatSessionDuration *int64 `json:"BotStatSessionDuration,omitempty" name:"BotStatSessionDuration"`
	// 是否命中智能统计的url种类异常

	BotStatReqUrlKind *int64 `json:"BotStatReqUrlKind,omitempty" name:"BotStatReqUrlKind"`
	// 是否命中智能统计的ua种类异常

	BotStatUaKindNums *int64 `json:"BotStatUaKindNums,omitempty" name:"BotStatUaKindNums"`
	// 是否命中智能统计的cookie种类异常

	BotStatCookieKindNums *int64 `json:"BotStatCookieKindNums,omitempty" name:"BotStatCookieKindNums"`
	// 是否命中智能统计的refer种类异常

	BotStatReferKindNums *int64 `json:"BotStatReferKindNums,omitempty" name:"BotStatReferKindNums"`
	// 规则id

	BotRuleId *string `json:"BotRuleId,omitempty" name:"BotRuleId"`
	// 场景ID

	BotSceneId *string `json:"BotSceneId,omitempty" name:"BotSceneId"`
	// 场景名称

	BotSceneName *string `json:"BotSceneName,omitempty" name:"BotSceneName"`
	// 场景动作策略ID

	BotSceneActionId *string `json:"BotSceneActionId,omitempty" name:"BotSceneActionId"`
	// 场景动作策略名称

	BotSceneActionName *string `json:"BotSceneActionName,omitempty" name:"BotSceneActionName"`
	// 是否命中UA策略模块

	BotUa *int64 `json:"BotUa,omitempty" name:"BotUa"`
}

type DescribeVIPRSResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名

		Domains []*string `json:"Domains,omitempty" name:"Domains"`
		// vip类型

		Type *string `json:"Type,omitempty" name:"Type"`
		// 源集群

		SrcGroupIds []*int64 `json:"SrcGroupIds,omitempty" name:"SrcGroupIds"`
		// 源rsip

		SrcRSIPs []*string `json:"SrcRSIPs,omitempty" name:"SrcRSIPs"`
		// 集群信息

		OptionalGroup []*OptionalGroup `json:"OptionalGroup,omitempty" name:"OptionalGroup"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVIPRSResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVIPRSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiSecEventChangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiSecEventChangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiSecEventChangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiLeakageItem struct {

	// 规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 状态值

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 匹配条件

	Strategies []*DescribeAntiInfoLeakRulesStrategyItem `json:"Strategies,omitempty" name:"Strategies"`
	// 匹配的URL

	Uri *string `json:"Uri,omitempty" name:"Uri"`
}

type DescribeTrafficMarkingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// HTTP头名称

		HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
		// HTTP头值

		HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`
		// Bot防护标记外层开关

		BotStatus *int64 `json:"BotStatus,omitempty" name:"BotStatus"`
		// Bot防护标记：bot得分开关

		BotScoreStatus *int64 `json:"BotScoreStatus,omitempty" name:"BotScoreStatus"`
		// Bot防护标记：客户端ID开关

		JsInjectStatus *int64 `json:"JsInjectStatus,omitempty" name:"JsInjectStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrafficMarkingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficMarkingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiStatPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势图参数

		Data []*ApiStatPointItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiStatPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiStatPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePocTestTasksRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页限制大小

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序 desc/asc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribePocTestTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePocTestTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessDownloadRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccessDownloadRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVIPConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录

		RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVIPConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVIPConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessHistogramItem struct {

	// 时间，单位ms

	BTime *int64 `json:"BTime,omitempty" name:"BTime"`
	// 日志条数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 时间，单位ms

	BeginTime *int64 `json:"BeginTime,omitempty" name:"BeginTime"`
}

type WafCreateResourceAfterPayRequest struct {
	*tchttp.BaseRequest

	// 区域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 购买模式

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// pid

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 接口名

	InterfaceName *string `json:"InterfaceName,omitempty" name:"InterfaceName"`
	// 类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 项目id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// goodsnum

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 实例名称111

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 客户等级：2-高级版 ，3-企业版，4-旗舰版，5-独享版

	Level *int64 `json:"Level,omitempty" name:"Level"`
}

func (r *WafCreateResourceAfterPayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafCreateResourceAfterPayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogSort struct {

	// 要排序的对象

	Key *string `json:"Key,omitempty" name:"Key"`
	// 排序方式

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeAttackDownloadRecordsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAttackDownloadRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackDownloadRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMainClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主类的总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 主类的详细信息

		MainClass []*MainClassItem `json:"MainClass,omitempty" name:"MainClass"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMainClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMainClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBatchAreaBanStatusRequest struct {
	*tchttp.BaseRequest

	// 规则Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 开关状态 0-关、1-开
	//

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyBatchAreaBanStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBatchAreaBanStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalancerInner struct {

	// 负载均衡LD的ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡LD的名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 负载均衡监听器的ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 负载均衡监听器的名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 负载均衡实例的IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 负载均衡实例的端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 负载均衡LD的地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 监听器协议，http、https

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 负载均衡监听器所在的zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 负载均衡的VPCID，公网为-1，内网按实际填写

	NumericalVpcId *int64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`
	// 负载均衡的网络类型

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
}

type MatchArgument struct {

	// 参数类型，枚举，支持：CUSTOM，METHOD，HEADER，QUERY，CALLER_SERVICE，CALLER_IP

	Type *string `json:"Type,omitempty" name:"Type"`
	// 参数键，对于HEADER、QUERY、CUSTOM，对应的是key值；对于CALLER_SERVICE，对应的是服务的命名空间值

	Key *string `json:"Key,omitempty" name:"Key"`
	// 参数值，对于HEADER、QUERY、CUSTOM，对应的是key所关联的value；对于CALLER_SERVICE，对应的是服务名，其他类型则是具体的值

	Value *MatchString `json:"Value,omitempty" name:"Value"`
}

type CreateAccessDownloadRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建下载任务返回值

		Data *CreateAccessDownloadRecordRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessDownloadRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyBotUCBFeatureRulesRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则名称

	Rules *string `json:"Rules,omitempty" name:"Rules"`
	// 目标域名

	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *CopyBotUCBFeatureRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBotUCBFeatureRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHybridClusterRequest struct {
	*tchttp.BaseRequest

	// 实例ID，可以为空。但是InstanceId和ClusterId不能同时为空

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 集群ID，可以为空。但是InstanceId和ClusterId不能同时为空

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeHybridClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHybridClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableClientMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：成功，1：失败

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableClientMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableClientMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击数量

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppletAccessV3 struct {

	// 小程序接入名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 微信小程序ID

	AppletID *string `json:"AppletID,omitempty" name:"AppletID"`
	// 详情

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 微信网关接入节点域名

	AccessDomain *string `json:"AccessDomain,omitempty" name:"AccessDomain"`
	// 授权接入阶段的状态,
	// 0 初始状态
	// 100 完成

	AuthAccessStatus *uint64 `json:"AuthAccessStatus,omitempty" name:"AuthAccessStatus"`
	// 客户体验阶段的状态
	// 0 初始状态
	// 100 完成

	ClientTrialStatus *uint64 `json:"ClientTrialStatus,omitempty" name:"ClientTrialStatus"`
	// 客户发布阶段的状态
	// 0 初始状态
	// 100 完成

	ClientPublishStatus *uint64 `json:"ClientPublishStatus,omitempty" name:"ClientPublishStatus"`
	// 小程序接入的生效版本

	EffectiveVersion *string `json:"EffectiveVersion,omitempty" name:"EffectiveVersion"`
	// 小程序业务域名信息

	Domains []*AppletAccessDomainV3 `json:"Domains,omitempty" name:"Domains"`
	// 网关ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 接入节点ID

	AccessID *string `json:"AccessID,omitempty" name:"AccessID"`
}

type DownloadRecordItem struct {

	// 任务名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务ID

	Flow *string `json:"Flow,omitempty" name:"Flow"`
	// 任务对应域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 任务涉及的记录条数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 任务运行状态，0运行中，1完成，2失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 任务结果的下载地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// 数据库自增ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type HostRecord struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 主域名，入参时为空

	MainDomain *string `json:"MainDomain,omitempty" name:"MainDomain"`
	// 规则引擎防护模式，0 观察模式，1拦截模式

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// waf和LD的绑定，0：没有绑定，1：绑定

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 域名状态，0：正常，1：未检测到流量，2：即将过期，3：过期

	State *uint64 `json:"State,omitempty" name:"State"`
	// 规则引擎和AI引擎防护模式联合状态,10规则引擎观察&&AI引擎关闭模式 11规则引擎观察&&AI引擎观察模式 12规则引擎观察&&AI引擎拦截模式 20规则引擎拦截&&AI引擎关闭模式 21规则引擎拦截&&AI引擎观察模式 22规则引擎拦截&&AI引擎拦截模式

	Engine *uint64 `json:"Engine,omitempty" name:"Engine"`
	// 是否开启代理，0：不开启，1：开启

	IsCdn *uint64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// 绑定的LB列表

	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
	// 域名绑定的LB的地域，以,分割多个地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 产品分类，取值为：sparta-waf、clb-waf、cdn-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// WAF的流量模式，1：清洗模式，0：镜像模式

	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`
	// 是否开启访问日志，1：开启，0：关闭

	ClsStatus *uint64 `json:"ClsStatus,omitempty" name:"ClsStatus"`
	// 防护等级，可选值100,200,300

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 域名需要下发到的cdc集群列表

	CdcClusters []*string `json:"CdcClusters,omitempty" name:"CdcClusters"`
	// 应用型负载均衡类型: clb或者apisix，默认clb

	AlbType *string `json:"AlbType,omitempty" name:"AlbType"`
	// IsCdn=3时，需要填此参数，表示自定义header

	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`
	// 规则引擎类型， 1: menshen,   2:tiga

	EngineType *int64 `json:"EngineType,omitempty" name:"EngineType"`
	// 云类型:public:公有云；private:私有云;hybrid:混合云

	CloudType *string `json:"CloudType,omitempty" name:"CloudType"`
}

type DeleteBotSessionKeyRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DeleteBotSessionKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBotSessionKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSpartaProtectionRequest struct {
	*tchttp.BaseRequest

	// 域名列表

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 实例类型

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DeleteSpartaProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSpartaProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpartaProtectionModeRequest struct {
	*tchttp.BaseRequest

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 0是修改规则引擎状态，1是修改AI的状态

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ModifySpartaProtectionModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpartaProtectionModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveBypassAllRuleRequest struct {
	*tchttp.BaseRequest
}

func (r *RemoveBypassAllRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveBypassAllRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CCRuleLists struct {

	// 总数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 规则

	Res []*CCRuleItems `json:"Res,omitempty" name:"Res"`
}

type DeleteDomainWhiteRulesRequest struct {
	*tchttp.BaseRequest

	// 需要删除的规则域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 需要删除的白名单规则

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteDomainWhiteRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDomainWhiteRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstancesResult struct {

	// 套餐版本

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 是否自动续费，1：自动续费

	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
	// 套餐的购买时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 购买的套餐信息

	Cls *ClsPackage `json:"Cls,omitempty" name:"Cls"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 套餐子域名个数限制

	DomainLimit *uint64 `json:"DomainLimit,omitempty" name:"DomainLimit"`
	// 套餐子域名已经使用的个数

	DomainCount *uint64 `json:"DomainCount,omitempty" name:"DomainCount"`
	// 套餐主域名使用个数

	MainDomainCount *uint64 `json:"MainDomainCount,omitempty" name:"MainDomainCount"`
	// 套餐主域名个数限制

	MainDomainLimit *uint64 `json:"MainDomainLimit,omitempty" name:"MainDomainLimit"`
	// qps套餐信息

	QPS *QPSPackage `json:"QPS,omitempty" name:"QPS"`
	// 当前的qps峰值

	MaxQPS *uint64 `json:"MaxQPS,omitempty" name:"MaxQPS"`
	// 购买的域名套餐

	DomainPkg *DomainPackage `json:"DomainPkg,omitempty" name:"DomainPkg"`
	// 计费id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 付费模式

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type PageInfo struct {

	// 页码

	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`
	// 页条目数量

	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
}

type SpartaProtectionPort struct {

	// nginx Id

	NginxServerId *uint64 `json:"NginxServerId,omitempty" name:"NginxServerId"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 后端端口

	UpstreamPort *string `json:"UpstreamPort,omitempty" name:"UpstreamPort"`
	// 后端协议

	UpstreamProtocol *string `json:"UpstreamProtocol,omitempty" name:"UpstreamProtocol"`
}

type DescribeRuleUpdateLogRequest struct {
	*tchttp.BaseRequest

	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// cn/en 默认为cn

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *DescribeRuleUpdateLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleUpdateLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostFlowModeInnerRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// WAF流量模式，1：清洗模式，0：镜像模式（默认）

	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *ModifyHostFlowModeInnerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostFlowModeInnerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserClbWafRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserClbWafRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserClbWafRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSpartaProtectionsAutoRequest struct {
	*tchttp.BaseRequest

	// 多域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *AddSpartaProtectionsAutoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSpartaProtectionsAutoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAggregateTopNRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 需要查询TOP的维度名，见AggregateDimension的key定义

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 查询过滤器

	Filters []*BotDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 开始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 需要的Top数，默认5， 最大值100

	TopN *uint64 `json:"TopN,omitempty" name:"TopN"`
}

func (r *DescribeBotAggregateTopNRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAggregateTopNRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainRulesEnResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表详情

		Rules []*RuleEn `json:"Rules,omitempty" name:"Rules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainRulesEnResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainRulesEnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例个数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 实例列表

		Instances []*MonitorInstanceItem `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPublishOrCancelRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPublishOrCancelRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPublishOrCancelRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstancePriceItem struct {

	// 商品的编码

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 折扣后的价格

	RealTotalCost *int64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 折扣前的价格

	TotalCost *int64 `json:"TotalCost,omitempty" name:"TotalCost"`
}

type CheckGlobalWhiteRuleNameRequest struct {
	*tchttp.BaseRequest

	// 要查询的名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CheckGlobalWhiteRuleNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckGlobalWhiteRuleNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名详情

		Host *HostRecord `json:"Host,omitempty" name:"Host"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCachePathPaging struct {

	// 起始页

	Index *string `json:"Index,omitempty" name:"Index"`
	// 页数目

	Count *string `json:"Count,omitempty" name:"Count"`
}

type DescribeParametersListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参数列表

		ParametersList []*string `json:"ParametersList,omitempty" name:"ParametersList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeParametersListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeParametersListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RateLimitCommonRsp struct {

	// 响应码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 提示信息

	Info *string `json:"Info,omitempty" name:"Info"`
}

type AddDomainSgTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDomainSgTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDomainSgTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddDomainWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则id

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDomainWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDomainWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleStatusRequest struct {
	*tchttp.BaseRequest

	// 需要设置的domain

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// WEB 安全模块开关，0或1

	WebSecurity *uint64 `json:"WebSecurity,omitempty" name:"WebSecurity"`
	// 访问控制模块开关，0或者1

	AccessControl *uint64 `json:"AccessControl,omitempty" name:"AccessControl"`
	// CC模块开关，0或者1

	CcProtection *uint64 `json:"CcProtection,omitempty" name:"CcProtection"`
	// 防篡改模块开关，0或者1

	AntiTamper *uint64 `json:"AntiTamper,omitempty" name:"AntiTamper"`
	// 防泄漏模块开关，0或者1

	AntiLeakage *uint64 `json:"AntiLeakage,omitempty" name:"AntiLeakage"`
	// API安全模块开关，0或者1

	ApiProtection *uint64 `json:"ApiProtection,omitempty" name:"ApiProtection"`
	// 限流模块开关，0或1

	RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`
}

func (r *ModifyModuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyModuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiListVersionTwoRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间

	StartTs *int64 `json:"StartTs,omitempty" name:"StartTs"`
	// 查询结束时间

	EndTs *int64 `json:"EndTs,omitempty" name:"EndTs"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 过滤条件

	Filters []*ApiDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 页面索引，第几页

	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页面大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 排序方法，1 升序，-1 降序

	Sort []*string `json:"Sort,omitempty" name:"Sort"`
	// 是否进行总数查询

	NeedTotalCount *bool `json:"NeedTotalCount,omitempty" name:"NeedTotalCount"`
}

func (r *DescribeApiListVersionTwoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiListVersionTwoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccessDownloadRecordsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAccessDownloadRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccessDownloadRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWafAutoDenyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// WAF 自动封禁配置项

		WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitempty" name:"WafAutoDenyDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWafAutoDenyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWafAutoDenyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则序号组

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 用户域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DeleteAttackWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttackWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PageInfoForInt struct {

	// 页码

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页个数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type PolicyDetail struct {

	// 总数

	Total *float64 `json:"Total,omitempty" name:"Total"`
	// 用户

	User *float64 `json:"User,omitempty" name:"User"`
	// 通用

	Common *float64 `json:"Common,omitempty" name:"Common"`
	// 激活

	Activity *float64 `json:"Activity,omitempty" name:"Activity"`
}

type DescribeBotGlobalUAStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全局配置详情

		Data []*UAStrategyDetail `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotGlobalUAStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotGlobalUAStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainEngineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// menshen,tiga

		EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainEngineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotIdDetail struct {

	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	BotId *string `json:"BotId,omitempty" name:"BotId"`
	// 规则开关

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 规则动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 风险等级

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 规则类型

	BotIdType *string `json:"BotIdType,omitempty" name:"BotIdType"`
	// 修改时间

	ModifyTime *int64 `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 插入时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 影响

	Influence *string `json:"Influence,omitempty" name:"Influence"`
}

type CCRuleData struct {

	// cc规则

	Res []*CCRuleItem `json:"Res,omitempty" name:"Res"`
	// 规则数目

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type PeakPointsItem struct {

	// 秒级别时间戳

	Time *uint64 `json:"Time,omitempty" name:"Time"`
	// QPS

	Access *uint64 `json:"Access,omitempty" name:"Access"`
	// 上行带宽峰值，单位B

	Up *uint64 `json:"Up,omitempty" name:"Up"`
	// 下行带宽峰值，单位B

	Down *uint64 `json:"Down,omitempty" name:"Down"`
	// Web攻击次数

	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`
	// CC攻击次数

	Cc *uint64 `json:"Cc,omitempty" name:"Cc"`
	// Bot qps

	BotAccess *uint64 `json:"BotAccess,omitempty" name:"BotAccess"`
	// WAF返回给客户端状态码5xx次数

	StatusServerError *uint64 `json:"StatusServerError,omitempty" name:"StatusServerError"`
	// WAF返回给客户端状态码4xx次数

	StatusClientError *uint64 `json:"StatusClientError,omitempty" name:"StatusClientError"`
	// WAF返回给客户端状态码302次数

	StatusRedirect *uint64 `json:"StatusRedirect,omitempty" name:"StatusRedirect"`
	// WAF返回给客户端状态码202次数

	StatusOk *uint64 `json:"StatusOk,omitempty" name:"StatusOk"`
	// 源站返回给WAF状态码5xx次数

	UpstreamServerError *uint64 `json:"UpstreamServerError,omitempty" name:"UpstreamServerError"`
	// 源站返回给WAF状态码4xx次数

	UpstreamClientError *uint64 `json:"UpstreamClientError,omitempty" name:"UpstreamClientError"`
	// 源站返回给WAF状态码302次数

	UpstreamRedirect *uint64 `json:"UpstreamRedirect,omitempty" name:"UpstreamRedirect"`
	// 黑名单次数

	BlackIP *uint64 `json:"BlackIP,omitempty" name:"BlackIP"`
	// 防篡改次数

	Tamper *uint64 `json:"Tamper,omitempty" name:"Tamper"`
	// 信息防泄露次数

	Leak *uint64 `json:"Leak,omitempty" name:"Leak"`
	// 访问控制

	ACL *uint64 `json:"ACL,omitempty" name:"ACL"`
	// 小程序 qps

	WxAccess *uint64 `json:"WxAccess,omitempty" name:"WxAccess"`
}

type DescribeAIModelStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模块状态

		UpdatedTimestamp *string `json:"UpdatedTimestamp,omitempty" name:"UpdatedTimestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAIModelStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAIModelStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostAccessLogStatusRequest struct {
	*tchttp.BaseRequest

	// 1表示开启功能；0表示关闭功能

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 负载均衡型waf取值clb-waf；SaaS型waf取值sparta-waf；不给出默认是sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 要修改状态的合法域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ModifyHostAccessLogStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostAccessLogStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockPage struct {

	// 唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 标记名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// HTTP状态码

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// HTTP的Content-Type

	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	// Html源码

	Html *string `json:"Html,omitempty" name:"Html"`
	// 0待审核，1审核通过，2审核拒绝

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 自定义头部信息

	CustomHeaders []*CustomHeader `json:"CustomHeaders,omitempty" name:"CustomHeaders"`
}

type DescribeApiSecEventOverviewRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 过滤条件

	Filters []*ApiDataFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiSecEventOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecEventOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotRequestLogRequest struct {
	*tchttp.BaseRequest

	// 非必需，因为第一页为空，后面的请求使用上一次响应中的值

	LastIndex *string `json:"LastIndex,omitempty" name:"LastIndex"`
	// 每页限制条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询过滤器

	Filters []*BotDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 排序

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 开始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
}

func (r *DescribeBotRequestLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRequestLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCdcResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 资源信息列表

		List []*CdcResource `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdcResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCdcResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpsertSessionRequest struct {
	*tchttp.BaseRequest

	// Session对应ID

	SessionID *int64 `json:"SessionID,omitempty" name:"SessionID"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// session来源位置

	Source *string `json:"Source,omitempty" name:"Source"`
	// 提取类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 提取key或者起始匹配模式

	KeyOrStartMat *string `json:"KeyOrStartMat,omitempty" name:"KeyOrStartMat"`
	// 起始偏移位置

	StartOffset *string `json:"StartOffset,omitempty" name:"StartOffset"`
	// 版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 结束匹配模式

	EndMat *string `json:"EndMat,omitempty" name:"EndMat"`
	// 结束偏移位置

	EndOffset *string `json:"EndOffset,omitempty" name:"EndOffset"`
	// Session名

	SessionName *string `json:"SessionName,omitempty" name:"SessionName"`
}

func (r *UpsertSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SceneJsInjectRule struct {

	// 规则对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 可选的值为：0（放行） 1（阻断） 2（人机识别） 3（观察） 4（重定向）

	RuleAction *int64 `json:"RuleAction,omitempty" name:"RuleAction"`
	// 动作为重定向的时候重定向的URL

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 是否开启自动化工具识别，true：是，false：否

	AutomatedTool *bool `json:"AutomatedTool,omitempty" name:"AutomatedTool"`
	// 是否开启页面防调试，true：是，false：否

	CrackBehavior *bool `json:"CrackBehavior,omitempty" name:"CrackBehavior"`
	// 搜索引擎白名单，可选的值为：["sogou", "baidu", "yandex", "360", "yahoo", "bing", "google"]

	GoodBot []*string `json:"GoodBot,omitempty" name:"GoodBot"`
	// 【场景化版本改为复用字段】前端对抗规则ID

	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
	// 规则内部ID

	RasSiteID *string `json:"RasSiteID,omitempty" name:"RasSiteID"`
	// 前端对抗的开关，0（关闭） 1（开启）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 请求白名单

	ReqWhiteList []*SceneWhiteListItem `json:"ReqWhiteList,omitempty" name:"ReqWhiteList"`
	// 响应白名单

	RspWhiteList []*SceneWhiteListItem `json:"RspWhiteList,omitempty" name:"RspWhiteList"`
	// 前端对抗的防护URL，老版本使用，已经废弃

	CheckURL *string `json:"CheckURL,omitempty" name:"CheckURL"`
	// 是否拦截扫描工具（目前暂不支持设置）

	PreventScanner *bool `json:"PreventScanner,omitempty" name:"PreventScanner"`
	// 前端对抗的防护URL，支持多个URL

	CheckUrls []*string `json:"CheckUrls,omitempty" name:"CheckUrls"`
	// 【场景化新增】当查询全局配置时，为空；否则为具体的场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

type ModifyWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadAttackRecordInfo struct {

	// 记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 下载任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 当前下载任务的日志条数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 下载任务运行状态：-1-下载超时，0-下载等待，1-下载完成，2-下载失败，4-正在下载

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 下载文件URL

	Url *string `json:"Url,omitempty" name:"Url"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后更新修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 下载任务需下载的日志总条数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type SessionData struct {

	// session定义

	Res []*SessionItem `json:"Res,omitempty" name:"Res"`
}

type DescribeBotDetailVersionTwoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// bot命中信息

		RequestBotInfo *BotItemRequestBotInfo `json:"RequestBotInfo,omitempty" name:"RequestBotInfo"`
		// 请求特征信息&会话基础信息

		RequestBaseInfo *BotItemRequestBaseInfo `json:"RequestBaseInfo,omitempty" name:"RequestBaseInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotDetailVersionTwoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotDetailVersionTwoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpAttackWhiteRuleNewRequest struct {
	*tchttp.BaseRequest

	// 排序字段，支持 user_id, signature_id, modify_time

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选字段支持 AppId，Domain

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOpAttackWhiteRuleNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpAttackWhiteRuleNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserSignatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserSignatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserSignatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddApiRuleRequest struct {
	*tchttp.BaseRequest

	// 该条api的请求方式

	Method *string `json:"Method,omitempty" name:"Method"`
	// 用户添加api具体的规则列表

	RuleList []*AddApiRule `json:"RuleList,omitempty" name:"RuleList"`
	// api规则对应的状态

	ApiAction *uint64 `json:"ApiAction,omitempty" name:"ApiAction"`
	// api规则的描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// api规则对应的url路径

	Url *string `json:"Url,omitempty" name:"Url"`
	// API规则对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// api对应的开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *AddApiRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddApiRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiSecEventChangeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更历史列表

		Data []*ApiSecEventChange `json:"Data,omitempty" name:"Data"`
		// 数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiSecEventChangeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecEventChangeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpSignatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则列表

		Rules []*UserSignatureRule `json:"Rules,omitempty" name:"Rules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpSignatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpSignatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessMetaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccessMetaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessMetaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchApiRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchApiRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchApiRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotSceneStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotSceneStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSceneStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VIPStateFilter struct {

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 无

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type DescribePostCKafkaFlowsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客户的投递流列表

		PostCKafkaFlows []*PostCKafkaFlowInfo `json:"PostCKafkaFlows,omitempty" name:"PostCKafkaFlows"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePostCKafkaFlowsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePostCKafkaFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotSceneActionRuleRequest struct {
	*tchttp.BaseRequest

	// 100-宽松、200-中等、300-严格、0-自定义

	Level *string `json:"Level,omitempty" name:"Level"`
	// 生效范围，为空[]表示全部生效

	Scope []*BotActionScopeRuleEntry `json:"Scope,omitempty" name:"Scope"`
	// default:默认场景 custom：自定义创建

	Type *string `json:"Type,omitempty" name:"Type"`
	// global-全局匹配 custom-自定义范围

	ScopeType *string `json:"ScopeType,omitempty" name:"ScopeType"`
	// 匹配条件间的与或关系

	ActionMatchType *string `json:"ActionMatchType,omitempty" name:"ActionMatchType"`
	// 动作策略ID,为""表示新增一条动作策略，需要判断当前动作策略是否满5条；不为""时表示修改动作策略

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 动作策略优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 分数范围

	Score []*BotScoreRuleEntry `json:"Score,omitempty" name:"Score"`
	// 动作策略开关状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 动作策略名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *ModifyBotSceneActionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSceneActionRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检查CAM角色服务是否开通

		Enable *bool `json:"Enable,omitempty" name:"Enable"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInstancesRequest struct {
	*tchttp.BaseRequest

	// 地域限制

	QueryRegion *string `json:"QueryRegion,omitempty" name:"QueryRegion"`
	// 版本：clb-waf，sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 对实例id或者实例名称进行查询

	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeUserInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiSecEventChangeListRequest struct {
	*tchttp.BaseRequest

	// 排序，第一个元素为排序的key，第二个元素为排序规则，其中1 为升序排列，而-1 是用于降序排列

	Sort []*string `json:"Sort,omitempty" name:"Sort"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 事件id，如果为空，则表示查询api维度的时间

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// api名称

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 请求方式

	Method *string `json:"Method,omitempty" name:"Method"`
	// 筛选条件

	Filters []*ApiDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 当前页索引

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 每一页显示多少条数据

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeApiSecEventChangeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecEventChangeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CCRuleItem struct {

	// 动作

	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`
	// 高级模式

	Advance *uint64 `json:"Advance,omitempty" name:"Advance"`
	// 时间周期

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 限制次数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 匹配方法

	MatchFunc *uint64 `json:"MatchFunc,omitempty" name:"MatchFunc"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 更新时间戳

	TsVersion *uint64 `json:"TsVersion,omitempty" name:"TsVersion"`
	// 匹配url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 策略动作有效时间

	ValidTime *uint64 `json:"ValidTime,omitempty" name:"ValidTime"`
	// 高级参数

	OptionsArr *string `json:"OptionsArr,omitempty" name:"OptionsArr"`
}

type DescribeUserSceneInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户的场景信息列表，返回[[{Domain,SceneId,SceneName, []ActionRuleInfo }…]列表，其中Domain:SceneId为1对多关系, []ActionRuleInfo 表示隶属于该场景下的所有动作策略

		UserSceneList []*UserSceneInfo `json:"UserSceneList,omitempty" name:"UserSceneList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserSceneInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserSceneInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDNSDetectDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDNSDetectDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDNSDetectDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpAccessControlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 输出

		Data *IpAccessControlData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpAccessControlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppletAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 小程序接入列表

		Applets []*AppletAccess `json:"Applets,omitempty" name:"Applets"`
		// 满足过滤条件的接入小程序的总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppletAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppletAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值

		Data *SearchLogsSimpleRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogCountRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 查询结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 需要查询的域名，全部填写all

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 风险等级

	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 拦截状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 自定义规则Id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 攻击者IP

	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`
	// 攻击类型

	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

func (r *DescribeAttackLogCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackLogCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TypeEn struct {

	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 规则英文名称

	NameEng *string `json:"NameEng,omitempty" name:"NameEng"`
	// 规则类型英文描述

	DescriptionEng *string `json:"DescriptionEng,omitempty" name:"DescriptionEng"`
}

type DescribeDomainBlockPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前域名使用的自定义拦截页面ID，没有使用自定义拦截页面，就为0

		PageId *uint64 `json:"PageId,omitempty" name:"PageId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainBlockPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainBlockPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FreshAntiFakeUrlRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *FreshAntiFakeUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FreshAntiFakeUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApi struct {

	// 用户删除指定url的api

	Url *string `json:"Url,omitempty" name:"Url"`
	// 用户删除指定请求方法的api

	Method *string `json:"Method,omitempty" name:"Method"`
}

type CreateWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 白名单内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 白名单的开关状态，0：关闭，1：开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 白名单的类型，jsinject_req：前端对抗规则的请求白名单，jsinject_rsp：前端对抗规则的响应白名单

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDownloadRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDownloadRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMiniProgramPeakPointRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	FromTime *uint64 `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间

	ToTime *uint64 `json:"ToTime,omitempty" name:"ToTime"`
	// WAF实例ID，不传则不过滤

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 查询的域名，如果查询所有域名数据，该参数不填写	-

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 安全网关接入节点ID

	AccessID *string `json:"AccessID,omitempty" name:"AccessID"`
	// 安全网关ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 7个值可选：
	// 1、AccessBefore 原生防护前总量
	// 2、AccessAfter 	原生防护后总量
	// 3、QpsBefore 原生防护前QPS
	// 4、QpsAfter  	原生防护后QPS
	// 5、BandWidth  带宽
	// 6、Attack  	web安全拦截量
	// 7、CC CC拦截量

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeMiniProgramPeakPointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMiniProgramPeakPointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostAccessLogStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostAccessLogStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostAccessLogStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessMetaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 清空次数

		RewriteCounts *int64 `json:"RewriteCounts,omitempty" name:"RewriteCounts"`
		// 日志主题ID

		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
		// 访问日志数据占用的存储量大小，单位GB

		UsedStorageSize *int64 `json:"UsedStorageSize,omitempty" name:"UsedStorageSize"`
		// 访问日志包到期时间

		ValidTime *int64 `json:"ValidTime,omitempty" name:"ValidTime"`
		// 自动续费标记，

		RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
		// 当用户超量存储，该值表示建议用户的存储天数。未超量，该字段无意义

		SuggestPeriod *int64 `json:"SuggestPeriod,omitempty" name:"SuggestPeriod"`
		// 日志包状态，-1: 日志包不存在 1： 日志包正常使用 2：日志包隔离期 3：日志包销毁期

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 存储时长修改时间

		ModifyPeriodTime *int64 `json:"ModifyPeriodTime,omitempty" name:"ModifyPeriodTime"`
		// 生效域名及大字段配置

		DomainsConfig []*DomainConfig `json:"DomainsConfig,omitempty" name:"DomainsConfig"`
		// 访问日志包总的存储量大小，单位GB

		TotalStorageSize *int64 `json:"TotalStorageSize,omitempty" name:"TotalStorageSize"`
		// 访问日志的存储时长，单位为天数

		Period *int64 `json:"Period,omitempty" name:"Period"`
		// 开启访问日志某些字段是否存储

		WriteConfig *FieldWriteConfig `json:"WriteConfig,omitempty" name:"WriteConfig"`
		// 插入时间

		InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessMetaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessMetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpartaProtectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySpartaProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GoodsDetail struct {

	// 时间间隔

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 单位，支持m、y、d

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 二级产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费策略id

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// waf产品码

	ProductInfo []*ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// waf实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// QPS数量

	ElasticQps *int64 `json:"ElasticQps,omitempty" name:"ElasticQps"`
	// 弹性账单

	FlexBill *int64 `json:"FlexBill,omitempty" name:"FlexBill"`
	// 1:自动续费，0:不自动续费

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// waf购买的实际地域信息

	RealRegion *int64 `json:"RealRegion,omitempty" name:"RealRegion"`
	// Waf实例对应的二级产品码

	Type *string `json:"Type,omitempty" name:"Type"`
	// 计费细项标签数组

	LabelTypes []*string `json:"LabelTypes,omitempty" name:"LabelTypes"`
	// 计费细项标签数量，一般和SvLabelType一一对应

	LabelCounts []*int64 `json:"LabelCounts,omitempty" name:"LabelCounts"`
	// 变配使用，实例到期时间

	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
	// 对存在的实例购买bot 或api 安全

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type LimitRule struct {

	// 规则名

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// api路径列表

	ApiPaths []*string `json:"ApiPaths,omitempty" name:"ApiPaths"`
	// 限流配额，包含限流周期和配额总数，可配置多个

	Amount *Amount `json:"Amount,omitempty" name:"Amount"`
	// 请求参数匹配条件，需全匹配才通过

	Arguments []*MatchArgument `json:"Arguments,omitempty" name:"Arguments"`
	// 优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 限流效果，支持REJECT（直接拒绝）,UNIRATE（匀速排队），默认REJECT

	LimitStrategy *int64 `json:"LimitStrategy,omitempty" name:"LimitStrategy"`
	// 限流对象，支持API,DOMAIN

	LimitObject *string `json:"LimitObject,omitempty" name:"LimitObject"`
	// 是否启用该限流规则，默认为1（启用）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 修改时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type AddApiRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddApiRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddApiRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyBotUCBPreinstallRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 输出

		Data *CopyBotsUCBPreinstallRuleRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyBotUCBPreinstallRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBotUCBPreinstallRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpartaPackageRenewRequest struct {
	*tchttp.BaseRequest

	// "true"表示开启自动续费功能；“false”表示关闭自动续费功能

	AutoRenew *string `json:"AutoRenew,omitempty" name:"AutoRenew"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *ModifySpartaPackageRenewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpartaPackageRenewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotUBRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Data *BotListData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotUBRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotUBRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 套餐详情

		Data *AppIdDetail `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApi struct {

	// api规则id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// api的url路径

	Url *string `json:"Url,omitempty" name:"Url"`
	// 该条api规则的描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 该条api规则的来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 该条api规则的请求方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 该条api对应的参数列表

	Parameters []*string `json:"Parameters,omitempty" name:"Parameters"`
	// 该条api对应的动作，包括观察和拦截

	Action *uint64 `json:"Action,omitempty" name:"Action"`
	// 该条api对应的开关，包括开和关

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 该条api添加的时间戳

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
}

type CreateAppletAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppletAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppletAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDNSDetectDomainRequest struct {
	*tchttp.BaseRequest

	// 域名的权威记录数组

	AuthIP []*string `json:"AuthIP,omitempty" name:"AuthIP"`
	// 添加的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *CreateDNSDetectDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDNSDetectDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotSceneInfo struct {

	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 场景类型，default:默认场景,custom:非默认场景

	Type *string `json:"Type,omitempty" name:"Type"`
	// 场景名

	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 场景模板类型，登录: login  秒杀:seckill  爬内容：crawl 自定义: custom

	BusinessType []*string `json:"BusinessType,omitempty" name:"BusinessType"`
	// 客户端类型，浏览器/H5 : browser  小程序: miniApp  App:

	ClientType []*string `json:"ClientType,omitempty" name:"ClientType"`
	// 优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 匹配范围

	MatchCondition []*BotSceneMatchCondition `json:"MatchCondition,omitempty" name:"MatchCondition"`
	// 场景开关

	SceneStatus *bool `json:"SceneStatus,omitempty" name:"SceneStatus"`
	// 前端对抗开关

	JsInjectStatus *bool `json:"JsInjectStatus,omitempty" name:"JsInjectStatus"`
	// AI开关

	AIStatus *bool `json:"AIStatus,omitempty" name:"AIStatus"`
	// TI开关

	TIStatus *bool `json:"TIStatus,omitempty" name:"TIStatus"`
	// 智能统计开关

	StatisticStatus *bool `json:"StatisticStatus,omitempty" name:"StatisticStatus"`
	// 动作策略数量

	ActionRuleCount *int64 `json:"ActionRuleCount,omitempty" name:"ActionRuleCount"`
	// 自定义规则数量

	UCBCount *int64 `json:"UCBCount,omitempty" name:"UCBCount"`
	// 场景的匹配范围，global-全部匹配 custom-自定义匹配范围

	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`
	// 匹配条件间的与或关系

	ActionMatchType *string `json:"ActionMatchType,omitempty" name:"ActionMatchType"`
	// UA模块开关

	UAStatus *bool `json:"UAStatus,omitempty" name:"UAStatus"`
	// 简易模式场景：前端对抗对应mysql的记录id

	JsInjectRuleId *int64 `json:"JsInjectRuleId,omitempty" name:"JsInjectRuleId"`
	// 简易模式场景：前端对抗配置动作

	JsInjectAction *int64 `json:"JsInjectAction,omitempty" name:"JsInjectAction"`
	// 简易模式场景：前端对抗重定向路径

	JsInjectRedirect *string `json:"JsInjectRedirect,omitempty" name:"JsInjectRedirect"`
	// 简易模式场景：动作策略信息  PS:简易模式只有一个动作策略

	ActionRuleList []*BotSceneActionRule `json:"ActionRuleList,omitempty" name:"ActionRuleList"`
	// 简易模式场景：monitor-观察 intercept-拦截 custom-自定义

	BotIdPattern *string `json:"BotIdPattern,omitempty" name:"BotIdPattern"`
	// 简易模式场景：bot_id规则总数

	BotIdCount *int64 `json:"BotIdCount,omitempty" name:"BotIdCount"`
	// 简易模式场景：观察动作的规则总数

	BotIdMonitorCount *int64 `json:"BotIdMonitorCount,omitempty" name:"BotIdMonitorCount"`
	// 简易模式场景：拦截动作的规则总数

	BotIdInterceptCount *int64 `json:"BotIdInterceptCount,omitempty" name:"BotIdInterceptCount"`
	// 创建场景时选择的规则集

	RuleSetSelection []*string `json:"RuleSetSelection,omitempty" name:"RuleSetSelection"`
}

type BatchAreaBanAppIdDomainData struct {

	// 域名列表

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

type DescribeSpartaProtectionListRequest struct {
	*tchttp.BaseRequest

	// 查询参数

	Search *string `json:"Search,omitempty" name:"Search"`
	// 分页参数

	Paging *PageForDescribeSpartaProtectionList `json:"Paging,omitempty" name:"Paging"`
	// 过滤条件

	Item *DescribeSpartaProtectionListItem `json:"Item,omitempty" name:"Item"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeSpartaProtectionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpartaProtectionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpAttackWhiteRuleNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		List []*OpUserWhiteRuleNew `json:"List,omitempty" name:"List"`
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpAttackWhiteRuleNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpAttackWhiteRuleNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVIPStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		VIPStates []*VIPStates `json:"VIPStates,omitempty" name:"VIPStates"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVIPStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVIPStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotSessionKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotSessionKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSessionKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAntiFakeUrlRequest struct {
	*tchttp.BaseRequest

	// uri

	Uri *string `json:"Uri,omitempty" name:"Uri"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *AddAntiFakeUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAntiFakeUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackDownloadRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载记录

		Records []*DonwloadRecordItem `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackDownloadRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeObjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 对象列表

		ClbObjects []*ClbObject `json:"ClbObjects,omitempty" name:"ClbObjects"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeObjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotUCBPreinstallRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正常情况下为null

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotUCBPreinstallRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotUCBPreinstallRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAutoStatisticRuleRequest struct {
	*tchttp.BaseRequest

	// 是否显示自定义列表标志

	RuleFlag *bool `json:"RuleFlag,omitempty" name:"RuleFlag"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 版本号，场景化版本为4.1.0

	BotVersion *string `json:"BotVersion,omitempty" name:"BotVersion"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotAutoStatisticRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAutoStatisticRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotTCBRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotTCBRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotTCBRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户配置的会话列表，最多5条

		RuleList []*BotToken `json:"RuleList,omitempty" name:"RuleList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostInnerRequest struct {
	*tchttp.BaseRequest

	// 编辑的域名配置信息

	Host *HostRecord `json:"Host,omitempty" name:"Host"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *ModifyHostInnerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostInnerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainWhiteRulesRequest struct {
	*tchttp.BaseRequest

	// 请求的白名单匹配路径

	Url *string `json:"Url,omitempty" name:"Url"`
	// 翻到多少页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页展示的条数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 排序方式,desc表示降序，asc表示升序

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 需要查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainWhiteRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainWhiteRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// instance列表

		Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainsCLSStatusRequest struct {
	*tchttp.BaseRequest

	// 需要修改的域名列表

	Domains []*DomainURI `json:"Domains,omitempty" name:"Domains"`
	// 修改域名的访问日志开关为Status

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDomainsCLSStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainsCLSStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrafficMarkingRequest struct {
	*tchttp.BaseRequest

	// bot: bot防护标记
	// traffic:流量标记

	Selection *string `json:"Selection,omitempty" name:"Selection"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// HTTP头名

	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
	// HTTP头值

	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`
	// Bot防护标记外层开关

	BotStatus *int64 `json:"BotStatus,omitempty" name:"BotStatus"`
	// Bot防护标记 Bot得分开关

	BotScoreStatus *int64 `json:"BotScoreStatus,omitempty" name:"BotScoreStatus"`
	// Bot防护标记 客户端ID开关

	JsInjectStatus *int64 `json:"JsInjectStatus,omitempty" name:"JsInjectStatus"`
}

func (r *ModifyTrafficMarkingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrafficMarkingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleTypesEnRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRuleTypesEnRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleTypesEnRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableRateLimitsRequest struct {
	*tchttp.BaseRequest

	// 规则单元

	RateLimits []*LimitRuleStatusUnit `json:"RateLimits,omitempty" name:"RateLimits"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *EnableRateLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableRateLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackOverviewRequest struct {
	*tchttp.BaseRequest

	// WAF实例ID，不传则不过滤

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 查询开始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 查询结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 客户的Appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 被查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeAttackOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRegionRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 地域属性

	RegionToModify *string `json:"RegionToModify,omitempty" name:"RegionToModify"`
}

func (r *ModifyInstanceRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpsertIpAccessControlRequest struct {
	*tchttp.BaseRequest

	// 具体域名如：test.qcloudwaf.com
	// 全局域名为：global

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// ip 参数列表，json数组由ip，source，note，action，valid_ts组成。ip对应配置的ip地址，source固定为custom值，note为注释，action值42为黑名单，40为白名单，valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））

	Items []*string `json:"Items,omitempty" name:"Items"`
	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 是否为多域名黑白名单，当为多域名的黑白名单时，取值为batch，否则为空

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *UpsertIpAccessControlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertIpAccessControlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHostRequest struct {
	*tchttp.BaseRequest

	// 防护域名配置信息

	Host *HostRecord `json:"Host,omitempty" name:"Host"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *CreateHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否生效

		Effective *bool `json:"Effective,omitempty" name:"Effective"`
		// 索引规则，当 effective 为 true 时返回

		Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotWhiteStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略开关列表

		BotWhiteList []*BotWhiteInfo `json:"BotWhiteList,omitempty" name:"BotWhiteList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotWhiteStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotWhiteStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbWafRegionItem struct {

	// 地域ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 地域中文说明

	Text *string `json:"Text,omitempty" name:"Text"`
	// 地域英文全拼

	Value *string `json:"Value,omitempty" name:"Value"`
	// 地域编码

	Code *string `json:"Code,omitempty" name:"Code"`
}

type SwitchTrafficMarkingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchTrafficMarkingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchTrafficMarkingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApiRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackHistogramResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计详情

		Data []*LogHistogramInfo `json:"Data,omitempty" name:"Data"`
		// 时间段大小

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 统计的条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttackHistogramResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWafThreatenIntelligenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前WAF威胁情报封禁模块详情

		WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitempty" name:"WafThreatenIntelligenceDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWafThreatenIntelligenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWafThreatenIntelligenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiFakeUrlPaging struct {

	// 页码

	Index *string `json:"Index,omitempty" name:"Index"`
	// 页条目数量

	Count *string `json:"Count,omitempty" name:"Count"`
}

type InstanceIDInfo struct {

	// 实例地区

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// sub_domain_limit

	DomainCount *uint64 `json:"DomainCount,omitempty" name:"DomainCount"`
}

type GetAllQpsConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAllQpsConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllQpsConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SceneWhiteListItem struct {

	// 白名单对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 白名单ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 白名单内容

	Content *WhiteListCondition `json:"Content,omitempty" name:"Content"`
	// 白名单的备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 白名单的开关状态，0：关闭，1：开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 白名单的类型，jsinject_req：动态风控的请求白名单，jsinject_rsp：动态风控的响应白名单

	Type *string `json:"Type,omitempty" name:"Type"`
}

type AddSpartaProtectionsAutoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败原因

		FailedInfos []*FailedInfo `json:"FailedInfos,omitempty" name:"FailedInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSpartaProtectionsAutoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSpartaProtectionsAutoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotRecordPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势图数据

		PointsTotal []*uint64 `json:"PointsTotal,omitempty" name:"PointsTotal"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotRecordPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRecordPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTsResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTsResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTsResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTsRegionRequest struct {
	*tchttp.BaseRequest

	// 一级区域

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 二级区域

	SubRegion *string `json:"SubRegion,omitempty" name:"SubRegion"`
}

func (r *DeleteTsRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTsRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAntiInfoLeakRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAntiInfoLeakRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAntiInfoLeakRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAIModelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAIModelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAIModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditHWRecordRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 位置

	Position *string `json:"Position,omitempty" name:"Position"`
	// Value

	Value *string `json:"Value,omitempty" name:"Value"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 执行动作

	ExecuteAction *uint64 `json:"ExecuteAction,omitempty" name:"ExecuteAction"`
	// 域名列表

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

func (r *EditHWRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditHWRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppletAuthorizeV3Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 小程序ID， 授权成功才有此字段

		AppletID *string `json:"AppletID,omitempty" name:"AppletID"`
		// 小程序授权状态
		// 0  失败
		// 1  成功

		AuthStatus *uint64 `json:"AuthStatus,omitempty" name:"AuthStatus"`
		// 授权失败原因

		AuthFailReason *string `json:"AuthFailReason,omitempty" name:"AuthFailReason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppletAuthorizeV3Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppletAuthorizeV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotStatisticPointsRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 步长

	Stride *uint64 `json:"Stride,omitempty" name:"Stride"`
	// 起始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
}

func (r *DescribeBotStatisticPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotStatisticPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleUpdateLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则日志总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则日志列表

		List []*RuleUpdateLog `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleUpdateLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleUpdateLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddPocTestTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddPocTestTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPocTestTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessExportRequest struct {
	*tchttp.BaseRequest

	// 日志导出数量，最大值100w

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 日志导出数据格式。json，csv，默认为json

	Format *string `json:"Format,omitempty" name:"Format"`
	// 日志导出时间排序。desc，asc，默认为desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 日志导出检索语句

	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *CreateAccessExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainDetailsClbRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDomainDetailsClbRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainDetailsClbRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCachePathRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 路径，以/开头

	Path *string `json:"Path,omitempty" name:"Path"`
	// 版本：sparta-waf, clb-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *CreateCachePathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCachePathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotGenerateDealsRequest struct {
	*tchttp.BaseRequest

	// 在该实例下购买bot

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// bot计费下单入参

	Goods []*Goods `json:"Goods,omitempty" name:"Goods"`
	// 1表示尾部;2表示腰部；3表示头部

	UserLevel *uint64 `json:"UserLevel,omitempty" name:"UserLevel"`
}

func (r *ModifyBotGenerateDealsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotGenerateDealsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOpSignatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOpSignatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOpSignatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBotSceneUCBRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正常情况下为null

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBotSceneUCBRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBotSceneUCBRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceElasticModeRequest struct {
	*tchttp.BaseRequest

	// 弹性计费开关

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyInstanceElasticModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceElasticModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyThreatIntelligenceStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 情报项的开关设置，增量更新，支持批量更新。

	Data []*ThreatIntelligenceStatus `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyThreatIntelligenceStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyThreatIntelligenceStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDetailRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePieChartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值

		Data *DescribePieChartRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePieChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePieChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1 和0 1新用户 ; 0 旧用户

		UserType *int64 `json:"UserType,omitempty" name:"UserType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpsertCCRuleRequest struct {
	*tchttp.BaseRequest

	// 动作有效时间

	ValidTime *int64 `json:"ValidTime,omitempty" name:"ValidTime"`
	// 规则ID，新增时填0

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 匹配方法

	MatchFunc *int64 `json:"MatchFunc,omitempty" name:"MatchFunc"`
	// 添加规则的来源事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 规则需要启用的SessionID

	SessionApplied []*int64 `json:"SessionApplied,omitempty" name:"SessionApplied"`
	// CC检测周期

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 动作

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 附加参数

	OptionsArr *string `json:"OptionsArr,omitempty" name:"OptionsArr"`
	// waf版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 检测Url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 操作类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 高级模式

	Advance *string `json:"Advance,omitempty" name:"Advance"`
	// CC检测阈值

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *UpsertCCRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertCCRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQpsWhiteListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeQpsWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQpsWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPackageRenewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功的状态码

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPackageRenewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPackageRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryClientMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0表示关闭，1表示开启

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryClientMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryClientMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCAutoStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置状态

		AutoCCSwitch *int64 `json:"AutoCCSwitch,omitempty" name:"AutoCCSwitch"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCAutoStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCAutoStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafAutoDenyRulesRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeWafAutoDenyRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafAutoDenyRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRateLimitsRequest struct {
	*tchttp.BaseRequest

	// 规则名

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则ID，创建时不用填写

	Id *string `json:"Id,omitempty" name:"Id"`
	// 规则所针对的服务接口

	ApiPaths []*string `json:"ApiPaths,omitempty" name:"ApiPaths"`
	// 限流配额，包含限流周期和配额总数

	Amount *Amount `json:"Amount,omitempty" name:"Amount"`
	// 请求参数匹配条件，需全匹配才通过

	Arguments []*MatchArgument `json:"Arguments,omitempty" name:"Arguments"`
	// 规则优先级，Priority越大，优先级越低

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 限流策略,0:观察,1:拦截，2:人机
	//

	LimitStrategy *int64 `json:"LimitStrategy,omitempty" name:"LimitStrategy"`
	// 默认为false

	Disable *bool `json:"Disable,omitempty" name:"Disable"`
	// 支持API/DOMAIN

	LimitObject *string `json:"LimitObject,omitempty" name:"LimitObject"`
}

func (r *CreateRateLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRateLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDNSDetectDomainListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`
		// 详细的数据

		List []*DescribeDNSDetectDomainListDomainListItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDNSDetectDomainListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSDetectDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleRollbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回滚的规则库id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleRollbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafRequestCountRequest struct {
	*tchttp.BaseRequest

	// domain，可填all

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开始时间

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
}

func (r *DescribeWafRequestCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafRequestCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 编辑的域名ID

		DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJsInjectRuleRequest struct {
	*tchttp.BaseRequest

	// 前端对抗规则对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeJsInjectRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJsInjectRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRateLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 基本返回

		BaseInfo *RateLimitCommonRsp `json:"BaseInfo,omitempty" name:"BaseInfo"`
		// 规则列表

		RateLimits []*LimitRule `json:"RateLimits,omitempty" name:"RateLimits"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRateLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRateLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDomainInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserDomainInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDomainInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotActionRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotActionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotActionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Data *CCRuleData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// bot总数量

		BotCount *uint64 `json:"BotCount,omitempty" name:"BotCount"`
		// 未知类型

		UnknownTypeCount *uint64 `json:"UnknownTypeCount,omitempty" name:"UnknownTypeCount"`
		// 公开类型

		PublicTypeCount *uint64 `json:"PublicTypeCount,omitempty" name:"PublicTypeCount"`
		// 自定义类型

		CustomTypeCount *uint64 `json:"CustomTypeCount,omitempty" name:"CustomTypeCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ThreatIntelligenceEntry struct {

	// 是否开启，true：开启，false：关闭

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 情报名称，以下是当前的取值以及跟情报内容的对应关系。
	// 'Value': 'amazon.com', 'Text': 'Aws'
	// 'Value': 'microsoft.com', 'Text': 'Azure'
	// 'Value': 'cloud.google.com', 'Text': 'Google'
	// 'Value': 'ucloud.cn', 'Text': 'UCloud'
	// 'Value': 'aliyun.com'), 'Text': '阿里云'
	// 'Value': 'baidu.com', 'Text': '百度云'
	// 'Value': 'huawei.com', 'Text': '华为云'
	// 'Value': 'ksyun.com', 'Text': '金山云'
	// 'Value': 'pubyun.com', 'Text': '京东云'
	// 'Value': 'qingcloud.com', 'Text': '青云'
	// 'Value': 'tencent.com', 'Text': '腾讯云'
	// 'Value': 'bot', 'Text': '恶意BOT'
	// 'Value': 'attack', 'Text': 'Web攻击'
	// 'Value': 'proxy', 'Text': '网络代理'

	Text *string `json:"Text,omitempty" name:"Text"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 情报内容，以下是情报内容和情报名称的对应关系。
	// 'Value': 'amazon.com', 'Text': 'Aws'
	// 'Value': 'microsoft.com', 'Text': 'Azure'
	// 'Value': 'cloud.google.com', 'Text': 'Google'
	// 'Value': 'ucloud.cn', 'Text': 'UCloud'
	// 'Value': 'aliyun.com'), 'Text': '阿里云'
	// 'Value': 'baidu.com', 'Text': '百度云'
	// 'Value': 'huawei.com', 'Text': '华为云'
	// 'Value': 'ksyun.com', 'Text': '金山云'
	// 'Value': 'pubyun.com', 'Text': '京东云'
	// 'Value': 'qingcloud.com', 'Text': '青云'
	// 'Value': 'tencent.com', 'Text': '腾讯云'
	// 'Value': 'bot', 'Text': '恶意BOT'
	// 'Value': 'attack', 'Text': 'Web攻击'
	// 'Value': 'proxy', 'Text': '网络代理'

	Value *string `json:"Value,omitempty" name:"Value"`
	// 情报分类

	Type *string `json:"Type,omitempty" name:"Type"`
	// 修改的时间戳，默认0表示没有进行修改

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// type的前端显示

	TypeText *string `json:"TypeText,omitempty" name:"TypeText"`
}

type DescribeBotTypeStatRequest struct {
	*tchttp.BaseRequest

	// 起始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotTypeStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotTypeStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSessionRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// clb-waf或者sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RedisVersionItem struct {

	// Redis集群地区标识

	Region *string `json:"Region,omitempty" name:"Region"`
	// Redis集群地区版本

	Version *uint64 `json:"Version,omitempty" name:"Version"`
}

type MajorEventsPkg struct {

	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 地域

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 申请数量

	InquireNum *int64 `json:"InquireNum,omitempty" name:"InquireNum"`
	// 使用数量

	UsedNum *int64 `json:"UsedNum,omitempty" name:"UsedNum"`
	// 续费标志

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 计费项

	BillingItem *string `json:"BillingItem,omitempty" name:"BillingItem"`
	// 护网包状态

	HWState *int64 `json:"HWState,omitempty" name:"HWState"`
}

type AddSpartaProtectionAutoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSpartaProtectionAutoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSpartaProtectionAutoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQPSResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最大QPS

		Value *uint64 `json:"Value,omitempty" name:"Value"`
		// 最大QPS发生的时间，2020-02-02 20:20:20

		Time *string `json:"Time,omitempty" name:"Time"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserQPSResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserQPSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBusinessRiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBusinessRiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBusinessRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMainClassRequest struct {
	*tchttp.BaseRequest

	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持 MainClassID

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMainClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMainClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryConstantEnumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模块类型数据

		Data *ModuleTypeData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryConstantEnumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryConstantEnumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRuleTagInfo struct {

	// 是否大小写敏感

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 标签索引配置中的字段信息

	KeyValues []*AccessKeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type DescribeUserEditionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserEditionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserEditionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddBatchAreaBanRequest struct {
	*tchttp.BaseRequest

	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据来源 custom-自定义(默认)、batch-批量防护
	//

	Source *string `json:"Source,omitempty" name:"Source"`
	// 域名列表

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
	// 地域列表

	Areas []*string `json:"Areas,omitempty" name:"Areas"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *AddBatchAreaBanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBatchAreaBanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotActionRuleRequest struct {
	*tchttp.BaseRequest

	// bot 等级模式，如严格等级，宽松等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 分数分布，包含分数区间，动作和标签

	Score []*BotScoreEntry `json:"Score,omitempty" name:"Score"`
	// 动作生效范围

	Scope []*BotActionScope `json:"Scope,omitempty" name:"Scope"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ModifyBotActionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotActionRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 索引修改时间，初始值为索引创建时间。

		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
		// 是否生效，true表示生效，false表示未生效

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 索引配置信息
		// 注意：此字段可能返回 null，表示取不到有效值。

		Rule *AccessRuleInfo `json:"Rule,omitempty" name:"Rule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防护状态

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势图的数据

		Points []*string `json:"Points,omitempty" name:"Points"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttackPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceCountKV struct {

	// 地域，如：gz

	Key *string `json:"Key,omitempty" name:"Key"`
	// 实例个数

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type ModifyCustomWhiteRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomWhiteRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomWhiteRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotSessionKeyRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotSessionKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSessionKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotActionScope struct {

	// 要匹配的key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 操作符号，包含或者不包含

	Op *string `json:"Op,omitempty" name:"Op"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type IpHitItem struct {

	// 动作

	Action *uint64 `json:"Action,omitempty" name:"Action"`
	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 时间戳

	TsVersion *uint64 `json:"TsVersion,omitempty" name:"TsVersion"`
	// 有效截止时间戳

	ValidTs *uint64 `json:"ValidTs,omitempty" name:"ValidTs"`
}

type DescribeUserLevelRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeUserLevelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserLevelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAiRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关状态 true为开 false为关

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 用户配置的ai自定义列表

		BotAiRule []*BotAiRule `json:"BotAiRule,omitempty" name:"BotAiRule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotAiRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAiRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDomainWhiteRulesRequest struct {
	*tchttp.BaseRequest

	// 需要切换的白名单开关的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 加白的规则id

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SwitchDomainWhiteRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDomainWhiteRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchIpAccessControlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 输出

		Data *BatchIpAccessControlData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBatchIpAccessControlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBatchIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HybridClusterNode struct {

	// 节点ID

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 端口列表

	Ports []*string `json:"Ports,omitempty" name:"Ports"`
	// 节点状态

	State *int64 `json:"State,omitempty" name:"State"`
	// 节点异常原因

	StateDetail *string `json:"StateDetail,omitempty" name:"StateDetail"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最新时间

	LatestUpdateTime *string `json:"LatestUpdateTime,omitempty" name:"LatestUpdateTime"`
}

type ModifyBotTiRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 2表示有实例未购买bot套餐

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotTiRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotTiRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiSecAttackSource struct {

	// 攻击来源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 威胁等级

	EventLevel *string `json:"EventLevel,omitempty" name:"EventLevel"`
	// BOT标签

	BotLabel *string `json:"BotLabel,omitempty" name:"BotLabel"`
	// 变更时间

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 地理位置

	City *string `json:"City,omitempty" name:"City"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
}

type MaxBodyDetect struct {

	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 执行动作

	Action *uint64 `json:"Action,omitempty" name:"Action"`
}

type DescribeHistogramRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 聚类字段，ip为ip聚合，art为响应耗时聚合，url为url聚合，local为ip转化的城市聚合

	QueryField *string `json:"QueryField,omitempty" name:"QueryField"`
	// 条件，access为访问日志，attack为攻击日志

	Source *string `json:"Source,omitempty" name:"Source"`
	// WAF实例ID，不传则不过滤

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 兼容Host，逐步淘汰Host字段

	Host *string `json:"Host,omitempty" name:"Host"`
	// 起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
}

func (r *DescribeHistogramRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHistogramRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAccessAggregateTopNResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// topN结果

		Data []*BotTopItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotAccessAggregateTopNResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAccessAggregateTopNResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotTiRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开关状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 运行状态

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 域名的InstanceID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 版本号，场景化版本为4.1.0

	BotVersion *string `json:"BotVersion,omitempty" name:"BotVersion"`
}

func (r *ModifyBotTiRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotTiRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotSceneUCBRuleRequest struct {
	*tchttp.BaseRequest

	// 兼容老数据和新旧版前端

	VersionFlag *string `json:"VersionFlag,omitempty" name:"VersionFlag"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 翻页组件的起始页

	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`
	// 翻页组件的页数据条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 需要过滤的动作

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 需要过滤的规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 1.BOT全局白名单处调用时，传"global";2.BOT场景配置处调用时，传具体的场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DescribeBotSceneUCBRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSceneUCBRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRuleListRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// asc或者desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// exp_ts或者mod_ts

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeCustomRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCdcClbWafRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CdcRegion的类型描述

		Data []*CdcRegion `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserCdcClbWafRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCdcClbWafRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BypassSpartaProtectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败原因

		FailedMsg *string `json:"FailedMsg,omitempty" name:"FailedMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BypassSpartaProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BypassSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVipInfoRequest struct {
	*tchttp.BaseRequest

	// waf实例id列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeVipInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVipInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPocTestAuthorizationRequest struct {
	*tchttp.BaseRequest

	// 授权状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyPocTestAuthorizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPocTestAuthorizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportApiRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportApiRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportApiRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FiltersItemNew struct {

	// 字段名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 是否精确查找

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type DescribeUserInstanceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各个地域的实例数量

		InstanceCount []*InstanceCountKV `json:"InstanceCount,omitempty" name:"InstanceCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInstanceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInstanceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotItemRequestBotInfo struct {

	// 命中模块

	BotModule *string `json:"BotModule,omitempty" name:"BotModule"`
	// 异常特征信息

	BotFeature *string `json:"BotFeature,omitempty" name:"BotFeature"`
	// 命中策略名称

	BotRuleName *string `json:"BotRuleName,omitempty" name:"BotRuleName"`
	// bot分数

	BotScore *int64 `json:"BotScore,omitempty" name:"BotScore"`
	// ip城市

	City *string `json:"City,omitempty" name:"City"`
	// ip国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// ip类型

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// ip所有者

	IpOwner *string `json:"IpOwner,omitempty" name:"IpOwner"`
	// 命中的威胁情报信息

	TiInfo []*TiInfo `json:"TiInfo,omitempty" name:"TiInfo"`
	// 命中的ai特征信息

	AiInfo []*AiInfo `json:"AiInfo,omitempty" name:"AiInfo"`
	// 命中的智能统计信息

	AutoStat []*AutoStat `json:"AutoStat,omitempty" name:"AutoStat"`
	// bot标签

	BotLabel *string `json:"BotLabel,omitempty" name:"BotLabel"`
	// 命中策略Id

	BotRuleId *string `json:"BotRuleId,omitempty" name:"BotRuleId"`
}

type CustomHeader struct {

	// 自定义头部Key信息

	Key *string `json:"Key,omitempty" name:"Key"`
	// 自定义头部Value信息

	Value *string `json:"Value,omitempty" name:"Value"`
	// 自定义头部类型：default: 内置; custom: 自定义

	Type *string `json:"Type,omitempty" name:"Type"`
}

type AddCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 需要添加策略的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 放行的详情

	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 优先级

	SortId *string `json:"SortId,omitempty" name:"SortId"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 策略详情

	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`
}

func (r *AddCustomWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCustomWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessIndexRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeAccessIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBotUCBFeatureRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteBotUCBFeatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBotUCBFeatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotRuleOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRuleOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessLogItem struct {

	// 日记Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 日志Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type BotDetail struct {

	// 该记录的唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// bot_id保留字段

	BotId *string `json:"BotId,omitempty" name:"BotId"`
	// 源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// bot的动作

	BotAction *string `json:"BotAction,omitempty" name:"BotAction"`
	// bot统计信息

	Stat *BotStat `json:"Stat,omitempty" name:"Stat"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 分数

	Score *BotScore `json:"Score,omitempty" name:"Score"`
	// 会话时长

	SessionDuration *float64 `json:"SessionDuration,omitempty" name:"SessionDuration"`
	// 时间戳

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// string数组

	BotFeature []*string `json:"BotFeature,omitempty" name:"BotFeature"`
	// ip信息

	IpInfo *IpInfo `json:"IpInfo,omitempty" name:"IpInfo"`
}

type CreateAttackDownloadTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		Flow *string `json:"Flow,omitempty" name:"Flow"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAttackDownloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAttackDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescirbeCCRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Data *CCRuleData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescirbeCCRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescirbeCCRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotCountRequest struct {
	*tchttp.BaseRequest

	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// WAF版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeBotCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOpRuleUpdateLogRequest struct {
	*tchttp.BaseRequest

	// 要删除的日志id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteOpRuleUpdateLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOpRuleUpdateLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePostPointsRequest struct {
	*tchttp.BaseRequest

	// 1-CLS投递，2-CKafka投递，默认CLS投递

	PostFlowType *int64 `json:"PostFlowType,omitempty" name:"PostFlowType"`
	// 1-访问日志，2-攻击日志，默认访问日志

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 时间间隔，单位秒，最小60，最大86400

	Interval *int64 `json:"Interval,omitempty" name:"Interval"`
	// 毫秒时间戳

	From *int64 `json:"From,omitempty" name:"From"`
	// 毫秒时间戳

	To *int64 `json:"To,omitempty" name:"To"`
}

func (r *DescribePostPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePostPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正常情况为null

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCiphersDetailRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCiphersDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCiphersDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceIdListRequest struct {
	*tchttp.BaseRequest

	// 版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeInstanceIdListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceIdListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDomainRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchDomainRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDomainRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAntiFakeUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAntiFakeUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAntiFakeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiSecAttackSourceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击源详情

		Data []*ApiSecAttackSource `json:"Data,omitempty" name:"Data"`
		// 数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 异常ua列表

		ApiSecUaAbnormalSource []*ApiSecUaAbnormalSource `json:"ApiSecUaAbnormalSource,omitempty" name:"ApiSecUaAbnormalSource"`
		// 异常ip段列表

		ApiSecIpAbnormalSource []*ApiSecIpAbnormalSource `json:"ApiSecIpAbnormalSource,omitempty" name:"ApiSecIpAbnormalSource"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiSecAttackSourceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecAttackSourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotListData struct {

	// bot列表

	Res []*BotItemsListItem `json:"Res,omitempty" name:"Res"`
	// 统计总数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type FraudPkg struct {

	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 地域

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 申请数量

	InquireNum *int64 `json:"InquireNum,omitempty" name:"InquireNum"`
	// 使用数量

	UsedNum *int64 `json:"UsedNum,omitempty" name:"UsedNum"`
	// 续费标志

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type OWASPCustomizedData struct {

	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 数量

	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type WhiteListCondition struct {

	// 匹配条件：相等、前缀等

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 匹配的值

	Content *string `json:"Content,omitempty" name:"Content"`
}

type DescribeAreaBanAreasRequest struct {
	*tchttp.BaseRequest

	// 需要查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeAreaBanAreasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreaBanAreasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotUCBPreinstallRuleRequest struct {
	*tchttp.BaseRequest

	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 动作

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 有效时间，分钟单位

	ValidTime *uint64 `json:"ValidTime,omitempty" name:"ValidTime"`
	// 优先级

	Prior *uint64 `json:"Prior,omitempty" name:"Prior"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 类别：1表示协议特征类型的内置规则；2表示IP情报类型的内置规则

	Category *uint64 `json:"Category,omitempty" name:"Category"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyBotUCBPreinstallRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotUCBPreinstallRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExpiredInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// instance列表

		ExpiredInstances []*ExpiredInstanceInfo `json:"ExpiredInstances,omitempty" name:"ExpiredInstances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExpiredInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExpiredInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotStat struct {

	// http信息

	HttpHeader *BotHttpHeader `json:"HttpHeader,omitempty" name:"HttpHeader"`
	// cookie信息

	CookieAnalyze *BotCookieAnalyze `json:"CookieAnalyze,omitempty" name:"CookieAnalyze"`
	// UA信息

	UAInfo *BotUaInfo `json:"UAInfo,omitempty" name:"UAInfo"`
	// 请求信息

	ReqAnalyzeRes *BotReqAnalyzeRes `json:"ReqAnalyzeRes,omitempty" name:"ReqAnalyzeRes"`
	// refer信息

	ReferAnalyzeRes *BotReferAnalyzeRes `json:"ReferAnalyzeRes,omitempty" name:"ReferAnalyzeRes"`
	// 协议统计

	BotProtocol []*BotKeyValueItem `json:"BotProtocol,omitempty" name:"BotProtocol"`
	// 状态码统计

	BotStatus []*BotKeyValueItem `json:"BotStatus,omitempty" name:"BotStatus"`
	// 请求方法统计

	BotMethod []*BotKeyValueItem `json:"BotMethod,omitempty" name:"BotMethod"`
	// 时序行为异常指数

	ReqVariance *float64 `json:"ReqVariance,omitempty" name:"ReqVariance"`
	// 时序熵异常指数

	ReqCCE *float64 `json:"ReqCCE,omitempty" name:"ReqCCE"`
	// 是否公开Bot伪造

	UaNotMatchIP *bool `json:"UaNotMatchIP,omitempty" name:"UaNotMatchIP"`
	// 敏感接口访问

	SensitiveRequestFlag *bool `json:"SensitiveRequestFlag,omitempty" name:"SensitiveRequestFlag"`
}

type VipInfo struct {

	// Virtual IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// waf实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CacheUrlItem struct {

	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// uri

	Uri *string `json:"Uri,omitempty" name:"Uri"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeBotIdTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则类型数组

		ValidTypes []*BotIdTypeEntity `json:"ValidTypes,omitempty" name:"ValidTypes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotIdTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotIdTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改状态：0为成功

		ModifyCode *int64 `json:"ModifyCode,omitempty" name:"ModifyCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityRuleCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 访问控制规则的数目

		AccessControl *uint64 `json:"AccessControl,omitempty" name:"AccessControl"`
		// CC防护规则的数目

		CcProtection *uint64 `json:"CcProtection,omitempty" name:"CcProtection"`
		// 防篡改规则的数目

		AntiTamper *uint64 `json:"AntiTamper,omitempty" name:"AntiTamper"`
		// 防泄漏规则的数目

		AntiLeakage *uint64 `json:"AntiLeakage,omitempty" name:"AntiLeakage"`
		// API防护规则的数目

		ApiProtection *uint64 `json:"ApiProtection,omitempty" name:"ApiProtection"`
		// WEB安全规则的数目

		WebSecurity *uint64 `json:"WebSecurity,omitempty" name:"WebSecurity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityRuleCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityRuleCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainsPartInfo struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 类型

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 证书

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// AI防御模式

	Engine *uint64 `json:"Engine,omitempty" name:"Engine"`
	// 是否开启GrayAreas

	GrayAreas []*string `json:"GrayAreas,omitempty" name:"GrayAreas"`
	// 是否开启httpRewrite

	HttpsRewrite *uint64 `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`
	// https回源端口

	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`
	// 是否是cdn

	IsCdn *uint64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// 是否开启gray

	IsGray *uint64 `json:"IsGray,omitempty" name:"IsGray"`
	// 是否是http2

	IsHttp2 *uint64 `json:"IsHttp2,omitempty" name:"IsHttp2"`
	// 是否开启websocket

	IsWebsocket *uint64 `json:"IsWebsocket,omitempty" name:"IsWebsocket"`
	// 负载均衡

	LoadBalance *uint64 `json:"LoadBalance,omitempty" name:"LoadBalance"`
	// 防御模式

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 私钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// ssl id

	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`
	// 回源域名

	UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`
	// 回源类型

	UpstreamType *uint64 `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// 回源ip

	SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`
	// 服务端口配置

	Ports []*PortInfo `json:"Ports,omitempty" name:"Ports"`
	// 证书类型

	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`
	// 回源方式

	UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`
	// 日志包

	Cls *uint64 `json:"Cls,omitempty" name:"Cls"`
	// 一级cname

	Cname *string `json:"Cname,omitempty" name:"Cname"`
	// 是否长连接

	IsKeepAlive *uint64 `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`
	// 是否开启主动健康检测，1表示开启，0表示不开启

	ActiveCheck *uint64 `json:"ActiveCheck,omitempty" name:"ActiveCheck"`
	// TLS版本信息

	TLSVersion *int64 `json:"TLSVersion,omitempty" name:"TLSVersion"`
	// 加密套件信息

	Ciphers []*int64 `json:"Ciphers,omitempty" name:"Ciphers"`
	// 模板

	CipherTemplate *int64 `json:"CipherTemplate,omitempty" name:"CipherTemplate"`
	// 300s

	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitempty" name:"ProxyReadTimeout"`
	// 300s

	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitempty" name:"ProxySendTimeout"`
	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；

	SniType *int64 `json:"SniType,omitempty" name:"SniType"`
	// SniType=3时，需要填此参数，表示自定义的host；

	SniHost *string `json:"SniHost,omitempty" name:"SniHost"`
	// 无

	Weights []*string `json:"Weights,omitempty" name:"Weights"`
	// IsCdn=3时，表示自定义header

	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`
	// 0:关闭xff重置；1:开启xff重置

	XFFReset *int64 `json:"XFFReset,omitempty" name:"XFFReset"`
}

type DescribeDomainEngineRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainEngineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainEngineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotItemTokenIp struct {

	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// key，默认为ip地址

	Key *string `json:"Key,omitempty" name:"Key"`
}

type DescribeApiSecEventTypeListRequest struct {
	*tchttp.BaseRequest

	// 查询结束时间

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 过滤条件

	Filters []*ApiDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 查询开始时间

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
}

func (r *DescribeApiSecEventTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecEventTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHWInfoRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeHWInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHWInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserSignaturePolicyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关状态0/1

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserSignaturePolicyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserSignaturePolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckGlobalWhiteRuleNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1=有重复 0=无重复

		Duplicate *int64 `json:"Duplicate,omitempty" name:"Duplicate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckGlobalWhiteRuleNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckGlobalWhiteRuleNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBatchAreaBanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBatchAreaBanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBatchAreaBanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QPSPackage struct {

	// 资源ID

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 是否自动续费，1：自动续费，0：不自动续费

	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
	// 套餐购买个数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 套餐购买地域，clb-waf暂时没有用到

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeUserInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的实例总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 实例信息列表

		InstanceList []*InstancesResult `json:"InstanceList,omitempty" name:"InstanceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyPostCKafkaFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyPostCKafkaFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyPostCKafkaFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackCountRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询的域名，全部域名不指定

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询的攻击类型，全部攻击类型不指定

	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

func (r *GetAttackCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddGroupProtectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddGroupProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddGroupProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessLogCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 满足条件的数量

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessLogCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessLogCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePieChartRsp struct {

	// 详细的内容

	Piechart []*string `json:"Piechart,omitempty" name:"Piechart"`
}

type DescribeBotSceneListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合筛选条件的场景数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 当TotalCount为0时，返回空

		BotSceneList []*BotSceneInfo `json:"BotSceneList,omitempty" name:"BotSceneList"`
		// true-简易模式

		SimpleFlag *bool `json:"SimpleFlag,omitempty" name:"SimpleFlag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotSceneListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSceneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainDetailsSaasRequest struct {
	*tchttp.BaseRequest

	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainDetailsSaasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainDetailsSaasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAccessLogsRequest struct {
	*tchttp.BaseRequest

	// 日志起始位置的游标，从DescribeAccessCursor获取

	Cursor *string `json:"Cursor,omitempty" name:"Cursor"`
	// 导出的日志条数，最大为1000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ExportAccessLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAccessLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomPayload struct {

	// 载荷类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 数据库id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 学习状态

	LearnStat *string `json:"LearnStat,omitempty" name:"LearnStat"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 更新时间戳

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 载荷值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteDNSDetectDomainRequest struct {
	*tchttp.BaseRequest

	// 删除记录的ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 删除的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DeleteDNSDetectDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDNSDetectDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 类型及个数

		Types []*StatisticType `json:"Types,omitempty" name:"Types"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatisticTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatisticTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyThreatIntelligenceStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyThreatIntelligenceStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyThreatIntelligenceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前返回的攻击日志条数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 翻页游标，如果没有下一页了，这个参数为空""

		Context *string `json:"Context,omitempty" name:"Context"`
		// 攻击日志数组

		Data []*AttackDetail `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopAttackIPRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 域名，不指定则为全部域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询的攻击类型，全部攻击类型不指定

	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

func (r *GetTopAttackIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopAttackIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiAnalyzeStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 需要批量开启的实体列表

	TargetList []*TargetEntity `json:"TargetList,omitempty" name:"TargetList"`
}

func (r *ModifyApiAnalyzeStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiAnalyzeStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionedIpItem struct {

	// 动作

	Action *uint64 `json:"Action,omitempty" name:"Action"`
	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 更新时间戳

	TsVersion *uint64 `json:"TsVersion,omitempty" name:"TsVersion"`
	// 有效时间戳

	ValidTs *uint64 `json:"ValidTs,omitempty" name:"ValidTs"`
}

type BotActionScopeRuleEntry struct {

	// 参数

	Key *string `json:"Key,omitempty" name:"Key"`
	// 匹配符

	Op *string `json:"Op,omitempty" name:"Op"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 对于头部字段匹配value的时候指定的头部名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 470后使用此字段存储多值

	ValueArray []*string `json:"ValueArray,omitempty" name:"ValueArray"`
}

type DescribeCCOpenCustomCountRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeCCOpenCustomCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCOpenCustomCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpAccessControlRequest struct {
	*tchttp.BaseRequest

	// 生效状态

	ValidStatus *int64 `json:"ValidStatus,omitempty" name:"ValidStatus"`
	// 最大有效时间的时间戳

	ValidTimeStampMax *string `json:"ValidTimeStampMax,omitempty" name:"ValidTimeStampMax"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 计数标识

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 最大创建时间的时间戳

	CtsMax *uint64 `json:"CtsMax,omitempty" name:"CtsMax"`
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 最小有效时间的时间戳

	ValidTimeStampMin *string `json:"ValidTimeStampMin,omitempty" name:"ValidTimeStampMin"`
	// 动作，40表示查询白名单，42表示查询黑名单

	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`
	// 最小有效时间的时间戳

	VtsMin *uint64 `json:"VtsMin,omitempty" name:"VtsMin"`
	// 每页返回的数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量，取Limit整数倍。最小值为0，最大值= Total/Limit向上取整

	OffSet *uint64 `json:"OffSet,omitempty" name:"OffSet"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 最大有效时间的时间戳

	VtsMax *uint64 `json:"VtsMax,omitempty" name:"VtsMax"`
	// 最小创建时间的时间戳

	CtsMin *uint64 `json:"CtsMin,omitempty" name:"CtsMin"`
}

func (r *DescribeIpAccessControlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpAccessControlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafAutoDenyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// WAF 自动封禁详情

		WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitempty" name:"WafAutoDenyDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafAutoDenyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafAutoDenyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttackLogPostRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 攻击日志投递开关

	AttackLogPost *int64 `json:"AttackLogPost,omitempty" name:"AttackLogPost"`
}

func (r *ModifyInstanceAttackLogPostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAttackLogPostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddBlockPageRequest struct {
	*tchttp.BaseRequest

	// 展示的名称，长度不能超过30个字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// HTTP状态码，可选的为2XX、4XX、624（CLB-WAF类型的域名只允许应用状态码为624的自定义拦截页面）

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// HTTP返回的Content-Type，可选的为application/json、application/html、text/html

	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	// 自定义页面的HTML代码，长度不能超过10240个字符

	Html *string `json:"Html,omitempty" name:"Html"`
	// 自定义头部信息

	CustomHeaders []*CustomHeader `json:"CustomHeaders,omitempty" name:"CustomHeaders"`
	// 记录ID，新增传空；编辑则传对应ID接口

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *AddBlockPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBlockPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotItemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志详情

		Data []*BotItem `json:"Data,omitempty" name:"Data"`
		// 日志条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotItemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpartUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 区域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 最大Qps

		MaxQps *string `json:"MaxQps,omitempty" name:"MaxQps"`
		// 是否是sparta

		IsSparta *string `json:"IsSparta,omitempty" name:"IsSparta"`
		// 自动续费开关 “true” ： 开启自动续费

		AutoRenew *string `json:"AutoRenew,omitempty" name:"AutoRenew"`
		// cls套餐包详情

		Cls *ClsPackage `json:"Cls,omitempty" name:"Cls"`
		// 套餐有效期

		Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
		// 一级域名使用数量

		MainDomainUsed *string `json:"MainDomainUsed,omitempty" name:"MainDomainUsed"`
		// ID，当未开通waf时，id为0

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 已使用域名数量

		DomainUsed *string `json:"DomainUsed,omitempty" name:"DomainUsed"`
		// qps套餐包详情

		Qps *QPSPackage `json:"Qps,omitempty" name:"Qps"`
		// 域名总数限制

		DomainCount *string `json:"DomainCount,omitempty" name:"DomainCount"`
		// 开始时间

		BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
		// 资源ID

		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
		// 套餐版本信息

		Type *string `json:"Type,omitempty" name:"Type"`
		// 一级域名总数限制

		MainDomainCount *string `json:"MainDomainCount,omitempty" name:"MainDomainCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpartUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpartUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafInfoRequest struct {
	*tchttp.BaseRequest

	// CLB回调WAF接口（获取、删除）的参数

	Params []*ClbHostsParams `json:"Params,omitempty" name:"Params"`
}

func (r *DescribeWafInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BusinessRisk struct {

	// 规则主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 规则动作

	RuleAction *int64 `json:"RuleAction,omitempty" name:"RuleAction"`
	// 拦截url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 分域信息

	DomainCode *string `json:"DomainCode,omitempty" name:"DomainCode"`
	// 策略开关

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 域名信息

	Host *string `json:"Host,omitempty" name:"Host"`
	// 风险等级，业务安全风险类型，包含：reject（高风险），review（人工审核）pass（无风险）

	Level *string `json:"Level,omitempty" name:"Level"`
	// 风险类型：
	// "1", "账号信用低"
	// "11", "低活跃账号"
	// "2", "垃圾账号"
	// "21", "疑似小号"
	// "22", "违规账号"
	// "3", "无效账号"
	// "4", "黑名单"
	// "5", "白名单"
	// "201", "环境异常"
	// "2011", "非常用IP"
	// "2012", "IP异常"
	// "205", "非公网有效ip"
	// "206", "设备异常"
	// "2061", "非常用设备"
	// "2063", "群控设备"

	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
	// 账户类型

	AccountTypes *string `json:"AccountTypes,omitempty" name:"AccountTypes"`
	// 条件信息

	Extend *string `json:"Extend,omitempty" name:"Extend"`
	// 账号提取规则

	Options *string `json:"Options,omitempty" name:"Options"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则UUid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 重定向链接

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 时间版本

	TsVersion *int64 `json:"TsVersion,omitempty" name:"TsVersion"`
	// 生效时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 发布的时间

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 修改的时间

	GmtModify *string `json:"GmtModify,omitempty" name:"GmtModify"`
	// 创建时间

	GmtCreate *string `json:"GmtCreate,omitempty" name:"GmtCreate"`
}

type BatchIpAccessControlItem struct {

	// mongo表自增Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 黑名单42或白名单40

	ActionType *int64 `json:"ActionType,omitempty" name:"ActionType"`
	// 黑白名单的IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 添加路径

	Source *string `json:"Source,omitempty" name:"Source"`
	// 修改时间

	TsVersion *uint64 `json:"TsVersion,omitempty" name:"TsVersion"`
	// 超时时间

	ValidTs *int64 `json:"ValidTs,omitempty" name:"ValidTs"`
	// 域名列表

	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
}

type DescribeApiRuleRequest struct {
	*tchttp.BaseRequest

	// 查询api规则的url路径

	Url *string `json:"Url,omitempty" name:"Url"`
	// 查询api规则的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询api规则的请求方式

	Method *string `json:"Method,omitempty" name:"Method"`
}

func (r *DescribeApiRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotIdTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBotIdTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotIdTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBatchAreaBanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBatchAreaBanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchAreaBanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiListRequest struct {
	*tchttp.BaseRequest

	// 用户的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 用户查询的url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 排序方式

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeApiListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackTypeRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// WAF实例ID，不传则不过滤

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 兼容Host，逐步淘汰Host字段

	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DescribeAttackTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainPostActionRequest struct {
	*tchttp.BaseRequest

	// 0-关闭投递，1-开启投递

	PostCLSAction *int64 `json:"PostCLSAction,omitempty" name:"PostCLSAction"`
	// 0-关闭投递，1-开启投递

	PostCKafkaAction *int64 `json:"PostCKafkaAction,omitempty" name:"PostCKafkaAction"`
	// www.tx.com

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ModifyDomainPostActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainPostActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotUCBRuleRsp struct {

	// 规则列表

	Res []*InOutputBotUCBRule `json:"Res,omitempty" name:"Res"`
	// 规则总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type WafRuleLimit struct {

	// 自定义CC的规格

	CC *uint64 `json:"CC,omitempty" name:"CC"`
	// 自定义策略的规格

	CustomRule *uint64 `json:"CustomRule,omitempty" name:"CustomRule"`
	// 黑白名单的规格

	IPControl *uint64 `json:"IPControl,omitempty" name:"IPControl"`
	// 信息防泄漏的规格

	AntiLeak *uint64 `json:"AntiLeak,omitempty" name:"AntiLeak"`
	// 防篡改的规格

	AntiTamper *uint64 `json:"AntiTamper,omitempty" name:"AntiTamper"`
	// 紧急CC的规格

	AutoCC *uint64 `json:"AutoCC,omitempty" name:"AutoCC"`
	// 地域封禁的规格

	AreaBan *uint64 `json:"AreaBan,omitempty" name:"AreaBan"`
	// 自定义CC中配置session

	CCSession *uint64 `json:"CCSession,omitempty" name:"CCSession"`
	// AI的规格

	AI *uint64 `json:"AI,omitempty" name:"AI"`
	// 精准白名单的规格

	CustomWhite *uint64 `json:"CustomWhite,omitempty" name:"CustomWhite"`
	// api安全的规格

	ApiSecurity *uint64 `json:"ApiSecurity,omitempty" name:"ApiSecurity"`
	// 客户端流量标记的规格

	ClientMsg *uint64 `json:"ClientMsg,omitempty" name:"ClientMsg"`
	// 流量标记的规格

	TrafficMarking *uint64 `json:"TrafficMarking,omitempty" name:"TrafficMarking"`
}

type DescribeApiRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询的api详细规则

		ApiRule *DescribeApiRule `json:"ApiRule,omitempty" name:"ApiRule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteAllDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白名单信息

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteAllDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteAllDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryConfigDeliveryRequest struct {
	*tchttp.BaseRequest

	// 检查AppId

	CheckAppId *uint64 `json:"CheckAppId,omitempty" name:"CheckAppId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 模块

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *QueryConfigDeliveryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryConfigDeliveryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePeakValueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下行带宽峰值，单位B

		Down *uint64 `json:"Down,omitempty" name:"Down"`
		// Web攻击总数

		Attack *uint64 `json:"Attack,omitempty" name:"Attack"`
		// CC攻击总数

		Cc *uint64 `json:"Cc,omitempty" name:"Cc"`
		// QPS峰值

		Access *uint64 `json:"Access,omitempty" name:"Access"`
		// 上行带宽峰值，单位B

		Up *uint64 `json:"Up,omitempty" name:"Up"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePeakValueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePeakValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafThreatenIntelligenceDetails struct {

	// 封禁属性标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 封禁模组启用状态

	DefenseStatus *int64 `json:"DefenseStatus,omitempty" name:"DefenseStatus"`
	// 最后更新时间

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type DescribeBotAiRuleRequest struct {
	*tchttp.BaseRequest

	// 是否显示ai自定义配置列表

	RuleFlag *bool `json:"RuleFlag,omitempty" name:"RuleFlag"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 版本号，场景化版本为4.1.0

	BotVersion *string `json:"BotVersion,omitempty" name:"BotVersion"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名的InstanceID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeBotAiRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAiRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotSceneUCBRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据包

		Data *DescribeBotUCBRuleRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotSceneUCBRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSceneUCBRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReqUserRule struct {

	// 特征序号

	Id *string `json:"Id,omitempty" name:"Id"`
	// 规则开关
	// 0：关
	// 1：开
	// 2：只告警

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 修改原因
	// 0：无(兼容记录为空)
	// 1：业务自身特性误报避免
	// 2：规则误报上报
	// 3：核心业务规则灰度
	// 4：其它

	Reason *int64 `json:"Reason,omitempty" name:"Reason"`
}

type DescribeGlobalWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则列表

		List []*GlobalWhiteRule `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGlobalWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGlobalWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainVerifyResultRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeDomainVerifyResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainVerifyResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleRequest struct {
	*tchttp.BaseRequest

	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MiniProgramItem struct {

	// 秒级别时间戳

	Time *uint64 `json:"Time,omitempty" name:"Time"`
	// 原生防护前总量

	AccessBefore *uint64 `json:"AccessBefore,omitempty" name:"AccessBefore"`
	// 原生防护后总量

	AccessAfter *uint64 `json:"AccessAfter,omitempty" name:"AccessAfter"`
	// 原生防护前QPS

	QpsBefore *uint64 `json:"QpsBefore,omitempty" name:"QpsBefore"`
	// 原生防护后QPS

	QpsAfter *uint64 `json:"QpsAfter,omitempty" name:"QpsAfter"`
	// 带宽

	BandWidth *uint64 `json:"BandWidth,omitempty" name:"BandWidth"`
	// web安全拦截量

	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`
	// CC拦截量

	CC *uint64 `json:"CC,omitempty" name:"CC"`
}

type DescribePostCKafkaFlowsRequest struct {
	*tchttp.BaseRequest

	// 1-访问日志，2-攻击日志，默认为访问日志。

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
}

func (r *DescribePostCKafkaFlowsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePostCKafkaFlowsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotTokenRequest struct {
	*tchttp.BaseRequest

	// 是否删除会话配置标志，true表示删除

	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`
	// 会话名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 会话描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 会话是否开启，true表示开启

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 会话位置 get/post/cookie/headers取值

	Location *string `json:"Location,omitempty" name:"Location"`
	// 会话的key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 会话的匹配方式

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ModifyBotTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogHistogramInfo struct {

	// 日志条数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 时间戳

	TimeStamp *int64 `json:"TimeStamp,omitempty" name:"TimeStamp"`
}

type DescribeCustomRulesPagingInfo struct {

	// 当前页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 当前页的最大数据条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type ModifyAreaBanStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAreaBanStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAreaBanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHostInnerRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 防护域名配置信息

	Host *HostRecord `json:"Host,omitempty" name:"Host"`
}

func (r *CreateHostInnerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostInnerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotUCBFeatureRuleRequest struct {
	*tchttp.BaseRequest

	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 需要过滤的动作

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 需要过滤的规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 翻页组件的起始页

	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`
	// 翻页组件的页数据条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBotUCBFeatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotUCBFeatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyPostCLSFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyPostCLSFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyPostCLSFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCopyCustomRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行状态，返回success或者error

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCopyCustomRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCopyCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDNSDetectDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDNSDetectDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDNSDetectDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePieChartRequest struct {
	*tchttp.BaseRequest

	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 兼容Host，逐步淘汰Host字段

	Host *string `json:"Host,omitempty" name:"Host"`
	// 起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// ua为user_agent，us为upstream_status

	QueryField *string `json:"QueryField,omitempty" name:"QueryField"`
	// WAF实例ID，不传则不过滤

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribePieChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePieChartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAggregateDomainStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 详细的值

		Data []*BotsDomainAggStatStyledItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotAggregateDomainStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAggregateDomainStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainSecurityGroupTipsCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组绑定失败的域名数量
		//

		FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`
		// 安全组是否发生变更的域名数量
		//

		ChangedCount *int64 `json:"ChangedCount,omitempty" name:"ChangedCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainSecurityGroupTipsCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainSecurityGroupTipsCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainEngineRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// menshen, tiga

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
}

func (r *ModifyDomainEngineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainEngineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotKeyValueItem struct {

	// key值

	Key *string `json:"Key,omitempty" name:"Key"`
	// value值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribePolicyStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// clb-waf或者saas-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribePolicyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePostCLSFlowsRequest struct {
	*tchttp.BaseRequest

	// 1-访问日志，2-攻击日志，默认为访问日志。

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
}

func (r *DescribePostCLSFlowsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePostCLSFlowsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescirbeCCRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 页码

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 页的数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// clb-waf 或者 sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 过滤条件

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescirbeCCRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescirbeCCRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiFakeRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值

		Data []*CacheUrlItems `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAntiFakeRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAntiFakeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiListVersionTwoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// api资产列表

		Data []*ApiAsset `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiListVersionTwoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiListVersionTwoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TIGARuleStatusItem struct {

	// 规则库的id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 规则库版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 是否使用中，1代表使用中，0代表没使用

	InUse *int64 `json:"InUse,omitempty" name:"InUse"`
	// 规则库来源，0代表cos

	Source *int64 `json:"Source,omitempty" name:"Source"`
	// 回滚标记，1代表已经下发回滚，0代表没有下发回滚

	Rollback *int64 `json:"Rollback,omitempty" name:"Rollback"`
	// 规则库状态，0代表新建，1代表下发升级，2代表升级中，3代表已完成升级，4代表已删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 升级进度

	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
	// 该规则库的错误码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 完成升级时间

	FinishedTime *string `json:"FinishedTime,omitempty" name:"FinishedTime"`
	// 规则库详细

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 规则库上次操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 规则库文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type DescribeAllDomainRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAllDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotThreatIntelligenceRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 情报的分类，可选字段，默认返回所有类型。

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeBotThreatIntelligenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotThreatIntelligenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotSceneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotSceneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSceneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainURI struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DescribeBotActionRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 动作生效范围

		Scope []*BotActionScope `json:"Scope,omitempty" name:"Scope"`
		// bot模式，严格模式，宽松模式等

		Level *string `json:"Level,omitempty" name:"Level"`
		// bot分数分布,默认是标准模式

		Score []*BotScoreEntry `json:"Score,omitempty" name:"Score"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotActionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotActionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserDomainInfo struct {

	// 用户id

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// waf类型

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 版本

	Level *string `json:"Level,omitempty" name:"Level"`
	// 指定域名访问日志字段的开关

	WriteConfig *string `json:"WriteConfig,omitempty" name:"WriteConfig"`
	// 指定域名是否写cls的开关 1:写 0:不写

	Cls *uint64 `json:"Cls,omitempty" name:"Cls"`
	// 标记是否是混合云接入。hybrid表示混合云接入域名

	CloudType *string `json:"CloudType,omitempty" name:"CloudType"`
}

type InstanceIDName struct {

	// 实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type AddCustomPayloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// id

		PayloadId *string `json:"PayloadId,omitempty" name:"PayloadId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCustomPayloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCustomPayloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest

	// 防护域名，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名，此参数不支持模糊搜索

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 防护域名ID，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名ID，此参数不支持模糊搜索

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 搜索条件，根据此参数对域名做模糊搜索

	Search *string `json:"Search,omitempty" name:"Search"`
	// 复杂的搜索条件

	Item *SearchItem `json:"Item,omitempty" name:"Item"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbHostsParams struct {

	// 负载均衡实例ID，如果不传次参数则默认认为操作的是整个AppId的监听器，如果此参数不为空则认为操作的是对应负载均衡实例。

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器ID，，如果不传次参数则默认认为操作的是整个负载均衡实例，如果此参数不为空则认为操作的是对应负载均衡监听器。

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// WAF实例ID，，如果不传次参数则默认认为操作的是整个负载均衡监听器实例，如果此参数不为空则认为操作的是对应负载均衡监听器的某一个具体的域名。

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

type ApiEvent struct {

	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 事件等级，100,200,300对应低中高

	Level *string `json:"Level,omitempty" name:"Level"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 处置状态，1:新发现，2，确认中，3，已确认，4，已下线，5，已忽略

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 发现时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 更新时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 关联的api

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 请求方式

	Method *string `json:"Method,omitempty" name:"Method"`
}

type BotSceneMatchCondition struct {

	// 匹配参数

	Key *string `json:"Key,omitempty" name:"Key"`
	// 匹配符

	Op *string `json:"Op,omitempty" name:"Op"`
	// 匹配值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 对于头部字段匹配value的时候指定的头部名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 470后使用此入参存在多值

	ValueArray []*string `json:"ValueArray,omitempty" name:"ValueArray"`
}

type DescribeAntiInfoLeakRulesStrategyItem struct {

	// 字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 条件

	CompareFunc *string `json:"CompareFunc,omitempty" name:"CompareFunc"`
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type DescribeAIVerifyResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参数主key

		ParamKey *string `json:"ParamKey,omitempty" name:"ParamKey"`
		// 参数类型

		ParamType *string `json:"ParamType,omitempty" name:"ParamType"`
		// 参数的原始值

		ParamValue *string `json:"ParamValue,omitempty" name:"ParamValue"`
		// 参数的解码后的值

		ParamValueDecode *string `json:"ParamValueDecode,omitempty" name:"ParamValueDecode"`
		// 耗时

		CostTime *float64 `json:"CostTime,omitempty" name:"CostTime"`
		// 返回值

		Return *int64 `json:"Return,omitempty" name:"Return"`
		// 域名

		Host *string `json:"Host,omitempty" name:"Host"`
		// 验证关键字

		Key *string `json:"Key,omitempty" name:"Key"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAIVerifyResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAIVerifyResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则库的描述信息

		Infos []*Info `json:"Infos,omitempty" name:"Infos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRulesInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRulesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyModuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotSceneRequest struct {
	*tchttp.BaseRequest

	// 匹配范围为全局时，传"global"，否则传"custom"

	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`
	// 各匹配条件直接的逻辑关系：and/or

	ActionMatchType *string `json:"ActionMatchType,omitempty" name:"ActionMatchType"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景名称

	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`
	// 场景模板类型，登录: login  秒杀:seckill  爬内容：crawl 自定义: custom

	BusinessType []*string `json:"BusinessType,omitempty" name:"BusinessType"`
	// 场景优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 匹配范围为全局时，传[{"Key": "url", "Op": "prefix", "Value": "/"}]

	MatchCondition []*BotSceneMatchCondition `json:"MatchCondition,omitempty" name:"MatchCondition"`
	//  创建场景时预选开启的规则集，["all"]-表示全部开启

	RuleSetSelection []*string `json:"RuleSetSelection,omitempty" name:"RuleSetSelection"`
	// 场景ID，当NewFlag为true时，为空

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 场景类型，default:默认场景,custom:非默认场景

	Type *string `json:"Type,omitempty" name:"Type"`
	// 客户端类型，浏览器/H5 : browser  小程序: miniApp  App: app

	ClientType []*string `json:"ClientType,omitempty" name:"ClientType"`
}

func (r *ModifyBotSceneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSceneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGlobalWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则序号

		RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGlobalWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGlobalWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbObject struct {

	// 对象ID

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 精准域名列表

	PreciseDomains []*string `json:"PreciseDomains,omitempty" name:"PreciseDomains"`
	// WAF功能开关状态，0关闭1开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// WAF日志开关状态，0关闭1开启

	ClsStatus *int64 `json:"ClsStatus,omitempty" name:"ClsStatus"`
	// CLB对象对应的虚拟域名

	VirtualDomain *string `json:"VirtualDomain,omitempty" name:"VirtualDomain"`
	// 对象名称

	ObjectName *string `json:"ObjectName,omitempty" name:"ObjectName"`
	// 公网地址

	PublicIp []*string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 内网地址

	PrivateIp []*string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// VPC ID

	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`
	// waf实例等级，如果未绑定实例为0

	InstanceLevel *int64 `json:"InstanceLevel,omitempty" name:"InstanceLevel"`
	// clb投递开关

	PostCLSStatus *int64 `json:"PostCLSStatus,omitempty" name:"PostCLSStatus"`
	// kafka投递开关

	PostCKafkaStatus *int64 `json:"PostCKafkaStatus,omitempty" name:"PostCKafkaStatus"`
	// 对象类型：CLB:负载均衡器，TSE:云原生网关

	Type *string `json:"Type,omitempty" name:"Type"`
	// 对象地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type StatisticType struct {

	// 攻击类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 攻击类型数目

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DescribeApiOverviewRequest struct {
	*tchttp.BaseRequest

	// 域名,all的时候，表示appid维度调用

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeApiOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchElasticModeRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 版本，只能是sparta-waf, clb-waf, cdn-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 0代表关闭，1代表打开

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
}

func (r *SwitchElasticModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchElasticModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PostCLSFlowInfo struct {

	// 投递流唯一ID

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
	// 1-访问日志 2-攻击日志

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 最后一次修改时间

	ModifyTime *int64 `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 状态 0-为关闭 1-为启用

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CLS所在区域

	CLSRegion *string `json:"CLSRegion,omitempty" name:"CLSRegion"`
	// CLS日志集合名称

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// CLS日志集合ID

	LogsetID *string `json:"LogsetID,omitempty" name:"LogsetID"`
	// CLS日志主题名称

	LogtopicName *string `json:"LogtopicName,omitempty" name:"LogtopicName"`
	// CLS日志主题ID

	LogtopicID *string `json:"LogtopicID,omitempty" name:"LogtopicID"`
	// 0-异步投递，1-同步投递，目前仅支持同步投递

	Sync *int64 `json:"Sync,omitempty" name:"Sync"`
	// 描述信息

	Content *string `json:"Content,omitempty" name:"Content"`
	// 投递的域名或实例个数

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DescribeBotPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分布数量

		Points []*uint64 `json:"Points,omitempty" name:"Points"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotRequestCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// bot总数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotRequestCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRequestCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiRequestParameterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiRequestParameterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiRequestParameterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 是否是bot4.0版本

	IsVersionFour *bool `json:"IsVersionFour,omitempty" name:"IsVersionFour"`
	// 传入Bot版本号，场景化版本为"4.1.0"

	BotVersion *string `json:"BotVersion,omitempty" name:"BotVersion"`
}

func (r *ModifyBotStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAccessLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新接口此字段失效，默认返回空字符串

		Context *string `json:"Context,omitempty" name:"Context"`
		// 日志查询结果是否全部返回，其中，“true”表示结果返回，“false”表示结果为返回

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 返回的是否为分析结果，其中，“true”表示返回分析结果，“false”表示未返回分析结果

		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`
		// 如果Analysis为True，则返回分析结果的列名，否则为空
		// 注意：此字段可能返回 null，表示取不到有效值。

		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`
		// 日志查询结果；当Analysis为True时，可能返回为null
		// 注意：此字段可能返回 null，表示取不到有效值

		Results []*AccessLogInfo `json:"Results,omitempty" name:"Results"`
		// 日志分析结果；当Analysis为False时，可能返回为null
		// 注意：此字段可能返回 null，表示取不到有效值

		AnalysisResults []*AccessLogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchAccessLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebshellStatus struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// webshell开关，1：开。0：关。2：观察

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotSceneListRequest struct {
	*tchttp.BaseRequest

	// 通过场景名称模糊搜索

	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 是否只显示默认场景

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
	// 是否仅显示生效场景

	IsValid *bool `json:"IsValid,omitempty" name:"IsValid"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景模板类型，通过此下拉字段进行场景筛选。全部: all 登录: login  秒杀:seckill  爬内容：crawl 自定义: custom

	BusinessType []*string `json:"BusinessType,omitempty" name:"BusinessType"`
}

func (r *DescribeBotSceneListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSceneListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessExportRequest struct {
	*tchttp.BaseRequest

	// 日志导出ID

	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
	// 日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteAccessExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticTypesRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开始日期

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 截止日期

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeStatisticTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatisticTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserEditionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 购买了SaaS WAF则数组中包含sparta-waf。
		// 购买了CLB WAF则数据中包干clb-waf。
		// 如果两者都有购买则数组中都包含，否则没有购买WAF则为null。

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserEditionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserEditionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGlobalWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 匹配规则项列表

	Rules []*GlobalWhiteCond `json:"Rules,omitempty" name:"Rules"`
	// 规则序号

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 特征序号

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 规则状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyGlobalWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGlobalWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCustomRuleListData struct {

	// 规则列表

	List []*BatchCustomRuleListItem `json:"List,omitempty" name:"List"`
	// 列表总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type DescribeBotAccessLogTwoRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间

	StartTs *int64 `json:"StartTs,omitempty" name:"StartTs"`
	// 查看结束时间

	EndTs *int64 `json:"EndTs,omitempty" name:"EndTs"`
	// 显示第几页，默认传1

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 每页显示的日志数量，默认传10

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序方式

	LogSort *LogSort `json:"LogSort,omitempty" name:"LogSort"`
	// 是否查询日志条数，默认不查询

	CountFlag *bool `json:"CountFlag,omitempty" name:"CountFlag"`
	// 要查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询过滤器

	Filters []*BotDataFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBotAccessLogTwoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAccessLogTwoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainDetailsSaasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名详情

		DomainsPartInfo *DomainsPartInfo `json:"DomainsPartInfo,omitempty" name:"DomainsPartInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainDetailsSaasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainDetailsSaasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebshellStatusRequest struct {
	*tchttp.BaseRequest

	// 域名webshell状态

	Webshell *WebshellStatus `json:"Webshell,omitempty" name:"Webshell"`
}

func (r *ModifyWebshellStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebshellStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiSecExtractRule struct {

	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// api名称

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 请求方法列表

	Methods []*string `json:"Methods,omitempty" name:"Methods"`
	// 开关状态，0是关，1是开

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 正则匹配内容

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 10更新时间戳

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ApiSecSensitiveRule struct {

	// 身份证号，唯一主键

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 表示OS系统内置，"custom"表示客户自定义

	Source *string `json:"Source,omitempty" name:"Source"`
	// 开关状态，0：表示关，1表示开

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 风险等级，100，200,300表示低中高三个等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 修改时间，默认0，表示没有进行修改

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 自定义规则部分

	CustomRule *ApiSecCustomSensitiveRule `json:"CustomRule,omitempty" name:"CustomRule"`
}

type ExpiredFraudPkg struct {

	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 续费标志

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type SwitchElasticModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchElasticModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchElasticModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBlockPageRequest struct {
	*tchttp.BaseRequest

	// 拦截页面的Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteBlockPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBlockPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiSecurityFirstPurchaseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -1=error, 1=未防护，是第一次购买，2未购买waf实例，3已购买api安全

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// true:首购用户，false:非首购用户

		First *bool `json:"First,omitempty" name:"First"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiSecurityFirstPurchaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecurityFirstPurchaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotDetailCountTopNRequest struct {
	*tchttp.BaseRequest

	// 结束时间戳

	EndTs *int64 `json:"EndTs,omitempty" name:"EndTs"`
	// topn，默认传10

	TopN *int64 `json:"TopN,omitempty" name:"TopN"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 要出的表的类型，ip表还是token表

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 要搜索的值，不填的话，默认是两张表全有，传的话，出对应的Dimension值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 搜索框内容，冒号前面是key, 冒号是操作，值是最后一位，操作（冒号）默认是相等

	Filters []*BotDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 开始时间戳

	StartTs *int64 `json:"StartTs,omitempty" name:"StartTs"`
}

func (r *DescribeBotDetailCountTopNRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotDetailCountTopNRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistogramResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计数据

		Histogram []*string `json:"Histogram,omitempty" name:"Histogram"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistogramResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FirstPurchaseData struct {

	// 状态码，0表示成功，非0表示是比啊

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 是否首购用户，0表示非首购用户，1表示首购用户

	FirstPurchase *uint64 `json:"FirstPurchase,omitempty" name:"FirstPurchase"`
}

type DescribeExpiredInstancesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeExpiredInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExpiredInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainCustomRuleInfoItem struct {

	// 规则总数

	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
	// 已过期数

	ExpiredNum *uint64 `json:"ExpiredNum,omitempty" name:"ExpiredNum"`
	// 阻断数

	DenyNum *uint64 `json:"DenyNum,omitempty" name:"DenyNum"`
	// 人机识别数

	CaptchaNum *uint64 `json:"CaptchaNum,omitempty" name:"CaptchaNum"`
	// 观察数

	LogNum *uint64 `json:"LogNum,omitempty" name:"LogNum"`
	// 重定向数

	RedirectNum *uint64 `json:"RedirectNum,omitempty" name:"RedirectNum"`
}

type DestroyPostCKafkaFlowRequest struct {
	*tchttp.BaseRequest

	// 投递流的流ID

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
	// 1-访问日志，2-攻击日志，默认为访问日志。

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
}

func (r *DestroyPostCKafkaFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyPostCKafkaFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CountTopNItem struct {

	// 要显示的访问源,取值为ip活token的取值

	SrcValue *string `json:"SrcValue,omitempty" name:"SrcValue"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// SrcValue对应的类型

	Label *string `json:"Label,omitempty" name:"Label"`
}

type MainClassItem struct {

	// 主类id

	MainClassID *string `json:"MainClassID,omitempty" name:"MainClassID"`
	// 主类名字

	MainClassName *string `json:"MainClassName,omitempty" name:"MainClassName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 子类

	SubClass []*SubClassItem `json:"SubClass,omitempty" name:"SubClass"`
	// 当前主类的规则个数

	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`
}

type ModifyAttackWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAttackWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAttackWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InOutputBotUCBRule struct {

	// 入参ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// UCB的具体规则项

	Rule []*InOutputUCBRuleEntry `json:"Rule,omitempty" name:"Rule"`
	// 处置动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 规则开关

	OnOff *string `json:"OnOff,omitempty" name:"OnOff"`
	// 规则类型

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 规则优先级

	Prior *int64 `json:"Prior,omitempty" name:"Prior"`
	// 修改时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 生效时间

	ValidTime *int64 `json:"ValidTime,omitempty" name:"ValidTime"`
	// 传入的appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 额外参数

	AdditionArg *string `json:"AdditionArg,omitempty" name:"AdditionArg"`
	// 规则描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 标签

	Label *string `json:"Label,omitempty" name:"Label"`
	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// true-系统预设规则 false-自定义规则

	PreDefine *bool `json:"PreDefine,omitempty" name:"PreDefine"`
}

type DeleteAppletAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAppletAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppletAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebshellStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关状态

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// webshell域名

		Domain *string `json:"Domain,omitempty" name:"Domain"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebshellStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebshellStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainsCLSStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainsCLSStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainsCLSStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpiredBotPkg struct {

	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 续费标志

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 计费项

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeDomainInstanceIDsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		InstanceIDs []*InstanceIDName `json:"InstanceIDs,omitempty" name:"InstanceIDs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainInstanceIDsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainInstanceIDsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFirstPurchaseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true:首购用户，false:非首购用户

		First *bool `json:"First,omitempty" name:"First"`
		// 1：执行成功 0：执行失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFirstPurchaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFirstPurchaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleStatusRequest struct {
	*tchttp.BaseRequest

	// 要查询状态的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeModuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GoodNews struct {

	// 订单类型ID，用来唯一标识一个业务的一种场景（总共三种场景：新购、配置变更、续费）
	// 高级版: 102375(新购),102376(续费),102377(变配)
	// 企业版 : 102378(新购),102379(续费),102380(变配)
	// 旗舰版 : 102369(新购),102370(续费),102371(变配)
	// 域名包 : 102372(新购),102373(续费),102374(变配)
	// 业务扩展包 : 101040(新购),101041(续费),101042(变配)
	//
	// 高级版-CLB: 新购 101198  续费 101199 变配 101200
	// 企业版-CLB 101204(新购),101205(续费),101206(变配)
	// 旗舰版-CLB : 101201(新购),101202(续费),101203(变配)
	// 域名包-CLB: 101207(新购),101208(续费),101209(变配)
	// 业务扩展包-CLB: 101210(新购),101211(续费),101212(变配)
	//

	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`
	// 付费类型，1:预付费，0:后付费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 平台类型，默认1

	Platform *int64 `json:"Platform,omitempty" name:"Platform"`
	// 商品数量

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 商品明细

	GoodsDetail *GoodsDetailNew `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 购买waf实例区域ID
	// 1 表示购买大陆资源;
	// 9表示购买非中国大陆资源

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 默认为0

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeAccessDownloadRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载记录

		Data []*DonwloadRecordItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessDownloadRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessDownloadRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotRecordItemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Res []*BotRecordItem `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotRecordItemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRecordItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOpSignatureRuleRequest struct {
	*tchttp.BaseRequest

	// 要修改的规则id

	SignatureIDs []*string `json:"SignatureIDs,omitempty" name:"SignatureIDs"`
	// 要修改的状态 只能为0/1/2三种状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyOpSignatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOpSignatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadTask struct {

	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务状态：1-排队、2-导出中、3-导出完成、4-任务过期、5-导出失败

	State *int64 `json:"State,omitempty" name:"State"`
	// 下载进度: 0-100整数

	Process *int64 `json:"Process,omitempty" name:"Process"`
	// 下载链接

	Uri *string `json:"Uri,omitempty" name:"Uri"`
	// 创建时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 下载任务的指定域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DeleteRateLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 通用返回

		BaseInfo *RateLimitCommonRsp `json:"BaseInfo,omitempty" name:"BaseInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRateLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRateLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostInnerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 编辑的域名ID

		DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostInnerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostInnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessValueInfo struct {

	// 字段类型，目前支持的类型有：long、text、double

	Type *string `json:"Type,omitempty" name:"Type"`
	// 字段的分词符，只有当字段类型为text时才有意义；输入字符串中的每个字符代表一个分词符

	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`
	// 字段是否开启分析功能

	SqlFlag *bool `json:"SqlFlag,omitempty" name:"SqlFlag"`
	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。

	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type UpsertIpAccessControlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加或修改失败的条目

		FailedItems *string `json:"FailedItems,omitempty" name:"FailedItems"`
		// 添加或修改失败的数目

		FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`
		// 添加或修改的IP数据Id列表

		Ids []*string `json:"Ids,omitempty" name:"Ids"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpsertIpAccessControlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBotSceneRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DeleteBotSceneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBotSceneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTsResourceRequest struct {
	*tchttp.BaseRequest

	// 分片偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 页面显示个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持OpAppId, ClusterId, Ip，Status，Type；Status = 0/1

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTsResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTsResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiSecEventKeyValue struct {

	// 事件类型

	Key *string `json:"Key,omitempty" name:"Key"`
	// 数量

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DeleteAppletAccessRequest struct {
	*tchttp.BaseRequest

	// 小程序接入名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteAppletAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppletAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCiphersDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 加密套件信息

		Ciphers []*TLSCiphers `json:"Ciphers,omitempty" name:"Ciphers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCiphersDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCiphersDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCdcResourceRequest struct {
	*tchttp.BaseRequest

	// 序号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteCdcResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCdcResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCachePathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCachePathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCachePathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCachePathRspListItem struct {

	// ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type DescribeFirstPurchaseRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFirstPurchaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFirstPurchaseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSceneJsInjectRuleRequest struct {
	*tchttp.BaseRequest

	// 前端对抗规则对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询全局规则时为空，否则必填

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DescribeSceneJsInjectRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSceneJsInjectRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserSignatureRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 主类id

	MainClassID *string `json:"MainClassID,omitempty" name:"MainClassID"`
	// 主类开关0=关闭，1=开启，2=只告警

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 下发修改的规则列表

	RuleID []*ReqUserRule `json:"RuleID,omitempty" name:"RuleID"`
}

func (r *ModifyUserSignatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserSignatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotScoreEntry struct {

	// 起始分数

	Upper *string `json:"Upper,omitempty" name:"Upper"`
	// 结束分数

	Lower *string `json:"Lower,omitempty" name:"Lower"`
	// 分数段对应的动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 分数段对应的bot标签

	Label *string `json:"Label,omitempty" name:"Label"`
	// 重定向地址

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
}

type GoodsDetailNew struct {

	// 时间间隔

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 单位，支持购买d、m、y 即（日、月、年）

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品标签,。新购，续费必传，变配时放在oldConfig newConfig里面
	//
	// Saas 高级版 ：sp_wsm_waf_premium
	// Saas企业版 ：sp_wsm_waf_enterprise
	// Saas旗舰版 ：sp_wsm_waf_ultimate
	// Saas 业务扩展包：sp_wsm_waf_qpsep
	// Saas 域名扩展包：sp_wsm_waf_domain
	//
	// 高级版-CLB:sp_wsm_waf_premium_clb
	// 企业版-CLB : sp_wsm_waf_enterprise_clb
	// 旗舰版-CLB:sp_wsm_waf_ultimate_clb
	//  业务扩展包-CLB：sp_wsm_waf_qpsep_clb
	// 域名扩展包-CLB：sp_wsm_waf_domain_clb
	//

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 业务产品申请的pid（对应一个定价公式），通过pid计费查询到定价模型
	// 高级版 ：1000827
	// 企业版 ：1000830
	// 旗舰版 ：1000832
	// 域名包 : 1000834
	// 业务扩展包 : 1000481
	// 高级版-CLB:1001150
	// 企业版-CLB : 1001152
	// 旗舰版-CLB:1001154
	// 域名包-CLB: 1001156
	// 业务扩展包-CLB : 1001160
	//

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// waf产品码

	ProductInfo []*ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// waf实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// QPS数量

	ElasticQps *int64 `json:"ElasticQps,omitempty" name:"ElasticQps"`
	// 弹性账单

	FlexBill *int64 `json:"FlexBill,omitempty" name:"FlexBill"`
	// 1:自动续费，0:不自动续费

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// waf购买的实际地域信息

	RealRegion *int64 `json:"RealRegion,omitempty" name:"RealRegion"`
	// Waf实例对应的二级产品码

	Type *string `json:"Type,omitempty" name:"Type"`
	// 计费细项标签数组
	// Saas 高级版  sv_wsm_waf_package_premium
	// Saas 企业版  sv_wsm_waf_package_enterprise
	// Saas 旗舰版  sv_wsm_waf_package_ultimate
	// Saas 非中国大陆高级版  sv_wsm_waf_package_premium_intl
	// Saas 非中国大陆企业版   sv_wsm_waf_package_enterprise_intl
	// Saas 非中国大陆旗舰版  sv_wsm_waf_package_ultimate _intl
	// Saas 业务扩展包  sv_wsm_waf_qps_ep
	// Saas 域名扩展包  sv_wsm_waf_domain
	//
	// 高级版CLB   sv_wsm_waf_package_premium_clb
	// 企业版CLB  sv_wsm_waf_package_enterprise_clb
	// 旗舰版CLB   sv_wsm_waf_package_ultimate_clb
	// 非中国大陆高级版 CLB sv_wsm_waf_package_premium_clb_intl
	// 非中国大陆企业版CLB   sv_wsm_waf_package_premium_clb_intl
	// 非中国大陆旗舰版CLB  sv_wsm_waf_package_ultimate_clb _intl
	// 业务扩展包CLB sv_wsm_waf_qps_ep_clb
	// 域名扩展包CLB  sv_wsm_waf_domain_clb
	//

	LabelTypes []*string `json:"LabelTypes,omitempty" name:"LabelTypes"`
	// 计费细项标签数量，一般和SvLabelType一一对应

	LabelCounts []*int64 `json:"LabelCounts,omitempty" name:"LabelCounts"`
	// 变配使用，实例到期时间

	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
	// 对存在的实例购买bot 或api 安全

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// sv_wsm_waf_package_micro计费项

	WafPackageMicroClb *int64 `json:"WafPackageMicroClb,omitempty" name:"WafPackageMicroClb"`
	// sv_wsm_waf_lightpackage_lightpackage

	WafPackageClb *int64 `json:"WafPackageClb,omitempty" name:"WafPackageClb"`
	// sv_wsm_waf_package_premium_clb

	WafPackagePremiumClb *int64 `json:"WafPackagePremiumClb,omitempty" name:"WafPackagePremiumClb"`
	// sv_wsm_waf_package_enterprise_clb

	WafPackageEnterpriseClb *int64 `json:"WafPackageEnterpriseClb,omitempty" name:"WafPackageEnterpriseClb"`
	// sv_wsm_waf_package_premium

	WafPackagePremium *int64 `json:"WafPackagePremium,omitempty" name:"WafPackagePremium"`
	// sv_wsm_waf_package_enterprise

	WafPackageEnterprise *int64 `json:"WafPackageEnterprise,omitempty" name:"WafPackageEnterprise"`
	// sv_wsm_waf_package_ultimate

	WafPackageUltimate *int64 `json:"WafPackageUltimate,omitempty" name:"WafPackageUltimate"`
	// sv_wsm_waf_scene_bot_protection

	WafSceneBotProtection *int64 `json:"WafSceneBotProtection,omitempty" name:"WafSceneBotProtection"`
	// sv_wsm_waf_scene_bot_protection_clb

	WafSceneBotProtectionClb *int64 `json:"WafSceneBotProtectionClb,omitempty" name:"WafSceneBotProtectionClb"`
	// sv_wsm_waf_domain

	WafDomain *int64 `json:"WafDomain,omitempty" name:"WafDomain"`
	// sv_wsm_waf_qps_ep

	WafQps *int64 `json:"WafQps,omitempty" name:"WafQps"`
	// sv_wsm_waf_scls

	WafCls *int64 `json:"WafCls,omitempty" name:"WafCls"`
	// sv_wsm_waf_package_ultimate_clb

	WafPackageUltimateClb *string `json:"WafPackageUltimateClb,omitempty" name:"WafPackageUltimateClb"`
	// 老配置

	OldConfig *ModifyNewOldConfigNew `json:"OldConfig,omitempty" name:"OldConfig"`
	// 新配置

	NewConfig *ModifyNewOldConfigNew `json:"NewConfig,omitempty" name:"NewConfig"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type CreateAccessDownloadRecordRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Query *string `json:"Query,omitempty" name:"Query"`
	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// WAF版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 排序方式，asc和desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 下载名字，方便记忆

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateAccessDownloadRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessDownloadRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCCRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一般为null

		Data *string `json:"Data,omitempty" name:"Data"`
		// 操作的规则Id

		RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCCRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCCRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSignatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		Rules []*UserSignatureRule `json:"Rules,omitempty" name:"Rules"`
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserSignatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserSignatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceWafInstanceRequest struct {
	*tchttp.BaseRequest

	// 商品信息，json string

	ResInfo *string `json:"ResInfo,omitempty" name:"ResInfo"`
}

func (r *InquiryPriceWafInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceWafInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiSecCustomSensitiveRule struct {

	// 参数位置

	Position []*string `json:"Position,omitempty" name:"Position"`
	// 匹配条件

	MatchKey *string `json:"MatchKey,omitempty" name:"MatchKey"`
	// 匹配符号，当匹配条件为关键字匹配和字符匹配的时候传该值,可传多个

	MatchCond []*string `json:"MatchCond,omitempty" name:"MatchCond"`
	// 匹配值

	MatchValue []*string `json:"MatchValue,omitempty" name:"MatchValue"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
}

type RegionItem struct {

	// 一级区域

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 二级区域

	SubRegion *string `json:"SubRegion,omitempty" name:"SubRegion"`
	// 接入域名数

	DomainCount *uint64 `json:"DomainCount,omitempty" name:"DomainCount"`
	// 标号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CopyBotUCBPreinstallRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 类别：1表示协议特征，2表示情报协议特征，其他无效

	Category *uint64 `json:"Category,omitempty" name:"Category"`
	// 目标域名

	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *CopyBotUCBPreinstallRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBotUCBPreinstallRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIpAccessControlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除失败的条目

		FailedItems *string `json:"FailedItems,omitempty" name:"FailedItems"`
		// 删除失败的条目数

		FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIpAccessControlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCRuleRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Name *string `json:"Name,omitempty" name:"Name"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 页码

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 页的数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// clb-waf 或者 sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeCCRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpartaProtectionModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySpartaProtectionModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpartaProtectionModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotAiRule struct {

	// 需要加白的url

	Uri *string `json:"Uri,omitempty" name:"Uri"`
	// 状态开关

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type BotItemRequestBaseInfo struct {

	// url重复比

	ReqUrlRepeat *float64 `json:"ReqUrlRepeat,omitempty" name:"ReqUrlRepeat"`
	// 请求参数比

	ReqQueryRepeat *float64 `json:"ReqQueryRepeat,omitempty" name:"ReqQueryRepeat"`
	// Cookie是否滥用

	CookieAbuse *bool `json:"CookieAbuse,omitempty" name:"CookieAbuse"`
	// cookie存在性

	CookieExist *bool `json:"CookieExist,omitempty" name:"CookieExist"`
	// cookie重复性

	CookieRepeat *float64 `json:"CookieRepeat,omitempty" name:"CookieRepeat"`
	// ua类型

	UaType *string `json:"UaType,omitempty" name:"UaType"`
	// ua存在比

	UaExist *bool `json:"UaExist,omitempty" name:"UaExist"`
	// ua随机性指数

	UaKindRandomRate *float64 `json:"UaKindRandomRate,omitempty" name:"UaKindRandomRate"`
	// ua种类

	UaKindNums *int64 `json:"UaKindNums,omitempty" name:"UaKindNums"`
	// refer重复比

	ReferRepeat *float64 `json:"ReferRepeat,omitempty" name:"ReferRepeat"`
	// Refer存在性

	ReferExist *bool `json:"ReferExist,omitempty" name:"ReferExist"`
	// refer是否滥用

	ReferAbuse *bool `json:"ReferAbuse,omitempty" name:"ReferAbuse"`
	// 会话平均速度

	AvgSpeed *float64 `json:"AvgSpeed,omitempty" name:"AvgSpeed"`
	// 会话总次数

	Nums *int64 `json:"Nums,omitempty" name:"Nums"`
	// 是否存在Robots.txt

	HasRobots *bool `json:"HasRobots,omitempty" name:"HasRobots"`
	// 会话持续时间

	SessionDuration *float64 `json:"SessionDuration,omitempty" name:"SessionDuration"`
	// cookie有效率

	CookieValidRate *float64 `json:"CookieValidRate,omitempty" name:"CookieValidRate"`
	// 出现最多的cookie

	CookieMaxItem *string `json:"CookieMaxItem,omitempty" name:"CookieMaxItem"`
	// 出现最多的cookie比率

	CookieMaxItemRate *float64 `json:"CookieMaxItemRate,omitempty" name:"CookieMaxItemRate"`
	// 出现最多的ua

	UaMaxItem *string `json:"UaMaxItem,omitempty" name:"UaMaxItem"`
	// 出现最多的ua占的比例

	UaMaxItemRate *float64 `json:"UaMaxItemRate,omitempty" name:"UaMaxItemRate"`
	// ua有效比

	UaValidRate *float64 `json:"UaValidRate,omitempty" name:"UaValidRate"`
	// ua相似率

	UaSim *float64 `json:"UaSim,omitempty" name:"UaSim"`
	// refer有效比

	ReferValidRate *float64 `json:"ReferValidRate,omitempty" name:"ReferValidRate"`
	// 出现最多的refer

	ReferMaxItem *string `json:"ReferMaxItem,omitempty" name:"ReferMaxItem"`
	// 出现最多的refer的比例

	ReferMaxItemRate *float64 `json:"ReferMaxItemRate,omitempty" name:"ReferMaxItemRate"`
	// url种类

	ReqUrlKind *int64 `json:"ReqUrlKind,omitempty" name:"ReqUrlKind"`
	// url最小深度

	UrlMinLevel *int64 `json:"UrlMinLevel,omitempty" name:"UrlMinLevel"`
	// url最大深度

	UrlMaxLevel *int64 `json:"UrlMaxLevel,omitempty" name:"UrlMaxLevel"`
	// url平均深度

	UrlAvgLevel *float64 `json:"UrlAvgLevel,omitempty" name:"UrlAvgLevel"`
	// url数量

	UrlCount *int64 `json:"UrlCount,omitempty" name:"UrlCount"`
	// 请求参数种类

	QueryKind *int64 `json:"QueryKind,omitempty" name:"QueryKind"`
	// token命中时，对应的ip列表

	BotItemTokenIp []*BotItemTokenIp `json:"BotItemTokenIp,omitempty" name:"BotItemTokenIp"`
	// BotItemTokenIp 的数量

	BotItemTokenIpCount *int64 `json:"BotItemTokenIpCount,omitempty" name:"BotItemTokenIpCount"`
}

type DescribeCachePathRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 翻页

	PageInfo *PageInfo `json:"PageInfo,omitempty" name:"PageInfo"`
}

func (r *DescribeCachePathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCachePathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例名称，这里是模糊匹配

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 版本，目前值为sparta-waf和clb-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 一页的数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移数量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 实例搜索，这里是模糊匹配

	Instance *string `json:"Instance,omitempty" name:"Instance"`
}

func (r *DescribeMonitorInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDefaultEngineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// menshen, tiga

		DefaultEngine *string `json:"DefaultEngine,omitempty" name:"DefaultEngine"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserDefaultEngineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDefaultEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessPeriodRequest struct {
	*tchttp.BaseRequest

	// 访问日志保存期限，范围为[1, 180]

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 日志主题，新版本不需要再传

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 开启访问日志某些字段是否存储

	WriteConfig *FieldWriteConfig `json:"WriteConfig,omitempty" name:"WriteConfig"`
	// 为空则为全局开关，不为空则为单域名开关

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 单域名设置下大字段生效范围，1为当前域名，2为全部域名

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyAccessPeriodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessPeriodRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAntiFakeUrlStatusRequest struct {
	*tchttp.BaseRequest

	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// Id列表

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ModifyAntiFakeUrlStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAntiFakeUrlStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命中为情报白IP的列表

		Result []*IPType `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceEvent struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例事件类型，1：带宽事件,2:qps事件

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 实例事件等级：低:：1，中：2，高：3，沙箱：4

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 2022-10-11 00:00:00

	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`
	// 当天内预警或超限次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 带宽峰值

	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// qps规格；带宽规格=qps规格/40

	QPSConf *uint64 `json:"QPSConf,omitempty" name:"QPSConf"`
	// 实例类型：clb-waf或sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// qps峰值

	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`
	// 实例购买日期

	InstanceTime *string `json:"InstanceTime,omitempty" name:"InstanceTime"`
	// 沙箱阈值

	SandboxValue *uint64 `json:"SandboxValue,omitempty" name:"SandboxValue"`
	// 事件日期

	DayTime *string `json:"DayTime,omitempty" name:"DayTime"`
}

type PriceData struct {

	// 是否有用户价格

	HasUserPrice *bool `json:"HasUserPrice,omitempty" name:"HasUserPrice"`
	// 价格

	Price *int64 `json:"Price,omitempty" name:"Price"`
	// 部件价格明细

	PartDetail []*PartDetail `json:"PartDetail,omitempty" name:"PartDetail"`
	// 总价

	TotalCost *int64 `json:"TotalCost,omitempty" name:"TotalCost"`
	// 金额单位

	Currency *string `json:"Currency,omitempty" name:"Currency"`
	// 商品的时间单位

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 商品的时间大小

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 资源实例个数

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 来源类型

	FromType *int64 `json:"FromType,omitempty" name:"FromType"`
	// 计费产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 计费二级产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费策略id

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 优惠后总费用，单位分

	RealTotalCost *int64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 价格单位，pent：分

	AmountUnit *string `json:"AmountUnit,omitempty" name:"AmountUnit"`
	// 高精度价格 主要用于国际站部分定价低于一分钱的产品购买页展示高精度价格（最大保留到小数点后八位）

	HighPrecisionPrice *HighPrecisionPrice `json:"HighPrecisionPrice,omitempty" name:"HighPrecisionPrice"`
	// 总的折扣，100表示100%不打折

	Policy *float64 `json:"Policy,omitempty" name:"Policy"`
	// 优惠明细

	PolicyDetail *PolicyDetail `json:"PolicyDetail,omitempty" name:"PolicyDetail"`
	// 单价

	UnitPrice *int64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
}

type DescribeTopAttackDomainRequest struct {
	*tchttp.BaseRequest

	// WAF实例ID，不传则不过滤

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 查询起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 查询结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// TOP N,可从0-10选择，默认是10

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeTopAttackDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopAttackDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiParameterType struct {

	// 参数名称

	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`
	// 参数类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 参数位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 数据标签(敏感字段)

	Label []*string `json:"Label,omitempty" name:"Label"`
	// 时间戳

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 来源是请求或者响应

	Source *string `json:"Source,omitempty" name:"Source"`
}

type DeleteBatchCustomRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBatchCustomRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryConfigDeliveryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// {"CheckAppId":251240571,"Domain":"config.qcloudwaf.com","Module":"areaban","DeliveryState":2,"DeliveryStateName":"下发成功","DeliveryMsg":"","DbVersion":21,"DbVersionTs":1686755452975,"DbVersionTsStr":"2023-06-14 23:10:52","RedisList":[{"Region":"cd","Version":21}],"InstanceList":[{"Ip":"30.188.190.224","Version":21,"Region":"cd","InstanceType":"container","EnginType":"nginx_lua"},{"Ip":"30.188.186.106","Version":21,"Region":"cd","InstanceType":"container","EnginType":"nginx_lua"}]}

		Data *ConfigDeliveryData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryConfigDeliveryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryConfigDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotsGeoStatStyledItem struct {

	// 地域

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeBotStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 包装参数

		Data *DomainBotStatusData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBotSceneActionRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBotSceneActionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBotSceneActionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainCls struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// clb开启状态

	Cls *uint64 `json:"Cls,omitempty" name:"Cls"`
	// waf开关

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 对外clbwaf域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

type DescribeAreaBanSupportAreasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域封禁的地域列表，要解析成json后使用

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAreaBanSupportAreasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreaBanSupportAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchAreaBanListItem struct {

	// 规则Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 数据来源 custom-自定义(默认)、batch-批量防护

	Source *string `json:"Source,omitempty" name:"Source"`
	// 地域列表

	Areas []*string `json:"Areas,omitempty" name:"Areas"`
	// 域名列表

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DeleteAccessDownloadRecordRequest struct {
	*tchttp.BaseRequest

	// 下载任务记录唯一标记

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteAccessDownloadRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessDownloadRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data *AllDomainQueryResult `json:"Data,omitempty" name:"Data"`
		// 域名总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiSecurityFirstPurchaseRequest struct {
	*tchttp.BaseRequest

	// 在该实例下购买api安全资格

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeApiSecurityFirstPurchaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecurityFirstPurchaseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Bot攻击总数

		BotCount *uint64 `json:"BotCount,omitempty" name:"BotCount"`
		// 访问请求总数周比

		AccessCircleCount *int64 `json:"AccessCircleCount,omitempty" name:"AccessCircleCount"`
		// 访问控制总数周比

		ACLCircleCount *int64 `json:"ACLCircleCount,omitempty" name:"ACLCircleCount"`
		// api资产总数周比

		ApiAssetsCircleCount *int64 `json:"ApiAssetsCircleCount,omitempty" name:"ApiAssetsCircleCount"`
		// 黑名单总数周比

		IPBlackCircleCount *int64 `json:"IPBlackCircleCount,omitempty" name:"IPBlackCircleCount"`
		// 信息泄露总数周比

		LeakCircleCount *int64 `json:"LeakCircleCount,omitempty" name:"LeakCircleCount"`
		// 防篡改总数

		TamperCount *uint64 `json:"TamperCount,omitempty" name:"TamperCount"`
		// 信息泄露总数

		LeakCount *uint64 `json:"LeakCount,omitempty" name:"LeakCount"`
		// Bot攻击总数周比

		BotCircleCount *int64 `json:"BotCircleCount,omitempty" name:"BotCircleCount"`
		// 防篡改总数周比

		TamperCircleCount *int64 `json:"TamperCircleCount,omitempty" name:"TamperCircleCount"`
		// CC攻击总数

		CCCount *uint64 `json:"CCCount,omitempty" name:"CCCount"`
		// 黑名单总数

		IPBlackCount *uint64 `json:"IPBlackCount,omitempty" name:"IPBlackCount"`
		// CC攻击总数周比

		CCCircleCount *int64 `json:"CCCircleCount,omitempty" name:"CCCircleCount"`
		// 访问请求总数

		AccessCount *uint64 `json:"AccessCount,omitempty" name:"AccessCount"`
		// Web攻击总数

		AttackCount *uint64 `json:"AttackCount,omitempty" name:"AttackCount"`
		// 访问控制总数

		ACLCount *uint64 `json:"ACLCount,omitempty" name:"ACLCount"`
		// api资产总数

		ApiAssetsCount *uint64 `json:"ApiAssetsCount,omitempty" name:"ApiAssetsCount"`
		// api风险事件数量

		ApiRiskEventCount *uint64 `json:"ApiRiskEventCount,omitempty" name:"ApiRiskEventCount"`
		// Web攻击总数周比

		AttackCircleCount *int64 `json:"AttackCircleCount,omitempty" name:"AttackCircleCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// instance事件列表

		InstanceEvents []*InstanceEvent `json:"InstanceEvents,omitempty" name:"InstanceEvents"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加成功的规则ID

		RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCustomWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpsertBotUCBFeatureRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则内容

	Rule *string `json:"Rule,omitempty" name:"Rule"`
}

func (r *UpsertBotUCBFeatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertBotUCBFeatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppletAuthorizeV3Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 小程序授权信息列表

		AppletAuthorizes []*AppletAuthInfoV3 `json:"AppletAuthorizes,omitempty" name:"AppletAuthorizes"`
		// 满足过滤条件的接入小程序的总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppletAuthorizeV3Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppletAuthorizeV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBatchCustomRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 批量Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 开关状态 0-关、1-开
	//

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyBatchCustomRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBatchCustomRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 白名单ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 白名单内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 白名单的备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 白名单的开关状态，0：关闭，1：开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 白名单的类型，jsinject_req：动态风控的请求白名单，jsinject_rsp：动态风控的响应白名单

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostDel struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 实例类型

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type ModifyBotAutoStatisticRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开关状态，默认为关

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 运行模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 智能统计用户自定义配置信息

	AutoStatRule *BotAutoStatRule `json:"AutoStatRule,omitempty" name:"AutoStatRule"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 版本号，场景化版本为4.1.0

	BotVersion *string `json:"BotVersion,omitempty" name:"BotVersion"`
}

func (r *ModifyBotAutoStatisticRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotAutoStatisticRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGenerateDealsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费下单响应结构体

		Data *DealData `json:"Data,omitempty" name:"Data"`
		// 1:成功，0:失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 返回message

		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGenerateDealsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGenerateDealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHWDomainStateRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *QueryHWDomainStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryHWDomainStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessFullTextInfo struct {

	// 是否大小写敏感

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 全文索引的分词符，字符串中每个字符代表一个分词符

	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`
	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。

	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type DescribePeakPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据点

		Points []*PeakPointsItem `json:"Points,omitempty" name:"Points"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePeakPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePeakPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDnsRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDnsRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDnsRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// bot详情

		Data []*BotDetail `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesRequest struct {
	*tchttp.BaseRequest

	// clb-waf或者sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 分页参数

	Paging *DescribeCustomRulesPagingInfo `json:"Paging,omitempty" name:"Paging"`
	// 过滤参数：动作类型：0放行，1阻断，2人机识别，3观察，4重定向

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 过滤参数：规则名称过滤条件

	Search *string `json:"Search,omitempty" name:"Search"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeCustomRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 匹配规则项列表

	Rules []*UserWhiteRuleItem `json:"Rules,omitempty" name:"Rules"`
	// 规则序号

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则Id

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
}

func (r *ModifyAttackWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAttackWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// api开通的数量

		SwitchDomainCount *int64 `json:"SwitchDomainCount,omitempty" name:"SwitchDomainCount"`
		// 识别出的场景列表

		SceneList []*string `json:"SceneList,omitempty" name:"SceneList"`
		// 风险api数量

		RiskCount *uint64 `json:"RiskCount,omitempty" name:"RiskCount"`
		// 事件总数

		EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
		// 已确认api

		ConfirmedCount *uint64 `json:"ConfirmedCount,omitempty" name:"ConfirmedCount"`
		// 新发现api数量

		NewCount *uint64 `json:"NewCount,omitempty" name:"NewCount"`
		// api总数,第一个元素表示当前周期，第二个元素表示上一周期，下同

		ApiCount []*uint64 `json:"ApiCount,omitempty" name:"ApiCount"`
		// 过去7日活跃api数量

		ActiveCount []*uint64 `json:"ActiveCount,omitempty" name:"ActiveCount"`
		// 确认中api

		ConfirmingCount *uint64 `json:"ConfirmingCount,omitempty" name:"ConfirmingCount"`
		// 已忽略api

		IgnoredCount *uint64 `json:"IgnoredCount,omitempty" name:"IgnoredCount"`
		// 敏感api数量

		SensitiveCount []*uint64 `json:"SensitiveCount,omitempty" name:"SensitiveCount"`
		// 过去7日不活跃api数量

		NoActiveCount []*uint64 `json:"NoActiveCount,omitempty" name:"NoActiveCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllGenerateDealsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费下单响应结构体

		Data *DealData `json:"Data,omitempty" name:"Data"`
		// 1:成功，0:失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 返回message

		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAllGenerateDealsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAllGenerateDealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGlobalWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则序号

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteGlobalWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGlobalWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiFakeRulesRequest struct {
	*tchttp.BaseRequest

	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// asc或者desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 目前支持根据ts排序

	By *string `json:"By,omitempty" name:"By"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAntiFakeRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAntiFakeRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpUserSignaturePolicyRequest struct {
	*tchttp.BaseRequest

	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持 AppId，Domain

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// 排序字段，仅支持 level

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeOpUserSignaturePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpUserSignaturePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVIPStateRequest struct {
	*tchttp.BaseRequest

	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Filters []*VIPStateFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVIPStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVIPStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackAnalysisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 聚合分析条目详情

		Data []*LogAnalysisInfo `json:"Data,omitempty" name:"Data"`
		// Data数组大小

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttackAnalysisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAccessStatPointsRequest struct {
	*tchttp.BaseRequest

	// 需要查询TOP的维度名，见tapd文档

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 查询过滤器

	Filters []*BotDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 开始时间

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 需要的top数量

	TopN *uint64 `json:"TopN,omitempty" name:"TopN"`
	// "asc"-升序  "desc"-降序

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotAccessStatPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAccessStatPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopStraceIPRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 域名，不指定则为全部域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 防护类型，不指定则默认为business_risk

	StraceType *string `json:"StraceType,omitempty" name:"StraceType"`
}

func (r *GetTopStraceIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopStraceIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpiredBotQPS struct {

	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 续费标志

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type ModifyDomainIpv6StatusRequest struct {
	*tchttp.BaseRequest

	// 需要修改的域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 修改域名的Ipv6开关为Status （1:开启 2:关闭）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 需要修改的域名所属的实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 需要修改的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ModifyDomainIpv6StatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainIpv6StatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotPointsRequest struct {
	*tchttp.BaseRequest

	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// WAF版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeBotPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrialQRCodeV3Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 二维码

		QRCode *string `json:"QRCode,omitempty" name:"QRCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrialQRCodeV3Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrialQRCodeV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyBlockPageRequest struct {
	*tchttp.BaseRequest

	// 拦截页面的ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 域名的版本信息，可取的值sparta-waf、clb-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ApplyBlockPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyBlockPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotLabelRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiSecEventChangeRequest struct {
	*tchttp.BaseRequest

	// 批量删除的事件列表

	EventIdList []*string `json:"EventIdList,omitempty" name:"EventIdList"`
	// 批量删除的Api列表

	ApiNameList []*ApiSecKey `json:"ApiNameList,omitempty" name:"ApiNameList"`
	// 判断是否删除，包括删除事件和删除资产

	IsDelete *bool `json:"IsDelete,omitempty" name:"IsDelete"`
	// 变更状态，1:新发现，2，确认中，3，已确认，4，已下线，5，已忽略

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 处理人

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 备注，有长度显示1k

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyApiSecEventChangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiSecEventChangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDomainRulesRequest struct {
	*tchttp.BaseRequest

	// 规则列表

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 开关状态，0表示关闭，1表示开启，2表示只观察

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 设置为观察模式原因，
	// 1表示业务自身原因观察，2表示系统规则误报上报，3表示核心业务灰度观察，4表示其他

	Reason *uint64 `json:"Reason,omitempty" name:"Reason"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *SwitchDomainRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDomainRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionedIpRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribeActionedIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeActionedIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrialQRCodeV3Request struct {
	*tchttp.BaseRequest

	// 小程序ID

	AppletId *string `json:"AppletId,omitempty" name:"AppletId"`
}

func (r *DescribeTrialQRCodeV3Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrialQRCodeV3Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllQpsConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 套餐类型弹性qps相关值

		Data *QpsConfig `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAllQpsConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllQpsConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询到的API规则的列表

		RuleList []*QueryListItems `json:"RuleList,omitempty" name:"RuleList"`
		// 规则的总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClsPackage struct {

	// 资源ID

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 是否自动续费，1：自动续费，0：不自动续费

	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
	// 套餐个数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 地域，目前在clb-waf中没有用到

	Region *string `json:"Region,omitempty" name:"Region"`
}

type PartDetailItem struct {

	// 计费策略ID

	Pid *string `json:"Pid,omitempty" name:"Pid"`
	// 价格

	Price *int64 `json:"Price,omitempty" name:"Price"`
	// 计费项数量

	Value *int64 `json:"Value,omitempty" name:"Value"`
	// 价格项明细

	PriceDetail []*PriceDetail `json:"PriceDetail,omitempty" name:"PriceDetail"`
	// 计费项CODE

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项二级CODE

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 总价

	TotalCost *int64 `json:"TotalCost,omitempty" name:"TotalCost"`
	// 真实总价

	RealTotalCost *int64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 总的折扣，100表示100%不打折

	Policy *float64 `json:"Policy,omitempty" name:"Policy"`
	// 折扣明细信息

	PolicyDetail *PolicyDetail `json:"PolicyDetail,omitempty" name:"PolicyDetail"`
}

type DescribeBatchAreaBanListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据信息

		Data *BatchAreaBanListData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBatchAreaBanListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBatchAreaBanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGenerateDealsRequest struct {
	*tchttp.BaseRequest

	// 计费下单入参

	Goods []*Goods `json:"Goods,omitempty" name:"Goods"`
}

func (r *ModifyGenerateDealsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGenerateDealsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTsRegionRequest struct {
	*tchttp.BaseRequest

	// 容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 一级区域

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 二级区域

	SubRegion *string `json:"SubRegion,omitempty" name:"SubRegion"`
	// 偏移,大于0的整数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTsRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTsRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAntiInfoLeakRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteAntiInfoLeakRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAntiInfoLeakRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotTiRuleRequest struct {
	*tchttp.BaseRequest

	// 版本号，场景化版本为4.1.0

	BotVersion *string `json:"BotVersion,omitempty" name:"BotVersion"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名的InstanceID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DescribeBotTiRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotTiRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTsResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 资源信息列表

		List []*TsResource `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTsResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTsResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiAnalyzeStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关状态

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiAnalyzeStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiAnalyzeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubClassItem struct {

	// 子类id

	SubClassID *string `json:"SubClassID,omitempty" name:"SubClassID"`
	// 子类名字

	SubClassName *string `json:"SubClassName,omitempty" name:"SubClassName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeRuleLimitRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeRuleLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceIdListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例总个数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 实例id信息列表

		Result []*InstanceIDInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceIdListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceIdListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpartaProtectionInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// upstreamType取值

		UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`
		// upstreamScheme取值

		UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`
		// port信息

		Ports []*PortItem `json:"Ports,omitempty" name:"Ports"`
		// 防御等级,100,200,300

		Level *string `json:"Level,omitempty" name:"Level"`
		// 证书

		Cert *string `json:"Cert,omitempty" name:"Cert"`
		// 私有密钥

		PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
		// 灰度区域列表

		GrayAreas []*string `json:"GrayAreas,omitempty" name:"GrayAreas"`
		// 是否含有websocket

		IsWebsocket *string `json:"IsWebsocket,omitempty" name:"IsWebsocket"`
		// loadBalance信息

		LoadBalance *string `json:"LoadBalance,omitempty" name:"LoadBalance"`
		// httpsUpstreamPort取值

		HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`
		// 是否灰度

		IsGray *string `json:"IsGray,omitempty" name:"IsGray"`
		// cname取值

		Cname *string `json:"Cname,omitempty" name:"Cname"`
		// 状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// ssl的id

		Sslid *string `json:"Sslid,omitempty" name:"Sslid"`
		// 与源站是否保持长连接

		IsKeepAlive *string `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`
		// 0：BGP 1：Anycast

		Anycast *string `json:"Anycast,omitempty" name:"Anycast"`
		// 源IP地址列表

		SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`
		// 证书类型

		CertType *string `json:"CertType,omitempty" name:"CertType"`
		// 引擎

		Engine *string `json:"Engine,omitempty" name:"Engine"`
		// HTTPS重写

		HttpsRewrite *string `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`
		// upstreamDomain取值

		UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`
		// 是否是HTTP2

		IsHttp2 *string `json:"IsHttp2,omitempty" name:"IsHttp2"`
		// 模式

		Mode *string `json:"Mode,omitempty" name:"Mode"`
		// 域名

		Domain *string `json:"Domain,omitempty" name:"Domain"`
		// 域名ID

		DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
		// 是否是cdn

		IsCdn *string `json:"IsCdn,omitempty" name:"IsCdn"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpartaProtectionInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpartaProtectionInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificatedDomainRequest struct {
	*tchttp.BaseRequest

	// 证书私钥

	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitempty" name:"CertificatePrivateKey"`
	// 证书公钥

	CertificatePublicKey *string `json:"CertificatePublicKey,omitempty" name:"CertificatePublicKey"`
	// 需要修改的域名列表

	DomainArray []*string `json:"DomainArray,omitempty" name:"DomainArray"`
	// ssl证书中的域名

	CertDomain *string `json:"CertDomain,omitempty" name:"CertDomain"`
	// 证书的CertificateId

	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *ModifyCertificatedDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCertificatedDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WhiteListItem struct {

	// 白名单对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 白名单ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 白名单内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 白名单的备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 白名单的开关状态，0：关闭，1：开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 白名单的类型，jsinject_req：动态风控的请求白名单，jsinject_rsp：动态风控的响应白名单

	Type *string `json:"Type,omitempty" name:"Type"`
}

type RuleUpgradeRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则库id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *RuleUpgradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleUpgradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePostCKafkaFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePostCKafkaFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePostCKafkaFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MapItem struct {

	// 城市

	Country *string `json:"Country,omitempty" name:"Country"`
	// 数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type ModifyAIModelRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 训练方式

	Method *uint64 `json:"Method,omitempty" name:"Method"`
}

func (r *ModifyAIModelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAIModelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleCosKeyRequest struct {
	*tchttp.BaseRequest

	// 可选参数，用于指定将文件上传到Cos的Key，留空则由后台随机生成

	Key *string `json:"Key,omitempty" name:"Key"`
	// 前端页面url地址，用于cos跨域上传，传空的话后台尝试从http header的Origin字段获取

	Origin *string `json:"Origin,omitempty" name:"Origin"`
	// 登录人

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *GetRuleCosKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleCosKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVIPConfigsRequest struct {
	*tchttp.BaseRequest

	// 无

	DstGroupIds []*int64 `json:"DstGroupIds,omitempty" name:"DstGroupIds"`
	// 无

	IsRetry *int64 `json:"IsRetry,omitempty" name:"IsRetry"`
	// 无

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
	// 无

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 无

	VIP *string `json:"VIP,omitempty" name:"VIP"`
	// 无

	Type *string `json:"Type,omitempty" name:"Type"`
	// 无

	SrcGroupIds []*int64 `json:"SrcGroupIds,omitempty" name:"SrcGroupIds"`
	// 无

	SrcRSIPs []*string `json:"SrcRSIPs,omitempty" name:"SrcRSIPs"`
}

func (r *ModifyVIPConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVIPConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotIdStat struct {

	// 模式：观察/拦截/自定义

	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`
	// 规则总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 配置观察的规则数

	MonitorCount *int64 `json:"MonitorCount,omitempty" name:"MonitorCount"`
	// 配置拦截的规则数

	InterceptCount *int64 `json:"InterceptCount,omitempty" name:"InterceptCount"`
}

type ModifyProtectionLevelRequest struct {
	*tchttp.BaseRequest

	// 客户域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 防护等级,100,200,300

	Level *int64 `json:"Level,omitempty" name:"Level"`
}

func (r *ModifyProtectionLevelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProtectionLevelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostFlowModeInnerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功的状态码

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostFlowModeInnerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostFlowModeInnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则类型描述列表

		Types []*Type `json:"Types,omitempty" name:"Types"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackPointsRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 域名，不指定则为全部域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询的攻击类型，全部攻击类型不指定

	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

func (r *GetAttackPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SessionItem struct {

	// 匹配类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// 起始模式

	KeyOrStartMat *string `json:"KeyOrStartMat,omitempty" name:"KeyOrStartMat"`
	// 结束模式

	EndMat *string `json:"EndMat,omitempty" name:"EndMat"`
	// 起始偏移

	StartOffset *string `json:"StartOffset,omitempty" name:"StartOffset"`
	// 结束偏移

	EndOffset *string `json:"EndOffset,omitempty" name:"EndOffset"`
	// 数据源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 更新时间戳

	TsVersion *string `json:"TsVersion,omitempty" name:"TsVersion"`
	// SessionID

	SessionId *int64 `json:"SessionId,omitempty" name:"SessionId"`
	// Session名

	SessionName *string `json:"SessionName,omitempty" name:"SessionName"`
	// Session是否正在被启用

	SessionInUsed *bool `json:"SessionInUsed,omitempty" name:"SessionInUsed"`
	// Session关联的CC规则ID

	RelatedRuleID []*int64 `json:"RelatedRuleID,omitempty" name:"RelatedRuleID"`
}

type GenerateDealsAndPayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费下单响应结构体

		Data *DealData `json:"Data,omitempty" name:"Data"`
		// 1:成功，0:失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 返回message

		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateDealsAndPayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateDealsAndPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSpartaWafRuleReqStrategy struct {

	// 匹配字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 匹配参数

	CompareFunc *string `json:"CompareFunc,omitempty" name:"CompareFunc"`
	// 逻辑符号

	Content *string `json:"Content,omitempty" name:"Content"`
	// 匹配内容

	Arg *string `json:"Arg,omitempty" name:"Arg"`
}

type BotDataFilter struct {

	// 查询维度

	Entity *string `json:"Entity,omitempty" name:"Entity"`
	// 操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 操作值，多个值用

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeCustomPayloadsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内容详情

		Resources []*string `json:"Resources,omitempty" name:"Resources"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomPayloadsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomPayloadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificatedDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Code Message

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCertificatedDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCertificatedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LimitRuleStatusUnit struct {

	// 规则ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否启用，0表示不启用，1表示启用

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeAccessDownloadRecordsRequest struct {
	*tchttp.BaseRequest

	// WAF版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeAccessDownloadRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessDownloadRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotRegionsStatRequest struct {
	*tchttp.BaseRequest

	// 地域范围，china表示中国，global表示全球，其他值无效

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// bot类型，可以是ALL，UB，TCB，UCB，其他值无效

	BotType *string `json:"BotType,omitempty" name:"BotType"`
	// 起始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 终止时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotRegionsStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRegionsStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotSceneUCBRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则内容, 增加编码SceneId信息,1.BOT全局白名单处调用时，SceneId为"global", RuleType传10, Action为"permit";2.BOT场景配置时，SceneId为场景ID

	Rule *InOutputBotUCBRule `json:"Rule,omitempty" name:"Rule"`
	// 1.BOT全局白名单处调用时，传"global";2.BOT场景配置时，传具体的场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *ModifyBotSceneUCBRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSceneUCBRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PortInfo struct {

	// Nginx的服务器id

	NginxServerId *uint64 `json:"NginxServerId,omitempty" name:"NginxServerId"`
	// 监听端口配置

	Port *string `json:"Port,omitempty" name:"Port"`
	// 与端口对应的协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 回源端口

	UpstreamPort *string `json:"UpstreamPort,omitempty" name:"UpstreamPort"`
	// 回源协议

	UpstreamProtocol *string `json:"UpstreamProtocol,omitempty" name:"UpstreamProtocol"`
}

type DescribeBotAggregateTopNResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// topN结果

		Data []*BotTopItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotAggregateTopNResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAggregateTopNResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogAnalysisInfo struct {

	// 日志条目数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 百分比

	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
	// Key对应的分类值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeCCAutoStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeCCAutoStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCAutoStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductInfo struct {

	// 产品名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateHybridClusterRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 集群名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 混合云集群地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

func (r *CreateHybridClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHybridClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCLSRequest struct {
	*tchttp.BaseRequest

	// WAF版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeCLSRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCLSRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWafAutoDenyRulesRequest struct {
	*tchttp.BaseRequest

	// IP封禁统计时间，范围为1-60分钟

	TimeThreshold *int64 `json:"TimeThreshold,omitempty" name:"TimeThreshold"`
	// 触发IP封禁后的封禁时间，范围为5~360分钟

	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitempty" name:"DenyTimeThreshold"`
	// 自动封禁状态，0表示关闭，1表示打开

	DefenseStatus *int64 `json:"DefenseStatus,omitempty" name:"DefenseStatus"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 触发IP封禁的攻击次数阈值，范围为2~100次

	AttackThreshold *int64 `json:"AttackThreshold,omitempty" name:"AttackThreshold"`
}

func (r *ModifyWafAutoDenyRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWafAutoDenyRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotStatPointsRequest struct {
	*tchttp.BaseRequest

	// 查询过滤器

	Filters []*BotDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 开始时间

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 需要的top数量

	TopN *uint64 `json:"TopN,omitempty" name:"TopN"`
	// 默认该接口获取top5的元素的趋势图，如果填充该值则表示取Dimension指定维度的PresetTargets的趋势图

	PresetTargets *string `json:"PresetTargets,omitempty" name:"PresetTargets"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 需要查询TOP的维度名，见AggregateDimension的key定义

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

func (r *DescribeBotStatPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotStatPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditOpRuleUpdateLogRequest struct {
	*tchttp.BaseRequest

	// 要编辑的日志id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// version

	LogVersion *string `json:"LogVersion,omitempty" name:"LogVersion"`
}

func (r *EditOpRuleUpdateLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditOpRuleUpdateLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJsInjectRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyJsInjectRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJsInjectRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTlsVersionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTlsVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTlsVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDomainCustomRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Data *DomainCustomRuleInfoData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDomainCustomRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDomainCustomRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddBlockPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddBlockPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBlockPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainRulesRequest struct {
	*tchttp.BaseRequest

	// 需要查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppletAuthorizeV3Request struct {
	*tchttp.BaseRequest

	// 小程序认证码

	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`
}

func (r *CreateAppletAuthorizeV3Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppletAuthorizeV3Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotUCBPreinstallRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data *BotUCBPreinstallRuleData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotUCBPreinstallRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotUCBPreinstallRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHybridClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHybridClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHybridClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebshellStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功的状态码，需要JSON解码后再使用，返回的格式是{"域名":"状态"}，成功的状态码为Success，其它的为失败的状态码（yunapi定义的错误码）

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWebshellStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebshellStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotDetailDownloadTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务列表

		Tasks []*DownloadTask `json:"Tasks,omitempty" name:"Tasks"`
		// 任务数量

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotDetailDownloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotDetailDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAntiInfoLeakRulesRequest struct {
	*tchttp.BaseRequest

	// 网址

	Uri *string `json:"Uri,omitempty" name:"Uri"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 动作，0（告警）、1（替换）、2（仅显示前四位）、3（仅显示后四位）、4（阻断）

	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`
	// 策略详情

	Strategies []*StrategyForAntiInfoLeak `json:"Strategies,omitempty" name:"Strategies"`
}

func (r *AddAntiInfoLeakRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAntiInfoLeakRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PiechartItem struct {

	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type ActionRuleInfo struct {

	// 动作策略ID

	ActionRuleId *string `json:"ActionRuleId,omitempty" name:"ActionRuleId"`
	// 动作策略名称

	ActionRuleName *string `json:"ActionRuleName,omitempty" name:"ActionRuleName"`
}

type BotsDomainAggStatStyledItem struct {

	// 域名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 攻击次数

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeInstanceEventsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotStatSt struct {

	// 会话数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 平均速度

	AvgSpeed *float64 `json:"AvgSpeed,omitempty" name:"AvgSpeed"`
}

type CopyBotUCBFeatureRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Data *CopyBotsUCBFeatureRuleRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyBotUCBFeatureRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBotUCBFeatureRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 标签list

		Label []*string `json:"Label,omitempty" name:"Label"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量趋势数据

		Data []*BotStatPointItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessFastAnalysisRequest struct {
	*tchttp.BaseRequest

	// 排序字段,升序asc,降序desc，默认降序desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 返回的top数，默认返回top5

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 查询语句，语句长度最大为4096，由于本接口是分析接口，如果无过滤条件，必须传 * 表示匹配所有，参考CLS的分析统计语句的文档

	Query *string `json:"Query,omitempty" name:"Query"`
	// 需要分析统计的字段名

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
}

func (r *DescribeAccessFastAnalysisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessFastAnalysisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiAggregateTopNResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// topN结果

		Data []*BotTopItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiAggregateTopNResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiAggregateTopNResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiSecEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表

		Data []*ApiEvent `json:"Data,omitempty" name:"Data"`
		// 事件总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiSecEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PostDomainsConfigTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PostDomainsConfigTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PostDomainsConfigTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiAsset struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 请求方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// api名称

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 场景

	Scene *string `json:"Scene,omitempty" name:"Scene"`
	// 数据标签

	Label []*string `json:"Label,omitempty" name:"Label"`
	// 过去7天是否活跃

	Active *bool `json:"Active,omitempty" name:"Active"`
	// 最近更新时间

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// api发现时间

	InsertTime *uint64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 资产状态，1:新发现，2，确认中，3，已确认，4，已下线，5，已忽略

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 风险等级，100,200,300对应低中高

	Level *string `json:"Level,omitempty" name:"Level"`
}

type DeleteCustomPayloadsRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 删除的payload的id数组

	Items []*string `json:"Items,omitempty" name:"Items"`
}

func (r *DeleteCustomPayloadsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomPayloadsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackWorldMapRequest struct {
	*tchttp.BaseRequest

	// WAF实例ID，不传则不过滤

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 兼容Host，逐步淘汰Host字段

	Host *string `json:"Host,omitempty" name:"Host"`
	// 起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeAttackWorldMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackWorldMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadAccessRecordInfo struct {

	// 记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 下载任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 查询条件query string

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 当前下载任务的日志条数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 下载任务运行状态：-1-下载超时，0-下载等待，1-下载完成，2-下载失败，4-正在下载

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 下载文件URL

	Url *string `json:"Url,omitempty" name:"Url"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后更新修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 下载任务需下载的日志总条数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ApiSecIpAbnormalSource struct {

	// 异常ip段

	IpSection *string `json:"IpSection,omitempty" name:"IpSection"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeApiRuleListRequest struct {
	*tchttp.BaseRequest

	// 需要查询的API所属的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// asc或者desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 目前支持根据ts排序

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeApiRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTsRegionRequest struct {
	*tchttp.BaseRequest

	// 二级区域

	SubRegion *string `json:"SubRegion,omitempty" name:"SubRegion"`
	// 一级区域

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *ModifyTsRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTsRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWafThreatenIntelligenceRequest struct {
	*tchttp.BaseRequest

	// 配置WAF威胁情报封禁模块详情

	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitempty" name:"WafThreatenIntelligenceDetails"`
}

func (r *ModifyWafThreatenIntelligenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWafThreatenIntelligenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalancer struct {

	// 负载均衡LD的ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡LD的名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 负载均衡监听器的ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 负载均衡监听器的名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 负载均衡实例的IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 负载均衡实例的端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 负载均衡LD的地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 监听器协议，http、https

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 负载均衡监听器所在的zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 负载均衡的VPCID，公网为-1，内网按实际填写

	NumericalVpcId *int64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`
	// 负载均衡的网络类型

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 负载均衡的域名

	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitempty" name:"LoadBalancerDomain"`
}

type DescribeBotRecordDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// bot详情数据

		Res []*string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotRecordDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRecordDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// exp_ts或者mod_ts

	By *string `json:"By,omitempty" name:"By"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// asc或者desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeCustomWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotOverviewRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
}

func (r *DescribeBotOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainIpv6StatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的状态 （0: 操作失败 1:操作成功 2:企业版以上不支持 3:企业版以下不支持 ）

		Ipv6Status *int64 `json:"Ipv6Status,omitempty" name:"Ipv6Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainIpv6StatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainIpv6StatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchApiRules struct {

	// 启用或者停用api规则的url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 启用或者停用api规则的请求方法

	Method *string `json:"Method,omitempty" name:"Method"`
}

type TiInfo struct {

	// 情报类别

	Type *string `json:"Type,omitempty" name:"Type"`
	// 情报值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 情报描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type GenerateDealsAndPayNewRequest struct {
	*tchttp.BaseRequest

	// 计费下单入参

	Goods []*GoodNews `json:"Goods,omitempty" name:"Goods"`
}

func (r *GenerateDealsAndPayNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateDealsAndPayNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbDomainsInfo struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// waf类型

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 是否是cdn

	IsCdn *uint64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// 负载均衡算法

	LoadBalancerSet []*LoadBalancerPackageNew `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
	// 镜像模式

	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`
	// 绑定clb状态

	State *int64 `json:"State,omitempty" name:"State"`
	// 负载均衡类型，clb或者apisix

	AlbType *string `json:"AlbType,omitempty" name:"AlbType"`
	// IsCdn=3时，表示自定义header

	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`
	// cdc类型会增加集群信息

	CdcClusters *string `json:"CdcClusters,omitempty" name:"CdcClusters"`
	// 云类型:public:公有云；private:私有云;hybrid:混合云

	CloudType *string `json:"CloudType,omitempty" name:"CloudType"`
}

type DescribeRequestCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求总数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRequestCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRequestCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotTypeStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量值

		TCB *uint64 `json:"TCB,omitempty" name:"TCB"`
		// 数量值

		UB *uint64 `json:"UB,omitempty" name:"UB"`
		// 数量值

		UCB *uint64 `json:"UCB,omitempty" name:"UCB"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotTypeStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotTypeStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableRateLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 通用返回

		BaseInfo *RateLimitCommonRsp `json:"BaseInfo,omitempty" name:"BaseInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableRateLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableRateLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpsertSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Data *string `json:"Data,omitempty" name:"Data"`
		// SessionID

		SessionID *int64 `json:"SessionID,omitempty" name:"SessionID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpsertSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBotSceneActionRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 动作策略ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteBotSceneActionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBotSceneActionRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotActionStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 验证码统计

		Captcha *uint64 `json:"Captcha,omitempty" name:"Captcha"`
		// 拦截统计

		Intercept *uint64 `json:"Intercept,omitempty" name:"Intercept"`
		// 监控统计

		Monitor *uint64 `json:"Monitor,omitempty" name:"Monitor"`
		// 重定向统计

		Redirect *uint64 `json:"Redirect,omitempty" name:"Redirect"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotActionStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotActionStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotThreatIntelligenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 威胁情报配置列表

		Data []*ThreatIntelligenceEntry `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotThreatIntelligenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotThreatIntelligenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpartUserInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSpartUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpartUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGlobalWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序参数

	By *string `json:"By,omitempty" name:"By"`
	// 支持 Status，SignatureId（规则id），Target, TargetValue (Target 支持字段为Host，URI，FullUri，Parameter，Cookie，TargetValue为对应的筛选值

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// 无用参数

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *DescribeGlobalWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGlobalWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyPostCLSFlowRequest struct {
	*tchttp.BaseRequest

	// 投递流的流ID

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
	// 1-访问日志，2-攻击日志，默认为访问日志。

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
}

func (r *DestroyPostCLSFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyPostCLSFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiInfoLeakRulesRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 动作类型

	ActionType *int64 `json:"ActionType,omitempty" name:"ActionType"`
	// 翻页

	PageInfo *PageInfo `json:"PageInfo,omitempty" name:"PageInfo"`
}

func (r *DescribeAntiInfoLeakRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAntiInfoLeakRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAccessStatPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势图参数

		Data []*BotStatPointItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotAccessStatPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAccessStatPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTsResourceRequest struct {
	*tchttp.BaseRequest

	// 序号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteTsResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTsResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDNSDetectDomainRequest struct {
	*tchttp.BaseRequest

	// 编辑的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 编辑的记录的ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 编辑的记录的权威记录的数组

	AuthIP []*string `json:"AuthIP,omitempty" name:"AuthIP"`
}

func (r *ModifyDNSDetectDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDNSDetectDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHWDomainStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名列表

		Data *QueryHWDomainStateData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryHWDomainStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryHWDomainStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppletAccessDomainV3 struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 服务协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 在WAF接入的状态
	// 0: 未接入
	// 1: 已接入
	//

	AccessStatus *uint64 `json:"AccessStatus,omitempty" name:"AccessStatus"`
}

type CopyCustomRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功复制的域名列表

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyCustomRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleDeleteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleDeleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleDeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddCustomRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 添加成功的规则ID

		RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCustomRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpMainClassRequest struct {
	*tchttp.BaseRequest

	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持 MainClassID

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// cn/en

	OpLanguage *string `json:"OpLanguage,omitempty" name:"OpLanguage"`
}

func (r *DescribeOpMainClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpMainClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditOpRuleUpdateLogStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditOpRuleUpdateLogStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditOpRuleUpdateLogStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAntiInfoLeakRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAntiInfoLeakRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAntiInfoLeakRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchAreaBanListData struct {

	// 规则列表

	List []*BatchAreaBanListItem `json:"List,omitempty" name:"List"`
	// 列表总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type DeleteBlockPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBlockPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBlockPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotSceneActionRuleRequest struct {
	*tchttp.BaseRequest

	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotSceneActionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSceneActionRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVIPRSRequest struct {
	*tchttp.BaseRequest

	// VIP信息

	VIP *string `json:"VIP,omitempty" name:"VIP"`
}

func (r *DescribeVIPRSRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVIPRSRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleUploadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据库中的id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleUploadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchTrafficMarkingRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// Bot防护标记外层开关

	BotStatus *int64 `json:"BotStatus,omitempty" name:"BotStatus"`
	// 接口复用用作选择：
	// traffic: 流量标记
	// bot: bot防护标记

	Selection *string `json:"Selection,omitempty" name:"Selection"`
}

func (r *SwitchTrafficMarkingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchTrafficMarkingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessLogCountRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 查询结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 查询条件

	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeAccessLogCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessLogCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则白名单列表

		List []*UserWhiteRule `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotIdRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 配置信息，支持批量

	Data []*BotIdConfig `json:"Data,omitempty" name:"Data"`
	// 0-全局设置不生效 1-全局开关配置字段生效 2-全局动作配置字段生效 3-全局开关和动作字段都生效

	GlobalSwitch *int64 `json:"GlobalSwitch,omitempty" name:"GlobalSwitch"`
	// 全局开关

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 全局动作

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
}

func (r *ModifyBotIdRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotIdRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiSecKey struct {

	// api名称

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 请求方法

	Method *string `json:"Method,omitempty" name:"Method"`
}

type BotUCBPreinstallRuleData struct {

	// 参数封装

	Res []*BotUCBPreinstallRule `json:"Res,omitempty" name:"Res"`
}

type ModifyCdcResourceRequest struct {
	*tchttp.BaseRequest

	// 序号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// App id

	OpAppId *uint64 `json:"OpAppId,omitempty" name:"OpAppId"`
	// Cluster id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// redis passwd

	Auth *string `json:"Auth,omitempty" name:"Auth"`
	// Resource Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// 0: disable 1:enable

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyCdcResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCdcResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TargetEntity struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeBotGlobalConfigRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotGlobalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpiredQPSPackage struct {

	// 资源ID

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 是否自动续费，1：自动续费，0：不自动续费

	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type HybridCluster struct {

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群状态：正常或异常，异常则通过StateStatus给出异常原因

	State *int64 `json:"State,omitempty" name:"State"`
	// 安全检测超时时间

	TimeoutThreshold *int64 `json:"TimeoutThreshold,omitempty" name:"TimeoutThreshold"`
	// 最大body检测长度

	BodyLengthThreshold *int64 `json:"BodyLengthThreshold,omitempty" name:"BodyLengthThreshold"`
	// 健康状态阈值

	HealthyThreshold *float64 `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
	// 集群异常状态说明

	StateDetail *string `json:"StateDetail,omitempty" name:"StateDetail"`
	// 集群接入方式

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// waf实例等级，如果未绑定实例为0

	InstanceLevel *int64 `json:"InstanceLevel,omitempty" name:"InstanceLevel"`
	// waf防御状态

	WafStatus *int64 `json:"WafStatus,omitempty" name:"WafStatus"`
	// 集群对应的虚拟域名

	VirtualDomain *string `json:"VirtualDomain,omitempty" name:"VirtualDomain"`
	// 精准域名列表

	PreciseDomains []*string `json:"PreciseDomains,omitempty" name:"PreciseDomains"`
	// WAF日志开关状态，0关闭1开启

	CLSStatus *int64 `json:"CLSStatus,omitempty" name:"CLSStatus"`
	// clb投递开关

	PostCLSStatus *int64 `json:"PostCLSStatus,omitempty" name:"PostCLSStatus"`
	// kafka投递开关

	PostCKafkaStatus *int64 `json:"PostCKafkaStatus,omitempty" name:"PostCKafkaStatus"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最新时间

	LatestUpdateTime *string `json:"LatestUpdateTime,omitempty" name:"LatestUpdateTime"`
	// 1: 初始化 2: 本地部署 3: 完成创建 4: 完成接入 5: 日常运维

	ClusterStage *int64 `json:"ClusterStage,omitempty" name:"ClusterStage"`
}

type DeleteAntiInfoLeakRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAntiInfoLeakRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAntiInfoLeakRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotIdRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		Data []*BotIdDetail `json:"Data,omitempty" name:"Data"`
		// 符合条件的规则数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Bot规则数量统计信息

		StatInfo *BotIdStat `json:"StatInfo,omitempty" name:"StatInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotIdRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotIdRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopAttackIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 柱状图

		Histogram []*string `json:"Histogram,omitempty" name:"Histogram"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopAttackIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopAttackIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TLSCiphers struct {

	// TLS版本ID

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 加密套件ID

	CipherId *int64 `json:"CipherId,omitempty" name:"CipherId"`
	// 加密套件

	CipherName *string `json:"CipherName,omitempty" name:"CipherName"`
}

type CreateHybridClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群id

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHybridClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHybridClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 删除的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 删除的规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteCustomWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiPkg struct {

	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 地域

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 申请数量

	InquireNum *int64 `json:"InquireNum,omitempty" name:"InquireNum"`
	// 使用数量

	UsedNum *int64 `json:"UsedNum,omitempty" name:"UsedNum"`
	// 续费标志

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 计费项

	BillingItem *string `json:"BillingItem,omitempty" name:"BillingItem"`
	// 1 API安全6折

	APICPWaf *int64 `json:"APICPWaf,omitempty" name:"APICPWaf"`
	// 1 表示5折折扣
	// 2 表示4折折扣

	APINPWaf *int64 `json:"APINPWaf,omitempty" name:"APINPWaf"`
	// api安全7天试用标识。1试用。0没试用

	IsAPISecurityTrial *int64 `json:"IsAPISecurityTrial,omitempty" name:"IsAPISecurityTrial"`
}

type DescribeApiSecAttackSourceListRequest struct {
	*tchttp.BaseRequest

	// 当前查询第几页,默认是0

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 每一页需要展示多少条数据

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤条件

	Filters []*ApiDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 排序，第一个元素为排序的key，第二个元素为排序规则，其中1 为升序排列，而-1 是用于降序排列

	Sort []*string `json:"Sort,omitempty" name:"Sort"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 是否返回总数，true表示返回

	NeedTotalCount *bool `json:"NeedTotalCount,omitempty" name:"NeedTotalCount"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeApiSecAttackSourceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecAttackSourceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePanDomainDetailRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePanDomainDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePanDomainDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAreaBanAreasRequest struct {
	*tchttp.BaseRequest

	// 需要修改的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// waf版本信息，spart-waf或者clb-waf，其他无效，请一定填写

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 需要新增的封禁地域

	Areas []*string `json:"Areas,omitempty" name:"Areas"`
}

func (r *AddAreaBanAreasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAreaBanAreasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCopyCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCopyCustomWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCopyCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBatchCustomRuleRequest struct {
	*tchttp.BaseRequest

	// 数据类型
	// "custom-rule"-自定义规则、"custom-white-rule"-精准白名单
	//

	DataType *string `json:"DataType,omitempty" name:"DataType"`
	// 0-指定Id删除、1-删除全部（除部分排除的Id）
	//

	IsDeleteAll *int64 `json:"IsDeleteAll,omitempty" name:"IsDeleteAll"`
	// 具体Ids 由IsDeleteAll而定

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 筛选条件

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DeleteBatchCustomRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchCustomRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryListItems struct {

	// API规则的描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 该API规则所使用的方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 该API的开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 该条api对应的动作，包括观察和拦截

	Action *uint64 `json:"Action,omitempty" name:"Action"`
	// API的URL地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// API的规则id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// API的来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// API子规则的名称列表

	Parameters []*string `json:"Parameters,omitempty" name:"Parameters"`
	// 时间戳

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type DescribePortsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// http端口列表

		HttpPorts []*string `json:"HttpPorts,omitempty" name:"HttpPorts"`
		// https端口列表

		HttpsPorts []*string `json:"HttpsPorts,omitempty" name:"HttpsPorts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePortsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackHistogramRequest struct {
	*tchttp.BaseRequest

	// Lucene语法

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 查询的域名，所有域名使用all

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetAttackHistogramRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackHistogramRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询到的CC规则的列表

		Data *CCRuleLists `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotUCBRecordsRequest struct {
	*tchttp.BaseRequest

	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 限制数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 过滤规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 偏移

	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`
	// 动作

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 起始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
}

func (r *DescribeBotUCBRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotUCBRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceItem struct {

	// Redis集群地区版本

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 版本

	Version *uint64 `json:"Version,omitempty" name:"Version"`
	// nginx地区

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 引擎类型

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
	// 实例大区域 chn(国内)、intl(国际)

	Area *string `json:"Area,omitempty" name:"Area"`
	// WAF类型 saas、clb、cdn、cdc

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 实例环境类型 Development、Test、Pre、Production

	Env *string `json:"Env,omitempty" name:"Env"`
	// 上报时间

	ReportTimeStr *string `json:"ReportTimeStr,omitempty" name:"ReportTimeStr"`
}

type DescribeAttackDownloadRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载记录类型

		Records []*DownloadRecordItem `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackDownloadRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackDownloadRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotActionRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotActionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotActionRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotQPS struct {

	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 有效时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 资源数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 资源所在地区

	Region *string `json:"Region,omitempty" name:"Region"`
	// 使用qps的最大值

	MaxBotQPS *uint64 `json:"MaxBotQPS,omitempty" name:"MaxBotQPS"`
	// 续费标志

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type ModifyJsInjectRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 规则对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 前端对抗规则ID，非必填，默认可以忽略。

	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
	// 前端对抗规则的开关，0（关闭） 1（开启）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 如果是V4版本则填写4，否则置空就行

	Ver *uint64 `json:"Ver,omitempty" name:"Ver"`
}

func (r *ModifyJsInjectRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJsInjectRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiAnalyzeStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已经开启的数量,如果返回值为3（大于支持的域名开启数量），则表示开启失败

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 不支持开启的域名列表

		UnSupportedList []*string `json:"UnSupportedList,omitempty" name:"UnSupportedList"`
		// 开启/关闭失败的域名列表

		FailDomainList []*string `json:"FailDomainList,omitempty" name:"FailDomainList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiAnalyzeStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiAnalyzeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleUploadRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Name *string `json:"Name,omitempty" name:"Name"`
	// 上传特征库文件名

	RuleFileName *string `json:"RuleFileName,omitempty" name:"RuleFileName"`
}

func (r *RuleUploadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleUploadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllGenerateDealsRequest struct {
	*tchttp.BaseRequest

	// 计费下单入参

	Goods []*Goods `json:"Goods,omitempty" name:"Goods"`
}

func (r *ModifyAllGenerateDealsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAllGenerateDealsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DNSDetectRecord struct {

	// 记录的ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 被劫持的记录数

	HijackRecords *uint64 `json:"HijackRecords,omitempty" name:"HijackRecords"`
	// 劫持的地域数

	HijackRegions *uint64 `json:"HijackRegions,omitempty" name:"HijackRegions"`
	// 记录新建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 权威记录数组

	AuthIP []*string `json:"AuthIP,omitempty" name:"AuthIP"`
}

type ModuleTypeData struct {

	// 出参数据

	ModuleTypeList []*ModuleTypeItem `json:"ModuleTypeList,omitempty" name:"ModuleTypeList"`
}

type DescribeBotGlobalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户开启的智能统计规则条数

		StatRuleNum *int64 `json:"StatRuleNum,omitempty" name:"StatRuleNum"`
		// 用户配置的会话管理规则条数

		TokenRuleNums *int64 `json:"TokenRuleNums,omitempty" name:"TokenRuleNums"`
		// 用户开启的合法爬虫数目

		GoodBotNums *int64 `json:"GoodBotNums,omitempty" name:"GoodBotNums"`
		// 用户开启的UA规则条数

		UaRuleNum *int64 `json:"UaRuleNum,omitempty" name:"UaRuleNum"`
		// 当前的bot_id规则数目

		BotIdRuleNums *int64 `json:"BotIdRuleNums,omitempty" name:"BotIdRuleNums"`
		// 用户配置的前端对抗规则数

		JsInjectRuleNum *int64 `json:"JsInjectRuleNum,omitempty" name:"JsInjectRuleNum"`
		// 用户配置的AI策略规则条数

		AiRuleNum *int64 `json:"AiRuleNum,omitempty" name:"AiRuleNum"`
		// 用户开启的TI规则条数

		TiRuleNum *int64 `json:"TiRuleNum,omitempty" name:"TiRuleNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertDomainItem struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 0未开启长链接 1 开启长链接

	Keepalive *int64 `json:"Keepalive,omitempty" name:"Keepalive"`
	// uin

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 域名ID

	DomainID *string `json:"DomainID,omitempty" name:"DomainID"`
}

type MonitorDomainItem struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 用户自己的Appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 域名对应的版本信息，0代表saas，1代表clb

	Edition *int64 `json:"Edition,omitempty" name:"Edition"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

type DescribeApiListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户配置的api详细信息

		RuleList []*DescribeApi `json:"RuleList,omitempty" name:"RuleList"`
		// 接口的数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchCustomRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Data *BatchCustomRuleListData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBatchCustomRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBatchCustomRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportApiRulesRequest struct {
	*tchttp.BaseRequest

	// 用户需要导入的api的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 用户需要导入的文件内容

	File *string `json:"File,omitempty" name:"File"`
	// 用户需要导入的文件类型

	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

func (r *ImportApiRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportApiRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryClientMsgRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *QueryClientMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryClientMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotAccessItem struct {

	// 源ip

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 命中的url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 地理位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 日志总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 命中模块

	BotModule *string `json:"BotModule,omitempty" name:"BotModule"`
	// 动作列表

	BotAction *string `json:"BotAction,omitempty" name:"BotAction"`
	// 结束时间

	MaxLogTime *string `json:"MaxLogTime,omitempty" name:"MaxLogTime"`
	// 开始时间

	MinLogTime *string `json:"MinLogTime,omitempty" name:"MinLogTime"`
	// 攻击次数

	AttackCount *int64 `json:"AttackCount,omitempty" name:"AttackCount"`
}

type BotAutoStatRule struct {

	// 策略名称 如avg_speed

	Feature *string `json:"Feature,omitempty" name:"Feature"`
	// 策略状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 策略等级，100（宽松），200（中等），300(严格),0表示智能推荐

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 特征描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 修改时间戳，默认为0，表示没有修改

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type BotTCBRule struct {

	// botTCB子类型

	FeedFetcher *BotTCBRuleItem `json:"FeedFetcher,omitempty" name:"FeedFetcher"`
	// botTCB子类型

	LinkChecker *BotTCBRuleItem `json:"LinkChecker,omitempty" name:"LinkChecker"`
	// botTCB子类型

	Marketing *BotTCBRuleItem `json:"Marketing,omitempty" name:"Marketing"`
	// botTCB子类型

	ScreenshotCreator *BotTCBRuleItem `json:"ScreenshotCreator,omitempty" name:"ScreenshotCreator"`
	// botTCB子类型

	SearchEngineBot *BotTCBRuleItem `json:"SearchEngineBot,omitempty" name:"SearchEngineBot"`
	// botTCB子类型

	SiteMonitor *BotTCBRuleItem `json:"SiteMonitor,omitempty" name:"SiteMonitor"`
	// botTCB子类型

	SpeedTester *BotTCBRuleItem `json:"SpeedTester,omitempty" name:"SpeedTester"`
	// botTCB子类型

	Tool *BotTCBRuleItem `json:"Tool,omitempty" name:"Tool"`
	// botTCB子类型

	Uncategorised *BotTCBRuleItem `json:"Uncategorised,omitempty" name:"Uncategorised"`
	// botTCB子类型

	VirusScanner *BotTCBRuleItem `json:"VirusScanner,omitempty" name:"VirusScanner"`
	// botTCB子类型

	VulnerabilityScanner *BotTCBRuleItem `json:"VulnerabilityScanner,omitempty" name:"VulnerabilityScanner"`
	// botTCB子类型

	WebScraper *BotTCBRuleItem `json:"WebScraper,omitempty" name:"WebScraper"`
	// appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 时间戳

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type DescribeUserGradeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserGradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppletAccess struct {

	// 小程序接入名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 小程序类型传入 1， 预留扩展用

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 小程序源站域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 服务协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 微信小程序ID

	AppletID *string `json:"AppletID,omitempty" name:"AppletID"`
	// 详情

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 微信网关接入节点域名

	AccessDomain *string `json:"AccessDomain,omitempty" name:"AccessDomain"`
	// 接入配置阶段的状态,
	// 0 初始状态
	// 100 完成

	AccessStatus *uint64 `json:"AccessStatus,omitempty" name:"AccessStatus"`
	// 客户端配置阶段的状态
	// 0 初始状态
	// 100 完成

	ClientConfigStatus *uint64 `json:"ClientConfigStatus,omitempty" name:"ClientConfigStatus"`
	// 客户端测试阶段的状态
	// 0 初始状态（未验证）
	// 1 已验证，待发布
	// 100 完成（已验证，已发布）

	ClientTestStatus *uint64 `json:"ClientTestStatus,omitempty" name:"ClientTestStatus"`
	// 是否要进行链路测试 0 不进行测试 1 进行测试

	LinkTest *uint64 `json:"LinkTest,omitempty" name:"LinkTest"`
	// 链路探测结果：
	// 0 初始状态 1-99 验证失败编码
	// 100 验证成功

	LinkTestStatus *uint64 `json:"LinkTestStatus,omitempty" name:"LinkTestStatus"`
	// 网关ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 接入节点ID

	AccessID *string `json:"AccessID,omitempty" name:"AccessID"`
}

type ModifyCdcResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCdcResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCdcResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRateLimitsRequest struct {
	*tchttp.BaseRequest

	// 规则名

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 规则所针对的服务接口

	ApiPaths []*string `json:"ApiPaths,omitempty" name:"ApiPaths"`
	// 限流配额，包含限流周期和配额总数

	Amount *Amount `json:"Amount,omitempty" name:"Amount"`
	// 请求参数匹配条件，需全匹配才通过

	Arguments []*MatchArgument `json:"Arguments,omitempty" name:"Arguments"`
	// 规则优先级，Priority越大，优先级越低

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 限流策略,0:观察,1:拦截，2:人机
	//

	LimitStrategy *int64 `json:"LimitStrategy,omitempty" name:"LimitStrategy"`
	// 默认为0

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 支持API/DOMAIN/SESSION

	LimitObject *string `json:"LimitObject,omitempty" name:"LimitObject"`
}

func (r *UpdateRateLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRateLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpUserSignatureRuleRequest struct {
	*tchttp.BaseRequest

	// 页面显示个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持MainClassID, SubClassID, CveID，Status，ID；ID为规则id, Status = 0/1

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// appid

	OpAppId *uint64 `json:"OpAppId,omitempty" name:"OpAppId"`
	// en/cn

	OpLanguage *string `json:"OpLanguage,omitempty" name:"OpLanguage"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 地域信息

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序条件，支持signature_id, modify_time

	By *string `json:"By,omitempty" name:"By"`
	// 分片偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeOpUserSignatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpUserSignatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotGlobalUAStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotGlobalUAStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotGlobalUAStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVIPRSResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录

		RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVIPRSResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVIPRSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotRecordPointsRequest struct {
	*tchttp.BaseRequest

	// bot的id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 步长

	Stride *uint64 `json:"Stride,omitempty" name:"Stride"`
	// 起始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotRecordPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRecordPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJsInjectRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 前端对抗的规则

		Rule *JsInjectRule `json:"Rule,omitempty" name:"Rule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJsInjectRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJsInjectRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HistogramItem struct {

	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 响应时间

	Time *float64 `json:"Time,omitempty" name:"Time"`
	// URL

	Url *string `json:"Url,omitempty" name:"Url"`
	// 城市

	City *string `json:"City,omitempty" name:"City"`
}

type ModifyBotAutoStatisticRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotAutoStatisticRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotAutoStatisticRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiAggregateTopNRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 需要查询TOP的维度名

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 过滤条件

	Filters []*BotDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 是否查询全域名的三个特殊图标

	GlobalFlag *bool `json:"GlobalFlag,omitempty" name:"GlobalFlag"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 需要的Top数，默认5， 最大值100

	TopN *uint64 `json:"TopN,omitempty" name:"TopN"`
}

func (r *DescribeApiAggregateTopNRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiAggregateTopNRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBatchAreaBanStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBatchAreaBanStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBatchAreaBanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBusinessRiskRequest struct {
	*tchttp.BaseRequest

	// 删除策略记录ID数组

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBusinessRiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBusinessRiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceWafInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 询价结果

		CostInfo []*InstancePriceItem `json:"CostInfo,omitempty" name:"CostInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceWafInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceWafInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 需要更改的规则的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 白名单id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则的id列表

	Rules []*uint64 `json:"Rules,omitempty" name:"Rules"`
	// 规则匹配路径

	Url *string `json:"Url,omitempty" name:"Url"`
	// 规则匹配方法

	Function *string `json:"Function,omitempty" name:"Function"`
	// 规则的开关状态，0表示关闭开关，1表示打开开关

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDomainWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHostInnerRequest struct {
	*tchttp.BaseRequest

	// 删除的域名列表

	HostsDel []*HostDel `json:"HostsDel,omitempty" name:"HostsDel"`
}

func (r *DeleteHostInnerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHostInnerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppletAuthorizeV3Request struct {
	*tchttp.BaseRequest

	// 支持的过滤器：
	// 1 AppletID: 小程序ID
	//
	// 2 AccessStatus: 小程序接入状态
	// 0 未接入waf

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAppletAuthorizeV3Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppletAuthorizeV3Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCustomRuleListItem struct {

	// 规则Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 执行动作

	ActionType *int64 `json:"ActionType,omitempty" name:"ActionType"`
	// 加白模块

	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
	// 有效期

	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 重定向地址

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 优先级

	SortId *int64 `json:"SortId,omitempty" name:"SortId"`
	// 开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 域名列表

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 策略列表

	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`
	// 事件Id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 生效状态

	ValidStatus *uint64 `json:"ValidStatus,omitempty" name:"ValidStatus"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DeleteGlobalWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGlobalWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGlobalWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAccessLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志显示内容

		Data []*BotAccessItem `json:"Data,omitempty" name:"Data"`
		// 日志条数，当传入CountFlag为true的时候显示为正确条数，其他情况显示0

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotAccessLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDNSDetectDataMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域详情

		Province []*DescribeDNSDetectDataMapRspItem `json:"Province,omitempty" name:"Province"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDNSDetectDataMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSDetectDataMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataSorter struct {

	// 排序字段

	Entity *string `json:"Entity,omitempty" name:"Entity"`
	//  1:降序 -1:升序

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeAccountAttributesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoStat struct {

	// 命中的维度特征名称

	Key *string `json:"Key,omitempty" name:"Key"`
	// 维度的实时特征值

	Value *float64 `json:"Value,omitempty" name:"Value"`
	// 维度的异常值

	Threshold *float64 `json:"Threshold,omitempty" name:"Threshold"`
}

type DescribeQpsWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0没有黑白名单
		// 1 白名单
		// 2 黑名单

		Type *int64 `json:"Type,omitempty" name:"Type"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQpsWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQpsWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableUserSignaturePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名规则策略id

		UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableUserSignaturePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableUserSignaturePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBotSessionKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBotSessionKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBotSessionKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppletAccessV3Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAppletAccessV3Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppletAccessV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProtectionStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *ModifyProtectionStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProtectionStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyBotsUCBPreinstallRuleRsp struct {

	// 0表示成功，其他表示失败

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 成功的数目

	SuccessNum *int64 `json:"SuccessNum,omitempty" name:"SuccessNum"`
	// 失败的数目

	FailedNum *int64 `json:"FailedNum,omitempty" name:"FailedNum"`
	// 成功的域名列表

	T []*string `json:"T,omitempty" name:"T"`
}

type DescribeUserAppletInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserAppletInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAppletInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClsPackageNew struct {

	// 资源ID

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 是否自动续费，1：自动续费，0：不自动续费

	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 套餐个数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 地域，目前在clb-waf中没有用到

	Region *string `json:"Region,omitempty" name:"Region"`
	// 已经使用的cls

	Used *float64 `json:"Used,omitempty" name:"Used"`
}

type DescribeAutoDenyIPRequest struct {
	*tchttp.BaseRequest

	// 查询IP自动封禁状态

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 有效时间最小时间戳

	VtsMin *uint64 `json:"VtsMin,omitempty" name:"VtsMin"`
	// 有效时间最大时间戳

	VtsMax *uint64 `json:"VtsMax,omitempty" name:"VtsMax"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 计数标识

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 创建时间最小时间戳

	CtsMin *uint64 `json:"CtsMin,omitempty" name:"CtsMin"`
	// 创建时间最大时间戳

	CtsMax *uint64 `json:"CtsMax,omitempty" name:"CtsMax"`
	// 偏移量

	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`
	// 限制条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 策略名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeAutoDenyIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoDenyIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSignatureRuleRequest struct {
	*tchttp.BaseRequest

	// 需要查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 排序字段，支持 signature_id, modify_time

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持 MainClassName，SubClassID ,CveID, Status, ID;  ID为规则id

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUserSignatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserSignatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBatchAreaBanAppIdDomainRequest struct {
	*tchttp.BaseRequest

	// 数据来源 custom-自定义、batch-批量防护（默认值）、all-全部

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *QueryBatchAreaBanAppIdDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBatchAreaBanAppIdDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessCursorRequest struct {
	*tchttp.BaseRequest

	// 指定时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
}

func (r *DescribeAccessCursorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessCursorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotTiRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关状态

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotTiRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotTiRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpUserSignaturePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则列表

		OpUserSigPolicy []*OpUserSignaturePolicy `json:"OpUserSigPolicy,omitempty" name:"OpUserSigPolicy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpUserSignaturePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpUserSignaturePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainEngineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainEngineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySceneJsInjectRuleRequest struct {
	*tchttp.BaseRequest

	// IsGlobal为0时，字段必填

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 规则对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 是否开启自动化工具识别，true：是，false：否

	AutomatedTool *bool `json:"AutomatedTool,omitempty" name:"AutomatedTool"`
	// 是否开启页面防调试，true：是，false：否

	CrackBehavior *bool `json:"CrackBehavior,omitempty" name:"CrackBehavior"`
	// 【复用字段】规则ID,可以是场景也可以是全局

	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
	// 前端对抗的开关，0（关闭） 1（开启）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 修改默认场景的前端对抗防护路径时传入

	CheckUrls []*string `json:"CheckUrls,omitempty" name:"CheckUrls"`
	// 可选的值为：0（放行） 1（阻断） 2（人机识别） 3（观察） 4（重定向）

	RuleAction *uint64 `json:"RuleAction,omitempty" name:"RuleAction"`
	// 动作为重定向的时候重定向的URL

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 搜索引擎白名单，可选的值为：["sogou", "baidu", "yandex", "360", "yahoo", "bing", "google"]

	GoodBot []*string `json:"GoodBot,omitempty" name:"GoodBot"`
	// 规则内部ID

	RasSiteID *string `json:"RasSiteID,omitempty" name:"RasSiteID"`
}

func (r *ModifySceneJsInjectRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySceneJsInjectRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDomainCustomRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 查询域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *QueryDomainCustomRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDomainCustomRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRuleTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBatchCustomRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBatchCustomRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBatchCustomRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotTCBRuleItem struct {

	// 动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 数目

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type ClbHostResult struct {

	// WAF绑定的监听器实例

	LoadBalancer *LoadBalancer `json:"LoadBalancer,omitempty" name:"LoadBalancer"`
	// WAF绑定的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// WAF绑定的实例ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 是否有绑定WAF，1：绑定了WAF，0：没有绑定WAF

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 绑定了WAF的情况下，WAF流量模式，1：清洗模式，0：镜像模式（默认）

	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`
}

type DescribeRuleLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// waf模块的规格

		Res *WafRuleLimit `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessExportsRequest struct {
	*tchttp.BaseRequest

	// 客户要查询的日志主题ID，每个客户都有对应的一个主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccessExportsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessExportsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiSecEventListRequest struct {
	*tchttp.BaseRequest

	// 排序，第一个元素为排序的key，第二个元素为排序规则，其中1 为升序排列，而-1 是用于降序排列

	Sort []*string `json:"Sort,omitempty" name:"Sort"`
	// 查询开始时间

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 查询结束时间

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 是否查询数量，默认不查询，为true则进行查询

	NeedTotalCount *bool `json:"NeedTotalCount,omitempty" name:"NeedTotalCount"`
	// 过滤条件

	Filters []*ApiDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 查询当前的页

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 每一页显示多少条数据

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeApiSecEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatedDomainRequest struct {
	*tchttp.BaseRequest

	// ssl证书id列表

	CertificateIds []*string `json:"CertificateIds,omitempty" name:"CertificateIds"`
	// 字符串枚举。可取值：
	// bindable: 查询所有可绑定的域名（即证书可以匹配到的域名）
	// binded: 查询所有已绑定的域名

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeCertificatedDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCertificatedDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiInfoLeakageRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则列表

		RuleList []*DescribeAntiLeakageItem `json:"RuleList,omitempty" name:"RuleList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAntiInfoLeakageRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAntiInfoLeakageRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CdcResource struct {

	// 资源id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则状态，0/1

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 用户id

	OpAppId *uint64 `json:"OpAppId,omitempty" name:"OpAppId"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 认证信息

	Auth *string `json:"Auth,omitempty" name:"Auth"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 资源类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeDNSDetectHijackDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 被劫持的记录数组

		List []*DNSDetectHijackData `json:"List,omitempty" name:"List"`
		// 被劫持的记录数

		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDNSDetectHijackDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSDetectHijackDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DNSDetectHijackData struct {

	// 被劫持的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 被劫持后的IP

	HijackIP *string `json:"HijackIP,omitempty" name:"HijackIP"`
	// 被劫持后的IP的运营商

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 被劫持的地域数

	HijackRegions *uint64 `json:"HijackRegions,omitempty" name:"HijackRegions"`
	// 被劫持的用户数

	HijackUsers *uint64 `json:"HijackUsers,omitempty" name:"HijackUsers"`
	// 权威记录

	AuthIP *string `json:"AuthIP,omitempty" name:"AuthIP"`
}

type CreatePostCLSFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePostCLSFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePostCLSFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHybridClusterNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息

		Nodes []*HybridClusterNode `json:"Nodes,omitempty" name:"Nodes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHybridClusterNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHybridClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBatchCustomRuleRequest struct {
	*tchttp.BaseRequest

	// 批量规则Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 有效期
	//

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 优先级

	SortId *int64 `json:"SortId,omitempty" name:"SortId"`
	// 执行动作

	ActionType *int64 `json:"ActionType,omitempty" name:"ActionType"`
	// 重定向地址

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 加白模块

	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 事件Id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 域名列表

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
	// 策略详情列表

	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`
}

func (r *ModifyBatchCustomRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBatchCustomRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostStatus struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// WAF的开关，1：开，0：关

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DescribeAccessHistogramRequest struct {
	*tchttp.BaseRequest

	// 查询语句，语句长度最大为4096

	Query *string `json:"Query,omitempty" name:"Query"`
	// 柱状图间隔时间差，单位ms

	Interval *int64 `json:"Interval,omitempty" name:"Interval"`
	// 老版本查询的日志主题ID，新版本传空字符串即可

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
}

func (r *DescribeAccessHistogramRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessHistogramRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotSceneOverviewRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotSceneOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSceneOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选字段支持 AppId，Domain

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// 排序字段，支持 user_id, signature_id, modify_time

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeOpAttackWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpAttackWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpSignatureRuleRequest struct {
	*tchttp.BaseRequest

	// 筛选条件，支持 MainClassID，SubClassID , Status, CveID, ID, Status; Status为必填字段 status=all，查询全部，status=0/1/2，查询对应status的规则

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// cn/en

	OpLanguage *string `json:"OpLanguage,omitempty" name:"OpLanguage"`
	// 排序字段，支持 signature_id, modify_time

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式，asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOpSignatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpSignatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 匹配条件数组

	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`
	// 编辑的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 编辑的规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 编辑的规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果勾选多个，则以“，”串接。

	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。

	SortId *uint64 `json:"SortId,omitempty" name:"SortId"`
	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。

	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *ModifyCustomWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomPayloadsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败个数

		FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`
		// 失败ID数组

		FailedItems []*string `json:"FailedItems,omitempty" name:"FailedItems"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomPayloadsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomPayloadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySceneJsInjectRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySceneJsInjectRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySceneJsInjectRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MatchString struct {

	// 匹配类型，枚举，支持：EXACT（全匹配，默认），REGEX（正则表达式匹配），NOT_EQUALS（不等于），IN（包含），NOT_IN（不包含）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 匹配的目标值，如果选择的是包含和不包含，则通过逗号分割多个值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeBypassSpartaProtectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名列表

		Domains []*string `json:"Domains,omitempty" name:"Domains"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBypassSpartaProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBypassSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePostFlowServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePostFlowServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePostFlowServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalWhiteCond struct {

	// 匹配目标，HTTP-Method,Host,URI,FULL-URL,Parameter,Cookie,HTTP-Header,JSON-Elements

	Target *string `json:"Target,omitempty" name:"Target"`
	// 匹配操作，String-Match 字符串匹配，支持通配符，例如/floder1/*  /floder1/*/index.htm；Regular-Expression-Match 正则表达式匹配， Include 包含（HTTP-Method），Exclude 不包含（HTTP-Method）

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 方法列表，"GET", "POST", "HEAD"

	HttpMethodList []*string `json:"HttpMethodList,omitempty" name:"HttpMethodList"`
	// 变量名称, 适用于Parameter,Cookie,HTTP-Header,JSON-Elements

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否检查变量值，0:disable, 1: enable。 适用于Parameter,Cookie,HTTP-Header,JSON-Elements

	CheckValue *uint64 `json:"CheckValue,omitempty" name:"CheckValue"`
	// 变量值，适用于Host,URI,FULL-URL,Parameter,Cookie,HTTP-Header,JSON-Elements

	Value *string `json:"Value,omitempty" name:"Value"`
	// 条件之间的链接方式，可以为AND/OR。AND 表示与上一个条件与的关系，OR表示与上一个条件是或的关系。多个条件之间的链接关系如下： C1.AND C2.AND C3.OR C4.OR C5.AND C6.AND, 则组合的逻辑关系是： c1 and (c2 or c3 or c4) and c5 and c6

	Concatenate *string `json:"Concatenate,omitempty" name:"Concatenate"`
}

type IPType struct {

	// IP地址

	IP *string `json:"IP,omitempty" name:"IP"`
	// IP类型描述

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeDomainCountInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名总数

		AllDomain *uint64 `json:"AllDomain,omitempty" name:"AllDomain"`
		// 最近发现时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 接入域名总数

		WafDomainCount *uint64 `json:"WafDomainCount,omitempty" name:"WafDomainCount"`
		// 剩下配额

		LeftDomainCount *uint64 `json:"LeftDomainCount,omitempty" name:"LeftDomainCount"`
		// 开启防护域名数

		OpenWafDomain *uint64 `json:"OpenWafDomain,omitempty" name:"OpenWafDomain"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainCountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainCountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRiskInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebshellStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeWebshellStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebshellStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotStatPointItem struct {

	// 横坐标

	TimeStamp *string `json:"TimeStamp,omitempty" name:"TimeStamp"`
	// value的所属对象

	Key *string `json:"Key,omitempty" name:"Key"`
	// 纵列表

	Value *int64 `json:"Value,omitempty" name:"Value"`
	// Key对应的页面展示内容

	Label *string `json:"Label,omitempty" name:"Label"`
}

type ApiSecUaAbnormalSource struct {

	// ua

	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
	// ua类型

	UaType *string `json:"UaType,omitempty" name:"UaType"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type QpsConfig struct {

	// sp_wsm_waf_premium_clb 类型 弹性qps相关值

	SpWafPremiumC *QpsData `json:"SpWafPremiumC,omitempty" name:"SpWafPremiumC"`
	// sp_wsm_waf_premium 类型 弹性qps相关值

	SpWafPremium *QpsData `json:"SpWafPremium,omitempty" name:"SpWafPremium"`
	// sp_wsm_waf_enterprise_clb 类型 弹性qps相关值

	SpWafEnterpriseC *QpsData `json:"SpWafEnterpriseC,omitempty" name:"SpWafEnterpriseC"`
	// sp_wsm_waf_enterprise 类型 弹性qps相关值

	SpWafEnterprise *QpsData `json:"SpWafEnterprise,omitempty" name:"SpWafEnterprise"`
	// sp_wsm_waf_ultimate_clb 类型 弹性qps相关值

	SpWafUltimateC *QpsData `json:"SpWafUltimateC,omitempty" name:"SpWafUltimateC"`
	// sp_wsm_waf_ultimate 类型 弹性qps相关值

	SpWafUltimate *QpsData `json:"SpWafUltimate,omitempty" name:"SpWafUltimate"`
	// sp_wsm_waf_exclusive 类型 弹性qps相关值

	SpWafExclusive *QpsData `json:"SpWafExclusive,omitempty" name:"SpWafExclusive"`
	// sp_wsm_waf_hybrid 类型 弹性qps相关值

	SpWafHybrid *QpsData `json:"SpWafHybrid,omitempty" name:"SpWafHybrid"`
	// sp_wsm_waf_hybriden 类型 弹性qps相关值

	SpWafHybridEn *QpsData `json:"SpWafHybridEn,omitempty" name:"SpWafHybridEn"`
	// sp_wsm_waf_hybrid_pr 类型 弹性qps相关值

	SpWafHybridPr *QpsData `json:"SpWafHybridPr,omitempty" name:"SpWafHybridPr"`
}

type CopyBotTCBRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值

		Data *CopyBotTCBRuleRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyBotTCBRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBotTCBRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessLogsRequest struct {
	*tchttp.BaseRequest

	// 查询结果的排序依据

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 查询内容过滤的关键词

	Query *string `json:"Query,omitempty" name:"Query"`
	// 版本，clb-waf或者sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 查询的日志量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 查询游标。第一次查询时不用传，查看更多时传上一次查询结果的最后一条记录的该字段的值即可。

	Context *string `json:"Context,omitempty" name:"Context"`
	// 查询其实时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 查询的结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
}

func (r *DescribeAccessLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CacheUrlItems struct {

	// 标识

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 网址

	Uri *string `json:"Uri,omitempty" name:"Uri"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DeleteHostRequest struct {
	*tchttp.BaseRequest

	// 删除的域名列表

	HostsDel []*HostDel `json:"HostsDel,omitempty" name:"HostsDel"`
}

func (r *DeleteHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBatchAreaBanRequest struct {
	*tchttp.BaseRequest

	// 规则Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据来源 custom-自定义(默认)、batch-批量防护
	//

	Source *string `json:"Source,omitempty" name:"Source"`
	// 域名列表

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
	// 地域列表

	Areas []*string `json:"Areas,omitempty" name:"Areas"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyBatchAreaBanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBatchAreaBanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomWhiteRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 开关的状态，1是开启、0是关闭

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyCustomWhiteRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomWhiteRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfoQPS struct {

	// 计数

	Count *string `json:"Count,omitempty" name:"Count"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 过期时间

	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
	// 续费标记

	Renew *string `json:"Renew,omitempty" name:"Renew"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeAppletAccessRequest struct {
	*tchttp.BaseRequest

	// 支持的过滤器：Name: 小程序接入Name

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 上限

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAppletAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppletAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRolePlayStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 暂未授权角色

		Enable *bool `json:"Enable,omitempty" name:"Enable"`
		// 暂未授权角色名称

		UnAuthRole []*string `json:"UnAuthRole,omitempty" name:"UnAuthRole"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRolePlayStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRolePlayStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAntiFakeUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// 规则ID

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAntiFakeUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAntiFakeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomRuleRequest struct {
	*tchttp.BaseRequest

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 删除的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 删除的规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteCustomRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotRequestLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 索引

		LastIndex *string `json:"LastIndex,omitempty" name:"LastIndex"`
		// 标记本次查询是否将数据查询完全

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 数据

		Data []*BotRequestLogItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotRequestLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRequestLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TsResource struct {

	// 资源id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则状态，0/1

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 子集群id

	SubRegion *string `json:"SubRegion,omitempty" name:"SubRegion"`
	// IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 认证信息

	Auth *string `json:"Auth,omitempty" name:"Auth"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 资源类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeBotRegionsStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各个地域的bot攻击量

		Data []*BotsGeoStatStyledItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotRegionsStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRegionsStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDNSDetectDomainListDomainListItem struct {

	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// hijackRecords

	HijackRecords *string `json:"HijackRecords,omitempty" name:"HijackRecords"`
	// hijackRegions

	HijackRegions *string `json:"HijackRegions,omitempty" name:"HijackRegions"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 授权的IP

	AuthIP []*string `json:"AuthIP,omitempty" name:"AuthIP"`
}

type ModifyJsInjectRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyJsInjectRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJsInjectRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainInfo struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// cname地址

	Cname *string `json:"Cname,omitempty" name:"Cname"`
	// 实例类型,sparta-waf表示saaswaf实例域名,clb-waf表示clbwaf实例域名,cdc-clb-waf表示CDC环境下clbwaf实例域名,cdn-waf表示cdnwaf实例域名

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 日志包

	ClsStatus *uint64 `json:"ClsStatus,omitempty" name:"ClsStatus"`
	// clbwaf使用模式,0镜像模式 1清洗模式

	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`
	// waf开关,0关闭 1开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 规则引擎防护模式,0观察模式 1拦截模式

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 规则引擎和AI引擎防护模式联合状态,10规则引擎观察&&AI引擎关闭模式 11规则引擎观察&&AI引擎观察模式 12规则引擎观察&&AI引擎拦截模式 20规则引擎拦截&&AI引擎关闭模式 21规则引擎拦截&&AI引擎观察模式 22规则引擎拦截&&AI引擎拦截模式

	Engine *uint64 `json:"Engine,omitempty" name:"Engine"`
	// CC列表

	CCList []*string `json:"CCList,omitempty" name:"CCList"`
	// 回源ip

	RsList []*string `json:"RsList,omitempty" name:"RsList"`
	// 服务端口配置

	Ports []*PortInfo `json:"Ports,omitempty" name:"Ports"`
	// 负载均衡器

	LoadBalancerSet []*LoadBalancerPackageNew `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
	// 用户id

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// clbwaf域名监听器状态,0操作成功 4正在绑定LB 6正在解绑LB 7解绑LB失败 8绑定LB失败 10内部错误

	State *int64 `json:"State,omitempty" name:"State"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Ipv6开关状态,0关闭 1开启

	Ipv6Status *int64 `json:"Ipv6Status,omitempty" name:"Ipv6Status"`
	// BOT开关状态,0关闭 1开启

	BotStatus *int64 `json:"BotStatus,omitempty" name:"BotStatus"`
	// 版本信息

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 是否开启投递CLS功能,0关闭 1开启

	PostCLSStatus *int64 `json:"PostCLSStatus,omitempty" name:"PostCLSStatus"`
	// 是否开启投递CKafka功能,0关闭 1开启

	PostCKafkaStatus *int64 `json:"PostCKafkaStatus,omitempty" name:"PostCKafkaStatus"`
	// cdc实例域名接入的集群信息,非cdc实例忽略

	CdcClusters *string `json:"CdcClusters,omitempty" name:"CdcClusters"`
	// api安全开关状态,0关闭 1开启

	ApiStatus *int64 `json:"ApiStatus,omitempty" name:"ApiStatus"`
	// 应用型负载均衡类型,clb或者apisix，默认clb

	AlbType *string `json:"AlbType,omitempty" name:"AlbType"`
	// 安全组状态,0不展示 1非腾讯云源站 2安全组绑定失败 3安全组发生变更

	SgState *int64 `json:"SgState,omitempty" name:"SgState"`
	// 安全组状态的详细解释

	SgDetail *string `json:"SgDetail,omitempty" name:"SgDetail"`
	// 域名类型:hybrid表示混合云域名，public表示公有云域名

	CloudType *string `json:"CloudType,omitempty" name:"CloudType"`
}

type DescribePostPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 时间间隔

		Interval *int64 `json:"Interval,omitempty" name:"Interval"`
		// 时间戳

		Points []*PostPointInfo `json:"Points,omitempty" name:"Points"`
		// 条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 未压缩前的总流量，单位：Bytes字节

		TotalRealTraffic *int64 `json:"TotalRealTraffic,omitempty" name:"TotalRealTraffic"`
		// 压缩后的总流量，单位：Bytes字节

		TotalCompressionTraffic *int64 `json:"TotalCompressionTraffic,omitempty" name:"TotalCompressionTraffic"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePostPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePostPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotTopItem struct {

	// 对应的key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 对应的值

	Value *uint64 `json:"Value,omitempty" name:"Value"`
	// key对应的展示描述语

	Label *string `json:"Label,omitempty" name:"Label"`
}

type UpsertCCAutoStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 状态值

	Value *int64 `json:"Value,omitempty" name:"Value"`
	// 版本：clb-waf, spart-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *UpsertCCAutoStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertCCAutoStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafRequestCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志条数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafRequestCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafRequestCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStracePointsRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 防护类型，不指定则默认为business_risk

	StraceType *string `json:"StraceType,omitempty" name:"StraceType"`
}

func (r *GetStracePointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStracePointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAccessLogTwoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志显示内容

		Data []*BotAccessItemTwo `json:"Data,omitempty" name:"Data"`
		// 日志条数，当传入CountFlag为true的时候显示为正确条数，其他情况显示0

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotAccessLogTwoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAccessLogTwoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防护域名列表的长度

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 防护域名的列表

		HostList []*HostRecord `json:"HostList,omitempty" name:"HostList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiSecurityGenerateDealsRequest struct {
	*tchttp.BaseRequest

	// api安全计费下单入参

	Goods []*Goods `json:"Goods,omitempty" name:"Goods"`
	// 在该实例下购买api安全

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyApiSecurityGenerateDealsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiSecurityGenerateDealsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBusinessRiskRuleRequest struct {
	*tchttp.BaseRequest

	// 风险等级，天御风险类型，包含：reject（高风险），review（人工审核）pass（无风险）

	Level *string `json:"Level,omitempty" name:"Level"`
	// 账户类型：微信、QQ、手机号、手机号MD5、其他

	AccountTypes *string `json:"AccountTypes,omitempty" name:"AccountTypes"`
	// 分类码

	DomainCode *string `json:"DomainCode,omitempty" name:"DomainCode"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 重定向url

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 条件匹配方式

	Extend *string `json:"Extend,omitempty" name:"Extend"`
	// 规则数据库记录Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 规则状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 拦截url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 规则名称，不容许修改

	Name *string `json:"Name,omitempty" name:"Name"`
	// 账户获取方式

	Options *string `json:"Options,omitempty" name:"Options"`
	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 风险类型

	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
	// 规则动作，50（观察）51（验证码）52（拦截）53（重定向）优先级：观察>验证码>重定向>拦截

	RuleAction *int64 `json:"RuleAction,omitempty" name:"RuleAction"`
}

func (r *ModifyBusinessRiskRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBusinessRiskRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PostDomainsConfigTaskRequest struct {
	*tchttp.BaseRequest

	// 数据偏移量，从0开始。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回域名的数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤数组

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *PostDomainsConfigTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PostDomainsConfigTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Amount struct {

	// 周期内最大配额数

	MaxAmount *uint64 `json:"MaxAmount,omitempty" name:"MaxAmount"`
	// 周期描述，支持duration类型的字符串，比如1s, 1m, 1h等

	ValidDuration *string `json:"ValidDuration,omitempty" name:"ValidDuration"`
}

type HavePostFlowServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开通投递服务接口

		HaveService *bool `json:"HaveService,omitempty" name:"HaveService"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HavePostFlowServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HavePostFlowServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserSignatureRule struct {

	// 特征ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 规则开关

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 主类ID

	MainClassID *string `json:"MainClassID,omitempty" name:"MainClassID"`
	// 子类ID

	SubClassID *string `json:"SubClassID,omitempty" name:"SubClassID"`
	// CVE ID

	CveID *string `json:"CveID,omitempty" name:"CveID"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 主类名字，根据Language字段输出中文/英文

	MainClassName *string `json:"MainClassName,omitempty" name:"MainClassName"`
	// 子类名字，根据Language字段输出中文/英文，若子类id为00000000，此字段为空

	SubClassName *string `json:"SubClassName,omitempty" name:"SubClassName"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 0/1

	Reason *int64 `json:"Reason,omitempty" name:"Reason"`
}

type DescribeBotLevelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// score分布

		Score []*BotScoreEntry `json:"Score,omitempty" name:"Score"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotLevelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDNSDetectDataMapRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开始日期

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 结束日期

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeDNSDetectDataMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSDetectDataMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// domain列表

		Domains []*DomainInfo `json:"Domains,omitempty" name:"Domains"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotSceneStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// true-开启 false-关闭

	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyBotSceneStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSceneStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotReqAnalyzeRes struct {

	// 请求重复比重

	ReqQueryRepeatRate *float64 `json:"ReqQueryRepeatRate,omitempty" name:"ReqQueryRepeatRate"`
	// 请求最多的url

	ReqMaxUrl *string `json:"ReqMaxUrl,omitempty" name:"ReqMaxUrl"`
	// url重复比重

	ReqUrlRepeatRate *float64 `json:"ReqUrlRepeatRate,omitempty" name:"ReqUrlRepeatRate"`
	// url种类

	ReqUrlKind *int64 `json:"ReqUrlKind,omitempty" name:"ReqUrlKind"`
	// 最多的请求参数

	ReqMaxQuery *string `json:"ReqMaxQuery,omitempty" name:"ReqMaxQuery"`
	// 是否发生在凌晨

	ReqInWeeHours *bool `json:"ReqInWeeHours,omitempty" name:"ReqInWeeHours"`
}

type CustomPayloadData struct {

	// 载荷

	Res []*CustomPayload `json:"Res,omitempty" name:"Res"`
	// 计数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type DescribeApiSecSensitiveRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自定义敏感检测规则总开关

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 非内置规则的rulename列表

		RuleNameList []*string `json:"RuleNameList,omitempty" name:"RuleNameList"`
		// api提取规则列表

		ApiExtractRule []*ApiSecExtractRule `json:"ApiExtractRule,omitempty" name:"ApiExtractRule"`
		// api敏感规则列表

		Data []*ApiSecSensitiveRule `json:"Data,omitempty" name:"Data"`
		// 规则数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiSecSensitiveRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecSensitiveRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomPayloadsRequest struct {
	*tchttp.BaseRequest

	// 类别

	LearnStat *string `json:"LearnStat,omitempty" name:"LearnStat"`
	// 筛选框列表

	Filters []*AiListFilter `json:"Filters,omitempty" name:"Filters"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 排序

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 偏移

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 翻页限制

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeCustomPayloadsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomPayloadsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitempty" name:"RuleList"`
		// 规则条数

		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Data struct {

	// {"data":{"AutoRenew":"false"}}}

	AutoRenew *string `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

type InOutputUCBRuleEntry struct {

	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 操作符

	Op *string `json:"Op,omitempty" name:"Op"`
	// 值

	Value *UCBEntryValue `json:"Value,omitempty" name:"Value"`
	// 可选的补充操作符

	OpOp *string `json:"OpOp,omitempty" name:"OpOp"`
	// 可选的补充参数

	OpArg []*string `json:"OpArg,omitempty" name:"OpArg"`
	// 可选的补充值

	OpValue *float64 `json:"OpValue,omitempty" name:"OpValue"`
	// Header参数值时使用

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeFindDomainListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 域名信息列表

		List []*FindAllDomainDetail `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFindDomainListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFindDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCopyAreaBanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行状态，返回success或者error

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCopyAreaBanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCopyAreaBanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：非腾讯云自研账号
		// 1:  腾讯云自研账号

		InnerUser *int64 `json:"InnerUser,omitempty" name:"InnerUser"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSpartaProtectionRequest struct {
	*tchttp.BaseRequest

	// 域名回源时的回源域名。UpstreamType为1时，需要填充此字段

	UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`
	// 回源IP列表各IP的权重，和SrcList一一对应。当且仅当UpstreamType为0，并且SrcList有多个IP，并且LoadBalance为2时需要填写，否则填 []

	Weights []*int64 `json:"Weights,omitempty" name:"Weights"`
	// 是否开启主动健康检测。
	// 0：不开启
	// 1：开启

	ActiveCheck *int64 `json:"ActiveCheck,omitempty" name:"ActiveCheck"`
	// 加密套件模板。
	// 0：不支持选择，使用默认模板
	// 1：通用型模板
	// 2：安全型模板
	// 3：自定义模板

	CipherTemplate *int64 `json:"CipherTemplate,omitempty" name:"CipherTemplate"`
	// WAF与源站的写超时时间，默认300s。

	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitempty" name:"ProxySendTimeout"`
	// WAF回源时的SNI类型。
	// 0：关闭SNI，不配置client_hello中的server_name
	// 1：开启SNI，client_hello中的server_name为防护域名
	// 2：开启SNI，SNI为域名回源时的源站域名
	// 3：开启SNI，SNI为自定义域名

	SniType *int64 `json:"SniType,omitempty" name:"SniType"`
	// CertType为1时，需要填充此参数，表示自有证书的私钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段

	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`
	// 是否开启XFF重置。
	// 0：关闭
	// 1：开启

	XFFReset *int64 `json:"XFFReset,omitempty" name:"XFFReset"`
	// 待废弃，可不填。灰度的地区

	GrayAreas []*string `json:"GrayAreas,omitempty" name:"GrayAreas"`
	// CertType为1时，需要填充此参数，表示自有证书的证书链

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// CertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id

	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`
	// TLS版本信息

	TLSVersion *int64 `json:"TLSVersion,omitempty" name:"TLSVersion"`
	// 是否开启WebSocket支持。
	// 0：关闭
	// 1：开启

	IsWebsocket *int64 `json:"IsWebsocket,omitempty" name:"IsWebsocket"`
	// 是否开启HTTP强制跳转到HTTPS。
	// 0：不强制跳转
	// 1：开启强制跳转

	HttpsRewrite *int64 `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`
	// 自定义的加密套件列表。CipherTemplate为3时需要填此字段，表示自定义的加密套件，值通过DescribeCiphersDetail接口获取。

	Ciphers []*int64 `json:"Ciphers,omitempty" name:"Ciphers"`
	// 服务配置有HTTPS端口时，HTTPS的回源协议。
	// http：使用http协议回源，和HttpsUpstreamPort配合使用
	// https：使用https协议回源

	UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`
	// 是否开启长连接。
	// 0： 短连接
	// 1： 长连接

	IsKeepAlive *string `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`
	// 证书类型。
	// 0：仅配置HTTP监听端口，没有证书
	// 1：证书来源为自有证书
	// 2：证书来源为托管证书

	CertType *int64 `json:"CertType,omitempty" name:"CertType"`
	// 待废弃，可不填。Waf的资源ID。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 待废弃，目前填0即可。anycast IP类型开关： 0 普通IP 1 Anycast IP

	Anycast *int64 `json:"Anycast,omitempty" name:"Anycast"`
	// waf前是否部署有七层代理服务。
	// 0：没有部署代理服务
	// 1：有部署代理服务，waf将使用XFF获取客户端IP
	// 2：有部署代理服务，waf将使用remote_addr获取客户端IP
	// 3：有部署代理服务，waf将使用ip_headers中的自定义header获取客户端IP

	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// 服务端口列表配置。
	// NginxServerId：新增域名时填'0'
	// Port：监听端口号
	// Protocol：端口协议
	// UpstreamPort：与Port相同
	// UpstreamProtocol：与Protocol相同

	Ports []*PortItem `json:"Ports,omitempty" name:"Ports"`
	// IP回源时的回源IP列表。UpstreamType为0时，需要填充此字段

	SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`
	// SniType为3时，需要填此参数，表示自定义的SNI；

	SniHost *string `json:"SniHost,omitempty" name:"SniHost"`
	// 回源类型。
	// 0：通过IP回源
	// 1：通过域名回源

	UpstreamType *int64 `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// 是否开启HTTP2，需要开启HTTPS协议支持。
	// 0：关闭
	// 1：开启

	IsHttp2 *int64 `json:"IsHttp2,omitempty" name:"IsHttp2"`
	// 回源负载均衡策略。
	// 0：轮询
	// 1：IP hash
	// 2：加权轮询

	LoadBalance *string `json:"LoadBalance,omitempty" name:"LoadBalance"`
	// 待废弃，可不填。WAF实例类型。
	// sparta-waf：SAAS型WAF
	// clb-waf：负载均衡型WAF
	// cdn-waf：CDN上的Web防护能力

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 域名所属实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// WAF与源站的读超时时间，默认300s。

	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitempty" name:"ProxyReadTimeout"`
	// 需要防护的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 待废弃，可不填。是否开启灰度，0表示不开启灰度。

	IsGray *int64 `json:"IsGray,omitempty" name:"IsGray"`
	// IsCdn为3时，需要填此参数，表示自定义header

	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`
}

func (r *AddSpartaProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSpartaProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FullTextInfo struct {

	// 是否大小写敏感

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
}

type SearchItem struct {

	// 日志开关

	ClsStatus *string `json:"ClsStatus,omitempty" name:"ClsStatus"`
	// waf开关

	Status *string `json:"Status,omitempty" name:"Status"`
	// 流量模式

	FlowMode *string `json:"FlowMode,omitempty" name:"FlowMode"`
}

type ModifyDNSDetectDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDNSDetectDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDNSDetectDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostModeRequest struct {
	*tchttp.BaseRequest

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 0:修改防护模式，1:修改AI

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 实例类型

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *ModifyHostModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotUCBRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 输出结果

		Data *BotListData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotUCBRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotUCBRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAntiFakeUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAntiFakeUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAntiFakeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAttackLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前返回的攻击日志条数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 接口升级，此字段无效，默认返回空字符串

		Context *string `json:"Context,omitempty" name:"Context"`
		// 攻击日志数组条目内容

		Data []*AttackLogInfo `json:"Data,omitempty" name:"Data"`
		// CLS接口返回内容

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// CLS接口返回内容，标志是否启动新版本索引

		SqlFlag *bool `json:"SqlFlag,omitempty" name:"SqlFlag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchAttackLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAttackLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleInfo struct {

	// 全文索引的相关配置

	FullText *FullTextInfo `json:"FullText,omitempty" name:"FullText"`
	// kv 索引的相关配置

	KeyValue *KeyValueInfo `json:"KeyValue,omitempty" name:"KeyValue"`
}

type DescribeAttackTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		Piechart []*PiechartItem `json:"Piechart,omitempty" name:"Piechart"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorDomainsRequest struct {
	*tchttp.BaseRequest

	// 偏移数量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 域名搜索，这里是前缀匹配

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 多个域名列表，这里是完全匹配

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
	// 版本，目前值为sparta-waf和clb-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 一页的数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMonitorDomainsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorDomainsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAreaBanAreasRequest struct {
	*tchttp.BaseRequest

	// 需要修改的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 需要调整的地域信息，一个字符串数组

	Areas []*string `json:"Areas,omitempty" name:"Areas"`
}

func (r *ModifyAreaBanAreasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAreaBanAreasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleDeleteRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Name *string `json:"Name,omitempty" name:"Name"`
	// 删除的规则库id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *RuleDeleteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleDeleteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRateLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 限流操作结果的基本信息

		BaseInfo *RateLimitCommonRsp `json:"BaseInfo,omitempty" name:"BaseInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRateLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRateLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotUCBFeatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据包

		Data *DescribeBotUCBFeatureRuleRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotUCBFeatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotUCBFeatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVipInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VIP信息

		VipInfo []*VipInfo `json:"VipInfo,omitempty" name:"VipInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVipInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVipInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotWhiteStatusRequest struct {
	*tchttp.BaseRequest

	// 上线模式，默认不传

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开关状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// searchengine(搜索引擎机器人)、feedfetcher(订阅机器人）

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyBotWhiteStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotWhiteStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合搜索条件攻击日志的数量

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackLogCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackLogCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteAllDomainRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWhiteAllDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteAllDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotItem struct {

	// 一条日志记录带的id，MongoDB返回

	Id *string `json:"Id,omitempty" name:"Id"`
	// botd的id

	BotId *string `json:"BotId,omitempty" name:"BotId"`
	// 攻击源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 会话数量

	Stat *BotStatSt `json:"Stat,omitempty" name:"Stat"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 分数

	Score *BotScore `json:"Score,omitempty" name:"Score"`
	// 时间戳

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 异常

	BotFeature []*string `json:"BotFeature,omitempty" name:"BotFeature"`
	// 检测时间

	SessionCheckTime *uint64 `json:"SessionCheckTime,omitempty" name:"SessionCheckTime"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeCdcResourceRequest struct {
	*tchttp.BaseRequest

	// 筛选条件，支持OpAppId, ClusterId, Ip，Status，Type；Status = 0/1

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// 分片偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 页面显示个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCdcResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCdcResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHWInfoData struct {

	// 数据状态

	HWState *uint64 `json:"HWState,omitempty" name:"HWState"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 有效域名数

	ValidDomainNum *uint64 `json:"ValidDomainNum,omitempty" name:"ValidDomainNum"`
	// 总域名数

	TotalDomainNum *uint64 `json:"TotalDomainNum,omitempty" name:"TotalDomainNum"`
	// IP黑白名单数上限

	IpBlackListLimit *uint64 `json:"IpBlackListLimit,omitempty" name:"IpBlackListLimit"`
	// 总状态

	TotalStatus *uint64 `json:"TotalStatus,omitempty" name:"TotalStatus"`
	// 有效域名列表

	ValidDomainList []*string `json:"ValidDomainList,omitempty" name:"ValidDomainList"`
	// 威胁情报

	ThreatIntelligence *ThreatIntelligenceData `json:"ThreatIntelligence,omitempty" name:"ThreatIntelligence"`
	// 防护能力定制

	AbilityCustomized *OWASPCustomizedData `json:"AbilityCustomized,omitempty" name:"AbilityCustomized"`
	// 防护深度定制

	MaxBodyDetect *MaxBodyDetect `json:"MaxBodyDetect,omitempty" name:"MaxBodyDetect"`
}

type IpHitItemsData struct {

	// 数组封装

	Res []*IpHitItem `json:"Res,omitempty" name:"Res"`
	// 总数目

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type RuleList struct {

	// 规则Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则列表的id

	Rules []*uint64 `json:"Rules,omitempty" name:"Rules"`
	// 请求url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 请求的方法

	Function *string `json:"Function,omitempty" name:"Function"`
	// 时间戳

	Time *string `json:"Time,omitempty" name:"Time"`
	// 开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type CreateAppletAccessV3Request struct {
	*tchttp.BaseRequest

	// 小程序ID

	AppletID *string `json:"AppletID,omitempty" name:"AppletID"`
}

func (r *CreateAppletAccessV3Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppletAccessV3Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterInstallPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterInstallPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterInstallPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAntiFakeUrlRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// uri

	Uri *string `json:"Uri,omitempty" name:"Uri"`
	// ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyAntiFakeUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAntiFakeUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoDenyDetail struct {

	// 攻击封禁类型标签

	AttackTags []*string `json:"AttackTags,omitempty" name:"AttackTags"`
	// 攻击次数阈值

	AttackThreshold *int64 `json:"AttackThreshold,omitempty" name:"AttackThreshold"`
	// 自动封禁状态

	DefenseStatus *int64 `json:"DefenseStatus,omitempty" name:"DefenseStatus"`
	// 攻击时间阈值

	TimeThreshold *int64 `json:"TimeThreshold,omitempty" name:"TimeThreshold"`
	// 最后更新时间

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 自动封禁时间

	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitempty" name:"DenyTimeThreshold"`
}

type DescribeSpartaProtectionListItem struct {

	// 是否是cdn

	IsCdn *string `json:"IsCdn,omitempty" name:"IsCdn"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 日志开关

	LogStatus *string `json:"LogStatus,omitempty" name:"LogStatus"`
	// AI模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
}

type Info struct {

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 规则描述信息

	Memos []*string `json:"Memos,omitempty" name:"Memos"`
	// 英文版本的规则描述信息

	MemosEng []*string `json:"MemosEng,omitempty" name:"MemosEng"`
}

type DescribeDomainRulesEnRequest struct {
	*tchttp.BaseRequest

	// 需要查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainRulesEnRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainRulesEnRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddCustomPayloadRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 类别，可选的值为good、bad，分别表示误报、漏报

	Category *string `json:"Category,omitempty" name:"Category"`
	// 载荷，最大长度限制为24KB

	Payload *string `json:"Payload,omitempty" name:"Payload"`
	// 来源，可选的值为manual、fp_auto，分别表示手动添加、自动添加

	Source *string `json:"Source,omitempty" name:"Source"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 添加漏报的时候载荷的标记，可选的值sqli、xss、other

	Flag *string `json:"Flag,omitempty" name:"Flag"`
}

func (r *AddCustomPayloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCustomPayloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrafficMarkingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTrafficMarkingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrafficMarkingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotTcbRuleData struct {

	// tcb规则

	Res []*string `json:"Res,omitempty" name:"Res"`
}

type ModifyApiRequestParameterRequest struct {
	*tchttp.BaseRequest

	// 参数名称

	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`
	// api名称

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 参数类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据标签

	Label []*string `json:"Label,omitempty" name:"Label"`
	// 请求方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 参数位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 场景，如果该字段有值，则表示修改场景信息

	Scene *string `json:"Scene,omitempty" name:"Scene"`
}

func (r *ModifyApiRequestParameterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiRequestParameterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesRspRuleListItem struct {

	// 动作类型

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 跳过的策略

	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 重定向地址

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 策略ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 优先级

	SortId *string `json:"SortId,omitempty" name:"SortId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 策略详情

	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 生效状态

	ValidStatus *int64 `json:"ValidStatus,omitempty" name:"ValidStatus"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
}

type AddBatchCustomRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddBatchCustomRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBatchCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackDownloadRecordsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAttackDownloadRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackDownloadRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackTotalCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击总次数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttackTotalCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackTotalCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserDefaultEngineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserDefaultEngineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserDefaultEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWafAutoDenyRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功的状态码，需要JSON解码后再使用，返回的格式是{"域名":"状态"}，成功的状态码为Success，其它的为失败的状态码（yunapi定义的错误码）

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWafAutoDenyRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWafAutoDenyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiSecSensitiveRuleListRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 是否查询api提取规则策略，true表示查询

	IsQueryApiExtractRule *bool `json:"IsQueryApiExtractRule,omitempty" name:"IsQueryApiExtractRule"`
}

func (r *DescribeApiSecSensitiveRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecSensitiveRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCOpenCustomCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCOpenCustomCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCOpenCustomCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableUserSignaturePolicyRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 地域信息

	Areas []*string `json:"Areas,omitempty" name:"Areas"`
}

func (r *DisableUserSignaturePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableUserSignaturePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1:成功，0:请求失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 计费询价响应体

		Data []*PriceData `json:"Data,omitempty" name:"Data"`
		// 返回消息

		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
		// true:首购，false:非首购

		FirstPurchase *bool `json:"FirstPurchase,omitempty" name:"FirstPurchase"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttackLogPostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceAttackLogPostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAttackLogPostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户级别

		UserGrade *string `json:"UserGrade,omitempty" name:"UserGrade"`
		// 1：企业；0：个人

		CustomerType *float64 `json:"CustomerType,omitempty" name:"CustomerType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserGradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllDomainQueryResult struct {

	// clbwaf列表

	Clb []*DomainCls `json:"Clb,omitempty" name:"Clb"`
	// saaswaf域名列表

	Saas []*DomainCls `json:"Saas,omitempty" name:"Saas"`
}

type DescribeMonitorDomainsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的域名个数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 域名列表

		Domains []*MonitorDomainItem `json:"Domains,omitempty" name:"Domains"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorDomainsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafThreatenIntelligenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// WAF 威胁情报封禁信息

		WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitempty" name:"WafThreatenIntelligenceDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafThreatenIntelligenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafThreatenIntelligenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditHWRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditHWRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditHWRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotIdTypeEntity struct {

	// 类型

	Key *string `json:"Key,omitempty" name:"Key"`
	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type ConfigDeliveryData struct {

	// 查询AppId

	CheckAppId *uint64 `json:"CheckAppId,omitempty" name:"CheckAppId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 配置下发状态 1-成功、2-失败

	DeliveryState *uint64 `json:"DeliveryState,omitempty" name:"DeliveryState"`
	// 配置下发状态名称

	DeliveryStateName *string `json:"DeliveryStateName,omitempty" name:"DeliveryStateName"`
	// 下发提示语

	DeliveryMsg *string `json:"DeliveryMsg,omitempty" name:"DeliveryMsg"`
	// DB版本

	DbVersion *uint64 `json:"DbVersion,omitempty" name:"DbVersion"`
	// DbVersion时间

	DbVersionTs *uint64 `json:"DbVersionTs,omitempty" name:"DbVersionTs"`
	// DbVersion时间标准格式

	DbVersionTsStr *string `json:"DbVersionTsStr,omitempty" name:"DbVersionTsStr"`
	// [{"Region":"cd","Version":21}]

	RedisList []*RedisVersionItem `json:"RedisList,omitempty" name:"RedisList"`
	// 实例列表

	InstanceList []*InstanceItem `json:"InstanceList,omitempty" name:"InstanceList"`
}

type AddHWRecordRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *AddHWRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddHWRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotUBRecordsRequest struct {
	*tchttp.BaseRequest

	// 过滤规则名称，""表示全部

	Name *string `json:"Name,omitempty" name:"Name"`
	// IP，""表示全部

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 起始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 偏移

	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`
	// 限制数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeBotUBRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotUBRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainVerifyResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检验状态：0表示可以添加，大于0为不能添加

		VerifyCode *int64 `json:"VerifyCode,omitempty" name:"VerifyCode"`
		// 结果描述；如果可以添加返回空字符串

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainVerifyResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainVerifyResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AiInfo struct {

	// 命中的key，请求特征的维度

	Key *string `json:"Key,omitempty" name:"Key"`
	// 异常值

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type BotRecordItemRes struct {

	// 访问序列数据

	Items *string `json:"Items,omitempty" name:"Items"`
}

type AccessRuleKeyValueInfo struct {

	// 是否大小写敏感

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 需要建立索引的键值对信息；最大只能配置100个键值对

	KeyValues []*AccessKeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type MiniPkg struct {

	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 地域

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 购买数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 续费标志

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 计费项

	BillingItem *string `json:"BillingItem,omitempty" name:"BillingItem"`
	// 小程序接入appid数量

	AccessAppidNum *int64 `json:"AccessAppidNum,omitempty" name:"AccessAppidNum"`
	// ddos防护阈值

	DdosThreshold *int64 `json:"DdosThreshold,omitempty" name:"DdosThreshold"`
}

type DescribeFlowTrendRequest struct {
	*tchttp.BaseRequest

	// 结束时间戳，精度秒

	EndTs *int64 `json:"EndTs,omitempty" name:"EndTs"`
	// 需要获取流量趋势的域名, all表示所有域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 起始时间戳，精度秒

	StartTs *int64 `json:"StartTs,omitempty" name:"StartTs"`
}

func (r *DescribeFlowTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHWInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *DescribeHWInfoData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHWInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHWInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDNSDetectHijackDataRequest struct {
	*tchttp.BaseRequest

	// 域名，如果是全部域名则为“all”

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询的开始日期，格式为：年-月-日

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 查询的结束日期，格式为：年-月-日

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 查询条件，只支持域名查询

	Search *string `json:"Search,omitempty" name:"Search"`
	// 分页参数

	PageInfo *PageInfo `json:"PageInfo,omitempty" name:"PageInfo"`
}

func (r *DescribeDNSDetectHijackDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSDetectHijackDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackAnalysisRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 需要聚合分析的参数key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Lucene语法

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 排序字段,升序asc,降序desc，默认降序desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 返回的top数，默认返回top5

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 查询的域名，所有域名使用all

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *GetAttackAnalysisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackAnalysisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskInfo struct {

	// 风险等级列表
	// 业务安全风险类型，包含：reject（高风险），review（人工审核）pass（无风险）

	RiskLevel []*Entry `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 风险类型列表
	// "1", "账号信用低"
	// "11", "低活跃账号"
	// "2", "垃圾账号"
	// "21", "疑似小号"
	// "22", "违规账号"
	// "3", "无效账号"
	// "4", "黑名单"
	// "5", "白名单"
	// "201", "环境异常"
	// "2011", "非常用IP"
	// "2012", "IP异常"
	// "205", "非公网有效ip"
	// "206", "设备异常"
	// "2061", "非常用设备"
	// "2063", "群控设备"

	RiskType []*Entry `json:"RiskType,omitempty" name:"RiskType"`
}

type CreateClsForAfterpayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -1代表是无效WAF用户，1代表已经购买cls，0代表新开通日志服务

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 错误信息或者日志服务的资源id

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClsForAfterpayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClsForAfterpayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotStatPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势图参数

		Data []*BotStatPointItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotStatPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotStatPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficMarkingRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeTrafficMarkingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficMarkingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BypassSpartaProtectionRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domains *string `json:"Domains,omitempty" name:"Domains"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 用户id

	AppIdInner *int64 `json:"AppIdInner,omitempty" name:"AppIdInner"`
	// 用户id

	UinInner *string `json:"UinInner,omitempty" name:"UinInner"`
}

func (r *BypassSpartaProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BypassSpartaProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiStatPointsRequest struct {
	*tchttp.BaseRequest

	// 过滤器

	Filters []*BotDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开始时间

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 趋势图维度，如attack_type

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

func (r *DescribeApiStatPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiStatPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAggregateDomainStatRequest struct {
	*tchttp.BaseRequest

	// topN的值

	TopNums *int64 `json:"TopNums,omitempty" name:"TopNums"`
	// 起始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 终止时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
}

func (r *DescribeBotAggregateDomainStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAggregateDomainStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAIModelStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeAIModelStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAIModelStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCLSResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志包数量

		Packages *uint64 `json:"Packages,omitempty" name:"Packages"`
		// 已使用容量，单位GB

		Used *float64 `json:"Used,omitempty" name:"Used"`
		// 开通访问日志的域名数量

		ClsDomain *uint64 `json:"ClsDomain,omitempty" name:"ClsDomain"`
		// 客户所有域名数量

		AllDomain *uint64 `json:"AllDomain,omitempty" name:"AllDomain"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCLSResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCLSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainPostActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainPostActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainPostActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostStatusInnerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功的状态码，需要JSON解码后再使用，返回的格式是{"域名":"状态"}，成功的状态码为Success，其它的为失败的状态码（yunapi定义的错误码）

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostStatusInnerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostStatusInnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductInfoRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeProductInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSpartaProtectionAutoRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *AddSpartaProtectionAutoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSpartaProtectionAutoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAccessAggregateTopNRequest struct {
	*tchttp.BaseRequest

	// 需要查询TOP的维度名，见AccessAggregateDimension的key定义

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 查询过滤器

	Filters []*BotDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 开始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 需要的Top数，默认5， 最大值100

	TopN *uint64 `json:"TopN,omitempty" name:"TopN"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotAccessAggregateTopNRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAccessAggregateTopNRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePocTestAuthorizationStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePocTestAuthorizationStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePocTestAuthorizationStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppletAccessRequest struct {
	*tchttp.BaseRequest

	// 小程序接入名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 1   接入配置阶段
	// 2   客户端配置阶段
	// 3   客户端测试阶段

	Stage *uint64 `json:"Stage,omitempty" name:"Stage"`
	// Stage =1 可修改，小程序源站域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Stage =1 可修改，服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// Stage =1 可修改，服务协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// Stage =1 可修改，小程序ID

	AppletID *string `json:"AppletID,omitempty" name:"AppletID"`
	// Stage =1 可修改，详情

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// Stage =2 可修改，客户端配置阶段的状态
	// 0  初始状态
	// 100  完成
	//

	ClientConfigStatus *uint64 `json:"ClientConfigStatus,omitempty" name:"ClientConfigStatus"`
	// Stage =3 可修改，客户端测试阶段的状态
	// 0  初始状态（未验证）
	// 1  已验证，待发布
	// 100  完成（已验证，已发布）

	ClientTestStatus *uint64 `json:"ClientTestStatus,omitempty" name:"ClientTestStatus"`
	// Stage =3 可修改，是否要进行链路测试
	// 0  不进行测试
	// 1  进行测试

	LinkTest *uint64 `json:"LinkTest,omitempty" name:"LinkTest"`
	// Stage =3 可修改，链路测试结果
	// 0  初始状态
	// 1-99  验证失败编码
	// 100  验证成功

	LinkTestStatus *uint64 `json:"LinkTestStatus,omitempty" name:"LinkTestStatus"`
}

func (r *ModifyAppletAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppletAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAccessLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出日志后游标

		Cursor *string `json:"Cursor,omitempty" name:"Cursor"`
		// 实际导出的日志数量

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 实际的访问日志,protobuf的格式，具体格式参考日志服务的，protobuf说明

		Logs *string `json:"Logs,omitempty" name:"Logs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAccessLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAccessLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyObjectRequest struct {
	*tchttp.BaseRequest

	// 修改对象标识

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
	// 新的Waf开关状态，如果和已有状态相同认为修改成功

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 新的实例ID，如果和已绑定的实例相同认为修改成功

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 改动作类型:Status修改开关，InstanceId绑定实例

	OpType *string `json:"OpType,omitempty" name:"OpType"`
}

func (r *ModifyObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCachePathRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteCachePathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCachePathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoDenyIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询IP封禁状态返回结果

		Data *IpHitItemsData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoDenyIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoDenyIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ThreatIntelligenceStatus struct {

	// 情报内容

	Value *string `json:"Value,omitempty" name:"Value"`
	// true：开启、false：关闭

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 情报分类

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeAttackCountRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 查询类型，web为Web攻击，cc是CC攻击

	QueryField *string `json:"QueryField,omitempty" name:"QueryField"`
	// 拦截类型，1为拦截，0为所有

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// WAF版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 获取的域名

	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DescribeAttackCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotRequestCountRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
}

func (r *DescribeBotRequestCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRequestCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDomainWhiteRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchDomainWhiteRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDomainWhiteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBusinessRiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 业务安全返回实体列表

		BusinessRiskList []*BusinessRisk `json:"BusinessRiskList,omitempty" name:"BusinessRiskList"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBusinessRiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBusinessRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityRuleCountRequest struct {
	*tchttp.BaseRequest

	// 需要查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeSecurityRuleCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityRuleCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiRule struct {

	// 用户要查询的api规则的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 该条api规则的描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 该条api的请求方式

	Method *string `json:"Method,omitempty" name:"Method"`
	// 该条api的开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 该条api对应的动作，包括观察和拦截

	Action *uint64 `json:"Action,omitempty" name:"Action"`
	// apid对应的具体的参数信息

	RuleList []*AddApiRule `json:"RuleList,omitempty" name:"RuleList"`
}

type ModifyBotSceneUAStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotSceneUAStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSceneUAStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainInstanceIDsRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainInstanceIDsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainInstanceIDsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePortsRequest struct {
	*tchttp.BaseRequest

	// 实例类型

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribePortsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePortsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPublishOrCancelRuleRequest struct {
	*tchttp.BaseRequest

	// publish时表示打开，offline时表示关闭

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 规则主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyPublishOrCancelRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPublishOrCancelRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveBypassAllRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveBypassAllRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveBypassAllRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotWhiteInfo struct {

	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 策略开关

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 策略描述

	WhiteInfo *string `json:"WhiteInfo,omitempty" name:"WhiteInfo"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest

	// waf的实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotUCBPreinstallRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 类别：1表示协议特征类型的内置规则；2表示IP情报类型的内置规则

	Category *uint64 `json:"Category,omitempty" name:"Category"`
}

func (r *DescribeBotUCBPreinstallRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotUCBPreinstallRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotDetailVersionTwoRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 要查询的type,ip或者token

	Type *string `json:"Type,omitempty" name:"Type"`
	// 要查询的type 对应的值(ip或者token对应的具体的值)

	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *DescribeBotDetailVersionTwoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotDetailVersionTwoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAreaBanStatusRequest struct {
	*tchttp.BaseRequest

	// 需要修改的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 状态值，0表示关闭，1表示开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAreaBanStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAreaBanStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceQpsLimitRequest struct {
	*tchttp.BaseRequest

	// 套餐实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// qps上限

	QpsLimit *int64 `json:"QpsLimit,omitempty" name:"QpsLimit"`
}

func (r *ModifyInstanceQpsLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceQpsLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchIpAccessControlData struct {

	// 总数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 黑白名单条目

	Res []*BatchIpAccessControlItem `json:"Res,omitempty" name:"Res"`
}

type DescribeApiDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 访问ip数

		IpCount *int64 `json:"IpCount,omitempty" name:"IpCount"`
		// 关联事件数

		EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
		// 涉敏数据条数

		SensitiveCount *uint64 `json:"SensitiveCount,omitempty" name:"SensitiveCount"`
		// 响应体

		RspLog *string `json:"RspLog,omitempty" name:"RspLog"`
		// 请求样例，json字符串格式

		Log *string `json:"Log,omitempty" name:"Log"`
		// 当前场景标签

		Scene *string `json:"Scene,omitempty" name:"Scene"`
		// 7天内是否活跃

		IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
		// 风险等级

		Level *uint64 `json:"Level,omitempty" name:"Level"`
		// 昨日访问峰值QPS

		MaxQPS *uint64 `json:"MaxQPS,omitempty" name:"MaxQPS"`
		// 请求参数样例列表

		ParameterList []*ApiParameterType `json:"ParameterList,omitempty" name:"ParameterList"`
		// 敏感字段

		SensitiveFields []*string `json:"SensitiveFields,omitempty" name:"SensitiveFields"`
		// 访问地域数量

		RegionCount *int64 `json:"RegionCount,omitempty" name:"RegionCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInstanceQpsLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹性qps相关值集合

		QpsData *QpsData `json:"QpsData,omitempty" name:"QpsData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInstanceQpsLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstanceQpsLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCachePathRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyCachePathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCachePathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostRequest struct {
	*tchttp.BaseRequest

	// 编辑的域名配置信息

	Host *HostRecord `json:"Host,omitempty" name:"Host"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *ModifyHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAreaBanAreasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAreaBanAreasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAreaBanAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCopyCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 复制类型

	CopyType *int64 `json:"CopyType,omitempty" name:"CopyType"`
	// 来源域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则id数组

	Rules *string `json:"Rules,omitempty" name:"Rules"`
	// 目标域名字符串数组

	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *CreateCopyCustomWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCopyCustomWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAttackDownloadTaskRequest struct {
	*tchttp.BaseRequest

	// 域名，所有域名填写all

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 下载任务名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 拦截状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 自定义策略ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 攻击类型

	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
	// 查询起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 查询结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 风险等级

	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 攻击者IP

	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`
}

func (r *CreateAttackDownloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAttackDownloadTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiFakeUrlRequest struct {
	*tchttp.BaseRequest

	// 翻页参数

	PageInfo *PageInfo `json:"PageInfo,omitempty" name:"PageInfo"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeAntiFakeUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAntiFakeUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotTCBRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// data包装

		Data *BotTcbRuleData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotTCBRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotTCBRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的WAF信息数组的长度，为0则表示没有查询到对应的信息

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 对应的WAF信息的数组。

		HostList []*ClbHostResult `json:"HostList,omitempty" name:"HostList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiDataFilter struct {

	// 数据标签，是否活跃，功能场景

	Entity *string `json:"Entity,omitempty" name:"Entity"`
	// 等于

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 日期，手机号，邮箱等

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeBotAccessLogRequest struct {
	*tchttp.BaseRequest

	// 查询过滤器

	Filters []*BotDataFilter `json:"Filters,omitempty" name:"Filters"`
	// 查询开始时间

	StartTs *int64 `json:"StartTs,omitempty" name:"StartTs"`
	// 查看结束时间

	EndTs *int64 `json:"EndTs,omitempty" name:"EndTs"`
	// 显示第几页，默认传1

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 每页显示的日志数量，默认传10

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否要查询总数，默认为false，即不显示日志总数，主要是为了性能考虑，不能每次查询都返回总数

	CountFlag *bool `json:"CountFlag,omitempty" name:"CountFlag"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotAccessLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAccessLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDomainInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// saas和clb域名信息

		UsersInfo []*UserDomainInfo `json:"UsersInfo,omitempty" name:"UsersInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserDomainInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDomainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserDefaultEngineRequest struct {
	*tchttp.BaseRequest

	// menshen, tiga

	DefaultEngine *string `json:"DefaultEngine,omitempty" name:"DefaultEngine"`
}

func (r *ModifyUserDefaultEngineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserDefaultEngineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotScore struct {

	// 总分数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type DescribeAccessExportsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志导出ID。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志导出列表

		Exports []*ExportAccessInfo `json:"Exports,omitempty" name:"Exports"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessExportsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiAnalyzeStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeApiAnalyzeStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiAnalyzeStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchCustomRuleListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 页尺寸
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段
	// "update_time"-更新时间、"expire_time"-过期时间、"sort_id"-优先级、"id"-规则Id、"create_time"-创建时间

	By *string `json:"By,omitempty" name:"By"`
	// 排序类型
	// desc-降序、asc-升序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 数据类型 "custom-rule"-自定义规则、"custom-white-rule"-精准白名单

	DataType *string `json:"DataType,omitempty" name:"DataType"`
	// 筛选列表

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBatchCustomRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBatchCustomRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostStatusRequest struct {
	*tchttp.BaseRequest

	// 域名状态列表

	HostsStatus []*HostStatus `json:"HostsStatus,omitempty" name:"HostsStatus"`
}

func (r *ModifyHostStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProtectionLevelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProtectionLevelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProtectionLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitempty" name:"RuleList"`
		// 规则条数

		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFindDomainListRequest struct {
	*tchttp.BaseRequest

	// 是否接入waf

	IsWafDomain *string `json:"IsWafDomain,omitempty" name:"IsWafDomain"`
	// 排序参数

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *DescribeFindDomainListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFindDomainListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAntiInfoLeakRulesRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Action 值

	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`
	// 策略数组

	Strategies []*StrategyForAntiInfoLeak `json:"Strategies,omitempty" name:"Strategies"`
	// 规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyAntiInfoLeakRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAntiInfoLeakRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotPkg struct {

	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 地域

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 申请数量

	InquireNum *int64 `json:"InquireNum,omitempty" name:"InquireNum"`
	// 使用数量

	UsedNum *int64 `json:"UsedNum,omitempty" name:"UsedNum"`
	// 子产品code

	Type *string `json:"Type,omitempty" name:"Type"`
	// 续费标志

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 购买页bot6折

	BotCPWaf *int64 `json:"BotCPWaf,omitempty" name:"BotCPWaf"`
	// 控制台买bot5折

	BotNPWaf *int64 `json:"BotNPWaf,omitempty" name:"BotNPWaf"`
	// 7天bot试用标识 1 试用 0 没有试用

	IsBotTrial *int64 `json:"IsBotTrial,omitempty" name:"IsBotTrial"`
}

type AddHWRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddHWRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddHWRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePostFlowServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检查是否开通投递服务接口

		HaveService *bool `json:"HaveService,omitempty" name:"HaveService"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePostFlowServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePostFlowServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotSceneUAStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 开启-true
	// 关闭-false

	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyBotSceneUAStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSceneUAStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRenewFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceRenewFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotIdRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	BotId *string `json:"BotId,omitempty" name:"BotId"`
	// 风险等级筛选

	Level []*int64 `json:"Level,omitempty" name:"Level"`
	// 规则类型筛选

	BotIdType []*string `json:"BotIdType,omitempty" name:"BotIdType"`
	// 规则开关-用于筛选: 0-全部 1-关闭 2-开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 动作类型-用于筛选

	RuleAction []*string `json:"RuleAction,omitempty" name:"RuleAction"`
}

func (r *DescribeBotIdRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotIdRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KVInt struct {

	// Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type UserWhiteRuleItem struct {

	// 匹配域

	MatchField *string `json:"MatchField,omitempty" name:"MatchField"`
	// 匹配方法

	MatchMethod *string `json:"MatchMethod,omitempty" name:"MatchMethod"`
	// 匹配内容

	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
}

type VIPStates struct {

	// 无

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
	// 无

	VIP *string `json:"VIP,omitempty" name:"VIP"`
	// 无

	Type *string `json:"Type,omitempty" name:"Type"`
	// 无

	DstGroupIds []*int64 `json:"DstGroupIds,omitempty" name:"DstGroupIds"`
	// 无

	SrcGroupIds []*int64 `json:"SrcGroupIds,omitempty" name:"SrcGroupIds"`
	// 无

	SrcRSIPs []*string `json:"SrcRSIPs,omitempty" name:"SrcRSIPs"`
	// 无

	DstRSIPs []*string `json:"DstRSIPs,omitempty" name:"DstRSIPs"`
	// 无

	State *int64 `json:"State,omitempty" name:"State"`
	// 无

	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
	// 无

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeBotActionStatRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 起始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 截止时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
}

func (r *DescribeBotActionStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotActionStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Goods struct {

	// 计费类目ID，对应cid

	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`
	// 付费类型，1:预付费，0:后付费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 平台类型，默认1

	Platform *int64 `json:"Platform,omitempty" name:"Platform"`
	// 商品数量

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 商品明细

	GoodsDetail *GoodsDetail `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 购买waf实例区域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 默认为0

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type MonitorInstanceItem struct {

	// 实例ID

	Instance *string `json:"Instance,omitempty" name:"Instance"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 用户自己的Appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 实例对应的版本信息，0代表saas，1代表clb

	Edition *int64 `json:"Edition,omitempty" name:"Edition"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type CopyBotsUCBFeatureRuleRsp struct {

	// 0表示成功，其他表示失败

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 成功的数目

	SuccessNum *int64 `json:"SuccessNum,omitempty" name:"SuccessNum"`
	// 失败的数目

	FailedNum *int64 `json:"FailedNum,omitempty" name:"FailedNum"`
	// 成功的域名列表

	T []*string `json:"T,omitempty" name:"T"`
}

type BotHttpHeader struct {

	// accept存在性信息

	AcceptExist *bool `json:"AcceptExist,omitempty" name:"AcceptExist"`
	// AcceptLanguage存在性

	AcceptLanguageExist *bool `json:"AcceptLanguageExist,omitempty" name:"AcceptLanguageExist"`
	// AcceptEncoding存在性

	AcceptEncodingExist *bool `json:"AcceptEncodingExist,omitempty" name:"AcceptEncodingExist"`
	// Connection存在性

	ConnectionExist *bool `json:"ConnectionExist,omitempty" name:"ConnectionExist"`
}

type LoadBalancerPackageNew struct {

	// 监听id

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听名

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 负载均衡id

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡名

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 地区

	Region *string `json:"Region,omitempty" name:"Region"`
	// 接入IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 接入端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// VPCID

	NumericalVpcId *int64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`
	// CLB类型

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 负载均衡器的域名

	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitempty" name:"LoadBalancerDomain"`
}

type DescribeBotSessionKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 时间戳

		TsVersion *int64 `json:"TsVersion,omitempty" name:"TsVersion"`
		// Session位置

		Category *string `json:"Category,omitempty" name:"Category"`
		// Session-key配置

		Key *string `json:"Key,omitempty" name:"Key"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotSessionKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSessionKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDNSDetectDataMapRspItem struct {

	// 地域

	Key *string `json:"Key,omitempty" name:"Key"`
	// 地域的劫持数量

	Value *string `json:"Value,omitempty" name:"Value"`
}

type IpInfo struct {

	// 攻击源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// ip所有者

	IpOwner *string `json:"IpOwner,omitempty" name:"IpOwner"`
	// ip类型

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// ip地域

	IpArea *string `json:"IpArea,omitempty" name:"IpArea"`
	// ip运营商信息

	IpOperator *string `json:"IpOperator,omitempty" name:"IpOperator"`
}

type Rule struct {

	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 规则等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则防护的CVE编号

	CVE *string `json:"CVE,omitempty" name:"CVE"`
	// 规则的状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 规则修改的时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 门神规则新增/更新时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
}

type DescribeHostLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回的状态码

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserSignaturePolicyStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *GetUserSignaturePolicyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserSignaturePolicyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTsResourceRequest struct {
	*tchttp.BaseRequest

	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// redis passwd

	Auth *string `json:"Auth,omitempty" name:"Auth"`
	// Resource Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// 0: disable 1:enable

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 序号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Sub Region id

	SubRegion *string `json:"SubRegion,omitempty" name:"SubRegion"`
	// Ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *ModifyTsResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTsResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotUaInfo struct {

	// ua存在性

	UaExist *bool `json:"UaExist,omitempty" name:"UaExist"`
	// ua随机数指数

	UaKindRandomRate *float64 `json:"UaKindRandomRate,omitempty" name:"UaKindRandomRate"`
	// ua种类

	UaKind *uint64 `json:"UaKind,omitempty" name:"UaKind"`
	// ua类型

	UaType *string `json:"UaType,omitempty" name:"UaType"`
	// 请求最多的ua

	UaMax *string `json:"UaMax,omitempty" name:"UaMax"`
	// ua存在比

	UaValidRate *float64 `json:"UaValidRate,omitempty" name:"UaValidRate"`
}

type GetStracePointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势图的数据

		Points []*string `json:"Points,omitempty" name:"Points"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStracePointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStracePointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpUserSignaturePolicy struct {

	// appid

	OpAppId *uint64 `json:"OpAppId,omitempty" name:"OpAppId"`
	// 域名

	OpDomain *string `json:"OpDomain,omitempty" name:"OpDomain"`
	// 防护等级 300=标准 400=扩展

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 防护状态 0=关闭，1=开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type DescribeBotCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpartaProtectionInfoRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 实例

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeSpartaProtectionInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpartaProtectionInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDomainListData struct {

	// 域名列表

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

type DeleteDomainWhiteRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 出参

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDomainWhiteRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDomainWhiteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotListScore struct {

	// 得分

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type BotRecordItem struct {

	// 压缩的字符串

	Items *string `json:"Items,omitempty" name:"Items"`
}

type ModifyBotSceneActionRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotSceneActionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSceneActionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshAccessCheckResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefreshAccessCheckResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshAccessCheckResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReloadConfigDeliveryRequest struct {
	*tchttp.BaseRequest

	// 要下发配置的AppId

	CheckAppId *uint64 `json:"CheckAppId,omitempty" name:"CheckAppId"`
	// 要下发配置的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ReloadConfigDeliveryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReloadConfigDeliveryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeObjectsRequest struct {
	*tchttp.BaseRequest

	// 支持的过滤器:
	// 	ObjectId: clb实例ID
	// 	VIP: clb实例的公网IP
	// 	InstanceId: waf实例ID
	// 	Domain: 精准域名
	// 	Status: waf防护开关状态: 0关闭，1开启
	// 	ClsStatus: waf日志开关: 0关闭，1开启

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeObjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRolePlayStateRequest struct {
	*tchttp.BaseRequest

	// 角色名称

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DescribeRolePlayStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRolePlayStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotSceneOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// BOT总开关

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 场景总数

		SceneCount *int64 `json:"SceneCount,omitempty" name:"SceneCount"`
		// 生效场景数

		ValidSceneCount *int64 `json:"ValidSceneCount,omitempty" name:"ValidSceneCount"`
		// 当前开启的、匹配范围为全局、优先级最高的场景

		CurrentGlobalScene *GlobalSceneInfo `json:"CurrentGlobalScene,omitempty" name:"CurrentGlobalScene"`
		// 自定义规则总数，不包括BOT白名单

		CustomRuleNums *int64 `json:"CustomRuleNums,omitempty" name:"CustomRuleNums"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotSceneOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSceneOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttackLogInfo struct {

	// 攻击日志的详情内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// CLS返回内容

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// CLS返回内容

	Source *string `json:"Source,omitempty" name:"Source"`
	// CLS返回内容

	TimeStamp *string `json:"TimeStamp,omitempty" name:"TimeStamp"`
}

type DomainBotStatus struct {

	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DescribeBlockPagesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBlockPagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockPagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiStatPointItem struct {

	// 横坐标

	TimeStamp *string `json:"TimeStamp,omitempty" name:"TimeStamp"`
	// 纵坐标

	Value *uint64 `json:"Value,omitempty" name:"Value"`
	// 类型

	Key *string `json:"Key,omitempty" name:"Key"`
}

type ModifyProtectionStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProtectionStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProtectionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAttackLogRequest struct {
	*tchttp.BaseRequest

	// 接口升级，这个字段传空字符串,翻页使用Page字段

	Context *string `json:"Context,omitempty" name:"Context"`
	// 查询的数量，默认10条，最多100条

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// Lucene语法

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 默认为desc，可以取值desc和asc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 第几页，从0开始

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 查询的域名，所有域名使用all

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *SearchAttackLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAttackLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpsertCCAutoStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正常情况为null

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpsertCCAutoStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertCCAutoStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FindAllDomainDetail struct {

	// 用户id

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ip

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
	// 发现时间

	FindTime *string `json:"FindTime,omitempty" name:"FindTime"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// waf类型

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 是否接入waf

	IsWafDomain *uint64 `json:"IsWafDomain,omitempty" name:"IsWafDomain"`
}

type AddPocTestTaskRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *AddPocTestTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPocTestTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIpAccessControlRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 删除的ip数组

	Items []*string `json:"Items,omitempty" name:"Items"`
	// 是否删除对应的域名下的所有黑/白IP名单，true表示全部删除，false表示只删除指定ip名单

	DeleteAll *bool `json:"DeleteAll,omitempty" name:"DeleteAll"`
	// 是否为多域名黑白名单

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 若IsId字段为True，则Items列表元素需为Id，否则为IP

	IsId *bool `json:"IsId,omitempty" name:"IsId"`
}

func (r *DeleteIpAccessControlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIpAccessControlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleUpdateLog struct {

	// id值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 详细修改动态

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// en/cn 描述信息语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 版本信息

	LogVersion *string `json:"LogVersion,omitempty" name:"LogVersion"`
	// 开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type CreateAccessExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志导出ID。

		ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePeakPointsRequest struct {
	*tchttp.BaseRequest

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// WAF实例ID，不传则不过滤

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 十二个值可选：
	// access-峰值qps趋势图
	// botAccess- bot峰值qps趋势图
	// down-下行峰值带宽趋势图
	// up-上行峰值带宽趋势图
	// attack-Web攻击总数趋势图
	// cc-CC攻击总数趋势图
	// bw-黑IP攻击总数趋势图
	// tamper-防篡改攻击总数趋势图
	// leak-防泄露攻击总数趋势图
	// acl-访问控制攻击总数趋势图
	// http_status-状态码各次数趋势图
	// wx_access-微信小程序峰值qps趋势图

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 查询起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 查询终止时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 查询的域名，如果查询所有域名数据，该参数不填写

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribePeakPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePeakPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotDetailDownloadTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建/删除的下载任务ID

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotDetailDownloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotDetailDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotUCBPreinstallRule struct {

	// 动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 附加参数

	AdditionArg *string `json:"AdditionArg,omitempty" name:"AdditionArg"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 开关

	OnOff *string `json:"OnOff,omitempty" name:"OnOff"`
	// 时间戳

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 有效时间

	ValidTime *uint64 `json:"ValidTime,omitempty" name:"ValidTime"`
	// 优先级

	Prior *uint64 `json:"Prior,omitempty" name:"Prior"`
}

type DescribeApiDetailRequest struct {
	*tchttp.BaseRequest

	// Api名称

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 请求方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeApiDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpUserSignatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则详细

		Rules []*UserSignatureRule `json:"Rules,omitempty" name:"Rules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpUserSignatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpUserSignatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Strategy struct {

	// 匹配字段
	//
	//     匹配字段不同，相应的匹配参数、逻辑符号、匹配内容有所不同具体如下所示：
	// <table><thead><tr><th>匹配字段</th><th>匹配参数</th><th>逻辑符号</th><th>匹配内容</th></tr></thead><tbody><tr><td>IP（来源IP）</td><td>不支持参数</td><td>ipmatch（匹配）<br/>ipnmatch（不匹配）</td><td>多个IP以英文逗号隔开,最多20个</td></tr><tr><td>IPV6（来源IPv6）</td><td>不支持参数</td><td>ipmatch（匹配）<br/>ipnmatch（不匹配）</td><td>支持单个IPV6地址</td></tr><tr><td>Referer（Referer）</td><td>不支持参数</td><td>empty（内容为空）<br/>null（不存在）<br/>eq（等于）<br/>neq（不等于）<br/>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）</td><td>请输入内容,512个字符以内</td></tr><tr><td>URL（请求路径）</td><td>不支持参数</td><td>eq（等于）<br/>neq（不等于）<br/>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）<br/></td><td>请以/开头,512个字符以内</td></tr><tr><td>UserAgent（UserAgent）</td><td>不支持参数</td><td>同匹配字段<font color="Red">Referer</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>HTTP_METHOD（HTTP请求方法）</td><td>不支持参数</td><td>eq（等于）<br/>neq（不等于）</td><td>请输入方法名称,建议大写</td></tr><tr><td>QUERY_STRING（请求字符串）</td><td>不支持参数</td><td>同匹配字段<font color="Red">请求路径</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET（GET参数值）</td><td>支持参数录入</td><td>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET_PARAMS_NAMES（GET参数名）</td><td>不支持参数</td><td>exsit（存在参数）<br/>nexsit（不存在参数）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）</td><td>请输入内容,512个字符以内</td></tr><tr><td>POST（POST参数值）</td><td>支持参数录入</td><td>同匹配字段<font color="Red">GET参数值</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET_POST_NAMES（POST参数名）</td><td>不支持参数</td><td>同匹配字段<font color="Red">GET参数名</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>POST_BODY（完整BODY）</td><td>不支持参数</td><td>同匹配字段<font color="Red">请求路径</font>逻辑符号</td><td>请输入BODY内容,512个字符以内</td></tr><tr><td>COOKIE（Cookie）</td><td>不支持参数</td><td>empty（内容为空）<br/>null（不存在）<br/>rematch（正则匹配）</td><td><font color="Red">暂不支持</font></td></tr><tr><td>GET_COOKIES_NAMES（Cookie参数名）</td><td>不支持参数</td><td>同匹配字段<font color="Red">GET参数名</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>ARGS_COOKIE（Cookie参数值）</td><td>支持参数录入</td><td>同匹配字段<font color="Red">GET参数值</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET_HEADERS_NAMES（Header参数名）</td><td>不支持参数</td><td>exsit（存在参数）<br/>nexsit（不存在参数）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）</td><td>请输入内容,建议小写,512个字符以内</td></tr><tr><td>ARGS_HEADER（Header参数值）</td><td>支持参数录入</td><td>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）</td><td>请输入内容,512个字符以内</td></tr></tbody></table>

	Field *string `json:"Field,omitempty" name:"Field"`
	// 逻辑符号
	//
	//     逻辑符号一共分为以下几种类型：
	//         empty （ 内容为空）
	//         null （不存在）
	//         eq （ 等于）
	//         neq （ 不等于）
	//         contains （ 包含）
	//         ncontains （ 不包含）
	//         strprefix （ 前缀匹配）
	//         strsuffix （ 后缀匹配）
	//         len_eq （ 长度等于）
	//         len_gt （ 长度大于）
	//         len_lt （ 长度小于）
	//         ipmatch （ 属于）
	//         ipnmatch （ 不属于）
	//     各匹配字段对应的逻辑符号不同，详见上述匹配字段表格
	//

	CompareFunc *string `json:"CompareFunc,omitempty" name:"CompareFunc"`
	// 匹配内容
	//
	//     目前 当匹配字段为COOKIE（Cookie）时，不需要输入 匹配内容其他都需要
	//

	Content *string `json:"Content,omitempty" name:"Content"`
	// 匹配参数
	//
	//     配置参数一共分2种类型 不支持参数与支持参数
	//     当匹配字段为以下4个时，匹配参数才能录入，否则不支持该参数
	//         GET（GET参数值）
	//         POST（POST参数值）
	//         ARGS_COOKIE（Cookie参数值）
	//         ARGS_HEADER（Header参数值）
	//

	Arg *string `json:"Arg,omitempty" name:"Arg"`
}

type DeleteBotSceneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBotSceneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBotSceneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotAiRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开关状态 true为开 false为关

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 模式，如试运行和运行状态

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 域名的InstanceID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 表示是否删除自定义策略规则

	DelFlag *bool `json:"DelFlag,omitempty" name:"DelFlag"`
	// 用户自定义ai策略

	BotAiRule *BotAiRule `json:"BotAiRule,omitempty" name:"BotAiRule"`
	// 场景化ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 版本号，场景化版本为"4.1.0"

	BotVersion *string `json:"BotVersion,omitempty" name:"BotVersion"`
}

func (r *ModifyBotAiRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotAiRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleRollbackRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Name *string `json:"Name,omitempty" name:"Name"`
	// 回滚的规则库id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *RuleRollbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleRollbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessHistogramResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 柱状图间隔时间差，单位ms

		Interval *int64 `json:"Interval,omitempty" name:"Interval"`
		// 满足条件的日志条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 注意：此字段可能返回 null，表示取不到有效值

		HistogramInfos []*AccessHistogramItem `json:"HistogramInfos,omitempty" name:"HistogramInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessHistogramResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainCountInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDomainCountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainCountInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiSecSensitiveRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiSecSensitiveRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiSecSensitiveRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyValueInfo struct {

	// 是否大小写敏感

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 支持检索的key的名字

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
	// 字段内容的正则匹配

	Tokenizers []*string `json:"Tokenizers,omitempty" name:"Tokenizers"`
	// 上面 key 对应的类型，一一对应，目前支持long double text

	Types []*string `json:"Types,omitempty" name:"Types"`
	// 是否开启统计功能，并用于控制是否支持搜索字段联想

	SqlFlags []*bool `json:"SqlFlags,omitempty" name:"SqlFlags"`
	// TemplateType模版类型

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
}

type UserSceneInfo struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
	// 场景名称

	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`
	// 场景的动作策略列表

	SceneActionList []*ActionRuleInfo `json:"SceneActionList,omitempty" name:"SceneActionList"`
	// 场景的优先级

	ScenePriority *int64 `json:"ScenePriority,omitempty" name:"ScenePriority"`
}

type DeleteSessionRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// clb-waf 或者 sprta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 要删除的SessionID

	SessionID *int64 `json:"SessionID,omitempty" name:"SessionID"`
}

func (r *DeleteSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotTCBRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 动作

	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

func (r *ModifyBotTCBRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotTCBRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInstanceQpsLimitRequest struct {
	*tchttp.BaseRequest

	// 套餐实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 套餐类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *GetInstanceQpsLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstanceQpsLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpiredDomainPackage struct {

	// 资源ID

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 是否自动续费，1：自动续费，0：不自动续费

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type DescribeHostLimitRequest struct {
	*tchttp.BaseRequest

	// 添加的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 流量来源

	AlbType *string `json:"AlbType,omitempty" name:"AlbType"`
}

func (r *DescribeHostLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiSecEventTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件类型列表显示

		Data []*ApiSecEventKeyValue `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiSecEventTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecEventTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotItemsRequest struct {
	*tchttp.BaseRequest

	// 动作

	BotAction *string `json:"BotAction,omitempty" name:"BotAction"`
	// 过滤条件

	BotFeature []*string `json:"BotFeature,omitempty" name:"BotFeature"`
	// 页面索引

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页面大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否需要查询总数

	NeedTotalCount *bool `json:"NeedTotalCount,omitempty" name:"NeedTotalCount"`
	// 开始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 类型

	BotType *string `json:"BotType,omitempty" name:"BotType"`
	// 排序方法

	Sort []*string `json:"Sort,omitempty" name:"Sort"`
	// 搜索条件

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *DescribeBotItemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotItemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotRecordItemsRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeBotRecordItemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRecordItemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCRuleListRequest struct {
	*tchttp.BaseRequest

	// 需要查询的API所属的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// asc或者desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 目前支持根据ts_version排序

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeCCRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafAutoDenyStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWafAutoDenyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafAutoDenyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyBotTCBRuleRsp struct {

	// 0表示成功，其他表示失败

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 成功的数目

	SuccessNum *int64 `json:"SuccessNum,omitempty" name:"SuccessNum"`
	// 失败的数目

	FailedNum *int64 `json:"FailedNum,omitempty" name:"FailedNum"`
	// 成功的域名列表

	T []*string `json:"T,omitempty" name:"T"`
}

type DeleteAppletAccessV3Request struct {
	*tchttp.BaseRequest

	// 小程序ID

	AppletID *string `json:"AppletID,omitempty" name:"AppletID"`
}

func (r *DeleteAppletAccessV3Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppletAccessV3Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTsRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTsRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTsRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UAStrategy struct {

	// UA策略名称

	UAName *string `json:"UAName,omitempty" name:"UAName"`
	// 对应开关状态

	Status *bool `json:"Status,omitempty" name:"Status"`
}

type DescribeOpAttackWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则列表

		List []*OpUserWhiteRule `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpAttackWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpAttackWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditOpRuleUpdateLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditOpRuleUpdateLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditOpRuleUpdateLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddDomainSgTaskRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 实例

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *AddDomainSgTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDomainSgTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackDownloadRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载攻击日志记录数组

		Records []*DownloadAttackRecordInfo `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttackDownloadRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackDownloadRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpartUserRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeSpartUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpartUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostFlowModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功的状态码

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostFlowModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostFlowModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccessExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHybridClusterNodesRequest struct {
	*tchttp.BaseRequest

	// 实例ID，可以为空。但是InstanceId和ClusterId不能同时为空

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 集群ID，可以为空。但是InstanceId和ClusterId不能同时为空

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeHybridClusterNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHybridClusterNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiSecSensitiveRuleRequest struct {
	*tchttp.BaseRequest

	// 1表示开，0表示关，3表示删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 客户自定义配置

	CustomRule *ApiSecCustomSensitiveRule `json:"CustomRule,omitempty" name:"CustomRule"`
	// rulename列表，批量操作的时候填改值

	RuleNameList []*string `json:"RuleNameList,omitempty" name:"RuleNameList"`
	// api提取规则内容

	CustomApiExtractRule *ApiSecExtractRule `json:"CustomApiExtractRule,omitempty" name:"CustomApiExtractRule"`
	// 批量操作的时候的api提取规则

	ApiExtractRuleName []*string `json:"ApiExtractRuleName,omitempty" name:"ApiExtractRuleName"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *ModifyApiSecSensitiveRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiSecSensitiveRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJsInjectRuleRequest struct {
	*tchttp.BaseRequest

	// 动作为重定向的时候重定向的URL

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 是否开启自动化工具识别，true：是，false：否

	AutomatedTool *bool `json:"AutomatedTool,omitempty" name:"AutomatedTool"`
	// 是否开启页面防调试，true：是，false：否

	CrackBehavior *bool `json:"CrackBehavior,omitempty" name:"CrackBehavior"`
	// 搜索引擎白名单，可选的值为：["sogou", "baidu", "yandex", "360", "yahoo", "bing", "google"]

	GoodBot []*string `json:"GoodBot,omitempty" name:"GoodBot"`
	// 前端对抗的开关，0（关闭） 1（开启）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 规则对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则ID

	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
	// 规则内部ID

	RasSiteID *string `json:"RasSiteID,omitempty" name:"RasSiteID"`
	// 前端对抗的防护URL，老版本使用，已经废弃

	CheckURL *string `json:"CheckURL,omitempty" name:"CheckURL"`
	// 前端对抗的防护URL，支持设置多个防护URL

	CheckUrls []*string `json:"CheckUrls,omitempty" name:"CheckUrls"`
	// 可选的值为：0（放行） 1（阻断） 2（人机识别） 3（观察） 4（重定向）

	RuleAction *uint64 `json:"RuleAction,omitempty" name:"RuleAction"`
}

func (r *ModifyJsInjectRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJsInjectRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotTCBRecordsRequest struct {
	*tchttp.BaseRequest

	// 结束时间戳

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 限制数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 动作

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 起始时间戳

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 偏移

	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`
	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 过滤规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeBotTCBRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotTCBRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量，取Limit整数倍。最小值为0，最大值= Total/Limit向上取整

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回域名的数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤数组

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DealData struct {

	// 订单号列表，元素个数与请求包的goods数组的元素个数一致，商品详情与订单按顺序对应

	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
	// 大订单号，一个大订单号下可以有多个子订单，说明是同一次下单[{},{}]

	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`
}

type DescribeTlsVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TLS key value

		TLS []*TLSVersion `json:"TLS,omitempty" name:"TLS"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTlsVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTlsVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBotUCBFeatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正常情况下为null

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBotUCBFeatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBotUCBFeatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAreaBanAreasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAreaBanAreasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAreaBanAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRewriteCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRewriteCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRewriteCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldWriteConfig struct {

	// 1:开启 0:不开启

	EnableHeaders *int64 `json:"EnableHeaders,omitempty" name:"EnableHeaders"`
	// 1:开启 0:不开启

	EnableBody *int64 `json:"EnableBody,omitempty" name:"EnableBody"`
}

type CreateWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 需要查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 排序字段，支持user_id, signature_id, modify_time

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持SignatureId, MatchContent

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAttackWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateDealsAndPayNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费下单响应结构体

		Data *DealData `json:"Data,omitempty" name:"Data"`
		// 1:成功，0:失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 返回message

		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
		// 购买的实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateDealsAndPayNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateDealsAndPayNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiSecurityGenerateDealsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// api安全计费下单响应结构体

		Data *DealData `json:"Data,omitempty" name:"Data"`
		// 1:成功；0:失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 返回message

		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiSecurityGenerateDealsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiSecurityGenerateDealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAntiFakeUrlRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteAntiFakeUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAntiFakeUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBusinessRiskRequest struct {
	*tchttp.BaseRequest

	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 策略开关，0（关闭），1（开启）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 拦截url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 规则动作，50（观察）51（验证码）52（拦截）53（重定向）优先级：观察>验证码>重定向>拦截

	RuleAction *int64 `json:"RuleAction,omitempty" name:"RuleAction"`
	// 分页页码

	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 分页每页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeBusinessRiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBusinessRiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePriceRequest struct {
	*tchttp.BaseRequest

	// 计费询价resInfo实体

	ResInfo []*Goods `json:"ResInfo,omitempty" name:"ResInfo"`
}

func (r *DescribePriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotGenerateDealsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// bot计费下单响应结构体

		Data *DealData `json:"Data,omitempty" name:"Data"`
		// 1:成功；0:失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 返回message

		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotGenerateDealsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotGenerateDealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserLevelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserLevelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHostInnerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHostInnerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHostInnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateVerifyResultRequest struct {
	*tchttp.BaseRequest

	// 证书类型

	CertType *int64 `json:"CertType,omitempty" name:"CertType"`
	// 证书公钥

	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`
	// 证书ID

	CertID *string `json:"CertID,omitempty" name:"CertID"`
	// 私钥信息

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeCertificateVerifyResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCertificateVerifyResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TLSVersion struct {

	// TLSVERSION的ID

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// TLSVERSION的NAME

	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

type GetStraceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防护总量

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStraceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStraceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotRequestLogItem struct {

	// 时间戳

	TimeStamp *string `json:"TimeStamp,omitempty" name:"TimeStamp"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
}

type RuleEn struct {

	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 规则等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则防护的CVE编号

	CVE *string `json:"CVE,omitempty" name:"CVE"`
	// 规则的状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 规则修改的时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 规则类型英文

	TypeEng *string `json:"TypeEng,omitempty" name:"TypeEng"`
	// 规则等级英文

	LevelEng *string `json:"LevelEng,omitempty" name:"LevelEng"`
	// 规则描述英文

	DescriptionEng *string `json:"DescriptionEng,omitempty" name:"DescriptionEng"`
	// 门神规则新增/更新时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
}

type CertDomain struct {

	// 客户的域名列表

	Domains []*CertDomainItem `json:"Domains,omitempty" name:"Domains"`
	// 该域名列表绑定的sslid

	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type DescribeBotFirstPurchaseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true:首购用户，false:非首购用户

		First *bool `json:"First,omitempty" name:"First"`
		// 1：执行成功 0：执行失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotFirstPurchaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotFirstPurchaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PriceDetail struct {

	// 单价

	SinglePrice *float64 `json:"SinglePrice,omitempty" name:"SinglePrice"`
	// 使用数量

	UsedAmount *int64 `json:"UsedAmount,omitempty" name:"UsedAmount"`
	// 价格

	Cost *float64 `json:"Cost,omitempty" name:"Cost"`
}

type DescribeHybridClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息. 为null表示集群不存在

		Cluster *HybridCluster `json:"Cluster,omitempty" name:"Cluster"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHybridClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHybridClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessFieldValueRatioInfo struct {

	// 日志条数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 对应的Value值的百分比

	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
	// 字段对应的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CdcRegion struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 该地域对应的集群信息

	Clusters []*CdcCluster `json:"Clusters,omitempty" name:"Clusters"`
}

type ModifyAppletAccessV3Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAppletAccessV3Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppletAccessV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoticeInstanceChangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NoticeInstanceChangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NoticeInstanceChangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotAutoStatisticRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关状态

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 用户自定义配置列表

		AutoStatRule []*BotAutoStatRule `json:"AutoStatRule,omitempty" name:"AutoStatRule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotAutoStatisticRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotAutoStatisticRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAntiInfoLeakRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAntiInfoLeakRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAntiInfoLeakRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomRuleRequest struct {
	*tchttp.BaseRequest

	// 编辑的规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	// 默认是100

	SortId *uint64 `json:"SortId,omitempty" name:"SortId"`
	// 匹配条件数组

	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`
	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	// 默认是0

	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 编辑的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 编辑的规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 执行动作，0：放行、1：阻断、2：人机识别、3：观察、4：重定向

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
	// 动作为重定向的时候重定向URL，默认为"/"

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果多个勾选那么以,串接。
	// 默认是"geoip,cc,owasp,ai,antileakage"

	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
}

func (r *ModifyCustomRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 续费开关

	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyInstanceRenewFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRenewFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddBatchAreaBanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作成功

		Res *string `json:"Res,omitempty" name:"Res"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddBatchAreaBanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBatchAreaBanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClsForAfterpayRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateClsForAfterpayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClsForAfterpayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInfo struct {

	// 全文索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。

	FullText *AccessFullTextInfo `json:"FullText,omitempty" name:"FullText"`
	// 键值索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。

	KeyValue *AccessRuleKeyValueInfo `json:"KeyValue,omitempty" name:"KeyValue"`
	// 元字段索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。

	Tag *AccessRuleTagInfo `json:"Tag,omitempty" name:"Tag"`
}

type DescribeDomainSecurityGroupTipsCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDomainSecurityGroupTipsCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainSecurityGroupTipsCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSceneJsInjectRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 前端对抗的规则

		Rule *SceneJsInjectRule `json:"Rule,omitempty" name:"Rule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSceneJsInjectRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSceneJsInjectRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotAiRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 2表示有实例未购买bot套餐，不能开启ai功能

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotAiRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotAiRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiSecEventOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 今日新增

		TodayCount *uint64 `json:"TodayCount,omitempty" name:"TodayCount"`
		// 处置中数量

		HandleCount *uint64 `json:"HandleCount,omitempty" name:"HandleCount"`
		// 已忽略数量

		IgnoreCount *uint64 `json:"IgnoreCount,omitempty" name:"IgnoreCount"`
		// 低危数量

		LowCount *uint64 `json:"LowCount,omitempty" name:"LowCount"`
		// 中等数量

		MediumCount *uint64 `json:"MediumCount,omitempty" name:"MediumCount"`
		// 高危数量

		HighCount *uint64 `json:"HighCount,omitempty" name:"HighCount"`
		// 事件总数

		EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
		// 新发现

		NewCount *uint64 `json:"NewCount,omitempty" name:"NewCount"`
		// 已处置数量

		HandledCount *uint64 `json:"HandledCount,omitempty" name:"HandledCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiSecEventOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiSecEventOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePocTestAuthorizationStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 授权状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePocTestAuthorizationStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePocTestAuthorizationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Entry struct {

	// key值

	Id *string `json:"Id,omitempty" name:"Id"`
	// value值

	Label *string `json:"Label,omitempty" name:"Label"`
}

type JsInjectRule struct {

	// 规则对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 可选的值为：0（放行） 1（阻断） 2（人机识别） 3（观察） 4（重定向）

	RuleAction *int64 `json:"RuleAction,omitempty" name:"RuleAction"`
	// 动作为重定向的时候重定向的URL

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 是否开启自动化工具识别，true：是，false：否

	AutomatedTool *bool `json:"AutomatedTool,omitempty" name:"AutomatedTool"`
	// 是否开启页面防调试，true：是，false：否

	CrackBehavior *bool `json:"CrackBehavior,omitempty" name:"CrackBehavior"`
	// 搜索引擎白名单，可选的值为：["sogou", "baidu", "yandex", "360", "yahoo", "bing", "google"]

	GoodBot []*string `json:"GoodBot,omitempty" name:"GoodBot"`
	// 前端对抗规则ID

	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
	// 规则内部ID

	RasSiteID *string `json:"RasSiteID,omitempty" name:"RasSiteID"`
	// 前端对抗的开关，0（关闭） 1（开启）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 请求白名单

	ReqWhiteList []*WhiteListItem `json:"ReqWhiteList,omitempty" name:"ReqWhiteList"`
	// 响应白名单

	RspWhiteList []*WhiteListItem `json:"RspWhiteList,omitempty" name:"RspWhiteList"`
	// 前端对抗的防护URL，多个URL以,串接

	CheckURL *string `json:"CheckURL,omitempty" name:"CheckURL"`
	// 是否拦截扫描工具（目前暂不支持设置）

	PreventScanner *bool `json:"PreventScanner,omitempty" name:"PreventScanner"`
}

type DescribeBypassSpartaProtectionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBypassSpartaProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBypassSpartaProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafThreatenIntelligenceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWafThreatenIntelligenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafThreatenIntelligenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVIPRSRequest struct {
	*tchttp.BaseRequest

	// 记录

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
	// 目标权重

	DstWeight *int64 `json:"DstWeight,omitempty" name:"DstWeight"`
	// 源权重

	SrcWeight *int64 `json:"SrcWeight,omitempty" name:"SrcWeight"`
}

func (r *ModifyVIPRSRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVIPRSRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotToken struct {

	// 会话名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 会话描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 会话id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 策略的开关状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 会话位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 会话key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 会话匹配方式，前缀匹配、后缀匹配等

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 会话更新的时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type ModifyInstanceQpsLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceQpsLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceQpsLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttackDetail struct {

	// 攻击命中位置

	ArgsName *string `json:"ArgsName,omitempty" name:"ArgsName"`
	// 攻击命中的内容

	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`
	// 攻击者IP

	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`
	// 攻击时间

	AttackTime *string `json:"AttackTime,omitempty" name:"AttackTime"`
	// 攻击类型

	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
	// 攻击聚合次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 被攻击的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 原始请求字符串

	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`
	// 攻击者城市

	IpinfoCity *string `json:"IpinfoCity,omitempty" name:"IpinfoCity"`
	// 攻击者运营商

	IpinfoDetail *string `json:"IpinfoDetail,omitempty" name:"IpinfoDetail"`
	// 经度

	IpinfoDimensionality *float64 `json:"IpinfoDimensionality,omitempty" name:"IpinfoDimensionality"`
	// 纬度

	IpinfoLongitude *float64 `json:"IpinfoLongitude,omitempty" name:"IpinfoLongitude"`
	// 线路

	IpinfoIsp *string `json:"IpinfoIsp,omitempty" name:"IpinfoIsp"`
	// 国家

	IpinfoNation *string `json:"IpinfoNation,omitempty" name:"IpinfoNation"`
	// 省份

	IpinfoProvince *string `json:"IpinfoProvince,omitempty" name:"IpinfoProvince"`
	// 国家简称

	IpinfoState *string `json:"IpinfoState,omitempty" name:"IpinfoState"`
	// HTTP方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 风险等级

	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 拦截状态，0是观察，1是放行

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 攻击的Uri

	Uri *string `json:"Uri,omitempty" name:"Uri"`
	// 攻击者浏览器User Agent

	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
	// 请求ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 如果是自定义规则，则显示规则的名字，如果不是则为-

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type DescribeBotSceneActionRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名

		Domain *string `json:"Domain,omitempty" name:"Domain"`
		// 场景ID

		SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
		// 动作策略列表，不为空，新建场景时，会默认创建一个系统宽松策略

		ActionRuleList []*BotSceneActionRule `json:"ActionRuleList,omitempty" name:"ActionRuleList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotSceneActionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSceneActionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤数组

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePanDomainDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0:未开通非标支持泛域名；1:非标开通泛域名；（主要针对高级别和超轻版实例的非标开通泛域名）

		PanDomainFlag *int64 `json:"PanDomainFlag,omitempty" name:"PanDomainFlag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePanDomainDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePanDomainDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessLogInfo struct {

	// 日志时间，单位ms

	Time *int64 `json:"Time,omitempty" name:"Time"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 日志来源IP

	Source *string `json:"Source,omitempty" name:"Source"`
	// 日志文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 日志上报请求包的ID

	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
	// 请求包内日志的ID

	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`
	// 日志内容的Json序列化字符串
	// 注意：此字段可能返回 null，表示取不到有效值。

	LogJson *string `json:"LogJson,omitempty" name:"LogJson"`
}

type AddGlobalWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则序号

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则Id

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 规则状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 匹配规则项列表

	Rules []*GlobalWhiteCond `json:"Rules,omitempty" name:"Rules"`
}

func (r *AddGlobalWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddGlobalWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackTotalCountRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询的域名，全部域名不指定

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询条件，默认为""

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
}

func (r *GetAttackTotalCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackTotalCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PostAccessDownloadTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务task id

		Flow *string `json:"Flow,omitempty" name:"Flow"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PostAccessDownloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PostAccessDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCdcClbWafRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserCdcClbWafRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCdcClbWafRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleCosKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Cos临时secret-id

		SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
		// Cos临时secret-key

		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
		// Cos临时token

		SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
		// 上传Cos的bucket

		Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
		// 上传Cos的region

		Region *string `json:"Region,omitempty" name:"Region"`
		// ExpiredTime

		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// Expiration

		Expiration *string `json:"Expiration,omitempty" name:"Expiration"`
		// 将文件上传到Cos的Key，可在请求参数中指定，留空由后台随机生成

		UploadKey *string `json:"UploadKey,omitempty" name:"UploadKey"`
		// CosUrl

		CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`
		// Domain

		Domain *string `json:"Domain,omitempty" name:"Domain"`
		// Schema

		Schema *string `json:"Schema,omitempty" name:"Schema"`
		// StartTime

		StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRuleCosKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleCosKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreaBanAreasRsp struct {

	// 状态 "0"：未开启地域封禁 "1"：开启地域封禁

	Status *string `json:"Status,omitempty" name:"Status"`
	// 数据来源 custom-自定义(默认)、batch-批量防护

	Source *string `json:"Source,omitempty" name:"Source"`
	// 字符串数据，配置的地域列表

	Areas []*string `json:"Areas,omitempty" name:"Areas"`
}

type DescribeAttackDownloadRecordRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAttackDownloadRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackDownloadRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserEngineTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：menshen 1: Tiga

		Type *uint64 `json:"Type,omitempty" name:"Type"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserEngineTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserEngineTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePostCLSFlowRequest struct {
	*tchttp.BaseRequest

	// 投递的CLS所在日志集合名称，默认为 waf_post_logset

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 投递的CLS所在日志主题的名称，默认为 waf_post_logtopic

	LogtopicName *string `json:"LogtopicName,omitempty" name:"LogtopicName"`
	// 1-访问日志，2-攻击日志，默认为访问日志。

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 开启访问日志某些字段是否投递

	WriteConfig *FieldWriteConfig `json:"WriteConfig,omitempty" name:"WriteConfig"`
	// 投递的CLS所在区域，默认为ap-shanghai

	CLSRegion *string `json:"CLSRegion,omitempty" name:"CLSRegion"`
}

func (r *CreatePostCLSFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePostCLSFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiInfoLeakRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录条数

		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`
		// 规则列表

		RuleList []*DescribeAntiInfoLeakRulesRuleItem `json:"RuleList,omitempty" name:"RuleList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAntiInfoLeakRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAntiInfoLeakRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotItemsListItem struct {

	// 动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// bot特征

	BotFeature []*string `json:"BotFeature,omitempty" name:"BotFeature"`
	// mongodb id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 数目

	Nums *int64 `json:"Nums,omitempty" name:"Nums"`
	// 关联规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 得分

	Score *BotListScore `json:"Score,omitempty" name:"Score"`
	// 持续时间

	SessionDuration *float64 `json:"SessionDuration,omitempty" name:"SessionDuration"`
	// 源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 统计数据

	Stat *BotListStat `json:"Stat,omitempty" name:"Stat"`
	// 时间戳

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 公开类型附加参数

	TcbDetail *string `json:"TcbDetail,omitempty" name:"TcbDetail"`
}

type AddGlobalWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddGlobalWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddGlobalWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCdcResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCdcResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCdcResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainDetailsClbResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// clb域名详情

		DomainsClbPartInfo *ClbDomainsInfo `json:"DomainsClbPartInfo,omitempty" name:"DomainsClbPartInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainDetailsClbResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainDetailsClbResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePostCKafkaFlowRequest struct {
	*tchttp.BaseRequest

	// SASL密码

	SASLPassword *string `json:"SASLPassword,omitempty" name:"SASLPassword"`
	// 目前无作用

	Sync *int64 `json:"Sync,omitempty" name:"Sync"`
	// 开启访问日志某些字段是否投递

	WriteConfig *FieldWriteConfig `json:"WriteConfig,omitempty" name:"WriteConfig"`
	// 投递的CKafka所在区域

	CKafkaRegion *string `json:"CKafkaRegion,omitempty" name:"CKafkaRegion"`
	// 默认为none，支持snappy、gzip和lz4压缩，推荐snappy

	Compression *string `json:"Compression,omitempty" name:"Compression"`
	// 1-访问日志，2-攻击日志，默认为访问日志

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 主题名称，默认不传或者传空字符串，默认值为waf_post_access_log

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 是否开启SASL校验，默认不开启，0-关闭，1-开启

	SASLEnable *int64 `json:"SASLEnable,omitempty" name:"SASLEnable"`
	// SASL用户名

	SASLUser *string `json:"SASLUser,omitempty" name:"SASLUser"`
	// kafka集群的版本号

	KafkaVersion *string `json:"KafkaVersion,omitempty" name:"KafkaVersion"`
	// 客户的CKafka 实例ID

	CKafkaID *string `json:"CKafkaID,omitempty" name:"CKafkaID"`
	// 支撑环境是IP:PORT，外网环境是domain:PORT

	Brokers *string `json:"Brokers,omitempty" name:"Brokers"`
	// 1-外网TGW，2-支撑环境，默认为支撑环境

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
}

func (r *CreatePostCKafkaFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePostCKafkaFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotReferAnalyzeRes struct {

	// refer存在性

	ReferExist *bool `json:"ReferExist,omitempty" name:"ReferExist"`
	// refer重复比

	ReferRepeat *float64 `json:"ReferRepeat,omitempty" name:"ReferRepeat"`
	// 出现最多的refer

	ReferMax *string `json:"ReferMax,omitempty" name:"ReferMax"`
	// refer存在比

	ReferValidRate *float64 `json:"ReferValidRate,omitempty" name:"ReferValidRate"`
	// refer种类

	ReferKind *uint64 `json:"ReferKind,omitempty" name:"ReferKind"`
	// refer滥用

	ReferAbuse *bool `json:"ReferAbuse,omitempty" name:"ReferAbuse"`
	// refer得分

	ReferScore *uint64 `json:"ReferScore,omitempty" name:"ReferScore"`
}

type DescribeBotRuleOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自定义规则数量

		CustomRuleNums *int64 `json:"CustomRuleNums,omitempty" name:"CustomRuleNums"`
		// bot 开关模块列表

		BotModuleStatus *BotModuleStatusEntry `json:"BotModuleStatus,omitempty" name:"BotModuleStatus"`
		// 用户自定义会话数量

		TokenRuleNums *int64 `json:"TokenRuleNums,omitempty" name:"TokenRuleNums"`
		// 用户配置的合法爬虫数量

		GoodBotNums *int64 `json:"GoodBotNums,omitempty" name:"GoodBotNums"`
		// bot开关状态

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotRuleOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppIdDetail struct {

	// 套餐版本，跟saas保持一致

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 是否自动续费，1：自动续费，0：不自动续费

	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
	// 套餐的购买时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 购买的日志套餐

	Cls *ClsPackage `json:"Cls,omitempty" name:"Cls"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 套餐子域名限制个数

	DomainLimit *uint64 `json:"DomainLimit,omitempty" name:"DomainLimit"`
	// 套餐子域名已经使用的个数

	DomainCount *uint64 `json:"DomainCount,omitempty" name:"DomainCount"`
	// 套餐主域名已经使用的个数

	MainDomainCount *uint64 `json:"MainDomainCount,omitempty" name:"MainDomainCount"`
	// 套餐主域名限制个数

	MainDomainLimit *uint64 `json:"MainDomainLimit,omitempty" name:"MainDomainLimit"`
	// 套餐的地域限制个数

	RegionLimit *uint64 `json:"RegionLimit,omitempty" name:"RegionLimit"`
	// 套餐的监听器限制个数

	LbLimit *uint64 `json:"LbLimit,omitempty" name:"LbLimit"`
	// 购买的QPS套餐

	QPS *QPSPackage `json:"QPS,omitempty" name:"QPS"`
	// 当前的QPS峰值

	MaxQPS *uint64 `json:"MaxQPS,omitempty" name:"MaxQPS"`
	// 资源ID

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 暂时未用到

	Type *string `json:"Type,omitempty" name:"Type"`
	// 购买的域名套餐

	DomainPkg *DomainPackage `json:"DomainPkg,omitempty" name:"DomainPkg"`
	// 账号开通了clbwaf的地域白名单，如果为空则表示开通了全部地域，否则则开通的地域以,串接。

	AllowRegions *string `json:"AllowRegions,omitempty" name:"AllowRegions"`
	// 账号开通了clbwaf的清洗模式的地域白名单，如果为空则表示没有地域开启了清洗模式，否则则开通的地域以,串接。

	CCGuardRegions *string `json:"CCGuardRegions,omitempty" name:"CCGuardRegions"`
}

type DescribeDNSDetectDataChartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 每一个时间点的被劫持用户个数

		Points []*uint64 `json:"Points,omitempty" name:"Points"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDNSDetectDataChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSDetectDataChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客户的产品信息，格式json string

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAppletInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关套餐ID

		PackageID *uint64 `json:"PackageID,omitempty" name:"PackageID"`
		// 是否独享网关套餐
		// 0  否
		// 1  是

		PackageExclusive *uint64 `json:"PackageExclusive,omitempty" name:"PackageExclusive"`
		// dount 网关ID

		GateWayID *string `json:"GateWayID,omitempty" name:"GateWayID"`
		// 备注信息

		Comment *string `json:"Comment,omitempty" name:"Comment"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserAppletInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAppletInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest

	// 新名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAccessLogRequest struct {
	*tchttp.BaseRequest

	// 第几页，从0开始。新版本接口字段

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题，新版本此字段填空字符串

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 查询语句，语句长度最大为4096

	Query *string `json:"Query,omitempty" name:"Query"`
	// 单次查询返回的日志条数，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 新版本此字段失效，填空字符串，翻页使用Page

	Context *string `json:"Context,omitempty" name:"Context"`
	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *SearchAccessLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAccessLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotDetailRequest struct {
	*tchttp.BaseRequest

	// 源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 趋势图开始时间

	StartTs *uint64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间

	EndTs *uint64 `json:"EndTs,omitempty" name:"EndTs"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotDetailDownloadTaskRequest struct {
	*tchttp.BaseRequest

	// 排序条件

	Sorter *DataSorter `json:"Sorter,omitempty" name:"Sorter"`
	// Limit跳跃: 从1开始

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset跳跃

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBotDetailDownloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotDetailDownloadTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotGlobalUAStrategyRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBotGlobalUAStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotGlobalUAStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BanAreaItem struct {

	// 地址的其他英文缩写

	EnAbbr *string `json:"EnAbbr,omitempty" name:"EnAbbr"`
	// 地址的英文缩写

	EN *string `json:"EN,omitempty" name:"EN"`
	// 地址的中文名

	ZH *string `json:"ZH,omitempty" name:"ZH"`
}

type CreateCopyCustomRuleRequest struct {
	*tchttp.BaseRequest

	// 目标域名字符串；

	Target *string `json:"Target,omitempty" name:"Target"`
	// 复制类型

	CopyType *int64 `json:"CopyType,omitempty" name:"CopyType"`
	// 来源域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则字符串；

	Rules *string `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateCopyCustomRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCopyCustomRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPackageRenewRequest struct {
	*tchttp.BaseRequest

	// 1：开启自动续费，0：关闭自动续费

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *ModifyPackageRenewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPackageRenewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApiRulesRequest struct {
	*tchttp.BaseRequest

	// 用户删除api规则的列表

	ApiList []*DeleteApi `json:"ApiList,omitempty" name:"ApiList"`
	// 用户需要删除规则的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DeleteApiRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiRulesRequest) FromJsonString(s string) error {
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

type CopyCustomRuleRequest struct {
	*tchttp.BaseRequest

	// 自定义策略的来源域名

	From *string `json:"From,omitempty" name:"From"`
	// 自定义策略的名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 客户使用的waf版本, 可选值为sparta-waf, clb-waf, cdn-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 自定义策略需要复制到哪些域名,使用逗号, 分隔多个域名,其他全部域名使用all

	Domains *string `json:"Domains,omitempty" name:"Domains"`
}

func (r *CopyCustomRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyCustomRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PostAttackDownloadTaskRequest struct {
	*tchttp.BaseRequest

	// 查询的域名，所有域名使用all

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// Lucene语法

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 默认为desc，可以取值desc和asc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 下载的日志条数

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

func (r *PostAttackDownloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PostAttackDownloadTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpartaProtectionRequest struct {
	*tchttp.BaseRequest

	// 表示是否开启了CDN代理

	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// 负载均衡策略，0表示轮徇，1表示IP hash

	LoadBalance *int64 `json:"LoadBalance,omitempty" name:"LoadBalance"`
	// 是否灰度

	IsGray *int64 `json:"IsGray,omitempty" name:"IsGray"`
	// 300s

	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitempty" name:"ProxySendTimeout"`
	// 是否开启HTTP2，1表示开启，0表示不开启http2。开启HTTP2需要HTTPS支持

	IsHttp2 *int64 `json:"IsHttp2,omitempty" name:"IsHttp2"`
	// IsCdn=3时，需要填此参数，表示自定义header

	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`
	// TLS版本信息

	TLSVersion *int64 `json:"TLSVersion,omitempty" name:"TLSVersion"`
	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；

	SniType *int64 `json:"SniType,omitempty" name:"SniType"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// HTTPS回源协议

	UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`
	// 回源类型，0表示通过IP回源,1 表示通过域名回源

	UpstreamType *int64 `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// UpstreamType=1时，填次字段表示回源域名

	UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`
	// WAF版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 端口信息

	Ports []*SpartaProtectionPort `json:"Ports,omitempty" name:"Ports"`
	// src的权重

	Weights []*int64 `json:"Weights,omitempty" name:"Weights"`
	// 表示是否强制跳转到HTTPS，1表示开启，0表示不开启

	HttpsRewrite *uint64 `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`
	// 加密套件信息

	Ciphers []*int64 `json:"Ciphers,omitempty" name:"Ciphers"`
	// 300s

	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitempty" name:"ProxyReadTimeout"`
	// 0:关闭xff重置；1:开启xff重置，只有在IsCdn=0时可以开启

	XFFReset *int64 `json:"XFFReset,omitempty" name:"XFFReset"`
	// UpstreamType=0时，填次字段表示回源ip

	SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`
	// 长短连接标志，仅IP回源时有效

	IsKeepAlive *string `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`
	// 0:不支持选择：默认模板  1:通用型模板 2:安全型模板 3:自定义模板

	CipherTemplate *int64 `json:"CipherTemplate,omitempty" name:"CipherTemplate"`
	// SniType=3时，需要填此参数，表示自定义的host；

	SniHost *string `json:"SniHost,omitempty" name:"SniHost"`
	// CertType=1时，需要填次参数，表示证书内容

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// CertType=1时，需要填次参数，表示证书的私钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// CertType=2时，需要填次参数，表示证书的ID

	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`
	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段

	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 是否为Anycast ip类型：1 Anycast 0 普通ip

	Anycast *int64 `json:"Anycast,omitempty" name:"Anycast"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 证书类型，0表示没有证书，CertType=1表示自有证书,2 为托管证书

	CertType *int64 `json:"CertType,omitempty" name:"CertType"`
	// 是否开启WebSocket， 1：开启WebSocket，0：不开启WebSocket

	IsWebsocket *int64 `json:"IsWebsocket,omitempty" name:"IsWebsocket"`
	// 是否开启源站的主动健康检测，1表示开启，0表示不开启

	ActiveCheck *int64 `json:"ActiveCheck,omitempty" name:"ActiveCheck"`
}

func (r *ModifySpartaProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpartaProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainPackageNew struct {

	// 资源ID

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 是否自动续费，1：自动续费，0：不自动续费

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 套餐购买个数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 套餐购买地域，clb-waf暂时没有用到

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeInstanceDomainListRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDomainListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDomainListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRateLimitsRequest struct {
	*tchttp.BaseRequest

	// 规则ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 规则名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 限流接口名

	Method *string `json:"Method,omitempty" name:"Method"`
	// 规则是否启用，0表示不启用，1表示启用

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 分页的起始位置，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页行数，默认为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 排序方式，支持desc,asc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序对象，支持priority,timestamp

	By *string `json:"By,omitempty" name:"By"`
	// 过滤

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRateLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRateLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrategyForAntiInfoLeak struct {

	// 匹配条件，returncode（响应码）、keywords（关键字）、information（敏感信息）

	Field *string `json:"Field,omitempty" name:"Field"`
	// 逻辑符号，固定取值为contains

	CompareFunc *string `json:"CompareFunc,omitempty" name:"CompareFunc"`
	// 匹配内容。
	// 以下三个对应Field为information时可取的匹配内容：
	// idcard（身份证）、phone（手机号）、bankcard（银行卡）。
	// 以下为对应Field为returncode时可取的匹配内容：
	// 400（状态码400）、403（状态码403）、404（状态码404）、4xx（其它4xx状态码）、500（状态码500）、501（状态码501）、502（状态码502）、504（状态码504）、5xx（其它5xx状态码）。
	// 当对应Field为keywords时由用户自己输入匹配内容。
	//

	Content *string `json:"Content,omitempty" name:"Content"`
}

type DeleteHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessFastAnalysisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 注意：此字段可能返回 null，表示取不到有效值

		FieldValueRatioInfos []*AccessFieldValueRatioInfo `json:"FieldValueRatioInfos,omitempty" name:"FieldValueRatioInfos"`
		// 日志条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessFastAnalysisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessFastAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPTypeRequest struct {
	*tchttp.BaseRequest

	// IP列表，支持单IP和CIDR

	Items []*string `json:"Items,omitempty" name:"Items"`
}

func (r *DescribeIPTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDNSDetectDataChartRequest struct {
	*tchttp.BaseRequest

	// 域名，如果是全部域名则为“all”

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询的开始日期，格式为：年-月-日

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 查询的结束日期，格式为：年-月-日

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 查询的粒度，单位为分。如果粒度是天则为1440，如果粒度是小时则为60。

	Stride *uint64 `json:"Stride,omitempty" name:"Stride"`
}

func (r *DescribeDNSDetectDataChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSDetectDataChartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainBlockPageRequest struct {
	*tchttp.BaseRequest

	// 查询的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainBlockPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainBlockPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppletAuthInfoV3 struct {

	// 微信小程序ID

	AppletID *string `json:"AppletID,omitempty" name:"AppletID"`
	// 授权状态
	// 0 未授权
	// 1 已授权

	AuthStatus *uint64 `json:"AuthStatus,omitempty" name:"AuthStatus"`
}

type AttackIpInfo struct {

	// 攻击者IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 攻击次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 攻击者所在城市json形式字符串

	City *string `json:"City,omitempty" name:"City"`
}

type DescribeBotRecordDetailRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 记录Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeBotRecordDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotRecordDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseCanBuyTceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLicenseCanBuyTceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseCanBuyTceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRateLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 限流操作结果的基本信息

		BaseInfo *RateLimitCommonRsp `json:"BaseInfo,omitempty" name:"BaseInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRateLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRateLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotRecordItemsData struct {

	// 数据

	Res *BotRecordItemRes `json:"Res,omitempty" name:"Res"`
}

type OptionalGroup struct {

	// 集群ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 集群描述

	Name *string `json:"Name,omitempty" name:"Name"`
}

type PortItem struct {

	// 监听端口配置

	Port *string `json:"Port,omitempty" name:"Port"`
	// 与Port一一对应，表示端口对应的协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 与Port一一对应,  表示回源端口

	UpstreamPort *string `json:"UpstreamPort,omitempty" name:"UpstreamPort"`
	// 与Port一一对应,  表示回源协议

	UpstreamProtocol *string `json:"UpstreamProtocol,omitempty" name:"UpstreamProtocol"`
	// Nginx的服务器ID

	NginxServerId *string `json:"NginxServerId,omitempty" name:"NginxServerId"`
}

type DeleteWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单对应的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 白名单ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 白名单内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 白名单的备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 白名单的开关状态，0：关闭，1：开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 白名单的类型，jsinject_req：动态风控的请求白名单，jsinject_rsp：动态风控的响应白名单

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DeleteWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则库总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 规则库详细信息

		Objects []*TIGARuleStatusItem `json:"Objects,omitempty" name:"Objects"`
		// 正在使用中的规则库

		InUseRule *TIGARuleStatusItem `json:"InUseRule,omitempty" name:"InUseRule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessKeyValueInfo struct {

	// 需要配置键值或者元字段索引的字段

	Key *string `json:"Key,omitempty" name:"Key"`
	// 字段的索引描述信息

	Value *AccessValueInfo `json:"Value,omitempty" name:"Value"`
}

type PageForDescribeSpartaProtectionList struct {

	// 每页个数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 页码

	Index *uint64 `json:"Index,omitempty" name:"Index"`
}

type QpsData struct {

	// 弹性qps默认值

	ElasticBillingDefault *uint64 `json:"ElasticBillingDefault,omitempty" name:"ElasticBillingDefault"`
	// 弹性qps最小值

	ElasticBillingMin *uint64 `json:"ElasticBillingMin,omitempty" name:"ElasticBillingMin"`
	// 弹性qps最大值

	ElasticBillingMax *uint64 `json:"ElasticBillingMax,omitempty" name:"ElasticBillingMax"`
	// 业务扩展包最大qps

	QPSExtendMax *uint64 `json:"QPSExtendMax,omitempty" name:"QPSExtendMax"`
	// 海外业务扩展包最大qps

	QPSExtendIntlMax *uint64 `json:"QPSExtendIntlMax,omitempty" name:"QPSExtendIntlMax"`
}

type AddBypassAllRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddBypassAllRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBypassAllRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreaBanSupportAreasRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAreaBanSupportAreasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreaBanSupportAreasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求总量

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttackCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppletAccessV3Request struct {
	*tchttp.BaseRequest

	// Stage =2 可修改小程序接入名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 1   授权接入阶段
	// 2   客户体验阶段
	// 3   客户发布阶段

	Stage *uint64 `json:"Stage,omitempty" name:"Stage"`
	// 小程序ID

	AppletID *string `json:"AppletID,omitempty" name:"AppletID"`
	// Stage =2 可修改，详情

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// Stage =2 可修改，客户体验阶段的状态
	// 0  初始状态
	// 100  完成
	//

	ClientTrialStatus *uint64 `json:"ClientTrialStatus,omitempty" name:"ClientTrialStatus"`
	// Stage =3 可修改，客户发布阶段的状态
	// 0  初始状态
	// 100  完成

	ClientPublishStatus *uint64 `json:"ClientPublishStatus,omitempty" name:"ClientPublishStatus"`
	// Stage =3 可修改， 小程序接入的生效版本
	// trial    体验版
	// any   开发版+体验版+发布版

	EffectiveVersion *string `json:"EffectiveVersion,omitempty" name:"EffectiveVersion"`
	// Stage =2 可修改， 全量传入所有的域名，协议，端口

	Domains []*AppletAccessDomainV3 `json:"Domains,omitempty" name:"Domains"`
}

func (r *ModifyAppletAccessV3Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppletAccessV3Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotTCBRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正常情况为null

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotTCBRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotTCBRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttackDownloadRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttackDownloadRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttackDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// instance总数

		TotalInstanceCount *uint64 `json:"TotalInstanceCount,omitempty" name:"TotalInstanceCount"`
		// QPS总量

		TotalQPS *uint64 `json:"TotalQPS,omitempty" name:"TotalQPS"`
		// CLS使用量

		UsedCLSStorage *uint64 `json:"UsedCLSStorage,omitempty" name:"UsedCLSStorage"`
		// cls的资源id

		CLSResourceIds *string `json:"CLSResourceIds,omitempty" name:"CLSResourceIds"`
		// 开启bot的域名数量

		BotCount *uint64 `json:"BotCount,omitempty" name:"BotCount"`
		// QPS包

		QPSPackageCount *uint64 `json:"QPSPackageCount,omitempty" name:"QPSPackageCount"`
		// CLS总量

		TotalCLSStorage *uint64 `json:"TotalCLSStorage,omitempty" name:"TotalCLSStorage"`
		// 开启api安全的域名数量

		ApiSecCount *uint64 `json:"ApiSecCount,omitempty" name:"ApiSecCount"`
		// 重保护网AppId域名数

		HWDomainCount *int64 `json:"HWDomainCount,omitempty" name:"HWDomainCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateDealsAndPayRequest struct {
	*tchttp.BaseRequest

	// 计费下单入参

	Goods []*Goods `json:"Goods,omitempty" name:"Goods"`
}

func (r *GenerateDealsAndPayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateDealsAndPayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotScoreRuleEntry struct {

	// 分数区间上限

	Upper *string `json:"Upper,omitempty" name:"Upper"`
	// 分数区间下限

	Lower *string `json:"Lower,omitempty" name:"Lower"`
	// 处置动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 流量标签

	Label *string `json:"Label,omitempty" name:"Label"`
	// 重定向

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
}

type DeleteAttackWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除失败的规则序号组

		FailIds []*uint64 `json:"FailIds,omitempty" name:"FailIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttackWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttackWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseCanBuyTceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SaaS是否可购买 0: 否 1:是

		SpartaBuy *int64 `json:"SpartaBuy,omitempty" name:"SpartaBuy"`
		// Clb是否可购买 0: 否 1:是

		CIbBuy *int64 `json:"CIbBuy,omitempty" name:"CIbBuy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseCanBuyTceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseCanBuyTceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInstanceCountRequest struct {
	*tchttp.BaseRequest

	// 版本：clb-waf，sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribeUserInstanceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInstanceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReloadConfigDeliveryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReloadConfigDeliveryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReloadConfigDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAIVerifyResultRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 内容

	Payload *string `json:"Payload,omitempty" name:"Payload"`
}

func (r *DescribeAIVerifyResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAIVerifyResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAntiFakeUrlStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAntiFakeUrlStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAntiFakeUrlStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWafAutoDenyStatusRequest struct {
	*tchttp.BaseRequest

	// WAF 自动封禁配置项

	WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitempty" name:"WafAutoDenyDetails"`
}

func (r *ModifyWafAutoDenyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWafAutoDenyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDownloadRecordRequest struct {
	*tchttp.BaseRequest

	// 记录id

	Flow *string `json:"Flow,omitempty" name:"Flow"`
}

func (r *DeleteDownloadRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDownloadRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotStatusRequest struct {
	*tchttp.BaseRequest

	// 类别-bot

	Category *string `json:"Category,omitempty" name:"Category"`
}

func (r *DescribeBotStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserClbWafRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域（标准的ap-格式）列表

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 包含详细属性的地域信息

		RichDatas []*ClbWafRegionItem `json:"RichDatas,omitempty" name:"RichDatas"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserClbWafRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserClbWafRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBatchAreaBanAppIdDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据信息

		Data *BatchAreaBanAppIdDomainData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBatchAreaBanAppIdDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBatchAreaBanAppIdDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBusinessRiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除成功总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBusinessRiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBusinessRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpHitItemsRequest struct {
	*tchttp.BaseRequest

	// 创建时间最小时间戳

	CtsMin *uint64 `json:"CtsMin,omitempty" name:"CtsMin"`
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 有效时间最大时间戳

	VtsMax *uint64 `json:"VtsMax,omitempty" name:"VtsMax"`
	// 偏移参数

	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`
	// 排序参数

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 计数标识

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 有效时间最小时间戳

	ValidTimeStampMin *uint64 `json:"ValidTimeStampMin,omitempty" name:"ValidTimeStampMin"`
	// 有效时间最大时间戳

	ValidTimeStampMax *uint64 `json:"ValidTimeStampMax,omitempty" name:"ValidTimeStampMax"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 有效时间最小时间戳

	VtsMin *uint64 `json:"VtsMin,omitempty" name:"VtsMin"`
	// 创建时间最大时间戳

	CtsMax *uint64 `json:"CtsMax,omitempty" name:"CtsMax"`
}

func (r *DescribeIpHitItemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpHitItemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableUserSignaturePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名规则策略id

		UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableUserSignaturePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableUserSignaturePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableClientMsgRequest struct {
	*tchttp.BaseRequest

	// 0表示关闭1表示开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *EnableClientMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableClientMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QPSPackageNew struct {

	// 资源ID

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 是否自动续费，1：自动续费，0：不自动续费

	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 套餐购买个数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 套餐购买地域，clb-waf暂时没有用到

	Region *string `json:"Region,omitempty" name:"Region"`
	// 计费项

	BillingItem *string `json:"BillingItem,omitempty" name:"BillingItem"`
}

type DescribeModuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CC防护是否开启

		CcProtection *uint64 `json:"CcProtection,omitempty" name:"CcProtection"`
		// 网页防篡改是否开启

		AntiTamper *uint64 `json:"AntiTamper,omitempty" name:"AntiTamper"`
		// 信息防泄漏是否开启

		AntiLeakage *uint64 `json:"AntiLeakage,omitempty" name:"AntiLeakage"`
		// API安全是否开启

		ApiProtection *uint64 `json:"ApiProtection,omitempty" name:"ApiProtection"`
		// 限流模块开关

		RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`
		// WEB安全规则是否开启

		WebSecurity *uint64 `json:"WebSecurity,omitempty" name:"WebSecurity"`
		// 访问控制规则是否开启

		AccessControl *int64 `json:"AccessControl,omitempty" name:"AccessControl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Data *SessionData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpartUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名总数限制

		DomainCount *string `json:"DomainCount,omitempty" name:"DomainCount"`
		// 区域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 自动续费开关 “true” ： 开启自动续费

		AutoRenew *string `json:"AutoRenew,omitempty" name:"AutoRenew"`
		// 一级域名总数限制

		MainDomainCount *string `json:"MainDomainCount,omitempty" name:"MainDomainCount"`
		// qps套餐包详情

		Qps *UserInfoQPS `json:"Qps,omitempty" name:"Qps"`
		// 资源ID

		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
		// 套餐版本信息

		Type *string `json:"Type,omitempty" name:"Type"`
		// 最大Qps

		MaxQps *string `json:"MaxQps,omitempty" name:"MaxQps"`
		// 已使用域名数量

		DomainUsed *string `json:"DomainUsed,omitempty" name:"DomainUsed"`
		// 一级域名使用数量

		MainDomainUsed *string `json:"MainDomainUsed,omitempty" name:"MainDomainUsed"`
		// cls套餐包详情

		Cls *UserInfoCls `json:"Cls,omitempty" name:"Cls"`
		// ID，当未开通waf时，id为0

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 套餐有效期

		Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
		// 开始时间

		BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
		// 是否是sparta

		IsSparta *string `json:"IsSparta,omitempty" name:"IsSparta"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpartUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpartUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackDetailRequest struct {
	*tchttp.BaseRequest

	// 查询结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 查询的数量，默认10条，最多100条

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 攻击者IP

	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`
	// 攻击类型

	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
	// 查询的域名，所有域名使用all

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 查询起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 拦截状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 自定义规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 查询的游标。第一次请求使用空字符串即可，后续请求使用上一次请求返回的最后一条记录的context的值即可。

	Context *string `json:"Context,omitempty" name:"Context"`
	// 风险等级

	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
}

func (r *DescribeAttackDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafCreateResourceAfterPayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *uint64 `json:"Code,omitempty" name:"Code"`
		// 状态码解释

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafCreateResourceAfterPayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafCreateResourceAfterPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyBotTCBRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 目标域名

	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *CopyBotTCBRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBotTCBRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockPagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户已经配置的自定义拦截页面数组

		Pages []*BlockPage `json:"Pages,omitempty" name:"Pages"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockPagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockPagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostFlowModeRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// WAF流量模式，1：清洗模式，0：镜像模式（默认）

	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *ModifyHostFlowModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostFlowModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainCustomRuleInfoData struct {

	// 自定义规则

	CustomRule *DomainCustomRuleInfoItem `json:"CustomRule,omitempty" name:"CustomRule"`
	// 批量域名规则

	BatchDomainRule *DomainCustomRuleInfoItem `json:"BatchDomainRule,omitempty" name:"BatchDomainRule"`
}

type DescribeDomainRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表详情

		Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleUpgradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 启动升级的规则库id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleUpgradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMiniProgramPeakPointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据点

		Points []*MiniProgramItem `json:"Points,omitempty" name:"Points"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMiniProgramPeakPointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMiniProgramPeakPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStraceCountRequest struct {
	*tchttp.BaseRequest

	// 防护类型，不指定则默认为business_risk

	StraceType *string `json:"StraceType,omitempty" name:"StraceType"`
	// 起始地址

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束地址

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 域名，不指定则为全部域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *GetStraceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStraceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionedIpData struct {

	// ip数据

	Res []*ActionedIpItem `json:"Res,omitempty" name:"Res"`
	// 计数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type DeleteCachePathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCachePathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCachePathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiFakeUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *string `json:"Total,omitempty" name:"Total"`
		// 信息

		List []*CacheUrlItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAntiFakeUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAntiFakeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserWhiteRule struct {

	// 白名单的id

	WhiteRuleId *uint64 `json:"WhiteRuleId,omitempty" name:"WhiteRuleId"`
	// 规则id

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 匹配域

	MatchField *string `json:"MatchField,omitempty" name:"MatchField"`
	// 匹配方法

	MatchMethod *string `json:"MatchMethod,omitempty" name:"MatchMethod"`
	// 匹配内容

	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type DeleteWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CdcCluster struct {

	// cdc的集群id

	Id *string `json:"Id,omitempty" name:"Id"`
	// cdc的集群名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type OpUserWhiteRule struct {

	// appid

	OpAppId *string `json:"OpAppId,omitempty" name:"OpAppId"`
	// 域名

	OpDomain *string `json:"OpDomain,omitempty" name:"OpDomain"`
	// id

	WhiteRuleId *uint64 `json:"WhiteRuleId,omitempty" name:"WhiteRuleId"`
	// 用户规则策略id

	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
	// 规则id

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 匹配域

	MatchField *string `json:"MatchField,omitempty" name:"MatchField"`
	// 匹配方法

	MatchMethod *string `json:"MatchMethod,omitempty" name:"MatchMethod"`
	// 匹配内容

	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type ModifyBotSceneUCBRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正常情况下为null

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBotSceneUCBRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotSceneUCBRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessLogItems struct {

	// 分析结果返回的KV数据对

	Data []*AccessLogItem `json:"Data,omitempty" name:"Data"`
}

type ResponseCode struct {

	// 如果成功则返回Success，失败则返回云api定义的错误码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 如果成功则返回Success，失败则返回WAF定义的二级错误码

	Message *string `json:"Message,omitempty" name:"Message"`
}

type Type struct {

	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type DeletePocTestTaskRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeletePocTestTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePocTestTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClsUsedResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// appid使用的cls量

		Used *float64 `json:"Used,omitempty" name:"Used"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClsUsedResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClsUsedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBotDetailDownloadTaskRequest struct {
	*tchttp.BaseRequest

	// 下载任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 导出日志字段

	Field []*string `json:"Field,omitempty" name:"Field"`
	// 查询条件

	Filter []*BotDataFilter `json:"Filter,omitempty" name:"Filter"`
	// 排序字段

	Sorter []*DataSorter `json:"Sorter,omitempty" name:"Sorter"`
	// 需导出的详情信息

	DetailModule []*string `json:"DetailModule,omitempty" name:"DetailModule"`
	// 删除标记

	DelFlag *bool `json:"DelFlag,omitempty" name:"DelFlag"`
	// 需删除的下载任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 指定下载域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 起始时间

	StartTs *int64 `json:"StartTs,omitempty" name:"StartTs"`
	// 结束时间

	EndTs *int64 `json:"EndTs,omitempty" name:"EndTs"`
}

func (r *ModifyBotDetailDownloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBotDetailDownloadTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiSecEventChange struct {

	// 变更人

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 变更的状态

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 时间戳

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type OpUserWhiteRuleNew struct {

	// appid

	OpAppId *string `json:"OpAppId,omitempty" name:"OpAppId"`
	// 域名

	OpDomain *string `json:"OpDomain,omitempty" name:"OpDomain"`
	// id

	WhiteRuleId *uint64 `json:"WhiteRuleId,omitempty" name:"WhiteRuleId"`
	// 规则id

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 匹配域

	MatchField *string `json:"MatchField,omitempty" name:"MatchField"`
	// 匹配方法

	MatchMethod *string `json:"MatchMethod,omitempty" name:"MatchMethod"`
	// 匹配内容

	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type DescribeClsUsedRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClsUsedRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClsUsedRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HavePostFlowServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *HavePostFlowServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HavePostFlowServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModuleTypeItem struct {

	// 模块名称

	Value *string `json:"Value,omitempty" name:"Value"`
	// 模块中文名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type AddAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则Id

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 规则状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 匹配规则项列表

	Rules []*UserWhiteRuleItem `json:"Rules,omitempty" name:"Rules"`
	// 规则序号

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *AddAttackWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAttackWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatedDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可替换证书的域名列表, Status:binded 已经绑定该证书的域名列表

		DomainList []*CertDomain `json:"DomainList,omitempty" name:"DomainList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificatedDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCertificatedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleTypesEnResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则类型描述列表

		Types []*TypeEn `json:"Types,omitempty" name:"Types"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleTypesEnResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleTypesEnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotSceneUAStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// BOT场景的开关状态

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotSceneUAStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotSceneUAStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 开关的状态，1是开启、0是关闭

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ModifyCustomRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PostAccessDownloadTaskRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间 形如：2021-01-15T00:00:00+08:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间 形如：2021-01-15T23:59:59+08:00

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// Lucene语法

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 默认为desc，可以取值desc和asc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *PostAccessDownloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PostAccessDownloadTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterInstallPackageRequest struct {
	*tchttp.BaseRequest

	// waf实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateClusterInstallPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterInstallPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBatchAreaBanRequest struct {
	*tchttp.BaseRequest

	// 规则Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteBatchAreaBanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchAreaBanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBotTCBRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 输出结果

		Data *BotListData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBotTCBRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBotTCBRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSpartaProtectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSpartaProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiInfoLeakageRulesRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 翻页支持，读取偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 翻页支持，读取长度限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序方式，asc或者desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序的根据

	By *string `json:"By,omitempty" name:"By"`
	// 过滤器,可以允许如下的值：
	// RuleId,Match_field,Name,Action,Status

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAntiInfoLeakageRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAntiInfoLeakageRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeParametersListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeParametersListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeParametersListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePocTestTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		Success *ResponseCode `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePocTestTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePocTestTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePeakValueRequest struct {
	*tchttp.BaseRequest

	// 五个值可选：
	// access-峰值qps
	// down-下行峰值带宽
	// up-上行峰值带宽
	// attack-Web攻击总数
	// cc-CC攻击总数趋势图

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 查询起始时间

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 查询结束时间

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 需要查询的域名，当前用户所有域名可以不传

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// WAF实例ID，不传则不过滤

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribePeakValueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePeakValueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCCRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// clb-waf或者sparta-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteCCRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCCRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DonwloadRecordItem struct {

	// 下载任务名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务ID

	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 记录条数

	Count *string `json:"Count,omitempty" name:"Count"`
	// 下载状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 下载文件URL

	Url *string `json:"Url,omitempty" name:"Url"`
	// 记录ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 产品ID

	Appid *string `json:"Appid,omitempty" name:"Appid"`
}

type DescribeAccessCursorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 游标

		Cursor *string `json:"Cursor,omitempty" name:"Cursor"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessCursorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessCursorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCachePathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置列表

		List []*DescribeCachePathRspListItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCachePathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCachePathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTsRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 区域列表

		Data []*RegionItem `json:"Data,omitempty" name:"Data"`
		// 区域列表总数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTsRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTsRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddApiRule struct {

	// api名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// api参数位置

	Position *string `json:"Position,omitempty" name:"Position"`
	// api类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 参数是否是必须

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 该参数的描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DeleteBotSceneUCBRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 自定义规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 1.BOT全局白名单处调用时，传"global";2.BOT场景配置时，传具体的场景ID

	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DeleteBotSceneUCBRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBotSceneUCBRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessSpartaProtectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败原因

		FailedMsg *string `json:"FailedMsg,omitempty" name:"FailedMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AccessSpartaProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBypassAllStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该用户是否被加入了全局的bypass列表

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBypassAllStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBypassAllStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PostPointInfo struct {

	// 该点的时间戳，单位为毫秒

	TimePoint *int64 `json:"TimePoint,omitempty" name:"TimePoint"`
	// 日志条数

	LogCount *int64 `json:"LogCount,omitempty" name:"LogCount"`
	// 未压缩的日志流量大小，单位为字节

	RealTraffic *int64 `json:"RealTraffic,omitempty" name:"RealTraffic"`
	// 压缩后的大小，单位为字节

	CompressionTraffic *int64 `json:"CompressionTraffic,omitempty" name:"CompressionTraffic"`
}

type CreateAppletAccessV3Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppletAccessV3Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppletAccessV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpiredInstanceInfo struct {

	// id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 自动续费

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 域名套餐

	DomainPkg *ExpiredDomainPackage `json:"DomainPkg,omitempty" name:"DomainPkg"`
	// 业务安全包

	FraudPkg *ExpiredFraudPkg `json:"FraudPkg,omitempty" name:"FraudPkg"`
	// Bot资源包

	BotPkg *ExpiredBotPkg `json:"BotPkg,omitempty" name:"BotPkg"`
	// bot的qps详情

	BotQPS *ExpiredBotQPS `json:"BotQPS,omitempty" name:"BotQPS"`
	// qps套餐

	QPS *ExpiredQPSPackage `json:"QPS,omitempty" name:"QPS"`
}

type DescribeCLSPackageRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCLSPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCLSPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
