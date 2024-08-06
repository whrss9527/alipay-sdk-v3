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

// checks if the DeliverAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliverAddress{}

// DeliverAddress struct for DeliverAddress
type DeliverAddress struct {
	// 地址
	Address *string `json:"address,omitempty"`
	// 区域编码
	AddressCode *string `json:"address_code,omitempty"`
	// 是否默认收货地址
	DefaultDeliverAddress *string `json:"default_deliver_address,omitempty"`
	// 收货人所在区县
	DeliverArea *string `json:"deliver_area,omitempty"`
	// 收货人所在城市
	DeliverCity *string `json:"deliver_city,omitempty"`
	// 收货人全名
	DeliverFullname *string `json:"deliver_fullname,omitempty"`
	// 收货地址的联系人移动电话
	DeliverMobile *string `json:"deliver_mobile,omitempty"`
	// 收货地址的联系人固定电话
	DeliverPhone *string `json:"deliver_phone,omitempty"`
	// 收货人所在省份
	DeliverProvince *string `json:"deliver_province,omitempty"`
	// 邮政编码
	Zip *string `json:"zip,omitempty"`
}

// NewDeliverAddress instantiates a new DeliverAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliverAddress() *DeliverAddress {
	this := DeliverAddress{}
	return &this
}

// NewDeliverAddressWithDefaults instantiates a new DeliverAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliverAddressWithDefaults() *DeliverAddress {
	this := DeliverAddress{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DeliverAddress) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverAddress) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DeliverAddress) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DeliverAddress) SetAddress(v string) {
	o.Address = &v
}

// GetAddressCode returns the AddressCode field value if set, zero value otherwise.
func (o *DeliverAddress) GetAddressCode() string {
	if o == nil || IsNil(o.AddressCode) {
		var ret string
		return ret
	}
	return *o.AddressCode
}

// GetAddressCodeOk returns a tuple with the AddressCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverAddress) GetAddressCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AddressCode) {
		return nil, false
	}
	return o.AddressCode, true
}

// HasAddressCode returns a boolean if a field has been set.
func (o *DeliverAddress) HasAddressCode() bool {
	if o != nil && !IsNil(o.AddressCode) {
		return true
	}

	return false
}

// SetAddressCode gets a reference to the given string and assigns it to the AddressCode field.
func (o *DeliverAddress) SetAddressCode(v string) {
	o.AddressCode = &v
}

// GetDefaultDeliverAddress returns the DefaultDeliverAddress field value if set, zero value otherwise.
func (o *DeliverAddress) GetDefaultDeliverAddress() string {
	if o == nil || IsNil(o.DefaultDeliverAddress) {
		var ret string
		return ret
	}
	return *o.DefaultDeliverAddress
}

// GetDefaultDeliverAddressOk returns a tuple with the DefaultDeliverAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverAddress) GetDefaultDeliverAddressOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultDeliverAddress) {
		return nil, false
	}
	return o.DefaultDeliverAddress, true
}

// HasDefaultDeliverAddress returns a boolean if a field has been set.
func (o *DeliverAddress) HasDefaultDeliverAddress() bool {
	if o != nil && !IsNil(o.DefaultDeliverAddress) {
		return true
	}

	return false
}

// SetDefaultDeliverAddress gets a reference to the given string and assigns it to the DefaultDeliverAddress field.
func (o *DeliverAddress) SetDefaultDeliverAddress(v string) {
	o.DefaultDeliverAddress = &v
}

// GetDeliverArea returns the DeliverArea field value if set, zero value otherwise.
func (o *DeliverAddress) GetDeliverArea() string {
	if o == nil || IsNil(o.DeliverArea) {
		var ret string
		return ret
	}
	return *o.DeliverArea
}

// GetDeliverAreaOk returns a tuple with the DeliverArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverAddress) GetDeliverAreaOk() (*string, bool) {
	if o == nil || IsNil(o.DeliverArea) {
		return nil, false
	}
	return o.DeliverArea, true
}

// HasDeliverArea returns a boolean if a field has been set.
func (o *DeliverAddress) HasDeliverArea() bool {
	if o != nil && !IsNil(o.DeliverArea) {
		return true
	}

	return false
}

// SetDeliverArea gets a reference to the given string and assigns it to the DeliverArea field.
func (o *DeliverAddress) SetDeliverArea(v string) {
	o.DeliverArea = &v
}

// GetDeliverCity returns the DeliverCity field value if set, zero value otherwise.
func (o *DeliverAddress) GetDeliverCity() string {
	if o == nil || IsNil(o.DeliverCity) {
		var ret string
		return ret
	}
	return *o.DeliverCity
}

// GetDeliverCityOk returns a tuple with the DeliverCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverAddress) GetDeliverCityOk() (*string, bool) {
	if o == nil || IsNil(o.DeliverCity) {
		return nil, false
	}
	return o.DeliverCity, true
}

// HasDeliverCity returns a boolean if a field has been set.
func (o *DeliverAddress) HasDeliverCity() bool {
	if o != nil && !IsNil(o.DeliverCity) {
		return true
	}

	return false
}

// SetDeliverCity gets a reference to the given string and assigns it to the DeliverCity field.
func (o *DeliverAddress) SetDeliverCity(v string) {
	o.DeliverCity = &v
}

