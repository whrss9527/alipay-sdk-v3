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

// checks if the AlipayFundTransUniTransferModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundTransUniTransferModel{}

// AlipayFundTransUniTransferModel struct for AlipayFundTransUniTransferModel
type AlipayFundTransUniTransferModel struct {
	// 描述特定的业务场景，可传的参数如下： DIRECT_TRANSFER：单笔无密转账到支付宝，B2C现金红包; PERSONAL_COLLECTION：C2C现金红包-领红包
	BizScene *string `json:"biz_scene,omitempty"`
	// 转账业务请求的扩展参数，支持传入的扩展参数如下： sub_biz_scene 子业务场景，红包业务必传，取值REDPACKET，C2C现金红包、B2C现金红包均需传入
	BusinessParams        *string                `json:"business_params,omitempty"`
	MutipleCurrencyDetail *MutipleCurrencyDetail `json:"mutiple_currency_detail,omitempty"`
	// 转账业务的标题，用于在支付宝用户的账单里显示
	OrderTitle *string `json:"order_title,omitempty"`
	// 原支付宝业务单号。
	OriginalOrderId *string `json:"original_order_id,omitempty"`
	// 商家侧唯一订单号，由商家自定义。对于不同转账请求，商家需保证该订单号在自身系统唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 公用回传参数，如果请求时传递了该参数，则异步通知商户时会回传该参数。
	PassbackParams *string      `json:"passback_params,omitempty"`
	PayeeInfo      *Participant `json:"payee_info,omitempty"`
	PayerInfo      *Participant `json:"payer_info,omitempty"`
	// 业务产品码， 单笔无密转账到支付宝账户固定为: TRANS_ACCOUNT_NO_PWD； 收发现金红包固定为: STD_RED_PACKET；
	ProductCode *string `json:"product_code,omitempty"`
	// 业务备注
	Remark   *string   `json:"remark,omitempty"`
	SignData *SignData `json:"sign_data,omitempty"`
	// 订单总金额，单位为元，精确到小数点后两位，STD_RED_PACKET产品取值范围[0.01,100000000]； TRANS_ACCOUNT_NO_PWD产品取值范围[0.1,100000000]
	TransAmount *string `json:"trans_amount,omitempty"`
}

// NewAlipayFundTransUniTransferModel instantiates a new AlipayFundTransUniTransferModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundTransUniTransferModel() *AlipayFundTransUniTransferModel {
	this := AlipayFundTransUniTransferModel{}
	return &this
}

// NewAlipayFundTransUniTransferModelWithDefaults instantiates a new AlipayFundTransUniTransferModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundTransUniTransferModelWithDefaults() *AlipayFundTransUniTransferModel {
	this := AlipayFundTransUniTransferModel{}
	return &this
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundTransUniTransferModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetBusinessParams returns the BusinessParams field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetBusinessParams() string {
	if o == nil || IsNil(o.BusinessParams) {
		var ret string
		return ret
	}
	return *o.BusinessParams
}

// GetBusinessParamsOk returns a tuple with the BusinessParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetBusinessParamsOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessParams) {
		return nil, false
	}
	return o.BusinessParams, true
}

// HasBusinessParams returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasBusinessParams() bool {
	if o != nil && !IsNil(o.BusinessParams) {
		return true
	}

	return false
}

// SetBusinessParams gets a reference to the given string and assigns it to the BusinessParams field.
func (o *AlipayFundTransUniTransferModel) SetBusinessParams(v string) {
	o.BusinessParams = &v
}

// GetMutipleCurrencyDetail returns the MutipleCurrencyDetail field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetMutipleCurrencyDetail() MutipleCurrencyDetail {
	if o == nil || IsNil(o.MutipleCurrencyDetail) {
		var ret MutipleCurrencyDetail
		return ret
	}
	return *o.MutipleCurrencyDetail
}

// GetMutipleCurrencyDetailOk returns a tuple with the MutipleCurrencyDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetMutipleCurrencyDetailOk() (*MutipleCurrencyDetail, bool) {
	if o == nil || IsNil(o.MutipleCurrencyDetail) {
		return nil, false
	}
	return o.MutipleCurrencyDetail, true
}

