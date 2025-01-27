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

// checks if the AlipayMarketingCampaignCashTriggerResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCampaignCashTriggerResponseModel{}

// AlipayMarketingCampaignCashTriggerResponseModel struct for AlipayMarketingCampaignCashTriggerResponseModel
type AlipayMarketingCampaignCashTriggerResponseModel struct {
	// 支付宝业务号,发奖成功时返回,调用者可提供此字符串进行问题排查与核对等
	BizNo *string `json:"biz_no,omitempty"`
	// 红包名称,创建活动时设置，用于账单列表和详情、红包列表和详情的展示
	CouponName *string `json:"coupon_name,omitempty"`
	// 用户领取失败的错误信息描述
	ErrorMsg *string `json:"error_msg,omitempty"`
	// 商户头像logo的图片url地址，用于账单和红包列表、详情的展示
	MerchantLogo *string `json:"merchant_logo,omitempty"`
	// 外部业务号,回填请求中的out_biz_no,请求参数中传了out_biz_no就会回传回去，如果不传就回传默认的\"default_out_biz_no\"，请求者可用于日志记录与核对等
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 发送红包商户签约pid，发奖成功时非空
	PartnerId *string `json:"partner_id,omitempty"`
	// 现金红包金额，发奖成功时非空，单位为元，保留两位小数
	PrizeAmount *string `json:"prize_amount,omitempty"`
	// 活动文案,用于账单的备注展示、红包详情页的文案展示
	PrizeMsg *string `json:"prize_msg,omitempty"`
	// 用户是否重复领取，如果重复领取，返回的是之前的中奖结果，如果是首次领取，则走发奖流程
	RepeatTriggerFlag *string `json:"repeat_trigger_flag,omitempty"`
	// 是否中奖结果状态，取值为true或false。如果为true表示发奖成功，这时返回的结果中的其他字段非空；如果为false表示发奖失败，这时返回的其他字段为空
	TriggerResult *string `json:"trigger_result,omitempty"`
}

// NewAlipayMarketingCampaignCashTriggerResponseModel instantiates a new AlipayMarketingCampaignCashTriggerResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCampaignCashTriggerResponseModel() *AlipayMarketingCampaignCashTriggerResponseModel {
	this := AlipayMarketingCampaignCashTriggerResponseModel{}
	return &this
}

// NewAlipayMarketingCampaignCashTriggerResponseModelWithDefaults instantiates a new AlipayMarketingCampaignCashTriggerResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCampaignCashTriggerResponseModelWithDefaults() *AlipayMarketingCampaignCashTriggerResponseModel {
	this := AlipayMarketingCampaignCashTriggerResponseModel{}
	return &this
}

// GetBizNo returns the BizNo field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetBizNo() string {
	if o == nil || IsNil(o.BizNo) {
		var ret string
		return ret
	}
	return *o.BizNo
}

// GetBizNoOk returns a tuple with the BizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.BizNo) {
		return nil, false
	}
	return o.BizNo, true
}

// HasBizNo returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) HasBizNo() bool {
	if o != nil && !IsNil(o.BizNo) {
		return true
	}

	return false
}

// SetBizNo gets a reference to the given string and assigns it to the BizNo field.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) SetBizNo(v string) {
	o.BizNo = &v
}

// GetCouponName returns the CouponName field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetCouponName() string {
	if o == nil || IsNil(o.CouponName) {
		var ret string
		return ret
	}
	return *o.CouponName
}

// GetCouponNameOk returns a tuple with the CouponName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetCouponNameOk() (*string, bool) {
	if o == nil || IsNil(o.CouponName) {
		return nil, false
	}
	return o.CouponName, true
}

// HasCouponName returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) HasCouponName() bool {
	if o != nil && !IsNil(o.CouponName) {
		return true
	}

	return false
}

// SetCouponName gets a reference to the given string and assigns it to the CouponName field.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) SetCouponName(v string) {
	o.CouponName = &v
}

// GetErrorMsg returns the ErrorMsg field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetErrorMsg() string {
	if o == nil || IsNil(o.ErrorMsg) {
		var ret string
		return ret
	}
	return *o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetErrorMsgOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMsg) {
		return nil, false
	}
	return o.ErrorMsg, true
}

// HasErrorMsg returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) HasErrorMsg() bool {
	if o != nil && !IsNil(o.ErrorMsg) {
		return true
	}

	return false
}

// SetErrorMsg gets a reference to the given string and assigns it to the ErrorMsg field.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) SetErrorMsg(v string) {
	o.ErrorMsg = &v
}

// GetMerchantLogo returns the MerchantLogo field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetMerchantLogo() string {
	if o == nil || IsNil(o.MerchantLogo) {
		var ret string
		return ret
	}
	return *o.MerchantLogo
}

// GetMerchantLogoOk returns a tuple with the MerchantLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetMerchantLogoOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantLogo) {
		return nil, false
	}
	return o.MerchantLogo, true
}

// HasMerchantLogo returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) HasMerchantLogo() bool {
	if o != nil && !IsNil(o.MerchantLogo) {
		return true
	}

	return false
}

