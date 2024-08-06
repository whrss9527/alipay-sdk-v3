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

// checks if the AlipayTradeAdvanceConsultResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayTradeAdvanceConsultResponseModel{}

// AlipayTradeAdvanceConsultResponseModel struct for AlipayTradeAdvanceConsultResponseModel
type AlipayTradeAdvanceConsultResponseModel struct {
	// true 代表当前时间点，用户允许垫资 false 代表当前时间，用户不允许垫资
	ReferResult *bool `json:"refer_result,omitempty"`
	// 用户被注销
	ResultCode *string `json:"result_code,omitempty"`
	// 返回用户不准入原因
	ResultMessage *string `json:"result_message,omitempty"`
	// 订单风险评估等级，在单笔订单风险预评估时返回。当基础风险校验通过时，可通过该值获取业务风险评估等级。取值：2-高风险；1-低风险。
	RiskLevel *string `json:"risk_level,omitempty"`
	UserRiskPrediction *UserRiskPrediction `json:"user_risk_prediction,omitempty"`
	// 用户剩余的总待还金额，无论当前用户是否允许垫资，都会返回该属性。
	WaitRepaymentAmount *string `json:"wait_repayment_amount,omitempty"`
	// 用户总的未还的垫资笔数，无论用户是否允许垫资，都会返回该属性
	WaitRepaymentOrderCount *int32 `json:"wait_repayment_order_count,omitempty"`
	// 待还订单列表，无论用户当前状态是否允许垫资，都会返回当前用户在商户下的待还订单信息
	WaitRepaymentOrderInfos []WaitRepaymentOrderInfo `json:"wait_repayment_order_infos,omitempty"`
}

// NewAlipayTradeAdvanceConsultResponseModel instantiates a new AlipayTradeAdvanceConsultResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayTradeAdvanceConsultResponseModel() *AlipayTradeAdvanceConsultResponseModel {
	this := AlipayTradeAdvanceConsultResponseModel{}
	return &this
}

// NewAlipayTradeAdvanceConsultResponseModelWithDefaults instantiates a new AlipayTradeAdvanceConsultResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayTradeAdvanceConsultResponseModelWithDefaults() *AlipayTradeAdvanceConsultResponseModel {
	this := AlipayTradeAdvanceConsultResponseModel{}
	return &this
}

// GetReferResult returns the ReferResult field value if set, zero value otherwise.
func (o *AlipayTradeAdvanceConsultResponseModel) GetReferResult() bool {
	if o == nil || IsNil(o.ReferResult) {
		var ret bool
		return ret
	}
	return *o.ReferResult
}

// GetReferResultOk returns a tuple with the ReferResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) GetReferResultOk() (*bool, bool) {
	if o == nil || IsNil(o.ReferResult) {
		return nil, false
	}
	return o.ReferResult, true
}

// HasReferResult returns a boolean if a field has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) HasReferResult() bool {
	if o != nil && !IsNil(o.ReferResult) {
		return true
	}

	return false
}

// SetReferResult gets a reference to the given bool and assigns it to the ReferResult field.
func (o *AlipayTradeAdvanceConsultResponseModel) SetReferResult(v bool) {
	o.ReferResult = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *AlipayTradeAdvanceConsultResponseModel) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *AlipayTradeAdvanceConsultResponseModel) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *AlipayTradeAdvanceConsultResponseModel) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *AlipayTradeAdvanceConsultResponseModel) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *AlipayTradeAdvanceConsultResponseModel) GetRiskLevel() string {
	if o == nil || IsNil(o.RiskLevel) {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) GetRiskLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RiskLevel) {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) HasRiskLevel() bool {
	if o != nil && !IsNil(o.RiskLevel) {
		return true
	}

	return false
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *AlipayTradeAdvanceConsultResponseModel) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// GetUserRiskPrediction returns the UserRiskPrediction field value if set, zero value otherwise.
func (o *AlipayTradeAdvanceConsultResponseModel) GetUserRiskPrediction() UserRiskPrediction {
	if o == nil || IsNil(o.UserRiskPrediction) {
		var ret UserRiskPrediction
		return ret
	}
	return *o.UserRiskPrediction
}

// GetUserRiskPredictionOk returns a tuple with the UserRiskPrediction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) GetUserRiskPredictionOk() (*UserRiskPrediction, bool) {
	if o == nil || IsNil(o.UserRiskPrediction) {
		return nil, false
	}
	return o.UserRiskPrediction, true
}

// HasUserRiskPrediction returns a boolean if a field has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) HasUserRiskPrediction() bool {
	if o != nil && !IsNil(o.UserRiskPrediction) {
		return true
	}

	return false
}

// SetUserRiskPrediction gets a reference to the given UserRiskPrediction and assigns it to the UserRiskPrediction field.
func (o *AlipayTradeAdvanceConsultResponseModel) SetUserRiskPrediction(v UserRiskPrediction) {
	o.UserRiskPrediction = &v
}

// GetWaitRepaymentAmount returns the WaitRepaymentAmount field value if set, zero value otherwise.
func (o *AlipayTradeAdvanceConsultResponseModel) GetWaitRepaymentAmount() string {
	if o == nil || IsNil(o.WaitRepaymentAmount) {
		var ret string
		return ret
	}
	return *o.WaitRepaymentAmount
}

