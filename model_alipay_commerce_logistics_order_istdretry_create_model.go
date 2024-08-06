/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AlipayCommerceLogisticsOrderIstdretryCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceLogisticsOrderIstdretryCreateModel{}

// AlipayCommerceLogisticsOrderIstdretryCreateModel struct for AlipayCommerceLogisticsOrderIstdretryCreateModel
type AlipayCommerceLogisticsOrderIstdretryCreateModel struct {
	// 消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户uid；如果consumer_source是wx，则consumer_id可以为空
	ConsumerId *string `json:"consumer_id,omitempty"`
	ConsumerNotify *ConsumerNotifyIstd `json:"consumer_notify,omitempty"`
	// 消费者来源,  支付宝：alipay, 微信：wx
	ConsumerSource *string `json:"consumer_source,omitempty"`
	// 商品明细
	GoodsDetails []GoodsDetailIstd `json:"goods_details,omitempty"`
	GoodsInfo *GoodsInfoIstd `json:"goods_info,omitempty"`
	// 即时配送公司编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// 下即时配送订单token, 配送公司可以返回此字段，当商家下单时候带上这个字段，保证在一段时间内运费不变
	LogisticsToken *string `json:"logistics_token,omitempty"`
	// 消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户openId；如果consumer_source是wx，则consumer_id可以为空
	OpenId *string `json:"open_id,omitempty"`
	OrderExtIstd *OrderExtIstd `json:"order_ext_istd,omitempty"`
	// 商家订单号
	OutOrderNo *string `json:"out_order_no,omitempty"`
	Receiver *ReceiverIstd `json:"receiver,omitempty"`
	Sender *SenderIstd `json:"sender,omitempty"`
	// 商家门店编号
	ShopNo *string `json:"shop_no,omitempty"`
}

// NewAlipayCommerceLogisticsOrderIstdretryCreateModel instantiates a new AlipayCommerceLogisticsOrderIstdretryCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceLogisticsOrderIstdretryCreateModel() *AlipayCommerceLogisticsOrderIstdretryCreateModel {
	this := AlipayCommerceLogisticsOrderIstdretryCreateModel{}
	return &this
}

// NewAlipayCommerceLogisticsOrderIstdretryCreateModelWithDefaults instantiates a new AlipayCommerceLogisticsOrderIstdretryCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceLogisticsOrderIstdretryCreateModelWithDefaults() *AlipayCommerceLogisticsOrderIstdretryCreateModel {
	this := AlipayCommerceLogisticsOrderIstdretryCreateModel{}
	return &this
}

// GetConsumerId returns the ConsumerId field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetConsumerId() string {
	if o == nil || IsNil(o.ConsumerId) {
		var ret string
		return ret
	}
	return *o.ConsumerId
}

// GetConsumerIdOk returns a tuple with the ConsumerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetConsumerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerId) {
		return nil, false
	}
	return o.ConsumerId, true
}

// HasConsumerId returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasConsumerId() bool {
	if o != nil && !IsNil(o.ConsumerId) {
		return true
	}

	return false
}

// SetConsumerId gets a reference to the given string and assigns it to the ConsumerId field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetConsumerId(v string) {
	o.ConsumerId = &v
}

// GetConsumerNotify returns the ConsumerNotify field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetConsumerNotify() ConsumerNotifyIstd {
	if o == nil || IsNil(o.ConsumerNotify) {
		var ret ConsumerNotifyIstd
		return ret
	}
	return *o.ConsumerNotify
}

// GetConsumerNotifyOk returns a tuple with the ConsumerNotify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetConsumerNotifyOk() (*ConsumerNotifyIstd, bool) {
	if o == nil || IsNil(o.ConsumerNotify) {
		return nil, false
	}
	return o.ConsumerNotify, true
}

// HasConsumerNotify returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasConsumerNotify() bool {
	if o != nil && !IsNil(o.ConsumerNotify) {
		return true
	}

	return false
}

// SetConsumerNotify gets a reference to the given ConsumerNotifyIstd and assigns it to the ConsumerNotify field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetConsumerNotify(v ConsumerNotifyIstd) {
	o.ConsumerNotify = &v
}

// GetConsumerSource returns the ConsumerSource field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetConsumerSource() string {
	if o == nil || IsNil(o.ConsumerSource) {
		var ret string
		return ret
	}
	return *o.ConsumerSource
}

// GetConsumerSourceOk returns a tuple with the ConsumerSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetConsumerSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerSource) {
		return nil, false
	}
	return o.ConsumerSource, true
}

