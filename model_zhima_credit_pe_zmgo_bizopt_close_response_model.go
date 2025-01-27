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

// checks if the ZhimaCreditPeZmgoBizoptCloseResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaCreditPeZmgoBizoptCloseResponseModel{}

// ZhimaCreditPeZmgoBizoptCloseResponseModel struct for ZhimaCreditPeZmgoBizoptCloseResponseModel
type ZhimaCreditPeZmgoBizoptCloseResponseModel struct {
	// 芝麻GO签约申请时生成的签约申请单据号
	BizOptNo *string `json:"biz_opt_no,omitempty"`
	// 蚂蚁统一会员ID
	OpenId *string `json:"open_id,omitempty"`
	// 商户本次操作的请求流水号
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 商户ID
	PartnerId *string `json:"partner_id,omitempty"`
	// 蚂蚁统一会员ID
	UserId *string `json:"user_id,omitempty"`
}

// NewZhimaCreditPeZmgoBizoptCloseResponseModel instantiates a new ZhimaCreditPeZmgoBizoptCloseResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaCreditPeZmgoBizoptCloseResponseModel() *ZhimaCreditPeZmgoBizoptCloseResponseModel {
	this := ZhimaCreditPeZmgoBizoptCloseResponseModel{}
	return &this
}

// NewZhimaCreditPeZmgoBizoptCloseResponseModelWithDefaults instantiates a new ZhimaCreditPeZmgoBizoptCloseResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaCreditPeZmgoBizoptCloseResponseModelWithDefaults() *ZhimaCreditPeZmgoBizoptCloseResponseModel {
	this := ZhimaCreditPeZmgoBizoptCloseResponseModel{}
	return &this
}

// GetBizOptNo returns the BizOptNo field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) GetBizOptNo() string {
	if o == nil || IsNil(o.BizOptNo) {
		var ret string
		return ret
	}
	return *o.BizOptNo
}

// GetBizOptNoOk returns a tuple with the BizOptNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) GetBizOptNoOk() (*string, bool) {
	if o == nil || IsNil(o.BizOptNo) {
		return nil, false
	}
	return o.BizOptNo, true
}

// HasBizOptNo returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) HasBizOptNo() bool {
	if o != nil && !IsNil(o.BizOptNo) {
		return true
	}

	return false
}

// SetBizOptNo gets a reference to the given string and assigns it to the BizOptNo field.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) SetBizOptNo(v string) {
	o.BizOptNo = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ZhimaCreditPeZmgoBizoptCloseResponseModel) SetUserId(v string) {
	o.UserId = &v
}

func (o ZhimaCreditPeZmgoBizoptCloseResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaCreditPeZmgoBizoptCloseResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizOptNo) {
		toSerialize["biz_opt_no"] = o.BizOptNo
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partner_id"] = o.PartnerId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableZhimaCreditPeZmgoBizoptCloseResponseModel struct {
	value *ZhimaCreditPeZmgoBizoptCloseResponseModel
	isSet bool
}

func (v NullableZhimaCreditPeZmgoBizoptCloseResponseModel) Get() *ZhimaCreditPeZmgoBizoptCloseResponseModel {
	return v.value
}

func (v *NullableZhimaCreditPeZmgoBizoptCloseResponseModel) Set(val *ZhimaCreditPeZmgoBizoptCloseResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCreditPeZmgoBizoptCloseResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCreditPeZmgoBizoptCloseResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCreditPeZmgoBizoptCloseResponseModel(val *ZhimaCreditPeZmgoBizoptCloseResponseModel) *NullableZhimaCreditPeZmgoBizoptCloseResponseModel {
	return &NullableZhimaCreditPeZmgoBizoptCloseResponseModel{value: val, isSet: true}
}

func (v NullableZhimaCreditPeZmgoBizoptCloseResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCreditPeZmgoBizoptCloseResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
