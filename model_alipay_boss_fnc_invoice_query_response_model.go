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

// checks if the AlipayBossFncInvoiceQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayBossFncInvoiceQueryResponseModel{}

// AlipayBossFncInvoiceQueryResponseModel struct for AlipayBossFncInvoiceQueryResponseModel
type AlipayBossFncInvoiceQueryResponseModel struct {
	ResultSet *ArInvoiceOpenApiResponse `json:"result_set,omitempty"`
}

// NewAlipayBossFncInvoiceQueryResponseModel instantiates a new AlipayBossFncInvoiceQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayBossFncInvoiceQueryResponseModel() *AlipayBossFncInvoiceQueryResponseModel {
	this := AlipayBossFncInvoiceQueryResponseModel{}
	return &this
}

// NewAlipayBossFncInvoiceQueryResponseModelWithDefaults instantiates a new AlipayBossFncInvoiceQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayBossFncInvoiceQueryResponseModelWithDefaults() *AlipayBossFncInvoiceQueryResponseModel {
	this := AlipayBossFncInvoiceQueryResponseModel{}
	return &this
}

// GetResultSet returns the ResultSet field value if set, zero value otherwise.
func (o *AlipayBossFncInvoiceQueryResponseModel) GetResultSet() ArInvoiceOpenApiResponse {
	if o == nil || IsNil(o.ResultSet) {
		var ret ArInvoiceOpenApiResponse
		return ret
	}
	return *o.ResultSet
}

// GetResultSetOk returns a tuple with the ResultSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncInvoiceQueryResponseModel) GetResultSetOk() (*ArInvoiceOpenApiResponse, bool) {
	if o == nil || IsNil(o.ResultSet) {
		return nil, false
	}
	return o.ResultSet, true
}

// HasResultSet returns a boolean if a field has been set.
func (o *AlipayBossFncInvoiceQueryResponseModel) HasResultSet() bool {
	if o != nil && !IsNil(o.ResultSet) {
		return true
	}

	return false
}

// SetResultSet gets a reference to the given ArInvoiceOpenApiResponse and assigns it to the ResultSet field.
func (o *AlipayBossFncInvoiceQueryResponseModel) SetResultSet(v ArInvoiceOpenApiResponse) {
	o.ResultSet = &v
}

func (o AlipayBossFncInvoiceQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayBossFncInvoiceQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultSet) {
		toSerialize["result_set"] = o.ResultSet
	}
	return toSerialize, nil
}

type NullableAlipayBossFncInvoiceQueryResponseModel struct {
	value *AlipayBossFncInvoiceQueryResponseModel
	isSet bool
}

func (v NullableAlipayBossFncInvoiceQueryResponseModel) Get() *AlipayBossFncInvoiceQueryResponseModel {
	return v.value
}

func (v *NullableAlipayBossFncInvoiceQueryResponseModel) Set(val *AlipayBossFncInvoiceQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayBossFncInvoiceQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayBossFncInvoiceQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayBossFncInvoiceQueryResponseModel(val *AlipayBossFncInvoiceQueryResponseModel) *NullableAlipayBossFncInvoiceQueryResponseModel {
	return &NullableAlipayBossFncInvoiceQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayBossFncInvoiceQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayBossFncInvoiceQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
