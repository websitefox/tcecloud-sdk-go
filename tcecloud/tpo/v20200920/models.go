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

package v20200920

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeProductsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接入产品列表

		Products []*Products `json:"Products,omitempty" name:"Products"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectNonMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 成员列表

		MemberSet []*ProjectMember `json:"MemberSet,omitempty" name:"MemberSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectNonMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectNonMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectMemberFilter struct {

	// 按用户ID或用户名模糊查询

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表

		RegionSet []*Region `json:"RegionSet,omitempty" name:"RegionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {

	// 产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 服务类型，对应cam serivceType

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type AddProjectMemberPolicyRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 用户ID列表

	Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
	// 待添加的策略名

	PolicyNames []*string `json:"PolicyNames,omitempty" name:"PolicyNames"`
}

func (r *AddProjectMemberPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProjectMemberPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Products struct {

	// Product名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 资源类型

	ResourceTypes []*Resource `json:"ResourceTypes,omitempty" name:"ResourceTypes"`
}

type DescribeProjectNonMembersRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 页码

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页数量

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 搜索参数

	Filter *ProjectMemberFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeProjectNonMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectNonMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资源列表

		ResourceSet []*ProjectResource `json:"ResourceSet,omitempty" name:"ResourceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProjectQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProjectQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest

	// 资源列表

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 产品名

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DescribeResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveProjectResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MoveProjectResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveProjectResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一级产品列表

		ProductSet []*Product `json:"ProductSet,omitempty" name:"ProductSet"`
		// 二级产品列表

		SubProductSet []*Product `json:"SubProductSet,omitempty" name:"SubProductSet"`
		// 三级产品列表

		BillingItemSet []*Product `json:"BillingItemSet,omitempty" name:"BillingItemSet"`
		// 四级产品列表

		SubBillingItemSet []*Product `json:"SubBillingItemSet,omitempty" name:"SubBillingItemSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeResourceRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddProjectQuotaRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 一级产品定义

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 二级产品定义

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 三级产品定义

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 四级产品定义

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 一级产品配额

	ProductQuota *uint64 `json:"ProductQuota,omitempty" name:"ProductQuota"`
	// 二级产品配额

	SubProductQuota *uint64 `json:"SubProductQuota,omitempty" name:"SubProductQuota"`
	// 三级产品配额

	BillingItemQuota *uint64 `json:"BillingItemQuota,omitempty" name:"BillingItemQuota"`
	// 四级产品配额

	SubBillingItemQuota *uint64 `json:"SubBillingItemQuota,omitempty" name:"SubBillingItemQuota"`
}

func (r *AddProjectQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProjectQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckProjectQuotasRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 配额名

	QuotaSet []*QuotaNames `json:"QuotaSet,omitempty" name:"QuotaSet"`
}

func (r *CheckProjectQuotasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckProjectQuotasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectPoliciesRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 页码

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页数量

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 模糊查询

	Filter *PolicyFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeProjectPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectMembersRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 页码

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页数量

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 搜索条件

	Filter *ProjectMemberFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeProjectMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransferProjectResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransferProjectResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TransferProjectResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchAddProjectQuotaRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 待添加的配额

	QuotaSet []*AddQuota `json:"QuotaSet,omitempty" name:"QuotaSet"`
}

func (r *BatchAddProjectQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchAddProjectQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectMember struct {

	// 用户Uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户Uid

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 具有的项目策略

	Policies []*Policy `json:"Policies,omitempty" name:"Policies"`
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目ID

		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// 项目名

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 归属组织

	Organization *string `json:"Organization,omitempty" name:"Organization"`
	// 项目描述

	ProjectDescription *string `json:"ProjectDescription,omitempty" name:"ProjectDescription"`
}

func (r *CreateProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 项目列表

		ProjectSet []*Project `json:"ProjectSet,omitempty" name:"ProjectSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceTypesRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeResourceTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectMemberPolicyRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 用户UID

	AccountUin *uint64 `json:"AccountUin,omitempty" name:"AccountUin"`
	// 待添加的策略名列表

	PolicyNames []*string `json:"PolicyNames,omitempty" name:"PolicyNames"`
}

