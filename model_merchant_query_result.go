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

// checks if the MerchantQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantQueryResult{}

// MerchantQueryResult struct for MerchantQueryResult
type MerchantQueryResult struct {
	// 商户的别名，用户商户对客展示的名称
	AliasName *string `json:"alias_name,omitempty"`
	// 商户认证证件号，企业营业执照号
	CertNo *string `json:"cert_no,omitempty"`
	// 市名称
	City *string `json:"city,omitempty"`
	// 商户经营详细地址
	DetailAddress *string `json:"detail_address,omitempty"`
	// 区县名称
	Distinct *string `json:"distinct,omitempty"`
	// 商户新版mcc code
	MccCode *string `json:"mcc_code,omitempty"`
	// 描述商户类型，个人-P/企业-B
	MerchantType *string `json:"merchant_type,omitempty"`
	// 商户认证名称信息
	Name *string `json:"name,omitempty"`
	// 省名称
	Province *string `json:"province,omitempty"`
}

// NewMerchantQueryResult instantiates a new MerchantQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantQueryResult() *MerchantQueryResult {
	this := MerchantQueryResult{}
	return &this
}

// NewMerchantQueryResultWithDefaults instantiates a new MerchantQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantQueryResultWithDefaults() *MerchantQueryResult {
	this := MerchantQueryResult{}
	return &this
}

// GetAliasName returns the AliasName field value if set, zero value otherwise.
func (o *MerchantQueryResult) GetAliasName() string {
	if o == nil || IsNil(o.AliasName) {
		var ret string
		return ret
	}
	return *o.AliasName
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantQueryResult) GetAliasNameOk() (*string, bool) {
	if o == nil || IsNil(o.AliasName) {
		return nil, false
	}
	return o.AliasName, true
}

// HasAliasName returns a boolean if a field has been set.
func (o *MerchantQueryResult) HasAliasName() bool {
	if o != nil && !IsNil(o.AliasName) {
		return true
	}

	return false
}

// SetAliasName gets a reference to the given string and assigns it to the AliasName field.
func (o *MerchantQueryResult) SetAliasName(v string) {
	o.AliasName = &v
}

// GetCertNo returns the CertNo field value if set, zero value otherwise.
func (o *MerchantQueryResult) GetCertNo() string {
	if o == nil || IsNil(o.CertNo) {
		var ret string
		return ret
	}
	return *o.CertNo
}

// GetCertNoOk returns a tuple with the CertNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantQueryResult) GetCertNoOk() (*string, bool) {
	if o == nil || IsNil(o.CertNo) {
		return nil, false
	}
	return o.CertNo, true
}

// HasCertNo returns a boolean if a field has been set.
func (o *MerchantQueryResult) HasCertNo() bool {
	if o != nil && !IsNil(o.CertNo) {
		return true
	}

	return false
}

// SetCertNo gets a reference to the given string and assigns it to the CertNo field.
func (o *MerchantQueryResult) SetCertNo(v string) {
	o.CertNo = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *MerchantQueryResult) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantQueryResult) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *MerchantQueryResult) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *MerchantQueryResult) SetCity(v string) {
	o.City = &v
}

// GetDetailAddress returns the DetailAddress field value if set, zero value otherwise.
func (o *MerchantQueryResult) GetDetailAddress() string {
	if o == nil || IsNil(o.DetailAddress) {
		var ret string
		return ret
	}
	return *o.DetailAddress
}

// GetDetailAddressOk returns a tuple with the DetailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantQueryResult) GetDetailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.DetailAddress) {
		return nil, false
	}
	return o.DetailAddress, true
}

// HasDetailAddress returns a boolean if a field has been set.
func (o *MerchantQueryResult) HasDetailAddress() bool {
	if o != nil && !IsNil(o.DetailAddress) {
		return true
	}

	return false
}

// SetDetailAddress gets a reference to the given string and assigns it to the DetailAddress field.
func (o *MerchantQueryResult) SetDetailAddress(v string) {
	o.DetailAddress = &v
}

