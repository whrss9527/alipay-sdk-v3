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

// checks if the ConsumeOutputInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsumeOutputInfo{}

// ConsumeOutputInfo struct for ConsumeOutputInfo
type ConsumeOutputInfo struct {
	// 支付宝交易号
	BillNo *string `json:"bill_no,omitempty"`
	// 账单分类
	CategoryName *string `json:"category_name,omitempty"`
	// 金额，单位元
	ConsumeAmount *string `json:"consume_amount,omitempty"`
	// 日期
	ConsumeDate *string `json:"consume_date,omitempty"`
	// 交易记录标题
	ConsumeTitle *string `json:"consume_title,omitempty"`
	// 商家名称
	PayeeName *string `json:"payee_name,omitempty"`
}

// NewConsumeOutputInfo instantiates a new ConsumeOutputInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumeOutputInfo() *ConsumeOutputInfo {
	this := ConsumeOutputInfo{}
	return &this
}

// NewConsumeOutputInfoWithDefaults instantiates a new ConsumeOutputInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumeOutputInfoWithDefaults() *ConsumeOutputInfo {
	this := ConsumeOutputInfo{}
	return &this
}

// GetBillNo returns the BillNo field value if set, zero value otherwise.
func (o *ConsumeOutputInfo) GetBillNo() string {
	if o == nil || IsNil(o.BillNo) {
		var ret string
		return ret
	}
	return *o.BillNo
}

// GetBillNoOk returns a tuple with the BillNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumeOutputInfo) GetBillNoOk() (*string, bool) {
	if o == nil || IsNil(o.BillNo) {
		return nil, false
	}
	return o.BillNo, true
}

// HasBillNo returns a boolean if a field has been set.
func (o *ConsumeOutputInfo) HasBillNo() bool {
	if o != nil && !IsNil(o.BillNo) {
		return true
	}

	return false
}

// SetBillNo gets a reference to the given string and assigns it to the BillNo field.
func (o *ConsumeOutputInfo) SetBillNo(v string) {
	o.BillNo = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *ConsumeOutputInfo) GetCategoryName() string {
	if o == nil || IsNil(o.CategoryName) {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumeOutputInfo) GetCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryName) {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *ConsumeOutputInfo) HasCategoryName() bool {
	if o != nil && !IsNil(o.CategoryName) {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *ConsumeOutputInfo) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetConsumeAmount returns the ConsumeAmount field value if set, zero value otherwise.
func (o *ConsumeOutputInfo) GetConsumeAmount() string {
	if o == nil || IsNil(o.ConsumeAmount) {
		var ret string
		return ret
	}
	return *o.ConsumeAmount
}

// GetConsumeAmountOk returns a tuple with the ConsumeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumeOutputInfo) GetConsumeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumeAmount) {
		return nil, false
	}
	return o.ConsumeAmount, true
}

// HasConsumeAmount returns a boolean if a field has been set.
func (o *ConsumeOutputInfo) HasConsumeAmount() bool {
	if o != nil && !IsNil(o.ConsumeAmount) {
		return true
	}

	return false
}

// SetConsumeAmount gets a reference to the given string and assigns it to the ConsumeAmount field.
func (o *ConsumeOutputInfo) SetConsumeAmount(v string) {
	o.ConsumeAmount = &v
}

// GetConsumeDate returns the ConsumeDate field value if set, zero value otherwise.
func (o *ConsumeOutputInfo) GetConsumeDate() string {
	if o == nil || IsNil(o.ConsumeDate) {
		var ret string
		return ret
	}
	return *o.ConsumeDate
}

// GetConsumeDateOk returns a tuple with the ConsumeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumeOutputInfo) GetConsumeDateOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumeDate) {
		return nil, false
	}
	return o.ConsumeDate, true
}

// HasConsumeDate returns a boolean if a field has been set.
func (o *ConsumeOutputInfo) HasConsumeDate() bool {
	if o != nil && !IsNil(o.ConsumeDate) {
		return true
	}

	return false
}

// SetConsumeDate gets a reference to the given string and assigns it to the ConsumeDate field.
func (o *ConsumeOutputInfo) SetConsumeDate(v string) {
	o.ConsumeDate = &v
}

// GetConsumeTitle returns the ConsumeTitle field value if set, zero value otherwise.
func (o *ConsumeOutputInfo) GetConsumeTitle() string {
	if o == nil || IsNil(o.ConsumeTitle) {
		var ret string
		return ret
	}
	return *o.ConsumeTitle
}

// GetConsumeTitleOk returns a tuple with the ConsumeTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumeOutputInfo) GetConsumeTitleOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumeTitle) {
		return nil, false
	}
	return o.ConsumeTitle, true
}

// HasConsumeTitle returns a boolean if a field has been set.
func (o *ConsumeOutputInfo) HasConsumeTitle() bool {
	if o != nil && !IsNil(o.ConsumeTitle) {
		return true
	}

	return false
}

// SetConsumeTitle gets a reference to the given string and assigns it to the ConsumeTitle field.
func (o *ConsumeOutputInfo) SetConsumeTitle(v string) {
	o.ConsumeTitle = &v
}

// GetPayeeName returns the PayeeName field value if set, zero value otherwise.
func (o *ConsumeOutputInfo) GetPayeeName() string {
	if o == nil || IsNil(o.PayeeName) {
		var ret string
		return ret
	}
	return *o.PayeeName
}

// GetPayeeNameOk returns a tuple with the PayeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumeOutputInfo) GetPayeeNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeName) {
		return nil, false
	}
	return o.PayeeName, true
}

// HasPayeeName returns a boolean if a field has been set.
func (o *ConsumeOutputInfo) HasPayeeName() bool {
	if o != nil && !IsNil(o.PayeeName) {
		return true
	}

	return false
}

// SetPayeeName gets a reference to the given string and assigns it to the PayeeName field.
func (o *ConsumeOutputInfo) SetPayeeName(v string) {
	o.PayeeName = &v
}

func (o ConsumeOutputInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsumeOutputInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillNo) {
		toSerialize["bill_no"] = o.BillNo
	}
	if !IsNil(o.CategoryName) {
		toSerialize["category_name"] = o.CategoryName
	}
	if !IsNil(o.ConsumeAmount) {
		toSerialize["consume_amount"] = o.ConsumeAmount
	}
	if !IsNil(o.ConsumeDate) {
		toSerialize["consume_date"] = o.ConsumeDate
	}
	if !IsNil(o.ConsumeTitle) {
		toSerialize["consume_title"] = o.ConsumeTitle
	}
	if !IsNil(o.PayeeName) {
		toSerialize["payee_name"] = o.PayeeName
	}
	return toSerialize, nil
}

type NullableConsumeOutputInfo struct {
	value *ConsumeOutputInfo
	isSet bool
}

func (v NullableConsumeOutputInfo) Get() *ConsumeOutputInfo {
	return v.value
}

func (v *NullableConsumeOutputInfo) Set(val *ConsumeOutputInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumeOutputInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumeOutputInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumeOutputInfo(val *ConsumeOutputInfo) *NullableConsumeOutputInfo {
	return &NullableConsumeOutputInfo{value: val, isSet: true}
}

func (v NullableConsumeOutputInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumeOutputInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
