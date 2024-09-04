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
package main

import (
	"fmt"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/errors"
	"github.com/websitefox/tcecloud-sdk-go/tcecloud/common/profile"
	location "github.com/websitefox/tcecloud-sdk-go/tcecloud/location/v20191128"
	"os"
)

func main() {
	// 必要步骤：
	// 实例化一个认证对象，入参需要传入账户密钥对secretId，secretKey。
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
	// 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	// 以免泄露密钥对危及你的财产安全。
	credential := common.NewCredential(
		os.Getenv("SECRET_ID"),
		os.Getenv("SECRET_KEY"),
	)

	// 非必要步骤
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// SDK默认使用HTTPS请求。
	cpf.HttpProfile.Scheme = "HTTP"
	// SDK默认使用POST方法。
	// 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
	//cpf.HttpProfile.ReqMethod = "GET"
	// SDK有默认的超时时间，非必要请不要进行调整。
	// 如有需要请在代码中查阅以获取最新的默认值。
	cpf.HttpProfile.ReqTimeout = 10
	// SDK会自动指定域名。通常是不需要特地指定域名的。
	// location是产品，api3是调用版本， cloud.sunhongs.com是主域名。
	cpf.HttpProfile.Endpoint = "location.api3.cloud.sunhongs.com"
	// SDK默认用TC3-HMAC-SHA256进行签名
	// 非必要请不要修改这个字段。
	// cpf.SignMethod = "TC3-HMAC-SHA256"

	// 实例化要请求产品(以location为例)的client对象
	// 第二个参数是地域信息
	client, _ := location.NewClient(credential, "", cpf)
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定DescribeRegionZoneRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进阅行开发，可以方便的跳转查各个接口和数据结构的文档说明。
	request := location.NewDescribeRegionZoneRequest()

	// 此接口指定产品查询
	request.ProductId = common.StringPtr("cvm")

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.DescribeRegionZone(request)
	// 处理异常
	if _, ok := err.(*errors.CloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())
}
