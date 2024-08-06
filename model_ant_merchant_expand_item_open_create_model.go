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

// checks if the AntMerchantExpandItemOpenCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntMerchantExpandItemOpenCreateModel{}

// AntMerchantExpandItemOpenCreateModel struct for AntMerchantExpandItemOpenCreateModel
type AntMerchantExpandItemOpenCreateModel struct {
	// 商品描述
	Description *string `json:"description,omitempty"`
	// 商品扩展信息（具体KEY请参见产品文档）。 小程序订单中心场景接入参见 <a href=\"https://opendocs.alipay.com/mini/024hj4\">接入指南</a>。
	ExtInfo []ItemExtInfo `json:"ext_info,omitempty"`
	// 素材列表（最多3个），素材内容为素材key（素材key为<a href=\"https://opendocs.alipay.com/apis/api_4/alipay.merchant.item.file.upload\">alipay.merchant.item.file.upload</a> 接口返回的 material_key）
	MaterialList []MaterialCreateInfo `json:"material_list,omitempty"`
	// 商品名称
	Name *string `json:"name,omitempty"`
	// 商品属性列表
	PropertyList []ItemPropertyInfo `json:"property_list,omitempty"`
	// 业务场景码。
	Scene *string `json:"scene,omitempty"`
	// 商品SKU列表（至少1个，最多20个）
	SkuList []SkuCreateInfo `json:"sku_list,omitempty"`
	// 商品所属标准类目ID，标品及非标品都需传入。 请填入 <a href=\"https://opendocs.alipay.com/mini/011lxt\">小程序商品类目表</a> 中三级类目ID。
	StandardCategoryId *string `json:"standard_category_id,omitempty"`
	// 商品归属主体ID 例：商品归属主体类型target_type为店铺，则商品归属主体ID为店铺ID（支付宝侧店铺ID）；归属主体类型target_type为小程序，则归属主体ID为小程序ID
	TargetId *string `json:"target_id,omitempty"`
	// 商品归属主体类型。枚举如下： 5：店铺。 8：小程序。
	TargetType *string `json:"target_type,omitempty"`
	// 商品类型。枚举支持： STANDARD_GOODS：标品，一般是具有明确、标准规格、型号、参数的商品，如：手机、数码产品、大多数的家电 。 NON_STANDARD_GOODS：非标品，则是在这些方面没有统一的行业标准和参数规格的商品，如：服装、鞋袜等。
	Type *string `json:"type,omitempty"`
}

// NewAntMerchantExpandItemOpenCreateModel instantiates a new AntMerchantExpandItemOpenCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntMerchantExpandItemOpenCreateModel() *AntMerchantExpandItemOpenCreateModel {
	this := AntMerchantExpandItemOpenCreateModel{}
	return &this
}

// NewAntMerchantExpandItemOpenCreateModelWithDefaults instantiates a new AntMerchantExpandItemOpenCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntMerchantExpandItemOpenCreateModelWithDefaults() *AntMerchantExpandItemOpenCreateModel {
	this := AntMerchantExpandItemOpenCreateModel{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AntMerchantExpandItemOpenCreateModel) SetDescription(v string) {
	o.Description = &v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetExtInfo() []ItemExtInfo {
	if o == nil || IsNil(o.ExtInfo) {
		var ret []ItemExtInfo
		return ret
	}
	return o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetExtInfoOk() ([]ItemExtInfo, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given []ItemExtInfo and assigns it to the ExtInfo field.
func (o *AntMerchantExpandItemOpenCreateModel) SetExtInfo(v []ItemExtInfo) {
	o.ExtInfo = v
}

// GetMaterialList returns the MaterialList field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetMaterialList() []MaterialCreateInfo {
	if o == nil || IsNil(o.MaterialList) {
		var ret []MaterialCreateInfo
		return ret
	}
	return o.MaterialList
}

// GetMaterialListOk returns a tuple with the MaterialList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetMaterialListOk() ([]MaterialCreateInfo, bool) {
	if o == nil || IsNil(o.MaterialList) {
		return nil, false
	}
	return o.MaterialList, true
}

// HasMaterialList returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasMaterialList() bool {
	if o != nil && !IsNil(o.MaterialList) {
		return true
	}

	return false
}

// SetMaterialList gets a reference to the given []MaterialCreateInfo and assigns it to the MaterialList field.
func (o *AntMerchantExpandItemOpenCreateModel) SetMaterialList(v []MaterialCreateInfo) {
	o.MaterialList = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AntMerchantExpandItemOpenCreateModel) SetName(v string) {
	o.Name = &v
}

