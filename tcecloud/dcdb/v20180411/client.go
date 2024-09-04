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

package v20180411

import (
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-04-11"

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

func NewDescribeDBLogFilesRequest() (request *DescribeDBLogFilesRequest) {
	request = &DescribeDBLogFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBLogFiles")
	return
}

func NewDescribeDBLogFilesResponse() (response *DescribeDBLogFilesResponse) {
	response = &DescribeDBLogFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDBLogFiles)用于获取数据库的各种日志列表，包括冷备、binlog、errlog和slowlog。
func (c *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) (response *DescribeDBLogFilesResponse, err error) {
	if request == nil {
		request = NewDescribeDBLogFilesRequest()
	}
	response = NewDescribeDBLogFilesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRealServerAccessStrategyRequest() (request *ModifyRealServerAccessStrategyRequest) {
	request = &ModifyRealServerAccessStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyRealServerAccessStrategy")
	return
}

func NewModifyRealServerAccessStrategyResponse() (response *ModifyRealServerAccessStrategyResponse) {
	response = &ModifyRealServerAccessStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyRealServerAccessStrategy）用于修改就近接入策略
func (c *Client) ModifyRealServerAccessStrategy(request *ModifyRealServerAccessStrategyRequest) (response *ModifyRealServerAccessStrategyResponse, err error) {
	if request == nil {
		request = NewModifyRealServerAccessStrategyRequest()
	}
	response = NewModifyRealServerAccessStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceRemarkRequest() (request *ModifyInstanceRemarkRequest) {
	request = &ModifyInstanceRemarkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceRemark")
	return
}

func NewModifyInstanceRemarkResponse() (response *ModifyInstanceRemarkResponse) {
	response = &ModifyInstanceRemarkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyInstanceRemark）用于修改实例备注。
func (c *Client) ModifyInstanceRemark(request *ModifyInstanceRemarkRequest) (response *ModifyInstanceRemarkResponse, err error) {
	if request == nil {
		request = NewModifyInstanceRemarkRequest()
	}
	response = NewModifyInstanceRemarkResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
	request = &DisassociateSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DisassociateSecurityGroups")
	return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
	response = &DisassociateSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DisassociateSecurityGroups)用于安全组批量解绑实例。
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewDisassociateSecurityGroupsRequest()
	}
	response = NewDisassociateSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDatabaseTableRequest() (request *DescribeDatabaseTableRequest) {
	request = &DescribeDatabaseTableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabaseTable")
	return
}

func NewDescribeDatabaseTableResponse() (response *DescribeDatabaseTableResponse) {
	response = &DescribeDatabaseTableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDatabaseTable）用于查询云数据库实例的表信息。
func (c *Client) DescribeDatabaseTable(request *DescribeDatabaseTableRequest) (response *DescribeDatabaseTableResponse, err error) {
	if request == nil {
		request = NewDescribeDatabaseTableRequest()
	}
	response = NewDescribeDatabaseTableResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogFileRetentionPeriodRequest() (request *ModifyLogFileRetentionPeriodRequest) {
	request = &ModifyLogFileRetentionPeriodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyLogFileRetentionPeriod")
	return
}

func NewModifyLogFileRetentionPeriodResponse() (response *ModifyLogFileRetentionPeriodResponse) {
	response = &ModifyLogFileRetentionPeriodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifyLogFileRetentionPeriod)用于修改数据库备份日志保存天数。
func (c *Client) ModifyLogFileRetentionPeriod(request *ModifyLogFileRetentionPeriodRequest) (response *ModifyLogFileRetentionPeriodResponse, err error) {
	if request == nil {
		request = NewModifyLogFileRetentionPeriodRequest()
	}
	response = NewModifyLogFileRetentionPeriodResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableExclusiveGroupsRequest() (request *DescribeAvailableExclusiveGroupsRequest) {
	request = &DescribeAvailableExclusiveGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAvailableExclusiveGroups")
	return
}

func NewDescribeAvailableExclusiveGroupsResponse() (response *DescribeAvailableExclusiveGroupsResponse) {
	response = &DescribeAvailableExclusiveGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeAvailableExclusiveGroups)用于拉取独享资源池信息
func (c *Client) DescribeAvailableExclusiveGroups(request *DescribeAvailableExclusiveGroupsRequest) (response *DescribeAvailableExclusiveGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableExclusiveGroupsRequest()
	}
	response = NewDescribeAvailableExclusiveGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
	request = &DescribeDatabasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabases")
	return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
	response = &DescribeDatabasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDatabases）用于查询云数据库实例的数据库列表。
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
	if request == nil {
		request = NewDescribeDatabasesRequest()
	}
	response = NewDescribeDatabasesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
	request = &ModifyDBInstanceSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBInstanceSecurityGroups")
	return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
	response = &ModifyDBInstanceSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDBInstanceSecurityGroups）用于修改云数据库安全组
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewModifyDBInstanceSecurityGroupsRequest()
	}
	response = NewModifyDBInstanceSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewIsolateHourDCDBInstanceRequest() (request *IsolateHourDCDBInstanceRequest) {
	request = &IsolateHourDCDBInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "IsolateHourDCDBInstance")
	return
}

func NewIsolateHourDCDBInstanceResponse() (response *IsolateHourDCDBInstanceResponse) {
	response = &IsolateHourDCDBInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（IsolateHourDCDBInstance）用于销毁/退货按量计费实例。
// 实例将会保留3天再进行真实的销毁动作，在彻底销毁前可以通过 ActiveHourDCDBInstance 恢复/开机。实例彻底销毁后数据将无法找回，请提前备份实例数据。实例彻底销毁后IP资源同时释放。
func (c *Client) IsolateHourDCDBInstance(request *IsolateHourDCDBInstanceRequest) (response *IsolateHourDCDBInstanceResponse, err error) {
	if request == nil {
		request = NewIsolateHourDCDBInstanceRequest()
	}
	response = NewIsolateHourDCDBInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBCharsetsRequest() (request *DescribeDBCharsetsRequest) {
	request = &DescribeDBCharsetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBCharsets")
	return
}

func NewDescribeDBCharsetsResponse() (response *DescribeDBCharsetsResponse) {
	response = &DescribeDBCharsetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DB字符集信息列表
func (c *Client) DescribeDBCharsets(request *DescribeDBCharsetsRequest) (response *DescribeDBCharsetsResponse, err error) {
	if request == nil {
		request = NewDescribeDBCharsetsRequest()
	}
	response = NewDescribeDBCharsetsResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
	request = &AssociateSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "AssociateSecurityGroups")
	return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
	response = &AssociateSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (AssociateSecurityGroups) 用于安全组批量绑定云资源。
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewAssociateSecurityGroupsRequest()
	}
	response = NewAssociateSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBSlowLogAnalysisRequest() (request *DescribeDBSlowLogAnalysisRequest) {
	request = &DescribeDBSlowLogAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSlowLogAnalysis")
	return
}

func NewDescribeDBSlowLogAnalysisResponse() (response *DescribeDBSlowLogAnalysisResponse) {
	response = &DescribeDBSlowLogAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDBSlowLogAnalysis)用于获取慢查询记录详情。
func (c *Client) DescribeDBSlowLogAnalysis(request *DescribeDBSlowLogAnalysisRequest) (response *DescribeDBSlowLogAnalysisResponse, err error) {
	if request == nil {
		request = NewDescribeDBSlowLogAnalysisRequest()
	}
	response = NewDescribeDBSlowLogAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewStartSmartDBARequest() (request *StartSmartDBARequest) {
	request = &StartSmartDBARequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "StartSmartDBA")
	return
}

func NewStartSmartDBAResponse() (response *StartSmartDBAResponse) {
	response = &StartSmartDBAResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（StartSmartDBA）用于启动性能检测任务。
func (c *Client) StartSmartDBA(request *StartSmartDBARequest) (response *StartSmartDBAResponse, err error) {
	if request == nil {
		request = NewStartSmartDBARequest()
	}
	response = NewStartSmartDBAResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDatabaseObjectsRequest() (request *DescribeDatabaseObjectsRequest) {
	request = &DescribeDatabaseObjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabaseObjects")
	return
}

func NewDescribeDatabaseObjectsResponse() (response *DescribeDatabaseObjectsResponse) {
	response = &DescribeDatabaseObjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDatabaseObjects）用于查询云数据库实例的数据库中的对象列表，包含表、存储过程、视图和函数。
func (c *Client) DescribeDatabaseObjects(request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
	if request == nil {
		request = NewDescribeDatabaseObjectsRequest()
	}
	response = NewDescribeDatabaseObjectsResponse()
	err = c.Send(request, response)
	return
}

func NewCancelDcnJobRequest() (request *CancelDcnJobRequest) {
	request = &CancelDcnJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "CancelDcnJob")
	return
}

func NewCancelDcnJobResponse() (response *CancelDcnJobResponse) {
	response = &CancelDcnJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消DCN同步
func (c *Client) CancelDcnJob(request *CancelDcnJobRequest) (response *CancelDcnJobResponse, err error) {
	if request == nil {
		request = NewCancelDcnJobRequest()
	}
	response = NewCancelDcnJobResponse()
	err = c.Send(request, response)
	return
}

func NewRestartDBInstancesRequest() (request *RestartDBInstancesRequest) {
	request = &RestartDBInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "RestartDBInstances")
	return
}

func NewRestartDBInstancesResponse() (response *RestartDBInstancesResponse) {
	response = &RestartDBInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（RestartDBInstances）用于重启数据库实例
func (c *Client) RestartDBInstances(request *RestartDBInstancesRequest) (response *RestartDBInstancesResponse, err error) {
	if request == nil {
		request = NewRestartDBInstancesRequest()
	}
	response = NewRestartDBInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDCDBBinlogTimeRequest() (request *DescribeDCDBBinlogTimeRequest) {
	request = &DescribeDCDBBinlogTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBBinlogTime")
	return
}

func NewDescribeDCDBBinlogTimeResponse() (response *DescribeDCDBBinlogTimeResponse) {
	response = &DescribeDCDBBinlogTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例回档时可选的时间范围
func (c *Client) DescribeDCDBBinlogTime(request *DescribeDCDBBinlogTimeRequest) (response *DescribeDCDBBinlogTimeResponse, err error) {
	if request == nil {
		request = NewDescribeDCDBBinlogTimeRequest()
	}
	response = NewDescribeDCDBBinlogTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBEnginesRequest() (request *DescribeDBEnginesRequest) {
	request = &DescribeDBEnginesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBEngines")
	return
}

func NewDescribeDBEnginesResponse() (response *DescribeDBEnginesResponse) {
	response = &DescribeDBEnginesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DB引擎版本列表
func (c *Client) DescribeDBEngines(request *DescribeDBEnginesRequest) (response *DescribeDBEnginesResponse, err error) {
	if request == nil {
		request = NewDescribeDBEnginesRequest()
	}
	response = NewDescribeDBEnginesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDCDBRollbackInstancesRequest() (request *DescribeDCDBRollbackInstancesRequest) {
	request = &DescribeDCDBRollbackInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBRollbackInstances")
	return
}

func NewDescribeDCDBRollbackInstancesResponse() (response *DescribeDCDBRollbackInstancesResponse) {
	response = &DescribeDCDBRollbackInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询分布式回档实例列表
func (c *Client) DescribeDCDBRollbackInstances(request *DescribeDCDBRollbackInstancesRequest) (response *DescribeDCDBRollbackInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeDCDBRollbackInstancesRequest()
	}
	response = NewDescribeDCDBRollbackInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBSyncModeRequest() (request *DescribeDBSyncModeRequest) {
	request = &DescribeDBSyncModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSyncMode")
	return
}

func NewDescribeDBSyncModeResponse() (response *DescribeDBSyncModeResponse) {
	response = &DescribeDBSyncModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDBSyncMode）用于查询云数据库实例的同步模式。
func (c *Client) DescribeDBSyncMode(request *DescribeDBSyncModeRequest) (response *DescribeDBSyncModeResponse, err error) {
	if request == nil {
		request = NewDescribeDBSyncModeRequest()
	}
	response = NewDescribeDBSyncModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeInstances）用于云监控拉取实例列表
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchRollbackInstanceRequest() (request *SwitchRollbackInstanceRequest) {
	request = &SwitchRollbackInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "SwitchRollbackInstance")
	return
}

func NewSwitchRollbackInstanceResponse() (response *SwitchRollbackInstanceResponse) {
	response = &SwitchRollbackInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SwitchRollbackInstance）用于切换回档实例。
func (c *Client) SwitchRollbackInstance(request *SwitchRollbackInstanceRequest) (response *SwitchRollbackInstanceResponse, err error) {
	if request == nil {
		request = NewSwitchRollbackInstanceRequest()
	}
	response = NewSwitchRollbackInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceNetworkRequest() (request *ModifyInstanceNetworkRequest) {
	request = &ModifyInstanceNetworkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceNetwork")
	return
}

func NewModifyInstanceNetworkResponse() (response *ModifyInstanceNetworkResponse) {
	response = &ModifyInstanceNetworkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyInstanceNetwork）用于修改实例所属网络。
func (c *Client) ModifyInstanceNetwork(request *ModifyInstanceNetworkRequest) (response *ModifyInstanceNetworkResponse, err error) {
	if request == nil {
		request = NewModifyInstanceNetworkRequest()
	}
	response = NewModifyInstanceNetworkResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBEncryptAttributesRequest() (request *DescribeDBEncryptAttributesRequest) {
	request = &DescribeDBEncryptAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBEncryptAttributes")
	return
}

func NewDescribeDBEncryptAttributesResponse() (response *DescribeDBEncryptAttributesResponse) {
	response = &DescribeDBEncryptAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDBEncryptAttributes)用于查询实例数据加密状态。
func (c *Client) DescribeDBEncryptAttributes(request *DescribeDBEncryptAttributesRequest) (response *DescribeDBEncryptAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeDBEncryptAttributesRequest()
	}
	response = NewDescribeDBEncryptAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
	request = &DescribeFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeFlow")
	return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
	response = &DescribeFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeFlow）用于查询流程状态
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
	if request == nil {
		request = NewDescribeFlowRequest()
	}
	response = NewDescribeFlowResponse()
	err = c.Send(request, response)
	return
}

func NewActiveHourDCDBInstanceRequest() (request *ActiveHourDCDBInstanceRequest) {
	request = &ActiveHourDCDBInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ActiveHourDCDBInstance")
	return
}

func NewActiveHourDCDBInstanceResponse() (response *ActiveHourDCDBInstanceResponse) {
	response = &ActiveHourDCDBInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解隔离后付费实例
func (c *Client) ActiveHourDCDBInstance(request *ActiveHourDCDBInstanceRequest) (response *ActiveHourDCDBInstanceResponse, err error) {
	if request == nil {
		request = NewActiveHourDCDBInstanceRequest()
	}
	response = NewActiveHourDCDBInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBSlowLogsRequest() (request *DescribeDBSlowLogsRequest) {
	request = &DescribeDBSlowLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSlowLogs")
	return
}

func NewDescribeDBSlowLogsResponse() (response *DescribeDBSlowLogsResponse) {
	response = &DescribeDBSlowLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDBSlowLogs)用于查询慢查询日志列表。
func (c *Client) DescribeDBSlowLogs(request *DescribeDBSlowLogsRequest) (response *DescribeDBSlowLogsResponse, err error) {
	if request == nil {
		request = NewDescribeDBSlowLogsRequest()
	}
	response = NewDescribeDBSlowLogsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTmpDCDBInstanceRequest() (request *CreateTmpDCDBInstanceRequest) {
	request = &CreateTmpDCDBInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "CreateTmpDCDBInstance")
	return
}

func NewCreateTmpDCDBInstanceResponse() (response *CreateTmpDCDBInstanceResponse) {
	response = &CreateTmpDCDBInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回档实例
func (c *Client) CreateTmpDCDBInstance(request *CreateTmpDCDBInstanceRequest) (response *CreateTmpDCDBInstanceResponse, err error) {
	if request == nil {
		request = NewCreateTmpDCDBInstanceRequest()
	}
	response = NewCreateTmpDCDBInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
	request = &DescribeProjectSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeProjectSecurityGroups")
	return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
	response = &DescribeProjectSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeProjectSecurityGroups）用于查询项目安全组信息
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeProjectSecurityGroupsRequest()
	}
	response = NewDescribeProjectSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceVportRequest() (request *ModifyInstanceVportRequest) {
	request = &ModifyInstanceVportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceVport")
	return
}

func NewModifyInstanceVportResponse() (response *ModifyInstanceVportResponse) {
	response = &ModifyInstanceVportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyInstanceVport）用于修改实例VPORT
func (c *Client) ModifyInstanceVport(request *ModifyInstanceVportRequest) (response *ModifyInstanceVportResponse, err error) {
	if request == nil {
		request = NewModifyInstanceVportRequest()
	}
	response = NewModifyInstanceVportResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceVipRequest() (request *ModifyInstanceVipRequest) {
	request = &ModifyInstanceVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceVip")
	return
}

func NewModifyInstanceVipResponse() (response *ModifyInstanceVipResponse) {
	response = &ModifyInstanceVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyInstanceVip）用于修改实例Vip
func (c *Client) ModifyInstanceVip(request *ModifyInstanceVipRequest) (response *ModifyInstanceVipResponse, err error) {
	if request == nil {
		request = NewModifyInstanceVipRequest()
	}
	response = NewModifyInstanceVipResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
	request = &DescribeAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAccounts")
	return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
	response = &DescribeAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAccounts）用于查询指定云数据库实例的账号列表。
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeAccountsRequest()
	}
	response = NewDescribeAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceSSLAttributesRequest() (request *ModifyInstanceSSLAttributesRequest) {
	request = &ModifyInstanceSSLAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceSSLAttributes")
	return
}

func NewModifyInstanceSSLAttributesResponse() (response *ModifyInstanceSSLAttributesResponse) {
	response = &ModifyInstanceSSLAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分布式实例开启SSL
func (c *Client) ModifyInstanceSSLAttributes(request *ModifyInstanceSSLAttributesRequest) (response *ModifyInstanceSSLAttributesResponse, err error) {
	if request == nil {
		request = NewModifyInstanceSSLAttributesRequest()
	}
	response = NewModifyInstanceSSLAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHourDCDBInstanceRequest() (request *CreateHourDCDBInstanceRequest) {
	request = &CreateHourDCDBInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "CreateHourDCDBInstance")
	return
}

func NewCreateHourDCDBInstanceResponse() (response *CreateHourDCDBInstanceResponse) {
	response = &CreateHourDCDBInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建后付费实例
func (c *Client) CreateHourDCDBInstance(request *CreateHourDCDBInstanceRequest) (response *CreateHourDCDBInstanceResponse, err error) {
	if request == nil {
		request = NewCreateHourDCDBInstanceRequest()
	}
	response = NewCreateHourDCDBInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLatestCloudDBAReportRequest() (request *DescribeLatestCloudDBAReportRequest) {
	request = &DescribeLatestCloudDBAReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeLatestCloudDBAReport")
	return
}

func NewDescribeLatestCloudDBAReportResponse() (response *DescribeLatestCloudDBAReportResponse) {
	response = &DescribeLatestCloudDBAReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeLatestCloudDBAReport）用于获取最新的性能检测报告
func (c *Client) DescribeLatestCloudDBAReport(request *DescribeLatestCloudDBAReportRequest) (response *DescribeLatestCloudDBAReportResponse, err error) {
	if request == nil {
		request = NewDescribeLatestCloudDBAReportRequest()
	}
	response = NewDescribeLatestCloudDBAReportResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccountPrivilegesRequest() (request *ModifyAccountPrivilegesRequest) {
	request = &ModifyAccountPrivilegesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyAccountPrivileges")
	return
}

func NewModifyAccountPrivilegesResponse() (response *ModifyAccountPrivilegesResponse) {
	response = &ModifyAccountPrivilegesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyAccountPrivileges）用于给云数据库账号赋权。
// 注意：相同用户名，不同Host是不同的账号。
func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
	if request == nil {
		request = NewModifyAccountPrivilegesRequest()
	}
	response = NewModifyAccountPrivilegesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDBParametersRequest() (request *ModifyDBParametersRequest) {
	request = &ModifyDBParametersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBParameters")
	return
}

func NewModifyDBParametersResponse() (response *ModifyDBParametersResponse) {
	response = &ModifyDBParametersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifyDBParameters)用于修改数据库参数。
func (c *Client) ModifyDBParameters(request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
	if request == nil {
		request = NewModifyDBParametersRequest()
	}
	response = NewModifyDBParametersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZonesArchInfoRequest() (request *DescribeZonesArchInfoRequest) {
	request = &DescribeZonesArchInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeZonesArchInfo")
	return
}

func NewDescribeZonesArchInfoResponse() (response *DescribeZonesArchInfoResponse) {
	response = &DescribeZonesArchInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据主从可用区返回可用Cpu架构查询
func (c *Client) DescribeZonesArchInfo(request *DescribeZonesArchInfoRequest) (response *DescribeZonesArchInfoResponse, err error) {
	if request == nil {
		request = NewDescribeZonesArchInfoRequest()
	}
	response = NewDescribeZonesArchInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
	request = &DescribeDBSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSecurityGroups")
	return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
	response = &DescribeDBSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDBSecurityGroups）用于查询实例安全组信息
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeDBSecurityGroupsRequest()
	}
	response = NewDescribeDBSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDcnDetailRequest() (request *DescribeDcnDetailRequest) {
	request = &DescribeDcnDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDcnDetail")
	return
}

func NewDescribeDcnDetailResponse() (response *DescribeDcnDetailResponse) {
	response = &DescribeDcnDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例灾备详情
func (c *Client) DescribeDcnDetail(request *DescribeDcnDetailRequest) (response *DescribeDcnDetailResponse, err error) {
	if request == nil {
		request = NewDescribeDcnDetailRequest()
	}
	response = NewDescribeDcnDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCommonDBInstancesRequest() (request *DescribeCommonDBInstancesRequest) {
	request = &DescribeCommonDBInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeCommonDBInstances")
	return
}

func NewDescribeCommonDBInstancesResponse() (response *DescribeCommonDBInstancesResponse) {
	response = &DescribeCommonDBInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 内部调用，为数据库产品中心对外的统一接口。用于拉取实例的信息，提供多个筛选参数。
func (c *Client) DescribeCommonDBInstances(request *DescribeCommonDBInstancesRequest) (response *DescribeCommonDBInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeCommonDBInstancesRequest()
	}
	response = NewDescribeCommonDBInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceProxyConfigRequest() (request *DescribeInstanceProxyConfigRequest) {
	request = &DescribeInstanceProxyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeInstanceProxyConfig")
	return
}

func NewDescribeInstanceProxyConfigResponse() (response *DescribeInstanceProxyConfigResponse) {
	response = &DescribeInstanceProxyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取实例网关配置
func (c *Client) DescribeInstanceProxyConfig(request *DescribeInstanceProxyConfigRequest) (response *DescribeInstanceProxyConfigResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceProxyConfigRequest()
	}
	response = NewDescribeInstanceProxyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewFlushBinlogRequest() (request *FlushBinlogRequest) {
	request = &FlushBinlogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "FlushBinlog")
	return
}

func NewFlushBinlogResponse() (response *FlushBinlogResponse) {
	response = &FlushBinlogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 相当于在所有分片的mysqld中执行flush logs，完成切分的binlog将展示在各个分片控制台binlog列表里。
func (c *Client) FlushBinlog(request *FlushBinlogRequest) (response *FlushBinlogResponse, err error) {
	if request == nil {
		request = NewFlushBinlogRequest()
	}
	response = NewFlushBinlogResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
	request = &CreateAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "CreateAccount")
	return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
	response = &CreateAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateAccount）用于创建云数据库账号。一个实例可以创建多个不同的账号，相同的用户名+不同的host是不同的账号。
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
	if request == nil {
		request = NewCreateAccountRequest()
	}
	response = NewCreateAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBParametersRequest() (request *DescribeDBParametersRequest) {
	request = &DescribeDBParametersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBParameters")
	return
}

func NewDescribeDBParametersResponse() (response *DescribeDBParametersResponse) {
	response = &DescribeDBParametersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDBParameters)用于获取数据库的当前参数设置。
func (c *Client) DescribeDBParameters(request *DescribeDBParametersRequest) (response *DescribeDBParametersResponse, err error) {
	if request == nil {
		request = NewDescribeDBParametersRequest()
	}
	response = NewDescribeDBParametersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogFileRetentionPeriodRequest() (request *DescribeLogFileRetentionPeriodRequest) {
	request = &DescribeLogFileRetentionPeriodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeLogFileRetentionPeriod")
	return
}

func NewDescribeLogFileRetentionPeriodResponse() (response *DescribeLogFileRetentionPeriodResponse) {
	response = &DescribeLogFileRetentionPeriodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeLogFileRetentionPeriod)用于查看数据库备份日志的备份天数的设置情况。
func (c *Client) DescribeLogFileRetentionPeriod(request *DescribeLogFileRetentionPeriodRequest) (response *DescribeLogFileRetentionPeriodResponse, err error) {
	if request == nil {
		request = NewDescribeLogFileRetentionPeriodRequest()
	}
	response = NewDescribeLogFileRetentionPeriodResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDCDBSaleInfoRequest() (request *DescribeDCDBSaleInfoRequest) {
	request = &DescribeDCDBSaleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBSaleInfo")
	return
}

func NewDescribeDCDBSaleInfoResponse() (response *DescribeDCDBSaleInfoResponse) {
	response = &DescribeDCDBSaleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDCDBSaleInfo)用于查询分布式数据库可售卖的地域和可用区信息。
func (c *Client) DescribeDCDBSaleInfo(request *DescribeDCDBSaleInfoRequest) (response *DescribeDCDBSaleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDCDBSaleInfoRequest()
	}
	response = NewDescribeDCDBSaleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewInitDCDBInstancesRequest() (request *InitDCDBInstancesRequest) {
	request = &InitDCDBInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "InitDCDBInstances")
	return
}

func NewInitDCDBInstancesResponse() (response *InitDCDBInstancesResponse) {
	response = &InitDCDBInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(InitDCDBInstances)用于初始化云数据库实例，包括设置默认字符集、表名大小写敏感等。
func (c *Client) InitDCDBInstances(request *InitDCDBInstancesRequest) (response *InitDCDBInstancesResponse, err error) {
	if request == nil {
		request = NewInitDCDBInstancesRequest()
	}
	response = NewInitDCDBInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTmpInstanceRequest() (request *DeleteTmpInstanceRequest) {
	request = &DeleteTmpInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DeleteTmpInstance")
	return
}

func NewDeleteTmpInstanceResponse() (response *DeleteTmpInstanceResponse) {
	response = &DeleteTmpInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteTmpInstance）用于删除临时实例。
func (c *Client) DeleteTmpInstance(request *DeleteTmpInstanceRequest) (response *DeleteTmpInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteTmpInstanceRequest()
	}
	response = NewDeleteTmpInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
	request = &ModifyDBInstanceNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBInstanceName")
	return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
	response = &ModifyDBInstanceNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDBInstanceName）用于修改实例名字
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
	if request == nil {
		request = NewModifyDBInstanceNameRequest()
	}
	response = NewModifyDBInstanceNameResponse()
	err = c.Send(request, response)
	return
}

func NewCopyAccountPrivilegesRequest() (request *CopyAccountPrivilegesRequest) {
	request = &CopyAccountPrivilegesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "CopyAccountPrivileges")
	return
}

func NewCopyAccountPrivilegesResponse() (response *CopyAccountPrivilegesResponse) {
	response = &CopyAccountPrivilegesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CopyAccountPrivileges）用于复制云数据库账号的权限。
// 注意：相同用户名，不同Host是不同的账号，Readonly属性相同的账号之间才能复制权限。
func (c *Client) CopyAccountPrivileges(request *CopyAccountPrivilegesRequest) (response *CopyAccountPrivilegesResponse, err error) {
	if request == nil {
		request = NewCopyAccountPrivilegesRequest()
	}
	response = NewCopyAccountPrivilegesResponse()
	err = c.Send(request, response)
	return
}

func NewCloneAccountRequest() (request *CloneAccountRequest) {
	request = &CloneAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "CloneAccount")
	return
}

func NewCloneAccountResponse() (response *CloneAccountResponse) {
	response = &CloneAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CloneAccount）用于克隆实例账户。
func (c *Client) CloneAccount(request *CloneAccountRequest) (response *CloneAccountResponse, err error) {
	if request == nil {
		request = NewCloneAccountRequest()
	}
	response = NewCloneAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDCDBInstancesRequest() (request *DescribeDCDBInstancesRequest) {
	request = &DescribeDCDBInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstances")
	return
}

func NewDescribeDCDBInstancesResponse() (response *DescribeDCDBInstancesResponse) {
	response = &DescribeDCDBInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云数据库实例列表，支持通过项目ID、实例ID、内网地址、实例名称等来筛选实例。
// 如果不指定任何筛选条件，则默认返回10条实例记录，单次请求最多支持返回100条实例记录。
func (c *Client) DescribeDCDBInstances(request *DescribeDCDBInstancesRequest) (response *DescribeDCDBInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeDCDBInstancesRequest()
	}
	response = NewDescribeDCDBInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
	request = &DeleteAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DeleteAccount")
	return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
	response = &DeleteAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteAccount）用于删除云数据库账号。用户名+host唯一确定一个账号。
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
	if request == nil {
		request = NewDeleteAccountRequest()
	}
	response = NewDeleteAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBTmpInstancesRequest() (request *DescribeDBTmpInstancesRequest) {
	request = &DescribeDBTmpInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBTmpInstances")
	return
}

func NewDescribeDBTmpInstancesResponse() (response *DescribeDBTmpInstancesResponse) {
	response = &DescribeDBTmpInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDBTmpInstances）用于获取实例回档生成的临时实例
func (c *Client) DescribeDBTmpInstances(request *DescribeDBTmpInstancesRequest) (response *DescribeDBTmpInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeDBTmpInstancesRequest()
	}
	response = NewDescribeDBTmpInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserTasksRequest() (request *DescribeUserTasksRequest) {
	request = &DescribeUserTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeUserTasks")
	return
}

func NewDescribeUserTasksResponse() (response *DescribeUserTasksResponse) {
	response = &DescribeUserTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeUserTasks）用于拉取用户任务列表
func (c *Client) DescribeUserTasks(request *DescribeUserTasksRequest) (response *DescribeUserTasksResponse, err error) {
	if request == nil {
		request = NewDescribeUserTasksRequest()
	}
	response = NewDescribeUserTasksResponse()
	err = c.Send(request, response)
	return
}

func NewAuthenticateCAMRequest() (request *AuthenticateCAMRequest) {
	request = &AuthenticateCAMRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "AuthenticateCAM")
	return
}

func NewAuthenticateCAMResponse() (response *AuthenticateCAMResponse) {
	response = &AuthenticateCAMResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AuthenticateCAM）为控制台提供预鉴权
func (c *Client) AuthenticateCAM(request *AuthenticateCAMRequest) (response *AuthenticateCAMResponse, err error) {
	if request == nil {
		request = NewAuthenticateCAMRequest()
	}
	response = NewAuthenticateCAMResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseNetResourceRequest() (request *ReleaseNetResourceRequest) {
	request = &ReleaseNetResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ReleaseNetResource")
	return
}

func NewReleaseNetResourceResponse() (response *ReleaseNetResourceResponse) {
	response = &ReleaseNetResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收保留的网络资源
func (c *Client) ReleaseNetResource(request *ReleaseNetResourceRequest) (response *ReleaseNetResourceResponse, err error) {
	if request == nil {
		request = NewReleaseNetResourceRequest()
	}
	response = NewReleaseNetResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDCDBShardsRequest() (request *DescribeDCDBShardsRequest) {
	request = &DescribeDCDBShardsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBShards")
	return
}

func NewDescribeDCDBShardsResponse() (response *DescribeDCDBShardsResponse) {
	response = &DescribeDCDBShardsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDCDBShards）用于查询云数据库实例的分片信息。
func (c *Client) DescribeDCDBShards(request *DescribeDCDBShardsRequest) (response *DescribeDCDBShardsResponse, err error) {
	if request == nil {
		request = NewDescribeDCDBShardsRequest()
	}
	response = NewDescribeDCDBShardsResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeHourDCDBInstanceRequest() (request *UpgradeHourDCDBInstanceRequest) {
	request = &UpgradeHourDCDBInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "UpgradeHourDCDBInstance")
	return
}

func NewUpgradeHourDCDBInstanceResponse() (response *UpgradeHourDCDBInstanceResponse) {
	response = &UpgradeHourDCDBInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpgradeHourDCDBInstance）用于升级后付费分布式数据库实例。
func (c *Client) UpgradeHourDCDBInstance(request *UpgradeHourDCDBInstanceRequest) (response *UpgradeHourDCDBInstanceResponse, err error) {
	if request == nil {
		request = NewUpgradeHourDCDBInstanceRequest()
	}
	response = NewUpgradeHourDCDBInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCheckIpStatusRequest() (request *CheckIpStatusRequest) {
	request = &CheckIpStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "CheckIpStatus")
	return
}

func NewCheckIpStatusResponse() (response *CheckIpStatusResponse) {
	response = &CheckIpStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CheckIpStatus)用于查询指定的私有网络中的虚拟IP是否可用。
func (c *Client) CheckIpStatus(request *CheckIpStatusRequest) (response *CheckIpStatusResponse, err error) {
	if request == nil {
		request = NewCheckIpStatusRequest()
	}
	response = NewCheckIpStatusResponse()
	err = c.Send(request, response)
	return
}

func NewKillSessionRequest() (request *KillSessionRequest) {
	request = &KillSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "KillSession")
	return
}

func NewKillSessionResponse() (response *KillSessionResponse) {
	response = &KillSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（KillSession）用于杀死指定会话。
func (c *Client) KillSession(request *KillSessionRequest) (response *KillSessionResponse, err error) {
	if request == nil {
		request = NewKillSessionRequest()
	}
	response = NewKillSessionResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchDcnJobRequest() (request *SwitchDcnJobRequest) {
	request = &SwitchDcnJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "SwitchDcnJob")
	return
}

func NewSwitchDcnJobResponse() (response *SwitchDcnJobResponse) {
	response = &SwitchDcnJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DCN主备切换
func (c *Client) SwitchDcnJob(request *SwitchDcnJobRequest) (response *SwitchDcnJobResponse, err error) {
	if request == nil {
		request = NewSwitchDcnJobRequest()
	}
	response = NewSwitchDcnJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDCDBInstanceNodeInfoRequest() (request *DescribeDCDBInstanceNodeInfoRequest) {
	request = &DescribeDCDBInstanceNodeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstanceNodeInfo")
	return
}

func NewDescribeDCDBInstanceNodeInfoResponse() (response *DescribeDCDBInstanceNodeInfoResponse) {
	response = &DescribeDCDBInstanceNodeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDCDBInstanceNodeInfo）用于获取实例节点信息
func (c *Client) DescribeDCDBInstanceNodeInfo(request *DescribeDCDBInstanceNodeInfoRequest) (response *DescribeDCDBInstanceNodeInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDCDBInstanceNodeInfoRequest()
	}
	response = NewDescribeDCDBInstanceNodeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupFilesRequest() (request *DescribeBackupFilesRequest) {
	request = &DescribeBackupFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeBackupFiles")
	return
}

func NewDescribeBackupFilesResponse() (response *DescribeBackupFilesResponse) {
	response = &DescribeBackupFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询备份文件列表
func (c *Client) DescribeBackupFiles(request *DescribeBackupFilesRequest) (response *DescribeBackupFilesResponse, err error) {
	if request == nil {
		request = NewDescribeBackupFilesRequest()
	}
	response = NewDescribeBackupFilesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceSSLAttributesRequest() (request *DescribeInstanceSSLAttributesRequest) {
	request = &DescribeInstanceSSLAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeInstanceSSLAttributes")
	return
}

func NewDescribeInstanceSSLAttributesResponse() (response *DescribeInstanceSSLAttributesResponse) {
	response = &DescribeInstanceSSLAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取实例SSL认证功能属性
func (c *Client) DescribeInstanceSSLAttributes(request *DescribeInstanceSSLAttributesRequest) (response *DescribeInstanceSSLAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceSSLAttributesRequest()
	}
	response = NewDescribeInstanceSSLAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
	request = &DescribeAccountPrivilegesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAccountPrivileges")
	return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
	response = &DescribeAccountPrivilegesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAccountPrivileges）用于查询云数据库账号权限。
// 注意：注意：相同用户名，不同Host是不同的账号。
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
	if request == nil {
		request = NewDescribeAccountPrivilegesRequest()
	}
	response = NewDescribeAccountPrivilegesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDCDBPriceRequest() (request *DescribeDCDBPriceRequest) {
	request = &DescribeDCDBPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBPrice")
	return
}

func NewDescribeDCDBPriceResponse() (response *DescribeDCDBPriceResponse) {
	response = &DescribeDCDBPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDCDBPrice）用于在购买实例前，查询实例的价格。
// 注意：当前 TCE 上 Paymode 参数需要传 postpaid。
func (c *Client) DescribeDCDBPrice(request *DescribeDCDBPriceRequest) (response *DescribeDCDBPriceResponse, err error) {
	if request == nil {
		request = NewDescribeDCDBPriceRequest()
	}
	response = NewDescribeDCDBPriceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBInstanceRsipRequest() (request *DescribeDBInstanceRsipRequest) {
	request = &DescribeDBInstanceRsipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBInstanceRsip")
	return
}

func NewDescribeDBInstanceRsipResponse() (response *DescribeDBInstanceRsipResponse) {
	response = &DescribeDBInstanceRsipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDBInstanceRsip）用于获取实例Rsip
func (c *Client) DescribeDBInstanceRsip(request *DescribeDBInstanceRsipRequest) (response *DescribeDBInstanceRsipResponse, err error) {
	if request == nil {
		request = NewDescribeDBInstanceRsipRequest()
	}
	response = NewDescribeDBInstanceRsipResponse()
	err = c.Send(request, response)
	return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
	request = &ResetAccountPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ResetAccountPassword")
	return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
	response = &ResetAccountPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResetAccountPassword）用于重置云数据库账号的密码。
// 注意：相同用户名，不同Host是不同的账号。
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
	if request == nil {
		request = NewResetAccountPasswordRequest()
	}
	response = NewResetAccountPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDBEncryptAttributesRequest() (request *ModifyDBEncryptAttributesRequest) {
	request = &ModifyDBEncryptAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBEncryptAttributes")
	return
}

func NewModifyDBEncryptAttributesResponse() (response *ModifyDBEncryptAttributesResponse) {
	response = &ModifyDBEncryptAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifyDBEncryptAttributes)用于修改实例数据加密。
func (c *Client) ModifyDBEncryptAttributes(request *ModifyDBEncryptAttributesRequest) (response *ModifyDBEncryptAttributesResponse, err error) {
	if request == nil {
		request = NewModifyDBEncryptAttributesRequest()
	}
	response = NewModifyDBEncryptAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDBSyncModeRequest() (request *ModifyDBSyncModeRequest) {
	request = &ModifyDBSyncModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBSyncMode")
	return
}

func NewModifyDBSyncModeResponse() (response *ModifyDBSyncModeResponse) {
	response = &ModifyDBSyncModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDBSyncMode）用于修改云数据库实例的同步模式。
func (c *Client) ModifyDBSyncMode(request *ModifyDBSyncModeRequest) (response *ModifyDBSyncModeResponse, err error) {
	if request == nil {
		request = NewModifyDBSyncModeRequest()
	}
	response = NewModifyDBSyncModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDCDBInstanceShardInfoRequest() (request *DescribeDCDBInstanceShardInfoRequest) {
	request = &DescribeDCDBInstanceShardInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstanceShardInfo")
	return
}

func NewDescribeDCDBInstanceShardInfoResponse() (response *DescribeDCDBInstanceShardInfoResponse) {
	response = &DescribeDCDBInstanceShardInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDCDBInstanceShardInfo）用于获取实例分片信息
func (c *Client) DescribeDCDBInstanceShardInfo(request *DescribeDCDBInstanceShardInfoRequest) (response *DescribeDCDBInstanceShardInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDCDBInstanceShardInfoRequest()
	}
	response = NewDescribeDCDBInstanceShardInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDCDBInstanceDetailRequest() (request *DescribeDCDBInstanceDetailRequest) {
	request = &DescribeDCDBInstanceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstanceDetail")
	return
}

func NewDescribeDCDBInstanceDetailResponse() (response *DescribeDCDBInstanceDetailResponse) {
	response = &DescribeDCDBInstanceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDCDBInstanceDetail）用于获取实例详情
func (c *Client) DescribeDCDBInstanceDetail(request *DescribeDCDBInstanceDetailRequest) (response *DescribeDCDBInstanceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeDCDBInstanceDetailRequest()
	}
	response = NewDescribeDCDBInstanceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGrantAccountPrivilegesRequest() (request *GrantAccountPrivilegesRequest) {
	request = &GrantAccountPrivilegesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "GrantAccountPrivileges")
	return
}

func NewGrantAccountPrivilegesResponse() (response *GrantAccountPrivilegesResponse) {
	response = &GrantAccountPrivilegesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（GrantAccountPrivileges）用于给云数据库账号赋权。
// 注意：相同用户名，不同Host是不同的账号。
func (c *Client) GrantAccountPrivileges(request *GrantAccountPrivilegesRequest) (response *GrantAccountPrivilegesResponse, err error) {
	if request == nil {
		request = NewGrantAccountPrivilegesRequest()
	}
	response = NewGrantAccountPrivilegesResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyHourDCDBInstanceRequest() (request *DestroyHourDCDBInstanceRequest) {
	request = &DestroyHourDCDBInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DestroyHourDCDBInstance")
	return
}

func NewDestroyHourDCDBInstanceResponse() (response *DestroyHourDCDBInstanceResponse) {
	response = &DestroyHourDCDBInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DestroyHourDCDBInstance）用于销毁按量计费实例。
// 注意：本接口在当前版本已废弃，后续版本可能会移除。
func (c *Client) DestroyHourDCDBInstance(request *DestroyHourDCDBInstanceRequest) (response *DestroyHourDCDBInstanceResponse, err error) {
	if request == nil {
		request = NewDestroyHourDCDBInstanceRequest()
	}
	response = NewDestroyHourDCDBInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeShardSpecRequest() (request *DescribeShardSpecRequest) {
	request = &DescribeShardSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "DescribeShardSpec")
	return
}

func NewDescribeShardSpecResponse() (response *DescribeShardSpecResponse) {
	response = &DescribeShardSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可创建的分布式数据库可售卖的分片规格配置。
func (c *Client) DescribeShardSpec(request *DescribeShardSpecRequest) (response *DescribeShardSpecResponse, err error) {
	if request == nil {
		request = NewDescribeShardSpecRequest()
	}
	response = NewDescribeShardSpecResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
	request = &ModifyAccountDescriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcdb", APIVersion, "ModifyAccountDescription")
	return
}

func NewModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
	response = &ModifyAccountDescriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyAccountDescription）用于修改云数据库账号备注。
// 注意：相同用户名，不同Host是不同的账号。
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
	if request == nil {
		request = NewModifyAccountDescriptionRequest()
	}
	response = NewModifyAccountDescriptionResponse()
	err = c.Send(request, response)
	return
}
