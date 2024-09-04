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

package v20200716

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type GetAnnounceClassRequest struct {
	*tchttp.BaseRequest

	// 状态 可传“1”

	Status *string `json:"Status,omitempty" name:"Status"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认返回所有

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetAnnounceClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告标题

		Toptic *string `json:"Toptic,omitempty" name:"Toptic"`
		// 地域id列表，多个用;号隔开

		CityIds *string `json:"CityIds,omitempty" name:"CityIds"`
		// 公告分类id

		ClassId *string `json:"ClassId,omitempty" name:"ClassId"`
		// 公告状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 生效起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 生效结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 发布时间

		PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
		// 公告内容

		Content *string `json:"Content,omitempty" name:"Content"`
		// 图片文件标识

		ImgMark []*string `json:"ImgMark,omitempty" name:"ImgMark"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAnnounceContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Annclass struct {

	// 公告分类中文名称

	NameCn *string `json:"NameCn,omitempty" name:"NameCn"`
	// 公告分类级别

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 公告分类唯一id

	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type GetAnnounceContentRequest struct {
	*tchttp.BaseRequest

	// 公告唯一ID

	AnnId *uint64 `json:"AnnId,omitempty" name:"AnnId"`
	// 状态，可传 “1”

	Status *string `json:"Status,omitempty" name:"Status"`
	// 可传1，返回生效时间内的公告

	TimeValid *uint64 `json:"TimeValid,omitempty" name:"TimeValid"`
}

func (r *GetAnnounceContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceListRequest struct {
	*tchttp.BaseRequest

	// 状态 可传“1”

	Status *string `json:"Status,omitempty" name:"Status"`
	// 公告分类id

	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
	// 地域id列表，多个用;号隔开

	CityIds *string `json:"CityIds,omitempty" name:"CityIds"`
	// 分类id列表，多个用;号隔开

	ClassIds *string `json:"ClassIds,omitempty" name:"ClassIds"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 显示条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索条件，按公告标题搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 可传1，返回生效时间内的公告

	TimeValid *uint64 `json:"TimeValid,omitempty" name:"TimeValid"`
}

func (r *GetAnnounceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告分类列表数组

		AnnclassSet []*Annclass `json:"AnnclassSet,omitempty" name:"AnnclassSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAnnounceClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告列表，默认按照发布时间倒序排序

		AnnounceSet []*Announce `json:"AnnounceSet,omitempty" name:"AnnounceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAnnounceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域信息列表数组

		HostAreaSet []*HostArea `json:"HostAreaSet,omitempty" name:"HostAreaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Announce struct {

	// 公告唯一id

	AnnId *uint64 `json:"AnnId,omitempty" name:"AnnId"`
	// 公告标题

	Toptic *string `json:"Toptic,omitempty" name:"Toptic"`
	// 地域id列表，多个用;号隔开

	CityIds *string `json:"CityIds,omitempty" name:"CityIds"`
	// 公告分类id

	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
	// 公告状态信息

	Status *string `json:"Status,omitempty" name:"Status"`
	// 生效起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 生效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 发布时间

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 图片文件标识

	ImgMark *string `json:"ImgMark,omitempty" name:"ImgMark"`
}

type HostArea struct {

	// 地域唯一id

	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
	// 地域中文名称

	CityCn *string `json:"CityCn,omitempty" name:"CityCn"`
	// 所在国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// 所有区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type GetHostAreaRequest struct {
	*tchttp.BaseRequest

	// 状态，可传 “1”

	Status *string `json:"Status,omitempty" name:"Status"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认返回所有

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetHostAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
