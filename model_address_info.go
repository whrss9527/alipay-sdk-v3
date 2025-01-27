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

// checks if the AddressInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressInfo{}

// AddressInfo struct for AddressInfo
type AddressInfo struct {
	// 地址。商户详细经营地址或人员所在地点
	Address *string `json:"address,omitempty"`
	// 城市编码。 蚂蚁店铺请按照<a href=\"https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx\" target=\"_blank\">蚂蚁店铺地区码</a> 表格中填写。 直付通商户请按照<a href=\"https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx \" target=\"_blank\">直付通商户地区码</a> 表格中内容填写。
	CityCode *string `json:"city_code,omitempty"`
	// 区县编码。 蚂蚁店铺请按照<a href=\"https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx\" target=\"_blank\">蚂蚁店铺地区码</a> 表格中填写。 直付通商户请按照<a href=\"https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx \" target=\"_blank\">直付通商户地区码</a> 表格中内容填写。
	DistrictCode *string `json:"district_code,omitempty"`
	// 纬度，浮点型,小数点后最多保留6位  如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。<a href=\"http://lbs.amap.com/console/show/picker\">高德经纬度查询</a>
	Latitude *string `json:"latitude,omitempty"`
	// 经度，浮点型, 小数点后最多保留6位。  如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。<a href=\"http://lbs.amap.com/console/show/picker\">高德经纬度查询</a>
	Longitude *string `json:"longitude,omitempty"`
	// 高德poiid
	Poiid *string `json:"poiid,omitempty"`
	// 省份编码。 蚂蚁店铺请按照<a href=\"https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx\" target=\"_blank\">蚂蚁店铺地区码</a> 表格中填写。 直付通商户请按照<a href=\"https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx \" target=\"_blank\">直付通商户地区码</a> 表格中内容填写。
	ProvinceCode *string `json:"province_code,omitempty"`
	// 地址类型。取值范围：BUSINESS_ADDRESS：经营地址（默认）
	Type *string `json:"type,omitempty"`
}

// NewAddressInfo instantiates a new AddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressInfo() *AddressInfo {
	this := AddressInfo{}
	return &this
}

// NewAddressInfoWithDefaults instantiates a new AddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressInfoWithDefaults() *AddressInfo {
	this := AddressInfo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AddressInfo) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AddressInfo) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AddressInfo) SetAddress(v string) {
	o.Address = &v
}

// GetCityCode returns the CityCode field value if set, zero value otherwise.
func (o *AddressInfo) GetCityCode() string {
	if o == nil || IsNil(o.CityCode) {
		var ret string
		return ret
	}
	return *o.CityCode
}

// GetCityCodeOk returns a tuple with the CityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetCityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CityCode) {
		return nil, false
	}
	return o.CityCode, true
}

// HasCityCode returns a boolean if a field has been set.
func (o *AddressInfo) HasCityCode() bool {
	if o != nil && !IsNil(o.CityCode) {
		return true
	}

	return false
}

// SetCityCode gets a reference to the given string and assigns it to the CityCode field.
func (o *AddressInfo) SetCityCode(v string) {
	o.CityCode = &v
}

// GetDistrictCode returns the DistrictCode field value if set, zero value otherwise.
func (o *AddressInfo) GetDistrictCode() string {
	if o == nil || IsNil(o.DistrictCode) {
		var ret string
		return ret
	}
	return *o.DistrictCode
}

// GetDistrictCodeOk returns a tuple with the DistrictCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetDistrictCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DistrictCode) {
		return nil, false
	}
	return o.DistrictCode, true
}

// HasDistrictCode returns a boolean if a field has been set.
func (o *AddressInfo) HasDistrictCode() bool {
	if o != nil && !IsNil(o.DistrictCode) {
		return true
	}

	return false
}

// SetDistrictCode gets a reference to the given string and assigns it to the DistrictCode field.
func (o *AddressInfo) SetDistrictCode(v string) {
	o.DistrictCode = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *AddressInfo) GetLatitude() string {
	if o == nil || IsNil(o.Latitude) {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *AddressInfo) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *AddressInfo) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *AddressInfo) GetLongitude() string {
	if o == nil || IsNil(o.Longitude) {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetLongitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *AddressInfo) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *AddressInfo) SetLongitude(v string) {
	o.Longitude = &v
}

// GetPoiid returns the Poiid field value if set, zero value otherwise.
func (o *AddressInfo) GetPoiid() string {
	if o == nil || IsNil(o.Poiid) {
		var ret string
		return ret
	}
	return *o.Poiid
}

// GetPoiidOk returns a tuple with the Poiid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetPoiidOk() (*string, bool) {
	if o == nil || IsNil(o.Poiid) {
		return nil, false
	}
	return o.Poiid, true
}

// HasPoiid returns a boolean if a field has been set.
func (o *AddressInfo) HasPoiid() bool {
	if o != nil && !IsNil(o.Poiid) {
		return true
	}

	return false
}

// SetPoiid gets a reference to the given string and assigns it to the Poiid field.
func (o *AddressInfo) SetPoiid(v string) {
	o.Poiid = &v
}

// GetProvinceCode returns the ProvinceCode field value if set, zero value otherwise.
func (o *AddressInfo) GetProvinceCode() string {
	if o == nil || IsNil(o.ProvinceCode) {
		var ret string
		return ret
	}
	return *o.ProvinceCode
}

// GetProvinceCodeOk returns a tuple with the ProvinceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetProvinceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProvinceCode) {
		return nil, false
	}
	return o.ProvinceCode, true
}

// HasProvinceCode returns a boolean if a field has been set.
func (o *AddressInfo) HasProvinceCode() bool {
	if o != nil && !IsNil(o.ProvinceCode) {
		return true
	}

	return false
}

// SetProvinceCode gets a reference to the given string and assigns it to the ProvinceCode field.
func (o *AddressInfo) SetProvinceCode(v string) {
	o.ProvinceCode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AddressInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AddressInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AddressInfo) SetType(v string) {
	o.Type = &v
}

func (o AddressInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.CityCode) {
		toSerialize["city_code"] = o.CityCode
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
	if !IsNil(o.Poiid) {
		toSerialize["poiid"] = o.Poiid
	}
	if !IsNil(o.ProvinceCode) {
		toSerialize["province_code"] = o.ProvinceCode
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAddressInfo struct {
	value *AddressInfo
	isSet bool
}

func (v NullableAddressInfo) Get() *AddressInfo {
	return v.value
}

func (v *NullableAddressInfo) Set(val *AddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressInfo(val *AddressInfo) *NullableAddressInfo {
	return &NullableAddressInfo{value: val, isSet: true}
}

func (v NullableAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
