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

package v20200107

import (
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-01-07"

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

func NewGetPriceRequest() (request *GetPriceRequest) {
	request = &GetPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetPrice")
	return
}

func NewGetPriceResponse() (response *GetPriceResponse) {
	response = &GetPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取计费明细
func (c *Client) GetPrice(request *GetPriceRequest) (response *GetPriceResponse, err error) {
	if request == nil {
		request = NewGetPriceRequest()
	}
	response = NewGetPriceResponse()
	err = c.Send(request, response)
	return
}

func NewListBucketsRequest() (request *ListBucketsRequest) {
	request = &ListBucketsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "ListBuckets")
	return
}

func NewListBucketsResponse() (response *ListBucketsResponse) {
	response = &ListBucketsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给控制台用的列存储桶接口
func (c *Client) ListBuckets(request *ListBucketsRequest) (response *ListBucketsResponse, err error) {
	if request == nil {
		request = NewListBucketsRequest()
	}
	response = NewListBucketsResponse()
	err = c.Send(request, response)
	return
}

func NewPutBucketDomainRequest() (request *PutBucketDomainRequest) {
	request = &PutBucketDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "PutBucketDomain")
	return
}

func NewPutBucketDomainResponse() (response *PutBucketDomainResponse) {
	response = &PutBucketDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置存储桶自定义域名
func (c *Client) PutBucketDomain(request *PutBucketDomainRequest) (response *PutBucketDomainResponse, err error) {
	if request == nil {
		request = NewPutBucketDomainRequest()
	}
	response = NewPutBucketDomainResponse()
	err = c.Send(request, response)
	return
}

func NewHeadObjectRequest() (request *HeadObjectRequest) {
	request = &HeadObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "HeadObject")
	return
}

func NewHeadObjectResponse() (response *HeadObjectResponse) {
	response = &HeadObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取对象属性
func (c *Client) HeadObject(request *HeadObjectRequest) (response *HeadObjectResponse, err error) {
	if request == nil {
		request = NewHeadObjectRequest()
	}
	response = NewHeadObjectResponse()
	err = c.Send(request, response)
	return
}

func NewSetBackupSettingRequest() (request *SetBackupSettingRequest) {
	request = &SetBackupSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetBackupSetting")
	return
}

func NewSetBackupSettingResponse() (response *SetBackupSettingResponse) {
	response = &SetBackupSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置存储桶备份配置
func (c *Client) SetBackupSetting(request *SetBackupSettingRequest) (response *SetBackupSettingResponse, err error) {
	if request == nil {
		request = NewSetBackupSettingRequest()
	}
	response = NewSetBackupSettingResponse()
	err = c.Send(request, response)
	return
}

func NewCheckKafkaRequest() (request *CheckKafkaRequest) {
	request = &CheckKafkaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "CheckKafka")
	return
}

func NewCheckKafkaResponse() (response *CheckKafkaResponse) {
	response = &CheckKafkaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查kafka连通性
func (c *Client) CheckKafka(request *CheckKafkaRequest) (response *CheckKafkaResponse, err error) {
	if request == nil {
		request = NewCheckKafkaRequest()
	}
	response = NewCheckKafkaResponse()
	err = c.Send(request, response)
	return
}

func NewGetBucketDomainRequest() (request *GetBucketDomainRequest) {
	request = &GetBucketDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBucketDomain")
	return
}

func NewGetBucketDomainResponse() (response *GetBucketDomainResponse) {
	response = &GetBucketDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储桶自定义域名
func (c *Client) GetBucketDomain(request *GetBucketDomainRequest) (response *GetBucketDomainResponse, err error) {
	if request == nil {
		request = NewGetBucketDomainRequest()
	}
	response = NewGetBucketDomainResponse()
	err = c.Send(request, response)
	return
}

func NewGetOverviewRequest() (request *GetOverviewRequest) {
	request = &GetOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetOverview")
	return
}

func NewGetOverviewResponse() (response *GetOverviewResponse) {
	response = &GetOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的概览信息
func (c *Client) GetOverview(request *GetOverviewRequest) (response *GetOverviewResponse, err error) {
	if request == nil {
		request = NewGetOverviewRequest()
	}
	response = NewGetOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteObjectRequest() (request *DeleteObjectRequest) {
	request = &DeleteObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DeleteObject")
	return
}

func NewDeleteObjectResponse() (response *DeleteObjectResponse) {
	response = &DeleteObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DELETE Object 接口请求可以在 COS 的存储桶中将一个对象（Object）删除。该操作需要请求者对存储桶有 WRITE 权限。
func (c *Client) DeleteObject(request *DeleteObjectRequest) (response *DeleteObjectResponse, err error) {
	if request == nil {
		request = NewDeleteObjectRequest()
	}
	response = NewDeleteObjectResponse()
	err = c.Send(request, response)
	return
}

func NewPutBucketRequest() (request *PutBucketRequest) {
	request = &PutBucketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "PutBucket")
	return
}

func NewPutBucketResponse() (response *PutBucketResponse) {
	response = &PutBucketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建存储桶
func (c *Client) PutBucket(request *PutBucketRequest) (response *PutBucketResponse, err error) {
	if request == nil {
		request = NewPutBucketRequest()
	}
	response = NewPutBucketResponse()
	err = c.Send(request, response)
	return
}

func NewPutBucketRefererRequest() (request *PutBucketRefererRequest) {
	request = &PutBucketRefererRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "PutBucketReferer")
	return
}

func NewPutBucketRefererResponse() (response *PutBucketRefererResponse) {
	response = &PutBucketRefererResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为存储桶设置 Referer 白名单或者黑名单
func (c *Client) PutBucketReferer(request *PutBucketRefererRequest) (response *PutBucketRefererResponse, err error) {
	if request == nil {
		request = NewPutBucketRefererRequest()
	}
	response = NewPutBucketRefererResponse()
	err = c.Send(request, response)
	return
}

func NewGetBackupDetailRequest() (request *GetBackupDetailRequest) {
	request = &GetBackupDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBackupDetail")
	return
}

func NewGetBackupDetailResponse() (response *GetBackupDetailResponse) {
	response = &GetBackupDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储桶备份详情
func (c *Client) GetBackupDetail(request *GetBackupDetailRequest) (response *GetBackupDetailResponse, err error) {
	if request == nil {
		request = NewGetBackupDetailRequest()
	}
	response = NewGetBackupDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetServiceRequest() (request *GetServiceRequest) {
	request = &GetServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetService")
	return
}

func NewGetServiceResponse() (response *GetServiceResponse) {
	response = &GetServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储桶列表
func (c *Client) GetService(request *GetServiceRequest) (response *GetServiceResponse, err error) {
	if request == nil {
		request = NewGetServiceRequest()
	}
	response = NewGetServiceResponse()
	err = c.Send(request, response)
	return
}

func NewGetStatDayRequest() (request *GetStatDayRequest) {
	request = &GetStatDayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetStatDay")
	return
}

func NewGetStatDayResponse() (response *GetStatDayResponse) {
	response = &GetStatDayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单日单存储桶统计信息
func (c *Client) GetStatDay(request *GetStatDayRequest) (response *GetStatDayResponse, err error) {
	if request == nil {
		request = NewGetStatDayRequest()
	}
	response = NewGetStatDayResponse()
	err = c.Send(request, response)
	return
}

func NewGetBucketRefererRequest() (request *GetBucketRefererRequest) {
	request = &GetBucketRefererRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBucketReferer")
	return
}

func NewGetBucketRefererResponse() (response *GetBucketRefererResponse) {
	response = &GetBucketRefererResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储桶 Referer 白名单或者黑名单
func (c *Client) GetBucketReferer(request *GetBucketRefererRequest) (response *GetBucketRefererResponse, err error) {
	if request == nil {
		request = NewGetBucketRefererRequest()
	}
	response = NewGetBucketRefererResponse()
	err = c.Send(request, response)
	return
}

func NewGetStatDaysRequest() (request *GetStatDaysRequest) {
	request = &GetStatDaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetStatDays")
	return
}

func NewGetStatDaysResponse() (response *GetStatDaysResponse) {
	response = &GetStatDaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取多日单存储桶的统计信息
func (c *Client) GetStatDays(request *GetStatDaysRequest) (response *GetStatDaysResponse, err error) {
	if request == nil {
		request = NewGetStatDaysRequest()
	}
	response = NewGetStatDaysResponse()
	err = c.Send(request, response)
	return
}

func NewGetRegionListRequest() (request *GetRegionListRequest) {
	request = &GetRegionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRegionList")
	return
}

func NewGetRegionListResponse() (response *GetRegionListResponse) {
	response = &GetRegionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户地域列表
func (c *Client) GetRegionList(request *GetRegionListRequest) (response *GetRegionListResponse, err error) {
	if request == nil {
		request = NewGetRegionListRequest()
	}
	response = NewGetRegionListResponse()
	err = c.Send(request, response)
	return
}

func NewGetStatHttpDayRequest() (request *GetStatHttpDayRequest) {
	request = &GetStatHttpDayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetStatHttpDay")
	return
}

func NewGetStatHttpDayResponse() (response *GetStatHttpDayResponse) {
	response = &GetStatHttpDayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单日单存储桶的http请求统计信息
func (c *Client) GetStatHttpDay(request *GetStatHttpDayRequest) (response *GetStatHttpDayResponse, err error) {
	if request == nil {
		request = NewGetStatHttpDayRequest()
	}
	response = NewGetStatHttpDayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBucketDomainRequest() (request *DeleteBucketDomainRequest) {
	request = &DeleteBucketDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DeleteBucketDomain")
	return
}

func NewDeleteBucketDomainResponse() (response *DeleteBucketDomainResponse) {
	response = &DeleteBucketDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除存储桶自定义域名
func (c *Client) DeleteBucketDomain(request *DeleteBucketDomainRequest) (response *DeleteBucketDomainResponse, err error) {
	if request == nil {
		request = NewDeleteBucketDomainRequest()
	}
	response = NewDeleteBucketDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMultipleObjectRequest() (request *DeleteMultipleObjectRequest) {
	request = &DeleteMultipleObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DeleteMultipleObject")
	return
}

func NewDeleteMultipleObjectResponse() (response *DeleteMultipleObjectResponse) {
	response = &DeleteMultipleObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除多个对象
func (c *Client) DeleteMultipleObject(request *DeleteMultipleObjectRequest) (response *DeleteMultipleObjectResponse, err error) {
	if request == nil {
		request = NewDeleteMultipleObjectRequest()
	}
	response = NewDeleteMultipleObjectResponse()
	err = c.Send(request, response)
	return
}

func NewGetBackupSettingRequest() (request *GetBackupSettingRequest) {
	request = &GetBackupSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBackupSetting")
	return
}

func NewGetBackupSettingResponse() (response *GetBackupSettingResponse) {
	response = &GetBackupSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储桶备份设置
func (c *Client) GetBackupSetting(request *GetBackupSettingRequest) (response *GetBackupSettingResponse, err error) {
	if request == nil {
		request = NewGetBackupSettingRequest()
	}
	response = NewGetBackupSettingResponse()
	err = c.Send(request, response)
	return
}

func NewOpenCosBillingRequest() (request *OpenCosBillingRequest) {
	request = &OpenCosBillingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "OpenCosBilling")
	return
}

func NewOpenCosBillingResponse() (response *OpenCosBillingResponse) {
	response = &OpenCosBillingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开通COS计费
func (c *Client) OpenCosBilling(request *OpenCosBillingRequest) (response *OpenCosBillingResponse, err error) {
	if request == nil {
		request = NewOpenCosBillingRequest()
	}
	response = NewOpenCosBillingResponse()
	err = c.Send(request, response)
	return
}

func NewTriggerBackupRequest() (request *TriggerBackupRequest) {
	request = &TriggerBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "TriggerBackup")
	return
}

func NewTriggerBackupResponse() (response *TriggerBackupResponse) {
	response = &TriggerBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 触发存储桶备份历史数据
func (c *Client) TriggerBackup(request *TriggerBackupRequest) (response *TriggerBackupResponse, err error) {
	if request == nil {
		request = NewTriggerBackupRequest()
	}
	response = NewTriggerBackupResponse()
	err = c.Send(request, response)
	return
}

func NewCosProxyRequest() (request *CosProxyRequest) {
	request = &CosProxyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "CosProxy")
	return
}

func NewCosProxyResponse() (response *CosProxyResponse) {
	response = &CosProxyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 代理 COS API 的接口
func (c *Client) CosProxy(request *CosProxyRequest) (response *CosProxyResponse, err error) {
	if request == nil {
		request = NewCosProxyRequest()
	}
	response = NewCosProxyResponse()
	err = c.Send(request, response)
	return
}

func NewGetRequestByIdRequest() (request *GetRequestByIdRequest) {
	request = &GetRequestByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRequestById")
	return
}

func NewGetRequestByIdResponse() (response *GetRequestByIdResponse) {
	response = &GetRequestByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过ID获取CSP长任务详情
func (c *Client) GetRequestById(request *GetRequestByIdRequest) (response *GetRequestByIdResponse, err error) {
	if request == nil {
		request = NewGetRequestByIdRequest()
	}
	response = NewGetRequestByIdResponse()
	err = c.Send(request, response)
	return
}

func NewGetStatHttpDaysRequest() (request *GetStatHttpDaysRequest) {
	request = &GetStatHttpDaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetStatHttpDays")
	return
}

func NewGetStatHttpDaysResponse() (response *GetStatHttpDaysResponse) {
	response = &GetStatHttpDaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取多日单存储桶的http请求统计信息
func (c *Client) GetStatHttpDays(request *GetStatHttpDaysRequest) (response *GetStatHttpDaysResponse, err error) {
	if request == nil {
		request = NewGetStatHttpDaysRequest()
	}
	response = NewGetStatHttpDaysResponse()
	err = c.Send(request, response)
	return
}

func NewGetWebsiteDomainPrefixRequest() (request *GetWebsiteDomainPrefixRequest) {
	request = &GetWebsiteDomainPrefixRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetWebsiteDomainPrefix")
	return
}

func NewGetWebsiteDomainPrefixResponse() (response *GetWebsiteDomainPrefixResponse) {
	response = &GetWebsiteDomainPrefixResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取静态网站域名前缀
func (c *Client) GetWebsiteDomainPrefix(request *GetWebsiteDomainPrefixRequest) (response *GetWebsiteDomainPrefixResponse, err error) {
	if request == nil {
		request = NewGetWebsiteDomainPrefixRequest()
	}
	response = NewGetWebsiteDomainPrefixResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserStatRequest() (request *GetUserStatRequest) {
	request = &GetUserStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetUserStat")
	return
}

func NewGetUserStatResponse() (response *GetUserStatResponse) {
	response = &GetUserStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户状态
func (c *Client) GetUserStat(request *GetUserStatRequest) (response *GetUserStatResponse, err error) {
	if request == nil {
		request = NewGetUserStatRequest()
	}
	response = NewGetUserStatResponse()
	err = c.Send(request, response)
	return
}
