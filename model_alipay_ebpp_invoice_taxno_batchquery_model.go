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

// checks if the AlipayEbppInvoiceTaxnoBatchqueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceTaxnoBatchqueryModel{}

// AlipayEbppInvoiceTaxnoBatchqueryModel struct for AlipayEbppInvoiceTaxnoBatchqueryModel
type AlipayEbppInvoiceTaxnoBatchqueryModel struct {
	// 默认值为false。true为输出交易信息，false为不输出交易信息。
	EnableTradeOut *string `json:"enable_trade_out,omitempty"`
	// 查询结束时间，精确到天（按开票日期查询） start_invoice_date和end_invoice_date传值要求 1.同时为空时，返回最近半年200条数据 2.必须同时为空 或 同时不为空 3.结束日期不能大于当前日期 4.开始时间和结束时间跨度不能超过6个月
	EndInvoiceDate *string `json:"end_invoice_date,omitempty"`
	// 查询票种列表。枚举值如下： *PLAIN：增值税电子普通发票； *SPECIAL：增值税专用发票； *ALL_ELECTRONIC_GENERAL： \"电子发票（普通发票）; *ALL_ELECTRONIC_SPECIAL： \"电子发票（专用发票）; *PLAIN_INVOICE：增值税普通发票； *PAPER_INVOICE：增值税普通发票（卷式）； *SALSE_INVOICE：机动车销售统一发票。
	InvoiceKindList []string `json:"invoice_kind_list,omitempty"`
	// 查询结果上限笔数，最大值20
	LimitSize *int32 `json:"limit_size,omitempty"`
	// 当前页码，为空时默认第一页
	PageNum *int32 `json:"page_num,omitempty"`
	// 发票要素获取应用场景。<a href=\"https://opendocs.alipay.com/open/10691/bv8s88\">\"拉\"模式报销</a> 固定为 INVOICE_EXPENSE 表示发票报销。
	Scene *string `json:"scene,omitempty"`
	// 查询起始时间，精确到天（按开票日期查询） start_invoice_date和end_invoice_date传值要求 1.同时为空时，返回最近半年200条数据 2.必须同时为空 或 同时不为空 3.结束日期不能大于当前日期 4.开始时间和结束时间跨度不能超过6个月
	StartInvoiceDate *string `json:"start_invoice_date,omitempty"`
	// 企业税号
	TaxNo *string `json:"tax_no,omitempty"`
}

// NewAlipayEbppInvoiceTaxnoBatchqueryModel instantiates a new AlipayEbppInvoiceTaxnoBatchqueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceTaxnoBatchqueryModel() *AlipayEbppInvoiceTaxnoBatchqueryModel {
	this := AlipayEbppInvoiceTaxnoBatchqueryModel{}
	return &this
}

// NewAlipayEbppInvoiceTaxnoBatchqueryModelWithDefaults instantiates a new AlipayEbppInvoiceTaxnoBatchqueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceTaxnoBatchqueryModelWithDefaults() *AlipayEbppInvoiceTaxnoBatchqueryModel {
	this := AlipayEbppInvoiceTaxnoBatchqueryModel{}
	return &this
}

// GetEnableTradeOut returns the EnableTradeOut field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetEnableTradeOut() string {
	if o == nil || IsNil(o.EnableTradeOut) {
		var ret string
		return ret
	}
	return *o.EnableTradeOut
}

// GetEnableTradeOutOk returns a tuple with the EnableTradeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetEnableTradeOutOk() (*string, bool) {
	if o == nil || IsNil(o.EnableTradeOut) {
		return nil, false
	}
	return o.EnableTradeOut, true
}

// HasEnableTradeOut returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) HasEnableTradeOut() bool {
	if o != nil && !IsNil(o.EnableTradeOut) {
		return true
	}

	return false
}

// SetEnableTradeOut gets a reference to the given string and assigns it to the EnableTradeOut field.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) SetEnableTradeOut(v string) {
	o.EnableTradeOut = &v
}

// GetEndInvoiceDate returns the EndInvoiceDate field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetEndInvoiceDate() string {
	if o == nil || IsNil(o.EndInvoiceDate) {
		var ret string
		return ret
	}
	return *o.EndInvoiceDate
}

// GetEndInvoiceDateOk returns a tuple with the EndInvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetEndInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndInvoiceDate) {
		return nil, false
	}
	return o.EndInvoiceDate, true
}

// HasEndInvoiceDate returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) HasEndInvoiceDate() bool {
	if o != nil && !IsNil(o.EndInvoiceDate) {
		return true
	}

	return false
}

// SetEndInvoiceDate gets a reference to the given string and assigns it to the EndInvoiceDate field.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) SetEndInvoiceDate(v string) {
	o.EndInvoiceDate = &v
}

// GetInvoiceKindList returns the InvoiceKindList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetInvoiceKindList() []string {
	if o == nil || IsNil(o.InvoiceKindList) {
		var ret []string
		return ret
	}
	return o.InvoiceKindList
}

