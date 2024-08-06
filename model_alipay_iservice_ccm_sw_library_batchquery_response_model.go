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

// checks if the AlipayIserviceCcmSwLibraryBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmSwLibraryBatchqueryResponseModel{}

// AlipayIserviceCcmSwLibraryBatchqueryResponseModel struct for AlipayIserviceCcmSwLibraryBatchqueryResponseModel
type AlipayIserviceCcmSwLibraryBatchqueryResponseModel struct {
	// 知识库集合
	Libraries []LibraryInfo `json:"libraries,omitempty"`
}

// NewAlipayIserviceCcmSwLibraryBatchqueryResponseModel instantiates a new AlipayIserviceCcmSwLibraryBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmSwLibraryBatchqueryResponseModel() *AlipayIserviceCcmSwLibraryBatchqueryResponseModel {
	this := AlipayIserviceCcmSwLibraryBatchqueryResponseModel{}
	return &this
}

// NewAlipayIserviceCcmSwLibraryBatchqueryResponseModelWithDefaults instantiates a new AlipayIserviceCcmSwLibraryBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmSwLibraryBatchqueryResponseModelWithDefaults() *AlipayIserviceCcmSwLibraryBatchqueryResponseModel {
	this := AlipayIserviceCcmSwLibraryBatchqueryResponseModel{}
	return &this
}

// GetLibraries returns the Libraries field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwLibraryBatchqueryResponseModel) GetLibraries() []LibraryInfo {
	if o == nil || IsNil(o.Libraries) {
		var ret []LibraryInfo
		return ret
	}
	return o.Libraries
}

// GetLibrariesOk returns a tuple with the Libraries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwLibraryBatchqueryResponseModel) GetLibrariesOk() ([]LibraryInfo, bool) {
	if o == nil || IsNil(o.Libraries) {
		return nil, false
	}
	return o.Libraries, true
}

// HasLibraries returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwLibraryBatchqueryResponseModel) HasLibraries() bool {
	if o != nil && !IsNil(o.Libraries) {
		return true
	}

	return false
}

// SetLibraries gets a reference to the given []LibraryInfo and assigns it to the Libraries field.
func (o *AlipayIserviceCcmSwLibraryBatchqueryResponseModel) SetLibraries(v []LibraryInfo) {
	o.Libraries = v
}

func (o AlipayIserviceCcmSwLibraryBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmSwLibraryBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Libraries) {
		toSerialize["libraries"] = o.Libraries
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmSwLibraryBatchqueryResponseModel struct {
	value *AlipayIserviceCcmSwLibraryBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmSwLibraryBatchqueryResponseModel) Get() *AlipayIserviceCcmSwLibraryBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwLibraryBatchqueryResponseModel) Set(val *AlipayIserviceCcmSwLibraryBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwLibraryBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwLibraryBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwLibraryBatchqueryResponseModel(val *AlipayIserviceCcmSwLibraryBatchqueryResponseModel) *NullableAlipayIserviceCcmSwLibraryBatchqueryResponseModel {
	return &NullableAlipayIserviceCcmSwLibraryBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwLibraryBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwLibraryBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

