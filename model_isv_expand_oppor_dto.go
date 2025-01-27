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

// checks if the IsvExpandOpporDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsvExpandOpporDTO{}

// IsvExpandOpporDTO struct for IsvExpandOpporDTO
type IsvExpandOpporDTO struct {
	// 商机地址
	Address *string `json:"address,omitempty"`
	// 商机Id
	LeadsId *string `json:"leads_id,omitempty"`
	// 商机名称
	Name *string `json:"name,omitempty"`
	// 商机作业Id
	OpporId *string `json:"oppor_id,omitempty"`
	// 外部幂等唯一键
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 商机联系电话
	Phone *string `json:"phone,omitempty"`
	// 商机作业状态
	Status *string `json:"status,omitempty"`
}

// NewIsvExpandOpporDTO instantiates a new IsvExpandOpporDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsvExpandOpporDTO() *IsvExpandOpporDTO {
	this := IsvExpandOpporDTO{}
	return &this
}

// NewIsvExpandOpporDTOWithDefaults instantiates a new IsvExpandOpporDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsvExpandOpporDTOWithDefaults() *IsvExpandOpporDTO {
	this := IsvExpandOpporDTO{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *IsvExpandOpporDTO) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvExpandOpporDTO) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *IsvExpandOpporDTO) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *IsvExpandOpporDTO) SetAddress(v string) {
	o.Address = &v
}

// GetLeadsId returns the LeadsId field value if set, zero value otherwise.
func (o *IsvExpandOpporDTO) GetLeadsId() string {
	if o == nil || IsNil(o.LeadsId) {
		var ret string
		return ret
	}
	return *o.LeadsId
}

// GetLeadsIdOk returns a tuple with the LeadsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvExpandOpporDTO) GetLeadsIdOk() (*string, bool) {
	if o == nil || IsNil(o.LeadsId) {
		return nil, false
	}
	return o.LeadsId, true
}

// HasLeadsId returns a boolean if a field has been set.
func (o *IsvExpandOpporDTO) HasLeadsId() bool {
	if o != nil && !IsNil(o.LeadsId) {
		return true
	}

	return false
}

// SetLeadsId gets a reference to the given string and assigns it to the LeadsId field.
func (o *IsvExpandOpporDTO) SetLeadsId(v string) {
	o.LeadsId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IsvExpandOpporDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvExpandOpporDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IsvExpandOpporDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IsvExpandOpporDTO) SetName(v string) {
	o.Name = &v
}

// GetOpporId returns the OpporId field value if set, zero value otherwise.
func (o *IsvExpandOpporDTO) GetOpporId() string {
	if o == nil || IsNil(o.OpporId) {
		var ret string
		return ret
	}
	return *o.OpporId
}

// GetOpporIdOk returns a tuple with the OpporId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvExpandOpporDTO) GetOpporIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpporId) {
		return nil, false
	}
	return o.OpporId, true
}

// HasOpporId returns a boolean if a field has been set.
func (o *IsvExpandOpporDTO) HasOpporId() bool {
	if o != nil && !IsNil(o.OpporId) {
		return true
	}

	return false
}

// SetOpporId gets a reference to the given string and assigns it to the OpporId field.
func (o *IsvExpandOpporDTO) SetOpporId(v string) {
	o.OpporId = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *IsvExpandOpporDTO) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvExpandOpporDTO) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *IsvExpandOpporDTO) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *IsvExpandOpporDTO) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *IsvExpandOpporDTO) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvExpandOpporDTO) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *IsvExpandOpporDTO) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *IsvExpandOpporDTO) SetPhone(v string) {
	o.Phone = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IsvExpandOpporDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvExpandOpporDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IsvExpandOpporDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IsvExpandOpporDTO) SetStatus(v string) {
	o.Status = &v
}

func (o IsvExpandOpporDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsvExpandOpporDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.LeadsId) {
		toSerialize["leads_id"] = o.LeadsId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OpporId) {
		toSerialize["oppor_id"] = o.OpporId
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableIsvExpandOpporDTO struct {
	value *IsvExpandOpporDTO
	isSet bool
}

func (v NullableIsvExpandOpporDTO) Get() *IsvExpandOpporDTO {
	return v.value
}

func (v *NullableIsvExpandOpporDTO) Set(val *IsvExpandOpporDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableIsvExpandOpporDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableIsvExpandOpporDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsvExpandOpporDTO(val *IsvExpandOpporDTO) *NullableIsvExpandOpporDTO {
	return &NullableIsvExpandOpporDTO{value: val, isSet: true}
}

func (v NullableIsvExpandOpporDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsvExpandOpporDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
