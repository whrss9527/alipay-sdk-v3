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

// checks if the CmItemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CmItemInfo{}

// CmItemInfo struct for CmItemInfo
type CmItemInfo struct {
	// 商品描述
	Description *string `json:"description,omitempty"`
	// 商品扩展信息
	ExtInfo []ItemExtInfo `json:"ext_info,omitempty"`
	// 商品所属前台类目ID列表
	FrontCategoryIdList []string `json:"front_category_id_list,omitempty"`
	// 商品创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// 商品更新时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// 商品ID
	ItemId *string `json:"item_id,omitempty"`
	// 商品素材列表
	MaterialList []MaterialInfo `json:"material_list,omitempty"`
	// 商品名称
	Name *string `json:"name,omitempty"`
	// 商品属性列表
	PropertyList []ItemPropertyInfo `json:"property_list,omitempty"`
	// 商品SKU列表
	SkuList []CmItemSkuInfo `json:"sku_list,omitempty"`
	// 商品所属标准类目ID（具体值请参见产品文档）
	StandardCategoryId *string `json:"standard_category_id,omitempty"`
	// 商品状态
	Status *string `json:"status,omitempty"`
	// 商品归属主体ID  例：商品归属主体类型为店铺，则商品归属主体ID为店铺ID；归属主体为小程序，则归属主体ID为小程序ID
	TargetId *string `json:"target_id,omitempty"`
	// 商品归属主体类型:  5（店铺）  8（小程序）
	TargetType *string `json:"target_type,omitempty"`
	// 商品类型：  STANDARD_GOODS（标品）、NON_STANDARD_GOODS（非标品）
	Type *string `json:"type,omitempty"`
}

// NewCmItemInfo instantiates a new CmItemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmItemInfo() *CmItemInfo {
	this := CmItemInfo{}
	return &this
}

// NewCmItemInfoWithDefaults instantiates a new CmItemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmItemInfoWithDefaults() *CmItemInfo {
	this := CmItemInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CmItemInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CmItemInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CmItemInfo) SetDescription(v string) {
	o.Description = &v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *CmItemInfo) GetExtInfo() []ItemExtInfo {
	if o == nil || IsNil(o.ExtInfo) {
		var ret []ItemExtInfo
		return ret
	}
	return o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetExtInfoOk() ([]ItemExtInfo, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *CmItemInfo) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given []ItemExtInfo and assigns it to the ExtInfo field.
func (o *CmItemInfo) SetExtInfo(v []ItemExtInfo) {
	o.ExtInfo = v
}

// GetFrontCategoryIdList returns the FrontCategoryIdList field value if set, zero value otherwise.
func (o *CmItemInfo) GetFrontCategoryIdList() []string {
	if o == nil || IsNil(o.FrontCategoryIdList) {
		var ret []string
		return ret
	}
	return o.FrontCategoryIdList
}

// GetFrontCategoryIdListOk returns a tuple with the FrontCategoryIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetFrontCategoryIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.FrontCategoryIdList) {
		return nil, false
	}
	return o.FrontCategoryIdList, true
}

// HasFrontCategoryIdList returns a boolean if a field has been set.
func (o *CmItemInfo) HasFrontCategoryIdList() bool {
	if o != nil && !IsNil(o.FrontCategoryIdList) {
		return true
	}

	return false
}

// SetFrontCategoryIdList gets a reference to the given []string and assigns it to the FrontCategoryIdList field.
func (o *CmItemInfo) SetFrontCategoryIdList(v []string) {
	o.FrontCategoryIdList = v
}

// GetGmtCreate returns the GmtCreate field value if set, zero value otherwise.
func (o *CmItemInfo) GetGmtCreate() string {
	if o == nil || IsNil(o.GmtCreate) {
		var ret string
		return ret
	}
	return *o.GmtCreate
}

// GetGmtCreateOk returns a tuple with the GmtCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetGmtCreateOk() (*string, bool) {
	if o == nil || IsNil(o.GmtCreate) {
		return nil, false
	}
	return o.GmtCreate, true
}

// HasGmtCreate returns a boolean if a field has been set.
func (o *CmItemInfo) HasGmtCreate() bool {
	if o != nil && !IsNil(o.GmtCreate) {
		return true
	}

	return false
}

// SetGmtCreate gets a reference to the given string and assigns it to the GmtCreate field.
func (o *CmItemInfo) SetGmtCreate(v string) {
	o.GmtCreate = &v
}

// GetGmtModified returns the GmtModified field value if set, zero value otherwise.
func (o *CmItemInfo) GetGmtModified() string {
	if o == nil || IsNil(o.GmtModified) {
		var ret string
		return ret
	}
	return *o.GmtModified
}

// GetGmtModifiedOk returns a tuple with the GmtModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetGmtModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModified) {
		return nil, false
	}
	return o.GmtModified, true
}

// HasGmtModified returns a boolean if a field has been set.
func (o *CmItemInfo) HasGmtModified() bool {
	if o != nil && !IsNil(o.GmtModified) {
		return true
	}

	return false
}

// SetGmtModified gets a reference to the given string and assigns it to the GmtModified field.
func (o *CmItemInfo) SetGmtModified(v string) {
	o.GmtModified = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *CmItemInfo) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *CmItemInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *CmItemInfo) SetItemId(v string) {
	o.ItemId = &v
}

