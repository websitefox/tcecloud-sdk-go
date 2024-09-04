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

package v20200726

import (
	"encoding/json"

	tchttp "github.com/websitefox/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DUserMsgSites struct {

	// 站内信id

	MessageId *int64 `json:"MessageId,omitempty" name:"MessageId"`
	// uin

	ReceiverId *int64 `json:"ReceiverId,omitempty" name:"ReceiverId"`
	// ownerUIn

	DeveloperId *int64 `json:"DeveloperId,omitempty" name:"DeveloperId"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 是否已读

	IsRead *int64 `json:"IsRead,omitempty" name:"IsRead"`
	// 是否删除

	IsDeleted *uint64 `json:"IsDeleted,omitempty" name:"IsDeleted"`
	// 一级消息分类id

	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 二级消息分类id

	SecondCategoryId *int64 `json:"SecondCategoryId,omitempty" name:"SecondCategoryId"`
	// 发送人

	SenderName *string `json:"SenderName,omitempty" name:"SenderName"`
	// 发送角色id

	SenderRoleId *int64 `json:"SenderRoleId,omitempty" name:"SenderRoleId"`
	// 消息类型id

	MessageType *int64 `json:"MessageType,omitempty" name:"MessageType"`
	// 发送方式

	SendWay *uint64 `json:"SendWay,omitempty" name:"SendWay"`
	// 是否重要

	IsImportant *int64 `json:"IsImportant,omitempty" name:"IsImportant"`
	// 重要原因

	ImportantReason *string `json:"ImportantReason,omitempty" name:"ImportantReason"`
	// 来源

	FromSource *string `json:"FromSource,omitempty" name:"FromSource"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 发送时间

	SendTime *string `json:"SendTime,omitempty" name:"SendTime"`
	// 是否是模版

	IsTemplate *int64 `json:"IsTemplate,omitempty" name:"IsTemplate"`
	// 下一条站内信id

	NextId *int64 `json:"NextId,omitempty" name:"NextId"`
	// 上一条站内信id

	PrevId *int64 `json:"PrevId,omitempty" name:"PrevId"`
}

type AddReceiverOnMsgTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddReceiverOnMsgTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddReceiverOnMsgTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMsgSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 站内信汇总信息

		Data *SiteSummary `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMsgSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMsgSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSiteMsgReadRequest struct {
	*tchttp.BaseRequest

	// 站内信id

	MessageId []*int64 `json:"MessageId,omitempty" name:"MessageId"`
}

func (r *SetSiteMsgReadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSiteMsgReadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResGetMsgList struct {

	// 消息id

	MsgId *int64 `json:"MsgId,omitempty" name:"MsgId"`
	// 一级消息分类名

	Category *string `json:"Category,omitempty" name:"Category"`
	// 一级消息分类id

	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 是否重要

	IsImportant *int64 `json:"IsImportant,omitempty" name:"IsImportant"`
	// 是否已读

	IsRead *int64 `json:"IsRead,omitempty" name:"IsRead"`
	// 发送时间

	SendTime *string `json:"SendTime,omitempty" name:"SendTime"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
}

type Users struct {

	// uid

	UID *uint64 `json:"UID,omitempty" name:"UID"`
	// 是否主账号

	IsOwner *uint64 `json:"IsOwner,omitempty" name:"IsOwner"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type GetMsgListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 未读消息

		Data []*ResGetMsgList `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMsgListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMsgListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSiteMsgDetailRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
	// 站内信id

	MessageId *int64 `json:"MessageId,omitempty" name:"MessageId"`
}

