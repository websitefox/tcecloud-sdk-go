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

package v20180411

import (
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeZonesArchInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用Cpu架构信息

		CpuArch []*string `json:"CpuArch,omitempty" name:"CpuArch"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesArchInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZonesArchInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlushBinlogRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *FlushBinlogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FlushBinlogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSmartDBARequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例的分片ID

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

func (r *StartSmartDBARequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSmartDBARequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceBackupFileItem struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例状态

	InstanceStatus *int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
	// 实例内的分片ID

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
	// 文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件下载链接

	Url *string `json:"Url,omitempty" name:"Url"`
	// 文件大小

	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`
	// 备份文件类型

	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
	// 手动备份(仅用于冷备)

	ManualBackup *int64 `json:"ManualBackup,omitempty" name:"ManualBackup"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type CopyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 源用户名

	SrcUserName *string `json:"SrcUserName,omitempty" name:"SrcUserName"`
	// 源用户允许的访问 host

	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`
	// 源账号的 ReadOnly 属性

	SrcReadOnly *string `json:"SrcReadOnly,omitempty" name:"SrcReadOnly"`
	// 目的用户名

	DstUserName *string `json:"DstUserName,omitempty" name:"DstUserName"`
	// 目的用户允许的访问 host

	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`
	// 目的账号的 ReadOnly 属性

	DstReadOnly *string `json:"DstReadOnly,omitempty" name:"DstReadOnly"`
}

func (r *CopyAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyAccountPrivilegesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 文件列表

		Files []*InstanceBackupFileItem `json:"Files,omitempty" name:"Files"`
		// 文件下载链接路径前缀

		UrlPrefix *string `json:"UrlPrefix,omitempty" name:"UrlPrefix"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNetworkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID，根据此FlowId通过DescribeFlow接口查询任务进行状态

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZonesInfo struct {

	// 可用区英文ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区数字ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区中文名

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type FlushBinlogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FlushBinlogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FlushBinlogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如dcdbt-87zif6ha

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroyHourDCDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyHourDCDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTmpInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteTmpInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTmpInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddShardConfig struct {

	// 新增分片的数量

	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`
	// 分片内存大小，单位 GB

	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`
	// 分片存储大小，单位 GB

	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
	// 分片cpu大小，单位：核，可以通过 DescribeShardSpec 查询实例规格获得。

	ShardCpu *int64 `json:"ShardCpu,omitempty" name:"ShardCpu"`
}

type ModifyInstanceSSLAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceSSLAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceSSLAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Project struct {

	// 应用Id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者uin

	CreatorUin *int64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 描述信息

	Info *string `json:"Info,omitempty" name:"Info"`
	// 是否默认项目，1 是，0 不是

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 项目名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源拥有者（主账号）uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 来源AppId

	SrcAppId *int64 `json:"SrcAppId,omitempty" name:"SrcAppId"`
	// 来源平台

	SrcPlat *string `json:"SrcPlat,omitempty" name:"SrcPlat"`
	// 项目状态,0正常，-1关闭。默认项目为3

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeDBSyncModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 同步模式：0 异步，1 强同步， 2 强同步可退化

		SyncMode *int64 `json:"SyncMode,omitempty" name:"SyncMode"`
		// 是否有修改流程在执行中：1 是， 0 否。

		IsModifying *int64 `json:"IsModifying,omitempty" name:"IsModifying"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSyncModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 销毁成功的实例id列表

		SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitempty" name:"SuccessInstanceIds"`
		// 销毁失败的实例id列表

		FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateHourDCDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEnginesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// items总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// DB引擎信息

		Items []*DBEngineInfo `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBEnginesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEnginesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Database struct {

	// 数据库名称

	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

type CloneAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 源用户账户名

	SrcUser *string `json:"SrcUser,omitempty" name:"SrcUser"`
	// 源用户HOST

	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`
	// 目的用户账户名

	DstUser *string `json:"DstUser,omitempty" name:"DstUser"`
	// 目的用户HOST

	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`
	// 目的用户账户描述

	DstDesc *string `json:"DstDesc,omitempty" name:"DstDesc"`
}

func (r *CloneAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBTmpInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBTmpInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBTmpInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlowLogData struct {

	// 语句校验和，用于查询详情

	CheckSum *string `json:"CheckSum,omitempty" name:"CheckSum"`
	// 数据库名称

	Db *string `json:"Db,omitempty" name:"Db"`
	// 样例Sql

	ExampleSql *string `json:"ExampleSql,omitempty" name:"ExampleSql"`
	// 抽象的SQL语句

	FingerPrint *string `json:"FingerPrint,omitempty" name:"FingerPrint"`
	// 平均的锁时间

	LockTimeAvg *string `json:"LockTimeAvg,omitempty" name:"LockTimeAvg"`
	// 最大锁时间

	LockTimeMax *string `json:"LockTimeMax,omitempty" name:"LockTimeMax"`
	// 最小锁时间

	LockTimeMin *string `json:"LockTimeMin,omitempty" name:"LockTimeMin"`
	// 锁时间总和

	LockTimeSum *string `json:"LockTimeSum,omitempty" name:"LockTimeSum"`
	// 查询次数

	QueryCount *string `json:"QueryCount,omitempty" name:"QueryCount"`
	// 平均查询时间

	QueryTimeAvg *string `json:"QueryTimeAvg,omitempty" name:"QueryTimeAvg"`
	// 最大查询时间

	QueryTimeMax *string `json:"QueryTimeMax,omitempty" name:"QueryTimeMax"`
	// 最小查询时间

	QueryTimeMin *string `json:"QueryTimeMin,omitempty" name:"QueryTimeMin"`
	// 查询时间总和

	QueryTimeSum *string `json:"QueryTimeSum,omitempty" name:"QueryTimeSum"`
	// 扫描行数

	RowsExaminedSum *string `json:"RowsExaminedSum,omitempty" name:"RowsExaminedSum"`
	// 发送行数

	RowsSentSum *string `json:"RowsSentSum,omitempty" name:"RowsSentSum"`
	// 最后执行时间

	TsMax *string `json:"TsMax,omitempty" name:"TsMax"`
	// 首次执行时间

	TsMin *string `json:"TsMin,omitempty" name:"TsMin"`
	// 帐号

	User *string `json:"User,omitempty" name:"User"`
}

type DescribeDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBEncryptAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEncryptAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例详细信息列表

		Instances []*DCDBInstanceInfo `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstanceShardInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分片总个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 分片信息

		ShardsInfo []*MonitorShardInfo `json:"ShardsInfo,omitempty" name:"ShardsInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBInstanceShardInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBInstanceShardInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceSSLAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否开启实例的SSL认证。0-关闭；1-开启

	SSLEnabled *int64 `json:"SSLEnabled,omitempty" name:"SSLEnabled"`
}

func (r *ModifyInstanceSSLAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceSSLAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseObjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 透传入参。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 数据库名称。

		DbName *string `json:"DbName,omitempty" name:"DbName"`
		// 表列表。

		Tables []*DatabaseTable `json:"Tables,omitempty" name:"Tables"`
		// 视图列表。

		Views []*DatabaseView `json:"Views,omitempty" name:"Views"`
		// 存储过程列表。

		Procs []*DatabaseProcedure `json:"Procs,omitempty" name:"Procs"`
		// 函数列表。

		Funcs []*DatabaseFunction `json:"Funcs,omitempty" name:"Funcs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseObjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabaseObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchRollbackInstanceRequest struct {
	*tchttp.BaseRequest

	// 源/旧实例ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 目标实例ID，形如：tdsql-ow728lmc。

	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`
}

func (r *SwitchRollbackInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchRollbackInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstanceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID，形如dcdbt-7oaxtcb7

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例名称

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 实例状态。0-实例创建中；1-异步任务处理中；2-运行中；3-实例未初始化；-1-实例已隔离

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 实例目前运行状态描述

		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
		// 实例内网IP地址

		Vip *string `json:"Vip,omitempty" name:"Vip"`
		// 实例内网端口

		Vport *int64 `json:"Vport,omitempty" name:"Vport"`
		// 实例节点数。值为2时表示一主一从，值为3时表示一主二从

		NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
		// 实例所在地域名称，形如ap-guangzhou

		Region *string `json:"Region,omitempty" name:"Region"`
		// 实例私有网络ID，形如vpc-r9jr0de3

		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
		// 实例私有网络子网ID，形如subnet-6rqs61o2

		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
		// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中；4-关闭中

		WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`
		// 外网访问的域名，公网可解析

		WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`
		// 外网IP地址，公网可访问

		WanVip *string `json:"WanVip,omitempty" name:"WanVip"`
		// 外网访问端口

		WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`
		// 实例所属项目ID

		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
		// 实例自动续费标志。0-正常续费；1-自动续费；2-到期不续费

		AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
		// 独享集群ID

		ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`
		// 付费模式。prepaid-预付费；postpaid-按量计费

		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
		// 实例创建时间，格式为 2006-01-02 15:04:05

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 实例到期时间，格式为 2006-01-02 15:04:05

		PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
		// 数据库版本信息

		DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`
		// 实例是否支持审计。0-不支持；1-支持

		IsAuditSupported *int64 `json:"IsAuditSupported,omitempty" name:"IsAuditSupported"`
		// 实例是否支持数据加密。0-不支持；1-支持

		IsEncryptSupported *int64 `json:"IsEncryptSupported,omitempty" name:"IsEncryptSupported"`
		// 实例母机机器型号

		Machine *string `json:"Machine,omitempty" name:"Machine"`
		// 实例内存大小，单位 GB，各个分片的内存大小的和

		Memory *int64 `json:"Memory,omitempty" name:"Memory"`
		// 实例磁盘存储大小，单位 GB，各个分片的磁盘大小的和

		Storage *int64 `json:"Storage,omitempty" name:"Storage"`
		// 实例存储空间使用率，计算方式为：各个分片已经使用的磁盘大小的和/各个分片的磁盘大小的和。

		StorageUsage *float64 `json:"StorageUsage,omitempty" name:"StorageUsage"`
		// 日志存储空间大小，单位GB

		LogStorage *int64 `json:"LogStorage,omitempty" name:"LogStorage"`
		// 产品类型ID

		Pid *int64 `json:"Pid,omitempty" name:"Pid"`
		// 主DB可用区

		MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`
		// 从DB可用区

		SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones"`
		// 分片信息

		Shards []*ShardBriefInfo `json:"Shards,omitempty" name:"Shards"`
		// 内网IPv6

		Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`
		// 实例Cpu核数

		Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
		// 实例QPS

		Qps *int64 `json:"Qps,omitempty" name:"Qps"`
		// DB引擎

		DbEngine *string `json:"DbEngine,omitempty" name:"DbEngine"`
		// 是否支持IPv6

		Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
		// 外网IPv6地址，公网可访问

		WanVipv6 *string `json:"WanVipv6,omitempty" name:"WanVipv6"`
		// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中；4-关闭中

		WanStatusIpv6 *int64 `json:"WanStatusIpv6,omitempty" name:"WanStatusIpv6"`
		// 外网IPv6端口

		WanPortIpv6 *int64 `json:"WanPortIpv6,omitempty" name:"WanPortIpv6"`
		// 标签信息

		ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
		// DCN标志，0-无，1-主实例，2-灾备实例

		DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`
		// DCN状态，0-无，1-创建中，2-同步中，3-已断开

		DcnStatus *int64 `json:"DcnStatus,omitempty" name:"DcnStatus"`
		// DCN灾备实例数

		DcnDstNum *int64 `json:"DcnDstNum,omitempty" name:"DcnDstNum"`
		// 平台项目ID

		PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
		// 平台项目名称

		PlatformProjectName *string `json:"PlatformProjectName,omitempty" name:"PlatformProjectName"`
		// 1： 主实例（独享型）, 2: 主实例, 3： 灾备实例, 4： 灾备实例（独享型）

		InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
		// cpu架构

		CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
		// 就近访问策略，0-无策略, 1-可用区就近访问

		RsAccessStrategy *int64 `json:"RsAccessStrategy,omitempty" name:"RsAccessStrategy"`
		// 资源池名称

		ExclusterName *string `json:"ExclusterName,omitempty" name:"ExclusterName"`
		// 尚未回收的网络资源

		ReservedNetResources []*ReservedNetResource `json:"ReservedNetResources,omitempty" name:"ReservedNetResources"`
		// 是否支持强同步dcn

		IsDcnSqlSyncSupported *int64 `json:"IsDcnSqlSyncSupported,omitempty" name:"IsDcnSqlSyncSupported"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProxyConfigParam struct {

	// 配置参数名称

	Param *string `json:"Param,omitempty" name:"Param"`
	// 配置参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type StartSmartDBAResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务流程ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartSmartDBAResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSmartDBAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTmpDCDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务流ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTmpDCDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTmpDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRemarkRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyInstanceRemarkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRemarkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务流程ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：dcdb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组详情。

		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`
		// 符合条件的安全组总数量。

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {

	// 标签键key

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值value

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID，透传入参。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 用户名，透传入参。

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 允许访问的 host，透传入参。

		Host *string `json:"Host,omitempty" name:"Host"`
		// 透传入参。

		ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户允许的访问 host

	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBBinlogTimeRequest struct {
	*tchttp.BaseRequest

	// 需要回档的实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDCDBBinlogTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBBinlogTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 参数列表，每一个元素是Param和Value的组合，可以通过 DescribeDBParameters 查询可配置的参数和值限制

	Params []*DBParamValue `json:"Params,omitempty" name:"Params"`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TablePrivilege struct {

	// 数据库

	Database *string `json:"Database,omitempty" name:"Database"`
	// 表名

	Table *string `json:"Table,omitempty" name:"Table"`
	// 权限集合

	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type CommonDBInstance struct {

	// 应用Id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 实例创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 0-基础网络，1-VPC网络

	NetType *int64 `json:"NetType,omitempty" name:"NetType"`
	// 计费类型，1-包年包月，0-按量计费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 项目Id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例状态信息，0-创建中,1-运行中，2-隔离中,下线后无法拉取实例信息

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网统一Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// VPC网络IP

	Vips []*string `json:"Vips,omitempty" name:"Vips"`
	// VPC网络统一Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPC网络端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DestroyHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID，形如dcdbt-87zif6ha

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 删除实例任务异步流程ID，以FlowId作为参数调用API 接口DescribeFlow 查询删除任务执行状态。

		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyHourDCDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {

	// DB节点ID

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// DB节点角色，取值为master或者slave

	Role *string `json:"Role,omitempty" name:"Role"`
}

type Rsip struct {

	// Ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DescribeLatestCloudDBAReportRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例分片ID

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

func (r *DescribeLatestCloudDBAReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLatestCloudDBAReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstanceShardInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tdsqlshard-r5h7x4x1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 单次请求最多返回多少条数据

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 从偏移量Offset处开始返回数据

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDCDBInstanceShardInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBInstanceShardInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KillSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *KillSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KillSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// db账号

	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`
	// 全局权限, 不传表示保留全局权限，传[]表示清除

	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges"`
	// 数据库权限, 不传表示保留数据库权限，传[]表示清除

	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges"`
	// 表权限, 不传表示保留表权限，传[]表示清除

	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges"`
	// 字段局权限, 不传表示保留字段权限，传[]表示清除

	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccountPrivilegesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLatestCloudDBAReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本次检测的有效分数

		Score *int64 `json:"Score,omitempty" name:"Score"`
		// 本次检测报告的名称

		ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
		// 本次检测报告的详情

		ReportDetail *string `json:"ReportDetail,omitempty" name:"ReportDetail"`
		// 本次检测报告的完成时间

		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLatestCloudDBAReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLatestCloudDBAReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceVportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点总个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 节点信息

		NodesInfo []*BriefNodeInfo `json:"NodesInfo,omitempty" name:"NodesInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBInstanceNodeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBInstanceNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableColumn struct {

	// 列名称

	Col *string `json:"Col,omitempty" name:"Col"`
	// 列类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DeleteTmpInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务流程ID。

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTmpInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTmpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：dcdb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 项目ID，TCE上恒传0。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 拉取数量限制。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索条件，支持安全组id或者安全组名称。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConstraintRange struct {

	// 约束类型为section时的最大值

	Max *string `json:"Max,omitempty" name:"Max"`
	// 约束类型为section时的最小值

	Min *string `json:"Min,omitempty" name:"Min"`
}

type SpecConfig struct {

	// 规格机型

	Machine *string `json:"Machine,omitempty" name:"Machine"`
	// 规格列表

	SpecConfigInfos []*SpecConfigInfo `json:"SpecConfigInfos,omitempty" name:"SpecConfigInfos"`
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 分片 ID，形如：shard-7noic7tv

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
	// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeHourDCDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID，透传入参。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例用户列表。

		Users []*DBAccount `json:"Users,omitempty" name:"Users"`
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

type ModifyDBSyncModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务Id，可通过 DescribeFlow 查询任务状态。

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBSyncModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例 ID，形如：dcdbt-ow728lmc。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 请求日志类型。1-binlog，2-冷备，3-errlog，4-slowlog。

		Type *uint64 `json:"Type,omitempty" name:"Type"`
		// 请求日志总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 日志文件列表

		Files []*LogFileInfo `json:"Files,omitempty" name:"Files"`
		// 如果是VPC网络的实例，做用本前缀加上URI为下载地址

		VpcPrefix *string `json:"VpcPrefix,omitempty" name:"VpcPrefix"`
		// 如果是普通网络的实例，做用本前缀加上URI为下载地址

		NormalPrefix *string `json:"NormalPrefix,omitempty" name:"NormalPrefix"`
		// 分片 ID，形如：shard-7noic7tv

		ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBSaleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可售卖地域信息列表

		RegionList []*RegionInfo `json:"RegionList,omitempty" name:"RegionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBSaleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBSaleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流程状态，0：成功，1：失败，2：运行中

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 登录用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示设置该数据库权限（即db.\*），此时忽略 Object 参数

	Type *string `json:"Type,omitempty" name:"Type"`
	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空

	Object *string `json:"Object,omitempty" name:"Object"`
	// 当 Type=table 时，ColName 为 \* 表示对表授权，如果为具体字段名，表示对字段授权

	ColName *string `json:"ColName,omitempty" name:"ColName"`
	// 全局权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER，SHOW DATABASES
	// 库权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER
	// 表/视图权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE VIEW，SHOW VIEW，TRIGGER
	// 存储过程/函数权限： ALTER ROUTINE，EXECUTE
	// 字段权限： INSERT，REFERENCES，SELECT，UPDATE

	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

func (r *GrantAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantAccountPrivilegesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActiveHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如dcdbt-87zif6ha

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ActiveHourDCDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActiveHourDCDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：mariadb,cdb,cynosdb,dcdb,redis,mongodb 等。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 要绑定的安全组ID，类似sg-efil73jd。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 被绑定的实例ID，类似tdsqlshard-lesecurk，支持指定多个实例。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：dcdbt-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRsipRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceRsipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceRsipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBBinlogTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBBinlogTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBBinlogTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBPriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 原价，单位：分

		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
		// 实际价格，单位：分。受折扣等影响，可能和原价不同。计量模式下返回值恒为0。

		Price *int64 `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBPriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例uuid列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *IsolateHourDCDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateHourDCDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBSyncModeRequest struct {
	*tchttp.BaseRequest

	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 同步模式：0 异步，1 强同步， 2 强同步可退化

	SyncMode *int64 `json:"SyncMode,omitempty" name:"SyncMode"`
}

func (r *ModifyDBSyncModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBSyncModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DCDBRollbackInstance struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 地域描述

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区描述

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 项目ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 节点数，2 为一主一从， 3 为一主二从。

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// Cpu核数。

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小，单位 GB。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 存储大小，单位 GB。

	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`
	// 字符串型的私有网络ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 字符串型的私有网络子网ID。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 内网IP。

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 内网端口。

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 实例IPv6标志。

	Ipv6Flag *uint64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// 内网IPv6。

	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`
	// 付费模式。

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 回档时间。

	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`
	// 到期时间。

	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
	// 实例创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 1：非临时实例，0：临时实例。

	IsTmp *uint64 `json:"IsTmp,omitempty" name:"IsTmp"`
	// 实例标签信息。

	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type DescribeDCDBInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 单次最多返回多少条，取值范围为(0-100]，默认为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 返回数据的偏移值，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDCDBInstanceNodeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBInstanceNodeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDcnDetailRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDcnDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDcnDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDcnJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流程ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchDcnJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDcnJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckIpStatusRequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如：dcdbt-ow728lmc

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 目标私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 目标子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 目标虚拟IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 是否IPv6，默认0

	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
}

func (r *CheckIpStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckIpStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorShardInfo struct {

	// 分片ID

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

type DescribeAvailableExclusiveGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 独享资源池数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 独享资源池信息

		Items []*FenceInfoItem `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableExclusiveGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableExclusiveGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GrantAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVportRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例VPORT

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
}

func (r *ModifyInstanceVportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceVportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Account struct {

	// 用户名

	User *string `json:"User,omitempty" name:"User"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。

	Host *string `json:"Host,omitempty" name:"Host"`
}

type BriefNodeInfo struct {

	// DB节点ID

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// DB节点角色，取值为master或者slave

	Role *string `json:"Role,omitempty" name:"Role"`
	// 节点所属分片的分片ID

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

type ParamConstraint struct {

	// 约束类型为enum时的可选值列表

	Enum *string `json:"Enum,omitempty" name:"Enum"`
	// 约束类型为section时的范围

	Range *ConstraintRange `json:"Range,omitempty" name:"Range"`
	// 约束类型为string时的可选值列表

	String *string `json:"String,omitempty" name:"String"`
	// 约束类型,如枚举enum，区间section

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ReservedNetResource struct {

	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 保留的私有网络ip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 保留的端口号

	Vports []*int64 `json:"Vports,omitempty" name:"Vports"`
	// 回收时间

	RecycleTime *string `json:"RecycleTime,omitempty" name:"RecycleTime"`
}

type CreateTmpDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 回档实例的ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 回档时间点

	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`
}

func (r *CreateTmpDCDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTmpDCDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActiveHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恢复成功的实例id列表

		SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitempty" name:"SuccessInstanceIds"`
		// 恢复失败的实例id列表

		FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActiveHourDCDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActiveHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBShardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的分片数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 分片信息列表

		Shards []*DCDBShardInfo `json:"Shards,omitempty" name:"Shards"`
		// 灾备标志，0-无，1-主实例，2-灾备实例

		DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBShardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBShardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：dcdb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 安全组Id。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 实例ID列表，一个或者多个实例Id组成的数组。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupFilesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 分片ID

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
	// 备份类型。目前支持Data,Binlog,Errlog,Slowlog这几种类型

	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
	// 开始时间。格式示例为：2023-02-03 04:05:06

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间。格式示例为：2023-02-03 04:05:06

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 页大小。分页查询参数，最大1000

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页码。分页查询参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询结果排序参数。支持Time、Size 这几种类型

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 查询结果排序类型参数。目前支持desc,asc,DESC,ASC这几种类型

	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
	// 返回的url格式类型。枚举类型，Base:无签名，PreSigned:带签名，控制台页面传Base

	Format *string `json:"Format,omitempty" name:"Format"`
}

func (r *DescribeBackupFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBSaleInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDCDBSaleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBSaleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseTable struct {

	// 表名

	Table *string `json:"Table,omitempty" name:"Table"`
}

type DescribeCommonDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例详细信息列表

		InstanceDetails []*CommonDBInstance `json:"InstanceDetails,omitempty" name:"InstanceDetails"`
		// 符合条件的实例数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCommonDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCommonDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNetworkRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 希望转到的VPC网络的VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 希望转到的VPC网络的子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 如果需要指定VIP，填上该字段

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 如果需要指定VIPv6，填上该字段

	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`
	// vip回收延时

	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitempty" name:"VipReleaseDelay"`
}

func (r *ModifyInstanceNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableExclusiveGroupsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableExclusiveGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableExclusiveGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseFunction struct {

	// 函数名称

	Func *string `json:"Func,omitempty" name:"Func"`
}

type DescribeDatabaseTableRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库名称，通过 DescribeDatabases 接口获取。

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 表名称，通过 DescribeDatabaseObjects 接口获取。

	Table *string `json:"Table,omitempty" name:"Table"`
}

func (r *DescribeDatabaseTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabaseTableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例 ID，形如：dcdbt-ow7t8lmc。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 请求DB的当前参数值

		Params []*ParamDesc `json:"Params,omitempty" name:"Params"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEnginesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDBEnginesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEnginesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例总的个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例列表

		Instances []*InstanceResource `json:"Instances,omitempty" name:"Instances"`
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

type ParamDesc struct {

	// 参数限制

	Constraint *ParamConstraint `json:"Constraint,omitempty" name:"Constraint"`
	// 系统默认值

	Default *string `json:"Default,omitempty" name:"Default"`
	// 是否有设置过值，false:没有设置过值，true:有设置过值。

	HaveSetValue *bool `json:"HaveSetValue,omitempty" name:"HaveSetValue"`
	// 参数名字

	Param *string `json:"Param,omitempty" name:"Param"`
	// 设置过的值，参数生效后，该值和value一样。未设置过就不返回该字段。

	SetValue *string `json:"SetValue,omitempty" name:"SetValue"`
	// 当前参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CloneAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务流程ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseNetResourceRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 私有网络

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 待回收的vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *ReleaseNetResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseNetResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecConfigInfo struct {

	// CPU核数

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 数据盘规格最大值，单位 GB

	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 内存大小，单位 GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 数据盘规格最小值，单位 GB

	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`
	// 节点个数，2 表示一主一从，3 表示一主二从

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 产品类型 Id

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 最大 Qps 值

	Qps *int64 `json:"Qps,omitempty" name:"Qps"`
	// 推荐的使用场景

	SuitInfo *string `json:"SuitInfo,omitempty" name:"SuitInfo"`
}

type DescribeDcnDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DCN同步详情

		DcnDetails []*DcnDetailItem `json:"DcnDetails,omitempty" name:"DcnDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDcnDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDcnDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBCharsetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DBCharsets总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// DB字符集信息

		DBCharsetItem *DBCharsetItem `json:"DBCharsetItem,omitempty" name:"DBCharsetItem"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBCharsetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBCharsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 可选择的主可用区和从可用区

	AvailableChoice []*ShardZoneChooseInfo `json:"AvailableChoice,omitempty" name:"AvailableChoice"`
	// 地域英文ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域数字ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域中文名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区列表

	ZoneList []*ZonesInfo `json:"ZoneList,omitempty" name:"ZoneList"`
}

type DescribeUserTasksRequest struct {
	*tchttp.BaseRequest

	// 任务的状态列表。0-任务启动中；1-任务运行中；2-任务成功；3-任务失败

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 实例ID列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 任务类型列表，当前支持的任务类型有：0-回档任务；1-创建实例任务；2-扩容任务；3-迁移任务；4-删除实例任务；5-重启任务

	FlowTypes []*int64 `json:"FlowTypes,omitempty" name:"FlowTypes"`
	// 任务的创建时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务的结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 任务ID的数组

	UTaskIds []*int64 `json:"UTaskIds,omitempty" name:"UTaskIds"`
	// 每次最多返回多少条任务，取值范围为(0,100]，不传的话默认值为20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回任务默认是按照创建时间降序排列，从偏移值Offset处开始返回

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUserTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRealServerAccessStrategyRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// RS就近模式, 0-无策略, 1-可用区就近访问

	RsAccessStrategy *int64 `json:"RsAccessStrategy,omitempty" name:"RsAccessStrategy"`
}

func (r *ModifyRealServerAccessStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRealServerAccessStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DcnDetailItem struct {

	// 实例DCN标志，1-主，2-备

	DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`
	// 实例DCN状态，0-无，1-创建中，2-同步中，3-已断开

	DcnStatus *int64 `json:"DcnStatus,omitempty" name:"DcnStatus"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 实例IP地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 实例IPv6地址

	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`
	// 实例端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 实例可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// KMS开启状态，0-未开启，1-已开启

	EncryptStatus *int64 `json:"EncryptStatus,omitempty" name:"EncryptStatus"`
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该实例上的数据库列表。

		Databases []*Database `json:"Databases,omitempty" name:"Databases"`
		// 透传入参。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBRollbackInstancesRequest struct {
	*tchttp.BaseRequest

	// 原实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 偏移量，默认为 0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为 10，最大值为 150。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDCDBRollbackInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBRollbackInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckIpStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 透传入参

		Vip *string `json:"Vip,omitempty" name:"Vip"`
		// Vip可用状态：0 - 可用，300 - 此ip已被子网内其他用户占用， 301 - 此ip不在子网内， 302 - ip格式错误， 303 - 该IP为实例当前IP。

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckIpStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckIpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID的数组

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 每次最多拉取多少个实例

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 拉取实例列表时候的偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 按照项目ID进行过滤

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 私有网络ID，形如 vpc-nobbvxkf

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// 子网ID，形如 subnet-dnn20os8

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
	// 实例的SetId

	SerialIds []*string `json:"SerialIds,omitempty" name:"SerialIds"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitDCDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 待初始化的实例ID列表，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步）。

	Params []*DBParamValue `json:"Params,omitempty" name:"Params"`
}

func (r *InitDCDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitDCDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KillSessionRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 会话ID列表

	SessionId []*int64 `json:"SessionId,omitempty" name:"SessionId"`
	// 分片ID，与ShardSerialId两个需要设置一个

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
	// 分片序列ID，与ShardId两者需要设置一个

	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`
}

func (r *KillSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KillSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否启用加密，1-已开启；0-未开启。

		EncryptStatus *int64 `json:"EncryptStatus,omitempty" name:"EncryptStatus"`
		// DEK密钥

		CipherText *string `json:"CipherText,omitempty" name:"CipherText"`
		// DEK密钥过期日期。

		ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBEncryptAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitShardConfig struct {

	// 分片ID数组

	ShardInstanceIds []*string `json:"ShardInstanceIds,omitempty" name:"ShardInstanceIds"`
	// 分片内存大小，单位 GB

	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`
	// 分片存储大小，单位 GB

	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
	// 数据切分比例

	SplitRate *int64 `json:"SplitRate,omitempty" name:"SplitRate"`
	// 分片cpu大小，单位：核，可以通过 DescribeShardSpec 查询实例规格获得。

	ShardCpu *int64 `json:"ShardCpu,omitempty" name:"ShardCpu"`
}

type UserTaskInfo struct {

	// 任务ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 用户账户ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 任务状态，0-任务初始化中；1-任务运行中；2-任务成功；3-任务失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务类型，0-实例回档；1-实例创建；2-实例扩容；3-实例迁移；4-实例删除；5-实例重启

	UserTaskType *int64 `json:"UserTaskType,omitempty" name:"UserTaskType"`
	// 任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 任务错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 客户端参数

	InputData *string `json:"InputData,omitempty" name:"InputData"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type DescribeCommonDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例Id过滤列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例名称过滤列表

	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`
	// 计费类型过滤列表，1-包年包月，0-按量计费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// VPC网络Id信息过滤列表

	VpcIds []*int64 `json:"VpcIds,omitempty" name:"VpcIds"`
	// 子网Id信息过滤列表

	SubnetIds []*int64 `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// VPC网络统一Id信息过滤列表

	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds"`
	// 子网统一Id信息过滤列表

	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`
	// 实例状态过滤列表

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 实例Vip信息列表

	Vips []*string `json:"Vips,omitempty" name:"Vips"`
	// 排序字段，支持ProjectId, InstanceName, CreateTime

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序方式，支持asc，desc

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 数量限制，默认推荐100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCommonDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCommonDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamModifyResult struct {

	// 参数修改结果。0表示修改成功；-1表示修改失败；-2表示该参数值非法

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 修改参数名字

	Param *string `json:"Param,omitempty" name:"Param"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否启用数据加密，开启后暂不支持关闭。本接口的可选值为：1-开启数据加密。

	EncryptEnabled *int64 `json:"EncryptEnabled,omitempty" name:"EncryptEnabled"`
}

func (r *ModifyDBEncryptAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBEncryptAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseView struct {

	// 视图名称

	View *string `json:"View,omitempty" name:"View"`
}

type DescribeDBCharsetsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDBCharsetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBCharsetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBCharsetItem struct {

	// DB版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// DB默认字符集

	DefaultCharset *string `json:"DefaultCharset,omitempty" name:"DefaultCharset"`
	// DB可选字符集

	OptionalCharsets *string `json:"OptionalCharsets,omitempty" name:"OptionalCharsets"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：dcdb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVipRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例VIP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// IPv6标志

	Ipv6Flag *uint64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// vip回收延时

	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitempty" name:"VipReleaseDelay"`
}

func (r *ModifyInstanceVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例 ID，形如：dcdbt-ow728lmc。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 各参数修改结果

		Result []*ParamModifyResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRealServerAccessStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRealServerAccessStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRealServerAccessStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TmpInstance struct {

	// 应用ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例备注

	InstanceRemark *string `json:"InstanceRemark,omitempty" name:"InstanceRemark"`
	// 实例IPv6标志

	Ipv6Flag *uint64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// 有效期结束时间

	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
	// 实例所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 源实例 ID，形如：tdsql-ow728lmc。

	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`
	// 实例状态,0:待初始化,1:流程处理中,2:有效状态,-1:已隔离，-2：已下线

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 0:非临时实例 ,1:无效临时实例, 2:回档成功的有效临时实例

	TempType *int64 `json:"TempType,omitempty" name:"TempType"`
	// 实例虚IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 实例虚IPv6

	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`
	// 实例虚端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 实例所在可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeDBSlowLogAnalysisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 慢查询SQL出现的开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 慢查询SQL出现的结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 返回的慢查询详情数据。每个点代表一个时间段内慢查询SQL出现的记录次数，一个时间段的长度由请求的开始时间和结束时间的差值决定，小于1天是5分钟一段，大于1天小于7天是30分钟一段，大于7天是2个小时一段。

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 采样周期

		Period *int64 `json:"Period,omitempty" name:"Period"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSlowLogAnalysisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSlowLogAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例 ID 查询。实例 ID 形如：dcdbt-2t4cf98d

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 搜索的字段名，当前支持的值有：instancename、vip、all。传 instancename 表示按实例名进行搜索；传 vip 表示按内网IP进行搜索；传 all 将会按实例ID、实例名和内网IP进行搜索。

	SearchName *string `json:"SearchName,omitempty" name:"SearchName"`
	// 搜索的关键字，支持模糊搜索。多个关键字使用换行符（'\n'）分割。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 按项目 ID 查询

	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
	// 是否根据 VPC 网络来搜索

	IsFilterVpc *bool `json:"IsFilterVpc,omitempty" name:"IsFilterVpc"`
	// 私有网络 ID， IsFilterVpc 为 1 时有效

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络的子网 ID， IsFilterVpc 为 1 时有效

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 排序字段， projectId， createtime， instancename 三者之一

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序类型， desc 或者 asc

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 偏移量，默认为 0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为 10，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 1非独享集群，2独享集群， 0全部

	ExclusterType *int64 `json:"ExclusterType,omitempty" name:"ExclusterType"`
	// 标识是否使用ExclusterType字段, false不使用，true使用

	IsFilterExcluster *bool `json:"IsFilterExcluster,omitempty" name:"IsFilterExcluster"`
	// 独享集群ID

	ExclusterIds []*string `json:"ExclusterIds,omitempty" name:"ExclusterIds"`
	// 按标签key查询

	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
	// 实例类型过滤，1-独享实例，2-主实例，3-灾备实例，多个按逗号分隔

	FilterInstanceType *string `json:"FilterInstanceType,omitempty" name:"FilterInstanceType"`
	// 筛选的实例状态

	Status []*int64 `json:"Status,omitempty" name:"Status"`
	// 过滤的实例状态

	ExcludeStatus []*int64 `json:"ExcludeStatus,omitempty" name:"ExcludeStatus"`
}

func (r *DescribeDCDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceProxyConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关配置参数

		ProxyConfigParamSet []*ProxyConfigParam `json:"ProxyConfigParamSet,omitempty" name:"ProxyConfigParamSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceProxyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceProxyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColumnPrivilege struct {

	// 数据库

	Database *string `json:"Database,omitempty" name:"Database"`
	// 表名

	Table *string `json:"Table,omitempty" name:"Table"`
	// 列名

	Column *string `json:"Column,omitempty" name:"Column"`
	// 权限集合

	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type DescribeZonesArchInfoRequest struct {
	*tchttp.BaseRequest

	// 可用区信息

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
}

func (r *DescribeZonesArchInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZonesArchInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelDcnJobRequest struct {
	*tchttp.BaseRequest

	// 灾备实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CancelDcnJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelDcnJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShardInfo struct {

	// Cpu核数

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 创建时间

	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
	// 内存大小，单位 GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 节点数，2 为一主一从， 3 为一主二从

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 产品类型 Id（过时字段，请勿依赖该值）

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 分片数字ID

	ShardId *int64 `json:"ShardId,omitempty" name:"ShardId"`
	// 分片ID

	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`
	// 分片Set ID

	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`
	// 状态：0 创建中，1 流程处理中， 2 运行中，3 分片未初始化，-2 分片已删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 存储大小，单位 GB

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 权限列表。

		Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
		// 数据库账号用户名

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 数据库账号Host

		Host *string `json:"Host,omitempty" name:"Host"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tdsql-hdaprz0v

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShardSpecRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeShardSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShardSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBParamValue struct {

	// 参数名称

	Param *string `json:"Param,omitempty" name:"Param"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CancelDcnJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流程ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelDcnJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelDcnJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Inbound struct {

	// 策略，ACCEPT或者DROP。

	Action *string `json:"Action,omitempty" name:"Action"`
	// 地址组id代表的地址集合。

	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`
	// 来源Ip或Ip段，例如192.168.0.0/16。

	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`
	// 描述。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 安全组id代表的地址集合。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 网络协议，支持udp、tcp等。

	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
	// 端口。

	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`
	// 服务组id代表的协议和端口集合。

	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBAccount struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障
	// 建议该参数值大于10。当ReadOnly选择1、2时该参数生效。

	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
	// 用户备注信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 用户可以从哪台主机登录（对应 MySQL 用户的 host 字段，UserName + Host 唯一标识一个用户，IP形式，IP段以%结尾；支持填入%；为空默认等于%）

	Host *string `json:"Host,omitempty" name:"Host"`
	// 只读标记，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败。

	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 最后更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type Outbound struct {

	// 策略，ACCEPT或者DROP。

	Action *string `json:"Action,omitempty" name:"Action"`
	// 地址组id代表的地址集合。

	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`
	// 来源Ip或Ip段，例如192.168.0.0/16。

	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`
	// 描述。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 安全组id代表的地址集合。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 网络协议，支持udp、tcp等。

	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
	// 端口。

	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`
	// 服务组id代表的协议和端口集合。

	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`
}

type DCDBShardInfo struct {

	// CPU核数

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 所属实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 内存大小，单位 GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 内存使用率，单位为 %

	MemoryUsage *float64 `json:"MemoryUsage,omitempty" name:"MemoryUsage"`
	// 节点数，2 为一主一从， 3 为一主二从

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 付费模型

	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`
	// 到期时间

	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
	// 产品ProductID

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// Proxy版本

	ProxyVersion *string `json:"ProxyVersion,omitempty" name:"ProxyVersion"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 数字分片Id（过时字段，请勿依赖该值）

	ShardId *int64 `json:"ShardId,omitempty" name:"ShardId"`
	// 全局唯一的分片Id

	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`
	// 分片的主可用区

	ShardMasterZone *string `json:"ShardMasterZone,omitempty" name:"ShardMasterZone"`
	// 分片SQL透传Id，用于将sql透传到指定分片执行

	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`
	// 分片的从可用区列表

	ShardSlaveZones []*string `json:"ShardSlaveZones,omitempty" name:"ShardSlaveZones"`
	// 状态：0 创建中，1 流程处理中， 2 运行中，3 分片未初始化

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 状态中文描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 存储大小，单位 GB

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 存储使用率，单位为 %

	StorageUsage *float64 `json:"StorageUsage,omitempty" name:"StorageUsage"`
	// 字符串格式的私有网络子网Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 字符串格式的私有网络Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 分片ShardKey的范围（总共64个哈希值），例如： 0-31，32-63

	Range *string `json:"Range,omitempty" name:"Range"`
	// cpu架构

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
}

type DescribeDBSlowLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有语句锁时间总和

		LockTimeSum *float64 `json:"LockTimeSum,omitempty" name:"LockTimeSum"`
		// 所有语句查询总次数

		QueryCount *int64 `json:"QueryCount,omitempty" name:"QueryCount"`
		// 总记录数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 所有语句查询时间总和

		QueryTimeSum *float64 `json:"QueryTimeSum,omitempty" name:"QueryTimeSum"`
		// 慢查询日志数据

		Data []*SlowLogData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSlowLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabasePrivilege struct {

	// 数据库

	Database *string `json:"Database,omitempty" name:"Database"`
	// 权限集合

	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type DescribeLogFileRetentionPeriodResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例 ID，形如：tdsql-ow728lmc。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 日志备份天数

		Days *uint64 `json:"Days,omitempty" name:"Days"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogFileRetentionPeriodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogFileRetentionPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthenticateCAMResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前请求账号是否为根账号

		IsRoot *int64 `json:"IsRoot,omitempty" name:"IsRoot"`
		// 当前请求账号是否有全部权限

		IsAllPass *int64 `json:"IsAllPass,omitempty" name:"IsAllPass"`
		// 主账号的appId

		OwnerAppId *int64 `json:"OwnerAppId,omitempty" name:"OwnerAppId"`
		// 权限详情

		Permissions []*Permission `json:"Permissions,omitempty" name:"Permissions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuthenticateCAMResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuthenticateCAMResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBPriceRequest struct {
	*tchttp.BaseRequest

	// 欲新购实例的可用区ID。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 欲购买实例的数量，目前支持购买1-10个实例

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 欲购买的时长，单位：月。不传时默认查询1个月的价格。付费模式为 postpaid 时该参数无效，固定查询第一计费阶梯使用1小时的费用。

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 单个分片节点个数大小，可以通过 DescribeShardSpec
	//  查询实例规格获得。

	ShardNodeCount *int64 `json:"ShardNodeCount,omitempty" name:"ShardNodeCount"`
	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。

	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`
	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。

	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。

	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`
	// 付费类型。postpaid：按量付费；prepaid：预付费。不传时默认为预付费（当前TCE暂不支持预付费功能）。

	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`
	// 分片CPU核数

	ShardCpu *int64 `json:"ShardCpu,omitempty" name:"ShardCpu"`
}

func (r *DescribeDCDBPriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBPriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpandShardConfig struct {

	// 分片ID数组

	ShardInstanceIds []*string `json:"ShardInstanceIds,omitempty" name:"ShardInstanceIds"`
	// 分片内存大小，单位 GB

	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`
	// 分片存储大小，单位 GB

	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
	// 分片cpu大小，单位：核，可以通过 DescribeShardSpec 查询实例规格获得。

	ShardCpu *int64 `json:"ShardCpu,omitempty" name:"ShardCpu"`
	// 分片节点数

	ShardNodeCount *int64 `json:"ShardNodeCount,omitempty" name:"ShardNodeCount"`
}

type CreateHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单对应的实例 ID 列表，如果此处没有返回实例 ID，可以通过订单查询接口获取。还可通过实例查询接口查询实例是否创建完成。

		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
		// 流程id，可以根据流程id查询创建进度

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHourDCDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogFileRetentionPeriodResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例 ID，形如：tdsql-ow728lmc。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogFileRetentionPeriodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogFileRetentionPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBEngineInfo struct {

	// 引擎描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 引擎名称，eg. 基于MariaDB 10.1.9 分支

	Name *string `json:"Name,omitempty" name:"Name"`
	// 引擎类型，MariaDB|Percona

	Type *string `json:"Type,omitempty" name:"Type"`
	// 引擎版本，eg. 10.1.9

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeDatabaseTableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例名称。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 数据库名称。

		DbName *string `json:"DbName,omitempty" name:"DbName"`
		// 表名称。

		Table *string `json:"Table,omitempty" name:"Table"`
		// 列信息。

		Cols []*TableColumn `json:"Cols,omitempty" name:"Cols"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabaseTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 登录用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 新密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShardSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 按机型分类的可售卖规格列表

		SpecConfig []*SpecConfig `json:"SpecConfig,omitempty" name:"SpecConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShardSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShardSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 入站规则。

	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound"`
	// 出站规则。

	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound"`
	// 项目ID。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 安全组ID。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组名称。

	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	// 安全组备注。

	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
}

type UpgradeHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 升级类型，取值范围:
	// <li> ADD: 新增分片 </li>
	//  <li> EXPAND: 升级实例中的已有分片 </li>
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>

	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`
	// 新增分片配置，当UpgradeType为ADD时生效。

	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitempty" name:"AddShardConfig"`
	// 扩容分片配置，当UpgradeType为EXPAND时生效。

	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitempty" name:"ExpandShardConfig"`
	// 切分分片配置，当UpgradeType为SPLIT时生效。

	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitempty" name:"SplitShardConfig"`
	// 切换开始时间，格式如: "2019-12-12 07:00:00"。开始时间必须在当前时间一个小时以后，3天以内。

	SwitchStartTime *string `json:"SwitchStartTime,omitempty" name:"SwitchStartTime"`
	// 切换结束时间,  格式如: "2019-12-12 07:15:00"，结束时间必须大于开始时间。

	SwitchEndTime *string `json:"SwitchEndTime,omitempty" name:"SwitchEndTime"`
	// 是否自动重试。 0：不自动重试  1：自动重试

	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitempty" name:"SwitchAutoRetry"`
	// 用于变更部署，变更配置传空

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
}

func (r *UpgradeHourDCDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeHourDCDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBShardsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：dcdbt-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 分片ID列表。

	ShardInstanceIds []*string `json:"ShardInstanceIds,omitempty" name:"ShardInstanceIds"`
	// 偏移量，默认为 0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段， 目前仅支持 createtime

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序类型， desc 或者 asc

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeDCDBShardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBShardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitDCDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID，可通过 DescribeFlow 查询任务状态。

		FlowIds []*uint64 `json:"FlowIds,omitempty" name:"FlowIds"`
		// 透传入参。

		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDCDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitDCDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseProcedure struct {

	// 存储过程名称

	Proc *string `json:"Proc,omitempty" name:"Proc"`
}

type DescribeDBSyncModeRequest struct {
	*tchttp.BaseRequest

	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSyncModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSyncModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest

	// 异步请求接口返回的任务流程号。

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Permission struct {

	// 是否有权限，1 - 有权限， 0 - 无权限

	IsPermitted *int64 `json:"IsPermitted,omitempty" name:"IsPermitted"`
	// 资源对象

	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type SwitchDcnJobRequest struct {
	*tchttp.BaseRequest

	// 灾备实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *SwitchDcnJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDcnJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRemarkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 是否修改

		IsModify *int64 `json:"IsModify,omitempty" name:"IsModify"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceRemarkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShardZoneChooseInfo struct {

	// 主可用区

	MasterZone *ZonesInfo `json:"MasterZone,omitempty" name:"MasterZone"`
	// 可选的从可用区

	SlaveZones []*ZonesInfo `json:"SlaveZones,omitempty" name:"SlaveZones"`
}

type FenceInfoItem struct {

	// 独享资源池ID

	FenceId *string `json:"FenceId,omitempty" name:"FenceId"`
	// 资源池名称

	FenceName *string `json:"FenceName,omitempty" name:"FenceName"`
	// 资源池是否所有人可见。0-所有人不可见；1-所有人可见

	Visibility *int64 `json:"Visibility,omitempty" name:"Visibility"`
}

type InstanceResource struct {

	// 账户的应用ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 实例所属集群的名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 实例的数据库版本

	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`
	// 数据库实例的数字ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 实例ID，形如 dcdbt-ifzc58fh或者tdsqlshard-ifzc58fh

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例的节点个数

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 实例的原始SerialId

	OriginSerialId *string `json:"OriginSerialId,omitempty" name:"OriginSerialId"`
	// 实例所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例的当前SerialId

	SerialId *string `json:"SerialId,omitempty" name:"SerialId"`
	// 分片信息

	ShardDetail []*InstanceShardInfo `json:"ShardDetail,omitempty" name:"ShardDetail"`
	// 实例运行状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例所属子网的数字ID

	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 实例所属子网的字符串唯一ID

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
	// 实例所属VPC的字符串唯一ID

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// 实例Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 实例所属VPC的数字ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 实例Vport

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 实例所属Zk的名字

	ZkName *string `json:"ZkName,omitempty" name:"ZkName"`
	// 实例所在可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 表示实例监控页面如何展示

	MonitorType *int64 `json:"MonitorType,omitempty" name:"MonitorType"`
}

type DescribeUserTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务列表

		FlowSet []*UserTaskInfo `json:"FlowSet,omitempty" name:"FlowSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogFileRetentionPeriodRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeLogFileRetentionPeriodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogFileRetentionPeriodRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库名称，通过 DescribeDatabases 接口获取。

	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *DescribeDatabaseObjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabaseObjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogAnalysisRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如 dcdbt-hw0qj6m1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 开始时间，形如 2006-01-02 15:04:05

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，形如 2006-01-02 15:04:05，不填的话默认是当前时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 要查询的慢查询语句对应的数据库名称

	Db *string `json:"Db,omitempty" name:"Db"`
	// 要查询的慢查询语句对应的用户名称

	User *string `json:"User,omitempty" name:"User"`
	// 要查询的慢查询语句的校验和，可以通过查询慢查询日志列表获得

	CheckSum *string `json:"CheckSum,omitempty" name:"CheckSum"`
	// 是否查询从机的慢查询。0-主机；1-从机，默认查询主机慢查询

	Slave *int64 `json:"Slave,omitempty" name:"Slave"`
	// 实例分片ID，形如 shard-nbchf68b

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

func (r *DescribeDBSlowLogAnalysisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSlowLogAnalysisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceProxyConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceProxyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceProxyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceShardInfo struct {

	// 分片ID

	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`
	// 分片的SerialId

	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`
	// 分片的运行状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 登录用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 新的账号备注，长度 0~256。

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchRollbackInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务流程ID。

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchRollbackInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchRollbackInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 虚拟私有网络 ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 虚拟私有网络子网 ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。

	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`
	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。

	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
	// 单个分片节点个数，可以通过 DescribeShardSpec
	//  查询实例规格获得。

	ShardNodeCount *int64 `json:"ShardNodeCount,omitempty" name:"ShardNodeCount"`
	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。

	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`
	// 数据库引擎版本，调用 DescribeDBEngines 查询可用的版本。

	DbVersionId *string `json:"DbVersionId,omitempty" name:"DbVersionId"`
	// 分片节点可用区分布，最多可填两个可用区。当分片规格为一主两从时，其中两个节点在第一个可用区。

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// 欲购买实例的数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 分片cpu大小，单位：核，可以通过 DescribeShardSpec
	//  查询实例规格获得。

	ShardCpu *int64 `json:"ShardCpu,omitempty" name:"ShardCpu"`
	// 安全组id

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 实例名称， 可以通过该字段自主的设置实例的名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 资源池 Id，TCE 专用

	ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`
	// 是否支持IPv6

	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// 标签键值对数组

	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
	// DCN源地域

	DcnRegion *string `json:"DcnRegion,omitempty" name:"DcnRegion"`
	// DCN源实例ID

	DcnInstanceId *string `json:"DcnInstanceId,omitempty" name:"DcnInstanceId"`
	// 平台项目ID

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// cpu架构

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步可退化）。

	InitParams []*DBParamValue `json:"InitParams,omitempty" name:"InitParams"`
	// RS访问策略, 0-无策略, 1-可用区就近访问

	RsAccessStrategy *int64 `json:"RsAccessStrategy,omitempty" name:"RsAccessStrategy"`
	// 需要回档的源实例ID

	RollbackInstanceId *string `json:"RollbackInstanceId,omitempty" name:"RollbackInstanceId"`
	// 回档时间

	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`
	// DCN同步模式 "0"代表普通DCN同步，"1"代表一致性同步

	DcnSyncMode *int64 `json:"DcnSyncMode,omitempty" name:"DcnSyncMode"`
}

func (r *CreateHourDCDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHourDCDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如dcdbt-7oaxtcb7

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDCDBInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBInstanceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组详情。

		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSSLAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceSSLAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceSSLAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID的数组

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 重启时间，

	RestartTime *string `json:"RestartTime,omitempty" name:"RestartTime"`
}

func (r *RestartDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRsipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 后端RS信息

		Rsips []*Rsip `json:"Rsips,omitempty" name:"Rsips"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceRsipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceRsipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShardBriefInfo struct {

	// 分片Cpu核数

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 分片创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 分片日志磁盘空间大小，单位GB

	LogDisk *int64 `json:"LogDisk,omitempty" name:"LogDisk"`
	// 分片内存大小，单位GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 分片节点个数

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 分片Proxy版本信息

	ProxyVersion *string `json:"ProxyVersion,omitempty" name:"ProxyVersion"`
	// 分片ID，形如shard-7vg1o339

	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`
	// 分片主DB可用区

	ShardMasterZone *string `json:"ShardMasterZone,omitempty" name:"ShardMasterZone"`
	// 分片SerialId

	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`
	// 分片从DB可用区

	ShardSlaveZones []*string `json:"ShardSlaveZones,omitempty" name:"ShardSlaveZones"`
	// 分片运行状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 分片运行状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 分片磁盘大小，单位GB

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 分片磁盘空间使用率

	StorageUsage *float64 `json:"StorageUsage,omitempty" name:"StorageUsage"`
	// DB节点信息

	NodesInfo []*NodeInfo `json:"NodesInfo,omitempty" name:"NodesInfo"`
}

type LogFileInfo struct {

	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件长度

	Length *uint64 `json:"Length,omitempty" name:"Length"`
	// Log最后修改时间

	Mtime *uint64 `json:"Mtime,omitempty" name:"Mtime"`
	// 下载Log时用到的统一资源标识符

	Uri *string `json:"Uri,omitempty" name:"Uri"`
}

type ReleaseNetResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseNetResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseNetResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogFileRetentionPeriodRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 保存的天数,不能超过30

	Days *uint64 `json:"Days,omitempty" name:"Days"`
}

func (r *ModifyLogFileRetentionPeriodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogFileRetentionPeriodRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBTmpInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 临时实例列表

		TmpInstances []*TmpInstance `json:"TmpInstances,omitempty" name:"TmpInstances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBTmpInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBTmpInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBEncryptAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 登录用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示查询该数据库权限（即db.\*），此时忽略 Object 参数

	Type *string `json:"Type,omitempty" name:"Type"`
	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空

	Object *string `json:"Object,omitempty" name:"Object"`
	// 当 Type=table 时，ColName 为 \* 表示查询表的权限，如果为具体字段名，表示查询对应字段的权限

	ColName *string `json:"ColName,omitempty" name:"ColName"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-hw0qj6m1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 从结果的第几条数据开始返回

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的结果条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询的起始时间，形如2016-07-23 14:55:20

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询的结束时间，形如2016-08-22 14:55:20。如果不填，那么查询结束时间就是当前时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 要查询的具体数据库名称

	Db *string `json:"Db,omitempty" name:"Db"`
	// 排序指标，取值为query_time_sum或者query_count。不填默认按照query_time_sum排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序类型，desc（降序）或者asc（升序）。不填默认desc排序

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 是否查询从机的慢查询，0-主机; 1-从机。不填默认查询主机慢查询

	Slave *int64 `json:"Slave,omitempty" name:"Slave"`
	// 实例的分片ID，形如shard-53ima8ln

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

func (r *DescribeDBSlowLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSlowLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBRollbackInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该实例的回档实例总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 回档实例列表。

		Instances []*DCDBRollbackInstance `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBRollbackInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDCDBRollbackInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// AccountName

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 可以登录的主机，与mysql 账号的 host 格式一致，可以支持通配符，例如 %，10.%，10.20.%。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 账号密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。

	Password *string `json:"Password,omitempty" name:"Password"`
	// 是否创建为只读账号，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败，3：只从备机读取。

	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 账号备注，可以包含中文、英文字符、常见符号和数字，长度为0~256字符

	Description *string `json:"Description,omitempty" name:"Description"`
	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障
	// 建议该参数值大于10。当ReadOnly选择1、2时该参数生效。

	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
}

func (r *CreateAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSSLAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例SSL认证功能当前状态。1-开启中；2-已开启；3-已关闭；4-关闭中

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceSSLAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceSSLAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DCDBInstanceInfo struct {

	// 应用ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 自动续费标志

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// Cpu核数

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 数据库引擎

	DbEngine *string `json:"DbEngine,omitempty" name:"DbEngine"`
	// 数据库引擎版本

	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`
	// DCN灾备实例数

	DcnDstNum *int64 `json:"DcnDstNum,omitempty" name:"DcnDstNum"`
	// DCN标志，0-无，1-主实例，2-灾备实例

	DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`
	// DCN状态，0-无，1-创建中，2-同步中，3-已断开

	DcnStatus *int64 `json:"DcnStatus,omitempty" name:"DcnStatus"`
	// 独享集群ID，为空表示非独享集群实例

	ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`
	// 数字实例ID（过时字段，请勿依赖该值）

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例IPv6标志

	Ipv6Flag *uint64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// 该实例是否支持审计。1-支持；0-不支持

	IsAuditSupported *uint64 `json:"IsAuditSupported,omitempty" name:"IsAuditSupported"`
	// 临时实例标记，0 为非临时实例

	IsTmp *int64 `json:"IsTmp,omitempty" name:"IsTmp"`
	// 隔离时间

	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitempty" name:"IsolatedTimestamp"`
	// 实例处于异步任务状态时，表示异步任务流程ID

	Locker *int64 `json:"Locker,omitempty" name:"Locker"`
	// 内存大小，单位 GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 节点数，2 为一主一从， 3 为一主二从

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 付费模式

	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`
	// 到期时间

	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
	// 产品类型 ID（过时字段，请勿依赖该值）

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 分片个数

	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`
	// 分片详情

	ShardDetail []*ShardInfo `json:"ShardDetail,omitempty" name:"ShardDetail"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 状态中文描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 存储大小，单位 GB

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// Subnet数字ID

	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 账号ID

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 字符串型的私有网络子网ID

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
	// 字符串型的私有网络ID

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// 实例最后更新时间，格式为 2006-01-02 15:04:05

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 内网IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 内网IPv6

	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`
	// VPC数字ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 内网端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 外网访问的域名，公网可解析

	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`
	// 外网端口

	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`
	// 外网IPv6端口

	WanPortIpv6 *uint64 `json:"WanPortIpv6,omitempty" name:"WanPortIpv6"`
	// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中

	WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`
	// 外网IPv6状态

	WanStatusIpv6 *uint64 `json:"WanStatusIpv6,omitempty" name:"WanStatusIpv6"`
	// 外网 IP 地址，公网可访问

	WanVip *string `json:"WanVip,omitempty" name:"WanVip"`
	// 外网IPv6

	WanVipv6 *string `json:"WanVipv6,omitempty" name:"WanVipv6"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 1： 主实例（独享型）, 2: 主实例, 3： 灾备实例, 4： 灾备实例（独享型）

	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
	// Cpu架构

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 资源池名称

	ExclusterName *string `json:"ExclusterName,omitempty" name:"ExclusterName"`
}

type AuthenticateCAMRequest struct {
	*tchttp.BaseRequest

	// 待鉴权的接口名，格式为产品名:API名称，如 dcdb:DescribeDCDBInstances

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 待鉴权资源信息，cam 六段式资源格式或 *

	Resources []*string `json:"Resources,omitempty" name:"Resources"`
}

func (r *AuthenticateCAMRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuthenticateCAMRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
