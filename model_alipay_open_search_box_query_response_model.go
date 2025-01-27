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

// checks if the AlipayOpenSearchBoxQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchBoxQueryResponseModel{}

// AlipayOpenSearchBoxQueryResponseModel struct for AlipayOpenSearchBoxQueryResponseModel
type AlipayOpenSearchBoxQueryResponseModel struct {
	AccountModule     *SearchBoxAccountModule     `json:"account_module,omitempty"`
	AreaKeywordModule *SearchBoxAreaKeyWordModule `json:"area_keyword_module,omitempty"`
	BasicInfoModule   *SearchBoxBasicInfoModule   `json:"basic_info_module,omitempty"`
	// 搜索直达配置id
	BoxId *string `json:"box_id,omitempty"`
	// 搜索直达配置状态，AUDIT-审核中/ONLINE-已上架/REJECT-驳回/OFFLINE-已下架
	BoxStatus              *string                    `json:"box_status,omitempty"`
	BusinessDistrictModule *BoxBusinessDistrictModule `json:"business_district_module,omitempty"`
	// 搜索直达默认触发词，由系统生成，无法修改
	DefaultKeywords  []string                `json:"default_keywords,omitempty"`
	KeywordModule    *SearchBoxKeyWordModule `json:"keyword_module,omitempty"`
	LatestAuditImage *SearchBoxImageModule   `json:"latest_audit_image,omitempty"`
	ServiceModule    *SearchBoxServiceModule `json:"service_module,omitempty"`
	ValidImage       *SearchBoxImageModule   `json:"valid_image,omitempty"`
}

// NewAlipayOpenSearchBoxQueryResponseModel instantiates a new AlipayOpenSearchBoxQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchBoxQueryResponseModel() *AlipayOpenSearchBoxQueryResponseModel {
	this := AlipayOpenSearchBoxQueryResponseModel{}
	return &this
}

// NewAlipayOpenSearchBoxQueryResponseModelWithDefaults instantiates a new AlipayOpenSearchBoxQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchBoxQueryResponseModelWithDefaults() *AlipayOpenSearchBoxQueryResponseModel {
	this := AlipayOpenSearchBoxQueryResponseModel{}
	return &this
}

// GetAccountModule returns the AccountModule field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetAccountModule() SearchBoxAccountModule {
	if o == nil || IsNil(o.AccountModule) {
		var ret SearchBoxAccountModule
		return ret
	}
	return *o.AccountModule
}

// GetAccountModuleOk returns a tuple with the AccountModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetAccountModuleOk() (*SearchBoxAccountModule, bool) {
	if o == nil || IsNil(o.AccountModule) {
		return nil, false
	}
	return o.AccountModule, true
}

// HasAccountModule returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasAccountModule() bool {
	if o != nil && !IsNil(o.AccountModule) {
		return true
	}

	return false
}

// SetAccountModule gets a reference to the given SearchBoxAccountModule and assigns it to the AccountModule field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetAccountModule(v SearchBoxAccountModule) {
	o.AccountModule = &v
}

// GetAreaKeywordModule returns the AreaKeywordModule field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetAreaKeywordModule() SearchBoxAreaKeyWordModule {
	if o == nil || IsNil(o.AreaKeywordModule) {
		var ret SearchBoxAreaKeyWordModule
		return ret
	}
	return *o.AreaKeywordModule
}

// GetAreaKeywordModuleOk returns a tuple with the AreaKeywordModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetAreaKeywordModuleOk() (*SearchBoxAreaKeyWordModule, bool) {
	if o == nil || IsNil(o.AreaKeywordModule) {
		return nil, false
	}
	return o.AreaKeywordModule, true
}

// HasAreaKeywordModule returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasAreaKeywordModule() bool {
	if o != nil && !IsNil(o.AreaKeywordModule) {
		return true
	}

	return false
}

// SetAreaKeywordModule gets a reference to the given SearchBoxAreaKeyWordModule and assigns it to the AreaKeywordModule field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetAreaKeywordModule(v SearchBoxAreaKeyWordModule) {
	o.AreaKeywordModule = &v
}

// GetBasicInfoModule returns the BasicInfoModule field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetBasicInfoModule() SearchBoxBasicInfoModule {
	if o == nil || IsNil(o.BasicInfoModule) {
		var ret SearchBoxBasicInfoModule
		return ret
	}
	return *o.BasicInfoModule
}

// GetBasicInfoModuleOk returns a tuple with the BasicInfoModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetBasicInfoModuleOk() (*SearchBoxBasicInfoModule, bool) {
	if o == nil || IsNil(o.BasicInfoModule) {
		return nil, false
	}
	return o.BasicInfoModule, true
}