// GetInvoiceKindListOk returns a tuple with the InvoiceKindList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetInvoiceKindListOk() ([]string, bool) {
	if o == nil || IsNil(o.InvoiceKindList) {
		return nil, false
	}
	return o.InvoiceKindList, true
}

// HasInvoiceKindList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) HasInvoiceKindList() bool {
	if o != nil && !IsNil(o.InvoiceKindList) {
		return true
	}

	return false
}

// SetInvoiceKindList gets a reference to the given []string and assigns it to the InvoiceKindList field.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) SetInvoiceKindList(v []string) {
	o.InvoiceKindList = v
}

// GetLimitSize returns the LimitSize field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetLimitSize() int32 {
	if o == nil || IsNil(o.LimitSize) {
		var ret int32
		return ret
	}
	return *o.LimitSize
}

// GetLimitSizeOk returns a tuple with the LimitSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetLimitSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitSize) {
		return nil, false
	}
	return o.LimitSize, true
}

// HasLimitSize returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) HasLimitSize() bool {
	if o != nil && !IsNil(o.LimitSize) {
		return true
	}

	return false
}

// SetLimitSize gets a reference to the given int32 and assigns it to the LimitSize field.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) SetLimitSize(v int32) {
	o.LimitSize = &v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetScene returns the Scene field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetScene() string {
	if o == nil || IsNil(o.Scene) {
		var ret string
		return ret
	}
	return *o.Scene
}

// GetSceneOk returns a tuple with the Scene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetSceneOk() (*string, bool) {
	if o == nil || IsNil(o.Scene) {
		return nil, false
	}
	return o.Scene, true
}

// HasScene returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) HasScene() bool {
	if o != nil && !IsNil(o.Scene) {
		return true
	}

	return false
}

// SetScene gets a reference to the given string and assigns it to the Scene field.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) SetScene(v string) {
	o.Scene = &v
}

// GetStartInvoiceDate returns the StartInvoiceDate field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetStartInvoiceDate() string {
	if o == nil || IsNil(o.StartInvoiceDate) {
		var ret string
		return ret
	}
	return *o.StartInvoiceDate
}

// GetStartInvoiceDateOk returns a tuple with the StartInvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetStartInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartInvoiceDate) {
		return nil, false
	}
	return o.StartInvoiceDate, true
}

// HasStartInvoiceDate returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) HasStartInvoiceDate() bool {
	if o != nil && !IsNil(o.StartInvoiceDate) {
		return true
	}

	return false
}

// SetStartInvoiceDate gets a reference to the given string and assigns it to the StartInvoiceDate field.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) SetStartInvoiceDate(v string) {
	o.StartInvoiceDate = &v
}

// GetTaxNo returns the TaxNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetTaxNo() string {
	if o == nil || IsNil(o.TaxNo) {
		var ret string
		return ret
	}
	return *o.TaxNo
}

// GetTaxNoOk returns a tuple with the TaxNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) GetTaxNoOk() (*string, bool) {
	if o == nil || IsNil(o.TaxNo) {
		return nil, false
	}
	return o.TaxNo, true
}

// HasTaxNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) HasTaxNo() bool {
	if o != nil && !IsNil(o.TaxNo) {
		return true
	}

	return false
}

// SetTaxNo gets a reference to the given string and assigns it to the TaxNo field.
func (o *AlipayEbppInvoiceTaxnoBatchqueryModel) SetTaxNo(v string) {
	o.TaxNo = &v
}

func (o AlipayEbppInvoiceTaxnoBatchqueryModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceTaxnoBatchqueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableTradeOut) {
		toSerialize["enable_trade_out"] = o.EnableTradeOut
	}
	if !IsNil(o.EndInvoiceDate) {
		toSerialize["end_invoice_date"] = o.EndInvoiceDate
	}
	if !IsNil(o.InvoiceKindList) {
		toSerialize["invoice_kind_list"] = o.InvoiceKindList
	}
	if !IsNil(o.LimitSize) {
		toSerialize["limit_size"] = o.LimitSize
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.Scene) {
		toSerialize["scene"] = o.Scene
	}
	if !IsNil(o.StartInvoiceDate) {
		toSerialize["start_invoice_date"] = o.StartInvoiceDate
	}
	if !IsNil(o.TaxNo) {
		toSerialize["tax_no"] = o.TaxNo
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceTaxnoBatchqueryModel struct {
	value *AlipayEbppInvoiceTaxnoBatchqueryModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceTaxnoBatchqueryModel) Get() *AlipayEbppInvoiceTaxnoBatchqueryModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceTaxnoBatchqueryModel) Set(val *AlipayEbppInvoiceTaxnoBatchqueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceTaxnoBatchqueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceTaxnoBatchqueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceTaxnoBatchqueryModel(val *AlipayEbppInvoiceTaxnoBatchqueryModel) *NullableAlipayEbppInvoiceTaxnoBatchqueryModel {
	return &NullableAlipayEbppInvoiceTaxnoBatchqueryModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceTaxnoBatchqueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceTaxnoBatchqueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
