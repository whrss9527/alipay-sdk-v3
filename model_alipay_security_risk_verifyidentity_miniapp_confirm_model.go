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

// checks if the AlipaySecurityRiskVerifyidentityMiniappConfirmModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipaySecurityRiskVerifyidentityMiniappConfirmModel{}

// AlipaySecurityRiskVerifyidentityMiniappConfirmModel struct for AlipaySecurityRiskVerifyidentityMiniappConfirmModel
type AlipaySecurityRiskVerifyidentityMiniappConfirmModel struct {
	// 业务的业务流水号，例如订单id
	Challenge *string `json:"challenge,omitempty"`
	// 蚂蚁统一会员ID，验证时会与做身份认证时候会话中的openid做比对，如不相符核验结果不通过
	OpenId *string `json:"open_id,omitempty"`
	// 蚂蚁统一会员ID，验证时会与做身份认证时候会话中的userId做比对，如不相符核验结果不通过
	UserId *string `json:"user_id,omitempty"`
	// 核身校验id，是一次核身校验服务中唯一性的id
	VerifyId *string `json:"verify_id,omitempty"`
}

// NewAlipaySecurityRiskVerifyidentityMiniappConfirmModel instantiates a new AlipaySecurityRiskVerifyidentityMiniappConfirmModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipaySecurityRiskVerifyidentityMiniappConfirmModel() *AlipaySecurityRiskVerifyidentityMiniappConfirmModel {
	this := AlipaySecurityRiskVerifyidentityMiniappConfirmModel{}
	return &this
}

// NewAlipaySecurityRiskVerifyidentityMiniappConfirmModelWithDefaults instantiates a new AlipaySecurityRiskVerifyidentityMiniappConfirmModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipaySecurityRiskVerifyidentityMiniappConfirmModelWithDefaults() *AlipaySecurityRiskVerifyidentityMiniappConfirmModel {
	this := AlipaySecurityRiskVerifyidentityMiniappConfirmModel{}
	return &this
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) GetChallenge() string {
	if o == nil || IsNil(o.Challenge) {
		var ret string
		return ret
	}
	return *o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) GetChallengeOk() (*string, bool) {
	if o == nil || IsNil(o.Challenge) {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) HasChallenge() bool {
	if o != nil && !IsNil(o.Challenge) {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given string and assigns it to the Challenge field.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) SetChallenge(v string) {
	o.Challenge = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) SetUserId(v string) {
	o.UserId = &v
}

// GetVerifyId returns the VerifyId field value if set, zero value otherwise.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) GetVerifyId() string {
	if o == nil || IsNil(o.VerifyId) {
		var ret string
		return ret
	}
	return *o.VerifyId
}

// GetVerifyIdOk returns a tuple with the VerifyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) GetVerifyIdOk() (*string, bool) {
	if o == nil || IsNil(o.VerifyId) {
		return nil, false
	}
	return o.VerifyId, true
}

// HasVerifyId returns a boolean if a field has been set.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) HasVerifyId() bool {
	if o != nil && !IsNil(o.VerifyId) {
		return true
	}

	return false
}

// SetVerifyId gets a reference to the given string and assigns it to the VerifyId field.
func (o *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) SetVerifyId(v string) {
	o.VerifyId = &v
}

func (o AlipaySecurityRiskVerifyidentityMiniappConfirmModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipaySecurityRiskVerifyidentityMiniappConfirmModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Challenge) {
		toSerialize["challenge"] = o.Challenge
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.VerifyId) {
		toSerialize["verify_id"] = o.VerifyId
	}
	return toSerialize, nil
}

type NullableAlipaySecurityRiskVerifyidentityMiniappConfirmModel struct {
	value *AlipaySecurityRiskVerifyidentityMiniappConfirmModel
	isSet bool
}

func (v NullableAlipaySecurityRiskVerifyidentityMiniappConfirmModel) Get() *AlipaySecurityRiskVerifyidentityMiniappConfirmModel {
	return v.value
}

func (v *NullableAlipaySecurityRiskVerifyidentityMiniappConfirmModel) Set(val *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipaySecurityRiskVerifyidentityMiniappConfirmModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipaySecurityRiskVerifyidentityMiniappConfirmModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipaySecurityRiskVerifyidentityMiniappConfirmModel(val *AlipaySecurityRiskVerifyidentityMiniappConfirmModel) *NullableAlipaySecurityRiskVerifyidentityMiniappConfirmModel {
	return &NullableAlipaySecurityRiskVerifyidentityMiniappConfirmModel{value: val, isSet: true}
}

func (v NullableAlipaySecurityRiskVerifyidentityMiniappConfirmModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipaySecurityRiskVerifyidentityMiniappConfirmModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
