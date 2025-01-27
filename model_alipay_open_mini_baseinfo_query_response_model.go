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

// checks if the AlipayOpenMiniBaseinfoQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniBaseinfoQueryResponseModel{}

// AlipayOpenMiniBaseinfoQueryResponseModel struct for AlipayOpenMiniBaseinfoQueryResponseModel
type AlipayOpenMiniBaseinfoQueryResponseModel struct {
	// 小程序应用描述
	AppDesc *string `json:"app_desc,omitempty"`
	// 小程序应用英文名称
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// 小程序应用logo图标
	AppLogo *string `json:"app_logo,omitempty"`
	// 小程序应用名称
	AppName *string `json:"app_name,omitempty"`
	// 小程序应用简介，一句话描述小程序功能
	AppSlogan *string `json:"app_slogan,omitempty"`
	// 类目名称，格式为第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目;
	CategoryNames *string `json:"category_names,omitempty"`
	// 功能包名称
	PackageNames []string `json:"package_names,omitempty"`
	// 域白名单
	SafeDomains []string `json:"safe_domains,omitempty"`
	// 小程序客服邮箱
	ServiceEmail *string `json:"service_email,omitempty"`
	// 小程序客服电话
	ServicePhone *string `json:"service_phone,omitempty"`
}

// NewAlipayOpenMiniBaseinfoQueryResponseModel instantiates a new AlipayOpenMiniBaseinfoQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniBaseinfoQueryResponseModel() *AlipayOpenMiniBaseinfoQueryResponseModel {
	this := AlipayOpenMiniBaseinfoQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniBaseinfoQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniBaseinfoQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniBaseinfoQueryResponseModelWithDefaults() *AlipayOpenMiniBaseinfoQueryResponseModel {
	this := AlipayOpenMiniBaseinfoQueryResponseModel{}
	return &this
}

// GetAppDesc returns the AppDesc field value if set, zero value otherwise.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetAppDesc() string {
	if o == nil || IsNil(o.AppDesc) {
		var ret string
		return ret
	}
	return *o.AppDesc
}

// GetAppDescOk returns a tuple with the AppDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetAppDescOk() (*string, bool) {
	if o == nil || IsNil(o.AppDesc) {
		return nil, false
	}
	return o.AppDesc, true
}

// HasAppDesc returns a boolean if a field has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) HasAppDesc() bool {
	if o != nil && !IsNil(o.AppDesc) {
		return true
	}

	return false
}

// SetAppDesc gets a reference to the given string and assigns it to the AppDesc field.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) SetAppDesc(v string) {
	o.AppDesc = &v
}

// GetAppEnglishName returns the AppEnglishName field value if set, zero value otherwise.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetAppEnglishName() string {
	if o == nil || IsNil(o.AppEnglishName) {
		var ret string
		return ret
	}
	return *o.AppEnglishName
}

// GetAppEnglishNameOk returns a tuple with the AppEnglishName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetAppEnglishNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppEnglishName) {
		return nil, false
	}
	return o.AppEnglishName, true
}

// HasAppEnglishName returns a boolean if a field has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) HasAppEnglishName() bool {
	if o != nil && !IsNil(o.AppEnglishName) {
		return true
	}

	return false
}

// SetAppEnglishName gets a reference to the given string and assigns it to the AppEnglishName field.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) SetAppEnglishName(v string) {
	o.AppEnglishName = &v
}

// GetAppLogo returns the AppLogo field value if set, zero value otherwise.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetAppLogo() string {
	if o == nil || IsNil(o.AppLogo) {
		var ret string
		return ret
	}
	return *o.AppLogo
}

// GetAppLogoOk returns a tuple with the AppLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetAppLogoOk() (*string, bool) {
	if o == nil || IsNil(o.AppLogo) {
		return nil, false
	}
	return o.AppLogo, true
}

// HasAppLogo returns a boolean if a field has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) HasAppLogo() bool {
	if o != nil && !IsNil(o.AppLogo) {
		return true
	}

	return false
}

// SetAppLogo gets a reference to the given string and assigns it to the AppLogo field.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) SetAppLogo(v string) {
	o.AppLogo = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) SetAppName(v string) {
	o.AppName = &v
}

// GetAppSlogan returns the AppSlogan field value if set, zero value otherwise.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetAppSlogan() string {
	if o == nil || IsNil(o.AppSlogan) {
		var ret string
		return ret
	}
	return *o.AppSlogan
}

// GetAppSloganOk returns a tuple with the AppSlogan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetAppSloganOk() (*string, bool) {
	if o == nil || IsNil(o.AppSlogan) {
		return nil, false
	}
	return o.AppSlogan, true
}

// HasAppSlogan returns a boolean if a field has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) HasAppSlogan() bool {
	if o != nil && !IsNil(o.AppSlogan) {
		return true
	}

	return false
}

// SetAppSlogan gets a reference to the given string and assigns it to the AppSlogan field.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) SetAppSlogan(v string) {
	o.AppSlogan = &v
}

// GetCategoryNames returns the CategoryNames field value if set, zero value otherwise.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetCategoryNames() string {
	if o == nil || IsNil(o.CategoryNames) {
		var ret string
		return ret
	}
	return *o.CategoryNames
}

