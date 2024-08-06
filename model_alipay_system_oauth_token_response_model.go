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

// checks if the AlipaySystemOauthTokenResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipaySystemOauthTokenResponseModel{}

// AlipaySystemOauthTokenResponseModel struct for AlipaySystemOauthTokenResponseModel
type AlipaySystemOauthTokenResponseModel struct {
	// 访问令牌。通过该令牌调用需要授权类接口
	AccessToken *string `json:"access_token,omitempty"`
	// 已废弃，请勿使用
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// 授权token开始时间，作为有效期计算的起点
	AuthStart *string `json:"auth_start,omitempty"`
	// 令牌类型，permanent表示返回的access_token和refresh_token永久有效，非永久令牌不返回该字段
	AuthTokenType *string `json:"auth_token_type,omitempty"`
	// 访问令牌的有效时间，单位是秒。
	ExpiresIn *string `json:"expires_in,omitempty"`
	// 支付宝用户唯一标识
	OpenId *string `json:"open_id,omitempty"`
	// 刷新令牌的有效时间，单位是秒。
	ReExpiresIn *string `json:"re_expires_in,omitempty"`
	// 刷新令牌。通过该令牌可以刷新access_token
	RefreshToken *string `json:"refresh_token,omitempty"`
	// union_id是支付宝用户在开放平台的唯一标识符，在配置应用分组后会返回该值。 同一用户的union_id在同一分组内应用保持一致。
	UnionId *string `json:"union_id,omitempty"`
	// 支付宝用户的唯一标识。以2088开头的16位数字。
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipaySystemOauthTokenResponseModel instantiates a new AlipaySystemOauthTokenResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipaySystemOauthTokenResponseModel() *AlipaySystemOauthTokenResponseModel {
	this := AlipaySystemOauthTokenResponseModel{}
	return &this
}

// NewAlipaySystemOauthTokenResponseModelWithDefaults instantiates a new AlipaySystemOauthTokenResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipaySystemOauthTokenResponseModelWithDefaults() *AlipaySystemOauthTokenResponseModel {
	this := AlipaySystemOauthTokenResponseModel{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AlipaySystemOauthTokenResponseModel) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySystemOauthTokenResponseModel) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AlipaySystemOauthTokenResponseModel) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AlipaySystemOauthTokenResponseModel) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetAlipayUserId returns the AlipayUserId field value if set, zero value otherwise.
func (o *AlipaySystemOauthTokenResponseModel) GetAlipayUserId() string {
	if o == nil || IsNil(o.AlipayUserId) {
		var ret string
		return ret
	}
	return *o.AlipayUserId
}

// GetAlipayUserIdOk returns a tuple with the AlipayUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySystemOauthTokenResponseModel) GetAlipayUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayUserId) {
		return nil, false
	}
	return o.AlipayUserId, true
}

// HasAlipayUserId returns a boolean if a field has been set.
func (o *AlipaySystemOauthTokenResponseModel) HasAlipayUserId() bool {
	if o != nil && !IsNil(o.AlipayUserId) {
		return true
	}

	return false
}

// SetAlipayUserId gets a reference to the given string and assigns it to the AlipayUserId field.
func (o *AlipaySystemOauthTokenResponseModel) SetAlipayUserId(v string) {
	o.AlipayUserId = &v
}

// GetAuthStart returns the AuthStart field value if set, zero value otherwise.
func (o *AlipaySystemOauthTokenResponseModel) GetAuthStart() string {
	if o == nil || IsNil(o.AuthStart) {
		var ret string
		return ret
	}
	return *o.AuthStart
}

// GetAuthStartOk returns a tuple with the AuthStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySystemOauthTokenResponseModel) GetAuthStartOk() (*string, bool) {
	if o == nil || IsNil(o.AuthStart) {
		return nil, false
	}
	return o.AuthStart, true
}

// HasAuthStart returns a boolean if a field has been set.
func (o *AlipaySystemOauthTokenResponseModel) HasAuthStart() bool {
	if o != nil && !IsNil(o.AuthStart) {
		return true
	}

	return false
}

// SetAuthStart gets a reference to the given string and assigns it to the AuthStart field.
func (o *AlipaySystemOauthTokenResponseModel) SetAuthStart(v string) {
	o.AuthStart = &v
}

// GetAuthTokenType returns the AuthTokenType field value if set, zero value otherwise.
func (o *AlipaySystemOauthTokenResponseModel) GetAuthTokenType() string {
	if o == nil || IsNil(o.AuthTokenType) {
		var ret string
		return ret
	}
	return *o.AuthTokenType
}

// GetAuthTokenTypeOk returns a tuple with the AuthTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySystemOauthTokenResponseModel) GetAuthTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthTokenType) {
		return nil, false
	}
	return o.AuthTokenType, true
}

