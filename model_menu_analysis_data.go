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

// checks if the MenuAnalysisData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MenuAnalysisData{}

// MenuAnalysisData struct for MenuAnalysisData
type MenuAnalysisData struct {
	// 人均点击次数
	AvgClickUserCnt *string `json:"avg_click_user_cnt,omitempty"`
	// 菜单点击次数
	ClickCnt *int32 `json:"click_cnt,omitempty"`
	// 菜单点击人数
	ClickUserCnt *int32 `json:"click_user_cnt,omitempty"`
	// 日期
	Date *string `json:"date,omitempty"`
	// 菜单类型 ，iconDefault ：图标菜单、default：文字菜单
	MenuType *string `json:"menu_type,omitempty"`
	// 菜单名称
	Name *string `json:"name,omitempty"`
	// 子菜单名称，文字菜单才有
	SubName *string `json:"sub_name,omitempty"`
}

// NewMenuAnalysisData instantiates a new MenuAnalysisData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMenuAnalysisData() *MenuAnalysisData {
	this := MenuAnalysisData{}
	return &this
}

// NewMenuAnalysisDataWithDefaults instantiates a new MenuAnalysisData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMenuAnalysisDataWithDefaults() *MenuAnalysisData {
	this := MenuAnalysisData{}
	return &this
}

// GetAvgClickUserCnt returns the AvgClickUserCnt field value if set, zero value otherwise.
func (o *MenuAnalysisData) GetAvgClickUserCnt() string {
	if o == nil || IsNil(o.AvgClickUserCnt) {
		var ret string
		return ret
	}
	return *o.AvgClickUserCnt
}

// GetAvgClickUserCntOk returns a tuple with the AvgClickUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuAnalysisData) GetAvgClickUserCntOk() (*string, bool) {
	if o == nil || IsNil(o.AvgClickUserCnt) {
		return nil, false
	}
	return o.AvgClickUserCnt, true
}

// HasAvgClickUserCnt returns a boolean if a field has been set.
func (o *MenuAnalysisData) HasAvgClickUserCnt() bool {
	if o != nil && !IsNil(o.AvgClickUserCnt) {
		return true
	}

	return false
}

// SetAvgClickUserCnt gets a reference to the given string and assigns it to the AvgClickUserCnt field.
func (o *MenuAnalysisData) SetAvgClickUserCnt(v string) {
	o.AvgClickUserCnt = &v
}

// GetClickCnt returns the ClickCnt field value if set, zero value otherwise.
func (o *MenuAnalysisData) GetClickCnt() int32 {
	if o == nil || IsNil(o.ClickCnt) {
		var ret int32
		return ret
	}
	return *o.ClickCnt
}

// GetClickCntOk returns a tuple with the ClickCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuAnalysisData) GetClickCntOk() (*int32, bool) {
	if o == nil || IsNil(o.ClickCnt) {
		return nil, false
	}
	return o.ClickCnt, true
}

// HasClickCnt returns a boolean if a field has been set.
func (o *MenuAnalysisData) HasClickCnt() bool {
	if o != nil && !IsNil(o.ClickCnt) {
		return true
	}

	return false
}

// SetClickCnt gets a reference to the given int32 and assigns it to the ClickCnt field.
func (o *MenuAnalysisData) SetClickCnt(v int32) {
	o.ClickCnt = &v
}

// GetClickUserCnt returns the ClickUserCnt field value if set, zero value otherwise.
func (o *MenuAnalysisData) GetClickUserCnt() int32 {
	if o == nil || IsNil(o.ClickUserCnt) {
		var ret int32
		return ret
	}
	return *o.ClickUserCnt
}

// GetClickUserCntOk returns a tuple with the ClickUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuAnalysisData) GetClickUserCntOk() (*int32, bool) {
	if o == nil || IsNil(o.ClickUserCnt) {
		return nil, false
	}
	return o.ClickUserCnt, true
}

// HasClickUserCnt returns a boolean if a field has been set.
func (o *MenuAnalysisData) HasClickUserCnt() bool {
	if o != nil && !IsNil(o.ClickUserCnt) {
		return true
	}

	return false
}

// SetClickUserCnt gets a reference to the given int32 and assigns it to the ClickUserCnt field.
func (o *MenuAnalysisData) SetClickUserCnt(v int32) {
	o.ClickUserCnt = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *MenuAnalysisData) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuAnalysisData) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *MenuAnalysisData) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *MenuAnalysisData) SetDate(v string) {
	o.Date = &v
}

// GetMenuType returns the MenuType field value if set, zero value otherwise.
func (o *MenuAnalysisData) GetMenuType() string {
	if o == nil || IsNil(o.MenuType) {
		var ret string
		return ret
	}
	return *o.MenuType
}

// GetMenuTypeOk returns a tuple with the MenuType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuAnalysisData) GetMenuTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MenuType) {
		return nil, false
	}
	return o.MenuType, true
}

// HasMenuType returns a boolean if a field has been set.
func (o *MenuAnalysisData) HasMenuType() bool {
	if o != nil && !IsNil(o.MenuType) {
		return true
	}

	return false
}

// SetMenuType gets a reference to the given string and assigns it to the MenuType field.
func (o *MenuAnalysisData) SetMenuType(v string) {
	o.MenuType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MenuAnalysisData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuAnalysisData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MenuAnalysisData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MenuAnalysisData) SetName(v string) {
	o.Name = &v
}

// GetSubName returns the SubName field value if set, zero value otherwise.
func (o *MenuAnalysisData) GetSubName() string {
	if o == nil || IsNil(o.SubName) {
		var ret string
		return ret
	}
	return *o.SubName
}

// GetSubNameOk returns a tuple with the SubName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuAnalysisData) GetSubNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubName) {
		return nil, false
	}
	return o.SubName, true
}

// HasSubName returns a boolean if a field has been set.
func (o *MenuAnalysisData) HasSubName() bool {
	if o != nil && !IsNil(o.SubName) {
		return true
	}

	return false
}

// SetSubName gets a reference to the given string and assigns it to the SubName field.
func (o *MenuAnalysisData) SetSubName(v string) {
	o.SubName = &v
}

func (o MenuAnalysisData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MenuAnalysisData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvgClickUserCnt) {
		toSerialize["avg_click_user_cnt"] = o.AvgClickUserCnt
	}
	if !IsNil(o.ClickCnt) {
		toSerialize["click_cnt"] = o.ClickCnt
	}
	if !IsNil(o.ClickUserCnt) {
		toSerialize["click_user_cnt"] = o.ClickUserCnt
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.MenuType) {
		toSerialize["menu_type"] = o.MenuType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SubName) {
		toSerialize["sub_name"] = o.SubName
	}
	return toSerialize, nil
}

type NullableMenuAnalysisData struct {
	value *MenuAnalysisData
	isSet bool
}

func (v NullableMenuAnalysisData) Get() *MenuAnalysisData {
	return v.value
}

func (v *NullableMenuAnalysisData) Set(val *MenuAnalysisData) {
	v.value = val
	v.isSet = true
}

func (v NullableMenuAnalysisData) IsSet() bool {
	return v.isSet
}

func (v *NullableMenuAnalysisData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMenuAnalysisData(val *MenuAnalysisData) *NullableMenuAnalysisData {
	return &NullableMenuAnalysisData{value: val, isSet: true}
}

func (v NullableMenuAnalysisData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMenuAnalysisData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