// HasBasicInfoModule returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasBasicInfoModule() bool {
	if o != nil && !IsNil(o.BasicInfoModule) {
		return true
	}

	return false
}

// SetBasicInfoModule gets a reference to the given SearchBoxBasicInfoModule and assigns it to the BasicInfoModule field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetBasicInfoModule(v SearchBoxBasicInfoModule) {
	o.BasicInfoModule = &v
}

// GetBoxId returns the BoxId field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetBoxId() string {
	if o == nil || IsNil(o.BoxId) {
		var ret string
		return ret
	}
	return *o.BoxId
}

// GetBoxIdOk returns a tuple with the BoxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetBoxIdOk() (*string, bool) {
	if o == nil || IsNil(o.BoxId) {
		return nil, false
	}
	return o.BoxId, true
}

// HasBoxId returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasBoxId() bool {
	if o != nil && !IsNil(o.BoxId) {
		return true
	}

	return false
}

// SetBoxId gets a reference to the given string and assigns it to the BoxId field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetBoxId(v string) {
	o.BoxId = &v
}

// GetBoxStatus returns the BoxStatus field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetBoxStatus() string {
	if o == nil || IsNil(o.BoxStatus) {
		var ret string
		return ret
	}
	return *o.BoxStatus
}

// GetBoxStatusOk returns a tuple with the BoxStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetBoxStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BoxStatus) {
		return nil, false
	}
	return o.BoxStatus, true
}

// HasBoxStatus returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasBoxStatus() bool {
	if o != nil && !IsNil(o.BoxStatus) {
		return true
	}

	return false
}

// SetBoxStatus gets a reference to the given string and assigns it to the BoxStatus field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetBoxStatus(v string) {
	o.BoxStatus = &v
}

// GetBusinessDistrictModule returns the BusinessDistrictModule field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetBusinessDistrictModule() BoxBusinessDistrictModule {
	if o == nil || IsNil(o.BusinessDistrictModule) {
		var ret BoxBusinessDistrictModule
		return ret
	}
	return *o.BusinessDistrictModule
}

// GetBusinessDistrictModuleOk returns a tuple with the BusinessDistrictModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetBusinessDistrictModuleOk() (*BoxBusinessDistrictModule, bool) {
	if o == nil || IsNil(o.BusinessDistrictModule) {
		return nil, false
	}
	return o.BusinessDistrictModule, true
}

// HasBusinessDistrictModule returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasBusinessDistrictModule() bool {
	if o != nil && !IsNil(o.BusinessDistrictModule) {
		return true
	}

	return false
}

// SetBusinessDistrictModule gets a reference to the given BoxBusinessDistrictModule and assigns it to the BusinessDistrictModule field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetBusinessDistrictModule(v BoxBusinessDistrictModule) {
	o.BusinessDistrictModule = &v
}

// GetDefaultKeywords returns the DefaultKeywords field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetDefaultKeywords() []string {
	if o == nil || IsNil(o.DefaultKeywords) {
		var ret []string
		return ret
	}
	return o.DefaultKeywords
}

// GetDefaultKeywordsOk returns a tuple with the DefaultKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetDefaultKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.DefaultKeywords) {
		return nil, false
	}
	return o.DefaultKeywords, true
}

// HasDefaultKeywords returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasDefaultKeywords() bool {
	if o != nil && !IsNil(o.DefaultKeywords) {
		return true
	}

	return false
}

// SetDefaultKeywords gets a reference to the given []string and assigns it to the DefaultKeywords field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetDefaultKeywords(v []string) {
	o.DefaultKeywords = v
}

// GetKeywordModule returns the KeywordModule field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetKeywordModule() SearchBoxKeyWordModule {
	if o == nil || IsNil(o.KeywordModule) {
		var ret SearchBoxKeyWordModule
		return ret
	}
	return *o.KeywordModule
}

// GetKeywordModuleOk returns a tuple with the KeywordModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetKeywordModuleOk() (*SearchBoxKeyWordModule, bool) {
	if o == nil || IsNil(o.KeywordModule) {
		return nil, false
	}
	return o.KeywordModule, true
}

// HasKeywordModule returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasKeywordModule() bool {
	if o != nil && !IsNil(o.KeywordModule) {
		return true
	}

	return false
}

// SetKeywordModule gets a reference to the given SearchBoxKeyWordModule and assigns it to the KeywordModule field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetKeywordModule(v SearchBoxKeyWordModule) {
	o.KeywordModule = &v
}