// GetWaitRepaymentAmountOk returns a tuple with the WaitRepaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) GetWaitRepaymentAmountOk() (*string, bool) {
	if o == nil || IsNil(o.WaitRepaymentAmount) {
		return nil, false
	}
	return o.WaitRepaymentAmount, true
}

// HasWaitRepaymentAmount returns a boolean if a field has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) HasWaitRepaymentAmount() bool {
	if o != nil && !IsNil(o.WaitRepaymentAmount) {
		return true
	}

	return false
}

// SetWaitRepaymentAmount gets a reference to the given string and assigns it to the WaitRepaymentAmount field.
func (o *AlipayTradeAdvanceConsultResponseModel) SetWaitRepaymentAmount(v string) {
	o.WaitRepaymentAmount = &v
}

// GetWaitRepaymentOrderCount returns the WaitRepaymentOrderCount field value if set, zero value otherwise.
func (o *AlipayTradeAdvanceConsultResponseModel) GetWaitRepaymentOrderCount() int32 {
	if o == nil || IsNil(o.WaitRepaymentOrderCount) {
		var ret int32
		return ret
	}
	return *o.WaitRepaymentOrderCount
}

// GetWaitRepaymentOrderCountOk returns a tuple with the WaitRepaymentOrderCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) GetWaitRepaymentOrderCountOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitRepaymentOrderCount) {
		return nil, false
	}
	return o.WaitRepaymentOrderCount, true
}

// HasWaitRepaymentOrderCount returns a boolean if a field has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) HasWaitRepaymentOrderCount() bool {
	if o != nil && !IsNil(o.WaitRepaymentOrderCount) {
		return true
	}

	return false
}

// SetWaitRepaymentOrderCount gets a reference to the given int32 and assigns it to the WaitRepaymentOrderCount field.
func (o *AlipayTradeAdvanceConsultResponseModel) SetWaitRepaymentOrderCount(v int32) {
	o.WaitRepaymentOrderCount = &v
}

// GetWaitRepaymentOrderInfos returns the WaitRepaymentOrderInfos field value if set, zero value otherwise.
func (o *AlipayTradeAdvanceConsultResponseModel) GetWaitRepaymentOrderInfos() []WaitRepaymentOrderInfo {
	if o == nil || IsNil(o.WaitRepaymentOrderInfos) {
		var ret []WaitRepaymentOrderInfo
		return ret
	}
	return o.WaitRepaymentOrderInfos
}

// GetWaitRepaymentOrderInfosOk returns a tuple with the WaitRepaymentOrderInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) GetWaitRepaymentOrderInfosOk() ([]WaitRepaymentOrderInfo, bool) {
	if o == nil || IsNil(o.WaitRepaymentOrderInfos) {
		return nil, false
	}
	return o.WaitRepaymentOrderInfos, true
}

// HasWaitRepaymentOrderInfos returns a boolean if a field has been set.
func (o *AlipayTradeAdvanceConsultResponseModel) HasWaitRepaymentOrderInfos() bool {
	if o != nil && !IsNil(o.WaitRepaymentOrderInfos) {
		return true
	}

	return false
}

// SetWaitRepaymentOrderInfos gets a reference to the given []WaitRepaymentOrderInfo and assigns it to the WaitRepaymentOrderInfos field.
func (o *AlipayTradeAdvanceConsultResponseModel) SetWaitRepaymentOrderInfos(v []WaitRepaymentOrderInfo) {
	o.WaitRepaymentOrderInfos = v
}

func (o AlipayTradeAdvanceConsultResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayTradeAdvanceConsultResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferResult) {
		toSerialize["refer_result"] = o.ReferResult
	}
	if !IsNil(o.ResultCode) {
		toSerialize["result_code"] = o.ResultCode
	}
	if !IsNil(o.ResultMessage) {
		toSerialize["result_message"] = o.ResultMessage
	}
	if !IsNil(o.RiskLevel) {
		toSerialize["risk_level"] = o.RiskLevel
	}
	if !IsNil(o.UserRiskPrediction) {
		toSerialize["user_risk_prediction"] = o.UserRiskPrediction
	}
	if !IsNil(o.WaitRepaymentAmount) {
		toSerialize["wait_repayment_amount"] = o.WaitRepaymentAmount
	}
	if !IsNil(o.WaitRepaymentOrderCount) {
		toSerialize["wait_repayment_order_count"] = o.WaitRepaymentOrderCount
	}
	if !IsNil(o.WaitRepaymentOrderInfos) {
		toSerialize["wait_repayment_order_infos"] = o.WaitRepaymentOrderInfos
	}
	return toSerialize, nil
}

type NullableAlipayTradeAdvanceConsultResponseModel struct {
	value *AlipayTradeAdvanceConsultResponseModel
	isSet bool
}

func (v NullableAlipayTradeAdvanceConsultResponseModel) Get() *AlipayTradeAdvanceConsultResponseModel {
	return v.value
}

func (v *NullableAlipayTradeAdvanceConsultResponseModel) Set(val *AlipayTradeAdvanceConsultResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeAdvanceConsultResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeAdvanceConsultResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeAdvanceConsultResponseModel(val *AlipayTradeAdvanceConsultResponseModel) *NullableAlipayTradeAdvanceConsultResponseModel {
	return &NullableAlipayTradeAdvanceConsultResponseModel{value: val, isSet: true}
}

func (v NullableAlipayTradeAdvanceConsultResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeAdvanceConsultResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