// HasConsumerSource returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasConsumerSource() bool {
	if o != nil && !IsNil(o.ConsumerSource) {
		return true
	}

	return false
}

// SetConsumerSource gets a reference to the given string and assigns it to the ConsumerSource field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetConsumerSource(v string) {
	o.ConsumerSource = &v
}

// GetGoodsDetails returns the GoodsDetails field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetGoodsDetails() []GoodsDetailIstd {
	if o == nil || IsNil(o.GoodsDetails) {
		var ret []GoodsDetailIstd
		return ret
	}
	return o.GoodsDetails
}

// GetGoodsDetailsOk returns a tuple with the GoodsDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetGoodsDetailsOk() ([]GoodsDetailIstd, bool) {
	if o == nil || IsNil(o.GoodsDetails) {
		return nil, false
	}
	return o.GoodsDetails, true
}

// HasGoodsDetails returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasGoodsDetails() bool {
	if o != nil && !IsNil(o.GoodsDetails) {
		return true
	}

	return false
}

// SetGoodsDetails gets a reference to the given []GoodsDetailIstd and assigns it to the GoodsDetails field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetGoodsDetails(v []GoodsDetailIstd) {
	o.GoodsDetails = v
}

// GetGoodsInfo returns the GoodsInfo field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetGoodsInfo() GoodsInfoIstd {
	if o == nil || IsNil(o.GoodsInfo) {
		var ret GoodsInfoIstd
		return ret
	}
	return *o.GoodsInfo
}

// GetGoodsInfoOk returns a tuple with the GoodsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetGoodsInfoOk() (*GoodsInfoIstd, bool) {
	if o == nil || IsNil(o.GoodsInfo) {
		return nil, false
	}
	return o.GoodsInfo, true
}

// HasGoodsInfo returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasGoodsInfo() bool {
	if o != nil && !IsNil(o.GoodsInfo) {
		return true
	}

	return false
}

// SetGoodsInfo gets a reference to the given GoodsInfoIstd and assigns it to the GoodsInfo field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetGoodsInfo(v GoodsInfoIstd) {
	o.GoodsInfo = &v
}

// GetLogisticsCode returns the LogisticsCode field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetLogisticsCode() string {
	if o == nil || IsNil(o.LogisticsCode) {
		var ret string
		return ret
	}
	return *o.LogisticsCode
}

// GetLogisticsCodeOk returns a tuple with the LogisticsCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetLogisticsCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LogisticsCode) {
		return nil, false
	}
	return o.LogisticsCode, true
}

// HasLogisticsCode returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasLogisticsCode() bool {
	if o != nil && !IsNil(o.LogisticsCode) {
		return true
	}

	return false
}

// SetLogisticsCode gets a reference to the given string and assigns it to the LogisticsCode field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetLogisticsCode(v string) {
	o.LogisticsCode = &v
}

// GetLogisticsToken returns the LogisticsToken field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetLogisticsToken() string {
	if o == nil || IsNil(o.LogisticsToken) {
		var ret string
		return ret
	}
	return *o.LogisticsToken
}

// GetLogisticsTokenOk returns a tuple with the LogisticsToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetLogisticsTokenOk() (*string, bool) {
	if o == nil || IsNil(o.LogisticsToken) {
		return nil, false
	}
	return o.LogisticsToken, true
}

// HasLogisticsToken returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasLogisticsToken() bool {
	if o != nil && !IsNil(o.LogisticsToken) {
		return true
	}

	return false
}

// SetLogisticsToken gets a reference to the given string and assigns it to the LogisticsToken field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetLogisticsToken(v string) {
	o.LogisticsToken = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOrderExtIstd returns the OrderExtIstd field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetOrderExtIstd() OrderExtIstd {
	if o == nil || IsNil(o.OrderExtIstd) {
		var ret OrderExtIstd
		return ret
	}
	return *o.OrderExtIstd
}

// GetOrderExtIstdOk returns a tuple with the OrderExtIstd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetOrderExtIstdOk() (*OrderExtIstd, bool) {
	if o == nil || IsNil(o.OrderExtIstd) {
		return nil, false
	}
	return o.OrderExtIstd, true
}

// HasOrderExtIstd returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasOrderExtIstd() bool {
	if o != nil && !IsNil(o.OrderExtIstd) {
		return true
	}

	return false
}

// SetOrderExtIstd gets a reference to the given OrderExtIstd and assigns it to the OrderExtIstd field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetOrderExtIstd(v OrderExtIstd) {
	o.OrderExtIstd = &v
}

