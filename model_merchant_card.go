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

// checks if the MerchantCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCard{}

// MerchantCard struct for MerchantCard
type MerchantCard struct {
	// 资金卡余额，单位：元，精确到小数点后两位。
	Balance *string `json:"balance,omitempty"`
	// 支付宝业务卡号  说明：  1、开卡成功后返回该参数，需要保存留用；  2、开卡/更新/删卡/查询卡接口请求中不需要传该参数；
	BizCardNo *string `json:"biz_card_no,omitempty"`
	// 会员卡自定义资产值，只供特定业务使用，通常接入无需关注
	CustomAssets *string `json:"custom_assets,omitempty"`
	// 用户在商户crm系统中的会员卡卡号，该参数必填。
	ExternalCardNo *string `json:"external_card_no,omitempty"`
	// 卡面展示图片的图片ID，通过接口（alipay.offline.material.image.upload）上传图片    这里预期展示的是个人照片；  图片说明：1M以内，格式bmp、png、jpeg、jpg、gif；  图片尺寸为230*295px，可等比放大；
	FrontImageId *string `json:"front_image_id,omitempty"`
	// 卡面文案列表，1项对应1行文案，最多只能传入4行文案；  单行文案展现分为左右两部分，左边对应label字段，右边对应value；  形如： 学院    新闻学院
	FrontTextList []CardFrontTextDTO `json:"front_text_list,omitempty"`
	// 会员卡等级（由商户自定义，并可以在卡模板创建时，定义等级信息）
	Level *string `json:"level,omitempty"`
	MdcodeInfo *MdCodeInfoDTO `json:"mdcode_info,omitempty"`
	// 会员卡开卡时间，格式为yyyy-MM-dd HH:mm:ss。会员卡更新时，该时间不支持修改。
	OpenDate *string `json:"open_date,omitempty"`
	// 会员卡积分，积分必须为数字型（可为浮点型，带2位小数点）
	Point *string `json:"point,omitempty"`
	// 会员卡更换不同的卡模板（该参数仅用在会员卡更新接口中）
	TemplateId *string `json:"template_id,omitempty"`
	// 会员卡有效期结束时间，格式为yyyy-MM-dd HH:mm:ss。会员卡更新时，该时间不支持修改。
	ValidDate *string `json:"valid_date,omitempty"`
}

// NewMerchantCard instantiates a new MerchantCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCard() *MerchantCard {
	this := MerchantCard{}
	return &this
}

// NewMerchantCardWithDefaults instantiates a new MerchantCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCardWithDefaults() *MerchantCard {
	this := MerchantCard{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *MerchantCard) GetBalance() string {
	if o == nil || IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *MerchantCard) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *MerchantCard) SetBalance(v string) {
	o.Balance = &v
}

// GetBizCardNo returns the BizCardNo field value if set, zero value otherwise.
func (o *MerchantCard) GetBizCardNo() string {
	if o == nil || IsNil(o.BizCardNo) {
		var ret string
		return ret
	}
	return *o.BizCardNo
}

// GetBizCardNoOk returns a tuple with the BizCardNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetBizCardNoOk() (*string, bool) {
	if o == nil || IsNil(o.BizCardNo) {
		return nil, false
	}
	return o.BizCardNo, true
}

// HasBizCardNo returns a boolean if a field has been set.
func (o *MerchantCard) HasBizCardNo() bool {
	if o != nil && !IsNil(o.BizCardNo) {
		return true
	}

	return false
}

// SetBizCardNo gets a reference to the given string and assigns it to the BizCardNo field.
func (o *MerchantCard) SetBizCardNo(v string) {
	o.BizCardNo = &v
}

// GetCustomAssets returns the CustomAssets field value if set, zero value otherwise.
func (o *MerchantCard) GetCustomAssets() string {
	if o == nil || IsNil(o.CustomAssets) {
		var ret string
		return ret
	}
	return *o.CustomAssets
}

// GetCustomAssetsOk returns a tuple with the CustomAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetCustomAssetsOk() (*string, bool) {
	if o == nil || IsNil(o.CustomAssets) {
		return nil, false
	}
	return o.CustomAssets, true
}

