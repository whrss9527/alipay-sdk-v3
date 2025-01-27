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

// checks if the SettleDetailInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettleDetailInfo{}

// SettleDetailInfo struct for SettleDetailInfo
type SettleDetailInfo struct {
	// 仅在直付通账期模式下，当一笔交易需要分多次发起部分确认结算时使用，表示本次确认结算的实际结算金额。传递本字段后，原amount字段不再生效，结算金额以本字段为准。如已经发生过部分确认结算、不传递本字段则默认按剩余待结算金额一次性结算。
	ActualAmount *string `json:"actual_amount,omitempty"`
	// 结算的金额，单位为元。在创建订单和支付接口时必须和交易金额相同。在结算确认接口时必须等于交易金额减去已退款金额。直付通账期模式下，如使用部分结算能力、传递了actual_amount字段，则忽略本字段的校验、可不传。
	Amount *string `json:"amount,omitempty"`
	// 结算主体标识。当结算主体类型为SecondMerchant时，为二级商户的SecondMerchantID；当结算主体类型为Store时，为门店的外标。
	SettleEntityId *string `json:"settle_entity_id,omitempty"`
	// 结算主体类型。
	SettleEntityType *string `json:"settle_entity_type,omitempty"`
	// 结算汇总维度，按照这个维度汇总成批次结算，由商户指定。  目前需要和结算收款方账户类型为cardAliasNo配合使用
	SummaryDimension *string `json:"summary_dimension,omitempty"`
	// 结算收款方。当结算收款方类型是cardAliasNo时，本参数为用户在支付宝绑定的卡编号；结算收款方类型是userId时，本参数为用户的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；当结算收款方类型是loginName时，本参数为用户的支付宝登录号；当结算收款方类型是defaultSettle时，本参数不能传值，保持为空。
	TransIn *string `json:"trans_in,omitempty"`
	// 结算收款方的账户类型。
	TransInType *string `json:"trans_in_type,omitempty"`
}

// NewSettleDetailInfo instantiates a new SettleDetailInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettleDetailInfo() *SettleDetailInfo {
	this := SettleDetailInfo{}
	return &this
}

// NewSettleDetailInfoWithDefaults instantiates a new SettleDetailInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettleDetailInfoWithDefaults() *SettleDetailInfo {
	this := SettleDetailInfo{}
	return &this
}

// GetActualAmount returns the ActualAmount field value if set, zero value otherwise.
func (o *SettleDetailInfo) GetActualAmount() string {
	if o == nil || IsNil(o.ActualAmount) {
		var ret string
		return ret
	}
	return *o.ActualAmount
}

// GetActualAmountOk returns a tuple with the ActualAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleDetailInfo) GetActualAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ActualAmount) {
		return nil, false
	}
	return o.ActualAmount, true
}

// HasActualAmount returns a boolean if a field has been set.
func (o *SettleDetailInfo) HasActualAmount() bool {
	if o != nil && !IsNil(o.ActualAmount) {
		return true
	}

	return false
}

// SetActualAmount gets a reference to the given string and assigns it to the ActualAmount field.
func (o *SettleDetailInfo) SetActualAmount(v string) {
	o.ActualAmount = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SettleDetailInfo) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleDetailInfo) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SettleDetailInfo) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *SettleDetailInfo) SetAmount(v string) {
	o.Amount = &v
}

// GetSettleEntityId returns the SettleEntityId field value if set, zero value otherwise.
func (o *SettleDetailInfo) GetSettleEntityId() string {
	if o == nil || IsNil(o.SettleEntityId) {
		var ret string
		return ret
	}
	return *o.SettleEntityId
}

// GetSettleEntityIdOk returns a tuple with the SettleEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleDetailInfo) GetSettleEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.SettleEntityId) {
		return nil, false
	}
	return o.SettleEntityId, true
}

// HasSettleEntityId returns a boolean if a field has been set.
func (o *SettleDetailInfo) HasSettleEntityId() bool {
	if o != nil && !IsNil(o.SettleEntityId) {
		return true
	}

	return false
}

// SetSettleEntityId gets a reference to the given string and assigns it to the SettleEntityId field.
func (o *SettleDetailInfo) SetSettleEntityId(v string) {
	o.SettleEntityId = &v
}

// GetSettleEntityType returns the SettleEntityType field value if set, zero value otherwise.
func (o *SettleDetailInfo) GetSettleEntityType() string {
	if o == nil || IsNil(o.SettleEntityType) {
		var ret string
		return ret
	}
	return *o.SettleEntityType
}

