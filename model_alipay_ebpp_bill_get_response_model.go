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

// checks if the AlipayEbppBillGetResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppBillGetResponseModel{}

// AlipayEbppBillGetResponseModel struct for AlipayEbppBillGetResponseModel
type AlipayEbppBillGetResponseModel struct {
	// 支付宝的业务订单号，具有唯一性。
	AlipayOrderNo *string `json:"alipay_order_no,omitempty"`
	// 账单的账期，例如201203表示2012年3月的账单。
	BillDate *string `json:"bill_date,omitempty"`
	// 账单单据号，例如水费单号，手机号，电费号，信用卡卡号。没有唯一性要求。
	BillKey *string `json:"bill_key,omitempty"`
	// 支付宝给每个出账机构指定了一个对应的英文短名称来唯一表示该收费单位。
	ChargeInst *string `json:"charge_inst,omitempty"`
	// 出账机构中文名称。
	ChargeInstName *string `json:"charge_inst_name,omitempty"`
	// 输出机构的业务流水号，需要保证唯一性。
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// 账单的状态。 INIT:等待付款，SUCCESS:成功,FAILED:失败。
	OrderStatus *string `json:"order_status,omitempty"`
	// 支付宝订单类型。公共事业缴纳JF,信用卡还款HK
	OrderType *string `json:"order_type,omitempty"`
	// 拥有该账单的用户姓名
	OwnerName *string `json:"owner_name,omitempty"`
	// 缴费金额。用户支付的总金额。单位为：RMB Yuan。取值范围为[0.01，100000000.00]，精确到小数点后两位。
	PayAmount *string `json:"pay_amount,omitempty"`
	// 付款时间
	PayTime *string `json:"pay_time,omitempty"`
	// 账单的服务费
	ServiceAmount *string `json:"service_amount,omitempty"`
	// 子业务类型是业务类型的下一级概念，例如：WATER表示JF下面的水费，ELECTRIC表示JF下面的电费，GAS表示JF下面的燃气费。
	SubOrderType *string `json:"sub_order_type,omitempty"`
	// 交通违章地点，sub_order_type=TRAFFIC时有值
	TrafficLocation *string `json:"traffic_location,omitempty"`
	// 违章行为，sub_order_type=TRAFFIC时有值。
	TrafficRegulations *string `json:"traffic_regulations,omitempty"`
}

// NewAlipayEbppBillGetResponseModel instantiates a new AlipayEbppBillGetResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppBillGetResponseModel() *AlipayEbppBillGetResponseModel {
	this := AlipayEbppBillGetResponseModel{}
	return &this
}

// NewAlipayEbppBillGetResponseModelWithDefaults instantiates a new AlipayEbppBillGetResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppBillGetResponseModelWithDefaults() *AlipayEbppBillGetResponseModel {
	this := AlipayEbppBillGetResponseModel{}
	return &this
}

// GetAlipayOrderNo returns the AlipayOrderNo field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetAlipayOrderNo() string {
	if o == nil || IsNil(o.AlipayOrderNo) {
		var ret string
		return ret
	}
	return *o.AlipayOrderNo
}

// GetAlipayOrderNoOk returns a tuple with the AlipayOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetAlipayOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayOrderNo) {
		return nil, false
	}
	return o.AlipayOrderNo, true
}

// HasAlipayOrderNo returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasAlipayOrderNo() bool {
	if o != nil && !IsNil(o.AlipayOrderNo) {
		return true
	}

	return false
}

// SetAlipayOrderNo gets a reference to the given string and assigns it to the AlipayOrderNo field.
func (o *AlipayEbppBillGetResponseModel) SetAlipayOrderNo(v string) {
	o.AlipayOrderNo = &v
}

// GetBillDate returns the BillDate field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetBillDate() string {
	if o == nil || IsNil(o.BillDate) {
		var ret string
		return ret
	}
	return *o.BillDate
}

// GetBillDateOk returns a tuple with the BillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetBillDateOk() (*string, bool) {
	if o == nil || IsNil(o.BillDate) {
		return nil, false
	}
	return o.BillDate, true
}

// HasBillDate returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasBillDate() bool {
	if o != nil && !IsNil(o.BillDate) {
		return true
	}

	return false
}

// SetBillDate gets a reference to the given string and assigns it to the BillDate field.
func (o *AlipayEbppBillGetResponseModel) SetBillDate(v string) {
	o.BillDate = &v
}

// GetBillKey returns the BillKey field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetBillKey() string {
	if o == nil || IsNil(o.BillKey) {
		var ret string
		return ret
	}
	return *o.BillKey
}

// GetBillKeyOk returns a tuple with the BillKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetBillKeyOk() (*string, bool) {
	if o == nil || IsNil(o.BillKey) {
		return nil, false
	}
	return o.BillKey, true
}

// HasBillKey returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasBillKey() bool {
	if o != nil && !IsNil(o.BillKey) {
		return true
	}

	return false
}