// GetMaterialList returns the MaterialList field value if set, zero value otherwise.
func (o *CmItemInfo) GetMaterialList() []MaterialInfo {
	if o == nil || IsNil(o.MaterialList) {
		var ret []MaterialInfo
		return ret
	}
	return o.MaterialList
}

// GetMaterialListOk returns a tuple with the MaterialList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetMaterialListOk() ([]MaterialInfo, bool) {
	if o == nil || IsNil(o.MaterialList) {
		return nil, false
	}
	return o.MaterialList, true
}

// HasMaterialList returns a boolean if a field has been set.
func (o *CmItemInfo) HasMaterialList() bool {
	if o != nil && !IsNil(o.MaterialList) {
		return true
	}

	return false
}

// SetMaterialList gets a reference to the given []MaterialInfo and assigns it to the MaterialList field.
func (o *CmItemInfo) SetMaterialList(v []MaterialInfo) {
	o.MaterialList = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CmItemInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CmItemInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CmItemInfo) SetName(v string) {
	o.Name = &v
}

// GetPropertyList returns the PropertyList field value if set, zero value otherwise.
func (o *CmItemInfo) GetPropertyList() []ItemPropertyInfo {
	if o == nil || IsNil(o.PropertyList) {
		var ret []ItemPropertyInfo
		return ret
	}
	return o.PropertyList
}

// GetPropertyListOk returns a tuple with the PropertyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetPropertyListOk() ([]ItemPropertyInfo, bool) {
	if o == nil || IsNil(o.PropertyList) {
		return nil, false
	}
	return o.PropertyList, true
}

// HasPropertyList returns a boolean if a field has been set.
func (o *CmItemInfo) HasPropertyList() bool {
	if o != nil && !IsNil(o.PropertyList) {
		return true
	}

	return false
}

// SetPropertyList gets a reference to the given []ItemPropertyInfo and assigns it to the PropertyList field.
func (o *CmItemInfo) SetPropertyList(v []ItemPropertyInfo) {
	o.PropertyList = v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *CmItemInfo) GetSkuList() []CmItemSkuInfo {
	if o == nil || IsNil(o.SkuList) {
		var ret []CmItemSkuInfo
		return ret
	}
	return o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetSkuListOk() ([]CmItemSkuInfo, bool) {
	if o == nil || IsNil(o.SkuList) {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *CmItemInfo) HasSkuList() bool {
	if o != nil && !IsNil(o.SkuList) {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given []CmItemSkuInfo and assigns it to the SkuList field.
func (o *CmItemInfo) SetSkuList(v []CmItemSkuInfo) {
	o.SkuList = v
}

// GetStandardCategoryId returns the StandardCategoryId field value if set, zero value otherwise.
func (o *CmItemInfo) GetStandardCategoryId() string {
	if o == nil || IsNil(o.StandardCategoryId) {
		var ret string
		return ret
	}
	return *o.StandardCategoryId
}

// GetStandardCategoryIdOk returns a tuple with the StandardCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetStandardCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.StandardCategoryId) {
		return nil, false
	}
	return o.StandardCategoryId, true
}

// HasStandardCategoryId returns a boolean if a field has been set.
func (o *CmItemInfo) HasStandardCategoryId() bool {
	if o != nil && !IsNil(o.StandardCategoryId) {
		return true
	}

	return false
}

// SetStandardCategoryId gets a reference to the given string and assigns it to the StandardCategoryId field.
func (o *CmItemInfo) SetStandardCategoryId(v string) {
	o.StandardCategoryId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CmItemInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CmItemInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CmItemInfo) SetStatus(v string) {
	o.Status = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *CmItemInfo) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *CmItemInfo) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *CmItemInfo) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *CmItemInfo) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *CmItemInfo) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *CmItemInfo) SetTargetType(v string) {
	o.TargetType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CmItemInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmItemInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CmItemInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CmItemInfo) SetType(v string) {
	o.Type = &v
}

func (o CmItemInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmItemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.FrontCategoryIdList) {
		toSerialize["front_category_id_list"] = o.FrontCategoryIdList
	}
	if !IsNil(o.GmtCreate) {
		toSerialize["gmt_create"] = o.GmtCreate
	}
	if !IsNil(o.GmtModified) {
		toSerialize["gmt_modified"] = o.GmtModified
	}
	if !IsNil(o.ItemId) {
		toSerialize["item_id"] = o.ItemId
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
	if !IsNil(o.SkuList) {
		toSerialize["sku_list"] = o.SkuList
	}
	if !IsNil(o.StandardCategoryId) {
		toSerialize["standard_category_id"] = o.StandardCategoryId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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

type NullableCmItemInfo struct {
	value *CmItemInfo
	isSet bool
}

func (v NullableCmItemInfo) Get() *CmItemInfo {
	return v.value
}

func (v *NullableCmItemInfo) Set(val *CmItemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCmItemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCmItemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmItemInfo(val *CmItemInfo) *NullableCmItemInfo {
	return &NullableCmItemInfo{value: val, isSet: true}
}

func (v NullableCmItemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmItemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