// GetPropertyList returns the PropertyList field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetPropertyList() []ItemPropertyInfo {
	if o == nil || IsNil(o.PropertyList) {
		var ret []ItemPropertyInfo
		return ret
	}
	return o.PropertyList
}

// GetPropertyListOk returns a tuple with the PropertyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetPropertyListOk() ([]ItemPropertyInfo, bool) {
	if o == nil || IsNil(o.PropertyList) {
		return nil, false
	}
	return o.PropertyList, true
}

// HasPropertyList returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasPropertyList() bool {
	if o != nil && !IsNil(o.PropertyList) {
		return true
	}

	return false
}

// SetPropertyList gets a reference to the given []ItemPropertyInfo and assigns it to the PropertyList field.
func (o *AntMerchantExpandItemOpenCreateModel) SetPropertyList(v []ItemPropertyInfo) {
	o.PropertyList = v
}

// GetScene returns the Scene field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetScene() string {
	if o == nil || IsNil(o.Scene) {
		var ret string
		return ret
	}
	return *o.Scene
}

// GetSceneOk returns a tuple with the Scene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetSceneOk() (*string, bool) {
	if o == nil || IsNil(o.Scene) {
		return nil, false
	}
	return o.Scene, true
}

// HasScene returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasScene() bool {
	if o != nil && !IsNil(o.Scene) {
		return true
	}

	return false
}

// SetScene gets a reference to the given string and assigns it to the Scene field.
func (o *AntMerchantExpandItemOpenCreateModel) SetScene(v string) {
	o.Scene = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetSkuList() []SkuCreateInfo {
	if o == nil || IsNil(o.SkuList) {
		var ret []SkuCreateInfo
		return ret
	}
	return o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetSkuListOk() ([]SkuCreateInfo, bool) {
	if o == nil || IsNil(o.SkuList) {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasSkuList() bool {
	if o != nil && !IsNil(o.SkuList) {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given []SkuCreateInfo and assigns it to the SkuList field.
func (o *AntMerchantExpandItemOpenCreateModel) SetSkuList(v []SkuCreateInfo) {
	o.SkuList = v
}

// GetStandardCategoryId returns the StandardCategoryId field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetStandardCategoryId() string {
	if o == nil || IsNil(o.StandardCategoryId) {
		var ret string
		return ret
	}
	return *o.StandardCategoryId
}

// GetStandardCategoryIdOk returns a tuple with the StandardCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetStandardCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.StandardCategoryId) {
		return nil, false
	}
	return o.StandardCategoryId, true
}

// HasStandardCategoryId returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasStandardCategoryId() bool {
	if o != nil && !IsNil(o.StandardCategoryId) {
		return true
	}

	return false
}

// SetStandardCategoryId gets a reference to the given string and assigns it to the StandardCategoryId field.
func (o *AntMerchantExpandItemOpenCreateModel) SetStandardCategoryId(v string) {
	o.StandardCategoryId = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AntMerchantExpandItemOpenCreateModel) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *AntMerchantExpandItemOpenCreateModel) SetTargetType(v string) {
	o.TargetType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AntMerchantExpandItemOpenCreateModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemOpenCreateModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AntMerchantExpandItemOpenCreateModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AntMerchantExpandItemOpenCreateModel) SetType(v string) {
	o.Type = &v
}

func (o AntMerchantExpandItemOpenCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntMerchantExpandItemOpenCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.MaterialList) {
		toSerialize["material_list"] = o.MaterialList
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PropertyList) {
		toSerialize["property_list"] = o.PropertyList
	}
	if !IsNil(o.Scene) {
		toSerialize["scene"] = o.Scene
	}
	if !IsNil(o.SkuList) {
		toSerialize["sku_list"] = o.SkuList
	}
	if !IsNil(o.StandardCategoryId) {
		toSerialize["standard_category_id"] = o.StandardCategoryId
	}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAntMerchantExpandItemOpenCreateModel struct {
	value *AntMerchantExpandItemOpenCreateModel
	isSet bool
}

func (v NullableAntMerchantExpandItemOpenCreateModel) Get() *AntMerchantExpandItemOpenCreateModel {
	return v.value
}

func (v *NullableAntMerchantExpandItemOpenCreateModel) Set(val *AntMerchantExpandItemOpenCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandItemOpenCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandItemOpenCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandItemOpenCreateModel(val *AntMerchantExpandItemOpenCreateModel) *NullableAntMerchantExpandItemOpenCreateModel {
	return &NullableAntMerchantExpandItemOpenCreateModel{value: val, isSet: true}
}

func (v NullableAntMerchantExpandItemOpenCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandItemOpenCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

