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

// checks if the EinvTrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EinvTrade{}

// EinvTrade struct for EinvTrade
type EinvTrade struct {
	// 订单编号
	BillNo *string `json:"bill_no,omitempty"`
	// 下单时间
	BillTime *string `json:"bill_time,omitempty"`
	// 商户所在城市(经营地址)
	CityName *string `json:"city_name,omitempty"`
	// 账单明细信息，酒店水单信息，行程单信息，餐饮小票信息
	DetailJson *string `json:"detail_json,omitempty"`
	// 账单明细信息，酒店水单，行程单，餐饮小票等pdf原件链接
	DownloadUrl *string `json:"download_url,omitempty"`
	// 扩展参数  不同组的k-v通过换行符区分
	ExtendMap *string `json:"extend_map,omitempty"`
	// 商家名称（显示名称，非企业名称，餐饮店、酒店、打车平台名称）
	MerchantName *string `json:"merchant_name,omitempty"`
	// 透传字段，不做处理，用于isv向后续报销税控方传递特殊信息标记
	OutJson *string `json:"out_json,omitempty"`
	// 销方名称
	PayeeName *string `json:"payee_name,omitempty"`
	// 支付金额，单位（元）， 对应账单中的交易金额
	PaymentAmount *string `json:"payment_amount,omitempty"`
	// 支付时间  对应账单中的账单日期
	PaymentTime *string `json:"payment_time,omitempty"`
	// 交易类型来源 需要按照枚举映射  consume 账单  hotel 酒店水单  itinerary 打车行程单  catering 餐饮小票
	Souce *string `json:"souce,omitempty"`
	// 交易类型/账单分类
	TradeType *string `json:"trade_type,omitempty"`
}

// NewEinvTrade instantiates a new EinvTrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEinvTrade() *EinvTrade {
	this := EinvTrade{}
	return &this
}

// NewEinvTradeWithDefaults instantiates a new EinvTrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEinvTradeWithDefaults() *EinvTrade {
	this := EinvTrade{}
	return &this
}

// GetBillNo returns the BillNo field value if set, zero value otherwise.
func (o *EinvTrade) GetBillNo() string {
	if o == nil || IsNil(o.BillNo) {
		var ret string
		return ret
	}
	return *o.BillNo
}

// GetBillNoOk returns a tuple with the BillNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetBillNoOk() (*string, bool) {
	if o == nil || IsNil(o.BillNo) {
		return nil, false
	}
	return o.BillNo, true
}

// HasBillNo returns a boolean if a field has been set.
func (o *EinvTrade) HasBillNo() bool {
	if o != nil && !IsNil(o.BillNo) {
		return true
	}

	return false
}

// SetBillNo gets a reference to the given string and assigns it to the BillNo field.
func (o *EinvTrade) SetBillNo(v string) {
	o.BillNo = &v
}

// GetBillTime returns the BillTime field value if set, zero value otherwise.
func (o *EinvTrade) GetBillTime() string {
	if o == nil || IsNil(o.BillTime) {
		var ret string
		return ret
	}
	return *o.BillTime
}

// GetBillTimeOk returns a tuple with the BillTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetBillTimeOk() (*string, bool) {
	if o == nil || IsNil(o.BillTime) {
		return nil, false
	}
	return o.BillTime, true
}

// HasBillTime returns a boolean if a field has been set.
func (o *EinvTrade) HasBillTime() bool {
	if o != nil && !IsNil(o.BillTime) {
		return true
	}

	return false
}

// SetBillTime gets a reference to the given string and assigns it to the BillTime field.
func (o *EinvTrade) SetBillTime(v string) {
	o.BillTime = &v
}

// GetCityName returns the CityName field value if set, zero value otherwise.
func (o *EinvTrade) GetCityName() string {
	if o == nil || IsNil(o.CityName) {
		var ret string
		return ret
	}
	return *o.CityName
}

// GetCityNameOk returns a tuple with the CityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetCityNameOk() (*string, bool) {
	if o == nil || IsNil(o.CityName) {
		return nil, false
	}
	return o.CityName, true
}

// HasCityName returns a boolean if a field has been set.
func (o *EinvTrade) HasCityName() bool {
	if o != nil && !IsNil(o.CityName) {
		return true
	}

	return false
}

// SetCityName gets a reference to the given string and assigns it to the CityName field.
func (o *EinvTrade) SetCityName(v string) {
	o.CityName = &v
}

