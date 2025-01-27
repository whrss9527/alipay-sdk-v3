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

// checks if the AlipayOpenSpMerchantInconsistentApproveModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSpMerchantInconsistentApproveModel{}

// AlipayOpenSpMerchantInconsistentApproveModel struct for AlipayOpenSpMerchantInconsistentApproveModel
type AlipayOpenSpMerchantInconsistentApproveModel struct {
	// 身份证背面照，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href=\"https://opendocs.alipay.com/apis/01ea4t\">alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。
	BackCard *string `json:"back_card,omitempty"`
	// 实际经营人承诺函照片，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href=\"https://opendocs.alipay.com/apis/01ea4t\">alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见报名资质要求。
	CommitmentLetter *string `json:"commitment_letter,omitempty"`
	// 身份证正面照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式. 请传<a href=\"https://opendocs.alipay.com/apis/01ea4t\">alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。
	FrontCard *string `json:"front_card,omitempty"`
	// 手持营业执照合照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href=\"https://opendocs.alipay.com/apis/01ea4t\">alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见报名资质要求。
	HandheldBusinessLicense *string `json:"handheld_business_license,omitempty"`
	// 手持身份证合照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href=\"https://opendocs.alipay.com/apis/01ea4t\">alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见报名资质要求。
	HandheldCard *string `json:"handheld_card,omitempty"`
	// 服务优选商品编码
	ItemCode *string `json:"item_code,omitempty"`
	// 商家pid
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// 商家小程序appId
	MiniAppid *string `json:"mini_appid,omitempty"`
	// 外部业务号，需不重复
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 推广服务商(S2)pid
	PromotorPid *string `json:"promotor_pid,omitempty"`
}

// NewAlipayOpenSpMerchantInconsistentApproveModel instantiates a new AlipayOpenSpMerchantInconsistentApproveModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSpMerchantInconsistentApproveModel() *AlipayOpenSpMerchantInconsistentApproveModel {
	this := AlipayOpenSpMerchantInconsistentApproveModel{}
	return &this
}

// NewAlipayOpenSpMerchantInconsistentApproveModelWithDefaults instantiates a new AlipayOpenSpMerchantInconsistentApproveModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSpMerchantInconsistentApproveModelWithDefaults() *AlipayOpenSpMerchantInconsistentApproveModel {
	this := AlipayOpenSpMerchantInconsistentApproveModel{}
	return &this
}

// GetBackCard returns the BackCard field value if set, zero value otherwise.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetBackCard() string {
	if o == nil || IsNil(o.BackCard) {
		var ret string
		return ret
	}
	return *o.BackCard
}

// GetBackCardOk returns a tuple with the BackCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetBackCardOk() (*string, bool) {
	if o == nil || IsNil(o.BackCard) {
		return nil, false
	}
	return o.BackCard, true
}

// HasBackCard returns a boolean if a field has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) HasBackCard() bool {
	if o != nil && !IsNil(o.BackCard) {
		return true
	}

	return false
}

// SetBackCard gets a reference to the given string and assigns it to the BackCard field.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) SetBackCard(v string) {
	o.BackCard = &v
}

// GetCommitmentLetter returns the CommitmentLetter field value if set, zero value otherwise.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetCommitmentLetter() string {
	if o == nil || IsNil(o.CommitmentLetter) {
		var ret string
		return ret
	}
	return *o.CommitmentLetter
}

// GetCommitmentLetterOk returns a tuple with the CommitmentLetter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetCommitmentLetterOk() (*string, bool) {
	if o == nil || IsNil(o.CommitmentLetter) {
		return nil, false
	}
	return o.CommitmentLetter, true
}

// HasCommitmentLetter returns a boolean if a field has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) HasCommitmentLetter() bool {
	if o != nil && !IsNil(o.CommitmentLetter) {
		return true
	}

	return false
}

