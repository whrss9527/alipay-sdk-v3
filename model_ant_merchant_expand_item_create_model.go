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

// checks if the AntMerchantExpandItemCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntMerchantExpandItemCreateModel{}

// AntMerchantExpandItemCreateModel struct for AntMerchantExpandItemCreateModel
type AntMerchantExpandItemCreateModel struct {
	// 详情地址
	DetailUrl *string `json:"detail_url,omitempty"`
	// 商品扩展信息：可以解析成 Map<String, String> 的 json string
	ExtInfo *string `json:"ext_info,omitempty"`
	// 外部商品ID
	ExternalItemId *string `json:"external_item_id,omitempty"`
	// 前台类目id：target_type + target_id 和 front_category_id 二选一
	FrontCategoryId *string `json:"front_category_id,omitempty"`
	// 商品标签列表
	LabelList []ItemLabelCreateInfo `json:"label_list,omitempty"`
	// 主图地址
	MainPic *string `json:"main_pic,omitempty"`
	// 商品名称
	Name *string `json:"name,omitempty"`
	// 场景：GAS_CHARGE（加油）
	Scene *string `json:"scene,omitempty"`
	// 商品sku列表，至少有一个
	SkuList []ItemSkuCreateInfo `json:"sku_list,omitempty"`
	// 商户归属主体id
	TargetId *string `json:"target_id,omitempty"`
	// 商品归属主体类型：target_type + target_id 和 front_category_id 二选一    商品归属主体类型:  5: 店铺  4: 主站MID  3: 参与者  2: 角色  1: 联系人
	TargetType *string `json:"target_type,omitempty"`
}

// NewAntMerchantExpandItemCreateModel instantiates a new AntMerchantExpandItemCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntMerchantExpandItemCreateModel() *AntMerchantExpandItemCreateModel {
	this := AntMerchantExpandItemCreateModel{}
	return &this
}

// NewAntMerchantExpandItemCreateModelWithDefaults instantiates a new AntMerchantExpandItemCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntMerchantExpandItemCreateModelWithDefaults() *AntMerchantExpandItemCreateModel {
	this := AntMerchantExpandItemCreateModel{}
	return &this
}

// GetDetailUrl returns the DetailUrl field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetDetailUrl() string {
	if o == nil || IsNil(o.DetailUrl) {
		var ret string
		return ret
	}
	return *o.DetailUrl
}

// GetDetailUrlOk returns a tuple with the DetailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetDetailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DetailUrl) {
		return nil, false
	}
	return o.DetailUrl, true
}

// HasDetailUrl returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasDetailUrl() bool {
	if o != nil && !IsNil(o.DetailUrl) {
		return true
	}

	return false
}

// SetDetailUrl gets a reference to the given string and assigns it to the DetailUrl field.
func (o *AntMerchantExpandItemCreateModel) SetDetailUrl(v string) {
	o.DetailUrl = &v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetExtInfo() string {
	if o == nil || IsNil(o.ExtInfo) {
		var ret string
		return ret
	}
	return *o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetExtInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given string and assigns it to the ExtInfo field.
func (o *AntMerchantExpandItemCreateModel) SetExtInfo(v string) {
	o.ExtInfo = &v
}

// GetExternalItemId returns the ExternalItemId field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetExternalItemId() string {
	if o == nil || IsNil(o.ExternalItemId) {
		var ret string
		return ret
	}
	return *o.ExternalItemId
}

// GetExternalItemIdOk returns a tuple with the ExternalItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetExternalItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalItemId) {
		return nil, false
	}
	return o.ExternalItemId, true
}

// HasExternalItemId returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasExternalItemId() bool {
	if o != nil && !IsNil(o.ExternalItemId) {
		return true
	}

	return false
}

// SetExternalItemId gets a reference to the given string and assigns it to the ExternalItemId field.
func (o *AntMerchantExpandItemCreateModel) SetExternalItemId(v string) {
	o.ExternalItemId = &v
}

// GetFrontCategoryId returns the FrontCategoryId field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetFrontCategoryId() string {
	if o == nil || IsNil(o.FrontCategoryId) {
		var ret string
		return ret
	}
	return *o.FrontCategoryId
}

// GetFrontCategoryIdOk returns a tuple with the FrontCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetFrontCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.FrontCategoryId) {
		return nil, false
	}
	return o.FrontCategoryId, true
}

// HasFrontCategoryId returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasFrontCategoryId() bool {
	if o != nil && !IsNil(o.FrontCategoryId) {
		return true
	}

	return false
}

// SetFrontCategoryId gets a reference to the given string and assigns it to the FrontCategoryId field.
func (o *AntMerchantExpandItemCreateModel) SetFrontCategoryId(v string) {
	o.FrontCategoryId = &v
}