// GetDistinct returns the Distinct field value if set, zero value otherwise.
func (o *MerchantQueryResult) GetDistinct() string {
	if o == nil || IsNil(o.Distinct) {
		var ret string
		return ret
	}
	return *o.Distinct
}

// GetDistinctOk returns a tuple with the Distinct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantQueryResult) GetDistinctOk() (*string, bool) {
	if o == nil || IsNil(o.Distinct) {
		return nil, false
	}
	return o.Distinct, true
}

// HasDistinct returns a boolean if a field has been set.
func (o *MerchantQueryResult) HasDistinct() bool {
	if o != nil && !IsNil(o.Distinct) {
		return true
	}

	return false
}

// SetDistinct gets a reference to the given string and assigns it to the Distinct field.
func (o *MerchantQueryResult) SetDistinct(v string) {
	o.Distinct = &v
}

// GetMccCode returns the MccCode field value if set, zero value otherwise.
func (o *MerchantQueryResult) GetMccCode() string {
	if o == nil || IsNil(o.MccCode) {
		var ret string
		return ret
	}
	return *o.MccCode
}

// GetMccCodeOk returns a tuple with the MccCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantQueryResult) GetMccCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MccCode) {
		return nil, false
	}
	return o.MccCode, true
}

// HasMccCode returns a boolean if a field has been set.
func (o *MerchantQueryResult) HasMccCode() bool {
	if o != nil && !IsNil(o.MccCode) {
		return true
	}

	return false
}

// SetMccCode gets a reference to the given string and assigns it to the MccCode field.
func (o *MerchantQueryResult) SetMccCode(v string) {
	o.MccCode = &v
}

// GetMerchantType returns the MerchantType field value if set, zero value otherwise.
func (o *MerchantQueryResult) GetMerchantType() string {
	if o == nil || IsNil(o.MerchantType) {
		var ret string
		return ret
	}
	return *o.MerchantType
}

// GetMerchantTypeOk returns a tuple with the MerchantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantQueryResult) GetMerchantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantType) {
		return nil, false
	}
	return o.MerchantType, true
}

// HasMerchantType returns a boolean if a field has been set.
func (o *MerchantQueryResult) HasMerchantType() bool {
	if o != nil && !IsNil(o.MerchantType) {
		return true
	}

	return false
}

// SetMerchantType gets a reference to the given string and assigns it to the MerchantType field.
func (o *MerchantQueryResult) SetMerchantType(v string) {
	o.MerchantType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MerchantQueryResult) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantQueryResult) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MerchantQueryResult) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MerchantQueryResult) SetName(v string) {
	o.Name = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *MerchantQueryResult) GetProvince() string {
	if o == nil || IsNil(o.Province) {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantQueryResult) GetProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.Province) {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *MerchantQueryResult) HasProvince() bool {
	if o != nil && !IsNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *MerchantQueryResult) SetProvince(v string) {
	o.Province = &v
}

func (o MerchantQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantQueryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AliasName) {
		toSerialize["alias_name"] = o.AliasName
	}
	if !IsNil(o.CertNo) {
		toSerialize["cert_no"] = o.CertNo
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.DetailAddress) {
		toSerialize["detail_address"] = o.DetailAddress
	}
	if !IsNil(o.Distinct) {
		toSerialize["distinct"] = o.Distinct
	}
	if !IsNil(o.MccCode) {
		toSerialize["mcc_code"] = o.MccCode
	}
	if !IsNil(o.MerchantType) {
		toSerialize["merchant_type"] = o.MerchantType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Province) {
		toSerialize["province"] = o.Province
	}
	return toSerialize, nil
}

type NullableMerchantQueryResult struct {
	value *MerchantQueryResult
	isSet bool
}

func (v NullableMerchantQueryResult) Get() *MerchantQueryResult {
	return v.value
}

func (v *NullableMerchantQueryResult) Set(val *MerchantQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantQueryResult(val *MerchantQueryResult) *NullableMerchantQueryResult {
	return &NullableMerchantQueryResult{value: val, isSet: true}
}

func (v NullableMerchantQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