// GetDeliverFullname returns the DeliverFullname field value if set, zero value otherwise.
func (o *DeliverAddress) GetDeliverFullname() string {
	if o == nil || IsNil(o.DeliverFullname) {
		var ret string
		return ret
	}
	return *o.DeliverFullname
}

// GetDeliverFullnameOk returns a tuple with the DeliverFullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverAddress) GetDeliverFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.DeliverFullname) {
		return nil, false
	}
	return o.DeliverFullname, true
}

// HasDeliverFullname returns a boolean if a field has been set.
func (o *DeliverAddress) HasDeliverFullname() bool {
	if o != nil && !IsNil(o.DeliverFullname) {
		return true
	}

	return false
}

// SetDeliverFullname gets a reference to the given string and assigns it to the DeliverFullname field.
func (o *DeliverAddress) SetDeliverFullname(v string) {
	o.DeliverFullname = &v
}

// GetDeliverMobile returns the DeliverMobile field value if set, zero value otherwise.
func (o *DeliverAddress) GetDeliverMobile() string {
	if o == nil || IsNil(o.DeliverMobile) {
		var ret string
		return ret
	}
	return *o.DeliverMobile
}

// GetDeliverMobileOk returns a tuple with the DeliverMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverAddress) GetDeliverMobileOk() (*string, bool) {
	if o == nil || IsNil(o.DeliverMobile) {
		return nil, false
	}
	return o.DeliverMobile, true
}

// HasDeliverMobile returns a boolean if a field has been set.
func (o *DeliverAddress) HasDeliverMobile() bool {
	if o != nil && !IsNil(o.DeliverMobile) {
		return true
	}

	return false
}

// SetDeliverMobile gets a reference to the given string and assigns it to the DeliverMobile field.
func (o *DeliverAddress) SetDeliverMobile(v string) {
	o.DeliverMobile = &v
}

// GetDeliverPhone returns the DeliverPhone field value if set, zero value otherwise.
func (o *DeliverAddress) GetDeliverPhone() string {
	if o == nil || IsNil(o.DeliverPhone) {
		var ret string
		return ret
	}
	return *o.DeliverPhone
}

// GetDeliverPhoneOk returns a tuple with the DeliverPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverAddress) GetDeliverPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.DeliverPhone) {
		return nil, false
	}
	return o.DeliverPhone, true
}

// HasDeliverPhone returns a boolean if a field has been set.
func (o *DeliverAddress) HasDeliverPhone() bool {
	if o != nil && !IsNil(o.DeliverPhone) {
		return true
	}

	return false
}

// SetDeliverPhone gets a reference to the given string and assigns it to the DeliverPhone field.
func (o *DeliverAddress) SetDeliverPhone(v string) {
	o.DeliverPhone = &v
}

// GetDeliverProvince returns the DeliverProvince field value if set, zero value otherwise.
func (o *DeliverAddress) GetDeliverProvince() string {
	if o == nil || IsNil(o.DeliverProvince) {
		var ret string
		return ret
	}
	return *o.DeliverProvince
}

// GetDeliverProvinceOk returns a tuple with the DeliverProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverAddress) GetDeliverProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.DeliverProvince) {
		return nil, false
	}
	return o.DeliverProvince, true
}

// HasDeliverProvince returns a boolean if a field has been set.
func (o *DeliverAddress) HasDeliverProvince() bool {
	if o != nil && !IsNil(o.DeliverProvince) {
		return true
	}

	return false
}

// SetDeliverProvince gets a reference to the given string and assigns it to the DeliverProvince field.
func (o *DeliverAddress) SetDeliverProvince(v string) {
	o.DeliverProvince = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *DeliverAddress) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverAddress) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *DeliverAddress) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *DeliverAddress) SetZip(v string) {
	o.Zip = &v
}

func (o DeliverAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliverAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AddressCode) {
		toSerialize["address_code"] = o.AddressCode
	}
	if !IsNil(o.DefaultDeliverAddress) {
		toSerialize["default_deliver_address"] = o.DefaultDeliverAddress
	}
	if !IsNil(o.DeliverArea) {
		toSerialize["deliver_area"] = o.DeliverArea
	}
	if !IsNil(o.DeliverCity) {
		toSerialize["deliver_city"] = o.DeliverCity
	}
	if !IsNil(o.DeliverFullname) {
		toSerialize["deliver_fullname"] = o.DeliverFullname
	}
	if !IsNil(o.DeliverMobile) {
		toSerialize["deliver_mobile"] = o.DeliverMobile
	}
	if !IsNil(o.DeliverPhone) {
		toSerialize["deliver_phone"] = o.DeliverPhone
	}
	if !IsNil(o.DeliverProvince) {
		toSerialize["deliver_province"] = o.DeliverProvince
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	return toSerialize, nil
}

type NullableDeliverAddress struct {
	value *DeliverAddress
	isSet bool
}

func (v NullableDeliverAddress) Get() *DeliverAddress {
	return v.value
}

func (v *NullableDeliverAddress) Set(val *DeliverAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliverAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliverAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliverAddress(val *DeliverAddress) *NullableDeliverAddress {
	return &NullableDeliverAddress{value: val, isSet: true}
}

func (v NullableDeliverAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliverAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

