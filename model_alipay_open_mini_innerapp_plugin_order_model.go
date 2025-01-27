/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
)

// checks if the AlipayOpenMiniInnerappPluginOrderModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerappPluginOrderModel{}

// AlipayOpenMiniInnerappPluginOrderModel struct for AlipayOpenMiniInnerappPluginOrderModel
type AlipayOpenMiniInnerappPluginOrderModel struct {
	// 业务来源
	AppOrigin *string `json:"app_origin,omitempty"`
	// 订购的服务商品ID
	MerchandiseId *string `json:"merchandise_id,omitempty"`
	// 一二方支持传入appId
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 业务幂等号，规则为商品id+appOrigin+requestId 作为幂等流水号 非必填,不传则默认生成一个8位requestId
	RequestId *string `json:"request_id,omitempty"`
}

// NewAlipayOpenMiniInnerappPluginOrderModel instantiates a new AlipayOpenMiniInnerappPluginOrderModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerappPluginOrderModel() *AlipayOpenMiniInnerappPluginOrderModel {
	this := AlipayOpenMiniInnerappPluginOrderModel{}
	return &this
}

// NewAlipayOpenMiniInnerappPluginOrderModelWithDefaults instantiates a new AlipayOpenMiniInnerappPluginOrderModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerappPluginOrderModelWithDefaults() *AlipayOpenMiniInnerappPluginOrderModel {
	this := AlipayOpenMiniInnerappPluginOrderModel{}
	return &this
}

// GetAppOrigin returns the AppOrigin field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginOrderModel) GetAppOrigin() string {
	if o == nil || IsNil(o.AppOrigin) {
		var ret string
		return ret
	}
	return *o.AppOrigin
}

// GetAppOriginOk returns a tuple with the AppOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginOrderModel) GetAppOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AppOrigin) {
		return nil, false
	}
	return o.AppOrigin, true
}

// HasAppOrigin returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginOrderModel) HasAppOrigin() bool {
	if o != nil && !IsNil(o.AppOrigin) {
		return true
	}

	return false
}

// SetAppOrigin gets a reference to the given string and assigns it to the AppOrigin field.
func (o *AlipayOpenMiniInnerappPluginOrderModel) SetAppOrigin(v string) {
	o.AppOrigin = &v
}

// GetMerchandiseId returns the MerchandiseId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginOrderModel) GetMerchandiseId() string {
	if o == nil || IsNil(o.MerchandiseId) {
		var ret string
		return ret
	}
	return *o.MerchandiseId
}

// GetMerchandiseIdOk returns a tuple with the MerchandiseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginOrderModel) GetMerchandiseIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchandiseId) {
		return nil, false
	}
	return o.MerchandiseId, true
}

// HasMerchandiseId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginOrderModel) HasMerchandiseId() bool {
	if o != nil && !IsNil(o.MerchandiseId) {
		return true
	}

	return false
}

// SetMerchandiseId gets a reference to the given string and assigns it to the MerchandiseId field.
func (o *AlipayOpenMiniInnerappPluginOrderModel) SetMerchandiseId(v string) {
	o.MerchandiseId = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginOrderModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginOrderModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginOrderModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniInnerappPluginOrderModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginOrderModel) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginOrderModel) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginOrderModel) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *AlipayOpenMiniInnerappPluginOrderModel) SetRequestId(v string) {
	o.RequestId = &v
}

func (o AlipayOpenMiniInnerappPluginOrderModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerappPluginOrderModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppOrigin) {
		toSerialize["app_origin"] = o.AppOrigin
	}
	if !IsNil(o.MerchandiseId) {
		toSerialize["merchandise_id"] = o.MerchandiseId
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerappPluginOrderModel struct {
	value *AlipayOpenMiniInnerappPluginOrderModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerappPluginOrderModel) Get() *AlipayOpenMiniInnerappPluginOrderModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerappPluginOrderModel) Set(val *AlipayOpenMiniInnerappPluginOrderModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerappPluginOrderModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerappPluginOrderModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerappPluginOrderModel(val *AlipayOpenMiniInnerappPluginOrderModel) *NullableAlipayOpenMiniInnerappPluginOrderModel {
	return &NullableAlipayOpenMiniInnerappPluginOrderModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerappPluginOrderModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerappPluginOrderModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
