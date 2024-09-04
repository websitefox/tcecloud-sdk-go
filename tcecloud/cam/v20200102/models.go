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

package v20200102

import (
	"encoding/json"

	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type CheckLoginDetail struct {

	// 登陆类型

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 登陆到到地址

	LoginTo *string `json:"LoginTo,omitempty" name:"LoginTo"`
	// 登陆平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 客户端类型

	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`
	// 是否使用多因子认证

	MFAUsed *string `json:"MFAUsed,omitempty" name:"MFAUsed"`
}

type UpdateRelatedRolePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRelatedRolePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRelatedRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUinListByRolePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uin列表

		UinList []*uint64 `json:"UinList,omitempty" name:"UinList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUinListByRolePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUinListByRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupUserInfo struct {

	// 接收者id

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 账户唯一id

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 手机号

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 国家编码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 手机号标识

	PhoneFlag *string `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 邮箱标识

	EmailFlag *string `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// 用户类型

	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否是主账户

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// 账户类型

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
}

type GetUidByUinArrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uin对应的Uid信息

		UinInfo []*UinInfo `json:"UinInfo,omitempty" name:"UinInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUidByUinArrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUidByUinArrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyMenuRequest struct {
	*tchttp.BaseRequest

	// 菜单

	Menu *string `json:"Menu,omitempty" name:"Menu"`
	// 子菜单

	Submenu *string `json:"Submenu,omitempty" name:"Submenu"`
}

func (r *VerifyMenuRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyMenuRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRelatedRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 需关联的uin列表

	TargetUin []*uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 需关联的GroupId列表

	TargetGroupId []*uint64 `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
}

func (r *UpdateRelatedRolePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRelatedRolePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRoleListByOwnerUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRoleListByOwnerUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRoleListByOwnerUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginData struct {

	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 登陆密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

type LoginDetail struct {

	// 登陆账户类型

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 登陆的平台地址

	LoginTo *string `json:"LoginTo,omitempty" name:"LoginTo"`
	// 登陆平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 客户端类型

	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`
}

type UinInfo struct {

	// uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户uid

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type GetUidByUinArrRequest struct {
	*tchttp.BaseRequest

	// uin列表

	UinList []*uint64 `json:"UinList,omitempty" name:"UinList"`
}

func (r *GetUidByUinArrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUidByUinArrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUinListByRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 角色名

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *GetUinListByRolePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUinListByRolePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyMenuResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyMenuResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyMenuResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRoleListByOwnerUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRoleListByOwnerUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRoleListByOwnerUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