// HasMutipleCurrencyDetail returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasMutipleCurrencyDetail() bool {
	if o != nil && !IsNil(o.MutipleCurrencyDetail) {
		return true
	}

	return false
}

// SetMutipleCurrencyDetail gets a reference to the given MutipleCurrencyDetail and assigns it to the MutipleCurrencyDetail field.
func (o *AlipayFundTransUniTransferModel) SetMutipleCurrencyDetail(v MutipleCurrencyDetail) {
	o.MutipleCurrencyDetail = &v
}

// GetOrderTitle returns the OrderTitle field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetOrderTitle() string {
	if o == nil || IsNil(o.OrderTitle) {
		var ret string
		return ret
	}
	return *o.OrderTitle
}

// GetOrderTitleOk returns a tuple with the OrderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetOrderTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OrderTitle) {
		return nil, false
	}
	return o.OrderTitle, true
}

// HasOrderTitle returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasOrderTitle() bool {
	if o != nil && !IsNil(o.OrderTitle) {
		return true
	}

	return false
}

// SetOrderTitle gets a reference to the given string and assigns it to the OrderTitle field.
func (o *AlipayFundTransUniTransferModel) SetOrderTitle(v string) {
	o.OrderTitle = &v
}

// GetOriginalOrderId returns the OriginalOrderId field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetOriginalOrderId() string {
	if o == nil || IsNil(o.OriginalOrderId) {
		var ret string
		return ret
	}
	return *o.OriginalOrderId
}

// GetOriginalOrderIdOk returns a tuple with the OriginalOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetOriginalOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalOrderId) {
		return nil, false
	}
	return o.OriginalOrderId, true
}

// HasOriginalOrderId returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasOriginalOrderId() bool {
	if o != nil && !IsNil(o.OriginalOrderId) {
		return true
	}

	return false
}

// SetOriginalOrderId gets a reference to the given string and assigns it to the OriginalOrderId field.
func (o *AlipayFundTransUniTransferModel) SetOriginalOrderId(v string) {
	o.OriginalOrderId = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayFundTransUniTransferModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPassbackParams returns the PassbackParams field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetPassbackParams() string {
	if o == nil || IsNil(o.PassbackParams) {
		var ret string
		return ret
	}
	return *o.PassbackParams
}

// GetPassbackParamsOk returns a tuple with the PassbackParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetPassbackParamsOk() (*string, bool) {
	if o == nil || IsNil(o.PassbackParams) {
		return nil, false
	}
	return o.PassbackParams, true
}

// HasPassbackParams returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasPassbackParams() bool {
	if o != nil && !IsNil(o.PassbackParams) {
		return true
	}

	return false
}

// SetPassbackParams gets a reference to the given string and assigns it to the PassbackParams field.
func (o *AlipayFundTransUniTransferModel) SetPassbackParams(v string) {
	o.PassbackParams = &v
}

// GetPayeeInfo returns the PayeeInfo field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetPayeeInfo() Participant {
	if o == nil || IsNil(o.PayeeInfo) {
		var ret Participant
		return ret
	}
	return *o.PayeeInfo
}

// GetPayeeInfoOk returns a tuple with the PayeeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetPayeeInfoOk() (*Participant, bool) {
	if o == nil || IsNil(o.PayeeInfo) {
		return nil, false
	}
	return o.PayeeInfo, true
}

// HasPayeeInfo returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasPayeeInfo() bool {
	if o != nil && !IsNil(o.PayeeInfo) {
		return true
	}

	return false
}

// SetPayeeInfo gets a reference to the given Participant and assigns it to the PayeeInfo field.
func (o *AlipayFundTransUniTransferModel) SetPayeeInfo(v Participant) {
	o.PayeeInfo = &v
}

// GetPayerInfo returns the PayerInfo field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetPayerInfo() Participant {
	if o == nil || IsNil(o.PayerInfo) {
		var ret Participant
		return ret
	}
	return *o.PayerInfo
}

