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

// checks if the SpuInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpuInfo{}

// SpuInfo struct for SpuInfo
type SpuInfo struct {
	// 品牌名称
	Brand *string `json:"brand,omitempty"`
	// 类目
	Category *string `json:"category,omitempty"`
	// 商品数量
	Count *int32 `json:"count,omitempty"`
	// 图片链接
	Icon *string `json:"icon,omitempty"`
	// 商品单价(单位:元)
	Price *string `json:"price,omitempty"`
	// 商品提供方，店铺或品牌方
	Provider *string `json:"provider,omitempty"`
	// 商品ID
	SpuId *string `json:"spu_id,omitempty"`
}

// NewSpuInfo instantiates a new SpuInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpuInfo() *SpuInfo {
	this := SpuInfo{}
	return &this
}

// NewSpuInfoWithDefaults instantiates a new SpuInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpuInfoWithDefaults() *SpuInfo {
	this := SpuInfo{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *SpuInfo) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpuInfo) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *SpuInfo) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *SpuInfo) SetBrand(v string) {
	o.Brand = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SpuInfo) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpuInfo) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SpuInfo) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SpuInfo) SetCategory(v string) {
	o.Category = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SpuInfo) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpuInfo) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SpuInfo) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SpuInfo) SetCount(v int32) {
	o.Count = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *SpuInfo) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpuInfo) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *SpuInfo) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *SpuInfo) SetIcon(v string) {
	o.Icon = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SpuInfo) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpuInfo) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SpuInfo) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SpuInfo) SetPrice(v string) {
	o.Price = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *SpuInfo) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpuInfo) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *SpuInfo) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *SpuInfo) SetProvider(v string) {
	o.Provider = &v
}

// GetSpuId returns the SpuId field value if set, zero value otherwise.
func (o *SpuInfo) GetSpuId() string {
	if o == nil || IsNil(o.SpuId) {
		var ret string
		return ret
	}
	return *o.SpuId
}

// GetSpuIdOk returns a tuple with the SpuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpuInfo) GetSpuIdOk() (*string, bool) {
	if o == nil || IsNil(o.SpuId) {
		return nil, false
	}
	return o.SpuId, true
}

// HasSpuId returns a boolean if a field has been set.
func (o *SpuInfo) HasSpuId() bool {
	if o != nil && !IsNil(o.SpuId) {
		return true
	}

	return false
}

// SetSpuId gets a reference to the given string and assigns it to the SpuId field.
func (o *SpuInfo) SetSpuId(v string) {
	o.SpuId = &v
}

func (o SpuInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpuInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.SpuId) {
		toSerialize["spu_id"] = o.SpuId
	}
	return toSerialize, nil
}

type NullableSpuInfo struct {
	value *SpuInfo
	isSet bool
}

func (v NullableSpuInfo) Get() *SpuInfo {
	return v.value
}

func (v *NullableSpuInfo) Set(val *SpuInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSpuInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSpuInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpuInfo(val *SpuInfo) *NullableSpuInfo {
	return &NullableSpuInfo{value: val, isSet: true}
}

func (v NullableSpuInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpuInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