// SetMerchantLogo gets a reference to the given string and assigns it to the MerchantLogo field.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) SetMerchantLogo(v string) {
	o.MerchantLogo = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetPrizeAmount returns the PrizeAmount field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetPrizeAmount() string {
	if o == nil || IsNil(o.PrizeAmount) {
		var ret string
		return ret
	}
	return *o.PrizeAmount
}

// GetPrizeAmountOk returns a tuple with the PrizeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetPrizeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PrizeAmount) {
		return nil, false
	}
	return o.PrizeAmount, true
}

// HasPrizeAmount returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) HasPrizeAmount() bool {
	if o != nil && !IsNil(o.PrizeAmount) {
		return true
	}

	return false
}

// SetPrizeAmount gets a reference to the given string and assigns it to the PrizeAmount field.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) SetPrizeAmount(v string) {
	o.PrizeAmount = &v
}

// GetPrizeMsg returns the PrizeMsg field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetPrizeMsg() string {
	if o == nil || IsNil(o.PrizeMsg) {
		var ret string
		return ret
	}
	return *o.PrizeMsg
}

// GetPrizeMsgOk returns a tuple with the PrizeMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetPrizeMsgOk() (*string, bool) {
	if o == nil || IsNil(o.PrizeMsg) {
		return nil, false
	}
	return o.PrizeMsg, true
}

// HasPrizeMsg returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) HasPrizeMsg() bool {
	if o != nil && !IsNil(o.PrizeMsg) {
		return true
	}

	return false
}

// SetPrizeMsg gets a reference to the given string and assigns it to the PrizeMsg field.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) SetPrizeMsg(v string) {
	o.PrizeMsg = &v
}

// GetRepeatTriggerFlag returns the RepeatTriggerFlag field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetRepeatTriggerFlag() string {
	if o == nil || IsNil(o.RepeatTriggerFlag) {
		var ret string
		return ret
	}
	return *o.RepeatTriggerFlag
}

// GetRepeatTriggerFlagOk returns a tuple with the RepeatTriggerFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetRepeatTriggerFlagOk() (*string, bool) {
	if o == nil || IsNil(o.RepeatTriggerFlag) {
		return nil, false
	}
	return o.RepeatTriggerFlag, true
}

// HasRepeatTriggerFlag returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) HasRepeatTriggerFlag() bool {
	if o != nil && !IsNil(o.RepeatTriggerFlag) {
		return true
	}

	return false
}

// SetRepeatTriggerFlag gets a reference to the given string and assigns it to the RepeatTriggerFlag field.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) SetRepeatTriggerFlag(v string) {
	o.RepeatTriggerFlag = &v
}

// GetTriggerResult returns the TriggerResult field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetTriggerResult() string {
	if o == nil || IsNil(o.TriggerResult) {
		var ret string
		return ret
	}
	return *o.TriggerResult
}

// GetTriggerResultOk returns a tuple with the TriggerResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) GetTriggerResultOk() (*string, bool) {
	if o == nil || IsNil(o.TriggerResult) {
		return nil, false
	}
	return o.TriggerResult, true
}

// HasTriggerResult returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) HasTriggerResult() bool {
	if o != nil && !IsNil(o.TriggerResult) {
		return true
	}

	return false
}

// SetTriggerResult gets a reference to the given string and assigns it to the TriggerResult field.
func (o *AlipayMarketingCampaignCashTriggerResponseModel) SetTriggerResult(v string) {
	o.TriggerResult = &v
}

func (o AlipayMarketingCampaignCashTriggerResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCampaignCashTriggerResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizNo) {
		toSerialize["biz_no"] = o.BizNo
	}
	if !IsNil(o.CouponName) {
		toSerialize["coupon_name"] = o.CouponName
	}
	if !IsNil(o.ErrorMsg) {
		toSerialize["error_msg"] = o.ErrorMsg
	}
	if !IsNil(o.MerchantLogo) {
		toSerialize["merchant_logo"] = o.MerchantLogo
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partner_id"] = o.PartnerId
	}
	if !IsNil(o.PrizeAmount) {
		toSerialize["prize_amount"] = o.PrizeAmount
	}
	if !IsNil(o.PrizeMsg) {
		toSerialize["prize_msg"] = o.PrizeMsg
	}
	if !IsNil(o.RepeatTriggerFlag) {
		toSerialize["repeat_trigger_flag"] = o.RepeatTriggerFlag
	}
	if !IsNil(o.TriggerResult) {
		toSerialize["trigger_result"] = o.TriggerResult
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCampaignCashTriggerResponseModel struct {
	value *AlipayMarketingCampaignCashTriggerResponseModel
	isSet bool
}

func (v NullableAlipayMarketingCampaignCashTriggerResponseModel) Get() *AlipayMarketingCampaignCashTriggerResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingCampaignCashTriggerResponseModel) Set(val *AlipayMarketingCampaignCashTriggerResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCampaignCashTriggerResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCampaignCashTriggerResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCampaignCashTriggerResponseModel(val *AlipayMarketingCampaignCashTriggerResponseModel) *NullableAlipayMarketingCampaignCashTriggerResponseModel {
	return &NullableAlipayMarketingCampaignCashTriggerResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCampaignCashTriggerResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCampaignCashTriggerResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
