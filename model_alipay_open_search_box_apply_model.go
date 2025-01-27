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

// checks if the AlipayOpenSearchBoxApplyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchBoxApplyModel{}

// AlipayOpenSearchBoxApplyModel struct for AlipayOpenSearchBoxApplyModel
type AlipayOpenSearchBoxApplyModel struct {
	// 小程序直达配置的常用服务中带有门店信息时，可添加简称触发词
	AreaKeywords []string `json:"area_keywords,omitempty"`
	// 品牌介绍，5-15个中文字符。 小程序直达不支持设置此项
	BoxDesc *string `json:"box_desc,omitempty"`
	// 品牌id，参考<a href=\"https://opendocs.alipay.com/rules/029uy4\"> 品牌认证说明 </a>
	BrandId *string `json:"brand_id,omitempty"`
	// 可通过配置来开启商圈权益模块，关闭后搜索直达不展示商圈权益模块
	BusinessBenefitSwitch *bool `json:"business_benefit_switch,omitempty"`
	// 小程序已关联商圈时，可添加商圈id（目前仅对品牌直达开放，小程序直达暂未开放）
	BusinessDistrictIds []string `json:"business_district_ids,omitempty"`
	// 自定义触发词，最多可配置6个，限1-8个中文字符。 小程序直达不支持设置此项
	CustomKeywords []string `json:"custom_keywords,omitempty"`
	// 氛围图片id，调用 <a href=\"https://opendocs.alipay.com/mini/03hvl1?ref=api\">支付宝文件上传接口</a> 上传图片获取图片id(bizCode：search_box_atmosphere)。 <a href=\"https://opendocs.alipay.com/b/03al6f\">图片规范</a>  小程序直达不支持设置此项。
	ImageId *string `json:"image_id,omitempty"`
	// 氛围图片名。 小程序直达不支持设置此项
	ImageName *string `json:"image_name,omitempty"`
	// 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家\"运营支付宝小程序\"授权。
	MerchantId *string `json:"merchant_id,omitempty"`
	// 关联账号信息，1-3个。 内部字段均需设置。当为品牌直达时，数组中的第1个账号会被设置为\"账号1\"，也就是将作为搜索直达专区头部的跳转地址
	RelatedAccounts []SearchBoxAppInfo `json:"related_accounts,omitempty"`
	// \"服务信息，服务必须审核通过才能申请搜索直达。品牌直达最多可配置同一品牌认证下的小程序4个，小程序直达最多可配置2个。 内部字段均需设置。\"
	ServiceInfos []SearchBoxServiceInfo `json:"service_infos,omitempty"`
	// 小程序id，小程序直达时必传，需要和申请的商户主体保持一致，且符合<a href=\"https://opendocs.alipay.com/b/03al6e\"> 准入类目 </a>
	TargetAppid *string `json:"target_appid,omitempty"`
}

// NewAlipayOpenSearchBoxApplyModel instantiates a new AlipayOpenSearchBoxApplyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchBoxApplyModel() *AlipayOpenSearchBoxApplyModel {
	this := AlipayOpenSearchBoxApplyModel{}
	return &this
}

// NewAlipayOpenSearchBoxApplyModelWithDefaults instantiates a new AlipayOpenSearchBoxApplyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchBoxApplyModelWithDefaults() *AlipayOpenSearchBoxApplyModel {
	this := AlipayOpenSearchBoxApplyModel{}
	return &this
}

// GetAreaKeywords returns the AreaKeywords field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetAreaKeywords() []string {
	if o == nil || IsNil(o.AreaKeywords) {
		var ret []string
		return ret
	}
	return o.AreaKeywords
}

// GetAreaKeywordsOk returns a tuple with the AreaKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetAreaKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.AreaKeywords) {
		return nil, false
	}
	return o.AreaKeywords, true
}

// HasAreaKeywords returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasAreaKeywords() bool {
	if o != nil && !IsNil(o.AreaKeywords) {
		return true
	}

	return false
}

// SetAreaKeywords gets a reference to the given []string and assigns it to the AreaKeywords field.
func (o *AlipayOpenSearchBoxApplyModel) SetAreaKeywords(v []string) {
	o.AreaKeywords = v
}

// GetBoxDesc returns the BoxDesc field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetBoxDesc() string {
	if o == nil || IsNil(o.BoxDesc) {
		var ret string
		return ret
	}
	return *o.BoxDesc
}

