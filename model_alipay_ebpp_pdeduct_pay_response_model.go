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

// checks if the AlipayEbppPdeductPayResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppPdeductPayResponseModel{}

// AlipayEbppPdeductPayResponseModel struct for AlipayEbppPdeductPayResponseModel
type AlipayEbppPdeductPayResponseModel struct {
	// 支付宝代扣协议ID
	AgreementId *string `json:"agreement_id,omitempty"`
	// 支付宝订单流水号
	BillNo *string `json:"bill_no,omitempty"`
	// 扩展参数
	ExtendField *string `json:"extend_field,omitempty"`
	// 商户代扣业务流水
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// 针对于支付失败时，给的对应错误明细，如果判断外围的错误码是INVOKE_PAYACCEPTANCE_EXCEPTION需要近一步再结合着结果模型中的result_code, result_msg来判断
	ResultCode *string `json:"result_code,omitempty"`
	// 针对于支付失败时，给的对应错误明细，如果判断外围的错误码是INVOKE_PAYACCEPTANCE_EXCEPTION需要近一步再结合着结果模型中的result_code, result_msg来判断
	ResultMsg *string `json:"result_msg,omitempty"`
	// 订单支付状态。  0：未知  1：成功  2：失败
	ResultStatus *string `json:"result_status,omitempty"`
}

// NewAlipayEbppPdeductPayResponseModel instantiates a new AlipayEbppPdeductPayResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppPdeductPayResponseModel() *AlipayEbppPdeductPayResponseModel {
	this := AlipayEbppPdeductPayResponseModel{}
	return &this
}

// NewAlipayEbppPdeductPayResponseModelWithDefaults instantiates a new AlipayEbppPdeductPayResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppPdeductPayResponseModelWithDefaults() *AlipayEbppPdeductPayResponseModel {
	this := AlipayEbppPdeductPayResponseModel{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductPayResponseModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductPayResponseModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductPayResponseModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *AlipayEbppPdeductPayResponseModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetBillNo returns the BillNo field value if set, zero value otherwise.
func (o *AlipayEbppPdeductPayResponseModel) GetBillNo() string {
	if o == nil || IsNil(o.BillNo) {
		var ret string
		return ret
	}
	return *o.BillNo
}

// GetBillNoOk returns a tuple with the BillNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductPayResponseModel) GetBillNoOk() (*string, bool) {
	if o == nil || IsNil(o.BillNo) {
		return nil, false
	}
	return o.BillNo, true
}

// HasBillNo returns a boolean if a field has been set.
func (o *AlipayEbppPdeductPayResponseModel) HasBillNo() bool {
	if o != nil && !IsNil(o.BillNo) {
		return true
	}

	return false
}

// SetBillNo gets a reference to the given string and assigns it to the BillNo field.
func (o *AlipayEbppPdeductPayResponseModel) SetBillNo(v string) {
	o.BillNo = &v
}

// GetExtendField returns the ExtendField field value if set, zero value otherwise.
func (o *AlipayEbppPdeductPayResponseModel) GetExtendField() string {
	if o == nil || IsNil(o.ExtendField) {
		var ret string
		return ret
	}
	return *o.ExtendField
}

// GetExtendFieldOk returns a tuple with the ExtendField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductPayResponseModel) GetExtendFieldOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendField) {
		return nil, false
	}
	return o.ExtendField, true
}

// HasExtendField returns a boolean if a field has been set.
func (o *AlipayEbppPdeductPayResponseModel) HasExtendField() bool {
	if o != nil && !IsNil(o.ExtendField) {
		return true
	}

	return false
}

// SetExtendField gets a reference to the given string and assigns it to the ExtendField field.
func (o *AlipayEbppPdeductPayResponseModel) SetExtendField(v string) {
	o.ExtendField = &v
}

// GetOutOrderNo returns the OutOrderNo field value if set, zero value otherwise.
func (o *AlipayEbppPdeductPayResponseModel) GetOutOrderNo() string {
	if o == nil || IsNil(o.OutOrderNo) {
		var ret string
		return ret
	}
	return *o.OutOrderNo
}

