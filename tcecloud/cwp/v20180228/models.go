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

package v20180228

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeAgentInstallCommandRequest struct {
	*tchttp.BaseRequest

	// 命令有效期，非腾讯云时必填

	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
	// 是否腾讯云

	IsCloud *bool `json:"IsCloud,omitempty" name:"IsCloud"`
	// 网络类型：basic-基础网络，private-VPC, public-公网，direct-专线

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 地域标示, NetType=direct时必填

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 标签ID列表，IsCloud=false时才会生效

	TagIds []*uint64 `json:"TagIds,omitempty" name:"TagIds"`
	// 代理方式接入的vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// VpcId, NetType=direct时必填

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeAgentInstallCommandRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentInstallCommandRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportEventlogRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>DateRange - String - 是否必填：否 - 日志时间范围 </li>
	// <li>SeverityClass - String  是否必填：否 - 威胁等级 <br/>"RISK":严重<br/>"HIGH":高危<br/>"MEDIUM":中危<br/>"LOW":低危<br/>"NOTICE":提示</li>
	// <li>EventType - String  是否必填：否 - 事件类型<br/>"MALWARE":木马文件<br/>"HOST_LOGIN":登录审计<br/>"BRUTE_ATTACK":暴力破解<br/>"MALICIOUS_REQUEST":恶意请求<br/>"ATTACK_LOG":网络攻击<br/>"HISH_RISK_BASH":高危命令<br/>"REVERSE_SHELL":反弹shell<br/>"PRIVILEGE_ESCALATION":本地提权<br/>"WEB_VUL":WEB漏洞<br/>"SYSTEM_COMPONENT_VUL":系统组件漏洞<br/>"EMERGENCY_VUL":应急漏洞<br/>"BASELINE":基线漏洞</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportEventlogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportEventlogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportScanTaskDetailsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数：ipOrAlias（服务器名/ip）

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 模块类型，当前提供：Malware 木马 , Vul 漏洞 , Baseline 基线

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 本次检测的任务id（不同于出参的导出本次检测Excel的任务Id）

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExportScanTaskDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportScanTaskDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttackLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttackLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportMalwaresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专业周报木马数据。

		WeeklyReportMalwares []*WeeklyReportMalware `json:"WeeklyReportMalwares,omitempty" name:"WeeklyReportMalwares"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportMalwaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefencePluginDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞防御插件信息详情列表

		List []*VulDefencePluginDetail `json:"List,omitempty" name:"List"`
		// 数据总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefencePluginDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefencePluginDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJavaMemShellPluginSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyJavaMemShellPluginSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJavaMemShellPluginSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetProcessCountRequest struct {
	*tchttp.BaseRequest

	// 搜索条件：搜索返回所有进程名包含Name的进程列表

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetProcessCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		List []*FileTamperRuleInfo `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileTamperRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseBindListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定机器列表信息

		List []*LicenseBindDetail `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseBindListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseBindListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExpertServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExpertServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExpertServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞防御事件列表

		List []*VulDefenceEvent `json:"List,omitempty" name:"List"`
		// 数据总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
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

type MalwareWhiteListAffectEvent struct {

	// 添加时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 唯一ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 影响的md5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type Component struct {

	// 组件名称。

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 组件类型。
	// <li>SYSTEM：系统组件</li>
	// <li>WEB：Web组件</li>

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 组件版本号。

	ComponentVersion *string `json:"ComponentVersion,omitempty" name:"ComponentVersion"`
	// 唯一ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机内网IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 组件检测更新时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 云镜客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeServersAndRiskAndFirstInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeServersAndRiskAndFirstInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServersAndRiskAndFirstInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessStatisticsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeProcessStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProcessStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanStateRequest struct {
	*tchttp.BaseRequest

	// 过滤参数;
	// <li>StrategyId 基线策略ID ,仅ModuleType 为 Baseline 时需要</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
}

func (r *DescribeScanStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellPluginInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// java内存马插件列表

		List []*JavaMemShellPluginInfo `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJavaMemShellPluginInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellPluginInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogDeliveryKafkaOptionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceList []*CKafkaInstanceInfo `json:"InstanceList,omitempty" name:"InstanceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogDeliveryKafkaOptionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDeliveryKafkaOptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebAppListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 列表

		WebApps []*AssetWebAppBaseInfo `json:"WebApps,omitempty" name:"WebApps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebAppListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableExpertServiceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应急响应可用次数

		EmergencyResponse *uint64 `json:"EmergencyResponse,omitempty" name:"EmergencyResponse"`
		// 是否购买过应急响应

		EmergencyResponseBuy *bool `json:"EmergencyResponseBuy,omitempty" name:"EmergencyResponseBuy"`
		// 安全管家订单

		ExpertService []*ExpertServiceOrderInfo `json:"ExpertService,omitempty" name:"ExpertService"`
		// 是否购买过安全管家

		ExpertServiceBuy *bool `json:"ExpertServiceBuy,omitempty" name:"ExpertServiceBuy"`
		// 旗舰护网可用次数

		ProtectNet *uint64 `json:"ProtectNet,omitempty" name:"ProtectNet"`
		// 是否购买过旗舰护网

		ProtectNetBuy *bool `json:"ProtectNetBuy,omitempty" name:"ProtectNetBuy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableExpertServiceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableExpertServiceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机信息与标签信息

		HostInfoList []*HostTagInfo `json:"HostInfoList,omitempty" name:"HostInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DuplicateHosts struct {

	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeBruteAttacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 暴力破解事件列表

		BruteAttacks []*BruteAttack `json:"BruteAttacks,omitempty" name:"BruteAttacks"`
		// 事件数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBruteAttacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetMachineListRequest struct {
	*tchttp.BaseRequest

	// 可选排序[FirstTime|PartitionCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>OsType - String - 是否必填：否 - windows或linux</li>
	// <li>CpuLoad - Int - 是否必填：否 -
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>
	// <li>DiskLoad - Int - 是否必填：否 -
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>
	// <li>MemLoad - Int - 是否必填：否 -
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportAssetMachineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetMachineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetAttackSettingRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeNetAttackSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetAttackSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRansomDefenseStrategyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRansomDefenseStrategyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRansomDefenseStrategyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 核心文件事件列表

		List []*FileTamperEvent `json:"List,omitempty" name:"List"`
		// 数据总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileTamperEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseStrategyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略详情

		Strategy *RansomDefenseStrategyDetail `json:"Strategy,omitempty" name:"Strategy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRansomDefenseStrategyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseStrategyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineAutoClearConfigRequest struct {
	*tchttp.BaseRequest

	// 清理规则时间, 0 表示关闭, 单位为天, 最大为 30天

	ClearRule *uint64 `json:"ClearRule,omitempty" name:"ClearRule"`
}

func (r *ModifyMachineAutoClearConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMachineAutoClearConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FullTextInfo struct {

	// 是否大小写敏感

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 是否包含中文

	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
	// 分词符

	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`
}

type VulHostInfo struct {

	// 主机别名

	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 是否是专业版主机

	IfProfessional *bool `json:"IfProfessional,omitempty" name:"IfProfessional"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机关联的Tag信息

	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAlarmAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 被暴力破解成功告警状态：
		// <li>OPEN：告警已开启</li>
		// <li>CLOSE： 告警已关闭</li>

		CrackSuccess *string `json:"CrackSuccess,omitempty" name:"CrackSuccess"`
		// 发现木马告警状态：
		// <li>OPEN：告警已开启</li>
		// <li>CLOSE： 告警已关闭</li>

		Malware *string `json:"Malware,omitempty" name:"Malware"`
		// 发现异地登录告警状态：
		// <li>OPEN：告警已开启</li>
		// <li>CLOSE： 告警已关闭</li>

		NonlocalLogin *string `json:"NonlocalLogin,omitempty" name:"NonlocalLogin"`
		// 防护软件离线告警状态：
		// <li>OPEN：告警已开启</li>
		// <li>CLOSE： 告警已关闭</li>

		Offline *string `json:"Offline,omitempty" name:"Offline"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件名称。

		ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
		// 组件类型。
		// <li>WEB：web组件</li>
		// <li>SYSTEM：系统组件</li>

		ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
		// 组件描述。

		Description *string `json:"Description,omitempty" name:"Description"`
		// 组件官网。

		Homepage *string `json:"Homepage,omitempty" name:"Homepage"`
		// 组件ID。

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanAssetRequest struct {
	*tchttp.BaseRequest

	// 资产指纹类型id列表

	AssetTypeIds []*uint64 `json:"AssetTypeIds,omitempty" name:"AssetTypeIds"`
	// Quuid列表

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *ScanAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefenceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件Id 可通过ExportTasks接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulDefenceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefenceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanVulResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防护状态及试用信息

		Data *ProductStatusInfo `json:"Data,omitempty" name:"Data"`
		// 接口调用返回状态码

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口调用返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentStatistics struct {

	// 组件名称。

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 组件类型。
	// <li>WEB：Web组件</li>
	// <li>SYSTEM：系统组件</li>

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 组件描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 组件ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机数量。

	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type VulLevelCountInfo struct {

	// 漏洞数量

	VulCount *uint64 `json:"VulCount,omitempty" name:"VulCount"`
	// 漏洞等级

	VulLevel *uint64 `json:"VulLevel,omitempty" name:"VulLevel"`
}

type WeeklyReportBruteAttack struct {

	// 攻击时间。

	AttackTime *string `json:"AttackTime,omitempty" name:"AttackTime"`
	// 尝试次数。

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 源IP。

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 被破解用户名。

	Username *string `json:"Username,omitempty" name:"Username"`
}

type CheckLogKafkaConnectionStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true连通，false不通

		IsConnect *bool `json:"IsConnect,omitempty" name:"IsConnect"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckLogKafkaConnectionStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLogKafkaConnectionStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检测进度

		Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportWebPageEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id 可通过 ExportTasks接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportWebPageEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportWebPageEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type City struct {

	// 城市Id

	CityId *string `json:"CityId,omitempty" name:"CityId"`
	// 省份Id

	ProvinceId *string `json:"ProvinceId,omitempty" name:"ProvinceId"`
}

type DescribeBanRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域信息列表

		RegionSet []*RegionSet `json:"RegionSet,omitempty" name:"RegionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBanRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBanRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineRule `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonthInspectionReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 巡检报告列表

		List []*MonthInspectionReport `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonthInspectionReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonthInspectionReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineFixListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBaselineFixListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineFixListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMachineTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMachineTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetAttackWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetAttackWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetAttackWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemDetectListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [HostCount|FirstTime|LastTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>PolicyId - int64 - 是否必填：否 - 规则Id</li>
	// <li>ItemName - string - 是否必填：否 - 项名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开始时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineItemDetectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemDetectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportInfoRequest struct {
	*tchttp.BaseRequest

	// 专业周报开始时间。

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
}

func (r *DescribeWeeklyReportInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBanModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBanModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBanModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReverseShellRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditPrivilegeRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditPrivilegeRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditPrivilegeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BillingDefineParams struct {

	// 自动绑定开关

	AutoBindSwitch *bool `json:"AutoBindSwitch,omitempty" name:"AutoBindSwitch"`
	// 自动加购订单是否自动续费

	AutoRepurchaseRenewSwitch *bool `json:"AutoRepurchaseRenewSwitch,omitempty" name:"AutoRepurchaseRenewSwitch"`
	// 自动加购开关

	AutoRepurchaseSwitch *bool `json:"AutoRepurchaseSwitch,omitempty" name:"AutoRepurchaseSwitch"`
	// 防护版本

	ProtectType *string `json:"ProtectType,omitempty" name:"ProtectType"`
}

type DescribeAttackSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击溯源数据

		AttackSource *AttackSource `json:"AttackSource,omitempty" name:"AttackSource"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLicenseOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 大订单号 , 后付费该字段空值

		BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`
		// 订单号列表

		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
		// 资源ID列表,预付费订单该字段空值

		ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLicenseOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLicenseOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsNewRequest struct {
	*tchttp.BaseRequest

	// 排序字段：CreateTime-发生时间。ModifyTime-处理时间

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>HostName - String - 是否必填：否 - 主机名</li>
	// <li>Hostip - String - 是否必填：否 - 主机内网IP</li>
	// <li>HostIp - String - 是否必填：否 - 主机内网IP</li>
	// <li>RuleCategory - Int - 是否必填：否 - 策略类型,全部或者单选(0:系统 1:用户)</li>
	// <li>RuleName - String - 是否必填：否 - 策略名称</li>
	// <li>RuleLevel - Int - 是否必填：否 - 威胁等级,可以多选</li>
	// <li>Status - Int - 是否必填：否 - 处理状态,可多选(0:待处理 1:已处理 2:已加白  3:已忽略 4:已删除 5:已拦截)</li>
	// <li>DetectBy - Int - 是否必填：否 - 数据来源,可多选(0:bash日志 1:实时监控)</li>
	// <li>StartTime - String - 是否必填：否 - 开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：根据请求次数排序：asc-升序/desc-降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBashEventsNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashEventsNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCanNotSeparateMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 不可隔离主机列表

		List []*CanNotSeparateInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCanNotSeparateMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCanNotSeparateMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHotVulTopRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeHotVulTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHotVulTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineRiskCntResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 高危命令

		Bash *uint64 `json:"Bash,omitempty" name:"Bash"`
		// 密码破解

		BruteAttack *uint64 `json:"BruteAttack,omitempty" name:"BruteAttack"`
		// 异地登录

		HostLogin *uint64 `json:"HostLogin,omitempty" name:"HostLogin"`
		// 恶意请求

		MaliciousRequest *uint64 `json:"MaliciousRequest,omitempty" name:"MaliciousRequest"`
		// 木马

		Malware *uint64 `json:"Malware,omitempty" name:"Malware"`
		// 本地提权

		PrivilegeEscalation *uint64 `json:"PrivilegeEscalation,omitempty" name:"PrivilegeEscalation"`
		// 反弹shell

		ReverseShell *uint64 `json:"ReverseShell,omitempty" name:"ReverseShell"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineRiskCntResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineRiskCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetInitServiceBaseInfo struct {

	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	//  附加信息
	//

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 服务器IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 服务器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 服务器外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 默认启用状态：0未启用，1启用

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 类型：
	// 1:编码器
	// 2:IE插件
	// 3:网络提供者
	// 4:镜像劫持
	// 5:LSA提供者
	// 6:KnownDLLs
	// 7:启动执行
	// 8:WMI
	// 9:计划任务
	// 10:Winsock提供者
	// 11:打印监控器
	// 12:资源管理器
	// 13:驱动服务
	// 14:登录

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 启动用户

	User *string `json:"User,omitempty" name:"User"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type OrderModifyObject struct {

	// 扩容/缩容数,变配子产品忽略该参数

	InquireNum *int64 `json:"InquireNum,omitempty" name:"InquireNum"`
	// 新产品标识,这里支持PRO_VERSION 专业版,FLAGSHIP 旗舰版

	NewSubProductCode *string `json:"NewSubProductCode,omitempty" name:"NewSubProductCode"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeRiskDnsEventListRequest struct {
	*tchttp.BaseRequest

	// 排序字段：[AccessCount:请求次数|LastTime:最近请求时间]

	By *string `json:"By,omitempty" name:"By"`
	// <li>IpOrName - String - 是否必填：否 - 主机Ip或别名筛选</li>
	// <li>HostId - String - 是否必填：否 - 主机Id</li>
	// <li>AgentId - String - 是否必填：否 - 客户端Id</li>
	// <li>PolicyType - String - 是否必填：否 - 策略类型,0:系统策略1:用户自定义策略</li>
	// <li>Domain - String - 是否必填：否 - 域名(先对域名做urlencode,再base64)</li>
	// <li>HandleStatus - String - 是否必填：否 - 状态筛选0:待处理；2:信任；3:不信任</li>
	// <li>BeginTime - String - 是否必填：否 - 最近访问开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 最近访问结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：根据请求次数排序：[asc:升序|desc:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRiskDnsEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseProVersionRequest struct {
	*tchttp.BaseRequest

	// 主机唯一标识Uuid数组。
	// 黑石的InstanceId，CVM的Uuid ,边缘计算的Uuid , 轻量应用服务器的Uuid ,混合云机器的Quuid 。 当前参数最大长度限制20

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *CloseProVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseProVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebLocationListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetWebLocationListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebLocationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportTasksRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExportTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineTagRequest struct {
	*tchttp.BaseRequest

	// 关联的标签ID

	Rid *uint64 `json:"Rid,omitempty" name:"Rid"`
}

func (r *DeleteMachineTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSystemPackageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Packages []*AssetSystemPackageInfo `json:"Packages,omitempty" name:"Packages"`
		// 记录总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSystemPackageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSystemPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalWareListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 木马列表

		MalWareList []*MalWareList `json:"MalWareList,omitempty" name:"MalWareList"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalWareListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalWareListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeEventsRequest struct {
	*tchttp.BaseRequest

	// 排序字段：CreateTime-发现时间

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机IP)</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：根据请求次数排序：asc-升序/desc-降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribePrivilegeEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivilegeEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseStrategyDetailRequest struct {
	*tchttp.BaseRequest

	// 策略ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeRansomDefenseStrategyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseStrategyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebServiceInfoListRequest struct {
	*tchttp.BaseRequest

	// 可选排序：[FirstTime|ProcessCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>User- string - 是否必填：否 - 运行用户</li>
	// <li>Name- string - 是否必填：否 - Web服务名：
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType- string - 是否必填：否 - Windows/linux</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetWebServiceInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebServiceInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogKafkaAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogKafkaAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogKafkaAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检测项Top列表

		RuleTopList []*BaselineRuleTopInfo `json:"RuleTopList,omitempty" name:"RuleTopList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCanFixVulMachineRequest struct {
	*tchttp.BaseRequest

	// 需要修复的主机，和VulIds是and 的关系

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 漏洞id 支持多个漏洞

	VulIds []*uint64 `json:"VulIds,omitempty" name:"VulIds"`
}

func (r *DescribeCanFixVulMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCanFixVulMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogExportsRequest struct {
	*tchttp.BaseRequest

	// 分页单页限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLogExportsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogExportsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackVulTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 威胁类型列表

		List []*string `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackVulTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackVulTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineDefenseCntResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击检测统计

		AttackLogs *uint64 `json:"AttackLogs,omitempty" name:"AttackLogs"`
		// 核心文件监控统计

		FileTamper *uint64 `json:"FileTamper,omitempty" name:"FileTamper"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineDefenseCntResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineDefenseCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseStrategyMachinesRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Ips- string- 主机名称</li>
	// <li>Status - Uint64 - 0未绑定，1已绑定</li>
	// <li>Names- String - 主机名称</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 防勒索策略ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRansomDefenseStrategyMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseStrategyMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMachineTagsRequest struct {
	*tchttp.BaseRequest

	// 服务器类型(CVM|BM|ECM|LH|Other)

	MachineArea *string `json:"MachineArea,omitempty" name:"MachineArea"`
	// 服务器地区 如: ap-guangzhou

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 机器 Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 标签ID，该操作会覆盖原有的标签列表

	TagIds []*uint64 `json:"TagIds,omitempty" name:"TagIds"`
}

func (r *UpdateMachineTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMachineTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BashEventsInfoNew struct {

	// 执行命令

	BashCmd *string `json:"BashCmd,omitempty" name:"BashCmd"`
	// 发生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 检测来源 0:bash日志 1:实时监控

	DetectBy *int64 `json:"DetectBy,omitempty" name:"DetectBy"`
	// 进程名称

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 描述

	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
	// 主机内网IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 数据ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机在线状态 OFFLINE  ONLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 0:普通 1:专业版 2:旗舰版

	MachineType *int64 `json:"MachineType,omitempty" name:"MachineType"`
	// 主机外网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 处理时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 进程号

	Pid *string `json:"Pid,omitempty" name:"Pid"`
	// 平台类型

	Platform *uint64 `json:"Platform,omitempty" name:"Platform"`
	// 进程树 json  pid:进程id，exe:文件路径 ，account:进程所属用组和用户 ,cmdline:执行命令，ssh_service: SSH服务ip, ssh_soure:登录源

	PsTree *string `json:"PsTree,omitempty" name:"PsTree"`
	// 主机ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 参考链接

	References []*string `json:"References,omitempty" name:"References"`
	// 自动生成的正则表达式

	RegexBashCmd *string `json:"RegexBashCmd,omitempty" name:"RegexBashCmd"`
	// 规则类别  0=系统规则，1=用户规则

	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`
	// 规则ID,等于0表示已规则已被删除或生效范围已修改

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则等级：1-高 2-中 3-低

	RuleLevel *uint64 `json:"RuleLevel,omitempty" name:"RuleLevel"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 处理状态： 0 = 待处理 1= 已处理, 2 = 已加白， 3= 已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 建议方案

	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 登录用户

	User *string `json:"User,omitempty" name:"User"`
	// 云镜ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeScanMalwareScheduleRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScanMalwareScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanMalwareScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMalwareWhiteListRequest struct {
	*tchttp.BaseRequest

	// 文件目录(正则)；长度不超过200个，内容base64转义

	FileDirectory []*string `json:"FileDirectory,omitempty" name:"FileDirectory"`
	// 文件后缀；长度不超过200个，内容base64转义（废弃）

	FileExtension []*string `json:"FileExtension,omitempty" name:"FileExtension"`
	// 文件名称(正则)；长度不超过200个

	FileName []*string `json:"FileName,omitempty" name:"FileName"`
	// 规则唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全部主机； 0否，1是。

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 匹配模式 ；0 精确匹配，1模糊匹配(废弃)

	MatchType *uint64 `json:"MatchType,omitempty" name:"MatchType"`
	// MD5列表

	Md5List []*string `json:"Md5List,omitempty" name:"Md5List"`
	// 白名单模式； 0MD5白名单，1自定义

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// quuid 列表

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
}

func (r *ModifyMalwareWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMalwareWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebPageProtectSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWebPageProtectSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebPageProtectSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基线信息列表

		BaselineList []*BaselineInfo `json:"BaselineList,omitempty" name:"BaselineList"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 标签列表

		List []*string `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaliciousRequestsRequest struct {
	*tchttp.BaseRequest

	// 恶意请求记录ID数组，(最大100条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteMaliciousRequestsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaliciousRequestsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBashPolicyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBashPolicyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBashPolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 查询关键字 </li>
	// <li>UserName - String - 是否必填：否 - 用户名筛选 </li>
	// <li>ModifyBeginTime - String - 是否必填：否 - 按照修改时间段筛选，开始时间 </li>
	// <li>ModifyEndTime - String - 是否必填：否 - 按照修改时间段筛选，结束时间 </li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLoginWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostInfoRequest struct {
	*tchttp.BaseRequest

	// 主机Quuid数组

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// Uuids 查询，Quuid查询时填空

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *DescribeHostInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenEventsCntRequest struct {
	*tchttp.BaseRequest

	// 数据类型：0:待处理风险总数 1:影响资产总数，默认为0

	BusinessType *uint64 `json:"BusinessType,omitempty" name:"BusinessType"`
}

func (r *DescribeScreenEventsCntRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenEventsCntRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebLocationListRequest struct {
	*tchttp.BaseRequest

	// 可选排序：[FirstTime|PathCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name - String - 是否必填：否 - 域名</li>
	// <li>User - String - 是否必填：否 - 运行用户</li>
	// <li>Port - uint64 - 是否必填：否 - 站点端口</li>
	// <li>Proto - uint64 - 是否必填：否 - 站点协议：1:HTTP,2:HTTPS</li>
	// <li>ServiceType - uint64 - 是否必填：否 - 服务类型：
	// 1:Tomcat
	// 2：Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetWebLocationListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebLocationListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageInfo struct {

	// agent镜像扫描错误

	AgentError *string `json:"AgentError,omitempty" name:"AgentError"`
	// 构建历史

	BuildHistory *string `json:"BuildHistory,omitempty" name:"BuildHistory"`
	// 关联容器个数

	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 关联主机数

	HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 是否可信镜像

	IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`
	// 镜像系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 风险行为数

	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
	// 后端镜像扫描错误

	ScanError *string `json:"ScanError,omitempty" name:"ScanError"`
	// 扫描状态

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 最近扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 敏感信息数

	SensitiveInfoCnt *uint64 `json:"SensitiveInfoCnt,omitempty" name:"SensitiveInfoCnt"`
	// 镜像大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 病毒木马数

	VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`
	// 漏洞数

	VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`
}

type StandardConfig struct {

	// 阻断时长，单位：秒

	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`
}

type DeleteMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马记录ID数组 (最大100条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteMalwaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMalwaresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAutoQuaraStatusRequest struct {
	*tchttp.BaseRequest

	// 是否开启自动隔离

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetAutoQuaraStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAutoQuaraStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordInfo struct {

	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 客户端离线时间

	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`
	// 客户端卸载调用链

	UninstallCmd *string `json:"UninstallCmd,omitempty" name:"UninstallCmd"`
	// 客户端卸载时间

	UninstallTime *string `json:"UninstallTime,omitempty" name:"UninstallTime"`
	// 客户端uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetUserInfoRequest struct {
	*tchttp.BaseRequest

	// 账户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 云服务器UUID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机安全UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttackRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 爆破阻断规则列表

		Rules []*BruteAttackRuleList `json:"Rules,omitempty" name:"Rules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBruteAttackRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBruteAttackRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebHookRuleDetail struct {

	// 机器人地址

	HookAddr *string `json:"HookAddr,omitempty" name:"HookAddr"`
	// 主机Id列表

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 主机范围

	HostLabels []*WebHookHostLabel `json:"HostLabels,omitempty" name:"HostLabels"`
	// 是否启用[1:禁用|0:启用]

	IsDisabled *int64 `json:"IsDisabled,omitempty" name:"IsDisabled"`
	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 事件类型

	RuleItems []*WebHookEventKv `json:"RuleItems,omitempty" name:"RuleItems"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 备注信息

	RuleRemark *string `json:"RuleRemark,omitempty" name:"RuleRemark"`
}

type DescribeAssetWebLocationListRequest struct {
	*tchttp.BaseRequest

	// 可选排序：[FirstTime|PathCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 域名</li>
	// <li>User - String - 是否必填：否 - 运行用户</li>
	// <li>Port - uint64 - 是否必填：否 - 站点端口</li>
	// <li>Proto - uint64 - 是否必填：否 - 站点协议：1:HTTP,2:HTTPS</li>
	// <li>ServiceType - uint64 - 是否必填：否 - 服务类型：
	// 1:Tomcat
	// 2：Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetWebLocationListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebServiceInfoListRequest struct {
	*tchttp.BaseRequest

	// 可选排序：[FirstTime|ProcessCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>User- string - 是否必填：否 - 运行用户</li>
	// <li>Name- string - 是否必填：否 - Web服务名：
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType- string - 是否必填：否 - Windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetWebServiceInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineEffectHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务id 可通过 ExportTasks接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBaselineEffectHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostEvent struct {

	// 主机名称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 用户ID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 防护目录ID

	ConfigId *uint64 `json:"ConfigId,omitempty" name:"ConfigId"`
	// 文件是否已恢复：0-未恢复；1-已恢复

	HasRecovered *uint64 `json:"HasRecovered,omitempty" name:"HasRecovered"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 时间主键ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 防护路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// CVM uuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 恢复类型

	RecoverType *uint64 `json:"RecoverType,omitempty" name:"RecoverType"`
	// 恢复时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type EditReverseShellRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditReverseShellRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditReverseShellRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionWhiteConfig struct {

	// 到期天数

	Deadline *uint64 `json:"Deadline,omitempty" name:"Deadline"`
	// 是否可申请

	IsApplyFor *bool `json:"IsApplyFor,omitempty" name:"IsApplyFor"`
	// 授权数量

	LicenseNum *uint64 `json:"LicenseNum,omitempty" name:"LicenseNum"`
	// 类型

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
}

type ModifyWebHookRuleRequest struct {
	*tchttp.BaseRequest

	// 规则内容

	Data *WebHookRuleDetail `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyWebHookRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebHookRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 索引修改时间，初始值为索引创建时间

		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
		// 索引规则

		Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
		// 是否生效

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenGeneralStatRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScreenGeneralStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenGeneralStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWebHookRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWebHookRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWebHookRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebAppPluginListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Plugins []*AssetWebAppPluginInfo `json:"Plugins,omitempty" name:"Plugins"`
		// 分区总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebAppPluginListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppPluginListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostLoginListRequest struct {
	*tchttp.BaseRequest

	// 排序字段：LoginTime-发生时间

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Uuid - String - 是否必填：否 - 云镜唯一Uuid</li>
	// <li>Quuid - String - 是否必填：否 - 云服务器uuid</li>
	// <li>UserName - String - 是否必填：否 - 用户名筛选</li>
	// <li>LoginTimeBegin - String - 是否必填：否 - 按照修改时间段筛选，开始时间</li>
	// <li>LoginTimeEnd - String - 是否必填：否 - 按照修改时间段筛选，结束时间</li>
	// <li>SrcIp - String - 是否必填：否 - 来源ip筛选</li>
	// <li>Status - int - 是否必填：否 - 状态筛选1:正常登录；5：已加白,14:已处理，15：已忽略</li>
	// <li>RiskLevel - int - 是否必填：否 - 状态筛选0:高危；1：可疑</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：根据请求次数排序：asc-升序/desc-降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeHostLoginListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostLoginListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CKafkaInstanceInfo struct {

	// 地域

	Az *string `json:"Az,omitempty" name:"Az"`
	// 实例带宽，单位Mbps

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 磁盘容量，单位GB

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 状态，1表示健康，2表示告警，3 表示实例状态异常

	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 版本号

	KafkaVersion *string `json:"KafkaVersion,omitempty" name:"KafkaVersion"`
	// 路由列表

	RouteList []*CKafkaRouteInfo `json:"RouteList,omitempty" name:"RouteList"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 主题列表

	TopicList []*CKafkaTopicInfo `json:"TopicList,omitempty" name:"TopicList"`
	// vpcId，如果为空，说明是基础网络

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ProtectHostConfig struct {

	// 自动恢复开关 0 关闭 1开启

	AutoRecovery *uint64 `json:"AutoRecovery,omitempty" name:"AutoRecovery"`
	// 防护开关 0  关闭 1开启

	ProtectSwitch *uint64 `json:"ProtectSwitch,omitempty" name:"ProtectSwitch"`
	// 机器唯一ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type DescribeWebHookRulesRequest struct {
	*tchttp.BaseRequest

	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 [asc:升序|desc:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeWebHookRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebHookRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportJavaMemShellPluginsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：Keywords: ip或者主机名模糊查询, Pid精确匹配，MainClass模糊匹配

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 导出字段

	Where []*string `json:"Where,omitempty" name:"Where"`
}

func (r *ExportJavaMemShellPluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportJavaMemShellPluginsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专业周报异地登录数据。

		WeeklyReportNonlocalLoginPlaces []*WeeklyReportNonlocalLoginPlace `json:"WeeklyReportNonlocalLoginPlaces,omitempty" name:"WeeklyReportNonlocalLoginPlaces"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBaselineWeakPasswordRequest struct {
	*tchttp.BaseRequest

	// 无

	Data []*BaselineWeakPassword `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyBaselineWeakPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBaselineWeakPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentInstallationTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安装命令token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentInstallationTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentInstallationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBanStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBanStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBanStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseListRequest struct {
	*tchttp.BaseRequest

	// 多个条件筛选时 LicenseStatus,DeadlineStatus,ResourceId,Keywords 取交集
	// <li> LicenseStatus 授权状态信息,0 未使用,1 部分使用, 2 已用完, 3 不可用  4 可使用</li>
	// <li> BuyTime 购买时间</li>
	// <li> LicenseType  授权类型, 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月</li>
	// <li>DeadlineStatus 到期状态 NotExpired 未过期, Expire 已过期(包含已销毁) NearExpiry 即将到期</li>
	// <li>ResourceId 资源ID</li>
	// <li>Keywords IP筛选</li>
	// <li>PayMode 付费模式 0 按量计费 , 1 包年包月</li>
	// <li>OrderStatus 订单状态 1 正常 2 隔离 3 销毁</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10.

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0.

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 资源ID列表筛选(与Filters.ResourceId 不能共存,同时出现则优先使用ResourceIds)

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 标签筛选,平台标签能力,这里传入 标签键,标签值作为一个对象

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeLicenseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBaselineRuleIgnoreResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBaselineRuleIgnoreResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBaselineRuleIgnoreResponse) FromJsonString(s string) error {
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

type ExportAssetSystemPackageListRequest struct {
	*tchttp.BaseRequest

	// 排序方式可选：[FistTime|InstallTime:安装时间]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name - String - 是否必填：否 - 包 名</li>
	// <li>StartTime - String - 是否必填：否 - 安装开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 安装开始时间</li>
	// <li>Type - int - 是否必填：否 - 安装包类型：
	// 1:rmp
	// 2:dpkg
	// 3:java
	// 4:system</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetSystemPackageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetSystemPackageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表内容

		List []*BashRule `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulFixStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修复结束时间，为空表示还没结束

		FixEndTime *string `json:"FixEndTime,omitempty" name:"FixEndTime"`
		// 修复失败的主机数

		FixFailCnt *uint64 `json:"FixFailCnt,omitempty" name:"FixFailCnt"`
		// 修复的任务id

		FixId *uint64 `json:"FixId,omitempty" name:"FixId"`
		// 修复精度 0-100

		FixProgress *uint64 `json:"FixProgress,omitempty" name:"FixProgress"`
		// 开始修复时间

		FixStartTime *string `json:"FixStartTime,omitempty" name:"FixStartTime"`
		// 修复成功的主机数

		FixSuccessCnt *uint64 `json:"FixSuccessCnt,omitempty" name:"FixSuccessCnt"`
		// 主机总是

		HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`
		// 是否允许重试 0:不允许 1：允许

		IsAllowRetry *uint64 `json:"IsAllowRetry,omitempty" name:"IsAllowRetry"`
		// 快照是否是重试状态 0=非重试  1=重试

		IsRetrySnapshot *uint64 `json:"IsRetrySnapshot,omitempty" name:"IsRetrySnapshot"`
		// 预计剩余时间（单位秒）

		RemainingTime *uint64 `json:"RemainingTime,omitempty" name:"RemainingTime"`
		// 快照创建失败数

		SnapshotFailCnt *uint64 `json:"SnapshotFailCnt,omitempty" name:"SnapshotFailCnt"`
		// 修复快照状态列表

		SnapshotList []*VulFixStatusSnapshotInfo `json:"SnapshotList,omitempty" name:"SnapshotList"`
		// 快照创建进度0-100

		SnapshotProgress *uint64 `json:"SnapshotProgress,omitempty" name:"SnapshotProgress"`
		// 修复漏洞详情列表

		VulFixList []*VulFixStatusInfo `json:"VulFixList,omitempty" name:"VulFixList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulFixStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulFixStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckFileTamperRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0=校验通过  1=规则名称校验不通过

		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`
		// 校验信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckFileTamperRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFileTamperRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOpenPortTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOpenPortTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOpenPortTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineDetectOverviewRequest struct {
	*tchttp.BaseRequest

	// 策略Id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeBaselineDetectOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineDetectOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentsRequest struct {
	*tchttp.BaseRequest

	// 组件ID。Uuid和ComponentId必填其一，使用ComponentId表示，查询该组件列表信息。

	ComponentId *uint64 `json:"ComponentId,omitempty" name:"ComponentId"`
	// 过滤条件。
	// <li>ComponentVersion - String - 是否必填：否 - 组件版本号</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 云镜客户端唯一Uuid。Uuid和ComponentId必填其一，使用Uuid表示，查询该主机列表信息。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenAttackHotspotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击热点列表

		List []*ScreenAttackHotspot `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenAttackHotspotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenAttackHotspotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginWhiteInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoginWhiteInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginWhiteInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionQuickRequest struct {
	*tchttp.BaseRequest

	// 活动ID

	ActivityId *uint64 `json:"ActivityId,omitempty" name:"ActivityId"`
}

func (r *OpenProVersionQuickRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenProVersionQuickRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBashPoliciesRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBashPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBashPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineRuleListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列

	By *string `json:"By,omitempty" name:"By"`
	// <li>RuleName - String - 是否必填：否 - 规则名称</li>
	// <li>CategoryId - int64 - 是否必填：否 自定义筛选为-1 - 规则分类</li>
	// <li>RuleType - int - 是否必填：否 0:系统 1:自定义 - 规则类型</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClientExceptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件详情

		Records []*RecordInfo `json:"Records,omitempty" name:"Records"`
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClientExceptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClientExceptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各软件应用数量

		Apps []*AssetKeyVal `json:"Apps,omitempty" name:"Apps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetAppCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgnoreHostAndItemConfigRequest struct {
	*tchttp.BaseRequest

	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>HostId - string - 是否必填：否 - 主机Id</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIgnoreHostAndItemConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgnoreHostAndItemConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTagsRequest struct {
	*tchttp.BaseRequest

	// 标签ID (最大100 条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetDatabaseListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>User- string - 是否必填：否 - 运行用户</li>
	// <li>Ip - String - 是否必填：否 - 绑定IP</li>
	// <li>Port - Int - 是否必填：否 - 端口</li>
	// <li>Name - Int - 是否必填：否 - 数据库名称
	// 0:全部
	// 1:MySQL
	// 2:Redis
	// 3:Oracle
	// 4:MongoDB
	// 5:MemCache
	// 6:PostgreSQL
	// 7:HBase
	// 8:DB2
	// 9:Sybase
	// 10:TiDB</li>
	// <li>Proto - String - 是否必填：否 - 协议：1:TCP, 2:UDP, 3:未知</li>
	// <li>OsType - String - 是否必填：否 - 操作系统: linux/windows</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetDatabaseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetDatabaseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeEventInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本地提权详情

		PrivilegeEventInfo *PrivilegeEventInfo `json:"PrivilegeEventInfo,omitempty" name:"PrivilegeEventInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivilegeEventInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivilegeEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeStrategyEnableStatusRequest struct {
	*tchttp.BaseRequest

	// 开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 基线策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *ChangeStrategyEnableStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeStrategyEnableStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetCoreModuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Modules []*AssetCoreModuleBaseInfo `json:"Modules,omitempty" name:"Modules"`
		// 总数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetCoreModuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetCoreModuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetAttackTopInfo struct {

	// 网络攻击主机维度top统计数据

	Agent []*TopInfo `json:"Agent,omitempty" name:"Agent"`
	// 网络攻击目标端口维度top统计数据

	DstPort []*TopInfo `json:"DstPort,omitempty" name:"DstPort"`
	// 网络攻击ip来源维度top统计数据

	SrcIp []*TopInfo `json:"SrcIp,omitempty" name:"SrcIp"`
	// 网络攻击漏洞维度top统计数据

	Vul []*TopInfo `json:"Vul,omitempty" name:"Vul"`
}

type DescribeBaselineHostDetectListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [LastTime|ItemCount|PassedItemCount|NotPassedItemCount|FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>HostName - string - 是否必填：否 - 主机名称</i>
	// <li>HostIp - string - 是否必填：否 - 主机Ip</i>
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开时时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineHostDetectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineHostDetectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenBroadcastsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScreenBroadcastsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenBroadcastsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebServiceInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 列表

		WebServices []*AssetWebServiceBaseInfo `json:"WebServices,omitempty" name:"WebServices"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebServiceInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGeneralResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 阿里云服务商机器数

		AliCloudMachineCnt *uint64 `json:"AliCloudMachineCnt,omitempty" name:"AliCloudMachineCnt"`
		// 百度云服务商机器数

		BaiduCloudMachineCnt *uint64 `json:"BaiduCloudMachineCnt,omitempty" name:"BaiduCloudMachineCnt"`
		// 已防护基础版机器数

		BaseMachineCnt *uint64 `json:"BaseMachineCnt,omitempty" name:"BaseMachineCnt"`
		// 比较昨日即将到期的机器数

		CompareYesterdayDeadlineMachineCnt *int64 `json:"CompareYesterdayDeadlineMachineCnt,omitempty" name:"CompareYesterdayDeadlineMachineCnt"`
		// 比较昨日新增的主机数

		CompareYesterdayMachineCnt *int64 `json:"CompareYesterdayMachineCnt,omitempty" name:"CompareYesterdayMachineCnt"`
		// 比较昨日未防护机器数

		CompareYesterdayNotProtectMachineCnt *int64 `json:"CompareYesterdayNotProtectMachineCnt,omitempty" name:"CompareYesterdayNotProtectMachineCnt"`
		// 比较昨日风险机器数

		CompareYesterdayRiskMachineCnt *int64 `json:"CompareYesterdayRiskMachineCnt,omitempty" name:"CompareYesterdayRiskMachineCnt"`
		// 即将到期的机器数

		DeadlineMachineCnt *uint64 `json:"DeadlineMachineCnt,omitempty" name:"DeadlineMachineCnt"`
		// 已防护旗舰版机器数

		FlagshipMachineCnt *uint64 `json:"FlagshipMachineCnt,omitempty" name:"FlagshipMachineCnt"`
		// IDC机器数

		IDCMachineCnt *uint64 `json:"IDCMachineCnt,omitempty" name:"IDCMachineCnt"`
		// 已防护普惠版机器数（Lighthouse机器）

		LHGeneralDiscountCnt *uint64 `json:"LHGeneralDiscountCnt,omitempty" name:"LHGeneralDiscountCnt"`
		// 资产总数

		MachineCnt *uint64 `json:"MachineCnt,omitempty" name:"MachineCnt"`
		// 自动清理时间,最大720小时,最小0, 默认0 ,0=关闭

		MachineDestroyAfterOfflineHours *uint64 `json:"MachineDestroyAfterOfflineHours,omitempty" name:"MachineDestroyAfterOfflineHours"`
		// 未防护机器数

		NotProtectMachineCnt *uint64 `json:"NotProtectMachineCnt,omitempty" name:"NotProtectMachineCnt"`
		// 其他云服务商机器数

		OtherCloudMachineCnt *uint64 `json:"OtherCloudMachineCnt,omitempty" name:"OtherCloudMachineCnt"`
		// 已防护机器数

		ProtectMachineCnt *uint64 `json:"ProtectMachineCnt,omitempty" name:"ProtectMachineCnt"`
		// 存在风险的机器数

		RiskMachineCnt *uint64 `json:"RiskMachineCnt,omitempty" name:"RiskMachineCnt"`
		// 已防护专业版机器数

		SpecialtyMachineCnt *uint64 `json:"SpecialtyMachineCnt,omitempty" name:"SpecialtyMachineCnt"`
		// 腾讯云服务商机器数

		TencentCloudMachineCnt *uint64 `json:"TencentCloudMachineCnt,omitempty" name:"TencentCloudMachineCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineGeneralResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineGeneralResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebPageGeneralizeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防护目录数

		ProtectDirNum *uint64 `json:"ProtectDirNum,omitempty" name:"ProtectDirNum"`
		// 防护文件数

		ProtectFileNum *uint64 `json:"ProtectFileNum,omitempty" name:"ProtectFileNum"`
		// 防护主机数

		ProtectHostNum *uint64 `json:"ProtectHostNum,omitempty" name:"ProtectHostNum"`
		// 防护监测 0 未开启 1 已开启 2 异常

		ProtectMonitor *uint64 `json:"ProtectMonitor,omitempty" name:"ProtectMonitor"`
		// 今日防护数

		ProtectToday *uint64 `json:"ProtectToday,omitempty" name:"ProtectToday"`
		// 篡改文件数

		TamperFileNum *uint64 `json:"TamperFileNum,omitempty" name:"TamperFileNum"`
		// 篡改数

		TamperNum *uint64 `json:"TamperNum,omitempty" name:"TamperNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebPageGeneralizeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebPageGeneralizeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExpertServiceTimeRequest struct {
	*tchttp.BaseRequest

	// 选中的订单

	OrderId *uint64 `json:"OrderId,omitempty" name:"OrderId"`
	// 选中的主机quuid

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *CreateExpertServiceTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExpertServiceTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 磁盘分区列表

		Disks []*AssetDiskPartitionInfo `json:"Disks,omitempty" name:"Disks"`
		// 分区总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulFixTaskQuuids struct {

	// 需要修复漏洞的主机，所有主机必须有VulId的这个漏洞且是待修复状态。

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

type DescribeAssetTotalCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各项资源数量
		// system : 资源监控
		// account: 账号
		// port: 端口
		// process: 进程
		// app: 应用软件
		// database:数据库
		// webapp: Web应用
		// webframe: Web框架
		// webservice: Web服务
		// weblocation: Web站点
		// jar: Jar包
		// initservice: 启动服务
		// planTask:计划任务
		// env:环境变量
		// coremoudle:内核模块

		Types []*AssetKeyVal `json:"Types,omitempty" name:"Types"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetTotalCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTotalCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineFixListRequest struct {
	*tchttp.BaseRequest

	// 0:过滤的结果导出；1:全部导出

	ExportAll *int64 `json:"ExportAll,omitempty" name:"ExportAll"`
	// <li>ItemName - String - 是否必填：否 - 项名称</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBaselineFixListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineFixListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivilegeEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulHostCountScanTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否第一次检测

		IfFirstScan *bool `json:"IfFirstScan,omitempty" name:"IfFirstScan"`
		// 最后一次修复漏洞的时间

		LastFixTime *string `json:"LastFixTime,omitempty" name:"LastFixTime"`
		// 扫描时间

		ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
		// 运行中的任务号, 没有任务则为0

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 总漏洞数

		TotalVulCount *uint64 `json:"TotalVulCount,omitempty" name:"TotalVulCount"`
		// 漏洞影响主机数

		VulHostCount *uint64 `json:"VulHostCount,omitempty" name:"VulHostCount"`
		// 是否有支持自动修复的漏洞事件

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulHostCountScanTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulHostCountScanTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchBashRulesRequest struct {
	*tchttp.BaseRequest

	// 是否禁用

	Disabled *uint64 `json:"Disabled,omitempty" name:"Disabled"`
	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *SwitchBashRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchBashRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseOrder struct {

	// 授权ID

	LicenseId *uint64 `json:"LicenseId,omitempty" name:"LicenseId"`
	// 授权类型

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 订单类型

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
	// 授权订单资源状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type MalWareList struct {

	// 主机别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// '木马检测平台用,分割 1云查杀引擎、2TAV、3binaryAi、4异常行为、5威胁情报

	CheckPlatform *string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 首次运行时间

	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`
	// 木马文件是否存在 0:不存在，1:存在

	FileExists *uint64 `json:"FileExists,omitempty" name:"FileExists"`
	// 最近运行时间

	FileModifierTime *string `json:"FileModifierTime,omitempty" name:"FileModifierTime"`
	// 路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 服务器ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 最近扫描时间

	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
	// 风险等级 0未知、1低、2中、3高、4严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 木马样本md5

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 木马进程是否存在 0:不存在，1:存在

	ProcessExists *uint64 `json:"ProcessExists,omitempty" name:"ProcessExists"`
	// cvm quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 状态；4-:待处理，5-已信任，6-已隔离，8-文件已删除, 14:已处理

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 特性标签，已废弃字段，不会再返回标签，详情中才会返回标签信息

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 唯一UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 描述

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
}

type DescribeUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeUsualLoginPlacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsualLoginPlacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJavaMemShellPluginSwitchRequest struct {
	*tchttp.BaseRequest

	// 插件目标状态：0: 关闭 1: 开启

	JavaShellStatus *uint64 `json:"JavaShellStatus,omitempty" name:"JavaShellStatus"`
	// 主机quuid数组

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *ModifyJavaMemShellPluginSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJavaMemShellPluginSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单ID (最大 100 条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteLoginWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoginWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerRelatedDirInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务器IP

		HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
		// 服务器名称

		HostName *string `json:"HostName,omitempty" name:"HostName"`
		// 防护目录数量

		ProtectDirNum *uint64 `json:"ProtectDirNum,omitempty" name:"ProtectDirNum"`
		// 防护文件数量

		ProtectFileNum *uint64 `json:"ProtectFileNum,omitempty" name:"ProtectFileNum"`
		// 防护软链数量

		ProtectLinkNum *uint64 `json:"ProtectLinkNum,omitempty" name:"ProtectLinkNum"`
		// 防篡改数量

		ProtectTamperNum *uint64 `json:"ProtectTamperNum,omitempty" name:"ProtectTamperNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerRelatedDirInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerRelatedDirInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBaselinePolicyStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBaselinePolicyStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBaselinePolicyStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebServiceProcessListRequest struct {
	*tchttp.BaseRequest

	// Web服务ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebServiceProcessListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceProcessListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsInfoNewRequest struct {
	*tchttp.BaseRequest

	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeBashEventsInfoNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashEventsInfoNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件列表数据。

		Components []*Component `json:"Components,omitempty" name:"Components"`
		// 组件列表记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckBashRuleParamsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0=校验通过  1=规则名称校验不通过 2=正则表达式校验不通过

		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`
		// 校验信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckBashRuleParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBashRuleParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞top列表

		VulTopList []*VulTopInfo `json:"VulTopList,omitempty" name:"VulTopList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CKafkaTopicInfo struct {

	// 主题ID

	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ExportMaliciousRequestsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务Id , 可通过ExportTasks 接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportMaliciousRequestsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出的文件下载url（已弃用！）

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出文件Id 可通过ExportTasks接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmergencyVul struct {

	// 漏洞分类

	Category *uint64 `json:"Category,omitempty" name:"Category"`
	// cve编号

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// CVSS评分

	CvssScore *float64 `json:"CvssScore,omitempty" name:"CvssScore"`
	// 已防御的攻击次数

	DefenseAttackCount *uint64 `json:"DefenseAttackCount,omitempty" name:"DefenseAttackCount"`
	// 影响机器数

	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`
	// 是否支持防御， 0:不支持 1:支持

	IsSupportDefense *uint64 `json:"IsSupportDefense,omitempty" name:"IsSupportDefense"`
	// 漏洞标签 多个逗号分割

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 最后扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 漏洞级别

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 扫描进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 发布日期

	PublishDate *string `json:"PublishDate,omitempty" name:"PublishDate"`
	// 漏洞状态 0未检测 1有风险 ，2无风险 ，3 检查中展示progress

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type AssetPlanTask struct {

	// 执行命令或脚本

	Command *string `json:"Command,omitempty" name:"Command"`
	// 配置文件路径

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 执行周期

	Cycle *string `json:"Cycle,omitempty" name:"Cycle"`
	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	//  附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 服务器IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 服务器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 服务器外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 操作系统

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 默认启用状态：1启用，2未启用

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 启动用户

	User *string `json:"User,omitempty" name:"User"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeBanWhiteListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBanWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBanWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetDatabaseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetDatabaseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetDatabaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CKafkaRouteInfo struct {

	// 接入类型

	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名端口

	DomainPort *uint64 `json:"DomainPort,omitempty" name:"DomainPort"`
	// 路由ID

	RouteID *int64 `json:"RouteID,omitempty" name:"RouteID"`
	// 虚拟ip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 虚拟ip类型

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
}

type ModifyBruteAttackRulesRequest struct {
	*tchttp.BaseRequest

	// 暴力破解判断规则

	Rules []*BruteAttackRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyBruteAttackRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBruteAttackRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProtectDirRelatedServer struct {

	// 是否已经授权

	Authorization *bool `json:"Authorization,omitempty" name:"Authorization"`
	// 自动恢复开关

	AutoRestoreSwitchStatus *uint64 `json:"AutoRestoreSwitchStatus,omitempty" name:"AutoRestoreSwitchStatus"`
	// 异常状态

	Exception *uint64 `json:"Exception,omitempty" name:"Exception"`
	// 异常信息

	ExceptionMessage *string `json:"ExceptionMessage,omitempty" name:"ExceptionMessage"`
	// 服务器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 服务器名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 唯一ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 服务器系统

	MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`
	// 过渡进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 防护状态

	ProtectStatus *uint64 `json:"ProtectStatus,omitempty" name:"ProtectStatus"`
	// 防护开关

	ProtectSwitch *uint64 `json:"ProtectSwitch,omitempty" name:"ProtectSwitch"`
	// 服务器唯一ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 关联目录数

	RelateDirNum *uint64 `json:"RelateDirNum,omitempty" name:"RelateDirNum"`
}

type RansomDefenseStrategy struct {

	// 备份模式： 0按周，1按天

	BackupType *uint64 `json:"BackupType,omitempty" name:"BackupType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 策略备注

	Description *string `json:"Description,omitempty" name:"Description"`
	// 包含目录，用;分隔

	ExcludeDir *string `json:"ExcludeDir,omitempty" name:"ExcludeDir"`
	// 备份执行时间点(0-23): 11:00;12:00

	Hour *string `json:"Hour,omitempty" name:"Hour"`
	// 策略id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 包含目录，用;分隔

	IncludeDir *string `json:"IncludeDir,omitempty" name:"IncludeDir"`
	// 是否对所有主机生效

	IsAll *uint64 `json:"IsAll,omitempty" name:"IsAll"`
	// 绑定机器数

	MachineCount *uint64 `json:"MachineCount,omitempty" name:"MachineCount"`
	// 最近修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 保存天数，0永久保存

	SaveDay *uint64 `json:"SaveDay,omitempty" name:"SaveDay"`
	// 开启状态：0关闭，1开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 操作uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 备份星期天数（1-7）：1;2;3;4

	Weekday *string `json:"Weekday,omitempty" name:"Weekday"`
}

type DescribeScanTaskDetailsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeScanTaskDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanTaskDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDurationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务下发需要的时长，单位为分钟

		Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskDurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulLevelCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计结果

		VulLevelList []*VulLevelInfo `json:"VulLevelList,omitempty" name:"VulLevelList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulLevelCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulLevelCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBaselineWeakPasswordRequest struct {
	*tchttp.BaseRequest

	// 弱口令Id

	PasswordIds []*int64 `json:"PasswordIds,omitempty" name:"PasswordIds"`
}

func (r *DeleteBaselineWeakPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBaselineWeakPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebFrameListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetWebFrameListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebFrameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsPolicyListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [HostScope|UpdateTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>PolicyType - int - 是否必填：否 - 策略类型</li>
	// <li>PolicyName - string - 是否必填：否 - 策略名称</li>
	// <li>Domain - string - 是否必填：否 - 域名(先对域名做urlencode,再base64)</li>
	// <li>PolicyAction- int - 是否必填：否 - 策略动作</li>
	// <li>IsEnabled - int - 是否必填：否 - 是否生效</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRiskDnsPolicyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsPolicyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineBasicInfo struct {

	// 基线id

	BaselineId *uint64 `json:"BaselineId,omitempty" name:"BaselineId"`
	// 基线名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 父级id

	ParentId *uint64 `json:"ParentId,omitempty" name:"ParentId"`
}

type DeleteRiskDnsPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRiskDnsPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRiskDnsPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserListRequest struct {
	*tchttp.BaseRequest

	// 可选排序：[FirstTime|LoginTime|PasswordChangeTime|PasswordDuaTime]
	// PasswordLockDays

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 账户名（模糊匹配）</li>
	// <li>NameStrict - String - 是否必填：否 - 账户名（严格匹配）</li>
	// <li>Uid - uint64 - 是否必填：否 - Uid</li>
	// <li>Guid - uint64 - 是否必填：否 - Guid</li>
	// <li>LoginTimeStart - String - 是否必填：否 - 开始时间，如：2021-01-11</li>
	// <li>LoginTimeEnd - String - 是否必填：否 - 结束时间，如：2021-01-11</li>
	// <li>LoginType - uint64 - 是否必填：否 - 0-不可登录；1-只允许key登录；2只允许密码登录；3-允许key和密码 仅linux</li>
	// <li>OsType - String - 是否必填：否 - windows或linux</li>
	// <li>Status - uint64 - 是否必填：否 - 账号状态：0-禁用；1-启用</li>
	// <li>UserType - uint64 - 是否必填：否 - 账号类型：0访客用户，1标准用户，2管理员用户 仅windows</li>
	// <li>IsDomain - uint64 - 是否必填：否 - 是否域账号：0 不是，1是 仅windows
	// <li>IsRoot - uint64 - 是否必填：否 - 是否Root权限：0 不是，1是 仅linux
	// <li>IsSudo - uint64 - 是否必填：否 - 是否Sudo权限：0 不是，1是 仅linux</li>
	// <li>IsSshLogin - uint64 - 是否必填：否 - 是否ssh登录：0 不是，1是 仅linux</li>
	// <li>ShellLoginStatus - uint64 - 是否必填：否 - 是否shell登录性，0不是；1是 仅linux</li>
	// <li>PasswordStatus - uint64 - 是否必填：否 - 密码状态：1正常 2即将过期 3已过期 4已锁定 仅linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetUserListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashPoliciesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Name - String - 是否必填：否 - 规则名称</li>
	// <li>Rule - String - 是否必填：否 - 规则内容</li>
	// <li>Level - Int - 是否必填：否 - 威胁等级</li>
	// <li>White - Int - 是否必填：否 - 白名单类型</li>
	// <li>Category - Int - 是否必填：否 - 策略类型</li>
	// <li>BashAction - Int - 是否必填：否 - 操作动作</li>
	// <li>Enable - Int - 是否必填：否 - 生效状态</li>
	// <li>Id - Int - 是否必填：否 - 策略ID</li>
	// <li>PolicyId - Int - 是否必填：否 - 策略ID</li>
	// <li>RuleId - Int - 是否必填：否 - 策略ID</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBashPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareTimingScanSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMalwareTimingScanSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareTimingScanSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulEffectHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 影响主机列表

		VulEffectHostList []*VulEffectHostList `json:"VulEffectHostList,omitempty" name:"VulEffectHostList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulEffectHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalWareListRequest struct {
	*tchttp.BaseRequest

	// 检测排序 CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>FilePath - String - 是否必填：否 - 路径筛选</li>
	// <li>VirusName - String - 是否必填：否 - 描述筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选-开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选-结束时间</li>
	// <li>Status - String - 是否必填：否 - 状态筛选 4待处理,5信任,6已隔离,10隔离中,11恢复隔离中</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 ASC,DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeMalWareListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalWareListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousRequestWhiteListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	//
	// <li>Domain  - String - 基线名称</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMaliciousRequestWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousRequestWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefenceRangeDetail struct {

	// cve id

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// cvss 分数

	CvssScore *float64 `json:"CvssScore,omitempty" name:"CvssScore"`
	// 标签

	Label *string `json:"Label,omitempty" name:"Label"`
	// 漏洞级别：  1低危 2中危 3高危 4严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 发布时间

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 漏洞id

	VulId *int64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type DescribeAttackEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击事件列表

		List []*NetAttackEvent `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 帐号变更历史数据数组。

		HistoryAccounts []*HistoryAccount `json:"HistoryAccounts,omitempty" name:"HistoryAccounts"`
		// 帐号变更历史列表记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHistoryAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseStateRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRansomDefenseStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteCombinedListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>UserName - String - 是否必填：否 - 用户名筛选</li>
	// <li>ModifyBeginTime - String - 是否必填：否 - 按照修改时间段筛选，开始时间</li>
	// <li>ModifyEndTime - String - 是否必填：否 - 按照修改时间段筛选，结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLoginWhiteCombinedListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginWhiteCombinedListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求事件详情

		RiskDnsInfo *RiskDnsList `json:"RiskDnsInfo,omitempty" name:"RiskDnsInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetSystemPackageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetSystemPackageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetSystemPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFileTamperRuleRequest struct {
	*tchttp.BaseRequest

	// cur or all 白名单时候用

	AddWhiteType *string `json:"AddWhiteType,omitempty" name:"AddWhiteType"`
	// 事件id ，提交规则后会讲此事件变成已加白状态

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 规则id  不填或者0表示新增

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则(默认否) 0：否 ，1：是，全局是Uuids 可为空

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 风险等级 0：无， 1: 高危， 2:中危， 3: 低危

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则内容

	Rules []*FileTamperRule `json:"Rules,omitempty" name:"Rules"`
	// 启用状态 0: 启用 1: 已关闭

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 影响的主机uuid集合

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *ModifyFileTamperRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFileTamperRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScreenRegionInfo struct {

	// 地域标志，如 ap-guangzhou，ap-shanghai，ap-beijing

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域代码，如 gz，sh，bj

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域中文名，如华南地区（广州），华东地区（上海金融），华北地区（北京）

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域英文名

	RegionNameEn *string `json:"RegionNameEn,omitempty" name:"RegionNameEn"`
}

type StopBaselineDetectRequest struct {
	*tchttp.BaseRequest

	// 取消任务ID集合

	TaskIds []*int64 `json:"TaskIds,omitempty" name:"TaskIds"`
}

func (r *StopBaselineDetectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopBaselineDetectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMachineTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMachineTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMachineTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列

	By *string `json:"By,omitempty" name:"By"`
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	// <li>HostName - string - 是否必填：否 - 主机名</li>
	// <li>HostIp - string - 是否必填：否 - 主机IP</li>
	// <li>ItemId - String - 是否必填：否 - 检测项Id</li>
	// <li>ItemName - String - 是否必填：否 - 项名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态[0:未通过|3:通过|5:检测中]</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开始时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineItemListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellRulesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeReverseShellRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExpertServiceTimeRequest struct {
	*tchttp.BaseRequest

	// 订单id

	OrderId *uint64 `json:"OrderId,omitempty" name:"OrderId"`
	// 机器quuid

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *ModifyExpertServiceTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExpertServiceTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBashEventsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBashEventsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBashEventsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagMachinesRequest struct {
	*tchttp.BaseRequest

	// 标签ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTagMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageVul struct {

	// 漏洞id

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// CVSS V3描述

	CVSSV3Desc *string `json:"CVSSV3Desc,omitempty" name:"CVSSV3Desc"`
	// CVSS V3分数

	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 分类2

	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
	// 组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 防御方案

	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`
	// 描述

	Des *string `json:"Des,omitempty" name:"Des"`
	// 风险等级

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 解决方案

	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`
	// 引用

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 提交时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type VulDefencePluginStatus struct {

	// 主机别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 插件状态：0 正常，1 异常

	Exception *int64 `json:"Exception,omitempty" name:"Exception"`
	// 最后跟新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type WarningObject struct {

	// 开始时间，格式: HH:mm

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 漏洞等级控制位二进制，每一位对应页面漏洞等级的开启关闭：低中高（0:关闭；1：开启），例如：101 → 同时勾选低+高；01→(登录审计)疑似不告警，高危告警

	ControlBits *string `json:"ControlBits,omitempty" name:"ControlBits"`
	// 1: 关闭告警 0: 开启告警

	DisablePhoneWarning *uint64 `json:"DisablePhoneWarning,omitempty" name:"DisablePhoneWarning"`
	// 结束时间，格式: HH:mm

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 告警主机范围类型，0:全部主机，1:按所属项目选，2:按腾讯云标签选，3:按主机安全标签选，4:自选主机

	HostRange *int64 `json:"HostRange,omitempty" name:"HostRange"`
	// 事件告警类型；1：离线，2：木马，3：异常登录，4：爆破，5：漏洞（已拆分为9-12四种类型）6：高位命令，7：反弹sell，8：本地提权，9：系统组件漏洞，10：web应用漏洞，11：应急漏洞，12：安全基线，14：恶意请求，15: 网络攻击，16：Windows系统漏洞，17：Linux软件漏洞

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type DescribeAllMachinesRequest struct {
	*tchttp.BaseRequest

	// 机器所属地域。
	// 如：ap-guangzhou，ap-shanghai

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

func (r *DescribeAllMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulStoreListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 免费搜索次数

		FreeSearchTimes *uint64 `json:"FreeSearchTimes,omitempty" name:"FreeSearchTimes"`
		// 漏洞信息

		List []*VulStoreListInfo `json:"List,omitempty" name:"List"`
		// 今日剩余搜索此时

		Remaining *uint64 `json:"Remaining,omitempty" name:"Remaining"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulStoreListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulStoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBruteAttacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBruteAttacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProtectMachineInfo struct {

	// 开通时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 到期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 机器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 机器名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

type ExportAssetWebFrameListRequest struct {
	*tchttp.BaseRequest

	// 可选排序：[FirstTime|JarCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name - String - 是否必填：否 - 框架名</li>
	// <li>Lang - String - 是否必填：否 - 框架语言:java/python</li>
	// <li>Type - String - 是否必填：否 - 服务类型：
	// 0：全部
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetWebFrameListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebFrameListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionQuickResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVersionQuickResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenProVersionQuickResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUndoVulCountsRequest struct {
	*tchttp.BaseRequest

	// 是否应急漏洞筛选, 是 : yes

	IfEmergency *string `json:"IfEmergency,omitempty" name:"IfEmergency"`
	// 漏洞分类，1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞

	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`
}

func (r *DescribeUndoVulCountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUndoVulCountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrustMaliciousRequestResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TrustMaliciousRequestResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TrustMaliciousRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCanNotSeparateMachineRequest struct {
	*tchttp.BaseRequest

	// 排除的事件id,当操作全部事件时，需要排除这次id

	ExcludeId []*uint64 `json:"ExcludeId,omitempty" name:"ExcludeId"`
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>FilePath - String - 是否必填：否 - 路径筛选</li>
	// <li>VirusName - String - 是否必填：否 - 描述筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选-开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选-结束时间</li>
	// <li>Status - String - 是否必填：否 - 状态筛选 4待处理,5信任,6已隔离,10隔离中,11恢复隔离中</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要修改的事件id 数组，支持批量

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 是否更新全部，即是否对所有的事件进行操作，当ids 不为空时，此参数无效

	UpdateAll *bool `json:"UpdateAll,omitempty" name:"UpdateAll"`
}

func (r *DescribeCanNotSeparateMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCanNotSeparateMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityBroadcastsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*Broadcasts `json:"List,omitempty" name:"List"`
		// 总共多少条

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityBroadcastsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityBroadcastsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseBindScheduleRequest struct {
	*tchttp.BaseRequest

	// 过滤参数
	// Status 绑定进度状态 0 进行中 1 已完成 2 失败

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10.

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeLicenseBindScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseBindScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsEventInfoRequest struct {
	*tchttp.BaseRequest

	// 恶意请求事件Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeRiskDnsEventInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsEventInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulCveIdInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 详情列表

		List []*VulInfoByCveId `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulCveIdInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulCveIdInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest

	// 是否开启木马自动隔离(0:关闭 1:开启)

	MalwareAutoIsolate *int64 `json:"MalwareAutoIsolate,omitempty" name:"MalwareAutoIsolate"`
	// 设置木马定时扫描间隔(天数)

	MalwareDiskScanInterDay *int64 `json:"MalwareDiskScanInterDay,omitempty" name:"MalwareDiskScanInterDay"`
	// 设置木马定时扫描时间(0-23整点时间)

	MalwareDiskScanTime *int64 `json:"MalwareDiskScanTime,omitempty" name:"MalwareDiskScanTime"`
	// 是否开启木马定时扫描(0:关闭 1:开启)

	MalwareEnableDiskScan *int64 `json:"MalwareEnableDiskScan,omitempty" name:"MalwareEnableDiskScan"`
	// 设置木马扫描全盘扫描忽略路径(linux), 换行符分割多个

	MalwareIgnoreWhiteListLinux *string `json:"MalwareIgnoreWhiteListLinux,omitempty" name:"MalwareIgnoreWhiteListLinux"`
	// 设置木马扫描全盘扫描忽略路径(windows), 换行符分割多个

	MalwareIgnoreWhiteListWindows *string `json:"MalwareIgnoreWhiteListWindows,omitempty" name:"MalwareIgnoreWhiteListWindows"`
	// 设置木马扫描实时监控忽略路径(linux), 换行符分割多个

	MalwareMonitorIgnoreWhiteListLinux *string `json:"MalwareMonitorIgnoreWhiteListLinux,omitempty" name:"MalwareMonitorIgnoreWhiteListLinux"`
	// 设置木马扫描实时监控忽略路径(windows), 换行符分割多个

	MalwareMonitorIgnoreWhiteListWindows *string `json:"MalwareMonitorIgnoreWhiteListWindows,omitempty" name:"MalwareMonitorIgnoreWhiteListWindows"`
	// 是否开启漏洞定时扫描(0:关闭 1:开启)

	VulEnableVulnerDetect *int64 `json:"VulEnableVulnerDetect,omitempty" name:"VulEnableVulnerDetect"`
	// 设置漏洞定时扫描间隔(天数)

	VulVulnerDetectDay *int64 `json:"VulVulnerDetectDay,omitempty" name:"VulVulnerDetectDay"`
	// 设置漏洞定时扫描时间(0-23整点时间)

	VulVulnerDetectTime *int64 `json:"VulVulnerDetectTime,omitempty" name:"VulVulnerDetectTime"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineDownloadListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [StartTime|EndTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>Status - int - 是否必填：否 - 0:导出中 1:已完成</li>
	// <li>StartTime - string - 是否必填：否 - 开始时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	// <li>TaskName - string - 是否必填：否 - 任务名称</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineDownloadListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineDownloadListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineFixListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [CreateTime|MoifyTime|FixTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>ItemName- string - 是否必填：否 - 项名称</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineFixListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineFixListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenHostInvasionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基线事件列表

		Baseline []*ScreenBaselineInfo `json:"Baseline,omitempty" name:"Baseline"`
		// 网络攻击事件列表

		DefendAttackLog []*ScreenDefendAttackLog `json:"DefendAttackLog,omitempty" name:"DefendAttackLog"`
		// 入侵检测事件列表

		InvasionEvents []*ScreenInvasion `json:"InvasionEvents,omitempty" name:"InvasionEvents"`
		// 漏洞事件列表

		Vul []*ScreenVulInfo `json:"Vul,omitempty" name:"Vul"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenHostInvasionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenHostInvasionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetLoadInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetLoadInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetLoadInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareRiskOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMalwareRiskOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareRiskOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineBasicInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基线基础信息列表

		BaselineBasicInfoList []*BaselineBasicInfo `json:"BaselineBasicInfoList,omitempty" name:"BaselineBasicInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineBasicInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetAttackWhiteListRequest struct {
	*tchttp.BaseRequest

	// 是否加白所有符合该规则的告警 ，1:处理,0:不处理

	DealOldEvents *uint64 `json:"DealOldEvents,omitempty" name:"DealOldEvents"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件id

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// quuid 列表

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// 是否全部主机； 0否，1是。

	Scope *uint64 `json:"Scope,omitempty" name:"Scope"`
	// 来源IP 单IP:1.1.1.1  IP范围:1.1.1.1-1.1.2.1  IP范围：1.1.1.0/24

	SrcIp []*string `json:"SrcIp,omitempty" name:"SrcIp"`
}

func (r *CreateNetAttackWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetAttackWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrivilegeEscalationProcess struct {

	// 执行命令

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 发生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 进程路径

	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`
	// 主机内网IP

	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
	// 数据ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 机器名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 父进程用户组

	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`
	// 父进程名

	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`
	// 父进程路径

	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`
	// 父进程用户名

	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`
	// 进程文件权限

	ProcFilePrivilege *string `json:"ProcFilePrivilege,omitempty" name:"ProcFilePrivilege"`
	// 进程树

	ProcTree *string `json:"ProcTree,omitempty" name:"ProcTree"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 主机ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 处理状态：0-待处理 2-白名单 3-已处理 4-已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 用户组

	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 云镜ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type MisAlarmNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MisAlarmNonlocalLoginPlacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MisAlarmNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBaselineStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBaselineStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBaselineStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanTaskDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险事件个数

		RiskEventCount *uint64 `json:"RiskEventCount,omitempty" name:"RiskEventCount"`
		// 发现风险机器数

		RiskMachineCount *uint64 `json:"RiskMachineCount,omitempty" name:"RiskMachineCount"`
		// 扫描开始时间

		ScanBeginTime *string `json:"ScanBeginTime,omitempty" name:"ScanBeginTime"`
		// 扫描内容

		ScanContent []*string `json:"ScanContent,omitempty" name:"ScanContent"`
		// 扫描结束时间

		ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`
		// 扫描剩余时间

		ScanLeftTime *uint64 `json:"ScanLeftTime,omitempty" name:"ScanLeftTime"`
		// 扫描机器总数

		ScanMachineCount *uint64 `json:"ScanMachineCount,omitempty" name:"ScanMachineCount"`
		// 扫描进度

		ScanProgress *uint64 `json:"ScanProgress,omitempty" name:"ScanProgress"`
		// 扫描任务信息列表

		ScanTaskDetailList []*ScanTaskDetails `json:"ScanTaskDetailList,omitempty" name:"ScanTaskDetailList"`
		// 检测时间

		ScanTime *uint64 `json:"ScanTime,omitempty" name:"ScanTime"`
		// 任务是否全部正在被停止 ture是

		StoppingAll *bool `json:"StoppingAll,omitempty" name:"StoppingAll"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 0一键检测 1定时检测

		Type *uint64 `json:"Type,omitempty" name:"Type"`
		// 扫描出漏洞个数

		VulCount *uint64 `json:"VulCount,omitempty" name:"VulCount"`
		// 漏洞信息

		VulInfo []*VulDetailInfo `json:"VulInfo,omitempty" name:"VulInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanTaskDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanTaskDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebFrameCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各Web框架数量

		WebFrames []*AssetKeyVal `json:"WebFrames,omitempty" name:"WebFrames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebFrameCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebFrameCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetLoadInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统负载

		CpuLoad *AssetLoadSummary `json:"CpuLoad,omitempty" name:"CpuLoad"`
		// 硬盘使用率

		DiskLoad *AssetLoadSummary `json:"DiskLoad,omitempty" name:"DiskLoad"`
		// 内存使用率

		MemLoad *AssetLoadSummary `json:"MemLoad,omitempty" name:"MemLoad"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetLoadInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetLoadInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareInfoRequest struct {
	*tchttp.BaseRequest

	// 唯一ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeMalwareInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskProcessEvent struct {

	// 木马检测平台 [1:云查杀引擎|2:TAV|3:binaryAi|4:异常行为|5:威胁情报]

	CheckPlatform []*string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
	// 执行命令

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 最近检测时间

	DetectTime *string `json:"DetectTime,omitempty" name:"DetectTime"`
	// 事件ID

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 处理状态[0待处理;1已处理;2查杀中;3已查杀;4已退出;5忽略]

	HandleStatus *int64 `json:"HandleStatus,omitempty" name:"HandleStatus"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机在线状态

	OnlineStatus *int64 `json:"OnlineStatus,omitempty" name:"OnlineStatus"`
	// 进程ID

	ProcessId *int64 `json:"ProcessId,omitempty" name:"ProcessId"`
	// 参考链接

	ReferenceLink *string `json:"ReferenceLink,omitempty" name:"ReferenceLink"`
	// 进程启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 建议方案

	SuggestSolution *string `json:"SuggestSolution,omitempty" name:"SuggestSolution"`
	// 威胁描述

	ThreatDesc *string `json:"ThreatDesc,omitempty" name:"ThreatDesc"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 病毒标签

	VirusTags []*string `json:"VirusTags,omitempty" name:"VirusTags"`
	// 外网IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}

type DeleteLicenseRecordAllRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteLicenseRecordAllRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLicenseRecordAllRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineInfo struct {

	// 基线风险项

	BaselineFailCount *uint64 `json:"BaselineFailCount,omitempty" name:"BaselineFailCount"`
	// 基线id

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 影响服务器数量

	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`
	// 最后检测时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 危害等级：1-低危；2-中危；3-高危；4-严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 检测中状态: 5

	MaxStatus *uint64 `json:"MaxStatus,omitempty" name:"MaxStatus"`
	// 基线名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 检测项数量

	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// 通过状态:0:未通过,1:已通过

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type RansomDefenseStrategyDetail struct {

	// 备份模式： 0按周，1按天

	BackupType *uint64 `json:"BackupType,omitempty" name:"BackupType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 策略备注

	Description *string `json:"Description,omitempty" name:"Description"`
	// 策略关联事件数

	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
	// 包含目录，用;分隔

	ExcludeDir *string `json:"ExcludeDir,omitempty" name:"ExcludeDir"`
	// 备份执行时间点(0-23): 11:00;12:00

	Hour *string `json:"Hour,omitempty" name:"Hour"`
	// 策略id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 包含目录，用;分隔

	IncludeDir *string `json:"IncludeDir,omitempty" name:"IncludeDir"`
	// 是否对所有主机生效

	IsAll *uint64 `json:"IsAll,omitempty" name:"IsAll"`
	// 绑定机器数

	MachineCount *uint64 `json:"MachineCount,omitempty" name:"MachineCount"`
	// 最近修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 保存天数，0永久保存

	SaveDay *uint64 `json:"SaveDay,omitempty" name:"SaveDay"`
	// 开启状态：0关闭，1开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 操作uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 备份星期天数（1-7）：1;2;3;4

	Weekday *string `json:"Weekday,omitempty" name:"Weekday"`
}

type ExportMaliciousRequestsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportMaliciousRequestsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportMaliciousRequestsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetNetworkCardInfo struct {

	// DNS服务器

	DnsServer *string `json:"DnsServer,omitempty" name:"DnsServer"`
	// 网关

	GateWay *string `json:"GateWay,omitempty" name:"GateWay"`
	// Ipv4对应IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Ipv6对应IP

	Ipv6 *string `json:"Ipv6,omitempty" name:"Ipv6"`
	// MAC地址

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// 网卡名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type MachineExtraInfo struct {

	// 主机名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 网络名，vpc网络情况下会返回vpc_id

	NetworkName *string `json:"NetworkName,omitempty" name:"NetworkName"`
	// 网络类型，1:vpc网络 2:基础网络 3:非腾讯云网络

	NetworkType *int64 `json:"NetworkType,omitempty" name:"NetworkType"`
	// 内网IP

	PrivateIP *string `json:"PrivateIP,omitempty" name:"PrivateIP"`
	// 公网IP

	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`
}

type ProtectEventLists struct {

	// 发现时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 事件地址

	EventDir *string `json:"EventDir,omitempty" name:"EventDir"`
	// 事件状态 1 已恢复 0 未恢复

	EventStatus *uint64 `json:"EventStatus,omitempty" name:"EventStatus"`
	// 事件类型 0-内容被修改恢复；1-权限被修改恢复；2-归属被修改恢复；3-被删除恢复；4-新增删除

	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`
	// 文件类型 0-常规文件；1-目录；2-软链

	FileType *uint64 `json:"FileType,omitempty" name:"FileType"`
	// 服务器ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 服务器名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 恢复时间

	RestoreTime *string `json:"RestoreTime,omitempty" name:"RestoreTime"`
}

type DescribeWebHookRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则内容

		Data *WebHookRuleDetail `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebHookRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebHookRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineHostDetectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBaselineHostDetectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineHostDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgnoreRuleEffectHostListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>AliasName- String- 主机别名</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 检测项id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 主机标签名

	TagNames []*string `json:"TagNames,omitempty" name:"TagNames"`
}

func (r *DescribeIgnoreRuleEffectHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgnoreRuleEffectHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulDefenceOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebPageProtectSwitchRequest struct {
	*tchttp.BaseRequest

	// 需要操作开关的网站 最大100条

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
	// 1 开启 0 关闭 SwitchType 为 1 | 2 必填;

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 开关类型 1 防护开关  2 自动恢复开关 3 移除防护目录

	SwitchType *uint64 `json:"SwitchType,omitempty" name:"SwitchType"`
}

func (r *ModifyWebPageProtectSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebPageProtectSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetMachineBaseInfo struct {

	// CPU信息

	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`
	// Cpu使用率百分比

	CpuLoad *string `json:"CpuLoad,omitempty" name:"CpuLoad"`
	// Cpu数量

	CpuSize *uint64 `json:"CpuSize,omitempty" name:"CpuSize"`
	// 硬盘使用率百分比

	DiskLoad *string `json:"DiskLoad,omitempty" name:"DiskLoad"`
	// 硬盘容量：单位G

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 服务器内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 服务器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 内存使用率百分比

	MemLoad *string `json:"MemLoad,omitempty" name:"MemLoad"`
	// 内存容量：单位G

	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`
	// 操作系统名称

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 分区数

	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`
	// 业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 服务器uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportInfo struct {

	// 日志导出路径

	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`
	// 日志导出数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 日志导出创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 日志导出结束时间，uinx毫秒时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 日志导出任务ID

	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
	// 日志导出文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 日志文件大小

	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`
	// 日志导出格式

	Format *string `json:"Format,omitempty" name:"Format"`
	// 日志导出时间排序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 日志导出查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 日志导出起始时间，uinx毫秒时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 日志下载状态。Processing:导出正在进行中，Complete:导出完成，Failed:导出失败，Expired:日志导出已过期（三天有效期）。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DeleteMaliciousRequestWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaliciousRequestWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaliciousRequestWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DefaultStrategyInfo struct {

	// 策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略名

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
}

type PriceDetail struct {

	// 订单分类

	GoodsCategoryId *uint64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`
	// 订单详情参数

	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 订单数

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 费用模式 0 后付费,1 预付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 地域

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 可用区

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type ModifyRansomDefenseEventsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRansomDefenseEventsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRansomDefenseEventsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RansomDefenseEvent struct {

	// 被篡改文件路径

	BaitFilePath *string `json:"BaitFilePath,omitempty" name:"BaitFilePath"`
	// 事件发送时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 恶意文件md5

	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`
	// 恶意文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 恶意文件大小

	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`
	// 主机外网ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// cvm 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 事件修改事件

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 恶意进程id

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 恶意进程参数

	PidParam *string `json:"PidParam,omitempty" name:"PidParam"`
	// 进程启动时间

	ProcessStartTime *string `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 进程树 base64 json

	PsTree *string `json:"PsTree,omitempty" name:"PsTree"`
	// cvm uuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机拥有快照备份数

	SnapshotNum *uint64 `json:"SnapshotNum,omitempty" name:"SnapshotNum"`
	// 事件状态 0待处理，1已处理，2已信任，3处理中，4已恢复备份

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 事件类型：0加密勒索，1文件篡改

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 主机内网ip

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}

type ScreenAttackHotspot struct {

	// 时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 受害者IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 事件名

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 攻击者IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
}

type DescribeBaselineWeakPasswordListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineWeakPassword `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineWeakPasswordListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineWeakPasswordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBaselineRuleIgnoreRequest struct {
	*tchttp.BaseRequest

	// 资产类型[0:所有专业版旗舰版|1:id|2:ip]

	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
	// <li>ItemName - string - 是否必填：否 - 项名称</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 主机Id

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 主机Ip

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 关联项

	ItemIds []*int64 `json:"ItemIds,omitempty" name:"ItemIds"`
	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 是否全选过滤

	SelectAll *int64 `json:"SelectAll,omitempty" name:"SelectAll"`
}

func (r *ModifyBaselineRuleIgnoreRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBaselineRuleIgnoreRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMalwareWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMalwareWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMalwareWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebAppListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetWebAppListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineSimple struct {

	// 云标签信息

	CloudTags []*Tags `json:"CloudTags,omitempty" name:"CloudTags"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例状态 TERMINATED_PRO_VERSION 已销毁

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 是否是专业版。
	// <li>true： 是</li>
	// <li>false：否</li>

	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 内核版本

	KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`
	// 授权订单对象

	LicenseOrder *LicenseOrder `json:"LicenseOrder,omitempty" name:"LicenseOrder"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机系统。

	MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`
	// 机器所属专区类型 CVM 云服务器, BM 黑石, ECM 边缘计算, LH 轻量应用服务器 ,Other 混合云专区

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 主机外网IP。

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 主机状态。
	// <li>POSTPAY: 表示后付费，即按量计费  </li>
	// <li>PREPAY: 表示预付费，即包年包月</li>

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 防护版本 BASIC_VERSION 基础版， PRO_VERSION 专业版，Flagship 旗舰版，GENERAL_DISCOUNT 普惠版.

	ProtectType *string `json:"ProtectType,omitempty" name:"ProtectType"`
	// CVM或BM机器唯一Uuid。

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 地域信息

	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`
	// 标签信息

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 云镜客户端唯一Uuid，若客户端长时间不在线将返回空字符。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type CreateLicenseOrderRequest struct {
	*tchttp.BaseRequest

	// 该字段作废

	AutoProtectOpenConfig *string `json:"AutoProtectOpenConfig,omitempty" name:"AutoProtectOpenConfig"`
	// 是否自动续费, 默认不自动续费.
	// 该参数仅包年包月生效

	AutoRenewFlag *bool `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 业务自定义参数

	BillingDefineParams *BillingDefineParams `json:"BillingDefineParams,omitempty" name:"BillingDefineParams"`
	// 授权数量 , 需要购买的数量.
	// 默认为1

	LicenseNum *uint64 `json:"LicenseNum,omitempty" name:"LicenseNum"`
	// 授权类型, 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月
	// 默认为0

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 变配参数

	ModifyConfig *OrderModifyObject `json:"ModifyConfig,omitempty" name:"ModifyConfig"`
	// 项目ID .
	// 默认为0

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 购买订单地域,这里仅支持 1 广州,9 新加坡. 推荐选择广州. 新加坡地域为白名单用户购买.
	// 默认为1

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 标签数组, 空则表示不需要绑定标签

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
	// 购买时间长度,默认1 , 可选值为1,2,3,4,5,6,7,8,9,10,11,12,24,36
	// 该参数仅包年包月生效

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

func (r *CreateLicenseOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLicenseOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取事件列表

		EventList []*VulEventList `json:"EventList,omitempty" name:"EventList"`
		// 获取事件总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineDefenseCntRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Uuids- String - 是否必填：否 - 主机uuid</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMachineDefenseCntRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineDefenseCntRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenMachineRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表详情

		List []*ScreenRegionInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenMachineRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenMachineRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldValueRatioInfo struct {

	// 个数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 比例

	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
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

type DescribeAssetUserCountRequest struct {
	*tchttp.BaseRequest

	// 搜索条件：搜索名称中包含name的所有账号列表

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetUserCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogDeliveryKafkaOptionsRequest struct {
	*tchttp.BaseRequest

	// kafka实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeLogDeliveryKafkaOptionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDeliveryKafkaOptionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JavaMemShellPluginSetting struct {

	// 服务器名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 插件是否存在异常 0: 正常 1: 异常

	Exception *uint64 `json:"Exception,omitempty" name:"Exception"`
	// 服务器ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// javashell插件开关 0: 关闭 1: 开启

	JavaShellStatus *uint64 `json:"JavaShellStatus,omitempty" name:"JavaShellStatus"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 容器quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetRecentMachineInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 在线数量列表

		LiveList []*AssetKeyVal `json:"LiveList,omitempty" name:"LiveList"`
		// 离线数量列表

		OfflineList []*AssetKeyVal `json:"OfflineList,omitempty" name:"OfflineList"`
		// 风险数量列表

		RiskList []*AssetKeyVal `json:"RiskList,omitempty" name:"RiskList"`
		// 总数量列表

		TotalList []*AssetKeyVal `json:"TotalList,omitempty" name:"TotalList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetRecentMachineInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetRecentMachineInfoResponse) FromJsonString(s string) error {
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

type DescribeAssetWebLocationInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 站点信息

		WebLocation *AssetWebLocationInfo `json:"WebLocation,omitempty" name:"WebLocation"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebLocationInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImpactedHost struct {

	// 漏洞描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 漏洞ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否为专业版。

	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 最后检测时间。

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 云镜客户端唯一标识UUID。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 漏洞种类ID。

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞状态。
	// <li>UN_OPERATED ：待处理</li>
	// <li>SCANING : 扫描中</li>
	// <li>FIXED : 已修复</li>

	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`
}

type ProtectStat struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数量

	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type DeleteLogExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountStatisticsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号用户名</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAccountStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineHostRiskTopRequest struct {
	*tchttp.BaseRequest

	// 策略ID

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeBaselineHostRiskTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineHostRiskTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineLicenseDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 授权信息

		MachineLicense []*MachineLicenseDetail `json:"MachineLicense,omitempty" name:"MachineLicense"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineLicenseDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineLicenseDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulCountByDatesRequest struct {
	*tchttp.BaseRequest

	// 是否为应急漏洞筛选  是: yes

	IfEmergency *string `json:"IfEmergency,omitempty" name:"IfEmergency"`
	// 需要查询最近几天的数据，需要都 -1后传入

	LastDays []*uint64 `json:"LastDays,omitempty" name:"LastDays"`
	// 漏洞的分类: 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞

	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`
}

func (r *DescribeVulCountByDatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulCountByDatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulInfoRequest struct {
	*tchttp.BaseRequest

	// 漏洞种类ID。

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *DescribeVulInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportJavaMemShellPluginsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportJavaMemShellPluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportJavaMemShellPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBaselineStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateBaselineStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBaselineStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileTamperEvent struct {

	// 发生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 危害描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件产生次数

	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
	// 进程名

	ExeMd5 *string `json:"ExeMd5,omitempty" name:"ExeMd5"`
	// 进程名称

	ExeName *string `json:"ExeName,omitempty" name:"ExeName"`
	// 进程权限

	ExePermission *string `json:"ExePermission,omitempty" name:"ExePermission"`
	// 进程pid

	ExePid *uint64 `json:"ExePid,omitempty" name:"ExePid"`
	// 进程文件大小

	ExeSize *uint64 `json:"ExeSize,omitempty" name:"ExeSize"`
	// 进程执行时长

	ExeTime *uint64 `json:"ExeTime,omitempty" name:"ExeTime"`
	// 机器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 机器名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 风险等级 0：无， 1: 高危， 2:中危， 3: 低危

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	//  主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机在线信息 ONLINE、OFFLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 最近发生时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 进程参数

	ProcessArgv *string `json:"ProcessArgv,omitempty" name:"ProcessArgv"`
	// 进程路径

	ProcessExe *string `json:"ProcessExe,omitempty" name:"ProcessExe"`
	// 事件详情: json格式

	Pstree *string `json:"Pstree,omitempty" name:"Pstree"`
	// cvm id

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 参考链接

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 规则类型 0系统规则 1自定义规则

	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`
	// 规则id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略 4-已手动处理

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 修护建议

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 目标文件路径

	Target *string `json:"Target,omitempty" name:"Target"`
	// 目标文件创建时间

	TargetCreatTime *string `json:"TargetCreatTime,omitempty" name:"TargetCreatTime"`
	// 目标文件更新时间

	TargetModifyTime *string `json:"TargetModifyTime,omitempty" name:"TargetModifyTime"`
	// 文件名称

	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`
	// 目标文件权限

	TargetPermission *string `json:"TargetPermission,omitempty" name:"TargetPermission"`
	// 目标文件大小

	TargetSize *uint64 `json:"TargetSize,omitempty" name:"TargetSize"`
	// 事件类型/动作  0 -- 告警

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 用户组

	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// PENDING：正在生成下载链接，FINISHED：下载链接已生成，ERROR：网络异常等异常情况

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefencePluginExceptionCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulDefencePluginExceptionCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefencePluginExceptionCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductStatusInfo struct {

	// 是否可以申请试用，true可以申请

	CanApplyTrial *bool `json:"CanApplyTrial,omitempty" name:"CanApplyTrial"`
	// 无法试用原因，可试用为空

	CanNotApplyReason *string `json:"CanNotApplyReason,omitempty" name:"CanNotApplyReason"`
	// 防护状态，1未防护，2防护中，3试用中，4已过期

	FWUserStatus *uint64 `json:"FWUserStatus,omitempty" name:"FWUserStatus"`
	// 上次试用结束时间（不存在试用记录则为空）

	LastTrialTime *string `json:"LastTrialTime,omitempty" name:"LastTrialTime"`
}

type BashEvent struct {

	// 执行命令

	BashCmd *string `json:"BashCmd,omitempty" name:"BashCmd"`
	// 发生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 0: bash日志 1: 实时监控(雷霆版)

	DetectBy *uint64 `json:"DetectBy,omitempty" name:"DetectBy"`
	// 进程名称

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 主机内网IP

	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
	// 数据ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 处理时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 进程id

	Pid *string `json:"Pid,omitempty" name:"Pid"`
	// 平台类型

	Platform *uint64 `json:"Platform,omitempty" name:"Platform"`
	// 主机ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 自动生成的正则表达式

	RegexBashCmd *string `json:"RegexBashCmd,omitempty" name:"RegexBashCmd"`
	// 规则类别  0=系统规则，1=用户规则

	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`
	// 规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则等级：1-高 2-中 3-低

	RuleLevel *uint64 `json:"RuleLevel,omitempty" name:"RuleLevel"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 处理状态： 0 = 待处理 1= 已处理, 2 = 已加白， 3 = 已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 执行用户名

	User *string `json:"User,omitempty" name:"User"`
	// 云镜ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type HostInfo struct {

	// Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAttackStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总攻击来源ip数量

		AttackSrcIpCount *uint64 `json:"AttackSrcIpCount,omitempty" name:"AttackSrcIpCount"`
		// 总受攻击资产数量

		AttackedAssetCount *uint64 `json:"AttackedAssetCount,omitempty" name:"AttackedAssetCount"`
		// 总受攻击端口数量

		AttackedPortCount *uint64 `json:"AttackedPortCount,omitempty" name:"AttackedPortCount"`
		// 今日新增攻击来源ip数量

		NewAttackSrcIpCount *uint64 `json:"NewAttackSrcIpCount,omitempty" name:"NewAttackSrcIpCount"`
		// 今日新增受攻击资产数量

		NewAttackedAssetCount *uint64 `json:"NewAttackedAssetCount,omitempty" name:"NewAttackedAssetCount"`
		// 今日新增受攻击端口数量

		NewAttackedPortCount *uint64 `json:"NewAttackedPortCount,omitempty" name:"NewAttackedPortCount"`
		// 总攻击次数

		PendingAttackCount *uint64 `json:"PendingAttackCount,omitempty" name:"PendingAttackCount"`
		// 今日新增攻击次数

		PendingNewAttackCount *uint64 `json:"PendingNewAttackCount,omitempty" name:"PendingNewAttackCount"`
		// 总攻击成功次数

		PendingSuccAttackCount *uint64 `json:"PendingSuccAttackCount,omitempty" name:"PendingSuccAttackCount"`
		// 总尝试攻击次数

		PendingTryAttackCount *uint64 `json:"PendingTryAttackCount,omitempty" name:"PendingTryAttackCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebPageServiceInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWebPageServiceInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebPageServiceInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentInfoRequest struct {
	*tchttp.BaseRequest

	// 组件ID。

	ComponentId *uint64 `json:"ComponentId,omitempty" name:"ComponentId"`
}

func (r *DescribeComponentInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttackSource struct {

	// 攻击溯源节点路径

	Edges []*AttackSourceEdge `json:"Edges,omitempty" name:"Edges"`
	// 请求节点相关事件详情的参数

	EventInfoParam *string `json:"EventInfoParam,omitempty" name:"EventInfoParam"`
	// 攻击溯源节点描述

	Nodes []*AttackSourceNode `json:"Nodes,omitempty" name:"Nodes"`
}

type DeletePrivilegeRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivilegeRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivilegeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPortCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各进程数量

		Ports []*AssetKeyVal `json:"Ports,omitempty" name:"Ports"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetPortCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebLocationInfoRequest struct {
	*tchttp.BaseRequest

	// 站点Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebLocationInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRiskDnsPolicyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRiskDnsPolicyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRiskDnsPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMalwareTimingScanSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMalwareTimingScanSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMalwareTimingScanSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerRelatedDirInfoRequest struct {
	*tchttp.BaseRequest

	// 唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeServerRelatedDirInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerRelatedDirInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞列表数组。

		Vuls []*Vul `json:"Vuls,omitempty" name:"Vuls"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefencePluginDetail struct {

	// 错误日志

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
	// 注入日志

	InjectLog *string `json:"InjectLog,omitempty" name:"InjectLog"`
	// 注入进程主类名

	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`
	// 注入进程Pid

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 插件状态：0: 注入中, 1: 注入成功, 2: 插件超时, 3: 插件退出, 4: 注入失败 5: 软删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type CancelIgnoreVulRequest struct {
	*tchttp.BaseRequest

	// 漏洞事件id串，多个用英文逗号分隔

	EventIds *string `json:"EventIds,omitempty" name:"EventIds"`
}

func (r *CancelIgnoreVulRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelIgnoreVulRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineHostDetectListRequest struct {
	*tchttp.BaseRequest

	// 0:过滤的结果导出；1:全部导出

	ExportAll *int64 `json:"ExportAll,omitempty" name:"ExportAll"`
	// <li>HostTag - string - 是否必填：否 - 主机标签</i>
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>IsPassed - int - 是否必填：否 - 是否通过</li>
	// <li>RiskTier - int - 是否必填：否 - 风险等级</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBaselineHostDetectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineHostDetectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRansomDefenseBackupListRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Status - Int - 是否必填：否 - 通过勒索状态查询：0未勒索，1已勒索 </li>
	// <li>CreateTimeBegin - string - 是否必填：否 - 创建时间开始 </li>
	// <li>CreateTimeEnd - string - 是否必填：否 - 创建时间结束 </li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportRansomDefenseBackupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRansomDefenseBackupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgnoreRuleEffectHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 忽略检测项影响主机列表

		IgnoreRuleEffectHostList []*IgnoreRuleEffectHostInfo `json:"IgnoreRuleEffectHostList,omitempty" name:"IgnoreRuleEffectHostList"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIgnoreRuleEffectHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgnoreRuleEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProtectPathInfo struct {

	// 主机名称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 用户Appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 文件自动恢复开关：0-未开启；1-已启动

	AutoRecover *uint64 `json:"AutoRecover,omitempty" name:"AutoRecover"`
	// 异常，0-无异常，1-超出限制；2-agent离线；3-超时；4-磁盘不足；5-机器已销毁；99-其他

	Exception *uint64 `json:"Exception,omitempty" name:"Exception"`
	// 防护文件类型列表，以;分隔

	FileTypesWhite *string `json:"FileTypesWhite,omitempty" name:"FileTypesWhite"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机操作系统

	HostOS *string `json:"HostOS,omitempty" name:"HostOS"`
	// 防护目录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 启动进度，百分数不含%

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 防护文件个数

	ProtectFileCount *uint64 `json:"ProtectFileCount,omitempty" name:"ProtectFileCount"`
	// 防护目录

	ProtectPath *string `json:"ProtectPath,omitempty" name:"ProtectPath"`
	// 腾迅云CVM uuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 防护状态：0-未开启；1-启动中；2-已启动；3-关闭中；4-已关闭；5-授权到期；6-已删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DescribeAssetCoreModuleListRequest struct {
	*tchttp.BaseRequest

	// 排序依据[Size|FirstTime|ProcessCount|ModuleCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetCoreModuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetCoreModuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMaliciousRequestWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 白名单id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 备注

	Mark *string `json:"Mark,omitempty" name:"Mark"`
}

func (r *ModifyMaliciousRequestWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMaliciousRequestWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filters struct {

	// 是否模糊匹配，前端框架会带上，可以不管

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeAutoQuaraInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAutoQuaraInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoQuaraInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAllJavaMemShellsRequest struct {
	*tchttp.BaseRequest

	// 服务器quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DeleteAllJavaMemShellsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAllJavaMemShellsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostTotalCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各项资源数量
		// system : 资源监控
		// account: 账号
		// port: 端口
		// process: 进程
		// app: 应用软件
		// database:数据库
		// webapp: Web应用
		// webframe: Web框架
		// webservice: Web服务
		// weblocation: Web站点
		// systempackage: 系统安装包
		// jar: jar包
		// initservice:启动服务
		// env: 环境变量
		// coremodule: 内核模块

		Types []*AssetKeyVal `json:"Types,omitempty" name:"Types"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetHostTotalCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostTotalCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopNoticeBanTipsRequest struct {
	*tchttp.BaseRequest
}

func (r *StopNoticeBanTipsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopNoticeBanTipsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BruteAttackList struct {

	// AgentId

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// HostName

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 事件Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 主机IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// Mid

	Mid *string `json:"Mid,omitempty" name:"Mid"`
	// 省份

	Province *string `json:"Province,omitempty" name:"Province"`
	// 来源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 用户名称

	Username *string `json:"Username,omitempty" name:"Username"`
}

type VulInfoList struct {

	// cve编号

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// CVSS评分

	CvssScore *float64 `json:"CvssScore,omitempty" name:"CvssScore"`
	// 已防御的攻击次数

	DefenseAttackCount *uint64 `json:"DefenseAttackCount,omitempty" name:"DefenseAttackCount"`
	// 描述

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// 废弃字段

	DescriptWisteria *string `json:"DescriptWisteria,omitempty" name:"DescriptWisteria"`
	// 首次出现时间

	FirstAppearTime *string `json:"FirstAppearTime,omitempty" name:"FirstAppearTime"`
	// 是否能自动修复且包含能自动修复的主机， 0=否  1=是

	FixSwitch *uint64 `json:"FixSwitch,omitempty" name:"FixSwitch"`
	// 废弃字段

	From *uint64 `json:"From,omitempty" name:"From"`
	// 影响主机数

	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`
	// 漏洞包含的事件id串，多个用“,”分割

	Ids *string `json:"Ids,omitempty" name:"Ids"`
	// 是否支持防御， 0:不支持 1:支持

	IsSupportDefense *uint64 `json:"IsSupportDefense,omitempty" name:"IsSupportDefense"`
	// 漏洞标签 多个逗号分割

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 最后检测时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 漏洞等级 1:低 2:中 3:高 4:严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 漏洞名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 废弃字段

	NameWisteria *string `json:"NameWisteria,omitempty" name:"NameWisteria"`
	// 漏洞披露事件

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 废弃字段

	PublishTimeWisteria *string `json:"PublishTimeWisteria,omitempty" name:"PublishTimeWisteria"`
	// 0: 待处理 1:忽略  3:已修复  5:检测中 6:修复中  8:修复失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 聚合后事件状态串

	StatusStr *string `json:"StatusStr,omitempty" name:"StatusStr"`
	// 最后扫描任务的id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 漏洞类别 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞

	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

type DescribeLicenseWhiteConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 旗舰版 配置信息

		FlagShip *VersionWhiteConfig `json:"FlagShip,omitempty" name:"FlagShip"`
		// 普惠版 配置信息

		PrattWhitney *VersionWhiteConfig `json:"PrattWhitney,omitempty" name:"PrattWhitney"`
		// 专业版 配置信息

		Professional *VersionWhiteConfig `json:"Professional,omitempty" name:"Professional"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseWhiteConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseWhiteConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanBaselineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正在扫描中的主机Quuid

		ScanningQuuids []*string `json:"ScanningQuuids,omitempty" name:"ScanningQuuids"`
		// 任务下发成功返回的TaskId

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanBaselineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanBaselineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPlanTaskListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	// <li>Status- int - 是否必填：否 - 默认启用状态：0未启用， 1启用 </li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetPlanTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPlanTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineRuleIgnoreListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*BaselineRule `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineRuleIgnoreListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetJarDetail struct {

	// 服务器IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 服务器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// Jar包Md5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 引用进程列表

	Process []*AssetAppProcessInfo `json:"Process,omitempty" name:"Process"`
	// 是否可执行：0未知，1是，2否

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 类型：1应用程序，2系统类库，3Web服务自带库，8:其他，

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type Data struct {

	// 分母

	Den *uint64 `json:"Den,omitempty" name:"Den"`
	// 分子

	Mol *uint64 `json:"Mol,omitempty" name:"Mol"`
	// 时间，如果不返回，将使用请求当前时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

type LicenseBindDetail struct {

	// 云镜客户端状态,OFFLINE 离线,ONLINE 在线,UNINSTALL 未安装

	AgentStatus *string `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// 是否允许换绑,false 不允许换绑

	IsSwitchBind *bool `json:"IsSwitchBind,omitempty" name:"IsSwitchBind"`
	// 是否允许解绑,false 不允许解绑

	IsUnBind *bool `json:"IsUnBind,omitempty" name:"IsUnBind"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 机器内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 机器别名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 机器公网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 云服务器UUID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 标签信息

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 云镜客户端UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetSystemPackageListRequest struct {
	*tchttp.BaseRequest

	// 排序方式可选：[FistTime|InstallTime:安装时间]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name - String - 是否必填：否 - 包 名</li>
	// <li>StartTime - String - 是否必填：否 - 安装开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 安装开始时间</li>
	// <li>Type - int - 是否必填：否 - 安装包类型：
	// 1:rmp
	// 2:dpkg
	// 3:java
	// 4:system</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc-升序 或 desc-降序。默认：desc-降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetSystemPackageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSystemPackageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityDynamicsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSecurityDynamicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityDynamicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 账号列表

		Users []*AssetUserBaseInfo `json:"Users,omitempty" name:"Users"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetUserListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenProtectionStatRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScreenProtectionStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenProtectionStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopNoticeBanTipsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopNoticeBanTipsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopNoticeBanTipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileTamperRuleInfo struct {

	// 白名单处理 当前:cur 所有符合历史:all

	AddWhiteType *string `json:"AddWhiteType,omitempty" name:"AddWhiteType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// FileAction

	FileAction *string `json:"FileAction,omitempty" name:"FileAction"`
	// 影响主机数

	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`
	// 规则id，系统的规则时为0。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否是全局的 0：否 ，1：是

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 风险等级 0：无， 1: 高危， 2:中危， 3: 低危

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// ReadRuleCount

	ReadRuleCount *uint64 `json:"ReadRuleCount,omitempty" name:"ReadRuleCount"`
	// ReadWriteRuleCount

	ReadWriteRuleCount *uint64 `json:"ReadWriteRuleCount,omitempty" name:"ReadWriteRuleCount"`
	// 规则类型 0 ：系统规则  1：用户规则

	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`
	// 状态 0: 启用 1: 已关闭

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// WriteRuleCount

	WriteRuleCount *uint64 `json:"WriteRuleCount,omitempty" name:"WriteRuleCount"`
}

type ImageVirus struct {

	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 风险等级

	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 标签

	Tags *string `json:"Tags,omitempty" name:"Tags"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
}

type DescribeLogKafkaDeliverInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogKafkaDeliverInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogKafkaDeliverInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebHookRuleRequest struct {
	*tchttp.BaseRequest

	// 规则Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeWebHookRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebHookRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditPrivilegeRuleRequest struct {
	*tchttp.BaseRequest

	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)

	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
	// 规则ID(新增时请留空)

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则(默认否)

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 是否S权限进程

	SMode *uint64 `json:"SMode,omitempty" name:"SMode"`
	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *EditPrivilegeRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditPrivilegeRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetTotalCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetTotalCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTotalCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDetectionReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出检测报告的任务Id（不同于入参的漏洞扫描任务id）

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulDetectionReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDetectionReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogKafkaDeliverTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogKafkaDeliverTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogKafkaDeliverTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwaresRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 木马状态（UN_OPERATED: 未处理 | SEGREGATED: 已隔离|TRUSTED：信任）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMalwaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwaresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulFixStatusRequest struct {
	*tchttp.BaseRequest

	// 任务id 传VulId可以不用传FixId

	FixId *uint64 `json:"FixId,omitempty" name:"FixId"`
	// 主机quuid 和VulId 组合可查 某主机最近一次修复任务详情

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 漏洞id 传FixId可以不用传VulId

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *DescribeVulFixStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulFixStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsInfoRequest struct {
	*tchttp.BaseRequest

	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeBashEventsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashEventsInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInquireRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 授权ID

	LicenseId *uint64 `json:"LicenseId,omitempty" name:"LicenseId"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 设置自动开通状态。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyInquireRenewFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInquireRenewFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebAppPluginInfo struct {

	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 链接

	Link *string `json:"Link,omitempty" name:"Link"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeWebPageGeneralizeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWebPageGeneralizeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebPageGeneralizeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 帐号统计列表。

		AccountStatistics []*AccountStatistics `json:"AccountStatistics,omitempty" name:"AccountStatistics"`
		// 帐号统计列表记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebFrameListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 列表

		WebFrames []*AssetWebFrameBaseInfo `json:"WebFrames,omitempty" name:"WebFrames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebFrameListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebFrameListResponse) FromJsonString(s string) error {
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

type DescribeEventListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>VulId- int - 是否必填：否 - 事件id</li>
	// <li>VulCategory- int - 是否必填：否 - 漏洞分类筛选 1: web应用漏洞 2:系统组件漏洞 3:安全基线 4: Linux系统漏洞 5: windows补丁</li>
	// <li>IfEmergency - String - 是否必填：否 - 是否为应急漏洞，查询应急漏洞传:yes</li>
	// <li>Status - String - 是否必填：是 - 漏洞状态筛选，0: 待处理 1:忽略  3:已修复  5:检测中， 控制台仅处理0,1,3,5四种状态</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareRiskOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Data *MalwareRiskOverview `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwareRiskOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareRiskOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseStrategyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略列表

		List []*RansomDefenseStrategy `json:"List,omitempty" name:"List"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRansomDefenseStrategyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseStrategyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBanWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重复机器的信息

		DuplicateHosts []*DuplicateHosts `json:"DuplicateHosts,omitempty" name:"DuplicateHosts"`
		// 添加规则是否重复

		IsDuplicate *bool `json:"IsDuplicate,omitempty" name:"IsDuplicate"`
		// 是否全局规则

		IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBanWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBanWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserKeyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公钥列表

		Keys []*AssetUserKeyInfo `json:"Keys,omitempty" name:"Keys"`
		// 分区总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetUserKeyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserKeyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineStrategyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户该策略下所有的基线id

		CategoryIds []*string `json:"CategoryIds,omitempty" name:"CategoryIds"`
		// 1 表示扫描过, 0没扫描过

		IfScanned *uint64 `json:"IfScanned,omitempty" name:"IfScanned"`
		// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机

		IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
		// 云服务器类型：
		// cvm：腾讯云服务器
		// bm：裸金属
		// ecm：边缘计算主机
		// lh: 轻量应用服务器
		// ohter: 混合云机器

		MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
		// 策略扫描通过率

		PassRate *uint64 `json:"PassRate,omitempty" name:"PassRate"`
		// 用户该策略下的所有主机id

		Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
		// 主机地域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 定期检测时间, 该时间下发扫描

		ScanAt *string `json:"ScanAt,omitempty" name:"ScanAt"`
		// 策略扫描周期(天)

		ScanCycle *string `json:"ScanCycle,omitempty" name:"ScanCycle"`
		// 策略名

		StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineStrategyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineStrategyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetAppProcessInfo struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 进程状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 进程版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeImpactedHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞影响机器列表数组

		ImpactedHosts []*ImpactedHost `json:"ImpactedHosts,omitempty" name:"ImpactedHosts"`
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImpactedHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImpactedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBaselineStrategyRequest struct {
	*tchttp.BaseRequest

	// 该策略下选择的基线id数组

	CategoryIds []*string `json:"CategoryIds,omitempty" name:"CategoryIds"`
	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 云主机类型：
	// cvm：腾讯云服务器
	// bm：裸金属
	// ecm：边缘计算主机
	// lh:轻量应用服务器
	// other:混合云机器

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 主机id数组

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 主机地域 ap-guangzhou

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 定期检测时间，该时间下发扫描

	ScanAt *string `json:"ScanAt,omitempty" name:"ScanAt"`
	// 检测周期

	ScanCycle *uint64 `json:"ScanCycle,omitempty" name:"ScanCycle"`
	// 策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
}

func (r *UpdateBaselineStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBaselineStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Strategy struct {

	// 基线id

	CategoryIds *string `json:"CategoryIds,omitempty" name:"CategoryIds"`
	// 是否可用

	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`
	// 主机数量

	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`
	// 是否默认策略

	IsDefault *uint64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 通过率

	PassRate *uint64 `json:"PassRate,omitempty" name:"PassRate"`
	// 基线检测项总数

	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// 扫描时间

	ScanAt *string `json:"ScanAt,omitempty" name:"ScanAt"`
	// 扫描周期

	ScanCycle *uint64 `json:"ScanCycle,omitempty" name:"ScanCycle"`
	// 策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略名

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
}

type CreateProtectServerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProtectServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProtectServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetUserListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetUserListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanVulSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanVulSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenRiskAssetsTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计详情图标数据 Name：展示主机ip 和地域， value：事件数量

		Chart []*ScreenNameValue `json:"Chart,omitempty" name:"Chart"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenRiskAssetsTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenRiskAssetsTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLicenseBindsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLicenseBindsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLicenseBindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateQuickProtectSettingRequest struct {
	*tchttp.BaseRequest

	// Quuid数组，云服务器uuid，最大100条

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 快捷防护类型：Vul-漏洞检测、Malware-文件查杀、Baseline-基线检测、AutoBan-自动阻断

	Type []*string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateQuickProtectSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQuickProtectSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MalwareInfo struct {

	// 影响广度 // 暂时不提供

	Breadth *string `json:"Breadth,omitempty" name:"Breadth"`
	// 木马检测平台用,分割 1云查杀引擎、2TAV、3binaryAi、4异常行为、5威胁情报

	CheckPlatform *string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
	// 首次发现时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 首次运行时间

	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`
	// 最近一次运行时间

	FileModifierTime *string `json:"FileModifierTime,omitempty" name:"FileModifierTime"`
	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件地址

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 文件大小

	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`
	// 危害描述

	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
	// 查询热度 // 暂时不提供

	Heat *string `json:"Heat,omitempty" name:"Heat"`
	// 服务器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 最近扫描时间

	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
	// 风险等级 0提示、1低、2中、3高、4严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 文件MD5

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机在线状态 OFFLINE  ONLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 外网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 最近修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 进程ID

	ProcessID *string `json:"ProcessID,omitempty" name:"ProcessID"`
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程树 json  pid:进程id，exe:文件路径 ，account:进程所属用组和用户 ,cmdline:执行命令，ssh_service: SSH服务ip, ssh_soure:登录源

	PsTree *string `json:"PsTree,omitempty" name:"PsTree"`
	// 参考链接

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 服务器名称

	ServersName *string `json:"ServersName,omitempty" name:"ServersName"`
	// 状态；4-:待处理，5-已信任，6-已隔离

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 最近访问时间

	StrFileAccessTime *string `json:"StrFileAccessTime,omitempty" name:"StrFileAccessTime"`
	// 建议方案

	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
	// 标签特性

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
}

type AssetEnvBaseInfo struct {

	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	//  附加信息
	//

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 服务器IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 服务器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 服务器外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 类型：
	// 0:用户变量
	// 1:系统变量

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 启动用户

	User *string `json:"User,omitempty" name:"User"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 环境变量值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeBaselineHostDetectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineHostDetect `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineHostDetectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineHostDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskProcessEventsRequest struct {
	*tchttp.BaseRequest

	// [StartTime:进程启动时间|DetectTime:最后检测时间]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>HostId - String - 是否必填：否 - 主机ID</li>
	// <li>IpOrName - String - 是否必填：否 - 主机IP或主机名</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名</li>
	// <li>ProcessId - String - 是否必填：否 - 进程ID</li>
	// <li>FilePath - String - 是否必填：否 - 进程路径</li>
	// <li>BeginTime - String - 是否必填：否 - 进程启动时间-开始</li>
	// <li>EndTime - String - 是否必填：否 - 进程启动时间-结束</li>
	// <li>Status - String - 是否必填：否 - 状态筛选 0待处理；1查杀中；2已查杀；3已退出;4已信任</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 [ASC|DESC]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRiskProcessEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskProcessEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProVersionRenewFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProVersionRenewFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProVersionRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户详细信息

		User *AssetUserDetail `json:"User,omitempty" name:"User"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectInstallCommandRequest struct {
	*tchttp.BaseRequest

	// 命令过期时间

	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
	// 地域标示

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// Vpc的ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeDirectConnectInstallCommandRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectInstallCommandRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProVersionInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增主机是否自动开通专业版

		IsAutoOpenProVersion *bool `json:"IsAutoOpenProVersion,omitempty" name:"IsAutoOpenProVersion"`
		// 后付费昨日扣费

		PostPayCost *uint64 `json:"PostPayCost,omitempty" name:"PostPayCost"`
		// 开通专业版主机数

		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`
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

type DescribeAssetInitServiceListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	// <li>Status- string - 是否必填：否 - 默认启用状态：0未启用， 1启用 仅linux</li>
	// <li>Type- string - 是否必填：否 - 类型：类型 仅windows：
	// 1:编码器
	// 2:IE插件
	// 3:网络提供者
	// 4:镜像劫持
	// 5:LSA提供者
	// 6:KnownDLLs
	// 7:启动执行
	// 8:WMI
	// 9:计划任务
	// 10:Winsock提供者
	// 11:打印监控器
	// 12:资源管理器
	// 13:驱动服务
	// 14:登录</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetInitServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetInitServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBanModeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBanModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBanModeRequest) FromJsonString(s string) error {
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

type ModifyEventAttackStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEventAttackStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEventAttackStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchBashRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchBashRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogStorageStatisticResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总容量（单位：GB）

		TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`
		// 已使用容量（单位：GB）

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

type DescribeRiskBatchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Handling:正在执行删除操作，
		// Pending：没有任务执行

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskBatchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskBatchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetAppListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetAppListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBanStatusRequest struct {
	*tchttp.BaseRequest

	// 阻断状态 0:关闭 1:开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyBanStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBanStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetLoadDetail struct {

	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 负载

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type CreateEmergencyResponseSosRequest struct {
	*tchttp.BaseRequest

	// 遇到的问题

	Problem *string `json:"Problem,omitempty" name:"Problem"`
	// 需要得到的帮助

	Support *string `json:"Support,omitempty" name:"Support"`
	// 0 应急响应 1旗舰护网

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *CreateEmergencyResponseSosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEmergencyResponseSosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineEffectHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 影响服务器列表

		EffectHostList []*BaselineEffectHost `json:"EffectHostList,omitempty" name:"EffectHostList"`
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineEffectHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetDiskPartitionInfo struct {

	// 分区名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 挂载目录

	Path *string `json:"Path,omitempty" name:"Path"`
	// 分区使用率

	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
	// 分区大小：单位G

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 文件系统类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 已使用空间：单位G

	Used *uint64 `json:"Used,omitempty" name:"Used"`
}

type DescribeAssetTypeTopRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetTypeTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTypeTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareFileRequest struct {
	*tchttp.BaseRequest

	// 木马记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeMalwareFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineFixListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineFix `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineFixListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineFixListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IgnoreRuleEffectHostInfo struct {

	// 事件id

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 最后检测时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 危害等级：1-低位，2-中危，3-高危，4-严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 状态：0-未通过，1-忽略，3-已通过，5-检测中

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 主机标签数组

	TagList []*string `json:"TagList,omitempty" name:"TagList"`
}

type EditPrivilegeRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditPrivilegeRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditPrivilegeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineInfo struct {

	// 机器ip。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 保护天数。

	ProtectDays *uint64 `json:"ProtectDays,omitempty" name:"ProtectDays"`
	// 防护级别。

	ProtectLevel *uint64 `json:"ProtectLevel,omitempty" name:"ProtectLevel"`
	// 地域。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 类型：bm或者cvm。

	RegionType *string `json:"RegionType,omitempty" name:"RegionType"`
	// 专业版开通时间。

	VipOpenTime *string `json:"VipOpenTime,omitempty" name:"VipOpenTime"`
}

type DescribeScreenDefenseTrendsRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeScreenDefenseTrendsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenDefenseTrendsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulCveIdInfoRequest struct {
	*tchttp.BaseRequest

	// 漏洞cve_id 列表

	CveIds []*string `json:"CveIds,omitempty" name:"CveIds"`
}

func (r *DescribeVulCveIdInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulCveIdInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderResource struct {

	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 到期时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 资源主键ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 授权类型

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeFileTamperRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情信息

		FileTamperRuleDetail *FileTamperRuleDetail `json:"FileTamperRuleDetail,omitempty" name:"FileTamperRuleDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileTamperRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperRuleInfoResponse) FromJsonString(s string) error {
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

type UntrustMalwaresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UntrustMalwaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UntrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryCreateSnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryCreateSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryCreateSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneInfo struct {

	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type DescribeDirectConnectInstallCommandResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安装命令的ip

		Ip *string `json:"Ip,omitempty" name:"Ip"`
		// 安装命令的token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectInstallCommandResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectInstallCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImpactedHostsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED：待处理 | FIXED：已修复）</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 漏洞种类ID。

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *DescribeImpactedHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImpactedHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | FIXED：已修复）
	//
	// Status过滤条件值只能取其一，不能是“或”逻辑。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 漏洞类型。
	// <li>WEB：Web应用漏洞</li>
	// <li>SYSTEM：系统组件漏洞</li>
	// <li>BASELINE：安全基线</li>

	VulType *string `json:"VulType,omitempty" name:"VulType"`
}

func (r *DescribeVulsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenBroadcastsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 播报文章列表

		List []*ScreenBroadcasts `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenBroadcastsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenBroadcastsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineItemDetectListRequest struct {
	*tchttp.BaseRequest

	// 0:过滤的结果导出；1:全部导出

	ExportAll *int64 `json:"ExportAll,omitempty" name:"ExportAll"`
	// <li>HostId - string - 是否必填：否 - 主机Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</i>
	// <li>IsPassed - int - 是否必填：否 - 是否通过</li>
	// <li>RiskTier - int - 是否必填：否 - 风险等级</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBaselineItemDetectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineItemDetectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginWhiteDimension struct {

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 地域字符串

	Locale *string `json:"Locale,omitempty" name:"Locale"`
	// 白名单地域

	Places []*Place `json:"Places,omitempty" name:"Places"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 白名单IP（多个IP逗号隔开）

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 白名单用户（多个用户逗号隔开）

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type WeeklyReport struct {

	// 周报开始时间。

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 周报结束时间。

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type CreateQuickProtectSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检测任务下发结果数组

		PushTaskResult []*PushTaskResult `json:"PushTaskResult,omitempty" name:"PushTaskResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateQuickProtectSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQuickProtectSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表查询结果列表

		HostList []*VulHostInfo `json:"HostList,omitempty" name:"HostList"`
		// 分页查询记录的总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// agent版本号

		AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
		// 免费木马剩余检测数量。

		FreeMalwaresLeft *uint64 `json:"FreeMalwaresLeft,omitempty" name:"FreeMalwaresLeft"`
		// 免费漏洞剩余检测数量。

		FreeVulsLeft *uint64 `json:"FreeVulsLeft,omitempty" name:"FreeVulsLeft"`
		// 是否有资产扫描记录，0无，1有

		HasAssetScan *uint64 `json:"HasAssetScan,omitempty" name:"HasAssetScan"`
		// CVM或BM主机唯一标识。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 是否开通专业版。
		// <li>true：是</li>
		// <li>false：否</li>

		IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
		// 机器ip。

		MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
		// 主机名称。

		MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
		// 操作系统。

		MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`
		// 机器所属地域。如：ap-guangzhou，ap-shanghai

		MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
		// 在线状态。
		// <li>ONLINE： 在线</li>
		// <li>OFFLINE：离线</li>

		MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
		// 云服务器类型。
		// <li>CVM: 腾讯云服务器</li>
		// <li>BM: 黑石物理机</li>
		// <li>ECM: 边缘计算服务器</li>
		// <li>LH: 轻量应用服务器</li>
		// <li>Other: 混合云机器</li>

		MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
		// 主机外网IP。

		MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
		// 主机状态。
		// <li>POSTPAY: 表示后付费，即按量计费  </li>
		// <li>PREPAY: 表示预付费，即包年包月</li>

		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
		// 专业版到期时间(仅预付费)

		ProVersionDeadline *string `json:"ProVersionDeadline,omitempty" name:"ProVersionDeadline"`
		// 专业版开通时间。

		ProVersionOpenDate *string `json:"ProVersionOpenDate,omitempty" name:"ProVersionOpenDate"`
		// 受云镜保护天数。

		ProtectDays *uint64 `json:"ProtectDays,omitempty" name:"ProtectDays"`
		// 防护版本 BASIC_VERSION 基础版, PRO_VERSION 专业版 Flagship 旗舰版.

		ProtectType *string `json:"ProtectType,omitempty" name:"ProtectType"`
		// CVM或BM主机唯一Uuid。

		Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
		// 云镜客户端唯一Uuid。

		Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckFileTamperRuleRequest struct {
	*tchttp.BaseRequest

	// 编辑时传的规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 填入的规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CheckFileTamperRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFileTamperRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportReverseShellEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportReverseShellEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineWeakPassword struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 密码Id

	PasswordId *int64 `json:"PasswordId,omitempty" name:"PasswordId"`
	// 密码

	WeakPassword *string `json:"WeakPassword,omitempty" name:"WeakPassword"`
}

type DescribeAssetWebFrameCountRequest struct {
	*tchttp.BaseRequest

	// 搜索条件：返回名称包含Name的所有Web框架列表

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetWebFrameCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebFrameCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表数据

		List []*TagMachine `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetProcessInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetProcessInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetProcessInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperEventsRequest struct {
	*tchttp.BaseRequest

	// 排序字段 CreateTime、ModifyTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略</li>
	// <li>ModifyTime - String - 是否必填：否 - 最近发生时间</li>
	// <li>Uuid- String - 是否必填：否 - 主机uuid查询</li>
	// <li>RuleCategory- string - 是否必填：否 - 规则类别 0 系统规则 1 自定义规则</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 ASC,DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeFileTamperEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncBaselineDetectSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正在检测的任务ID

		DetectingTaskIds []*int64 `json:"DetectingTaskIds,omitempty" name:"DetectingTaskIds"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 主机总数

		HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
		// 扫描中剩余时间(分钟)

		LeftMins *int64 `json:"LeftMins,omitempty" name:"LeftMins"`
		// 未通过策略总数

		NotPassPolicyCount *int64 `json:"NotPassPolicyCount,omitempty" name:"NotPassPolicyCount"`
		// 处理进度

		ProgressRate *int64 `json:"ProgressRate,omitempty" name:"ProgressRate"`
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 1:即将进行首次扫描   0:已经扫描过了

		WillFirstScan *int64 `json:"WillFirstScan,omitempty" name:"WillFirstScan"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncBaselineDetectSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncBaselineDetectSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetCoreModuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetCoreModuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetCoreModuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryCreateSnapshotRequest struct {
	*tchttp.BaseRequest

	// 修复任务id

	FixId *uint64 `json:"FixId,omitempty" name:"FixId"`
	// 任务进度返回的快照唯一Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *RetryCreateSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryCreateSnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetCoreModuleBaseInfo struct {

	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 模块ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	//  附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 服务器IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 服务器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 服务器外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 依赖模块数

	ModuleCount *uint64 `json:"ModuleCount,omitempty" name:"ModuleCount"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 依赖进程数

	ProcessCount *uint64 `json:"ProcessCount,omitempty" name:"ProcessCount"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 模块大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeBaselineRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基线检测项列表

		BaselineRuleList []*BaselineRuleInfo `json:"BaselineRuleList,omitempty" name:"BaselineRuleList"`
		// 是否显示说明列：true-是，false-否

		ShowRuleRemark *bool `json:"ShowRuleRemark,omitempty" name:"ShowRuleRemark"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeESHitsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// ES查询条件JSON

	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeESHitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeESHitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseRollBackTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务列表

		List []*RansomDefenseRollbackTask `json:"List,omitempty" name:"List"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRansomDefenseRollBackTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseRollBackTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulOverview struct {

	// 今日新增数量

	TodayCount *int64 `json:"TodayCount,omitempty" name:"TodayCount"`
	// 总数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type WebHookHostLabel struct {

	// 主机范围[1:所属项目|2:腾讯云标签|3:主机安全标签|4:自选]空数组为全部

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 主机项目或标签内容

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeSafeInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSafeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSafeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineTag struct {

	// 标签名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 关联标签ID

	Rid *int64 `json:"Rid,omitempty" name:"Rid"`
	// 标签ID

	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
}

type DeleteLicenseRecordAllResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLicenseRecordAllResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLicenseRecordAllResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RansomDefenseStrategyMachineDetail struct {

	// 云标签

	CloudTags []*Tag `json:"CloudTags,omitempty" name:"CloudTags"`
	// 硬盘信息，为空时所有硬盘生效：
	// ;分割 diskId1|diskName1;diskId2|diskName2

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// 版本信息：0-基础版 1-专业版 2-旗舰版 3-普惠版

	HostVersion *uint64 `json:"HostVersion,omitempty" name:"HostVersion"`
	// 主机实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 内网ip

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 外网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 可用区信息

	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`
	// 防护状态：0关闭，1开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 策略id，为0时未绑定策略

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 主机安全标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeLoginWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异地登录白名单数组

		LoginWhiteLists []*LoginWhiteLists `json:"LoginWhiteLists,omitempty" name:"LoginWhiteLists"`
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetLoadSummary struct {

	// 负载量数组，依次为：
	// [
	// 0%或未知数量，
	// 0%～20%，
	// 20%～50%，
	// 50%～80%，
	// 80%～100%
	// ]

	Counts []*uint64 `json:"Counts,omitempty" name:"Counts"`
	// 负载Top5

	Top5 []*AssetLoadDetail `json:"Top5,omitempty" name:"Top5"`
}

type DeleteLicenseRecordRequest struct {
	*tchttp.BaseRequest

	// 授权ID ,可以用授权订单列表获取.

	LicenseId *uint64 `json:"LicenseId,omitempty" name:"LicenseId"`
	// 授权类型

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DeleteLicenseRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLicenseRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeHistoryServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHistoryServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOrderAttributeRequest struct {
	*tchttp.BaseRequest

	// 可编辑的属性名称 ,当前支持的有: alias 资源别名

	AttrName *string `json:"AttrName,omitempty" name:"AttrName"`
	// 属性值

	AttrValue *string `json:"AttrValue,omitempty" name:"AttrValue"`
	// 授权类型 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *ModifyOrderAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOrderAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineLicenseDetail struct {

	// xxx

	InquireKey *string `json:"InquireKey,omitempty" name:"InquireKey"`
	// xx

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// xxx

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// xxx

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
}

type DescribeFastAnalysisRequest struct {
	*tchttp.BaseRequest

	// 日志字段名称

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 起始时间，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 查询语句，语句长度最大为4096

	Query *string `json:"Query,omitempty" name:"Query"`
	// 结束时间，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
}

func (r *DescribeFastAnalysisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFastAnalysisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulFixedInfo struct {

	// 修护时间

	FixTime *string `json:"FixTime,omitempty" name:"FixTime"`
	// 1:快照备份修护，2:直接修护

	FixType *uint64 `json:"FixType,omitempty" name:"FixType"`
	// 漏洞等级 1:低 2:中 3:高 4:提示

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

type AssetMachineDetail struct {

	// agent版本

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// 系统启动时间

	BootTime *string `json:"BootTime,omitempty" name:"BootTime"`
	// 专业版开通时间

	BuyTime *string `json:"BuyTime,omitempty" name:"BuyTime"`
	// 内核版本

	CoreVersion *string `json:"CoreVersion,omitempty" name:"CoreVersion"`
	// CPU信息

	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`
	// Cpu使用率百分比

	CpuLoad *string `json:"CpuLoad,omitempty" name:"CpuLoad"`
	// CpuLoadVul

	CpuLoadVul *string `json:"CpuLoadVul,omitempty" name:"CpuLoadVul"`
	// Cpu数量

	CpuSize *uint64 `json:"CpuSize,omitempty" name:"CpuSize"`
	// 设备型号

	DeviceVersion *string `json:"DeviceVersion,omitempty" name:"DeviceVersion"`
	// 硬盘使用率百分比

	DiskLoad *string `json:"DiskLoad,omitempty" name:"DiskLoad"`
	// 硬盘容量：单位G

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 分区

	Disks []*AssetDiskPartitionInfo `json:"Disks,omitempty" name:"Disks"`
	// 专业版到期时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// FirstTime

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 安装时间

	InstallTime *string `json:"InstallTime,omitempty" name:"InstallTime"`
	// 主机ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 最后上线时间

	LastLiveTime *string `json:"LastLiveTime,omitempty" name:"LastLiveTime"`
	// 主机二外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 服务器内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 服务器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 内存使用率百分比

	MemLoad *string `json:"MemLoad,omitempty" name:"MemLoad"`
	// 内存容量：单位G

	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`
	// 网卡

	NetCards []*AssetNetworkCardInfo `json:"NetCards,omitempty" name:"NetCards"`
	// 离线时间

	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`
	// 操作系统名称

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// linux/windows

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 分区数

	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`
	// 生产商

	Producer *string `json:"Producer,omitempty" name:"Producer"`
	// 业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 已防护天数

	ProtectDays *uint64 `json:"ProtectDays,omitempty" name:"ProtectDays"`
	// 防护级别：0基础版，1专业版，2旗舰版，3普惠版

	ProtectLevel *uint64 `json:"ProtectLevel,omitempty" name:"ProtectLevel"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 风险状态：UNKNOW-未知，RISK-风险，SAFT-安全

	RiskStatus *string `json:"RiskStatus,omitempty" name:"RiskStatus"`
	// 序列号

	SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`
	// 0在线，1已离线

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 服务器uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeMonthInspectionReportRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页步长

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMonthInspectionReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonthInspectionReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProtectNetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全管家数据

		List []*ProtectNetInfo `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProtectNetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectNetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDatabaseCountRequest struct {
	*tchttp.BaseRequest

	// 搜索条件：返回名称包含Name的数据库列表

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetDatabaseCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebAppPluginListRequest struct {
	*tchttp.BaseRequest

	// Web应用ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebAppPluginListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppPluginListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrustMaliciousRequestRequest struct {
	*tchttp.BaseRequest

	// 恶意请求记录ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *TrustMaliciousRequestRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TrustMaliciousRequestRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTagsRequest struct {
	*tchttp.BaseRequest

	// 标签ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 标签名

	Name *string `json:"Name,omitempty" name:"Name"`
	// Quuid

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *EditTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UntrustMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马ID数组 (最大100条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *UntrustMalwaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UntrustMalwaresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScreenVulInfo struct {

	// 漏洞类型 1: web-cms漏洞, 2:应用漏洞, 4: Linux软件漏洞, 5: Windows系统漏洞

	Category *uint64 `json:"Category,omitempty" name:"Category"`
	// 漏洞事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 最后检测时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 漏洞等级 1:低 2:中 3:高 4:提示

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 漏洞名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主机UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

type CreateBanWhiteListRequest struct {
	*tchttp.BaseRequest

	// 事件ID，事件列表加白名单时传递，白名单添加成功后，会将事件变成已加白状态

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 阻断规则

	Rules *BanWhiteList `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateBanWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBanWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 机器所属地域。如：ap-guangzhou，ap-shanghai

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	// <li>ECM:  表示边缘计算服务器</li>
	// <li>LH:  表示轻量应用服务器</li>
	// <li>Other:  表示混合云机器</li>

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMachineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetScanTaskRequest struct {
	*tchttp.BaseRequest

	// 未设置Quuid时默认扫描所有主机，否则扫描指定主机

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 未任务ID时创建新任务，设置TaskId时为指定TaskId的重试任务

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CreateAssetScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 地域标志，如 ap-guangzhou，ap-shanghai，ap-beijing

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域代码，如 gz，sh，bj

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域中文名，如华南地区（广州），华东地区（上海金融），华北地区（北京）

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域英文名

	RegionNameEn *string `json:"RegionNameEn,omitempty" name:"RegionNameEn"`
}

type CreateProcessTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProcessTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProcessTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRansomDefenseEventsListRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>HostName- string- 主机名称</li>
	// <li>Status - Uint64 - 0待处理，1已处理，2已信任</li>
	// <li>HostIp- String - 主机ip</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportRansomDefenseEventsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRansomDefenseEventsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基础版机器数。

		BasicVersionNum *uint64 `json:"BasicVersionNum,omitempty" name:"BasicVersionNum"`
		// 主机总数。

		HostNum *uint64 `json:"HostNum,omitempty" name:"HostNum"`
		// 受影响的专业版主机数。

		ImpactedHostNum *uint64 `json:"ImpactedHostNum,omitempty" name:"ImpactedHostNum"`
		// 专业版机器数。

		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`
		// 漏洞数量。

		VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulScanResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBashEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBashEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicProxyInstallCommandRequest struct {
	*tchttp.BaseRequest

	// nginx主机ip列表，逗号分隔

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribePublicProxyInstallCommandRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicProxyInstallCommandRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginWhiteLists struct {

	// 创建白名单时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 机器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否为全局规则

	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 机器名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 修改白名单时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 白名单地域

	Places []*Place `json:"Places,omitempty" name:"Places"`
	// 白名单IP（多个IP逗号隔开）

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 白名单用户（多个用户逗号隔开）

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 云镜客户端ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ResponseList struct {

	// 文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 主机ip

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
}

type UsualPlace struct {

	// 城市 ID。

	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
	// 国家 ID。

	CountryId *uint64 `json:"CountryId,omitempty" name:"CountryId"`
	// ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 省份 ID。

	ProvinceId *uint64 `json:"ProvinceId,omitempty" name:"ProvinceId"`
	// 云镜客户端唯一标识UUID。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeSecurityProtectionStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0 ：0付费资产情况， 1：存在未安装agent情况 ，2：部分资产已是专业版/旗舰版， 3：全部资产已是专业版/旗舰版

		AssetManageStat *uint64 `json:"AssetManageStat,omitempty" name:"AssetManageStat"`
		// 密码破解是否开启防护 0:未开启防护或0付费资产情况 1:已开启防护 2:存在带处理事件

		DefenseBruteAttackStat *uint64 `json:"DefenseBruteAttackStat,omitempty" name:"DefenseBruteAttackStat"`
		// 检测--密码破解有无存在风险  0:存在未处理风险, 1：无风险，正常检测

		DiscoverBruteAttackStat *uint64 `json:"DiscoverBruteAttackStat,omitempty" name:"DiscoverBruteAttackStat"`
		// 高危命令，0：0台开通专业版/旗舰版， 1：存在未处理风险， 2：已有付费资产，无风险

		EventBashStat *uint64 `json:"EventBashStat,omitempty" name:"EventBashStat"`
		// 专家服务 0:未开通服务, 1:已开通

		ExpertServiceStat *uint64 `json:"ExpertServiceStat,omitempty" name:"ExpertServiceStat"`
		// 核心文件监控 0:未开启防护（0付费资产情况）,1: 已开通

		FileTamperStat *uint64 `json:"FileTamperStat,omitempty" name:"FileTamperStat"`
		// 日志分析 0:未开通服务, 1:已开通

		LogAnalysisStat *uint64 `json:"LogAnalysisStat,omitempty" name:"LogAnalysisStat"`
		// 异常登录 0:存在未处理风险,1:无风险，未配置白名单,2:无风险，已配置

		LoginLogStat *uint64 `json:"LoginLogStat,omitempty" name:"LoginLogStat"`
		// 恶意请求 0 : 0台开通专业版/旗舰版, 1: 恶意请求 存在未处理风险, 2:已有付费资产，无风险

		MaliciousRequestStat *uint64 `json:"MaliciousRequestStat,omitempty" name:"MaliciousRequestStat"`
		// 0:从未检测过，或0资产付费情况, 1:已检测，存在恶意文件, 2:已检测，未开启隔离防护, 3:已检测且已开启防护且无风险

		MalwareScanStat *uint64 `json:"MalwareScanStat,omitempty" name:"MalwareScanStat"`
		// 本地提权 0:0台开通专业版/旗舰版, 1:存在未处理风险 2:已有付费资产，无风险

		PrivilegeStat *uint64 `json:"PrivilegeStat,omitempty" name:"PrivilegeStat"`
		// 反弹shell 0:0台开通专业版/旗舰版, 1:存在未处理风险 2:已有付费资产，无风险

		ReverseShellStat *uint64 `json:"ReverseShellStat,omitempty" name:"ReverseShellStat"`
		// 0:从未检测过，或0资产付费情况, 1:存在基线风险,2:无风险

		SecureBasicLineStat *uint64 `json:"SecureBasicLineStat,omitempty" name:"SecureBasicLineStat"`
		// 0:从未检测过，或0资产付费情况, 1:存在漏洞风险, 2:无风险

		VulManageStat *uint64 `json:"VulManageStat,omitempty" name:"VulManageStat"`
		// 安全告警 0:未开通设置（全部关闭） 1:已开通（只要开启1个就算）

		WarningSetStat *uint64 `json:"WarningSetStat,omitempty" name:"WarningSetStat"`
		// 网页防篡改  0:未开通, 1:已开通

		WebPageStat *uint64 `json:"WebPageStat,omitempty" name:"WebPageStat"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityProtectionStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityProtectionStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetAttackWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetAttackWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetAttackWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetAppBaseInfo struct {

	// 二进制路径

	BinPath *string `json:"BinPath,omitempty" name:"BinPath"`
	// 配置文件路径

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 应用描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 关联进程数

	ProcessCount *uint64 `json:"ProcessCount,omitempty" name:"ProcessCount"`
	// 主机业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 应用类型
	// 1: 运维
	// 2 : 数据库
	// 3 : 安全
	// 4 : 可疑应用
	// 5 : 系统架构
	// 6 : 系统应用
	// 7 : WEB服务
	// 99: 其他

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 版本号

	Version *string `json:"Version,omitempty" name:"Version"`
}

type SearchTemplate struct {

	// 检索语句

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 展示数据

	DisplayData *string `json:"DisplayData,omitempty" name:"DisplayData"`
	// 检索方式。输入框检索：standard,过滤，检索：simple

	Flag *string `json:"Flag,omitempty" name:"Flag"`
	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 检索索引类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 检索名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 转换的检索语句内容

	Query *string `json:"Query,omitempty" name:"Query"`
	// 时间范围

	TimeRange *string `json:"TimeRange,omitempty" name:"TimeRange"`
}

type DescribeEventlogRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>DateRange - String - 是否必填：否 - 日志时间范围 </li>
	// <li>SeverityClass - String  是否必填：否 - 威胁等级(10:严重 20:高危 30:中危 40:低危)</li>
	// <li>EventType - String  是否必填：否 - 事件类型</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeEventlogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventlogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetMachineDetailRequest struct {
	*tchttp.BaseRequest

	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetMachineDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetMachineDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefencePluginEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件Id 可通过ExportTasks接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulDefencePluginEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefencePluginEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRiskProcessEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRiskProcessEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRiskProcessEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselinePolicy struct {

	// 资产类型[0:所有专业版旗舰版|1:id|2:ip]

	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
	// 检测间隔[1:1天|3:3天|5:5天|7:7天]

	DetectInterval *int64 `json:"DetectInterval,omitempty" name:"DetectInterval"`
	// 检测时间

	DetectTime *string `json:"DetectTime,omitempty" name:"DetectTime"`
	// 关联基线主机数目

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 主机Id

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 主机Ip

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 是否是系统默认

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 是否开启[0:未开启|1:开启]

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 关联基线项数目

	ItemCount *int64 `json:"ItemCount,omitempty" name:"ItemCount"`
	// 策略Id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称,长度不超过128英文字符

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 关联基线项数目

	RuleCount *int64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// 规则Id

	RuleIds []*int64 `json:"RuleIds,omitempty" name:"RuleIds"`
}

type RansomDefenseRollbackTask struct {

	// 快照时间

	BackupTime *string `json:"BackupTime,omitempty" name:"BackupTime"`
	// 操作时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 硬盘id列表，;分隔

	Disks *string `json:"Disks,omitempty" name:"Disks"`
	// 任务ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// Status!=0时为完成时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 可用区信息

	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`
	// 回滚任务状态：0进行中，1成功，2失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
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

type DescribeMalwareTimingScanSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否自动隔离：1-是，0-否

		AutoIsolation *uint64 `json:"AutoIsolation,omitempty" name:"AutoIsolation"`
		// 检测模式 0 全盘检测  1快速检测

		CheckPattern *uint64 `json:"CheckPattern,omitempty" name:"CheckPattern"`
		// 一键扫描超时时长，如：1800秒（s）

		ClickTimeout *uint64 `json:"ClickTimeout,omitempty" name:"ClickTimeout"`
		// 周期 1每天

		Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`
		// 启发引擎 0 关闭 1开启

		EnableInspiredEngine *uint64 `json:"EnableInspiredEngine,omitempty" name:"EnableInspiredEngine"`
		// 是否开启恶意进程查杀[0:未开启,1:开启]

		EnableMemShellScan *uint64 `json:"EnableMemShellScan,omitempty" name:"EnableMemShellScan"`
		// 定时检测开关 0 关闭1 开启

		EnableScan *int64 `json:"EnableScan,omitempty" name:"EnableScan"`
		// 检测周期 超时结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 1标准模式（只报严重、高危）、2增强模式（报严重、高危、中危）、3严格模式（报严重、高、中、低、提示）

		EngineType *uint64 `json:"EngineType,omitempty" name:"EngineType"`
		// 唯一ID

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 是否全部服务器 1 全部 2 自选

		IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
		// 是否杀掉进程 1杀掉 0不杀掉 只有开启自动隔离才生效

		KillProcess *uint64 `json:"KillProcess,omitempty" name:"KillProcess"`
		// 监控模式 0 标准 1深度

		MonitoringPattern *uint64 `json:"MonitoringPattern,omitempty" name:"MonitoringPattern"`
		// 自选服务器时必须 主机quuid的string数组

		QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
		// 实时监控0 关闭 1开启

		RealTimeMonitoring *int64 `json:"RealTimeMonitoring,omitempty" name:"RealTimeMonitoring"`
		// 检测周期 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwareTimingScanSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareTimingScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBanWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 白名单列表

		WhiteList []*BanWhiteListDetail `json:"WhiteList,omitempty" name:"WhiteList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBanWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBanWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReverseShell struct {

	// 命令详情

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 产生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 检测方法

	DetectBy *uint64 `json:"DetectBy,omitempty" name:"DetectBy"`
	// 目标IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 目标端口

	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`
	// 进程路径

	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`
	// 主机内网IP

	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
	// ID 主键

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	//  主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 父进程用户组

	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`
	// 父进程名

	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`
	// 父进程路径

	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`
	// 父进程用户

	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`
	// 进程树

	ProcTree *string `json:"ProcTree,omitempty" name:"ProcTree"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 主机ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 处理状态：0-待处理 2-白名单 3-已处理 4-已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 执行用户组

	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`
	// 执行用户

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 云镜UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeMachineOsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作系统列表

		List []*OsName `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineOsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineOsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRansomDefenseStrategyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRansomDefenseStrategyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRansomDefenseStrategyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持功能

		Functions []*string `json:"Functions,omitempty" name:"Functions"`
		// 支持功能英文翻译

		FunctionsEn []*string `json:"FunctionsEn,omitempty" name:"FunctionsEn"`
		// 授权总数

		PurchaseCount *int64 `json:"PurchaseCount,omitempty" name:"PurchaseCount"`
		// 授权类型, 专业版：Pro 旗舰版：Flagship

		Type *string `json:"Type,omitempty" name:"Type"`
		// 可用授权数

		UnusedCount *int64 `json:"UnusedCount,omitempty" name:"UnusedCount"`
		// 已用授权数

		UsedCount *int64 `json:"UsedCount,omitempty" name:"UsedCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteCombinedListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 合并后的白名单列表

		LoginWhiteCombinedInfos []*LoginWhiteCombinedInfo `json:"LoginWhiteCombinedInfos,omitempty" name:"LoginWhiteCombinedInfos"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginWhiteCombinedListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginWhiteCombinedListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文本内容

		Context *string `json:"Context,omitempty" name:"Context"`
		// 受影响机器数

		EffectHostCount *uint64 `json:"EffectHostCount,omitempty" name:"EffectHostCount"`
		// 受影响事件类型 0 无 1 木马 2 漏洞 3基线

		EventCategory *uint64 `json:"EventCategory,omitempty" name:"EventCategory"`
		// 受影响事件名称

		EventName *string `json:"EventName,omitempty" name:"EventName"`
		// 是否展示通知

		IsShow *bool `json:"IsShow,omitempty" name:"IsShow"`
		// 标题

		Title *string `json:"Title,omitempty" name:"Title"`
		// 超链接地址

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSafeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSafeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BashEventNew struct {

	// 执行命令

	BashCmd *string `json:"BashCmd,omitempty" name:"BashCmd"`
	// 发生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 0: bash日志 1: 实时监控(雷霆版)

	DetectBy *uint64 `json:"DetectBy,omitempty" name:"DetectBy"`
	// 进程名称

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 主机内网IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 数据ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 机器额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 0:普通 1:专业版 2:旗舰版

	MachineType *int64 `json:"MachineType,omitempty" name:"MachineType"`
	// 处理时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 进程id

	Pid *string `json:"Pid,omitempty" name:"Pid"`
	// 平台类型

	Platform *uint64 `json:"Platform,omitempty" name:"Platform"`
	// 主机ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 自动生成的正则表达式

	RegexBashCmd *string `json:"RegexBashCmd,omitempty" name:"RegexBashCmd"`
	// 规则类别  0=系统规则，1=用户规则

	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`
	// 规则ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则等级：1-高 2-中 3-低

	RuleLevel *uint64 `json:"RuleLevel,omitempty" name:"RuleLevel"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 处理状态： 0 = 待处理 1= 已处理, 2 = 已加白， 3 = 已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 执行用户名

	User *string `json:"User,omitempty" name:"User"`
	// 云镜ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeLicenseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 授权数列表信息

		List []*LicenseDetail `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表

		Machines []*Machine `json:"Machines,omitempty" name:"Machines"`
		// 主机数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenProtectionCntResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机安全防护引擎介绍 内容

		List []*ScreenProtectionCnt `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenProtectionCntResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenProtectionCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenRiskAssetsTopRequest struct {
	*tchttp.BaseRequest

	// 统计类型：0:潜在威胁 1:失陷 2:漏洞 3:基线

	BusinessType *uint64 `json:"BusinessType,omitempty" name:"BusinessType"`
}

func (r *DescribeScreenRiskAssetsTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenRiskAssetsTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineDefaultStrategyListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBaselineDefaultStrategyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineDefaultStrategyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基线详情

		BaselineDetail *BaselineDetail `json:"BaselineDetail,omitempty" name:"BaselineDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperRuleCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机关联核心文件规则数量信息

		List []*FileTamperRuleCount `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileTamperRuleCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperRuleCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportIgnoreRuleEffectHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务Id , 可通过ExportTasks 接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportIgnoreRuleEffectHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportIgnoreRuleEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbTestStatusRequest struct {
	*tchttp.BaseRequest

	// 灰度ID。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeAbTestStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbTestStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefencePluginStatusRequest struct {
	*tchttp.BaseRequest

	// 排序列，严格相等：CreateTime创建时间，ModifyTime更新时间

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Exception - String - 是否必填：否 - 插件状态 0:正常,1:异常,2:无java进程注入</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 数据限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 数据偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序，大小写无关：asc 升序，desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVulDefencePluginStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefencePluginStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HistoryAccount struct {

	// 唯一ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机内网IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 变更时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 帐号变更类型。
	// <li>CREATE：表示新增帐号</li>
	// <li>MODIFY：表示修改帐号</li>
	// <li>DELETE：表示删除帐号</li>

	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`
	// 帐号名。

	Username *string `json:"Username,omitempty" name:"Username"`
	// 云镜客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetWebLocationPathListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Paths []*AssetWebLocationPath `json:"Paths,omitempty" name:"Paths"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebLocationPathListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationPathListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetProcessInfoListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime|StartTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name - String - 是否必填：否 - 进程名</li>
	// <li>User - String - 是否必填：否 - 进程用户</li>
	// <li>Group - String - 是否必填：否 - 进程用户组</li>
	// <li>Pid - uint64 - 是否必填：否 - 进程ID</li>
	// <li>Ppid - uint64 - 是否必填：否 - 父进程ID</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>Status - string - 是否必填：否 - 进程状态：
	// 0:全部
	// 1:R 可执行
	// 2:S 可中断
	// 3:不可中断</li>
	// <li>RunTimeStart - String - 是否必填：否 - 运行开始时间</li>
	// <li>RunTimeEnd - String - 是否必填：否 - 运行结束时间</li>
	// <li>InstallByPackage - uint64 - 是否必填：否 - 是否包安装：0否，1是</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetProcessInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetProcessInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanMalwareScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否正在扫描中

		IsSchedule *bool `json:"IsSchedule,omitempty" name:"IsSchedule"`
		// 风险文件数,当进度满了以后才有该值

		RiskFileNumber *int64 `json:"RiskFileNumber,omitempty" name:"RiskFileNumber"`
		// 0 从未扫描过、 1 扫描中、 2扫描完成、 3停止中、 4停止完成

		ScanStatus *uint64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
		// 扫描进度（单位：%）

		Schedule *int64 `json:"Schedule,omitempty" name:"Schedule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanMalwareScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanMalwareScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProcessTaskRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *CreateProcessTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProcessTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWebHookRuleRequest struct {
	*tchttp.BaseRequest

	// 规则Id列表

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteWebHookRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWebHookRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DescribeAllMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机信息数组。

		Machines []*FullMachine `json:"Machines,omitempty" name:"Machines"`
		// 记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProVersionStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开通状态。
		// <li>UNOPENED：未开通专业版</li>
		// <li>OPENED：已开通专业版</li>

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProVersionStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProVersionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportVulsRequest struct {
	*tchttp.BaseRequest

	// 专业版周报开始时间。

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportVulsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportVulsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JavaMemShellPluginStatus struct {

	// 服务器名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 错误日志

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
	// 服务器ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 注入进程主类名

	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`
	// 注入进程pid

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 插件状态：0: 注入中, 1: 注入成功, 2: 插件超时, 3: 插件退出, 4: 注入失败 5: 软删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DescribeAgentInstallCommandResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Arm安装命令

		ARMCommand *string `json:"ARMCommand,omitempty" name:"ARMCommand"`
		// linux系统安装命令

		LinuxCommand *string `json:"LinuxCommand,omitempty" name:"LinuxCommand"`
		// windows系统安装命令（windows2008及以上）

		WindowsCommand *string `json:"WindowsCommand,omitempty" name:"WindowsCommand"`
		// windows版agent下载链接

		WindowsDownloadUrl *string `json:"WindowsDownloadUrl,omitempty" name:"WindowsDownloadUrl"`
		// windows系统安装命令第一步（windows2003）

		WindowsStepOne *string `json:"WindowsStepOne,omitempty" name:"WindowsStepOne"`
		// windows系统安装命令第二步（windows2003）

		WindowsStepTwo *string `json:"WindowsStepTwo,omitempty" name:"WindowsStepTwo"`
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

type ExportRiskDnsEventListRequest struct {
	*tchttp.BaseRequest

	// 排序字段：[AccessCount:请求次数|LastTime:最近请求时间]

	By *string `json:"By,omitempty" name:"By"`
	// <li>IpOrName - String - 是否必填：否 - 主机Ip或别名筛选</li>
	// <li>HostId - String - 是否必填：否 - 主机Id</li>
	// <li>AgentId - String - 是否必填：否 - 客户端Id</li>
	// <li>PolicyType - String - 是否必填：否 - 策略类型,0:系统策略1:用户自定义策略</li>
	// <li>Domain - String - 是否必填：否 - 域名(先对域名做urlencode,再base64)</li>
	// <li>HandleStatus - String - 是否必填：否 - 状态筛选0:待处理；2:信任；3:不信任</li>
	// <li>BeginTime - String - 是否必填：否 - 最近访问开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 最近访问结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式：[ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportRiskDnsEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRiskDnsEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MaliciousRequest struct {

	// 执行命令行。

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 恶意请求数。

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 发现时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 恶意请求域名描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 恶意请求域名。

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 记录ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机内网IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 记录合并时间。

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// 进程PID。

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 进程MD5
	// 值。

	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
	// 进程名。

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 参考地址。

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 记录状态。
	// <li>UN_OPERATED：待处理</li>
	// <li>TRUSTED：已信任</li>
	// <li>UN_TRUSTED：已取消信任</li>

	Status *string `json:"Status,omitempty" name:"Status"`
	// 云镜客户端UUID。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportFileTamperEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID 可通过ExportTasks接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportFileTamperEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportFileTamperEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmSettings struct {

	// 邮件告警。

	DisableMailWarning *string `json:"DisableMailWarning,omitempty" name:"DisableMailWarning"`
	// 手机告警。

	DisablePhoneWarning *string `json:"DisablePhoneWarning,omitempty" name:"DisablePhoneWarning"`
	// 网页告警。

	DisableWebWarning *string `json:"DisableWebWarning,omitempty" name:"DisableWebWarning"`
}

type DescribeBashRulesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(规则名称)</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 0-系统规则; 1-用户规则

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeBashRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRansomDefenseEventsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRansomDefenseEventsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRansomDefenseEventsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostLoginWhiteObj struct {

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 白名单生效的机器信息列表

	HostInfos []*HostInfo `json:"HostInfos,omitempty" name:"HostInfos"`
	// 是否对全局生效, 1：全局有效 0: 仅针对单台主机'

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 加白地域

	Places []*Place `json:"Places,omitempty" name:"Places"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 加白源IP，支持网段，多个IP以逗号隔开

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 加白用户名，多个用户名以逗号隔开

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type ScreenRegionMachines struct {

	// 潜在风险主机数

	AttackCnt *uint64 `json:"AttackCnt,omitempty" name:"AttackCnt"`
	// 省略展示多少主机，等于0时没有省略展示

	IgnoreCnt *uint64 `json:"IgnoreCnt,omitempty" name:"IgnoreCnt"`
	// 主机列表

	Machines []*ScreenMachine `json:"Machines,omitempty" name:"Machines"`
	// 所有区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 区域中文描述

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 风险主机数量

	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
	// 无风险主机数

	SafetyCnt *uint64 `json:"SafetyCnt,omitempty" name:"SafetyCnt"`
	// 此区域的主机总数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 离线/未安装主机数

	UnAgentOfflineCnt *uint64 `json:"UnAgentOfflineCnt,omitempty" name:"UnAgentOfflineCnt"`
}

type DescribeBaselineDetailRequest struct {
	*tchttp.BaseRequest

	// 基线id

	BaselineId *uint64 `json:"BaselineId,omitempty" name:"BaselineId"`
}

func (r *DescribeBaselineDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrivilegeRule struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 主机IP

	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 是否S权限

	SMode *uint64 `json:"SMode,omitempty" name:"SMode"`
	// 状态(0: 有效 1: 无效)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 客户端ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribePrivilegeEventInfoRequest struct {
	*tchttp.BaseRequest

	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribePrivilegeEventInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivilegeEventInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProtectDirResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProtectDirResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProtectDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineTopRequest struct {
	*tchttp.BaseRequest

	// 策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 动态top值

	Top *uint64 `json:"Top,omitempty" name:"Top"`
}

func (r *DescribeBaselineTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulStoreListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [PublishDate]

	By *string `json:"By,omitempty" name:"By"`
	//
	// <li>VulName- string - 是否必填：否 - 漏洞名称</li>
	// <li>CveId- string - 是否必填：否 - cveid</li>
	// <li>VulCategory- string - 是否必填：否 - 漏洞分类  1 Web-CMS漏洞 ,2 应用漏洞 ,4 Linux软件漏洞,5 Windows系统漏洞</li>
	// <li>Method- string - 是否必填：否 - 检测方法 0版本对比,1 poc检测 </li>
	// <li>SupportDefense- string - 是否必填：否 - 是否支持防御 0不支持,1支持</li>
	// <li>FixSwitch- string - 是否必填：否 - 是否支持自动修复 0不支持,1支持</li>
	//

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVulStoreListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulStoreListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebPageProtectSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWebPageProtectSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebPageProtectSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWhiteListOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源对象

		Resource *OrderResource `json:"Resource,omitempty" name:"Resource"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWhiteListOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWhiteListOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBashEventsRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBashEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBashEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BashPolicy struct {

	// 0:告警 1:白名单 2:拦截

	BashAction *int64 `json:"BashAction,omitempty" name:"BashAction"`
	// 策略类型，0:系统  1:用户

	Category *int64 `json:"Category,omitempty" name:"Category"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否处理旧事件为白名单 0=不处理 1=处理

	DealOldEvents *int64 `json:"DealOldEvents,omitempty" name:"DealOldEvents"`
	// 策略描述

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// 1:有效 0:无效

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 事件列表点击“加入白名单”时,需要传EventId 事件的id

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 策略ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 危险等级(0:无，1: 高危 2:中危 3: 低危)

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 生效主机的QUUID集合

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 正则表达式

	Rule *string `json:"Rule,omitempty" name:"Rule"`
	// 生效范围（0:一组quuid 1:所有专业版(包含旗舰版) 2:所有旗舰版 3:所有主机）

	Scope *int64 `json:"Scope,omitempty" name:"Scope"`
	// 老版本兼容可能会用到

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
	// 0:黑名单 1:白名单

	White *int64 `json:"White,omitempty" name:"White"`
}

type JavaMemShellDetail struct {

	// 注释

	Annotations *string `json:"Annotations,omitempty" name:"Annotations"`
	// java进程命令行参数

	Args *string `json:"Args,omitempty" name:"Args"`
	// java内存马二进制代码(base64)

	ClassContent *string `json:"ClassContent,omitempty" name:"ClassContent"`
	// java内存马反编译代码

	ClassContentPretty *string `json:"ClassContentPretty,omitempty" name:"ClassContentPretty"`
	// java加载器类名

	ClassLoaderName *string `json:"ClassLoaderName,omitempty" name:"ClassLoaderName"`
	// 类名

	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`
	// 首次发现时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件描述

	EventDescription *string `json:"EventDescription,omitempty" name:"EventDescription"`
	// java进程路径

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 容器名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例状态：RUNNING,STOPPED,SHUTDOWN...

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 继承的接口

	Interfaces *string `json:"Interfaces,omitempty" name:"Interfaces"`
	//  主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 实例状态：RUNNING,STOPPED,SHUTDOWN...

	MachineState *string `json:"MachineState,omitempty" name:"MachineState"`
	// 类文件MD5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 进程pid

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公共ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 最近检测时间

	RecentFoundTime *string `json:"RecentFoundTime,omitempty" name:"RecentFoundTime"`
	// 安全建议

	SecurityAdvice *string `json:"SecurityAdvice,omitempty" name:"SecurityAdvice"`
	// 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略  4 - 已手动处理

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 父类名

	SuperClassName *string `json:"SuperClassName,omitempty" name:"SuperClassName"`
	// 内存马类型  0:Filter型 1:Listener型 2:Servlet型 3:Interceptors型 4:Agent型 5:其他

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type DescribeAssetWebServiceProcessListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 进程列表

		Process []*AssetAppProcessInfo `json:"Process,omitempty" name:"Process"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebServiceProcessListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemIgnoreListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineItemInfo `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineItemIgnoreListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFileTamperEventsRequest struct {
	*tchttp.BaseRequest

	// 对应事件id

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// Status 1 -- 加白 2 -- 删除 3 - 忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyFileTamperEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFileTamperEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventStat struct {

	// 事件数

	EventsNum *uint64 `json:"EventsNum,omitempty" name:"EventsNum"`
	// 受影响的主机数

	MachineAffectNum *uint64 `json:"MachineAffectNum,omitempty" name:"MachineAffectNum"`
}

type BatchAddMachineTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchAddMachineTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchAddMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// 暴力破解事件Id数组。(最大 100条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBruteAttacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBruteAttacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求列表数组

		RiskDnsList []*RiskDnsList `json:"RiskDnsList,omitempty" name:"RiskDnsList"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
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

type DescribeTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表信息

		List []*Tag `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞CVE。

		CveId *string `json:"CveId,omitempty" name:"CveId"`
		// 漏洞描述。

		Description *string `json:"Description,omitempty" name:"Description"`
		// 参考链接。

		Reference *string `json:"Reference,omitempty" name:"Reference"`
		// 修复方案。

		RepairPlan *string `json:"RepairPlan,omitempty" name:"RepairPlan"`
		// 漏洞种类ID。

		VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
		// 漏洞等级。

		VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`
		// 漏洞名称。

		VulName *string `json:"VulName,omitempty" name:"VulName"`
		// 漏洞类型。

		VulType *string `json:"VulType,omitempty" name:"VulType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineRuleTopInfo struct {

	// 事件总数

	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
	// 检测项危害等级

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 检测项id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 基线检测项名

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type DescribeAttackSourceRequest struct {
	*tchttp.BaseRequest

	// 开始日期

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 结束日期

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAttackSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据列表

		List []*PrivilegeEscalationProcess `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivilegeEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsInfoRequest struct {
	*tchttp.BaseRequest

	// 恶意请求-事件Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeRiskDnsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefenceEvent struct {

	// 主机名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 攻击源ip地址所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 事件发生次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 创建事件时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// cve编号

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// 0: 尝试攻击(WeDetect) 1:尝试攻击成功(WeDetect) 2:rasp防御事件

	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`
	// 0 不支持修复，1 支持修复

	FixType *int64 `json:"FixType,omitempty" name:"FixType"`
	// 漏洞事件id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 更新事件时间

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 攻击源ip

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 攻击源端口

	SourcePort []*uint64 `json:"SourcePort,omitempty" name:"SourcePort"`
	// 状态 0: 待处理 1:已防御 2:已处理 3: 已忽略 4: 已删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 0 专业版,1 旗舰版,2 LH普惠版（仅限LH使用）,3  CVM普惠版（仅限CVM使用）

	UpgradeType *int64 `json:"UpgradeType,omitempty" name:"UpgradeType"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 漏洞ID

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type DescribeAssetJarInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Jar包详情

		Jar *AssetJarDetail `json:"Jar,omitempty" name:"Jar"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetJarInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetJarInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulLabelsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetDatabaseDetail struct {

	// 二进制路径

	BinPath *string `json:"BinPath,omitempty" name:"BinPath"`
	// 配置文件路径

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 数据路径

	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`
	// 错误日志路径

	ErrorLogPath *string `json:"ErrorLogPath,omitempty" name:"ErrorLogPath"`
	// 绑定IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 日志文件路径

	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 数据库名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 启动参数

	Param *string `json:"Param,omitempty" name:"Param"`
	// 运行权限

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 插件路径

	PlugInPath *string `json:"PlugInPath,omitempty" name:"PlugInPath"`
	// 监听端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 运行用户

	User *string `json:"User,omitempty" name:"User"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type OpenPort struct {

	// 记录创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 唯一ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 记录更新时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 端口对应进程Pid。

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 开放端口号。

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 端口对应进程名。

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 云镜客户端唯一UUID。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ModifyLogKafkaDeliverTypeRequest struct {
	*tchttp.BaseRequest

	// 安全模块下的日志类型，http://tapd.woa.com/Teneyes/markdown_wikis/show/#1210131751002328905

	LogType []*uint64 `json:"LogType,omitempty" name:"LogType"`
	// 安全模块类型 1: 入侵检测 2: 漏洞管理 3: 基线管理 4: 高级防御 5:客户端相关 6: 资产指纹

	SecurityType *uint64 `json:"SecurityType,omitempty" name:"SecurityType"`
	// 投递开关 0关闭 1开启

	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`
	// kafka topic id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// kafka topic name

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *ModifyLogKafkaDeliverTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogKafkaDeliverTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopAssetScanRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopAssetScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopAssetScanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDatabaseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Databases []*AssetDatabaseBaseInfo `json:"Databases,omitempty" name:"Databases"`
		// 总数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDatabaseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperEventRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 一页多少个 控制返回的uuids条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 控制返回的uuids条数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFileTamperEventRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperEventRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestWebHookRuleRequest struct {
	*tchttp.BaseRequest

	// 测试内容

	Data *string `json:"Data,omitempty" name:"Data"`
	// 规则Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *TestWebHookRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestWebHookRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellPluginInfoRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：Pid精确匹配，MainClass模糊匹配

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeJavaMemShellPluginInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellPluginInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionPrepaidRequest struct {
	*tchttp.BaseRequest

	// 购买相关参数。

	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`
	// 需要开通专业版主机信息数组。

	Machines []*ProVersionMachine `json:"Machines,omitempty" name:"Machines"`
}

func (r *OpenProVersionPrepaidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenProVersionPrepaidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebLocationBaseInfo struct {

	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// Web站点Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	//  附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 主目录

	MainPath *string `json:"MainPath,omitempty" name:"MainPath"`
	// 主目录所有者

	MainPathOwner *string `json:"MainPathOwner,omitempty" name:"MainPathOwner"`
	// 域名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 站点路经数

	PathCount *uint64 `json:"PathCount,omitempty" name:"PathCount"`
	// 拥有者权限

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 站点端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 主机业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 站点协议

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 主机标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 运行用户

	User *string `json:"User,omitempty" name:"User"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DeliverTypeDetails struct {

	// 错误信息

	ErrInfo *string `json:"ErrInfo,omitempty" name:"ErrInfo"`
	// 安全模块下的日志类型，http://tapd.woa.com/Teneyes/markdown_wikis/show/#1210131751002328905

	LogType []*int64 `json:"LogType,omitempty" name:"LogType"`
	// 安全模块类型 1: 入侵检测 2: 漏洞管理 3: 基线管理 4: 高级防御 5:客户端相关 6: 资产指纹

	SecurityType *uint64 `json:"SecurityType,omitempty" name:"SecurityType"`
	// 投递状态，0未开启 1正常 2异常

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 最近一次状态上报时间戳，s

	StatusTime *int64 `json:"StatusTime,omitempty" name:"StatusTime"`
	// 投递开关 0关闭 1开启

	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`
	// kafka topic id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// kafka topic name

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type DescribeAssetDiskListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetDiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// cls日志类型信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditBashRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditBashRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefenceEventDetail struct {

	// 主机名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 攻击源ip地址所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 事件发生次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 创建事件时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// cve编号

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// 漏洞描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 漏洞ID相关的事件详情(json array格式 rasp特有)

	EventDetail *string `json:"EventDetail,omitempty" name:"EventDetail"`
	// 0: 尝试攻击(WeDetect) 1:尝试攻击成功(WeDetect) 2:rasp防御事件

	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`
	// 主机失陷事件进程树(json格式 WeDetect特有)

	ExceptionPstree *string `json:"ExceptionPstree,omitempty" name:"ExceptionPstree"`
	// 修复建议

	Fix *string `json:"Fix,omitempty" name:"Fix"`
	// 漏洞事件id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// ONLINE OFFLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 关联进程主类名

	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`
	// 更新事件时间

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// 攻击payload

	NetworkPayload *string `json:"NetworkPayload,omitempty" name:"NetworkPayload"`
	// 关联进程pid

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 攻击源ip

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 攻击源端口

	SourcePort []*uint64 `json:"SourcePort,omitempty" name:"SourcePort"`
	// 堆栈信息(rasp特有)

	StackTrace *string `json:"StackTrace,omitempty" name:"StackTrace"`
	// 状态 0: 待处理 1:已防御 2:已处理 3: 已忽略 4: 已删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type DescribeProVersionStatusRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端UUID、填写"all"表示所有主机。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeProVersionStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProVersionStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceEventRequest struct {
	*tchttp.BaseRequest

	// 排序列，严格相等：CreateTime创建时间，MergeTime合并时间，Count事件数量

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件：Keywords: ip或者主机名, VulKeywords漏洞名或者CveId模糊查询; Quuid，VulId，EventType，Status精确匹配，CreateBeginTime，CreateEndTime时间段查询

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 数据限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 数据偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序，大小写无关：asc 升序，desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVulDefenceEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemIgnoreListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列 [ID]

	By *string `json:"By,omitempty" name:"By"`
	// <li>CatgoryId - int64 - 是否必填：否 - 规则Id</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求偏移默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
	// 忽略规则ID

	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
}

func (r *DescribeBaselineItemIgnoreListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemIgnoreListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLicenseRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLicenseRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLicenseRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmAttributeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAlarmAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineDownloadListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineDownload `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineDownloadListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineDownloadListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞防御概览信息

		Overview *VulDefenceOverview `json:"Overview,omitempty" name:"Overview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefenceOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FixBaselineDetectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FixBaselineDetectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FixBaselineDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 模糊搜索

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ExportAssetPlanTaskListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>User- string - 是否必填：否 - 用户</li>
	// <li>Status- int - 是否必填：否 - 默认启用状态：0未启用， 1启用 </li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetPlanTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetPlanTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDetectionReportRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 漏洞扫描任务id（不同于出参的导出检测报告的任务Id）

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExportVulDetectionReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDetectionReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetMachineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录列表

		Machines []*AssetMachineBaseInfo `json:"Machines,omitempty" name:"Machines"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetMachineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// java内存马事件详细信息

		Info *JavaMemShellDetail `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJavaMemShellInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineRuleCategoryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineCategory `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineRuleCategoryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleCategoryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 高危命令事件列表

		List []*BashEventNew `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashEventsNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashEventsNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DescribeBanModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 阻断模式，STANDARD_MODE：标准阻断，DEEP_MODE：深度阻断

		Mode *string `json:"Mode,omitempty" name:"Mode"`
		// 标准阻断模式的配置

		StandardModeConfig *StandardModeConfig `json:"StandardModeConfig,omitempty" name:"StandardModeConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBanModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBanModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttackRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBruteAttackRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBruteAttackRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsEventInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求事件详情

		Info *RiskDnsEvent `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsEventInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetAttackSettingRequest struct {
	*tchttp.BaseRequest

	// 新增资产自动包含 0 不包含 1包含

	AutoInclude *uint64 `json:"AutoInclude,omitempty" name:"AutoInclude"`
	// 自选排除的主机

	ExcludeInstanceIds []*string `json:"ExcludeInstanceIds,omitempty" name:"ExcludeInstanceIds"`
	// 自选主机

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 0 新增告警事件默认待处理，1新增告警事件默认已处理，3新增告警事件默认忽略

	NetAttackAlarmStatus *uint64 `json:"NetAttackAlarmStatus,omitempty" name:"NetAttackAlarmStatus"`
	// 0 关闭网络攻击检测，1开启网络攻击检测

	NetAttackEnable *uint64 `json:"NetAttackEnable,omitempty" name:"NetAttackEnable"`
	// 1 全部旗舰版主机，0 Quuids列表主机

	Scope *uint64 `json:"Scope,omitempty" name:"Scope"`
}

func (r *ModifyNetAttackSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetAttackSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UntrustMaliciousRequestResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UntrustMaliciousRequestResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UntrustMaliciousRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoQuaraInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启

		AppidQuaraAll *uint64 `json:"AppidQuaraAll,omitempty" name:"AppidQuaraAll"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoQuaraInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoQuaraInfoResponse) FromJsonString(s string) error {
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

type VulDetailInfo struct {

	// cve编号

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// CVSS详情

	Cvss *string `json:"Cvss,omitempty" name:"Cvss"`
	// CVSS评分

	CvssScore *float64 `json:"CvssScore,omitempty" name:"CvssScore"`
	// 漏洞描述

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// 修复建议

	Fix *string `json:"Fix,omitempty" name:"Fix"`
	// 漏洞级别

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 发布时间

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 参考链接

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞 0= 应急漏洞

	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`
	// 漏洞ID

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

type DescribeSecurityBroadcastInfoRequest struct {
	*tchttp.BaseRequest

	// 文章id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeSecurityBroadcastInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityBroadcastInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportFileTamperRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportFileTamperRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportFileTamperRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetAttackEventInfo struct {

	// 异常行为

	AbnormalAction *string `json:"AbnormalAction,omitempty" name:"AbnormalAction"`
	// 漏洞攻击热度

	AttackLevel *uint64 `json:"AttackLevel,omitempty" name:"AttackLevel"`
	// 漏洞CVE编号

	CVEId *string `json:"CVEId,omitempty" name:"CVEId"`
	// 攻击次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 目标端口

	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`
	// 进程树,需要用base64 解码

	HostOpProcessTree *string `json:"HostOpProcessTree,omitempty" name:"HostOpProcessTree"`
	// 0:无失陷行为 1: rce(命令执行) 2: dnslog 3: writefile

	HostOpType *uint64 `json:"HostOpType,omitempty" name:"HostOpType"`
	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 攻击源地

	Location *string `json:"Location,omitempty" name:"Location"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 攻击发生时间

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// 攻击数据包

	NetPayload *string `json:"NetPayload,omitempty" name:"NetPayload"`
	// 机器付费版本，0 基础版，1专业版，2旗舰版，3普惠版

	PayVersion *uint64 `json:"PayVersion,omitempty" name:"PayVersion"`
	// cvm uuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 攻击源ip

	SrcIP *string `json:"SrcIP,omitempty" name:"SrcIP"`
	// 处理状态，0 待处理 1 已处理 2 已加白  3 已忽略 4 已删除 5: 已开启防御

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 服务进程 base64

	SvcPs *string `json:"SvcPs,omitempty" name:"SvcPs"`
	// 0: 尝试攻击 1:攻击成功

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 漏洞防御状态，0关闭，1开启

	VulDefenceStatus *uint64 `json:"VulDefenceStatus,omitempty" name:"VulDefenceStatus"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
	// 漏洞是否支持防御，0:不支持 1:支持

	VulSupportDefense *uint64 `json:"VulSupportDefense,omitempty" name:"VulSupportDefense"`
}

type OpenProVersionRequest struct {
	*tchttp.BaseRequest

	// 活动ID。

	ActivityId *uint64 `json:"ActivityId,omitempty" name:"ActivityId"`
	// 机器所属地域（当前字段已作废，暂时填string类型空字符串代替，例如：""，英文双引号）

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 云服务器类型（当前字段已作废，暂时填string类型空字符串代替，例如：""，英文双引号）

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 主机唯一标识Uuid数组。
	// 黑石的InstanceId，CVM的Uuid ,边缘计算的Uuid , 轻量应用服务器的Uuid ,混合云机器的Quuid 。 当前参数最大长度限制20

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *OpenProVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenProVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityEventStatRequest struct {
	*tchttp.BaseRequest

	// 该接口无过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSecurityEventStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityEventStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportFileTamperRulesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>OsType - String - 是否必填：否 - windows或linux</li>
	// <li>CpuLoad - Int - 是否必填：否 -
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>
	// <li>DiskLoad - Int - 是否必填：否 -
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>
	// <li>MemLoad - Int - 是否必填：否 -
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportFileTamperRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportFileTamperRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProtectDirInfo struct {

	// 自动恢复开关 (Filters 过滤Quuid 时 返回) 默认0

	AutoRestoreSwitchStatus *uint64 `json:"AutoRestoreSwitchStatus,omitempty" name:"AutoRestoreSwitchStatus"`
	// 网站名称

	DirName *string `json:"DirName,omitempty" name:"DirName"`
	// 网站防护目录地址

	DirPath *string `json:"DirPath,omitempty" name:"DirPath"`
	// 唯一ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 未防护服务器数

	NoProtectServerNum *uint64 `json:"NoProtectServerNum,omitempty" name:"NoProtectServerNum"`
	// 防护异常

	ProtectException *uint64 `json:"ProtectException,omitempty" name:"ProtectException"`
	// 防护服务器数

	ProtectServerNum *uint64 `json:"ProtectServerNum,omitempty" name:"ProtectServerNum"`
	// 防护状态

	ProtectStatus *uint64 `json:"ProtectStatus,omitempty" name:"ProtectStatus"`
	// 关联服务器数

	RelatedServerNum *uint64 `json:"RelatedServerNum,omitempty" name:"RelatedServerNum"`
}

type DescribeLogIndexRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMalwareScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMalwareScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMalwareScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：Keywords: ip或者主机名模糊查询, Type，Status精确匹配，CreateBeginTime，CreateEndTime时间段

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeJavaMemShellListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityTrendsRequest struct {
	*tchttp.BaseRequest

	// 开始时间，如：2021-07-10

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 结束时间，如：2021-07-10

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeSecurityTrendsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityTrendsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRiskProcessEventsRequest struct {
	*tchttp.BaseRequest

	// [StartTime:进程启动时间|DetectTime:最后检测时间]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>HostId - String - 是否必填：否 - 主机ID</li>
	// <li>IpOrName - String - 是否必填：否 - 主机IP或主机名</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名</li>
	// <li>ProcessId - String - 是否必填：否 - 进程ID</li>
	// <li>ProcessPath - String - 是否必填：否 - 进程路径</li>
	// <li>BeginTime - String - 是否必填：否 - 进程启动时间-开始</li>
	// <li>EndTime - String - 是否必填：否 - 进程启动时间-结束</li>
	// <li>Status - String - 是否必填：否 - 状态筛选 0待处理；1查杀中；2已查杀；3已退出;4已信任</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式：[ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportRiskProcessEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRiskProcessEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMaliciousRequestWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMaliciousRequestWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMaliciousRequestWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetScanTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 状态筛选，0:全部，1进行中，2成功，3异常

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeAssetScanTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetScanTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExpertServiceOrderListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单列表

		List []*ExpertServiceOrderInfo `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExpertServiceOrderListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExpertServiceOrderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogHistogramResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 周期内统计结果详情

		Data []*LogHistogram `json:"Data,omitempty" name:"Data"`
		// 统计周期： 单位ms

		Period *int64 `json:"Period,omitempty" name:"Period"`
		// 命中关键字的日志总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogHistogramResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulEmergentMsgInfo struct {

	// 漏洞名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 漏洞纰漏时间

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

type DescribeWeeklyReportsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JavaMemShellPluginInfo struct {

	// 错误日志

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
	// 注入进程主类

	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`
	// 注入进程pid

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 注入状态：0: 注入中, 1: 注入成功, 2: 插件超时, 3: 插件退出, 4: 注入失败 5: 软删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DescribeAssetMachineTagTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Top5标签及数量

		Tags []*AssetKeyVal `json:"Tags,omitempty" name:"Tags"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetMachineTagTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineTagTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulHostTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务器风险top列表

		VulHostTopList []*VulHostTopInfo `json:"VulHostTopList,omitempty" name:"VulHostTopList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulHostTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulHostTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetEnvListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name- string - 是否必填：否 - 环境变量名</li>
	// <li>Type- int - 是否必填：否 - 类型：0用户变量，1系统变量</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetEnvListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetEnvListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebHookRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 规则Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 开启状态[0:开启|1:关闭]

	IsDisabled *int64 `json:"IsDisabled,omitempty" name:"IsDisabled"`
}

func (r *ModifyWebHookRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebHookRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewProVersionRequest struct {
	*tchttp.BaseRequest

	// 购买相关参数。

	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`
	// 主机唯一ID，对应CVM的uuid、BM的InstanceId。

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *RenewProVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewProVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBashEventsStatusRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 新状态(0-待处理 1-高危 2-正常)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetBashEventsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBashEventsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogInfo struct {

	// 日志内容的Json序列化字符串

	Content *string `json:"Content,omitempty" name:"Content"`
	// 日志文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 日志来源IP

	Source *string `json:"Source,omitempty" name:"Source"`
	// 日志时间，单位ms

	TimeStamp *int64 `json:"TimeStamp,omitempty" name:"TimeStamp"`
}

type CancelIgnoreVulResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelIgnoreVulResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelIgnoreVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeStrategyEnableStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeStrategyEnableStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeStrategyEnableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineClearHistoryRequest struct {
	*tchttp.BaseRequest

	// AgentLastOfflineTime 客户端最后离线时间
	// AutoClearTime 清理时间

	By *string `json:"By,omitempty" name:"By"`
	// 筛选条件
	// 多个条件筛选时 Keywords,TimeBetween,取交集
	// <li> Keywords 实例名称/内网/公网IP</li>
	// <li> TimeBetween 时间区间</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10 ,最大100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 正序, ASC , 倒序 DESC , 默认ASC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeMachineClearHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineClearHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserKeyListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 账号名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetUserKeyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserKeyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetPlanTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetPlanTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetPlanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceApiInfo struct {

	// 123

	ApiList *int64 `json:"ApiList,omitempty" name:"ApiList"`
	// 服务介绍文档链接

	ArnDocument *string `json:"ArnDocument,omitempty" name:"ArnDocument"`
	// 条件规则列表

	ConditionKeyList []*string `json:"ConditionKeyList,omitempty" name:"ConditionKeyList"`
	// 服务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务ID

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type DescribeProtectDirListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防护目录列表信息

		List []*ProtectDirInfo `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProtectDirListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectDirListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBuyBindTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBuyBindTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBuyBindTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 攻击目标IP

		DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
		// 攻击目标端口

		DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`
		// 攻击路径

		HttpCgi *string `json:"HttpCgi,omitempty" name:"HttpCgi"`
		// 攻击内容

		HttpContent *string `json:"HttpContent,omitempty" name:"HttpContent"`
		// 攻击头信息

		HttpHead *string `json:"HttpHead,omitempty" name:"HttpHead"`
		// 攻击目标主机

		HttpHost *string `json:"HttpHost,omitempty" name:"HttpHost"`
		// 攻击方法

		HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
		// 攻击参数

		HttpParam *string `json:"HttpParam,omitempty" name:"HttpParam"`
		// 请求源

		HttpReferer *string `json:"HttpReferer,omitempty" name:"HttpReferer"`
		// 攻击者浏览器标识

		HttpUserAgent *string `json:"HttpUserAgent,omitempty" name:"HttpUserAgent"`
		// 日志ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 主机ID

		Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
		// 攻击来源IP

		SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
		// 攻击来源端口

		SrcPort *uint64 `json:"SrcPort,omitempty" name:"SrcPort"`
		// 威胁类型

		VulType *string `json:"VulType,omitempty" name:"VulType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackLogInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackLogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全基线数。

		BaseLineNum *uint64 `json:"BaseLineNum,omitempty" name:"BaseLineNum"`
		// 暴力破解成功数。

		BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitempty" name:"BruteAttackSuccessNum"`
		// 木马文件数。

		MalwareNum *uint64 `json:"MalwareNum,omitempty" name:"MalwareNum"`
		// 异地登录数。

		NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitempty" name:"NonlocalLoginNum"`
		// 服务器在线数。

		OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitempty" name:"OnlineMachineNum"`
		// 专业服务器数。

		ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitempty" name:"ProVersionMachineNum"`
		// 漏洞数。

		VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBaselinePolicyRequest struct {
	*tchttp.BaseRequest

	// 无

	Data *BaselinePolicy `json:"Data,omitempty" name:"Data"`
	// <li>RuleName - String - 是否必填：否 - 规则名称</li>
	// <li>CategoryId - int64 - 是否必填：否 自定义筛选为-1 - 规则分类</li>
	// <li>RuleType - int - 是否必填：否 0:系统 1:自定义 - 规则类型</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 是否按照过滤的全选

	SelectAll *int64 `json:"SelectAll,omitempty" name:"SelectAll"`
}

func (r *ModifyBaselinePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBaselinePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAESKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AES IV

		AESIV *string `json:"AESIV,omitempty" name:"AESIV"`
		// AES key

		AESKey *string `json:"AESKey,omitempty" name:"AESKey"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAESKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAESKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebServiceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各Web服务数量

		WebServices []*AssetKeyVal `json:"WebServices,omitempty" name:"WebServices"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebServiceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulHostCountScanTimeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulHostCountScanTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulHostCountScanTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulEffectHostList struct {

	// 主机别名

	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`
	// 云标签信息

	CloudTags []*Tags `json:"CloudTags,omitempty" name:"CloudTags"`
	// 说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件id

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 首次发现时间

	FirstDiscoveryTime *string `json:"FirstDiscoveryTime,omitempty" name:"FirstDiscoveryTime"`
	// 失败原因

	FixStatusMsg *string `json:"FixStatusMsg,omitempty" name:"FixStatusMsg"`
	// 主机HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 版本信息：0-基础版 1-专业版 2-旗舰版 3-普惠版

	HostVersion *uint64 `json:"HostVersion,omitempty" name:"HostVersion"`
	// 实例状态："PENDING"-创建中 "LAUNCH_FAILED"-创建失败 "RUNNING"-运行中 "STOPPED"-关机 "STARTING"-表示开机中 "STOPPING"-表示关机中 "REBOOTING"-重启中 "SHUTDOWN"-表示停止待销毁 "TERMINATING"-表示销毁中 "

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 是否能自动修复 0 :漏洞不可自动修复，  1：可自动修复， 2：客户端已离线， 3：主机不是旗舰版只能手动修复， 4：机型不允许 ，5：修复中 ，6：已修复， 7：检测中  9:修复失败，10:已忽略 11:漏洞只支持linux不支持Windows 12：漏洞只支持Windows不支持linux，13:修复失败但此时主机已离线，14:修复失败但此时主机不是旗舰版， 15:已手动修复

	IsSupportAutoFix *uint64 `json:"IsSupportAutoFix,omitempty" name:"IsSupportAutoFix"`
	// 最后检测时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 危害等级：1-低危；2-中危；3-高危；4-严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 外网ip

	PublicIpAddresses *string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 状态：0: 待处理 1:忽略  3:已修复  5:检测中 6:修复中 7: 回滚中 8:修复失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 主机标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetAppListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime|ProcessCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>AppName- string - 是否必填：否 - 应用名搜索</li>
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Type - int - 是否必填：否 - 类型	: 仅linux
	// 0: 全部
	// 1: 运维
	// 2 : 数据库
	// 3 : 安全
	// 4 : 可疑应用
	// 5 : 系统架构
	// 6 : 系统应用
	// 7 : WEB服务
	// 99:其他</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetAppListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBashPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBashPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBashPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWarningSettingRequest struct {
	*tchttp.BaseRequest

	// 告警设置的修改内容

	WarningObjects []*WarningObject `json:"WarningObjects,omitempty" name:"WarningObjects"`
}

func (r *ModifyWarningSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWarningSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareRiskWarningResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启自动扫描：true-开启，false-未开启

		IsCheckRisk *bool `json:"IsCheckRisk,omitempty" name:"IsCheckRisk"`
		// 是否弹出提示 true 弹出, false不弹

		IsPop *bool `json:"IsPop,omitempty" name:"IsPop"`
		// 风险文件列表信息

		List []*MalwareRisk `json:"List,omitempty" name:"List"`
		// 异常进程列表信息

		ProcessList []*MalwareRisk `json:"ProcessList,omitempty" name:"ProcessList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwareRiskWarningResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareRiskWarningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRansomDefenseStrategyListRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime、MachineCount

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Ips - String - 是否必填：否 - 通过ip查询 </li>
	// <li>MachineNames - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>Names - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>Dirs - String - 是否必填：否 - 诱饵目录 </li>
	// <li>Status - String - 是否必填：否 - 策略状态：0关闭，1开启 </li>
	// <li>BackupType - String - 是否必填：否 - 备份模式：0-按周;1-按天 </li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportRansomDefenseStrategyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRansomDefenseStrategyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditPrivilegeRulesRequest struct {
	*tchttp.BaseRequest

	// 事件列表和详情点击加白时关联的事件id (新增规则时请留空)

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 规则ID(新增时请留空)

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则(默认否)

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 是否S权限进程

	SMode *uint64 `json:"SMode,omitempty" name:"SMode"`
	// 客户端ID数组

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *EditPrivilegeRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditPrivilegeRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ABTestConfig struct {

	// 灰度项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// true：正在灰度，false：不在灰度

	Status *bool `json:"Status,omitempty" name:"Status"`
}

type AssetFilters struct {

	// 是否模糊查询

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ModifyRiskDnsPolicyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRiskDnsPolicyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskDnsPolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineRuleDetectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBaselineRuleDetectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineRuleDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulFixStatusSnapshotInfo struct {

	// 快照创建失败原因

	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 记录唯一id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 快照创建时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// cvm id

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 快照id

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 快照名称

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	// 快照状态 0-初始状态1-快照创建成功；2-快照创建失败；

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type MachineSnapshotInfo struct {

	// 快照创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 磁盘id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// cvm id

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 地区id

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 快照id

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 快照名称

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
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

type DescribeProcessTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态。
		// <li>COMPLETE：完成（此时可以调用DescribeProcesses接口获取实时进程列表）</li>
		// <li>AGENT_OFFLINE：云镜客户端离线</li>
		// <li>COLLECTING：进程获取中</li>
		// <li>FAILED：进程获取失败</li>

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcessTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProcessTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineWeakPasswordListRequest struct {
	*tchttp.BaseRequest

	// 0:过滤的结果导出；1:全部导出

	ExportAll *int64 `json:"ExportAll,omitempty" name:"ExportAll"`
	// <li>WeakPassword - string - 是否必填：否 - 弱口令</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBaselineWeakPasswordListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineWeakPasswordListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFileTamperRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFileTamperRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFileTamperRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMalwareWhiteListRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteMalwareWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMalwareWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetProcessInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Process []*AssetProcessBaseInfo `json:"Process,omitempty" name:"Process"`
		// 记录总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetProcessInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBaselineWeakPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBaselineWeakPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBaselineWeakPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScanTaskRequest struct {
	*tchttp.BaseRequest

	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 自选服务器时生效，主机quuid的string数组

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetScanTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败的主机数量

		FailCount *uint64 `json:"FailCount,omitempty" name:"FailCount"`
		// 扫描的主机总数

		HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`
		// 主机任务详情列表

		Hosts *AssetScanTaskHostDetail `json:"Hosts,omitempty" name:"Hosts"`
		// 成功的主机数量

		SuccCount *uint64 `json:"SuccCount,omitempty" name:"SuccCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetScanTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetScanTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全部漏洞统计

		AllVul *VulOverview `json:"AllVul,omitempty" name:"AllVul"`
		// 影响主机统计

		EffectHost *VulOverview `json:"EffectHost,omitempty" name:"EffectHost"`
		// 重点关注漏洞统计

		FollowVul *VulOverview `json:"FollowVul,omitempty" name:"FollowVul"`
		// 漏洞攻击事件（近1月）统计

		VulAttackEvent *VulOverview `json:"VulAttackEvent,omitempty" name:"VulAttackEvent"`
		// 已防御攻击（近1月）统计

		VulDefenceEvent *VulOverview `json:"VulDefenceEvent,omitempty" name:"VulDefenceEvent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载地址（已弃用）

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出文件Id 可通过ExportTasks接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBaselineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJavaMemShellsStatusRequest struct {
	*tchttp.BaseRequest

	// 事件Id数组

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 目标处理状态： 0 - 待处理 1 - 已加白 2 - 已删除 3 - 已忽略 4 - 已手动处理

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyJavaMemShellsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJavaMemShellsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JavaMemShellEvent struct {

	// 服务器名称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 注释

	Annotations *string `json:"Annotations,omitempty" name:"Annotations"`
	// 进程命令行参数

	Args *string `json:"Args,omitempty" name:"Args"`
	// 类名

	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`
	// 首次发现时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件描述

	EventDescription *string `json:"EventDescription,omitempty" name:"EventDescription"`
	// 进程路径

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 服务器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 继承的接口

	Interfaces *string `json:"Interfaces,omitempty" name:"Interfaces"`
	// 类文件MD5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 进程pid

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 最近检测时间

	RecentFoundTime *string `json:"RecentFoundTime,omitempty" name:"RecentFoundTime"`
	// 建议方案

	SecurityAdvice *string `json:"SecurityAdvice,omitempty" name:"SecurityAdvice"`
	// 父类名

	SuperClassName *string `json:"SuperClassName,omitempty" name:"SuperClassName"`
	// 内存马类型  0:Filter型 1:Listener型 2:Servlet型 3:Interceptors型 4:Agent型 5:其他

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type ScanTaskDetails struct {

	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 失败详情

	FailType *uint64 `json:"FailType,omitempty" name:"FailType"`
	// 服务器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 服务器名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// id唯一

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 外网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 操作系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 唯一Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 风险数量

	RiskNum *uint64 `json:"RiskNum,omitempty" name:"RiskNum"`
	// 扫描开始时间

	ScanBeginTime *string `json:"ScanBeginTime,omitempty" name:"ScanBeginTime"`
	// 扫描结束时间

	ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`
	// 状态码

	Status *string `json:"Status,omitempty" name:"Status"`
	// 唯一Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAttackEventInfoRequest struct {
	*tchttp.BaseRequest

	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAttackEventInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackEventInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareWhiteListAffectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白名单规则影响事件列表

		AffectList []*MalwareWhiteListAffectEvent `json:"AffectList,omitempty" name:"AffectList"`
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwareWhiteListAffectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareWhiteListAffectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板列表

		List []*SearchTemplate `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
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

type DescribeRansomDefenseMachineListRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime、LastBackupTime、BackupCount

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Ips - String - 是否必填：否 - 通过ip查询 </li>
	// <li>MachineNames - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>Names - String - 是否必填：否 - 通过策略名查询 </li>
	// <li>Status - String - 是否必填：否 - 策略状态：0备份中，1备份成功，2备份失败, 9暂无备份 </li>
	// <li>LastBackupStatus - String - 是否必填：否 - 上次备份状态：0备份中，1备份成功，2备份失败, 9暂无备份 </li>
	// <li>LastBackupTimeBegin - String - 是否必填：否 - 最近一次备份时间开始</li>
	// <li>LastBackupTimeEnd - String - 是否必填：否 - 最近一次备份时间结束</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRansomDefenseMachineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseMachineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulEffectModuleInfo struct {

	// 修复命令

	FixCmd *string `json:"FixCmd,omitempty" name:"FixCmd"`
	// 组件名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 影响的主机quuid

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 组件影响版本

	Rule *string `json:"Rule,omitempty" name:"Rule"`
	// 影响的主机uuid

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
	// 组件版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeAssetMachineTagTopRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetMachineTagTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineTagTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAttackStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetMachineInfoRequest struct {
	*tchttp.BaseRequest

	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetMachineInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmergencyVulScanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEmergencyVulScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEmergencyVulScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExpertServiceTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExpertServiceTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExpertServiceTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表内容

		List []*PrivilegeRule `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivilegeRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivilegeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulInfoRequest struct {
	*tchttp.BaseRequest

	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *ExportVulInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWarningHostConfigRequest struct {
	*tchttp.BaseRequest

	// 告警主机范围类型，0:全部主机，1:按所属项目选，2:按腾讯云标签选，3:按主机安全标签选，4:自选主机

	HostRange *int64 `json:"HostRange,omitempty" name:"HostRange"`
	// 项目或标签的id列表，自选主机时为空

	ItemLabelIds []*string `json:"ItemLabelIds,omitempty" name:"ItemLabelIds"`
	// 项目或标签的名称列表，自选主机时为空

	ItemLabels []*string `json:"ItemLabels,omitempty" name:"ItemLabels"`
	// 机器列表

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 告警类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyWarningHostConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWarningHostConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FullMachine struct {

	// 是否是专业版
	// * true：是
	// * false：否

	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP。

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 主机唯一标识Uuid。

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type DescribeAssetWebServiceCountRequest struct {
	*tchttp.BaseRequest

	// 搜索条件：返回名称包含Name的所有Web服务列表

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetWebServiceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineBasicInfoRequest struct {
	*tchttp.BaseRequest

	// 基线名称

	BaselineName *string `json:"BaselineName,omitempty" name:"BaselineName"`
}

func (r *DescribeBaselineBasicInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineBasicInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityEventInfo struct {

	// 安全事件数

	EventCnt *uint64 `json:"EventCnt,omitempty" name:"EventCnt"`
	// 受影响机器数

	UuidCnt *uint64 `json:"UuidCnt,omitempty" name:"UuidCnt"`
}

type DescribeDefenceEventDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞事件详细信息

		Data *VulDefenceEventDetail `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefenceEventDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenceEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineClearHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表数据

		List []*MachineClearHistory `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineClearHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineClearHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRansomDefenseStrategyMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRansomDefenseStrategyMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRansomDefenseStrategyMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebHookRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWebHookRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebHookRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWarningHostConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警主机范围类型，0:全部主机，1:按所属项目选，2:按腾讯云标签选，3:按主机安全标签选，4:自选主机

		HostRange *int64 `json:"HostRange,omitempty" name:"HostRange"`
		// 项目或标签的id列表，自选主机时为空

		ItemLabelIds []*string `json:"ItemLabelIds,omitempty" name:"ItemLabelIds"`
		// 项目或标签的名称列表，自选主机时为空

		ItemLabels []*string `json:"ItemLabels,omitempty" name:"ItemLabels"`
		// 机器列表

		Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
		// 机器列表总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWarningHostConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarningHostConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportJavaMemShellsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务Id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportJavaMemShellsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportJavaMemShellsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttackSourceEdge struct {

	// 出发节点

	From *string `json:"From,omitempty" name:"From"`
	// 目标节点

	To *string `json:"To,omitempty" name:"To"`
}

type DescribeAssetWebAppCountRequest struct {
	*tchttp.BaseRequest

	// 搜索条件：返回名称包含Name的所有Web应用列表

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetWebAppCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRequest struct {
	*tchttp.BaseRequest

	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时

	Context *string `json:"Context,omitempty" name:"Context"`
	// 表示单次查询返回的原始日志条数，最大值为1000，获取后续日志需使用Context参数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 要检索分析的日志的结束时间，Unix时间戳（毫秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 检索分析语句，最大长度为12KB

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 原始日志是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 要检索分析的日志的起始时间，Unix时间戳（毫秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *SearchLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportSecurityTrendsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportSecurityTrendsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSecurityTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineHostDetect struct {

	// 0:未通过 1:忽略 3:通过 5:检测中

	DetectStatus *int64 `json:"DetectStatus,omitempty" name:"DetectStatus"`
	// 首次检测时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 主机Id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 内网Ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 关联检测项数

	ItemCount *int64 `json:"ItemCount,omitempty" name:"ItemCount"`
	// 最后检测时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 检测未通过数

	NotPassedItemCount *int64 `json:"NotPassedItemCount,omitempty" name:"NotPassedItemCount"`
	// 检测通过数

	PassedItemCount *int64 `json:"PassedItemCount,omitempty" name:"PassedItemCount"`
	// 主机安全UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 外网Ip

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}

type DescribeAssetAppCountRequest struct {
	*tchttp.BaseRequest

	// 搜索条件：搜索返回所有软件名包含Name的进程列表

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetAppCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogKafkaDeliverInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接入地址

		AccessAddr *string `json:"AccessAddr,omitempty" name:"AccessAddr"`
		// 接入方式，1公网域名接入，2支撑环境接入

		AccessType *uint64 `json:"AccessType,omitempty" name:"AccessType"`
		// 可用区

		Az *string `json:"Az,omitempty" name:"Az"`
		// 峰值带宽

		BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`
		// 投递状态，1：健康，2：告警，3：异常

		DeliverStatus *uint64 `json:"DeliverStatus,omitempty" name:"DeliverStatus"`
		// xx

		DeliverTypeDetails []*DeliverTypeDetails `json:"DeliverTypeDetails,omitempty" name:"DeliverTypeDetails"`
		// 磁盘容量

		DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
		// kafka版本

		InsVersion *string `json:"InsVersion,omitempty" name:"InsVersion"`
		// 实例环境

		KafkaEnvName *string `json:"KafkaEnvName,omitempty" name:"KafkaEnvName"`
		// 实例id

		KafkaId *string `json:"KafkaId,omitempty" name:"KafkaId"`
		// 所在子网

		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
		// 用户名

		Username *string `json:"Username,omitempty" name:"Username"`
		// 所属网络

		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
		// 地域

		Zone *string `json:"Zone,omitempty" name:"Zone"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogKafkaDeliverInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogKafkaDeliverInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityEventsCntRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecurityEventsCntRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityEventsCntRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportIgnoreBaselineRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务Id , 可通过ExportTasks 接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportIgnoreBaselineRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportIgnoreBaselineRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetAttackEvent struct {

	// 攻击次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 目标端口

	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`
	// 日志ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 来源地

	Location *string `json:"Location,omitempty" name:"Location"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 攻击时间

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// 是否今日新增主机

	New *bool `json:"New,omitempty" name:"New"`
	// 机器付费版本，0 基础版，1专业版，2旗舰版，3普惠版

	PayVersion *uint64 `json:"PayVersion,omitempty" name:"PayVersion"`
	// cvm uuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 来源IP

	SrcIP *string `json:"SrcIP,omitempty" name:"SrcIP"`
	// 处理状态，0 待处理 1 已处理 2 已加白  3 已忽略 4 已删除 5: 已开启防御

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 攻击状态，0: 尝试攻击 1: 实锤攻击(攻击成功)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 客户端ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 是否开启漏洞防御，0关1开

	VulDefenceStatus *uint64 `json:"VulDefenceStatus,omitempty" name:"VulDefenceStatus"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
	// 漏洞是否支持防御，0:不支持 1:支持

	VulSupportDefense *uint64 `json:"VulSupportDefense,omitempty" name:"VulSupportDefense"`
}

type DeleteBaselineRuleIgnoreRequest struct {
	*tchttp.BaseRequest

	// 规则Id

	RuleIds []*int64 `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *DeleteBaselineRuleIgnoreRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBaselineRuleIgnoreRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表

		List []*JavaMemShellInfo `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJavaMemShellListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAutoQuaraStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAutoQuaraStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAutoQuaraStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务扫描状态列表

		State *TaskStatus `json:"State,omitempty" name:"State"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BruteAttack struct {

	// 阻断状态。

	BanStatus *string `json:"BanStatus,omitempty" name:"BanStatus"`
	// 城市ID。

	City *uint64 `json:"City,omitempty" name:"City"`
	// 尝试破解次数。

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 国家ID。

	Country *uint64 `json:"Country,omitempty" name:"Country"`
	// 发生时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 事件ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否专业版。

	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 省份ID。

	Province *uint64 `json:"Province,omitempty" name:"Province"`
	// 机器UUID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 来源IP。

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 破解事件状态
	// <li>BRUTEATTACK_FAIL_ACCOUNT： 暴力破解事件-失败(存在帐号)  </li>
	// <li>BRUTEATTACK_FAIL_NOACCOUNT：暴力破解事件-失败(帐号不存在)</li>
	// <li>BRUTEATTACK_SUCCESS：暴力破解事件-成功</li>

	Status *string `json:"Status,omitempty" name:"Status"`
	// 用户名称。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 云镜客户端唯一标识UUID。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ScreenProtection struct {

	// 类型值：文件查杀，暴力破解，漏洞扫描，基线检测

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件查杀:  0:从未检测过，或0资产付费情况, 1:已检测，存在恶意文件, 2:已检测，未开启隔离防护, 3:已检测且已开启防护且无风险；
	// 暴力破解: 0:未开启防护（0付费资产情况）1:已开启自动阻断；
	// 漏洞扫描: 0:从未检测过，或0资产付费情况, 1:存在漏洞风险, 2:无风险；
	// 基线检测: 0:从未检测过，或0资产付费情况, 1:存在基线风险,2:无风险；

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type CreateUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 登录地域信息数组。

	Places []*Place `json:"Places,omitempty" name:"Places"`
	// 云镜客户端UUID数组。

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *CreateUsualLoginPlacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUsualLoginPlacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralStatRequest struct {
	*tchttp.BaseRequest

	// 机器所属地域。如：ap-guangzhou，ap-shanghai

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 云主机类型。
	// <li>CVM：表示腾讯云服务器</li>
	// <li>BM:  表示黑石物理机</li>
	// <li>ECM:  表示边缘计算服务器</li>
	// <li>LH:  表示轻量应用服务器</li>
	// <li>Other:  表示混合云机器</li>

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

func (r *DescribeGeneralStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousRequestsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | TRUSTED：已信任 | UN_TRUSTED：已取消信任）</li>
	// <li>Domain - String - 是否必填：否 - 恶意请求的域名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 云镜客户端唯一UUID。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMaliciousRequestsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousRequestsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseProVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseProVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞防御列表

		List []*VulDefenceRangeDetail `json:"List,omitempty" name:"List"`
		// 数据总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefenceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebHookRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则内容

		List []*WebHookRuleSummary `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebHookRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebHookRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefencePluginExceptionCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前异常插件数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefencePluginExceptionCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefencePluginExceptionCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScreenProtectionCnt struct {

	// 总数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// cloud：云查杀引擎，detect：检测引擎，defend：攻击防御，threat：威胁情报，analysis：异常分析，ai：AI引擎

	Name *string `json:"Name,omitempty" name:"Name"`
	// cloud：云查杀引擎，detect：检测引擎，defend：攻击防御，threat：威胁情报，analysis：异常分析，ai：AI引擎

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeJavaMemShellInfoRequest struct {
	*tchttp.BaseRequest

	// 事件Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeJavaMemShellInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityEventsCntResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击检测相关风险事件

		AttackLogs *SecurityEventInfo `json:"AttackLogs,omitempty" name:"AttackLogs"`
		// 安全基线相关风险事件

		BaseLine *SecurityEventInfo `json:"BaseLine,omitempty" name:"BaseLine"`
		// 高危命令相关风险事件

		Bash *SecurityEventInfo `json:"Bash,omitempty" name:"Bash"`
		// 密码破解相关风险事件

		BruteAttack *SecurityEventInfo `json:"BruteAttack,omitempty" name:"BruteAttack"`
		// 受影响机器数

		EffectMachineCount *uint64 `json:"EffectMachineCount,omitempty" name:"EffectMachineCount"`
		// 应急漏洞相关风险事件

		EmergencyVul *SecurityEventInfo `json:"EmergencyVul,omitempty" name:"EmergencyVul"`
		// 所有事件总数

		EventsCount *uint64 `json:"EventsCount,omitempty" name:"EventsCount"`
		// 登录审计相关风险事件

		HostLogin *SecurityEventInfo `json:"HostLogin,omitempty" name:"HostLogin"`
		// linux系统漏洞事件总数

		LinuxVul *SecurityEventInfo `json:"LinuxVul,omitempty" name:"LinuxVul"`
		// 木马文件相关风险事件

		Malware *SecurityEventInfo `json:"Malware,omitempty" name:"Malware"`
		// 本地提权相关风险事件

		PrivilegeRules *SecurityEventInfo `json:"PrivilegeRules,omitempty" name:"PrivilegeRules"`
		// 反弹Shell相关风险事件

		ReverseShell *SecurityEventInfo `json:"ReverseShell,omitempty" name:"ReverseShell"`
		// 恶意请求相关风险事件

		RiskDns *SecurityEventInfo `json:"RiskDns,omitempty" name:"RiskDns"`
		// 应用漏洞风险事件

		SysVul *SecurityEventInfo `json:"SysVul,omitempty" name:"SysVul"`
		// Web应用漏洞相关风险事件

		WebVul *SecurityEventInfo `json:"WebVul,omitempty" name:"WebVul"`
		// window 系统漏洞事件总数

		WindowVul *SecurityEventInfo `json:"WindowVul,omitempty" name:"WindowVul"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityEventsCntResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityEventsCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSaveOrUpdateWarningsRequest struct {
	*tchttp.BaseRequest

	// 告警设置的修改内容

	WarningObjects []*WarningObject `json:"WarningObjects,omitempty" name:"WarningObjects"`
}

func (r *DescribeSaveOrUpdateWarningsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSaveOrUpdateWarningsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackTopRequest struct {
	*tchttp.BaseRequest

	//  过滤条件。
	// <li>BeginTime - String 起始时间,默认近7天- 是否必填: 否</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAttackTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStrategyExistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略是否存在, 1是 0否

		IfExist *uint64 `json:"IfExist,omitempty" name:"IfExist"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStrategyExistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStrategyExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>VulCategory - int - 是否必填：否 - 漏洞分类筛选1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞</li>
	// <li>IfEmergency - String - 是否必填：否 - 是否为应急漏洞，查询应急漏洞传:yes</li>
	// <li>Status - String - 是否必填：是 - 漏洞状态筛选，0: 待处理 1:忽略  3:已修复  5:检测中， 控制台仅处理0,1,3,5四种状态</li>
	// <li>Level - String - 是否必填：否 - 漏洞等级筛选 1:低 2:中 3:高 4:提示</li>
	// <li>VulName- String - 是否必填：否 - 漏洞名称搜索</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 是否导出详情,1是 0不是

	IfDetail *uint64 `json:"IfDetail,omitempty" name:"IfDetail"`
}

func (r *ExportVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyValueInfo struct {

	// 是否大小写敏感

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 需要建立索引的键值对信息

	KeyValues []*KeyValueArrayInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type Machine struct {

	// 基线风险数。

	BaselineNum *int64 `json:"BaselineNum,omitempty" name:"BaselineNum"`
	// 云标签信息

	CloudTags []*Tags `json:"CloudTags,omitempty" name:"CloudTags"`
	// 网络风险数。

	CyberAttackNum *int64 `json:"CyberAttackNum,omitempty" name:"CyberAttackNum"`
	// 是否有资产扫描接口，0无，1有

	HasAssetScan *uint64 `json:"HasAssetScan,omitempty" name:"HasAssetScan"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例状态 TERMINATED_PRO_VERSION 已销毁

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 入侵事件数

	InvasionNum *int64 `json:"InvasionNum,omitempty" name:"InvasionNum"`
	// 主机ip列表

	IpList *string `json:"IpList,omitempty" name:"IpList"`
	// 是否15天内新增的主机 0：非15天内新增的主机，1：15天内增加的主机

	IsAddedOnTheFifteen *uint64 `json:"IsAddedOnTheFifteen,omitempty" name:"IsAddedOnTheFifteen"`
	// 是否是专业版。
	// <li>true： 是</li>
	// <li>false：否</li>

	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 内核版本

	KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`
	// 防篡改 授权状态 1 授权 0 未授权

	LicenseStatus *uint64 `json:"LicenseStatus,omitempty" name:"LicenseStatus"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机系统。

	MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`
	// 主机状态。
	// <li>OFFLINE: 离线  </li>
	// <li>ONLINE: 在线</li>
	// <li>SHUTDOWN: 已关机</li>
	// <li>UNINSTALLED: 未防护</li>

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 机器所属专区类型 CVM 云服务器, BM 黑石, ECM 边缘计算, LH 轻量应用服务器 ,Other 混合云专区

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 主机外网IP。

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 木马数。

	MalwareNum *int64 `json:"MalwareNum,omitempty" name:"MalwareNum"`
	// 主机状态。
	// <li>POSTPAY: 表示后付费，即按量计费  </li>
	// <li>PREPAY: 表示预付费，即包年包月</li>

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 防护版本：BASIC_VERSION 基础版， PRO_VERSION 专业版，Flagship 旗舰版，GENERAL_DISCOUNT 普惠版

	ProtectType *string `json:"ProtectType,omitempty" name:"ProtectType"`
	// CVM或BM机器唯一Uuid。

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 地域信息

	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 风险状态。
	// <li>SAFE：安全</li>
	// <li>RISK：风险</li>
	// <li>UNKNOWN：未知</li>

	SecurityStatus *string `json:"SecurityStatus,omitempty" name:"SecurityStatus"`
	// 标签信息

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 云镜客户端唯一Uuid，若客户端长时间不在线将返回空字符。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 所属网络

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 漏洞数。

	VulNum *int64 `json:"VulNum,omitempty" name:"VulNum"`
}

type AddLoginWhiteListsRequest struct {
	*tchttp.BaseRequest

	// 异地登录事件ID，当ProcessType为Id时此项必填

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 异地登录白名单实体

	HostLoginWhiteObj *HostLoginWhiteObj `json:"HostLoginWhiteObj,omitempty" name:"HostLoginWhiteObj"`
	// 事件同步处理方式：
	//   "" -- 不操作
	//   "All" -- 将符合此配置的所有事件记录加白
	//   "Id" -- 将EventId对应的事件记录加白

	ProcessType *string `json:"ProcessType,omitempty" name:"ProcessType"`
}

func (r *AddLoginWhiteListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLoginWhiteListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeEventsRequest struct {
	*tchttp.BaseRequest

	// ID数组. (最大100条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeletePrivilegeEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivilegeEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulLevelCountRequest struct {
	*tchttp.BaseRequest

	// 是否仅统计重点关注漏洞 1=仅统计重点关注漏洞, 0=统计全部漏洞

	IsFollowVul *uint64 `json:"IsFollowVul,omitempty" name:"IsFollowVul"`
	// 1:web-cms 漏洞，2.应用漏洞 3:安全基线 4: Linux软件漏洞 5: windows系统漏洞 6:应急漏洞，不填或者填0时返回 1，2，4，5 的总统计数据

	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`
}

func (r *DescribeVulLevelCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulLevelCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportPrivilegeEventsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportPrivilegeEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportPrivilegeEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBaselineRuleRequest struct {
	*tchttp.BaseRequest

	// 无

	Data *BaselineRule `json:"Data,omitempty" name:"Data"`
	// <li>ItemName - string - 是否必填：否 - 项名称</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 是否过滤全选

	SelectAll *int64 `json:"SelectAll,omitempty" name:"SelectAll"`
}

func (r *ModifyBaselineRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBaselineRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUsualLoginPlacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebLocationPath struct {

	// 文件所属组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 文件权限

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 物理路径

	RealPath *string `json:"RealPath,omitempty" name:"RealPath"`
	// 文件所有者

	User *string `json:"User,omitempty" name:"User"`
	// 虚拟路径

	VirtualPath *string `json:"VirtualPath,omitempty" name:"VirtualPath"`
}

type DescribeAssetJarInfoRequest struct {
	*tchttp.BaseRequest

	// Jar包ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetJarInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetJarInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVdbAndPocInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞库更新时间。

		PocUpdateTime *string `json:"PocUpdateTime,omitempty" name:"PocUpdateTime"`
		// 病毒库更新时间。

		VdbUpdateTime *string `json:"VdbUpdateTime,omitempty" name:"VdbUpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVdbAndPocInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVdbAndPocInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 控制返回的uuids 数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 控制返回的uuids 数量，起始位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFileTamperRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateHostLoginWhiteObj struct {

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 要更新的数据id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 地域信息数组

	Places []*Place `json:"Places,omitempty" name:"Places"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 来源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type VulTopInfo struct {

	// 漏洞数量

	VulCount *uint64 `json:"VulCount,omitempty" name:"VulCount"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 危害等级：1-低危；2-中危；3-高危；4-严重

	VulLevel *uint64 `json:"VulLevel,omitempty" name:"VulLevel"`
	// 漏洞 名

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type DescribeBaselineEffectHostListRequest struct {
	*tchttp.BaseRequest

	// 基线id

	BaselineId *uint64 `json:"BaselineId,omitempty" name:"BaselineId"`
	// 过滤条件。
	// <li>AliasName- String- 主机别名</li>
	// <li>Status- Uint- 1已通过  0未通过 5检测中</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 主机uuid数组

	UuidList []*string `json:"UuidList,omitempty" name:"UuidList"`
}

func (r *DescribeBaselineEffectHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineEffectHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表

		Machines []*Machine `json:"Machines,omitempty" name:"Machines"`
		// 主机数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetJarListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>Type- uint - 是否必填：否 - 类型
	// 1: 应用程序
	// 2 : 系统类库
	// 3 : Web服务自带库
	// 4 : 其他依赖包</li>
	// <li>Status- string - 是否必填：否 - 是否可执行：0否，1是</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetJarListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetJarListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetRecentMachineInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetRecentMachineInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetRecentMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RansomDefenseRollbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RansomDefenseRollbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RansomDefenseRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncBaselineDetectSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *SyncBaselineDetectSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncBaselineDetectSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineClearHistoryRequest struct {
	*tchttp.BaseRequest

	// 需要删除的记录id值,最大长度100个

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteMachineClearHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineClearHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BruteAttackRule struct {

	// 爆破事件失败次数

	LoginFailTimes *uint64 `json:"LoginFailTimes,omitempty" name:"LoginFailTimes"`
	// 爆破事件发生的时间范围，单位：秒

	TimeRange *uint64 `json:"TimeRange,omitempty" name:"TimeRange"`
}

type DescribeVersionCompareChartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json 字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVersionCompareChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVersionCompareChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportSecurityTrendsRequest struct {
	*tchttp.BaseRequest

	// 开始时间。

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 结束时间。

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *ExportSecurityTrendsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSecurityTrendsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号数

		AccountCount *uint64 `json:"AccountCount,omitempty" name:"AccountCount"`
		// 账号今日新增

		AccountNewCount *int64 `json:"AccountNewCount,omitempty" name:"AccountNewCount"`
		// 数据库数

		DatabaseCount *uint64 `json:"DatabaseCount,omitempty" name:"DatabaseCount"`
		// 数据库今日新增

		DatabaseNewCount *int64 `json:"DatabaseNewCount,omitempty" name:"DatabaseNewCount"`
		// 主机数

		MachineCount *uint64 `json:"MachineCount,omitempty" name:"MachineCount"`
		// 主机今日新增

		MachineNewCount *int64 `json:"MachineNewCount,omitempty" name:"MachineNewCount"`
		// 端口数

		PortCount *uint64 `json:"PortCount,omitempty" name:"PortCount"`
		// 端口今日新增

		PortNewCount *int64 `json:"PortNewCount,omitempty" name:"PortNewCount"`
		// 进程数

		ProcessCount *uint64 `json:"ProcessCount,omitempty" name:"ProcessCount"`
		// 进程今日新增

		ProcessNewCount *int64 `json:"ProcessNewCount,omitempty" name:"ProcessNewCount"`
		// 软件数

		SoftwareCount *uint64 `json:"SoftwareCount,omitempty" name:"SoftwareCount"`
		// 软件今日新增

		SoftwareNewCount *int64 `json:"SoftwareNewCount,omitempty" name:"SoftwareNewCount"`
		// Web应用数

		WebAppCount *uint64 `json:"WebAppCount,omitempty" name:"WebAppCount"`
		// Web应用今日新增

		WebAppNewCount *int64 `json:"WebAppNewCount,omitempty" name:"WebAppNewCount"`
		// Web框架数

		WebFrameCount *uint64 `json:"WebFrameCount,omitempty" name:"WebFrameCount"`
		// Web框架今日新增

		WebFrameNewCount *int64 `json:"WebFrameNewCount,omitempty" name:"WebFrameNewCount"`
		// Web站点数

		WebLocationCount *uint64 `json:"WebLocationCount,omitempty" name:"WebLocationCount"`
		// Web站点今日新增

		WebLocationNewCount *int64 `json:"WebLocationNewCount,omitempty" name:"WebLocationNewCount"`
		// Web服务数

		WebServiceCount *uint64 `json:"WebServiceCount,omitempty" name:"WebServiceCount"`
		// Web服务今日新增

		WebServiceNewCount *int64 `json:"WebServiceNewCount,omitempty" name:"WebServiceNewCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulEffectHostListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>AliasName - String - 主机名筛选</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *ExportVulEffectHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulEffectHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebHookRuleSummary struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 机器人地址

	HookAddr *string `json:"HookAddr,omitempty" name:"HookAddr"`
	// 主机数目

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 主机范围

	HostLabels []*WebHookHostLabel `json:"HostLabels,omitempty" name:"HostLabels"`
	// 是否启用[1:禁用|0:启用]

	IsDisabled *int64 `json:"IsDisabled,omitempty" name:"IsDisabled"`
	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 事件类型

	RuleItems []*WebHookEventKv `json:"RuleItems,omitempty" name:"RuleItems"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 备注信息

	RuleRemark *string `json:"RuleRemark,omitempty" name:"RuleRemark"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
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

type DescribeAssetEnvListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 环境变量名</li>
	// <li>Type- int - 是否必填：否 - 类型：0用户变量，1系统变量</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 该字段已废弃，由Filters代替

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetEnvListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetEnvListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogInfoRequest struct {
	*tchttp.BaseRequest

	// 日志ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAttackLogInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackLogInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ExportVulDefenceEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件Id 可通过ExportTasks接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulDefenceEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefenceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetCoreModuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内核模块详情

		Module *AssetCoreModuleDetail `json:"Module,omitempty" name:"Module"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetCoreModuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetCoreModuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostTotalCountRequest struct {
	*tchttp.BaseRequest

	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetHostTotalCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostTotalCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseStrategyMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表

		List []*RansomDefenseStrategyMachineDetail `json:"List,omitempty" name:"List"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRansomDefenseStrategyMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseStrategyMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMalwareWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMalwareWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMalwareWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLicenseDetailRequest struct {
	*tchttp.BaseRequest

	// 导出月份, 该参数仅在IsHistory 时可选.

	ExportMonth *string `json:"ExportMonth,omitempty" name:"ExportMonth"`
	// 多个条件筛选时 LicenseStatus,DeadlineStatus,ResourceId,Keywords 取交集
	// <li> LicenseType  授权类型, 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月</li>
	// <li>ResourceId 资源ID</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 是否导出全部授权详情

	IsHistory *bool `json:"IsHistory,omitempty" name:"IsHistory"`
	// 标签筛选,平台标签能力,这里传入 标签键,标签值作为一个对象

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

func (r *ExportLicenseDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLicenseDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Malware struct {

	// 木马描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 木马文件创建时间。

	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`
	// 木马所在的路径。

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 事件ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 木马文件修改时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 当前木马状态。
	// <li>UN_OPERATED：未处理</li><li>SEGREGATED：已隔离</li><li>TRUSTED：已信任</li>
	// <li>SEPARATING：隔离中</li><li>RECOVERING：恢复中</li>

	Status *string `json:"Status,omitempty" name:"Status"`
	// 云镜客户端唯一标识UUID。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ScreenEventsCnt struct {

	// name 具体展示内容类型： 攻击事件, 潜在风险, 失陷资产, 潜在风险资产
	// Value: 事件统计数

	Category []*ScreenNameValue `json:"Category,omitempty" name:"Category"`
	// 展示内容：待处理风险总数，影响资产总数

	Title *string `json:"Title,omitempty" name:"Title"`
	// 事件总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type Vul struct {

	// 受影响机器数量

	ImpactedHostNum *uint64 `json:"ImpactedHostNum,omitempty" name:"ImpactedHostNum"`
	// 最后扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 漏洞种类ID

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞危害等级:
	// HIGH：高危
	// MIDDLE：中危
	// LOW：低危
	// NOTICE：提示

	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
	// 漏洞状态
	// * UN_OPERATED : 待处理
	// * FIXED : 已修复

	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`
}

type HostLoginList struct {

	// 城市id

	City *uint64 `json:"City,omitempty" name:"City"`
	// 国家id

	Country *uint64 `json:"Country,omitempty" name:"Country"`
	// 高危信息说明：
	// ABROAD - 海外IP；
	// XTI - 威胁情报

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 记录Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否命中异地登录异常  1表示命中此类异常, 0表示未命中

	IsRiskArea *uint64 `json:"IsRiskArea,omitempty" name:"IsRiskArea"`
	// 是否命中异常IP异常 1表示命中此类异常, 0表示未命中

	IsRiskSrcIp *uint64 `json:"IsRiskSrcIp,omitempty" name:"IsRiskSrcIp"`
	// 是否命中异常时间异常 1表示命中此类异常, 0表示未命中

	IsRiskTime *uint64 `json:"IsRiskTime,omitempty" name:"IsRiskTime"`
	// 是否命中异常用户异常 1表示命中此类异常, 0表示未命中

	IsRiskUser *uint64 `json:"IsRiskUser,omitempty" name:"IsRiskUser"`
	// 位置名称

	Location *string `json:"Location,omitempty" name:"Location"`
	// 登录时间

	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机ip

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 省份id

	Province *uint64 `json:"Province,omitempty" name:"Province"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 危险等级：
	// 0 高危
	// 1 可疑

	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 来源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 1:正常登录；2异地登录； 5已加白； 14：已处理；15：已忽略。

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// Uuid串

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetPortCountRequest struct {
	*tchttp.BaseRequest

	// 搜索条件：返回指定端口号数据

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeAssetPortCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineCancelParam struct {

	// 取消的策略ID

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 取消的任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeWeeklyReportInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 周报开始时间。

		BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
		// 密码破解成功数。

		BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitempty" name:"BruteAttackSuccessNum"`
		// 账号所属公司或个人名称。

		CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`
		// 导出文件下载地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 周报结束时间。

		EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
		// 安全等级。
		// <li>HIGH：高</li>
		// <li>MIDDLE：中</li>
		// <li>LOW：低</li>

		Level *string `json:"Level,omitempty" name:"Level"`
		// 机器总数。

		MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
		// 木马记录数。

		MalwareNum *uint64 `json:"MalwareNum,omitempty" name:"MalwareNum"`
		// 异地登录数。

		NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitempty" name:"NonlocalLoginNum"`
		// 云镜客户端离线数。

		OfflineMachineNum *uint64 `json:"OfflineMachineNum,omitempty" name:"OfflineMachineNum"`
		// 云镜客户端在线数。

		OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitempty" name:"OnlineMachineNum"`
		// 开通云镜专业版数量。

		ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitempty" name:"ProVersionMachineNum"`
		// 漏洞数。

		VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLocalStorageItemRequest struct {
	*tchttp.BaseRequest

	// 失效时间（单位；秒）

	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`
	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *SetLocalStorageItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLocalStorageItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetKeyVal struct {

	// 描述信息

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 标签

	Key *string `json:"Key,omitempty" name:"Key"`
	// 今日新增数量

	NewCount *int64 `json:"NewCount,omitempty" name:"NewCount"`
	// 数量

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type BaselineRuleDetect struct {

	// 0:未通过 1:忽略 3:通过 5:检测中

	DetectStatus *int64 `json:"DetectStatus,omitempty" name:"DetectStatus"`
	// 首次检测时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 关联主机数

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 关联项数

	ItemCount *int64 `json:"ItemCount,omitempty" name:"ItemCount"`
	// ItemID集合

	ItemIds []*int64 `json:"ItemIds,omitempty" name:"ItemIds"`
	// string

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 规则描述

	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`
	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type ChangeRuleEventsIgnoreStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeRuleEventsIgnoreStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeRuleEventsIgnoreStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceOpenProVersionPrepaidRequest struct {
	*tchttp.BaseRequest

	// 预付费模式(包年包月)参数设置。

	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`
	// 需要开通专业版机器列表数组。

	Machines []*ProVersionMachine `json:"Machines,omitempty" name:"Machines"`
}

func (r *InquiryPriceOpenProVersionPrepaidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceOpenProVersionPrepaidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RansomDefenseStrategyMachineBackupInfo struct {

	// 备份数量

	BackupCount *uint64 `json:"BackupCount,omitempty" name:"BackupCount"`
	// 备份成功次数

	BackupSuccessCount *uint64 `json:"BackupSuccessCount,omitempty" name:"BackupSuccessCount"`
	// 云标签

	CloudTags []*Tag `json:"CloudTags,omitempty" name:"CloudTags"`
	// 硬盘信息，为空时所有硬盘生效：
	// ;分割 diskId1|diskName1;diskId2|diskName2

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// 主机实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 最近一次备份失败原因

	LastBackupMessage *string `json:"LastBackupMessage,omitempty" name:"LastBackupMessage"`
	// 最近一次备份状态：0备份中，1正常，2失败，9暂无备份

	LastBackupStatus *uint64 `json:"LastBackupStatus,omitempty" name:"LastBackupStatus"`
	// 最近一次备份时间

	LastBackupTime *string `json:"LastBackupTime,omitempty" name:"LastBackupTime"`
	// 内网ip

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 外网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 可用区信息

	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`
	// 最近一次回滚进度百分比

	RollBackPercent *uint64 `json:"RollBackPercent,omitempty" name:"RollBackPercent"`
	// 最近一次回滚状态：0进行中，1成功，2失败

	RollBackStatus *uint64 `json:"RollBackStatus,omitempty" name:"RollBackStatus"`
	// 防护状态：0关闭，1开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 策略id，为0时未绑定策略

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 主机安全标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type SyncAssetScanRequest struct {
	*tchttp.BaseRequest

	// 是否同步：true-是 false-否；默认false

	Sync *bool `json:"Sync,omitempty" name:"Sync"`
}

func (r *SyncAssetScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetScanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineWeakPasswordListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBaselineWeakPasswordListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineWeakPasswordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMalwareWhiteListRequest struct {
	*tchttp.BaseRequest

	// 木马事件ID

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 文件目录(正则)；长度不超过200个,内容base64转义

	FileDirectory []*string `json:"FileDirectory,omitempty" name:"FileDirectory"`
	// 文件后缀；长度不超过200个,内容base64转义（废弃）

	FileExtension []*string `json:"FileExtension,omitempty" name:"FileExtension"`
	// 文件名称(正则)；长度不超过200个

	FileName []*string `json:"FileName,omitempty" name:"FileName"`
	// 是否全部主机； 0否，1是。

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 匹配模式 ；0 精确匹配，1模糊匹配（废弃）

	MatchType *uint64 `json:"MatchType,omitempty" name:"MatchType"`
	// MD5列表

	Md5List []*string `json:"Md5List,omitempty" name:"Md5List"`
	// 白名单模式； 0MD5白名单，1自定义

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// quuid 列表

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
}

func (r *CreateMalwareWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMalwareWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortStatisticsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeOpenPortStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpenPortStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeProcessTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProcessTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IgnoreImpactedHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IgnoreImpactedHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IgnoreImpactedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebServiceInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetWebServiceInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebServiceInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPlanTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Tasks []*AssetPlanTask `json:"Tasks,omitempty" name:"Tasks"`
		// 总数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetPlanTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPlanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果列表

		List []*BaselineItemInfo `json:"List,omitempty" name:"List"`
		// 总条目数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineItemInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWarningListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWarningListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarningListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBanWhiteListRequest struct {
	*tchttp.BaseRequest

	// 修改白名单规则项

	Rules *BanWhiteList `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyBanWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBanWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartBaselineDetectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 扫描任务ID

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartBaselineDetectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartBaselineDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePayPageBrowsingRecordsRequest struct {
	*tchttp.BaseRequest
}

func (r *CreatePayPageBrowsingRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePayPageBrowsingRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenEmergentMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 通知内容

		MessageInfo []*ScreenEmergentMsg `json:"MessageInfo,omitempty" name:"MessageInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenEmergentMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenEmergentMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgnoreBaselineRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 忽略基线检测项列表信息

		IgnoreBaselineRuleList []*IgnoreBaselineRule `json:"IgnoreBaselineRuleList,omitempty" name:"IgnoreBaselineRuleList"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIgnoreBaselineRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgnoreBaselineRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetProcessCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各进程数量

		Process []*AssetKeyVal `json:"Process,omitempty" name:"Process"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetProcessCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
}

func (r *DescribeScanTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRansomDefenseBackupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRansomDefenseBackupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRansomDefenseBackupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskDnsPolicy struct {

	// 域名,作为入参时需要进行base64 encode

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
	// 事件ID

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 主机ID

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 主机范围[1: 所有专业版+旗舰版|2:所有旗舰版|0: 部分主机]

	HostScope *int64 `json:"HostScope,omitempty" name:"HostScope"`
	// 是否处理之前的事件[0:不处理|1:处理]

	IsDealOldEvent *int64 `json:"IsDealOldEvent,omitempty" name:"IsDealOldEvent"`
	// 是否生效[0:生效,1:不生效]

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 策略动作[0:告警,1:放行,2:拦截+告警]

	PolicyAction *int64 `json:"PolicyAction,omitempty" name:"PolicyAction"`
	// 策略描述

	PolicyDesc *string `json:"PolicyDesc,omitempty" name:"PolicyDesc"`
	// 策略ID

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略类型[0:系统,1:用户]

	PolicyType *int64 `json:"PolicyType,omitempty" name:"PolicyType"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ModifyRiskDnsPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略

	Data *RiskDnsPolicy `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyRiskDnsPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskDnsPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportProtectDirListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportProtectDirListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportProtectDirListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineRuleCategoryListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBaselineRuleCategoryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleCategoryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineFileTamperRulesRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMachineFileTamperRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineFileTamperRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// 排序字段：CreateTime-发生时间

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(主机内网IP|进程名)</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：根据请求次数排序：asc-升序/desc-降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeReverseShellEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWebPageEventLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWebPageEventLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWebPageEventLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetMachineInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CPU信息

		CPU *string `json:"CPU,omitempty" name:"CPU"`
		// 主机内网IP

		MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
		// 主机名称

		MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
		// 内存生产商

		MemManufacturer *string `json:"MemManufacturer,omitempty" name:"MemManufacturer"`
		// 内存使用率

		MemPercent *float64 `json:"MemPercent,omitempty" name:"MemPercent"`
		// 内存容量：单位G

		MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`
		// 网卡信息列表

		NetworkCards []*AssetNetworkCardInfo `json:"NetworkCards,omitempty" name:"NetworkCards"`
		// 操作系统名称

		OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
		// 设备型号

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// 序列号

		SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetMachineInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetCoreModuleDetail struct {

	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 被依赖模块

	Modules *string `json:"Modules,omitempty" name:"Modules"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数信息

	Params []*AssetCoreModuleParam `json:"Params,omitempty" name:"Params"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 依赖进程

	Processes *string `json:"Processes,omitempty" name:"Processes"`
	// 大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type CreateEmergencyVulScanRequest struct {
	*tchttp.BaseRequest

	// 扫描超时时长 ，单位秒

	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitempty" name:"TimeoutPeriod"`
	// 自选服务器时生效，主机uuid的string数组

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *CreateEmergencyVulScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEmergencyVulScanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FixBaselineDetectRequest struct {
	*tchttp.BaseRequest

	// 修复内容

	Data []*string `json:"Data,omitempty" name:"Data"`
	// 主机Id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 项Id

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
}

func (r *FixBaselineDetectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FixBaselineDetectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulRequest struct {
	*tchttp.BaseRequest

	// 服务器分类：1:专业版服务器；2:自选服务器

	HostType *uint64 `json:"HostType,omitempty" name:"HostType"`
	// 自选服务器时生效，主机quuid的string数组

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// 超时时长 单位秒 默认 3600 秒

	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitempty" name:"TimeoutPeriod"`
	// 漏洞类型：1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞 (多选英文;分隔)

	VulCategories *string `json:"VulCategories,omitempty" name:"VulCategories"`
	// 是否是应急漏洞 0 否 1 是

	VulEmergency *uint64 `json:"VulEmergency,omitempty" name:"VulEmergency"`
	// 需要扫描的漏洞id

	VulIds []*uint64 `json:"VulIds,omitempty" name:"VulIds"`
	// 危害等级：1-低危；2-中危；3-高危；4-严重 (多选英文;分隔)

	VulLevels *string `json:"VulLevels,omitempty" name:"VulLevels"`
}

func (r *ScanVulRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanVulRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebFrameBaseInfo struct {

	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	//  附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 数据库名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 主机业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 主机标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeBaselineRuleDetectListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [HostCount|FirstTime|LastTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>ItemId - int64 - 是否必填：否 - 策略Id</li>
	// <li>RuleName - string - 是否必填：否 - 规则名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态</li>
	// <li>StartTime - string - 是否必填：否 - 开时时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineRuleDetectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleDetectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 反弹shell详情信息

		ReverseShellEventInfo *ReverseShellEventInfo `json:"ReverseShellEventInfo,omitempty" name:"ReverseShellEventInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellEventInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoNotNoticeBanTipsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoNotNoticeBanTipsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoNotNoticeBanTipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebAppBaseInfo struct {

	// 应用描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 站点域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 应用ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	//  附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 应用名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 插件数

	PluginCount *uint64 `json:"PluginCount,omitempty" name:"PluginCount"`
	// 主机业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 根路径

	RootPath *string `json:"RootPath,omitempty" name:"RootPath"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 主机标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 虚拟路径

	VirtualPath *string `json:"VirtualPath,omitempty" name:"VirtualPath"`
}

type AttackSourceEvent struct {

	// 【文件查杀】病毒名 VirusName、文件名 FileName、文件路径 FilePath、文件大小 FileSize、文件MD5 MD5、首次发现时间 CreateTime、最近检测时间LatestScanTime、危害描述 HarmDescribe、修复建议SuggestScheme
	// 【异常登录】来源IP SrcIp、来源地 Location、登录用户名 UserName、登录时间 LoginTime
	// 【密码破解】来源IP SrcIp、来源地 City,Country 、协议 Protocol、登录用户名UserName 、端口 Port、尝试次数 Count、首次攻击时间 CreateTime、最近攻击时间 ModifyTime
	// 【恶意请求】恶意请求域名 Url、进程ProcessName 、MD5 ProcessMd5、PID Pid、请求次数 AccessCount、最近请求时间 MergeTime、危害描述 HarmDescribe、修复建议SuggestScheme
	// 【高危命令】命中规则名 RuleName、规则类别 RuleCategory、命令内容 BashCmd、数据来源 DetectBy、登录用户 User、PID Pid、发生时间 CreateTime 、危害描述 HarmDescribe、修复建议SuggestScheme

	Content *string `json:"Content,omitempty" name:"Content"`
	// 入侵时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 事件类型：0：文件查杀，1：异常登录， 2：密码破解，3：恶意请求，4：高危命令

	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`
	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 等级 事件统一等级 0：提示，1：低危,  2：中危,  3：高危,  4：严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 等级中文展示字符串

	LevelZh *string `json:"LevelZh,omitempty" name:"LevelZh"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type BaselineDownload struct {

	// 下载地址

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// 完成时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 状态0:未完成 1:完成

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务Id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type BaselineFix struct {

	// 首次检测时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修复时间

	FixTime *string `json:"FixTime,omitempty" name:"FixTime"`
	// 主机Ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 基线检测项结果ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 修复项名称

	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 最后检测时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type RansomDefenseBackup struct {

	// 备份状态：0备份中，1正常，2、3失败，4快照已过期，9快照已删除

	BackupStatus *uint64 `json:"BackupStatus,omitempty" name:"BackupStatus"`
	// 备份时间

	BackupTime *string `json:"BackupTime,omitempty" name:"BackupTime"`
	// 备份磁盘数量

	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`
	// 硬盘信息，；分隔

	Disks *string `json:"Disks,omitempty" name:"Disks"`
	// 勒索状态：0无告警, 1有告警

	EventStatus *uint64 `json:"EventStatus,omitempty" name:"EventStatus"`
	// 快照列表，；分隔

	SnapshotIds *string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
	// 策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 策略状态:0关闭，1开启，9已删除

	StrategyStatus *uint64 `json:"StrategyStatus,omitempty" name:"StrategyStatus"`
}

type DeleteAllJavaMemShellsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAllJavaMemShellsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAllJavaMemShellsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentVul struct {

	// 漏洞描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 漏洞ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 最后扫描时间。

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 漏洞种类ID。

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞危害等级。
	// <li>HIGH：高危</li>
	// <li>MIDDLE：中危</li>
	// <li>LOW：低危</li>
	// <li>NOTICE：提示</li>

	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`
	// 漏洞名称。

	VulName *string `json:"VulName,omitempty" name:"VulName"`
	// 漏洞状态。
	// <li>UN_OPERATED : 待处理</li>
	// <li>FIXED : 已修复</li>

	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`
}

type ScanAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MalwareWhiteListInfo struct {

	// 规则创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 受影响记录

	EventsCount *uint64 `json:"EventsCount,omitempty" name:"EventsCount"`
	// 文件目录；按,分割

	FileDirectory *string `json:"FileDirectory,omitempty" name:"FileDirectory"`
	// 文件后缀；按,分割

	FileExtension *string `json:"FileExtension,omitempty" name:"FileExtension"`
	// 文件名；按,分割

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 唯一ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全部主机； 0否，1是

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 匹配模式；0精确匹配，1模糊匹配

	MatchType *uint64 `json:"MatchType,omitempty" name:"MatchType"`
	// md5列表 按,分割

	Md5List *string `json:"Md5List,omitempty" name:"Md5List"`
	// 白名单模式；0 MD5 ，1自定义

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// cvm quuid 按,分割。

	QuuidList *string `json:"QuuidList,omitempty" name:"QuuidList"`
}

type DeleteNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 删除异地登录事件的方式，可选值："Ids"、"Ip"、"All"，默认为Ids

	DelType *string `json:"DelType,omitempty" name:"DelType"`
	// 异地登录事件ID数组。DelType为Ids或DelType未填时此项必填

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 异地登录事件的Ip。DelType为Ip时必填

	Ip []*string `json:"Ip,omitempty" name:"Ip"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DeleteNonlocalLoginPlacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNonlocalLoginPlacesRequest) FromJsonString(s string) error {
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

type ScanTaskAgainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanTaskAgainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanTaskAgainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskDnsList struct {

	// 访问次数

	AccessCount *uint64 `json:"AccessCount,omitempty" name:"AccessCount"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 命令行

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 首次访问时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否为全局规则，0否，1是

	GlobalRuleId *uint64 `json:"GlobalRuleId,omitempty" name:"GlobalRuleId"`
	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机在线状态[OFFLINE:离线|ONLINE:在线|UNKNOWN:未知]

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 外网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 最近访问时间

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// 进程号

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 进程MD5

	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 唯一 Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 参考

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 状态；0-待处理，2-已加白，3-非信任状态，4-已处理，5-已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 建议方案

	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
	// 标签特性

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 对外访问域名

	Url *string `json:"Url,omitempty" name:"Url"`
	// 用户规则id

	UserRuleId *uint64 `json:"UserRuleId,omitempty" name:"UserRuleId"`
	// 唯一UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetScanHostDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异常信息

		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
		// 任务状态：0进行中，1已完成，2异常

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetScanHostDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetScanHostDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BruteAttackInfo struct {

	// 阻断状态：1-阻断成功；非1-阻断失败

	BanStatus *uint64 `json:"BanStatus,omitempty" name:"BanStatus"`
	// 城市id

	City *uint64 `json:"City,omitempty" name:"City"`
	// 发生次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 国家id

	Country *uint64 `json:"Country,omitempty" name:"Country"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 0：待处理，1：忽略，5：已处理，6：加入白名单

	DataStatus *uint64 `json:"DataStatus,omitempty" name:"DataStatus"`
	// 事件类型：200-暴力破解事件，300-暴力破解成功事件（页面展示），400-暴力破解不存在的帐号事件

	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`
	// 唯一Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否为专业版（true/false）

	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机ip

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 最近攻击时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 被攻击的服务的用户名

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 省份id

	Province *uint64 `json:"Province,omitempty" name:"Province"`
	// 机器UUID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 来源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// SUCCESS：破解成功；FAILED：破解失败

	Status *string `json:"Status,omitempty" name:"Status"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 云镜客户端唯一标识UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type MalwareRisk struct {

	// 发现时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 机器IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 病毒名

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
}

type DescribeNetAttackWhiteListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [CreateTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IP - String - 是否必填：否 - 主机ip查询 </li>
	// <li>SrcIP- String - 是否必填：否 - 白名单ip查询 </li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetAttackWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetAttackWhiteListRequest) FromJsonString(s string) error {
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

type DescribeRiskDnsEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求事件列表

		List []*RiskDnsEvent `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrustMalwaresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TrustMalwaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperRulesRequest struct {
	*tchttp.BaseRequest

	// 排序字段 CreateTime、ModifyTime、HostCount

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>RuleCategory- string- 规则类别  0=系统规则，1=用户规则</li>
	// <li>Name- String - 规则名称/li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 ASC,DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeFileTamperRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBaselineStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBaselineStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBaselineStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSaveOrUpdateWarningsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSaveOrUpdateWarningsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSaveOrUpdateWarningsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OsName struct {

	// 操作系统类型枚举值

	MachineOSType *uint64 `json:"MachineOSType,omitempty" name:"MachineOSType"`
	// 系统名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeWarningHostConfigRequest struct {
	*tchttp.BaseRequest

	// 分页单页限制数目, 0表示不分页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 告警类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeWarningHostConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarningHostConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditBashRulesRequest struct {
	*tchttp.BaseRequest

	// 是否处理旧事件为白名单 0=不处理 1=处理

	DealOldEvents *uint64 `json:"DealOldEvents,omitempty" name:"DealOldEvents"`
	// 事件列表点击“加入白名单”时,需要传EventId 事件的id

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 规则ID（新增时不填）

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则(默认否)：1-全局，0-非全局

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 危险等级(0:无，1: 高危 2:中危 3: 低危)

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 规则名称，编辑时不可修改规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 正则表达式 ，编辑时不可修改正则表达式，需要对内容QueryEscape后再base64

	Rule *string `json:"Rule,omitempty" name:"Rule"`
	// 客户端ID数组

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
	// 0=黑名单， 1=白名单

	White *uint64 `json:"White,omitempty" name:"White"`
}

func (r *EditBashRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditBashRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLocalStorageExpireRequest struct {
	*tchttp.BaseRequest

	// 过期时间（单位：秒）

	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`
	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *SetLocalStorageExpireRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLocalStorageExpireRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DefendAttackLog struct {

	// 攻击时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 目标IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 目标端口

	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`
	// 攻击描述

	HttpCgi *string `json:"HttpCgi,omitempty" name:"HttpCgi"`
	// 攻击内容

	HttpContent *string `json:"HttpContent,omitempty" name:"HttpContent"`
	// 攻击方式

	HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
	// 攻击参数

	HttpParam *string `json:"HttpParam,omitempty" name:"HttpParam"`
	// 日志ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 目标服务器IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 目标服务器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 来源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 来源端口

	SrcPort *uint64 `json:"SrcPort,omitempty" name:"SrcPort"`
	// 客户端ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 威胁类型

	VulType *string `json:"VulType,omitempty" name:"VulType"`
}

type DescribeExpertServiceListRequest struct {
	*tchttp.BaseRequest

	// 排序字段 StartTime，EndTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Keyword- String - 是否必填：否 - 关键词过滤，</li>
	// <li>Uuids - String - 是否必填：否 - 主机id过滤</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序步长

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方法

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeExpertServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExpertServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各账号数量

		Users []*AssetKeyVal `json:"Users,omitempty" name:"Users"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetUserCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenProtectionStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件查杀 status:  0:从未检测过，或0资产付费情况, 1:已检测，存在恶意文件, 2:已检测，未开启隔离防护, 3:已检测且已开启防护且无风险
		// 暴力破解status: 0:未开启防护或0付费资产情况 1:已开启防护 2:存在带处理事件
		// 漏洞扫描status: 0:从未检测过，或0资产付费情况, 1:存在漏洞风险, 2:无风险
		// 基线检测status: 0:从未检测过，或0资产付费情况, 1:存在基线风险,2:无风险

		Info []*ScreenProtection `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenProtectionStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenProtectionStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulFixResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		FixId *uint64 `json:"FixId,omitempty" name:"FixId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulFixResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulFixResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineHostIgnoreListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineHost `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineHostIgnoreListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineHostIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyOrderRequest struct {
	*tchttp.BaseRequest

	// 授权类型 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DestroyOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoginWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 进程列表数据数组。

		Processes []*Process `json:"Processes,omitempty" name:"Processes"`
		// 进程列表记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcessesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProcessesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用列表

		Apps []*AssetAppBaseInfo `json:"Apps,omitempty" name:"Apps"`
		// 总数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetAppListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRansomDefenseEventsStatusRequest struct {
	*tchttp.BaseRequest

	// 是否更新所有同路径事件

	All *bool `json:"All,omitempty" name:"All"`
	// 需要修改的事件id 数组，支持批量

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 操作 0待处理，1已处理,2信任,9:删除记录

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyRansomDefenseEventsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRansomDefenseEventsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartBaselineDetectRequest struct {
	*tchttp.BaseRequest

	// 基线检测参数

	Param *BaselineDetectParam `json:"Param,omitempty" name:"Param"`
}

func (r *StartBaselineDetectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartBaselineDetectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopBaselineDetectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopBaselineDetectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopBaselineDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentVuls struct {

	// 漏洞描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 漏洞ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 最后扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 漏洞种类ID

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞危害等级:
	// HIGH：高危
	// MIDDLE：中危
	// LOW：低危
	// NOTICE：提示

	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type DescribeSecurityBroadcastInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全播报文章详情

		BroadcastInfo *BroadcastInfo `json:"BroadcastInfo,omitempty" name:"BroadcastInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityBroadcastInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityBroadcastInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBanStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBanStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportEventlogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportEventlogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportEventlogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebPageProtectSettingRequest struct {
	*tchttp.BaseRequest

	// 配置对应的protect_path

	Id *string `json:"Id,omitempty" name:"Id"`
	// 需要操作的类型1 目录名称 2 防护文件类型

	ModifyType *uint64 `json:"ModifyType,omitempty" name:"ModifyType"`
	// 提交值

	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyWebPageProtectSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebPageProtectSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSearchTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：成功，非0：失败

		Status *uint64 `json:"Status,omitempty" name:"Status"`
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

type DescribeAssetTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产指纹类型列表

		Types []*AssetType `json:"Types,omitempty" name:"Types"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRansomDefenseMachineListRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime、MachineCount

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Ips - String - 是否必填：否 - 通过ip查询 </li>
	// <li>MachineNames - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>Names - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>Status - String - 是否必填：否 - 策略状态：0备份中，1备份成功，2备份失败 </li>
	// <li>LastBackupTimeBegin - String - 是否必填：否 - 最近一次备份时间开始</li>
	// <li>LastBackupTimeEnd - String - 是否必填：否 - 最近一次备份时间结束</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportRansomDefenseMachineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRansomDefenseMachineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBanModeRequest struct {
	*tchttp.BaseRequest

	// 阻断模式，STANDARD_MODE：标准阻断，DEEP_MODE：深度阻断

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 阻断时间，用于标准阻断模式

	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`
}

func (r *ModifyBanModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBanModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBaselinePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略Id

	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`
}

func (r *DeleteBaselinePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBaselinePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetCoreModuleInfoRequest struct {
	*tchttp.BaseRequest

	// 内核模块ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetCoreModuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetCoreModuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebAppCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各Web应用数量

		WebApps []*AssetKeyVal `json:"WebApps,omitempty" name:"WebApps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebAppCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeparateMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马事件ID数组。(最大100条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 是否杀掉进程

	KillProcess *bool `json:"KillProcess,omitempty" name:"KillProcess"`
}

func (r *SeparateMalwaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeparateMalwaresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineDetail struct {

	// 基线描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 危害等级

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 基线名

	Name *string `json:"Name,omitempty" name:"Name"`
	// package名

	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
	// 父级id

	ParentId *uint64 `json:"ParentId,omitempty" name:"ParentId"`
}

type SecurityDynamic struct {

	// 安全事件发生时间。

	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`
	// 安全事件类型。
	// <li>MALWARE：木马事件</li>
	// <li>NON_LOCAL_LOGIN：异地登录</li>
	// <li>BRUTEATTACK_SUCCESS：密码破解成功</li>
	// <li>VUL：漏洞</li>
	// <li>BASELINE：安全基线</li>

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 安全事件消息。

	Message *string `json:"Message,omitempty" name:"Message"`
	// 安全事件等级。
	// <li>RISK: 严重</li>
	// <li>HIGH: 高危</li>
	// <li>NORMAL: 中危</li>
	// <li>LOW: 低危</li>
	// <li>UNKNOWNED: 可疑</li>

	SecurityLevel *string `json:"SecurityLevel,omitempty" name:"SecurityLevel"`
	// 云镜客户端UUID。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeEventTrendRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>DateRange - String - 是否必填：否 - 日志时间范围 </li>
	// <li>SeverityClass - String  是否必填：否 - 威胁等级 <br/>"RISK":严重<br/>"HIGH":高危<br/>"MEDIUM":中危<br/>"LOW":低危<br/>"NOTICE":提示</li>
	// <li>EventType - String  是否必填：否 - 事件类型<br/>"MALWARE":木马文件<br/>"HOST_LOGIN":登录审计<br/>"BRUTE_ATTACK":暴力破解<br/>"MALICIOUS_REQUEST":恶意请求<br/>"ATTACK_LOG":网络攻击<br/>"HISH_RISK_BASH":高危命令<br/>"REVERSE_SHELL":反弹shell<br/>"PRIVILEGE_ESCALATION":本地提权<br/>"WEB_VUL":WEB漏洞<br/>"SYSTEM_COMPONENT_VUL":系统组件漏洞<br/>"EMERGENCY_VUL":应急漏洞<br/>"BASELINE":基线漏洞</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeEventTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwaresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 木马相关信息。

		Malwares []*Malware `json:"Malwares,omitempty" name:"Malwares"`
		// 木马总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulCountByDatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 批量获得对应天数的主机数量

		HostCount []*uint64 `json:"HostCount,omitempty" name:"HostCount"`
		// 批量获得对应天数的漏洞数量

		VulCount []*uint64 `json:"VulCount,omitempty" name:"VulCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulCountByDatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulCountByDatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 -  查询关键字</li>
	// <li>Status - String - 是否必填：否 -  登录状态（NON_LOCAL_LOGIN: 异地登录 | NORMAL_LOGIN : 正常登录）</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeNonlocalLoginPlacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulAgainRequest struct {
	*tchttp.BaseRequest

	// 漏洞事件id串，多个用英文逗号分隔

	EventIds *string `json:"EventIds,omitempty" name:"EventIds"`
	// 重新检查的机器uuid,多个逗号分隔

	Uuids *string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *ScanVulAgainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanVulAgainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackTrendsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击趋势统计数据（天）

		NetAttackTrend []*NetAttackTrend `json:"NetAttackTrend,omitempty" name:"NetAttackTrend"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackTrendsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImportMachineInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 有效的机器信息列表：机器名称、机器公网/内网ip、机器标签

		EffectiveMachineInfoList []*EffectiveMachineInfo `json:"EffectiveMachineInfoList,omitempty" name:"EffectiveMachineInfoList"`
		// 用户批量导入失败的机器列表（例如机器不存在等...）

		InvalidMachineList []*string `json:"InvalidMachineList,omitempty" name:"InvalidMachineList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImportMachineInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetInitServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetInitServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetInitServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulFixRequest struct {
	*tchttp.BaseRequest

	// 漏洞对应要修复的主机列表

	CreateVulFixTaskQuuids []*CreateVulFixTaskQuuids `json:"CreateVulFixTaskQuuids,omitempty" name:"CreateVulFixTaskQuuids"`
	// 快照保存天数，  0 天表示不创建快照。24小时没有快照的主机必须要创建快照才能修复

	SaveDays *uint64 `json:"SaveDays,omitempty" name:"SaveDays"`
	// 快照名称

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}

func (r *CreateVulFixRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulFixRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineAnalysisDataRequest struct {
	*tchttp.BaseRequest

	// 基线策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DescribeBaselineAnalysisDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineAnalysisDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulTopRequest struct {
	*tchttp.BaseRequest

	// 是否仅统计重点关注漏洞 1=仅统计重点关注漏洞, 0=统计全部漏洞

	IsFollowVul *uint64 `json:"IsFollowVul,omitempty" name:"IsFollowVul"`
	// 漏洞风险服务器top，1-100

	Top *uint64 `json:"Top,omitempty" name:"Top"`
	// 1:web-cms 漏洞，2.应用漏洞 4: Linux软件漏洞 5: windows系统漏洞 6:应急漏洞，不填或者填0时返回 1，2，4，5 的总统计数据

	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`
}

func (r *DescribeVulTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineRuleDetectListRequest struct {
	*tchttp.BaseRequest

	// 0:过滤的结果导出；1:全部导出

	ExportAll *int64 `json:"ExportAll,omitempty" name:"ExportAll"`
	// <li>RuleName - string - 是否必填：否 - 规则名称</i>
	// <li>IsPassed - int - 是否必填：否 - 是否通过</li>
	// <li>RiskTier - int - 是否必填：否 - 风险等级</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBaselineRuleDetectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineRuleDetectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemRiskTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果数组

		RiskItemTop5 []*BaselineRiskItem `json:"RiskItemTop5,omitempty" name:"RiskItemTop5"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineItemRiskTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemRiskTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsInfoNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件详情

		BashEventsInfo *BashEventsInfoNew `json:"BashEventsInfo,omitempty" name:"BashEventsInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashEventsInfoNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashEventsInfoNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态。
		// <li>COMPLETE：完成（此时可以调用DescribeOpenPorts接口获取实时进程列表）</li>
		// <li>AGENT_OFFLINE：云镜客户端离线</li>
		// <li>COLLECTING：端口获取中</li>
		// <li>FAILED：进程获取失败</li>

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenPortTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpenPortTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulFixStatusHostInfo struct {

	// 修复失败原因

	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 修复时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 主机的quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 状态：0-初始状态；1-已下发任务（修复中）2-完成（成功）；3-修复失败（失败）4-快照创建失败 导致修复失败（未修复）；

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DescribeAttackLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志列表

		AttackLogs []*DefendAttackLog `json:"AttackLogs,omitempty" name:"AttackLogs"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginWhiteRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoginWhiteRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginWhiteRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {

	// 主机TagId

	TagId *string `json:"TagId,omitempty" name:"TagId"`
	// 主机TagName

	TagName *string `json:"TagName,omitempty" name:"TagName"`
}

type DeleteMachineClearHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineClearHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineClearHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskStatus struct {

	// 扫描失败

	Fail *string `json:"Fail,omitempty" name:"Fail"`
	// 扫描终止（包含终止中）

	Ok *string `json:"Ok,omitempty" name:"Ok"`
	// 扫描中（包含初始化）

	Scanning *string `json:"Scanning,omitempty" name:"Scanning"`
	// 扫描失败（提示具体原因：扫描超时、客户端版本低、客户端离线）

	Stop *string `json:"Stop,omitempty" name:"Stop"`
}

type DeleteBaselineRuleRequest struct {
	*tchttp.BaseRequest

	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteBaselineRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBaselineRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineRuleType struct {

	// 类型Id

	TypeId *int64 `json:"TypeId,omitempty" name:"TypeId"`
	// 类型名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
}

type ExportVulInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务Id , 可通过ExportTasks 接口下载

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetCoreModuleParam struct {

	// 数据

	Data *string `json:"Data,omitempty" name:"Data"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteMaliciousRequestWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单id (最大100条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteMaliciousRequestWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaliciousRequestWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEmergencyResponseListRequest struct {
	*tchttp.BaseRequest

	// 排序字段 StartTime，EndTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Keyword- String - 是否必填：否 - 关键词过滤，</li>
	// <li>Uuids - String - 是否必填：否 - 主机id过滤</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序步长

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方法

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeEmergencyResponseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEmergencyResponseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DescribeImportMachineInfoRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版 | Flagship：旗舰版 | ProtectedMachines：专业版+旗舰版） | BASIC_PROPOST_GENERAL_DISCOUNT：普惠版+专业版按量计费+基础版主机 | UnFlagship：专业版预付费+专业版后付费+基础版+普惠版</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 批量导入的数据类型：Ip、Name、Id 三选一

	ImportType *string `json:"ImportType,omitempty" name:"ImportType"`
	// 该参数已作废.

	IsQueryProMachine *bool `json:"IsQueryProMachine,omitempty" name:"IsQueryProMachine"`
	// 服务器内网IP（默认）/ 服务器名称 / 服务器ID 数组 (最大 1000条)

	MachineList []*string `json:"MachineList,omitempty" name:"MachineList"`
}

func (r *DescribeImportMachineInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportMachineInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenEventsCntResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件统计详情

		Info []*ScreenEventsCnt `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenEventsCntResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenEventsCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionActivityRequest struct {
	*tchttp.BaseRequest

	// 活动ID

	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`
	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *OpenProVersionActivityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenProVersionActivityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineItemInfo struct {

	// 检测项的修复方法

	FixMethod *string `json:"FixMethod,omitempty" name:"FixMethod"`
	// 检测项描述

	ItemDesc *string `json:"ItemDesc,omitempty" name:"ItemDesc"`
	// 基线检测项ID

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
	// 检测项名字

	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`
	// 危险等级

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 被引自定义规则信息

	RelatedCustomRuleInfo []*BaselineCustomRuleIdName `json:"RelatedCustomRuleInfo,omitempty" name:"RelatedCustomRuleInfo"`
	// 检测项所属规则的ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 检测项所属规则名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 系统规则ID

	SysRuleId *int64 `json:"SysRuleId,omitempty" name:"SysRuleId"`
}

type ValueInfo struct {

	// 是否包含中文

	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
	// 字段是否开启分析功能

	SqlFlag *bool `json:"SqlFlag,omitempty" name:"SqlFlag"`
	// 字段的分词符

	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`
	// 字段类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DeleteUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 已添加常用登录地城市ID数组

	CityIds []*uint64 `json:"CityIds,omitempty" name:"CityIds"`
	// 云镜客户端Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DeleteUsualLoginPlacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUsualLoginPlacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWebPageEventLogRequest struct {
	*tchttp.BaseRequest

	// 需要删除的事件ID 集合 最大100条

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteWebPageEventLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWebPageEventLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseBindListRequest struct {
	*tchttp.BaseRequest

	// <li>Keywords 机器别名/公私IP 模糊查询</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 授权ID

	LicenseId *uint64 `json:"LicenseId,omitempty" name:"LicenseId"`
	// 授权类型

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 限制条数,默认10.

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0.

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DescribeLicenseBindListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseBindListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BanWhiteList struct {

	// 创建白名单时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 白名单ID。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 白名单是否全局

	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 修改白名单时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 白名单所属机器列表

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 白名单别名。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 阻断来源IP。

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 白名单所属机器。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeSecurityEventStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络攻击事件统计

		AttackLogsStat *EventStat `json:"AttackLogsStat,omitempty" name:"AttackLogsStat"`
		// 有未处理基线安全事件的机器总数

		BaseLineTotalAffectNum *uint64 `json:"BaseLineTotalAffectNum,omitempty" name:"BaseLineTotalAffectNum"`
		// 高危基线漏洞事件统计

		BaselineHighStat *EventStat `json:"BaselineHighStat,omitempty" name:"BaselineHighStat"`
		// 低危基线漏事件统计

		BaselineLowStat *EventStat `json:"BaselineLowStat,omitempty" name:"BaselineLowStat"`
		// 中危基线漏事件统计

		BaselineNormalStat *EventStat `json:"BaselineNormalStat,omitempty" name:"BaselineNormalStat"`
		// 严重基线漏洞事件统计

		BaselineRiskStat *EventStat `json:"BaselineRiskStat,omitempty" name:"BaselineRiskStat"`
		// 爆破事件统计

		BruteAttackStat *EventStat `json:"BruteAttackStat,omitempty" name:"BruteAttackStat"`
		// 有未处理网络攻击安全事件的机器总数

		CyberAttackTotalAffectNum *uint64 `json:"CyberAttackTotalAffectNum,omitempty" name:"CyberAttackTotalAffectNum"`
		// 高危命令事件统计

		HighRiskBashStat *EventStat `json:"HighRiskBashStat,omitempty" name:"HighRiskBashStat"`
		// 异地事件统计

		HostLoginStat *EventStat `json:"HostLoginStat,omitempty" name:"HostLoginStat"`
		// 有未处理入侵安全事件的机器总数

		InvasionTotalAffectNum *uint64 `json:"InvasionTotalAffectNum,omitempty" name:"InvasionTotalAffectNum"`
		// 有未处理安全事件的机器总数

		MachineTotalAffectNum *uint64 `json:"MachineTotalAffectNum,omitempty" name:"MachineTotalAffectNum"`
		// 恶意请求事件统计

		MaliciousRequestStat *EventStat `json:"MaliciousRequestStat,omitempty" name:"MaliciousRequestStat"`
		// 木马事件统计

		MalwareStat *EventStat `json:"MalwareStat,omitempty" name:"MalwareStat"`
		// 本地提权事件统计

		PrivilegeStat *EventStat `json:"PrivilegeStat,omitempty" name:"PrivilegeStat"`
		// 反弹Shell事件统计

		ReverseShellStat *EventStat `json:"ReverseShellStat,omitempty" name:"ReverseShellStat"`
		// 安全得分

		Score *uint64 `json:"Score,omitempty" name:"Score"`
		// 高危漏洞事件统计

		VulHighStat *EventStat `json:"VulHighStat,omitempty" name:"VulHighStat"`
		// 低危漏洞事件统计

		VulLowStat *EventStat `json:"VulLowStat,omitempty" name:"VulLowStat"`
		// 中危漏洞事件统计

		VulNormalStat *EventStat `json:"VulNormalStat,omitempty" name:"VulNormalStat"`
		// 严重漏洞事件统计

		VulRiskStat *EventStat `json:"VulRiskStat,omitempty" name:"VulRiskStat"`
		// 漏洞数统计

		VulStat *EventStat `json:"VulStat,omitempty" name:"VulStat"`
		// 有未处理漏洞安全事件的机器总数

		VulTotalAffectNum *uint64 `json:"VulTotalAffectNum,omitempty" name:"VulTotalAffectNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityEventStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityEventStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanBaselineRequest struct {
	*tchttp.BaseRequest

	// 基线id数组(StrategyIdList与CategoryIdList和RuleIdList三选一)

	CategoryIdList []*uint64 `json:"CategoryIdList,omitempty" name:"CategoryIdList"`
	// 选择StrategyIdList时，不需要填写，其他情况必填

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// 检测项id数组(StrategyIdList与CategoryIdList和RuleIdList三选一)

	RuleIdList []*uint64 `json:"RuleIdList,omitempty" name:"RuleIdList"`
	// 策略id数组(StrategyIdList与CategoryIdList和RuleIdList三选一)

	StrategyIdList []*uint64 `json:"StrategyIdList,omitempty" name:"StrategyIdList"`
	// 主机Uuid数组

	UuidList []*string `json:"UuidList,omitempty" name:"UuidList"`
}

func (r *ScanBaselineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanBaselineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagMachine struct {

	// ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机区域

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 主机区域类型

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 主机ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type DescribeAssetDatabaseInfoRequest struct {
	*tchttp.BaseRequest

	// 数据库ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetDatabaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineHostIgnoreListRequest struct {
	*tchttp.BaseRequest

	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 请求的规则

	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
}

func (r *DescribeBaselineHostIgnoreListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineHostIgnoreListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperEventRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情信息

		FileTamperRuleDetail *FileTamperRuleDetail `json:"FileTamperRuleDetail,omitempty" name:"FileTamperRuleDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileTamperEventRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperEventRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmergencyResponseInfo struct {

	// 服务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 主机个数

	HostNum *uint64 `json:"HostNum,omitempty" name:"HostNum"`
	// 报告下载地址

	ReportPath *string `json:"ReportPath,omitempty" name:"ReportPath"`
	// 服务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 服务状态 0未启动，·响应中，2响应完成

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type MalwareListInfo struct {

	// 主机别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 首次运行时间

	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`
	// 最近运行时间

	FileModifierTime *string `json:"FileModifierTime,omitempty" name:"FileModifierTime"`
	// 路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 服务器ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 最近扫描时间

	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
	// 风险等级 0未知、1低、2中、3高、4严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 状态；4-:待处理，5-已信任，6-已隔离，8-文件已删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 特性标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 唯一UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 木马名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
}

type DescribeAssetWebLocationCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各Web站点数量

		WebLocations []*AssetKeyVal `json:"WebLocations,omitempty" name:"WebLocations"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebLocationCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseGeneralRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLicenseGeneralRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseGeneralRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络攻击事件统计。

		AttackLogsStat *EventStat `json:"AttackLogsStat,omitempty" name:"AttackLogsStat"`
		// 有未处理基线安全事件的机器总数。

		BaseLineTotalAffectNum *uint64 `json:"BaseLineTotalAffectNum,omitempty" name:"BaseLineTotalAffectNum"`
		// 高危基线漏洞事件统计。

		BaselineHighStat *EventStat `json:"BaselineHighStat,omitempty" name:"BaselineHighStat"`
		// 低危基线漏事件统计。

		BaselineLowStat *EventStat `json:"BaselineLowStat,omitempty" name:"BaselineLowStat"`
		// 中危基线漏事件统计。

		BaselineNormalStat *EventStat `json:"BaselineNormalStat,omitempty" name:"BaselineNormalStat"`
		// 爆破事件统计。

		BruteAttackStat *EventStat `json:"BruteAttackStat,omitempty" name:"BruteAttackStat"`
		// 有未处理网络攻击安全事件的机器总数。

		CyberAttackTotalAffectNum *uint64 `json:"CyberAttackTotalAffectNum,omitempty" name:"CyberAttackTotalAffectNum"`
		// 高危指令事件统计。

		HighRiskBashStat *EventStat `json:"HighRiskBashStat,omitempty" name:"HighRiskBashStat"`
		// 异地事件统计。

		HostLoginStat *EventStat `json:"HostLoginStat,omitempty" name:"HostLoginStat"`
		// 有未处理入侵安全事件的机器总数。

		InvasionTotalAffectNum *uint64 `json:"InvasionTotalAffectNum,omitempty" name:"InvasionTotalAffectNum"`
		// 有未处理安全事件的机器总数。

		MachineTotalAffectNum *uint64 `json:"MachineTotalAffectNum,omitempty" name:"MachineTotalAffectNum"`
		// 恶意请求事件统计。

		MaliciousRequestStat *EventStat `json:"MaliciousRequestStat,omitempty" name:"MaliciousRequestStat"`
		// 木马事件统计。

		MalwareStat *EventStat `json:"MalwareStat,omitempty" name:"MalwareStat"`
		// 服务器在线数。

		OnlineNum *uint64 `json:"OnlineNum,omitempty" name:"OnlineNum"`
		// 漏洞库更新时间。

		PocUpdateTime *string `json:"PocUpdateTime,omitempty" name:"PocUpdateTime"`
		// 本地提权事件统计。

		PrivilegeStat *EventStat `json:"PrivilegeStat,omitempty" name:"PrivilegeStat"`
		// 专业服务器数。

		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`
		// 反弹Shell事件统计。

		ReverseShellStat *EventStat `json:"ReverseShellStat,omitempty" name:"ReverseShellStat"`
		// 病毒库更新时间。

		VdbUpdateTime *string `json:"VdbUpdateTime,omitempty" name:"VdbUpdateTime"`
		// 高危漏洞事件统计。

		VulHighStat *EventStat `json:"VulHighStat,omitempty" name:"VulHighStat"`
		// 低危漏洞事件统计。

		VulLowStat *EventStat `json:"VulLowStat,omitempty" name:"VulLowStat"`
		// 中危漏洞事件统计。

		VulNormalStat *EventStat `json:"VulNormalStat,omitempty" name:"VulNormalStat"`
		// 有未处理漏洞安全事件的机器总数。

		VulTotalAffectNum *uint64 `json:"VulTotalAffectNum,omitempty" name:"VulTotalAffectNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionSet struct {

	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区信息

	ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet"`
}

type DescribeBaselineRuleDetectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineRuleDetect `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineRuleDetectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefencePluginStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各主机漏洞防御插件信息列表

		List []*VulDefencePluginStatus `json:"List,omitempty" name:"List"`
		// 数据总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefencePluginStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefencePluginStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenPortStatistics struct {

	// 主机数量

	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
	// 端口号

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

type DescribeVulDefenceSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防御开关，0 关闭 1 开启

		Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
		// 当前旗舰版主机数量

		FlagshipCount *uint64 `json:"FlagshipCount,omitempty" name:"FlagshipCount"`
		// 影响主机quuid列表

		Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
		// 影响范围：1 全网旗舰版主机，0 quuid列表主机

		Scope *uint64 `json:"Scope,omitempty" name:"Scope"`
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

type ExportAttackEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAttackEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAttackEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgnoreHostAndItemConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受影响主机

		HostSet []*BaselineHost `json:"HostSet,omitempty" name:"HostSet"`
		// 受影响检测项

		ItemSet []*BaselineItemInfo `json:"ItemSet,omitempty" name:"ItemSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIgnoreHostAndItemConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgnoreHostAndItemConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoOpenProVersionConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoOpenProVersionConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoOpenProVersionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScanMalwareSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScanMalwareSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScanMalwareSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteListNewRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>UserName - String - 是否必填：否 - 用户名筛选</li>
	// <li>ModifyBeginTime - String - 是否必填：否 - 按照修改时间段筛选，开始时间</li>
	// <li>ModifyEndTime - String - 是否必填：否 - 按照修改时间段筛选，结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLoginWhiteListNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginWhiteListNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchExportListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出的任务号

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
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

type DescribeBaselinePolicyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselinePolicy `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselinePolicyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselinePolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意文件详情信息

		MalwareInfo *MalwareInfo `json:"MalwareInfo,omitempty" name:"MalwareInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwareInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// ID数组. (最大100条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteReverseShellEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebFrameJarListRequest struct {
	*tchttp.BaseRequest

	// Web框架ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebFrameJarListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebFrameJarListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemInfoRequest struct {
	*tchttp.BaseRequest

	// 可选排序列

	By *string `json:"By,omitempty" name:"By"`
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>PolicyId - int64 - 是否必填：否 - 项Id</i>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>ItemName - string - 是否必填：否 - 检测项名字</li>
	// <li>RuleId - int - 是否必填：否 - 规则Id</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineItemInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineStrategyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户策略信息列表

		StrategyList []*Strategy `json:"StrategyList,omitempty" name:"StrategyList"`
		// 分页查询记录的总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineStrategyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineStrategyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditReverseShellRuleRequest struct {
	*tchttp.BaseRequest

	// 目标IP

	DestIp *string `json:"DestIp,omitempty" name:"DestIp"`
	// 目标端口

	DestPort *string `json:"DestPort,omitempty" name:"DestPort"`
	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)

	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
	// 规则ID(新增时请留空)

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则(默认否)

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *EditReverseShellRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditReverseShellRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDetectionExcelRequest struct {
	*tchttp.BaseRequest

	// 本次漏洞检测任务id（不同于出参的导出本次漏洞检测Excel的任务Id）

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExportVulDetectionExcelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDetectionExcelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenMachineRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScreenMachineRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenMachineRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBashEventsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBashEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBashEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表详情

		List []*ScreenRegionMachines `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFileTamperRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFileTamperRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFileTamperRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineEffectHost struct {

	// 主机别名

	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`
	// 风险项

	FailCount *uint64 `json:"FailCount,omitempty" name:"FailCount"`
	// 首次检测事件

	FirstScanTime *string `json:"FirstScanTime,omitempty" name:"FirstScanTime"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 最后检测时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 检测中状态

	MaxStatus *uint64 `json:"MaxStatus,omitempty" name:"MaxStatus"`
	// 通过项

	PassCount *uint64 `json:"PassCount,omitempty" name:"PassCount"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 风险项处理状态状态：0-未通过，1-通过

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type RiskDnsEvent struct {

	// 访问次数

	AccessCount *int64 `json:"AccessCount,omitempty" name:"AccessCount"`
	// 客户端ID

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// 命令行

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 访问域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 首次访问时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 处理状态；[0:待处理|2:已加白|3:非信任状态|4:已处理|5:已忽略]

	HandleStatus *int64 `json:"HandleStatus,omitempty" name:"HandleStatus"`
	// 主机ID

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机在线状态[OFFLINE:离线|ONLINE:在线|UNKNOWN:未知]

	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`
	// 事件Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 最近访问时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 进程ID

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 策略ID

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 命中策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 命中策略类型[-1:未知|0系统|1:用户]

	PolicyType *int64 `json:"PolicyType,omitempty" name:"PolicyType"`
	// 进程MD5

	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 保护级别[0:基础版|1:专业版|2:旗舰版]

	ProtectLevel *int64 `json:"ProtectLevel,omitempty" name:"ProtectLevel"`
	// 参考链接

	ReferenceLink *string `json:"ReferenceLink,omitempty" name:"ReferenceLink"`
	// 修复方案

	SuggestSolution *string `json:"SuggestSolution,omitempty" name:"SuggestSolution"`
	// 标签特性

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 威胁描述

	ThreatDesc *string `json:"ThreatDesc,omitempty" name:"ThreatDesc"`
	// 外网IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}

type DescribeExportMachinesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 机器所属地域。如：ap-guangzhou，ap-shanghai

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 机器所属业务ID列表

	ProjectIds []*uint64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
}

func (r *DescribeExportMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineInfoRequest struct {
	*tchttp.BaseRequest

	// Quuid , Uuid 必填一项

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 云镜客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMachineInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Broadcasts struct {

	// 发布时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 文章唯一id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 危险程度  0：无， 1：严重， 2: 高危， 3:中危， 4: 低危

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 副标题

	Subtitle *string `json:"Subtitle,omitempty" name:"Subtitle"`
	// 文章名字

	Title *string `json:"Title,omitempty" name:"Title"`
	// 类型：0=紧急通知，1=功能更新，2=行业荣誉，3=版本发布

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type ExportVulDetectionExcelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出本次漏洞检测Excel的任务Id（不同于入参的本次漏洞检测任务id）

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulDetectionExcelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDetectionExcelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineItemListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBaselineItemListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulHostTopInfo struct {

	// 主机名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// top评分

	Score *uint64 `json:"Score,omitempty" name:"Score"`
	// 漏洞等级与数量统计列表

	VulLevelList []*VulLevelCountInfo `json:"VulLevelList,omitempty" name:"VulLevelList"`
}

type DeleteLoginWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoginWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWarningHostConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWarningHostConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWarningHostConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineFileTamperRule struct {

	// 唯一id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则

	Rule []*FileTamperRule `json:"Rule,omitempty" name:"Rule"`
	// 规则类型 0 ：系统规则  1：用户规则

	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`
}

type VulInfoHostInfo struct {

	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机InstanceId

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 0 :漏洞不可自动修复，  1：可自动修复， 2：客户端已离线， 3：主机不是旗舰版只能手动修复， 4：机型不允许 ，5：修复中 ，6：已修复， 7：检测中, 9:修复失败, 10:已忽略 ,11:漏洞只支持linux不支持Windows, 12：漏洞只支持Windows不支持linux

	IsSupportAutoFix *uint64 `json:"IsSupportAutoFix,omitempty" name:"IsSupportAutoFix"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DeletePrivilegeRulesRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeletePrivilegeRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivilegeRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityDynamicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全事件消息数组。

		SecurityDynamics []*SecurityDynamic `json:"SecurityDynamics,omitempty" name:"SecurityDynamics"`
		// 记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityDynamicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityDynamicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWhiteListOrderRequest struct {
	*tchttp.BaseRequest

	// 到期时间,最小为1

	Deadline *uint64 `json:"Deadline,omitempty" name:"Deadline"`
	// 授权数量,最小为1 最大99999

	LicenseNum *uint64 `json:"LicenseNum,omitempty" name:"LicenseNum"`
	// 授权类型

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 地域, 1 广州 9新加坡, 默认为 1. 非必要情况不要选9

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 规则名称,大资产中心:asset_center

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 订单类型, 1 试用 2 赠送 3 体验 4 SSL-证书赠送 5 cvm赠送

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
}

func (r *CreateWhiteListOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWhiteListOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableExpertServiceDetailRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableExpertServiceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableExpertServiceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetEnvListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetEnvListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetEnvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryVulFixResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryVulFixResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryVulFixResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NonLocalLoginPlace struct {

	// 城市ID。

	City *uint64 `json:"City,omitempty" name:"City"`
	// 国家ID。

	Country *uint64 `json:"Country,omitempty" name:"Country"`
	// 事件ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 登录时间。

	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 机器名称。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 省份ID。

	Province *uint64 `json:"Province,omitempty" name:"Province"`
	// 登录IP。

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 登录状态
	// <li>NON_LOCAL_LOGIN：异地登录</li>
	// <li>NORMAL_LOGIN：正常登录</li>

	Status *string `json:"Status,omitempty" name:"Status"`
	// 用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 云镜客户端唯一标识Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeCanFixVulMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机漏洞修护信息列表

		VulInfo []*CanFixVulInfo `json:"VulInfo,omitempty" name:"VulInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCanFixVulMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCanFixVulMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WarningInfoObj struct {

	// 开始时间，格式: HH:mm

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 漏洞等级控制位（对应DB的十进制存储）

	ControlBit *uint64 `json:"ControlBit,omitempty" name:"ControlBit"`
	// 漏洞等级控制位二进制，每一位对应页面漏洞等级的开启关闭：低中高（0:关闭；1：开启），例如：101 → 同时勾选低+高

	ControlBits *string `json:"ControlBits,omitempty" name:"ControlBits"`
	// 配置的告警范围主机个数，前端用此判断展示提示信息

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 1: 关闭告警 0: 开启告警

	DisablePhoneWarning *uint64 `json:"DisablePhoneWarning,omitempty" name:"DisablePhoneWarning"`
	// 结束时间，格式: HH:mm

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 告警主机范围类型，0:全部主机，1:按所属项目选，2:按腾讯云标签选，3:按主机安全标签选，4:自选主机

	HostRange *int64 `json:"HostRange,omitempty" name:"HostRange"`
	// 离线类型判断时间 单位秒

	OfflineInterval *int64 `json:"OfflineInterval,omitempty" name:"OfflineInterval"`
	// 时区信息

	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
	// 事件告警类型；1：离线，2：木马，3：异常登录，4：爆破，5：漏洞（已拆分为9-12四种类型）6：高危命令，7：反弹sell，8：本地提权，9：应用漏洞，10：web-cms漏洞，11：应急漏洞，12：安全基线 ,13: 防篡改，14：恶意请求，15: 网络攻击，16：Windows系统漏洞，17：Linux软件漏洞，18：核心文件监控告警，19：客户端卸载告警。20：客户端离线告警

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type DescribeComponentStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件统计列表数据数组。

		ComponentStatistics []*ComponentStatistics `json:"ComponentStatistics,omitempty" name:"ComponentStatistics"`
		// 组件统计列表记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportVulsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专业周报漏洞数据数组。

		WeeklyReportVuls []*WeeklyReportVul `json:"WeeklyReportVuls,omitempty" name:"WeeklyReportVuls"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportVulsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageHost struct {

	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

type Tag struct {

	// 服务器数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 标签ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 标签名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CheckBashPolicyParamsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0=校验通过  1=规则名称校验不通过 2=正则表达式校验不通过

		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`
		// 校验信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckBashPolicyParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBashPolicyParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetRecentMachineInfoRequest struct {
	*tchttp.BaseRequest

	// 开始时间，如：2020-09-22

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 结束时间，如：2020-09-22

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeAssetRecentMachineInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetRecentMachineInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineDetectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselinePolicyDetect `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineDetectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbTestStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该用户当前是否正在灰度。

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbTestStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbTestStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表内容

		List []*ReverseShellRule `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineWeakPasswordListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列 [CreateTime|ModifyTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>WeakPassword - string - 是否必填：否 - 弱口令</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineWeakPasswordListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineWeakPasswordListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBanStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否弹窗提示信息 false: 关闭，true: 开启

		ShowTips *bool `json:"ShowTips,omitempty" name:"ShowTips"`
		// 阻断开关状态 0:关闭 1:开启

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBanStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesSimpleRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版 | Flagship : 旗舰版 | ProtectedMachines: 专业版+旗舰版 | UnFlagship : 非旗舰版 | PRO_POST_PAY：专业版按量计费 | PRO_PRE_PAY：专业版包年包月）</li>
	// <li>TagId - String - 是否必填：否 - 标签ID </li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 机器所属地域。如：ap-guangzhou，ap-shanghai

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 机器所属专区类型
	// CVM 云服务器
	// BM 黑石
	// ECM 边缘计算
	// LH 轻量应用服务器
	// Other 混合云专区

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 机器所属业务ID列表

	ProjectIds []*uint64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
}

func (r *DescribeMachinesSimpleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesSimpleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebHookEventKv struct {

	// 事件内容

	ControlBit *string `json:"ControlBit,omitempty" name:"ControlBit"`
	// 事件类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type DescribeMachineSnapshotRequest struct {
	*tchttp.BaseRequest

	// 分页个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页步长

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// cvm id 集合

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 查询快照类型0 最近一个 1所有

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeMachineSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineSnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBashPoliciesRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBashPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBashPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBaselinePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBaselinePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBaselinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBaselineRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBaselineRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBaselineRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmSettingstType struct {

	// 告警设置。

	AlarmSettings *AlarmSettings `json:"AlarmSettings,omitempty" name:"AlarmSettings"`
	// 告警类型。

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeAssetPortInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Ports []*AssetPortBaseInfo `json:"Ports,omitempty" name:"Ports"`
		// 记录总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetPortInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselinePolicyListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [RuleCount|ItemCount|HostCount]

	By *string `json:"By,omitempty" name:"By"`
	// <li>PolicyName - String - 是否必填：否 - 策略名称</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselinePolicyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselinePolicyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineOsListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMachineOsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineOsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表内容

		List []*BashPolicy `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRiskDnsPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRiskDnsPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskDnsPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogHistogram struct {

	// 统计周期内的日志条数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 按 period 取整后的 unix timestamp： 单位毫秒

	TimeStamp *int64 `json:"TimeStamp,omitempty" name:"TimeStamp"`
}

type AddLoginWhiteListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重复添加的提示列表

		DuplicateHosts []*DuplicateHosts `json:"DuplicateHosts,omitempty" name:"DuplicateHosts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLoginWhiteListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLoginWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmergencyResponseSosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEmergencyResponseSosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEmergencyResponseSosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBanRegionsRequest struct {
	*tchttp.BaseRequest

	// 阻断模式，STANDARD_MODE：标准阻断，DEEP_MODE：深度阻断

	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *DescribeBanRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBanRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulAgainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanVulAgainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanVulAgainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineRiskCntRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Uuids- String - 是否必填：否 - 主机uuid</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMachineRiskCntRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineRiskCntRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBaselinePolicyStateRequest struct {
	*tchttp.BaseRequest

	// 开启状态[1:开启|0:未开启]

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 策略Id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *ModifyBaselinePolicyStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBaselinePolicyStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLicenseBindsRequest struct {
	*tchttp.BaseRequest

	// 是否全部机器(当全部机器数大于当前订单可用授权数时,多余机器会被跳过)

	IsAll *bool `json:"IsAll,omitempty" name:"IsAll"`
	// 授权类型

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 需要绑定的机器quuid列表, 当IsAll = false 时必填,反之忽略该参数. 最大长度=2000

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *ModifyLicenseBindsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLicenseBindsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginWhiteCombinedInfo struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 仅在单台服务器时，返回服务器名称

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 白名单ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否对全局生效, 1：全局有效 0: 对指定主机列表生效'

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 地域字符串

	Locale *string `json:"Locale,omitempty" name:"Locale"`
	// 登陆地

	Locations *string `json:"Locations,omitempty" name:"Locations"`
	// 最近修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 白名单名字：IsLocal=1时固定为：全部服务器；单台机器时为机器内网IP，多台服务器时为服务器数量，如：11台

	Name *string `json:"Name,omitempty" name:"Name"`
	// 白名单地域

	Places []*Place `json:"Places,omitempty" name:"Places"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 白名单IP（多个IP逗号隔开）

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 白名单用户（多个用户逗号隔开）

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type AddMachineTagRequest struct {
	*tchttp.BaseRequest

	// 云服务器类型(CVM|BM)

	MArea *string `json:"MArea,omitempty" name:"MArea"`
	// 云服务器地区

	MRegion *string `json:"MRegion,omitempty" name:"MRegion"`
	// 云服务器ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 标签ID

	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
}

func (r *AddMachineTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMachineTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentInstallationTokenRequest struct {
	*tchttp.BaseRequest

	// token 过期时间

	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
}

func (r *DescribeAgentInstallationTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentInstallationTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostRiskLevelCount struct {

	// 高危个数

	HighCount *int64 `json:"HighCount,omitempty" name:"HighCount"`
	// 主机ID

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 主机名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 低危个数

	LowCount *int64 `json:"LowCount,omitempty" name:"LowCount"`
	// 中危个数

	MediumCount *int64 `json:"MediumCount,omitempty" name:"MediumCount"`
	// 严重个数

	SeriousCount *int64 `json:"SeriousCount,omitempty" name:"SeriousCount"`
}

type LicenseDetail struct {

	// 资源别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 0 初始化,1 自动续费,2 不自动续费

	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 购买时间

	BuyTime *string `json:"BuyTime,omitempty" name:"BuyTime"`
	// 截止日期

	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
	// 总授权数

	LicenseCnt *uint64 `json:"LicenseCnt,omitempty" name:"LicenseCnt"`
	// 授权ID

	LicenseId *uint64 `json:"LicenseId,omitempty" name:"LicenseId"`
	// 授权状态 0 未使用,1 部分使用, 2 已用完, 3 不可用

	LicenseStatus *uint64 `json:"LicenseStatus,omitempty" name:"LicenseStatus"`
	// 授权类型,0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 订单状态 1 正常 2隔离, 3销毁

	OrderStatus *uint64 `json:"OrderStatus,omitempty" name:"OrderStatus"`
	// 项目ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 订单资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 是否试用订单.

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
	// 平台标签

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
	// 任务ID ,默认0 ,查询绑定进度用

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 已使用授权数

	UsedLicenseCnt *uint64 `json:"UsedLicenseCnt,omitempty" name:"UsedLicenseCnt"`
}

type DescribeVersionStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基础版数量

		BasicVersionNum *uint64 `json:"BasicVersionNum,omitempty" name:"BasicVersionNum"`
		// 普惠版数量

		GeneralVersionNum *uint64 `json:"GeneralVersionNum,omitempty" name:"GeneralVersionNum"`
		// 专业版数量

		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`
		// 旗舰版数量

		UltimateVersionNum *uint64 `json:"UltimateVersionNum,omitempty" name:"UltimateVersionNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVersionStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVersionStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAttackEventsRequest struct {
	*tchttp.BaseRequest

	// 排序值 CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤参数。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>FilePath - String - 是否必填：否 - 路径筛选</li>
	// <li>VirusName - String - 是否必填：否 - 描述筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选-开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选-结束时间</li>
	// <li>Status - String - 是否必填：否 - 状态筛选</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 排序 方式 ，ASC，DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportAttackEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAttackEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetPortBaseInfo struct {

	// 绑定IP

	BindIp *string `json:"BindIp,omitempty" name:"BindIp"`
	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 所属用户组

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	//  附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 进程MD5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 操作系统信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 启动参数

	Param *string `json:"Param,omitempty" name:"Param"`
	// 父进程名称

	ParentProcessName *string `json:"ParentProcessName,omitempty" name:"ParentProcessName"`
	// 进程ID

	Pid *string `json:"Pid,omitempty" name:"Pid"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 父进程ID

	Ppid *string `json:"Ppid,omitempty" name:"Ppid"`
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 进程版本

	ProcessVersion *string `json:"ProcessVersion,omitempty" name:"ProcessVersion"`
	// 主机业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 端口协议

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 主机标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 进程TTY

	Teletype *string `json:"Teletype,omitempty" name:"Teletype"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 运行用户

	User *string `json:"User,omitempty" name:"User"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeProtectDirRelatedServerRequest struct {
	*tchttp.BaseRequest

	// 排序值

	By *string `json:"By,omitempty" name:"By"`
	// 过滤参数 ProtectStatus

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 唯一ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 分页条数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeProtectDirRelatedServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectDirRelatedServerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProVersionRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 主机唯一ID，对应CVM的uuid、BM的instanceId。

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 自动续费标识。取值范围：
	// <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	// <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyProVersionRenewFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProVersionRenewFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulInfoByCveId struct {

	// 修复支持情况：0-windows/linux均不支持修复 ;1-windows/linux 均支持修复 ;2-仅linux支持修复;3-仅windows支持修复

	FixSwitch *uint64 `json:"FixSwitch,omitempty" name:"FixSwitch"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

type DescribeAttackEventInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络攻击事件详情

		NetAttackEventInfo *NetAttackEventInfo `json:"NetAttackEventInfo,omitempty" name:"NetAttackEventInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackEventInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetMachineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetMachineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulEmergentMsgRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulEmergentMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulEmergentMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 开放端口号。Port和Uuid必填其一，使用Port表示查询该端口的列表信息。

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 云镜客户端唯一Uuid。Port和Uuid必填其一，使用Uuid表示，查询该主机列表信息。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeOpenPortsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpenPortsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProtectDirListRequest struct {
	*tchttp.BaseRequest

	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// DirName 网站名称
	// DirPath 网站防护目录地址

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 分页条数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// asc：升序/desc：降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeProtectDirListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectDirListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProtectMachine struct {

	// 机器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 机器名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 防护目录数

	SafeguardDirNum *uint64 `json:"SafeguardDirNum,omitempty" name:"SafeguardDirNum"`
}

type AddLoginWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLoginWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemRiskTopRequest struct {
	*tchttp.BaseRequest

	// 策略ID

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeBaselineItemRiskTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemRiskTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MisAlarmNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 异地登录事件Id数组。

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *MisAlarmNonlocalLoginPlacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MisAlarmNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetScanStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最近同步时间

		LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
		// 同步任务ID，任务ID为0表示没有进行中的同步任务

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetScanStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetScanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefencePluginEventRequest struct {
	*tchttp.BaseRequest

	// Keywords: ip或者主机名模糊匹配，Quuid，Exception、Status精确匹配

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 导出字段，默认全导出

	Where []*string `json:"Where,omitempty" name:"Where"`
}

func (r *ExportVulDefencePluginEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefencePluginEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProtectDirRelatedServerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网站关联服务器列表信息

		List []*ProtectDirRelatedServer `json:"List,omitempty" name:"List"`
		// 已开启防护总数

		ProtectServerCount *uint64 `json:"ProtectServerCount,omitempty" name:"ProtectServerCount"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProtectDirRelatedServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectDirRelatedServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeparateMalwaresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 隔离失败的id数组，若无则返回空数组

		FailedIds []*uint64 `json:"FailedIds,omitempty" name:"FailedIds"`
		// 隔离成功的id数组，若无则返回空数组

		SuccessIds []*uint64 `json:"SuccessIds,omitempty" name:"SuccessIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SeparateMalwaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeparateMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttackLogsRequest struct {
	*tchttp.BaseRequest

	// 日志ID数组，最大100条。

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 是否全部删除

	IsAll *bool `json:"IsAll,omitempty" name:"IsAll"`
}

func (r *DeleteAttackLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttackLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// top统计数据

		NetAttackTopInfo *NetAttackTopInfo `json:"NetAttackTopInfo,omitempty" name:"NetAttackTopInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackTopResponse) FromJsonString(s string) error {
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

type ExportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportNonlocalLoginPlacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulEventList struct {

	// 主机别名

	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`
	// 漏洞描述

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 漏洞事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否为专业版防护

	IfProfession *bool `json:"IfProfession,omitempty" name:"IfProfession"`
	// 最后检测时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 漏洞等级 1:低 2:中 3:高 4:提示

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 漏洞披露时间

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 事件状态:0: 待处理 1:忽略  3:已修复  5:检测中， 控制台仅处理0,1,3,5四种状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

type DeleteBaselineRuleIgnoreResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBaselineRuleIgnoreResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBaselineRuleIgnoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetMachineDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机详情

		MachineDetail *AssetMachineDetail `json:"MachineDetail,omitempty" name:"MachineDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetMachineDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	// <li>VulType - String 威胁类型 - 是否必填: 否</li>
	// <li>SrcIp - String 攻击源IP - 是否必填: 否</li>
	// <li>DstIp - String 攻击目标IP - 是否必填: 否</li>
	// <li>SrcPort - String 攻击源端口 - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 云主机机器ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机安全客户端ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAttackLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEmergencyResponseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应急响应列表

		List []*EmergencyResponseInfo `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEmergencyResponseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEmergencyResponseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDurationRequest struct {
	*tchttp.BaseRequest

	// 是否定时扫描

	TimingScan *bool `json:"TimingScan,omitempty" name:"TimingScan"`
	// 需要执行任务的主机数

	UuidCnt *uint64 `json:"UuidCnt,omitempty" name:"UuidCnt"`
}

func (r *DescribeTaskDurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskDurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePayPageBrowsingRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePayPageBrowsingRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePayPageBrowsingRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScanMalwareSettingRequest struct {
	*tchttp.BaseRequest

	// 是否开启恶意进程查杀[0:未开启,1:开启]

	EnableMemShellScan *int64 `json:"EnableMemShellScan,omitempty" name:"EnableMemShellScan"`
	// 1标准模式（只报严重、高危）、2增强模式（报严重、高危、中危）、3严格模式（报严重、高、中、低、提示）

	EngineType *uint64 `json:"EngineType,omitempty" name:"EngineType"`
	// 服务器分类：1:专业版服务器；2:自选服务器

	HostType *int64 `json:"HostType,omitempty" name:"HostType"`
	// 自选服务器时生效，主机quuid的string数组

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// 扫描模式 0 全盘扫描, 1 快速扫描

	ScanPattern *uint64 `json:"ScanPattern,omitempty" name:"ScanPattern"`
	// 超时时间单位 秒 默认3600 秒

	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitempty" name:"TimeoutPeriod"`
}

func (r *CreateScanMalwareSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScanMalwareSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDatabaseListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>User- string - 是否必填：否 - 运行用户</li>
	// <li>Ip - String - 是否必填：否 - 绑定IP</li>
	// <li>Port - Int - 是否必填：否 - 端口</li>
	// <li>Name - Int - 是否必填：否 - 数据库名称
	// 0:全部
	// 1:MySQL
	// 2:Redis
	// 3:Oracle
	// 4:MongoDB
	// 5:MemCache
	// 6:PostgreSQL
	// 7:HBase
	// 8:DB2
	// 9:Sybase
	// 10:TiDB</li>
	// <li>Proto - String - 是否必填：否 - 协议：1:TCP, 2:UDP, 3:未知</li>
	// <li>OsType - String - 是否必填：否 - 操作系统: linux/windows</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetDatabaseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MalwareRiskOverview struct {

	// 恶意文件数

	FileCount *int64 `json:"FileCount,omitempty" name:"FileCount"`
	// 影响主机数

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 是否首次扫描[false:否|true:是]

	IsFirstScan *bool `json:"IsFirstScan,omitempty" name:"IsFirstScan"`
	// 异常进程数

	ProcessCount *int64 `json:"ProcessCount,omitempty" name:"ProcessCount"`
	// 最后扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	// <li>Privilege - String - 是否必填：否 - 帐号类型（ORDINARY: 普通帐号 | SUPPER: 超级管理员帐号）</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Username表示，查询该用户名下列表信息。

	Username *string `json:"Username,omitempty" name:"Username"`
	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Uuid表示，查询该主机下列表信息。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 进程名。Uuid和ProcessName必填其一，使用ProcessName表示，查询该进程列表信息。

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 云镜客户端唯一Uuid。Uuid和ProcessName必填其一，使用Uuid表示，查询该主机列表信息。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeProcessesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProcessesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogKafkaAccessRequest struct {
	*tchttp.BaseRequest

	// 接入地址

	AccessAddr *string `json:"AccessAddr,omitempty" name:"AccessAddr"`
	// 接入方式，1公网域名接入，2支撑环境接入

	AccessType *uint64 `json:"AccessType,omitempty" name:"AccessType"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// 峰值带宽

	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`
	// 投递状态，1：健康，2：告警，3：异常

	DeliverStatus *int64 `json:"DeliverStatus,omitempty" name:"DeliverStatus"`
	// 日志投递类型配置细节

	DeliverTypeDetails []*DeliverTypeDetails `json:"DeliverTypeDetails,omitempty" name:"DeliverTypeDetails"`
	// 磁盘容量

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 0不用密码，1有密码

	HasPwd *uint64 `json:"HasPwd,omitempty" name:"HasPwd"`
	// kafka版本

	InsVersion *string `json:"InsVersion,omitempty" name:"InsVersion"`
	// 实例名称 如 云镜测试环境

	KafkaEnvName *string `json:"KafkaEnvName,omitempty" name:"KafkaEnvName"`
	// 实例ID

	KafkaId *string `json:"KafkaId,omitempty" name:"KafkaId"`
	// 密码，aes加密

	Pwd *string `json:"Pwd,omitempty" name:"Pwd"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// vpcid地址

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *ModifyLogKafkaAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogKafkaAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostDesc struct {

	// 机器IP:已销毁的服务器IP为空

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 机器名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 公网IP:已销毁的服务器IP为空

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 云镜客户端ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 标签信息数组

	Tags []*MachineTag `json:"Tags,omitempty" name:"Tags"`
	// 主机ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetPortInfoListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime|StartTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Port - uint64 - 是否必填：否 - 端口</li>
	// <li>Ip - String - 是否必填：否 - 绑定IP</li>
	// <li>ProcessName - String - 是否必填：否 - 监听进程</li>
	// <li>Pid - uint64 - 是否必填：否 - PID</li>
	// <li>User - String - 是否必填：否 - 运行用户</li>
	// <li>Group - String - 是否必填：否 - 所属用户组</li>
	// <li>Ppid - uint64 - 是否必填：否 - PPID</li>
	// <li>Proto - string - 是否必填：否 - tcp/udp或“”(空字符串筛选未知状态)</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>RunTimeStart - String - 是否必填：否 - 运行开始时间</li>
	// <li>RunTimeEnd - String - 是否必填：否 - 运行结束时间</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetPortInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ModifyMachineRemarkRequest struct {
	*tchttp.BaseRequest

	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyMachineRemarkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMachineRemarkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulFixStatusInfo struct {

	// 漏洞修复失败主机数量

	FailCnt *uint64 `json:"FailCnt,omitempty" name:"FailCnt"`
	// 修复成功的数量

	FixSuccessCnt *uint64 `json:"FixSuccessCnt,omitempty" name:"FixSuccessCnt"`
	// 漏洞对应主机修复状态

	HostList []*VulFixStatusHostInfo `json:"HostList,omitempty" name:"HostList"`
	// 漏洞修复进度 1-100；

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type DeleteMaliciousRequestsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaliciousRequestsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExpertServiceOrderListRequest struct {
	*tchttp.BaseRequest

	// <li>InquireType- String - 是否必填：否 - 订单类型过滤，</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 分页条数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页步长

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeExpertServiceOrderListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExpertServiceOrderListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMalwareTimingScanSettingsRequest struct {
	*tchttp.BaseRequest

	// 是否自动隔离 1隔离 0 不隔离

	AutoIsolation *uint64 `json:"AutoIsolation,omitempty" name:"AutoIsolation"`
	// 检测模式 0 全盘检测  1快速检测

	CheckPattern *uint64 `json:"CheckPattern,omitempty" name:"CheckPattern"`
	// 扫描周期 默认每天 1

	Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`
	// 启发引擎开关 0 关闭 1开启

	EnableInspiredEngine *uint64 `json:"EnableInspiredEngine,omitempty" name:"EnableInspiredEngine"`
	// 是否开启恶意进程查杀[0:未开启,1:开启]

	EnableMemShellScan *uint64 `json:"EnableMemShellScan,omitempty" name:"EnableMemShellScan"`
	// 定时检测开关 0 关闭 1开启

	EnableScan *uint64 `json:"EnableScan,omitempty" name:"EnableScan"`
	// 检测周期 超时结束时间，如：04:00:00

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 1标准模式（只报严重、高危）、2增强模式（报严重、高危、中危）、3严格模式（报严重、高、中、低、提示）

	EngineType *uint64 `json:"EngineType,omitempty" name:"EngineType"`
	// 是否全部服务器 1 全部 2 自选

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 是否杀掉进程 1杀掉 0不杀掉

	KillProcess *uint64 `json:"KillProcess,omitempty" name:"KillProcess"`
	// 监控模式 0 标准 1深度

	MonitoringPattern *uint64 `json:"MonitoringPattern,omitempty" name:"MonitoringPattern"`
	// 自选服务器时必须 主机quuid的string数组

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// 实时监控 0 关闭 1开启

	RealTimeMonitoring *uint64 `json:"RealTimeMonitoring,omitempty" name:"RealTimeMonitoring"`
	// 检测周期 开始时间，如：02:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *ModifyMalwareTimingScanSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMalwareTimingScanSettingsRequest) FromJsonString(s string) error {
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

type DescribeVulEffectModulesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// <li>Uuid - String数组 - Uuid串数组</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页limit 最大100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *DescribeVulEffectModulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulEffectModulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetMachineDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetMachineDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetMachineDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessStatistics struct {

	// 主机数量。

	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
	// 进程名。

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
}

type WeeklyReportNonlocalLoginPlace struct {

	// 城市ID。

	City *uint64 `json:"City,omitempty" name:"City"`
	// 国家ID。

	Country *uint64 `json:"Country,omitempty" name:"Country"`
	// 登录时间。

	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 省份ID。

	Province *uint64 `json:"Province,omitempty" name:"Province"`
	// 源IP。

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 用户名。

	Username *string `json:"Username,omitempty" name:"Username"`
}

type DescribeAssetProcessInfoListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime|StartTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 进程名</li>
	// <li>User - String - 是否必填：否 - 进程用户</li>
	// <li>Group - String - 是否必填：否 - 进程用户组</li>
	// <li>Pid - uint64 - 是否必填：否 - 进程ID</li>
	// <li>Ppid - uint64 - 是否必填：否 - 父进程ID</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>Status - string - 是否必填：否 - 进程状态：
	// 1:R 可执行
	// 2:S 可中断
	// 3:D 不可中断
	// 4:T 暂停状态或跟踪状态
	// 5:Z 僵尸状态
	// 6:X 将被销毁</li>
	// <li>RunTimeStart - String - 是否必填：否 - 运行开始时间</li>
	// <li>RunTimeEnd - String - 是否必填：否 - 运行结束时间</li>
	// <li>InstallByPackage - uint64 - 是否必填：否 - 是否包安装：0否，1是</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetProcessInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareWhiteListAffectListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [CreateTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>HostIp - String - 是否必填：否 - 主机ip查询 </li>
	// <li>FileName - String - 是否必填：否 - 文件名称查询 </li>
	// <li>FileDirectory - String - 是否必填：否 - 文件目录查询 </li>
	// <li>FileExtension - String - 是否必填：否 - 文件后缀查询 </li>
	// <li>Mode - String - 是否必填：否 - 规则类型 0 MD5,1自定义</li>
	// <li>Md5 - String - 是否必填：否 - MD5查询 </li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
	// 白名单规则id

	WhiteListId *uint64 `json:"WhiteListId,omitempty" name:"WhiteListId"`
}

func (r *DescribeMalwareWhiteListAffectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareWhiteListAffectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// <li>Status - int - 是否必填：否 - 状态筛选1:正常登录；2：异地登录</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportNonlocalLoginPlacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRiskDnsEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRiskDnsEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRiskDnsEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceOpenProVersionPrepaidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预支费用的折扣价，单位：元。

		DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
		// 预支费用的原价，单位：元。

		OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceOpenProVersionPrepaidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceOpenProVersionPrepaidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileTamperRule struct {

	// 执行动作 跳过：skip，告警：alert

	Action *string `json:"Action,omitempty" name:"Action"`
	// 文件监控内容 write read read-write

	FileAction *string `json:"FileAction,omitempty" name:"FileAction"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 被访问文件路径

	Target *string `json:"Target,omitempty" name:"Target"`
}

type DescribeAssetInitServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Services []*AssetInitServiceBaseInfo `json:"Services,omitempty" name:"Services"`
		// 总数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetInitServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetInitServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAttackLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已废弃

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务ID 可通过ExportTasks接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAttackLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionCompareChartRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVersionCompareChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVersionCompareChartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 白名单列表

		WhiteList []*MalwareWhiteListInfo `json:"WhiteList,omitempty" name:"WhiteList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwareWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulEffectHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已废弃

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务Id , 可通过ExportTasks 接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulEffectHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BashRule struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否处理之前的事件 0: 不处理 1:处理

	DealOldEvents *uint64 `json:"DealOldEvents,omitempty" name:"DealOldEvents"`
	// 规则描述

	Decription *string `json:"Decription,omitempty" name:"Decription"`
	// 主机IP

	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 危险等级(0 ：无 1: 高危 2:中危 3: 低危)

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 正则表达式

	Rule *string `json:"Rule,omitempty" name:"Rule"`
	// 状态 (0: 有效 1: 无效)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 客户端ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 生效服务器的uuid数组

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
	// 0=黑名单 1=白名单

	White *uint64 `json:"White,omitempty" name:"White"`
}

type DescribeStrategyExistRequest struct {
	*tchttp.BaseRequest

	// 策略名

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
}

func (r *DescribeStrategyExistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStrategyExistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetMachineDetailRequest struct {
	*tchttp.BaseRequest

	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetMachineDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有地域列表(包含以上所有地域)

		ALL []*RegionInfo `json:"ALL,omitempty" name:"ALL"`
		// BM 黑石机器地域列表

		BM []*RegionInfo `json:"BM,omitempty" name:"BM"`
		// CVM 云服务器地域列表

		CVM []*RegionInfo `json:"CVM,omitempty" name:"CVM"`
		// ECM 边缘计算服务器地域列表

		ECM []*RegionInfo `json:"ECM,omitempty" name:"ECM"`
		// LH 轻量应用服务器地域列表

		LH []*RegionInfo `json:"LH,omitempty" name:"LH"`
		// Other 混合云地域列表

		Other []*RegionInfo `json:"Other,omitempty" name:"Other"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseBackupListRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Status - Int - 是否必填：否 - 通过勒索状态查询：0未勒索，1已勒索 </li>
	// <li>CreateTimeBegin - string - 是否必填：否 - 创建时间开始 </li>
	// <li>CreateTimeEnd - string - 是否必填：否 - 创建时间结束 </li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeRansomDefenseBackupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseBackupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBanWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBanWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBanWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProtectNetListRequest struct {
	*tchttp.BaseRequest

	// 排序字段 StartTime，EndTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Keyword- String - 是否必填：否 - 关键词过滤，</li>
	// <li>Uuids - String - 是否必填：否 - 主机id过滤</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序步长

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方法

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeProtectNetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectNetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentVulsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机漏洞信息

		AgentVuls []*AgentVul `json:"AgentVuls,omitempty" name:"AgentVuls"`
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentVulsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsListRequest struct {
	*tchttp.BaseRequest

	// 排序字段：AccessCount-请求次数。MergeTime-最近请求时间

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Url - String - 是否必填：否 - Url筛选</li>
	// <li>Status - String - 是否必填：否 - 状态筛选0:待处理；2:信任；3:不信任</li>
	// <li>MergeBeginTime - String - 是否必填：否 - 最近访问开始时间</li>
	// <li>MergeEndTime - String - 是否必填：否 - 最近访问结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：根据请求次数排序：asc-升序/desc-降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRiskDnsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulEffectModulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 影响主机列表

		VulEffectModuleInfo []*VulEffectModuleInfo `json:"VulEffectModuleInfo,omitempty" name:"VulEffectModuleInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulEffectModulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulEffectModulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWarningListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取告警列表

		WarningInfoList []*WarningInfoObj `json:"WarningInfoList,omitempty" name:"WarningInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWarningListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarningListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 专业周报开始时间。

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionPrepaidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单ID列表。

		DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVersionPrepaidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenProVersionPrepaidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BruteAttackRuleList struct {

	// 规则是否为空，为空则填充默认规则

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 爆破事件失败次数

	LoginFailTimes *uint64 `json:"LoginFailTimes,omitempty" name:"LoginFailTimes"`
	// 爆破事件失败次数（默认规则）

	LoginFailTimesDefault *uint64 `json:"LoginFailTimesDefault,omitempty" name:"LoginFailTimesDefault"`
	// 爆破事件发生的时间范围，单位：秒

	TimeRange *uint64 `json:"TimeRange,omitempty" name:"TimeRange"`
	// 爆破事件发生的时间范围，单位：秒（默认规则）

	TimeRangeDefault *uint64 `json:"TimeRangeDefault,omitempty" name:"TimeRangeDefault"`
}

type DescribeRansomDefenseMachineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表

		List []*RansomDefenseStrategyMachineBackupInfo `json:"List,omitempty" name:"List"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRansomDefenseMachineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseEventsListRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>HostName- string- 主机名称</li>
	// <li>Status - Uint64 - 0待处理，1已处理，2已信任</li>
	// <li>HostIp- String - 主机ip</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRansomDefenseEventsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseEventsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseBindTaskDetail struct {

	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 云服务器UUID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 0 执行中, 1 成功,2失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type ApiKey struct {

	// 创建时间(时间戳)

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 密钥ID

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 状态(2:有效, 3:禁用, 4:已删除)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type LicenseUnBindRsp struct {

	// 失败原因

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// QUUID 云服务器uuid,轻量服务器uuid,边缘计算 uuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type ModifyAutoOpenProVersionConfigRequest struct {
	*tchttp.BaseRequest

	// 自动加购的订单是否自动续费,默认0 ,0关闭, 1开启

	AutoRepurchaseRenewSwitch *uint64 `json:"AutoRepurchaseRenewSwitch,omitempty" name:"AutoRepurchaseRenewSwitch"`
	// 自动加购/扩容授权开关,默认 1, 0关闭, 1开启

	AutoRepurchaseSwitch *uint64 `json:"AutoRepurchaseSwitch,omitempty" name:"AutoRepurchaseSwitch"`
	// 加固防护模式
	// PROVERSION_POSTPAY 专业版-按量计费
	// PROVERSION_PREPAY 专业版-包年包月
	// FLAGSHIP_PREPAY 旗舰版-包年包月

	ProtectType *string `json:"ProtectType,omitempty" name:"ProtectType"`
	// 设置自动开通状态。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAutoOpenProVersionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoOpenProVersionConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVerWithQuuidsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVerWithQuuidsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenProVerWithQuuidsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetAttackSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增资产自动包含 0 不包含 1包含

		AutoInclude *uint64 `json:"AutoInclude,omitempty" name:"AutoInclude"`
		// 自选排除主机

		ExcludeInstanceIds []*string `json:"ExcludeInstanceIds,omitempty" name:"ExcludeInstanceIds"`
		// 自选主机

		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
		// 0 新增告警事件默认待处理，1新增告警事件默认已处理，3新增告警事件默认忽略

		NetAttackAlarmStatus *uint64 `json:"NetAttackAlarmStatus,omitempty" name:"NetAttackAlarmStatus"`
		// 0 关闭网络攻击检测，1开启网络攻击检测

		NetAttackEnable *uint64 `json:"NetAttackEnable,omitempty" name:"NetAttackEnable"`
		// 1 全部旗舰版主机，0 InstanceIds列表主机

		Scope *uint64 `json:"Scope,omitempty" name:"Scope"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetAttackSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetAttackSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetInitServiceListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	// <li>Status- string - 是否必填：否 - 默认启用状态：0未启用， 1启用 仅linux</li>
	// <li>Type- string - 是否必填：否 - 类型：类型 仅windows：
	// 1:编码器
	// 2:IE插件
	// 3:网络提供者
	// 4:镜像劫持
	// 5:LSA提供者
	// 6:KnownDLLs
	// 7:启动执行
	// 8:WMI
	// 9:计划任务
	// 10:Winsock提供者
	// 11:打印监控器
	// 12:资源管理器
	// 13:驱动服务
	// 14:登录</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetInitServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetInitServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditBashRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditBashRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditBashRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportProtectDirListRequest struct {
	*tchttp.BaseRequest

	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// DirName 网站名称
	// DirPath 网站防护目录地址

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// asc：升序/desc：降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportProtectDirListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportProtectDirListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>StrategyId- Uint64 - 基线策略id</li>
	// <li>Status - Uint64 - 处理状态1已通过 0未通过</li>
	// <li>Level - Uint64[] - 处理状态1已通过 0未通过</li>BaselineName
	// <li>BaselineName  - String - 基线名称</li>
	// <li>Quuid- String - 主机quuid</li>
	// <li>Uuid- String - 主机uuid</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBaselineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProductStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryAccountsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 云镜客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeHistoryAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHistoryAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetAttackWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 白名单列表

		WhiteList []*NetAttackWhiteRule `json:"WhiteList,omitempty" name:"WhiteList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetAttackWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetAttackWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebPageProtectStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件篡改信息

		FileTamperNum []*ProtectStat `json:"FileTamperNum,omitempty" name:"FileTamperNum"`
		// 防护文件分类信息

		ProtectFileType []*ProtectStat `json:"ProtectFileType,omitempty" name:"ProtectFileType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebPageProtectStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebPageProtectStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostLoginListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 登录审计列表

		HostLoginList []*HostLoginList `json:"HostLoginList,omitempty" name:"HostLoginList"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostLoginListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostLoginListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口统计数据列表

		OpenPortStatistics []*OpenPortStatistics `json:"OpenPortStatistics,omitempty" name:"OpenPortStatistics"`
		// 端口统计列表总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenPortStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpenPortStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityBroadcastsRequest struct {
	*tchttp.BaseRequest

	// 筛选发布日期：开始时间

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 过滤安全播报类型：0-紧急通知，1-功能更新，2-行业荣誉，3-版本发布

	BroadcastType *string `json:"BroadcastType,omitempty" name:"BroadcastType"`
	// 筛选发布日期：结束时间

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 需要返回的数量，默认为10 ，0=全部

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSecurityBroadcastsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityBroadcastsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetAttackWhiteRule struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否处理之前的事件 0: 不处理 1:处理

	DealOldEvents *uint64 `json:"DealOldEvents,omitempty" name:"DealOldEvents"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 主机quuid 多个用;隔开

	Quuids *string `json:"Quuids,omitempty" name:"Quuids"`
	// 0: 一组quuid 1: 所有主机

	Scope *uint64 `json:"Scope,omitempty" name:"Scope"`
	// 来源IP 单IP:1.1.1.1  IP范围:1.1.1.1-1.1.2.1  IP范围：1.1.1.0/24 多个用;隔开

	SrcIP *string `json:"SrcIP,omitempty" name:"SrcIP"`
}

type CreateRansomDefenseStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRansomDefenseStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRansomDefenseStrategyResponse) FromJsonString(s string) error {
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

type DeleteNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNonlocalLoginPlacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WeeklyReportMalware struct {

	// 木马文件路径。

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 木马发现时间。

	FindTime *string `json:"FindTime,omitempty" name:"FindTime"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 木马文件MD5值。

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 当前木马状态。
	// <li>UN_OPERATED：未处理</li>
	// <li>SEGREGATED：已隔离</li>
	// <li>TRUSTED：已信任</li>
	// <li>SEPARATING：隔离中</li>
	// <li>RECOVERING：恢复中</li>

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeAssetEnvListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Envs []*AssetEnvBaseInfo `json:"Envs,omitempty" name:"Envs"`
		// 总数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetEnvListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetEnvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBashEventsNewRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBashEventsNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBashEventsNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Process struct {

	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 进程路径。

	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`
	// 唯一ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机内网IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 进程Pid。

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 所属平台。
	// <li>WIN32：windows32位</li>
	// <li>WIN64：windows64位</li>
	// <li>LINUX32：Linux32位</li>
	// <li>LINUX64：Linux64位</li>

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 进程Ppid。

	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`
	// 进程名。

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程用户名。

	Username *string `json:"Username,omitempty" name:"Username"`
	// 云镜客户端唯一UUID。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeBashEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 高危命令事件列表

		List []*BashEvent `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBashPolicyStatusRequest struct {
	*tchttp.BaseRequest

	// 是否禁用

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyBashPolicyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBashPolicyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExpertServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全管家数据

		List []*SecurityButlerInfo `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExpertServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExpertServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenAttackHotspotRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScreenAttackHotspotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenAttackHotspotRequest) FromJsonString(s string) error {
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

type TestWebHookRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestWebHookRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestWebHookRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PushTaskResult struct {

	// 下发失败原因

	FailureReason *string `json:"FailureReason,omitempty" name:"FailureReason"`
	// 是否下发成功

	IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`
	// 对应任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务种类：Vul-检测漏洞，Malware-文件查杀，Baseline-基线检测，AutoBan-自动阻断

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

type DeleteRiskDnsEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRiskDnsEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRiskDnsEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppProcessListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 进程列表

		Process []*AssetAppProcessInfo `json:"Process,omitempty" name:"Process"`
		// 分区总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetAppProcessListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Id struct {

	// 对象id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type PrivilegeEventInfo struct {

	// 执行命令

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 发生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 进程路径

	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`
	// 危害描述信息

	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
	// 主机内网IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 数据ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 机器名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机在线状态 OFFLINE  ONLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 主机外网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 处理时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 权限列表|隔开

	NewCaps *string `json:"NewCaps,omitempty" name:"NewCaps"`
	// 父进程用户组

	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`
	// 父进程名

	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`
	// 父进程路径

	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`
	// 父进程用户名

	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`
	// 进程文件权限

	ProcFilePrivilege *string `json:"ProcFilePrivilege,omitempty" name:"ProcFilePrivilege"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程树 json  pid:进程id，exe:文件路径 ，account:进程所属用组和用户 ,cmdline:执行命令，ssh_service: SSH服务ip, ssh_soure:登录源

	PsTree *string `json:"PsTree,omitempty" name:"PsTree"`
	// 主机ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 参考链接

	References []*string `json:"References,omitempty" name:"References"`
	// 处理状态：0-待处理 2-白名单 3-已处理 4-已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 建议方案

	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 用户组

	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 云镜ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type AssetJarBaseInfo struct {

	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// Jar包ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	//  附加信息
	//

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 服务器IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 服务器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 服务器外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// Jar包Md5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 是否可执行：0未知，1是，2否

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 类型：1应用程序，2系统类库，3Web服务自带库，8:其他，

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type CreateMaliciousRequestWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaliciousRequestWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaliciousRequestWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异地登录信息数组。

		NonLocalLoginPlaces []*NonLocalLoginPlace `json:"NonLocalLoginPlaces,omitempty" name:"NonLocalLoginPlaces"`
		// 记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNonlocalLoginPlacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityButlerInfo struct {

	// 服务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 主机Ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 数据id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 订单id

	OrderId *uint64 `json:"OrderId,omitempty" name:"OrderId"`
	// cvm id

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机风险数

	RiskCount *uint64 `json:"RiskCount,omitempty" name:"RiskCount"`
	// 服务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 服务状态 0-服务中,1-已到期 2已销毁

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 主机 uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeScreenMachinesRequest struct {
	*tchttp.BaseRequest

	// IP 支持内网IP,主机别名

	MachineIpOrAlias *string `json:"MachineIpOrAlias,omitempty" name:"MachineIpOrAlias"`
	// 主机区域

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 风险状态类型：0：全部类型，1：风险主机，2：潜在风险主机 ， 3：已关机/离线主机，4：无风险主机

	SecurityStatus *uint64 `json:"SecurityStatus,omitempty" name:"SecurityStatus"`
}

func (r *DescribeScreenMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefencePluginDetailRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：Keywords: ip或者主机名,Exception，Status精确匹配

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 数据限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 数据偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeVulDefencePluginDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefencePluginDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBashEventsNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBashEventsNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBashEventsNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportIgnoreBaselineRuleRequest struct {
	*tchttp.BaseRequest

	// 检测项名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *ExportIgnoreBaselineRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportIgnoreBaselineRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetScanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最新结束同步时间

		LatestEndTime *string `json:"LatestEndTime,omitempty" name:"LatestEndTime"`
		// 最新开始同步时间

		LatestStartTime *string `json:"LatestStartTime,omitempty" name:"LatestStartTime"`
		// 枚举值有(大写)：NOTASK（没有同步任务），SYNCING（同步中），FINISHED（同步完成）

		State *string `json:"State,omitempty" name:"State"`
		// 任务ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncAssetScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventlogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表数据

		List []*Eventlog `json:"List,omitempty" name:"List"`
		// 事件数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventlogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventlogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginWhiteListsRule struct {

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 白名单生效的机器

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 规则ID，用于更新规则

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否对全局生效

	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 加白地域

	Places []*Place `json:"Places,omitempty" name:"Places"`
	// 加白源IP，支持网段，多个IP以逗号隔开

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 加白用户名，多个用户名以逗号隔开

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type DeleteBruteAttacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBruteAttacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackSourceEventsRequest struct {
	*tchttp.BaseRequest

	// 开始日期

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 结束日期

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 接口DescribeAttackSource 返回的EventInfoParam

	EventInfoParam *string `json:"EventInfoParam,omitempty" name:"EventInfoParam"`
	// 限制分页条数默认10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 起始步长默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAttackSourceEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackSourceEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineRuleIgnoreListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [HostCount]

	By *string `json:"By,omitempty" name:"By"`
	// <li>RuleName - String - 是否必填：否 - 规则名称</li>
	// <li>ItemId- int - 是否必填：否 - 检测项Id</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineRuleIgnoreListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleIgnoreListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceListRequest struct {
	*tchttp.BaseRequest

	// 排序列，严格相等：PublishTime发布时间

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件：Level、Keywords

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 数据限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 数据偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序，大小写无关：asc 升序，desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVulDefenceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditReverseShellRulesRequest struct {
	*tchttp.BaseRequest

	// 目标IP

	DestIp *string `json:"DestIp,omitempty" name:"DestIp"`
	// 目标端口

	DestPort *string `json:"DestPort,omitempty" name:"DestPort"`
	// 事件列表和详情点击加白时关联的事件id (新增规则时请留空)

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 规则ID(新增时请留空)

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则(默认否)

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 客户端ID数组

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *EditReverseShellRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditReverseShellRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineRiskItem struct {

	// 影响服务器数

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 检测项Id

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
	// 检测项名字

	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`
	// 风险等级

	Level *int64 `json:"Level,omitempty" name:"Level"`
}

type DeleteReverseShellRulesRequest struct {
	*tchttp.BaseRequest

	// ID数组. (最大100条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteReverseShellRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRansomDefenseStrategyRequest struct {
	*tchttp.BaseRequest

	// 备份模式： 0按周，1按天

	BackupType *uint64 `json:"BackupType,omitempty" name:"BackupType"`
	// 策略备注

	Description *string `json:"Description,omitempty" name:"Description"`
	// 排除目录，;分隔

	ExcludeDir *string `json:"ExcludeDir,omitempty" name:"ExcludeDir"`
	// 定时快照执行时间（0-23）：01:00;23:00

	Hour *string `json:"Hour,omitempty" name:"Hour"`
	// 策略ID，填ID时修改策略，否则新增策略

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 包含目录，;分隔

	IncludeDir *string `json:"IncludeDir,omitempty" name:"IncludeDir"`
	// 是否全部机器生效：0否，1是

	IsAll *uint64 `json:"IsAll,omitempty" name:"IsAll"`
	// 绑定主机列表

	Machines []*RansomDefenseStrategyMachineInfo `json:"Machines,omitempty" name:"Machines"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 保留天数：0永久

	SaveDay *uint64 `json:"SaveDay,omitempty" name:"SaveDay"`
	// 是否开启：0关，1开

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 定时任务执行星期天数（1-7）：1;2;7

	Weekday *string `json:"Weekday,omitempty" name:"Weekday"`
}

func (r *CreateRansomDefenseStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRansomDefenseStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchTemplatesRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSearchTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAttackLogsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	// <li>VulType - String 威胁类型 - 是否必填: 否</li>
	// <li>SrcIp - String 攻击源IP - 是否必填: 否</li>
	// <li>DstIp - String 攻击目标IP - 是否必填: 否</li>
	// <li>SrcPort - String 攻击源端口 - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 云主机机器ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机安全客户端ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAttackLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAttackLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeysLocalStorageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 键列表
		// 注意：此字段可能返回 null，表示取不到有效值。

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

type ModifyFileTamperEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFileTamperEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFileTamperEventsResponse) FromJsonString(s string) error {
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

type ExportBruteAttacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBruteAttacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRiskEventsStatusRequest struct {
	*tchttp.BaseRequest

	// 排除的事件id,当操作全部事件时，需要排除这次id

	ExcludeId []*uint64 `json:"ExcludeId,omitempty" name:"ExcludeId"`
	// 过滤条件。RiskType为 MALWARE时
	// 1、当RiskType为 MALWARE时：
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>FilePath - String - 是否必填：否 - 路径筛选</li>
	// <li>VirusName - String - 是否必填：否 - 描述筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选-开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选-结束时间</li>
	// <li>Status - String - 是否必填：否 - 状态筛选 4待处理,5信任,6已隔离,10隔离中,11恢复隔离中</li>
	// RiskType 为PROCESS时:
	// 过滤条件。
	// <li>IpOrName - String - 是否必填：否 - 主机IP或主机名</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名</li>
	// <li>BeginTime - String - 是否必填：否 - 进程启动时间-开始</li>
	// <li>EndTime - String - 是否必填：否 - 进程启动时间-结束</li>
	// <li>Status - String - 是否必填：否 - 状态筛选 0待处理；1查杀中;2已查杀3已退出;4已信任</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要修改的事件id 数组，支持批量

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 当RiskType 为异地登录且ids为空时，可以修改所有来源ip的事件的状态

	Ip []*string `json:"Ip,omitempty" name:"Ip"`
	// 当Operate 是木马隔离时，表示是否要杀进程，其他操作无效

	KillProcess *bool `json:"KillProcess,omitempty" name:"KillProcess"`
	// 操作-0:标记已处理,1:忽略,2:删除记录,3:木马隔离,4:木马恢复隔离,5:木马信任,6:木马取消信任,7:查杀异常进程

	Operate *uint64 `json:"Operate,omitempty" name:"Operate"`
	// 操作事件类型，文件查杀：MALWARE，异常登录：HOST_LOGIN，密码破解：BRUTE_ATTACK，恶意请求：MALICIOUS_REQUEST，高危命令：BASH_EVENT，本地提权：PRIVILEGE_EVENT，反弹shell：REVERSE_SHELL. 异常进程:PROCESS

	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
	// 是否更新全部，即是否对所有的事件进行操作，当ids 不为空时，此参数无效

	UpdateAll *bool `json:"UpdateAll,omitempty" name:"UpdateAll"`
}

func (r *ModifyRiskEventsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskEventsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDatabaseCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各数据库数量

		Databases []*AssetKeyVal `json:"Databases,omitempty" name:"Databases"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDatabaseCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLicenseUnBindsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 只有解绑失败的才有该值.

		ErrMsg []*LicenseUnBindRsp `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLicenseUnBindsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLicenseUnBindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务器列表

		Hosts []*HostDesc `json:"Hosts,omitempty" name:"Hosts"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginWhiteHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginWhiteHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulInfoCvssResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CVSS信息

		CVSS *string `json:"CVSS,omitempty" name:"CVSS"`
		// 漏洞CVEID

		CveId *string `json:"CveId,omitempty" name:"CveId"`
		// cvss详情

		CveInfo *string `json:"CveInfo,omitempty" name:"CveInfo"`
		// Cvss分数

		CvssScore *uint64 `json:"CvssScore,omitempty" name:"CvssScore"`
		// cvss 分数 浮点型

		CvssScoreFloat *float64 `json:"CvssScoreFloat,omitempty" name:"CvssScoreFloat"`
		// 已防御的攻击次数

		DefenseAttackCount *uint64 `json:"DefenseAttackCount,omitempty" name:"DefenseAttackCount"`
		// 漏洞描述信息

		Description *string `json:"Description,omitempty" name:"Description"`
		// 修复是否支持：0-windows/linux均不支持修复 ;1-windows/linux 均支持修复 ;2-仅linux支持修复;3-仅windows支持修复

		FixSwitch *int64 `json:"FixSwitch,omitempty" name:"FixSwitch"`
		// 漏洞标签 多个逗号分割

		Labels *string `json:"Labels,omitempty" name:"Labels"`
		// 发布时间

		PublicDate *string `json:"PublicDate,omitempty" name:"PublicDate"`
		// 参考链接

		Reference *string `json:"Reference,omitempty" name:"Reference"`
		// 修复方案

		RepairPlan *string `json:"RepairPlan,omitempty" name:"RepairPlan"`
		// 全网修复成功次数, 不支持自动修复的漏洞默认返回0

		SuccessFixCount *uint64 `json:"SuccessFixCount,omitempty" name:"SuccessFixCount"`
		// 漏洞id

		VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
		// 危害等级：1-低危；2-中危；3-高危；4-严重

		VulLevel *uint64 `json:"VulLevel,omitempty" name:"VulLevel"`
		// 漏洞名称

		VulName *string `json:"VulName,omitempty" name:"VulName"`
		// 漏洞分类 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞

		VulType *uint64 `json:"VulType,omitempty" name:"VulType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulInfoCvssResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulInfoCvssResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebPageProtectStatRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWebPageProtectStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebPageProtectStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MiniSite struct {

	// 站点ID。

	SiteId *uint64 `json:"SiteId,omitempty" name:"SiteId"`
	// 站点Url。

	Url *string `json:"Url,omitempty" name:"Url"`
}

type DescribeMachineFileTamperRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询主机相关核心文件监控规则详情

		List []*MachineFileTamperRule `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineFileTamperRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineFileTamperRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChargePrepaid struct {

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标识。取值范围：
	// <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	// <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>
	//
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type DescribeJavaMemShellPluginListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：Keywords: ip或者主机名模糊查询, JavaShellStatus，Exception精确匹配

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeJavaMemShellPluginListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellPluginListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulHostTopRequest struct {
	*tchttp.BaseRequest

	// 是否仅统计重点关注漏洞 1=仅统计重点关注漏洞, 0=统计全部漏洞

	IsFollowVul *uint64 `json:"IsFollowVul,omitempty" name:"IsFollowVul"`
	// 获取top值，1-100

	Top *uint64 `json:"Top,omitempty" name:"Top"`
	// 1:web-cms 漏洞，2.应用漏洞   4: Linux软件漏洞 5: windows系统漏洞 6:应急漏洞，不填或者填0时返回 1，2，4，5 的总统计数据

	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`
}

func (r *DescribeVulHostTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulHostTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRiskDnsPolicyListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [PolicyName|HostType]

	By *string `json:"By,omitempty" name:"By"`
	// <li>PolicyType - int - 是否必填：否 - 策略类型</li>
	// <li>PolicyName - string - 是否必填：否 - 策略名称</li>
	// <li>Domain - string - 是否必填：否 - 域名(先对域名做urlencode,再base64)</li>
	// <li>PolicyAction- int - 是否必填：否 - 策略动作</li>
	// <li>IsEnabled - int - 是否必填：否 - 是否生效</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportRiskDnsPolicyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRiskDnsPolicyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CanFixVulInfo struct {

	// 修复提示tag

	FixTag []*string `json:"FixTag,omitempty" name:"FixTag"`
	// 该漏洞可修复的主机信息

	HostList []*VulInfoHostInfo `json:"HostList,omitempty" name:"HostList"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type CheckFirstScanBaselineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否是第一次检测 0：不是；1：是

		FirstScan *uint64 `json:"FirstScan,omitempty" name:"FirstScan"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckFirstScanBaselineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFirstScanBaselineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAESKeyRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAESKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAESKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClientExceptionRequest struct {
	*tchttp.BaseRequest

	// 结束时间 `2006-01-02 15:04:05` 格式

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 客户端异常类型 1:客户端离线，2:客户端卸载

	ExceptionType *int64 `json:"ExceptionType,omitempty" name:"ExceptionType"`
	// 分页单页限制数目，不为0，最大值100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页的偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 起始时间 `2006-01-02 15:04:05` 格式

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeClientExceptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClientExceptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportFileTamperEventsRequest struct {
	*tchttp.BaseRequest

	// 需要导出的字段

	Fileds []*string `json:"Fileds,omitempty" name:"Fileds"`
	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略</li>
	// <li>ModifyTime - String - 是否必填：否 - 最近发生时间</li>
	// <li>Uuid- String - 是否必填：否 - 主机uuid查询</li>
	// <li>RuleCategory- string - 是否必填：否 - 规则类别 0 系统规则 1 自定义规则</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportFileTamperEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportFileTamperEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckBashPolicyParamsRequest struct {
	*tchttp.BaseRequest

	// 校验内容 Name或Rule ，两个都要校验时逗号分割

	CheckField *string `json:"CheckField,omitempty" name:"CheckField"`
	// 在事件列表中新增白名时需要提交事件ID

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 编辑时传的规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 填入的规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户填入的正则表达式："正则表达式" 需与 "提交EventId对应的命令内容" 相匹配

	Rule *string `json:"Rule,omitempty" name:"Rule"`
}

func (r *CheckBashPolicyParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBashPolicyParamsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOverviewStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineScanScheduleRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeBaselineScanScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineScanScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulDefenceEventStatusRequest struct {
	*tchttp.BaseRequest

	// 事件id

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 0: 待处理  2:已处理 3: 已忽略 4: 已删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyVulDefenceEventStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceEventStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Eventlog struct {

	// 详情原文

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 事件操作状态

	EventOperateStatus *string `json:"EventOperateStatus,omitempty" name:"EventOperateStatus"`
	// 事件状态

	EventStatus *string `json:"EventStatus,omitempty" name:"EventStatus"`
	// 发生时间

	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`
	// 事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 主机IP

	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
	// 事件ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 关键信息

	Memo *string `json:"Memo,omitempty" name:"Memo"`
	// 主机ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 威胁等级

	SeverityClass *string `json:"SeverityClass,omitempty" name:"SeverityClass"`
	// 云镜ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeRiskDnsPolicyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*RiskDnsPolicy `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsPolicyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字(机器名称/机器IP </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装 | SHUTDOWN 已关机）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// <li>Risk - String 是否必填: 否 - 风险主机( yes ) </li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 机器所属地域。如：ap-guangzhou

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 云主机类型。
	// <li>CVM：表示云服务器</li>
	// <li>BM:  表示黑石物理机</li>
	// <li>ECM:  表示边缘计算服务器</li>
	// <li>LH:  表示轻量应用服务器</li>
	// <li>Other:  表示混合云服务器</li>

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaliciousRequestWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 备注

	Mark *string `json:"Mark,omitempty" name:"Mark"`
}

func (r *CreateMaliciousRequestWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaliciousRequestWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineAnalysisDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否是第一次检测  1是 0不是

		IfFirstScan *uint64 `json:"IfFirstScan,omitempty" name:"IfFirstScan"`
		// 是否全部服务器：1-是 0-否

		IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
		// 最后检测时间

		LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
		// 服务器总数

		ScanHostCount *uint64 `json:"ScanHostCount,omitempty" name:"ScanHostCount"`
		// 检测项总数

		ScanRuleCount *uint64 `json:"ScanRuleCount,omitempty" name:"ScanRuleCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineAnalysisDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineAnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefenceEventRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：Keywords: ip或者主机名, VulKeywords漏洞名或者CveId模糊查询; Quuid，VulId，EventType，Status精确匹配，CreateBeginTime，CreateEndTime时间段查询

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 导出字段，默认全导出

	Where []*string `json:"Where,omitempty" name:"Where"`
}

func (r *ExportVulDefenceEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefenceEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenHostInvasionRequest struct {
	*tchttp.BaseRequest

	// 主机的网络攻击数

	AttackLogCnt *uint64 `json:"AttackLogCnt,omitempty" name:"AttackLogCnt"`
	// 主机的基线数

	BaselineCnt *uint64 `json:"BaselineCnt,omitempty" name:"BaselineCnt"`
	// 主机的入侵事件数

	InvasionEventsCnt *uint64 `json:"InvasionEventsCnt,omitempty" name:"InvasionEventsCnt"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机的漏洞数

	VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`
}

func (r *DescribeScreenHostInvasionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenHostInvasionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetUserListRequest struct {
	*tchttp.BaseRequest

	// 可选排序：[FirstTime|LoginTime|ChangePasswordTime|PasswordDuaTime]
	// PasswordLockDays

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name - String - 是否必填：否 - 账户名</li>
	// <li>Uid - uint64 - 是否必填：否 - Uid</li>
	// <li>Guid - uint64 - 是否必填：否 - Guid</li>
	// <li>LoginTimeStart - String - 是否必填：否 - 开始时间，如：2021-01-11</li>
	// <li>LoginTimeEnd - String - 是否必填：否 - 结束时间，如：2021-01-11</li>
	// <li>LoginType - uint64 - 是否必填：否 - 0-不可登录；1-只允许key登录；2只允许密码登录；3-允许key和密码 仅linux</li>
	// <li>OsType - String - 是否必填：否 - windows或linux</li>
	// <li>Status - uint64 - 是否必填：否 - 账号状态：0-禁用；1-启用</li>
	// <li>Type - uint64 - 是否必填：否 - 账号类型：0访客用户，1标准用户，2管理员用户 仅windows</li>
	// <li>IsDomain - uint64 - 是否必填：否 - 是否域账号：0 不是，1是 仅windows
	// <li>IsRoot - uint64 - 是否必填：否 - 是否Root权限：0 不是，1是 仅linux
	// <li>IsSudo - uint64 - 是否必填：否 - 是否Sudo权限：0 不是，1是 仅linux</li>
	// <li>IsSshLogin - uint64 - 是否必填：否 - 是否ssh登录：0 不是，1是 仅linux</li>
	// <li>ShellLoginStatus - uint64 - 是否必填：否 - 是否shell登录性，0不是；1是 仅linux</li>
	// <li>PasswordStatus - uint64 - 是否必填：否 - 密码状态：1正常 2即将过期 3已过期 4已锁定 仅linux</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetUserListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetUserListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseTrendRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRansomDefenseTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebServiceBaseInfo struct {

	// 二进制路径

	BinPath *string `json:"BinPath,omitempty" name:"BinPath"`
	// 配置路径

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// Web服务ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 安装路径

	InstallPath *string `json:"InstallPath,omitempty" name:"InstallPath"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	//  附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 数据库名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 关联进程数

	ProcessCount *uint64 `json:"ProcessCount,omitempty" name:"ProcessCount"`
	// 主机业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 启动用户

	User *string `json:"User,omitempty" name:"User"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type BaselineDetectParam struct {

	// 检测的主机ID集合

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 检测项集合

	ItemIds []*int64 `json:"ItemIds,omitempty" name:"ItemIds"`
	// 检测的策略集合

	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`
	// 检测的规则集合

	RuleIds []*int64 `json:"RuleIds,omitempty" name:"RuleIds"`
}

type PortInfo struct {

	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 容器内PID

	ContainerPID *uint64 `json:"ContainerPID,omitempty" name:"ContainerPID"`
	// 容器端口

	ContainerPort *uint64 `json:"ContainerPort,omitempty" name:"ContainerPort"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 容器内监听地址

	ListenContainer *string `json:"ListenContainer,omitempty" name:"ListenContainer"`
	// 容器外监听地址

	ListenHost *string `json:"ListenHost,omitempty" name:"ListenHost"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 对外ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 主机端口

	PublicPort *uint64 `json:"PublicPort,omitempty" name:"PublicPort"`
	// 运行账号

	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DeleteMalwaresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMalwaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseBindScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定任务详情

		List []*LicenseBindTaskDetail `json:"List,omitempty" name:"List"`
		// 进度

		Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseBindScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseBindScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMachineRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的是否为统计分析（即SQL）结果

		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`
		// 透传本次接口返回的Context值，可获取后续更多日志，过期时间1小时

		Context *string `json:"Context,omitempty" name:"Context"`
		// 匹配检索条件的原始日志的数量

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 匹配检索条件的原始日志

		Data []*LogInfo `json:"Data,omitempty" name:"Data"`
		// 符合检索条件的日志是否已全部返回，如未全部返回可使用Context参数获取后续更多日志

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tags struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type BatchAddMachineTagRequest struct {
	*tchttp.BaseRequest

	// 云服务器类型(CVM|BM)

	MachineArea *string `json:"MachineArea,omitempty" name:"MachineArea"`
	// 云服务器地区

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 云服务器ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 标签ID列表

	TagIds []*uint64 `json:"TagIds,omitempty" name:"TagIds"`
}

func (r *BatchAddMachineTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchAddMachineTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表内容

		List []*ReverseShell `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
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

type DescribeVdbAndPocInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVdbAndPocInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVdbAndPocInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBaselineRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBaselineRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBaselineRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetMachineListRequest struct {
	*tchttp.BaseRequest

	// 可选排序[FirstTime|PartitionCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>OsType - String - 是否必填：否 - windows或linux</li>
	// <li>CpuLoad - Int - 是否必填：否 -
	// 0: 未知  1: 低负载
	// 2: 中负载  3: 高负载</li>
	// <li>DiskLoad - Int - 是否必填：否 -
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>
	// <li>MemLoad - Int - 是否必填：否 -
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>
	// <li>Quuid：主机Quuid</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetMachineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetAttackSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetAttackSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetAttackSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewProVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewProVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogExportRequest struct {
	*tchttp.BaseRequest

	// 日志导出ID

	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
}

func (r *DeleteLogExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProtectDirRequest struct {
	*tchttp.BaseRequest

	// 删除的目录ID 最大100条

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteProtectDirRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProtectDirRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefenceListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：Keywords: 漏洞名称或CVE编号模糊匹配, Level精确匹配

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 导出字段，默认全导出

	Where []*string `json:"Where,omitempty" name:"Where"`
}

func (r *ExportVulDefenceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefenceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineItem `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineItemListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteHostListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件 <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 白名单ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 需要返回的数量，最大值为1000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLoginWhiteHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginWhiteHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesSimpleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表

		Machines []*MachineSimple `json:"Machines,omitempty" name:"Machines"`
		// 主机数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachinesSimpleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesSimpleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUndoVulCountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 普通版主机数

		NotProfessionCount *uint64 `json:"NotProfessionCount,omitempty" name:"NotProfessionCount"`
		// 未处理的主机数

		UndoHostCount *int64 `json:"UndoHostCount,omitempty" name:"UndoHostCount"`
		// 未处理的漏洞数

		UndoVulCount *uint64 `json:"UndoVulCount,omitempty" name:"UndoVulCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUndoVulCountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUndoVulCountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonthInspectionReport struct {

	// 巡检报告更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 巡检报告名称

	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
	// 巡检报告下载地址

	ReportPath *string `json:"ReportPath,omitempty" name:"ReportPath"`
}

type DeleteBanWhiteListRequest struct {
	*tchttp.BaseRequest

	// 要删除的白名单ID列表 (最大100条)

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBanWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBanWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebLocationCountRequest struct {
	*tchttp.BaseRequest

	// 搜索条件：返回名称包含Name的所有Web站点列表

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetWebLocationCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityTrendsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基线统计数据数组。

		BaseLines []*SecurityTrend `json:"BaseLines,omitempty" name:"BaseLines"`
		// 密码破解事件统计数据数组。

		BruteAttacks []*SecurityTrend `json:"BruteAttacks,omitempty" name:"BruteAttacks"`
		// 网络攻击统计数据数组。

		CyberAttacks []*SecurityTrend `json:"CyberAttacks,omitempty" name:"CyberAttacks"`
		// 高危命令统计数据数组。

		HighRiskBashs []*SecurityTrend `json:"HighRiskBashs,omitempty" name:"HighRiskBashs"`
		// 恶意请求统计数据数组。

		MaliciousRequests []*SecurityTrend `json:"MaliciousRequests,omitempty" name:"MaliciousRequests"`
		// 木马事件统计数据数组。

		Malwares []*SecurityTrend `json:"Malwares,omitempty" name:"Malwares"`
		// 异地登录事件统计数据数组。

		NonLocalLoginPlaces []*SecurityTrend `json:"NonLocalLoginPlaces,omitempty" name:"NonLocalLoginPlaces"`
		// 本地提权统计数据数组。

		PrivilegeEscalations []*SecurityTrend `json:"PrivilegeEscalations,omitempty" name:"PrivilegeEscalations"`
		// 反弹shell统计数据数组。

		ReverseShells []*SecurityTrend `json:"ReverseShells,omitempty" name:"ReverseShells"`
		// 漏洞统计数据数组。

		Vuls []*SecurityTrend `json:"Vuls,omitempty" name:"Vuls"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityTrendsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityTrendsResponse) FromJsonString(s string) error {
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

type DescribeLogExportsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志导出列表

		Exports []*ExportInfo `json:"Exports,omitempty" name:"Exports"`
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogExportsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBashPolicyRequest struct {
	*tchttp.BaseRequest

	// 具体的策略配置

	Policy *BashPolicy `json:"Policy,omitempty" name:"Policy"`
}

func (r *ModifyBashPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBashPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopAssetScanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopAssetScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopAssetScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgnoreBaselineRuleRequest struct {
	*tchttp.BaseRequest

	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 检测项名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *DescribeIgnoreBaselineRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgnoreBaselineRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePriceRequest struct {
	*tchttp.BaseRequest

	// 询价参数

	ResInfo []*PriceDetail `json:"ResInfo,omitempty" name:"ResInfo"`
}

func (r *DescribePriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenEmergentMsgRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScreenEmergentMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenEmergentMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigUUIDInfo struct {

	// 文件防护ID

	ConfigId *uint64 `json:"ConfigId,omitempty" name:"ConfigId"`
	// 服务器ID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

type ReverseShellRule struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 目标IP

	DestIp *string `json:"DestIp,omitempty" name:"DestIp"`
	// 目标端口

	DestPort *string `json:"DestPort,omitempty" name:"DestPort"`
	// 主机IP

	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 状态 (0: 有效 1: 无效)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 客户端ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetAppProcessListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// App名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetAppProcessListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppProcessListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackEventsRequest struct {
	*tchttp.BaseRequest

	// 排序

	By *string `json:"By,omitempty" name:"By"`
	//  过滤条件。
	// <li>Type - String 攻击状态 0: 尝试攻击 1: 攻击成功 - 是否必填: 否</li>
	// <li>Status - String 事件处理状态 0：待处理 1：已处理 2： 已加白 3： 已忽略 4：已删除  - 是否必填: 否</li>
	// <li>SrcIP - String 来源IP - 是否必填: 否</li>
	// <li>Uuids - String 主机安全uuid - 是否必填: 否</li>
	// <li>Quuids - String cvm uuid - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>
	// <li>MachineName - String 主机名称 - 是否必填: 否</li>
	// <li>InstanceID - String 主机实例ID - 是否必填: 否</li>
	// <li>AttackTimeBegin - String 攻击开始时间 - 是否必填: 否</li>
	// <li>AttackTimeEnd - String 攻击结束时间 - 是否必填: 否</li>
	// <li>VulSupportDefense - String 漏洞是否支持防御 0不支持，1支持 - 是否必填: 否</li>
	//

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 ASC,DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAttackEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件详情

		BashEventsInfo *BashEventsInfo `json:"BashEventsInfo,omitempty" name:"BashEventsInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashEventsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashEventsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DeleteMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseRollBackTaskListRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime、ModifyTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Ips- string- 主机名称</li>
	// <li>Status - Uint64 - 0进行中，1成功，2失败</li>
	// <li>Names- String - 主机名称</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRansomDefenseRollBackTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseRollBackTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetPortInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetPortInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetPortInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetSystemPackageInfo struct {

	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 安装时间

	InstallTime *string `json:"InstallTime,omitempty" name:"InstallTime"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 主机IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 数据库名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeScreenProtectionCntRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScreenProtectionCntRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenProtectionCntRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 常用登录地数组

		UsualLoginPlaces []*UsualPlace `json:"UsualLoginPlaces,omitempty" name:"UsualLoginPlaces"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsualLoginPlacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJavaMemShellsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyJavaMemShellsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJavaMemShellsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLicenseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RescanImpactedHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RescanImpactedHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RescanImpactedHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// 专业周报开始时间。

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportBruteAttacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportBruteAttacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BroadcastInfo struct {

	// 富文本内容信息

	Content *string `json:"Content,omitempty" name:"Content"`
	// 发布时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 跳转位置：0=不跳转，1=文件查杀，2=漏洞扫描，3=安全基线

	GotoType *uint64 `json:"GotoType,omitempty" name:"GotoType"`
	// 文章唯一Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 副标题

	Subtitle *string `json:"Subtitle,omitempty" name:"Subtitle"`
	// 文章名字

	Title *string `json:"Title,omitempty" name:"Title"`
	// 类型：0=紧急通知，1=功能更新，2=行业荣誉，3=版本发布

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type EffectiveMachineInfo struct {

	// 云标签信息

	CloudTags []*Tags `json:"CloudTags,omitempty" name:"CloudTags"`
	// 机器instance ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 内核版本号

	KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`
	// 授权订单对象

	LicenseOrder *LicenseOrder `json:"LicenseOrder,omitempty" name:"LicenseOrder"`
	// 机器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 机器内网ip

	MachinePrivateIp *string `json:"MachinePrivateIp,omitempty" name:"MachinePrivateIp"`
	// 机器公网ip

	MachinePublicIp *string `json:"MachinePublicIp,omitempty" name:"MachinePublicIp"`
	// 在线状态 OFFLINE，ONLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 机器标签

	MachineTag []*MachineTag `json:"MachineTag,omitempty" name:"MachineTag"`
	// 机器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 云镜Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 漏洞数量

	VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`
}

type DescribeAssetScanHostDetailRequest struct {
	*tchttp.BaseRequest

	// 主机的Quuid

	Quuid *uint64 `json:"Quuid,omitempty" name:"Quuid"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeAssetScanHostDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetScanHostDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEmergencyVulListRequest struct {
	*tchttp.BaseRequest

	// 排序字段 PublishDate  LastScanTime HostCount

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Status - String - 是否必填：是 - 漏洞状态筛选，0//未检测 1有风险 ，2无风险 ，3 检查中展示progress</li>
	// <li>Level - String - 是否必填：否 - 漏洞等级筛选 1:低 2:中 3:高 4:提示</li>
	// <li>VulName- String - 是否必填：否 - 漏洞名称搜索</li>
	// <li>Uuids- String - 是否必填：否 - 主机uuid</li>
	// <li>IsSupportDefense - int- 是否必填：否 - 是否支持防御 0:不支持 1:支持</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 desc , asc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeEmergencyVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEmergencyVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseWhiteConfigRequest struct {
	*tchttp.BaseRequest

	// 规则名称,例如: cwp

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *DescribeLicenseWhiteConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseWhiteConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBashPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBashPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBashPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBruteAttackRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBruteAttackRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBruteAttackRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络攻击

		AttackLog []*SecurityTrend `json:"AttackLog,omitempty" name:"AttackLog"`
		// 基线漏洞

		BaseLine []*SecurityTrend `json:"BaseLine,omitempty" name:"BaseLine"`
		// 暴力破解

		BruteAttack []*SecurityTrend `json:"BruteAttack,omitempty" name:"BruteAttack"`
		// 应急漏洞

		EmergencyVul []*SecurityTrend `json:"EmergencyVul,omitempty" name:"EmergencyVul"`
		// 高危命令

		HighRiskBash []*SecurityTrend `json:"HighRiskBash,omitempty" name:"HighRiskBash"`
		// 登录审计

		HostLogin []*SecurityTrend `json:"HostLogin,omitempty" name:"HostLogin"`
		// 恶意请求

		MaliciousRequest []*SecurityTrend `json:"MaliciousRequest,omitempty" name:"MaliciousRequest"`
		// 木马文件

		Malware []*SecurityTrend `json:"Malware,omitempty" name:"Malware"`
		// 本地提权

		PrivilegeEscalation []*SecurityTrend `json:"PrivilegeEscalation,omitempty" name:"PrivilegeEscalation"`
		// 反弹shell

		ReverseShell []*SecurityTrend `json:"ReverseShell,omitempty" name:"ReverseShell"`
		// 系统组件漏洞

		SystemComponentVul []*SecurityTrend `json:"SystemComponentVul,omitempty" name:"SystemComponentVul"`
		// WEB漏洞

		WebVul []*SecurityTrend `json:"WebVul,omitempty" name:"WebVul"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启木马自动隔离(0:关闭 1:开启)

		MalwareAutoIsolate *int64 `json:"MalwareAutoIsolate,omitempty" name:"MalwareAutoIsolate"`
		// 设置木马定时扫描间隔(天数)

		MalwareDiskScanInterDay *int64 `json:"MalwareDiskScanInterDay,omitempty" name:"MalwareDiskScanInterDay"`
		// 设置木马定时扫描时间(0-23整点时间)

		MalwareDiskScanTime *int64 `json:"MalwareDiskScanTime,omitempty" name:"MalwareDiskScanTime"`
		// 是否开启木马定时扫描(0:关闭 1:开启)

		MalwareEnableDiskScan *int64 `json:"MalwareEnableDiskScan,omitempty" name:"MalwareEnableDiskScan"`
		// 设置木马扫描全盘扫描忽略路径(linux), 换行符分割多个

		MalwareIgnoreWhiteListLinux *string `json:"MalwareIgnoreWhiteListLinux,omitempty" name:"MalwareIgnoreWhiteListLinux"`
		// 设置木马扫描全盘扫描忽略路径(windows), 换行符分割多个

		MalwareIgnoreWhiteListWindows *string `json:"MalwareIgnoreWhiteListWindows,omitempty" name:"MalwareIgnoreWhiteListWindows"`
		// 设置木马扫描实时监控忽略路径(linux), 换行符分割多个

		MalwareMonitorIgnoreWhiteListLinux *string `json:"MalwareMonitorIgnoreWhiteListLinux,omitempty" name:"MalwareMonitorIgnoreWhiteListLinux"`
		// 设置木马扫描实时监控忽略路径(windows), 换行符分割多个

		MalwareMonitorIgnoreWhiteListWindows *string `json:"MalwareMonitorIgnoreWhiteListWindows,omitempty" name:"MalwareMonitorIgnoreWhiteListWindows"`
		// 是否开启漏洞定时扫描(0:关闭 1:开启)

		VulEnableVulnerDetect *int64 `json:"VulEnableVulnerDetect,omitempty" name:"VulEnableVulnerDetect"`
		// 设置漏洞定时扫描间隔(天数)

		VulVulnerDetectDay *int64 `json:"VulVulnerDetectDay,omitempty" name:"VulVulnerDetectDay"`
		// 设置漏洞定时扫描时间(0-23整点时间)

		VulVulnerDetectTime *int64 `json:"VulVulnerDetectTime,omitempty" name:"VulVulnerDetectTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFastAnalysisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分析统计信息

		FieldValueRatioInfos []*FieldValueRatioInfo `json:"FieldValueRatioInfos,omitempty" name:"FieldValueRatioInfos"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFastAnalysisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFastAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 扫描漏洞数

		RiskEventCount *uint64 `json:"RiskEventCount,omitempty" name:"RiskEventCount"`
		// 开始扫描时间

		ScanBeginTime *string `json:"ScanBeginTime,omitempty" name:"ScanBeginTime"`
		// 扫描结束时间

		ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`
		// 0 从未扫描过、 1 扫描中、 2扫描完成、 3停止中、 4停止完成

		ScanState *uint64 `json:"ScanState,omitempty" name:"ScanState"`
		// 扫描进度

		Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`
		// 任务Id

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 0一键检测 1定时检测

		Type *uint64 `json:"Type,omitempty" name:"Type"`
		// 任务扫描的漏洞id

		VulId []*uint64 `json:"VulId,omitempty" name:"VulId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProVersionMachine struct {

	// 主机所在地域。
	// 如：ap-guangzhou、ap-beijing

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 主机类型。
	// <li>CVM: 云服务器</li>
	// <li>BM: 黑石物理机</li>
	// <li>ECM: 边缘计算服务器</li>
	// <li>LH: 轻量应用服务器</li>
	// <li>Other: 混合云机器</li>

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 主机唯一标识Uuid数组。
	// 黑石的InstanceId，CVM的Uuid ,边缘计算的Uuid , 轻量应用服务器的Uuid ,混合云机器的Quuid 。 当前参数最大长度限制20

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type VulLevelInfo struct {

	// 数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// // 危害等级：1-低危；2-中危；3-高危；4-严重

	VulLevel *uint64 `json:"VulLevel,omitempty" name:"VulLevel"`
}

type DeleteUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsualLoginPlacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportWebPageEventListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：CreateTime 或 RestoreTime，默认为CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>EventType - String - 是否必填：否 - 事件类型</li>
	// <li>EventStatus - String - 是否必填：否 - 事件状态</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，0降序，1升序，默认为0

	Order *uint64 `json:"Order,omitempty" name:"Order"`
}

func (r *ExportWebPageEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportWebPageEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScreenBroadcasts struct {

	// 文章ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 播报文章危险程度  0：无， 1：严重， 2: 高危， 3:中危， 4: 低危

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 发布时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 播报文章标题

	Title *string `json:"Title,omitempty" name:"Title"`
}

type CreateAssetScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebPageServiceInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有授权机器信息

		AllAuthorizedMachines []*ProtectMachineInfo `json:"AllAuthorizedMachines,omitempty" name:"AllAuthorizedMachines"`
		// 已购授权数

		BuyNum *uint64 `json:"BuyNum,omitempty" name:"BuyNum"`
		// 临近到期授权机器信息

		ExpireAuthorizedMachines []*ProtectMachine `json:"ExpireAuthorizedMachines,omitempty" name:"ExpireAuthorizedMachines"`
		// 临近到期数量

		ExpireNum *uint64 `json:"ExpireNum,omitempty" name:"ExpireNum"`
		// 已过期授权数

		ExpiredNum *uint64 `json:"ExpiredNum,omitempty" name:"ExpiredNum"`
		// 防护目录数

		ProtectDirNum *uint64 `json:"ProtectDirNum,omitempty" name:"ProtectDirNum"`
		// 剩余授权数

		ResidueNum *uint64 `json:"ResidueNum,omitempty" name:"ResidueNum"`
		// 是否已购服务：true-是，false-否

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 已使用授权数

		UsedNum *uint64 `json:"UsedNum,omitempty" name:"UsedNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebPageServiceInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebPageServiceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineHostRiskTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险主机top5

		HostRiskTop5 []*HostRiskLevelCount `json:"HostRiskTop5,omitempty" name:"HostRiskTop5"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineHostRiskTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineHostRiskTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebAppListRequest struct {
	*tchttp.BaseRequest

	// 可选排序：[FirstTime|PluginCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name - String - 是否必填：否 - 应用名</li>
	// <li>Domain - String - 是否必填：否 - 站点域名</li>
	// <li>Type - int - 是否必填：否 - 服务类型：
	// 0：全部
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:Jetty
	// 8:IHS
	// 9:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetWebAppListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebAppListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineEventLevelInfo struct {

	// 漏洞数量

	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
	// 危害等级：1-低危；2-中危；3-高危；4-严重

	EventLevel *uint64 `json:"EventLevel,omitempty" name:"EventLevel"`
}

type DeleteBaselineStrategyRequest struct {
	*tchttp.BaseRequest

	// 基线策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DeleteBaselineStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBaselineStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDatabaseInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据库详情

		Database *AssetDatabaseDetail `json:"Database,omitempty" name:"Database"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDatabaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExpertServiceTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyExpertServiceTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExpertServiceTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertProtectPathInfo struct {

	// 失败的防护目录ID和UUID

	Fail []*ConfigUUIDInfo `json:"Fail,omitempty" name:"Fail"`
	// 离线的防护目录ID和UUID

	Offline []*ConfigUUIDInfo `json:"Offline,omitempty" name:"Offline"`
	// 响应成功的防护目录ID和UUID

	Success []*ConfigUUIDInfo `json:"Success,omitempty" name:"Success"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 帐号数据列表。

		Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`
		// 帐号列表记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportMalwaresRequest struct {
	*tchttp.BaseRequest

	// 专业周报开始时间。

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportMalwaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportMalwaresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最新同步结束时间

		LatestEndTime *string `json:"LatestEndTime,omitempty" name:"LatestEndTime"`
		// 最新开始同步时间

		LatestStartTime *string `json:"LatestStartTime,omitempty" name:"LatestStartTime"`
		// 枚举值有(大写)：NOTASK（没有同步任务），SYNCING（同步中），FINISHED（同步完成）

		State *string `json:"State,omitempty" name:"State"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单规则

	Rules *LoginWhiteListsRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *AddLoginWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLoginWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineSnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机快照列表

		List []*MachineSnapshotInfo `json:"List,omitempty" name:"List"`
		// 是否都有24小时内创建的快照

		SnapshotCheck *bool `json:"SnapshotCheck,omitempty" name:"SnapshotCheck"`
		// 总个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetCoreModuleListRequest struct {
	*tchttp.BaseRequest

	// 排序依据[FirstTime|Size|ProcessCount|ModuleCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetCoreModuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetCoreModuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetScanStatusRequest struct {
	*tchttp.BaseRequest

	// 主机的Quuid, 如果未设置，则获取的是整体的扫描状态

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetScanStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetScanStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanResultRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulScanResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileTamperRuleCount struct {

	// 关联规则的数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 关联规则的名称（仅展示其中一条）

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAttackSourceEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击溯源事件列表

		List []*AttackSourceEvent `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackSourceEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackSourceEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineLicenseDetailRequest struct {
	*tchttp.BaseRequest

	// 主机quuid

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *DescribeMachineLicenseDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineLicenseDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseMachineStrategyInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略ID列表,0表示未绑定任何策略

		StrategyIds []*uint64 `json:"StrategyIds,omitempty" name:"StrategyIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRansomDefenseMachineStrategyInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseMachineStrategyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineItem struct {

	// 是否可以修复

	CanBeFixed *int64 `json:"CanBeFixed,omitempty" name:"CanBeFixed"`
	// 检测项分类

	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 检测结果描述

	DetectResultDesc *string `json:"DetectResultDesc,omitempty" name:"DetectResultDesc"`
	// 检测状态：0 未通过，1：忽略，3：通过，5：检测中

	DetectStatus *int64 `json:"DetectStatus,omitempty" name:"DetectStatus"`
	// 第一次出现时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 修复方法

	FixMethod *string `json:"FixMethod,omitempty" name:"FixMethod"`
	// 主机ID

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 项描述

	ItemDesc *string `json:"ItemDesc,omitempty" name:"ItemDesc"`
	// 项Id

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
	// 项名称

	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`
	// 最近出现时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 危险等级

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 所属规则

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 主机安全uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 外网IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Ips - String - 是否必填：否 - 通过ip查询 </li>
	// <li>Names - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>InstanceIds - String - 是否必填：否 - 通过实例id查询 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线/关机 | ONLINE: 在线 | UNINSTALLED：未安装 | AGENT_OFFLINE 离线| AGENT_SHUTDOWN 已关机）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版 | Flagship : 旗舰版 | ProtectedMachines: 专业版+旗舰版）</li>
	// <li>Risk - String 是否必填: 否 - 风险主机( yes ) </li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	// <li>Quuid - String - 是否必填: 否 - 云服务器uuid  最大100条.</li>
	// <li>AddedOnTheFifteen- String 是否必填: 否 - 是否只查询15天内新增的主机( 1：是) </li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 机器所属地域。如：ap-guangzhou，ap-shanghai

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 机器所属专区类型
	// CVM 云服务器
	// BM 黑石
	// ECM 边缘计算
	// LH 轻量应用服务器
	// Other 混合云专区

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 机器所属业务ID列表

	ProjectIds []*uint64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseEventsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表

		List []*RansomDefenseEvent `json:"List,omitempty" name:"List"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRansomDefenseEventsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseEventsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselinePolicyDetect struct {

	// 失败主技数

	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`
	// 结束时间

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
	// 关联主机数

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 1:检测中 2:检测完成

	PolicyDetectStatus *int64 `json:"PolicyDetectStatus,omitempty" name:"PolicyDetectStatus"`
	// 策略Id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 成功主机数

	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
	// 检测任务Id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 失败主机数

	TimeoutCount *int64 `json:"TimeoutCount,omitempty" name:"TimeoutCount"`
}

type DescribeRansomDefenseTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开通勒索的公司数量

		CompanyCount *uint64 `json:"CompanyCount,omitempty" name:"CompanyCount"`
		// 勒索事件数量

		EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
		// 影响行业数量

		IndustryCount *uint64 `json:"IndustryCount,omitempty" name:"IndustryCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRansomDefenseTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportIgnoreRuleEffectHostListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>AliasName- String- 主机别名</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 检测项id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ExportIgnoreRuleEffectHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportIgnoreRuleEffectHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebPageProtectDirResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWebPageProtectDirResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebPageProtectDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RescanImpactedHostRequest struct {
	*tchttp.BaseRequest

	// 漏洞ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *RescanImpactedHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RescanImpactedHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetScanTaskHostDetail struct {

	// 异常信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	// 主机IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 任务状态：0进行中，1成功，2失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type WeeklyReportVul struct {

	// 漏洞描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 最后扫描时间。

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 主机内网IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 漏洞名称。

	VulName *string `json:"VulName,omitempty" name:"VulName"`
	// 漏洞状态。
	// <li> UN_OPERATED : 待处理</li>
	// <li> SCANING : 扫描中</li>
	// <li> FIXED : 已修复</li>

	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`
	// 漏洞类型。
	// <li> WEB : Web漏洞</li>
	// <li> SYSTEM :系统组件漏洞</li>
	// <li> BASELINE : 安全基线</li>

	VulType *string `json:"VulType,omitempty" name:"VulType"`
}

type DescribeScanScheduleRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeScanScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportScanTaskDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出本次检测Excel的任务Id（不同于入参的本次检测任务id）

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportScanTaskDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportScanTaskDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineAutoClearConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMachineAutoClearConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMachineAutoClearConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverMalwaresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恢复失败id数组，若无则返回空数组

		FailedIds []*uint64 `json:"FailedIds,omitempty" name:"FailedIds"`
		// 恢复成功id数组，若无则返回空数组

		SuccessIds []*uint64 `json:"SuccessIds,omitempty" name:"SuccessIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecoverMalwaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecoverMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicProxyInstallCommandResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Keepalived安装命令

		KeepAliveCommand *string `json:"KeepAliveCommand,omitempty" name:"KeepAliveCommand"`
		// Nginx安装命令

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

type DescribeRansomDefenseMachineStrategyInfoRequest struct {
	*tchttp.BaseRequest

	// 主机Quuid列表

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *DescribeRansomDefenseMachineStrategyInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseMachineStrategyInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineItemDetect struct {

	// 检测结果,Json字符串

	DetectResult *string `json:"DetectResult,omitempty" name:"DetectResult"`
	// 0:未通过 1:忽略 3:通过 5:检测中

	DetectStatus *int64 `json:"DetectStatus,omitempty" name:"DetectStatus"`
	// 首次检测时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 修复方法

	FixMethod *string `json:"FixMethod,omitempty" name:"FixMethod"`
	// 影响服务器数

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 项描述

	ItemDesc *string `json:"ItemDesc,omitempty" name:"ItemDesc"`
	// 项Id

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
	// 项名称

	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`
	// 最后检测时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 风险等级

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 未通过的服务器数

	NotPassedHostCount *int64 `json:"NotPassedHostCount,omitempty" name:"NotPassedHostCount"`
	// 通过的服务器数

	PassedHostCount *int64 `json:"PassedHostCount,omitempty" name:"PassedHostCount"`
	// 所属规则ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 所属规则

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type DescribeAttackTrendsRequest struct {
	*tchttp.BaseRequest

	//  过滤条件。
	// <li>BeginTime - String 起始时间,默认近7天- 是否必填: 否</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAttackTrendsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackTrendsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncMachinesRequest struct {
	*tchttp.BaseRequest

	// 是否同步

	Sync *bool `json:"Sync,omitempty" name:"Sync"`
}

func (r *SyncMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRansomDefenseStrategyMachinesRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime、MachineCount

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Ips - String - 是否必填：否 - 通过ip查询 </li>
	// <li>MachineNames - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>Names - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>Status - String - 是否必填：否 - 策略状态：0备份中，1备份成功，2备份失败 </li>
	// <li>LastBackupTimeBegin - String - 是否必填：否 - 最近一次备份时间开始</li>
	// <li>LastBackupTimeEnd - String - 是否必填：否 - 最近一次备份时间结束</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 策略id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportRansomDefenseStrategyMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRansomDefenseStrategyMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSearchLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：成功，非0：失败

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSearchLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEmergencyVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否存在风险

		ExistsRisk *bool `json:"ExistsRisk,omitempty" name:"ExistsRisk"`
		// 漏洞列表

		List []*EmergencyVul `json:"List,omitempty" name:"List"`
		// 漏洞总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
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

type HostTagInfo struct {

	// 主机名

	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`
	// 云标签信息

	CloudTags []*Tags `json:"CloudTags,omitempty" name:"CloudTags"`
	// 主机内网Ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机instance ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 内核版本号

	KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`
	// 主机在线状态 ONLINE，OFFLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 主机公网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 防护版本 BASIC_VERSION 基础版, PRO_VERSION 专业版 Flagship 旗舰版

	ProtectType *string `json:"ProtectType,omitempty" name:"ProtectType"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机标签名数组

	TagList []*string `json:"TagList,omitempty" name:"TagList"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 漏洞数

	VulNum *int64 `json:"VulNum,omitempty" name:"VulNum"`
}

type OpenProVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID, 可查看授权绑定进度.

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineCustomRuleIdName struct {

	// 自定义规则ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 自定义规则名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type FileProtectAutoRecoverInfo struct {

	// 失败的防护目录ID

	Fail []*uint64 `json:"Fail,omitempty" name:"Fail"`
	// 离线的防护目录ID

	Offline []*uint64 `json:"Offline,omitempty" name:"Offline"`
	// 响应成功的防护目录ID

	Success []*uint64 `json:"Success,omitempty" name:"Success"`
}

type DeleteRiskDnsPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略ID

	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`
}

func (r *DeleteRiskDnsPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRiskDnsPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebLocationListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 站点列表

		Locations []*AssetWebLocationBaseInfo `json:"Locations,omitempty" name:"Locations"`
		// 记录总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebLocationListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 -  查询关键字</li>
	// <li>Status - String - 是否必填：否 -  查询状态（FAILED：破解失败 |SUCCESS：破解成功）</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeBruteAttacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBruteAttacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttackListRequest struct {
	*tchttp.BaseRequest

	// 排序字段：CreateTime-首次攻击时间

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Uuid - String - 是否必填：否 - 云镜唯一Uuid</li>
	// <li>Quuid - String - 是否必填：否 - 云服务器uuid</li>
	// <li>Status - String - 是否必填：否 - 状态筛选：失败：FAILED 成功：SUCCESS</li>
	// <li>UserName - String - 是否必填：否 - UserName筛选</li>
	// <li>SrcIp - String - 是否必填：否 - 来源ip筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 首次攻击时间筛选，开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 首次攻击时间筛选，结束时间</li>
	// <li>ModifyBeginTime - String - 是否必填：否 - 最近攻击时间筛选，开始时间</li>
	// <li>ModifyEndTime - String - 是否必填：否 - 最近攻击时间筛选，结束时间</li>
	// <li>Banned - String - 是否必填：否 - 阻断状态筛选，多个用","分割：0-未阻断（全局ZK开关关闭），82-未阻断(非专业版)，83-未阻断(已加白名单)，1-已阻断，2-未阻断-程序异常，3-未阻断-内网攻击暂不支持阻断，4-未阻断-安平暂不支持阻断</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：根据请求次数排序：asc-升序/desc-降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBruteAttackListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBruteAttackListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefenceEventDetailRequest struct {
	*tchttp.BaseRequest

	// 漏洞事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDefenceEventDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenceEventDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineHost struct {

	// 主机Id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 内网Ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机标签

	HostTag *string `json:"HostTag,omitempty" name:"HostTag"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 外网Ip

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}

type LoginWhiteListsNew struct {

	// 创建白名单时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 机器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否为全局规则,1： 全局有效 0: 仅针对单台主机

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 地域字符串

	Locale *string `json:"Locale,omitempty" name:"Locale"`
	// 机器名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 修改白名单时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 白名单地域

	Places []*Place `json:"Places,omitempty" name:"Places"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 白名单IP（多个IP逗号隔开）

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 白名单用户（多个用户逗号隔开）

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 云镜客户端ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type BaselinePassword struct {

	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
}

type DescribeAssetJarListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用列表

		Jars []*AssetJarBaseInfo `json:"Jars,omitempty" name:"Jars"`
		// 总数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetJarListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetJarListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportPrivilegeEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportPrivilegeEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportPrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebHookRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWebHookRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebHookRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JavaMemShellInfo struct {

	// 服务器名称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 首次发现时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// 服务器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 事件ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 服务器quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 最近检测时间

	RecentFoundTime *string `json:"RecentFoundTime,omitempty" name:"RecentFoundTime"`
	// 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略  4 - 已手动处理

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 内存马类型  0:Filter型 1:Listener型 2:Servlet型 3:Interceptors型 4:Agent型 5:其他

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 服务器uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExpertServiceOrderInfo struct {

	// 服务开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 服务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 服务数量

	InquireNum *uint64 `json:"InquireNum,omitempty" name:"InquireNum"`
	// 订单类型 1应急 2 旗舰重保 3 安全管家

	InquireType *uint64 `json:"InquireType,omitempty" name:"InquireType"`
	// 订单id

	OrderId *uint64 `json:"OrderId,omitempty" name:"OrderId"`
	// 服务时长几个月

	ServiceTime *uint64 `json:"ServiceTime,omitempty" name:"ServiceTime"`
	// 订单状态 0 未启动 1 服务中 2已过期 3完成，4退费销毁

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type CreateBaselineStrategyRequest struct {
	*tchttp.BaseRequest

	// 该策略下选择的基线id数组. 示例: [1,3,5,7]

	CategoryIds []*uint64 `json:"CategoryIds,omitempty" name:"CategoryIds"`
	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 云主机类型：
	// CVM：虚拟主机
	// BM：裸金属
	// ECM：边缘计算主机
	// LH：轻量应用服务器
	// Other：混合云机器

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 主机id数组. 示例: ["quuid1","quuid2"]

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 主机地域. 示例: "ap-guangzhou"

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 定期检测时间，该时间下发扫描. 示例:“22:00”, 表示在22:00下发检测

	ScanAt *string `json:"ScanAt,omitempty" name:"ScanAt"`
	// 检测周期, 表示每隔多少天进行检测.示例: 2, 表示每2天进行检测一次.

	ScanCycle *uint64 `json:"ScanCycle,omitempty" name:"ScanCycle"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
}

func (r *CreateBaselineStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBaselineStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineDefaultStrategyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 默认策略基础信息列表

		StrategyList []*DefaultStrategyInfo `json:"StrategyList,omitempty" name:"StrategyList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineDefaultStrategyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineDefaultStrategyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareRiskWarningRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMalwareRiskWarningRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareRiskWarningRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBashRulesRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBashRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBashRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetPortInfoListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime|StartTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Port - uint64 - 是否必填：否 - 端口</li>
	// <li>Ip - String - 是否必填：否 - 绑定IP</li>
	// <li>ProcessName - String - 是否必填：否 - 监听进程</li>
	// <li>Pid - uint64 - 是否必填：否 - PID</li>
	// <li>User - String - 是否必填：否 - 运行用户</li>
	// <li>Group - String - 是否必填：否 - 所属用户组</li>
	// <li>Ppid - uint64 - 是否必填：否 - PPID</li>
	// <li>Proto - string - 是否必填：否 - tcp/udp或“”(空字符串筛选未知状态)</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>RunTimeStart - String - 是否必填：否 - 运行开始时间</li>
	// <li>RunTimeEnd - String - 是否必填：否 - 运行结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetPortInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetPortInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRansomDefenseMachineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRansomDefenseMachineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRansomDefenseMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IgnoreImpactedHostsRequest struct {
	*tchttp.BaseRequest

	// 漏洞ID数组。

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *IgnoreImpactedHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IgnoreImpactedHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFileTamperRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 对应事件id

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 是否是系统规则 0=系统规则 1=用户自定义规则，系统规则Status 不支持删除

	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`
	// 0: 启用1: 关闭2：删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyFileTamperRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFileTamperRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogExportRequest struct {
	*tchttp.BaseRequest

	// 日志导出数量, 最大值5000万

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 日志导出结束时间，毫秒时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 日志导出数据格式。json，csv，默认为json

	Format *string `json:"Format,omitempty" name:"Format"`
	// 日志导出检索语句，不支持[SQL语句]

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 日志导出时间排序。desc，asc，默认为desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 日志导出起始时间，毫秒时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *CreateLogExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSearchLogRequest struct {
	*tchttp.BaseRequest

	// 搜索内容

	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
}

func (r *CreateSearchLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSearchLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineRuleRequest struct {
	*tchttp.BaseRequest

	// 基线id

	BaselineId *uint64 `json:"BaselineId,omitempty" name:"BaselineId"`
	// 危害等级

	Level []*uint64 `json:"Level,omitempty" name:"Level"`
	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeBaselineRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1 可购买 2 只能升降配 3 只能跳到续费管理页

		BuyStatus *uint64 `json:"BuyStatus,omitempty" name:"BuyStatus"`
		// 到期时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 用户已购容量 单位 G

		InquireNum *uint64 `json:"InquireNum,omitempty" name:"InquireNum"`
		// 是否自动续费,0 初始值, 1 开通 2 没开通

		IsAutoOpenRenew *uint64 `json:"IsAutoOpenRenew,omitempty" name:"IsAutoOpenRenew"`
		// 资源ID

		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 0 没开通 1 正常 2隔离 3销毁

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHistoryServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebFrameListRequest struct {
	*tchttp.BaseRequest

	// 可选排序：[FirstTime|JarCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 框架名</li>
	// <li>NameStrict - String - 是否必填：否 - 框架名（严格匹配）</li>
	// <li>Lang - String - 是否必填：否 - 框架语言:java/python</li>
	// <li>Type - String - 是否必填：否 - 服务类型：
	// 0：全部
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetWebFrameListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebFrameListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebPageEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防护事件列表信息

		List []*ProtectEventLists `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebPageEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebPageEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProtectServerRequest struct {
	*tchttp.BaseRequest

	// 防护目录地址

	ProtectDir *string `json:"ProtectDir,omitempty" name:"ProtectDir"`
	// 防护机器 信息

	ProtectHostConfig []*ProtectHostConfig `json:"ProtectHostConfig,omitempty" name:"ProtectHostConfig"`
}

func (r *CreateProtectServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProtectServerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVerWithQuuidsRequest struct {
	*tchttp.BaseRequest

	// 主机唯一标识Uuid数组。
	// 黑石的InstanceId，CVM/ECM的Uuid

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *OpenProVerWithQuuidsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenProVerWithQuuidsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScreenTrendsChart struct {

	// 统计分类类型 值：防御次数，攻击次数

	Type *string `json:"Type,omitempty" name:"Type"`
	// X轴 时间

	X *string `json:"X,omitempty" name:"X"`
	// Y轴 数值

	Y *uint64 `json:"Y,omitempty" name:"Y"`
}

type Vuls struct {

	// 漏洞ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 受影响机器数量

	ImpactHostCount *uint64 `json:"ImpactHostCount,omitempty" name:"ImpactHostCount"`
	// 最后扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 漏洞危害等级:
	// HIGH：高危
	// MIDDLE：中危
	// LOW：低危
	// NOTICE：提示

	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type DescribeBaselineDetectListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [HostCount|StartTime|StopTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>PolicyName - string - 是否必填：否 - 策略名称</li>
	// <li>PolicyDetectStatus - int - 是否必填：否 - 1:检测中 2:检测完成</li>
	// <li>FirstTime - string - 是否必填：否 - 开始时间</li>
	// <li>LastTime - string - 是否必填：否 - 结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineDetectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineDetectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsRequest struct {
	*tchttp.BaseRequest

	// 排序字段：CreateTime-发生时间。ModifyTime-处理时间

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机内网IP)</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：根据请求次数排序：asc-升序/desc-降序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBashEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeRulesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePrivilegeRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivilegeRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckLogKafkaConnectionStateRequest struct {
	*tchttp.BaseRequest

	// 接入地址，域名或ip，带端口

	AccessAddr *string `json:"AccessAddr,omitempty" name:"AccessAddr"`
	// 接入方式,1公网域名接入，2支撑环境接入

	AccessType *uint64 `json:"AccessType,omitempty" name:"AccessType"`
	// 0不用密码，1有密码

	HasPwd *uint64 `json:"HasPwd,omitempty" name:"HasPwd"`
	// kafka版本

	InsVersion *string `json:"InsVersion,omitempty" name:"InsVersion"`
	// kafka实例id

	KafkaId *string `json:"KafkaId,omitempty" name:"KafkaId"`
	// 密码，使用aes加密

	Pwd *string `json:"Pwd,omitempty" name:"Pwd"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
}

func (r *CheckLogKafkaConnectionStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLogKafkaConnectionStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMalwareWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMalwareWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMalwareWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryVulFixRequest struct {
	*tchttp.BaseRequest

	// 漏洞id

	FixId *uint64 `json:"FixId,omitempty" name:"FixId"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *RetryVulFixRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryVulFixRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetType struct {

	// 类型ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 类型名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type Item struct {

	// Id

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
	// 名称

	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`
}

type DescribeAssetJarListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>Type- uint - 是否必填：否 - 类型
	// 1: 应用程序
	// 2 : 系统类库
	// 3 : Web服务自带库
	// 4 : 其他依赖包</li>
	// <li>Status- string - 是否必填：否 - 是否可执行：0否，1是</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetJarListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetJarListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProtectNetInfo struct {

	// 护网完成时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 护网天数

	ProtectDays *uint64 `json:"ProtectDays,omitempty" name:"ProtectDays"`
	// 报告下载地址

	ReportPath *string `json:"ReportPath,omitempty" name:"ReportPath"`
	// 护网启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 护网状态 0未启动，1护网中，2已完成

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type CheckFirstScanBaselineRequest struct {
	*tchttp.BaseRequest
}

func (r *CheckFirstScanBaselineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFirstScanBaselineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrustMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马ID数组（单次不超过的最大条数：100）

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *TrustMalwaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TrustMalwaresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventInfoRequest struct {
	*tchttp.BaseRequest

	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeReverseShellEventInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScreenBaselineInfo struct {

	// 基线风险项

	BaselineFailCount *uint64 `json:"BaselineFailCount,omitempty" name:"BaselineFailCount"`
	// 基线id

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 最后检测时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 危害等级：1-低危；2-中危；3-高危；4-严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 基线名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ScreenDefendAttackLog struct {

	// 攻击时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 目标IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 目标端口

	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`
	// 攻击方式

	HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
	// 日志ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机 quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 来源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 来源端口

	SrcPort *uint64 `json:"SrcPort,omitempty" name:"SrcPort"`
	// 客户端ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 威胁类型

	VulType *string `json:"VulType,omitempty" name:"VulType"`
}

type DescribeComponentStatisticsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// ComponentName - String - 是否必填：否 - 组件名称

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeComponentStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseBackupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份列表

		List []*RansomDefenseBackup `json:"List,omitempty" name:"List"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRansomDefenseBackupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseBackupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetUserDetail struct {

	// 账号到期时间

	DisableTime *string `json:"DisableTime,omitempty" name:"DisableTime"`
	// 账号GID

	Gid *string `json:"Gid,omitempty" name:"Gid"`
	// 用户组名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// Home目录

	HomePath *string `json:"HomePath,omitempty" name:"HomePath"`
	// 是否域账号：0否， 1是, 999为空  仅windows

	IsDomain *uint64 `json:"IsDomain,omitempty" name:"IsDomain"`
	// 是否有root权限：0-否；1是，999为空: 仅linux

	IsRoot *uint64 `json:"IsRoot,omitempty" name:"IsRoot"`
	// 是否允许ssh登录，1是，0否, 999为空, 仅linux

	IsSshLogin *uint64 `json:"IsSshLogin,omitempty" name:"IsSshLogin"`
	// 用户公钥列表

	Keys []*AssetUserKeyInfo `json:"Keys,omitempty" name:"Keys"`
	// 最近登录IP

	LastLoginIp *string `json:"LastLoginIp,omitempty" name:"LastLoginIp"`
	// 最近登录位置

	LastLoginLoc *string `json:"LastLoginLoc,omitempty" name:"LastLoginLoc"`
	// 最近登录终端

	LastLoginTerminal *string `json:"LastLoginTerminal,omitempty" name:"LastLoginTerminal"`
	// 上次登录时间

	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 账号名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 密码修改时间

	PasswordChangeTime *string `json:"PasswordChangeTime,omitempty" name:"PasswordChangeTime"`
	// 密码修改设置：0-不可修改，1-可修改

	PasswordChangeType *uint64 `json:"PasswordChangeType,omitempty" name:"PasswordChangeType"`
	// 密码过期时间  仅linux

	PasswordDueTime *string `json:"PasswordDueTime,omitempty" name:"PasswordDueTime"`
	// 密码锁定时间：单位天, -1为永不锁定 999为空，仅linux

	PasswordLockDays *int64 `json:"PasswordLockDays,omitempty" name:"PasswordLockDays"`
	// 密码过期提醒：单位天

	PasswordWarnDays *uint64 `json:"PasswordWarnDays,omitempty" name:"PasswordWarnDays"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// Shell路径  仅linux

	Shell *string `json:"Shell,omitempty" name:"Shell"`
	// 是否shell登录性，0不是；1是 仅linux

	ShellLoginStatus *uint64 `json:"ShellLoginStatus,omitempty" name:"ShellLoginStatus"`
	// 账号状态：0-禁用；1-启用

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 账号UID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 账号类型：0访客用户，1标准用户，2管理员用户 ,999为空,仅windows

	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ModifyInquireRenewFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInquireRenewFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInquireRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebLocationInfo struct {

	// 启动命令

	Command *string `json:"Command,omitempty" name:"Command"`
	// 绑定IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 主目录

	MainPath *string `json:"MainPath,omitempty" name:"MainPath"`
	// 域名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 站点端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 站点协议

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 安全模块状态：0未启用，1启用，999空，仅nginx

	SafeStatus *uint64 `json:"SafeStatus,omitempty" name:"SafeStatus"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 运行用户

	User *string `json:"User,omitempty" name:"User"`
}

type BaselineCategory struct {

	// 分类Id

	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 分类名称

	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`
	// 父分类ID,如果为0则没有父分类

	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" name:"ParentCategoryId"`
}

type DescribeLogHistogramRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志的结束时间，Unix时间戳，单位ms

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 时间间隔: 单位ms

	Interval *int64 `json:"Interval,omitempty" name:"Interval"`
	// 查询语句

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeLogHistogramRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogHistogramRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyValueArrayInfo struct {

	// 需要配置键值或者元字段索引的字段

	Key *string `json:"Key,omitempty" name:"Key"`
	// 字段的索引描述信息

	Value *ValueInfo `json:"Value,omitempty" name:"Value"`
}

type NetAttackTrend struct {

	// 攻击次数

	AttackCount *uint64 `json:"AttackCount,omitempty" name:"AttackCount"`
	// 时间点，如 2023-05-06

	DateTime *string `json:"DateTime,omitempty" name:"DateTime"`
	// 攻击成功次数

	SuccAttackCount *uint64 `json:"SuccAttackCount,omitempty" name:"SuccAttackCount"`
	// 尝试攻击次数

	TryAttackCount *uint64 `json:"TryAttackCount,omitempty" name:"TryAttackCount"`
}

type ExportLicenseDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址,该字段废弃

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 任务ID,可通过任务ID去查下载任务

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportLicenseDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLicenseDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanVulSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScanVulSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanVulSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperRuleCountRequest struct {
	*tchttp.BaseRequest

	// 查询的主机uuids 一次性最多只能查100个

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *DescribeFileTamperRuleCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperRuleCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVersionStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVersionStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebPageProtectDirRequest struct {
	*tchttp.BaseRequest

	// 防护机器列表信息

	HostConfig []*ProtectHostConfig `json:"HostConfig,omitempty" name:"HostConfig"`
	// 网站防护目录地址

	ProtectDirAddr *string `json:"ProtectDirAddr,omitempty" name:"ProtectDirAddr"`
	// 网站防护目录名称

	ProtectDirName *string `json:"ProtectDirName,omitempty" name:"ProtectDirName"`
	// 防护文件类型,分号分割 ;

	ProtectFileType *string `json:"ProtectFileType,omitempty" name:"ProtectFileType"`
}

func (r *ModifyWebPageProtectDirRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebPageProtectDirRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeOpenPortTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpenPortTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 进程统计列表数据数组。

		ProcessStatistics []*ProcessStatistics `json:"ProcessStatistics,omitempty" name:"ProcessStatistics"`
		// 进程统计列表记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcessStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProcessStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenDefenseTrendsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计详情图标数据

		TrendsChart []*ScreenTrendsChart `json:"TrendsChart,omitempty" name:"TrendsChart"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenDefenseTrendsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenDefenseTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBanWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBanWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBanWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetAttackWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetAttackWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetAttackWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMalwareScanTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteMalwareScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMalwareScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单规则

	Rules *LoginWhiteListsRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyLoginWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScreenMachine struct {

	// 基线风险数。

	BaselineNum *int64 `json:"BaselineNum,omitempty" name:"BaselineNum"`
	// 内核版本

	CoreVersion *string `json:"CoreVersion,omitempty" name:"CoreVersion"`
	// cpu 负载状态

	CpuLoad *string `json:"CpuLoad,omitempty" name:"CpuLoad"`
	// cpu 核数

	CpuSize *float64 `json:"CpuSize,omitempty" name:"CpuSize"`
	// 网络风险数。

	CyberAttackNum *int64 `json:"CyberAttackNum,omitempty" name:"CyberAttackNum"`
	// 硬盘使用率 %

	DiskLoad *string `json:"DiskLoad,omitempty" name:"DiskLoad"`
	// 硬盘容量GB

	DiskSize *float64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 入侵事件数

	InvasionNum *int64 `json:"InvasionNum,omitempty" name:"InvasionNum"`
	// 附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机系统。

	MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`
	// 大屏主机状态 0：未安装agent，1：离线状态，2:离线-风险，3：离线-严重
	// 4：安装设备-正常，5：安装设备-正常 且是专业版或旗舰版，6：安装设备-风险（网络攻击事件>0） ，
	//  7：安装设备-风险（网络攻击事件>0 且是专业版或旗舰版，8：安装设备-严重（入侵检测事件>0），
	// 9：安装设备-严重（入侵检测事件>0）且是专业版或旗舰版

	MachineStatus *uint64 `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 机器所属专区类型 CVM 云服务器, BM 黑石, ECM 边缘计算, LH 轻量应用服务器 ,Other 混合云专区

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 主机外网IP。

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 内存负载率%

	MemLoad *string `json:"MemLoad,omitempty" name:"MemLoad"`
	// 内存容量 GB

	MemSize *float64 `json:"MemSize,omitempty" name:"MemSize"`
	// CVM或BM机器唯一Uuid。

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 风险状态。
	// <li>SAFE：安全</li>
	// <li>RISK：风险</li>
	// <li>UNKNOWN：未知</li>

	SecurityStatus *string `json:"SecurityStatus,omitempty" name:"SecurityStatus"`
	// 云镜客户端唯一Uuid，若客户端长时间不在线将返回空字符。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 漏洞数。

	VulNum *int64 `json:"VulNum,omitempty" name:"VulNum"`
}

type DescribeAssetTypeTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据库Top5

		Database []*AssetKeyVal `json:"Database,omitempty" name:"Database"`
		// 端口Top5

		Port []*AssetKeyVal `json:"Port,omitempty" name:"Port"`
		// 进程Top5

		Process []*AssetKeyVal `json:"Process,omitempty" name:"Process"`
		// 软件Top5

		Software []*AssetKeyVal `json:"Software,omitempty" name:"Software"`
		// 账号Top5

		User []*AssetKeyVal `json:"User,omitempty" name:"User"`
		// Web应用Top5

		WebApp []*AssetKeyVal `json:"WebApp,omitempty" name:"WebApp"`
		// Web框架Top5

		WebFrame []*AssetKeyVal `json:"WebFrame,omitempty" name:"WebFrame"`
		// Web站点Top5

		WebLocation []*AssetKeyVal `json:"WebLocation,omitempty" name:"WebLocation"`
		// Web服务Top5

		WebService []*AssetKeyVal `json:"WebService,omitempty" name:"WebService"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetTypeTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTypeTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineHostTopRequest struct {
	*tchttp.BaseRequest

	// 策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 动态top值

	Top *uint64 `json:"Top,omitempty" name:"Top"`
}

func (r *DescribeBaselineHostTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineHostTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重点关注漏洞总数

		FollowVulCount *uint64 `json:"FollowVulCount,omitempty" name:"FollowVulCount"`
		// 漏洞总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞列表

		VulInfoList []*VulInfoList `json:"VulInfoList,omitempty" name:"VulInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebAppListRequest struct {
	*tchttp.BaseRequest

	// 可选排序：[FirstTime|PluginCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 应用名</li>
	// <li>Domain - String - 是否必填：否 - 站点域名</li>
	// <li>Type - int - 是否必填：否 - 服务类型：
	// 0：全部
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:Jetty
	// 8:IHS
	// 9:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetWebAppListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseStrategyListRequest struct {
	*tchttp.BaseRequest

	// 排序字段支持CreateTime, MachineCount

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Ips - String - 是否必填：否 - 通过ip查询 </li>
	// <li>MachineNames - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>Names - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>Dirs - String - 是否必填：否 - 诱饵目录 </li>
	// <li>Status - String - 是否必填：否 - 策略状态：0关闭，1开启 </li>
	// <li>BackupType - String - 是否必填：否 - 备份模式：0-按周;1-按天 </li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 分页参数 最大100条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方法 ASC DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRansomDefenseStrategyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseStrategyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportBruteAttacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专业周报密码破解数组。

		WeeklyReportBruteAttacks []*WeeklyReportBruteAttack `json:"WeeklyReportBruteAttacks,omitempty" name:"WeeklyReportBruteAttacks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportBruteAttacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBashRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBashRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttackListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密码破解列表

		BruteAttackList []*BruteAttackInfo `json:"BruteAttackList,omitempty" name:"BruteAttackList"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBruteAttackListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBruteAttackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousRequestsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求记录数组。

		MaliciousRequests []*MaliciousRequest `json:"MaliciousRequests,omitempty" name:"MaliciousRequests"`
		// 记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaliciousRequestsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportJavaMemShellsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：Keywords: ip或者主机名模糊查询, Type，Status精确匹配，CreateBeginTime，CreateEndTime时间段

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 导出字段

	Where []*string `json:"Where,omitempty" name:"Where"`
}

func (r *ExportJavaMemShellsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportJavaMemShellsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 15天内新增的主机数

		AddedOnTheFifteen *uint64 `json:"AddedOnTheFifteen,omitempty" name:"AddedOnTheFifteen"`
		// 主机安全客户端总数的总数

		AgentsAll *uint64 `json:"AgentsAll,omitempty" name:"AgentsAll"`
		// 主机安全客户端基础版的总数

		AgentsBasic *uint64 `json:"AgentsBasic,omitempty" name:"AgentsBasic"`
		// 主机安全客户端 离线+关机 的总数

		AgentsOffline *uint64 `json:"AgentsOffline,omitempty" name:"AgentsOffline"`
		// 主机安全客户端在线的总数

		AgentsOnline *uint64 `json:"AgentsOnline,omitempty" name:"AgentsOnline"`
		// 主机安全客户端专业版的总数

		AgentsPro *uint64 `json:"AgentsPro,omitempty" name:"AgentsPro"`
		// 7天内到期的预付费专业版总数

		AgentsProExpireWithInSevenDays *uint64 `json:"AgentsProExpireWithInSevenDays,omitempty" name:"AgentsProExpireWithInSevenDays"`
		// 旗舰版主机数

		FlagshipMachineCnt *uint64 `json:"FlagshipMachineCnt,omitempty" name:"FlagshipMachineCnt"`
		// 云主机总数

		MachinesAll *uint64 `json:"MachinesAll,omitempty" name:"MachinesAll"`
		// 云主机没有安装主机安全客户端的总数

		MachinesUninstalled *uint64 `json:"MachinesUninstalled,omitempty" name:"MachinesUninstalled"`
		// 已离线总数

		Offline *uint64 `json:"Offline,omitempty" name:"Offline"`
		// 保护天数

		ProtectDays *uint64 `json:"ProtectDays,omitempty" name:"ProtectDays"`
		// 风险主机总数

		RiskMachine *uint64 `json:"RiskMachine,omitempty" name:"RiskMachine"`
		// 已关机总数

		Shutdown *uint64 `json:"Shutdown,omitempty" name:"Shutdown"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetUserBaseInfo struct {

	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 账号GID

	Gid *string `json:"Gid,omitempty" name:"Gid"`
	// Home目录

	HomePath *string `json:"HomePath,omitempty" name:"HomePath"`
	// 是否域账号：0否， 1是，2否, 999为空  仅windows

	IsDomain *uint64 `json:"IsDomain,omitempty" name:"IsDomain"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 是否有root权限：0-否；1是，999为空: 仅linux

	IsRoot *uint64 `json:"IsRoot,omitempty" name:"IsRoot"`
	// 是否允许ssh登录，1是，0否, 999为空, 仅linux

	IsSshLogin *uint64 `json:"IsSshLogin,omitempty" name:"IsSshLogin"`
	// 是否有sudo权限，1是，0否, 999为空, 仅linux

	IsSudo *uint64 `json:"IsSudo,omitempty" name:"IsSudo"`
	// 上次登录时间

	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
	// 登录方式：0-不可登录；1-只允许key登录；2只允许密码登录；3-允许key和密码，999为空，仅linux

	LoginType *uint64 `json:"LoginType,omitempty" name:"LoginType"`
	//
	//  附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 账号名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 密码修改时间

	PasswordChangeTime *string `json:"PasswordChangeTime,omitempty" name:"PasswordChangeTime"`
	// 密码过期时间  仅linux

	PasswordDueTime *string `json:"PasswordDueTime,omitempty" name:"PasswordDueTime"`
	// 密码锁定时间：单位天, -1为永不锁定 999为空，仅linux

	PasswordLockDays *int64 `json:"PasswordLockDays,omitempty" name:"PasswordLockDays"`
	// 密码状态：1正常 2即将过期 3已过期 4已锁定 999为空 仅linux

	PasswordStatus *int64 `json:"PasswordStatus,omitempty" name:"PasswordStatus"`
	// 主机业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// Shell路径  仅linux

	Shell *string `json:"Shell,omitempty" name:"Shell"`
	// 是否shell登录性，0不是；1是 仅linux

	ShellLoginStatus *uint64 `json:"ShellLoginStatus,omitempty" name:"ShellLoginStatus"`
	// 账号状态：0-禁用；1-启用

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 账号UID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 账号类型：0访客用户，1标准用户，2管理员用户 ,999为空,仅windows

	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type VulStoreListInfo struct {

	// 漏洞攻击热度

	AttackLevel *uint64 `json:"AttackLevel,omitempty" name:"AttackLevel"`
	// cve编号

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// 漏洞是否支持自动修复
	// 0-windows/linux均关闭; 1-windows/linux均打开; 2-仅linux; 3-仅windows

	FixSwitch *uint64 `json:"FixSwitch,omitempty" name:"FixSwitch"`
	// 漏洞级别

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 漏洞检测方法 0 - 版本比对, 1 - POC验证

	Method *uint64 `json:"Method,omitempty" name:"Method"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 发布时间

	PublishDate *string `json:"PublishDate,omitempty" name:"PublishDate"`
	// 漏洞是否支持防御
	// 0:不支持 1:支持

	SupportDefense *uint64 `json:"SupportDefense,omitempty" name:"SupportDefense"`
	// 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞 0= 应急漏洞

	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`
	// 漏洞ID

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

type ExportMalwaresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportMalwaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportReverseShellEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportReverseShellEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWarningSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWarningSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWarningSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RansomDefenseRollbackRequest struct {
	*tchttp.BaseRequest

	// 快照时间

	BackupTime *string `json:"BackupTime,omitempty" name:"BackupTime"`
	// 需要回滚的硬盘信息，硬盘直接用;分隔，留空为全部已快照磁盘： disk-id1|disk-name1;disk-id2|disk-name2

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *RansomDefenseRollbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RansomDefenseRollbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineHostTopList struct {

	// 事件等级与次数列表

	EventLevelList []*BaselineEventLevelInfo `json:"EventLevelList,omitempty" name:"EventLevelList"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 计算权重的分数

	Score *uint64 `json:"Score,omitempty" name:"Score"`
}

type DescribeHostListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>AliasNameOrIp- String - 主机名或ip筛选</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 地域,对应 DescribeMachineRegions 的 RegionCode

	HostArea *string `json:"HostArea,omitempty" name:"HostArea"`
	// 机器类型专区：公有云cvm专区:CVM，公有云黑石专区：BM，TCE裸金属：TCE，公有云ECM：ECM

	HostRegion *string `json:"HostRegion,omitempty" name:"HostRegion"`
	// 分页参数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机标签id串,多个用英文逗号分隔

	TagIds *string `json:"TagIds,omitempty" name:"TagIds"`
}

func (r *DescribeHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RansomDefenseStrategyMachineInfo struct {

	// 指定硬盘列表，为空时表示所有硬盘：disk_id1|disk_name1;disk_id2|disk_name2

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportAssetAppListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：[FirstTime|ProcessCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>AppName- string - 是否必填：否 - 应用名搜索</li>
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Type - int - 是否必填：否 - 类型	: 仅linux
	// 0: 全部
	// 1: 运维
	// 2 : 数据库
	// 3 : 安全
	// 4 : 可疑应用
	// 5 : 系统架构
	// 6 : 系统应用
	// 7 : WEB服务
	// 99:其他</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式，asc升序 或 desc降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询指定Quuid主机的信息

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetAppListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetAppListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志导出ID

		ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOpenPortTaskRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *CreateOpenPortTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOpenPortTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRiskDnsEventRequest struct {
	*tchttp.BaseRequest

	// 恶意请求记录ID数组，(最大100条), 为空时全部删除

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteRiskDnsEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRiskDnsEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginWhiteRecordRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 服务器列表

	Hosts []*HostInfo `json:"Hosts,omitempty" name:"Hosts"`
	// 白名单ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否对所有服务器生效，0-否，1-是

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 地域列表

	Places []*Place `json:"Places,omitempty" name:"Places"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 更新后记录的白名单维度信息

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 白名单用户（多个用户逗号隔开）

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *ModifyLoginWhiteRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginWhiteRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileTamperRuleDetail struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则(默认否) 0：否 ，1：是

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 风险等级 0：无， 1: 高危， 2:中危， 3: 低危

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则

	Rule []*FileTamperRule `json:"Rule,omitempty" name:"Rule"`
	// 状态 0: 启用 1: 已关闭

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 生效主机的总数

	UuidTotalCount *uint64 `json:"UuidTotalCount,omitempty" name:"UuidTotalCount"`
	// 生效主机uuid,空表示全部主机，通过参数可控制返回的条数

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
}

type StandardModeConfig struct {

	// 阻断时长，单位：秒

	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`
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

type EditReverseShellRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditReverseShellRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditReverseShellRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditBashRuleRequest struct {
	*tchttp.BaseRequest

	// 主机IP(IsGlobal为0时，Uuid或Hostip必填一个)

	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
	// 规则ID(新增时不填)

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则(默认否)

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 危险等级(1: 高危 2:中危 3: 低危)

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 正则表达式

	Rule *string `json:"Rule,omitempty" name:"Rule"`
	// 客户端ID(IsGlobal为0时，Uuid或Hostip必填一个)

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *EditBashRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditBashRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulEffectHostListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// <li>AliasName - String - 主机名筛选</li>
	// <li>TagIds - String - 主机标签id串，多个用英文用逗号分隔</li>
	// <li>Status - String - 状态：0-待处理 1-忽略  3-已修复  5-检测中  6-修复中  8-修复失败</li>
	// <li>Uuid - String数组 - Uuid串数组</li>
	// <li>Version - String数组 - 付费版本数组："Flagship"-旗舰版 "PRO_VERSION"-专业版 "BASIC_VERSION"-基础版</li>
	// <li>InstanceState - String数组 - 实例状态数组："PENDING"-创建中 "LAUNCH_FAILED"-创建失败 "RUNNING"-运行中 "STOPPED"-关机 "STARTING"-开机中 "STOPPING"-关机中 "REBOOTING"-重启中 "SHUTDOWN"-待销毁 "TERMINATING"-销毁中 "UNKNOWN"-未知（针对非腾讯云机器，且客户端离线的场景） </li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页limit 最大100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *DescribeVulEffectHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulEffectHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEventAttackStatusRequest struct {
	*tchttp.BaseRequest

	// 是否更新全部，即是否对所有的事件进行操作，当ids 不为空时，此参数无效

	All *bool `json:"All,omitempty" name:"All"`
	// 排除的id

	ExcludeId []*uint64 `json:"ExcludeId,omitempty" name:"ExcludeId"`
	// 过滤条件。
	// <li>Type - String 攻击状态 0: 尝试攻击 1: 攻击成功 - 是否必填: 否</li>
	// <li>Status - String 事件处理状态 0：待处理 1：已处理 2： 已加白 3： 已忽略 4：已删除  - 是否必填: 否</li>
	// <li>SrcIP - String 来源IP - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>
	// <li>MachineName - String 主机名称 - 是否必填: 否</li>
	// <li>InstanceID - String 主机实例ID - 是否必填: 否</li>
	// <li>Quuids - String 主机cvm uuid - 是否必填: 否</li>
	//

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 需要修改的事件id 数组，支持批量

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 0：待处理 1：已处理 2： 已加白 3： 已忽略 4：已删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyEventAttackStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEventAttackStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLicenseUnBindsRequest struct {
	*tchttp.BaseRequest

	// 是否全部机器(当全部机器数大于当前订单可用授权数时,多余机器会被跳过)

	IsAll *bool `json:"IsAll,omitempty" name:"IsAll"`
	// 是否自动缩容,如果授权数剩余1,则销毁订单

	IsShrink *bool `json:"IsShrink,omitempty" name:"IsShrink"`
	// 授权类型

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 需要绑定的机器quuid列表, 当IsAll = false 时必填,反之忽略该参数.
	// 最大长度=100

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *ModifyLicenseUnBindsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLicenseUnBindsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetJarListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步下载任务ID，需要配合ExportTasks接口使用

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetJarListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetJarListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmAttributeRequest struct {
	*tchttp.BaseRequest

	// 告警项目。
	// <li>Offline：防护软件离线</li>
	// <li>Malware：发现木马文件</li>
	// <li>NonlocalLogin：发现异地登录行为</li>
	// <li>CrackSuccess：被暴力破解成功</li>

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 告警项目属性。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>

	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyAlarmAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// <li>StrategyId- Uint64 - 基线策略id</li>
	// <li>Status - Uint64 - 事件状态：0-未通过，1-忽略，3-通过，5-检测中</li>
	// <li>BaselineName  - String - 基线名称</li>
	// <li>AliasName- String - 服务器名称/服务器ip</li>
	// <li>Uuid- String - 主机uuid</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 已废弃

	IfDetail *uint64 `json:"IfDetail,omitempty" name:"IfDetail"`
}

func (r *ExportBaselineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountStatistics struct {

	// 主机数量。

	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
	// 用户名。

	Username *string `json:"Username,omitempty" name:"Username"`
}

type AssetUserKeyInfo struct {

	// 公钥备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 加密方式

	EncryptType *string `json:"EncryptType,omitempty" name:"EncryptType"`
	// 公钥值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type AssetDatabaseBaseInfo struct {

	// 二进制路径

	BinPath *string `json:"BinPath,omitempty" name:"BinPath"`
	// 配置文件路径

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 数据路径

	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`
	// 错误日志路径

	ErrorLogPath *string `json:"ErrorLogPath,omitempty" name:"ErrorLogPath"`
	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 数据库ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 绑定IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 日志文件路径

	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`
	//  附加信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 数据库名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 启动参数

	Param *string `json:"Param,omitempty" name:"Param"`
	// 运行权限

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 插件路径

	PlugInPath *string `json:"PlugInPath,omitempty" name:"PlugInPath"`
	// 监听端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 主机业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 协议

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 运行用户

	User *string `json:"User,omitempty" name:"User"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeWebPageEventListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：CreateTime 或 RestoreTime，默认为CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>EventType - String - 是否必填：否 - 事件类型</li>
	// <li>EventStatus - String - 是否必填：否 - 事件状态</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，0降序，1升序，默认为0

	Order *uint64 `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeWebPageEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebPageEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马Id数组（最大100条）

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *RecoverMalwaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecoverMalwaresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulSettingRequest struct {
	*tchttp.BaseRequest

	// 是否开启扫描 1开启 0不开启

	EnableScan *uint64 `json:"EnableScan,omitempty" name:"EnableScan"`
	// 扫描结束时间，如：08:00

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 扫描开始时间，如：00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 定期检测间隔时间（天）

	TimerInterval *uint64 `json:"TimerInterval,omitempty" name:"TimerInterval"`
	// 定期检测时间，如：02:10:50

	TimerTime *string `json:"TimerTime,omitempty" name:"TimerTime"`
	// 为空默认扫描全部专业版、旗舰版、普惠版主机，不为空只扫描选中主机

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
	// 漏洞类型：1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞, 以数组方式传参[1,2]

	VulCategories []*uint64 `json:"VulCategories,omitempty" name:"VulCategories"`
	// 是否是应急漏洞 0 否 1 是

	VulEmergency *uint64 `json:"VulEmergency,omitempty" name:"VulEmergency"`
	// 危害等级：1-低危；2-中危；3-高危；4-严重,以数组方式传参[1,2,3]

	VulLevels []*uint64 `json:"VulLevels,omitempty" name:"VulLevels"`
}

func (r *ScanVulSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanVulSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UntrustMaliciousRequestRequest struct {
	*tchttp.BaseRequest

	// 受信任记录ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *UntrustMaliciousRequestRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UntrustMaliciousRequestRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineRuleInfo struct {

	// 检测项描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 唯一事件ID

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 修复建议

	FixMessage *string `json:"FixMessage,omitempty" name:"FixMessage"`
	// 最后检测时间

	LastScanAt *string `json:"LastScanAt,omitempty" name:"LastScanAt"`
	// 危害等级

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 检测项id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 检测项名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 具体原因说明

	RuleRemark *string `json:"RuleRemark,omitempty" name:"RuleRemark"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 唯一Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type Place struct {

	// 城市 ID。

	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
	// 国家ID，暂只支持国内：1。

	CountryId *uint64 `json:"CountryId,omitempty" name:"CountryId"`
	// 位置名称

	Location *string `json:"Location,omitempty" name:"Location"`
	// 省份 ID。

	ProvinceId *uint64 `json:"ProvinceId,omitempty" name:"ProvinceId"`
}

type DescribeMalwareFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 木马文件下载地址

		FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwareFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanVulSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一键扫描超时时长，如：1800秒（s）

		ClickTimeout *uint64 `json:"ClickTimeout,omitempty" name:"ClickTimeout"`
		// 是否开启

		EnableScan *uint64 `json:"EnableScan,omitempty" name:"EnableScan"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 定期检测间隔时间（天）

		TimerInterval *uint64 `json:"TimerInterval,omitempty" name:"TimerInterval"`
		// 定期检测时间，如：00:00

		TimerTime *string `json:"TimerTime,omitempty" name:"TimerTime"`
		// 为空默认扫描全部专业版、旗舰版、普惠版主机，不为空只扫描选中主机

		Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
		// 漏洞类型：1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞

		VulCategories *string `json:"VulCategories,omitempty" name:"VulCategories"`
		// 是否紧急漏洞：0-否 1-是

		VulEmergency *uint64 `json:"VulEmergency,omitempty" name:"VulEmergency"`
		// 危害等级：1-低危；2-中危；3-高危；4-严重 (多选英文逗号分隔)

		VulLevels *string `json:"VulLevels,omitempty" name:"VulLevels"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanVulSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanVulSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulEmergentMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞紧急通知数据

		EmergentMsgList []*VulEmergentMsgInfo `json:"EmergentMsgList,omitempty" name:"EmergentMsgList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulEmergentMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulEmergentMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogKafkaStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogKafkaStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogKafkaStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteListNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白名单列表

		LoginWhiteLists []*LoginWhiteListsNew `json:"LoginWhiteLists,omitempty" name:"LoginWhiteLists"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginWhiteListNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginWhiteListNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionActivityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVersionActivityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenProVersionActivityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineItemListRequest struct {
	*tchttp.BaseRequest

	// 0:过滤的结果导出；1:全部导出

	ExportAll *int64 `json:"ExportAll,omitempty" name:"ExportAll"`
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	// <li>HostName - string - 是否必填：否 - 主机名</li>
	// <li>HostIp - string - 是否必填：否 - 主机IP</li>
	// <li>ItemId - String - 是否必填：否 - 检测项Id</li>
	// <li>ItemName - String - 是否必填：否 - 项名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态[0:未通过|3:通过|5:检测中]</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开时时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBaselineItemListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineItemListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineRemarkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMachineRemarkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMachineRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityTrend struct {

	// 事件时间。

	Date *string `json:"Date,omitempty" name:"Date"`
	// 事件数量。

	EventNum *uint64 `json:"EventNum,omitempty" name:"EventNum"`
}

type CheckBashRuleParamsRequest struct {
	*tchttp.BaseRequest

	// 校验内容 Name或Rule ，两个都要校验时逗号分割

	CheckField *string `json:"CheckField,omitempty" name:"CheckField"`
	// 在事件列表中新增白名时需要提交事件ID

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 编辑时传的规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 填入的规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户填入的正则表达式："正则表达式" 需与 "提交EventId对应的命令内容" 相匹配

	Rule *string `json:"Rule,omitempty" name:"Rule"`
}

func (r *CheckBashRuleParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBashRuleParamsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityProtectionStatRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecurityProtectionStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityProtectionStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServersAndRiskAndFirstInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 今日新增风险文件数

		AddRiskFileCount *uint64 `json:"AddRiskFileCount,omitempty" name:"AddRiskFileCount"`
		// 是否试用：true-是，false-否

		IsFirstCheck *bool `json:"IsFirstCheck,omitempty" name:"IsFirstCheck"`
		// 风险文件数

		RiskFileCount *uint64 `json:"RiskFileCount,omitempty" name:"RiskFileCount"`
		// 木马最近检测时间

		ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
		// 受影响服务器台数

		ServersCount *uint64 `json:"ServersCount,omitempty" name:"ServersCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServersAndRiskAndFirstInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServersAndRiskAndFirstInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanTaskAgainRequest struct {
	*tchttp.BaseRequest

	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 自选服务器时生效，主机quuid的string数组

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 扫描超时时长

	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitempty" name:"TimeoutPeriod"`
}

func (r *ScanTaskAgainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanTaskAgainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulInfoCvssRequest struct {
	*tchttp.BaseRequest

	// 漏洞id

	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *DescribeVulInfoCvssRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulInfoCvssRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetProcessBaseInfo struct {

	// 进程说明

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 首次采集时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 进程用户组

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 数字签名:0无，1有， 999 空，仅windows

	HasSign *uint64 `json:"HasSign,omitempty" name:"HasSign"`
	// 是否通过安装包安装：:0否，1是， 999 空，仅linux

	InstallByPackage *uint64 `json:"InstallByPackage,omitempty" name:"InstallByPackage"`
	// 是否新增[0:否|1:是]

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	//
	//  附加信息
	//

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机内网IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机外网IP

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 进程MD5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 进程名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 软件包名

	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
	// 启动参数

	Param *string `json:"Param,omitempty" name:"Param"`
	// 父进程名称

	ParentProcessName *string `json:"ParentProcessName,omitempty" name:"ParentProcessName"`
	// 进程路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 进程ID

	Pid *string `json:"Pid,omitempty" name:"Pid"`
	// 父进程ID

	Ppid *string `json:"Ppid,omitempty" name:"Ppid"`
	// 主机业务组ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 进程状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 主机标签

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 进程TTY

	Tty *string `json:"Tty,omitempty" name:"Tty"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 运行用户

	User *string `json:"User,omitempty" name:"User"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 进程版本

	Version *string `json:"Version,omitempty" name:"Version"`
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

type CreateBuyBindTaskRequest struct {
	*tchttp.BaseRequest

	// 订单号

	DealName *string `json:"DealName,omitempty" name:"DealName"`
	// 是否全选机器

	IsAll *bool `json:"IsAll,omitempty" name:"IsAll"`
	// 可选参数: 1专业版-包年包月 , 2 旗舰版-包年包月

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 机器列表

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
}

func (r *CreateBuyBindTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBuyBindTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskProcessEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异常进程列表

		List []*RiskProcessEvent `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskProcessEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskProcessEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoNotNoticeBanTipsRequest struct {
	*tchttp.BaseRequest
}

func (r *DoNotNoticeBanTipsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoNotNoticeBanTipsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetAttackWhiteListRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteNetAttackWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetAttackWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScreenEmergentMsg struct {

	// 通知内容

	Text *string `json:"Text,omitempty" name:"Text"`
	// 通知标签/标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 跳转类型：0=漏洞管理

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type ExportBaselineItemDetectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBaselineItemDetectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineItemDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MaliciousRequestWhiteListInfo struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 白名单id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 备注

	Mark *string `json:"Mark,omitempty" name:"Mark"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type ReverseShellEventInfo struct {

	// 命令详情

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 产生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 检测方法

	DetectBy *uint64 `json:"DetectBy,omitempty" name:"DetectBy"`
	// 目标IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 目标端口

	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`
	// 进程路径

	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`
	// 描述

	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
	// 主机内网IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// ID 主键

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机在线状态 OFFLINE  ONLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 主机外网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 处理时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 父进程用户组

	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`
	// 父进程名

	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`
	// 父进程路径

	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`
	// 父进程用户

	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`
	// 进程名

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程树 json  pid:进程id，exe:文件路径 ，account:进程所属用组和用户 ,cmdline:执行命令，ssh_service: SSH服务ip, ssh_soure:登录源

	PsTree *string `json:"PsTree,omitempty" name:"PsTree"`
	// 主机ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 参考链接

	References []*string `json:"References,omitempty" name:"References"`
	// 处理状态：0-待处理 2-白名单 3-已处理 4-已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 建议方案

	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 执行用户组

	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`
	// 执行用户

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 云镜UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type StrategyDetailInfo struct {

	// 添加时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 策略详情

	StrategyInfo *bool `json:"StrategyInfo,omitempty" name:"StrategyInfo"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 策略类型

	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeBaselineStrategyListRequest struct {
	*tchttp.BaseRequest

	// 规则开关，1：打开 0：关闭  2:全部

	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`
	// 分页参数 最大100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBaselineStrategyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineStrategyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulListRequest struct {
	*tchttp.BaseRequest

	// 可选排序字段 Level，LastTime，HostCount

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略</li>
	// <li>ModifyTime - String - 是否必填：否 - 最近发生时间</li>
	// <li>Uuid- String - 是否必填：否 - 主机uuid查询</li>
	// <li>VulName- string -</li>
	// <li>VulCategory- string - 是否必填：否 - 漏洞类别 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞</li>
	// <li>IsSupportDefense - int- 是否必填：否 - 是否支持防御 0:不支持 1:支持</li>
	// <li>Labels- string- 是否必填：否 - 标签搜索</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序顺序：desc  默认asc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBashEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件下载链接地址。

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBashEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellPluginListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// java内存马插件列表

		List []*JavaMemShellPluginSetting `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJavaMemShellPluginListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellPluginListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenGeneralStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// name 的值: 在线，关机/离线,未安装，
		// value : 表示对应的数量

		Machines []*ScreenNameValue `json:"Machines,omitempty" name:"Machines"`
		// name 的值: 旗舰版，专业版，基础版
		// value : 表示对应的数量

		Protection []*ScreenNameValue `json:"Protection,omitempty" name:"Protection"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenGeneralStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScreenGeneralStatResponse) FromJsonString(s string) error {
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

type ModifyRiskEventsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 此次操作是否是异步操作，0：操作已完成，1：正在异步操作中，前端需要通过DescribeRiskBatchStatus 查询操作是否完成

		IsSync *uint64 `json:"IsSync,omitempty" name:"IsSync"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRiskEventsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskEventsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BashEventsInfo struct {

	// 执行命令

	BashCmd *string `json:"BashCmd,omitempty" name:"BashCmd"`
	// 发生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 检测来源 0:bash日志 1:实时监控

	DetectBy *int64 `json:"DetectBy,omitempty" name:"DetectBy"`
	// 进程名称

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 描述

	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
	// 主机内网IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 数据ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 主机在线状态 OFFLINE  ONLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 主机外网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 处理时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 进程号

	Pid *string `json:"Pid,omitempty" name:"Pid"`
	// 平台类型

	Platform *uint64 `json:"Platform,omitempty" name:"Platform"`
	// 进程树 json  pid:进程id，exe:文件路径 ，account:进程所属用组和用户 ,cmdline:执行命令，ssh_service: SSH服务ip, ssh_soure:登录源

	PsTree *string `json:"PsTree,omitempty" name:"PsTree"`
	// 主机ID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 参考链接

	References []*string `json:"References,omitempty" name:"References"`
	// 自动生成的正则表达式

	RegexBashCmd *string `json:"RegexBashCmd,omitempty" name:"RegexBashCmd"`
	// 规则类别  0=系统规则，1=用户规则

	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`
	// 规则ID,等于0表示已规则已被删除或生效范围已修改

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则等级：1-高 2-中 3-低

	RuleLevel *uint64 `json:"RuleLevel,omitempty" name:"RuleLevel"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 处理状态： 0 = 待处理 1= 已处理, 2 = 已加白， 3= 已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 建议方案

	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 登录用户

	User *string `json:"User,omitempty" name:"User"`
	// 云镜ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAgentVulsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | FIXED：已修复）

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 客户端UUID。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 漏洞类型。
	// <li>WEB: Web应用漏洞</li>
	// <li>SYSTEM：系统组件漏洞</li>
	// <li>BASELINE：安全基线</li>

	VulType *string `json:"VulType,omitempty" name:"VulType"`
}

func (r *DescribeAgentVulsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentVulsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousRequestWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白名单信息列表

		List []*MaliciousRequestWhiteListInfo `json:"List,omitempty" name:"List"`
		// 分页查询记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaliciousRequestWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousRequestWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineHostTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机基线策略事件Top

		BaselineHostTopList []*BaselineHostTopList `json:"BaselineHostTopList,omitempty" name:"BaselineHostTopList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineHostTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineHostTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineEffectHostListRequest struct {
	*tchttp.BaseRequest

	// 基线id

	BaselineId *uint64 `json:"BaselineId,omitempty" name:"BaselineId"`
	// 基线名称

	BaselineName *string `json:"BaselineName,omitempty" name:"BaselineName"`
	// 筛选条件
	// <li>AliasName- String- 主机别名</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 主机uuid数组

	UuidList []*string `json:"UuidList,omitempty" name:"UuidList"`
}

func (r *ExportBaselineEffectHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineEffectHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BanWhiteListDetail struct {

	// 创建白名单时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 白名单ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 白名单是否全局

	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 机器IP

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 机器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 修改白名单时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 机器的UUID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 白名单别名

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 阻断来源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 主机安全程序的UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ChangeRuleEventsIgnoreStatusRequest struct {
	*tchttp.BaseRequest

	// 事件id数组

	EventIdList []*uint64 `json:"EventIdList,omitempty" name:"EventIdList"`
	// 忽略状态 0:取消忽略 ； 1:忽略

	IgnoreStatus *uint64 `json:"IgnoreStatus,omitempty" name:"IgnoreStatus"`
	// 检测项id数组

	RuleIdList []*uint64 `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

func (r *ChangeRuleEventsIgnoreStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeRuleEventsIgnoreStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginWhiteInfoRequest struct {
	*tchttp.BaseRequest

	// 更新白名单信息实体

	HostLoginWhiteObj *UpdateHostLoginWhiteObj `json:"HostLoginWhiteObj,omitempty" name:"HostLoginWhiteObj"`
}

func (r *ModifyLoginWhiteInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginWhiteInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportMalwaresRequest struct {
	*tchttp.BaseRequest

	// 排序值 CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 过滤参数。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>FilePath - String - 是否必填：否 - 路径筛选</li>
	// <li>VirusName - String - 是否必填：否 - 描述筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选-开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选-结束时间</li>
	// <li>Status - String - 是否必填：否 - 状态筛选</li>

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量 默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序 方式 ，ASC，DESC

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportMalwaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportMalwaresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetAttackWhiteListRequest struct {
	*tchttp.BaseRequest

	// 是否加白所有符合该规则的告警 ，1:处理,0:不处理

	DealOldEvents *uint64 `json:"DealOldEvents,omitempty" name:"DealOldEvents"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// quuid 列表

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
	// 是否全部主机； 0否，1是。

	Scope *uint64 `json:"Scope,omitempty" name:"Scope"`
	// 来源IP 单IP:1.1.1.1  IP范围:1.1.1.1-1.1.2.1  IP范围：1.1.1.0/24

	SrcIp []*string `json:"SrcIp,omitempty" name:"SrcIp"`
}

func (r *ModifyNetAttackWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetAttackWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口列表。

		OpenPorts []*OpenPort `json:"OpenPorts,omitempty" name:"OpenPorts"`
		// 端口列表记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenPortsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpenPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 询价返回数据

		Data *string `json:"Data,omitempty" name:"Data"`
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

type ModifyVulDefenceSettingRequest struct {
	*tchttp.BaseRequest

	// 防御开关，0 关闭 1 开启

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 作用弄范围内旗舰版主机列表

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
	// 1 全部旗舰版主机，0 Quuids列表主机

	Scope *uint64 `json:"Scope,omitempty" name:"Scope"`
}

func (r *ModifyVulDefenceSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSearchTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：成功，非0：失败

		Status *uint64 `json:"Status,omitempty" name:"Status"`
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

type DescribeBaselineDetectOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检测服务器数

		HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
		// 检测项

		ItemCount *int64 `json:"ItemCount,omitempty" name:"ItemCount"`
		// 最近一次检测未通过个数

		LatestNotPassCount *int64 `json:"LatestNotPassCount,omitempty" name:"LatestNotPassCount"`
		// 最近一次检测通过个数

		LatestPassCount *int64 `json:"LatestPassCount,omitempty" name:"LatestPassCount"`
		// 通过率*100%

		PassRate *int64 `json:"PassRate,omitempty" name:"PassRate"`
		// 检测策略项

		PolicyCount *int64 `json:"PolicyCount,omitempty" name:"PolicyCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineDetectOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineDetectOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareWhiteListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [EventsCount]

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件。
	// <li>HostIp - String - 是否必填：否 - 主机ip查询 </li>
	// <li>FileName - String - 是否必填：否 - 文件名称查询 </li>
	// <li>FileDirectory - String - 是否必填：否 - 文件目录查询 </li>
	// <li>FileExtension - String - 是否必填：否 - 文件后缀查询 </li>
	// <li>Mode - String - 是否必填：否 - 规则类型 0 MD5,1自定义</li>
	// <li>Md5 - String - 是否必填：否 - MD5查询 </li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeMalwareWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CanNotSeparateInfo struct {

	// 主机名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 隔离失败原因 1:agent离线

	Reason *uint64 `json:"Reason,omitempty" name:"Reason"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DeleteBaselineWeakPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBaselineWeakPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBaselineWeakPasswordResponse) FromJsonString(s string) error {
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

type AttackSourceNode struct {

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 事件ID，为空的时候表示没有对应事件

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// BRUTEFORCE:密码破解、MALWARE:木马、BASH:高危命令、RISK_DNS:恶意请求、LOGIN:异地登录、HOST:主机节点, TIME_ORDER：通用节点

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 节点ip 当节点为HOST时

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 等级  0：提示，1：低危,  2：中危,  3：高危,  4：严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 通用节点的描述

	NodeDesc *string `json:"NodeDesc,omitempty" name:"NodeDesc"`
	// 节点详情

	NodeDetail *string `json:"NodeDetail,omitempty" name:"NodeDetail"`
	// 节点ID

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 时间线编号，同一个编号的节点属于同一个时间线

	TimeLineNum *uint64 `json:"TimeLineNum,omitempty" name:"TimeLineNum"`
}

type IgnoreBaselineRule struct {

	// 影响主机数

	EffectHostCount *uint64 `json:"EffectHostCount,omitempty" name:"EffectHostCount"`
	// 修复建议

	Fix *string `json:"Fix,omitempty" name:"Fix"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 基线检测项id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 基线检测项名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type MachineClearHistory struct {

	// 客户端最后离线时间

	AgentLastOfflineTime *string `json:"AgentLastOfflineTime,omitempty" name:"AgentLastOfflineTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// ID值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
}

type RuleInfo struct {

	// 全文索引的相关配置

	FullText *FullTextInfo `json:"FullText,omitempty" name:"FullText"`
	// 键值索引的相关配置

	KeyValue *KeyValueInfo `json:"KeyValue,omitempty" name:"KeyValue"`
	// 元字段索引配置

	Tag *KeyValueInfo `json:"Tag,omitempty" name:"Tag"`
}

type DeleteExpertServiceRequest struct {
	*tchttp.BaseRequest

	// 需要删除的数据id 最大100条

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteExpertServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExpertServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHotVulTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞信息

		List []*VulStoreListInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHotVulTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHotVulTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseGeneralResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自动升级开关状态,默认 false,  true 开启, false 关闭

		AutoOpenStatus *bool `json:"AutoOpenStatus,omitempty" name:"AutoOpenStatus"`
		// 自动加购订单是否自动续费 ,true 开启, false 关闭

		AutoRepurchaseRenewSwitch *bool `json:"AutoRepurchaseRenewSwitch,omitempty" name:"AutoRepurchaseRenewSwitch"`
		// 自动加购开关, true 开启, false 关闭

		AutoRepurchaseSwitch *bool `json:"AutoRepurchaseSwitch,omitempty" name:"AutoRepurchaseSwitch"`
		// 可用旗舰版授权数

		AvailableFlagshipVersionLicenseCnt *uint64 `json:"AvailableFlagshipVersionLicenseCnt,omitempty" name:"AvailableFlagshipVersionLicenseCnt"`
		// 可用惠普版授权数

		AvailableLHLicenseCnt *uint64 `json:"AvailableLHLicenseCnt,omitempty" name:"AvailableLHLicenseCnt"`
		// 可用授权数

		AvailableLicenseCnt *uint64 `json:"AvailableLicenseCnt,omitempty" name:"AvailableLicenseCnt"`
		// 可用专业版授权数(包含后付费).

		AvailableProVersionLicenseCnt *uint64 `json:"AvailableProVersionLicenseCnt,omitempty" name:"AvailableProVersionLicenseCnt"`
		// 普惠版总授权数(有效订单的授权数)

		CwpVersionLicenseCnt *uint64 `json:"CwpVersionLicenseCnt,omitempty" name:"CwpVersionLicenseCnt"`
		// 已到期授权数(不包含已删除的记录)

		ExpireLicenseCnt *uint64 `json:"ExpireLicenseCnt,omitempty" name:"ExpireLicenseCnt"`
		// 旗舰版总授权数(有效订单)

		FlagshipVersionLicenseCnt *uint64 `json:"FlagshipVersionLicenseCnt,omitempty" name:"FlagshipVersionLicenseCnt"`
		// 历史是否开通过自动升级开关

		IsOpenStatusHistory *bool `json:"IsOpenStatusHistory,omitempty" name:"IsOpenStatusHistory"`
		// 总授权数 (包含隔离,过期等不可用状态)

		LicenseCnt *uint64 `json:"LicenseCnt,omitempty" name:"LicenseCnt"`
		// 即将到期授权数 (15天内到期的)

		NearExpiryLicenseCnt *uint64 `json:"NearExpiryLicenseCnt,omitempty" name:"NearExpiryLicenseCnt"`
		// 未到期授权数

		NotExpiredLicenseCnt *uint64 `json:"NotExpiredLicenseCnt,omitempty" name:"NotExpiredLicenseCnt"`
		// 专业版总授权数(有效订单)

		ProVersionLicenseCnt *uint64 `json:"ProVersionLicenseCnt,omitempty" name:"ProVersionLicenseCnt"`
		// PROVERSION_POSTPAY 专业版-后付费, PROVERSION_PREPAY 专业版-预付费, FLAGSHIP_PREPAY 旗舰版-预付费

		ProtectType *string `json:"ProtectType,omitempty" name:"ProtectType"`
		// 已使用授权数

		UsedLicenseCnt *uint64 `json:"UsedLicenseCnt,omitempty" name:"UsedLicenseCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseGeneralResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseGeneralResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineStrategyDetailRequest struct {
	*tchttp.BaseRequest

	// 用户基线策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DescribeBaselineStrategyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineStrategyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRansomDefenseStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已备份主机数

		BackupMachineCount *uint64 `json:"BackupMachineCount,omitempty" name:"BackupMachineCount"`
		// 账户状态：0未欠费，1已欠费

		BalanceStatus *uint64 `json:"BalanceStatus,omitempty" name:"BalanceStatus"`
		// 已开启防御机器数

		MachineCount *uint64 `json:"MachineCount,omitempty" name:"MachineCount"`
		// 机器总数

		MachineTotal *uint64 `json:"MachineTotal,omitempty" name:"MachineTotal"`
		// 进行中的恢复备份任务数

		ProgressingRollBackTaskCount *uint64 `json:"ProgressingRollBackTaskCount,omitempty" name:"ProgressingRollBackTaskCount"`
		// 进行中的创建快照任务数

		ProgressingSnapshotTaskCount *uint64 `json:"ProgressingSnapshotTaskCount,omitempty" name:"ProgressingSnapshotTaskCount"`
		// 恢复备份任务数量

		RollBackTaskCount *uint64 `json:"RollBackTaskCount,omitempty" name:"RollBackTaskCount"`
		// 快照总容量

		SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`
		// 已开启防御策略数量

		StrategyCount *uint64 `json:"StrategyCount,omitempty" name:"StrategyCount"`
		// 策略总数

		StrategyTotal *uint64 `json:"StrategyTotal,omitempty" name:"StrategyTotal"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRansomDefenseStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRansomDefenseStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专业周报列表数组。

		WeeklyReports []*WeeklyReport `json:"WeeklyReports,omitempty" name:"WeeklyReports"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeklyReportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogKafkaStateRequest struct {
	*tchttp.BaseRequest

	// 接入地址

	AccessAddr *string `json:"AccessAddr,omitempty" name:"AccessAddr"`
	// 接入方式，1公网域名接入，2支撑环境接入

	AccessType *uint64 `json:"AccessType,omitempty" name:"AccessType"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// 峰值带宽

	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`
	// 投递状态，1：健康，2：告警，3：异常

	DeliverStatus *uint64 `json:"DeliverStatus,omitempty" name:"DeliverStatus"`
	// 磁盘容量

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// kafka版本

	InsVersion *string `json:"InsVersion,omitempty" name:"InsVersion"`
	// 实例名称 如 云镜测试环境

	KafkaEnvName *string `json:"KafkaEnvName,omitempty" name:"KafkaEnvName"`
	// 实例ID

	KafkaId *string `json:"KafkaId,omitempty" name:"KafkaId"`
	// 所在子网

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 所属网络

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 区域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *ModifyLogKafkaStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogKafkaStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineRule struct {

	// [0:所有专业版旗舰版|1:hostID|2:ip]

	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
	// 规则分类

	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 主机数

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 主机Id集合

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 主机IP

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 适配项ID列表

	Items []*Item `json:"Items,omitempty" name:"Items"`
	// 规则描述

	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`
	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称,长度不超过128英文字符

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则类型 [0:系统|1:自定义]

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
}

type ScreenNameValue struct {

	// 统计类型 不同接口对应不同的内容

	Name *string `json:"Name,omitempty" name:"Name"`
	// 统计数量

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DescribeBaselineItemDetectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineItemDetect `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineItemDetectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskBatchStatusRequest struct {
	*tchttp.BaseRequest

	// 操作事件类型，文件查杀：MALWARE，异常登录：HOST_LOGIN，密码破解：BRUTE_ATTACK，恶意请求：MALICIOUS_REQUEST，高危命令：BASH_EVENT，本地提权：PRIVILEGE_EVENT，反弹shell：REVERSE_SHELL

	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
}

func (r *DescribeRiskBatchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskBatchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetRecentMachineInfoRequest struct {
	*tchttp.BaseRequest

	// 开始时间。

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 结束时间。

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *ExportAssetRecentMachineInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetRecentMachineInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopInfo struct {

	// top统计计数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// top统计数据，如ip、漏洞名等

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeMachineGeneralRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMachineGeneralRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineGeneralRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOrderAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOrderAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOrderAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRansomDefenseStrategyStatusRequest struct {
	*tchttp.BaseRequest

	// 策略ID列表

	IdList []*uint64 `json:"IdList,omitempty" name:"IdList"`
	// 是否对全部策略生效: 0否，1是

	IsAll *uint64 `json:"IsAll,omitempty" name:"IsAll"`
	// 0关闭，1开启，9删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyRansomDefenseStrategyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRansomDefenseStrategyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveMachineRequest struct {
	*tchttp.BaseRequest

	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *RemoveMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Account struct {

	// 帐号创建时间。

	AccountCreateTime *string `json:"AccountCreateTime,omitempty" name:"AccountCreateTime"`
	// 帐号所属组。

	Groups *string `json:"Groups,omitempty" name:"Groups"`
	// 唯一ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 帐号最后登录时间。

	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
	// 主机内网IP。

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 主机名称。

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 帐号类型。
	// <li>ORDINARY：普通帐号</li>
	// <li>SUPPER：超级管理员帐号</li>

	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`
	// 帐号名。

	Username *string `json:"Username,omitempty" name:"Username"`
	// 云镜客户端唯一Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeVulOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScreenInvasion struct {

	// 事件数据的json, 每种事件不同，
	// 【文件查杀】病毒名 VirusName、文件名 FileName、文件路径 FilePath、文件大小 FileSize、文件MD5 MD5、首次发现时间 CreateTime、最近检测时间LatestScanTime、危害描述 HarmDescribe、修复建议SuggestScheme
	// 【异常登录】来源IP SrcIp、来源地 Location、登录用户名 UserName、登录时间 LoginTime
	// 【密码破解】来源IP SrcIp、来源地 City,Country、协议 Protocol、登录用户名UserName 、端口 Port、尝试次数 Count、首次攻击时间 CreateTime、最近攻击时间 ModifyTime
	// 【恶意请求】恶意请求域名 Url、进程ProcessName 、MD5 ProcessMd5、PID Pid、请求次数 AccessCount、最近请求时间 MergeTime、危害描述 HarmDescribe、修复建议SuggestScheme
	// 【高危命令】命中规则名 RuleName、规则类别 RuleCategory、命令内容 BashCmd、数据来源 DetectBy、登录用户 User、PID Pid、发生时间 CreateTime 、危害描述 HarmDescribe、修复建议SuggestScheme
	// 【本地提权】提权用户 UserName、父进程 ParentProcName 、父进程所属用户 ParentProcGroup、发现时间 CreateTime、危害描述 HarmDescribe、修复建议SuggestScheme
	// 【反弹shell】连接进程 ProcessName、执行命令CmdLine、父进程ParentProcName、目标主机DstIp、目标端口DstPort、发现时间 CreateTime、危害描述 HarmDescribe、修复建议SuggestScheme

	Content *string `json:"Content,omitempty" name:"Content"`
	// 入侵时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 事件类型：0：文件查杀，1：异常登录， 2：密码破解，3：恶意请求，4：高危命令，5：本地提权， 6：反弹shell

	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`
	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 事件统一等级 0：提示，1：低危,  2：中危,  3：高危,  4：严重

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 等级中文展示

	LevelZh *string `json:"LevelZh,omitempty" name:"LevelZh"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type VulDefenceOverview struct {

	// 每日攻击趋势

	AttackCounts []*int64 `json:"AttackCounts,omitempty" name:"AttackCounts"`
	// 日期

	Date []*string `json:"Date,omitempty" name:"Date"`
	// 每日防御趋势

	DefendCounts []*int64 `json:"DefendCounts,omitempty" name:"DefendCounts"`
	// 已开启防御主机数

	DefendHostCount *int64 `json:"DefendHostCount,omitempty" name:"DefendHostCount"`
	// 防御开关：0 关闭 1 开启

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 插件异常数

	ExceptionCount *int64 `json:"ExceptionCount,omitempty" name:"ExceptionCount"`
}

type DeleteBashPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBashPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBashPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRiskDnsPolicyStatusRequest struct {
	*tchttp.BaseRequest

	// 是否生效[0:生效,1:不生效]

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 策略ID

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *ModifyRiskDnsPolicyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskDnsPolicyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebLocationPathListRequest struct {
	*tchttp.BaseRequest

	// Web站点Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 返回列表数量，最多100，默认10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 服务器Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebLocationPathListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationPathListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackVulTypeListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAttackVulTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackVulTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineScanScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检测进度(百分比)

		Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineScanScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineScanScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBaselinePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBaselinePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBaselinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebFrameJarListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Jars []*AssetJarBaseInfo `json:"Jars,omitempty" name:"Jars"`
		// 分区总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebFrameJarListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebFrameJarListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
