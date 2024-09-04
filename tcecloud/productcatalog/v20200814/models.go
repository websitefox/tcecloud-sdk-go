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

package v20200814

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type VerifyUserProductCatalogsRequest struct {
	*tchttp.BaseRequest

	// 产品Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *VerifyUserProductCatalogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyUserProductCatalogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsersProductCatalogsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUsersProductCatalogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsersProductCatalogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsersProductCatalogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户产品目录

		DataList []*string `json:"DataList,omitempty" name:"DataList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsersProductCatalogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsersProductCatalogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyUserProductCatalogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyUserProductCatalogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyUserProductCatalogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductCatalog struct {

	// 产品Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}
