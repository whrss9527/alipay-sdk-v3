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

// checks if the AlipayOpenPublicMessageContentCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicMessageContentCreateResponseModel{}

// AlipayOpenPublicMessageContentCreateResponseModel struct for AlipayOpenPublicMessageContentCreateResponseModel
type AlipayOpenPublicMessageContentCreateResponseModel struct {
	// 内容id
	ContentId *string `json:"content_id,omitempty"`
	// 内容详情内链url
	ContentUrl *string `json:"content_url,omitempty"`
}

// NewAlipayOpenPublicMessageContentCreateResponseModel instantiates a new AlipayOpenPublicMessageContentCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicMessageContentCreateResponseModel() *AlipayOpenPublicMessageContentCreateResponseModel {
	this := AlipayOpenPublicMessageContentCreateResponseModel{}
	return &this
}

// NewAlipayOpenPublicMessageContentCreateResponseModelWithDefaults instantiates a new AlipayOpenPublicMessageContentCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicMessageContentCreateResponseModelWithDefaults() *AlipayOpenPublicMessageContentCreateResponseModel {
	this := AlipayOpenPublicMessageContentCreateResponseModel{}
	return &this
}

// GetContentId returns the ContentId field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentCreateResponseModel) GetContentId() string {
	if o == nil || IsNil(o.ContentId) {
		var ret string
		return ret
	}
	return *o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentCreateResponseModel) GetContentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContentId) {
		return nil, false
	}
	return o.ContentId, true
}

// HasContentId returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentCreateResponseModel) HasContentId() bool {
	if o != nil && !IsNil(o.ContentId) {
		return true
	}

	return false
}

// SetContentId gets a reference to the given string and assigns it to the ContentId field.
func (o *AlipayOpenPublicMessageContentCreateResponseModel) SetContentId(v string) {
	o.ContentId = &v
}

// GetContentUrl returns the ContentUrl field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentCreateResponseModel) GetContentUrl() string {
	if o == nil || IsNil(o.ContentUrl) {
		var ret string
		return ret
	}
	return *o.ContentUrl
}

// GetContentUrlOk returns a tuple with the ContentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentCreateResponseModel) GetContentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ContentUrl) {
		return nil, false
	}
	return o.ContentUrl, true
}

// HasContentUrl returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentCreateResponseModel) HasContentUrl() bool {
	if o != nil && !IsNil(o.ContentUrl) {
		return true
	}

	return false
}

// SetContentUrl gets a reference to the given string and assigns it to the ContentUrl field.
func (o *AlipayOpenPublicMessageContentCreateResponseModel) SetContentUrl(v string) {
	o.ContentUrl = &v
}

func (o AlipayOpenPublicMessageContentCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicMessageContentCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentId) {
		toSerialize["content_id"] = o.ContentId
	}
	if !IsNil(o.ContentUrl) {
		toSerialize["content_url"] = o.ContentUrl
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicMessageContentCreateResponseModel struct {
	value *AlipayOpenPublicMessageContentCreateResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicMessageContentCreateResponseModel) Get() *AlipayOpenPublicMessageContentCreateResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicMessageContentCreateResponseModel) Set(val *AlipayOpenPublicMessageContentCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMessageContentCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMessageContentCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMessageContentCreateResponseModel(val *AlipayOpenPublicMessageContentCreateResponseModel) *NullableAlipayOpenPublicMessageContentCreateResponseModel {
	return &NullableAlipayOpenPublicMessageContentCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMessageContentCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMessageContentCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
