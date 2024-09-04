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

package v20200814

import (
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/tcecloud/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-08-14"

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

func NewDescribeUsersProductCatalogsRequest() (request *DescribeUsersProductCatalogsRequest) {
	request = &DescribeUsersProductCatalogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productcatalog", APIVersion, "DescribeUsersProductCatalogs")
	return
}

func NewDescribeUsersProductCatalogsResponse() (response *DescribeUsersProductCatalogsResponse) {
	response = &DescribeUsersProductCatalogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定租户的产品目录
func (c *Client) DescribeUsersProductCatalogs(request *DescribeUsersProductCatalogsRequest) (response *DescribeUsersProductCatalogsResponse, err error) {
	if request == nil {
		request = NewDescribeUsersProductCatalogsRequest()
	}
	response = NewDescribeUsersProductCatalogsResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyUserProductCatalogsRequest() (request *VerifyUserProductCatalogsRequest) {
	request = &VerifyUserProductCatalogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productcatalog", APIVersion, "VerifyUserProductCatalogs")
	return
}

func NewVerifyUserProductCatalogsResponse() (response *VerifyUserProductCatalogsResponse) {
	response = &VerifyUserProductCatalogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 鉴权（判断特定appid是否有权限访问特定产品uuid ）
func (c *Client) VerifyUserProductCatalogs(request *VerifyUserProductCatalogsRequest) (response *VerifyUserProductCatalogsResponse, err error) {
	if request == nil {
		request = NewVerifyUserProductCatalogsRequest()
	}
	response = NewVerifyUserProductCatalogsResponse()
	err = c.Send(request, response)
	return
}
