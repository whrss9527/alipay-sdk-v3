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

// checks if the OpenCertifyIdentifyInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenCertifyIdentifyInfo{}

// OpenCertifyIdentifyInfo struct for OpenCertifyIdentifyInfo
type OpenCertifyIdentifyInfo struct {
	// 填入真实姓名
	CertName *string `json:"cert_name,omitempty"`
	// 填入姓名相匹配的证件号码
	CertNo *string `json:"cert_no,omitempty"`
	// 当前仅支持IDENTITY_CARD
	CertType *string `json:"cert_type,omitempty"`
	// 当前仅支持CERT_INFO
	IdentityType *string `json:"identity_type,omitempty"`
	// 选填手机号
	PhoneNo *string `json:"phone_no,omitempty"`
}

// NewOpenCertifyIdentifyInfo instantiates a new OpenCertifyIdentifyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenCertifyIdentifyInfo() *OpenCertifyIdentifyInfo {
	this := OpenCertifyIdentifyInfo{}
	return &this
}

// NewOpenCertifyIdentifyInfoWithDefaults instantiates a new OpenCertifyIdentifyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenCertifyIdentifyInfoWithDefaults() *OpenCertifyIdentifyInfo {
	this := OpenCertifyIdentifyInfo{}
	return &this
}

// GetCertName returns the CertName field value if set, zero value otherwise.
func (o *OpenCertifyIdentifyInfo) GetCertName() string {
	if o == nil || IsNil(o.CertName) {
		var ret string
		return ret
	}
	return *o.CertName
}

// GetCertNameOk returns a tuple with the CertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyIdentifyInfo) GetCertNameOk() (*string, bool) {
	if o == nil || IsNil(o.CertName) {
		return nil, false
	}
	return o.CertName, true
}

// HasCertName returns a boolean if a field has been set.
func (o *OpenCertifyIdentifyInfo) HasCertName() bool {
	if o != nil && !IsNil(o.CertName) {
		return true
	}

	return false
}

// SetCertName gets a reference to the given string and assigns it to the CertName field.
func (o *OpenCertifyIdentifyInfo) SetCertName(v string) {
	o.CertName = &v
}

// GetCertNo returns the CertNo field value if set, zero value otherwise.
func (o *OpenCertifyIdentifyInfo) GetCertNo() string {
	if o == nil || IsNil(o.CertNo) {
		var ret string
		return ret
	}
	return *o.CertNo
}

// GetCertNoOk returns a tuple with the CertNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyIdentifyInfo) GetCertNoOk() (*string, bool) {
	if o == nil || IsNil(o.CertNo) {
		return nil, false
	}
	return o.CertNo, true
}

// HasCertNo returns a boolean if a field has been set.
func (o *OpenCertifyIdentifyInfo) HasCertNo() bool {
	if o != nil && !IsNil(o.CertNo) {
		return true
	}

	return false
}

// SetCertNo gets a reference to the given string and assigns it to the CertNo field.
func (o *OpenCertifyIdentifyInfo) SetCertNo(v string) {
	o.CertNo = &v
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *OpenCertifyIdentifyInfo) GetCertType() string {
	if o == nil || IsNil(o.CertType) {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyIdentifyInfo) GetCertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CertType) {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *OpenCertifyIdentifyInfo) HasCertType() bool {
	if o != nil && !IsNil(o.CertType) {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *OpenCertifyIdentifyInfo) SetCertType(v string) {
	o.CertType = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *OpenCertifyIdentifyInfo) GetIdentityType() string {
	if o == nil || IsNil(o.IdentityType) {
		var ret string
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyIdentifyInfo) GetIdentityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityType) {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *OpenCertifyIdentifyInfo) HasIdentityType() bool {
	if o != nil && !IsNil(o.IdentityType) {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given string and assigns it to the IdentityType field.
func (o *OpenCertifyIdentifyInfo) SetIdentityType(v string) {
	o.IdentityType = &v
}

// GetPhoneNo returns the PhoneNo field value if set, zero value otherwise.
func (o *OpenCertifyIdentifyInfo) GetPhoneNo() string {
	if o == nil || IsNil(o.PhoneNo) {
		var ret string
		return ret
	}
	return *o.PhoneNo
}

// GetPhoneNoOk returns a tuple with the PhoneNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyIdentifyInfo) GetPhoneNoOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNo) {
		return nil, false
	}
	return o.PhoneNo, true
}

// HasPhoneNo returns a boolean if a field has been set.
func (o *OpenCertifyIdentifyInfo) HasPhoneNo() bool {
	if o != nil && !IsNil(o.PhoneNo) {
		return true
	}

	return false
}

// SetPhoneNo gets a reference to the given string and assigns it to the PhoneNo field.
func (o *OpenCertifyIdentifyInfo) SetPhoneNo(v string) {
	o.PhoneNo = &v
}

func (o OpenCertifyIdentifyInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenCertifyIdentifyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertName) {
		toSerialize["cert_name"] = o.CertName
	}
	if !IsNil(o.CertNo) {
		toSerialize["cert_no"] = o.CertNo
	}
	if !IsNil(o.CertType) {
		toSerialize["cert_type"] = o.CertType
	}
	if !IsNil(o.IdentityType) {
		toSerialize["identity_type"] = o.IdentityType
	}
	if !IsNil(o.PhoneNo) {
		toSerialize["phone_no"] = o.PhoneNo
	}
	return toSerialize, nil
}

type NullableOpenCertifyIdentifyInfo struct {
	value *OpenCertifyIdentifyInfo
	isSet bool
}

func (v NullableOpenCertifyIdentifyInfo) Get() *OpenCertifyIdentifyInfo {
	return v.value
}

func (v *NullableOpenCertifyIdentifyInfo) Set(val *OpenCertifyIdentifyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenCertifyIdentifyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenCertifyIdentifyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenCertifyIdentifyInfo(val *OpenCertifyIdentifyInfo) *NullableOpenCertifyIdentifyInfo {
	return &NullableOpenCertifyIdentifyInfo{value: val, isSet: true}
}

func (v NullableOpenCertifyIdentifyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenCertifyIdentifyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
