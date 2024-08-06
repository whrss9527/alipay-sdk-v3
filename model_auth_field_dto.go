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

// checks if the AuthFieldDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthFieldDTO{}

// AuthFieldDTO struct for AuthFieldDTO
type AuthFieldDTO struct {
	// 接口英文名称
	ApiName *string `json:"api_name,omitempty"`
	// 接口字段英文名称
	FieldName *string `json:"field_name,omitempty"`
	// 接口归属的功能code
	PackageCode *string `json:"package_code,omitempty"`
	// 驳回原因
	Reason *string `json:"reason,omitempty"`
	// AUDIT 审核中，AGREE通过，REJECT驳回，INVALID无效（isv代申请场景)
	Status *string `json:"status,omitempty"`
	// 用户应用app_id
	UserAppId *string `json:"user_app_id,omitempty"`
}

// NewAuthFieldDTO instantiates a new AuthFieldDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthFieldDTO() *AuthFieldDTO {
	this := AuthFieldDTO{}
	return &this
}

// NewAuthFieldDTOWithDefaults instantiates a new AuthFieldDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthFieldDTOWithDefaults() *AuthFieldDTO {
	this := AuthFieldDTO{}
	return &this
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *AuthFieldDTO) GetApiName() string {
	if o == nil || IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldDTO) GetApiNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *AuthFieldDTO) HasApiName() bool {
	if o != nil && !IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *AuthFieldDTO) SetApiName(v string) {
	o.ApiName = &v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *AuthFieldDTO) GetFieldName() string {
	if o == nil || IsNil(o.FieldName) {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldDTO) GetFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *AuthFieldDTO) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *AuthFieldDTO) SetFieldName(v string) {
	o.FieldName = &v
}

// GetPackageCode returns the PackageCode field value if set, zero value otherwise.
func (o *AuthFieldDTO) GetPackageCode() string {
	if o == nil || IsNil(o.PackageCode) {
		var ret string
		return ret
	}
	return *o.PackageCode
}

// GetPackageCodeOk returns a tuple with the PackageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldDTO) GetPackageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PackageCode) {
		return nil, false
	}
	return o.PackageCode, true
}

// HasPackageCode returns a boolean if a field has been set.
func (o *AuthFieldDTO) HasPackageCode() bool {
	if o != nil && !IsNil(o.PackageCode) {
		return true
	}

	return false
}

// SetPackageCode gets a reference to the given string and assigns it to the PackageCode field.
func (o *AuthFieldDTO) SetPackageCode(v string) {
	o.PackageCode = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AuthFieldDTO) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldDTO) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AuthFieldDTO) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AuthFieldDTO) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthFieldDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthFieldDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthFieldDTO) SetStatus(v string) {
	o.Status = &v
}

// GetUserAppId returns the UserAppId field value if set, zero value otherwise.
func (o *AuthFieldDTO) GetUserAppId() string {
	if o == nil || IsNil(o.UserAppId) {
		var ret string
		return ret
	}
	return *o.UserAppId
}

// GetUserAppIdOk returns a tuple with the UserAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldDTO) GetUserAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserAppId) {
		return nil, false
	}
	return o.UserAppId, true
}

// HasUserAppId returns a boolean if a field has been set.
func (o *AuthFieldDTO) HasUserAppId() bool {
	if o != nil && !IsNil(o.UserAppId) {
		return true
	}

	return false
}

// SetUserAppId gets a reference to the given string and assigns it to the UserAppId field.
func (o *AuthFieldDTO) SetUserAppId(v string) {
	o.UserAppId = &v
}

func (o AuthFieldDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthFieldDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiName) {
		toSerialize["api_name"] = o.ApiName
	}
	if !IsNil(o.FieldName) {
		toSerialize["field_name"] = o.FieldName
	}
	if !IsNil(o.PackageCode) {
		toSerialize["package_code"] = o.PackageCode
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UserAppId) {
		toSerialize["user_app_id"] = o.UserAppId
	}
	return toSerialize, nil
}

type NullableAuthFieldDTO struct {
	value *AuthFieldDTO
	isSet bool
}

func (v NullableAuthFieldDTO) Get() *AuthFieldDTO {
	return v.value
}

func (v *NullableAuthFieldDTO) Set(val *AuthFieldDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthFieldDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthFieldDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthFieldDTO(val *AuthFieldDTO) *NullableAuthFieldDTO {
	return &NullableAuthFieldDTO{value: val, isSet: true}
}

func (v NullableAuthFieldDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthFieldDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

