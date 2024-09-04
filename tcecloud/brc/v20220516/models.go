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

package v20220516

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeBackupGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 备份列表详情。

		BackupGroupSet []*BackupGroup `json:"BackupGroupSet,omitempty" name:"BackupGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsBackupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>backup-id - Array of String - 是否必填：否 -（过滤条件）按照备份的ID过滤。快照ID形如：`backup-11112222`。<br><li>backup-name - Array of String - 是否必填：否 -（过滤条件）按照备份名称过滤。<br><li>backup-state - Array of String - 是否必填：否 -（过滤条件）按照备份状态过滤。 (NORMAL：正常 | CREATING：创建中 | ROLLBACKING：回滚中。)<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按创建快照的云盘类型过滤。 (SYSTEM_DISK：代表系统盘 | DATA_DISK：代表数据盘。)<br><li>platform-project-id  - Array of String - 是否必填：否 -（过滤条件）按备份所属项目ID过滤。<br><li>file-system-id  - Array of String - 是否必填：否 -（过滤条件）按照创建备份的文件系统ID过滤。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照可用区过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 备份列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据备份的创建时间排序<br>默认按创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeCfsBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstancesWithBackupGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当通过本接口来创建实例时会返回该参数，表示一个或多个实例`ID`。返回实例`ID`列表并不代表实例创建成功，可根据 [DescribeInstancesStatus](DescribeInstancesStatus) 接口查询返回的InstancesSet中对应实例的`ID`的状态来判断创建是否完成；如果实例状态由“准备中”变为“正在运行”，则为创建成功。

		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
		// 本次创建实例使用的镜像ID。

		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunInstancesWithBackupGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstancesWithBackupGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvancedRetentionPolicy struct {

	// 保留Days天内的每天最新的一个备份

	Days *uint64 `json:"Days,omitempty" name:"Days"`
	// 保留Weeks周内的每周最新的一个备份

	Weeks *uint64 `json:"Weeks,omitempty" name:"Weeks"`
	// 保留Months月内的每月最新的一个备份

	Months *uint64 `json:"Months,omitempty" name:"Months"`
	// 保留Years年内的每年最新的一个备份

	Years *uint64 `json:"Years,omitempty" name:"Years"`
}

type Filter struct {

	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeBackupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份的数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云硬盘备份列表详情。

		BackupSet []*Backup `json:"BackupSet,omitempty" name:"BackupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsDeniedActionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份的操作掩码。

		BackupDeniedActionSet []*BackupDeniedAction `json:"BackupDeniedActionSet,omitempty" name:"BackupDeniedActionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupsDeniedActionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationLog struct {

	// 操日志任务ID。

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 定期备份策略ID。

	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	// 操作日志任务结果。<br><li>SUCCESS: 成功<br><li>FAILED: 失败

	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`
	// 子任务执行情况统计。

	SubTaskStatistic *SubTaskStatistic `json:"SubTaskStatistic,omitempty" name:"SubTaskStatistic"`
	// 任务流水ID。

	LogTaskId *string `json:"LogTaskId,omitempty" name:"LogTaskId"`
	// 备份组ID。

	BackupGroupId *string `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
	// 任务开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 用户AppId。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 操作者。

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
	// 任务名称。

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 当前任务的父任务ID。如果无父任务，则默认为0。

	ParentTaskId *uint64 `json:"ParentTaskId,omitempty" name:"ParentTaskId"`
	// 云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 任务结果描述。

	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
	// 文件系统ID。

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type Policy struct {

	// 选定周一到周日中需要创建快照的日期，取值范围：[0, 6]。0表示周日触发，1表示周一触发，依次类推。

	DayOfWeek []*uint64 `json:"DayOfWeek,omitempty" name:"DayOfWeek"`
	// 指定定期快照策略的触发时间。单位为小时，取值范围：[0, 23]。00:00 ~ 23:00 共 24 个时间点可选，1表示 01:00，依此类推。

	Hour []*int64 `json:"Hour,omitempty" name:"Hour"`
	// 指定每月从月初到月底需要触发定期备份的日期,取值范围：[1, 31]，1-31分别表示每月的具体日期，比如5表示每月的5号。注：若设置29、30、31等部分月份不存在的日期，则对应不存在日期的月份会跳过不打定期备份。

	DayOfMonth []*uint64 `json:"DayOfMonth,omitempty" name:"DayOfMonth"`
	// 指定创建定期快照的间隔天数，取值范围：[1, 365]，例如设置为5，则间隔5天即触发定期快照创建。注：当选择按天备份时，理论上第一次备份的时间为备份策略创建当天。如果当天备份策略创建的时间已经晚于设置的备份时间，那么将会等到第二个备份周期再进行第一次备份。

	IntervalDays *uint64 `json:"IntervalDays,omitempty" name:"IntervalDays"`
}

type CreateBackupGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份组ID。

		BackupGroupId *string `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupResourceOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份资源概览

		BackupOverview *BackupOverview `json:"BackupOverview,omitempty" name:"BackupOverview"`
		// 资源概览详情列表

		OverviewDetailSet []*OverviewDetail `json:"OverviewDetailSet,omitempty" name:"OverviewDetailSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupResourceOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupResourceOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileSystemBackup struct {

	// 备份名称。

	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
	// 是否为远程复制的备份。

	CopyFromRemote *bool `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`
	// 是否永久保留。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 备份的状态。取值范围：<br><li>NORMAL：正常<br><li>CREATING：创建中<br><li>ROLLBACKING：回滚中<br><li>COPYING_FROM_REMOTE：跨地域复制拷贝中。

	BackupState *string `json:"BackupState,omitempty" name:"BackupState"`
	// 备份的到期时间。

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 备份创建的进度。

	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
	// 备份被共享的次数。

	ShareReference *uint64 `json:"ShareReference,omitempty" name:"ShareReference"`
	// 备份类型，目前该项取值可以为PRIVATE_BACKUP或者SHARED_BACKUP。

	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
	// 创建备份的云硬盘大小，单位GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 创建备份的云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 备份所属项目ID。

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// 备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
	// 备份当前正在远程复制的目标地域列表。

	CopyingToRegions []*string `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`
	// 备份是否为加密备份。

	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`
	// 备份的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建此备份的云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 备份所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 备份绑定的标签列表。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 用户AppId。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 云盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 创建备份时刻，文件系统各属性的详情。

	FileSystemDetails *string `json:"FileSystemDetails,omitempty" name:"FileSystemDetails"`
	// 全量、增量备份信息；FULL表示全量备份，INC表示增量备份。

	BackupClass *string `json:"BackupClass,omitempty" name:"BackupClass"`
	// 创建备份的文件系统ID。

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type InternetAccessible struct {

	// 网络计费类型。取值范围：<br><li>BANDWIDTH_PREPAID：预付费按带宽结算<br><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费<br><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费<br><li>BANDWIDTH_PACKAGE：带宽包用户<br>默认取值：TRAFFIC_POSTPAID_BY_HOUR。

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致。

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>公网带宽大于0时必须设置为True,默认开通公网IP；当公网带宽为0，则不允许分配公网IP。

	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
}

