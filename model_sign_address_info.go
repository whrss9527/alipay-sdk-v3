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

// checks if the SignAddressInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignAddressInfo{}

// SignAddressInfo struct for SignAddressInfo
type SignAddressInfo struct {
	// 城市编码。请按照<a href=\"https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx\">表格</a> 中内容填写。 （参考资料： <a href=\"http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/\">参考资料</a>）
	CityCode *string `json:"city_code,omitempty"`
	// 国家编码，中国默认：156
	CountryCode *string `json:"country_code,omitempty"`
	// 详细地址
	DetailAddress *string `json:"detail_address,omitempty"`
	// 区县编码。请按照<a href=\"https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx\">表格</a> 中内容填写。 （参考资料： <a href=\"http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/\">参考资料</a>）
	DistrictCode *string `json:"district_code,omitempty"`
	// 纬度，浮点型,小数点后最多保留6位 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：<a href=\"http://lbs.amap.com/console/show/picker\">查询地址</a>
	Latitude *string `json:"latitude,omitempty"`
	// 经度，浮点型, 小数点后最多保留6位。 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：<a href=\"http://lbs.amap.com/console/show/picker\">查询地址</a>
	Longitude *string `json:"longitude,omitempty"`
	// 省份编码。请按照<a href=\"https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx\">表格</a> 中内容填写。 （参考资料： <a href=\"http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/\">参考资料</a>）
	ProvinceCode *string `json:"province_code,omitempty"`
}

// NewSignAddressInfo instantiates a new SignAddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignAddressInfo() *SignAddressInfo {
	this := SignAddressInfo{}
	return &this
}

// NewSignAddressInfoWithDefaults instantiates a new SignAddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignAddressInfoWithDefaults() *SignAddressInfo {
	this := SignAddressInfo{}
	return &this
}

// GetCityCode returns the CityCode field value if set, zero value otherwise.
func (o *SignAddressInfo) GetCityCode() string {
	if o == nil || IsNil(o.CityCode) {
		var ret string
		return ret
	}
	return *o.CityCode
}

// GetCityCodeOk returns a tuple with the CityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignAddressInfo) GetCityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CityCode) {
		return nil, false
	}
	return o.CityCode, true
}

// HasCityCode returns a boolean if a field has been set.
func (o *SignAddressInfo) HasCityCode() bool {
	if o != nil && !IsNil(o.CityCode) {
		return true
	}

	return false
}

// SetCityCode gets a reference to the given string and assigns it to the CityCode field.
func (o *SignAddressInfo) SetCityCode(v string) {
	o.CityCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *SignAddressInfo) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignAddressInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *SignAddressInfo) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *SignAddressInfo) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetDetailAddress returns the DetailAddress field value if set, zero value otherwise.
func (o *SignAddressInfo) GetDetailAddress() string {
	if o == nil || IsNil(o.DetailAddress) {
		var ret string
		return ret
	}
	return *o.DetailAddress
}

// GetDetailAddressOk returns a tuple with the DetailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignAddressInfo) GetDetailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.DetailAddress) {
		return nil, false
	}
	return o.DetailAddress, true
}

// HasDetailAddress returns a boolean if a field has been set.
func (o *SignAddressInfo) HasDetailAddress() bool {
	if o != nil && !IsNil(o.DetailAddress) {
		return true
	}

	return false
}

// SetDetailAddress gets a reference to the given string and assigns it to the DetailAddress field.
func (o *SignAddressInfo) SetDetailAddress(v string) {
	o.DetailAddress = &v
}

// GetDistrictCode returns the DistrictCode field value if set, zero value otherwise.
func (o *SignAddressInfo) GetDistrictCode() string {
	if o == nil || IsNil(o.DistrictCode) {
		var ret string
		return ret
	}
	return *o.DistrictCode
}

// GetDistrictCodeOk returns a tuple with the DistrictCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignAddressInfo) GetDistrictCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DistrictCode) {
		return nil, false
	}
	return o.DistrictCode, true
}

// HasDistrictCode returns a boolean if a field has been set.
func (o *SignAddressInfo) HasDistrictCode() bool {
	if o != nil && !IsNil(o.DistrictCode) {
		return true
	}

	return false
}

// SetDistrictCode gets a reference to the given string and assigns it to the DistrictCode field.
func (o *SignAddressInfo) SetDistrictCode(v string) {
	o.DistrictCode = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *SignAddressInfo) GetLatitude() string {
	if o == nil || IsNil(o.Latitude) {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignAddressInfo) GetLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *SignAddressInfo) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *SignAddressInfo) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *SignAddressInfo) GetLongitude() string {
	if o == nil || IsNil(o.Longitude) {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignAddressInfo) GetLongitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *SignAddressInfo) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *SignAddressInfo) SetLongitude(v string) {
	o.Longitude = &v
}

// GetProvinceCode returns the ProvinceCode field value if set, zero value otherwise.
func (o *SignAddressInfo) GetProvinceCode() string {
	if o == nil || IsNil(o.ProvinceCode) {
		var ret string
		return ret
	}
	return *o.ProvinceCode
}

// GetProvinceCodeOk returns a tuple with the ProvinceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignAddressInfo) GetProvinceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProvinceCode) {
		return nil, false
	}
	return o.ProvinceCode, true
}

// HasProvinceCode returns a boolean if a field has been set.
func (o *SignAddressInfo) HasProvinceCode() bool {
	if o != nil && !IsNil(o.ProvinceCode) {
		return true
	}

	return false
}

// SetProvinceCode gets a reference to the given string and assigns it to the ProvinceCode field.
func (o *SignAddressInfo) SetProvinceCode(v string) {
	o.ProvinceCode = &v
}

func (o SignAddressInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignAddressInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CityCode) {
		toSerialize["city_code"] = o.CityCode
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.DetailAddress) {
		toSerialize["detail_address"] = o.DetailAddress
	}
	if !IsNil(o.DistrictCode) {
		toSerialize["district_code"] = o.DistrictCode
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.ProvinceCode) {
		toSerialize["province_code"] = o.ProvinceCode
	}
	return toSerialize, nil
}

type NullableSignAddressInfo struct {
	value *SignAddressInfo
	isSet bool
}

func (v NullableSignAddressInfo) Get() *SignAddressInfo {
	return v.value
}

func (v *NullableSignAddressInfo) Set(val *SignAddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSignAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSignAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignAddressInfo(val *SignAddressInfo) *NullableSignAddressInfo {
	return &NullableSignAddressInfo{value: val, isSet: true}
}

func (v NullableSignAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
