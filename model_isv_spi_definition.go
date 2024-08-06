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

// checks if the IsvSpiDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsvSpiDefinition{}

// IsvSpiDefinition struct for IsvSpiDefinition
type IsvSpiDefinition struct {
	// ISV自定义的标识功能的业务代码，不可重复
	BizCode *string `json:"biz_code,omitempty"`
	// spi功能描述
	Description *string `json:"description,omitempty"`
	// 功能图标文件url
	Icon *string `json:"icon,omitempty"`
	// spi接口服务地址
	SpiEndpoint *string `json:"spi_endpoint,omitempty"`
	// spi接口扩展参数，json格式字符串
	SpiExtProperty *string `json:"spi_ext_property,omitempty"`
	// CCM预先定义的spi key，与插件位置有关
	SpiKey *string `json:"spi_key,omitempty"`
	// SPI 名称
	SpiName *string `json:"spi_name,omitempty"`
}

// NewIsvSpiDefinition instantiates a new IsvSpiDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsvSpiDefinition() *IsvSpiDefinition {
	this := IsvSpiDefinition{}
	return &this
}

// NewIsvSpiDefinitionWithDefaults instantiates a new IsvSpiDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsvSpiDefinitionWithDefaults() *IsvSpiDefinition {
	this := IsvSpiDefinition{}
	return &this
}

// GetBizCode returns the BizCode field value if set, zero value otherwise.
func (o *IsvSpiDefinition) GetBizCode() string {
	if o == nil || IsNil(o.BizCode) {
		var ret string
		return ret
	}
	return *o.BizCode
}

// GetBizCodeOk returns a tuple with the BizCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvSpiDefinition) GetBizCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BizCode) {
		return nil, false
	}
	return o.BizCode, true
}

// HasBizCode returns a boolean if a field has been set.
func (o *IsvSpiDefinition) HasBizCode() bool {
	if o != nil && !IsNil(o.BizCode) {
		return true
	}

	return false
}

// SetBizCode gets a reference to the given string and assigns it to the BizCode field.
func (o *IsvSpiDefinition) SetBizCode(v string) {
	o.BizCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IsvSpiDefinition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvSpiDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IsvSpiDefinition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IsvSpiDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *IsvSpiDefinition) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvSpiDefinition) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *IsvSpiDefinition) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *IsvSpiDefinition) SetIcon(v string) {
	o.Icon = &v
}

// GetSpiEndpoint returns the SpiEndpoint field value if set, zero value otherwise.
func (o *IsvSpiDefinition) GetSpiEndpoint() string {
	if o == nil || IsNil(o.SpiEndpoint) {
		var ret string
		return ret
	}
	return *o.SpiEndpoint
}

// GetSpiEndpointOk returns a tuple with the SpiEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvSpiDefinition) GetSpiEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.SpiEndpoint) {
		return nil, false
	}
	return o.SpiEndpoint, true
}

// HasSpiEndpoint returns a boolean if a field has been set.
func (o *IsvSpiDefinition) HasSpiEndpoint() bool {
	if o != nil && !IsNil(o.SpiEndpoint) {
		return true
	}

	return false
}

// SetSpiEndpoint gets a reference to the given string and assigns it to the SpiEndpoint field.
func (o *IsvSpiDefinition) SetSpiEndpoint(v string) {
	o.SpiEndpoint = &v
}

// GetSpiExtProperty returns the SpiExtProperty field value if set, zero value otherwise.
func (o *IsvSpiDefinition) GetSpiExtProperty() string {
	if o == nil || IsNil(o.SpiExtProperty) {
		var ret string
		return ret
	}
	return *o.SpiExtProperty
}

// GetSpiExtPropertyOk returns a tuple with the SpiExtProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvSpiDefinition) GetSpiExtPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.SpiExtProperty) {
		return nil, false
	}
	return o.SpiExtProperty, true
}

// HasSpiExtProperty returns a boolean if a field has been set.
func (o *IsvSpiDefinition) HasSpiExtProperty() bool {
	if o != nil && !IsNil(o.SpiExtProperty) {
		return true
	}

	return false
}

// SetSpiExtProperty gets a reference to the given string and assigns it to the SpiExtProperty field.
func (o *IsvSpiDefinition) SetSpiExtProperty(v string) {
	o.SpiExtProperty = &v
}

// GetSpiKey returns the SpiKey field value if set, zero value otherwise.
func (o *IsvSpiDefinition) GetSpiKey() string {
	if o == nil || IsNil(o.SpiKey) {
		var ret string
		return ret
	}
	return *o.SpiKey
}

// GetSpiKeyOk returns a tuple with the SpiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvSpiDefinition) GetSpiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SpiKey) {
		return nil, false
	}
	return o.SpiKey, true
}

// HasSpiKey returns a boolean if a field has been set.
func (o *IsvSpiDefinition) HasSpiKey() bool {
	if o != nil && !IsNil(o.SpiKey) {
		return true
	}

	return false
}

// SetSpiKey gets a reference to the given string and assigns it to the SpiKey field.
func (o *IsvSpiDefinition) SetSpiKey(v string) {
	o.SpiKey = &v
}

// GetSpiName returns the SpiName field value if set, zero value otherwise.
func (o *IsvSpiDefinition) GetSpiName() string {
	if o == nil || IsNil(o.SpiName) {
		var ret string
		return ret
	}
	return *o.SpiName
}

// GetSpiNameOk returns a tuple with the SpiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvSpiDefinition) GetSpiNameOk() (*string, bool) {
	if o == nil || IsNil(o.SpiName) {
		return nil, false
	}
	return o.SpiName, true
}

// HasSpiName returns a boolean if a field has been set.
func (o *IsvSpiDefinition) HasSpiName() bool {
	if o != nil && !IsNil(o.SpiName) {
		return true
	}

	return false
}

// SetSpiName gets a reference to the given string and assigns it to the SpiName field.
func (o *IsvSpiDefinition) SetSpiName(v string) {
	o.SpiName = &v
}

func (o IsvSpiDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsvSpiDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizCode) {
		toSerialize["biz_code"] = o.BizCode
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.SpiEndpoint) {
		toSerialize["spi_endpoint"] = o.SpiEndpoint
	}
	if !IsNil(o.SpiExtProperty) {
		toSerialize["spi_ext_property"] = o.SpiExtProperty
	}
	if !IsNil(o.SpiKey) {
		toSerialize["spi_key"] = o.SpiKey
	}
	if !IsNil(o.SpiName) {
		toSerialize["spi_name"] = o.SpiName
	}
	return toSerialize, nil
}

type NullableIsvSpiDefinition struct {
	value *IsvSpiDefinition
	isSet bool
}

func (v NullableIsvSpiDefinition) Get() *IsvSpiDefinition {
	return v.value
}

func (v *NullableIsvSpiDefinition) Set(val *IsvSpiDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableIsvSpiDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableIsvSpiDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsvSpiDefinition(val *IsvSpiDefinition) *NullableIsvSpiDefinition {
	return &NullableIsvSpiDefinition{value: val, isSet: true}
}

func (v NullableIsvSpiDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsvSpiDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