// GetCategoryNamesOk returns a tuple with the CategoryNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetCategoryNamesOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryNames) {
		return nil, false
	}
	return o.CategoryNames, true
}

// HasCategoryNames returns a boolean if a field has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) HasCategoryNames() bool {
	if o != nil && !IsNil(o.CategoryNames) {
		return true
	}

	return false
}

// SetCategoryNames gets a reference to the given string and assigns it to the CategoryNames field.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) SetCategoryNames(v string) {
	o.CategoryNames = &v
}

// GetPackageNames returns the PackageNames field value if set, zero value otherwise.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetPackageNames() []string {
	if o == nil || IsNil(o.PackageNames) {
		var ret []string
		return ret
	}
	return o.PackageNames
}

// GetPackageNamesOk returns a tuple with the PackageNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetPackageNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.PackageNames) {
		return nil, false
	}
	return o.PackageNames, true
}

// HasPackageNames returns a boolean if a field has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) HasPackageNames() bool {
	if o != nil && !IsNil(o.PackageNames) {
		return true
	}

	return false
}

// SetPackageNames gets a reference to the given []string and assigns it to the PackageNames field.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) SetPackageNames(v []string) {
	o.PackageNames = v
}

// GetSafeDomains returns the SafeDomains field value if set, zero value otherwise.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetSafeDomains() []string {
	if o == nil || IsNil(o.SafeDomains) {
		var ret []string
		return ret
	}
	return o.SafeDomains
}

// GetSafeDomainsOk returns a tuple with the SafeDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetSafeDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.SafeDomains) {
		return nil, false
	}
	return o.SafeDomains, true
}

// HasSafeDomains returns a boolean if a field has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) HasSafeDomains() bool {
	if o != nil && !IsNil(o.SafeDomains) {
		return true
	}

	return false
}

// SetSafeDomains gets a reference to the given []string and assigns it to the SafeDomains field.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) SetSafeDomains(v []string) {
	o.SafeDomains = v
}

// GetServiceEmail returns the ServiceEmail field value if set, zero value otherwise.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetServiceEmail() string {
	if o == nil || IsNil(o.ServiceEmail) {
		var ret string
		return ret
	}
	return *o.ServiceEmail
}

// GetServiceEmailOk returns a tuple with the ServiceEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetServiceEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceEmail) {
		return nil, false
	}
	return o.ServiceEmail, true
}

// HasServiceEmail returns a boolean if a field has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) HasServiceEmail() bool {
	if o != nil && !IsNil(o.ServiceEmail) {
		return true
	}

	return false
}

// SetServiceEmail gets a reference to the given string and assigns it to the ServiceEmail field.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) SetServiceEmail(v string) {
	o.ServiceEmail = &v
}

// GetServicePhone returns the ServicePhone field value if set, zero value otherwise.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetServicePhone() string {
	if o == nil || IsNil(o.ServicePhone) {
		var ret string
		return ret
	}
	return *o.ServicePhone
}

// GetServicePhoneOk returns a tuple with the ServicePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) GetServicePhoneOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePhone) {
		return nil, false
	}
	return o.ServicePhone, true
}

// HasServicePhone returns a boolean if a field has been set.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) HasServicePhone() bool {
	if o != nil && !IsNil(o.ServicePhone) {
		return true
	}

	return false
}

// SetServicePhone gets a reference to the given string and assigns it to the ServicePhone field.
func (o *AlipayOpenMiniBaseinfoQueryResponseModel) SetServicePhone(v string) {
	o.ServicePhone = &v
}

func (o AlipayOpenMiniBaseinfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniBaseinfoQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppDesc) {
		toSerialize["app_desc"] = o.AppDesc
	}
	if !IsNil(o.AppEnglishName) {
		toSerialize["app_english_name"] = o.AppEnglishName
	}
	if !IsNil(o.AppLogo) {
		toSerialize["app_logo"] = o.AppLogo
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AppSlogan) {
		toSerialize["app_slogan"] = o.AppSlogan
	}
	if !IsNil(o.CategoryNames) {
		toSerialize["category_names"] = o.CategoryNames
	}
	if !IsNil(o.PackageNames) {
		toSerialize["package_names"] = o.PackageNames
	}
	if !IsNil(o.SafeDomains) {
		toSerialize["safe_domains"] = o.SafeDomains
	}
	if !IsNil(o.ServiceEmail) {
		toSerialize["service_email"] = o.ServiceEmail
	}
	if !IsNil(o.ServicePhone) {
		toSerialize["service_phone"] = o.ServicePhone
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniBaseinfoQueryResponseModel struct {
	value *AlipayOpenMiniBaseinfoQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniBaseinfoQueryResponseModel) Get() *AlipayOpenMiniBaseinfoQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniBaseinfoQueryResponseModel) Set(val *AlipayOpenMiniBaseinfoQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniBaseinfoQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniBaseinfoQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniBaseinfoQueryResponseModel(val *AlipayOpenMiniBaseinfoQueryResponseModel) *NullableAlipayOpenMiniBaseinfoQueryResponseModel {
	return &NullableAlipayOpenMiniBaseinfoQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniBaseinfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniBaseinfoQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