// GetSettleEntityTypeOk returns a tuple with the SettleEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleDetailInfo) GetSettleEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SettleEntityType) {
		return nil, false
	}
	return o.SettleEntityType, true
}

// HasSettleEntityType returns a boolean if a field has been set.
func (o *SettleDetailInfo) HasSettleEntityType() bool {
	if o != nil && !IsNil(o.SettleEntityType) {
		return true
	}

	return false
}

// SetSettleEntityType gets a reference to the given string and assigns it to the SettleEntityType field.
func (o *SettleDetailInfo) SetSettleEntityType(v string) {
	o.SettleEntityType = &v
}

// GetSummaryDimension returns the SummaryDimension field value if set, zero value otherwise.
func (o *SettleDetailInfo) GetSummaryDimension() string {
	if o == nil || IsNil(o.SummaryDimension) {
		var ret string
		return ret
	}
	return *o.SummaryDimension
}

// GetSummaryDimensionOk returns a tuple with the SummaryDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleDetailInfo) GetSummaryDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryDimension) {
		return nil, false
	}
	return o.SummaryDimension, true
}

// HasSummaryDimension returns a boolean if a field has been set.
func (o *SettleDetailInfo) HasSummaryDimension() bool {
	if o != nil && !IsNil(o.SummaryDimension) {
		return true
	}

	return false
}

// SetSummaryDimension gets a reference to the given string and assigns it to the SummaryDimension field.
func (o *SettleDetailInfo) SetSummaryDimension(v string) {
	o.SummaryDimension = &v
}

// GetTransIn returns the TransIn field value if set, zero value otherwise.
func (o *SettleDetailInfo) GetTransIn() string {
	if o == nil || IsNil(o.TransIn) {
		var ret string
		return ret
	}
	return *o.TransIn
}

// GetTransInOk returns a tuple with the TransIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleDetailInfo) GetTransInOk() (*string, bool) {
	if o == nil || IsNil(o.TransIn) {
		return nil, false
	}
	return o.TransIn, true
}

// HasTransIn returns a boolean if a field has been set.
func (o *SettleDetailInfo) HasTransIn() bool {
	if o != nil && !IsNil(o.TransIn) {
		return true
	}

	return false
}

// SetTransIn gets a reference to the given string and assigns it to the TransIn field.
func (o *SettleDetailInfo) SetTransIn(v string) {
	o.TransIn = &v
}

// GetTransInType returns the TransInType field value if set, zero value otherwise.
func (o *SettleDetailInfo) GetTransInType() string {
	if o == nil || IsNil(o.TransInType) {
		var ret string
		return ret
	}
	return *o.TransInType
}

// GetTransInTypeOk returns a tuple with the TransInType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleDetailInfo) GetTransInTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransInType) {
		return nil, false
	}
	return o.TransInType, true
}

// HasTransInType returns a boolean if a field has been set.
func (o *SettleDetailInfo) HasTransInType() bool {
	if o != nil && !IsNil(o.TransInType) {
		return true
	}

	return false
}

// SetTransInType gets a reference to the given string and assigns it to the TransInType field.
func (o *SettleDetailInfo) SetTransInType(v string) {
	o.TransInType = &v
}

func (o SettleDetailInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettleDetailInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActualAmount) {
		toSerialize["actual_amount"] = o.ActualAmount
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.SettleEntityId) {
		toSerialize["settle_entity_id"] = o.SettleEntityId
	}
	if !IsNil(o.SettleEntityType) {
		toSerialize["settle_entity_type"] = o.SettleEntityType
	}
	if !IsNil(o.SummaryDimension) {
		toSerialize["summary_dimension"] = o.SummaryDimension
	}
	if !IsNil(o.TransIn) {
		toSerialize["trans_in"] = o.TransIn
	}
	if !IsNil(o.TransInType) {
		toSerialize["trans_in_type"] = o.TransInType
	}
	return toSerialize, nil
}

type NullableSettleDetailInfo struct {
	value *SettleDetailInfo
	isSet bool
}

func (v NullableSettleDetailInfo) Get() *SettleDetailInfo {
	return v.value
}

func (v *NullableSettleDetailInfo) Set(val *SettleDetailInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSettleDetailInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSettleDetailInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettleDetailInfo(val *SettleDetailInfo) *NullableSettleDetailInfo {
	return &NullableSettleDetailInfo{value: val, isSet: true}
}

func (v NullableSettleDetailInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettleDetailInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
