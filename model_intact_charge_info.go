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

// checks if the IntactChargeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntactChargeInfo{}

// IntactChargeInfo struct for IntactChargeInfo
type IntactChargeInfo struct {
	// 实际收费金额，单位元
	ActualAmount *string `json:"actual_amount,omitempty"`
	// 收费类型
	BillType *string `json:"bill_type,omitempty"`
	// 收费时间,时间精确到秒
	GmtPay *string `json:"gmt_pay,omitempty"`
	// 是否退费
	IsRefund *bool `json:"is_refund,omitempty"`
	// 外部请求号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 应收费金额，单位元
	PlanAmount *string `json:"plan_amount,omitempty"`
	// 收费产品
	ProductName *string `json:"product_name,omitempty"`
	// 收费唯一id
	ServiceTarget *string `json:"service_target,omitempty"`
	// 收费类型
	ServiceType *string `json:"service_type,omitempty"`
	// 状态
	Status *string `json:"status,omitempty"`
	// 收费目标账号
	TargetAccountNo *string `json:"target_account_no,omitempty"`
	// 收费目标uid
	TargetUserId *string `json:"target_user_id,omitempty"`
}

// NewIntactChargeInfo instantiates a new IntactChargeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntactChargeInfo() *IntactChargeInfo {
	this := IntactChargeInfo{}
	return &this
}

// NewIntactChargeInfoWithDefaults instantiates a new IntactChargeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntactChargeInfoWithDefaults() *IntactChargeInfo {
	this := IntactChargeInfo{}
	return &this
}

// GetActualAmount returns the ActualAmount field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetActualAmount() string {
	if o == nil || IsNil(o.ActualAmount) {
		var ret string
		return ret
	}
	return *o.ActualAmount
}

// GetActualAmountOk returns a tuple with the ActualAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetActualAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ActualAmount) {
		return nil, false
	}
	return o.ActualAmount, true
}

// HasActualAmount returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasActualAmount() bool {
	if o != nil && !IsNil(o.ActualAmount) {
		return true
	}

	return false
}

// SetActualAmount gets a reference to the given string and assigns it to the ActualAmount field.
func (o *IntactChargeInfo) SetActualAmount(v string) {
	o.ActualAmount = &v
}

// GetBillType returns the BillType field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetBillType() string {
	if o == nil || IsNil(o.BillType) {
		var ret string
		return ret
	}
	return *o.BillType
}

// GetBillTypeOk returns a tuple with the BillType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetBillTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BillType) {
		return nil, false
	}
	return o.BillType, true
}

// HasBillType returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasBillType() bool {
	if o != nil && !IsNil(o.BillType) {
		return true
	}

	return false
}

// SetBillType gets a reference to the given string and assigns it to the BillType field.
func (o *IntactChargeInfo) SetBillType(v string) {
	o.BillType = &v
}

// GetGmtPay returns the GmtPay field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetGmtPay() string {
	if o == nil || IsNil(o.GmtPay) {
		var ret string
		return ret
	}
	return *o.GmtPay
}

// GetGmtPayOk returns a tuple with the GmtPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetGmtPayOk() (*string, bool) {
	if o == nil || IsNil(o.GmtPay) {
		return nil, false
	}
	return o.GmtPay, true
}

// HasGmtPay returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasGmtPay() bool {
	if o != nil && !IsNil(o.GmtPay) {
		return true
	}

	return false
}

// SetGmtPay gets a reference to the given string and assigns it to the GmtPay field.
func (o *IntactChargeInfo) SetGmtPay(v string) {
	o.GmtPay = &v
}

// GetIsRefund returns the IsRefund field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetIsRefund() bool {
	if o == nil || IsNil(o.IsRefund) {
		var ret bool
		return ret
	}
	return *o.IsRefund
}

// GetIsRefundOk returns a tuple with the IsRefund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetIsRefundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRefund) {
		return nil, false
	}
	return o.IsRefund, true
}

// HasIsRefund returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasIsRefund() bool {
	if o != nil && !IsNil(o.IsRefund) {
		return true
	}

	return false
}

// SetIsRefund gets a reference to the given bool and assigns it to the IsRefund field.
func (o *IntactChargeInfo) SetIsRefund(v bool) {
	o.IsRefund = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *IntactChargeInfo) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPlanAmount returns the PlanAmount field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetPlanAmount() string {
	if o == nil || IsNil(o.PlanAmount) {
		var ret string
		return ret
	}
	return *o.PlanAmount
}

// GetPlanAmountOk returns a tuple with the PlanAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetPlanAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PlanAmount) {
		return nil, false
	}
	return o.PlanAmount, true
}

