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

package v20210428

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type Action1Request struct {
	*tchttp.BaseRequest
}

func (r *Action1Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *Action1Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySwitchRequest struct {
	*tchttp.BaseRequest

	// input

	Params *QuerySwitchParams `json:"Params,omitempty" name:"Params"`
}

func (r *QuerySwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySwitchParams struct {

	// key

	Key *string `json:"Key,omitempty" name:"Key"`
	// ns

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type Action1Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *Action1Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *Action1Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

		// 开关数据信息
		Data *SwitchData `json:"Data,omitempty" name:"Data"`
	} `json:"Response"`
}

// SwitchData 开关数据详情
type SwitchData struct {
	// 条目ID
	EntryID int64 `json:"EntryID,omitempty" name:"EntryID"`

	// 开关键名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 原始值（字符串形式）
	RawValue *string `json:"RawValue,omitempty" name:"RawValue"`

	// 解析后的值
	Value []int `json:"Value,omitempty" name:"Value"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 版本号
	Version int64 `json:"Version,omitempty" name:"Version"`

	// 修改时间
	ModifyTime int64 `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 来源
	Source *string `json:"Source,omitempty" name:"Source"`

	// 扩展信息
	Ext *SwitchExtInfo `json:"Ext,omitempty" name:"Ext"`
}

// SwitchExtInfo 开关扩展信息
type SwitchExtInfo struct {
	// 标签信息
	Label *SwitchLabelInfo `json:"Label,omitempty" name:"Label"`

	// 值类型
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// 扩展版本
	Version *string `json:"Version,omitempty" name:"Version"`
}

// SwitchLabelInfo 开关标签信息
type SwitchLabelInfo struct {
	// 英文标签
	EnUS *string `json:"en-US,omitempty" name:"en-US"`

	// 中文标签
	ZhCN *string `json:"zh-CN,omitempty" name:"zh-CN"`
}

func (r *QuerySwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
