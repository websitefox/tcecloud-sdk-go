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
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 字段列表

		Fields []*RuleInfoFields `json:"Fields,omitempty" name:"Fields"`
		// 规则信息

		Rule *AuditRules `json:"Rule,omitempty" name:"Rule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleSwitchListRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 规则类型

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 规则属性

	IsInner *int64 `json:"IsInner,omitempty" name:"IsInner"`
	// 风险等级

	DangerLevel *int64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 规则开关

	IsOpened *int64 `json:"IsOpened,omitempty" name:"IsOpened"`
	// 资产Id

	AssetsId *int64 `json:"AssetsId,omitempty" name:"AssetsId"`
	// 规则分类

	RuleClassify *int64 `json:"RuleClassify,omitempty" name:"RuleClassify"`
	// 规则类别（全部 -1 默认 0 常规 1  宽松 2  严格 3 用户配置4 ）

	ConfigType *int64 `json:"ConfigType,omitempty" name:"ConfigType"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *DescribeRuleSwitchListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleSwitchListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmSettingSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmSettingSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmSettingSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRestoreLogTaskRequest struct {
	*tchttp.BaseRequest

	// 备份日志Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyRestoreLogTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRestoreLogTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerAdditionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否主动展示新手引导

		IsGuideOpen *bool `json:"IsGuideOpen,omitempty" name:"IsGuideOpen"`
		// 日志投递引导状态（1关闭、2开启 ）

		LogDeliveryGuide *int64 `json:"LogDeliveryGuide,omitempty" name:"LogDeliveryGuide"`
		// 报表引导状态 1.关闭 2.开启

		ReportGuide *int64 `json:"ReportGuide,omitempty" name:"ReportGuide"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerAdditionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerAdditionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcAbnormalLog struct {

	// 数据库用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 客户端IP

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 异常类型，取值1~4，分别对应：异常时间登录、 异常登录ip 、  异常操作、   异常资源访问

	AbnormalType *uint64 `json:"AbnormalType,omitempty" name:"AbnormalType"`
	// 操作时间

	OpTime *string `json:"OpTime,omitempty" name:"OpTime"`
	// SQL内容

	SqlContent *string `json:"SqlContent,omitempty" name:"SqlContent"`
}

type DescribeOldAlarmStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOldAlarmStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOldAlarmStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQpsAndStoreStatsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeQpsAndStoreStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQpsAndStoreStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DasbInstance struct {

	// 实例key

	PincKey *string `json:"PincKey,omitempty" name:"PincKey"`
	// 实例名称

	PsysName *string `json:"PsysName,omitempty" name:"PsysName"`
	// 公网ip

	MpublicIp *string `json:"MpublicIp,omitempty" name:"MpublicIp"`
	// 产品型号

	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`
	// 购买时间

	PaddTime *uint64 `json:"PaddTime,omitempty" name:"PaddTime"`
	// 内网IP

	MprivateIp *string `json:"MprivateIp,omitempty" name:"MprivateIp"`
	// 实例状态

	PsysStatus *uint64 `json:"PsysStatus,omitempty" name:"PsysStatus"`
	// 实例开始时间

	PstartTime *uint64 `json:"PstartTime,omitempty" name:"PstartTime"`
	// 实例结束时间

	PendTime *uint64 `json:"PendTime,omitempty" name:"PendTime"`
	// 实例所在地域ID

	Pregion *uint64 `json:"Pregion,omitempty" name:"Pregion"`
	// 实例所在地区ID

	PzoneId *uint64 `json:"PzoneId,omitempty" name:"PzoneId"`
	// CVM实例名列表

	CvmInstanceNames []*string `json:"CvmInstanceNames,omitempty" name:"CvmInstanceNames"`
	// CVM实例ID列表

	CvmInstanceIds []*string `json:"CvmInstanceIds,omitempty" name:"CvmInstanceIds"`
}

