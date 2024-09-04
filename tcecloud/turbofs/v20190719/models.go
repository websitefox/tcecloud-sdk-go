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

package v20190719

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type VersionCtrlRegionZones struct {

	// 区域可用情况

	RegionZones []*AvailableRegion `json:"RegionZones,omitempty" name:"RegionZones"`
}

type CreateCfsFileSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件系统创建时间

		CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
		// 用户自定义文件系统名称

		CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
		// 文件系统 ID

		FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
		// 文件系统状态

		LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`
		// 文件系统已使用容量大小

		SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`
		// 可用区 ID

		ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
		// 用户自定义文件系统名称

		FsName *string `json:"FsName,omitempty" name:"FsName"`
		// 文件系统是否加密

		Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`
		// 文件系统ID

		CfsId *string `json:"CfsId,omitempty" name:"CfsId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCfsFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleUpFileSystemRequest struct {
	*tchttp.BaseRequest

	// 扩容大小

	TargetCapacity *uint64 `json:"TargetCapacity,omitempty" name:"TargetCapacity"`
	// fsid

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *ScaleUpFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleUpFileSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsFileSystemsRequest struct {
	*tchttp.BaseRequest

	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 此选项值为1时，未命名文件系统名使用英文

	InternationalFlag *uint64 `json:"InternationalFlag,omitempty" name:"InternationalFlag"`
	// 私有网络（VPC） ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网 ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 用户自定义名称

	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
	// 通过字段进行过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCfsFileSystemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsFileSystemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVipStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 子网ID

		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
		// 已查询Vip

		Vip *string `json:"Vip,omitempty" name:"Vip"`
		// 查询结果，取值可为“available”或“unavailable”

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVipStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVipStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCfsFileSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCfsFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCfsFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfsSnapJobHistory struct {

	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 任务类型

	JobType *string `json:"JobType,omitempty" name:"JobType"`
	// 操作人

	Operater *string `json:"Operater,omitempty" name:"Operater"`
	// 任务状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 任务详情

	Message *string `json:"Message,omitempty" name:"Message"`
}

type FileSystemClient struct {

	// 文件系统IP地址

	CfsVip *string `json:"CfsVip,omitempty" name:"CfsVip"`
	// 客户端IP地址

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 文件系统所属VPCID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 可用区名称，例如ap-beijing-1

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区中文名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 该文件系统被挂载到客户端上的路径信息

	MountDirectory *string `json:"MountDirectory,omitempty" name:"MountDirectory"`
	// 可用起ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 链接状态：1：已链接，2:链接断开

	LinkStatus *uint64 `json:"LinkStatus,omitempty" name:"LinkStatus"`
	// 监控状态：1:已安装监控，2：未安装监控

	MonitorStatus *uint64 `json:"MonitorStatus,omitempty" name:"MonitorStatus"`
	// 协议：v3：nfsv3挂载，v4:nfsv4挂载

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeMountTargetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 挂载点详情

		MountTargets []*MountInfo `json:"MountTargets,omitempty" name:"MountTargets"`
		// 挂载点数量

		NumberOfMountTargets *int64 `json:"NumberOfMountTargets,omitempty" name:"NumberOfMountTargets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMountTargetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMountTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 周期统计信息

		PeriodStatisticsList []*PeriodStatisticsSet `json:"PeriodStatisticsList,omitempty" name:"PeriodStatisticsList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AvailableZone struct {

	// 可用区名称

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区中文名称

	ZoneCnName *string `json:"ZoneCnName,omitempty" name:"ZoneCnName"`
	// Type数组

	Types []*AvailableType `json:"Types,omitempty" name:"Types"`
	// 可用区英文名

	ZoneEnName *string `json:"ZoneEnName,omitempty" name:"ZoneEnName"`
}

type CfsTagUser struct {

	// tag id

	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
	// 标识名称

	TagName *string `json:"TagName,omitempty" name:"TagName"`
	// 打开状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// SD是否可售

	SdSaleStatus *string `json:"SdSaleStatus,omitempty" name:"SdSaleStatus"`
	// 是否可售

	HpSaleStatus *string `json:"HpSaleStatus,omitempty" name:"HpSaleStatus"`
	// 可用区id

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 支持SMB

	ISSMB *uint64 `json:"ISSMB,omitempty" name:"ISSMB"`
	// 支持NFS

	ISNFS *uint64 `json:"ISNFS,omitempty" name:"ISNFS"`
}

type SignUpCfsServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *SignUpCfsServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SignUpCfsServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsFileSystemPGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 权限组 ID

		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
		// 文件系统 ID

		FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCfsFileSystemPGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCfsFileSystemPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticsRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PGroupRuleInfo struct {

	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 允许访问的客户端IP

	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`
	// 读写权限, ro为只读，rw为读写

	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`
	// 用户权限。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。

	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`
	// 规则优先级，1-100。 其中 1 为最高，100为最低

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

type UpdateCfsFileSystemNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户自定义文件系统名称

		CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
		// 文件系统ID

		FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
		// 用户自定义文件系统名称

		FsName *string `json:"FsName,omitempty" name:"FsName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCfsFileSystemNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCfsFileSystemNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PGroupInfo struct {

	// 权限组ID

	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
	// 权限组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述信息

	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
	// 创建时间

	CDate *string `json:"CDate,omitempty" name:"CDate"`
	// 关联文件系统个数

	BindCfsNum *int64 `json:"BindCfsNum,omitempty" name:"BindCfsNum"`
}

type DescribeCfsFileSystemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件系统信息

		FileSystems []*FileSystemInfo `json:"FileSystems,omitempty" name:"FileSystems"`
		// 文件系统总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfsFileSystemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsFileSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVipStatusRequest struct {
	*tchttp.BaseRequest

	// 可用区名称

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 待查询Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *DescribeVipStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVipStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileSystemByVpcInfo struct {

	// 文件系统ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 挂载根目录

	FSID *string `json:"FSID,omitempty" name:"FSID"`
	// 挂载点状态

	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`
	// 挂载点 ID

	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`
	// 网络类型

	NetworkInterface *string `json:"NetworkInterface,omitempty" name:"NetworkInterface"`
	// 私有网络 ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网 ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 可用区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type UpdateCfsFileSystemPGroupRequest struct {
	*tchttp.BaseRequest

	// 权限组 ID

	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *UpdateCfsFileSystemPGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCfsFileSystemPGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoSnapshotPolicyInfo struct {

	// 快照策略ID

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 关联的文件系统个数

	FileSystemNums *uint64 `json:"FileSystemNums,omitempty" name:"FileSystemNums"`
	// 星期

	DayOfWeek *string `json:"DayOfWeek,omitempty" name:"DayOfWeek"`
	// 小时

	Hour *string `json:"Hour,omitempty" name:"Hour"`
	// 是否激活

	IsActivated *uint64 `json:"IsActivated,omitempty" name:"IsActivated"`
	// 下次执行时间

	NextActiveTime *string `json:"NextActiveTime,omitempty" name:"NextActiveTime"`
	// 策略状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 保留时间

	AliveDays *uint64 `json:"AliveDays,omitempty" name:"AliveDays"`
	// 地域

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 关联文件系统信息

	FileSystems []*FileSystemByPolicy `json:"FileSystems,omitempty" name:"FileSystems"`
	// appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type AvailableRegion struct {

	// 区域名称，如“ap-beijing”

	Region *string `json:"Region,omitempty" name:"Region"`
	// 区域名称，如“bj”

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 区域可用情况，当区域内至少有一个可用区处于可售状态时，取值为AVAILABLE，否则为UNAVAILABLE

	RegionStatus *string `json:"RegionStatus,omitempty" name:"RegionStatus"`
	// 可用区数组

	Zones []*AvailableZone `json:"Zones,omitempty" name:"Zones"`
	// 区域中文名称，如“广州”

	RegionCnName *string `json:"RegionCnName,omitempty" name:"RegionCnName"`
}

type FileSystemByPolicy struct {

	// 文件系统名称

	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
	// 文件系统ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 文件系统大小

	SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`
	// 文件系统存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 快照总大小

	TotalSnapshotSize *uint64 `json:"TotalSnapshotSize,omitempty" name:"TotalSnapshotSize"`
	// 协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DeleteCfsFileSystemRequest struct {
	*tchttp.BaseRequest

	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DeleteCfsFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCfsFileSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapList struct {

	// 地域名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 快照id

	SnapId *string `json:"SnapId,omitempty" name:"SnapId"`
	// 快照名称

	SnapName *string `json:"SnapName,omitempty" name:"SnapName"`
	// 文件系统id

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 文件系统名

	FsName *string `json:"FsName,omitempty" name:"FsName"`
	// 快照大小B

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 保留天数

	AliveDay *uint64 `json:"AliveDay,omitempty" name:"AliveDay"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 快照状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 快照进度

	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
	// 操作历史

	JobHistorys []*CfsSnapJobHistory `json:"JobHistorys,omitempty" name:"JobHistorys"`
	// 快照任务id

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type UpdateCfsFileSystemNameRequest struct {
	*tchttp.BaseRequest

	// 旧版本用户自定义文件系统名称，优先级低于FsName

	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 用户自定义文件系统名称,与CreationToken 两者必须填一项

	FsName *string `json:"FsName,omitempty" name:"FsName"`
}

func (r *UpdateCfsFileSystemNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCfsFileSystemNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PeriodStatisticsSet struct {

	// 标准型文件系统数量

	StandardFilesystemCount *uint64 `json:"StandardFilesystemCount,omitempty" name:"StandardFilesystemCount"`
	// 性能型文件系统数量

	HighPerformanceFilesystemCount *uint64 `json:"HighPerformanceFilesystemCount,omitempty" name:"HighPerformanceFilesystemCount"`
	// 标准型文件系统存储量

	StandardFilesystemStorage *float64 `json:"StandardFilesystemStorage,omitempty" name:"StandardFilesystemStorage"`
	// 性能型文件系统存储量

	HighPerformanceFilesystemStorage *float64 `json:"HighPerformanceFilesystemStorage,omitempty" name:"HighPerformanceFilesystemStorage"`
	// 统计周期

	StatisticalPeriod *string `json:"StatisticalPeriod,omitempty" name:"StatisticalPeriod"`
}

type CreateCfsFileSystemRequest struct {
	*tchttp.BaseRequest

	// 可用区 ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
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
	// 权限组 ID

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
	// 文件系统是否加密，若留空则默认为不加密

	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`
	// 加密密钥 ID

	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`
	// cfs资源池id

	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
	// 快照id

	SnapId *string `json:"SnapId,omitempty" name:"SnapId"`
	// 标签

	ResourceTags []*ResourceTags `json:"ResourceTags,omitempty" name:"ResourceTags"`
	// 项目id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// CFS 文件系统版本

	CfsVersion *string `json:"CfsVersion,omitempty" name:"CfsVersion"`
	// TURBO文件系统的容量（TiB）

	Capacity *uint64 `json:"Capacity,omitempty" name:"Capacity"`
}

func (r *CreateCfsFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsFileSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsFileSystemClientsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客户端列表

		ClientList []*FileSystemClient `json:"ClientList,omitempty" name:"ClientList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfsFileSystemClientsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsFileSystemClientsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InputPermissionGroupRules struct {

	// 允许访问的客户端IP

	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`
	// 读写权限, ro为只读，rw为读写

	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`
	// 用户权限。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。

	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`
	// 规则优先级，1-100。 其中 1 为最高，100为最低

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

type MountInfo struct {

	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 挂载点 ID

	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`
	// 挂载点 IP

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 挂载根目录

	FSID *string `json:"FSID,omitempty" name:"FSID"`
	// 挂载点状态

	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`
	// 网络类型

	NetworkInterface *string `json:"NetworkInterface,omitempty" name:"NetworkInterface"`
	// 私有网络 ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 子网 Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 权限组名称

	PGroupName *string `json:"PGroupName,omitempty" name:"PGroupName"`
}

type DescribeCfsFileSystemClientsRequest struct {
	*tchttp.BaseRequest

	// 文件系统 ID。

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeCfsFileSystemClientsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsFileSystemClientsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsServiceStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该用户当前 CFS 服务的状态，none 为未开通，creating 为开通中，created 为已开通

		CfsServiceStatus *string `json:"CfsServiceStatus,omitempty" name:"CfsServiceStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfsServiceStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SignUpCfsServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该用户当前 CFS 服务的状态，none 是未开通，creating 是开通中，created 是已开通

		CfsServiceStatus *string `json:"CfsServiceStatus,omitempty" name:"CfsServiceStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SignUpCfsServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SignUpCfsServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 变量

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type PGroup struct {

	// 权限组ID

	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
	// 权限组名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeCfsTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// tag列表

		Tags []*CfsTagUser `json:"Tags,omitempty" name:"Tags"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfsTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsTagsRequest struct {
	*tchttp.BaseRequest

	// 国际化

	InternationalFlag *int64 `json:"InternationalFlag,omitempty" name:"InternationalFlag"`
	// 可用区id

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeCfsTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountTargetCollectionWithRegion struct {

	// 挂载点数量

	NumberOfMountTargets *int64 `json:"NumberOfMountTargets,omitempty" name:"NumberOfMountTargets"`
	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 挂载点详情

	MountTargets []*MountTargetsWithRegion `json:"MountTargets,omitempty" name:"MountTargets"`
}

type DescribeMountTargetsWithRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 挂载点列表

		MountTargetCollection []*MountTargetCollectionWithRegion `json:"MountTargetCollection,omitempty" name:"MountTargetCollection"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMountTargetsWithRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMountTargetsWithRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountTargetsWithRegion struct {

	// 私有网络 ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 挂载点 IP

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
}

type DescribeCfsServiceStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCfsServiceStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsServiceStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMountTargetsWithRegionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMountTargetsWithRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMountTargetsWithRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KmsKey struct {

	// 下次轮转时间戳

	NextRotateTime *int64 `json:"NextRotateTime,omitempty" name:"NextRotateTime"`
	// 密钥属主

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 密钥轮转是否开启

	KeyRotationEnabled *bool `json:"KeyRotationEnabled,omitempty" name:"KeyRotationEnabled"`
	// 密钥ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 密钥别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 创建时间戳

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 密钥描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 密钥状态，可选值为“Enabled“，”Disabled“

	KeyState *string `json:"KeyState,omitempty" name:"KeyState"`
	// 密钥使用方

	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`
	// 密钥类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 创建者UIN

	CreatorUin *int64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
}

type ScaleUpFileSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 扩容大小

		TargetCapacity *uint64 `json:"TargetCapacity,omitempty" name:"TargetCapacity"`
		// fsid

		FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScaleUpFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleUpFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileSystemInfo struct {

	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 用户自定义名称

	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 文件系统状态

	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`
	// 文件系统已使用容量

	SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`
	// 文件系统最大空间限制

	SizeLimit *uint64 `json:"SizeLimit,omitempty" name:"SizeLimit"`
	// 区域 ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 区域名称

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 文件系统协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 文件系统存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 文件系统绑定的预付费存储包（暂未支持）

	StorageResourcePkg *string `json:"StorageResourcePkg,omitempty" name:"StorageResourcePkg"`
	// 文件系统绑定的预付费带宽包（暂未支持）

	BandwidthResourcePkg *string `json:"BandwidthResourcePkg,omitempty" name:"BandwidthResourcePkg"`
	// 文件系统绑定权限组信息

	PGroup *PGroup `json:"PGroup,omitempty" name:"PGroup"`
	// 用户自定义名称

	FsName *string `json:"FsName,omitempty" name:"FsName"`
	// 文件系统是否加密

	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`
	// 加密所使用的密钥，可以为密钥的 ID 或者 ARN

	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`
	// 容量限额最大值

	SizeLimitMax *uint64 `json:"SizeLimitMax,omitempty" name:"SizeLimitMax"`
	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 挂载点IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 已分配空间

	AllocedSpace *uint64 `json:"AllocedSpace,omitempty" name:"AllocedSpace"`
	// cfs资源池tagid

	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
	// cfs资源池tag名称

	TagName *string `json:"TagName,omitempty" name:"TagName"`
	// 快照id

	SnapId *string `json:"SnapId,omitempty" name:"SnapId"`
	// 快照状态

	SnapStatus *string `json:"SnapStatus,omitempty" name:"SnapStatus"`
	// 带宽上限

	BandwidthLimit *float64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type ResourceTags struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeMountTargetsRequest struct {
	*tchttp.BaseRequest

	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeMountTargetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMountTargetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AvailableType struct {

	// 协议与售卖详情

	Protocols []*AvailableProtoStatus `json:"Protocols,omitempty" name:"Protocols"`
	// 存储类型。可选值有 SD 标准型存储、HP性能型存储

	Type *string `json:"Type,omitempty" name:"Type"`
}

type AvailableProtoStatus struct {

	// 售卖状态。可选值有 sale_out 售罄、saling可售、no_saling不可销售

	SaleStatus *string `json:"SaleStatus,omitempty" name:"SaleStatus"`
	// 协议类型。可选值有 NFS、CIFS

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type MigratableCfsFileSystemId struct {

	// 文件系统ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type TagInfo struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}
