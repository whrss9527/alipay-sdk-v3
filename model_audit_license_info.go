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

// checks if the AuditLicenseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLicenseInfo{}

// AuditLicenseInfo struct for AuditLicenseInfo
type AuditLicenseInfo struct {
	// 营业执照名称，需要与营业执照保持一致
	LicenseName *string `json:"license_name,omitempty"`
	// 营业执照号，部分小程序类目需要提交，参照 <a href=‘https://opendocs.alipay.com/b/03al2m’>开放服务类目</a> 中是否需要营业执照信息，如果不填默认采用当前小程序应用营业执照号。
	LicenseNo *string `json:"license_no,omitempty"`
	// 营业执照照片地址列表
	LicensePicList []string `json:"license_pic_list,omitempty"`
	// 营业执有效期，格式为yyyy-MM-dd，9999-12-31表示长期
	LicenseValidDate *string `json:"license_valid_date,omitempty"`
	// 门头照图片地址，部分小程序类目需要提交，参照 <a href=‘https://opendocs.alipay.com/b/03al2m’>开放服务类目</a> 中是否需要营业执照信息，如果不填默认采用当前小程序门头照图片
	OutDoorPic *string `json:"out_door_pic,omitempty"`
}

// NewAuditLicenseInfo instantiates a new AuditLicenseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLicenseInfo() *AuditLicenseInfo {
	this := AuditLicenseInfo{}
	return &this
}

// NewAuditLicenseInfoWithDefaults instantiates a new AuditLicenseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLicenseInfoWithDefaults() *AuditLicenseInfo {
	this := AuditLicenseInfo{}
	return &this
}

// GetLicenseName returns the LicenseName field value if set, zero value otherwise.
func (o *AuditLicenseInfo) GetLicenseName() string {
	if o == nil || IsNil(o.LicenseName) {
		var ret string
		return ret
	}
	return *o.LicenseName
}

// GetLicenseNameOk returns a tuple with the LicenseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLicenseInfo) GetLicenseNameOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseName) {
		return nil, false
	}
	return o.LicenseName, true
}

// HasLicenseName returns a boolean if a field has been set.
func (o *AuditLicenseInfo) HasLicenseName() bool {
	if o != nil && !IsNil(o.LicenseName) {
		return true
	}

	return false
}

// SetLicenseName gets a reference to the given string and assigns it to the LicenseName field.
func (o *AuditLicenseInfo) SetLicenseName(v string) {
	o.LicenseName = &v
}

// GetLicenseNo returns the LicenseNo field value if set, zero value otherwise.
func (o *AuditLicenseInfo) GetLicenseNo() string {
	if o == nil || IsNil(o.LicenseNo) {
		var ret string
		return ret
	}
	return *o.LicenseNo
}

// GetLicenseNoOk returns a tuple with the LicenseNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLicenseInfo) GetLicenseNoOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseNo) {
		return nil, false
	}
	return o.LicenseNo, true
}

// HasLicenseNo returns a boolean if a field has been set.
func (o *AuditLicenseInfo) HasLicenseNo() bool {
	if o != nil && !IsNil(o.LicenseNo) {
		return true
	}

	return false
}

// SetLicenseNo gets a reference to the given string and assigns it to the LicenseNo field.
func (o *AuditLicenseInfo) SetLicenseNo(v string) {
	o.LicenseNo = &v
}

// GetLicensePicList returns the LicensePicList field value if set, zero value otherwise.
func (o *AuditLicenseInfo) GetLicensePicList() []string {
	if o == nil || IsNil(o.LicensePicList) {
		var ret []string
		return ret
	}
	return o.LicensePicList
}

// GetLicensePicListOk returns a tuple with the LicensePicList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLicenseInfo) GetLicensePicListOk() ([]string, bool) {
	if o == nil || IsNil(o.LicensePicList) {
		return nil, false
	}
	return o.LicensePicList, true
}

// HasLicensePicList returns a boolean if a field has been set.
func (o *AuditLicenseInfo) HasLicensePicList() bool {
	if o != nil && !IsNil(o.LicensePicList) {
		return true
	}

	return false
}

// SetLicensePicList gets a reference to the given []string and assigns it to the LicensePicList field.
func (o *AuditLicenseInfo) SetLicensePicList(v []string) {
	o.LicensePicList = v
}

// GetLicenseValidDate returns the LicenseValidDate field value if set, zero value otherwise.
func (o *AuditLicenseInfo) GetLicenseValidDate() string {
	if o == nil || IsNil(o.LicenseValidDate) {
		var ret string
		return ret
	}
	return *o.LicenseValidDate
}

// GetLicenseValidDateOk returns a tuple with the LicenseValidDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLicenseInfo) GetLicenseValidDateOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseValidDate) {
		return nil, false
	}
	return o.LicenseValidDate, true
}

// HasLicenseValidDate returns a boolean if a field has been set.
func (o *AuditLicenseInfo) HasLicenseValidDate() bool {
	if o != nil && !IsNil(o.LicenseValidDate) {
		return true
	}

	return false
}

// SetLicenseValidDate gets a reference to the given string and assigns it to the LicenseValidDate field.
func (o *AuditLicenseInfo) SetLicenseValidDate(v string) {
	o.LicenseValidDate = &v
}

// GetOutDoorPic returns the OutDoorPic field value if set, zero value otherwise.
func (o *AuditLicenseInfo) GetOutDoorPic() string {
	if o == nil || IsNil(o.OutDoorPic) {
		var ret string
		return ret
	}
	return *o.OutDoorPic
}

// GetOutDoorPicOk returns a tuple with the OutDoorPic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLicenseInfo) GetOutDoorPicOk() (*string, bool) {
	if o == nil || IsNil(o.OutDoorPic) {
		return nil, false
	}
	return o.OutDoorPic, true
}

// HasOutDoorPic returns a boolean if a field has been set.
func (o *AuditLicenseInfo) HasOutDoorPic() bool {
	if o != nil && !IsNil(o.OutDoorPic) {
		return true
	}

	return false
}

// SetOutDoorPic gets a reference to the given string and assigns it to the OutDoorPic field.
func (o *AuditLicenseInfo) SetOutDoorPic(v string) {
	o.OutDoorPic = &v
}

func (o AuditLicenseInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLicenseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LicenseName) {
		toSerialize["license_name"] = o.LicenseName
	}
	if !IsNil(o.LicenseNo) {
		toSerialize["license_no"] = o.LicenseNo
	}
	if !IsNil(o.LicensePicList) {
		toSerialize["license_pic_list"] = o.LicensePicList
	}
	if !IsNil(o.LicenseValidDate) {
		toSerialize["license_valid_date"] = o.LicenseValidDate
	}
	if !IsNil(o.OutDoorPic) {
		toSerialize["out_door_pic"] = o.OutDoorPic
	}
	return toSerialize, nil
}

type NullableAuditLicenseInfo struct {
	value *AuditLicenseInfo
	isSet bool
}

func (v NullableAuditLicenseInfo) Get() *AuditLicenseInfo {
	return v.value
}

func (v *NullableAuditLicenseInfo) Set(val *AuditLicenseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLicenseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLicenseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLicenseInfo(val *AuditLicenseInfo) *NullableAuditLicenseInfo {
	return &NullableAuditLicenseInfo{value: val, isSet: true}
}

func (v NullableAuditLicenseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLicenseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
