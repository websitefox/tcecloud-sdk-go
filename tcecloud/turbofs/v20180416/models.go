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

package v20180416

import (
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type MountTargetCollectionWithRegion struct {

	// 挂载点数量

	NumberOfMountTargets *int64 `json:"NumberOfMountTargets,omitempty" name:"NumberOfMountTargets"`
	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 挂载点详情

	MountTargets []*MountTargetsWithRegion `json:"MountTargets,omitempty" name:"MountTargets"`
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

type CreateCfsFileSystemRequest struct {
	*tchttp.BaseRequest

	// 可用区 ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区缩写

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 用户自定义文件系统名称

	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
	// 文件系统协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 文件系统存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 网络类型

	NetInterface *string `json:"NetInterface,omitempty" name:"NetInterface"`
	// 权限组 ID

	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
	// VPC ID

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 系统分配的VPC统一ID

	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
	// 子网 ID

	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 系统分配的子网统一 ID

	UnSubnetId *string `json:"UnSubnetId,omitempty" name:"UnSubnetId"`
	// 指定IP地址，仅VPC网络支持

	MountIP *string `json:"MountIP,omitempty" name:"MountIP"`
	// 文件系统绑定的存储包，每个文件系统只能绑定一个。低频文件系统该字段为必填

	StorageResourcePkgId *string `json:"StorageResourcePkgId,omitempty" name:"StorageResourcePkgId"`
	// 文件系统绑定的带宽包，每个文件系统只能绑定一个

	BandwidthResourcePkgId *string `json:"BandwidthResourcePkgId,omitempty" name:"BandwidthResourcePkgId"`
	// 用户自定义文件系统名称，优先级优于CreationToken

	FsName *string `json:"FsName,omitempty" name:"FsName"`
	// 文件系统协议类型，输入值 NFS，CIFS；不填默认为 NFS

	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`
	// 文件系统是否加密

	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`
}

func (r *CreateCfsFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsFileSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionCtrlProtoStatus struct {

	// 售卖状态。可选值有“sale_out”：售罄、“saling”：可售、“no_saling”：不可销售

	SaleStatus *string `json:"SaleStatus,omitempty" name:"SaleStatus"`
	// 协议类型。可选值有“NFS”、“CIFS”

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type UpdateCfsFileSystemPGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 权限组ID

		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
		// 文件系统ID

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

type ZoneSet struct {

	// 可用区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区名称

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 存储类型。可选值有”SD“：标准、“EP”：极速型、“HP”：高性能、“IA”：低频

	Type *string `json:"Type,omitempty" name:"Type"`
	// 协议类型。可选值有“NFS”、“CIFS”

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 售卖状态。可选值有“sale_out”：售罄、“saling”：可售、“no_saling”：不可销售、“reserved”：保留。

	SaleStatus *string `json:"SaleStatus,omitempty" name:"SaleStatus"`
	// 区域名称，如“ap-beijing”

	Region *string `json:"Region,omitempty" name:"Region"`
	// 区域名称，如“bj”

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type SizeInBytes struct {

	// 时间戳

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// 文件系统用量

	Value *string `json:"Value,omitempty" name:"Value"`
}

type VersionCtrlType struct {

	// 协议与售卖详情

	Protocols []*VersionCtrlProtoStatus `json:"Protocols,omitempty" name:"Protocols"`
	// 存储类型。可选值有”SD“：标准、“EP”：极速型、“HP”：高性能、“IA”：低频

	Type *string `json:"Type,omitempty" name:"Type"`
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

type VersionCtrlRegionZones struct {

	// 区域可用情况

	RegionZones []*VersionCtrlRegion `json:"RegionZones,omitempty" name:"RegionZones"`
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

type FileSystemInfo struct {

	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 用户自定义名称

	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
	// 文件系统ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 文件系统状态

	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`
	// 文件系统已使用容量

	SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`
	// 文件系统最大空间限制

	SizeLimit *uint64 `json:"SizeLimit,omitempty" name:"SizeLimit"`
	// 区域ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 区域名称

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 文件系统协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 文件系统存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 文件系统绑定的存储包

	StorageResourcePkg *string `json:"StorageResourcePkg,omitempty" name:"StorageResourcePkg"`
	// 文件系统绑定的带宽包

	BandwidthResourcePkg *string `json:"BandwidthResourcePkg,omitempty" name:"BandwidthResourcePkg"`
	// 文件系统绑定权限组信息

	PGroup *PGroup `json:"PGroup,omitempty" name:"PGroup"`
	// 用户自定义名称

	FsName *string `json:"FsName,omitempty" name:"FsName"`
	// 文件系统是否加密

	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`
	// 加密所使用的密钥，可以为密钥的ID或者ARN

	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`
}

type PGroupRules struct {

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

type CreateCfsRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则ID

		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
		// 权限组ID

		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
		// 客户端IP

		AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`
		// 读写权限

		RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`
		// 用户权限

		UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`
		// 优先级

		Priority *int64 `json:"Priority,omitempty" name:"Priority"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCfsRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountTargetCollection struct {

	// 挂载点数量

	NumberOfMountTargets *int64 `json:"NumberOfMountTargets,omitempty" name:"NumberOfMountTargets"`
	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 挂载点详情

	MountTargets []*MountInfo `json:"MountTargets,omitempty" name:"MountTargets"`
}

type CreateCfsPGroupRequest struct {
	*tchttp.BaseRequest

	// 权限组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 权限组描述信息

	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
}

func (r *CreateCfsPGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsPGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SignUpCfsServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该用户当前CFS服务的状态，none是未开通，creating是开通中，created是已开通

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
}

type CfsKmsKeys struct {

	// 下次轮转时间戳

	NextRotateTime *int64 `json:"NextRotateTime,omitempty" name:"NextRotateTime"`
	// 密钥属主

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 密钥轮转是否开启

	KeyRotationEnabled *bool `json:"KeyRotationEnabled,omitempty" name:"KeyRotationEnabled"`
	// Key的ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 密钥别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 创建时间戳

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 密钥描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Key状态，可选值为“Enabled“，”Disabled“

	KeyState *string `json:"KeyState,omitempty" name:"KeyState"`
	// 密钥使用方

	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`
	// Key类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 创建者UIN

	CreatorUin *int64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
}

type CreateCfsPGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 权限组ID

		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
		// 权限组名字

		Name *string `json:"Name,omitempty" name:"Name"`
		// 权限组描述信息

		DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
		// 权限组关联文件系统个数

		BindCfsNum *int64 `json:"BindCfsNum,omitempty" name:"BindCfsNum"`
		// 权限组创建时间

		CDate *string `json:"CDate,omitempty" name:"CDate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCfsPGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsFileSystemNameRequest struct {
	*tchttp.BaseRequest

	// 用户自定义文件系统名称

	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
	// 文件系统ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 用户自定义文件系统名称，优先级优于CreationToken

	FsName *string `json:"FsName,omitempty" name:"FsName"`
}

func (r *UpdateCfsFileSystemNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCfsFileSystemNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PGroups struct {

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

type UpdateCfsFileSystemPGroupRequest struct {
	*tchttp.BaseRequest

	// 权限组ID

	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
	// 文件系统ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *UpdateCfsFileSystemPGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCfsFileSystemPGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionCtrlZone struct {

	// 可用区名称

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区中文名称

	ZoneCnName *string `json:"ZoneCnName,omitempty" name:"ZoneCnName"`
	// Type数组

	Types []*VersionCtrlType `json:"Types,omitempty" name:"Types"`
}

type FileSystems struct {

	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 用户自定义名称

	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`
	// 文件系统ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	// 文进系统状态

	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`
	// 文进系统已使用容量

	SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`
	// 文件系统最大空间限制

	SizeLimit *uint64 `json:"SizeLimit,omitempty" name:"SizeLimit"`
	// 区域ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 区域名称

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 文件系统协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 文件系统存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 文件系统绑定的存储包

	StorageResourcePkg *string `json:"StorageResourcePkg,omitempty" name:"StorageResourcePkg"`
	// 文件系统绑定的带宽包

	BandwidthResourcePkg *string `json:"BandwidthResourcePkg,omitempty" name:"BandwidthResourcePkg"`
	// 用户自定义名称

	FsName *string `json:"FsName,omitempty" name:"FsName"`
}

type PGroup struct {

	// 权限组ID

	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
	// 权限组名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateCfsRuleRequest struct {
	*tchttp.BaseRequest

	// 权限组ID

	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
	// 读写权限, 可选参数：ro, rw。ro为只读，rw为读写，不填默认为读写

	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`
	// 用户权限，可选参数：all_squash，no_all_squash，root_squash，no_root_squash。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。不填默认为no_root_squash。

	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`
	// 允许访问的客户端IP地址或地址段

	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`
	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *CreateCfsRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionCtrlRegion struct {

	// 区域名称，如“ap-beijing”

	Region *string `json:"Region,omitempty" name:"Region"`
	// 区域名称，如“bj”

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 区域可用情况，当区域内至少有一个可用区处于可售状态时，取值为AVAILABLE，否则为UNAVAILABLE

	RegionStatus *string `json:"RegionStatus,omitempty" name:"RegionStatus"`
	// 可用区数组

	Zones []*VersionCtrlZone `json:"Zones,omitempty" name:"Zones"`
}

type MountTargetsWithRegion struct {

	// 私有网络 ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 挂载点 IP

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
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

type TagRows struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
	// 文件系统ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type CreateCfsFileSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建时间

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
