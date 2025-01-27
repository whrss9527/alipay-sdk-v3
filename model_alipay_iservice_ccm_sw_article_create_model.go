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

// checks if the AlipayIserviceCcmSwArticleCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmSwArticleCreateModel{}

// AlipayIserviceCcmSwArticleCreateModel struct for AlipayIserviceCcmSwArticleCreateModel
type AlipayIserviceCcmSwArticleCreateModel struct {
	// 所属类目ID，如果search_all_category为true则不用填
	CategoryId *int32 `json:"category_id,omitempty"`
	// 子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// 内容
	Content *string `json:"content,omitempty"`
	// 扩展标题
	ExtendTitles []string `json:"extend_titles,omitempty"`
	// 标签
	Keywords []string `json:"keywords,omitempty"`
	// 知识库ID
	LibraryId *string `json:"library_id,omitempty"`
	// 场景ID。KNOWLEDGE（内部知识库）；ROBOT（机器人）;HELP（帮助中心）；WHELP（无线帮助中心）
	SceneCodes []string `json:"scene_codes,omitempty"`
	// 标题
	Title *string `json:"title,omitempty"`
}

// NewAlipayIserviceCcmSwArticleCreateModel instantiates a new AlipayIserviceCcmSwArticleCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmSwArticleCreateModel() *AlipayIserviceCcmSwArticleCreateModel {
	this := AlipayIserviceCcmSwArticleCreateModel{}
	return &this
}

// NewAlipayIserviceCcmSwArticleCreateModelWithDefaults instantiates a new AlipayIserviceCcmSwArticleCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmSwArticleCreateModelWithDefaults() *AlipayIserviceCcmSwArticleCreateModel {
	this := AlipayIserviceCcmSwArticleCreateModel{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetCategoryId() int32 {
	if o == nil || IsNil(o.CategoryId) {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *AlipayIserviceCcmSwArticleCreateModel) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetCcsInstanceId returns the CcsInstanceId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetCcsInstanceId() string {
	if o == nil || IsNil(o.CcsInstanceId) {
		var ret string
		return ret
	}
	return *o.CcsInstanceId
}

// GetCcsInstanceIdOk returns a tuple with the CcsInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetCcsInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CcsInstanceId) {
		return nil, false
	}
	return o.CcsInstanceId, true
}

// HasCcsInstanceId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) HasCcsInstanceId() bool {
	if o != nil && !IsNil(o.CcsInstanceId) {
		return true
	}

	return false
}

// SetCcsInstanceId gets a reference to the given string and assigns it to the CcsInstanceId field.
func (o *AlipayIserviceCcmSwArticleCreateModel) SetCcsInstanceId(v string) {
	o.CcsInstanceId = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AlipayIserviceCcmSwArticleCreateModel) SetContent(v string) {
	o.Content = &v
}

// GetExtendTitles returns the ExtendTitles field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetExtendTitles() []string {
	if o == nil || IsNil(o.ExtendTitles) {
		var ret []string
		return ret
	}
	return o.ExtendTitles
}

// GetExtendTitlesOk returns a tuple with the ExtendTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetExtendTitlesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtendTitles) {
		return nil, false
	}
	return o.ExtendTitles, true
}

// HasExtendTitles returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) HasExtendTitles() bool {
	if o != nil && !IsNil(o.ExtendTitles) {
		return true
	}

	return false
}

// SetExtendTitles gets a reference to the given []string and assigns it to the ExtendTitles field.
func (o *AlipayIserviceCcmSwArticleCreateModel) SetExtendTitles(v []string) {
	o.ExtendTitles = v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetKeywords() []string {
	if o == nil || IsNil(o.Keywords) {
		var ret []string
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []string and assigns it to the Keywords field.
func (o *AlipayIserviceCcmSwArticleCreateModel) SetKeywords(v []string) {
	o.Keywords = v
}

// GetLibraryId returns the LibraryId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetLibraryId() string {
	if o == nil || IsNil(o.LibraryId) {
		var ret string
		return ret
	}
	return *o.LibraryId
}

// GetLibraryIdOk returns a tuple with the LibraryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetLibraryIdOk() (*string, bool) {
	if o == nil || IsNil(o.LibraryId) {
		return nil, false
	}
	return o.LibraryId, true
}

// HasLibraryId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) HasLibraryId() bool {
	if o != nil && !IsNil(o.LibraryId) {
		return true
	}

	return false
}

// SetLibraryId gets a reference to the given string and assigns it to the LibraryId field.
func (o *AlipayIserviceCcmSwArticleCreateModel) SetLibraryId(v string) {
	o.LibraryId = &v
}

// GetSceneCodes returns the SceneCodes field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetSceneCodes() []string {
	if o == nil || IsNil(o.SceneCodes) {
		var ret []string
		return ret
	}
	return o.SceneCodes
}

// GetSceneCodesOk returns a tuple with the SceneCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetSceneCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.SceneCodes) {
		return nil, false
	}
	return o.SceneCodes, true
}

// HasSceneCodes returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) HasSceneCodes() bool {
	if o != nil && !IsNil(o.SceneCodes) {
		return true
	}

	return false
}

// SetSceneCodes gets a reference to the given []string and assigns it to the SceneCodes field.
func (o *AlipayIserviceCcmSwArticleCreateModel) SetSceneCodes(v []string) {
	o.SceneCodes = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleCreateModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlipayIserviceCcmSwArticleCreateModel) SetTitle(v string) {
	o.Title = &v
}

func (o AlipayIserviceCcmSwArticleCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmSwArticleCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.CcsInstanceId) {
		toSerialize["ccs_instance_id"] = o.CcsInstanceId
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.ExtendTitles) {
		toSerialize["extend_titles"] = o.ExtendTitles
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.LibraryId) {
		toSerialize["library_id"] = o.LibraryId
	}
	if !IsNil(o.SceneCodes) {
		toSerialize["scene_codes"] = o.SceneCodes
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmSwArticleCreateModel struct {
	value *AlipayIserviceCcmSwArticleCreateModel
	isSet bool
}

func (v NullableAlipayIserviceCcmSwArticleCreateModel) Get() *AlipayIserviceCcmSwArticleCreateModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwArticleCreateModel) Set(val *AlipayIserviceCcmSwArticleCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwArticleCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwArticleCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwArticleCreateModel(val *AlipayIserviceCcmSwArticleCreateModel) *NullableAlipayIserviceCcmSwArticleCreateModel {
	return &NullableAlipayIserviceCcmSwArticleCreateModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwArticleCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwArticleCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
