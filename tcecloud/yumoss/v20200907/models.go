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

package v20200907

import (
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type SearchSyncConfigRequest struct {
	*tchttp.BaseRequest

	// 查询数量（用于分页）

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值（用于分页）

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 筛选条件：每个筛选条件为一个Name-Value组合，如果本list为空则代表查全量数据

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *SearchSyncConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSyncConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSyncConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSyncConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSyncConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchYumCatalogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchYumCatalogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchYumCatalogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 查询条件的key

	Name *string `json:"Name,omitempty" name:"Name"`
	// 查询条件的value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type SearchYumCatalogRequest struct {
	*tchttp.BaseRequest

	// 要查询的详细路径，注意要以“/”开头

	Route *string `json:"Route,omitempty" name:"Route"`
	// 查询数量（用于分页）

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值（用于分页）

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchYumCatalogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchYumCatalogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