func (r *GetSiteMsgDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSiteMsgDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterString struct {

	// key过滤字段

	Name *string `json:"Name,omitempty" name:"Name"`
	// value过滤字段的值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 多选过滤字段

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DelReceiverOnMsgTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelReceiverOnMsgTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelReceiverOnMsgTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSiteMsgDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 站内信

		Data *DUserMsgSites `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSiteMsgDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSiteMsgDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFirstMsgTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *ListFirstMsgTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListFirstMsgTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SiteSummary struct {

	// 总计

	TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
	// 总计未读

	UnRead *int64 `json:"UnRead,omitempty" name:"UnRead"`
	// 重要类型的

	Important *ReadAndUnRead `json:"Important,omitempty" name:"Important"`
	// 非重要类型的

	Other *ReadAndUnRead `json:"Other,omitempty" name:"Other"`
}

type SearchSiteMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 站内信

		Data []*DUserMsgSites `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSiteMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSiteMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResMessageTypeList struct {

	// 二级

	SecondType *uint64 `json:"SecondType,omitempty" name:"SecondType"`
	// 一级

	FirstType *uint64 `json:"FirstType,omitempty" name:"FirstType"`
	// 类型名

	Name *Languagefield `json:"Name,omitempty" name:"Name"`
	// 描述

	Desc *Languagefield `json:"Desc,omitempty" name:"Desc"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 发送通道

	Channels *int64 `json:"Channels,omitempty" name:"Channels"`
	// 默认通道

	DefaultChannels *int64 `json:"DefaultChannels,omitempty" name:"DefaultChannels"`
	// 默认策略

	DefaultStrategy *string `json:"DefaultStrategy,omitempty" name:"DefaultStrategy"`
	// 是否订阅

	IsDisplay *uint64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 订阅权重

	DisplayWeight *uint64 `json:"DisplayWeight,omitempty" name:"DisplayWeight"`
	// 调用类型

	InvokePattern *uint64 `json:"InvokePattern,omitempty" name:"InvokePattern"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 站点

	Station *int64 `json:"Station,omitempty" name:"Station"`
}

type ListFirstMsgTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一级消息分类详情

		Data []*FirstType `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListFirstMsgTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListFirstMsgTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSecondaryMsgTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *ListSecondaryMsgTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSecondaryMsgTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAllSiteMsgReadRequest struct {
	*tchttp.BaseRequest
}

func (r *SetAllSiteMsgReadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAllSiteMsgReadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddReceiverOnMsgTypeRequest struct {
	*tchttp.BaseRequest

	// 订阅类型

	Types []*uint64 `json:"Types,omitempty" name:"Types"`
	// 接收人uin

	SubIDs []*uint64 `json:"SubIDs,omitempty" name:"SubIDs"`
}

func (r *AddReceiverOnMsgTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddReceiverOnMsgTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReceiversOnAllTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetReceiversOnAllTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReceiversOnAllTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSiteMsgRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤查询

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *SearchSiteMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSiteMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelReceiverOnMsgTypeRequest struct {
	*tchttp.BaseRequest

	// 子账号subid

	SubIDs []*uint64 `json:"SubIDs,omitempty" name:"SubIDs"`
	// 二级消息分类

	Types []*uint64 `json:"Types,omitempty" name:"Types"`
}

func (r *DelReceiverOnMsgTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelReceiverOnMsgTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReceiversOnAllTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消息订阅信息

		Data []*MsgTypeSubscribe `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetReceiversOnAllTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReceiversOnAllTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReceiverOnMsgTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReceiverOnMsgTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReceiverOnMsgTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSiteMsgRequest struct {
	*tchttp.BaseRequest

	// 站内信id

	MessageId []*int64 `json:"MessageId,omitempty" name:"MessageId"`
}

func (r *DeleteSiteMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSiteMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMsgListRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *GetMsgListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMsgListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSecondaryMsgTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 二级分类

		Data []*ResMessageTypeList `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSecondaryMsgTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSecondaryMsgTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAllSiteMsgReadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAllSiteMsgReadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAllSiteMsgReadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgTypeSubscribe struct {

	// 二级分类id

	SecondType *uint64 `json:"SecondType,omitempty" name:"SecondType"`
	// 一级分类id

	FirstType *uint64 `json:"FirstType,omitempty" name:"FirstType"`
	// 类型名国际化

	Name *Languagefield `json:"Name,omitempty" name:"Name"`
	// 描述国际化

	Desc *Languagefield `json:"Desc,omitempty" name:"Desc"`
	// 发送通道

	Channels *int64 `json:"Channels,omitempty" name:"Channels"`
	// 通道

	ChannelsSet *int64 `json:"ChannelsSet,omitempty" name:"ChannelsSet"`
	// 默认通道

	DefaultChannels *int64 `json:"DefaultChannels,omitempty" name:"DefaultChannels"`
	// 订阅权重

	DisplayWeight *uint64 `json:"DisplayWeight,omitempty" name:"DisplayWeight"`
	// 用户信息

	Users []*Users `json:"Users,omitempty" name:"Users"`
}

type ReadAndUnRead struct {

	// 已读数量

	Read *int64 `json:"Read,omitempty" name:"Read"`
	// 未读数量

	UnRead *int64 `json:"UnRead,omitempty" name:"UnRead"`
}

type GetMsgSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMsgSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMsgSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FirstType struct {

	// 一级分类id

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 分类权重

	DisplayWeight *uint64 `json:"DisplayWeight,omitempty" name:"DisplayWeight"`
	// 分类名

	CategoryName *Languagefield `json:"CategoryName,omitempty" name:"CategoryName"`
}

type Languagefield struct {

	// 英文

	En *string `json:"En,omitempty" name:"En"`
	// 中文

	Zh *string `json:"Zh,omitempty" name:"Zh"`
}

type DeleteSiteMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除失败的站内信id

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSiteMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSiteMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReceiverOnMsgTypeRequest struct {
	*tchttp.BaseRequest

	// 订阅类型id

	MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`
	// uin

	SubIDs []*uint64 `json:"SubIDs,omitempty" name:"SubIDs"`
	// 发送通道

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
}

func (r *ModifyReceiverOnMsgTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReceiverOnMsgTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSiteMsgReadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 设置已读失败的站内信

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetSiteMsgReadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSiteMsgReadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
