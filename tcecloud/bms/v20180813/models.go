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

package v20180813

import (
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type CreateOutbandIPRequest struct {
	*tchttp.BaseRequest

	// BMS实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台]查询；也可以调用接口  [DescribeSubnetEx]，从接口返回中的`unSubnetId`字段获取。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台]查询；也可以调用接口 [DescribeVpcEx] ，从接口返回中的`unVpcId`字段获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 服务器上映射的端口号范围在0-65535，注意避开服务器正在使用的端口号

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 被映射的overlay IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *CreateOutbandIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOutbandIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`]API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 内部参数，释放弹性IP。

	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`
	// 试运行。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回任务的taskid,可通过taskid查询任务运行情况

		TaskId []*uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionTimer struct {

	// 定时器

	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`
	// 执行时间

	ActionTime *string `json:"ActionTime,omitempty" name:"ActionTime"`
	// 扩展数据

	Externals *Externals `json:"Externals,omitempty" name:"Externals"`
}

type OutbandInfos struct {

	// 带外密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 服务器sn

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 带外用户名

	User *string `json:"User,omitempty" name:"User"`
}

type DeleteOutbandIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOutbandIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOutbandIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 分散置放群组名称，长度1-60个字符，支持中、英文。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分散置放群组类型，取值范围：<br><li>RACK：机架层级（跨交换机）<br><li>RACK_SAME_SW：机架层级（同交换机)

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateDisasterRecoverGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDisasterRecoverGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回任务的taskid,可通过taskid查询任务运行情况

		TaskId []*uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateHeartbeatRequest struct {
	*tchttp.BaseRequest

	// 心跳网络组的新名称，长度在1-60

	Name *string `json:"Name,omitempty" name:"Name"`
	// 心跳网络组ID，形如HA-XXX

	HeartbeatId *string `json:"HeartbeatId,omitempty" name:"HeartbeatId"`
}

func (r *UpdateHeartbeatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateHeartbeatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 分散置放群组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 分散置放群组名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *UpdateDisasterRecoverGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDisasterRecoverGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginSettings struct {

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。

	Password *string `json:"Password,omitempty" name:"Password"`
	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。

	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type DescribeOutbandIPRequest struct {
	*tchttp.BaseRequest

	// BMS实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOutbandIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOutbandIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesGroupIdRequest struct {
	*tchttp.BaseRequest

	// 一个待操作的实例ID。可通过[`DescribeInstances`]API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 分散置放群组ID，可通过[DescribeDisasterRecoverGroups]接口获取。

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ModifyInstancesGroupIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesGroupIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnBackupIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReturnBackupIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnBackupIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest

	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。<br><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。<br><li>其他类型（Linux 等）实例：字符长度为[2, 30]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 套餐ID。

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// 系统类型，支持Linux/Windows。

	OperatingSystemType *string `json:"OperatingSystemType,omitempty" name:"OperatingSystemType"`
	// 系统的发行版本号。

	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若在此参数中指定了私有网络ip，InstanceCount参数必须与私有网络ip的个数一致。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 实例数量。

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 实例显示名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例登录设置，通过该参数可以设置实例的登录方式密码。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// RAID类型

	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`
	// 绑定标签参数

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 分散置放群组ID列表，可通过[DescribeDisasterRecoverGroups]接口获取。

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *RunInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`]API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例名称。可任意命名，但不得超过60个字符。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`]API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StartInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDisasterRecoverGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutbandInfo struct {

	// BMS实例ID

	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	// BMS实例名称。

	InstanceName []*string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台]查询；也可以调用接口 [DescribeVpcEx] ，从接口返回中的`unVpcId`字段获取。

	UniqVpcId []*string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 被映射的overlay IP

	Vip []*string `json:"Vip,omitempty" name:"Vip"`
	// 被映射的服务器端口号

	Vport []*int64 `json:"Vport,omitempty" name:"Vport"`
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个硬盘ID查询。硬盘ID形如：`disk-11112222`。参数不支持同时指定`DiskIds`和`Filters`。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 过滤条件。参数不支持同时指定`DiskIds`和`Filters`。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照硬盘ID过滤。硬盘ID形如：`disk-11112222`<br><li>instance-id - Array of String - 是否必填：否 -（过滤条件）按照硬盘挂载的裸金属实例ID过滤。可根据此参数查询挂载在指定机下的硬盘。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOutbandInfoRequest struct {
	*tchttp.BaseRequest

	// BMS实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeOutbandInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOutbandInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`]API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Flavor struct {

	// 套餐ID。

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// 套餐名称。

	FlavorName *string `json:"FlavorName,omitempty" name:"FlavorName"`
	// 实例所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 支持的Raid类型。

	RaidType []*string `json:"RaidType,omitempty" name:"RaidType"`
	// 支持的系统列表。

	OperatingSystem *OperatingSystem `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
	// cpu信息。

	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`
	// 内存信息。

	Memory *string `json:"Memory,omitempty" name:"Memory"`
	// 硬盘信息。

	SystemDisk *string `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 网卡信息

	NetSpeed *string `json:"NetSpeed,omitempty" name:"NetSpeed"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 计费机型

	FlavorType *string `json:"FlavorType,omitempty" name:"FlavorType"`
	// CPU机型，X86/ARM

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 售卖状态

	Soldout *uint64 `json:"Soldout,omitempty" name:"Soldout"`
	// 是否用户自定义机型

	UserDefined *uint64 `json:"UserDefined,omitempty" name:"UserDefined"`
	// 机型网口数

	NetworkPorts *int64 `json:"NetworkPorts,omitempty" name:"NetworkPorts"`
	// dcos后置脚本

	DcosShell *string `json:"DcosShell,omitempty" name:"DcosShell"`
}

type DescribeOutbandIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// BMS映射详细信息

		InstanceSet []*OutbandInfo `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 信息条目总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOutbandIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOutbandIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceBmsInstanceForTradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 单价

		UnitPrice *string `json:"UnitPrice,omitempty" name:"UnitPrice"`
		// 商品的时间单位

		ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
		// 总费用

		OriginalPrice *string `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
		// 优惠后总价

		DiscountPrice *string `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceBmsInstanceForTradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceBmsInstanceForTradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelOutbandInfo struct {

	// BMS实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 被映射端口号

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 被映射的underlay IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type EnhancedService struct {

	// 开启云安全服务。若不指定该参数，则默认不开启云安全服务

	SecurityService *bool `json:"SecurityService,omitempty" name:"SecurityService"`
	// 开启云监控服务。若不指定该参数，则默认不开启云监控服务。

	MonitorService *bool `json:"MonitorService,omitempty" name:"MonitorService"`
	// 开启云哨监控服务。若不指定该参数，则默认不开启云哨监控服务。

	WhistleService *bool `json:"WhistleService,omitempty" name:"WhistleService"`
}

type CreateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDisasterRecoverGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例的详细信息列表

		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`
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

type Placement struct {

	// 实例所属的可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例所属项目ID。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type QueryTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务执行结果

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 任务返回结果的信息

		Output *string `json:"Output,omitempty" name:"Output"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DeleteHeartbeatsRequest struct {
	*tchttp.BaseRequest

	// 删除的心跳网络组ID

	HeartbeatIds []*string `json:"HeartbeatIds,omitempty" name:"HeartbeatIds"`
}

func (r *DeleteHeartbeatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHeartbeatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的硬盘数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 硬盘的详细信息列表。

		DiskSet []*Disk `json:"DiskSet,omitempty" name:"DiskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID。可通过 [DescribeInstances]查看

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// raid类型

	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`
	// 操作系统类型

	OperatingSystemType *string `json:"OperatingSystemType,omitempty" name:"OperatingSystemType"`
	// 操作系统发行版本

	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
	// 实例登录设置。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则关闭开启云监控、云安全服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
}

func (r *ResetInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlavorsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 套餐详情列表

		FlavorSet []*Flavor `json:"FlavorSet,omitempty" name:"FlavorSet"`
		// 套餐总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlavorsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlavorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOutbandIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOutbandIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOutbandIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 置放群组信息列表

		GroupId []*DisasterRecoverGroup `json:"GroupId,omitempty" name:"GroupId"`
		// 用户置放群组总量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisasterRecoverGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：`bms-11112222`。（此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的`id.N`一节）。每次请求的实例的上限为100。参数不支持同时指定`InstanceIds`和`Filters`。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 过滤条件。<li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li><li> instance-id - String - 是否必填：否 - （过滤条件）按照实例ID过滤。实例ID形如：ins-11112222。</li><li> instance-name - String - 是否必填：否 - （过滤条件）按照实例名称过滤。</li><li> instance-state - String - 是否必填：否 -（过滤条件）按照实例的状态过滤 </li><li> private-ip-address - String - 是否必填：否 - （过滤条件）按照实例主网卡的内网IP过滤。</li><li> vpc-id - String - 是否必填：否 - （过滤条件）按照实例所属的vpcid进行过滤 </li><li> subnet-id - String - 是否必填：否 - （过滤条件）按照实例所属的子网id进行过滤 -（过滤条件）按照实例所属的groupid进行过滤 </li><li> groupId - String - 是否必填：否 。</li><li> cpuArch - String - 是否必填：否 -（过滤条件）按照CPU架构过滤。</li><li> operating-system-type - String - 是否必填：否 -（过滤条件）按照操作系统过滤。</li>每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagSpecification struct {

	// 标签绑定的资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 标签对列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateHeartbeatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHeartbeatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHeartbeatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOutbandIPRequest struct {
	*tchttp.BaseRequest

	// 删除映射关系参数

	OutbandInfos []*DelOutbandInfo `json:"OutbandInfos,omitempty" name:"OutbandInfos"`
}

func (r *DeleteOutbandIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOutbandIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHeartbeatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHeartbeatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHeartbeatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTaskRequest struct {
	*tchttp.BaseRequest

	// 异步任务请求返回的RequestId。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *QueryTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHeartbeatRequest struct {
	*tchttp.BaseRequest

	// 心跳网络组名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateHeartbeatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHeartbeatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest

	// 分散置放群组ID列表，可通过[DescribeDisasterRecoverGroups]接口获取。每次请求允许操作的分散置放群组数量上限是10。

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

func (r *DeleteDisasterRecoverGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDisasterRecoverGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHeartbeatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 心跳网络组详细信息

		HeartbeatSet []*HeartbeatSet `json:"HeartbeatSet,omitempty" name:"HeartbeatSet"`
		// 心跳网络组条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHeartbeatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHeartbeatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHeartbeatRequest struct {
	*tchttp.BaseRequest

	// BMS服务器ID，形如BMS-xxx

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 心跳网络组ID，形如HA-xxx

	HeartbeatId *string `json:"HeartbeatId,omitempty" name:"HeartbeatId"`
}

func (r *ModifyHeartbeatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHeartbeatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperatingSystem struct {

	// 支持的linux系统列表

	Linux []*string `json:"Linux,omitempty" name:"Linux"`
	// 支持的windows系统列表

	Windows []*string `json:"Windows,omitempty" name:"Windows"`
	// 支持的Other系统

	Other []*string `json:"Other,omitempty" name:"Other"`
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回任务的taskid,可通过taskid查询任务运行情况

		TaskId []*uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualPrivateCloud struct {

	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台]查询；也可以调用接口 [DescribeVpcEx] ，从接口返回中的`unVpcId`字段获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台]查询；也可以调用接口  [DescribeSubnetEx]，从接口返回中的`unSubnetId`字段获取。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 是否分配IPv6地址

	Ipv6Address *bool `json:"Ipv6Address,omitempty" name:"Ipv6Address"`
}

type DescribeHeartbeatsRequest struct {
	*tchttp.BaseRequest

	// 心跳网络组ID

	HeartbeatIds []*string `json:"HeartbeatIds,omitempty" name:"HeartbeatIds"`
	// 过滤条件。<li> Name - String - 是否必填：否 -（过滤条件）按照名称过滤。</li><li> VlanId - String - 是否必填：否 -（过滤条件）按照vlanid过滤。</li>每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`HeartbeatIds`和`Filters`。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeHeartbeatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHeartbeatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateHeartbeatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateHeartbeatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateHeartbeatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOutbandInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 带外信息详情

		OutbandInfo []*OutbandInfos `json:"OutbandInfo,omitempty" name:"OutbandInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOutbandInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOutbandInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回任务的taskid,可通过taskid查询任务运行情况

		TaskId []*uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupIPRequest struct {
	*tchttp.BaseRequest

	// BMS服务器的ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指定备用IP

	BackupIP *string `json:"BackupIP,omitempty" name:"BackupIP"`
}

func (r *CreateBackupIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlavorsRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：`flavor-11112222`,每次请求的实例的上限为100。参数不支持同时指定`FlavorIds`和`Filters`。

	FlavorIds []*string `json:"FlavorIds,omitempty" name:"FlavorIds"`
	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>
	// <li> flavor-id - String - 是否必填：否 - （过滤条件）按照实例ID过滤。实例ID形如：flavor-11112222。</li>
	// <li> flavor-name - String - 是否必填：否 - （过滤条件）按照实例名称过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`FlavorIds`和`Filters`。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，指定Offset时，必须同时指定Limit。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFlavorsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlavorsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 实例`ID`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// Raid类型。

	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`
	// 操作系统类型。

	OperatingSystemType *string `json:"OperatingSystemType,omitempty" name:"OperatingSystemType"`
	// 操作系统发行版本

	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
	// 实例主网卡的内网`IP`列表。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 实例所属虚拟私有网络信息。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 套餐信息。

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 实例状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// CPU类型，支持X86/ARM

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 实例标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// 实例用户uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 实例用户AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 是否自定义机型

	UserDefined *uint64 `json:"UserDefined,omitempty" name:"UserDefined"`
}

type InternetAccessible struct {

	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。

	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
	// 取值范围{"CTCC":电信，”CUCC“：联通，”CMCC“：移动， ”BGP“：外网CAP}

	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
}

type DescribeDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest

	// 置放群组id列表。

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
	// 过滤条件。<li> Name - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`GroupIds`和`Filters`。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDisasterRecoverGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisasterRecoverGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回任务的taskid,可通过taskid查询任务运行情况

		TaskId []*uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Externals struct {
}

type HeartbeatSet struct {

	// 心跳网络组ID

	HeartbeatId *string `json:"HeartbeatId,omitempty" name:"HeartbeatId"`
	// 心跳网络组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 心跳网络组所属VLAN

	VlanId *string `json:"VlanId,omitempty" name:"VlanId"`
	// 绑定此心跳网络的BMS资源数

	CurrentNum *string `json:"CurrentNum,omitempty" name:"CurrentNum"`
	// 心跳网络组创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 心跳网络组更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DeleteHeartbeatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHeartbeatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHeartbeatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回改任务的taskid,可通过taskid查询任务运行情况

		TaskId []*string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDisasterRecoverGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDisasterRecoverGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesGroupIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesGroupIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesGroupIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`]API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StopInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisasterRecoverGroup struct {

	// 分散置放群组id。

	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`
	// 分散置放群组名称，长度1-60个字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分散置放群组类型，取值范围：<br><li>RACK：机架层级（跨交换机）<br><li>RACK_SAME_SW：机架层级（同交换机)

	Type *string `json:"Type,omitempty" name:"Type"`
	// 分散置放群组内bms服务器当前数量。

	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`
	// 分散置放群组创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 分散置放群组更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type InquiryPriceBmsInstanceForTradeRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品编码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费机型

	BmsFlavor *string `json:"BmsFlavor,omitempty" name:"BmsFlavor"`
}

func (r *InquiryPriceBmsInstanceForTradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceBmsInstanceForTradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnBackupIPRequest struct {
	*tchttp.BaseRequest

	// BMS服务器的ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 要退还的备用IP，如果OverlayIp存在，那么是要退还的主IP

	BackupIP *string `json:"BackupIP,omitempty" name:"BackupIP"`
	// 新的主IP

	OverlayIp *string `json:"OverlayIp,omitempty" name:"OverlayIp"`
}

func (r *ReturnBackupIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnBackupIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Disk struct {

	// 硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 硬盘所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 硬盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 硬盘大小。

	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`
	// 硬盘挂载的云主机ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：SSD表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 硬盘的创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}