type DsgcClientInfo struct {

	// 腾讯云用户UIN

	TcloudUserId *string `json:"TcloudUserId,omitempty" name:"TcloudUserId"`
	// 腾讯云用户AppId

	TcloudAppId *string `json:"TcloudAppId,omitempty" name:"TcloudAppId"`
	// secret_id

	ServerSecretId *string `json:"ServerSecretId,omitempty" name:"ServerSecretId"`
	// secret_key

	ServerSecretKey *string `json:"ServerSecretKey,omitempty" name:"ServerSecretKey"`
	// 用户状态

	UsedStatus *int64 `json:"UsedStatus,omitempty" name:"UsedStatus"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeAnalysisQpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据列表

		Data []*TimeValue `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAnalysisQpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAnalysisQpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupSettingRequest struct {
	*tchttp.BaseRequest

	// 备份日志保留时长

	BackupLogSaveTime *int64 `json:"BackupLogSaveTime,omitempty" name:"BackupLogSaveTime"`
	// 恢复日志保留时长

	RestoreLogSaveTime *int64 `json:"RestoreLogSaveTime,omitempty" name:"RestoreLogSaveTime"`
}

func (r *ModifyBackupSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBackupSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaRouteListRequest struct {
	*tchttp.BaseRequest

	// Ckafka接入类型

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
	// Ckafka实例的地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Ckafka实例的id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Ckafka实例的名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *DescribeCkafkaRouteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaRouteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据列表

		List []*AssetsInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DasbCvmInstanceType struct {

	// 主机实例类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 主机实例名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 系统盘类型

	SystemDiskType *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`
	// 数据盘类型

	DataDiskType *string `json:"DataDiskType,omitempty" name:"DataDiskType"`
	// 所属区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DasbInstanceLog struct {

	// 日志ID

	LogId *uint64 `json:"LogId,omitempty" name:"LogId"`
	// 日志内容

	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
	// 日志记录时间

	LogTime *uint64 `json:"LogTime,omitempty" name:"LogTime"`
}

type CreateNewProduceReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNewProduceReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNewProduceReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatcdsAuditSaasResourceRequest struct {
	*tchttp.BaseRequest

	//  资源类型 基础版，高级版，企业版，旗舰版

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	// 资产扩展包

	AssetExp *int64 `json:"AssetExp,omitempty" name:"AssetExp"`
	//  // 存储扩展包

	StorageExp *int64 `json:"StorageExp,omitempty" name:"StorageExp"`
	//  //日志投递服务

	DeliveryExp *int64 `json:"DeliveryExp,omitempty" name:"DeliveryExp"`
	//  // 计费周期 只支持 "m"

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	//  // 计费时长

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// // 计费模式 0后付费，1预付费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *CreatcdsAuditSaasResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatcdsAuditSaasResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域描述

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域可用状态

	RegionState *int64 `json:"RegionState,omitempty" name:"RegionState"`
}

type DescribeBackupLogListRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式(desc=倒叙,asc=升序)

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 排序字段(支持'StartTime')

	Field *string `json:"Field,omitempty" name:"Field"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 备份日志状态0未完成的,1备份文件，2恢复中，3已恢复，4.已删除,全部查询-1

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeBackupLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcPrivateLinkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcPrivateLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcPrivateLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// VPC通道列表数据

		List []*AuditVpcAccess `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志信息列表

		List []*AuditLogInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmSmsConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发送间隔

		Interval *string `json:"Interval,omitempty" name:"Interval"`
		// 下发手机号码

		PhoneNumberSet []*string `json:"PhoneNumberSet,omitempty" name:"PhoneNumberSet"`
		// 短信告警SdkAppId

		SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`
		// 短信告警SmsSecretId

		SmsSecretId *string `json:"SmsSecretId,omitempty" name:"SmsSecretId"`
		// 短信告警SecretKey

		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
		// 签名

		Sign *string `json:"Sign,omitempty" name:"Sign"`
		// 地域

		SmsRegion *string `json:"SmsRegion,omitempty" name:"SmsRegion"`
		// 模版 ID

		TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmSmsConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmSmsConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquireCreatecdsAuditSaasResourceRequest struct {
	*tchttp.BaseRequest

	// saas数审的规格版本

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	// 是否扩展资产包

	AssetExp *uint64 `json:"AssetExp,omitempty" name:"AssetExp"`
	// 是否存储扩展包

	StorageExp *uint64 `json:"StorageExp,omitempty" name:"StorageExp"`
	// 是否日志投递资产包

	DeliveryExp *uint64 `json:"DeliveryExp,omitempty" name:"DeliveryExp"`
	// 计费周期 只支持 "m"

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 计费时长

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 计费模式 0后付费，1预付费 }

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *InquireCreatecdsAuditSaasResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquireCreatecdsAuditSaasResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditBehaviour struct {

	// 创建时间

	AddTime *uint64 `json:"AddTime,omitempty" name:"AddTime"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 修改人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// 是否删除

	IsDelete *uint64 `json:"IsDelete,omitempty" name:"IsDelete"`
	// 操作Id

	OperateId *uint64 `json:"OperateId,omitempty" name:"OperateId"`
	// 操作名称

	OperateName *string `json:"OperateName,omitempty" name:"OperateName"`
	// 备注

	OperateRemark *string `json:"OperateRemark,omitempty" name:"OperateRemark"`
	// 操作action

	OperateUrl *string `json:"OperateUrl,omitempty" name:"OperateUrl"`
	// 操作父Id

	ParentId *uint64 `json:"ParentId,omitempty" name:"ParentId"`
	// 操作参数

	QueryParams *string `json:"QueryParams,omitempty" name:"QueryParams"`
	// 修改时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CreatcdsAuditSaasResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatcdsAuditSaasResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatcdsAuditSaasResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TimeValue struct {

	// 时间

	Time *int64 `json:"Time,omitempty" name:"Time"`
	// 统计数值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DeleteVpcPrivateLinkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcPrivateLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcPrivateLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCkafkaConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCkafkaConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCkafkaConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExternalAgentRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeExternalAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExternalAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogStoreTrendRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeLogStoreTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogStoreTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReportPdfRequest struct {
	*tchttp.BaseRequest

	// 报表 Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *CreateReportPdfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateReportPdfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcStatisticsItem struct {

	// 统计标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 今天数量

	TodayCount *uint64 `json:"TodayCount,omitempty" name:"TodayCount"`
	// 昨天数量

	YestodayCount *uint64 `json:"YestodayCount,omitempty" name:"YestodayCount"`
	// 类型：0，自定义；非0为预定义

	DefineType *uint64 `json:"DefineType,omitempty" name:"DefineType"`
	// 排序号

	Sort *uint64 `json:"Sort,omitempty" name:"Sort"`
}

type DeleteCkafkaConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteCkafkaConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCkafkaConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigFields struct {

	// 客户端IP

	ClientIp *RuleConfigFieldType `json:"ClientIp,omitempty" name:"ClientIp"`
	// 数据库IP

	DbIp *RuleConfigFieldType `json:"DbIp,omitempty" name:"DbIp"`
	// 数据库名

	DbName *RuleConfigFieldType `json:"DbName,omitempty" name:"DbName"`
	// 数据库端口

	DbPort *RuleConfigFieldType `json:"DbPort,omitempty" name:"DbPort"`
	// 数据库用户

	DbUser *RuleConfigFieldType `json:"DbUser,omitempty" name:"DbUser"`
	// 影响行数

	EffectRow *RuleConfigFieldType `json:"EffectRow,omitempty" name:"EffectRow"`
	// 执行时间

	ExecTime *RuleConfigFieldType `json:"ExecTime,omitempty" name:"ExecTime"`
	// 操作语句

	OpSql *RuleConfigFieldType `json:"OpSql,omitempty" name:"OpSql"`
	// 操作时间

	OpTime *RuleConfigFieldType `json:"OpTime,omitempty" name:"OpTime"`
	// 操作类型

	SqlType *RuleConfigFieldType `json:"SqlType,omitempty" name:"SqlType"`
	// 表名

	TableName *RuleConfigFieldType `json:"TableName,omitempty" name:"TableName"`
}

type DescribeServiceStateRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeServiceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReportMissionRequest struct {
	*tchttp.BaseRequest

	// 报表Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 任务名称 不变更为""

	TplName *string `json:"TplName,omitempty" name:"TplName"`
	// 报表类型 1:单次报表 2:周期报表 不变更为0

	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`
	// 任务起停 1关闭 2开启 单次报表默认为2'

	MissionStart *int64 `json:"MissionStart,omitempty" name:"MissionStart"`
	// 报告说明 不变更为""

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 报表模板 1:综合分析报告 2:等保合规报告 不变更为0

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 包含资产 0:全部资产 其余为资产id 不变更为[]

	AssetsId []*int64 `json:"AssetsId,omitempty" name:"AssetsId"`
	// 统计周期 1:24小时 7:近一周 30:近30天 90:近90天 180:近180天 不变

	CntDay *uint64 `json:"CntDay,omitempty" name:"CntDay"`
	// 重复周期 1:每天 2:每周 3:每月 不变更为0

	CntCycle *uint64 `json:"CntCycle,omitempty" name:"CntCycle"`
	// 执行时间 格式15:04 到分钟

	CntDate *string `json:"CntDate,omitempty" name:"CntDate"`
	// 执行日期 重复周期为天：无意义 周：星期几 1-7  月每月

	CntTime *uint64 `json:"CntTime,omitempty" name:"CntTime"`
	// 报表通知  int  1关闭 2开启 不变更为0

	Notification *int64 `json:"Notification,omitempty" name:"Notification"`
}

func (r *ModifyReportMissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReportMissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CvmInfo struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// cvm所在vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// cvm所在子网

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// cvm私有ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 操作系统

	SystemType *int64 `json:"SystemType,omitempty" name:"SystemType"`
}

type RouteInfo struct {

	// 接入ip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 接入端口

	Vport *string `json:"Vport,omitempty" name:"Vport"`
	// 接入域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 接入域名端口

	DomainPort *string `json:"DomainPort,omitempty" name:"DomainPort"`
}

type CreateAgentDeployResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAgentDeployResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgentDeployResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaTopicListRequest struct {
	*tchttp.BaseRequest

	// ckafka的实例接入类型

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
	// ckafka的实例地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// ckafka的实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// ckafka的实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *DescribeCkafkaTopicListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaTopicListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGuideOpenRequest struct {
	*tchttp.BaseRequest

	// 打开类型 1 用户主动关闭 2 用户查看了

	GuideOpenType *int64 `json:"GuideOpenType,omitempty" name:"GuideOpenType"`
	// 日志投递引导状态（1关闭、2开启 ）

	LogDeliveryGuide *int64 `json:"LogDeliveryGuide,omitempty" name:"LogDeliveryGuide"`
	// 报表引导状态 1.关闭 2.开启

	ReportGuide *int64 `json:"ReportGuide,omitempty" name:"ReportGuide"`
}

func (r *ModifyCustomerGuideOpenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomerGuideOpenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCvmAssetsSaveRequest struct {
	*tchttp.BaseRequest

	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 添加方式(2: cvm;3: ip)

	AssetsAddType *int64 `json:"AssetsAddType,omitempty" name:"AssetsAddType"`
	// cvm 实例 ID (AssetsAddType === 2)

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// VpcId (AssetsAddType===3)

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 数据资产列表

	Assets []*AssetsForm `json:"Assets,omitempty" name:"Assets"`
	// cvm实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 添加方式

	AddType *int64 `json:"AddType,omitempty" name:"AddType"`
	// 子网

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 安全证书密码

	PemSecretKey *string `json:"PemSecretKey,omitempty" name:"PemSecretKey"`
	// 是否需要安全证书

	NeedPem *bool `json:"NeedPem,omitempty" name:"NeedPem"`
	// 安全证书

	DBSecretPem *string `json:"DBSecretPem,omitempty" name:"DBSecretPem"`
	// IP/域名

	AssetsIp *string `json:"AssetsIp,omitempty" name:"AssetsIp"`
	// 端口

	AssetsPort *int64 `json:"AssetsPort,omitempty" name:"AssetsPort"`
	// 数据资产名称

	AssetsName *string `json:"AssetsName,omitempty" name:"AssetsName"`
	// 数据资产类型版本

	AssetsVersion *string `json:"AssetsVersion,omitempty" name:"AssetsVersion"`
	// 数据资产类型

	AssetsType *string `json:"AssetsType,omitempty" name:"AssetsType"`
}

func (r *CreateCvmAssetsSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCvmAssetsSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentStopRequest struct {
	*tchttp.BaseRequest

	// 标识Id

	AgentDeployId *uint64 `json:"AgentDeployId,omitempty" name:"AgentDeployId"`
}

func (r *ModifyAgentStopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAgentStopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DbauditTypesInfo struct {

	// 规格描述

	InstanceVersionName *string `json:"InstanceVersionName,omitempty" name:"InstanceVersionName"`
	// 规格名称

	InstanceVersionKey *string `json:"InstanceVersionKey,omitempty" name:"InstanceVersionKey"`
	// 最大吞吐量

	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`
	// 最大实例数

	MaxInstances *uint64 `json:"MaxInstances,omitempty" name:"MaxInstances"`
	// 入库速率（每小时）

	InsertSpeed *uint64 `json:"InsertSpeed,omitempty" name:"InsertSpeed"`
	// 最大在线存储量，单位：条

	OnlineStorageCapacity *uint64 `json:"OnlineStorageCapacity,omitempty" name:"OnlineStorageCapacity"`
	// 最大归档存储量，单位：条

	ArchivingStorageCapacity *uint64 `json:"ArchivingStorageCapacity,omitempty" name:"ArchivingStorageCapacity"`
}

type DeleteBackupLogListRequest struct {
	*tchttp.BaseRequest

	// 备份日志Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteBackupLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcSensitiveDataItem struct {

	// 数据库名

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 表名

	TableName *string `json:"TableName,omitempty" name:"TableName"`
	// 数据摘要

	DataSummary *string `json:"DataSummary,omitempty" name:"DataSummary"`
	// 敏感级别

	SensitiveLvel *uint64 `json:"SensitiveLvel,omitempty" name:"SensitiveLvel"`
	// 数据类型

	DataType *string `json:"DataType,omitempty" name:"DataType"`
	// 最近访问时间

	LatestOpTime *string `json:"LatestOpTime,omitempty" name:"LatestOpTime"`
}

type LogDeliveryInfo struct {

	// 日志类型

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 投递的topicid

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递的topicname

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type DescribeCkafkaTopicListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ckafka实例的主题列表

		List []*TopicInfo `json:"List,omitempty" name:"List"`
		// ckafka实例的主题列表的总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCkafkaTopicListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaTopicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProduceReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProduceReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProduceReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogListRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式(desc=倒叙,asc=升序)

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 排序字段(opTime=时间,dangerLvl=风险等级)

	Field *string `json:"Field,omitempty" name:"Field"`
	// 风险等级(0-安全,1-低风险,2-中风险,3-高风险,不传全部)

	DangerLevel *string `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 数据库名称

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 数据库端口

	DbPort *int64 `json:"DbPort,omitempty" name:"DbPort"`
	// 数据库 IP

	DbIp *string `json:"DbIp,omitempty" name:"DbIp"`
	// 资产 ID

	AssetsId *int64 `json:"AssetsId,omitempty" name:"AssetsId"`
	// 会话 ID

	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
	// 客户端 IP

	ClientSideIp *string `json:"ClientSideIp,omitempty" name:"ClientSideIp"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 命中规则

	HitRule *int64 `json:"HitRule,omitempty" name:"HitRule"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 恢复日志id

	RestoreLogId *int64 `json:"RestoreLogId,omitempty" name:"RestoreLogId"`
}

func (r *DescribeLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogDeliveryCkafkaConfig struct {

	// 接入类型

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 虚拟ip  VipType 为7 有效

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 虚拟端口 VipType 为7有效

	Vport *string `json:"Vport,omitempty" name:"Vport"`
	// 域名  VipType 为1有效

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名端口 VipType 为1有效

	DomainPort *string `json:"DomainPort,omitempty" name:"DomainPort"`
	// 实例地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 实例vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 实例子网

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 实例健康状态， 1：健康，2：告警，3：异常',

	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`
	// 日志类型

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 投递的topicid

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递的topicname

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 投递状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 状态信息

	StatusMessages *string `json:"StatusMessages,omitempty" name:"StatusMessages"`
	// 开启or关闭，投递状态1为开启，0 关闭，默认开启，但是如果没有topic 则关闭

	IsOpen *int64 `json:"IsOpen,omitempty" name:"IsOpen"`
}

type LogDeliveryType struct {

	// 日志投递类型

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 投递日志类型的名称

	LogTypeName *string `json:"LogTypeName,omitempty" name:"LogTypeName"`
	// 投递日志类型的描述

	LogTypeDesc *string `json:"LogTypeDesc,omitempty" name:"LogTypeDesc"`
}

type DescribeLogDeliveryTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持日志投递的类型列表

		List []*LogDeliveryType `json:"List,omitempty" name:"List"`
		// 支持日志投递类型的总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogDeliveryTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDeliveryTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCkafkaStopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCkafkaStopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCkafkaStopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeAgentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Agent列表数组

		List []*AuditAgentInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAllRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentDeployResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Agent部署列表信息

		List []*AgentDeployInfo `json:"List,omitempty" name:"List"`
		// 是否打通VPC

		VpcAccessStatus *uint64 `json:"VpcAccessStatus,omitempty" name:"VpcAccessStatus"`
		// 0: 任务未执行 1 任务执行中

		JobStatus *int64 `json:"JobStatus,omitempty" name:"JobStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentDeployResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentDeployResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQpsAndStoreStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// qps是否开启告警

		QpsAlarm *int64 `json:"QpsAlarm,omitempty" name:"QpsAlarm"`
		// qps当前值

		QpsUse *int64 `json:"QpsUse,omitempty" name:"QpsUse"`
		// qps规格

		QpsSpec *int64 `json:"QpsSpec,omitempty" name:"QpsSpec"`
		// 存储是否开启告警，0不开启，1开启

		StoreAlarm *int64 `json:"StoreAlarm,omitempty" name:"StoreAlarm"`
		// 存储当前值

		StoreUse *int64 `json:"StoreUse,omitempty" name:"StoreUse"`
		// 存储规格

		StoreSpec *int64 `json:"StoreSpec,omitempty" name:"StoreSpec"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQpsAndStoreStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQpsAndStoreStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStoreTopLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可存储天数

		LeftOverDay *int64 `json:"LeftOverDay,omitempty" name:"LeftOverDay"`
		// 需扩展的空间

		NeedExtendStore *int64 `json:"NeedExtendStore,omitempty" name:"NeedExtendStore"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStoreTopLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStoreTopLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRuleSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRuleSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditIpDeploy struct {

	// 服务器 IP（按IP部署）

	ServiceIp *string `json:"ServiceIp,omitempty" name:"ServiceIp"`
	// Ssh 端口号

	SshPort *uint64 `json:"SshPort,omitempty" name:"SshPort"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 实例 Id（CVM部署）

	Id *string `json:"Id,omitempty" name:"Id"`
	// 起始 IP（按IP段部署）

	StartIp *string `json:"StartIp,omitempty" name:"StartIp"`
	// 结束 IP（按IP段部署）

	EndIp *string `json:"EndIp,omitempty" name:"EndIp"`
}

type SendAlarmEnterpriseWechatTestResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendAlarmEnterpriseWechatTestResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendAlarmEnterpriseWechatTestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemVersionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSystemVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportRequest struct {
	*tchttp.BaseRequest

	// 报表 Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentUninstallResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAgentUninstallResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAgentUninstallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditRuleSwitch struct {

	// 风险等级

	DangerLevel *int64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 是否为内置规则

	IsInner *int64 `json:"IsInner,omitempty" name:"IsInner"`
	// 规则是否开启

	IsOpened *int64 `json:"IsOpened,omitempty" name:"IsOpened"`
	// 规则备注

	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`
	// 规则 ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则类型

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 删除状态

	IsDelete *int64 `json:"IsDelete,omitempty" name:"IsDelete"`
	// 规则分类

	RuleClassify *int64 `json:"RuleClassify,omitempty" name:"RuleClassify"`
	// 规则类别（ 默认 0 常规 1  宽松 2  严格 3 用户配置4 ）

	ConfigType *int64 `json:"ConfigType,omitempty" name:"ConfigType"`
}

type DeleteReportMissionRequest struct {
	*tchttp.BaseRequest

	// 报表任务id列表

	Id []*int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteReportMissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReportMissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetsTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产类型列表

		List []*string `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetsTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志信息列表

		List []*AuditLogInfo `json:"List,omitempty" name:"List"`
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

type AuditOperateLogInfo struct {

	// 操作时间

	AddTime *uint64 `json:"AddTime,omitempty" name:"AddTime"`
	// 标识

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 操作 IP

	OpIp *string `json:"OpIp,omitempty" name:"OpIp"`
	// 操作行为

	Op *string `json:"Op,omitempty" name:"Op"`
	// 行为 ID

	OperateId *uint64 `json:"OperateId,omitempty" name:"OperateId"`
	// 操作日志的操作参数

	ParamData *string `json:"ParamData,omitempty" name:"ParamData"`
	// 行为分类

	TypeStr *string `json:"TypeStr,omitempty" name:"TypeStr"`
	// 操作账户

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type RuleConfigFieldType struct {

	// 英文规则名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 中文规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ModifyAssetsSaveRequest struct {
	*tchttp.BaseRequest

	// 资产 ID

	Aid *int64 `json:"Aid,omitempty" name:"Aid"`
	// 数据资产名称

	AssetsName *string `json:"AssetsName,omitempty" name:"AssetsName"`
	// 数据资产类型版本

	AssetsVersion *string `json:"AssetsVersion,omitempty" name:"AssetsVersion"`
	// IP/域名

	AssetsIp *string `json:"AssetsIp,omitempty" name:"AssetsIp"`
	// 端口

	AssetsPort *int64 `json:"AssetsPort,omitempty" name:"AssetsPort"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 资产类型

	AssetsType *string `json:"AssetsType,omitempty" name:"AssetsType"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 是否需要安全证书

	NeedPem *bool `json:"NeedPem,omitempty" name:"NeedPem"`
	// 安全证书内容

	DBSecretPem *string `json:"DBSecretPem,omitempty" name:"DBSecretPem"`
	// 安全证书密码

	PemSecretKey *string `json:"PemSecretKey,omitempty" name:"PemSecretKey"`
}

func (r *ModifyAssetsSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetsSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CkafkaInstance struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

type DescribeAnalysisQpsRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAnalysisQpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAnalysisQpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmSettingSaveRequest struct {
	*tchttp.BaseRequest

	// agent 掉线告警

	AlarmSwitchAgentLogout *uint64 `json:"AlarmSwitchAgentLogout,omitempty" name:"AlarmSwitchAgentLogout"`
	// 企业微信告警

	AlarmSwitchEnterpriseWechat *uint64 `json:"AlarmSwitchEnterpriseWechat,omitempty" name:"AlarmSwitchEnterpriseWechat"`
	// 邮件告警

	AlarmSwitchMail *uint64 `json:"AlarmSwitchMail,omitempty" name:"AlarmSwitchMail"`
	// QPS 告警

	AlarmSwitchQps *uint64 `json:"AlarmSwitchQps,omitempty" name:"AlarmSwitchQps"`
	// 行为规则告警

	AlarmSwitchRule *uint64 `json:"AlarmSwitchRule,omitempty" name:"AlarmSwitchRule"`
	// 短信告警

	AlarmSwitchSms *uint64 `json:"AlarmSwitchSms,omitempty" name:"AlarmSwitchSms"`
	// SQL 风险告警

	AlarmSwitchSql *uint64 `json:"AlarmSwitchSql,omitempty" name:"AlarmSwitchSql"`
	// 时间间隔

	AlarmDuration *uint64 `json:"AlarmDuration,omitempty" name:"AlarmDuration"`
	// 报表通知  int  1关闭 2开启 不变更为0

	ReportNotification *int64 `json:"ReportNotification,omitempty" name:"ReportNotification"`
}

func (r *ModifyAlarmSettingSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmSettingSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetsInfo struct {

	// 创建时间

	AddTime *uint64 `json:"AddTime,omitempty" name:"AddTime"`
	// 资产 ID

	Aid *uint64 `json:"Aid,omitempty" name:"Aid"`
	// 数据资产 IP

	AssetsIp *string `json:"AssetsIp,omitempty" name:"AssetsIp"`
	// 数据资产名称

	AssetsName *string `json:"AssetsName,omitempty" name:"AssetsName"`
	// 数据资产端口

	AssetsPort *uint64 `json:"AssetsPort,omitempty" name:"AssetsPort"`
	// 数据资产类型

	AssetsType *string `json:"AssetsType,omitempty" name:"AssetsType"`
	// 资产版本

	AssetsVersion *string `json:"AssetsVersion,omitempty" name:"AssetsVersion"`
	// 是否动态

	AssetsAddType *uint64 `json:"AssetsAddType,omitempty" name:"AssetsAddType"`
	// 是否删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 最后一次修改时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 资产的vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 审计权限

	Permission *int64 `json:"Permission,omitempty" name:"Permission"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 用来区分自建资产是已通过cvm还是添加ip的方式

	AddType *uint64 `json:"AddType,omitempty" name:"AddType"`
	// 子网Id

	AssetSubnetId *string `json:"AssetSubnetId,omitempty" name:"AssetSubnetId"`
	// 是否已上传数据库私钥（0 否，1 是）

	UploadPem *int64 `json:"UploadPem,omitempty" name:"UploadPem"`
	// 资产状态栏 0:正常 1:已删除（目前仅对tencentDB有效）

	AliveStatus *int64 `json:"AliveStatus,omitempty" name:"AliveStatus"`
}

type BackupLog struct {

	// 索引

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 索引开始时间

	IndexStartTime *int64 `json:"IndexStartTime,omitempty" name:"IndexStartTime"`
	// 索引结束时间

	IndexEndTime *int64 `json:"IndexEndTime,omitempty" name:"IndexEndTime"`
	// 备份后压缩的大小，单位M

	BackupSize *int64 `json:"BackupSize,omitempty" name:"BackupSize"`
	// 日志状态 0备份未完成， 1备份文件，2恢复中，3已恢复，4.已删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeBackupLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总共多少条

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 备份日志列表

		List []*BackupLog `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcPrivateLinkRequest struct {
	*tchttp.BaseRequest

	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DeleteVpcPrivateLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcPrivateLinkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBehaviourRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBehaviourRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBehaviourRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportTplRequest struct {
	*tchttp.BaseRequest

	// 报表id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeReportTplRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportTplRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRestoreLogTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRestoreLogTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRestoreLogTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
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

type DescribeAllDbCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAllDbCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllDbCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifycdsAuditSaasResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifycdsAuditSaasResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifycdsAuditSaasResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditExternalAgent struct {

	// 标识

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Ip 地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type AgentDeployInfo struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// Vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 操作系统

	SystemType *int64 `json:"SystemType,omitempty" name:"SystemType"`
	// 自动化助手

	AutoAssistant *int64 `json:"AutoAssistant,omitempty" name:"AutoAssistant"`
	// Agent状态

	AgentStatus *int64 `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// id

	AgentId *int64 `json:"AgentId,omitempty" name:"AgentId"`
}

type DsgcAssociateSetting struct {

	// 数据库类型

	DbType *string `json:"DbType,omitempty" name:"DbType"`
	// 设置名称

	SettingName *string `json:"SettingName,omitempty" name:"SettingName"`
	// 所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 数据库版本

	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeReportListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据列表

		List []*Reports `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReportListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmRulesRequest struct {
	*tchttp.BaseRequest

	// 告警类型，不填默认获取所有类型。取值：sql_risk_high, sql_risk_middle, sql_risk_low, operation_risk_high, operation_risk_middle, operation_risk_low, qps_limit, storage_capacity_limit, agent_offline, log_delivery_abnormal, report_notification

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
}

func (r *DescribeAlarmRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGuideOpenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomerGuideOpenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomerGuideOpenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgentDeployRequest struct {
	*tchttp.BaseRequest

	// 地域(仅腾讯云内网 Agent 需要)

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// VPC(仅腾讯云内网 Agent 需要)

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 部署方式

	DeployType *uint64 `json:"DeployType,omitempty" name:"DeployType"`
	// 同时开通白名单(仅腾讯云外 Agent 需要，1 开启 0 关闭)

	WhiteListType *uint64 `json:"WhiteListType,omitempty" name:"WhiteListType"`
	// 判断是腾讯云内网Agent还是腾讯云外Agent(1-腾讯云内网Agent,2-腾讯云外Agent)

	NetWorkType *uint64 `json:"NetWorkType,omitempty" name:"NetWorkType"`
	// 部署信息数组

	IpDeploy *AuditIpDeploy `json:"IpDeploy,omitempty" name:"IpDeploy"`
}

func (r *CreateAgentDeployRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgentDeployRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcDataTrendItem struct {

	// 日期

	Day *string `json:"Day,omitempty" name:"Day"`
	// 数据量

	RiskCount *uint64 `json:"RiskCount,omitempty" name:"RiskCount"`
	// 资产类型：1， cvm总数；2，保护cvm数；5，waf总数；11，数据库总数；12，木马文件数；13，异地登录数；14，暴力破解数；15，漏洞数。

	RiskType *uint64 `json:"RiskType,omitempty" name:"RiskType"`
}

type DescribeAgentListRequest struct {
	*tchttp.BaseRequest

	// 流量来源

	TrafficType *int64 `json:"TrafficType,omitempty" name:"TrafficType"`
	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 资产 ID

	AssetsId *int64 `json:"AssetsId,omitempty" name:"AssetsId"`
	// 运行状态

	AgentStatus *int64 `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// 部署ip

	DeployIp *string `json:"DeployIp,omitempty" name:"DeployIp"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAgentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmEnterpriseWechatConfigRequest struct {
	*tchttp.BaseRequest

	// WebHook 地址

	WebUrl *string `json:"WebUrl,omitempty" name:"WebUrl"`
}

func (r *ModifyAlarmEnterpriseWechatConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmEnterpriseWechatConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Reports struct {

	// 生成时间

	AddTime *int64 `json:"AddTime,omitempty" name:"AddTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 报告 ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 审计 ID

	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否已删除

	IsDelete *int64 `json:"IsDelete,omitempty" name:"IsDelete"`
	// 发送目标

	Receivers *string `json:"Receivers,omitempty" name:"Receivers"`
	// 报告说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 报告文件

	ReportFile *string `json:"ReportFile,omitempty" name:"ReportFile"`
	// 状态

	ReportStatus *int64 `json:"ReportStatus,omitempty" name:"ReportStatus"`
	// 状态

	ReportTmpStatus *int64 `json:"ReportTmpStatus,omitempty" name:"ReportTmpStatus"`
	// 报告类型

	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`
	// 发送结果

	SendResult *string `json:"SendResult,omitempty" name:"SendResult"`
	// 发送类型

	SendType *string `json:"SendType,omitempty" name:"SendType"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 报告名称

	Title *string `json:"Title,omitempty" name:"Title"`
	// 报表模板

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 包含资产

	AssetsList []*AssetsInfo `json:"AssetsList,omitempty" name:"AssetsList"`
	// 时间范围 1:24小时 7:近一周 30:近30天 90:近90天 180:近180天 不变更为0

	CntDay *int64 `json:"CntDay,omitempty" name:"CntDay"`
}

type CreateTimerReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTimerReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTimerReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceList struct {

	// 签名时间

	List []*Services `json:"List,omitempty" name:"List"`
	// 总数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type SendCkafkaTestRequest struct {
	*tchttp.BaseRequest

	// 1为外网，7为支撑环境接入

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
	// ckafka实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 为7有效

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 为7有效

	Vport *string `json:"Vport,omitempty" name:"Vport"`
	// 域名，VipType为1有效

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名端口，VipType为1有效

	DomainPort *string `json:"DomainPort,omitempty" name:"DomainPort"`
	// 用户名，VipType为1有效

	Username *string `json:"Username,omitempty" name:"Username"`
	// 用户密码，VipType为1有效

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *SendCkafkaTestRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendCkafkaTestRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigFieldTypes struct {

	// Double 类型

	Double []*RuleConfigFieldType `json:"Double,omitempty" name:"Double"`
	// Int 类型

	Int []*RuleConfigFieldType `json:"Int,omitempty" name:"Int"`
	// String 类型

	String []*RuleConfigFieldType `json:"String,omitempty" name:"String"`
}

type DsgcRiskDataTopItem struct {

	// 资产类型

	DataType *string `json:"DataType,omitempty" name:"DataType"`
	// 资产名称

	DataName *string `json:"DataName,omitempty" name:"DataName"`
	// 敏感数据访问量

	SensitiveAccessCount *uint64 `json:"SensitiveAccessCount,omitempty" name:"SensitiveAccessCount"`
	// 异常行为数

	AbnormalActionCount *uint64 `json:"AbnormalActionCount,omitempty" name:"AbnormalActionCount"`
}

type DsgcSafeTrendItem struct {

	// 统计日期

	Day *string `json:"Day,omitempty" name:"Day"`
	// 敏感数据访问量

	SensitiveAccessCount *uint64 `json:"SensitiveAccessCount,omitempty" name:"SensitiveAccessCount"`
	// 异常行为数

	AbnormalOperationCount *uint64 `json:"AbnormalOperationCount,omitempty" name:"AbnormalOperationCount"`
	// 异常用户数

	AbnormalUserCount *uint64 `json:"AbnormalUserCount,omitempty" name:"AbnormalUserCount"`
}

type AuditLogInfo struct {

	// ai分数

	AiScore *float64 `json:"AiScore,omitempty" name:"AiScore"`
	// 应用用户

	AppUser *string `json:"AppUser,omitempty" name:"AppUser"`
	// 备份数据包

	BackPacket *string `json:"BackPacket,omitempty" name:"BackPacket"`
	// 客户端 IP

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 客户端 Mac

	ClientMac *string `json:"ClientMac,omitempty" name:"ClientMac"`
	// 终端名称

	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`
	// 客户端用户

	ClientUser *string `json:"ClientUser,omitempty" name:"ClientUser"`
	// 客户端端口

	ClientPort *uint64 `json:"ClientPort,omitempty" name:"ClientPort"`
	// 风险等级

	DangerLevel *uint64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 数据库 IP

	DbIp *string `json:"DbIp,omitempty" name:"DbIp"`
	// 数据库名称

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 数据库端口

	DbPort *uint64 `json:"DbPort,omitempty" name:"DbPort"`
	// 数据库用户

	DbUser *string `json:"DbUser,omitempty" name:"DbUser"`
	// 影响行数

	EffectRow *uint64 `json:"EffectRow,omitempty" name:"EffectRow"`
	// 执行时间

	ExecTime *uint64 `json:"ExecTime,omitempty" name:"ExecTime"`
	// 命中规则

	HitRule *string `json:"HitRule,omitempty" name:"HitRule"`
	// 日志 ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 数据资产名称

	InstanceId *uint64 `json:"InstanceId,omitempty" name:"InstanceId"`
	// 审计单元名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 操作语句(sql 语句)

	OpSql *string `json:"OpSql,omitempty" name:"OpSql"`
	// 操作时间(时间)

	OpTime *uint64 `json:"OpTime,omitempty" name:"OpTime"`
	// 返回消息

	RetMsg *string `json:"RetMsg,omitempty" name:"RetMsg"`
	// 返回码

	RetNo *uint64 `json:"RetNo,omitempty" name:"RetNo"`
	// 会话ID

	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
	// 操作类型

	SqlType *string `json:"SqlType,omitempty" name:"SqlType"`
	// 表名

	TableName *string `json:"TableName,omitempty" name:"TableName"`
	// 数据资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 规则集合

	HitRules []*HitRules `json:"HitRules,omitempty" name:"HitRules"`
}

type CreateAgentBatchDeployResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAgentBatchDeployResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgentBatchDeployResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单 ID

	Id []*string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductCost struct {

	//  //总费用

	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
	//  //优惠后总价

	RealCost *float64 `json:"RealCost,omitempty" name:"RealCost"`
}

type ModifyCkafkaStopRequest struct {
	*tchttp.BaseRequest

	// 日志投递类型的配置

	LogDeliveryInfo []*LogDeliveryInfo `json:"LogDeliveryInfo,omitempty" name:"LogDeliveryInfo"`
}

func (r *ModifyCkafkaStopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCkafkaStopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReportMissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReportMissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReportMissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentDeployRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 操作系统(-1 全部，0 linux，1 windows)

	SystemType *int64 `json:"SystemType,omitempty" name:"SystemType"`
	// Agent状态 （0 已离线，1 运行中，2 过载停止，3 手动停止，4 已卸载，5 未部署，6 部署中，8 卸载中，9 卸载失败）

	AgentStatus *int64 `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// 自动化助手（-1 全部 ，0 离线，1 在线，2 未安装）

	AutoAssistant *int64 `json:"AutoAssistant,omitempty" name:"AutoAssistant"`
	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// CVM IP/名称 搜索

	SearchValues []*NameValueString `json:"SearchValues,omitempty" name:"SearchValues"`
}

func (r *DescribeAgentDeployRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentDeployRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// agent掉线告警，值为0或者1

		AlarmSwitchAgentLogout *uint64 `json:"AlarmSwitchAgentLogout,omitempty" name:"AlarmSwitchAgentLogout"`
		// 邮件告警，值为0或者1

		AlarmSwitchMail *uint64 `json:"AlarmSwitchMail,omitempty" name:"AlarmSwitchMail"`
		// QPS告警，值为0或者1

		AlarmSwitchQps *uint64 `json:"AlarmSwitchQps,omitempty" name:"AlarmSwitchQps"`
		// 行为规则告警，值为0或者1

		AlarmSwitchRule *uint64 `json:"AlarmSwitchRule,omitempty" name:"AlarmSwitchRule"`
		// sql风险告警，值为0或者1

		AlarmSwitchSql *uint64 `json:"AlarmSwitchSql,omitempty" name:"AlarmSwitchSql"`
		// 短信告警，值为0或者1

		AlarmSwitchSms *uint64 `json:"AlarmSwitchSms,omitempty" name:"AlarmSwitchSms"`
		// 企业微信告警，值为0或者1

		AlarmSwitchEnterpriseWechat *uint64 `json:"AlarmSwitchEnterpriseWechat,omitempty" name:"AlarmSwitchEnterpriseWechat"`
		// 时间间隔

		AlarmDuration *uint64 `json:"AlarmDuration,omitempty" name:"AlarmDuration"`
		// 报表通知  int  1关闭 2开启 不变更为0

		ReportNotification *int64 `json:"ReportNotification,omitempty" name:"ReportNotification"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 规则列表数据

		List []*AuditRules `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRulesListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRulesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditVpcAccess struct {

	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 资产列表

	AssetsList []*AssetsInfo `json:"AssetsList,omitempty" name:"AssetsList"`
	// 标识

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 类型

	VpcType *int64 `json:"VpcType,omitempty" name:"VpcType"`
	// 无用参数

	DeployAgentType *int64 `json:"DeployAgentType,omitempty" name:"DeployAgentType"`
	// 部署Agent数量

	DeployAgentNumber *int64 `json:"DeployAgentNumber,omitempty" name:"DeployAgentNumber"`
	// ip地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type SignData struct {

	// 签名时间

	SignTime *uint64 `json:"SignTime,omitempty" name:"SignTime"`
	// 签名

	Sign *string `json:"Sign,omitempty" name:"Sign"`
}

type DescribeDashBoardAnalysisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DELETE 操作(TrendType == 1)

		DeleteData []*TimeValue `json:"DeleteData,omitempty" name:"DeleteData"`
		// INSERT 操作(TrendType == 1)

		InsertData []*TimeValue `json:"InsertData,omitempty" name:"InsertData"`
		// SELECT 操作(TrendType == 1)

		SelectData []*TimeValue `json:"SelectData,omitempty" name:"SelectData"`
		// SQL 语句总量(TrendType == 1)

		SqlData []*TimeValue `json:"SqlData,omitempty" name:"SqlData"`
		// UPDATE 操作(TrendType == 1)

		UpdateData []*TimeValue `json:"UpdateData,omitempty" name:"UpdateData"`
		// 高风险(TrendType == 2)

		HighRisk []*TimeValue `json:"HighRisk,omitempty" name:"HighRisk"`
		// 低风险(TrendType == 2)

		LowRisk []*TimeValue `json:"LowRisk,omitempty" name:"LowRisk"`
		// 活跃会话(TrendType == 3)

		ActiveSession []*TimeValue `json:"ActiveSession,omitempty" name:"ActiveSession"`
		// 失败会话(TrendType == 3)

		FailSession []*TimeValue `json:"FailSession,omitempty" name:"FailSession"`
		// 新建会话(TrendType == 1)

		NewSession []*TimeValue `json:"NewSession,omitempty" name:"NewSession"`
		// 中风险(TrendType == 2)

		MiddleRisk []*TimeValue `json:"MiddleRisk,omitempty" name:"MiddleRisk"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashBoardAnalysisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashBoardAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateRuleSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperateRuleSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateRuleSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditInstance struct {

	// 实例id

	Psid *uint64 `json:"Psid,omitempty" name:"Psid"`
	// 实例名字

	PsysName *string `json:"PsysName,omitempty" name:"PsysName"`
	// 公网ip

	MpublicIp *string `json:"MpublicIp,omitempty" name:"MpublicIp"`
	// 内网ip

	MprivateIp *string `json:"MprivateIp,omitempty" name:"MprivateIp"`
	// 规格pid

	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`
	// 域名

	Pdomain *string `json:"Pdomain,omitempty" name:"Pdomain"`
	// 添加时间

	PaddTime *uint64 `json:"PaddTime,omitempty" name:"PaddTime"`
	// 地区id

	Pregion *uint64 `json:"Pregion,omitempty" name:"Pregion"`
	// 网络id

	PuniqVpcId *string `json:"PuniqVpcId,omitempty" name:"PuniqVpcId"`
	// 子网id

	PuniqSubnetId *string `json:"PuniqSubnetId,omitempty" name:"PuniqSubnetId"`
}

type DescribeAllDbCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 关系型数据库数量

		Relational *int64 `json:"Relational,omitempty" name:"Relational"`
		// Nosql 数据库数量

		NoSql *int64 `json:"NoSql,omitempty" name:"NoSql"`
		// 分布式数据库数量

		Distributed *int64 `json:"Distributed,omitempty" name:"Distributed"`
		// 自建数据库数量

		SelfBuild *int64 `json:"SelfBuild,omitempty" name:"SelfBuild"`
		// 腾讯云外数据库数量

		Outside *int64 `json:"Outside,omitempty" name:"Outside"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllDbCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllDbCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditSessionInfo struct {

	// 审计日志数

	SqlCount *int64 `json:"SqlCount,omitempty" name:"SqlCount"`
	// 登出时间

	LogoutTime *int64 `json:"LogoutTime,omitempty" name:"LogoutTime"`
	// 客户端 IP

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 客户端端口

	ClientPort *int64 `json:"ClientPort,omitempty" name:"ClientPort"`
	// 数据库 IP

	DbIp *string `json:"DbIp,omitempty" name:"DbIp"`
	// 数据库名称

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 数据库端口

	DbPort *int64 `json:"DbPort,omitempty" name:"DbPort"`
	// 数据库用户

	DbUser *string `json:"DbUser,omitempty" name:"DbUser"`
	// 审计单元 ID

	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`
	// 操作时间(时间)

	OpTime *int64 `json:"OpTime,omitempty" name:"OpTime"`
	// 返回码

	RetNo *int64 `json:"RetNo,omitempty" name:"RetNo"`
	// 会话Id

	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
	// 登录时间

	LoginTime *int64 `json:"LoginTime,omitempty" name:"LoginTime"`
	// 数据资产名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 数据资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
}

type NameValue struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type ModifyCkafkaStartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCkafkaStartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCkafkaStartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashBoardLogRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDashBoardLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashBoardLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcSensitiveDataDetailItem struct {

	// 字段名

	ColName *string `json:"ColName,omitempty" name:"ColName"`
	// 数据摘要

	DataSummary *string `json:"DataSummary,omitempty" name:"DataSummary"`
	// 类型分组

	TypeGroup *string `json:"TypeGroup,omitempty" name:"TypeGroup"`
	// 敏感级别

	DataLevel *uint64 `json:"DataLevel,omitempty" name:"DataLevel"`
	// 敏感字段

	DataType *string `json:"DataType,omitempty" name:"DataType"`
}

type CreateCvmAssetsSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCvmAssetsSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCvmAssetsSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSessionListRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 登陆状态(0 全部 1 成功 2 失败)

	LoginType *int64 `json:"LoginType,omitempty" name:"LoginType"`
	// 数据库端口

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 数据库名称

	DbPort *int64 `json:"DbPort,omitempty" name:"DbPort"`
	// 数据库 IP

	DbIp *string `json:"DbIp,omitempty" name:"DbIp"`
	// 资产 ID

	AssetsId *int64 `json:"AssetsId,omitempty" name:"AssetsId"`
	// 会话 ID

	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
	// 客户端 IP

	ClientSideIp *string `json:"ClientSideIp,omitempty" name:"ClientSideIp"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeSessionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSessionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetsInstResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产信息

		List []*AssetsInst `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetsInstResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsInstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// cvm列表

		List []*CvmInfo `json:"List,omitempty" name:"List"`
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCvmListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCvmListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquireModifycdsAuditSaasResourceRequest struct {
	*tchttp.BaseRequest

	// saas数审实例id  bh-saas-xxxx

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源类型 基础版，高级版，企业版，旗舰版//sv_cds_audit_saas_pg_base，sv_cds_audit_saas_pg_adv，sv_cds_audit_saas_pg_env，sv_cds_audit_saas_pg_ult，

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	//  资产扩展包 sv_cds_audit_saas_exp_asset

	AssetExp *int64 `json:"AssetExp,omitempty" name:"AssetExp"`
	// 存储扩展包  sv_cds_audit_saas_exp_storage

	StorageExp *int64 `json:"StorageExp,omitempty" name:"StorageExp"`
	// 日志投递服务  sv_cds_audit_log_delivery

	DeliveryExp *int64 `json:"DeliveryExp,omitempty" name:"DeliveryExp"`
}

func (r *InquireModifycdsAuditSaasResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquireModifycdsAuditSaasResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRuleSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRuleSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReportMissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReportMissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReportMissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmSmsConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAlarmSmsConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmSmsConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetsInstRequest struct {
	*tchttp.BaseRequest

	// 1页的数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAssetsInstRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsInstRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBehaviourResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作行为分类列表

		List []*AuditBehaviour `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBehaviourResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBehaviourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportTplResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建时间

		AddTime *int64 `json:"AddTime,omitempty" name:"AddTime"`
		// 执行周期

		CntCycle *int64 `json:"CntCycle,omitempty" name:"CntCycle"`
		// 执行日期

		CntDate *string `json:"CntDate,omitempty" name:"CntDate"`
		// 执行日(星期)

		CntDay *int64 `json:"CntDay,omitempty" name:"CntDay"`
		// 执行时间

		CntTime *int64 `json:"CntTime,omitempty" name:"CntTime"`
		// 标识

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 接收者

		Receivers *string `json:"Receivers,omitempty" name:"Receivers"`
		// 发送类型

		SendType *string `json:"SendType,omitempty" name:"SendType"`
		// 模板名称

		TplName *string `json:"TplName,omitempty" name:"TplName"`
		// 报告说明

		TplRemark *string `json:"TplRemark,omitempty" name:"TplRemark"`
		// 最后一次修改时间

		UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReportTplResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportTplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentStartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAgentStartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAgentStartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleSaveRequest struct {
	*tchttp.BaseRequest

	// 数据资产

	AssetsId []*int64 `json:"AssetsId,omitempty" name:"AssetsId"`
	// 告警

	Behaviour []*string `json:"Behaviour,omitempty" name:"Behaviour"`
	// 风险等级

	DangerLevel *int64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则备注

	RuleRemark *string `json:"RuleRemark,omitempty" name:"RuleRemark"`
	// 规则类型

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 规则字段

	FieldList []*RuleFields `json:"FieldList,omitempty" name:"FieldList"`
}

func (r *CreateRuleSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRuleSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExternalAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 腾讯云外 Agent 名单列表数据

		List []*AuditExternalAgent `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExternalAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExternalAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleQuickConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRuleQuickConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRuleQuickConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendAlarmEnterpriseWechatTestRequest struct {
	*tchttp.BaseRequest

	// WebHook 地址

	WebUrl *string `json:"WebUrl,omitempty" name:"WebUrl"`
}

func (r *SendAlarmEnterpriseWechatTestRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendAlarmEnterpriseWechatTestRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnData struct {

	// 签名时间

	SignTime *uint64 `json:"SignTime,omitempty" name:"SignTime"`
	// 签名

	Sign *string `json:"Sign,omitempty" name:"Sign"`
}

type ModifyAlarmSmsConfigSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmSmsConfigSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmSmsConfigSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleSaveRequest struct {
	*tchttp.BaseRequest

	// 规则 Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 数据资产

	AssetsId []*int64 `json:"AssetsId,omitempty" name:"AssetsId"`
	// 告警

	Behaviour []*string `json:"Behaviour,omitempty" name:"Behaviour"`
	// 风险等级

	DangerLevel *int64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则备注

	RuleRemark *string `json:"RuleRemark,omitempty" name:"RuleRemark"`
	// 规则类型

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 规则字段

	FieldList []*RuleFields `json:"FieldList,omitempty" name:"FieldList"`
}

func (r *ModifyRuleSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRuleSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMailConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMailConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMailConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendCkafkaTestResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendCkafkaTestResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendCkafkaTestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTypeConfigListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogTypeConfigListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogTypeConfigListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateRuleSaveRequest struct {
	*tchttp.BaseRequest

	// 告警模版

	AlarmContent *string `json:"AlarmContent,omitempty" name:"AlarmContent"`
	// 风险等级

	DangerLevel *uint64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 是否开启告警

	IsOpen *uint64 `json:"IsOpen,omitempty" name:"IsOpen"`
	// 行为规则Id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 备注

	RuleRemark *string `json:"RuleRemark,omitempty" name:"RuleRemark"`
}

func (r *ModifyOperateRuleSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateRuleSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportMissionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 报表列表

		List []*ReportMission `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReportMissionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportMissionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBackupSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBackupSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcRiskUserTopItem struct {

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 是否有公网访问（1：是；0：否）

	HasPublicNetworkAccess *int64 `json:"HasPublicNetworkAccess,omitempty" name:"HasPublicNetworkAccess"`
	// 访问敏感数据量

	SenstiveAccessCount *uint64 `json:"SenstiveAccessCount,omitempty" name:"SenstiveAccessCount"`
	// 异常行为次数

	AbnormalOperationCount *uint64 `json:"AbnormalOperationCount,omitempty" name:"AbnormalOperationCount"`
}

type DeleteReportRequest struct {
	*tchttp.BaseRequest

	// 报表 Id（此字段废弃，不再支持，请使用Ids）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 报表Id数组，支持批量删除

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroycdsAuditSaasResourceRequest struct {
	*tchttp.BaseRequest

	// 数审实例id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DestroycdsAuditSaasResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroycdsAuditSaasResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmListRequest struct {
	*tchttp.BaseRequest

	// 搜索对象

	FuzzyQuery []*NameValueString `json:"FuzzyQuery,omitempty" name:"FuzzyQuery"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *DescribeCvmListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCvmListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCkafkaSaveRequest struct {
	*tchttp.BaseRequest

	// 接入类型，当前支持 1和7, 类型vip网络类型（1:外网TGW 2:基础网络 3:VPC网络 4:支撑网络(idc 环境) 5:SSL外网访问方式访问 6:黑石环境vpc 7:支撑网络(cvm 环境）

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
	// 实例的地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 实例的id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例的接入信息

	RouteInfo *RouteInfo `json:"RouteInfo,omitempty" name:"RouteInfo"`
	// 接入为域名的时候，有效

	Username *string `json:"Username,omitempty" name:"Username"`
	// 接入为域名的时候，有效

	Password *string `json:"Password,omitempty" name:"Password"`
	// 日志投递的主题配置

	LogDeliveryInfo []*LogDeliveryInfo `json:"LogDeliveryInfo,omitempty" name:"LogDeliveryInfo"`
}

func (r *ModifyCkafkaSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCkafkaSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleQuickConfigRequest struct {
	*tchttp.BaseRequest

	// 规则类别（全部 -1 默认 0 常规 1  宽松 2  严格 3 用户配置4 ）

	ConfigType *int64 `json:"ConfigType,omitempty" name:"ConfigType"`
	// 资产Id列表

	AssetIds []*int64 `json:"AssetIds,omitempty" name:"AssetIds"`
}

func (r *ModifyRuleQuickConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRuleQuickConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogDeliveryTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogDeliveryTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDeliveryTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMailConfigSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMailConfigSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMailConfigSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcAbnormalOperationStatistics struct {

	// 异常时间登录次数

	AbnormalTimeLoginCount *uint64 `json:"AbnormalTimeLoginCount,omitempty" name:"AbnormalTimeLoginCount"`
	// 异常操作次数

	AbnormalOperationCount *uint64 `json:"AbnormalOperationCount,omitempty" name:"AbnormalOperationCount"`
	// 异常ip登录次数

	AbnormalIpLoginCount *uint64 `json:"AbnormalIpLoginCount,omitempty" name:"AbnormalIpLoginCount"`
	// 异常资源访问次数

	AbnormalDataAccessCount *uint64 `json:"AbnormalDataAccessCount,omitempty" name:"AbnormalDataAccessCount"`
}

type TopicInfo struct {

	// ckafka主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// ckafka主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type DescribeAlarmEnterpriseWechatConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAlarmEnterpriseWechatConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmEnterpriseWechatConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgentRequest struct {
	*tchttp.BaseRequest

	// 标识Id

	AgentDeployId *uint64 `json:"AgentDeployId,omitempty" name:"AgentDeployId"`
}

func (r *DeleteAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendAlarmSmsTestResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendAlarmSmsTestResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendAlarmSmsTestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportListRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 报告名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 报告类型

	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`
	// 报告状态

	ReportStatus *int64 `json:"ReportStatus,omitempty" name:"ReportStatus"`
	// 报表模版id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 需要排序的字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 排序顺序 asc desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 时间范围 1:24小时 7:近一周 30:近30天 90:近90天 180:近180天 不变更为0

	CntDay *int64 `json:"CntDay,omitempty" name:"CntDay"`
}

func (r *DescribeReportListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTencentDbResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTencentDbResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTencentDbResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBackupLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBackupSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroycdsAuditSaasResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroycdsAuditSaasResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroycdsAuditSaasResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquireModifycdsAuditSaasResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 询价费用

		ProductCost *ProductCost `json:"ProductCost,omitempty" name:"ProductCost"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquireModifycdsAuditSaasResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquireModifycdsAuditSaasResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCkafkaStartRequest struct {
	*tchttp.BaseRequest

	// 日志类型的主题投递

	LogDeliveryInfo []*LogDeliveryInfo `json:"LogDeliveryInfo,omitempty" name:"LogDeliveryInfo"`
}

func (r *ModifyCkafkaStartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCkafkaStartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAssetsRequest struct {
	*tchttp.BaseRequest

	// 资产ID

	Aid *int64 `json:"Aid,omitempty" name:"Aid"`
}

func (r *DeleteAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmEnterpriseWechatConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmEnterpriseWechatConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmEnterpriseWechatConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTencentDbRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTencentDbRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTencentDbRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquireCreatecdsAuditSaasResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 询价结果

		ProductCost *ProductCost `json:"ProductCost,omitempty" name:"ProductCost"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquireCreatecdsAuditSaasResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquireCreatecdsAuditSaasResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetsPermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetsPermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetsPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户告警时间规则列表

		DbAuditAlarmRuleList []*DbAuditAlarmRule `json:"DbAuditAlarmRuleList,omitempty" name:"DbAuditAlarmRuleList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditInstanceList struct {

	// 总数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 实例列表

	List []*AuditInstance `json:"List,omitempty" name:"List"`
}

type KeyValueItem struct {

	// 键名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateReportPdfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateReportPdfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateReportPdfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		List []*AuditRules `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAgentSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAgentSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashBoardRiskRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDashBoardRiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashBoardRiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogStoreTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据列表

		Data []*TimeValue `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogStoreTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogStoreTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
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

type ModifyAlarmSmsConfigSaveRequest struct {
	*tchttp.BaseRequest

	// 发送间隔

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 下发手机号码

	PhoneNumberSet []*string `json:"PhoneNumberSet,omitempty" name:"PhoneNumberSet"`
	// 短信告警SdkAppId

	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`
	// 短信告警SecretId

	SmsSecretId *string `json:"SmsSecretId,omitempty" name:"SmsSecretId"`
	// 短信告警SecretKey

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// 签名

	Sign *string `json:"Sign,omitempty" name:"Sign"`
	// 地域

	SmsRegion *string `json:"SmsRegion,omitempty" name:"SmsRegion"`
	// 模版 ID

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *ModifyAlarmSmsConfigSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmSmsConfigSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 行为规则列表

		List []*AuditOperateRuleInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentStopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAgentStopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAgentStopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcDataTypeGroup struct {

	// 分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type CreateTimerReportRequest struct {
	*tchttp.BaseRequest

	// 任务名称 不变更为""

	TplName *string `json:"TplName,omitempty" name:"TplName"`
	// 执行日期 重复周期为天：无意义周：星期几1-7月每月几号 1-31

	CntTime *int64 `json:"CntTime,omitempty" name:"CntTime"`
	// 重复周期

	CntCycle *int64 `json:"CntCycle,omitempty" name:"CntCycle"`
	// 发送目标

	Receivers *string `json:"Receivers,omitempty" name:"Receivers"`
	// 时间范围 1:24小时 7:近一周 30:近30天 90:近90天 180:近180天 不变更为0

	CntDay *int64 `json:"CntDay,omitempty" name:"CntDay"`
	// 执行时间 格式15:04 到分钟

	CntDate *string `json:"CntDate,omitempty" name:"CntDate"`
	// 报告说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 模版Id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 报表类型

	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`
	// 关联的资产数组

	AssetsId []*int64 `json:"AssetsId,omitempty" name:"AssetsId"`
	// 报表通知 1关闭 2开启 不变更为0

	Notification *int64 `json:"Notification,omitempty" name:"Notification"`
	// 任务起停 1:关闭 2:开启 单次报表默认为2

	MissionStart *int64 `json:"MissionStart,omitempty" name:"MissionStart"`
}

func (r *CreateTimerReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTimerReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashBoardLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 今日日志-日环比,默认0

		LogDayChain *int64 `json:"LogDayChain,omitempty" name:"LogDayChain"`
		// 日志总数,默认0

		LogTotal *int64 `json:"LogTotal,omitempty" name:"LogTotal"`
		// 今日日志-周环比,默认0

		LogWeekChain *int64 `json:"LogWeekChain,omitempty" name:"LogWeekChain"`
		// 今日风险-日环比,默认0

		RiskDayChain *int64 `json:"RiskDayChain,omitempty" name:"RiskDayChain"`
		// 今日风险-周环比,默认0

		RiskWeekChain *int64 `json:"RiskWeekChain,omitempty" name:"RiskWeekChain"`
		// 今日会话-日环比,默认0

		SessionDayChain *int64 `json:"SessionDayChain,omitempty" name:"SessionDayChain"`
		// 今日会话-周环比,默认0

		SessionWeekChain *int64 `json:"SessionWeekChain,omitempty" name:"SessionWeekChain"`
		// 今日日志总数,默认0

		TodayLog *int64 `json:"TodayLog,omitempty" name:"TodayLog"`
		// 今日风险总数,默认0

		TodayRisk *int64 `json:"TodayRisk,omitempty" name:"TodayRisk"`
		// 今日会话总数,默认0

		TodaySession *int64 `json:"TodaySession,omitempty" name:"TodaySession"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashBoardLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashBoardLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDsgcAbnormalOperationTrendStatistics struct {

	// 统计日期

	Day *string `json:"Day,omitempty" name:"Day"`
	// 统计数据

	StatisticsSet *DsgcAbnormalOperationStatistics `json:"StatisticsSet,omitempty" name:"StatisticsSet"`
}

type DescribeOperateLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 操作日志列表

		List []*AuditOperateLogInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表

		Regions []*Regions `json:"Regions,omitempty" name:"Regions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleSwitchRequest struct {
	*tchttp.BaseRequest

	// 规则Id列表

	RuleId []*int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则启用状态

	RuleStatus *int64 `json:"RuleStatus,omitempty" name:"RuleStatus"`
	// 发件人(仅限 ruleID 为 10000 和 20000)

	RuleReceivers *string `json:"RuleReceivers,omitempty" name:"RuleReceivers"`
	// 资产 ID(除 ruleID 为 10000 和 20000)

	AssetsId *int64 `json:"AssetsId,omitempty" name:"AssetsId"`
}

func (r *ModifyRuleSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRuleSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditRules struct {

	// 风险等级(不包括 ruleId 为 10000 和 20000)

	DangerLevel *int64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 是否为内置规则

	IsInner *int64 `json:"IsInner,omitempty" name:"IsInner"`
	// 规则是否开启(仅限 ruleId 为 10000 和 20000)

	IsOpened *int64 `json:"IsOpened,omitempty" name:"IsOpened"`
	// 规则备注

	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`
	// 规则 ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则类型

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 删除状态

	IsDelete *int64 `json:"IsDelete,omitempty" name:"IsDelete"`
	// 告警相关数据

	AlarmEmail *int64 `json:"AlarmEmail,omitempty" name:"AlarmEmail"`
	// 告警相关数据

	AlarmPhone *int64 `json:"AlarmPhone,omitempty" name:"AlarmPhone"`
	// 告警相关数据

	AlarmSms *int64 `json:"AlarmSms,omitempty" name:"AlarmSms"`
	// 告警相关数据

	AlarmWx *int64 `json:"AlarmWx,omitempty" name:"AlarmWx"`
	// 资产列表

	AssetsList []*AssetsInfo `json:"AssetsList,omitempty" name:"AssetsList"`
	// 审计 ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 自定义规则接收人

	Receivers *string `json:"Receivers,omitempty" name:"Receivers"`
	// 标识是否写数据库

	SaveToDb *int64 `json:"SaveToDb,omitempty" name:"SaveToDb"`
	// 最后一次修改时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 审计名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// AI和CVE规则接收人

	RuleReceiver *string `json:"RuleReceiver,omitempty" name:"RuleReceiver"`
	// 规则启用状态

	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`
	// 规则分类

	RuleClassify *int64 `json:"RuleClassify,omitempty" name:"RuleClassify"`
	// 规则类别（ 默认 0 常规 1  宽松 2  严格 3 用户配置4 ）

	ConfigType *int64 `json:"ConfigType,omitempty" name:"ConfigType"`
}

type DescribeDashBoardRiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险列表(RiskType=1)

		Risks []*NameValue `json:"Risks,omitempty" name:"Risks"`
		// 资产列表(RiskType=2)

		Assets []*NameValue `json:"Assets,omitempty" name:"Assets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashBoardRiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashBoardRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportMission struct {

	// 报表任务id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 任务名称

	TplName *string `json:"TplName,omitempty" name:"TplName"`
	// 报表类型 1:单次报表 2:周期报表

	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`
	// 报告说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 报表模板 1:综合分析报告 2:等保合规报告

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 包含资产

	AssetsList []*AssetsInfo `json:"AssetsList,omitempty" name:"AssetsList"`
	// 下次启动时间

	NextStartTime *int64 `json:"NextStartTime,omitempty" name:"NextStartTime"`
	// 任务状态 1:生成中 2:待生成3:已生成4:生成失败5:已暂停

	MissionStatus *int64 `json:"MissionStatus,omitempty" name:"MissionStatus"`
	// 任务状态说明 仅生成中和生成失败有效

	MissionStatusMessage *string `json:"MissionStatusMessage,omitempty" name:"MissionStatusMessage"`
	// 已生成报表数

	ReportCount *int64 `json:"ReportCount,omitempty" name:"ReportCount"`
	// 任务起停 1:关闭 2:开启 仅周期报表有效

	MissionStart *int64 `json:"MissionStart,omitempty" name:"MissionStart"`
	// 统计周期 1:24小时 7:近一周 30:近30天 90:近90天 180:

	CntDay *int64 `json:"CntDay,omitempty" name:"CntDay"`
	// 重复周期 1:每天 2:每周 3:每月

	CntCycle *uint64 `json:"CntCycle,omitempty" name:"CntCycle"`
	// 执行日期 重复周期为天：无意义 周：星期几 1-7  月每月

	CntTime *uint64 `json:"CntTime,omitempty" name:"CntTime"`
	// 执行时间 格式15:04 到分钟

	CntDate *string `json:"CntDate,omitempty" name:"CntDate"`
	// 创建者 0:内置 其余存放用户(uin)

	Receivers *string `json:"Receivers,omitempty" name:"Receivers"`
	// Notification  int  1关闭 2开启 不变更为0

	Notification *int64 `json:"Notification,omitempty" name:"Notification"`
}

type DescribeNewReportTplRequest struct {
	*tchttp.BaseRequest

	// 报表id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeNewReportTplRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewReportTplRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionListsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 报表预览地址

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 到期时间

		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// 资产扩展包数量

		ExpandNum *int64 `json:"ExpandNum,omitempty" name:"ExpandNum"`
		// 是否存在传统型

		HasOldVersion *bool `json:"HasOldVersion,omitempty" name:"HasOldVersion"`
		// 状态（0 未开通 1正常，2隔离，3销毁）

		OpenState *int64 `json:"OpenState,omitempty" name:"OpenState"`
		// 资源 Id

		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
		// 版本编码

		VersionCode *string `json:"VersionCode,omitempty" name:"VersionCode"`
		// 存储扩展数

		StorageExpandNum *int64 `json:"StorageExpandNum,omitempty" name:"StorageExpandNum"`
		// 日志投递开通状态（0:未开通，1:已开通）

		SvCdsAuditLogDelivery *int64 `json:"SvCdsAuditLogDelivery,omitempty" name:"SvCdsAuditLogDelivery"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentUninstallRequest struct {
	*tchttp.BaseRequest

	// 标识Id

	AgentDeployId *uint64 `json:"AgentDeployId,omitempty" name:"AgentDeployId"`
}

func (r *ModifyAgentUninstallRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAgentUninstallRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcSensitiveDataInfoSet struct {

	// 资产类型

	PropertyType *string `json:"PropertyType,omitempty" name:"PropertyType"`
	// 数据库名称

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 数据表名称

	TableName *string `json:"TableName,omitempty" name:"TableName"`
	// 数据级别

	DataLevel *uint64 `json:"DataLevel,omitempty" name:"DataLevel"`
	// 数据类型集合

	DataTypes *string `json:"DataTypes,omitempty" name:"DataTypes"`
	// 数据级别集合，以逗号分隔

	DataLevels *string `json:"DataLevels,omitempty" name:"DataLevels"`
}

type DescribeBackupSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份日志保留时长

		BackupLogSaveTime *int64 `json:"BackupLogSaveTime,omitempty" name:"BackupLogSaveTime"`
		// 恢复日志保留时长

		RestoreLogSaveTime *int64 `json:"RestoreLogSaveTime,omitempty" name:"RestoreLogSaveTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerAdditionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCustomerAdditionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerAdditionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFieldConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFieldConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFieldConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNewReportTplResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加时间

		AddTime *int64 `json:"AddTime,omitempty" name:"AddTime"`
		// 执行周期

		CntCycle *int64 `json:"CntCycle,omitempty" name:"CntCycle"`
		// 执行日期

		CntDate *string `json:"CntDate,omitempty" name:"CntDate"`
		// 执行日(星期)

		CntDay *int64 `json:"CntDay,omitempty" name:"CntDay"`
		// 执行时间

		CntTime *int64 `json:"CntTime,omitempty" name:"CntTime"`
		// 标识

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 接收者

		Receivers *string `json:"Receivers,omitempty" name:"Receivers"`
		// 发送类型

		SendType *string `json:"SendType,omitempty" name:"SendType"`
		// 模板名称

		TplName *string `json:"TplName,omitempty" name:"TplName"`
		// 报告说明

		TplRemark *string `json:"TplRemark,omitempty" name:"TplRemark"`
		// 最后一次修改时间

		UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNewReportTplResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewReportTplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgentDownloadRequest struct {
	*tchttp.BaseRequest

	// 区分腾讯云内网还是腾讯云外(0 腾讯云内网 1 腾讯云外)

	TrafficType *uint64 `json:"TrafficType,omitempty" name:"TrafficType"`
	// 操作系统类型(0 linux 1 windows)

	SystemType *uint64 `json:"SystemType,omitempty" name:"SystemType"`
}

func (r *CreateAgentDownloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgentDownloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateRuleListRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 风险等级

	DangerLevel *int64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 行为 ID

	OperateIds []*uint64 `json:"OperateIds,omitempty" name:"OperateIds"`
}

func (r *DescribeOperateRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleInfoFields struct {

	// 浮点型数值

	DoubleValue *int64 `json:"DoubleValue,omitempty" name:"DoubleValue"`
	// 字段名称

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 字段类型

	FieldType *string `json:"FieldType,omitempty" name:"FieldType"`
	// 标识id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 整型数值

	IntValue *int64 `json:"IntValue,omitempty" name:"IntValue"`
	// 是否已删除

	IsDelete *int64 `json:"IsDelete,omitempty" name:"IsDelete"`
	// 逻辑

	Logic *string `json:"Logic,omitempty" name:"Logic"`
	// 规则ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 字符型数值

	StringValue *string `json:"StringValue,omitempty" name:"StringValue"`
	// 最后一次修改时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeAssetsTypeRequest struct {
	*tchttp.BaseRequest

	// 数据资产类型

	AssetsAddType *int64 `json:"AssetsAddType,omitempty" name:"AssetsAddType"`
}

func (r *DescribeAssetsTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesListRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 规则类型

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 规则属性

	IsInner *int64 `json:"IsInner,omitempty" name:"IsInner"`
	// 风险等级

	DangerLevel *int64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 规则分类

	RuleClassify *int64 `json:"RuleClassify,omitempty" name:"RuleClassify"`
	// 规则类别（全部 -1 默认 0 常规 1  宽松 2  严格 3 用户配置4 ）

	ConfigType *int64 `json:"ConfigType,omitempty" name:"ConfigType"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *DescribeRulesListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRulesListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcPrivateLinkRequest struct {
	*tchttp.BaseRequest

	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateVpcPrivateLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcPrivateLinkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmRuleRequest struct {
	*tchttp.BaseRequest

	// 告警类型，取值：sql_risk_high, sql_risk_middle, sql_risk_low, operation_risk_high, operation_risk_middle, operation_risk_low, qps_limit, storage_capacity_limit, agent_offline, log_delivery_abnormal, report_notification

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
	// 允许告警起始时间-小时，取值 0-23

	AllowedStartHour *int64 `json:"AllowedStartHour,omitempty" name:"AllowedStartHour"`
	// 允许告警结束时间-小时，取值 0-23

	AllowedEndHour *int64 `json:"AllowedEndHour,omitempty" name:"AllowedEndHour"`
	// 允许告警起始时间-分钟，取值 0-59

	AllowedStartMinute *int64 `json:"AllowedStartMinute,omitempty" name:"AllowedStartMinute"`
	// 允许告警结束时间-分钟，取值 0-59

	AllowedEndMinute *int64 `json:"AllowedEndMinute,omitempty" name:"AllowedEndMinute"`
	// 该告警类型的告警开关，取值为0或1，0代表关闭，1代表开启

	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`
}

func (r *ModifyAlarmRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendAlarmSmsTestRequest struct {
	*tchttp.BaseRequest

	// 发送间隔

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 下发手机号码

	PhoneNumberSet []*string `json:"PhoneNumberSet,omitempty" name:"PhoneNumberSet"`
	// 短信告警SdkAppId

	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`
	// 短信告警SecretId

	SmsSecretId *string `json:"SmsSecretId,omitempty" name:"SmsSecretId"`
	// 短信告警SecretKey

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// 签名

	Sign *string `json:"Sign,omitempty" name:"Sign"`
	// 地域

	SmsRegion *string `json:"SmsRegion,omitempty" name:"SmsRegion"`
	// 模版 ID

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *SendAlarmSmsTestRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendAlarmSmsTestRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetsSaveRequest struct {
	*tchttp.BaseRequest

	// 数据资产名称

	AssetsName *string `json:"AssetsName,omitempty" name:"AssetsName"`
	// 数据资产类型

	AssetsType *string `json:"AssetsType,omitempty" name:"AssetsType"`
	// 数据资产类型版本

	AssetsVersion *string `json:"AssetsVersion,omitempty" name:"AssetsVersion"`
	// IP/域名

	AssetsIp *string `json:"AssetsIp,omitempty" name:"AssetsIp"`
	// 端口

	AssetsPort *int64 `json:"AssetsPort,omitempty" name:"AssetsPort"`
	// 是否需要安全证书

	NeedPem *bool `json:"NeedPem,omitempty" name:"NeedPem"`
	// 安全证书内容

	DBSecretPem *string `json:"DBSecretPem,omitempty" name:"DBSecretPem"`
	// 安全证书密码

	PemSecretKey *string `json:"PemSecretKey,omitempty" name:"PemSecretKey"`
}

func (r *CreateAssetsSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetsSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogDeliveryTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogDeliveryTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogDeliveryTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogDeliveryTypeRequest struct {
	*tchttp.BaseRequest

	// 日志类型

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 投递的topicid

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递的topicname

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *ModifyLogDeliveryTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogDeliveryTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportMissionListRequest struct {
	*tchttp.BaseRequest

	// 报表名 可模糊查询

	TplName *string `json:"TplName,omitempty" name:"TplName"`
	// 报表类型 1:单次报表 2:周期报表 0全查

	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`
	// 报表模板 1:综合分析报告 2:等保合规报告 0全查

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 任务状态0全查 1:生成中 2:待生成 3:已生成 4:生成失败 5:已暂停

	MissionStatus *int64 `json:"MissionStatus,omitempty" name:"MissionStatus"`
	// 排序字段 支持“NextStartTime” 与 “MissionStatus”

	Field *string `json:"Field,omitempty" name:"Field"`
	// ‘desc' | 'asc'

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeReportMissionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportMissionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublickeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公钥

		PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublickeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublickeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFieldConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 字段类型

		FieldTypes *RuleConfigFieldTypes `json:"FieldTypes,omitempty" name:"FieldTypes"`
		// 字段信息

		Fields *RuleConfigFields `json:"Fields,omitempty" name:"Fields"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFieldConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFieldConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetsForm struct {

	// 数据资产名称

	AssetsName *string `json:"AssetsName,omitempty" name:"AssetsName"`
	// 数据资产类型

	AssetsType *string `json:"AssetsType,omitempty" name:"AssetsType"`
	// 数据资产类型版本

	AssetsVersion *string `json:"AssetsVersion,omitempty" name:"AssetsVersion"`
	// IP/域名

	AssetsIp *string `json:"AssetsIp,omitempty" name:"AssetsIp"`
	// 端口

	AssetsPort *int64 `json:"AssetsPort,omitempty" name:"AssetsPort"`
}

type DbAuditAlarmRule struct {

	// 告警规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则所属用户的AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 告警类型，string字符串，取值：sql_risk_high, sql_risk_middle, sql_risk_low, operation_risk_high, operation_risk_middle, operation_risk_low, qps_limit, storage_capacity_limit, agent_offline

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
	// 告警规则开关，0关闭1开启

	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`
	// 允许告警的起始时间-小时，取值0-23

	AllowedStartHour *int64 `json:"AllowedStartHour,omitempty" name:"AllowedStartHour"`
	// 允许告警的结束时间-小时，取值0-23

	AllowedEndHour *int64 `json:"AllowedEndHour,omitempty" name:"AllowedEndHour"`
	// 允许告警的起始时间-分钟，取值0-59

	AllowedStartMinute *int64 `json:"AllowedStartMinute,omitempty" name:"AllowedStartMinute"`
	// 允许告警的结束时间-分钟，取值0-59

	AllowedEndMinute *int64 `json:"AllowedEndMinute,omitempty" name:"AllowedEndMinute"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type DescribeRuleSwitchListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 规则列表数据

		List []*AuditRuleSwitch `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleSwitchListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleSwitchListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCkafkaInstanceListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserCkafkaInstanceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCkafkaInstanceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CdsAuditInstance struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 项目ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 续费标识

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 付费模式（数据安全审计只支持预付费：1）

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 实例状态： 0，未生效；1：正常运行； 2：被隔离； 3，已过期

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 实例被隔离时间，格式：yyyy-mm-dd HH:ii:ss

	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitempty" name:"IsolatedTimestamp"`
	// 实例创建时间，格式： yyyy-mm-dd HH:ii:ss

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例过期时间，格式：yyyy-mm-dd HH:ii:ss

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 实例私网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 实例类型（版本）

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例域名

	Pdomain *string `json:"Pdomain,omitempty" name:"Pdomain"`
}

type DsgcDataLevelItem struct {

	// ID，0为新增，非0表示修改

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 敏感数据类型

	SensitiveType *string `json:"SensitiveType,omitempty" name:"SensitiveType"`
	// 正则表达式

	RegularExpression *string `json:"RegularExpression,omitempty" name:"RegularExpression"`
	// 敏感数据分级 10-低 20-中 30-高

	SensitiveLevel *uint64 `json:"SensitiveLevel,omitempty" name:"SensitiveLevel"`
	// 是否启用（0否；1是），只针对预定义有效

	IsEnabled *uint64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 1-预定义；2-用户自定义

	DefineType *uint64 `json:"DefineType,omitempty" name:"DefineType"`
	// 类型标签，以 ; （英文半角分号）分隔多个

	DataTypeGroups *string `json:"DataTypeGroups,omitempty" name:"DataTypeGroups"`
}

type Services struct {

	// 服务id

	Sid *uint64 `json:"Sid,omitempty" name:"Sid"`
	// 服务名称

	SysName *string `json:"SysName,omitempty" name:"SysName"`
}

type DescribeOperateLogListRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 操作 IP

	OpIp *string `json:"OpIp,omitempty" name:"OpIp"`
	// 操作行为

	OpType *uint64 `json:"OpType,omitempty" name:"OpType"`
	// 操作账户

	OpUser *string `json:"OpUser,omitempty" name:"OpUser"`
	// 排列顺序

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 排序字段

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *DescribeOperateLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgentDownloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载地址

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAgentDownloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgentDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetsVersionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetsVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmEnterpriseWechatConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// WebHook 地址

		WebUrl *string `json:"WebUrl,omitempty" name:"WebUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmEnterpriseWechatConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmEnterpriseWechatConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetsInst struct {

	// 数据资产Id

	AssetsKey *int64 `json:"AssetsKey,omitempty" name:"AssetsKey"`
	// 审计单元Id

	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产IP

	InstanceIp *string `json:"InstanceIp,omitempty" name:"InstanceIp"`
	// 资产名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 端口

	InstancePort *int64 `json:"InstancePort,omitempty" name:"InstancePort"`
	// 资产类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// DDL操作数量

	DdlCount *int64 `json:"DdlCount,omitempty" name:"DdlCount"`
	// 失败会话数量

	ErrorCount *int64 `json:"ErrorCount,omitempty" name:"ErrorCount"`
	// 会话总数

	InstCount *int64 `json:"InstCount,omitempty" name:"InstCount"`
	// 最大并发会话

	MaxSession *int64 `json:"MaxSession,omitempty" name:"MaxSession"`
	// 风险数量

	WarnCount *int64 `json:"WarnCount,omitempty" name:"WarnCount"`
	// 日志数量

	LogAmount *int64 `json:"LogAmount,omitempty" name:"LogAmount"`
}

type CreateWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单数组

	WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
}

func (r *CreateWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentInstanceInfo struct {

	// 内网IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 系统类型

	SystemType *uint64 `json:"SystemType,omitempty" name:"SystemType"`
}

type DsgcAlertDetails struct {

	// 告警ID

	AlertId *uint64 `json:"AlertId,omitempty" name:"AlertId"`
	// 告警类型

	AlertType *string `json:"AlertType,omitempty" name:"AlertType"`
	// 告警时间

	AlertTime *string `json:"AlertTime,omitempty" name:"AlertTime"`
	// 涉及到的资源

	AssociateResource *string `json:"AssociateResource,omitempty" name:"AssociateResource"`
	// 告警状态：0-未处理 1-已处理 2-已忽略

	AlertStatus *int64 `json:"AlertStatus,omitempty" name:"AlertStatus"`
	// 告警的详情

	AlertDetails *string `json:"AlertDetails,omitempty" name:"AlertDetails"`
	// 风险级别：10-低 20-中 30-高

	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 风险描述

	RiskDesc *string `json:"RiskDesc,omitempty" name:"RiskDesc"`
	// 建议应对措施

	HandleAdvice *string `json:"HandleAdvice,omitempty" name:"HandleAdvice"`
}

type RuleFields struct {

	// 浮点型数值

	DoubleValue *int64 `json:"DoubleValue,omitempty" name:"DoubleValue"`
	// 字段名称

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 字段类型

	FieldType *string `json:"FieldType,omitempty" name:"FieldType"`
	// 标识 id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 整型数值

	IntValue *int64 `json:"IntValue,omitempty" name:"IntValue"`
	// 是否已删除

	IsDelete *int64 `json:"IsDelete,omitempty" name:"IsDelete"`
	// 逻辑

	Logic *string `json:"Logic,omitempty" name:"Logic"`
	// 规则 ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 字符型数值

	StringValue *string `json:"StringValue,omitempty" name:"StringValue"`
	// 最后一次修改时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CreateProduceReportRequest struct {
	*tchttp.BaseRequest

	// 报告名称

	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 报告说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 发送目标

	Receivers *string `json:"Receivers,omitempty" name:"Receivers"`
	// 审计单元 ID

	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CreateProduceReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProduceReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetsPermissionRequest struct {
	*tchttp.BaseRequest

	// 资产 ID

	Aid *int64 `json:"Aid,omitempty" name:"Aid"`
	// 审计权限

	Permission *int64 `json:"Permission,omitempty" name:"Permission"`
}

func (r *ModifyAssetsPermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetsPermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcAbnormalOperationInstanceStatistics struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 异常统计元素

	StatisticsSet *DsgcAbnormalOperationStatistics `json:"StatisticsSet,omitempty" name:"StatisticsSet"`
}

type DeleteRulesRequest struct {
	*tchttp.BaseRequest

	// 规则ID

	RuleId []*int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HitRules struct {

	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type DescribeMailConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 端口

		Port *uint64 `json:"Port,omitempty" name:"Port"`
		// 收件人

		Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
		// 发件人

		Sender *string `json:"Sender,omitempty" name:"Sender"`
		// SMTP 服务器

		SmtpHost *string `json:"SmtpHost,omitempty" name:"SmtpHost"`
		// 安全传输层协议

		Tls *uint64 `json:"Tls,omitempty" name:"Tls"`
		// 邮件账号

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMailConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMailConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTypeConfigListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户ckafka配置，以及日志投递配置

		List []*LogDeliveryCkafkaConfig `json:"List,omitempty" name:"List"`
		// 配置总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogTypeConfigListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogTypeConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublickeyRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePublickeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublickeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditOperateRuleInfo struct {

	// 创建时间

	AddTime *uint64 `json:"AddTime,omitempty" name:"AddTime"`
	// 告警模版

	AlarmContent *string `json:"AlarmContent,omitempty" name:"AlarmContent"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 风险等级

	DangerLevel *uint64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 修改人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// 是否删除

	IsDelete *uint64 `json:"IsDelete,omitempty" name:"IsDelete"`
	// 是否开启告警

	IsOpen *uint64 `json:"IsOpen,omitempty" name:"IsOpen"`
	// 操作Id

	OperateId *uint64 `json:"OperateId,omitempty" name:"OperateId"`
	// 操作名称

	OperateName *string `json:"OperateName,omitempty" name:"OperateName"`
	// 操作action

	OperateUrl *string `json:"OperateUrl,omitempty" name:"OperateUrl"`
	// 父操作名称

	ParentOperateName *string `json:"ParentOperateName,omitempty" name:"ParentOperateName"`
	// 行为规则Id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 备注

	RuleRemark *string `json:"RuleRemark,omitempty" name:"RuleRemark"`
	// 修改时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DasbCvmConfig struct {

	// 主机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主机核数

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 外网带宽

	NetBand *uint64 `json:"NetBand,omitempty" name:"NetBand"`
	// 系统盘大小

	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`
	// 数据盘大小

	DataDiskSize *uint64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`
	// 购买月份

	MonthSpan *uint64 `json:"MonthSpan,omitempty" name:"MonthSpan"`
	// 所属商品模型ID

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名

	Region *string `json:"Region,omitempty" name:"Region"`
	// 使用的镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

type DsgcAbnormalUserInfos struct {

	// 用户名

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 综合评分

	Score *float64 `json:"Score,omitempty" name:"Score"`
	// 学习进度

	LearningProgress *float64 `json:"LearningProgress,omitempty" name:"LearningProgress"`
	// 异常情况序列 - 异常操作

	AbnormalSequenceOperation *string `json:"AbnormalSequenceOperation,omitempty" name:"AbnormalSequenceOperation"`
	// 异常情况序列 - 异常资源访问

	AbnormalSequenceAccess *string `json:"AbnormalSequenceAccess,omitempty" name:"AbnormalSequenceAccess"`
	// 异常情况序列 - 异常ip登录

	AbnormalSequenceIp *string `json:"AbnormalSequenceIp,omitempty" name:"AbnormalSequenceIp"`
	// 异常情况序列 - 异常时间登录

	AbnormalSequenceTime *string `json:"AbnormalSequenceTime,omitempty" name:"AbnormalSequenceTime"`
}

type DeleteAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcSensitiveType struct {

	// 级别

	SensitiveLevel *uint64 `json:"SensitiveLevel,omitempty" name:"SensitiveLevel"`
	// 类型名称

	SensitiveType *string `json:"SensitiveType,omitempty" name:"SensitiveType"`
	// 类型分组

	SensitiveGroup *string `json:"SensitiveGroup,omitempty" name:"SensitiveGroup"`
}

type CreateNewProduceReportRequest struct {
	*tchttp.BaseRequest

	// 报表名称

	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 接受者

	Receivers *string `json:"Receivers,omitempty" name:"Receivers"`
	// 审计资产id

	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CreateNewProduceReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNewProduceReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRuleSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRuleSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogDownloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载地址

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogDownloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetsListRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 实例Id/实例名称/资产名称

	SearchValues []*NameValueString `json:"SearchValues,omitempty" name:"SearchValues"`
	// 数据资产类型

	AssetsType *string `json:"AssetsType,omitempty" name:"AssetsType"`
	// 查询的资产类型（1:cdb、2:cvm、3:others

	AssetsAddType *int64 `json:"AssetsAddType,omitempty" name:"AssetsAddType"`
	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 审计权限

	Permission *int64 `json:"Permission,omitempty" name:"Permission"`
	// 状态

	AliveStatus *int64 `json:"AliveStatus,omitempty" name:"AliveStatus"`
}

func (r *DescribeAssetsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaRouteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Ckafka实例的路由列表

		List []*RouteInfo `json:"List,omitempty" name:"List"`
		// Ckafka实例的路由列表个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCkafkaRouteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaRouteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总资产数，默认0

		MaxInstance *int64 `json:"MaxInstance,omitempty" name:"MaxInstance"`
		// 版本名称,默认''

		VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
		// 已授权的资产数，默认0

		AuthedInstance *int64 `json:"AuthedInstance,omitempty" name:"AuthedInstance"`
		// 存储空间总容量，单位G

		AllStorage *int64 `json:"AllStorage,omitempty" name:"AllStorage"`
		// 已使用的存储空间容量，单位G

		UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`
		// 过期时间，默认为0

		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// 添加时间，默认为0

		AddTime *int64 `json:"AddTime,omitempty" name:"AddTime"`
		// 服务器时间，默认为0

		ServerTime *int64 `json:"ServerTime,omitempty" name:"ServerTime"`
		// 总资产数

		UserInstanceCount *int64 `json:"UserInstanceCount,omitempty" name:"UserInstanceCount"`
		// 未开通资产数

		UnAuthedInstanceCount *int64 `json:"UnAuthedInstanceCount,omitempty" name:"UnAuthedInstanceCount"`
		// 已使用在线日志空间

		OnlineStorage *int64 `json:"OnlineStorage,omitempty" name:"OnlineStorage"`
		// 总在线日志空间

		SpecificationOnlineStorage *int64 `json:"SpecificationOnlineStorage,omitempty" name:"SpecificationOnlineStorage"`
		// 备份日志大小

		BackupStorage *int64 `json:"BackupStorage,omitempty" name:"BackupStorage"`
		// 备份日志大小总空间

		SpecificationBackupStorage *int64 `json:"SpecificationBackupStorage,omitempty" name:"SpecificationBackupStorage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcAccessRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVpcAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {

	// 实例id

	Psid *uint64 `json:"Psid,omitempty" name:"Psid"`
	// 实例名称

	PsysName *string `json:"PsysName,omitempty" name:"PsysName"`
	// 公网ip

	MpublicIp *string `json:"MpublicIp,omitempty" name:"MpublicIp"`
	// 内网ip

	MprivateIp *string `json:"MprivateIp,omitempty" name:"MprivateIp"`
	// 产品型号

	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`
	// 购买时间

	PaddTime *uint64 `json:"PaddTime,omitempty" name:"PaddTime"`
	// 到期时间

	PendTime *uint64 `json:"PendTime,omitempty" name:"PendTime"`
	// 域名

	Pdomain *string `json:"Pdomain,omitempty" name:"Pdomain"`
	// 地区id

	Pregion *uint64 `json:"Pregion,omitempty" name:"Pregion"`
	// 网络id

	PuniqVpcId *string `json:"PuniqVpcId,omitempty" name:"PuniqVpcId"`
	// 子网id

	PuniqSubnetId *string `json:"PuniqSubnetId,omitempty" name:"PuniqSubnetId"`
	// 系统状态

	PsysStatus *uint64 `json:"PsysStatus,omitempty" name:"PsysStatus"`
	// 网络类型

	PnetType *uint64 `json:"PnetType,omitempty" name:"PnetType"`
}

type DescribeRiskListRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式(desc=倒叙,asc=升序)

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 排序字段(opTime=时间,dangerLvl=风险等级)

	Field *string `json:"Field,omitempty" name:"Field"`
	// 风险等级(1-低风险,2-中风险,3-高风险,不传全部)

	DangerLevel *string `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 数据库名称

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 数据库端口

	DbPort *int64 `json:"DbPort,omitempty" name:"DbPort"`
	// 数据库 IP

	DbIp *string `json:"DbIp,omitempty" name:"DbIp"`
	// 资产 ID

	AssetsId *int64 `json:"AssetsId,omitempty" name:"AssetsId"`
	// 会话 ID

	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
	// 客户端 IP

	ClientSideIp *string `json:"ClientSideIp,omitempty" name:"ClientSideIp"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 命中规则

	HitRule *int64 `json:"HitRule,omitempty" name:"HitRule"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *DescribeRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetIdName struct {

	// 资产Id

	Aid *int64 `json:"Aid,omitempty" name:"Aid"`
	// 资产名称

	AssetsName *string `json:"AssetsName,omitempty" name:"AssetsName"`
}

type Regions struct {

	// 中心地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 中心地域的id

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域总称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 支持的地域子集

	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`
	// 地区描述

	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeStoreTopLimitRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeStoreTopLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStoreTopLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentSaveRequest struct {
	*tchttp.BaseRequest

	// 标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 资产 ID数组

	AssetsId []*string `json:"AssetsId,omitempty" name:"AssetsId"`
	// CPU 利用率

	CpuUse *uint64 `json:"CpuUse,omitempty" name:"CpuUse"`
	// 超过阈值停止审计

	OverloadCheck *uint64 `json:"OverloadCheck,omitempty" name:"OverloadCheck"`
	// 内存利用率

	RamUse *uint64 `json:"RamUse,omitempty" name:"RamUse"`
}

func (r *ModifyAgentSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAgentSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditAgentInfo struct {

	// 部署 IP

	DeployIp *string `json:"DeployIp,omitempty" name:"DeployIp"`
	// 流量来源

	TrafficType *uint64 `json:"TrafficType,omitempty" name:"TrafficType"`
	// VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Agent部署的系统类型

	SystemType *uint64 `json:"SystemType,omitempty" name:"SystemType"`
	// 内存利用率

	MemoryUsage *uint64 `json:"MemoryUsage,omitempty" name:"MemoryUsage"`
	// CPU 利用率

	CpuUsage *uint64 `json:"CpuUsage,omitempty" name:"CpuUsage"`
	// 运行状态

	AgentStatus *uint64 `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// 创建时间

	AddTime *uint64 `json:"AddTime,omitempty" name:"AddTime"`
	// Agent的Id

	AgentId *uint64 `json:"AgentId,omitempty" name:"AgentId"`
	// Agent 名称

	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`
	// Agent进程id

	AgentPid *uint64 `json:"AgentPid,omitempty" name:"AgentPid"`
	// Agent心跳运行时间

	AgentRuntimeTime *uint64 `json:"AgentRuntimeTime,omitempty" name:"AgentRuntimeTime"`
	// agent手动关闭时间

	AgentShutdownTime *uint64 `json:"AgentShutdownTime,omitempty" name:"AgentShutdownTime"`
	// agent的修改状态标识

	AgentUpdate *uint64 `json:"AgentUpdate,omitempty" name:"AgentUpdate"`
	// 资产 ID

	AssetsIds *string `json:"AssetsIds,omitempty" name:"AssetsIds"`
	// 资产列表

	AssetsList []*AssetsInfo `json:"AssetsList,omitempty" name:"AssetsList"`
	// agent是否开启负载均衡

	CheckOn *uint64 `json:"CheckOn,omitempty" name:"CheckOn"`
	// 部署错误信息

	DeployErrorMsg *string `json:"DeployErrorMsg,omitempty" name:"DeployErrorMsg"`
	// 部署 Mac

	DeployMac *string `json:"DeployMac,omitempty" name:"DeployMac"`
	// 部署路径

	DeployPath *string `json:"DeployPath,omitempty" name:"DeployPath"`
	// 部署状态

	DeployStatus *uint64 `json:"DeployStatus,omitempty" name:"DeployStatus"`
	// 部署类型

	DeployType *uint64 `json:"DeployType,omitempty" name:"DeployType"`
	// 心跳时间

	HeartbeatTime *uint64 `json:"HeartbeatTime,omitempty" name:"HeartbeatTime"`
	// 标识

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// agent部署任务id

	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
	// 离线时间

	OfflineTime *uint64 `json:"OfflineTime,omitempty" name:"OfflineTime"`
	// 服务器 IP

	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`
	// 服务器端口

	ServerPort *uint64 `json:"ServerPort,omitempty" name:"ServerPort"`
	// Agent负载CPU利用率

	AgentCpuUsage *int64 `json:"AgentCpuUsage,omitempty" name:"AgentCpuUsage"`
	// Agent负载内存利用率

	AgentMemoryUsage *int64 `json:"AgentMemoryUsage,omitempty" name:"AgentMemoryUsage"`
	// 机器负载CPU利用率

	MachineCpuUsage *int64 `json:"MachineCpuUsage,omitempty" name:"MachineCpuUsage"`
	// 机器负载内存利用率

	MachineMemoryUsage *int64 `json:"MachineMemoryUsage,omitempty" name:"MachineMemoryUsage"`
}

type DescribeOldAlarmStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户老的告警流程开启状态，0关闭1开启

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOldAlarmStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOldAlarmStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCkafkaInstanceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户ckafka列表

		List []*CkafkaInstance `json:"List,omitempty" name:"List"`
		// 返回租户ckafka数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserCkafkaInstanceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCkafkaInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCkafkaSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCkafkaSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCkafkaSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetsVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// MYSQL 版本列表

		Mysql []*string `json:"Mysql,omitempty" name:"Mysql"`
		// MARIADB 版本列表

		MariaDb []*string `json:"MariaDb,omitempty" name:"MariaDb"`
		// SQLSERVER 版本列表

		SqlServer []*string `json:"SqlServer,omitempty" name:"SqlServer"`
		// POSTGRESQL 版本列表

		PostGreSql []*string `json:"PostGreSql,omitempty" name:"PostGreSql"`
		// ORACLE 版本列表

		Oracle []*string `json:"Oracle,omitempty" name:"Oracle"`
		// REDIS 版本列表

		Redis []*string `json:"Redis,omitempty" name:"Redis"`
		// MONGODB 版本列表

		MongoDb []*string `json:"MongoDb,omitempty" name:"MongoDb"`
		// HIVE 版本列表

		Hive []*string `json:"Hive,omitempty" name:"Hive"`
		// HBASE 版本列表

		Hbase []*string `json:"Hbase,omitempty" name:"Hbase"`
		// TdSql 版本列表

		TDSql []*string `json:"TDSql,omitempty" name:"TDSql"`
		// tdSql-C Mysql 版本列表

		TDSqlMysql []*string `json:"TDSqlMysql,omitempty" name:"TDSqlMysql"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetsVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 规则 ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSessionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 审计会话列表信息

		List []*AuditSessionInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSessionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSessionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentStartRequest struct {
	*tchttp.BaseRequest

	// 标识Id

	AgentDeployId *uint64 `json:"AgentDeployId,omitempty" name:"AgentDeployId"`
}

func (r *ModifyAgentStartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAgentStartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogDownloadRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 日志类型（0 审计日志 1 审计风险 2 审计会话）

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyLogDownloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogDownloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NameValueString struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeAlarmSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAlarmSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogExportJobInfo struct {

	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 修改人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// 导出文件的文件名

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 日志总数

	FinishCount *uint64 `json:"FinishCount,omitempty" name:"FinishCount"`
	// 是否删除

	IsDelete *uint64 `json:"IsDelete,omitempty" name:"IsDelete"`
	// 日志导出任务ID

	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
	// 导出状态

	JobStatus *uint64 `json:"JobStatus,omitempty" name:"JobStatus"`
	// 导出进度

	JobSpeed *uint64 `json:"JobSpeed,omitempty" name:"JobSpeed"`
	// 返回日志数量

	LogCount *uint64 `json:"LogCount,omitempty" name:"LogCount"`
	// 日志数据

	PostData *string `json:"PostData,omitempty" name:"PostData"`
	// 查询日志结束id

	SearchIdRangeEnd *string `json:"SearchIdRangeEnd,omitempty" name:"SearchIdRangeEnd"`
	// 查询日志的起始id

	SearchIdRangeStart *string `json:"SearchIdRangeStart,omitempty" name:"SearchIdRangeStart"`
	// 查询时间

	SearchTime *uint64 `json:"SearchTime,omitempty" name:"SearchTime"`
	// 更新时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	AddTime *uint64 `json:"AddTime,omitempty" name:"AddTime"`
}

type DeleteRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifycdsAuditSaasResourceRequest struct {
	*tchttp.BaseRequest

	// saas数审实例id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源类型 基础版，高级版，企业版，旗舰版

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	// 资产扩展包

	AssetExp *int64 `json:"AssetExp,omitempty" name:"AssetExp"`
	// 存储扩展包

	StorageExp *int64 `json:"StorageExp,omitempty" name:"StorageExp"`
	// 日志投递服务

	DeliveryExp *int64 `json:"DeliveryExp,omitempty" name:"DeliveryExp"`
}

func (r *ModifycdsAuditSaasResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifycdsAuditSaasResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashBoardAnalysisRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDashBoardAnalysisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashBoardAnalysisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMailConfigSaveRequest struct {
	*tchttp.BaseRequest

	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 收件人

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 发件人

	Sender *string `json:"Sender,omitempty" name:"Sender"`
	// SMTP 服务器

	SmtpHost *string `json:"SmtpHost,omitempty" name:"SmtpHost"`
	// 安全传输层协议

	Tls *uint64 `json:"Tls,omitempty" name:"Tls"`
	// 邮件账号

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 测试参数

	TestConfig *uint64 `json:"TestConfig,omitempty" name:"TestConfig"`
}

func (r *ModifyMailConfigSaveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMailConfigSaveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetsSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetsSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetsSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetsSaveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetsSaveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetsSaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgentBatchDeployRequest struct {
	*tchttp.BaseRequest

	// Agent部署参数

	InstanceInfos []*AgentInstanceInfo `json:"InstanceInfos,omitempty" name:"InstanceInfos"`
	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 0 部署  1 卸载

	DeployType *int64 `json:"DeployType,omitempty" name:"DeployType"`
}

func (r *CreateAgentBatchDeployRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgentBatchDeployRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
