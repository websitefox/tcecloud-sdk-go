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

package v20191025

import (
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ModifyVpcDnsRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcDnsRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainListFilters struct {

	// 过滤类型

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type RecordListFilters struct {

	// 过滤类型

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type CreateVpcDnsDomainRemarkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcDnsDomainRemarkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsDomainRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
	// 记录ID，逗号分隔

	RecordIds *string `json:"RecordIds,omitempty" name:"RecordIds"`
}

func (r *DeleteVpcDnsRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainDetail struct {

	// 域名id

	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
	// 域名所有者uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 创建时间

	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
	// 最后修改时间

	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
	// 域名记录数量

	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`
	// 域名备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否开启子域名递归（“ENABLED”和“DISABLED”）

	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
	// VPC信息

	VpcInfos []*VpcInfos `json:"VpcInfos,omitempty" name:"VpcInfos"`
}

type RecordId struct {

	// 记录id

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
}

type VpcInfos struct {

	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// RegionId

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// UnVpcId

	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
}

type RecordDetail struct {

	// 记录ID

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
	// 记录对应的域名ID

	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
	// 主机记录

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 记录类型

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 记录值

	Value *string `json:"Value,omitempty" name:"Value"`
	// TTL

	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`
	// MX优先级

	Mx *int64 `json:"Mx,omitempty" name:"Mx"`
	// 记录是否可用，0不可用，1可用

	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`
	// 记录状态，启用，暂停等

	Status *string `json:"Status,omitempty" name:"Status"`
	// 其他信息

	Extra *string `json:"Extra,omitempty" name:"Extra"`
	// 记录创建时间

	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
	// 记录最后修改时间

	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
	// 记录权重

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type CreateVpcDnsDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcDnsDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainCountInfo struct {

	// 全部域名数量

	AllTotal *int64 `json:"AllTotal,omitempty" name:"AllTotal"`
	// 返回的域名数量

	DomainTotal *int64 `json:"DomainTotal,omitempty" name:"DomainTotal"`
}

type CreateVpcDnsRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 记录Id

		Data *RecordId `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcDnsRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcDnsDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindVpcDnsDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindVpcDnsDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindVpcDnsDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名ID，以逗号分隔

	DomainIds *string `json:"DomainIds,omitempty" name:"DomainIds"`
}

func (r *DeleteVpcDnsDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcDnsRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
	// 子域名

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 记录类型

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 记录值

	Value *string `json:"Value,omitempty" name:"Value"`
	// MX优先级

	Mx *uint64 `json:"Mx,omitempty" name:"Mx"`
	// 权重

	Weight *string `json:"Weight,omitempty" name:"Weight"`
}

func (r *CreateVpcDnsRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsDomainRemarkRequest struct {
	*tchttp.BaseRequest

	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateVpcDnsDomainRemarkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsDomainRemarkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsDomainListRequest struct {
	*tchttp.BaseRequest

	// 长度

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤

	Filters []*DomainListFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVpcDnsDomainListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsDomainListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsRecordListRequest struct {
	*tchttp.BaseRequest

	// 长度

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤

	Filters []*RecordListFilters `json:"Filters,omitempty" name:"Filters"`
	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeVpcDnsRecordListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsRecordListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MessageDefault struct {

	// 返回码，正常为0

	Code *string `json:"Code,omitempty" name:"Code"`
	// 操作码说明

	Message *string `json:"Message,omitempty" name:"Message"`
}

type BindVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
	// VPC信息

	VpcInfos []*VpcInfos `json:"VpcInfos,omitempty" name:"VpcInfos"`
}

func (r *BindVpcDnsDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindVpcDnsDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 标签数组

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 是否开启子域名递归，只接受"ENABLED”和“DISABLED”

	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
}

func (r *CreateVpcDnsDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordListCountInfo struct {

	// 该域名的记录总数量

	AllTotal *int64 `json:"AllTotal,omitempty" name:"AllTotal"`
	// 本次返回的记录数量

	RecordTotal *int64 `json:"RecordTotal,omitempty" name:"RecordTotal"`
}

type DescribeVpcDnsDomainListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 域名数量信息

		Info *DomainCountInfo `json:"Info,omitempty" name:"Info"`
		// 域名信息

		Domains []*DomainDetail `json:"Domains,omitempty" name:"Domains"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcDnsDomainListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsRecordListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 记录数量描述

		Info *RecordListCountInfo `json:"Info,omitempty" name:"Info"`
		// 解析记录信息

		Records []*RecordDetail `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcDnsRecordListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
	// 记录ID

	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`
	// 子域名

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 记录类型

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 记录值

	Value *string `json:"Value,omitempty" name:"Value"`
	// MX优先级

	Mx *uint64 `json:"Mx,omitempty" name:"Mx"`
	// 权重

	Weight *string `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyVpcDnsRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Record struct {

	// 记录id

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
}

type DeleteVpcDnsDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcDnsDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名ID，以逗号分割

	DomainIds *string `json:"DomainIds,omitempty" name:"DomainIds"`
	// 是否开启子域名递归，只接受"ENABLED"和"DISABLED"两种

	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
}

func (r *ModifyVpcDnsDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