// GetDetailJson returns the DetailJson field value if set, zero value otherwise.
func (o *EinvTrade) GetDetailJson() string {
	if o == nil || IsNil(o.DetailJson) {
		var ret string
		return ret
	}
	return *o.DetailJson
}

// GetDetailJsonOk returns a tuple with the DetailJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetDetailJsonOk() (*string, bool) {
	if o == nil || IsNil(o.DetailJson) {
		return nil, false
	}
	return o.DetailJson, true
}

// HasDetailJson returns a boolean if a field has been set.
func (o *EinvTrade) HasDetailJson() bool {
	if o != nil && !IsNil(o.DetailJson) {
		return true
	}

	return false
}

// SetDetailJson gets a reference to the given string and assigns it to the DetailJson field.
func (o *EinvTrade) SetDetailJson(v string) {
	o.DetailJson = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *EinvTrade) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *EinvTrade) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *EinvTrade) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetExtendMap returns the ExtendMap field value if set, zero value otherwise.
func (o *EinvTrade) GetExtendMap() string {
	if o == nil || IsNil(o.ExtendMap) {
		var ret string
		return ret
	}
	return *o.ExtendMap
}

// GetExtendMapOk returns a tuple with the ExtendMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetExtendMapOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendMap) {
		return nil, false
	}
	return o.ExtendMap, true
}

// HasExtendMap returns a boolean if a field has been set.
func (o *EinvTrade) HasExtendMap() bool {
	if o != nil && !IsNil(o.ExtendMap) {
		return true
	}

	return false
}

// SetExtendMap gets a reference to the given string and assigns it to the ExtendMap field.
func (o *EinvTrade) SetExtendMap(v string) {
	o.ExtendMap = &v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *EinvTrade) GetMerchantName() string {
	if o == nil || IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *EinvTrade) HasMerchantName() bool {
	if o != nil && !IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *EinvTrade) SetMerchantName(v string) {
	o.MerchantName = &v
}

// GetOutJson returns the OutJson field value if set, zero value otherwise.
func (o *EinvTrade) GetOutJson() string {
	if o == nil || IsNil(o.OutJson) {
		var ret string
		return ret
	}
	return *o.OutJson
}

// GetOutJsonOk returns a tuple with the OutJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetOutJsonOk() (*string, bool) {
	if o == nil || IsNil(o.OutJson) {
		return nil, false
	}
	return o.OutJson, true
}

// HasOutJson returns a boolean if a field has been set.
func (o *EinvTrade) HasOutJson() bool {
	if o != nil && !IsNil(o.OutJson) {
		return true
	}

	return false
}

// SetOutJson gets a reference to the given string and assigns it to the OutJson field.
func (o *EinvTrade) SetOutJson(v string) {
	o.OutJson = &v
}

// GetPayeeName returns the PayeeName field value if set, zero value otherwise.
func (o *EinvTrade) GetPayeeName() string {
	if o == nil || IsNil(o.PayeeName) {
		var ret string
		return ret
	}
	return *o.PayeeName
}

// GetPayeeNameOk returns a tuple with the PayeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetPayeeNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeName) {
		return nil, false
	}
	return o.PayeeName, true
}

// HasPayeeName returns a boolean if a field has been set.
func (o *EinvTrade) HasPayeeName() bool {
	if o != nil && !IsNil(o.PayeeName) {
		return true
	}

	return false
}

// SetPayeeName gets a reference to the given string and assigns it to the PayeeName field.
func (o *EinvTrade) SetPayeeName(v string) {
	o.PayeeName = &v
}

// GetPaymentAmount returns the PaymentAmount field value if set, zero value otherwise.
func (o *EinvTrade) GetPaymentAmount() string {
	if o == nil || IsNil(o.PaymentAmount) {
		var ret string
		return ret
	}
	return *o.PaymentAmount
}

// GetPaymentAmountOk returns a tuple with the PaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetPaymentAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentAmount) {
		return nil, false
	}
	return o.PaymentAmount, true
}

// HasPaymentAmount returns a boolean if a field has been set.
func (o *EinvTrade) HasPaymentAmount() bool {
	if o != nil && !IsNil(o.PaymentAmount) {
		return true
	}

	return false
}

// SetPaymentAmount gets a reference to the given string and assigns it to the PaymentAmount field.
func (o *EinvTrade) SetPaymentAmount(v string) {
	o.PaymentAmount = &v
}

