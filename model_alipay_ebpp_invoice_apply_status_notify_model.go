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

// checks if the AlipayEbppInvoiceApplyStatusNotifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceApplyStatusNotifyModel{}

// AlipayEbppInvoiceApplyStatusNotifyModel struct for AlipayEbppInvoiceApplyStatusNotifyModel
type AlipayEbppInvoiceApplyStatusNotifyModel struct {
	// 发票申请ID，由阿里发票平台生成。字母或数字组成。 申请单的唯一标识，幂等字段。
	ApplyId *string `json:"apply_id,omitempty"`
	// 申请状态，可选值： applying: 申请中，初始状态； cancelled: 申请已取消、或商户已驳回； creating_inv: 商户已受理/开票中，待发票结果回传； inv_failed: 开票失败； inv_success: 开票成功； inv_part_success: 部分成功（拆单场景下存在。举例：发票申请拆单之后有10张票，其中有1张开票成功时，此时申请状态为inv_part_success，当10张票全部成功申请状态则为inv_success）
	ApplyStatus *string `json:"apply_status,omitempty"`
	// 该申请下所有已开具成功的发票。 状态变更为 apply_status=inv_success 时该字段必传
	InvoiceUks []InvoiceUkDTO `json:"invoice_uks,omitempty"`
	// 说明信息：驳回或失败原因 apply_status=inv_failed 或 apply_status=cancelled 时必传
	Message *string `json:"message,omitempty"`
}

// NewAlipayEbppInvoiceApplyStatusNotifyModel instantiates a new AlipayEbppInvoiceApplyStatusNotifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceApplyStatusNotifyModel() *AlipayEbppInvoiceApplyStatusNotifyModel {
	this := AlipayEbppInvoiceApplyStatusNotifyModel{}
	return &this
}

// NewAlipayEbppInvoiceApplyStatusNotifyModelWithDefaults instantiates a new AlipayEbppInvoiceApplyStatusNotifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceApplyStatusNotifyModelWithDefaults() *AlipayEbppInvoiceApplyStatusNotifyModel {
	this := AlipayEbppInvoiceApplyStatusNotifyModel{}
	return &this
}

// GetApplyId returns the ApplyId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) GetApplyId() string {
	if o == nil || IsNil(o.ApplyId) {
		var ret string
		return ret
	}
	return *o.ApplyId
}

// GetApplyIdOk returns a tuple with the ApplyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) GetApplyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyId) {
		return nil, false
	}
	return o.ApplyId, true
}

// HasApplyId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) HasApplyId() bool {
	if o != nil && !IsNil(o.ApplyId) {
		return true
	}

	return false
}

// SetApplyId gets a reference to the given string and assigns it to the ApplyId field.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) SetApplyId(v string) {
	o.ApplyId = &v
}

// GetApplyStatus returns the ApplyStatus field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) GetApplyStatus() string {
	if o == nil || IsNil(o.ApplyStatus) {
		var ret string
		return ret
	}
	return *o.ApplyStatus
}

// GetApplyStatusOk returns a tuple with the ApplyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) GetApplyStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyStatus) {
		return nil, false
	}
	return o.ApplyStatus, true
}

// HasApplyStatus returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) HasApplyStatus() bool {
	if o != nil && !IsNil(o.ApplyStatus) {
		return true
	}

	return false
}

// SetApplyStatus gets a reference to the given string and assigns it to the ApplyStatus field.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) SetApplyStatus(v string) {
	o.ApplyStatus = &v
}

// GetInvoiceUks returns the InvoiceUks field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) GetInvoiceUks() []InvoiceUkDTO {
	if o == nil || IsNil(o.InvoiceUks) {
		var ret []InvoiceUkDTO
		return ret
	}
	return o.InvoiceUks
}

// GetInvoiceUksOk returns a tuple with the InvoiceUks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) GetInvoiceUksOk() ([]InvoiceUkDTO, bool) {
	if o == nil || IsNil(o.InvoiceUks) {
		return nil, false
	}
	return o.InvoiceUks, true
}

// HasInvoiceUks returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) HasInvoiceUks() bool {
	if o != nil && !IsNil(o.InvoiceUks) {
		return true
	}

	return false
}

// SetInvoiceUks gets a reference to the given []InvoiceUkDTO and assigns it to the InvoiceUks field.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) SetInvoiceUks(v []InvoiceUkDTO) {
	o.InvoiceUks = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AlipayEbppInvoiceApplyStatusNotifyModel) SetMessage(v string) {
	o.Message = &v
}

func (o AlipayEbppInvoiceApplyStatusNotifyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceApplyStatusNotifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyId) {
		toSerialize["apply_id"] = o.ApplyId
	}
	if !IsNil(o.ApplyStatus) {
		toSerialize["apply_status"] = o.ApplyStatus
	}
	if !IsNil(o.InvoiceUks) {
		toSerialize["invoice_uks"] = o.InvoiceUks
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceApplyStatusNotifyModel struct {
	value *AlipayEbppInvoiceApplyStatusNotifyModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceApplyStatusNotifyModel) Get() *AlipayEbppInvoiceApplyStatusNotifyModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceApplyStatusNotifyModel) Set(val *AlipayEbppInvoiceApplyStatusNotifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceApplyStatusNotifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceApplyStatusNotifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceApplyStatusNotifyModel(val *AlipayEbppInvoiceApplyStatusNotifyModel) *NullableAlipayEbppInvoiceApplyStatusNotifyModel {
	return &NullableAlipayEbppInvoiceApplyStatusNotifyModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceApplyStatusNotifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceApplyStatusNotifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
