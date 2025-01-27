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

// checks if the UserTradeInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTradeInfoDTO{}

// UserTradeInfoDTO struct for UserTradeInfoDTO
type UserTradeInfoDTO struct {
	// 商品信息列表
	GoodsInfoList []GoodsInfoDTO `json:"goods_info_list,omitempty"`
	// 门店ISV的PID
	PartnerId *string `json:"partner_id,omitempty"`
	// 是否是风险交易：NO_RISK-无风险；POTENTIAL_RISK-潜在风险（中等风险）；HIGH_RISK-高风险
	RiskLevel *string `json:"risk_level,omitempty"`
	// 交易总金额，单位元，精确两位小数点
	TradeAmount *string `json:"trade_amount,omitempty"`
	// 交易单号
	TradeNo *string `json:"trade_no,omitempty"`
	// 交易时间
	TradeTime *string `json:"trade_time,omitempty"`
	// 商品数据会根据活动商品列表进行过滤，该字段代表未过滤的商品列表大小
	UnfilteredTotalGoodsCount *int32 `json:"unfiltered_total_goods_count,omitempty"`
	// 消费者支付宝ID
	UserId *string `json:"user_id,omitempty"`
}

// NewUserTradeInfoDTO instantiates a new UserTradeInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTradeInfoDTO() *UserTradeInfoDTO {
	this := UserTradeInfoDTO{}
	return &this
}

// NewUserTradeInfoDTOWithDefaults instantiates a new UserTradeInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTradeInfoDTOWithDefaults() *UserTradeInfoDTO {
	this := UserTradeInfoDTO{}
	return &this
}

// GetGoodsInfoList returns the GoodsInfoList field value if set, zero value otherwise.
func (o *UserTradeInfoDTO) GetGoodsInfoList() []GoodsInfoDTO {
	if o == nil || IsNil(o.GoodsInfoList) {
		var ret []GoodsInfoDTO
		return ret
	}
	return o.GoodsInfoList
}

// GetGoodsInfoListOk returns a tuple with the GoodsInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTradeInfoDTO) GetGoodsInfoListOk() ([]GoodsInfoDTO, bool) {
	if o == nil || IsNil(o.GoodsInfoList) {
		return nil, false
	}
	return o.GoodsInfoList, true
}

// HasGoodsInfoList returns a boolean if a field has been set.
func (o *UserTradeInfoDTO) HasGoodsInfoList() bool {
	if o != nil && !IsNil(o.GoodsInfoList) {
		return true
	}

	return false
}

// SetGoodsInfoList gets a reference to the given []GoodsInfoDTO and assigns it to the GoodsInfoList field.
func (o *UserTradeInfoDTO) SetGoodsInfoList(v []GoodsInfoDTO) {
	o.GoodsInfoList = v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *UserTradeInfoDTO) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTradeInfoDTO) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *UserTradeInfoDTO) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *UserTradeInfoDTO) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *UserTradeInfoDTO) GetRiskLevel() string {
	if o == nil || IsNil(o.RiskLevel) {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTradeInfoDTO) GetRiskLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RiskLevel) {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *UserTradeInfoDTO) HasRiskLevel() bool {
	if o != nil && !IsNil(o.RiskLevel) {
		return true
	}

	return false
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *UserTradeInfoDTO) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// GetTradeAmount returns the TradeAmount field value if set, zero value otherwise.
func (o *UserTradeInfoDTO) GetTradeAmount() string {
	if o == nil || IsNil(o.TradeAmount) {
		var ret string
		return ret
	}
	return *o.TradeAmount
}

// GetTradeAmountOk returns a tuple with the TradeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTradeInfoDTO) GetTradeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TradeAmount) {
		return nil, false
	}
	return o.TradeAmount, true
}

// HasTradeAmount returns a boolean if a field has been set.
func (o *UserTradeInfoDTO) HasTradeAmount() bool {
	if o != nil && !IsNil(o.TradeAmount) {
		return true
	}

	return false
}