// SetCommitmentLetter gets a reference to the given string and assigns it to the CommitmentLetter field.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) SetCommitmentLetter(v string) {
	o.CommitmentLetter = &v
}

// GetFrontCard returns the FrontCard field value if set, zero value otherwise.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetFrontCard() string {
	if o == nil || IsNil(o.FrontCard) {
		var ret string
		return ret
	}
	return *o.FrontCard
}

// GetFrontCardOk returns a tuple with the FrontCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetFrontCardOk() (*string, bool) {
	if o == nil || IsNil(o.FrontCard) {
		return nil, false
	}
	return o.FrontCard, true
}

// HasFrontCard returns a boolean if a field has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) HasFrontCard() bool {
	if o != nil && !IsNil(o.FrontCard) {
		return true
	}

	return false
}

// SetFrontCard gets a reference to the given string and assigns it to the FrontCard field.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) SetFrontCard(v string) {
	o.FrontCard = &v
}

// GetHandheldBusinessLicense returns the HandheldBusinessLicense field value if set, zero value otherwise.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetHandheldBusinessLicense() string {
	if o == nil || IsNil(o.HandheldBusinessLicense) {
		var ret string
		return ret
	}
	return *o.HandheldBusinessLicense
}

// GetHandheldBusinessLicenseOk returns a tuple with the HandheldBusinessLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetHandheldBusinessLicenseOk() (*string, bool) {
	if o == nil || IsNil(o.HandheldBusinessLicense) {
		return nil, false
	}
	return o.HandheldBusinessLicense, true
}

// HasHandheldBusinessLicense returns a boolean if a field has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) HasHandheldBusinessLicense() bool {
	if o != nil && !IsNil(o.HandheldBusinessLicense) {
		return true
	}

	return false
}

// SetHandheldBusinessLicense gets a reference to the given string and assigns it to the HandheldBusinessLicense field.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) SetHandheldBusinessLicense(v string) {
	o.HandheldBusinessLicense = &v
}

// GetHandheldCard returns the HandheldCard field value if set, zero value otherwise.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetHandheldCard() string {
	if o == nil || IsNil(o.HandheldCard) {
		var ret string
		return ret
	}
	return *o.HandheldCard
}

// GetHandheldCardOk returns a tuple with the HandheldCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetHandheldCardOk() (*string, bool) {
	if o == nil || IsNil(o.HandheldCard) {
		return nil, false
	}
	return o.HandheldCard, true
}

// HasHandheldCard returns a boolean if a field has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) HasHandheldCard() bool {
	if o != nil && !IsNil(o.HandheldCard) {
		return true
	}

	return false
}

// SetHandheldCard gets a reference to the given string and assigns it to the HandheldCard field.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) SetHandheldCard(v string) {
	o.HandheldCard = &v
}

// GetItemCode returns the ItemCode field value if set, zero value otherwise.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetItemCode() string {
	if o == nil || IsNil(o.ItemCode) {
		var ret string
		return ret
	}
	return *o.ItemCode
}

// GetItemCodeOk returns a tuple with the ItemCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetItemCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ItemCode) {
		return nil, false
	}
	return o.ItemCode, true
}

// HasItemCode returns a boolean if a field has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) HasItemCode() bool {
	if o != nil && !IsNil(o.ItemCode) {
		return true
	}

	return false
}

// SetItemCode gets a reference to the given string and assigns it to the ItemCode field.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) SetItemCode(v string) {
	o.ItemCode = &v
}

// GetMerchantPid returns the MerchantPid field value if set, zero value otherwise.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetMerchantPid() string {
	if o == nil || IsNil(o.MerchantPid) {
		var ret string
		return ret
	}
	return *o.MerchantPid
}

// GetMerchantPidOk returns a tuple with the MerchantPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetMerchantPidOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantPid) {
		return nil, false
	}
	return o.MerchantPid, true
}

// HasMerchantPid returns a boolean if a field has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) HasMerchantPid() bool {
	if o != nil && !IsNil(o.MerchantPid) {
		return true
	}

	return false
}

