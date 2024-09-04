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

package v20211001

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DeleteOrganizationRequest struct {
	*tchttp.BaseRequest

	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
}

func (r *DeleteOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uin列表

		Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationQuotasRequest struct {
	*tchttp.BaseRequest

	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 页数

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页显示条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤条件

	Filter *ProjectQuotaFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeOrganizationQuotasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationQuotasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrganizationMemberPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加失败的用户列表

		FailedUins *AddOrgMemberPolicyRes `json:"FailedUins,omitempty" name:"FailedUins"`
		// 添加成功的用户列表

		SuccessfulUins *AddOrgMemberPolicyRes `json:"SuccessfulUins,omitempty" name:"SuccessfulUins"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOrganizationMemberPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrganizationMemberPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationProjectsRequest struct {
	*tchttp.BaseRequest

	// 页数，例如：1

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页显示条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
}

func (r *DescribeOrganizationProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationProjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOrganizationMemberPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOrganizationMemberPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOrganizationMemberPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrgMemberPolicyRes struct {

	// 用户uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

type AddOrganizationMemberPolicyRequest struct {
	*tchttp.BaseRequest

	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 用户uin列表

	Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
	// 策略名称列表

	PolicyNames []*string `json:"PolicyNames,omitempty" name:"PolicyNames"`
}

func (r *AddOrganizationMemberPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrganizationMemberPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductInfo struct {

	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type ProjectQuota struct {

	// 项目id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品code

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

	QuotaLeft *uint64 `json:"QuotaLeft,omitempty" name:"QuotaLeft"`
	// 已使用配额

	QuotaUsed *uint64 `json:"QuotaUsed,omitempty" name:"QuotaUsed"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 子产品名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 产品项

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 产品细项

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
}

type Resource struct {

	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品大类名称

	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 项目id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域英文名

	RegionEnName *string `json:"RegionEnName,omitempty" name:"RegionEnName"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// ServiceType

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type Filter struct {

	// 过滤条件

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 组织id列表

	OrgIds []*string `json:"OrgIds,omitempty" name:"OrgIds"`
}

type OrganizationMember struct {

	// 用户uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 具有的组织策略

	OwnedPolicies []*OwnedPolicies `json:"OwnedPolicies,omitempty" name:"OwnedPolicies"`
	// 加入组织时间

	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`
}

type ProjectResourceFilter struct {

	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 组织id列表

	OrgIds []*string `json:"OrgIds,omitempty" name:"OrgIds"`
}

type DescribeOrganizationPoliciesTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略列表

		PolicySet *Policy `json:"PolicySet,omitempty" name:"PolicySet"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationPoliciesTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationPoliciesTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Policy struct {

	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type AddOrganizationRequest struct {
	*tchttp.BaseRequest

	// 父组织id，如果创建一级组织则填写“root”

	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`
	// 组织名称，长度最大64

	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
}

func (r *AddOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// uin列表

	Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
}

func (r *DeleteOrganizationMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNonMembersRequest struct {
	*tchttp.BaseRequest

	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 页数

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页显示条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeOrganizationNonMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNonMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNonMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 非组织成员

		MemberSet *OrganizationNonMember `json:"MemberSet,omitempty" name:"MemberSet"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNonMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNonMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOrganizationRequest struct {
	*tchttp.BaseRequest

	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 组织名称

	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
}

func (r *ModifyOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectQuotaFilter struct {

	// 一级产品

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 组织id列表

	OrgIds []*string `json:"OrgIds,omitempty" name:"OrgIds"`
}

type DescribeOrganizationProjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目列表

		ProjectSet []*Project `json:"ProjectSet,omitempty" name:"ProjectSet"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationProjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationResourcesRequest struct {
	*tchttp.BaseRequest

	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 页数

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤条件

	Filter *ProjectResourceFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeOrganizationResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OwnedPolicies struct {

	// 策略id

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DeleteOrganizationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织id

		OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOrganizationProjectsRequest struct {
	*tchttp.BaseRequest

	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 操作类型，只能是“Add”或“Move”

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 需要添加/移除的项目

	Projects []*string `json:"Projects,omitempty" name:"Projects"`
}

func (r *ModifyOrganizationProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOrganizationProjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationsFilter struct {

	// 级别，默认为3 表示搜索第一级、第二级和第三级符合搜索条件的组织

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 搜索条件

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
}

type DescribeOrganizationPoliciesTemplateRequest struct {
	*tchttp.BaseRequest

	// 页数

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页显示条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeOrganizationPoliciesTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationPoliciesTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源列表

		ResourceSet *Resource `json:"ResourceSet,omitempty" name:"ResourceSet"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOrganizationMemberPolicyRequest struct {
	*tchttp.BaseRequest

	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 账号uin

	AccountUin *uint64 `json:"AccountUin,omitempty" name:"AccountUin"`
	// 需要具有的策略列表

	PolicyNames []*string `json:"PolicyNames,omitempty" name:"PolicyNames"`
}

func (r *ModifyOrganizationMemberPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOrganizationMemberPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOrganizationProjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加/移除失败的项目列表

		FailedProjects []*string `json:"FailedProjects,omitempty" name:"FailedProjects"`
		// 添加/移除成功的项目列表

		SuccessfulProjects []*string `json:"SuccessfulProjects,omitempty" name:"SuccessfulProjects"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOrganizationProjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOrganizationProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织列表

		OrgSet []*Organization `json:"OrgSet,omitempty" name:"OrgSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Organization struct {

	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 组织名称

	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
	// 创建人uin

	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 子组织

	Children []*Organization `json:"Children,omitempty" name:"Children"`
}

type DescribeOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 页数

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页显示条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
}

func (r *DescribeOrganizationMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织成员

		MemberSet *OrganizationMember `json:"MemberSet,omitempty" name:"MemberSet"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
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

type DescribeOrganizationQuotasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额列表

		QuotaSet []*ProjectQuota `json:"QuotaSet,omitempty" name:"QuotaSet"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 产品信息列表

		ProductSet []*ProductInfo `json:"ProductSet,omitempty" name:"ProductSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationQuotasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filter *DescribeOrganizationsFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeOrganizationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrganizationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织id，例如：org-7737f12a

		OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrganizationNonMember struct {

	// 用户uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type Project struct {

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者uin

	CreatorUin *uint64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 组织id，不为空说明已经加入其它组织

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 组织名称

	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
	// 转入组织操作人

	OrgOperator *string `json:"OrgOperator,omitempty" name:"OrgOperator"`
	// 转入组织时间

	OrgOperationTime *string `json:"OrgOperationTime,omitempty" name:"OrgOperationTime"`
}

type ModifyOrganizationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织id

		OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
