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

// checks if the SignerBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignerBean{}

// SignerBean struct for SignerBean
type SignerBean struct {
	// 个人邮箱
	Email *string `json:"email,omitempty"`
	// 个人证件号
	IdNumber *string `json:"id_number,omitempty"`
	// 个人证件类型（证件号不为空时必填，详见个人证件类型说明 ）
	IdType *string `json:"id_type,omitempty"`
	// 个人手机号
	Mobile *string `json:"mobile,omitempty"`
	// 个人姓名
	Name *string `json:"name,omitempty"`
	Org  *Org    `json:"org,omitempty"`
	// 个人唯一标识：可传入平台的个人用户id、支付宝userid、证件号、手机号、邮箱等。
	ThirdPartyUserId *string `json:"third_party_user_id,omitempty"`
}

// NewSignerBean instantiates a new SignerBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignerBean() *SignerBean {
	this := SignerBean{}
	return &this
}

// NewSignerBeanWithDefaults instantiates a new SignerBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignerBeanWithDefaults() *SignerBean {
	this := SignerBean{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SignerBean) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignerBean) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SignerBean) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SignerBean) SetEmail(v string) {
	o.Email = &v
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise.
func (o *SignerBean) GetIdNumber() string {
	if o == nil || IsNil(o.IdNumber) {
		var ret string
		return ret
	}
	return *o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignerBean) GetIdNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IdNumber) {
		return nil, false
	}
	return o.IdNumber, true
}

// HasIdNumber returns a boolean if a field has been set.
func (o *SignerBean) HasIdNumber() bool {
	if o != nil && !IsNil(o.IdNumber) {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given string and assigns it to the IdNumber field.
func (o *SignerBean) SetIdNumber(v string) {
	o.IdNumber = &v
}

// GetIdType returns the IdType field value if set, zero value otherwise.
func (o *SignerBean) GetIdType() string {
	if o == nil || IsNil(o.IdType) {
		var ret string
		return ret
	}
	return *o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignerBean) GetIdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdType) {
		return nil, false
	}
	return o.IdType, true
}

// HasIdType returns a boolean if a field has been set.
func (o *SignerBean) HasIdType() bool {
	if o != nil && !IsNil(o.IdType) {
		return true
	}

	return false
}

// SetIdType gets a reference to the given string and assigns it to the IdType field.
func (o *SignerBean) SetIdType(v string) {
	o.IdType = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *SignerBean) GetMobile() string {
	if o == nil || IsNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignerBean) GetMobileOk() (*string, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *SignerBean) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *SignerBean) SetMobile(v string) {
	o.Mobile = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SignerBean) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignerBean) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SignerBean) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SignerBean) SetName(v string) {
	o.Name = &v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *SignerBean) GetOrg() Org {
	if o == nil || IsNil(o.Org) {
		var ret Org
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignerBean) GetOrgOk() (*Org, bool) {
	if o == nil || IsNil(o.Org) {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *SignerBean) HasOrg() bool {
	if o != nil && !IsNil(o.Org) {
		return true
	}

	return false
}

// SetOrg gets a reference to the given Org and assigns it to the Org field.
func (o *SignerBean) SetOrg(v Org) {
	o.Org = &v
}

// GetThirdPartyUserId returns the ThirdPartyUserId field value if set, zero value otherwise.
func (o *SignerBean) GetThirdPartyUserId() string {
	if o == nil || IsNil(o.ThirdPartyUserId) {
		var ret string
		return ret
	}
	return *o.ThirdPartyUserId
}

// GetThirdPartyUserIdOk returns a tuple with the ThirdPartyUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignerBean) GetThirdPartyUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThirdPartyUserId) {
		return nil, false
	}
	return o.ThirdPartyUserId, true
}

// HasThirdPartyUserId returns a boolean if a field has been set.
func (o *SignerBean) HasThirdPartyUserId() bool {
	if o != nil && !IsNil(o.ThirdPartyUserId) {
		return true
	}

	return false
}

// SetThirdPartyUserId gets a reference to the given string and assigns it to the ThirdPartyUserId field.
func (o *SignerBean) SetThirdPartyUserId(v string) {
	o.ThirdPartyUserId = &v
}

func (o SignerBean) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignerBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.IdNumber) {
		toSerialize["id_number"] = o.IdNumber
	}
	if !IsNil(o.IdType) {
		toSerialize["id_type"] = o.IdType
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Org) {
		toSerialize["org"] = o.Org
	}
	if !IsNil(o.ThirdPartyUserId) {
		toSerialize["third_party_user_id"] = o.ThirdPartyUserId
	}
	return toSerialize, nil
}

type NullableSignerBean struct {
	value *SignerBean
	isSet bool
}

func (v NullableSignerBean) Get() *SignerBean {
	return v.value
}

func (v *NullableSignerBean) Set(val *SignerBean) {
	v.value = val
	v.isSet = true
}

func (v NullableSignerBean) IsSet() bool {
	return v.isSet
}

func (v *NullableSignerBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignerBean(val *SignerBean) *NullableSignerBean {
	return &NullableSignerBean{value: val, isSet: true}
}

func (v NullableSignerBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignerBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
