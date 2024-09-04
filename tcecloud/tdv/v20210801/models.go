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

package v20210801

import (
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type QueryDataSourceRequest struct {
	*tchttp.BaseRequest

	// 数据源uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
}

func (r *QueryDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 页面列表

		List []*Page `json:"List,omitempty" name:"List"`
		// 环境信息

		TceEnv *bool `json:"TceEnv,omitempty" name:"TceEnv"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filters struct {

	// 搜索key

	Name *string `json:"Name,omitempty" name:"Name"`
	// 搜索值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type QueryDataSouceDataRequest struct {
	*tchttp.BaseRequest

	// 独立数据源Id

	DsId *string `json:"DsId,omitempty" name:"DsId"`
	// 数据源类型，云api: yunApi, 云哨：yunShao ; 中台： middlePlatform

	DsType *string `json:"DsType,omitempty" name:"DsType"`
	// 数据源所引用变量json字符串数据

	VariableList *string `json:"VariableList,omitempty" name:"VariableList"`
}

func (r *QueryDataSouceDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDataSouceDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDataSouceDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDataSouceDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDataSouceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPagesRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否包含配置文件

	InCludeConf *bool `json:"InCludeConf,omitempty" name:"InCludeConf"`
	// 搜索参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 支持国际化

	I18n *bool `json:"I18n,omitempty" name:"I18n"`
}

func (r *QueryPagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Page struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 修改者

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 唯一code

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 名称

	Label *string `json:"Label,omitempty" name:"Label"`
	// 描述详情

	Description *string `json:"Description,omitempty" name:"Description"`
	// 页面类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 页面json

	ComponentConf *string `json:"ComponentConf,omitempty" name:"ComponentConf"`
	// 页面中文名称

	LabelCn *string `json:"LabelCn,omitempty" name:"LabelCn"`
	// 页面英文名称

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
}
