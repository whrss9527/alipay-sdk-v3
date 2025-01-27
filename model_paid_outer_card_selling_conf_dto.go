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

// checks if the PaidOuterCardSellingConfDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaidOuterCardSellingConfDTO{}

// PaidOuterCardSellingConfDTO struct for PaidOuterCardSellingConfDTO
type PaidOuterCardSellingConfDTO struct {
	// 售卖结束时间。 格式：yyyy-MM-dd HH:mm:ss 如果永久有效则不传。
	EndDate *string `json:"end_date,omitempty"`
	// 售卖方案列表
	PriceDetail []PaidOuterCardPriceDetailDTO `json:"price_detail,omitempty"`
	// 售卖地址
	SellingUrl *string `json:"selling_url,omitempty"`
	// 售卖开始时间。格式：yyyy-MM-dd HH:mm:ss
	StartDate *string `json:"start_date,omitempty"`
}

// NewPaidOuterCardSellingConfDTO instantiates a new PaidOuterCardSellingConfDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaidOuterCardSellingConfDTO() *PaidOuterCardSellingConfDTO {
	this := PaidOuterCardSellingConfDTO{}
	return &this
}

// NewPaidOuterCardSellingConfDTOWithDefaults instantiates a new PaidOuterCardSellingConfDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaidOuterCardSellingConfDTOWithDefaults() *PaidOuterCardSellingConfDTO {
	this := PaidOuterCardSellingConfDTO{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *PaidOuterCardSellingConfDTO) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidOuterCardSellingConfDTO) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *PaidOuterCardSellingConfDTO) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *PaidOuterCardSellingConfDTO) SetEndDate(v string) {
	o.EndDate = &v
}

// GetPriceDetail returns the PriceDetail field value if set, zero value otherwise.
func (o *PaidOuterCardSellingConfDTO) GetPriceDetail() []PaidOuterCardPriceDetailDTO {
	if o == nil || IsNil(o.PriceDetail) {
		var ret []PaidOuterCardPriceDetailDTO
		return ret
	}
	return o.PriceDetail
}

// GetPriceDetailOk returns a tuple with the PriceDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidOuterCardSellingConfDTO) GetPriceDetailOk() ([]PaidOuterCardPriceDetailDTO, bool) {
	if o == nil || IsNil(o.PriceDetail) {
		return nil, false
	}
	return o.PriceDetail, true
}

// HasPriceDetail returns a boolean if a field has been set.
func (o *PaidOuterCardSellingConfDTO) HasPriceDetail() bool {
	if o != nil && !IsNil(o.PriceDetail) {
		return true
	}

	return false
}

// SetPriceDetail gets a reference to the given []PaidOuterCardPriceDetailDTO and assigns it to the PriceDetail field.
func (o *PaidOuterCardSellingConfDTO) SetPriceDetail(v []PaidOuterCardPriceDetailDTO) {
	o.PriceDetail = v
}

// GetSellingUrl returns the SellingUrl field value if set, zero value otherwise.
func (o *PaidOuterCardSellingConfDTO) GetSellingUrl() string {
	if o == nil || IsNil(o.SellingUrl) {
		var ret string
		return ret
	}
	return *o.SellingUrl
}

// GetSellingUrlOk returns a tuple with the SellingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidOuterCardSellingConfDTO) GetSellingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SellingUrl) {
		return nil, false
	}
	return o.SellingUrl, true
}

// HasSellingUrl returns a boolean if a field has been set.
func (o *PaidOuterCardSellingConfDTO) HasSellingUrl() bool {
	if o != nil && !IsNil(o.SellingUrl) {
		return true
	}

	return false
}

// SetSellingUrl gets a reference to the given string and assigns it to the SellingUrl field.
func (o *PaidOuterCardSellingConfDTO) SetSellingUrl(v string) {
	o.SellingUrl = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *PaidOuterCardSellingConfDTO) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidOuterCardSellingConfDTO) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *PaidOuterCardSellingConfDTO) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *PaidOuterCardSellingConfDTO) SetStartDate(v string) {
	o.StartDate = &v
}

func (o PaidOuterCardSellingConfDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaidOuterCardSellingConfDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.PriceDetail) {
		toSerialize["price_detail"] = o.PriceDetail
	}
	if !IsNil(o.SellingUrl) {
		toSerialize["selling_url"] = o.SellingUrl
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	return toSerialize, nil
}

type NullablePaidOuterCardSellingConfDTO struct {
	value *PaidOuterCardSellingConfDTO
	isSet bool
}

func (v NullablePaidOuterCardSellingConfDTO) Get() *PaidOuterCardSellingConfDTO {
	return v.value
}

func (v *NullablePaidOuterCardSellingConfDTO) Set(val *PaidOuterCardSellingConfDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePaidOuterCardSellingConfDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePaidOuterCardSellingConfDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaidOuterCardSellingConfDTO(val *PaidOuterCardSellingConfDTO) *NullablePaidOuterCardSellingConfDTO {
	return &NullablePaidOuterCardSellingConfDTO{value: val, isSet: true}
}

func (v NullablePaidOuterCardSellingConfDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaidOuterCardSellingConfDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