// HasCustomAssets returns a boolean if a field has been set.
func (o *MerchantCard) HasCustomAssets() bool {
	if o != nil && !IsNil(o.CustomAssets) {
		return true
	}

	return false
}

// SetCustomAssets gets a reference to the given string and assigns it to the CustomAssets field.
func (o *MerchantCard) SetCustomAssets(v string) {
	o.CustomAssets = &v
}

// GetExternalCardNo returns the ExternalCardNo field value if set, zero value otherwise.
func (o *MerchantCard) GetExternalCardNo() string {
	if o == nil || IsNil(o.ExternalCardNo) {
		var ret string
		return ret
	}
	return *o.ExternalCardNo
}

// GetExternalCardNoOk returns a tuple with the ExternalCardNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetExternalCardNoOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalCardNo) {
		return nil, false
	}
	return o.ExternalCardNo, true
}

// HasExternalCardNo returns a boolean if a field has been set.
func (o *MerchantCard) HasExternalCardNo() bool {
	if o != nil && !IsNil(o.ExternalCardNo) {
		return true
	}

	return false
}

// SetExternalCardNo gets a reference to the given string and assigns it to the ExternalCardNo field.
func (o *MerchantCard) SetExternalCardNo(v string) {
	o.ExternalCardNo = &v
}

// GetFrontImageId returns the FrontImageId field value if set, zero value otherwise.
func (o *MerchantCard) GetFrontImageId() string {
	if o == nil || IsNil(o.FrontImageId) {
		var ret string
		return ret
	}
	return *o.FrontImageId
}

// GetFrontImageIdOk returns a tuple with the FrontImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetFrontImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.FrontImageId) {
		return nil, false
	}
	return o.FrontImageId, true
}

// HasFrontImageId returns a boolean if a field has been set.
func (o *MerchantCard) HasFrontImageId() bool {
	if o != nil && !IsNil(o.FrontImageId) {
		return true
	}

	return false
}

// SetFrontImageId gets a reference to the given string and assigns it to the FrontImageId field.
func (o *MerchantCard) SetFrontImageId(v string) {
	o.FrontImageId = &v
}

// GetFrontTextList returns the FrontTextList field value if set, zero value otherwise.
func (o *MerchantCard) GetFrontTextList() []CardFrontTextDTO {
	if o == nil || IsNil(o.FrontTextList) {
		var ret []CardFrontTextDTO
		return ret
	}
	return o.FrontTextList
}

// GetFrontTextListOk returns a tuple with the FrontTextList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetFrontTextListOk() ([]CardFrontTextDTO, bool) {
	if o == nil || IsNil(o.FrontTextList) {
		return nil, false
	}
	return o.FrontTextList, true
}

// HasFrontTextList returns a boolean if a field has been set.
func (o *MerchantCard) HasFrontTextList() bool {
	if o != nil && !IsNil(o.FrontTextList) {
		return true
	}

	return false
}

// SetFrontTextList gets a reference to the given []CardFrontTextDTO and assigns it to the FrontTextList field.
func (o *MerchantCard) SetFrontTextList(v []CardFrontTextDTO) {
	o.FrontTextList = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *MerchantCard) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *MerchantCard) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *MerchantCard) SetLevel(v string) {
	o.Level = &v
}

// GetMdcodeInfo returns the MdcodeInfo field value if set, zero value otherwise.
func (o *MerchantCard) GetMdcodeInfo() MdCodeInfoDTO {
	if o == nil || IsNil(o.MdcodeInfo) {
		var ret MdCodeInfoDTO
		return ret
	}
	return *o.MdcodeInfo
}

// GetMdcodeInfoOk returns a tuple with the MdcodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetMdcodeInfoOk() (*MdCodeInfoDTO, bool) {
	if o == nil || IsNil(o.MdcodeInfo) {
		return nil, false
	}
	return o.MdcodeInfo, true
}

// HasMdcodeInfo returns a boolean if a field has been set.
func (o *MerchantCard) HasMdcodeInfo() bool {
	if o != nil && !IsNil(o.MdcodeInfo) {
		return true
	}

	return false
}