// HasPlanAmount returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasPlanAmount() bool {
	if o != nil && !IsNil(o.PlanAmount) {
		return true
	}

	return false
}

// SetPlanAmount gets a reference to the given string and assigns it to the PlanAmount field.
func (o *IntactChargeInfo) SetPlanAmount(v string) {
	o.PlanAmount = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *IntactChargeInfo) SetProductName(v string) {
	o.ProductName = &v
}

// GetServiceTarget returns the ServiceTarget field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetServiceTarget() string {
	if o == nil || IsNil(o.ServiceTarget) {
		var ret string
		return ret
	}
	return *o.ServiceTarget
}

// GetServiceTargetOk returns a tuple with the ServiceTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetServiceTargetOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceTarget) {
		return nil, false
	}
	return o.ServiceTarget, true
}

// HasServiceTarget returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasServiceTarget() bool {
	if o != nil && !IsNil(o.ServiceTarget) {
		return true
	}

	return false
}

// SetServiceTarget gets a reference to the given string and assigns it to the ServiceTarget field.
func (o *IntactChargeInfo) SetServiceTarget(v string) {
	o.ServiceTarget = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetServiceType() string {
	if o == nil || IsNil(o.ServiceType) {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *IntactChargeInfo) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IntactChargeInfo) SetStatus(v string) {
	o.Status = &v
}

// GetTargetAccountNo returns the TargetAccountNo field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetTargetAccountNo() string {
	if o == nil || IsNil(o.TargetAccountNo) {
		var ret string
		return ret
	}
	return *o.TargetAccountNo
}

// GetTargetAccountNoOk returns a tuple with the TargetAccountNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetTargetAccountNoOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAccountNo) {
		return nil, false
	}
	return o.TargetAccountNo, true
}

// HasTargetAccountNo returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasTargetAccountNo() bool {
	if o != nil && !IsNil(o.TargetAccountNo) {
		return true
	}

	return false
}

// SetTargetAccountNo gets a reference to the given string and assigns it to the TargetAccountNo field.
func (o *IntactChargeInfo) SetTargetAccountNo(v string) {
	o.TargetAccountNo = &v
}

// GetTargetUserId returns the TargetUserId field value if set, zero value otherwise.
func (o *IntactChargeInfo) GetTargetUserId() string {
	if o == nil || IsNil(o.TargetUserId) {
		var ret string
		return ret
	}
	return *o.TargetUserId
}

// GetTargetUserIdOk returns a tuple with the TargetUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntactChargeInfo) GetTargetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUserId) {
		return nil, false
	}
	return o.TargetUserId, true
}

// HasTargetUserId returns a boolean if a field has been set.
func (o *IntactChargeInfo) HasTargetUserId() bool {
	if o != nil && !IsNil(o.TargetUserId) {
		return true
	}

	return false
}

// SetTargetUserId gets a reference to the given string and assigns it to the TargetUserId field.
func (o *IntactChargeInfo) SetTargetUserId(v string) {
	o.TargetUserId = &v
}

func (o IntactChargeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntactChargeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActualAmount) {
		toSerialize["actual_amount"] = o.ActualAmount
	}
	if !IsNil(o.BillType) {
		toSerialize["bill_type"] = o.BillType
	}
	if !IsNil(o.GmtPay) {
		toSerialize["gmt_pay"] = o.GmtPay
	}
	if !IsNil(o.IsRefund) {
		toSerialize["is_refund"] = o.IsRefund
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PlanAmount) {
		toSerialize["plan_amount"] = o.PlanAmount
	}
	if !IsNil(o.ProductName) {
		toSerialize["product_name"] = o.ProductName
	}
	if !IsNil(o.ServiceTarget) {
		toSerialize["service_target"] = o.ServiceTarget
	}
	if !IsNil(o.ServiceType) {
		toSerialize["service_type"] = o.ServiceType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TargetAccountNo) {
		toSerialize["target_account_no"] = o.TargetAccountNo
	}
	if !IsNil(o.TargetUserId) {
		toSerialize["target_user_id"] = o.TargetUserId
	}
	return toSerialize, nil
}

type NullableIntactChargeInfo struct {
	value *IntactChargeInfo
	isSet bool
}

func (v NullableIntactChargeInfo) Get() *IntactChargeInfo {
	return v.value
}

func (v *NullableIntactChargeInfo) Set(val *IntactChargeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIntactChargeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIntactChargeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntactChargeInfo(val *IntactChargeInfo) *NullableIntactChargeInfo {
	return &NullableIntactChargeInfo{value: val, isSet: true}
}

func (v NullableIntactChargeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntactChargeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

