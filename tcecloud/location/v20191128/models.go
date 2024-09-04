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

package v20191128

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ZoneEx struct {

	// 所属Region编号

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// 可用区名称(固定为英文字符串, API使用)

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区编号

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// 可用区名称(国际化支持, 前端展示使用)

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区状态(AVAILABLE/UNAVAILABLE)

	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
	// 可用区状态(原因描述)

	ZoneStateRemark *string `json:"ZoneStateRemark,omitempty" name:"ZoneStateRemark"`
	// 可用区角色（MAZ:主可用区, SAZ:从可用区, OAZ:其它可用区）

	ZoneRole *string `json:"ZoneRole,omitempty" name:"ZoneRole"`
}

type RegionInfo struct {

	// 地域名称，例如，ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域描述，例如，华南地区(广州)

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域是否可用状态

	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用区列表信息

		ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionEx struct {

	// 地域名称(固定为英文字符串, API使用)

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域编号(国际化支持, 前端展示使用)

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域状态(AVAILABLE/UNAVAILABLE)

	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
	// 地域状态描述

	RegionStateRemark *string `json:"RegionStateRemark,omitempty" name:"RegionStateRemark"`
	// 可用区数量

	ZoneCount *int64 `json:"ZoneCount,omitempty" name:"ZoneCount"`
	// 可用区列表

	ZoneSet []*ZoneEx `json:"ZoneSet,omitempty" name:"ZoneSet"`
	// 地域角色(MR:主地域, SR:从地域, OR:其它地域)

	RegionRole *string `json:"RegionRole,omitempty" name:"RegionRole"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域列表信息

		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionZoneRequest struct {
	*tchttp.BaseRequest

	// TCE 产品Code

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// TCE 子产品Code

	SubProductId *string `json:"SubProductId,omitempty" name:"SubProductId"`
	// 待过滤的Region列表

	Regions []*string `json:"Regions,omitempty" name:"Regions"`
}

func (r *DescribeRegionZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域数量

		RegionCount *int64 `json:"RegionCount,omitempty" name:"RegionCount"`
		// 地域信息

		RegionSet []*RegionEx `json:"RegionSet,omitempty" name:"RegionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneInfo struct {

	// 可用区名称，例如，ap-guangzhou-3

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区描述，例如，广州三区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区ID

	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}
