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

// checks if the AliosOpenAutoInfoQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AliosOpenAutoInfoQueryResponseModel{}

// AliosOpenAutoInfoQueryResponseModel struct for AliosOpenAutoInfoQueryResponseModel
type AliosOpenAutoInfoQueryResponseModel struct {
	// 发动机号
	EngineNo *string `json:"engine_no,omitempty"`
	// 车牌号
	LicenseNo *string `json:"license_no,omitempty"`
	// 用户激活时间
	UserActivationTime *string `json:"user_activation_time,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicle_type,omitempty"`
	// 车辆识别号码
	Vin *string `json:"vin,omitempty"`
}

// NewAliosOpenAutoInfoQueryResponseModel instantiates a new AliosOpenAutoInfoQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAliosOpenAutoInfoQueryResponseModel() *AliosOpenAutoInfoQueryResponseModel {
	this := AliosOpenAutoInfoQueryResponseModel{}
	return &this
}

// NewAliosOpenAutoInfoQueryResponseModelWithDefaults instantiates a new AliosOpenAutoInfoQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliosOpenAutoInfoQueryResponseModelWithDefaults() *AliosOpenAutoInfoQueryResponseModel {
	this := AliosOpenAutoInfoQueryResponseModel{}
	return &this
}

// GetEngineNo returns the EngineNo field value if set, zero value otherwise.
func (o *AliosOpenAutoInfoQueryResponseModel) GetEngineNo() string {
	if o == nil || IsNil(o.EngineNo) {
		var ret string
		return ret
	}
	return *o.EngineNo
}

// GetEngineNoOk returns a tuple with the EngineNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AliosOpenAutoInfoQueryResponseModel) GetEngineNoOk() (*string, bool) {
	if o == nil || IsNil(o.EngineNo) {
		return nil, false
	}
	return o.EngineNo, true
}

// HasEngineNo returns a boolean if a field has been set.
func (o *AliosOpenAutoInfoQueryResponseModel) HasEngineNo() bool {
	if o != nil && !IsNil(o.EngineNo) {
		return true
	}

	return false
}

// SetEngineNo gets a reference to the given string and assigns it to the EngineNo field.
func (o *AliosOpenAutoInfoQueryResponseModel) SetEngineNo(v string) {
	o.EngineNo = &v
}

// GetLicenseNo returns the LicenseNo field value if set, zero value otherwise.
func (o *AliosOpenAutoInfoQueryResponseModel) GetLicenseNo() string {
	if o == nil || IsNil(o.LicenseNo) {
		var ret string
		return ret
	}
	return *o.LicenseNo
}

// GetLicenseNoOk returns a tuple with the LicenseNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AliosOpenAutoInfoQueryResponseModel) GetLicenseNoOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseNo) {
		return nil, false
	}
	return o.LicenseNo, true
}

// HasLicenseNo returns a boolean if a field has been set.
func (o *AliosOpenAutoInfoQueryResponseModel) HasLicenseNo() bool {
	if o != nil && !IsNil(o.LicenseNo) {
		return true
	}

	return false
}

// SetLicenseNo gets a reference to the given string and assigns it to the LicenseNo field.
func (o *AliosOpenAutoInfoQueryResponseModel) SetLicenseNo(v string) {
	o.LicenseNo = &v
}

// GetUserActivationTime returns the UserActivationTime field value if set, zero value otherwise.
func (o *AliosOpenAutoInfoQueryResponseModel) GetUserActivationTime() string {
	if o == nil || IsNil(o.UserActivationTime) {
		var ret string
		return ret
	}
	return *o.UserActivationTime
}

// GetUserActivationTimeOk returns a tuple with the UserActivationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AliosOpenAutoInfoQueryResponseModel) GetUserActivationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UserActivationTime) {
		return nil, false
	}
	return o.UserActivationTime, true
}

// HasUserActivationTime returns a boolean if a field has been set.
func (o *AliosOpenAutoInfoQueryResponseModel) HasUserActivationTime() bool {
	if o != nil && !IsNil(o.UserActivationTime) {
		return true
	}

	return false
}

// SetUserActivationTime gets a reference to the given string and assigns it to the UserActivationTime field.
func (o *AliosOpenAutoInfoQueryResponseModel) SetUserActivationTime(v string) {
	o.UserActivationTime = &v
}

// GetVehicleType returns the VehicleType field value if set, zero value otherwise.
func (o *AliosOpenAutoInfoQueryResponseModel) GetVehicleType() string {
	if o == nil || IsNil(o.VehicleType) {
		var ret string
		return ret
	}
	return *o.VehicleType
}

// GetVehicleTypeOk returns a tuple with the VehicleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AliosOpenAutoInfoQueryResponseModel) GetVehicleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VehicleType) {
		return nil, false
	}
	return o.VehicleType, true
}

// HasVehicleType returns a boolean if a field has been set.
func (o *AliosOpenAutoInfoQueryResponseModel) HasVehicleType() bool {
	if o != nil && !IsNil(o.VehicleType) {
		return true
	}

	return false
}

// SetVehicleType gets a reference to the given string and assigns it to the VehicleType field.
func (o *AliosOpenAutoInfoQueryResponseModel) SetVehicleType(v string) {
	o.VehicleType = &v
}

// GetVin returns the Vin field value if set, zero value otherwise.
func (o *AliosOpenAutoInfoQueryResponseModel) GetVin() string {
	if o == nil || IsNil(o.Vin) {
		var ret string
		return ret
	}
	return *o.Vin
}

// GetVinOk returns a tuple with the Vin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AliosOpenAutoInfoQueryResponseModel) GetVinOk() (*string, bool) {
	if o == nil || IsNil(o.Vin) {
		return nil, false
	}
	return o.Vin, true
}

// HasVin returns a boolean if a field has been set.
func (o *AliosOpenAutoInfoQueryResponseModel) HasVin() bool {
	if o != nil && !IsNil(o.Vin) {
		return true
	}

	return false
}

// SetVin gets a reference to the given string and assigns it to the Vin field.
func (o *AliosOpenAutoInfoQueryResponseModel) SetVin(v string) {
	o.Vin = &v
}

func (o AliosOpenAutoInfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AliosOpenAutoInfoQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineNo) {
		toSerialize["engine_no"] = o.EngineNo
	}
	if !IsNil(o.LicenseNo) {
		toSerialize["license_no"] = o.LicenseNo
	}
	if !IsNil(o.UserActivationTime) {
		toSerialize["user_activation_time"] = o.UserActivationTime
	}
	if !IsNil(o.VehicleType) {
		toSerialize["vehicle_type"] = o.VehicleType
	}
	if !IsNil(o.Vin) {
		toSerialize["vin"] = o.Vin
	}
	return toSerialize, nil
}

type NullableAliosOpenAutoInfoQueryResponseModel struct {
	value *AliosOpenAutoInfoQueryResponseModel
	isSet bool
}

func (v NullableAliosOpenAutoInfoQueryResponseModel) Get() *AliosOpenAutoInfoQueryResponseModel {
	return v.value
}

func (v *NullableAliosOpenAutoInfoQueryResponseModel) Set(val *AliosOpenAutoInfoQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAliosOpenAutoInfoQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAliosOpenAutoInfoQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAliosOpenAutoInfoQueryResponseModel(val *AliosOpenAutoInfoQueryResponseModel) *NullableAliosOpenAutoInfoQueryResponseModel {
	return &NullableAliosOpenAutoInfoQueryResponseModel{value: val, isSet: true}
}

func (v NullableAliosOpenAutoInfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAliosOpenAutoInfoQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
