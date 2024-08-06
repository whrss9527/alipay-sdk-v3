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

// checks if the AntMerchantExpandShopReceiptaccountSaveModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntMerchantExpandShopReceiptaccountSaveModel{}

// AntMerchantExpandShopReceiptaccountSaveModel struct for AntMerchantExpandShopReceiptaccountSaveModel
type AntMerchantExpandShopReceiptaccountSaveModel struct {
	// 是否承诺收单账号信息准确。具体承诺信息可查看 <a href=\"https://gw.alipayobjects.com/os/bmw-prod/922bafa8-a712-4f79-aa32-6f08d7359a5c.docx\">门店信息承诺函</a>。 Y 是 N 否
	Promise *string `json:"promise,omitempty"`
	// 收单账号
	ReceiptAccountId *string `json:"receipt_account_id,omitempty"`
	// 店铺ID
	ShopId *string `json:"shop_id,omitempty"`
}

// NewAntMerchantExpandShopReceiptaccountSaveModel instantiates a new AntMerchantExpandShopReceiptaccountSaveModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntMerchantExpandShopReceiptaccountSaveModel() *AntMerchantExpandShopReceiptaccountSaveModel {
	this := AntMerchantExpandShopReceiptaccountSaveModel{}
	return &this
}

// NewAntMerchantExpandShopReceiptaccountSaveModelWithDefaults instantiates a new AntMerchantExpandShopReceiptaccountSaveModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntMerchantExpandShopReceiptaccountSaveModelWithDefaults() *AntMerchantExpandShopReceiptaccountSaveModel {
	this := AntMerchantExpandShopReceiptaccountSaveModel{}
	return &this
}

// GetPromise returns the Promise field value if set, zero value otherwise.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) GetPromise() string {
	if o == nil || IsNil(o.Promise) {
		var ret string
		return ret
	}
	return *o.Promise
}

// GetPromiseOk returns a tuple with the Promise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) GetPromiseOk() (*string, bool) {
	if o == nil || IsNil(o.Promise) {
		return nil, false
	}
	return o.Promise, true
}

// HasPromise returns a boolean if a field has been set.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) HasPromise() bool {
	if o != nil && !IsNil(o.Promise) {
		return true
	}

	return false
}

// SetPromise gets a reference to the given string and assigns it to the Promise field.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) SetPromise(v string) {
	o.Promise = &v
}

// GetReceiptAccountId returns the ReceiptAccountId field value if set, zero value otherwise.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) GetReceiptAccountId() string {
	if o == nil || IsNil(o.ReceiptAccountId) {
		var ret string
		return ret
	}
	return *o.ReceiptAccountId
}

// GetReceiptAccountIdOk returns a tuple with the ReceiptAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) GetReceiptAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptAccountId) {
		return nil, false
	}
	return o.ReceiptAccountId, true
}

// HasReceiptAccountId returns a boolean if a field has been set.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) HasReceiptAccountId() bool {
	if o != nil && !IsNil(o.ReceiptAccountId) {
		return true
	}

	return false
}

// SetReceiptAccountId gets a reference to the given string and assigns it to the ReceiptAccountId field.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) SetReceiptAccountId(v string) {
	o.ReceiptAccountId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) GetShopId() string {
	if o == nil || IsNil(o.ShopId) {
		var ret string
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) GetShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given string and assigns it to the ShopId field.
func (o *AntMerchantExpandShopReceiptaccountSaveModel) SetShopId(v string) {
	o.ShopId = &v
}

func (o AntMerchantExpandShopReceiptaccountSaveModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntMerchantExpandShopReceiptaccountSaveModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Promise) {
		toSerialize["promise"] = o.Promise
	}
	if !IsNil(o.ReceiptAccountId) {
		toSerialize["receipt_account_id"] = o.ReceiptAccountId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	return toSerialize, nil
}

type NullableAntMerchantExpandShopReceiptaccountSaveModel struct {
	value *AntMerchantExpandShopReceiptaccountSaveModel
	isSet bool
}

func (v NullableAntMerchantExpandShopReceiptaccountSaveModel) Get() *AntMerchantExpandShopReceiptaccountSaveModel {
	return v.value
}

func (v *NullableAntMerchantExpandShopReceiptaccountSaveModel) Set(val *AntMerchantExpandShopReceiptaccountSaveModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandShopReceiptaccountSaveModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandShopReceiptaccountSaveModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandShopReceiptaccountSaveModel(val *AntMerchantExpandShopReceiptaccountSaveModel) *NullableAntMerchantExpandShopReceiptaccountSaveModel {
	return &NullableAntMerchantExpandShopReceiptaccountSaveModel{value: val, isSet: true}
}

func (v NullableAntMerchantExpandShopReceiptaccountSaveModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandShopReceiptaccountSaveModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