// SetTradeAmount gets a reference to the given string and assigns it to the TradeAmount field.
func (o *UserTradeInfoDTO) SetTradeAmount(v string) {
	o.TradeAmount = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *UserTradeInfoDTO) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTradeInfoDTO) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *UserTradeInfoDTO) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *UserTradeInfoDTO) SetTradeNo(v string) {
	o.TradeNo = &v
}

// GetTradeTime returns the TradeTime field value if set, zero value otherwise.
func (o *UserTradeInfoDTO) GetTradeTime() string {
	if o == nil || IsNil(o.TradeTime) {
		var ret string
		return ret
	}
	return *o.TradeTime
}

// GetTradeTimeOk returns a tuple with the TradeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTradeInfoDTO) GetTradeTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TradeTime) {
		return nil, false
	}
	return o.TradeTime, true
}

// HasTradeTime returns a boolean if a field has been set.
func (o *UserTradeInfoDTO) HasTradeTime() bool {
	if o != nil && !IsNil(o.TradeTime) {
		return true
	}

	return false
}

// SetTradeTime gets a reference to the given string and assigns it to the TradeTime field.
func (o *UserTradeInfoDTO) SetTradeTime(v string) {
	o.TradeTime = &v
}

// GetUnfilteredTotalGoodsCount returns the UnfilteredTotalGoodsCount field value if set, zero value otherwise.
func (o *UserTradeInfoDTO) GetUnfilteredTotalGoodsCount() int32 {
	if o == nil || IsNil(o.UnfilteredTotalGoodsCount) {
		var ret int32
		return ret
	}
	return *o.UnfilteredTotalGoodsCount
}

// GetUnfilteredTotalGoodsCountOk returns a tuple with the UnfilteredTotalGoodsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTradeInfoDTO) GetUnfilteredTotalGoodsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UnfilteredTotalGoodsCount) {
		return nil, false
	}
	return o.UnfilteredTotalGoodsCount, true
}

// HasUnfilteredTotalGoodsCount returns a boolean if a field has been set.
func (o *UserTradeInfoDTO) HasUnfilteredTotalGoodsCount() bool {
	if o != nil && !IsNil(o.UnfilteredTotalGoodsCount) {
		return true
	}

	return false
}

// SetUnfilteredTotalGoodsCount gets a reference to the given int32 and assigns it to the UnfilteredTotalGoodsCount field.
func (o *UserTradeInfoDTO) SetUnfilteredTotalGoodsCount(v int32) {
	o.UnfilteredTotalGoodsCount = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserTradeInfoDTO) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTradeInfoDTO) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserTradeInfoDTO) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserTradeInfoDTO) SetUserId(v string) {
	o.UserId = &v
}

func (o UserTradeInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTradeInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GoodsInfoList) {
		toSerialize["goods_info_list"] = o.GoodsInfoList
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partner_id"] = o.PartnerId
	}
	if !IsNil(o.RiskLevel) {
		toSerialize["risk_level"] = o.RiskLevel
	}
	if !IsNil(o.TradeAmount) {
		toSerialize["trade_amount"] = o.TradeAmount
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	if !IsNil(o.TradeTime) {
		toSerialize["trade_time"] = o.TradeTime
	}
	if !IsNil(o.UnfilteredTotalGoodsCount) {
		toSerialize["unfiltered_total_goods_count"] = o.UnfilteredTotalGoodsCount
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUserTradeInfoDTO struct {
	value *UserTradeInfoDTO
	isSet bool
}

func (v NullableUserTradeInfoDTO) Get() *UserTradeInfoDTO {
	return v.value
}

func (v *NullableUserTradeInfoDTO) Set(val *UserTradeInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTradeInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTradeInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTradeInfoDTO(val *UserTradeInfoDTO) *NullableUserTradeInfoDTO {
	return &NullableUserTradeInfoDTO{value: val, isSet: true}
}

func (v NullableUserTradeInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTradeInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
