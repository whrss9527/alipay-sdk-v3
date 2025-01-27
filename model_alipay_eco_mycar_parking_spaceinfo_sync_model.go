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

// checks if the AlipayEcoMycarParkingSpaceinfoSyncModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoMycarParkingSpaceinfoSyncModel{}

// AlipayEcoMycarParkingSpaceinfoSyncModel struct for AlipayEcoMycarParkingSpaceinfoSyncModel
type AlipayEcoMycarParkingSpaceinfoSyncModel struct {
	// 空闲充电桩车位数
	FreeChargingPile *int32 `json:"free_charging_pile,omitempty"`
	// 车场空闲可用的车位数
	FreeParkingSpace *int32 `json:"free_parking_space,omitempty"`
	// 停车场ID
	ParkingId *string `json:"parking_id,omitempty"`
	// 停车场实时状态: 0 拥堵,1 正常，2 空闲
	ParkingSpaceStatus *string `json:"parking_space_status,omitempty"`
}

// NewAlipayEcoMycarParkingSpaceinfoSyncModel instantiates a new AlipayEcoMycarParkingSpaceinfoSyncModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoMycarParkingSpaceinfoSyncModel() *AlipayEcoMycarParkingSpaceinfoSyncModel {
	this := AlipayEcoMycarParkingSpaceinfoSyncModel{}
	return &this
}

// NewAlipayEcoMycarParkingSpaceinfoSyncModelWithDefaults instantiates a new AlipayEcoMycarParkingSpaceinfoSyncModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoMycarParkingSpaceinfoSyncModelWithDefaults() *AlipayEcoMycarParkingSpaceinfoSyncModel {
	this := AlipayEcoMycarParkingSpaceinfoSyncModel{}
	return &this
}

// GetFreeChargingPile returns the FreeChargingPile field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) GetFreeChargingPile() int32 {
	if o == nil || IsNil(o.FreeChargingPile) {
		var ret int32
		return ret
	}
	return *o.FreeChargingPile
}

// GetFreeChargingPileOk returns a tuple with the FreeChargingPile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) GetFreeChargingPileOk() (*int32, bool) {
	if o == nil || IsNil(o.FreeChargingPile) {
		return nil, false
	}
	return o.FreeChargingPile, true
}

// HasFreeChargingPile returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) HasFreeChargingPile() bool {
	if o != nil && !IsNil(o.FreeChargingPile) {
		return true
	}

	return false
}

// SetFreeChargingPile gets a reference to the given int32 and assigns it to the FreeChargingPile field.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) SetFreeChargingPile(v int32) {
	o.FreeChargingPile = &v
}

// GetFreeParkingSpace returns the FreeParkingSpace field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) GetFreeParkingSpace() int32 {
	if o == nil || IsNil(o.FreeParkingSpace) {
		var ret int32
		return ret
	}
	return *o.FreeParkingSpace
}

// GetFreeParkingSpaceOk returns a tuple with the FreeParkingSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) GetFreeParkingSpaceOk() (*int32, bool) {
	if o == nil || IsNil(o.FreeParkingSpace) {
		return nil, false
	}
	return o.FreeParkingSpace, true
}

// HasFreeParkingSpace returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) HasFreeParkingSpace() bool {
	if o != nil && !IsNil(o.FreeParkingSpace) {
		return true
	}

	return false
}

// SetFreeParkingSpace gets a reference to the given int32 and assigns it to the FreeParkingSpace field.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) SetFreeParkingSpace(v int32) {
	o.FreeParkingSpace = &v
}

// GetParkingId returns the ParkingId field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) GetParkingId() string {
	if o == nil || IsNil(o.ParkingId) {
		var ret string
		return ret
	}
	return *o.ParkingId
}

// GetParkingIdOk returns a tuple with the ParkingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) GetParkingIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingId) {
		return nil, false
	}
	return o.ParkingId, true
}

// HasParkingId returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) HasParkingId() bool {
	if o != nil && !IsNil(o.ParkingId) {
		return true
	}

	return false
}

// SetParkingId gets a reference to the given string and assigns it to the ParkingId field.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) SetParkingId(v string) {
	o.ParkingId = &v
}

// GetParkingSpaceStatus returns the ParkingSpaceStatus field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) GetParkingSpaceStatus() string {
	if o == nil || IsNil(o.ParkingSpaceStatus) {
		var ret string
		return ret
	}
	return *o.ParkingSpaceStatus
}

// GetParkingSpaceStatusOk returns a tuple with the ParkingSpaceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) GetParkingSpaceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingSpaceStatus) {
		return nil, false
	}
	return o.ParkingSpaceStatus, true
}

// HasParkingSpaceStatus returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) HasParkingSpaceStatus() bool {
	if o != nil && !IsNil(o.ParkingSpaceStatus) {
		return true
	}

	return false
}

// SetParkingSpaceStatus gets a reference to the given string and assigns it to the ParkingSpaceStatus field.
func (o *AlipayEcoMycarParkingSpaceinfoSyncModel) SetParkingSpaceStatus(v string) {
	o.ParkingSpaceStatus = &v
}

func (o AlipayEcoMycarParkingSpaceinfoSyncModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoMycarParkingSpaceinfoSyncModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FreeChargingPile) {
		toSerialize["free_charging_pile"] = o.FreeChargingPile
	}
	if !IsNil(o.FreeParkingSpace) {
		toSerialize["free_parking_space"] = o.FreeParkingSpace
	}
	if !IsNil(o.ParkingId) {
		toSerialize["parking_id"] = o.ParkingId
	}
	if !IsNil(o.ParkingSpaceStatus) {
		toSerialize["parking_space_status"] = o.ParkingSpaceStatus
	}
	return toSerialize, nil
}

type NullableAlipayEcoMycarParkingSpaceinfoSyncModel struct {
	value *AlipayEcoMycarParkingSpaceinfoSyncModel
	isSet bool
}

func (v NullableAlipayEcoMycarParkingSpaceinfoSyncModel) Get() *AlipayEcoMycarParkingSpaceinfoSyncModel {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingSpaceinfoSyncModel) Set(val *AlipayEcoMycarParkingSpaceinfoSyncModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingSpaceinfoSyncModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingSpaceinfoSyncModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingSpaceinfoSyncModel(val *AlipayEcoMycarParkingSpaceinfoSyncModel) *NullableAlipayEcoMycarParkingSpaceinfoSyncModel {
	return &NullableAlipayEcoMycarParkingSpaceinfoSyncModel{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingSpaceinfoSyncModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingSpaceinfoSyncModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
