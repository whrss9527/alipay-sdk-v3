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

// checks if the RecruitMiniApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecruitMiniApp{}

// RecruitMiniApp struct for RecruitMiniApp
type RecruitMiniApp struct {
	// 小程序ID
	MiniAppId *string `json:"mini_app_id,omitempty"`
}

// NewRecruitMiniApp instantiates a new RecruitMiniApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecruitMiniApp() *RecruitMiniApp {
	this := RecruitMiniApp{}
	return &this
}

// NewRecruitMiniAppWithDefaults instantiates a new RecruitMiniApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecruitMiniAppWithDefaults() *RecruitMiniApp {
	this := RecruitMiniApp{}
	return &this
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *RecruitMiniApp) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitMiniApp) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *RecruitMiniApp) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *RecruitMiniApp) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

func (o RecruitMiniApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecruitMiniApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	return toSerialize, nil
}

type NullableRecruitMiniApp struct {
	value *RecruitMiniApp
	isSet bool
}

func (v NullableRecruitMiniApp) Get() *RecruitMiniApp {
	return v.value
}

func (v *NullableRecruitMiniApp) Set(val *RecruitMiniApp) {
	v.value = val
	v.isSet = true
}

func (v NullableRecruitMiniApp) IsSet() bool {
	return v.isSet
}

func (v *NullableRecruitMiniApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecruitMiniApp(val *RecruitMiniApp) *NullableRecruitMiniApp {
	return &NullableRecruitMiniApp{value: val, isSet: true}
}

func (v NullableRecruitMiniApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecruitMiniApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