// GetLabelList returns the LabelList field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetLabelList() []ItemLabelCreateInfo {
	if o == nil || IsNil(o.LabelList) {
		var ret []ItemLabelCreateInfo
		return ret
	}
	return o.LabelList
}

// GetLabelListOk returns a tuple with the LabelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetLabelListOk() ([]ItemLabelCreateInfo, bool) {
	if o == nil || IsNil(o.LabelList) {
		return nil, false
	}
	return o.LabelList, true
}

// HasLabelList returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasLabelList() bool {
	if o != nil && !IsNil(o.LabelList) {
		return true
	}

	return false
}

// SetLabelList gets a reference to the given []ItemLabelCreateInfo and assigns it to the LabelList field.
func (o *AntMerchantExpandItemCreateModel) SetLabelList(v []ItemLabelCreateInfo) {
	o.LabelList = v
}

// GetMainPic returns the MainPic field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetMainPic() string {
	if o == nil || IsNil(o.MainPic) {
		var ret string
		return ret
	}
	return *o.MainPic
}

// GetMainPicOk returns a tuple with the MainPic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetMainPicOk() (*string, bool) {
	if o == nil || IsNil(o.MainPic) {
		return nil, false
	}
	return o.MainPic, true
}

// HasMainPic returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasMainPic() bool {
	if o != nil && !IsNil(o.MainPic) {
		return true
	}

	return false
}

// SetMainPic gets a reference to the given string and assigns it to the MainPic field.
func (o *AntMerchantExpandItemCreateModel) SetMainPic(v string) {
	o.MainPic = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AntMerchantExpandItemCreateModel) SetName(v string) {
	o.Name = &v
}

// GetScene returns the Scene field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetScene() string {
	if o == nil || IsNil(o.Scene) {
		var ret string
		return ret
	}
	return *o.Scene
}

// GetSceneOk returns a tuple with the Scene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetSceneOk() (*string, bool) {
	if o == nil || IsNil(o.Scene) {
		return nil, false
	}
	return o.Scene, true
}

// HasScene returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasScene() bool {
	if o != nil && !IsNil(o.Scene) {
		return true
	}

	return false
}

// SetScene gets a reference to the given string and assigns it to the Scene field.
func (o *AntMerchantExpandItemCreateModel) SetScene(v string) {
	o.Scene = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetSkuList() []ItemSkuCreateInfo {
	if o == nil || IsNil(o.SkuList) {
		var ret []ItemSkuCreateInfo
		return ret
	}
	return o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetSkuListOk() ([]ItemSkuCreateInfo, bool) {
	if o == nil || IsNil(o.SkuList) {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasSkuList() bool {
	if o != nil && !IsNil(o.SkuList) {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given []ItemSkuCreateInfo and assigns it to the SkuList field.
func (o *AntMerchantExpandItemCreateModel) SetSkuList(v []ItemSkuCreateInfo) {
	o.SkuList = v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AntMerchantExpandItemCreateModel) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateModel) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateModel) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateModel) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *AntMerchantExpandItemCreateModel) SetTargetType(v string) {
	o.TargetType = &v
}

func (o AntMerchantExpandItemCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntMerchantExpandItemCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DetailUrl) {
		toSerialize["detail_url"] = o.DetailUrl
	}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.ExternalItemId) {
		toSerialize["external_item_id"] = o.ExternalItemId
	}
	if !IsNil(o.FrontCategoryId) {
		toSerialize["front_category_id"] = o.FrontCategoryId
	}
	if !IsNil(o.LabelList) {
		toSerialize["label_list"] = o.LabelList
	}
	if !IsNil(o.MainPic) {
		toSerialize["main_pic"] = o.MainPic
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Scene) {
		toSerialize["scene"] = o.Scene
	}
	if !IsNil(o.SkuList) {
		toSerialize["sku_list"] = o.SkuList
	}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	return toSerialize, nil
}

type NullableAntMerchantExpandItemCreateModel struct {
	value *AntMerchantExpandItemCreateModel
	isSet bool
}

func (v NullableAntMerchantExpandItemCreateModel) Get() *AntMerchantExpandItemCreateModel {
	return v.value
}

func (v *NullableAntMerchantExpandItemCreateModel) Set(val *AntMerchantExpandItemCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandItemCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandItemCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandItemCreateModel(val *AntMerchantExpandItemCreateModel) *NullableAntMerchantExpandItemCreateModel {
	return &NullableAntMerchantExpandItemCreateModel{value: val, isSet: true}
}

func (v NullableAntMerchantExpandItemCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandItemCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