// GetBoxDescOk returns a tuple with the BoxDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetBoxDescOk() (*string, bool) {
	if o == nil || IsNil(o.BoxDesc) {
		return nil, false
	}
	return o.BoxDesc, true
}

// HasBoxDesc returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasBoxDesc() bool {
	if o != nil && !IsNil(o.BoxDesc) {
		return true
	}

	return false
}

// SetBoxDesc gets a reference to the given string and assigns it to the BoxDesc field.
func (o *AlipayOpenSearchBoxApplyModel) SetBoxDesc(v string) {
	o.BoxDesc = &v
}

// GetBrandId returns the BrandId field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetBrandId() string {
	if o == nil || IsNil(o.BrandId) {
		var ret string
		return ret
	}
	return *o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetBrandIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandId) {
		return nil, false
	}
	return o.BrandId, true
}

// HasBrandId returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasBrandId() bool {
	if o != nil && !IsNil(o.BrandId) {
		return true
	}

	return false
}

// SetBrandId gets a reference to the given string and assigns it to the BrandId field.
func (o *AlipayOpenSearchBoxApplyModel) SetBrandId(v string) {
	o.BrandId = &v
}

// GetBusinessBenefitSwitch returns the BusinessBenefitSwitch field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetBusinessBenefitSwitch() bool {
	if o == nil || IsNil(o.BusinessBenefitSwitch) {
		var ret bool
		return ret
	}
	return *o.BusinessBenefitSwitch
}

// GetBusinessBenefitSwitchOk returns a tuple with the BusinessBenefitSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetBusinessBenefitSwitchOk() (*bool, bool) {
	if o == nil || IsNil(o.BusinessBenefitSwitch) {
		return nil, false
	}
	return o.BusinessBenefitSwitch, true
}

// HasBusinessBenefitSwitch returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasBusinessBenefitSwitch() bool {
	if o != nil && !IsNil(o.BusinessBenefitSwitch) {
		return true
	}

	return false
}

// SetBusinessBenefitSwitch gets a reference to the given bool and assigns it to the BusinessBenefitSwitch field.
func (o *AlipayOpenSearchBoxApplyModel) SetBusinessBenefitSwitch(v bool) {
	o.BusinessBenefitSwitch = &v
}

// GetBusinessDistrictIds returns the BusinessDistrictIds field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetBusinessDistrictIds() []string {
	if o == nil || IsNil(o.BusinessDistrictIds) {
		var ret []string
		return ret
	}
	return o.BusinessDistrictIds
}

// GetBusinessDistrictIdsOk returns a tuple with the BusinessDistrictIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetBusinessDistrictIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.BusinessDistrictIds) {
		return nil, false
	}
	return o.BusinessDistrictIds, true
}

// HasBusinessDistrictIds returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasBusinessDistrictIds() bool {
	if o != nil && !IsNil(o.BusinessDistrictIds) {
		return true
	}

	return false
}

// SetBusinessDistrictIds gets a reference to the given []string and assigns it to the BusinessDistrictIds field.
func (o *AlipayOpenSearchBoxApplyModel) SetBusinessDistrictIds(v []string) {
	o.BusinessDistrictIds = v
}

// GetCustomKeywords returns the CustomKeywords field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetCustomKeywords() []string {
	if o == nil || IsNil(o.CustomKeywords) {
		var ret []string
		return ret
	}
	return o.CustomKeywords
}

// GetCustomKeywordsOk returns a tuple with the CustomKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetCustomKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomKeywords) {
		return nil, false
	}
	return o.CustomKeywords, true
}

// HasCustomKeywords returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasCustomKeywords() bool {
	if o != nil && !IsNil(o.CustomKeywords) {
		return true
	}

	return false
}

// SetCustomKeywords gets a reference to the given []string and assigns it to the CustomKeywords field.
func (o *AlipayOpenSearchBoxApplyModel) SetCustomKeywords(v []string) {
	o.CustomKeywords = v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *AlipayOpenSearchBoxApplyModel) SetImageId(v string) {
	o.ImageId = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetImageName() string {
	if o == nil || IsNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasImageName() bool {
	if o != nil && !IsNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *AlipayOpenSearchBoxApplyModel) SetImageName(v string) {
	o.ImageName = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *AlipayOpenSearchBoxApplyModel) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetRelatedAccounts returns the RelatedAccounts field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetRelatedAccounts() []SearchBoxAppInfo {
	if o == nil || IsNil(o.RelatedAccounts) {
		var ret []SearchBoxAppInfo
		return ret
	}
	return o.RelatedAccounts
}

// GetRelatedAccountsOk returns a tuple with the RelatedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetRelatedAccountsOk() ([]SearchBoxAppInfo, bool) {
	if o == nil || IsNil(o.RelatedAccounts) {
		return nil, false
	}
	return o.RelatedAccounts, true
}

// HasRelatedAccounts returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasRelatedAccounts() bool {
	if o != nil && !IsNil(o.RelatedAccounts) {
		return true
	}

	return false
}

