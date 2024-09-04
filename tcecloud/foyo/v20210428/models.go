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
	} `json:"Response"`
}

func (r *QuerySwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