// GetLatestAuditImage returns the LatestAuditImage field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetLatestAuditImage() SearchBoxImageModule {
	if o == nil || IsNil(o.LatestAuditImage) {
		var ret SearchBoxImageModule
		return ret
	}
	return *o.LatestAuditImage
}

// GetLatestAuditImageOk returns a tuple with the LatestAuditImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetLatestAuditImageOk() (*SearchBoxImageModule, bool) {
	if o == nil || IsNil(o.LatestAuditImage) {
		return nil, false
	}
	return o.LatestAuditImage, true
}

// HasLatestAuditImage returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasLatestAuditImage() bool {
	if o != nil && !IsNil(o.LatestAuditImage) {
		return true
	}

	return false
}

// SetLatestAuditImage gets a reference to the given SearchBoxImageModule and assigns it to the LatestAuditImage field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetLatestAuditImage(v SearchBoxImageModule) {
	o.LatestAuditImage = &v
}

// GetServiceModule returns the ServiceModule field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetServiceModule() SearchBoxServiceModule {
	if o == nil || IsNil(o.ServiceModule) {
		var ret SearchBoxServiceModule
		return ret
	}
	return *o.ServiceModule
}

// GetServiceModuleOk returns a tuple with the ServiceModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetServiceModuleOk() (*SearchBoxServiceModule, bool) {
	if o == nil || IsNil(o.ServiceModule) {
		return nil, false
	}
	return o.ServiceModule, true
}

// HasServiceModule returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasServiceModule() bool {
	if o != nil && !IsNil(o.ServiceModule) {
		return true
	}

	return false
}

// SetServiceModule gets a reference to the given SearchBoxServiceModule and assigns it to the ServiceModule field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetServiceModule(v SearchBoxServiceModule) {
	o.ServiceModule = &v
}

// GetValidImage returns the ValidImage field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetValidImage() SearchBoxImageModule {
	if o == nil || IsNil(o.ValidImage) {
		var ret SearchBoxImageModule
		return ret
	}
	return *o.ValidImage
}

// GetValidImageOk returns a tuple with the ValidImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) GetValidImageOk() (*SearchBoxImageModule, bool) {
	if o == nil || IsNil(o.ValidImage) {
		return nil, false
	}
	return o.ValidImage, true
}

// HasValidImage returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxQueryResponseModel) HasValidImage() bool {
	if o != nil && !IsNil(o.ValidImage) {
		return true
	}

	return false
}

// SetValidImage gets a reference to the given SearchBoxImageModule and assigns it to the ValidImage field.
func (o *AlipayOpenSearchBoxQueryResponseModel) SetValidImage(v SearchBoxImageModule) {
	o.ValidImage = &v
}

func (o AlipayOpenSearchBoxQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchBoxQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountModule) {
		toSerialize["account_module"] = o.AccountModule
	}
	if !IsNil(o.AreaKeywordModule) {
		toSerialize["area_keyword_module"] = o.AreaKeywordModule
	}
	if !IsNil(o.BasicInfoModule) {
		toSerialize["basic_info_module"] = o.BasicInfoModule
	}
	if !IsNil(o.BoxId) {
		toSerialize["box_id"] = o.BoxId
	}
	if !IsNil(o.BoxStatus) {
		toSerialize["box_status"] = o.BoxStatus
	}
	if !IsNil(o.BusinessDistrictModule) {
		toSerialize["business_district_module"] = o.BusinessDistrictModule
	}
	if !IsNil(o.DefaultKeywords) {
		toSerialize["default_keywords"] = o.DefaultKeywords
	}
	if !IsNil(o.KeywordModule) {
		toSerialize["keyword_module"] = o.KeywordModule
	}
	if !IsNil(o.LatestAuditImage) {
		toSerialize["latest_audit_image"] = o.LatestAuditImage
	}
	if !IsNil(o.ServiceModule) {
		toSerialize["service_module"] = o.ServiceModule
	}
	if !IsNil(o.ValidImage) {
		toSerialize["valid_image"] = o.ValidImage
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchBoxQueryResponseModel struct {
	value *AlipayOpenSearchBoxQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenSearchBoxQueryResponseModel) Get() *AlipayOpenSearchBoxQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSearchBoxQueryResponseModel) Set(val *AlipayOpenSearchBoxQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchBoxQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchBoxQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchBoxQueryResponseModel(val *AlipayOpenSearchBoxQueryResponseModel) *NullableAlipayOpenSearchBoxQueryResponseModel {
	return &NullableAlipayOpenSearchBoxQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchBoxQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchBoxQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
