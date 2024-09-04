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

package v20190116

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ListGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组总数。

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 用户组数组信息。

		GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiKeyDetail struct {

	// 密钥ID

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 密钥Key

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// 创建时间(时间戳)

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 状态(2:有效, 3:禁用)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 默认0

	Source *uint64 `json:"Source,omitempty" name:"Source"`
	// 描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type CreateApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 持久密钥

		IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachRolesPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachRolesPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachRolesPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceRoleInfoRequest struct {
	*tchttp.BaseRequest

	// 角色名

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 策略名列表

	PolicyName []*string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *GetServiceRoleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceRoleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAttachedUserAllPoliciesRequest struct {
	*tchttp.BaseRequest

	// 目标用户ID

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 0:返回直接关联和随组关联策略，1:只返回直接关联策略，2:只返回随组关联策略

	AttachType *uint64 `json:"AttachType,omitempty" name:"AttachType"`
	// 策略类型

	StrategyType *uint64 `json:"StrategyType,omitempty" name:"StrategyType"`
	// 每页数量，默认值是 20，必须大于 0 且小于或等于 200

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 页码，默认值是 1，从 1开始，不能大于 200

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 过滤条件user or group

	Filter *string `json:"Filter,omitempty" name:"Filter"`
}

func (r *ListAttachedUserAllPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAttachedUserAllPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedStrategyInfo struct {

	// 策略ID。

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称。

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略创建时间。

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 创建来源，1 通过控制台创建, 2 通过策略语法创建。

	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
	// 策略描述。

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeAttachedEntityPoliciesRequest struct {
	*tchttp.BaseRequest

	// 目标用户ID (当Type=1时必传)

	TargetUin []*uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 用户组ID (当Type=2时必传)

	GroupId []*uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 查询类型(1:用户，2:用户组)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 每页数量，默认值是 20，必须大于 0 且小于或等于 200

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 页码，默认值是 1，从 1开始，不能大于 200

	Page *uint64 `json:"Page,omitempty" name:"Page"`
}

func (r *DescribeAttachedEntityPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttachedEntityPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachUserPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachUserPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachUserPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServicePermListRequest struct {
	*tchttp.BaseRequest

	// 业务标识 如：cvm

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 接口名

	InterfaceEnName *string `json:"InterfaceEnName,omitempty" name:"InterfaceEnName"`
	// 排序字段 默认updateTime

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段desc/asc  默认desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *GetServicePermListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServicePermListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserGroupListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUserGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserGroupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAttributeValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAttributeValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAttributeValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSubAccountRequest struct {
	*tchttp.BaseRequest

	// 用户信息

	UserInfo []*UserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 是否在白名单

	InWhiteList *bool `json:"InWhiteList,omitempty" name:"InWhiteList"`
	// 从Api创建

	FromAPI *int64 `json:"FromAPI,omitempty" name:"FromAPI"`
}

func (r *AddSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachGroupsPolicyRequest struct {
	*tchttp.BaseRequest

	// 用户组ID list

	GroupId []*uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略ID

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DetachGroupsPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachGroupsPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceApiListRequest struct {
	*tchttp.BaseRequest

	// 是否返回api

	IsReturnApi *uint64 `json:"IsReturnApi,omitempty" name:"IsReturnApi"`
}

func (r *GetServiceApiListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceApiListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryKeyBySecretIdRequest struct {
	*tchttp.BaseRequest

	// 秘钥ID

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
	// 平台 (1:表示是云api ; 2:表示是多项目)

	Platform []*uint64 `json:"Platform,omitempty" name:"Platform"`
}

func (r *QueryKeyBySecretIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryKeyBySecretIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PasswordRules struct {

	// 最小密码长度

	MinimumLength *int64 `json:"MinimumLength,omitempty" name:"MinimumLength"`
	// 最少包含

	MustContain *string `json:"MustContain,omitempty" name:"MustContain"`
	// 密码有效期

	ForcePasswordChange *int64 `json:"ForcePasswordChange,omitempty" name:"ForcePasswordChange"`
	// 密码重复次数

	ReusePasswordLimit *int64 `json:"ReusePasswordLimit,omitempty" name:"ReusePasswordLimit"`
	// 登陆最大密码失败次数

	RetryPasswordLimit *int64 `json:"RetryPasswordLimit,omitempty" name:"RetryPasswordLimit"`
	// 是否只有admin可以重置密码

	OnlyAdminCanResetPassword *int64 `json:"OnlyAdminCanResetPassword,omitempty" name:"OnlyAdminCanResetPassword"`
}

type ListGroupsRequest struct {
	*tchttp.BaseRequest

	// 页码。默认为1。

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页数量。默认为20。

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 按用户组名称匹配。

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 用户组类型，0-自定义用户组，1-预设用户组

	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`
}

func (r *ListGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesByActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 策略数组，数组每个成员包括 policyId、policyName、addTime、type、description、 createMode 字段。其中： policyId：策略 id policyName：策略名addTime：策略创建时间type：1 表示自定义策略，2 表示预设策略 description：策略描述 createMode：1 表示按业务权限创建的策略，其他值表示可以查看策略语法和通过策略语法更新策略Attachments: 关联的用户数ServiceType: 策略关联的产品IsAttached: 当需要查询标记实体是否已经关联策略时不为null。0表示未关联策略，1表示已关联策略

		List *StrategyInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPoliciesByActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPoliciesByActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyUserAccessTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 第三方平台openId

		UserOpenId *string `json:"UserOpenId,omitempty" name:"UserOpenId"`
		// 第三方平台unionId

		UserUnionId *string `json:"UserUnionId,omitempty" name:"UserUnionId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyUserAccessTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyUserAccessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiKey struct {

	// 密钥ID

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 创建时间(时间戳)

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 状态(2:有效, 3:禁用)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DisableApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddReceiverRequest struct {
	*tchttp.BaseRequest

	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 手机号

	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	// 国家编码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 电子邮件

	Email *string `json:"Email,omitempty" name:"Email"`
	// 是否是微信接收者

	IsReceiverWechat *int64 `json:"IsReceiverWechat,omitempty" name:"IsReceiverWechat"`
	// 身份认证类型

	IdentifyType *int64 `json:"IdentifyType,omitempty" name:"IdentifyType"`
	// 是否需要重置密码

	NeedResetPassword *int64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
	// 是否是主账户

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
}

func (r *AddReceiverRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddReceiverRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchOperateCamStrategyRequest struct {
	*tchttp.BaseRequest

	// 策略ID

	StrategyId []*uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 操作类型：1绑定，2解绑

	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`
	// 子账号列表

	RelateUin []*uint64 `json:"RelateUin,omitempty" name:"RelateUin"`
	// 语言：en，中文空

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *BatchOperateCamStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchOperateCamStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥所属用户Uin

	ApiUin *uint64 `json:"ApiUin,omitempty" name:"ApiUin"`
	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *DisableApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCollApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCollApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCollApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAttachedGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	TargetGroupId *uint64 `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
	// 页码，默认值是 1，从 1 开始

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页大小，默认值是 20

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListAttachedGroupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAttachedGroupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersForGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户数量

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 用户信息列表

		UserInfo []*GroupUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersForGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUsersForGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttachedEntityPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略分组列表

		List []*AttachedStrategyInfoPack `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttachedEntityPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttachedEntityPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCamServiceAndPermResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 业务信息列表

		List []*ServiceAndPerm `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCamServiceAndPermResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCamServiceAndPermResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEmailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEmailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IsOwner

		IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`
		// 是否是主账号

		IsOwner *bool `json:"IsOwner,omitempty" name:"IsOwner"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfo struct {

	// 组id

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 息接收渠道 0:无 1: 短信 2：邮件 3：短信+邮件

	Channel *int64 `json:"Channel,omitempty" name:"Channel"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 用户组成员信息

	UserInfo []*GroupMemberInfo `json:"UserInfo,omitempty" name:"UserInfo"`
	// 用户组类型，0-自定义，1-预设

	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`
}

type CreateOauthProviderRequest struct {
	*tchttp.BaseRequest

	// 身份提供商（企业）名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 注册应用的id

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// 注册应用的密钥

	ClientSecret *string `json:"ClientSecret,omitempty" name:"ClientSecret"`
	// oauth验证授权信息url

	AuthorizeUri *string `json:"AuthorizeUri,omitempty" name:"AuthorizeUri"`
	// 获取access_token url

	AccessTokenUri *string `json:"AccessTokenUri,omitempty" name:"AccessTokenUri"`
	// 获取用户信息url

	GetUserInfoUri *string `json:"GetUserInfoUri,omitempty" name:"GetUserInfoUri"`
	// 登录账号对应字段名称

	UserNameField *string `json:"UserNameField,omitempty" name:"UserNameField"`
	// 昵称对应字段名称

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// 手机号对应字段名称

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// 邮箱对应字段名称

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// 是否同步 idp 用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

func (r *CreateOauthProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOauthProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOauthProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOauthProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOauthProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountDetail struct {

	// 敏感操作标识

	ActionFlag *ActionLoginFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`
	// 是否允许控制台登录

	ConsoleLogin *string `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
	// 登录保护

	LoginFlag *ActionLoginFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`
	// 是否需要重置密码

	NeedResetPassword *string `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
	// 用户密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 使用Api

	UseApi *string `json:"UseApi,omitempty" name:"UseApi"`
	// 分配到设备类型

	TokenType *int64 `json:"TokenType,omitempty" name:"TokenType"`
}

type RolePolicyList struct {

	// 状态

	IsHidden *uint64 `json:"IsHidden,omitempty" name:"IsHidden"`
	// 策略Id

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

type CreateOauthProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// SAMLProviderArn

		SAMLProviderArn *string `json:"SAMLProviderArn,omitempty" name:"SAMLProviderArn"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOauthProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOauthProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRolePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRolePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserGroupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		GroupInfo []*UserGroup `json:"GroupInfo,omitempty" name:"GroupInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachRolePoliciesRequest struct {
	*tchttp.BaseRequest

	// 角色ID(与角色名称必传一项)

	RoleId *uint64 `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名称(与角色ID必传一项)

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 策略ID list(与策略名 list必传一项)

	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名 list(与策略ID list必传一项)

	PolicyName []*string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *AttachRolePoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachRolePoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupsSubAccountRequest struct {
	*tchttp.BaseRequest

	// 用户组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 单页数量

	Rp *string `json:"Rp,omitempty" name:"Rp"`
	// 分页数

	Page *string `json:"Page,omitempty" name:"Page"`
}

func (r *GetGroupsSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupsSubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesRequest struct {
	*tchttp.BaseRequest

	// 每页数量，默认值是 20，必须大于 0 且小于或等于 200

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 页码，默认值是 1，从 1开始，不能大于 200

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 可取值 'All'、'QCS' 和 'Local'，'All' 获取所有策略，'QCS' 只获取预设策略，'Local' 只获取自定义策略，默认取 'All'

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 按策略名匹配

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 按Uin匹配

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 按组Id匹配

	TargetGroupId *uint64 `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
	// 按角色Id匹配

	TargetRoleId *uint64 `json:"TargetRoleId,omitempty" name:"TargetRoleId"`
	// 按产品Id匹配，如cvm

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 调用来源，控制台为"console"

	Client *string `json:"Client,omitempty" name:"Client"`
	// 按Uin标记关联

	FlagUin *uint64 `json:"FlagUin,omitempty" name:"FlagUin"`
	// 按GroupId标记关联

	FlagGroupId *uint64 `json:"FlagGroupId,omitempty" name:"FlagGroupId"`
	// 按角色Id标记关联

	FlagRoleId *uint64 `json:"FlagRoleId,omitempty" name:"FlagRoleId"`
	// 创建类型：1.按产品功能或项目权限创建；2.按策略语法创建；3.按策略生成器创建；4.按标签授权创建；5.按权限边界规则创建

	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
}

func (r *ListPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedStrategyInfoPack struct {

	// 策略数组，数组每个成员包括 policyId、policyName、addTime、type、description、 createMode 字段。其中：

	List []*AttachedStrategyInfo `json:"List,omitempty" name:"List"`
	// 策略数

	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
	// 入参Type=1时表示uin，2时表示groupId

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type ModifyEmailRequest struct {
	*tchttp.BaseRequest

	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 验证码

	Code *string `json:"Code,omitempty" name:"Code"`
}

func (r *ModifyEmailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEmailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachUserPoliciesRequest struct {
	*tchttp.BaseRequest

	// 策略ID list

	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 用户ID

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *AttachUserPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachUserPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProviderList struct {

	// cas Id号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 创建者uin

	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`
	// 主账号

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 身份提供商名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 身份提供商描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 身份提供商类型

	ProviderType *uint64 `json:"ProviderType,omitempty" name:"ProviderType"`
	// 身份提供商状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Cas属性

	Cas *string `json:"Cas,omitempty" name:"Cas"`
}

type AccountAttributeAndValue struct {

	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 属性id

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 值id

	ValueId *int64 `json:"ValueId,omitempty" name:"ValueId"`
	// uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetPrincipalServiceRequest struct {
	*tchttp.BaseRequest

	// 类型（0：不隐藏；1：隐藏；缺省：所有）

	Shield *uint64 `json:"Shield,omitempty" name:"Shield"`
}

func (r *GetPrincipalServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPrincipalServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeInfo struct {

	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 属性值

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

type ServiceItem struct {

	// 创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// ArnDocument

	ArnDocument *string `json:"ArnDocument,omitempty" name:"ArnDocument"`
	// ColConf

	ColConf *string `json:"ColConf,omitempty" name:"ColConf"`
	// DefAddr

	DefAddr *string `json:"DefAddr,omitempty" name:"DefAddr"`
	// 默认策略

	DefaultStrategyList *string `json:"DefaultStrategyList,omitempty" name:"DefaultStrategyList"`
	// IsAllowDefProj

	IsAllowDefProj *string `json:"IsAllowDefProj,omitempty" name:"IsAllowDefProj"`
	// IsDisProject

	IsDisProject *string `json:"IsDisProject,omitempty" name:"IsDisProject"`
	// IsDisZone

	IsDisZone *string `json:"IsDisZone,omitempty" name:"IsDisZone"`
	// 是否可见

	IsSeen *string `json:"IsSeen,omitempty" name:"IsSeen"`
	// Online

	Online *string `json:"Online,omitempty" name:"Online"`
	// QueryAddr

	QueryAddr *string `json:"QueryAddr,omitempty" name:"QueryAddr"`
	// QueryInterface

	QueryInterface *string `json:"QueryInterface,omitempty" name:"QueryInterface"`
	// 服务英文名

	ServiceEnName *string `json:"ServiceEnName,omitempty" name:"ServiceEnName"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// SynInterface

	SynInterface *string `json:"SynInterface,omitempty" name:"SynInterface"`
	// 变更时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// Weight

	Weight *string `json:"Weight,omitempty" name:"Weight"`
	// WhiteKey

	WhiteKey *string `json:"WhiteKey,omitempty" name:"WhiteKey"`
	// 创建人

	Writter *string `json:"Writter,omitempty" name:"Writter"`
	// 资源类型数组

	ResourceTypeList *ResourceTypeItem `json:"ResourceTypeList,omitempty" name:"ResourceTypeList"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ConfirmCASProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmCASProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmCASProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersForPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 子账号列表

		List []*UserList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersForPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUsersForPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleDescriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRoleDescriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceApiInfo struct {

	// 服务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务ID

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 服务介绍文档链接

	ArnDocument *string `json:"ArnDocument,omitempty" name:"ArnDocument"`
	// API信息列表

	ApiList []*ServiceApiListInfo `json:"ApiList,omitempty" name:"ApiList"`
	// 条件规则列表

	ConditionKeyList []*string `json:"ConditionKeyList,omitempty" name:"ConditionKeyList"`
}

type CreateCollApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 持久密钥

		IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCollApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCollApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupListRequest struct {
	*tchttp.BaseRequest

	// 每页行数

	Rp *int64 `json:"Rp,omitempty" name:"Rp"`
	// 分页数

	Page *int64 `json:"Page,omitempty" name:"Page"`
}

func (r *GetGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserToGroupRequest struct {
	*tchttp.BaseRequest

	// 更新类型，1-增加，2-删除

	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`
	// 待更新用户列表

	Info []*GroupMember `json:"Info,omitempty" name:"Info"`
}

func (r *AddUserToGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserToGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCASProviderRequest struct {
	*tchttp.BaseRequest

	// CAS登录页面url

	CasLoginUrl *string `json:"CasLoginUrl,omitempty" name:"CasLoginUrl"`
	// CAS退出页面url

	CasLogoutUrl *string `json:"CasLogoutUrl,omitempty" name:"CasLogoutUrl"`
	// CasRoot

	CasRoot *string `json:"CasRoot,omitempty" name:"CasRoot"`
	// CAS校验ticket url

	CasValidateUrl *string `json:"CasValidateUrl,omitempty" name:"CasValidateUrl"`
	// 备注

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 身份提供商（企业）名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 自定义昵称字段

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// 自定义用户名字段

	UserNameField *string `json:"UserNameField,omitempty" name:"UserNameField"`
	// 自定义邮箱字段

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// 自定义手机字段

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// 是否同步 idp 用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

func (r *CreateCASProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCASProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllSubUserRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAllSubUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllSubUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOwnerAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOwnerAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOwnerAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CasProviderItem struct {

	// provider id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建账户uin

	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`
	// 主账户uin

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// provider类型

	ProviderType *int64 `json:"ProviderType,omitempty" name:"ProviderType"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// SAML元数据

	SAMLMetadata *string `json:"SAMLMetadata,omitempty" name:"SAMLMetadata"`
	// SAML实例id

	SAMLEntityId *string `json:"SAMLEntityId,omitempty" name:"SAMLEntityId"`
	// SAML登陆跳转

	SAMLSingleSignOn *string `json:"SAMLSingleSignOn,omitempty" name:"SAMLSingleSignOn"`
	// SAML登出跳转

	SAMLSingleLogout *string `json:"SAMLSingleLogout,omitempty" name:"SAMLSingleLogout"`
	// SAML关键字

	SAMLKeys *string `json:"SAMLKeys,omitempty" name:"SAMLKeys"`
	// Cas

	Cas *string `json:"Cas,omitempty" name:"Cas"`
	// cas根地址

	CasRoot *string `json:"CasRoot,omitempty" name:"CasRoot"`
	// cas登陆url

	CasLoginUrl *string `json:"CasLoginUrl,omitempty" name:"CasLoginUrl"`
	// cas校验url

	CasValidateUrl *string `json:"CasValidateUrl,omitempty" name:"CasValidateUrl"`
	// cas登出url

	CasLogoutUrl *string `json:"CasLogoutUrl,omitempty" name:"CasLogoutUrl"`
	// oauth配置

	Oauth *string `json:"Oauth,omitempty" name:"Oauth"`
}

type GroupMemberInfo struct {

	// 子用户 Uid。

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 子用户 Uin。

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 子用户名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 手机号。

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 手机区域代码。

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 是否已验证手机。

	PhoneFlag *uint64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// 邮箱地址。

	Email *string `json:"Email,omitempty" name:"Email"`
	// 是否已验证邮箱。

	EmailFlag *uint64 `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// 用户类型。

	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否为主消息接收人。

	IsReceiverOwner *uint64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// 企业微信id

	QywxUserId *string `json:"QywxUserId,omitempty" name:"QywxUserId"`
}

type ServicePermItem struct {

	// 创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// ApiAddr

	ApiAddr *string `json:"ApiAddr,omitempty" name:"ApiAddr"`
	// 中文描述

	ApiZhName *string `json:"ApiZhName,omitempty" name:"ApiZhName"`
	// 鉴权接口

	AuthFunction *string `json:"AuthFunction,omitempty" name:"AuthFunction"`
	// CWildcardName

	CWildcardName *string `json:"CWildcardName,omitempty" name:"CWildcardName"`
	// 接口名

	InterfaceEnName *string `json:"InterfaceEnName,omitempty" name:"InterfaceEnName"`
	// 鉴权粒度，0:接口级别、1:资源级别

	InterfaceLevel *string `json:"InterfaceLevel,omitempty" name:"InterfaceLevel"`
	// 鉴权方式，0:由云API转发鉴权、1:业务自行调用鉴权接口

	IsAuthBusiness *string `json:"IsAuthBusiness,omitempty" name:"IsAuthBusiness"`
	// IsNeedObject

	IsNeedObject *string `json:"IsNeedObject,omitempty" name:"IsNeedObject"`
	// IsSeen

	IsSeen *string `json:"IsSeen,omitempty" name:"IsSeen"`
	// 策略生成器是否可见

	IsSeenAtGenerator *string `json:"IsSeenAtGenerator,omitempty" name:"IsSeenAtGenerator"`
	// IsSpResource

	IsSpResource *string `json:"IsSpResource,omitempty" name:"IsSpResource"`
	// IsUserSet

	IsUserSet *string `json:"IsUserSet,omitempty" name:"IsUserSet"`
	// Id

	PermId *string `json:"PermId,omitempty" name:"PermId"`
	// 接口类别

	ReadWriteDetail *string `json:"ReadWriteDetail,omitempty" name:"ReadWriteDetail"`
	// 资源类别

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// Weight

	Weight *string `json:"Weight,omitempty" name:"Weight"`
	// 操作者

	Writter *string `json:"Writter,omitempty" name:"Writter"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// ProductShortCode

	ProductShortCode *string `json:"ProductShortCode,omitempty" name:"ProductShortCode"`
	// ProductShortName

	ProductShortName *string `json:"ProductShortName,omitempty" name:"ProductShortName"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type GroupMeta struct {

	// 用户组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 用户组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略ID

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo

	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachUsersPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略ID

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 用户ID list

	TargetUin []*uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *AttachUsersPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachUsersPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAttributeValuesRequest struct {
	*tchttp.BaseRequest

	// 属性列表

	Attributes []*AttributeInfo `json:"Attributes,omitempty" name:"Attributes"`
}

func (r *AddAttributeValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAttributeValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCollApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *DisableCollApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCollApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色详情列表。

		List []*RoleInfo `json:"List,omitempty" name:"List"`
		// 角色总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRoleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePasswordRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePasswordRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePasswordRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户唯一id

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 用户名

		Name *string `json:"Name,omitempty" name:"Name"`
		// 密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 密钥id

		SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
		// 密钥

		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubAccountFilter struct {

	// 子用户Uid

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 用户Uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是都允许登录

	CanLogin *uint64 `json:"CanLogin,omitempty" name:"CanLogin"`
	// 电话号码

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 区号

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 电话号码是否验证

	PhoneFlag *int64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 邮箱是否验证

	EmailFlag *int64 `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// 用户类型

	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否消息接收人

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// 类型

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
	// 是否需要重置密码

	NeedResetPassword *int64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
	// 是否允许控制台登录

	ConsoleLogin *int64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
	// 微信公众号关注状态

	WxzsStatus *int64 `json:"WxzsStatus,omitempty" name:"WxzsStatus"`
	// 权限类型

	PermType []*string `json:"PermType,omitempty" name:"PermType"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 企业微信用户id

	QywxUserId *string `json:"QywxUserId,omitempty" name:"QywxUserId"`
	// 扩展属性

	UserAttributeAndValues []*AccountAttributeAndValue `json:"UserAttributeAndValues,omitempty" name:"UserAttributeAndValues"`
}

type UpdateSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥所属用户Uin

	ApiUin *uint64 `json:"ApiUin,omitempty" name:"ApiUin"`
	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *DeleteApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshUserTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 第三方access_token

		UserAccessToken *string `json:"UserAccessToken,omitempty" name:"UserAccessToken"`
		// 过期时间

		ExpiresAt *int64 `json:"ExpiresAt,omitempty" name:"ExpiresAt"`
		// appId

		AppId *string `json:"AppId,omitempty" name:"AppId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefreshUserTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshUserTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiEnNameList struct {

	// API名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否需要关联对象

	IsNeedObject *string `json:"IsNeedObject,omitempty" name:"IsNeedObject"`
	// 是否可见

	IsSeen *string `json:"IsSeen,omitempty" name:"IsSeen"`
}

type ListOpenPlatformsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 单页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListOpenPlatformsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOpenPlatformsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPhoneNumRequest struct {
	*tchttp.BaseRequest

	// 手机号

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 手机号归属地编码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 验证码

	Code *string `json:"Code,omitempty" name:"Code"`
}

func (r *ModifyPhoneNumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPhoneNumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCASProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCASProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCASProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAccountLoginStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAccountLoginStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAccountLoginStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLoginRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLoginRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoginRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleSessionDurationRequest struct {
	*tchttp.BaseRequest

	// 角色名(与角色ID必填一个)

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 角色ID(与角色名必填一个)

	RoleId *uint64 `json:"RoleId,omitempty" name:"RoleId"`
	// 时长(秒)

	SessionDuration *uint64 `json:"SessionDuration,omitempty" name:"SessionDuration"`
	// 是否允许控制台登录

	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
}

func (r *UpdateRoleSessionDurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleSessionDurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServicePermListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *uint64 `json:"Code,omitempty" name:"Code"`
		// 响应信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 接口列表

		Data []*ServicePermItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServicePermListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServicePermListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubAccountContactsRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubAccountContactsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubAccountContactsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupRequest struct {
	*tchttp.BaseRequest

	// 待更新的组id，非0，-1表示创建一个新的用户组

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 用户组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 消息接收渠道 0:无 1: 短信 2：邮件 3：短信+邮件

	Channel *int64 `json:"Channel,omitempty" name:"Channel"`
}

func (r *UpdateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachRolesPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachRolesPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachRolesPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubAccountContactsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 联系方式列表。

		Contacts []*Receiver `json:"Contacts,omitempty" name:"Contacts"`
		// 总数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubAccountContactsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubAccountContactsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachRolesPolicyRequest struct {
	*tchttp.BaseRequest

	// 角色ID(与角色名称必填一个)

	RoleId []*uint64 `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名称(与角色ID必填一个)

	RoleName []*string `json:"RoleName,omitempty" name:"RoleName"`
	// 策略ID(与策略名必填一个)

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名 list(与策略ID必填一个)

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *DetachRolesPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachRolesPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAccessTokenRequest struct {
	*tchttp.BaseRequest

	// auth code授权码

	UserAuthCode *string `json:"UserAuthCode,omitempty" name:"UserAuthCode"`
	// 第三方access token，复杂授权使用。

	OpenAccessToken *string `json:"OpenAccessToken,omitempty" name:"OpenAccessToken"`
}

func (r *GetUserAccessTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAccessTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCollApiKeyRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryCollApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCollApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {

	// 子账号类型

	CanLogin *string `json:"CanLogin,omitempty" name:"CanLogin"`
	// 区号

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 详情

	Detail *AccountDetail `json:"Detail,omitempty" name:"Detail"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 电话号码

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 系统类型

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
	// 安全邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 微信登陆状态

	WxzsStatus *int64 `json:"WxzsStatus,omitempty" name:"WxzsStatus"`
	// 联系邮箱

	ContactMail *string `json:"ContactMail,omitempty" name:"ContactMail"`
	// 是否是主账号

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// 身份类型

	IdentifyType *int64 `json:"IdentifyType,omitempty" name:"IdentifyType"`
}

type DeleteRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 角色名

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 角色Id

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
}

func (r *DeleteRolePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRolePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendPhoneVerifyLinkRequest struct {
	*tchttp.BaseRequest

	// 目标账户uin，不填则发送到登录账户

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 区号

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *SendPhoneVerifyLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendPhoneVerifyLinkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserGroup struct {

	// 接收者用户id

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 账户唯一id

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否是主账户

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// 用户组信息

	Group []*GroupMeta `json:"Group,omitempty" name:"Group"`
}

type UpdateRoleSessionDurationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRoleSessionDurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleSessionDurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedPolicyOfRole struct {

	// 策略ID

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 绑定时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 策略类型，User表示自定义策略，QCS表示预设策略

	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
	// 策略创建方式，1表示按产品功能或项目权限创建，其他表示按策略语法创建

	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
}

type CheckAccountRequest struct {
	*tchttp.BaseRequest

	// 检查的uin

	OpUin *uint64 `json:"OpUin,omitempty" name:"OpUin"`
}

func (r *CheckAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOpenPlatform struct {

	// openid

	OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
	// app id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// app name

	OpenName *string `json:"OpenName,omitempty" name:"OpenName"`
	// open logo

	OpenLogo *string `json:"OpenLogo,omitempty" name:"OpenLogo"`
	// 第三方平台主页

	OpenHome *string `json:"OpenHome,omitempty" name:"OpenHome"`
	// 授权类型

	OpenType *int64 `json:"OpenType,omitempty" name:"OpenType"`
	// 申请账号

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 第三方平台域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 冻结状态，0-非冻结，1-冻结

	State *int64 `json:"State,omitempty" name:"State"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 备注

	Memo *string `json:"Memo,omitempty" name:"Memo"`
}

type GetLoginRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetLoginRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLoginRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCASProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCASProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCASProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListEntitiesForPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 页码，默认值是 1，从 1 开始

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页大小，默认值是 20

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 可取值 'All'、'User'、'Group' 和 'Role'，'All' 表示获取所有实体类型，'User' 表示只获取子账号，'Group' 表示只获取用户组，'Role' 表示只获取角色，默认取 'All'

	EntityFilter *string `json:"EntityFilter,omitempty" name:"EntityFilter"`
}

func (r *ListEntitiesForPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListEntitiesForPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户近

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// uin列表

		List []*uint64 `json:"List,omitempty" name:"List"`
		// uid列表

		Uids []*uint64 `json:"Uids,omitempty" name:"Uids"`
		// 账号详情

		SubAccounts []*AddSubAccountDetail `json:"SubAccounts,omitempty" name:"SubAccounts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAllUserGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		GroupInfo *UserGroup `json:"GroupInfo,omitempty" name:"GroupInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAllUserGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAllUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachUsersPolicyRequest struct {
	*tchttp.BaseRequest

	// 目标用户ID list

	TargetUin []*uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 策略ID

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DetachUsersPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachUsersPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCollApiKeyRequest struct {
	*tchttp.BaseRequest

	// 持久密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *EnableCollApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCollApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyUserPolicyRequest struct {
	*tchttp.BaseRequest

	// 来源子账号uin

	FromUin *uint64 `json:"FromUin,omitempty" name:"FromUin"`
	// 目标子账号uin

	ToUin []*uint64 `json:"ToUin,omitempty" name:"ToUin"`
}

func (r *CopyUserPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyUserPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest

	// 角色名称

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo

	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
	// 角色描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否允许登录 1 为允许 0 为不允许

	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
	// 申请角色临时密钥的最长有效期限制(范围：0~43200)

	SessionDuration *uint64 `json:"SessionDuration,omitempty" name:"SessionDuration"`
	// 角色类型(system|user)

	RoleType *string `json:"RoleType,omitempty" name:"RoleType"`
}

func (r *CreateRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAttachedUserAllPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略列表数据

		PolicyList []*AttachedUserPolicy `json:"PolicyList,omitempty" name:"PolicyList"`
		// 策略总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAttachedUserAllPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAttachedUserAllPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSubAccountsRequest struct {
	*tchttp.BaseRequest

	// 查询关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 子用户Uid

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 过滤类型

	IsFilter *int64 `json:"IsFilter,omitempty" name:"IsFilter"`
	// 过滤的用户组

	FilterGroups []*uint64 `json:"FilterGroups,omitempty" name:"FilterGroups"`
	// 过滤的策略ID

	FilterPolicyId *uint64 `json:"FilterPolicyId,omitempty" name:"FilterPolicyId"`
	// 项目ID

	ProjectId []*uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 用户类型

	UserType []*string `json:"UserType,omitempty" name:"UserType"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 起始条数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排查的用户

	ExcludeUids []*uint64 `json:"ExcludeUids,omitempty" name:"ExcludeUids"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
	// 检索关键字，支持name和uin

	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`
}

func (r *ListSubAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSubAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OwnerInfo struct {

	// 主帐号Uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 校验状态

	CheckStatus *uint64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
}

type ServiceApiListInfo struct {

	// API名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否需要关联对象

	IsNeedObject *string `json:"IsNeedObject,omitempty" name:"IsNeedObject"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 接口类别：0.读取，1.写入，2.标记，3.列表

	ReadWriteDetail *uint64 `json:"ReadWriteDetail,omitempty" name:"ReadWriteDetail"`
	// 授权粒度：0.接口级，1.资源级

	InterfaceLevel *uint64 `json:"InterfaceLevel,omitempty" name:"InterfaceLevel"`
	// 资源六段式范例

	ResourceExample *string `json:"ResourceExample,omitempty" name:"ResourceExample"`
}

type AddReceiverResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uid

		Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddReceiverResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Name struct {

	// 中文

	Zh *string `json:"Zh,omitempty" name:"Zh"`
	// 英文

	En *string `json:"En,omitempty" name:"En"`
}

type QueryCollApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥列表

		IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCollApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCollApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoleInfo struct {

	// 角色ID

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名称

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 角色的策略文档

	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
	// 角色描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 角色的创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 角色的最近一次时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 兼容公有云字段，无含义

	DeletionTaskId *string `json:"DeletionTaskId,omitempty" name:"DeletionTaskId"`
	// 角色是否允许登录

	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
	// 角色类型，取user、system或service_linked

	RoleType *string `json:"RoleType,omitempty" name:"RoleType"`
	// 有效时间

	SessionDuration *uint64 `json:"SessionDuration,omitempty" name:"SessionDuration"`
}

type CheckUserPolicyAttachmentRequest struct {
	*tchttp.BaseRequest

	// 子账号uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 策略id

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *CheckUserPolicyAttachmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckUserPolicyAttachmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLoginRulesRequest struct {
	*tchttp.BaseRequest

	// 登录态持续时间，单位分钟

	SessionDuration *int64 `json:"SessionDuration,omitempty" name:"SessionDuration"`
}

func (r *SetLoginRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoginRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupsSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户数量

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 用户信息列表

		UserInfo []*GroupUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupsSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupsSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 策略数组，数组每个成员包括 policyId、policyName、addTime、type、description、 createMode 字段。其中：
		// policyId：策略 id
		// policyName：策略名
		// addTime：策略创建时间
		// type：1 表示自定义策略，2 表示预设策略
		// description：策略描述
		// createMode：1 表示按业务权限创建的策略，其他值表示可以查看策略语法和通过策略语法更新策略
		// Attachments: 关联的用户数
		// ServiceType: 策略关联的产品
		// IsAttached: 当需要查询标记实体是否已经关联策略时不为null。0表示未关联策略，1表示已关联策略

		List []*StrategyInfo `json:"List,omitempty" name:"List"`
		// 服务列表

		ServiceTypeList []*string `json:"ServiceTypeList,omitempty" name:"ServiceTypeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachUserPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachUserPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachUserPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelModifyContactResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelModifyContactResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelModifyContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachGroupsPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略ID

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 用户组ID list

	GroupId []*uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *AttachGroupsPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachGroupsPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥所属用户Uin

	ApiUin *uint64 `json:"ApiUin,omitempty" name:"ApiUin"`
	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *EnableApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略Id

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *GetPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAccessTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// app id

		AppId *string `json:"AppId,omitempty" name:"AppId"`
		// 第三方openId

		UserOpenId *string `json:"UserOpenId,omitempty" name:"UserOpenId"`
		// 第三方unionId

		UserUnionId *string `json:"UserUnionId,omitempty" name:"UserUnionId"`
		// 第三方access token

		UserAccessToken *string `json:"UserAccessToken,omitempty" name:"UserAccessToken"`
		// 过期时间

		ExpiresAt *int64 `json:"ExpiresAt,omitempty" name:"ExpiresAt"`
		// refresh token

		UserRefreshToken *string `json:"UserRefreshToken,omitempty" name:"UserRefreshToken"`
		// 授权范围

		Scope *string `json:"Scope,omitempty" name:"Scope"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserAccessTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAccessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSubAccountRequest struct {
	*tchttp.BaseRequest

	// 账户变更类型，1-变更用户类型，2-变更用户权限，3-变更用户的项目权限，4-变更用户个人信息

	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`
	// 用户信息

	UserInfo []*SubAccountInfo `json:"UserInfo,omitempty" name:"UserInfo"`
	// 是否在白名单列表中， 是1

	InWhiteList *int64 `json:"InWhiteList,omitempty" name:"InWhiteList"`
	// 变更的用户类型

	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
	// sessId

	SessId *string `json:"SessId,omitempty" name:"SessId"`
	// vcode

	Vcode *string `json:"Vcode,omitempty" name:"Vcode"`
	// ptappid

	PtAppId *string `json:"PtAppId,omitempty" name:"PtAppId"`
	// 用户组信息

	GroupInfo []*UpdateGroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`
	// skey

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *UpdateSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsRequest struct {
	*tchttp.BaseRequest

	// 用户组列表，列表长度小于等于10

	GroupId []*uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 分页偏移量，缺省为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页条数，缺省为10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 用户组类型，0-自定义用户组，1-预设用户组

	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`
}

func (r *DescribeGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSubAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 用户信息

		UserInfo []*SubAccountFilter `json:"UserInfo,omitempty" name:"UserInfo"`
		// 主帐号信息

		OwnerInfo []*OwnerInfo `json:"OwnerInfo,omitempty" name:"OwnerInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSubAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSubAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubAccountRequest struct {
	*tchttp.BaseRequest

	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 用户信息

	UserInfo []*GroupUidUinInfo `json:"UserInfo,omitempty" name:"UserInfo"`
	// Api调用

	FromAPI *uint64 `json:"FromAPI,omitempty" name:"FromAPI"`
}

func (r *DeleteSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupListRequest struct {
	*tchttp.BaseRequest

	// 用户组id

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachGroupsPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachGroupsPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachGroupsPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeedOIDCLoginTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// token信息，提供到oidc

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SeedOIDCLoginTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeedOIDCLoginTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OwnerAccountAttribute struct {

	// 属性

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ListAllUserGroupRequest struct {
	*tchttp.BaseRequest

	// 登录帐号Uin

	LoginUin *string `json:"LoginUin,omitempty" name:"LoginUin"`
	// Rp

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// Page

	Page *uint64 `json:"Page,omitempty" name:"Page"`
}

func (r *ListAllUserGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAllUserGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleRequest struct {
	*tchttp.BaseRequest

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DeleteRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAssumeRolePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAssumeRolePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAssumeRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略名

		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
		// 策略描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 1 表示自定义策略，2 表示预设策略

		Type *uint64 `json:"Type,omitempty" name:"Type"`
		// 创建时间

		AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
		// 最近更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 策略文档

		PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
		// 备注

		PresetAlias *string `json:"PresetAlias,omitempty" name:"PresetAlias"`
		// 是否服务相关策略

		IsServiceLinkedRolePolicy *uint64 `json:"IsServiceLinkedRolePolicy,omitempty" name:"IsServiceLinkedRolePolicy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApiKeyRequest struct {
	*tchttp.BaseRequest

	// 用户Uin

	ApiUin *uint64 `json:"ApiUin,omitempty" name:"ApiUin"`
	// 自定义SecretId

	CustomSecretId *string `json:"CustomSecretId,omitempty" name:"CustomSecretId"`
	// 自定义SecretKey

	CustomSecretKey *string `json:"CustomSecretKey,omitempty" name:"CustomSecretKey"`
}

func (r *CreateApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendEmailVerifyLinkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendEmailVerifyLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendEmailVerifyLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceApiListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务以及API列表信息

		List []*ServiceApiInfo `json:"List,omitempty" name:"List"`
		// 策略条件

		Condition []*string `json:"Condition,omitempty" name:"Condition"`
		// 服务列表

		ServiceList []*ServiceApiInfo `json:"ServiceList,omitempty" name:"ServiceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceApiListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceApiListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOwnerAccountRequest struct {
	*tchttp.BaseRequest

	// 待更新待账户属性

	Detail *OwnerAccountAttribute `json:"Detail,omitempty" name:"Detail"`
}

func (r *UpdateOwnerAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOwnerAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupUidUinInfo struct {

	// 子用户Uid

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 子用户Uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户组ID 如果没有任何组传递-1,传入指定组id表示将用户从组删除

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

type AttachPolicyInfo struct {

	// 策略id

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 创建来源，1 通过控制台创建, 2 通过策略语法创建。

	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
}

type QueryApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API密钥数据列表

		IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAccountLoginStatusRequest struct {
	*tchttp.BaseRequest

	// 账户列表

	Accounts []*string `json:"Accounts,omitempty" name:"Accounts"`
	// 登陆状态

	LoginStatus *int64 `json:"LoginStatus,omitempty" name:"LoginStatus"`
}

func (r *SetAccountLoginStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAccountLoginStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCASProviderRequest struct {
	*tchttp.BaseRequest

	// CAS登录页面url

	CasLoginUrl *string `json:"CasLoginUrl,omitempty" name:"CasLoginUrl"`
	// CAS退出页面url

	CasLogoutUrl *string `json:"CasLogoutUrl,omitempty" name:"CasLogoutUrl"`
	// CasRoot

	CasRoot *string `json:"CasRoot,omitempty" name:"CasRoot"`
	// CAS校验ticket url

	CasValidateUrl *string `json:"CasValidateUrl,omitempty" name:"CasValidateUrl"`
	// 备注

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 身份提供商（企业）名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// OwnerUin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 自定义昵称字段

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// 自定义用户名字段

	UserNameField *string `json:"UserNameField,omitempty" name:"UserNameField"`
	// 自定义邮箱字段

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// 自定义手机字段

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// 是否同步 idp 用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

func (r *UpdateCASProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCASProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApiKeyRequest struct {
	*tchttp.BaseRequest

	// 用户Uin

	ApiUin *uint64 `json:"ApiUin,omitempty" name:"ApiUin"`
	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *GetApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupData struct {

	// 用户组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 用户组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 用户组成员数量

	GroupNum *int64 `json:"GroupNum,omitempty" name:"GroupNum"`
	// 创建渠道

	Channel *int64 `json:"Channel,omitempty" name:"Channel"`
	// 组成员uid

	GroupMem []*uint64 `json:"GroupMem,omitempty" name:"GroupMem"`
}

type GroupMember struct {

	// 用户id

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DetachRolePoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachRolePoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachRolePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PermList struct {

	// Id

	PermId *string `json:"PermId,omitempty" name:"PermId"`
	// API中文名

	ApiZhName *string `json:"ApiZhName,omitempty" name:"ApiZhName"`
	// API英文名信息列表

	ApiEnName []*ApiEnNameList `json:"ApiEnName,omitempty" name:"ApiEnName"`
	// 是否需要关联对象

	IsNeedObject *string `json:"IsNeedObject,omitempty" name:"IsNeedObject"`
	// 是否支持实例级鉴权

	IsSpResource *string `json:"IsSpResource,omitempty" name:"IsSpResource"`
	// pmService

	PmService *string `json:"PmService,omitempty" name:"PmService"`
}

type AttachRolePoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachRolePoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachRolePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubAccountInfo struct {

	// 能否登陆，0-否，1-可

	CanLogin *string `json:"CanLogin,omitempty" name:"CanLogin"`
	// 是否是控制台登陆，1-是

	ConsoleLogin *string `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
	// 国家编码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否需要重置密码，1-是

	NeedResetPassword *string `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
	// 手机号

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 账户类型

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
	// 接收者用户ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 账户唯一id

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 微信消息状态

	WxzsStatus *int64 `json:"WxzsStatus,omitempty" name:"WxzsStatus"`
	// 用户类型

	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
	// 联系邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
}

type GetPasswordRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密码规则列表

		Rules *PasswordRules `json:"Rules,omitempty" name:"Rules"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 更新用户

		Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
		// 黑名单字符串列表json string

		BlackList *string `json:"BlackList,omitempty" name:"BlackList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPasswordRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPasswordRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachRolesPolicyRequest struct {
	*tchttp.BaseRequest

	// 角色ID(与角色名称必传一项)

	RoleId []*uint64 `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名称(与角色ID必传一项)

	RoleName []*string `json:"RoleName,omitempty" name:"RoleName"`
	// 策略ID(与策略名必传一项)

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名(与策略ID必传一项)

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *AttachRolesPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachRolesPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupInfo struct {

	// 用户组id

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 用户id

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 用户组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 息接收渠道 0:无 1: 短信 2：邮件 3：短信+邮件

	Channel *int64 `json:"Channel,omitempty" name:"Channel"`
}

type AttachEntityOfPolicy struct {

	// 实体ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 实体名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实体Uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 关联类型。1 用户关联 ； 2 用户组关联

	RelatedType *uint64 `json:"RelatedType,omitempty" name:"RelatedType"`
}

type UserList struct {

	// 子账号名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 子账号uin

	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}

type CheckResourceRequest struct {
	*tchttp.BaseRequest

	// Arn资源描述符

	Arn *string `json:"Arn,omitempty" name:"Arn"`
}

func (r *CheckResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyUserAccessTokenRequest struct {
	*tchttp.BaseRequest

	// access token

	UserAccessToken *string `json:"UserAccessToken,omitempty" name:"UserAccessToken"`
	// open id

	UserOpenId *string `json:"UserOpenId,omitempty" name:"UserOpenId"`
}

func (r *VerifyUserAccessTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyUserAccessTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllSubUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组列表

		GroupInfo []*GroupData `json:"GroupInfo,omitempty" name:"GroupInfo"`
		// 用户信息列表

		UserInfo []*UserData `json:"UserInfo,omitempty" name:"UserInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAllSubUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllSubUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListEntitiesForPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实体总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 实体列表

		List []*AttachEntityOfPolicy `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEntitiesForPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListEntitiesForPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachUserPoliciesRequest struct {
	*tchttp.BaseRequest

	// 策略ID list

	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 目标用户ID

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *DetachUserPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachUserPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedUserPolicyGroupInfo struct {

	// 分组ID

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type Receiver struct {

	// id

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 手机号码

	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	// 手机号码是否验证

	PhoneFlag *int64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 邮箱是否验证

	EmailFlag *int64 `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// 是否主联系人

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// 是否允许微信接收通知

	WechatFlag *int64 `json:"WechatFlag,omitempty" name:"WechatFlag"`
	// 账号uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
}

type CheckResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否正确，正确：1，不正确：0

		Right *int64 `json:"Right,omitempty" name:"Right"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCollApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCollApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCollApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterItem struct {

	// 操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 属性

	Attr *string `json:"Attr,omitempty" name:"Attr"`
	// 匹配值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type AttachGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 策略ID list

	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 用户组ID

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *AttachGroupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachGroupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleListRequest struct {
	*tchttp.BaseRequest

	// 页码，从1开始

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页行数，不能大于200

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 按角色的服务账号载体过滤

	Service *string `json:"Service,omitempty" name:"Service"`
	// 按角色名或角色描述过滤

	Keyword []*string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeRoleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRoleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersForPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略id

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 页码

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页行数

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListUsersForPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUsersForPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtAttr struct {

	// 需要重置mfa的token

	NeedResetToken *int64 `json:"NeedResetToken,omitempty" name:"NeedResetToken"`
	// 需要重置mfa的stoken

	NeedResetStoken *int64 `json:"NeedResetStoken,omitempty" name:"NeedResetStoken"`
}

type StrategyInfo struct {

	// 策略ID。

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称。

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略创建时间。

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 策略类型。1 表示自定义策略，2 表示预设策略。

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 策略描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建来源，1 通过控制台创建, 2 通过策略语法创建。

	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
	// 关联的用户数

	Attachments *uint64 `json:"Attachments,omitempty" name:"Attachments"`
	// 策略关联的产品

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 当需要查询标记实体是否已经关联策略时不为null。0表示未关联策略，1表示已关联策略

	IsAttached *uint64 `json:"IsAttached,omitempty" name:"IsAttached"`
	// 是否已下线

	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`
	// 已下线产品列表

	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail"`
	// 是否进行安全性校验

	IsCheck *uint64 `json:"IsCheck,omitempty" name:"IsCheck"`
}

type DeletePolicyRequest struct {
	*tchttp.BaseRequest

	// 数组，数组成员是策略 id，支持批量删除策略

	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeletePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubsGroupRequest struct {
	*tchttp.BaseRequest

	// 接收者用户id

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 单页数量

	Rp *int64 `json:"Rp,omitempty" name:"Rp"`
	// 分页数

	Page *int64 `json:"Page,omitempty" name:"Page"`
}

func (r *GetSubsGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubsGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API密钥数据列表

		IdKeys []*ApiKeyDetail `json:"IdKeys,omitempty" name:"IdKeys"`
		// 用户uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 默认0

		SecretProjectId *uint64 `json:"SecretProjectId,omitempty" name:"SecretProjectId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersForGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 单页数量

	Rp *int64 `json:"Rp,omitempty" name:"Rp"`
	// 分页数

	Page *int64 `json:"Page,omitempty" name:"Page"`
}

func (r *ListUsersForGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUsersForGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略ID，入参PolicyId与PolicyName二选一

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 角色ID，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一

	AttachRoleId *string `json:"AttachRoleId,omitempty" name:"AttachRoleId"`
	// 角色名称，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一

	AttachRoleName *string `json:"AttachRoleName,omitempty" name:"AttachRoleName"`
	// 策略名，入参PolicyId与PolicyName二选一

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *AttachRolePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachRolePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// 待更新的组id，-1表示创建一个新的用户组

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 用户组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 消息接收渠道 0:无 1: 短信 2：邮件 3：短信+邮件

	Channel *int64 `json:"Channel,omitempty" name:"Channel"`
}

func (r *CreateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceRoleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色名

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// 业务类型

		ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
		// 英文业务类型

		ServiceTypeEn *string `json:"ServiceTypeEn,omitempty" name:"ServiceTypeEn"`
		// 角色描述

		RoleDesc *string `json:"RoleDesc,omitempty" name:"RoleDesc"`
		// 角色英文描述

		RoleDescEn *string `json:"RoleDescEn,omitempty" name:"RoleDescEn"`
		// 预设策略名

		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
		// 描述

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 英文描述

		EnRemark *string `json:"EnRemark,omitempty" name:"EnRemark"`
		// 策略列表

		PolicyList []*RolePolicyList `json:"PolicyList,omitempty" name:"PolicyList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceRoleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceRoleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelModifyContactRequest struct {
	*tchttp.BaseRequest

	// 目标账户uin，不传则操作登录账户

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 信息类型，1-手机号，2-邮箱

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *CancelModifyContactRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelModifyContactRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLoginRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 登陆h会话持续时间，单位分钟

		SessionDuration *int64 `json:"SessionDuration,omitempty" name:"SessionDuration"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLoginRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLoginRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCASProviderRequest struct {
	*tchttp.BaseRequest

	// 身份提供商（企业）名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// OwnerUin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *DisableCASProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCASProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPasswordRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetPasswordRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPasswordRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrinciPalService struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 名称

	Name *Name `json:"Name,omitempty" name:"Name"`
	// 是否存在服务相关角色

	ServiceLinkedRole *uint64 `json:"ServiceLinkedRole,omitempty" name:"ServiceLinkedRole"`
	// 角色场景信息

	RoleCases *PrincipalServiceRoleCases `json:"RoleCases,omitempty" name:"RoleCases"`
	// 业务模块名

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type CreatePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo

	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
	// 创建模式

	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
}

func (r *CreatePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleStatusRequest struct {
	*tchttp.BaseRequest

	// 角色ID

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
	// 角色状态，1：启用，2：禁用

	RoleStatus *int64 `json:"RoleStatus,omitempty" name:"RoleStatus"`
}

func (r *UpdateRoleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *uint64 `json:"Code,omitempty" name:"Code"`
		// 响应信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 服务信息

		Data []*ServiceItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCasProviderRequest struct {
	*tchttp.BaseRequest

	// Fields

	Fields *string `json:"Fields,omitempty" name:"Fields"`
	// 条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 第几页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCasProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCasProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCollApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCollApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCollApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserToGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserToGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserToGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceAndPerm struct {

	// 业务ID

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 业务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 是否分地域、分区

	IsDisZone *string `json:"IsDisZone,omitempty" name:"IsDisZone"`
	// 是否分项目

	IsDisProject *string `json:"IsDisProject,omitempty" name:"IsDisProject"`
	// 是否允许关联默认项目

	IsAllowDefProj *string `json:"IsAllowDefProj,omitempty" name:"IsAllowDefProj"`
	// 灰度白名单 key

	WhiteKey *string `json:"WhiteKey,omitempty" name:"WhiteKey"`
	// 相关API列表

	PermList []*PermList `json:"PermList,omitempty" name:"PermList"`
}

type AttachGroupPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachGroupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserRequest struct {
	*tchttp.BaseRequest

	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 手机号

	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	// 控制台登陆，是-1

	ConsoleLogin *int64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
	// 是否需要重置密码

	NeedResetPassword *int64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
	// 邮箱

	Mail *string `json:"Mail,omitempty" name:"Mail"`
	// 是否通过api请求，是-1

	UseApi *int64 `json:"UseApi,omitempty" name:"UseApi"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 手机号归属地国家编码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
}

func (r *AddUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupBatchRequest struct {
	*tchttp.BaseRequest

	// 用户组Id列表

	GroupId []*uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteGroupBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryKeyBySecretIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户Uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 秘钥列表

		IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryKeyBySecretIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryKeyBySecretIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApiKeyRequest struct {
	*tchttp.BaseRequest

	// 待查询的账号(不填默认查当前账号)

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *QueryApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGroupMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCASProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCASProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCASProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCasProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据总量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 配置信息列表

		List []*CasProviderItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCasProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCasProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略ID list

	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DetachGroupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachGroupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchOperateCamStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchOperateCamStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchOperateCamStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubReceiverResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uid

		Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
		// uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 主账号uin

		OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
		// 名字

		Name *string `json:"Name,omitempty" name:"Name"`
		// 手机号码

		PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
		// 手机号码是否验证 0：未验证  1：验证

		PhoneFlag *int64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
		// 邮箱

		Email *string `json:"Email,omitempty" name:"Email"`
		// 邮箱是否验证 0：未验证  1：验证

		EmailFlag *int64 `json:"EmailFlag,omitempty" name:"EmailFlag"`
		// 国家代码

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// 能否登录

		CanLogin *int64 `json:"CanLogin,omitempty" name:"CanLogin"`
		// 控制台登录

		ConsoleLogin *int64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
		// 账号类型

		SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
		// 微信通知状态

		WxzsStatus *int64 `json:"WxzsStatus,omitempty" name:"WxzsStatus"`
		// 是否需要重置密码

		NeedResetPassword *int64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
		// 用户属性

		ExtAttr *ExtAttr `json:"ExtAttr,omitempty" name:"ExtAttr"`
		// 用户昵称

		Nick *string `json:"Nick,omitempty" name:"Nick"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubReceiverResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePasswordRulesRequest struct {
	*tchttp.BaseRequest

	// 密码设置规则

	Rules *PasswordRules `json:"Rules,omitempty" name:"Rules"`
	// 黑名单字符串列表json string

	BlackList *string `json:"BlackList,omitempty" name:"BlackList"`
}

func (r *UpdatePasswordRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePasswordRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserData struct {

	// 子用户id

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 账号唯一序列号

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 电话号码

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 区号

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 电话认证标志

	PhoneFlag *int64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// 邮箱地址

	Email *string `json:"Email,omitempty" name:"Email"`
	// 邮箱是否认证

	EmailFlag *int64 `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// 用户类型

	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 微信标识

	WechatFlag *int64 `json:"WechatFlag,omitempty" name:"WechatFlag"`
	// 账号系统类型

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
	// 是否为主账号

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
}

type AttachRolePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachRolePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCamServiceAndPermRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCamServiceAndPermRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCamServiceAndPermRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOpenPlatformsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Cnt *int64 `json:"Cnt,omitempty" name:"Cnt"`
		// 开发平台列表

		List []*ListOpenPlatform `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOpenPlatformsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOpenPlatformsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleDescriptionRequest struct {
	*tchttp.BaseRequest

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 角色描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否控制台登录

	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
}

func (r *UpdateRoleDescriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleDescriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCollApiKeyRequest struct {
	*tchttp.BaseRequest

	// 持久密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *DeleteCollApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCollApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色详情

		RoleInfo *RoleInfo `json:"RoleInfo,omitempty" name:"RoleInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAttachedRolePoliciesRequest struct {
	*tchttp.BaseRequest

	// 角色 ID。用于指定角色，入参 RoleId 与 RoleName 二选一

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名。用于指定角色，入参 RoleId 与 RoleName 二选一

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 按策略类型过滤，User表示仅查询自定义策略，QCS表示仅查询预设策略

	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
	// 页码，从 1 开始

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页行数，不能大于200

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListAttachedRolePoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAttachedRolePoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOauthProviderRequest struct {
	*tchttp.BaseRequest

	// 备注

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 身份提供商（企业）名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// OwnerUin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 注册应用的id

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// 注册应用的密钥

	ClientSecret *string `json:"ClientSecret,omitempty" name:"ClientSecret"`
	// oauth验证授权信息url

	AuthorizeUri *string `json:"AuthorizeUri,omitempty" name:"AuthorizeUri"`
	// 获取access_token url

	AccessTokenUri *string `json:"AccessTokenUri,omitempty" name:"AccessTokenUri"`
	// 获取用户信息url

	GetUserInfoUri *string `json:"GetUserInfoUri,omitempty" name:"GetUserInfoUri"`
	// 登录账号对应字段名称

	UserNameField *string `json:"UserNameField,omitempty" name:"UserNameField"`
	// 昵称对应字段名称

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// 手机号对应字段名称

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// 邮箱对应字段名称

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// 是否同步 idp 用户数据

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

func (r *UpdateOauthProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOauthProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRoleRequest struct {
	*tchttp.BaseRequest

	// 角色 ID，用于指定角色，入参 RoleId 与 RoleName 二选一

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名，用于指定角色，入参 RoleId 与 RoleName 二选一

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *GetRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTypeItem struct {

	// 英文名

	ResourceEnName *string `json:"ResourceEnName,omitempty" name:"ResourceEnName"`
	// 中文名

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type AttachUsersPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachUsersPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachUsersPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveUserFromGroupRequest struct {
	*tchttp.BaseRequest

	// 传2

	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`
	// 待更新用户列表

	Info []*GroupMember `json:"Info,omitempty" name:"Info"`
}

func (r *RemoveUserFromGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveUserFromGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendPhoneVerifyLinkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendPhoneVerifyLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendPhoneVerifyLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedUserPolicy struct {

	// 策略ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 策略类型(1表示自定义策略，2表示预设策略)

	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`
	// 创建模式(1表示按产品或项目权限创建的策略，其他表示策略语法创建的策略)

	CreateMode *string `json:"CreateMode,omitempty" name:"CreateMode"`
	// 随组关联信息

	Groups []*AttachedUserPolicyGroupInfo `json:"Groups,omitempty" name:"Groups"`
}

type PrincipalServiceRoleCases struct {

	// 角色名

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 服务载体名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 场景中文名

	CaseName *string `json:"CaseName,omitempty" name:"CaseName"`
	// 场景英文名

	CaseEnName *string `json:"CaseEnName,omitempty" name:"CaseEnName"`
	// 角色描述

	RoleDesc *string `json:"RoleDesc,omitempty" name:"RoleDesc"`
	// 角色英文描述

	RoleEnDesc *string `json:"RoleEnDesc,omitempty" name:"RoleEnDesc"`
}

type RemoveUserFromGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveUserFromGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveUserFromGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceListRequest struct {
	*tchttp.BaseRequest

	// 服务名

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 排序字段 默认updateTime

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段desc/asc 默认desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *GetServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAttachedRolePoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色关联的策略列表

		List []*AttachedPolicyOfRole `json:"List,omitempty" name:"List"`
		// 角色关联的策略总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAttachedRolePoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAttachedRolePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一共有多少行数据

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 组信息列表

		GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshUserTokenRequest struct {
	*tchttp.BaseRequest

	// 用户刷新token

	UserRefreshToken *string `json:"UserRefreshToken,omitempty" name:"UserRefreshToken"`
	// 用户openId

	UserOpenId *string `json:"UserOpenId,omitempty" name:"UserOpenId"`
}

func (r *RefreshUserTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshUserTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeedOIDCLoginTokenRequest struct {
	*tchttp.BaseRequest

	// 用户uin

	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`
}

func (r *SeedOIDCLoginTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeedOIDCLoginTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCollApiKeyRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateCollApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCollApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachRolePoliciesRequest struct {
	*tchttp.BaseRequest

	// 角色ID(与角色名称必填一个)

	RoleId *uint64 `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名称(与角色ID必填一个)

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 策略ID list(与策略名 list必填一个)

	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名 list(与策略ID list必填一个)

	PolicyName []*string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *DetachRolePoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachRolePoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo

	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
}

func (r *CheckPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubsGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总体数量

		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
		// 用户信息列表

		GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubsGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubsGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAssumeRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo

	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
}

func (r *UpdateAssumeRolePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAssumeRolePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckUserPolicyAttachmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0表示未关联，1表示关联

		ExistAttachment *uint64 `json:"ExistAttachment,omitempty" name:"ExistAttachment"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckUserPolicyAttachmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckUserPolicyAttachmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组id，创建新组的时候返回

		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupUserInfo struct {

	// 接收者用户id

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

	IsReceiverOwner *string `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// 账户类型

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
}

type ActionLoginFlag struct {

	// 电话

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 软Token

	Stoken *string `json:"Stoken,omitempty" name:"Stoken"`
	// 硬Token

	Token *string `json:"Token,omitempty" name:"Token"`
	// ukey

	Ukey *string `json:"Ukey,omitempty" name:"Ukey"`
}

type Filter struct {

	// 过滤条件组合

	Keywords []*FilterItem `json:"Keywords,omitempty" name:"Keywords"`
	// 操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type CopyUserPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyUserPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略id

		PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRoleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAttachedGroupPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 策略列表

		List []*AttachPolicyInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAttachedGroupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAttachedGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesByActionRequest struct {
	*tchttp.BaseRequest

	// 每页数量，默认值是 20，必须大于 0 且小于或等于 200

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 页码，默认值是 1，从 1开始，不能大于 200

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 可取值 'All'、'QCS' 和 'Local'，'All' 获取所有策略，'QCS' 只获取预设策略，'Local' 只获取自定义策略，默认取 'All'

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 按接口名匹配

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 排序

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ListPoliciesByActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPoliciesByActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSubAccountDetail struct {

	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// token

	Token *string `json:"Token,omitempty" name:"Token"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// secretid

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// secretKey

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
}

type DescribeGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组信息

		GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`
		// 用户组数量

		GroupNum *int64 `json:"GroupNum,omitempty" name:"GroupNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachGroupPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachGroupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendEmailVerifyLinkRequest struct {
	*tchttp.BaseRequest

	// 目标账户uin，不填则发送登录账户

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *SendEmailVerifyLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendEmailVerifyLinkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色ID

		RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPrincipalServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号列表

		PrinciPalService []*PrinciPalService `json:"PrinciPalService,omitempty" name:"PrinciPalService"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPrincipalServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPrincipalServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachGroupsPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachGroupsPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachGroupsPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubReceiverRequest struct {
	*tchttp.BaseRequest

	// 是否是主接受者

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// 查询目标账号uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *DescribeSubReceiverRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubReceiverRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachUsersPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachUsersPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachUsersPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组id，创建新组的时候返回

		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupMemberRequest struct {
	*tchttp.BaseRequest

	// 更新类型，1-增加，2-删除

	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`
	// 待更新用户列表

	Info []*GroupMember `json:"Info,omitempty" name:"Info"`
}

func (r *UpdateGroupMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGroupMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPhoneNumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPhoneNumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPhoneNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmCASProviderRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// OwnerUin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *ConfirmCASProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmCASProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
