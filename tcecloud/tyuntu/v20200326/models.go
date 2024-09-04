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

package v20200326

import (
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type QueryModuleFilterRequest struct {
	*tchttp.BaseRequest

	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页号

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryModuleFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryModuleFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApiInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApiInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApiInfoFilterRequest struct {
	*tchttp.BaseRequest

	// 过滤分页查询字段

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
	// 分页每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页当前页号

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *QueryApiInfoFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApiInfoFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回

		Data *DescribeActionData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDataTypesRequest struct {
	*tchttp.BaseRequest

	// 指定模块名和版本和参数标志

	DataParam *GetApiInfoParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetDataTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunActionRequest struct {
	*tchttp.BaseRequest

	// TargetModule

	TargetModule *string `json:"TargetModule,omitempty" name:"TargetModule"`
	// TargetAction

	TargetAction *string `json:"TargetAction,omitempty" name:"TargetAction"`
	// TargetActionParams

	TargetActionParams *string `json:"TargetActionParams,omitempty" name:"TargetActionParams"`
	// TargetRegion

	TargetRegion *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
	// TargetSecretId

	TargetSecretId *string `json:"TargetSecretId,omitempty" name:"TargetSecretId"`
	// TargetSecretKey

	TargetSecretKey *string `json:"TargetSecretKey,omitempty" name:"TargetSecretKey"`
	// TargetVersion

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
}

func (r *RunActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunActionData struct {

	// response body

	Body *string `json:"Body,omitempty" name:"Body"`
	// request headers

	Headers *string `json:"Headers,omitempty" name:"Headers"`
	// request curl command

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
}

type GetApiInfoRequest struct {
	*tchttp.BaseRequest

	// 获取API详情参数

	DataParam *GetApiInfoParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetApiInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryModuleFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表模块数据

		Data []*ModuleModule `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryModuleFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryModuleFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionData struct {

	// 接口文档，markdown格式

	Document *string `json:"Document,omitempty" name:"Document"`
}

type DescribeSdkDemosData struct {
}

type ModuleModule struct {

	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 权限人

	PermissionUser *string `json:"PermissionUser,omitempty" name:"PermissionUser"`
	// 负责人

	SuperUser *string `json:"SuperUser,omitempty" name:"SuperUser"`
	// url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 行业分类

	CategoryID *int64 `json:"CategoryID,omitempty" name:"CategoryID"`
}

type GetDataTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回模块下数据类型数组

		Data []*DataType `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDataTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApiInfoFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryApiInfoFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApiInfoFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterString struct {

	// key过滤字段

	Name *string `json:"Name,omitempty" name:"Name"`
	// value过滤字段的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetApiInfoParam struct {

	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
}

type DescribeSdkDemosRequest struct {
	*tchttp.BaseRequest

	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 接口名称

	ProductAction *string `json:"ProductAction,omitempty" name:"ProductAction"`
	// 接口版本

	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`
}

func (r *DescribeSdkDemosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSdkDemosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataType struct {

	// 数据类型名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据类型说明

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DescribeActionRequest struct {
	*tchttp.BaseRequest

	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 接口名称

	ProductAction *string `json:"ProductAction,omitempty" name:"ProductAction"`
	// 接口版本

	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`
}

func (r *DescribeActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSdkDemosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回

		Data *DescribeSdkDemosData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSdkDemosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSdkDemosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API Explorer 调用其他接口的返回

		Data *RunActionData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