// GetOutOrderNo returns the OutOrderNo field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetOutOrderNo() string {
	if o == nil || IsNil(o.OutOrderNo) {
		var ret string
		return ret
	}
	return *o.OutOrderNo
}

// GetOutOrderNoOk returns a tuple with the OutOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetOutOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutOrderNo) {
		return nil, false
	}
	return o.OutOrderNo, true
}

// HasOutOrderNo returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasOutOrderNo() bool {
	if o != nil && !IsNil(o.OutOrderNo) {
		return true
	}

	return false
}

// SetOutOrderNo gets a reference to the given string and assigns it to the OutOrderNo field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetOutOrderNo(v string) {
	o.OutOrderNo = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetReceiver() ReceiverIstd {
	if o == nil || IsNil(o.Receiver) {
		var ret ReceiverIstd
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetReceiverOk() (*ReceiverIstd, bool) {
	if o == nil || IsNil(o.Receiver) {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasReceiver() bool {
	if o != nil && !IsNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given ReceiverIstd and assigns it to the Receiver field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetReceiver(v ReceiverIstd) {
	o.Receiver = &v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetSender() SenderIstd {
	if o == nil || IsNil(o.Sender) {
		var ret SenderIstd
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetSenderOk() (*SenderIstd, bool) {
	if o == nil || IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasSender() bool {
	if o != nil && !IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given SenderIstd and assigns it to the Sender field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetSender(v SenderIstd) {
	o.Sender = &v
}

// GetShopNo returns the ShopNo field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetShopNo() string {
	if o == nil || IsNil(o.ShopNo) {
		var ret string
		return ret
	}
	return *o.ShopNo
}

// GetShopNoOk returns a tuple with the ShopNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) GetShopNoOk() (*string, bool) {
	if o == nil || IsNil(o.ShopNo) {
		return nil, false
	}
	return o.ShopNo, true
}

// HasShopNo returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) HasShopNo() bool {
	if o != nil && !IsNil(o.ShopNo) {
		return true
	}

	return false
}

// SetShopNo gets a reference to the given string and assigns it to the ShopNo field.
func (o *AlipayCommerceLogisticsOrderIstdretryCreateModel) SetShopNo(v string) {
	o.ShopNo = &v
}

func (o AlipayCommerceLogisticsOrderIstdretryCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceLogisticsOrderIstdretryCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumerId) {
		toSerialize["consumer_id"] = o.ConsumerId
	}
	if !IsNil(o.ConsumerNotify) {
		toSerialize["consumer_notify"] = o.ConsumerNotify
	}
	if !IsNil(o.ConsumerSource) {
		toSerialize["consumer_source"] = o.ConsumerSource
	}
	if !IsNil(o.GoodsDetails) {
		toSerialize["goods_details"] = o.GoodsDetails
	}
	if !IsNil(o.GoodsInfo) {
		toSerialize["goods_info"] = o.GoodsInfo
	}
	if !IsNil(o.LogisticsCode) {
		toSerialize["logistics_code"] = o.LogisticsCode
	}
	if !IsNil(o.LogisticsToken) {
		toSerialize["logistics_token"] = o.LogisticsToken
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OrderExtIstd) {
		toSerialize["order_ext_istd"] = o.OrderExtIstd
	}
	if !IsNil(o.OutOrderNo) {
		toSerialize["out_order_no"] = o.OutOrderNo
	}
	if !IsNil(o.Receiver) {
		toSerialize["receiver"] = o.Receiver
	}
	if !IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !IsNil(o.ShopNo) {
		toSerialize["shop_no"] = o.ShopNo
	}
	return toSerialize, nil
}

type NullableAlipayCommerceLogisticsOrderIstdretryCreateModel struct {
	value *AlipayCommerceLogisticsOrderIstdretryCreateModel
	isSet bool
}

func (v NullableAlipayCommerceLogisticsOrderIstdretryCreateModel) Get() *AlipayCommerceLogisticsOrderIstdretryCreateModel {
	return v.value
}

func (v *NullableAlipayCommerceLogisticsOrderIstdretryCreateModel) Set(val *AlipayCommerceLogisticsOrderIstdretryCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceLogisticsOrderIstdretryCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceLogisticsOrderIstdretryCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceLogisticsOrderIstdretryCreateModel(val *AlipayCommerceLogisticsOrderIstdretryCreateModel) *NullableAlipayCommerceLogisticsOrderIstdretryCreateModel {
	return &NullableAlipayCommerceLogisticsOrderIstdretryCreateModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceLogisticsOrderIstdretryCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceLogisticsOrderIstdretryCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

