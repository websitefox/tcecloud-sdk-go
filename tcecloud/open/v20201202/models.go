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

package v20201202

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type GetLdapIdpConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetLdapIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLdapIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindWorkWeixinAccountRequest struct {
	*tchttp.BaseRequest

	// 子账号uin

	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`
	// 企业微信userId

	WechatUserId *string `json:"WechatUserId,omitempty" name:"WechatUserId"`
}

func (r *BindWorkWeixinAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindWorkWeixinAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIdentityProviderRequest struct {
	*tchttp.BaseRequest
}

func (r *ListIdentityProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIdentityProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIdentityProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListIdentityProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIdentityProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestLdapRequest struct {
	*tchttp.BaseRequest

	// 认证类型

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// 提供商名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// ldap类型

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// 地址

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// 连接类型

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 证书

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// basedn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// 管理账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理账号密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 过滤条件

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 用户账号字段

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// 邮箱字段

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// 昵称字段

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
	// 电话字段

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// 测试账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 测试账号密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *TestLdapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestLdapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLdapIdpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLdapIdpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLdapIdpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinOpenAppMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 企业微信部门成员信息

		DepartmentList []*DepartmentAttr `json:"DepartmentList,omitempty" name:"DepartmentList"`
		// 企业微信可见成员信息

		VisibleUserList []*WechatUserAttr `json:"VisibleUserList,omitempty" name:"VisibleUserList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkWeixinOpenAppMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkWeixinOpenAppConfigRequest struct {
	*tchttp.BaseRequest

	// 内部应用名

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 内部应用id

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// 内部应用secret

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// 企业id

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
	// 配置id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *UpdateWorkWeixinOpenAppConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkWeixinOpenAppConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinOpenAppConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内部应用id

		AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
		// 内部应用名称

		AppName *string `json:"AppName,omitempty" name:"AppName"`
		// 企业id

		CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
		// 关联时间

		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
		// 应用存储id

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 内部应用秘钥

		Secret *string `json:"Secret,omitempty" name:"Secret"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkWeixinOpenAppConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisabledIdpConfigRequest struct {
	*tchttp.BaseRequest

	// 主键id

	IdpId *int64 `json:"IdpId,omitempty" name:"IdpId"`
}

func (r *DisabledIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisabledIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepartmentAttr struct {

	// 部门id

	DepartmentId *uint64 `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 部门名称

	DepartmentName *string `json:"DepartmentName,omitempty" name:"DepartmentName"`
	// 父部门id

	ParentDepartmentId *uint64 `json:"ParentDepartmentId,omitempty" name:"ParentDepartmentId"`
	// 部门成员列表

	UserList []*WechatUserAttr `json:"UserList,omitempty" name:"UserList"`
}

type CreateLdapIdpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLdapIdpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLdapIdpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkWeixinOpenAppConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWorkWeixinOpenAppConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkWeixinOpenAppConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisabledIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisabledIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisabledIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLdapIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLdapIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLdapIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinOpenAppMemberRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWorkWeixinOpenAppMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProviderList struct {

	// 身份提供商列表

	List []*Item `json:"List,omitempty" name:"List"`
}

type BindWorkWeixinAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindWorkWeixinAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindWorkWeixinAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnabledIdpConfigRequest struct {
	*tchttp.BaseRequest

	// 主键id

	IdpId *int64 `json:"IdpId,omitempty" name:"IdpId"`
	// 用户登录名

	TuserId *string `json:"TuserId,omitempty" name:"TuserId"`
}

func (r *EnabledIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnabledIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkWeixinOpenAppConfigRequest struct {
	*tchttp.BaseRequest

	// 内部应用名

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 内部应用id

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// 内部应用secret

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// 企业id

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
}

func (r *CreateWorkWeixinOpenAppConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkWeixinOpenAppConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinOpenAppConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWorkWeixinOpenAppConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Item struct {

	// 配置 id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	//  配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主账号 uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 身份提供商类型

	ProviderType *int64 `json:"ProviderType,omitempty" name:"ProviderType"`
	// 企业认证状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	//  oauth协议映射

	Oauth *string `json:"Oauth,omitempty" name:"Oauth"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type EnabledIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnabledIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnabledIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestLdapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestLdapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestLdapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLdapIdpRequest struct {
	*tchttp.BaseRequest

	// 认证类型

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// 提供商名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// ldap类型

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// 地址

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// 连接类型

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 证书

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// basedn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// 管理账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理账号密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 过滤条件

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 用户账号字段

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// 邮箱字段

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// 昵称字段

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
	// 电话字段

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否同步 idp 用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

func (r *UpdateLdapIdpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLdapIdpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkWeixinOpenAppConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateWorkWeixinOpenAppConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkWeixinOpenAppConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProviderItem struct {

	//  身份提供商配置

	Item *ProviderData `json:"Item,omitempty" name:"Item"`
}

type WechatUserAttr struct {

	// 企业微信userId

	WechatUserId *string `json:"WechatUserId,omitempty" name:"WechatUserId"`
	// 企业微信用户名

	WechatUserName *string `json:"WechatUserName,omitempty" name:"WechatUserName"`
}

type CreateLdapIdpRequest struct {
	*tchttp.BaseRequest

	// 认证类型

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// 提供商名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// ldap类型

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// 地址

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// 连接类型

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 证书

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// basedn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// 管理账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理账号密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 过滤条件

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 用户账号字段

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// 邮箱字段

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// 昵称字段

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
	// 电话字段

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// 是否同步 idp 用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

func (r *CreateLdapIdpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLdapIdpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProviderData struct {

	// 身份提供商数据

	Data *ProviderList `json:"Data,omitempty" name:"Data"`
}