// SetMerchantPid gets a reference to the given string and assigns it to the MerchantPid field.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) SetMerchantPid(v string) {
	o.MerchantPid = &v
}

// GetMiniAppid returns the MiniAppid field value if set, zero value otherwise.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetMiniAppid() string {
	if o == nil || IsNil(o.MiniAppid) {
		var ret string
		return ret
	}
	return *o.MiniAppid
}

// GetMiniAppidOk returns a tuple with the MiniAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetMiniAppidOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppid) {
		return nil, false
	}
	return o.MiniAppid, true
}

// HasMiniAppid returns a boolean if a field has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) HasMiniAppid() bool {
	if o != nil && !IsNil(o.MiniAppid) {
		return true
	}

	return false
}

// SetMiniAppid gets a reference to the given string and assigns it to the MiniAppid field.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) SetMiniAppid(v string) {
	o.MiniAppid = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPromotorPid returns the PromotorPid field value if set, zero value otherwise.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetPromotorPid() string {
	if o == nil || IsNil(o.PromotorPid) {
		var ret string
		return ret
	}
	return *o.PromotorPid
}

// GetPromotorPidOk returns a tuple with the PromotorPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) GetPromotorPidOk() (*string, bool) {
	if o == nil || IsNil(o.PromotorPid) {
		return nil, false
	}
	return o.PromotorPid, true
}

// HasPromotorPid returns a boolean if a field has been set.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) HasPromotorPid() bool {
	if o != nil && !IsNil(o.PromotorPid) {
		return true
	}

	return false
}

// SetPromotorPid gets a reference to the given string and assigns it to the PromotorPid field.
func (o *AlipayOpenSpMerchantInconsistentApproveModel) SetPromotorPid(v string) {
	o.PromotorPid = &v
}

func (o AlipayOpenSpMerchantInconsistentApproveModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSpMerchantInconsistentApproveModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackCard) {
		toSerialize["back_card"] = o.BackCard
	}
	if !IsNil(o.CommitmentLetter) {
		toSerialize["commitment_letter"] = o.CommitmentLetter
	}
	if !IsNil(o.FrontCard) {
		toSerialize["front_card"] = o.FrontCard
	}
	if !IsNil(o.HandheldBusinessLicense) {
		toSerialize["handheld_business_license"] = o.HandheldBusinessLicense
	}
	if !IsNil(o.HandheldCard) {
		toSerialize["handheld_card"] = o.HandheldCard
	}
	if !IsNil(o.ItemCode) {
		toSerialize["item_code"] = o.ItemCode
	}
	if !IsNil(o.MerchantPid) {
		toSerialize["merchant_pid"] = o.MerchantPid
	}
	if !IsNil(o.MiniAppid) {
		toSerialize["mini_appid"] = o.MiniAppid
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PromotorPid) {
		toSerialize["promotor_pid"] = o.PromotorPid
	}
	return toSerialize, nil
}

type NullableAlipayOpenSpMerchantInconsistentApproveModel struct {
	value *AlipayOpenSpMerchantInconsistentApproveModel
	isSet bool
}

func (v NullableAlipayOpenSpMerchantInconsistentApproveModel) Get() *AlipayOpenSpMerchantInconsistentApproveModel {
	return v.value
}

func (v *NullableAlipayOpenSpMerchantInconsistentApproveModel) Set(val *AlipayOpenSpMerchantInconsistentApproveModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpMerchantInconsistentApproveModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpMerchantInconsistentApproveModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpMerchantInconsistentApproveModel(val *AlipayOpenSpMerchantInconsistentApproveModel) *NullableAlipayOpenSpMerchantInconsistentApproveModel {
	return &NullableAlipayOpenSpMerchantInconsistentApproveModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSpMerchantInconsistentApproveModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpMerchantInconsistentApproveModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
