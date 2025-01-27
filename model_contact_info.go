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

// checks if the ContactInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactInfo{}

// ContactInfo struct for ContactInfo
type ContactInfo struct {
	// 电子邮箱
	Email *string `json:"email,omitempty"`
	// 身份证号
	IdCardNo *string `json:"id_card_no,omitempty"`
	// 手机号
	Mobile *string `json:"mobile,omitempty"`
	// 联系人名字
	Name *string `json:"name,omitempty"`
	// 电话
	Phone *string `json:"phone,omitempty"`
	// 商户联系人业务标识枚举，表示商户联系人的职责
	Tag []string `json:"tag,omitempty"`
	// 联系人类型，取值范围：LEGAL_PERSON：法人；CONTROLLER：实际控制人；AGENT：代理人；OTHER：其他
	Type *string `json:"type,omitempty"`
}

// NewContactInfo instantiates a new ContactInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactInfo() *ContactInfo {
	this := ContactInfo{}
	return &this
}

// NewContactInfoWithDefaults instantiates a new ContactInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactInfoWithDefaults() *ContactInfo {
	this := ContactInfo{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ContactInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ContactInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ContactInfo) SetEmail(v string) {
	o.Email = &v
}

// GetIdCardNo returns the IdCardNo field value if set, zero value otherwise.
func (o *ContactInfo) GetIdCardNo() string {
	if o == nil || IsNil(o.IdCardNo) {
		var ret string
		return ret
	}
	return *o.IdCardNo
}

// GetIdCardNoOk returns a tuple with the IdCardNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactInfo) GetIdCardNoOk() (*string, bool) {
	if o == nil || IsNil(o.IdCardNo) {
		return nil, false
	}
	return o.IdCardNo, true
}

// HasIdCardNo returns a boolean if a field has been set.
func (o *ContactInfo) HasIdCardNo() bool {
	if o != nil && !IsNil(o.IdCardNo) {
		return true
	}

	return false
}

// SetIdCardNo gets a reference to the given string and assigns it to the IdCardNo field.
func (o *ContactInfo) SetIdCardNo(v string) {
	o.IdCardNo = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *ContactInfo) GetMobile() string {
	if o == nil || IsNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactInfo) GetMobileOk() (*string, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *ContactInfo) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *ContactInfo) SetMobile(v string) {
	o.Mobile = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContactInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContactInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContactInfo) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ContactInfo) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactInfo) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ContactInfo) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ContactInfo) SetPhone(v string) {
	o.Phone = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ContactInfo) GetTag() []string {
	if o == nil || IsNil(o.Tag) {
		var ret []string
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactInfo) GetTagOk() ([]string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ContactInfo) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given []string and assigns it to the Tag field.
func (o *ContactInfo) SetTag(v []string) {
	o.Tag = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContactInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContactInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContactInfo) SetType(v string) {
	o.Type = &v
}

func (o ContactInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.IdCardNo) {
		toSerialize["id_card_no"] = o.IdCardNo
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableContactInfo struct {
	value *ContactInfo
	isSet bool
}

func (v NullableContactInfo) Get() *ContactInfo {
	return v.value
}

func (v *NullableContactInfo) Set(val *ContactInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableContactInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableContactInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactInfo(val *ContactInfo) *NullableContactInfo {
	return &NullableContactInfo{value: val, isSet: true}
}

func (v NullableContactInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