type CreateBackupGroupRequest struct {
	*tchttp.BaseRequest

	// 需要创建备份组的云硬盘ID列表。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 备份组的名称

	BackupGroupName *string `json:"BackupGroupName,omitempty" name:"BackupGroupName"`
	// 指定备份组到期时间，如果未传入该参数，默认为永久保留。

	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
	// 创建备份的带宽上限，范围：[0, 100]

	CreateSpeed *uint64 `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
}

func (r *CreateBackupGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCfsFileSystemWithBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建的文件系统ID。

		FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCfsFileSystemWithBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsFileSystemWithBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoBackupPolicyAttributeRequest struct {
	*tchttp.BaseRequest

	// 定期备份策略ID。

	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	// 定期备份的执行策略。

	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`
	// 定期备份策略的名称。

	AutoBackupPolicyName *string `json:"AutoBackupPolicyName,omitempty" name:"AutoBackupPolicyName"`
	// 是否激活定期备份策略。

	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`
	// 通过该定期备份策略创建的备份是否永久保留。false表示非永久保留，true表示永久保留，默认为false。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 使用该定期备份策略创建出来的备份保留天数。

	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
	// 该定期快照策略创建的快照可以保留的月数，该参数不可与IsPermanent/RetentionDays参数冲突。

	RetentionMonths *uint64 `json:"RetentionMonths,omitempty" name:"RetentionMonths"`
	// 通过该定期快照策略最多保留的快照个数，超过该个数限制后自动删除最先创建的快照，该参数不可与IsPermanent参数冲突。

	RetentionAmount *uint64 `json:"RetentionAmount,omitempty" name:"RetentionAmount"`
	// 定期快照高级保留策略，该参数不可与IsPermanent参数冲突。

	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitempty" name:"AdvancedRetentionPolicy"`
	// 创建备份的带宽上限，范围：[0, 100]

	CreateSpeed *uint64 `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	// 每隔几个备份做一个全量备份，0表示全部做全量备份。

	FullBackupInterval *int64 `json:"FullBackupInterval,omitempty" name:"FullBackupInterval"`
}

func (r *ModifyAutoBackupPolicyAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoBackupPolicyAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupDisk struct {

	// 云硬盘最新一次备份的时间。

	LatestBackupTime *string `json:"LatestBackupTime,omitempty" name:"LatestBackupTime"`
	// 云硬盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 云硬盘绑定的定期备份策略列表。

	AutoBackupPolicyIdSet []*string `json:"AutoBackupPolicyIdSet,omitempty" name:"AutoBackupPolicyIdSet"`
	// 云硬盘所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 付费模式。取值范围：<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：后付费，即按量计费。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 用户AppId。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 云盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：SSD表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云硬盘的备份ID列表。

	BackupIdSet []*string `json:"BackupIdSet,omitempty" name:"BackupIdSet"`
	// 云硬盘创建时间。

	DiskCreateTime *string `json:"DiskCreateTime,omitempty" name:"DiskCreateTime"`
	// 云盘状态。取值范围：<br><li>UNATTACHED：未挂载<br><li>ATTACHING：挂载中<br><li>ATTACHED：已挂载<br><li>DETACHING：解挂中<br><li>EXPANDING：扩容中<br><li>ROLLBACKING：回滚中<br><li>DELETED：已删除。

	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`
	// 云硬盘大小，单位GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 最新修改时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 云硬盘当前存在的最早备份的时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 云盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

type DeniedAction struct {

	// 不能操作接口名。

	Action *string `json:"Action,omitempty" name:"Action"`
	// 接口不能操作的原因。

	Message *string `json:"Message,omitempty" name:"Message"`
	// 接口不能操作对应提示的错误码。

	Code *string `json:"Code,omitempty" name:"Code"`
}

type OverviewDetail struct {

	// 备份资源概览

	BackupOverview *BackupOverview `json:"BackupOverview,omitempty" name:"BackupOverview"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeBackupGroupsDeniedActionsRequest struct {
	*tchttp.BaseRequest

	// 备份组ID列表。

	BackupGroupIds []*string `json:"BackupGroupIds,omitempty" name:"BackupGroupIds"`
}

func (r *DescribeBackupGroupsDeniedActionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupGroupsDeniedActionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyDisk struct {

	// 备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
	// 云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

type DeleteAutoBackupPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAutoBackupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoBackupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAutoBackupPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 定期备份策略ID。

		AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
		// 首次开始备份的时间。

		NextTriggerTime *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAutoBackupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupOperationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 操作日志列表。

		OperationLogSet []*OperationLog `json:"OperationLogSet,omitempty" name:"OperationLogSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupOperationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoBackupPolicy struct {

	// 备份策略绑定的云硬盘列表。

	DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`
	// 定期备份策略是否激活。

	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`
	// 使用该定期备份策略创建出来的备份是否永久保留。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 定期备份下次触发的时间。

	NextTriggerTime *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`
	// 定期备份策略的状态。取值范围：<br><li>NORMAL：正常<br><li>ISOLATED：已隔离。

	AutoBackupPolicyState *string `json:"AutoBackupPolicyState,omitempty" name:"AutoBackupPolicyState"`
	// 备份策略的名称。

	AutoBackupPolicyName *string `json:"AutoBackupPolicyName,omitempty" name:"AutoBackupPolicyName"`
	// 定期备份的执行策略。

	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`
	// 备份策略ID。

	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	// 备份策略的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 使用该定期备份策略创建出来的备份保留天数。

	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
	// 用户AppId。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 定期备份策略绑定的实例ID列表。

	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
	// 备份策略绑定的文件系统ID列表。

	FileSystemIdSet []*string `json:"FileSystemIdSet,omitempty" name:"FileSystemIdSet"`
	// 该定期快照创建的快照最大保留月数

	RetentionMonths *uint64 `json:"RetentionMonths,omitempty" name:"RetentionMonths"`
	// 该定期快照创建的快照最大保留数量

	RetentionAmount *uint64 `json:"RetentionAmount,omitempty" name:"RetentionAmount"`
	// 高级保留策略

	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitempty" name:"AdvancedRetentionPolicy"`
	// 创建备份的带宽上限，范围：[0, 100]

	CreateSpeed *uint64 `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
}

type BackupOverview struct {

	// 备份云硬盘个数

	BackupDiskCount *uint64 `json:"BackupDiskCount,omitempty" name:"BackupDiskCount"`
	// 备份云硬盘容量(MiB)

	BackupDiskSizeMb *uint64 `json:"BackupDiskSizeMb,omitempty" name:"BackupDiskSizeMb"`
	// 备份云主机个数

	BackupInstanceCount *uint64 `json:"BackupInstanceCount,omitempty" name:"BackupInstanceCount"`
	// 备份云主机容量(MiB)

	BackupInstanceSizeMb *uint64 `json:"BackupInstanceSizeMb,omitempty" name:"BackupInstanceSizeMb"`
	// 备份文件系统数量。

	BackupFileSystemCount *uint64 `json:"BackupFileSystemCount,omitempty" name:"BackupFileSystemCount"`
	// 备份文件系统容量(MiB)。

	BackupFileSystemSizeMb *uint64 `json:"BackupFileSystemSizeMb,omitempty" name:"BackupFileSystemSizeMb"`
}

type ApplyCfsBackupRequest struct {
	*tchttp.BaseRequest

	// 回滚的备份ID，必须是文件系统创建的备份。。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
	// 回滚的原文件系统ID。

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *ApplyCfsBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyCfsBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstancesWithBackupGroupRequest struct {
	*tchttp.BaseRequest

	// 实例计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDHPAID：独享母机付费（基于专用宿主机创建，宿主机部分的资源不收费），该付费模式下必须填写placement.hostid参数<br>默认值：POSTPAID_BY_HOUR。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，专用宿主机（对于独享母机付费模式的子机创建）等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 实例机型。不同实例机型指定了不同的资源规格。<br><li>对于付费模式为PREPAID或POSTPAID_BY_HOUR的子机创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](DescribeInstanceTypeConfigs)来获得最新的规格表或参见[实例类型](/tcloud/Compute/CVM/292128/484318/specification)描述。若不指定该参数，则默认机型为S1.SMALL1。<br><li>对于付费模式为CDHPAID的子机创建，该参数以"CDH_"为前缀，根据cpu和内存配置生成，具体形式为：CDH_XCXG，例如对于创建cpu为1核，内存为1G大小的专用宿主机的子机，该参数应该为CDH_1C1G。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例系统盘配置信息。系统盘的大小必须大于等于备份组系统盘备份的大小。

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，支持购买时指定多个数据盘。数据盘大小必须大于等于备份组数据盘备份的大小。

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，默认使用vpc网络。若在此参数中指定了私有网络ip，那么InstanceCount参数可以填1或2。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 购买实例数量。取值范围：[1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量。

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 实例显示名称。如果不指定则默认显示`FROM cbackup-xxxxxxxx`. 最多只支持60个字符，点后面的名字都会过滤掉。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 实例所属安全组。若不指定该参数，则绑定默认安全组。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，默认关闭云监控和云安全服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。<br><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。<br><li>其他类型（Linux 等）实例：字符长度为[2, 31]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成，不支持全数字;不支持.-(点和短横线放在一起)。

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。

	UserData *string `json:"UserData,omitempty" name:"UserData"`
	// 置放群组id，仅支持指定一个。

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
	// 指定要回滚的备份组ID。

	BackupGroupId *string `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
	// 是否执行真正的创建，默认为false。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *RunInstancesWithBackupGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstancesWithBackupGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeBackupDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云硬盘备份详情。

		BackupDiskSet []*BackupDisk `json:"BackupDiskSet,omitempty" name:"BackupDiskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyBackupGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyBackupGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyBackupGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoBackupPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的定期备份策略数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 定期备份策略列表。

		AutoBackupPolicySet []*AutoBackupPolicy `json:"AutoBackupPolicySet,omitempty" name:"AutoBackupPolicySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoBackupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoBackupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindAutoBackupPolicyRequest struct {
	*tchttp.BaseRequest

	// 要解绑定期备份策略的云硬盘ID列表。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 要解绑的定期备份策略ID。

	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	// 要解绑定期备份策略的实例ID列表。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 要解绑定期备份策略的文件系统列表。

	FileSystemIds []*string `json:"FileSystemIds,omitempty" name:"FileSystemIds"`
}

func (r *UnbindAutoBackupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindAutoBackupPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份ID

		BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoBackupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 要删除的定期备份策略列表。

	AutoBackupPolicyIds []*string `json:"AutoBackupPolicyIds,omitempty" name:"AutoBackupPolicyIds"`
}

func (r *DeleteAutoBackupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoBackupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupCfsFileSystemsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>file-system-id - Array of String - 是否必填：否 -（过滤条件）按文件系统ID过滤。<br><li>auto-backup-policy-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘绑定的定期备份策略ID过滤。<br><li>auto-backup-policy-name - Array of String - 是否必填：否 -（过滤条件）按照云硬盘绑定的定期备份策略名称过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出云硬盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 备份云硬盘列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云硬盘首次备份的时间排序<br>默认按例首次备份的时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeBackupCfsFileSystemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupCfsFileSystemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupResourceOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBackupResourceOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupResourceOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyBackupRequest struct {
	*tchttp.BaseRequest

	// 回滚的备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
	// 回滚的云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 回滚备份前是否执行自动关机，如果回滚的盘挂载在实例上且实例处于运行状态，可传入该参数。

	AutoStopInstance *bool `json:"AutoStopInstance,omitempty" name:"AutoStopInstance"`
	// 回滚备份完成后是否执行自动开机。

	AutoStartInstance *bool `json:"AutoStartInstance,omitempty" name:"AutoStartInstance"`
}

func (r *ApplyBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupCfsFileSystemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 文件系统备份详情。

		BackupFileSystemSet []*BackupFileSystem `json:"BackupFileSystemSet,omitempty" name:"BackupFileSystemSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupCfsFileSystemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupCfsFileSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBackupAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBackupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupDeniedAction struct {

	// 备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
	// 具体的备份操作掩码列表。

	SnapshotDeniedAction []*DeniedAction `json:"SnapshotDeniedAction,omitempty" name:"SnapshotDeniedAction"`
}

type Placement struct {

	// 所属的可用区ID。该参数也可以通过调用DescribeZones的返回值中的Zone字段来获取。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](/document/api/378/4400) 的返回值中的 projectId 字段来获取。不填为默认项目。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 实例所属的独享集群ID。作为入参时，表示对指定的CdcId独享集群的资源进行操作，可为空。 作为出参时，表示资源所属的独享集群的ID，可为空。

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
	// 围笼Id。作为入参时，表示对指定的CageId的资源进行操作，可为空。 作为出参时，表示资源所属围笼ID，可为空。

	CageId *string `json:"CageId,omitempty" name:"CageId"`
	// 独享集群名字。作为入参时，忽略。作为出参时，表示云硬盘所属的独享集群名，可为空。

	CdcName *string `json:"CdcName,omitempty" name:"CdcName"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type SystemDisk struct {

	// 系统盘类型。取值范围：<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><li>CLOUD_PREMIUM：高性能云盘

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘大小，单位：GB。默认值为 50

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 系统盘指定的资源池。

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
}

type ApplyCfsBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyCfsBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyCfsBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyBackupGroupToSnapshotGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 转换后的快照组Id。

		SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" name:"SnapshotGroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyBackupGroupToSnapshotGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBackupGroupToSnapshotGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupAttributeRequest struct {
	*tchttp.BaseRequest

	// 备份的名称。

	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
	// 是否为永久保留的备份。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 备份到期时间。

	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
	// 修改的备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *ModifyBackupAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBackupAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyBackupGroupRequest struct {
	*tchttp.BaseRequest

	// 回滚的备份组ID。

	BackupGroupId *string `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
	// 回滚的备份ID、云硬盘ID列表。

	ApplyDisks []*ApplyDisk `json:"ApplyDisks,omitempty" name:"ApplyDisks"`
	// 回滚备份前是否执行自动关机，如果回滚的盘挂载在实例上且实例处于运行状态，可传入该参数。

	AutoStopInstance *bool `json:"AutoStopInstance,omitempty" name:"AutoStopInstance"`
	// 回滚备份完成后是否执行自动开机。

	AutoStartInstance *bool `json:"AutoStartInstance,omitempty" name:"AutoStartInstance"`
}

func (r *ApplyBackupGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyBackupGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupInstancesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>instance-id - Array of String - 是否必填：否 -（过滤条件）按实例ID过滤。<br><li>auto-backup-policy-id - Array of String - 是否必填：否 -（过滤条件）按照实例绑定的定期备份策略过滤。<br><li>auto-backup-policy-name - Array of String - 是否必填：否 -（过滤条件）按照云硬盘绑定的定期备份策略名称过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出实例列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 备份实例列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据实例首次备份的时间排序<br>默认按例首次备份的时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeBackupInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunSecurityServiceEnabled struct {

	// 是否开启[云安全](/tcloud/Security/CWP)服务。取值范围：<br><li>TRUE：表示开启云安全服务<br><li>FALSE：表示不开启云安全服务<br><br>默认取值：TRUE。

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type DescribeCfsBackupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份的数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 文件系统备份列表详情

		BackupSet []*FileSystemBackup `json:"BackupSet,omitempty" name:"BackupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfsBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBackupGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyBackupToSnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份转快照的快照ID。

		SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyBackupToSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBackupToSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>backup-id - Array of String - 是否必填：否 -（过滤条件）按照备份的ID过滤。快照ID形如：`backup-11112222`。<br><li>backup-name - Array of String - 是否必填：否 -（过滤条件）按照备份名称过滤。<br><li>backup-state - Array of String - 是否必填：否 -（过滤条件）按照备份状态过滤。 (NORMAL：正常 | CREATING：创建中 | ROLLBACKING：回滚中。)<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按创建快照的云盘类型过滤。 (SYSTEM_DISK：代表系统盘 | DATA_DISK：代表数据盘。)<br><li>platform-project-id  - Array of String - 是否必填：否 -（过滤条件）按备份所属项目ID过滤。<br><li>disk-id  - Array of String - 是否必填：否 -（过滤条件）按照创建备份的云硬盘ID过滤。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照可用区过滤。<br><li>backup-group-id - Array of String - 是否必填：否 -（过滤条件）按照备份绑定的备份组ID过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 备份列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据备份的创建时间排序<br>默认按创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupGroup struct {

	// 备份组是否包含系统盘备份。

	ContainRootBackup *bool `json:"ContainRootBackup,omitempty" name:"ContainRootBackup"`
	// 备份组ID。

	BackupGroupId *string `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
	// 备份组类型。NORMAL: 普通备份组，非一致性备份。

	BackupGroupType *string `json:"BackupGroupType,omitempty" name:"BackupGroupType"`
	// 备份组创建进度。

	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
	// 备份组包含的备份ID列表。

	BackupIdSet []*string `json:"BackupIdSet,omitempty" name:"BackupIdSet"`
	// 备份组名称。

	BackupGroupName *string `json:"BackupGroupName,omitempty" name:"BackupGroupName"`
	// 备份组状态。<br><li>NORMAL: 正常<br><li>CREATING:创建中<br><li>DELETED:已删除<br><li>FAILED:创建失败<br><li>DISMISS:已解散<br><li>ROLLBACKING:回滚中

	BackupGroupState *string `json:"BackupGroupState,omitempty" name:"BackupGroupState"`
	// 最新修改时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 备份组创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 用户AppId。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 是否为永久备份组。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 备份组的到期时间。如果为永久备份组，则取值为null。

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 创建备份组的实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 创建备份组时刻实例的详情。

	InstanceDetails *string `json:"InstanceDetails,omitempty" name:"InstanceDetails"`
}

type SubTaskStatistic struct {

	// 子任务成功数量，null表示没有子任务。

	SucceedCount *uint64 `json:"SucceedCount,omitempty" name:"SucceedCount"`
	// 子任务失败数量，null表示没有子任务。

	FailedCount *uint64 `json:"FailedCount,omitempty" name:"FailedCount"`
}

type VirtualPrivateCloud struct {

	// 私有网络ID，形如`vpc-xxxxxxx`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络子网ID，形如`subnet-xxxxxxx`。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：<br><li>TRUE：表示用作公网网关<br><li>FALSE：表示不用作公网网关<br><br>默认取值：FALSE。

	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`
	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 为弹性网卡指定随机生成的 IPv6 地址数量。

	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

type DeleteBackupGroupsRequest struct {
	*tchttp.BaseRequest

	// 备份组ID列表。

	BackupGroupIds []*string `json:"BackupGroupIds,omitempty" name:"BackupGroupIds"`
	// 解散备份组。取值为true，表示仅解散备份组，不删除备份组关联的快照；取值为false，则会同时删除备份组关联的备份。默认取值为false

	OnlyDismiss *bool `json:"OnlyDismiss,omitempty" name:"OnlyDismiss"`
}

func (r *DeleteBackupGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupDisksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按云硬盘ID过滤。<br><li>auto-backup-policy-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘绑定的定期备份策略ID过滤。<br><li>auto-backup-policy-name - Array of String - 是否必填：否 -（过滤条件）按照云硬盘绑定的定期备份策略名称过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出云硬盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 备份云硬盘列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云硬盘首次备份的时间排序<br>默认按例首次备份的时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeBackupDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyBackupCrossRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份跨地域复制的结果，如果请求下发成功，则返回相应地地域的新备份ID，否则返回Error。

		BackupCopyResultSet []*BackupCopyResult `json:"BackupCopyResultSet,omitempty" name:"BackupCopyResultSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyBackupCrossRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBackupCrossRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupGroupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>backup-group-id - Array of String - 是否必填：否 -（过滤条件）按备份组ID过滤 <br><li>backup-group-state - Array of String - 是否必填：否 -（过滤条件）按备份组状态过滤。(NORMAL: 正常 | CREATING:创建中 | ROLLBACKING:回滚中) <br><li>backup-group-name - Array of String - 是否必填：否 -（过滤条件）按备份组名称过滤 <br><li>backup-id - Array of String - 是否必填：否 -（过滤条件）按备份组内的备份ID过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据备份组的创建时间排序<br>默认按创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeBackupGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDisksWithBackupRequest struct {
	*tchttp.BaseRequest

	// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘显示名称。不传则默认为“FROM backup-11112222”。最大长度不能超60个字节。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 创建云硬盘数量，不传则默认为1。单次请求最多可创建的云盘数有限制。

	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`
	// 付费模式，目前只有后付费，取值为POSTPAID_BY_HOUR。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 云硬盘大小，单位为GB。云硬盘大小必须大于等于备份大小，备份大小可通过DescribeBackups接口查询。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 可选参数，默认为False。传入True时，云盘将创建为共享型云盘。

	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`
	// 定期快照策略ID。传入该参数时，云硬盘创建成功后将会自动绑定该定期快照策略。

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 存储资源池组

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 可选参数。使用此参数可给云硬盘购买额外的性能。<br>当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）。

	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
	// 备份ID，本接口会将此备份的数据回滚到新创建的云硬盘上。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *CreateDisksWithBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDisksWithBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsDeniedActionsRequest struct {
	*tchttp.BaseRequest

	// 备份ID列表。

	BackupIds []*string `json:"BackupIds,omitempty" name:"BackupIds"`
}

func (r *DescribeBackupsDeniedActionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupsDeniedActionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// 系统盘类型。取值范围：<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><li>CLOUD_PREMIUM：高性能云盘

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同。默认值为0，表示不购买数据盘。更多限制详见产品文档。

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 数据盘指定的资源池。

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 指定回滚到盘上的备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
}

type BackupCopyResult struct {

	// 跨地域复制的目的地域。

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	// 复制到目的地域的新备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
}

type CreateCfsBackupRequest struct {
	*tchttp.BaseRequest

	// 需要创建备份的文件系统ID。

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 备份名称。

	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
	// 指定备份到期时间，如果未传入该参数，默认为永久保留。

	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
	// 创建备份的带宽上限，范围：[0, 100]

	CreateSpeed *uint64 `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
}

func (r *CreateCfsBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupOperationsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>task-id - Array of Integer - 是否必填：否 -（过滤条件）按任务ID过滤。<br><li>parent-task-id - Array of Integer - 是否必填：否 -（过滤条件）按父任务ID过滤。<br><li>task-name - Array of String - 是否必填：否 -（过滤条件）按任务名称过滤。<br><li>task-state - Array of String - 是否必填：否 -（过滤条件）按任务状态过滤。<br><li>backup-id - Array of String - 是否必填：否 -（过滤条件）按备份D过滤。<br><li>backup-group-id - Array of String - 是否必填：否 -（过滤条件）按备份组ID过滤。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按云盘ID过滤。<br><li>instance-id - Array of String - 是否必填：否 -（过滤条件）按实例ID过滤。<br><li>auto-backup-policy-id - Array of String - 是否必填：否 -（过滤条件）按定期备份策略ID过滤。<br><li>file-system-id - Array of String - 是否必填：否 -（过滤条件）按文件系统ID过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据备份的创建时间排序<br>默认按创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeBackupOperationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupOperationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupsRequest struct {
	*tchttp.BaseRequest

	// 要删除的备份ID列表。

	BackupIds []*string `json:"BackupIds,omitempty" name:"BackupIds"`
}

func (r *DeleteBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupGroupsDeniedActionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询备份组的操作掩码。

		BackupGroupDeniedActionSet []*BackupGroupDeniedAction `json:"BackupGroupDeniedActionSet,omitempty" name:"BackupGroupDeniedActionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupGroupsDeniedActionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupGroupsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAutoBackupPolicyRequest struct {
	*tchttp.BaseRequest

	// 定期备份的执行策略。

	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`
	// 通过该定期备份策略创建的备份是否永久保留。false表示非永久保留，true表示永久保留，默认为false。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 定期备份策略的名称。

	AutoBackupPolicyName *string `json:"AutoBackupPolicyName,omitempty" name:"AutoBackupPolicyName"`
	// 是否激活定期备份策略。

	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`
	// 通过定期备份策略创建出的备份保留时间。

	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
	// 是否实际创建定期备份策略。true表示只需获取首次开始备份的时间，不实际创建定期备份策略，false表示创建，默认为false。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 该定期备份策略创建的备份可以保留的月数，该参数不可与IsPermanent/RetentionDays参数冲突。

	RetentionMonths *uint64 `json:"RetentionMonths,omitempty" name:"RetentionMonths"`
	// 通过该定期备份策略最多保留的备份个数，超过该个数限制后自动删除最先创建的备份，该参数不可与IsPermanent参数冲突。

	RetentionAmount *uint64 `json:"RetentionAmount,omitempty" name:"RetentionAmount"`
	// 定期备份高级保留策略，该参数不可与IsPermanent参数冲突。

	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitempty" name:"AdvancedRetentionPolicy"`
	// 创建备份的带宽上限，范围:[0, 100]

	CreateSpeed *uint64 `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	// 每隔几个备份做一个全量备份，0表示全部做全量备份。

	FullBackupInterval *int64 `json:"FullBackupInterval,omitempty" name:"FullBackupInterval"`
}

func (r *CreateAutoBackupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoBackupPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCfsBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份ID

		BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCfsBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupInstance struct {

	// 实例绑定的定期备份策略列表。

	AutoBackupPolicyIdSet []*string `json:"AutoBackupPolicyIdSet,omitempty" name:"AutoBackupPolicyIdSet"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户AppId。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 实例最新备份时间。

	LatestBackupTime *string `json:"LatestBackupTime,omitempty" name:"LatestBackupTime"`
	// 实例的备份组ID列表。

	BackupGroupIdSet []*string `json:"BackupGroupIdSet,omitempty" name:"BackupGroupIdSet"`
	// 修改时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 实例当前存在的最早备份的时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CopyBackupGroupToSnapshotGroupRequest struct {
	*tchttp.BaseRequest

	// 转换的源备份组Id。

	BackupGroupId *string `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
	// 转换后的目的快照组名称。

	SnapshotGroupName *string `json:"SnapshotGroupName,omitempty" name:"SnapshotGroupName"`
}

func (r *CopyBackupGroupToSnapshotGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBackupGroupToSnapshotGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindAutoBackupPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAutoBackupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAutoBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyBackupToSnapshotRequest struct {
	*tchttp.BaseRequest

	// 需要转快照的备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
	// 备份转快照的快照名称。

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}

func (r *CopyBackupToSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBackupToSnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCfsFileSystemWithBackupRequest struct {
	*tchttp.BaseRequest

	// 可用区名称，例如:ap-beijing-1

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 用户自定义文件系统名称，优先级低于 FSNAME

	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
	// 文件系统协议类型， 值为 NFS、CIFS; 若留空则默认为 NFS协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 文件系统存储类型，值为 SD ；其中 SD 为标准型存储

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 网络类型，值为 VPC，BASIC；其中 VPC 为私有网络，BASIC 为基础网络

	NetInterface *string `json:"NetInterface,omitempty" name:"NetInterface"`
	// 权限组ID。

	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
	// 私有网路（VPC） ID;当网络类型值为 VPC时，与UnVpcId 两者必须填一项

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 系统分配的VPC统一ID

	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
	// 子网， 当网络类型值为 VPC时，与UnSubnetId 两者必须填一项

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 系统分配的子网统一 ID

	UnSubnetId *string `json:"UnSubnetId,omitempty" name:"UnSubnetId"`
	// 指定IP地址，仅VPC网络支持；若不填写、将在该子网下随机分配 IP

	MountIP *string `json:"MountIP,omitempty" name:"MountIP"`
	// 文件系统绑定的存储包，每个文件系统只能绑定一个

	StorageResourcePkgId *string `json:"StorageResourcePkgId,omitempty" name:"StorageResourcePkgId"`
	// 文件系统绑定的带宽包，每个文件系统只能绑定一个

	BandwidthResourcePkgId *string `json:"BandwidthResourcePkgId,omitempty" name:"BandwidthResourcePkgId"`
	// 用户自定义文件系统名称,与CreationToken 两者必须填一项

	FsName *string `json:"FsName,omitempty" name:"FsName"`
	// cfs资源池id。

	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
	// CFS文件系统版本。

	CfsVersion *string `json:"CfsVersion,omitempty" name:"CfsVersion"`
	// 用于新建文件系统的备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *CreateCfsFileSystemWithBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsFileSystemWithBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {

	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。

	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`
	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。

	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type RunMonitorServiceEnabled struct {

	// 是否开启[云监控](/tcloud/omtool/BARAD)服务。取值范围：<br><li>TRUE：表示开启云监控服务<br><li>FALSE：表示不开启云监控服务<br><br>默认取值：TRUE。

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type BindAutoBackupPolicyRequest struct {
	*tchttp.BaseRequest

	// 要绑定的定期备份策略ID。

	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	// 要绑定的云硬盘ID列表。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 要绑定的实例ID列表。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 要绑定的文件系统ID列表。

	FileSystemIds []*string `json:"FileSystemIds,omitempty" name:"FileSystemIds"`
}

func (r *BindAutoBackupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAutoBackupPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDisksWithBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的云硬盘ID列表。

		DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDisksWithBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDisksWithBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoBackupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>auto-backup-policy-id - Array of String - 是否必填：否 -（过滤条件）按定期备份策略ID进行过滤。定期备份策略ID形如：`abp-11112222`。<br><li>auto-backup-policy-state - Array of String - 是否必填：否 -（过滤条件）按定期备份策略的状态进行过滤。定期备份策略ID形如：`abp-11112222`。(NORMAL：正常 | ISOLATED：已隔离。)<br><li>auto-backup-policy-name - Array of String - 是否必填：否 -（过滤条件）按定期备份策略名称进行过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 输出定期备份策略列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 定期备份策略列表排序的依据字段。取值范围：<br><li>CREATETIME：依据定期备份策略的创建时间排序<br>默认按创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeAutoBackupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoBackupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupGroupDeniedAction struct {

	// 备份组ID。

	BackupGroupId *string `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
	// 具体的备份操作掩码列表。

	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type LoginSettings struct {

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux机器密码需10到30位，至少包括三项([a-z],[A-Z],[0-9]和[()`~!@#$%^&*-+=_|{}[]:;'<>,.?/]的特殊符号)。<br><li>Windows机器密码需12到30位，至少包括三项([a-z],[A-Z],[0-9]和[()`~!@#$%^&*-+=_|{}[]:;'<>,.?/]的特殊符号),密码不允许包含用户名密码不允许以`/`符号开头。<br><li>如果实例即包含`Linux`实例又包含`Windows`实例，则密码复杂度限制按照`Windows`实例的限制

	Password *string `json:"Password,omitempty" name:"Password"`
	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。

	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type UnbindAutoBackupPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindAutoBackupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindAutoBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Backup struct {

	// 备份名称。

	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
	// 是否为远程复制的备份。

	CopyFromRemote *bool `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`
	// 是否永久保留。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 备份的状态。取值范围：<br><li>NORMAL：正常<br><li>CREATING：创建中<br><li>ROLLBACKING：回滚中<br><li>COPYING_FROM_REMOTE：跨地域复制拷贝中。

	BackupState *string `json:"BackupState,omitempty" name:"BackupState"`
	// 备份的到期时间。

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 备份创建的进度。

	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
	// 备份被共享的次数。

	ShareReference *uint64 `json:"ShareReference,omitempty" name:"ShareReference"`
	// 备份类型，目前该项取值可以为PRIVATE_BACKUP或者SHARED_BACKUP。

	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
	// 创建备份的云硬盘大小，单位GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 创建备份的云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 备份所属项目ID。

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// 备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
	// 备份当前正在远程复制的目标地域列表。

	CopyingToRegions []*string `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`
	// 备份是否为加密备份。

	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`
	// 备份的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建此备份的云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 备份所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 备份绑定的标签列表。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 用户AppId。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 云盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 创建备份时刻，云硬盘各属性的详情。

	DiskDetails *string `json:"DiskDetails,omitempty" name:"DiskDetails"`
	// 全量、增量备份信息；FULL表示全量备份，INC表示增量备份。

	BackupClass *string `json:"BackupClass,omitempty" name:"BackupClass"`
	// 备份关联的备份组。

	BackupGroupId *string `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest

	// 需要创建备份的云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 备份名称。

	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
	// 指定备份到期时间，如果未传入该参数，默认为永久保留。

	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
	// 创建备份的带宽上限，范围：[0, 100]

	CreateSpeed *uint64 `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
}

func (r *CreateBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagSpecification struct {

	// 标签绑定的资源类型。

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 标签对列表。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CopyBackupCrossRegionsRequest struct {
	*tchttp.BaseRequest

	// 需要跨地域复制的源备份ID。

	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
	// 备份需要复制的目的地域列表。

	DestinationRegions []*string `json:"DestinationRegions,omitempty" name:"DestinationRegions"`
	// 新复制的备份名称。

	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

func (r *CopyBackupCrossRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyBackupCrossRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoBackupPolicyAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoBackupPolicyAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoBackupPolicyAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupFileSystem struct {

	// 实例绑定的定期备份策略列表。

	AutoBackupPolicyIdSet []*string `json:"AutoBackupPolicyIdSet,omitempty" name:"AutoBackupPolicyIdSet"`
	// 文件系统ID。

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 用户AppId。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 文件系统的最新备份时间。

	LatestBackupTime *string `json:"LatestBackupTime,omitempty" name:"LatestBackupTime"`
	// 文件系统的备份ID列表。

	BackupIdSet []*string `json:"BackupIdSet,omitempty" name:"BackupIdSet"`
	// 修改时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 文件系统当前存在的最早备份的时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 文件系统关联的云盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

type ProjectSpecification struct {

	// 资源类型,默认instance

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 项目Id

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
}

type DescribeBackupInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 满足条件的数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例的备份详情。

		BackupInstanceSet []*BackupInstance `json:"BackupInstanceSet,omitempty" name:"BackupInstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