// GetOutOrderNoOk returns a tuple with the OutOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductPayResponseModel) GetOutOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutOrderNo) {
		return nil, false
	}
	return o.OutOrderNo, true
}

// HasOutOrderNo returns a boolean if a field has been set.
func (o *AlipayEbppPdeductPayResponseModel) HasOutOrderNo() bool {
	if o != nil && !IsNil(o.OutOrderNo) {
		return true
	}

	return false
}

// SetOutOrderNo gets a reference to the given string and assigns it to the OutOrderNo field.
func (o *AlipayEbppPdeductPayResponseModel) SetOutOrderNo(v string) {
	o.OutOrderNo = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *AlipayEbppPdeductPayResponseModel) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductPayResponseModel) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *AlipayEbppPdeductPayResponseModel) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *AlipayEbppPdeductPayResponseModel) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMsg returns the ResultMsg field value if set, zero value otherwise.
func (o *AlipayEbppPdeductPayResponseModel) GetResultMsg() string {
	if o == nil || IsNil(o.ResultMsg) {
		var ret string
		return ret
	}
	return *o.ResultMsg
}

// GetResultMsgOk returns a tuple with the ResultMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductPayResponseModel) GetResultMsgOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMsg) {
		return nil, false
	}
	return o.ResultMsg, true
}

// HasResultMsg returns a boolean if a field has been set.
func (o *AlipayEbppPdeductPayResponseModel) HasResultMsg() bool {
	if o != nil && !IsNil(o.ResultMsg) {
		return true
	}

	return false
}

// SetResultMsg gets a reference to the given string and assigns it to the ResultMsg field.
func (o *AlipayEbppPdeductPayResponseModel) SetResultMsg(v string) {
	o.ResultMsg = &v
}

// GetResultStatus returns the ResultStatus field value if set, zero value otherwise.
func (o *AlipayEbppPdeductPayResponseModel) GetResultStatus() string {
	if o == nil || IsNil(o.ResultStatus) {
		var ret string
		return ret
	}
	return *o.ResultStatus
}

// GetResultStatusOk returns a tuple with the ResultStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductPayResponseModel) GetResultStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ResultStatus) {
		return nil, false
	}
	return o.ResultStatus, true
}

// HasResultStatus returns a boolean if a field has been set.
func (o *AlipayEbppPdeductPayResponseModel) HasResultStatus() bool {
	if o != nil && !IsNil(o.ResultStatus) {
		return true
	}

	return false
}

// SetResultStatus gets a reference to the given string and assigns it to the ResultStatus field.
func (o *AlipayEbppPdeductPayResponseModel) SetResultStatus(v string) {
	o.ResultStatus = &v
}

func (o AlipayEbppPdeductPayResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppPdeductPayResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.BillNo) {
		toSerialize["bill_no"] = o.BillNo
	}
	if !IsNil(o.ExtendField) {
		toSerialize["extend_field"] = o.ExtendField
	}
	if !IsNil(o.OutOrderNo) {
		toSerialize["out_order_no"] = o.OutOrderNo
	}
	if !IsNil(o.ResultCode) {
		toSerialize["result_code"] = o.ResultCode
	}
	if !IsNil(o.ResultMsg) {
		toSerialize["result_msg"] = o.ResultMsg
	}
	if !IsNil(o.ResultStatus) {
		toSerialize["result_status"] = o.ResultStatus
	}
	return toSerialize, nil
}

type NullableAlipayEbppPdeductPayResponseModel struct {
	value *AlipayEbppPdeductPayResponseModel
	isSet bool
}

func (v NullableAlipayEbppPdeductPayResponseModel) Get() *AlipayEbppPdeductPayResponseModel {
	return v.value
}

func (v *NullableAlipayEbppPdeductPayResponseModel) Set(val *AlipayEbppPdeductPayResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppPdeductPayResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppPdeductPayResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppPdeductPayResponseModel(val *AlipayEbppPdeductPayResponseModel) *NullableAlipayEbppPdeductPayResponseModel {
	return &NullableAlipayEbppPdeductPayResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppPdeductPayResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppPdeductPayResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
