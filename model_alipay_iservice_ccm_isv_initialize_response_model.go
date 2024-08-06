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

// checks if the AlipayIserviceCcmIsvInitializeResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmIsvInitializeResponseModel{}

// AlipayIserviceCcmIsvInitializeResponseModel struct for AlipayIserviceCcmIsvInitializeResponseModel
type AlipayIserviceCcmIsvInitializeResponseModel struct {
	// Base64编码CCM公钥：CCM公私钥对由CCM自动生成，用于Iframe spi接口安全认证
	CcmPubKey *string `json:"ccm_pub_key,omitempty"`
}

// NewAlipayIserviceCcmIsvInitializeResponseModel instantiates a new AlipayIserviceCcmIsvInitializeResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmIsvInitializeResponseModel() *AlipayIserviceCcmIsvInitializeResponseModel {
	this := AlipayIserviceCcmIsvInitializeResponseModel{}
	return &this
}

// NewAlipayIserviceCcmIsvInitializeResponseModelWithDefaults instantiates a new AlipayIserviceCcmIsvInitializeResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmIsvInitializeResponseModelWithDefaults() *AlipayIserviceCcmIsvInitializeResponseModel {
	this := AlipayIserviceCcmIsvInitializeResponseModel{}
	return &this
}

// GetCcmPubKey returns the CcmPubKey field value if set, zero value otherwise.
func (o *AlipayIserviceCcmIsvInitializeResponseModel) GetCcmPubKey() string {
	if o == nil || IsNil(o.CcmPubKey) {
		var ret string
		return ret
	}
	return *o.CcmPubKey
}

// GetCcmPubKeyOk returns a tuple with the CcmPubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmIsvInitializeResponseModel) GetCcmPubKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CcmPubKey) {
		return nil, false
	}
	return o.CcmPubKey, true
}

// HasCcmPubKey returns a boolean if a field has been set.
func (o *AlipayIserviceCcmIsvInitializeResponseModel) HasCcmPubKey() bool {
	if o != nil && !IsNil(o.CcmPubKey) {
		return true
	}

	return false
}

// SetCcmPubKey gets a reference to the given string and assigns it to the CcmPubKey field.
func (o *AlipayIserviceCcmIsvInitializeResponseModel) SetCcmPubKey(v string) {
	o.CcmPubKey = &v
}

func (o AlipayIserviceCcmIsvInitializeResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmIsvInitializeResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CcmPubKey) {
		toSerialize["ccm_pub_key"] = o.CcmPubKey
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmIsvInitializeResponseModel struct {
	value *AlipayIserviceCcmIsvInitializeResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmIsvInitializeResponseModel) Get() *AlipayIserviceCcmIsvInitializeResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmIsvInitializeResponseModel) Set(val *AlipayIserviceCcmIsvInitializeResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmIsvInitializeResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmIsvInitializeResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmIsvInitializeResponseModel(val *AlipayIserviceCcmIsvInitializeResponseModel) *NullableAlipayIserviceCcmIsvInitializeResponseModel {
	return &NullableAlipayIserviceCcmIsvInitializeResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmIsvInitializeResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmIsvInitializeResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