// HasAuthTokenType returns a boolean if a field has been set.
func (o *AlipaySystemOauthTokenResponseModel) HasAuthTokenType() bool {
	if o != nil && !IsNil(o.AuthTokenType) {
		return true
	}

	return false
}

// SetAuthTokenType gets a reference to the given string and assigns it to the AuthTokenType field.
func (o *AlipaySystemOauthTokenResponseModel) SetAuthTokenType(v string) {
	o.AuthTokenType = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *AlipaySystemOauthTokenResponseModel) GetExpiresIn() string {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret string
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySystemOauthTokenResponseModel) GetExpiresInOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *AlipaySystemOauthTokenResponseModel) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given string and assigns it to the ExpiresIn field.
func (o *AlipaySystemOauthTokenResponseModel) SetExpiresIn(v string) {
	o.ExpiresIn = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipaySystemOauthTokenResponseModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySystemOauthTokenResponseModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipaySystemOauthTokenResponseModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipaySystemOauthTokenResponseModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetReExpiresIn returns the ReExpiresIn field value if set, zero value otherwise.
func (o *AlipaySystemOauthTokenResponseModel) GetReExpiresIn() string {
	if o == nil || IsNil(o.ReExpiresIn) {
		var ret string
		return ret
	}
	return *o.ReExpiresIn
}

// GetReExpiresInOk returns a tuple with the ReExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySystemOauthTokenResponseModel) GetReExpiresInOk() (*string, bool) {
	if o == nil || IsNil(o.ReExpiresIn) {
		return nil, false
	}
	return o.ReExpiresIn, true
}

// HasReExpiresIn returns a boolean if a field has been set.
func (o *AlipaySystemOauthTokenResponseModel) HasReExpiresIn() bool {
	if o != nil && !IsNil(o.ReExpiresIn) {
		return true
	}

	return false
}

// SetReExpiresIn gets a reference to the given string and assigns it to the ReExpiresIn field.
func (o *AlipaySystemOauthTokenResponseModel) SetReExpiresIn(v string) {
	o.ReExpiresIn = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *AlipaySystemOauthTokenResponseModel) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySystemOauthTokenResponseModel) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *AlipaySystemOauthTokenResponseModel) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *AlipaySystemOauthTokenResponseModel) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetUnionId returns the UnionId field value if set, zero value otherwise.
func (o *AlipaySystemOauthTokenResponseModel) GetUnionId() string {
	if o == nil || IsNil(o.UnionId) {
		var ret string
		return ret
	}
	return *o.UnionId
}

// GetUnionIdOk returns a tuple with the UnionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySystemOauthTokenResponseModel) GetUnionIdOk() (*string, bool) {
	if o == nil || IsNil(o.UnionId) {
		return nil, false
	}
	return o.UnionId, true
}

// HasUnionId returns a boolean if a field has been set.
func (o *AlipaySystemOauthTokenResponseModel) HasUnionId() bool {
	if o != nil && !IsNil(o.UnionId) {
		return true
	}

	return false
}

// SetUnionId gets a reference to the given string and assigns it to the UnionId field.
func (o *AlipaySystemOauthTokenResponseModel) SetUnionId(v string) {
	o.UnionId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipaySystemOauthTokenResponseModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySystemOauthTokenResponseModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipaySystemOauthTokenResponseModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipaySystemOauthTokenResponseModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipaySystemOauthTokenResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipaySystemOauthTokenResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !IsNil(o.AlipayUserId) {
		toSerialize["alipay_user_id"] = o.AlipayUserId
	}
	if !IsNil(o.AuthStart) {
		toSerialize["auth_start"] = o.AuthStart
	}
	if !IsNil(o.AuthTokenType) {
		toSerialize["auth_token_type"] = o.AuthTokenType
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.ReExpiresIn) {
		toSerialize["re_expires_in"] = o.ReExpiresIn
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if !IsNil(o.UnionId) {
		toSerialize["union_id"] = o.UnionId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipaySystemOauthTokenResponseModel struct {
	value *AlipaySystemOauthTokenResponseModel
	isSet bool
}

func (v NullableAlipaySystemOauthTokenResponseModel) Get() *AlipaySystemOauthTokenResponseModel {
	return v.value
}

func (v *NullableAlipaySystemOauthTokenResponseModel) Set(val *AlipaySystemOauthTokenResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipaySystemOauthTokenResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipaySystemOauthTokenResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipaySystemOauthTokenResponseModel(val *AlipaySystemOauthTokenResponseModel) *NullableAlipaySystemOauthTokenResponseModel {
	return &NullableAlipaySystemOauthTokenResponseModel{value: val, isSet: true}
}

func (v NullableAlipaySystemOauthTokenResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipaySystemOauthTokenResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

