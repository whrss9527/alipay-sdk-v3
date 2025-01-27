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

// checks if the AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel{}

// AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel struct for AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel
type AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel struct {
	// 共同账号ID
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 余额失效时间（默认有效期100年）
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// 余额生效时间(默认当前时间)
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// 外部操作幂等ID（接入方接口调用幂等控制ID）
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// 切换open_id前请使用：余额所属者ID owner_type为员工时为员工支付宝ID
	OwnerId *string `json:"owner_id,omitempty"`
	// 切换open_id后请使用：余额所属者ID owner_type为员工时为员工open_id
	OwnerOpenId *string `json:"owner_open_id,omitempty"`
	// 余额所属者类型 EMPLOYEE: 员工
	OwnerType *string `json:"owner_type,omitempty"`
	// 外部平台编码（通常为接入方大写英文缩写）
	Platform *string `json:"platform,omitempty"`
	// 余额同步模式，默认 DEFAULT_STANDARD
	QuotaModel *string `json:"quota_model,omitempty"`
	// 余额，以（分）为单位 特殊说明：余额不超过100000元
	QuotaValue *string `json:"quota_value,omitempty"`
}

// NewAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel instantiates a new AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel() *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel {
	this := AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel{}
	return &this
}

// NewAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModelWithDefaults instantiates a new AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModelWithDefaults() *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel {
	this := AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetEffectiveEndDate returns the EffectiveEndDate field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetEffectiveEndDate() string {
	if o == nil || IsNil(o.EffectiveEndDate) {
		var ret string
		return ret
	}
	return *o.EffectiveEndDate
}

// GetEffectiveEndDateOk returns a tuple with the EffectiveEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetEffectiveEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveEndDate) {
		return nil, false
	}
	return o.EffectiveEndDate, true
}

// HasEffectiveEndDate returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasEffectiveEndDate() bool {
	if o != nil && !IsNil(o.EffectiveEndDate) {
		return true
	}

	return false
}

// SetEffectiveEndDate gets a reference to the given string and assigns it to the EffectiveEndDate field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetEffectiveEndDate(v string) {
	o.EffectiveEndDate = &v
}

// GetEffectiveStartDate returns the EffectiveStartDate field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetEffectiveStartDate() string {
	if o == nil || IsNil(o.EffectiveStartDate) {
		var ret string
		return ret
	}
	return *o.EffectiveStartDate
}

// GetEffectiveStartDateOk returns a tuple with the EffectiveStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetEffectiveStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveStartDate) {
		return nil, false
	}
	return o.EffectiveStartDate, true
}

// HasEffectiveStartDate returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasEffectiveStartDate() bool {
	if o != nil && !IsNil(o.EffectiveStartDate) {
		return true
	}

	return false
}

// SetEffectiveStartDate gets a reference to the given string and assigns it to the EffectiveStartDate field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetEffectiveStartDate(v string) {
	o.EffectiveStartDate = &v
}

// GetOuterSourceId returns the OuterSourceId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetOuterSourceId() string {
	if o == nil || IsNil(o.OuterSourceId) {
		var ret string
		return ret
	}
	return *o.OuterSourceId
}

// GetOuterSourceIdOk returns a tuple with the OuterSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetOuterSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.OuterSourceId) {
		return nil, false
	}
	return o.OuterSourceId, true
}

// HasOuterSourceId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasOuterSourceId() bool {
	if o != nil && !IsNil(o.OuterSourceId) {
		return true
	}

	return false
}

// SetOuterSourceId gets a reference to the given string and assigns it to the OuterSourceId field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetOuterSourceId(v string) {
	o.OuterSourceId = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerOpenId returns the OwnerOpenId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetOwnerOpenId() string {
	if o == nil || IsNil(o.OwnerOpenId) {
		var ret string
		return ret
	}
	return *o.OwnerOpenId
}

// GetOwnerOpenIdOk returns a tuple with the OwnerOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetOwnerOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerOpenId) {
		return nil, false
	}
	return o.OwnerOpenId, true
}

// HasOwnerOpenId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasOwnerOpenId() bool {
	if o != nil && !IsNil(o.OwnerOpenId) {
		return true
	}

	return false
}

// SetOwnerOpenId gets a reference to the given string and assigns it to the OwnerOpenId field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetOwnerOpenId(v string) {
	o.OwnerOpenId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetOwnerType(v string) {
	o.OwnerType = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetPlatform(v string) {
	o.Platform = &v
}

// GetQuotaModel returns the QuotaModel field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetQuotaModel() string {
	if o == nil || IsNil(o.QuotaModel) {
		var ret string
		return ret
	}
	return *o.QuotaModel
}

// GetQuotaModelOk returns a tuple with the QuotaModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetQuotaModelOk() (*string, bool) {
	if o == nil || IsNil(o.QuotaModel) {
		return nil, false
	}
	return o.QuotaModel, true
}

// HasQuotaModel returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasQuotaModel() bool {
	if o != nil && !IsNil(o.QuotaModel) {
		return true
	}

	return false
}

// SetQuotaModel gets a reference to the given string and assigns it to the QuotaModel field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetQuotaModel(v string) {
	o.QuotaModel = &v
}

// GetQuotaValue returns the QuotaValue field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetQuotaValue() string {
	if o == nil || IsNil(o.QuotaValue) {
		var ret string
		return ret
	}
	return *o.QuotaValue
}

// GetQuotaValueOk returns a tuple with the QuotaValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) GetQuotaValueOk() (*string, bool) {
	if o == nil || IsNil(o.QuotaValue) {
		return nil, false
	}
	return o.QuotaValue, true
}

// HasQuotaValue returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) HasQuotaValue() bool {
	if o != nil && !IsNil(o.QuotaValue) {
		return true
	}

	return false
}

// SetQuotaValue gets a reference to the given string and assigns it to the QuotaValue field.
func (o *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) SetQuotaValue(v string) {
	o.QuotaValue = &v
}

func (o AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.EffectiveEndDate) {
		toSerialize["effective_end_date"] = o.EffectiveEndDate
	}
	if !IsNil(o.EffectiveStartDate) {
		toSerialize["effective_start_date"] = o.EffectiveStartDate
	}
	if !IsNil(o.OuterSourceId) {
		toSerialize["outer_source_id"] = o.OuterSourceId
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.OwnerOpenId) {
		toSerialize["owner_open_id"] = o.OwnerOpenId
	}
	if !IsNil(o.OwnerType) {
		toSerialize["owner_type"] = o.OwnerType
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.QuotaModel) {
		toSerialize["quota_model"] = o.QuotaModel
	}
	if !IsNil(o.QuotaValue) {
		toSerialize["quota_value"] = o.QuotaValue
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel struct {
	value *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) Get() *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) Set(val *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel(val *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) *NullableAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel {
	return &NullableAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
