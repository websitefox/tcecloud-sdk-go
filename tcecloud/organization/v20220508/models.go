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

package v20220508

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type MoveOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest

	// 新的节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 成员账号uin列表

	MemberUin []*uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *MoveOrganizationNodeMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveOrganizationNodeMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMemberPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略id

		PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrganizationMemberPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrganizationMemberPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMemberPolicyRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 策略名称， tce后端自动生成

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 身份id， 固定传1

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateOrganizationMemberPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrganizationMemberPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodesByParentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织节点列表

		Nodes []*OrgNode `json:"Nodes,omitempty" name:"Nodes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodesByParentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodesByParentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrganizationNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的节点id

		NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOrganizationNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodesRequest struct {
	*tchttp.BaseRequest

	// 节点id列表

	NodeId []*int64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DeleteOrganizationNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeRequest struct {
	*tchttp.BaseRequest

	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DescribeOrganizationNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全部数据量大小

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 节点列表

		Items []*OrgNode `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationNodeMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest

	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 成员账号列表

	MemberUin []*uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *DeleteOrganizationNodeMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationNodeMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindOrganizationMemberAuthAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationNodeRequest struct {
	*tchttp.BaseRequest

	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *UpdateOrganizationNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgMember struct {

	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 成员类型

	MemberType *string `json:"MemberType,omitempty" name:"MemberType"`
	// 集团策略类型

	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	// 集团策略名称

	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
	// 集团权限

	OrgPermission []*Permission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	// 代付uin

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 代付用户名称

	PayName *string `json:"PayName,omitempty" name:"PayName"`
	// 节点id

	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否允许退出

	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
	// 权限状态

	PermissionStatus *string `json:"PermissionStatus,omitempty" name:"PermissionStatus"`
	// 授权身份列表

	OrgIdentity []*Identity `json:"OrgIdentity,omitempty" name:"OrgIdentity"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否是管理员

	IsManager *int64 `json:"IsManager,omitempty" name:"IsManager"`
}

type DescribeOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小[0,50]

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索关键字，支持账号名称和uin

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest struct {
	*tchttp.BaseRequest

	// 子账号uin

	OrgSubAccountUin *uint64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
}

func (r *CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成员账号uin

		MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
		// 名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 成员类型

		MemberType *string `json:"MemberType,omitempty" name:"MemberType"`
		// 集团策略类型

		OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
		// 集团策略名称

		OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
		// 集团权限

		OrgPermission []*Permission `json:"OrgPermission,omitempty" name:"OrgPermission"`
		// 代付uin

		PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
		// 代付用户名称

		PayName *string `json:"PayName,omitempty" name:"PayName"`
		// 节点id

		NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
		// 节点名称

		NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 是否允许退出

		IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
		// 授权身份列表

		OrgIdentity []*Identity `json:"OrgIdentity,omitempty" name:"OrgIdentity"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveOrganizationNodeRequest struct {
	*tchttp.BaseRequest

	// 新的父节点id

	ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
	// 被移动的节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *MoveOrganizationNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveOrganizationNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrganizationNodeRequest struct {
	*tchttp.BaseRequest

	// 父节点id

	ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *AddOrganizationNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrganizationNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 策略id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 子账号uin

	OrgSubAccountUin *uint64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
}

func (r *CancelOrganizationMemberAuthAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Identity struct {

	// id

	IdentityId *uint64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 别名

	IdentityAliasName *string `json:"IdentityAliasName,omitempty" name:"IdentityAliasName"`
}

type DescribeOrganizationNodesByParentRequest struct {
	*tchttp.BaseRequest

	// 父节点id

	ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
}

func (r *DescribeOrganizationNodesByParentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodesByParentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 策略id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 授权子账号uin列表。 <=5

	OrgSubAccountUins []*uint64 `json:"OrgSubAccountUins,omitempty" name:"OrgSubAccountUins"`
}

func (r *BindOrganizationMemberAuthAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMemberAuthAccountsRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin， 传0则不使用成员账号进行过滤

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 策略id，传0则不指定策略

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小，[0,50]

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 子账号uin， 传0则不根据子账号进行过滤

	TargetSubAccountUin *int64 `json:"TargetSubAccountUin,omitempty" name:"TargetSubAccountUin"`
}

func (r *DescribeOrganizationMemberAuthAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberAuthAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据量大小

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 成员账号列表

		Items []*OrgMember `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveOrganizationNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MoveOrganizationNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MoveOrganizationNodeMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckMemberExistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存在的成员账号uin列表

		ExistMemberUin []*uint64 `json:"ExistMemberUin,omitempty" name:"ExistMemberUin"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckMemberExistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMemberExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询数量[0-50]

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOrganizationNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeMember struct {

	// 成员账号uin

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 成员类型，Create,Invite,Admin等

	MemberType *string `json:"MemberType,omitempty" name:"MemberType"`
	// 组织策略类型，Finance

	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	// 权限列表

	OrgPermission []*Permission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	// 代付账户

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 代付用户名

	PayName *string `json:"PayName,omitempty" name:"PayName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CancelOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelOrganizationMemberAuthAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest

	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 查询偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索关键字

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *DescribeOrganizationNodeMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgNode struct {

	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 父节点id

	ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type Permission struct {

	// id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeOrganizationRequest struct {
	*tchttp.BaseRequest

	// 语言，支持cn,en

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *DescribeOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *DescribeOrganizationMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点id

		NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
		// 父节点id

		ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
		// 名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMemberAuthAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全部数据量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 授权子账号列表

		Items []*OrgMemberAuthAccount `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberAuthAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 成员列表

		Items []*NodeMember `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodeMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrganizationNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgMemberAuthAccount struct {

	// 子账号uin

	OrgSubAccountUin *uint64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
	// 策略id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 身份id

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 身份角色名称

	IdentityRoleName *string `json:"IdentityRoleName,omitempty" name:"IdentityRoleName"`
	// 身份角色别名

	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitempty" name:"IdentityRoleAliasName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 被授权的成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

type CheckMemberExistRequest struct {
	*tchttp.BaseRequest

	// 待检测的账号uin列表

	MemberUin []*uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *CheckMemberExistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMemberExistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集团组织id

		OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`
		// 组织类型

		OrgType *int64 `json:"OrgType,omitempty" name:"OrgType"`
		// 管理员账号uin

		HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
		// 管理员账号昵称

		NickName *string `json:"NickName,omitempty" name:"NickName"`
		// 请求账号是否为管理员

		IsManager *bool `json:"IsManager,omitempty" name:"IsManager"`
		// 请求账号是否允许退出组织

		IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
		// 根几点id

		RootNodeId *int64 `json:"RootNodeId,omitempty" name:"RootNodeId"`
		// 组织关系策略类型，如Financial

		OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
		// 关系策略名称，如财务管理

		OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
		// 关联策略权限

		OrgPermission []*Permission `json:"OrgPermission,omitempty" name:"OrgPermission"`
		// 代付账号uin

		PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
		// 代付账号用户名

		PayName *string `json:"PayName,omitempty" name:"PayName"`
		// 组织创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 请求账号加入集团组织时间

		JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`
		// 管理员用户名

		ManagerName *string `json:"ManagerName,omitempty" name:"ManagerName"`
		// 集团组织名称

		OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
