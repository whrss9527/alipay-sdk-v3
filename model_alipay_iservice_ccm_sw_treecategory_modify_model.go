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

// checks if the AlipayIserviceCcmSwTreecategoryModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmSwTreecategoryModifyModel{}

// AlipayIserviceCcmSwTreecategoryModifyModel struct for AlipayIserviceCcmSwTreecategoryModifyModel
type AlipayIserviceCcmSwTreecategoryModifyModel struct {
	// 子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// 描述
	Description *string `json:"description,omitempty"`
	// 父节点ID
	FatherId *string `json:"father_id,omitempty"`
	// 节点ID
	Id *int32 `json:"id,omitempty"`
	// 节点名称
	Name *string `json:"name,omitempty"`
	// 标签。KNOWLEDGE（知识库）；PLATFORM（公有云工作台）；HELP（公有云帮助中心）
	Tags *string `json:"tags,omitempty"`
}

// NewAlipayIserviceCcmSwTreecategoryModifyModel instantiates a new AlipayIserviceCcmSwTreecategoryModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmSwTreecategoryModifyModel() *AlipayIserviceCcmSwTreecategoryModifyModel {
	this := AlipayIserviceCcmSwTreecategoryModifyModel{}
	return &this
}

// NewAlipayIserviceCcmSwTreecategoryModifyModelWithDefaults instantiates a new AlipayIserviceCcmSwTreecategoryModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmSwTreecategoryModifyModelWithDefaults() *AlipayIserviceCcmSwTreecategoryModifyModel {
	this := AlipayIserviceCcmSwTreecategoryModifyModel{}
	return &this
}

// GetCcsInstanceId returns the CcsInstanceId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetCcsInstanceId() string {
	if o == nil || IsNil(o.CcsInstanceId) {
		var ret string
		return ret
	}
	return *o.CcsInstanceId
}

// GetCcsInstanceIdOk returns a tuple with the CcsInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetCcsInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CcsInstanceId) {
		return nil, false
	}
	return o.CcsInstanceId, true
}

// HasCcsInstanceId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) HasCcsInstanceId() bool {
	if o != nil && !IsNil(o.CcsInstanceId) {
		return true
	}

	return false
}

// SetCcsInstanceId gets a reference to the given string and assigns it to the CcsInstanceId field.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) SetCcsInstanceId(v string) {
	o.CcsInstanceId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) SetDescription(v string) {
	o.Description = &v
}

// GetFatherId returns the FatherId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetFatherId() string {
	if o == nil || IsNil(o.FatherId) {
		var ret string
		return ret
	}
	return *o.FatherId
}

// GetFatherIdOk returns a tuple with the FatherId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetFatherIdOk() (*string, bool) {
	if o == nil || IsNil(o.FatherId) {
		return nil, false
	}
	return o.FatherId, true
}

// HasFatherId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) HasFatherId() bool {
	if o != nil && !IsNil(o.FatherId) {
		return true
	}

	return false
}

// SetFatherId gets a reference to the given string and assigns it to the FatherId field.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) SetFatherId(v string) {
	o.FatherId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *AlipayIserviceCcmSwTreecategoryModifyModel) SetTags(v string) {
	o.Tags = &v
}

func (o AlipayIserviceCcmSwTreecategoryModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmSwTreecategoryModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CcsInstanceId) {
		toSerialize["ccs_instance_id"] = o.CcsInstanceId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FatherId) {
		toSerialize["father_id"] = o.FatherId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmSwTreecategoryModifyModel struct {
	value *AlipayIserviceCcmSwTreecategoryModifyModel
	isSet bool
}

func (v NullableAlipayIserviceCcmSwTreecategoryModifyModel) Get() *AlipayIserviceCcmSwTreecategoryModifyModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwTreecategoryModifyModel) Set(val *AlipayIserviceCcmSwTreecategoryModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwTreecategoryModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwTreecategoryModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwTreecategoryModifyModel(val *AlipayIserviceCcmSwTreecategoryModifyModel) *NullableAlipayIserviceCcmSwTreecategoryModifyModel {
	return &NullableAlipayIserviceCcmSwTreecategoryModifyModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwTreecategoryModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwTreecategoryModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
