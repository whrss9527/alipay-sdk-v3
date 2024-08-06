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

// checks if the CommonVoucherDisplayInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonVoucherDisplayInfo{}

// CommonVoucherDisplayInfo struct for CommonVoucherDisplayInfo
type CommonVoucherDisplayInfo struct {
	// 商家品牌 logo 链接
	BrandLogo *string `json:"brand_logo,omitempty"`
	// 商户品牌名称。
	BrandName *string `json:"brand_name,omitempty"`
	// 用于说明详细的活动规则，会展示在支付宝卡包券详情页。
	VoucherDescription *string `json:"voucher_description,omitempty"`
	// 券详情图片链接。
	VoucherDetailImages []string `json:"voucher_detail_images,omitempty"`
	// 券封面图链接。
	VoucherImage *string `json:"voucher_image,omitempty"`
}

// NewCommonVoucherDisplayInfo instantiates a new CommonVoucherDisplayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonVoucherDisplayInfo() *CommonVoucherDisplayInfo {
	this := CommonVoucherDisplayInfo{}
	return &this
}

// NewCommonVoucherDisplayInfoWithDefaults instantiates a new CommonVoucherDisplayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonVoucherDisplayInfoWithDefaults() *CommonVoucherDisplayInfo {
	this := CommonVoucherDisplayInfo{}
	return &this
}

// GetBrandLogo returns the BrandLogo field value if set, zero value otherwise.
func (o *CommonVoucherDisplayInfo) GetBrandLogo() string {
	if o == nil || IsNil(o.BrandLogo) {
		var ret string
		return ret
	}
	return *o.BrandLogo
}

// GetBrandLogoOk returns a tuple with the BrandLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonVoucherDisplayInfo) GetBrandLogoOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogo) {
		return nil, false
	}
	return o.BrandLogo, true
}

// HasBrandLogo returns a boolean if a field has been set.
func (o *CommonVoucherDisplayInfo) HasBrandLogo() bool {
	if o != nil && !IsNil(o.BrandLogo) {
		return true
	}

	return false
}

// SetBrandLogo gets a reference to the given string and assigns it to the BrandLogo field.
func (o *CommonVoucherDisplayInfo) SetBrandLogo(v string) {
	o.BrandLogo = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *CommonVoucherDisplayInfo) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonVoucherDisplayInfo) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *CommonVoucherDisplayInfo) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *CommonVoucherDisplayInfo) SetBrandName(v string) {
	o.BrandName = &v
}

// GetVoucherDescription returns the VoucherDescription field value if set, zero value otherwise.
func (o *CommonVoucherDisplayInfo) GetVoucherDescription() string {
	if o == nil || IsNil(o.VoucherDescription) {
		var ret string
		return ret
	}
	return *o.VoucherDescription
}

// GetVoucherDescriptionOk returns a tuple with the VoucherDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonVoucherDisplayInfo) GetVoucherDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherDescription) {
		return nil, false
	}
	return o.VoucherDescription, true
}

// HasVoucherDescription returns a boolean if a field has been set.
func (o *CommonVoucherDisplayInfo) HasVoucherDescription() bool {
	if o != nil && !IsNil(o.VoucherDescription) {
		return true
	}

	return false
}

// SetVoucherDescription gets a reference to the given string and assigns it to the VoucherDescription field.
func (o *CommonVoucherDisplayInfo) SetVoucherDescription(v string) {
	o.VoucherDescription = &v
}

// GetVoucherDetailImages returns the VoucherDetailImages field value if set, zero value otherwise.
func (o *CommonVoucherDisplayInfo) GetVoucherDetailImages() []string {
	if o == nil || IsNil(o.VoucherDetailImages) {
		var ret []string
		return ret
	}
	return o.VoucherDetailImages
}

// GetVoucherDetailImagesOk returns a tuple with the VoucherDetailImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonVoucherDisplayInfo) GetVoucherDetailImagesOk() ([]string, bool) {
	if o == nil || IsNil(o.VoucherDetailImages) {
		return nil, false
	}
	return o.VoucherDetailImages, true
}

// HasVoucherDetailImages returns a boolean if a field has been set.
func (o *CommonVoucherDisplayInfo) HasVoucherDetailImages() bool {
	if o != nil && !IsNil(o.VoucherDetailImages) {
		return true
	}

	return false
}

// SetVoucherDetailImages gets a reference to the given []string and assigns it to the VoucherDetailImages field.
func (o *CommonVoucherDisplayInfo) SetVoucherDetailImages(v []string) {
	o.VoucherDetailImages = v
}

// GetVoucherImage returns the VoucherImage field value if set, zero value otherwise.
func (o *CommonVoucherDisplayInfo) GetVoucherImage() string {
	if o == nil || IsNil(o.VoucherImage) {
		var ret string
		return ret
	}
	return *o.VoucherImage
}

// GetVoucherImageOk returns a tuple with the VoucherImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonVoucherDisplayInfo) GetVoucherImageOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherImage) {
		return nil, false
	}
	return o.VoucherImage, true
}

// HasVoucherImage returns a boolean if a field has been set.
func (o *CommonVoucherDisplayInfo) HasVoucherImage() bool {
	if o != nil && !IsNil(o.VoucherImage) {
		return true
	}

	return false
}

// SetVoucherImage gets a reference to the given string and assigns it to the VoucherImage field.
func (o *CommonVoucherDisplayInfo) SetVoucherImage(v string) {
	o.VoucherImage = &v
}

func (o CommonVoucherDisplayInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonVoucherDisplayInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandLogo) {
		toSerialize["brand_logo"] = o.BrandLogo
	}
	if !IsNil(o.BrandName) {
		toSerialize["brand_name"] = o.BrandName
	}
	if !IsNil(o.VoucherDescription) {
		toSerialize["voucher_description"] = o.VoucherDescription
	}
	if !IsNil(o.VoucherDetailImages) {
		toSerialize["voucher_detail_images"] = o.VoucherDetailImages
	}
	if !IsNil(o.VoucherImage) {
		toSerialize["voucher_image"] = o.VoucherImage
	}
	return toSerialize, nil
}

type NullableCommonVoucherDisplayInfo struct {
	value *CommonVoucherDisplayInfo
	isSet bool
}

func (v NullableCommonVoucherDisplayInfo) Get() *CommonVoucherDisplayInfo {
	return v.value
}

func (v *NullableCommonVoucherDisplayInfo) Set(val *CommonVoucherDisplayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonVoucherDisplayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonVoucherDisplayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonVoucherDisplayInfo(val *CommonVoucherDisplayInfo) *NullableCommonVoucherDisplayInfo {
	return &NullableCommonVoucherDisplayInfo{value: val, isSet: true}
}

func (v NullableCommonVoucherDisplayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonVoucherDisplayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