// SetRelatedAccounts gets a reference to the given []SearchBoxAppInfo and assigns it to the RelatedAccounts field.
func (o *AlipayOpenSearchBoxApplyModel) SetRelatedAccounts(v []SearchBoxAppInfo) {
	o.RelatedAccounts = v
}

// GetServiceInfos returns the ServiceInfos field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetServiceInfos() []SearchBoxServiceInfo {
	if o == nil || IsNil(o.ServiceInfos) {
		var ret []SearchBoxServiceInfo
		return ret
	}
	return o.ServiceInfos
}

// GetServiceInfosOk returns a tuple with the ServiceInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetServiceInfosOk() ([]SearchBoxServiceInfo, bool) {
	if o == nil || IsNil(o.ServiceInfos) {
		return nil, false
	}
	return o.ServiceInfos, true
}

// HasServiceInfos returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasServiceInfos() bool {
	if o != nil && !IsNil(o.ServiceInfos) {
		return true
	}

	return false
}

// SetServiceInfos gets a reference to the given []SearchBoxServiceInfo and assigns it to the ServiceInfos field.
func (o *AlipayOpenSearchBoxApplyModel) SetServiceInfos(v []SearchBoxServiceInfo) {
	o.ServiceInfos = v
}

// GetTargetAppid returns the TargetAppid field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyModel) GetTargetAppid() string {
	if o == nil || IsNil(o.TargetAppid) {
		var ret string
		return ret
	}
	return *o.TargetAppid
}

// GetTargetAppidOk returns a tuple with the TargetAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyModel) GetTargetAppidOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAppid) {
		return nil, false
	}
	return o.TargetAppid, true
}

// HasTargetAppid returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyModel) HasTargetAppid() bool {
	if o != nil && !IsNil(o.TargetAppid) {
		return true
	}

	return false
}

// SetTargetAppid gets a reference to the given string and assigns it to the TargetAppid field.
func (o *AlipayOpenSearchBoxApplyModel) SetTargetAppid(v string) {
	o.TargetAppid = &v
}

func (o AlipayOpenSearchBoxApplyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchBoxApplyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AreaKeywords) {
		toSerialize["area_keywords"] = o.AreaKeywords
	}
	if !IsNil(o.BoxDesc) {
		toSerialize["box_desc"] = o.BoxDesc
	}
	if !IsNil(o.BrandId) {
		toSerialize["brand_id"] = o.BrandId
	}
	if !IsNil(o.BusinessBenefitSwitch) {
		toSerialize["business_benefit_switch"] = o.BusinessBenefitSwitch
	}
	if !IsNil(o.BusinessDistrictIds) {
		toSerialize["business_district_ids"] = o.BusinessDistrictIds
	}
	if !IsNil(o.CustomKeywords) {
		toSerialize["custom_keywords"] = o.CustomKeywords
	}
	if !IsNil(o.ImageId) {
		toSerialize["image_id"] = o.ImageId
	}
	if !IsNil(o.ImageName) {
		toSerialize["image_name"] = o.ImageName
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchant_id"] = o.MerchantId
	}
	if !IsNil(o.RelatedAccounts) {
		toSerialize["related_accounts"] = o.RelatedAccounts
	}
	if !IsNil(o.ServiceInfos) {
		toSerialize["service_infos"] = o.ServiceInfos
	}
	if !IsNil(o.TargetAppid) {
		toSerialize["target_appid"] = o.TargetAppid
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchBoxApplyModel struct {
	value *AlipayOpenSearchBoxApplyModel
	isSet bool
}

func (v NullableAlipayOpenSearchBoxApplyModel) Get() *AlipayOpenSearchBoxApplyModel {
	return v.value
}

func (v *NullableAlipayOpenSearchBoxApplyModel) Set(val *AlipayOpenSearchBoxApplyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchBoxApplyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchBoxApplyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchBoxApplyModel(val *AlipayOpenSearchBoxApplyModel) *NullableAlipayOpenSearchBoxApplyModel {
	return &NullableAlipayOpenSearchBoxApplyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchBoxApplyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchBoxApplyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