func (r *ModifyProjectMemberPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProjectMemberPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectMemberPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略名称列表

		PolicyNames []*string `json:"PolicyNames,omitempty" name:"PolicyNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProjectMemberPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProjectMemberPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveProjectResourceRequest struct {
	*tchttp.BaseRequest

	// 原项目

	OldProjectId *string `json:"OldProjectId,omitempty" name:"OldProjectId"`
	// 新项目

	NewProjectId *string `json:"NewProjectId,omitempty" name:"NewProjectId"`
	// 资源

	ResourceList []*TransferResource `json:"ResourceList,omitempty" name:"ResourceList"`
}

func (r *MoveProjectResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveProjectResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectQuotasRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 页码

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页数量

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 搜索条件

	Filter *ProjectQuotaFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeProjectQuotasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectQuotasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectQuotaFilter struct {

	// 一级产品

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 二级产品

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 三级产品

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 四级产品

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 配额项名称模糊查询

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type QuotaChecking struct {

	// 三级产品

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 产品项

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 错误内容

	Error *string `json:"Error,omitempty" name:"Error"`
	// 是否已有配额

	Exists *bool `json:"Exists,omitempty" name:"Exists"`
	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 配额键

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// 剩余配额

	QuotaLeft *string `json:"QuotaLeft,omitempty" name:"QuotaLeft"`
	// 已使用配额

	QuotaUsed *string `json:"QuotaUsed,omitempty" name:"QuotaUsed"`
	// 配额值

	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`
	// 四级产品

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 产品细项

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 二级产品

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 是否合法

	Success *bool `json:"Success,omitempty" name:"Success"`
}

type RemoveProjectMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 移除的用户uin列表

		Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveProjectMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveProjectMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Product struct {

	// 产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type ModifyProjectNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目id

		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProjectNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProjectNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Project struct {

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 组织

	Organization *string `json:"Organization,omitempty" name:"Organization"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者Uin

	CreatorUin *uint64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 组织名称

	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
	// 项目转入项目目录时间

	OrgOperationTime *string `json:"OrgOperationTime,omitempty" name:"OrgOperationTime"`
	// 项目转入项目目录操作人

	OrgOperator *string `json:"OrgOperator,omitempty" name:"OrgOperator"`
	// 项目描述

	ProjectDescription *string `json:"ProjectDescription,omitempty" name:"ProjectDescription"`
}

type AddProjectMemberPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功列表

		SuccessfulUins []*UinPolicy `json:"SuccessfulUins,omitempty" name:"SuccessfulUins"`
		// 失败列表

		FailedUins []*UinPolicy `json:"FailedUins,omitempty" name:"FailedUins"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddProjectMemberPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProjectMemberPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveProjectMemberRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 待移除的用户UID列表

	Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
}

func (r *RemoveProjectMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveProjectMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceAdminProject struct {

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 项目名

	Organization *string `json:"Organization,omitempty" name:"Organization"`
}

type AddProjectResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddProjectResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProjectResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchAddProjectQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加列表

		AddList *AddQuota `json:"AddList,omitempty" name:"AddList"`
		// 是否添加成功

		AddSuccess *bool `json:"AddSuccess,omitempty" name:"AddSuccess"`
		// 更新列表

		UpdateList *AddQuota `json:"UpdateList,omitempty" name:"UpdateList"`
		// 是否更新成功

		UpdateSuccess *bool `json:"UpdateSuccess,omitempty" name:"UpdateSuccess"`
		// 错误列表

		ErrorList *AddQuota `json:"ErrorList,omitempty" name:"ErrorList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchAddProjectQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchAddProjectQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckProjectQuotasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额列表

		QuotaSet []*QuotaChecking `json:"QuotaSet,omitempty" name:"QuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckProjectQuotasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckProjectQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectResourceRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 资源列表

	ResourceList []*TransferResource `json:"ResourceList,omitempty" name:"ResourceList"`
}

func (r *DeleteProjectResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 策略列表

		PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceAdminProjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 项目列表

		ProjectSet []*ResourceAdminProject `json:"ProjectSet,omitempty" name:"ProjectSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceAdminProjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceAdminProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectQuotaRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 一级产品定义

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 配额键值

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// 配额数

	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`
}

func (r *ModifyProjectQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProjectQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectNameExistsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否存在

		Exist *bool `json:"Exist,omitempty" name:"Exist"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProjectNameExistsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProjectNameExistsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectResource struct {

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品大类名称

	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`
	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 地域英文名

	RegionEnName *string `json:"RegionEnName,omitempty" name:"RegionEnName"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 产品对应的serviceType

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type DescribeResourceTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源类型列表

		ResourceTypeSet []*string `json:"ResourceTypeSet,omitempty" name:"ResourceTypeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectResourcesRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 页码

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页数量

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 搜索条件

	Filter *ProjectResourceFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeProjectResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductTreeRequest struct {
	*tchttp.BaseRequest

	// 一级产品

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 二级产品

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 三级产品

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 四级产品

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 资源获取方式，billing表示从计费同步资源ID，local表示直接从垂直产品同步资源ID

	AccessTypes []*string `json:"AccessTypes,omitempty" name:"AccessTypes"`
}

func (r *DescribeProductTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectNameExistsRequest struct {
	*tchttp.BaseRequest

	// 项目名

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *ProjectNameExistsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProjectNameExistsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UinPolicy struct {

	// uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

type AddQuota struct {

	// 一级产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 二级产品编码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 三级产品编码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 四级产品编码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 配额值

	QuotaValue *int64 `json:"QuotaValue,omitempty" name:"QuotaValue"`
	// 配额键

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
}

type TransferResource struct {

	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 地域ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 用户Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type ProjectQuota struct {

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 二级产品

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 三级产品

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 四级产品

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 配额键

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// 配额项名称

	QuotaName *string `json:"QuotaName,omitempty" name:"QuotaName"`
	// 总配额

	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`
	// 剩余配额

	QuotaLeft *string `json:"QuotaLeft,omitempty" name:"QuotaLeft"`
	// 已使用配额

	QuotaUsed *string `json:"QuotaUsed,omitempty" name:"QuotaUsed"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 产品项

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 产品细项

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 子产品名

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type ProjectResourceFilter struct {

	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 资源类型（暂不支持）

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源ID或资源名称模糊查询

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 产品Code

	Product *string `json:"Product,omitempty" name:"Product"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type AddProjectResourceRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 添加的资源列表

	ResourceList []*TransferResource `json:"ResourceList,omitempty" name:"ResourceList"`
}

func (r *AddProjectResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProjectResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectQuotaRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 资源列表

	ResourceList []*DeleteQuota `json:"ResourceList,omitempty" name:"ResourceList"`
}

func (r *DeleteProjectQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 成员列表

		MemberSet []*ProjectMember `json:"MemberSet,omitempty" name:"MemberSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectQuotasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配额列表

		QuotaSet []*ProjectQuota `json:"QuotaSet,omitempty" name:"QuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectQuotasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransferProjectResourceRequest struct {
	*tchttp.BaseRequest

	// 原项目ID

	OldProjectId *string `json:"OldProjectId,omitempty" name:"OldProjectId"`
	// 新项目ID

	NewProjectId *string `json:"NewProjectId,omitempty" name:"NewProjectId"`
	// 变动的资源列表

	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`
}

func (r *TransferProjectResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TransferProjectResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsFilter struct {

	// 按项目ID或者项目名模糊搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type Region struct {

	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProductsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectNameRequest struct {
	*tchttp.BaseRequest

	// 项目id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 项目描述

	ProjectDescription *string `json:"ProjectDescription,omitempty" name:"ProjectDescription"`
}

func (r *ModifyProjectNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProjectNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目ID

		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectMemberPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已添加策略

		OwnedPolicies []*Policy `json:"OwnedPolicies,omitempty" name:"OwnedPolicies"`
		// 未添加策略

		Policies []*Policy `json:"Policies,omitempty" name:"Policies"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectMemberPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectMemberPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest

	// 页码

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页数量

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤条件

	Filter *DescribeProjectsFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuotaNames struct {

	// 一级产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 二级产品名

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 三级产品名

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 四级产品名

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
}

type AddProjectQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddProjectQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProjectQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectMemberPoliciesRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 用户Uin

	AccountUin *uint64 `json:"AccountUin,omitempty" name:"AccountUin"`
	// 过滤参数

	Filter *PolicyFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeProjectMemberPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectMemberPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceAdminProjectsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeResourceAdminProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceAdminProjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源列表

		ResourceSet []*ProjectResource `json:"ResourceSet,omitempty" name:"ResourceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQuota struct {

	// 产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 配额键

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
}

type Policy struct {

	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 策略ID

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type PolicyFilter struct {

	// 按策略名模糊查询

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}
