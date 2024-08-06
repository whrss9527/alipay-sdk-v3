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

// checks if the MiniAppCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiniAppCategory{}

// MiniAppCategory struct for MiniAppCategory
type MiniAppCategory struct {
	// 类目id
	CategoryId *string `json:"category_id,omitempty"`
	// 类目名称
	CategoryName *string `json:"category_name,omitempty"`
	// 是否有子类目
	HasChild *bool `json:"has_child,omitempty"`
	// 是否需要营业执照
	NeedLicense *bool `json:"need_license,omitempty"`
	// 是否需要门头照
	NeedOutDoorPic *bool `json:"need_out_door_pic,omitempty"`
	// 是否需要特许营业执照
	NeedSpecialLicense *bool `json:"need_special_license,omitempty"`
	// 父类目id
	ParentCategoryId *string `json:"parent_category_id,omitempty"`
}

// NewMiniAppCategory instantiates a new MiniAppCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiniAppCategory() *MiniAppCategory {
	this := MiniAppCategory{}
	return &this
}

// NewMiniAppCategoryWithDefaults instantiates a new MiniAppCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiniAppCategoryWithDefaults() *MiniAppCategory {
	this := MiniAppCategory{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *MiniAppCategory) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategory) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *MiniAppCategory) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *MiniAppCategory) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *MiniAppCategory) GetCategoryName() string {
	if o == nil || IsNil(o.CategoryName) {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategory) GetCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryName) {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *MiniAppCategory) HasCategoryName() bool {
	if o != nil && !IsNil(o.CategoryName) {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *MiniAppCategory) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetHasChild returns the HasChild field value if set, zero value otherwise.
func (o *MiniAppCategory) GetHasChild() bool {
	if o == nil || IsNil(o.HasChild) {
		var ret bool
		return ret
	}
	return *o.HasChild
}

// GetHasChildOk returns a tuple with the HasChild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategory) GetHasChildOk() (*bool, bool) {
	if o == nil || IsNil(o.HasChild) {
		return nil, false
	}
	return o.HasChild, true
}

// HasHasChild returns a boolean if a field has been set.
func (o *MiniAppCategory) HasHasChild() bool {
	if o != nil && !IsNil(o.HasChild) {
		return true
	}

	return false
}

// SetHasChild gets a reference to the given bool and assigns it to the HasChild field.
func (o *MiniAppCategory) SetHasChild(v bool) {
	o.HasChild = &v
}

// GetNeedLicense returns the NeedLicense field value if set, zero value otherwise.
func (o *MiniAppCategory) GetNeedLicense() bool {
	if o == nil || IsNil(o.NeedLicense) {
		var ret bool
		return ret
	}
	return *o.NeedLicense
}

// GetNeedLicenseOk returns a tuple with the NeedLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategory) GetNeedLicenseOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedLicense) {
		return nil, false
	}
	return o.NeedLicense, true
}

// HasNeedLicense returns a boolean if a field has been set.
func (o *MiniAppCategory) HasNeedLicense() bool {
	if o != nil && !IsNil(o.NeedLicense) {
		return true
	}

	return false
}

// SetNeedLicense gets a reference to the given bool and assigns it to the NeedLicense field.
func (o *MiniAppCategory) SetNeedLicense(v bool) {
	o.NeedLicense = &v
}

// GetNeedOutDoorPic returns the NeedOutDoorPic field value if set, zero value otherwise.
func (o *MiniAppCategory) GetNeedOutDoorPic() bool {
	if o == nil || IsNil(o.NeedOutDoorPic) {
		var ret bool
		return ret
	}
	return *o.NeedOutDoorPic
}

// GetNeedOutDoorPicOk returns a tuple with the NeedOutDoorPic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategory) GetNeedOutDoorPicOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedOutDoorPic) {
		return nil, false
	}
	return o.NeedOutDoorPic, true
}

// HasNeedOutDoorPic returns a boolean if a field has been set.
func (o *MiniAppCategory) HasNeedOutDoorPic() bool {
	if o != nil && !IsNil(o.NeedOutDoorPic) {
		return true
	}

	return false
}

// SetNeedOutDoorPic gets a reference to the given bool and assigns it to the NeedOutDoorPic field.
func (o *MiniAppCategory) SetNeedOutDoorPic(v bool) {
	o.NeedOutDoorPic = &v
}

// GetNeedSpecialLicense returns the NeedSpecialLicense field value if set, zero value otherwise.
func (o *MiniAppCategory) GetNeedSpecialLicense() bool {
	if o == nil || IsNil(o.NeedSpecialLicense) {
		var ret bool
		return ret
	}
	return *o.NeedSpecialLicense
}

// GetNeedSpecialLicenseOk returns a tuple with the NeedSpecialLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategory) GetNeedSpecialLicenseOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedSpecialLicense) {
		return nil, false
	}
	return o.NeedSpecialLicense, true
}

// HasNeedSpecialLicense returns a boolean if a field has been set.
func (o *MiniAppCategory) HasNeedSpecialLicense() bool {
	if o != nil && !IsNil(o.NeedSpecialLicense) {
		return true
	}

	return false
}

// SetNeedSpecialLicense gets a reference to the given bool and assigns it to the NeedSpecialLicense field.
func (o *MiniAppCategory) SetNeedSpecialLicense(v bool) {
	o.NeedSpecialLicense = &v
}

// GetParentCategoryId returns the ParentCategoryId field value if set, zero value otherwise.
func (o *MiniAppCategory) GetParentCategoryId() string {
	if o == nil || IsNil(o.ParentCategoryId) {
		var ret string
		return ret
	}
	return *o.ParentCategoryId
}

// GetParentCategoryIdOk returns a tuple with the ParentCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategory) GetParentCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCategoryId) {
		return nil, false
	}
	return o.ParentCategoryId, true
}

// HasParentCategoryId returns a boolean if a field has been set.
func (o *MiniAppCategory) HasParentCategoryId() bool {
	if o != nil && !IsNil(o.ParentCategoryId) {
		return true
	}

	return false
}

// SetParentCategoryId gets a reference to the given string and assigns it to the ParentCategoryId field.
func (o *MiniAppCategory) SetParentCategoryId(v string) {
	o.ParentCategoryId = &v
}

func (o MiniAppCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiniAppCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.CategoryName) {
		toSerialize["category_name"] = o.CategoryName
	}
	if !IsNil(o.HasChild) {
		toSerialize["has_child"] = o.HasChild
	}
	if !IsNil(o.NeedLicense) {
		toSerialize["need_license"] = o.NeedLicense
	}
	if !IsNil(o.NeedOutDoorPic) {
		toSerialize["need_out_door_pic"] = o.NeedOutDoorPic
	}
	if !IsNil(o.NeedSpecialLicense) {
		toSerialize["need_special_license"] = o.NeedSpecialLicense
	}
	if !IsNil(o.ParentCategoryId) {
		toSerialize["parent_category_id"] = o.ParentCategoryId
	}
	return toSerialize, nil
}

type NullableMiniAppCategory struct {
	value *MiniAppCategory
	isSet bool
}

func (v NullableMiniAppCategory) Get() *MiniAppCategory {
	return v.value
}

func (v *NullableMiniAppCategory) Set(val *MiniAppCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableMiniAppCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableMiniAppCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiniAppCategory(val *MiniAppCategory) *NullableMiniAppCategory {
	return &NullableMiniAppCategory{value: val, isSet: true}
}

func (v NullableMiniAppCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiniAppCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