// SetBillKey gets a reference to the given string and assigns it to the BillKey field.
func (o *AlipayEbppBillGetResponseModel) SetBillKey(v string) {
	o.BillKey = &v
}

// GetChargeInst returns the ChargeInst field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetChargeInst() string {
	if o == nil || IsNil(o.ChargeInst) {
		var ret string
		return ret
	}
	return *o.ChargeInst
}

// GetChargeInstOk returns a tuple with the ChargeInst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetChargeInstOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeInst) {
		return nil, false
	}
	return o.ChargeInst, true
}

// HasChargeInst returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasChargeInst() bool {
	if o != nil && !IsNil(o.ChargeInst) {
		return true
	}

	return false
}

// SetChargeInst gets a reference to the given string and assigns it to the ChargeInst field.
func (o *AlipayEbppBillGetResponseModel) SetChargeInst(v string) {
	o.ChargeInst = &v
}

// GetChargeInstName returns the ChargeInstName field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetChargeInstName() string {
	if o == nil || IsNil(o.ChargeInstName) {
		var ret string
		return ret
	}
	return *o.ChargeInstName
}

// GetChargeInstNameOk returns a tuple with the ChargeInstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetChargeInstNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeInstName) {
		return nil, false
	}
	return o.ChargeInstName, true
}

// HasChargeInstName returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasChargeInstName() bool {
	if o != nil && !IsNil(o.ChargeInstName) {
		return true
	}

	return false
}

// SetChargeInstName gets a reference to the given string and assigns it to the ChargeInstName field.
func (o *AlipayEbppBillGetResponseModel) SetChargeInstName(v string) {
	o.ChargeInstName = &v
}

// GetMerchantOrderNo returns the MerchantOrderNo field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetMerchantOrderNo() string {
	if o == nil || IsNil(o.MerchantOrderNo) {
		var ret string
		return ret
	}
	return *o.MerchantOrderNo
}

// GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetMerchantOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantOrderNo) {
		return nil, false
	}
	return o.MerchantOrderNo, true
}

// HasMerchantOrderNo returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasMerchantOrderNo() bool {
	if o != nil && !IsNil(o.MerchantOrderNo) {
		return true
	}

	return false
}

// SetMerchantOrderNo gets a reference to the given string and assigns it to the MerchantOrderNo field.
func (o *AlipayEbppBillGetResponseModel) SetMerchantOrderNo(v string) {
	o.MerchantOrderNo = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *AlipayEbppBillGetResponseModel) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetOrderType() string {
	if o == nil || IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *AlipayEbppBillGetResponseModel) SetOrderType(v string) {
	o.OrderType = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *AlipayEbppBillGetResponseModel) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetPayAmount returns the PayAmount field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetPayAmount() string {
	if o == nil || IsNil(o.PayAmount) {
		var ret string
		return ret
	}
	return *o.PayAmount
}

// GetPayAmountOk returns a tuple with the PayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PayAmount) {
		return nil, false
	}
	return o.PayAmount, true
}

// HasPayAmount returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasPayAmount() bool {
	if o != nil && !IsNil(o.PayAmount) {
		return true
	}

	return false
}

// SetPayAmount gets a reference to the given string and assigns it to the PayAmount field.
func (o *AlipayEbppBillGetResponseModel) SetPayAmount(v string) {
	o.PayAmount = &v
}

// GetPayTime returns the PayTime field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetPayTime() string {
	if o == nil || IsNil(o.PayTime) {
		var ret string
		return ret
	}
	return *o.PayTime
}

// GetPayTimeOk returns a tuple with the PayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetPayTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PayTime) {
		return nil, false
	}
	return o.PayTime, true
}

// HasPayTime returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasPayTime() bool {
	if o != nil && !IsNil(o.PayTime) {
		return true
	}

	return false
}

// SetPayTime gets a reference to the given string and assigns it to the PayTime field.
func (o *AlipayEbppBillGetResponseModel) SetPayTime(v string) {
	o.PayTime = &v
}

// GetServiceAmount returns the ServiceAmount field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetServiceAmount() string {
	if o == nil || IsNil(o.ServiceAmount) {
		var ret string
		return ret
	}
	return *o.ServiceAmount
}

// GetServiceAmountOk returns a tuple with the ServiceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetServiceAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAmount) {
		return nil, false
	}
	return o.ServiceAmount, true
}

// HasServiceAmount returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasServiceAmount() bool {
	if o != nil && !IsNil(o.ServiceAmount) {
		return true
	}

	return false
}

// SetServiceAmount gets a reference to the given string and assigns it to the ServiceAmount field.
func (o *AlipayEbppBillGetResponseModel) SetServiceAmount(v string) {
	o.ServiceAmount = &v
}

// GetSubOrderType returns the SubOrderType field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetSubOrderType() string {
	if o == nil || IsNil(o.SubOrderType) {
		var ret string
		return ret
	}
	return *o.SubOrderType
}

