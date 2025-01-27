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

// checks if the AlipayBossFncUserinvoiceinfoQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayBossFncUserinvoiceinfoQueryResponseModel{}

// AlipayBossFncUserinvoiceinfoQueryResponseModel struct for AlipayBossFncUserinvoiceinfoQueryResponseModel
type AlipayBossFncUserinvoiceinfoQueryResponseModel struct {
	ResultSet *UserInvoiceInfoOpenApiResponse `json:"result_set,omitempty"`
}

// NewAlipayBossFncUserinvoiceinfoQueryResponseModel instantiates a new AlipayBossFncUserinvoiceinfoQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayBossFncUserinvoiceinfoQueryResponseModel() *AlipayBossFncUserinvoiceinfoQueryResponseModel {
	this := AlipayBossFncUserinvoiceinfoQueryResponseModel{}
	return &this
}

// NewAlipayBossFncUserinvoiceinfoQueryResponseModelWithDefaults instantiates a new AlipayBossFncUserinvoiceinfoQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayBossFncUserinvoiceinfoQueryResponseModelWithDefaults() *AlipayBossFncUserinvoiceinfoQueryResponseModel {
	this := AlipayBossFncUserinvoiceinfoQueryResponseModel{}
	return &this
}

// GetResultSet returns the ResultSet field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoQueryResponseModel) GetResultSet() UserInvoiceInfoOpenApiResponse {
	if o == nil || IsNil(o.ResultSet) {
		var ret UserInvoiceInfoOpenApiResponse
		return ret
	}
	return *o.ResultSet
}

// GetResultSetOk returns a tuple with the ResultSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoQueryResponseModel) GetResultSetOk() (*UserInvoiceInfoOpenApiResponse, bool) {
	if o == nil || IsNil(o.ResultSet) {
		return nil, false
	}
	return o.ResultSet, true
}

// HasResultSet returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoQueryResponseModel) HasResultSet() bool {
	if o != nil && !IsNil(o.ResultSet) {
		return true
	}

	return false
}

// SetResultSet gets a reference to the given UserInvoiceInfoOpenApiResponse and assigns it to the ResultSet field.
func (o *AlipayBossFncUserinvoiceinfoQueryResponseModel) SetResultSet(v UserInvoiceInfoOpenApiResponse) {
	o.ResultSet = &v
}

func (o AlipayBossFncUserinvoiceinfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayBossFncUserinvoiceinfoQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultSet) {
		toSerialize["result_set"] = o.ResultSet
	}
	return toSerialize, nil
}

type NullableAlipayBossFncUserinvoiceinfoQueryResponseModel struct {
	value *AlipayBossFncUserinvoiceinfoQueryResponseModel
	isSet bool
}

func (v NullableAlipayBossFncUserinvoiceinfoQueryResponseModel) Get() *AlipayBossFncUserinvoiceinfoQueryResponseModel {
	return v.value
}

func (v *NullableAlipayBossFncUserinvoiceinfoQueryResponseModel) Set(val *AlipayBossFncUserinvoiceinfoQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayBossFncUserinvoiceinfoQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayBossFncUserinvoiceinfoQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayBossFncUserinvoiceinfoQueryResponseModel(val *AlipayBossFncUserinvoiceinfoQueryResponseModel) *NullableAlipayBossFncUserinvoiceinfoQueryResponseModel {
	return &NullableAlipayBossFncUserinvoiceinfoQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayBossFncUserinvoiceinfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayBossFncUserinvoiceinfoQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