// GetPayerInfoOk returns a tuple with the PayerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetPayerInfoOk() (*Participant, bool) {
	if o == nil || IsNil(o.PayerInfo) {
		return nil, false
	}
	return o.PayerInfo, true
}

// HasPayerInfo returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasPayerInfo() bool {
	if o != nil && !IsNil(o.PayerInfo) {
		return true
	}

	return false
}

// SetPayerInfo gets a reference to the given Participant and assigns it to the PayerInfo field.
func (o *AlipayFundTransUniTransferModel) SetPayerInfo(v Participant) {
	o.PayerInfo = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundTransUniTransferModel) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *AlipayFundTransUniTransferModel) SetRemark(v string) {
	o.Remark = &v
}

// GetSignData returns the SignData field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetSignData() SignData {
	if o == nil || IsNil(o.SignData) {
		var ret SignData
		return ret
	}
	return *o.SignData
}

// GetSignDataOk returns a tuple with the SignData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetSignDataOk() (*SignData, bool) {
	if o == nil || IsNil(o.SignData) {
		return nil, false
	}
	return o.SignData, true
}

// HasSignData returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasSignData() bool {
	if o != nil && !IsNil(o.SignData) {
		return true
	}

	return false
}

// SetSignData gets a reference to the given SignData and assigns it to the SignData field.
func (o *AlipayFundTransUniTransferModel) SetSignData(v SignData) {
	o.SignData = &v
}

// GetTransAmount returns the TransAmount field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferModel) GetTransAmount() string {
	if o == nil || IsNil(o.TransAmount) {
		var ret string
		return ret
	}
	return *o.TransAmount
}

// GetTransAmountOk returns a tuple with the TransAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferModel) GetTransAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TransAmount) {
		return nil, false
	}
	return o.TransAmount, true
}

// HasTransAmount returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferModel) HasTransAmount() bool {
	if o != nil && !IsNil(o.TransAmount) {
		return true
	}

	return false
}

// SetTransAmount gets a reference to the given string and assigns it to the TransAmount field.
func (o *AlipayFundTransUniTransferModel) SetTransAmount(v string) {
	o.TransAmount = &v
}

func (o AlipayFundTransUniTransferModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundTransUniTransferModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.BusinessParams) {
		toSerialize["business_params"] = o.BusinessParams
	}
	if !IsNil(o.MutipleCurrencyDetail) {
		toSerialize["mutiple_currency_detail"] = o.MutipleCurrencyDetail
	}
	if !IsNil(o.OrderTitle) {
		toSerialize["order_title"] = o.OrderTitle
	}
	if !IsNil(o.OriginalOrderId) {
		toSerialize["original_order_id"] = o.OriginalOrderId
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PassbackParams) {
		toSerialize["passback_params"] = o.PassbackParams
	}
	if !IsNil(o.PayeeInfo) {
		toSerialize["payee_info"] = o.PayeeInfo
	}
	if !IsNil(o.PayerInfo) {
		toSerialize["payer_info"] = o.PayerInfo
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	if !IsNil(o.SignData) {
		toSerialize["sign_data"] = o.SignData
	}
	if !IsNil(o.TransAmount) {
		toSerialize["trans_amount"] = o.TransAmount
	}
	return toSerialize, nil
}

type NullableAlipayFundTransUniTransferModel struct {
	value *AlipayFundTransUniTransferModel
	isSet bool
}

func (v NullableAlipayFundTransUniTransferModel) Get() *AlipayFundTransUniTransferModel {
	return v.value
}

func (v *NullableAlipayFundTransUniTransferModel) Set(val *AlipayFundTransUniTransferModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundTransUniTransferModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundTransUniTransferModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundTransUniTransferModel(val *AlipayFundTransUniTransferModel) *NullableAlipayFundTransUniTransferModel {
	return &NullableAlipayFundTransUniTransferModel{value: val, isSet: true}
}

func (v NullableAlipayFundTransUniTransferModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundTransUniTransferModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
