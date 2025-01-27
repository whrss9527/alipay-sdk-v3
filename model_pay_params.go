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

// checks if the PayParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayParams{}

// PayParams struct for PayParams
type PayParams struct {
	// 普通异步支付, 传入该参数时，如果满足受理条件，会先同步受理支付，然后在异步调度推进支付  NORMAL_ASYNC: 普通异步，受理成功之后，会在交易关单之前通过一定的策略重试  NEAR_REAL_TIME_ASYNC: 准实时异步，受理成功之后，会准实时发起1次调度
	AsyncType *string `json:"async_type,omitempty"`
	// 重试类型，当async_type传入NORMAL_ASYNC时，可以设置该参数，选择是否要重试，retry_type 可选，不设置时，可重试。 ● NONE_AND_CLOSETRADE：不重试，支付请求只会被执行1次，执行完成后如果交易未成功，会关闭交易 ● NONE：不重试，支付请求只会被执行1次，执行完成后，不做任何处理。交易到达了timeout_express指定的时间后，关闭交易。 ● RETY: 重试，支付请求在超时关单前，会按照策略重试
	RetryType *string `json:"retry_type,omitempty"`
}

// NewPayParams instantiates a new PayParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayParams() *PayParams {
	this := PayParams{}
	return &this
}

// NewPayParamsWithDefaults instantiates a new PayParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayParamsWithDefaults() *PayParams {
	this := PayParams{}
	return &this
}

// GetAsyncType returns the AsyncType field value if set, zero value otherwise.
func (o *PayParams) GetAsyncType() string {
	if o == nil || IsNil(o.AsyncType) {
		var ret string
		return ret
	}
	return *o.AsyncType
}

// GetAsyncTypeOk returns a tuple with the AsyncType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayParams) GetAsyncTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AsyncType) {
		return nil, false
	}
	return o.AsyncType, true
}

// HasAsyncType returns a boolean if a field has been set.
func (o *PayParams) HasAsyncType() bool {
	if o != nil && !IsNil(o.AsyncType) {
		return true
	}

	return false
}

// SetAsyncType gets a reference to the given string and assigns it to the AsyncType field.
func (o *PayParams) SetAsyncType(v string) {
	o.AsyncType = &v
}

// GetRetryType returns the RetryType field value if set, zero value otherwise.
func (o *PayParams) GetRetryType() string {
	if o == nil || IsNil(o.RetryType) {
		var ret string
		return ret
	}
	return *o.RetryType
}

// GetRetryTypeOk returns a tuple with the RetryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayParams) GetRetryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RetryType) {
		return nil, false
	}
	return o.RetryType, true
}

// HasRetryType returns a boolean if a field has been set.
func (o *PayParams) HasRetryType() bool {
	if o != nil && !IsNil(o.RetryType) {
		return true
	}

	return false
}

// SetRetryType gets a reference to the given string and assigns it to the RetryType field.
func (o *PayParams) SetRetryType(v string) {
	o.RetryType = &v
}

func (o PayParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AsyncType) {
		toSerialize["async_type"] = o.AsyncType
	}
	if !IsNil(o.RetryType) {
		toSerialize["retry_type"] = o.RetryType
	}
	return toSerialize, nil
}

type NullablePayParams struct {
	value *PayParams
	isSet bool
}

func (v NullablePayParams) Get() *PayParams {
	return v.value
}

func (v *NullablePayParams) Set(val *PayParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePayParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePayParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayParams(val *PayParams) *NullablePayParams {
	return &NullablePayParams{value: val, isSet: true}
}

func (v NullablePayParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
