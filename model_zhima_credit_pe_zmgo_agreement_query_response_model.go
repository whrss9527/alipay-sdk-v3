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

// checks if the ZhimaCreditPeZmgoAgreementQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaCreditPeZmgoAgreementQueryResponseModel{}

// ZhimaCreditPeZmgoAgreementQueryResponseModel struct for ZhimaCreditPeZmgoAgreementQueryResponseModel
type ZhimaCreditPeZmgoAgreementQueryResponseModel struct {
	// 支付宝系统中用以唯一标识用户签约记录的编号，即花芝轻会员协议号。
	AgreementId *string `json:"agreement_id,omitempty"`
	// 协议名称
	AgreementName *string `json:"agreement_name,omitempty"`
	// 协议状态。Y表示状态有效，P表示状态失效中，N表示状态已失效
	AgreementStatus *string `json:"agreement_status,omitempty"`
	// 支付宝用户userId
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// 代表签约芝麻GO所属业务类型
	BizType *string `json:"biz_type,omitempty"`
	// 支付宝用户userId
	OpenId *string `json:"open_id,omitempty"`
	// 该条芝麻GO协议签约时间
	SignTime *string `json:"sign_time,omitempty"`
	// 该芝麻GO协议开始时间
	StartTime *string `json:"start_time,omitempty"`
}

// NewZhimaCreditPeZmgoAgreementQueryResponseModel instantiates a new ZhimaCreditPeZmgoAgreementQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaCreditPeZmgoAgreementQueryResponseModel() *ZhimaCreditPeZmgoAgreementQueryResponseModel {
	this := ZhimaCreditPeZmgoAgreementQueryResponseModel{}
	return &this
}

// NewZhimaCreditPeZmgoAgreementQueryResponseModelWithDefaults instantiates a new ZhimaCreditPeZmgoAgreementQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaCreditPeZmgoAgreementQueryResponseModelWithDefaults() *ZhimaCreditPeZmgoAgreementQueryResponseModel {
	this := ZhimaCreditPeZmgoAgreementQueryResponseModel{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetAgreementName returns the AgreementName field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetAgreementName() string {
	if o == nil || IsNil(o.AgreementName) {
		var ret string
		return ret
	}
	return *o.AgreementName
}

// GetAgreementNameOk returns a tuple with the AgreementName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetAgreementNameOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementName) {
		return nil, false
	}
	return o.AgreementName, true
}

// HasAgreementName returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) HasAgreementName() bool {
	if o != nil && !IsNil(o.AgreementName) {
		return true
	}

	return false
}

// SetAgreementName gets a reference to the given string and assigns it to the AgreementName field.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) SetAgreementName(v string) {
	o.AgreementName = &v
}

// GetAgreementStatus returns the AgreementStatus field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetAgreementStatus() string {
	if o == nil || IsNil(o.AgreementStatus) {
		var ret string
		return ret
	}
	return *o.AgreementStatus
}

// GetAgreementStatusOk returns a tuple with the AgreementStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetAgreementStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementStatus) {
		return nil, false
	}
	return o.AgreementStatus, true
}

// HasAgreementStatus returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) HasAgreementStatus() bool {
	if o != nil && !IsNil(o.AgreementStatus) {
		return true
	}

	return false
}

// SetAgreementStatus gets a reference to the given string and assigns it to the AgreementStatus field.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) SetAgreementStatus(v string) {
	o.AgreementStatus = &v
}

// GetAlipayUserId returns the AlipayUserId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetAlipayUserId() string {
	if o == nil || IsNil(o.AlipayUserId) {
		var ret string
		return ret
	}
	return *o.AlipayUserId
}

// GetAlipayUserIdOk returns a tuple with the AlipayUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetAlipayUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayUserId) {
		return nil, false
	}
	return o.AlipayUserId, true
}

// HasAlipayUserId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) HasAlipayUserId() bool {
	if o != nil && !IsNil(o.AlipayUserId) {
		return true
	}

	return false
}

// SetAlipayUserId gets a reference to the given string and assigns it to the AlipayUserId field.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) SetAlipayUserId(v string) {
	o.AlipayUserId = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetBizType() string {
	if o == nil || IsNil(o.BizType) {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) HasBizType() bool {
	if o != nil && !IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) SetBizType(v string) {
	o.BizType = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetSignTime returns the SignTime field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetSignTime() string {
	if o == nil || IsNil(o.SignTime) {
		var ret string
		return ret
	}
	return *o.SignTime
}

// GetSignTimeOk returns a tuple with the SignTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetSignTimeOk() (*string, bool) {
	if o == nil || IsNil(o.SignTime) {
		return nil, false
	}
	return o.SignTime, true
}

// HasSignTime returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) HasSignTime() bool {
	if o != nil && !IsNil(o.SignTime) {
		return true
	}

	return false
}

// SetSignTime gets a reference to the given string and assigns it to the SignTime field.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) SetSignTime(v string) {
	o.SignTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ZhimaCreditPeZmgoAgreementQueryResponseModel) SetStartTime(v string) {
	o.StartTime = &v
}

func (o ZhimaCreditPeZmgoAgreementQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaCreditPeZmgoAgreementQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.AgreementName) {
		toSerialize["agreement_name"] = o.AgreementName
	}
	if !IsNil(o.AgreementStatus) {
		toSerialize["agreement_status"] = o.AgreementStatus
	}
	if !IsNil(o.AlipayUserId) {
		toSerialize["alipay_user_id"] = o.AlipayUserId
	}
	if !IsNil(o.BizType) {
		toSerialize["biz_type"] = o.BizType
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.SignTime) {
		toSerialize["sign_time"] = o.SignTime
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableZhimaCreditPeZmgoAgreementQueryResponseModel struct {
	value *ZhimaCreditPeZmgoAgreementQueryResponseModel
	isSet bool
}

func (v NullableZhimaCreditPeZmgoAgreementQueryResponseModel) Get() *ZhimaCreditPeZmgoAgreementQueryResponseModel {
	return v.value
}

func (v *NullableZhimaCreditPeZmgoAgreementQueryResponseModel) Set(val *ZhimaCreditPeZmgoAgreementQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCreditPeZmgoAgreementQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCreditPeZmgoAgreementQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCreditPeZmgoAgreementQueryResponseModel(val *ZhimaCreditPeZmgoAgreementQueryResponseModel) *NullableZhimaCreditPeZmgoAgreementQueryResponseModel {
	return &NullableZhimaCreditPeZmgoAgreementQueryResponseModel{value: val, isSet: true}
}

func (v NullableZhimaCreditPeZmgoAgreementQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCreditPeZmgoAgreementQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

