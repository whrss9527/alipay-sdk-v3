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

// checks if the ZhimaMerchantSubsidiariesCloseResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaMerchantSubsidiariesCloseResponseModel{}

// ZhimaMerchantSubsidiariesCloseResponseModel struct for ZhimaMerchantSubsidiariesCloseResponseModel
type ZhimaMerchantSubsidiariesCloseResponseModel struct {
	// 业务错误码
	BizErrorCode *string `json:"biz_error_code,omitempty"`
	// 业务错误信息
	BizErrorMessage *string `json:"biz_error_message,omitempty"`
	// 工单审核备注
	OrderMemo *string `json:"order_memo,omitempty"`
	// 工单标识，业务成功时返回工单号
	OrderNo *string `json:"order_no,omitempty"`
	// 工单状态
	OrderStatus *string `json:"order_status,omitempty"`
}

// NewZhimaMerchantSubsidiariesCloseResponseModel instantiates a new ZhimaMerchantSubsidiariesCloseResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaMerchantSubsidiariesCloseResponseModel() *ZhimaMerchantSubsidiariesCloseResponseModel {
	this := ZhimaMerchantSubsidiariesCloseResponseModel{}
	return &this
}

// NewZhimaMerchantSubsidiariesCloseResponseModelWithDefaults instantiates a new ZhimaMerchantSubsidiariesCloseResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaMerchantSubsidiariesCloseResponseModelWithDefaults() *ZhimaMerchantSubsidiariesCloseResponseModel {
	this := ZhimaMerchantSubsidiariesCloseResponseModel{}
	return &this
}

// GetBizErrorCode returns the BizErrorCode field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) GetBizErrorCode() string {
	if o == nil || IsNil(o.BizErrorCode) {
		var ret string
		return ret
	}
	return *o.BizErrorCode
}

// GetBizErrorCodeOk returns a tuple with the BizErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) GetBizErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BizErrorCode) {
		return nil, false
	}
	return o.BizErrorCode, true
}

// HasBizErrorCode returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) HasBizErrorCode() bool {
	if o != nil && !IsNil(o.BizErrorCode) {
		return true
	}

	return false
}

// SetBizErrorCode gets a reference to the given string and assigns it to the BizErrorCode field.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) SetBizErrorCode(v string) {
	o.BizErrorCode = &v
}

// GetBizErrorMessage returns the BizErrorMessage field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) GetBizErrorMessage() string {
	if o == nil || IsNil(o.BizErrorMessage) {
		var ret string
		return ret
	}
	return *o.BizErrorMessage
}

// GetBizErrorMessageOk returns a tuple with the BizErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) GetBizErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.BizErrorMessage) {
		return nil, false
	}
	return o.BizErrorMessage, true
}

// HasBizErrorMessage returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) HasBizErrorMessage() bool {
	if o != nil && !IsNil(o.BizErrorMessage) {
		return true
	}

	return false
}

// SetBizErrorMessage gets a reference to the given string and assigns it to the BizErrorMessage field.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) SetBizErrorMessage(v string) {
	o.BizErrorMessage = &v
}

// GetOrderMemo returns the OrderMemo field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) GetOrderMemo() string {
	if o == nil || IsNil(o.OrderMemo) {
		var ret string
		return ret
	}
	return *o.OrderMemo
}

// GetOrderMemoOk returns a tuple with the OrderMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) GetOrderMemoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderMemo) {
		return nil, false
	}
	return o.OrderMemo, true
}

// HasOrderMemo returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) HasOrderMemo() bool {
	if o != nil && !IsNil(o.OrderMemo) {
		return true
	}

	return false
}

// SetOrderMemo gets a reference to the given string and assigns it to the OrderMemo field.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) SetOrderMemo(v string) {
	o.OrderMemo = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *ZhimaMerchantSubsidiariesCloseResponseModel) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

func (o ZhimaMerchantSubsidiariesCloseResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaMerchantSubsidiariesCloseResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizErrorCode) {
		toSerialize["biz_error_code"] = o.BizErrorCode
	}
	if !IsNil(o.BizErrorMessage) {
		toSerialize["biz_error_message"] = o.BizErrorMessage
	}
	if !IsNil(o.OrderMemo) {
		toSerialize["order_memo"] = o.OrderMemo
	}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["order_status"] = o.OrderStatus
	}
	return toSerialize, nil
}

type NullableZhimaMerchantSubsidiariesCloseResponseModel struct {
	value *ZhimaMerchantSubsidiariesCloseResponseModel
	isSet bool
}

func (v NullableZhimaMerchantSubsidiariesCloseResponseModel) Get() *ZhimaMerchantSubsidiariesCloseResponseModel {
	return v.value
}

func (v *NullableZhimaMerchantSubsidiariesCloseResponseModel) Set(val *ZhimaMerchantSubsidiariesCloseResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaMerchantSubsidiariesCloseResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaMerchantSubsidiariesCloseResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaMerchantSubsidiariesCloseResponseModel(val *ZhimaMerchantSubsidiariesCloseResponseModel) *NullableZhimaMerchantSubsidiariesCloseResponseModel {
	return &NullableZhimaMerchantSubsidiariesCloseResponseModel{value: val, isSet: true}
}

func (v NullableZhimaMerchantSubsidiariesCloseResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaMerchantSubsidiariesCloseResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