// SetMdcodeInfo gets a reference to the given MdCodeInfoDTO and assigns it to the MdcodeInfo field.
func (o *MerchantCard) SetMdcodeInfo(v MdCodeInfoDTO) {
	o.MdcodeInfo = &v
}

// GetOpenDate returns the OpenDate field value if set, zero value otherwise.
func (o *MerchantCard) GetOpenDate() string {
	if o == nil || IsNil(o.OpenDate) {
		var ret string
		return ret
	}
	return *o.OpenDate
}

// GetOpenDateOk returns a tuple with the OpenDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetOpenDateOk() (*string, bool) {
	if o == nil || IsNil(o.OpenDate) {
		return nil, false
	}
	return o.OpenDate, true
}

// HasOpenDate returns a boolean if a field has been set.
func (o *MerchantCard) HasOpenDate() bool {
	if o != nil && !IsNil(o.OpenDate) {
		return true
	}

	return false
}

// SetOpenDate gets a reference to the given string and assigns it to the OpenDate field.
func (o *MerchantCard) SetOpenDate(v string) {
	o.OpenDate = &v
}

// GetPoint returns the Point field value if set, zero value otherwise.
func (o *MerchantCard) GetPoint() string {
	if o == nil || IsNil(o.Point) {
		var ret string
		return ret
	}
	return *o.Point
}

// GetPointOk returns a tuple with the Point field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetPointOk() (*string, bool) {
	if o == nil || IsNil(o.Point) {
		return nil, false
	}
	return o.Point, true
}

// HasPoint returns a boolean if a field has been set.
func (o *MerchantCard) HasPoint() bool {
	if o != nil && !IsNil(o.Point) {
		return true
	}

	return false
}

// SetPoint gets a reference to the given string and assigns it to the Point field.
func (o *MerchantCard) SetPoint(v string) {
	o.Point = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *MerchantCard) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *MerchantCard) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *MerchantCard) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetValidDate returns the ValidDate field value if set, zero value otherwise.
func (o *MerchantCard) GetValidDate() string {
	if o == nil || IsNil(o.ValidDate) {
		var ret string
		return ret
	}
	return *o.ValidDate
}

// GetValidDateOk returns a tuple with the ValidDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCard) GetValidDateOk() (*string, bool) {
	if o == nil || IsNil(o.ValidDate) {
		return nil, false
	}
	return o.ValidDate, true
}

// HasValidDate returns a boolean if a field has been set.
func (o *MerchantCard) HasValidDate() bool {
	if o != nil && !IsNil(o.ValidDate) {
		return true
	}

	return false
}

// SetValidDate gets a reference to the given string and assigns it to the ValidDate field.
func (o *MerchantCard) SetValidDate(v string) {
	o.ValidDate = &v
}

func (o MerchantCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.BizCardNo) {
		toSerialize["biz_card_no"] = o.BizCardNo
	}
	if !IsNil(o.CustomAssets) {
		toSerialize["custom_assets"] = o.CustomAssets
	}
	if !IsNil(o.ExternalCardNo) {
		toSerialize["external_card_no"] = o.ExternalCardNo
	}
	if !IsNil(o.FrontImageId) {
		toSerialize["front_image_id"] = o.FrontImageId
	}
	if !IsNil(o.FrontTextList) {
		toSerialize["front_text_list"] = o.FrontTextList
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.MdcodeInfo) {
		toSerialize["mdcode_info"] = o.MdcodeInfo
	}
	if !IsNil(o.OpenDate) {
		toSerialize["open_date"] = o.OpenDate
	}
	if !IsNil(o.Point) {
		toSerialize["point"] = o.Point
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.ValidDate) {
		toSerialize["valid_date"] = o.ValidDate
	}
	return toSerialize, nil
}

type NullableMerchantCard struct {
	value *MerchantCard
	isSet bool
}

func (v NullableMerchantCard) Get() *MerchantCard {
	return v.value
}

func (v *NullableMerchantCard) Set(val *MerchantCard) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCard) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCard(val *MerchantCard) *NullableMerchantCard {
	return &NullableMerchantCard{value: val, isSet: true}
}

func (v NullableMerchantCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

