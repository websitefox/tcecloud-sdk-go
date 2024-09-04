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

package v20200907

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-09-07"

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

func NewSearchSyncConfigRequest() (request *SearchSyncConfigRequest) {
	request = &SearchSyncConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "SearchSyncConfig")
	return
}

func NewSearchSyncConfigResponse() (response *SearchSyncConfigResponse) {
	response = &SearchSyncConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询同步配置
func (c *Client) SearchSyncConfig(request *SearchSyncConfigRequest) (response *SearchSyncConfigResponse, err error) {
	if request == nil {
		request = NewSearchSyncConfigRequest()
	}
	response = NewSearchSyncConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSearchYumCatalogRequest() (request *SearchYumCatalogRequest) {
	request = &SearchYumCatalogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "SearchYumCatalog")
	return
}

func NewSearchYumCatalogResponse() (response *SearchYumCatalogResponse) {
	response = &SearchYumCatalogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件目录
func (c *Client) SearchYumCatalog(request *SearchYumCatalogRequest) (response *SearchYumCatalogResponse, err error) {
	if request == nil {
		request = NewSearchYumCatalogRequest()
	}
	response = NewSearchYumCatalogResponse()
	err = c.Send(request, response)
	return
}
