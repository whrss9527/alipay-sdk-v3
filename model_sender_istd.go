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

// checks if the SenderIstd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SenderIstd{}

// SenderIstd struct for SenderIstd
type SenderIstd struct {
	// 地址(街道、小区、大厦等，用于定位)
	Address *string `json:"address,omitempty"`
	// 地址详情(楼号、单元号、层号)
	AddressDetail *string `json:"address_detail,omitempty"`
	// 城市名称，如杭州市
	City *string `json:"city,omitempty"`
	// 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1:百度坐标， 目前只支持0
	CoordinateType *int32 `json:"coordinate_type,omitempty"`
	// 纬度
	Lat *string `json:"lat,omitempty"`
	// 经度
	Lng *string `json:"lng,omitempty"`
	// 手机号
	MobileNo *string `json:"mobile_no,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty"`
}

// NewSenderIstd instantiates a new SenderIstd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSenderIstd() *SenderIstd {
	this := SenderIstd{}
	return &this
}

// NewSenderIstdWithDefaults instantiates a new SenderIstd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSenderIstdWithDefaults() *SenderIstd {
	this := SenderIstd{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *SenderIstd) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderIstd) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *SenderIstd) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *SenderIstd) SetAddress(v string) {
	o.Address = &v
}

// GetAddressDetail returns the AddressDetail field value if set, zero value otherwise.
func (o *SenderIstd) GetAddressDetail() string {
	if o == nil || IsNil(o.AddressDetail) {
		var ret string
		return ret
	}
	return *o.AddressDetail
}

// GetAddressDetailOk returns a tuple with the AddressDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderIstd) GetAddressDetailOk() (*string, bool) {
	if o == nil || IsNil(o.AddressDetail) {
		return nil, false
	}
	return o.AddressDetail, true
}

// HasAddressDetail returns a boolean if a field has been set.
func (o *SenderIstd) HasAddressDetail() bool {
	if o != nil && !IsNil(o.AddressDetail) {
		return true
	}

	return false
}

// SetAddressDetail gets a reference to the given string and assigns it to the AddressDetail field.
func (o *SenderIstd) SetAddressDetail(v string) {
	o.AddressDetail = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *SenderIstd) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderIstd) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *SenderIstd) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *SenderIstd) SetCity(v string) {
	o.City = &v
}

// GetCoordinateType returns the CoordinateType field value if set, zero value otherwise.
func (o *SenderIstd) GetCoordinateType() int32 {
	if o == nil || IsNil(o.CoordinateType) {
		var ret int32
		return ret
	}
	return *o.CoordinateType
}

// GetCoordinateTypeOk returns a tuple with the CoordinateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderIstd) GetCoordinateTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.CoordinateType) {
		return nil, false
	}
	return o.CoordinateType, true
}

// HasCoordinateType returns a boolean if a field has been set.
func (o *SenderIstd) HasCoordinateType() bool {
	if o != nil && !IsNil(o.CoordinateType) {
		return true
	}

	return false
}

// SetCoordinateType gets a reference to the given int32 and assigns it to the CoordinateType field.
func (o *SenderIstd) SetCoordinateType(v int32) {
	o.CoordinateType = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *SenderIstd) GetLat() string {
	if o == nil || IsNil(o.Lat) {
		var ret string
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderIstd) GetLatOk() (*string, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *SenderIstd) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given string and assigns it to the Lat field.
func (o *SenderIstd) SetLat(v string) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *SenderIstd) GetLng() string {
	if o == nil || IsNil(o.Lng) {
		var ret string
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderIstd) GetLngOk() (*string, bool) {
	if o == nil || IsNil(o.Lng) {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *SenderIstd) HasLng() bool {
	if o != nil && !IsNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given string and assigns it to the Lng field.
func (o *SenderIstd) SetLng(v string) {
	o.Lng = &v
}

// GetMobileNo returns the MobileNo field value if set, zero value otherwise.
func (o *SenderIstd) GetMobileNo() string {
	if o == nil || IsNil(o.MobileNo) {
		var ret string
		return ret
	}
	return *o.MobileNo
}

// GetMobileNoOk returns a tuple with the MobileNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderIstd) GetMobileNoOk() (*string, bool) {
	if o == nil || IsNil(o.MobileNo) {
		return nil, false
	}
	return o.MobileNo, true
}

// HasMobileNo returns a boolean if a field has been set.
func (o *SenderIstd) HasMobileNo() bool {
	if o != nil && !IsNil(o.MobileNo) {
		return true
	}

	return false
}

// SetMobileNo gets a reference to the given string and assigns it to the MobileNo field.
func (o *SenderIstd) SetMobileNo(v string) {
	o.MobileNo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SenderIstd) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderIstd) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SenderIstd) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SenderIstd) SetName(v string) {
	o.Name = &v
}

func (o SenderIstd) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SenderIstd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AddressDetail) {
		toSerialize["address_detail"] = o.AddressDetail
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.CoordinateType) {
		toSerialize["coordinate_type"] = o.CoordinateType
	}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	if !IsNil(o.MobileNo) {
		toSerialize["mobile_no"] = o.MobileNo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSenderIstd struct {
	value *SenderIstd
	isSet bool
}

func (v NullableSenderIstd) Get() *SenderIstd {
	return v.value
}

func (v *NullableSenderIstd) Set(val *SenderIstd) {
	v.value = val
	v.isSet = true
}

func (v NullableSenderIstd) IsSet() bool {
	return v.isSet
}

func (v *NullableSenderIstd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSenderIstd(val *SenderIstd) *NullableSenderIstd {
	return &NullableSenderIstd{value: val, isSet: true}
}

func (v NullableSenderIstd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSenderIstd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