// GetPaymentTime returns the PaymentTime field value if set, zero value otherwise.
func (o *EinvTrade) GetPaymentTime() string {
	if o == nil || IsNil(o.PaymentTime) {
		var ret string
		return ret
	}
	return *o.PaymentTime
}

// GetPaymentTimeOk returns a tuple with the PaymentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetPaymentTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTime) {
		return nil, false
	}
	return o.PaymentTime, true
}

// HasPaymentTime returns a boolean if a field has been set.
func (o *EinvTrade) HasPaymentTime() bool {
	if o != nil && !IsNil(o.PaymentTime) {
		return true
	}

	return false
}

// SetPaymentTime gets a reference to the given string and assigns it to the PaymentTime field.
func (o *EinvTrade) SetPaymentTime(v string) {
	o.PaymentTime = &v
}

// GetSouce returns the Souce field value if set, zero value otherwise.
func (o *EinvTrade) GetSouce() string {
	if o == nil || IsNil(o.Souce) {
		var ret string
		return ret
	}
	return *o.Souce
}

// GetSouceOk returns a tuple with the Souce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetSouceOk() (*string, bool) {
	if o == nil || IsNil(o.Souce) {
		return nil, false
	}
	return o.Souce, true
}

// HasSouce returns a boolean if a field has been set.
func (o *EinvTrade) HasSouce() bool {
	if o != nil && !IsNil(o.Souce) {
		return true
	}

	return false
}

// SetSouce gets a reference to the given string and assigns it to the Souce field.
func (o *EinvTrade) SetSouce(v string) {
	o.Souce = &v
}

// GetTradeType returns the TradeType field value if set, zero value otherwise.
func (o *EinvTrade) GetTradeType() string {
	if o == nil || IsNil(o.TradeType) {
		var ret string
		return ret
	}
	return *o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EinvTrade) GetTradeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TradeType) {
		return nil, false
	}
	return o.TradeType, true
}

// HasTradeType returns a boolean if a field has been set.
func (o *EinvTrade) HasTradeType() bool {
	if o != nil && !IsNil(o.TradeType) {
		return true
	}

	return false
}

// SetTradeType gets a reference to the given string and assigns it to the TradeType field.
func (o *EinvTrade) SetTradeType(v string) {
	o.TradeType = &v
}

func (o EinvTrade) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EinvTrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillNo) {
		toSerialize["bill_no"] = o.BillNo
	}
	if !IsNil(o.BillTime) {
		toSerialize["bill_time"] = o.BillTime
	}
	if !IsNil(o.CityName) {
		toSerialize["city_name"] = o.CityName
	}
	if !IsNil(o.DetailJson) {
		toSerialize["detail_json"] = o.DetailJson
	}
	if !IsNil(o.DownloadUrl) {
		toSerialize["download_url"] = o.DownloadUrl
	}
	if !IsNil(o.ExtendMap) {
		toSerialize["extend_map"] = o.ExtendMap
	}
	if !IsNil(o.MerchantName) {
		toSerialize["merchant_name"] = o.MerchantName
	}
	if !IsNil(o.OutJson) {
		toSerialize["out_json"] = o.OutJson
	}
	if !IsNil(o.PayeeName) {
		toSerialize["payee_name"] = o.PayeeName
	}
	if !IsNil(o.PaymentAmount) {
		toSerialize["payment_amount"] = o.PaymentAmount
	}
	if !IsNil(o.PaymentTime) {
		toSerialize["payment_time"] = o.PaymentTime
	}
	if !IsNil(o.Souce) {
		toSerialize["souce"] = o.Souce
	}
	if !IsNil(o.TradeType) {
		toSerialize["trade_type"] = o.TradeType
	}
	return toSerialize, nil
}

type NullableEinvTrade struct {
	value *EinvTrade
	isSet bool
}

func (v NullableEinvTrade) Get() *EinvTrade {
	return v.value
}

func (v *NullableEinvTrade) Set(val *EinvTrade) {
	v.value = val
	v.isSet = true
}

func (v NullableEinvTrade) IsSet() bool {
	return v.isSet
}

func (v *NullableEinvTrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEinvTrade(val *EinvTrade) *NullableEinvTrade {
	return &NullableEinvTrade{value: val, isSet: true}
}

func (v NullableEinvTrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEinvTrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