// GetSubOrderTypeOk returns a tuple with the SubOrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetSubOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubOrderType) {
		return nil, false
	}
	return o.SubOrderType, true
}

// HasSubOrderType returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasSubOrderType() bool {
	if o != nil && !IsNil(o.SubOrderType) {
		return true
	}

	return false
}

// SetSubOrderType gets a reference to the given string and assigns it to the SubOrderType field.
func (o *AlipayEbppBillGetResponseModel) SetSubOrderType(v string) {
	o.SubOrderType = &v
}

// GetTrafficLocation returns the TrafficLocation field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetTrafficLocation() string {
	if o == nil || IsNil(o.TrafficLocation) {
		var ret string
		return ret
	}
	return *o.TrafficLocation
}

// GetTrafficLocationOk returns a tuple with the TrafficLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetTrafficLocationOk() (*string, bool) {
	if o == nil || IsNil(o.TrafficLocation) {
		return nil, false
	}
	return o.TrafficLocation, true
}

// HasTrafficLocation returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasTrafficLocation() bool {
	if o != nil && !IsNil(o.TrafficLocation) {
		return true
	}

	return false
}

// SetTrafficLocation gets a reference to the given string and assigns it to the TrafficLocation field.
func (o *AlipayEbppBillGetResponseModel) SetTrafficLocation(v string) {
	o.TrafficLocation = &v
}

// GetTrafficRegulations returns the TrafficRegulations field value if set, zero value otherwise.
func (o *AlipayEbppBillGetResponseModel) GetTrafficRegulations() string {
	if o == nil || IsNil(o.TrafficRegulations) {
		var ret string
		return ret
	}
	return *o.TrafficRegulations
}

// GetTrafficRegulationsOk returns a tuple with the TrafficRegulations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppBillGetResponseModel) GetTrafficRegulationsOk() (*string, bool) {
	if o == nil || IsNil(o.TrafficRegulations) {
		return nil, false
	}
	return o.TrafficRegulations, true
}

// HasTrafficRegulations returns a boolean if a field has been set.
func (o *AlipayEbppBillGetResponseModel) HasTrafficRegulations() bool {
	if o != nil && !IsNil(o.TrafficRegulations) {
		return true
	}

	return false
}

// SetTrafficRegulations gets a reference to the given string and assigns it to the TrafficRegulations field.
func (o *AlipayEbppBillGetResponseModel) SetTrafficRegulations(v string) {
	o.TrafficRegulations = &v
}

func (o AlipayEbppBillGetResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppBillGetResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlipayOrderNo) {
		toSerialize["alipay_order_no"] = o.AlipayOrderNo
	}
	if !IsNil(o.BillDate) {
		toSerialize["bill_date"] = o.BillDate
	}
	if !IsNil(o.BillKey) {
		toSerialize["bill_key"] = o.BillKey
	}
	if !IsNil(o.ChargeInst) {
		toSerialize["charge_inst"] = o.ChargeInst
	}
	if !IsNil(o.ChargeInstName) {
		toSerialize["charge_inst_name"] = o.ChargeInstName
	}
	if !IsNil(o.MerchantOrderNo) {
		toSerialize["merchant_order_no"] = o.MerchantOrderNo
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["order_status"] = o.OrderStatus
	}
	if !IsNil(o.OrderType) {
		toSerialize["order_type"] = o.OrderType
	}
	if !IsNil(o.OwnerName) {
		toSerialize["owner_name"] = o.OwnerName
	}
	if !IsNil(o.PayAmount) {
		toSerialize["pay_amount"] = o.PayAmount
	}
	if !IsNil(o.PayTime) {
		toSerialize["pay_time"] = o.PayTime
	}
	if !IsNil(o.ServiceAmount) {
		toSerialize["service_amount"] = o.ServiceAmount
	}
	if !IsNil(o.SubOrderType) {
		toSerialize["sub_order_type"] = o.SubOrderType
	}
	if !IsNil(o.TrafficLocation) {
		toSerialize["traffic_location"] = o.TrafficLocation
	}
	if !IsNil(o.TrafficRegulations) {
		toSerialize["traffic_regulations"] = o.TrafficRegulations
	}
	return toSerialize, nil
}

type NullableAlipayEbppBillGetResponseModel struct {
	value *AlipayEbppBillGetResponseModel
	isSet bool
}

func (v NullableAlipayEbppBillGetResponseModel) Get() *AlipayEbppBillGetResponseModel {
	return v.value
}

func (v *NullableAlipayEbppBillGetResponseModel) Set(val *AlipayEbppBillGetResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppBillGetResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppBillGetResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppBillGetResponseModel(val *AlipayEbppBillGetResponseModel) *NullableAlipayEbppBillGetResponseModel {
	return &NullableAlipayEbppBillGetResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppBillGetResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppBillGetResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
