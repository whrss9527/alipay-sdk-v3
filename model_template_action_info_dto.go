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

// checks if the TemplateActionInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateActionInfoDTO{}

// TemplateActionInfoDTO struct for TemplateActionInfoDTO
type TemplateActionInfoDTO struct {
	// 行动点业务CODE，商户自定义
	Code       *string                      `json:"code,omitempty"`
	MiniAppUrl *TemplateActionMiniAppUrlDTO `json:"mini_app_url,omitempty"`
	// 行动点展示文案
	Text *string `json:"text,omitempty"`
	// 行动点跳转链接，当url_type填\"url\"或不填时必填，支持http(s)和支付宝schema地址等
	Url *string `json:"url,omitempty"`
	// 跳转链接类型，不填则默认为url类型：  url：对应填写url参数  miniAppUrl: 对应填写mini_app_url参数，跳转至指定的支付宝小程序页面
	UrlType *string `json:"url_type,omitempty"`
}

// NewTemplateActionInfoDTO instantiates a new TemplateActionInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateActionInfoDTO() *TemplateActionInfoDTO {
	this := TemplateActionInfoDTO{}
	return &this
}

// NewTemplateActionInfoDTOWithDefaults instantiates a new TemplateActionInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateActionInfoDTOWithDefaults() *TemplateActionInfoDTO {
	this := TemplateActionInfoDTO{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TemplateActionInfoDTO) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateActionInfoDTO) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TemplateActionInfoDTO) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TemplateActionInfoDTO) SetCode(v string) {
	o.Code = &v
}

// GetMiniAppUrl returns the MiniAppUrl field value if set, zero value otherwise.
func (o *TemplateActionInfoDTO) GetMiniAppUrl() TemplateActionMiniAppUrlDTO {
	if o == nil || IsNil(o.MiniAppUrl) {
		var ret TemplateActionMiniAppUrlDTO
		return ret
	}
	return *o.MiniAppUrl
}

// GetMiniAppUrlOk returns a tuple with the MiniAppUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateActionInfoDTO) GetMiniAppUrlOk() (*TemplateActionMiniAppUrlDTO, bool) {
	if o == nil || IsNil(o.MiniAppUrl) {
		return nil, false
	}
	return o.MiniAppUrl, true
}

// HasMiniAppUrl returns a boolean if a field has been set.
func (o *TemplateActionInfoDTO) HasMiniAppUrl() bool {
	if o != nil && !IsNil(o.MiniAppUrl) {
		return true
	}

	return false
}

// SetMiniAppUrl gets a reference to the given TemplateActionMiniAppUrlDTO and assigns it to the MiniAppUrl field.
func (o *TemplateActionInfoDTO) SetMiniAppUrl(v TemplateActionMiniAppUrlDTO) {
	o.MiniAppUrl = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TemplateActionInfoDTO) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateActionInfoDTO) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TemplateActionInfoDTO) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *TemplateActionInfoDTO) SetText(v string) {
	o.Text = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TemplateActionInfoDTO) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateActionInfoDTO) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TemplateActionInfoDTO) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TemplateActionInfoDTO) SetUrl(v string) {
	o.Url = &v
}

// GetUrlType returns the UrlType field value if set, zero value otherwise.
func (o *TemplateActionInfoDTO) GetUrlType() string {
	if o == nil || IsNil(o.UrlType) {
		var ret string
		return ret
	}
	return *o.UrlType
}

// GetUrlTypeOk returns a tuple with the UrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateActionInfoDTO) GetUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UrlType) {
		return nil, false
	}
	return o.UrlType, true
}

// HasUrlType returns a boolean if a field has been set.
func (o *TemplateActionInfoDTO) HasUrlType() bool {
	if o != nil && !IsNil(o.UrlType) {
		return true
	}

	return false
}

// SetUrlType gets a reference to the given string and assigns it to the UrlType field.
func (o *TemplateActionInfoDTO) SetUrlType(v string) {
	o.UrlType = &v
}

func (o TemplateActionInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateActionInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.MiniAppUrl) {
		toSerialize["mini_app_url"] = o.MiniAppUrl
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.UrlType) {
		toSerialize["url_type"] = o.UrlType
	}
	return toSerialize, nil
}

type NullableTemplateActionInfoDTO struct {
	value *TemplateActionInfoDTO
	isSet bool
}

func (v NullableTemplateActionInfoDTO) Get() *TemplateActionInfoDTO {
	return v.value
}

func (v *NullableTemplateActionInfoDTO) Set(val *TemplateActionInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateActionInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateActionInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateActionInfoDTO(val *TemplateActionInfoDTO) *NullableTemplateActionInfoDTO {
	return &NullableTemplateActionInfoDTO{value: val, isSet: true}
}

func (v NullableTemplateActionInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateActionInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
