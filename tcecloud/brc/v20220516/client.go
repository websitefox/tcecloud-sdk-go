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

package v20220516

import (
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2022-05-16"

var _ = tchttp.POST

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewDescribeBackupGroupsRequest() (request *DescribeBackupGroupsRequest) {
	request = &DescribeBackupGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackupGroups")
	return
}

func NewDescribeBackupGroupsResponse() (response *DescribeBackupGroupsResponse) {
	response = &DescribeBackupGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询备份组列表，返回状态正常、创建中、回滚中的备份组列表。
func (c *Client) DescribeBackupGroups(request *DescribeBackupGroupsRequest) (response *DescribeBackupGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeBackupGroupsRequest()
	}
	response = NewDescribeBackupGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupResourceOverviewRequest() (request *DescribeBackupResourceOverviewRequest) {
	request = &DescribeBackupResourceOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackupResourceOverview")
	return
}

func NewDescribeBackupResourceOverviewResponse() (response *DescribeBackupResourceOverviewResponse) {
	response = &DescribeBackupResourceOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询备份资源概览数据。
func (c *Client) DescribeBackupResourceOverview(request *DescribeBackupResourceOverviewRequest) (response *DescribeBackupResourceOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBackupResourceOverviewRequest()
	}
	response = NewDescribeBackupResourceOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBackupsRequest() (request *DeleteBackupsRequest) {
	request = &DeleteBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DeleteBackups")
	return
}

func NewDeleteBackupsResponse() (response *DeleteBackupsResponse) {
	response = &DeleteBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteBackups）用于删除备份。
//
// * 备份必须处于NORMAL状态，快照状态可以通过DescribeBackups接口查询，见输出参数中BackupState字段解释；
// * 支持批量操作。如果多个备份存在无法删除的快照，则操作不执行，以返回特定的错误码返回。
func (c *Client) DeleteBackups(request *DeleteBackupsRequest) (response *DeleteBackupsResponse, err error) {
	if request == nil {
		request = NewDeleteBackupsRequest()
	}
	response = NewDeleteBackupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAutoBackupPoliciesRequest() (request *DeleteAutoBackupPoliciesRequest) {
	request = &DeleteAutoBackupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DeleteAutoBackupPolicies")
	return
}

func NewDeleteAutoBackupPoliciesResponse() (response *DeleteAutoBackupPoliciesResponse) {
	response = &DeleteAutoBackupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteAutoBackupPolicies）用于删除定期备份策略。
//
// *  支持批量操作。如果多个定期备份策略存在无法删除的，则操作不执行，以特定错误码返回。
func (c *Client) DeleteAutoBackupPolicies(request *DeleteAutoBackupPoliciesRequest) (response *DeleteAutoBackupPoliciesResponse, err error) {
	if request == nil {
		request = NewDeleteAutoBackupPoliciesRequest()
	}
	response = NewDeleteAutoBackupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAutoBackupPolicyRequest() (request *CreateAutoBackupPolicyRequest) {
	request = &CreateAutoBackupPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateAutoBackupPolicy")
	return
}

func NewCreateAutoBackupPolicyResponse() (response *CreateAutoBackupPolicyResponse) {
	response = &CreateAutoBackupPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateAutoBackupPolicy）用于创建定期备份策略。
//
func (c *Client) CreateAutoBackupPolicy(request *CreateAutoBackupPolicyRequest) (response *CreateAutoBackupPolicyResponse, err error) {
	if request == nil {
		request = NewCreateAutoBackupPolicyRequest()
	}
	response = NewCreateAutoBackupPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewApplyCfsBackupRequest() (request *ApplyCfsBackupRequest) {
	request = &ApplyCfsBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "ApplyCfsBackup")
	return
}

func NewApplyCfsBackupResponse() (response *ApplyCfsBackupResponse) {
	response = &ApplyCfsBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ApplyCfsBackup）用于回滚备份到原文件系统上。
//
// * 仅支持回滚到原文件系统上；
// * 用于回滚的备份必须处于NORMAL状态，备份状态可以通过DescribeBackups接口查询，见输出参数中BackupState字段解释；
func (c *Client) ApplyCfsBackup(request *ApplyCfsBackupRequest) (response *ApplyCfsBackupResponse, err error) {
	if request == nil {
		request = NewApplyCfsBackupRequest()
	}
	response = NewApplyCfsBackupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
	request = &CreateBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateBackup")
	return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
	response = &CreateBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateBackup）用于对指定云硬盘创建备份。
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
	if request == nil {
		request = NewCreateBackupRequest()
	}
	response = NewCreateBackupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupInstancesRequest() (request *DescribeBackupInstancesRequest) {
	request = &DescribeBackupInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackupInstances")
	return
}

func NewDescribeBackupInstancesResponse() (response *DescribeBackupInstancesResponse) {
	response = &DescribeBackupInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询受备份保护的实例列表。
// * 本接口用于查询实例的备份组信息，包含实例当前的备份列表，绑定的定期备份策略；
// * 本接口只会返回当前存在备份组的实例。
func (c *Client) DescribeBackupInstances(request *DescribeBackupInstancesRequest) (response *DescribeBackupInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeBackupInstancesRequest()
	}
	response = NewDescribeBackupInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewBindAutoBackupPolicyRequest() (request *BindAutoBackupPolicyRequest) {
	request = &BindAutoBackupPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "BindAutoBackupPolicy")
	return
}

func NewBindAutoBackupPolicyResponse() (response *BindAutoBackupPolicyResponse) {
	response = &BindAutoBackupPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（BindAutoBackupPolicy）用于绑定资源到指定的定期快照策略。
//
// * 当已绑定定期备份策略的云硬盘处于未使用状态（即弹性云盘未挂载或非弹性云盘的主机处于关机状态）将不会创建定期备份。
// * 云硬盘与云硬盘挂载的实例不能绑定到同一个ABP上。
func (c *Client) BindAutoBackupPolicy(request *BindAutoBackupPolicyRequest) (response *BindAutoBackupPolicyResponse, err error) {
	if request == nil {
		request = NewBindAutoBackupPolicyRequest()
	}
	response = NewBindAutoBackupPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfsBackupsRequest() (request *DescribeCfsBackupsRequest) {
	request = &DescribeCfsBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeCfsBackups")
	return
}

func NewDescribeCfsBackupsResponse() (response *DescribeCfsBackupsResponse) {
	response = &DescribeCfsBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询文件系统备份列表。
// * 可以根据备份ID、备份状态等信息来查询文件系统备份的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的文件系统备份列表。
func (c *Client) DescribeCfsBackups(request *DescribeCfsBackupsRequest) (response *DescribeCfsBackupsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsBackupsRequest()
	}
	response = NewDescribeCfsBackupsResponse()
	err = c.Send(request, response)
	return
}

func NewApplyBackupGroupRequest() (request *ApplyBackupGroupRequest) {
	request = &ApplyBackupGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "ApplyBackupGroup")
	return
}

func NewApplyBackupGroupResponse() (response *ApplyBackupGroupResponse) {
	response = &ApplyBackupGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于回滚备份组
// * 1.可选择备份组全部或部分盘进行回滚；
// * 2.如果回滚的盘中包含已挂载的盘，要求这些盘必须挂载在同一实例上，且要求该实例已关机才能回滚；
// * 3.回滚为异步操作，接口返回成功不代表回滚成功。
func (c *Client) ApplyBackupGroup(request *ApplyBackupGroupRequest) (response *ApplyBackupGroupResponse, err error) {
	if request == nil {
		request = NewApplyBackupGroupRequest()
	}
	response = NewApplyBackupGroupResponse()
	err = c.Send(request, response)
	return
}

func NewApplyBackupRequest() (request *ApplyBackupRequest) {
	request = &ApplyBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "ApplyBackup")
	return
}

func NewApplyBackupResponse() (response *ApplyBackupResponse) {
	response = &ApplyBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ApplyBackup）用于回滚备份到原云硬盘。
//
// * 仅支持回滚到原云硬盘上；
// * 用于回滚的备份必须处于NORMAL状态，备份状态可以通过DescribeBackups接口查询，见输出参数中BackupState字段解释；
func (c *Client) ApplyBackup(request *ApplyBackupRequest) (response *ApplyBackupResponse, err error) {
	if request == nil {
		request = NewApplyBackupRequest()
	}
	response = NewApplyBackupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupsRequest() (request *DescribeBackupsRequest) {
	request = &DescribeBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackups")
	return
}

func NewDescribeBackupsResponse() (response *DescribeBackupsResponse) {
	response = &DescribeBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询云硬盘备份列表。
// * 可以根据备份ID、备份状态等信息来查询云硬盘备份的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的云硬盘备份列表。
func (c *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
	if request == nil {
		request = NewDescribeBackupsRequest()
	}
	response = NewDescribeBackupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBackupAttributeRequest() (request *ModifyBackupAttributeRequest) {
	request = &ModifyBackupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "ModifyBackupAttribute")
	return
}

func NewModifyBackupAttributeResponse() (response *ModifyBackupAttributeResponse) {
	response = &ModifyBackupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyBackupAttribute）用于修改指定备份的属性。
//
// * 当前仅支持修改备份名称及将及备份的到期时间。
func (c *Client) ModifyBackupAttribute(request *ModifyBackupAttributeRequest) (response *ModifyBackupAttributeResponse, err error) {
	if request == nil {
		request = NewModifyBackupAttributeRequest()
	}
	response = NewModifyBackupAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindAutoBackupPolicyRequest() (request *UnbindAutoBackupPolicyRequest) {
	request = &UnbindAutoBackupPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "UnbindAutoBackupPolicy")
	return
}

func NewUnbindAutoBackupPolicyResponse() (response *UnbindAutoBackupPolicyResponse) {
	response = &UnbindAutoBackupPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnbindAutoBackupPolicy）用于解除资源绑定的定期备份策略。
//
// * 支持批量操作，可一次解除多个云硬盘与同一定期备份策略的绑定；
// * 如果传入的云硬盘未绑定到当前定期备份策略，接口将自动跳过，仅解绑与当前定期备份策略绑定的云硬盘。
func (c *Client) UnbindAutoBackupPolicy(request *UnbindAutoBackupPolicyRequest) (response *UnbindAutoBackupPolicyResponse, err error) {
	if request == nil {
		request = NewUnbindAutoBackupPolicyRequest()
	}
	response = NewUnbindAutoBackupPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCopyBackupToSnapshotRequest() (request *CopyBackupToSnapshotRequest) {
	request = &CopyBackupToSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CopyBackupToSnapshot")
	return
}

func NewCopyBackupToSnapshotResponse() (response *CopyBackupToSnapshotResponse) {
	response = &CopyBackupToSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CopyBackupToSnapshot）用于将备份转成快照。
// * 如果是系统盘备份转成的快照，则可用快照创建镜像后，再用镜像去新建实例；
// * 如果是数据盘备份，则转成快照后，可直接用于新建数据盘；
// * 支持将文件系统备份转成文件系统快照。
func (c *Client) CopyBackupToSnapshot(request *CopyBackupToSnapshotRequest) (response *CopyBackupToSnapshotResponse, err error) {
	if request == nil {
		request = NewCopyBackupToSnapshotRequest()
	}
	response = NewCopyBackupToSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDisksWithBackupRequest() (request *CreateDisksWithBackupRequest) {
	request = &CreateDisksWithBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateDisksWithBackup")
	return
}

func NewCreateDisksWithBackupResponse() (response *CreateDisksWithBackupResponse) {
	response = &CreateDisksWithBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateDisksWithBackup）用于创建从备份新建云硬盘。
//
// * 只有数据盘备份才支持新建云硬盘；
// * 本接口会将备份的数据回滚到新建云硬盘上；
// * 新建云硬盘的大小必须大于等于备份的大小；
// * 本接口为异步接口，接口成功返回时，说明回滚任务下发成功，但不代表数据已回滚完，需要通过DescribeDiks接口查询盘状态，字段Rollbacking为true，说明盘仍在回滚中。
func (c *Client) CreateDisksWithBackup(request *CreateDisksWithBackupRequest) (response *CreateDisksWithBackupResponse, err error) {
	if request == nil {
		request = NewCreateDisksWithBackupRequest()
	}
	response = NewCreateDisksWithBackupResponse()
	err = c.Send(request, response)
	return
}

func NewRunInstancesWithBackupGroupRequest() (request *RunInstancesWithBackupGroupRequest) {
	request = &RunInstancesWithBackupGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "RunInstancesWithBackupGroup")
	return
}

func NewRunInstancesWithBackupGroupResponse() (response *RunInstancesWithBackupGroupResponse) {
	response = &RunInstancesWithBackupGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于备份组新建云服器，将备份组的备份回滚到新建云服务器的云硬盘上；
// * 备份回滚是异步操作，云服务器创建成功后，并不代表备份回滚完成了，需要判断云服务器关联云硬盘的状态，看盘是否正在回滚中；
//
func (c *Client) RunInstancesWithBackupGroup(request *RunInstancesWithBackupGroupRequest) (response *RunInstancesWithBackupGroupResponse, err error) {
	if request == nil {
		request = NewRunInstancesWithBackupGroupRequest()
	}
	response = NewRunInstancesWithBackupGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCopyBackupCrossRegionsRequest() (request *CopyBackupCrossRegionsRequest) {
	request = &CopyBackupCrossRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CopyBackupCrossRegions")
	return
}

func NewCopyBackupCrossRegionsResponse() (response *CopyBackupCrossRegionsResponse) {
	response = &CopyBackupCrossRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CopyBackupCrossRegions）用于备份跨地域复制。
//
// * 本接口为异步接口，当跨地域复制的请求下发成功后会返回一个新的备份ID，此时备份未立即复制到目标地域，可请求目标地域的DescribeBackups接口获取新备份的状态，判断是否复制完成。如果备份的状态为“NORMAL”，表示备份复制完成。
func (c *Client) CopyBackupCrossRegions(request *CopyBackupCrossRegionsRequest) (response *CopyBackupCrossRegionsResponse, err error) {
	if request == nil {
		request = NewCopyBackupCrossRegionsRequest()
	}
	response = NewCopyBackupCrossRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupsDeniedActionsRequest() (request *DescribeBackupsDeniedActionsRequest) {
	request = &DescribeBackupsDeniedActionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackupsDeniedActions")
	return
}

func NewDescribeBackupsDeniedActionsResponse() (response *DescribeBackupsDeniedActionsResponse) {
	response = &DescribeBackupsDeniedActionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeBackupsDeniedActions）用于查询备份的操作掩码。
//
// * 根据入参BackupIds提供的备份ID列表，返回对应备份的操作掩码。
func (c *Client) DescribeBackupsDeniedActions(request *DescribeBackupsDeniedActionsRequest) (response *DescribeBackupsDeniedActionsResponse, err error) {
	if request == nil {
		request = NewDescribeBackupsDeniedActionsRequest()
	}
	response = NewDescribeBackupsDeniedActionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBackupGroupsRequest() (request *DeleteBackupGroupsRequest) {
	request = &DeleteBackupGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DeleteBackupGroups")
	return
}

func NewDeleteBackupGroupsResponse() (response *DeleteBackupGroupsResponse) {
	response = &DeleteBackupGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除备份组，一次调用仅支持删除一个备份组
// * 默认会删除备份组内的所有快照
//
func (c *Client) DeleteBackupGroups(request *DeleteBackupGroupsRequest) (response *DeleteBackupGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteBackupGroupsRequest()
	}
	response = NewDeleteBackupGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBackupGroupRequest() (request *CreateBackupGroupRequest) {
	request = &CreateBackupGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateBackupGroup")
	return
}

func NewCreateBackupGroupResponse() (response *CreateBackupGroupResponse) {
	response = &CreateBackupGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建备份组：
// * 创建备份组的盘列表必须挂载在同一实例上，可选择挂载在实例上的全部或部分盘创建备份组
func (c *Client) CreateBackupGroup(request *CreateBackupGroupRequest) (response *CreateBackupGroupResponse, err error) {
	if request == nil {
		request = NewCreateBackupGroupRequest()
	}
	response = NewCreateBackupGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoBackupPoliciesRequest() (request *DescribeAutoBackupPoliciesRequest) {
	request = &DescribeAutoBackupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeAutoBackupPolicies")
	return
}

func NewDescribeAutoBackupPoliciesResponse() (response *DescribeAutoBackupPoliciesResponse) {
	response = &DescribeAutoBackupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAutoBackupPolicies）用于查询定期备份策略列表。
// * 可以根据定期备份策略ID、名称或者状态等信息来查询定期备份策略的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的定期备份策略表。
func (c *Client) DescribeAutoBackupPolicies(request *DescribeAutoBackupPoliciesRequest) (response *DescribeAutoBackupPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeAutoBackupPoliciesRequest()
	}
	response = NewDescribeAutoBackupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupGroupsDeniedActionsRequest() (request *DescribeBackupGroupsDeniedActionsRequest) {
	request = &DescribeBackupGroupsDeniedActionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackupGroupsDeniedActions")
	return
}

func NewDescribeBackupGroupsDeniedActionsResponse() (response *DescribeBackupGroupsDeniedActionsResponse) {
	response = &DescribeBackupGroupsDeniedActionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeBackupGroupsDeniedActions）用于查询备份组的操作掩码。
//
// * 根据入参BackupGroupIds提供的备份组ID，返回对应备份组的操作掩码。
func (c *Client) DescribeBackupGroupsDeniedActions(request *DescribeBackupGroupsDeniedActionsRequest) (response *DescribeBackupGroupsDeniedActionsResponse, err error) {
	if request == nil {
		request = NewDescribeBackupGroupsDeniedActionsRequest()
	}
	response = NewDescribeBackupGroupsDeniedActionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupCfsFileSystemsRequest() (request *DescribeBackupCfsFileSystemsRequest) {
	request = &DescribeBackupCfsFileSystemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackupCfsFileSystems")
	return
}

func NewDescribeBackupCfsFileSystemsResponse() (response *DescribeBackupCfsFileSystemsResponse) {
	response = &DescribeBackupCfsFileSystemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询当前受备份保护的文件系统列表。
// * 本接口用于查询文件系统的备份信息，包含文件系统当前的备份列表，绑定的定期备份策略；
// * 本接口只会返回当前存在备份或绑定备份策略的文件系统。
func (c *Client) DescribeBackupCfsFileSystems(request *DescribeBackupCfsFileSystemsRequest) (response *DescribeBackupCfsFileSystemsResponse, err error) {
	if request == nil {
		request = NewDescribeBackupCfsFileSystemsRequest()
	}
	response = NewDescribeBackupCfsFileSystemsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupOperationsRequest() (request *DescribeBackupOperationsRequest) {
	request = &DescribeBackupOperationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackupOperations")
	return
}

func NewDescribeBackupOperationsResponse() (response *DescribeBackupOperationsResponse) {
	response = &DescribeBackupOperationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于操作与备份相关的操作日志
// * 所有与备份相关的修改操作均会记录操作日志；
func (c *Client) DescribeBackupOperations(request *DescribeBackupOperationsRequest) (response *DescribeBackupOperationsResponse, err error) {
	if request == nil {
		request = NewDescribeBackupOperationsRequest()
	}
	response = NewDescribeBackupOperationsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCfsFileSystemWithBackupRequest() (request *CreateCfsFileSystemWithBackupRequest) {
	request = &CreateCfsFileSystemWithBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateCfsFileSystemWithBackup")
	return
}

func NewCreateCfsFileSystemWithBackupResponse() (response *CreateCfsFileSystemWithBackupResponse) {
	response = &CreateCfsFileSystemWithBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCfsFileSystemWithBackup）用于通过备份新建文件系统。
// * 新建文件系统后，会把备份的数据回滚到文件系统上；
// * 调用此接口需要先通过授权创建角色BRC_CfsRole，才有权限新建文件系统。
func (c *Client) CreateCfsFileSystemWithBackup(request *CreateCfsFileSystemWithBackupRequest) (response *CreateCfsFileSystemWithBackupResponse, err error) {
	if request == nil {
		request = NewCreateCfsFileSystemWithBackupRequest()
	}
	response = NewCreateCfsFileSystemWithBackupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoBackupPolicyAttributeRequest() (request *ModifyAutoBackupPolicyAttributeRequest) {
	request = &ModifyAutoBackupPolicyAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "ModifyAutoBackupPolicyAttribute")
	return
}

func NewModifyAutoBackupPolicyAttributeResponse() (response *ModifyAutoBackupPolicyAttributeResponse) {
	response = &ModifyAutoBackupPolicyAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyAutoBackupPolicyAttribute）用于修改定期备份策略属性。
//
// * 可通过该接口修改定期备份策略的执行策略、名称、是否激活等属性；
// * 修改保留天数时必须保证不与是否永久保留属性冲突，否则整个操作失败，以特定的错误码返回。
func (c *Client) ModifyAutoBackupPolicyAttribute(request *ModifyAutoBackupPolicyAttributeRequest) (response *ModifyAutoBackupPolicyAttributeResponse, err error) {
	if request == nil {
		request = NewModifyAutoBackupPolicyAttributeRequest()
	}
	response = NewModifyAutoBackupPolicyAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewCopyBackupGroupToSnapshotGroupRequest() (request *CopyBackupGroupToSnapshotGroupRequest) {
	request = &CopyBackupGroupToSnapshotGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CopyBackupGroupToSnapshotGroup")
	return
}

func NewCopyBackupGroupToSnapshotGroupResponse() (response *CopyBackupGroupToSnapshotGroupResponse) {
	response = &CopyBackupGroupToSnapshotGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于将备份组转成快照组。
func (c *Client) CopyBackupGroupToSnapshotGroup(request *CopyBackupGroupToSnapshotGroupRequest) (response *CopyBackupGroupToSnapshotGroupResponse, err error) {
	if request == nil {
		request = NewCopyBackupGroupToSnapshotGroupRequest()
	}
	response = NewCopyBackupGroupToSnapshotGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCfsBackupRequest() (request *CreateCfsBackupRequest) {
	request = &CreateCfsBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateCfsBackup")
	return
}

func NewCreateCfsBackupResponse() (response *CreateCfsBackupResponse) {
	response = &CreateCfsBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCfsBackup）用于对指定文件系统创建备份。
func (c *Client) CreateCfsBackup(request *CreateCfsBackupRequest) (response *CreateCfsBackupResponse, err error) {
	if request == nil {
		request = NewCreateCfsBackupRequest()
	}
	response = NewCreateCfsBackupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupDisksRequest() (request *DescribeBackupDisksRequest) {
	request = &DescribeBackupDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackupDisks")
	return
}

func NewDescribeBackupDisksResponse() (response *DescribeBackupDisksResponse) {
	response = &DescribeBackupDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询受备份保护的云硬盘列表。
// * 本接口用于查询云硬盘的备份信息，包含实例当前的备份列表，绑定的定期备份策略；
// * 本接口只会返回当前存在备份的云硬盘。
func (c *Client) DescribeBackupDisks(request *DescribeBackupDisksRequest) (response *DescribeBackupDisksResponse, err error) {
	if request == nil {
		request = NewDescribeBackupDisksRequest()
	}
	response = NewDescribeBackupDisksResponse()
	err = c.Send(request, response)
	return
}
