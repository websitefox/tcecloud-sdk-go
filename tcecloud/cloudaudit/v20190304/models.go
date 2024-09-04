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

package v20190304

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type LookupEventsRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 每次最大查询数量

	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询分页

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 查询用户uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 查询属性细节

	LookupAttributes []*Attr `json:"LookupAttributes,omitempty" name:"LookupAttributes"`
	// 搜索类型

	LookupType *string `json:"LookupType,omitempty" name:"LookupType"`
	// 全局内容模糊查询

	ContentValue *string `json:"ContentValue,omitempty" name:"ContentValue"`
}

func (r *LookupEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LookupEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExporJobRequest struct {
	*tchttp.BaseRequest

	// IDs

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteExporJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExporJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExporJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExporJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExporJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Events struct {

	// 事件内容

	CloudAuditEvent *string `json:"CloudAuditEvent,omitempty" name:"CloudAuditEvent"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 事件名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 事件时间

	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`
	// SecretId

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 错误码

	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 请求ID

	RequestID *uint64 `json:"RequestID,omitempty" name:"RequestID"`
	// 账户ID

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
	// 源ip地址

	SourceIPAddress *string `json:"SourceIPAddress,omitempty" name:"SourceIPAddress"`
	// 事件源

	EventSource *string `json:"EventSource,omitempty" name:"EventSource"`
	// 事件域

	EventRegion *string `json:"EventRegion,omitempty" name:"EventRegion"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 资源内容

	Resources *Resources `json:"Resources,omitempty" name:"Resources"`
}

type GetExporLoggingListRequest struct {
	*tchttp.BaseRequest

	// 页数

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页显示数量

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *GetExporLoggingListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetExporLoggingListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetExporLoggingListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetExporLoggingListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetExporLoggingListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBucketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBucketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBucketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Attribute struct {

	// 查询属性

	AttributeKey *string `json:"AttributeKey,omitempty" name:"AttributeKey"`
	// 查询内容

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

type Data struct {

	// 事件

	Events []*Events `json:"Events,omitempty" name:"Events"`
	// 分页

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 分页是否结束

	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
}

type Resources struct {

	// 资源名字

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type Attr struct {

	// 查询属性

	AttributeKey *string `json:"AttributeKey,omitempty" name:"AttributeKey"`
	// 查询内容

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

type CreateBucketRequest struct {
	*tchttp.BaseRequest

	// 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶名称

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
}

func (r *CreateBucketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBucketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatExporJobRequest struct {
	*tchttp.BaseRequest

	// 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶名称

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *CreatExporJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatExporJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatExporJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatExporJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatExporJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LookupEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LookupEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LookupEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
