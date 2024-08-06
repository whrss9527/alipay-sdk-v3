/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AlipayCommerceTransportParkingEnterinfoSyncModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceTransportParkingEnterinfoSyncModel{}

// AlipayCommerceTransportParkingEnterinfoSyncModel struct for AlipayCommerceTransportParkingEnterinfoSyncModel
type AlipayCommerceTransportParkingEnterinfoSyncModel struct {
	// 是否启用车牌代扣状态查询功能，true为启用，false为停用
	AgreementQuery *bool `json:"agreement_query,omitempty"`
	// 当前停车场的入场免费时长分钟数
	FreeEnterMinutes *string `json:"free_enter_minutes,omitempty"`
	// 车辆入场的时间，格式\"YYYY-MM-DD HH:mm:ss\"，24小时制，请保证服务器时间准确，入场时间不应晚于当前网络时间
	InTime *string `json:"in_time,omitempty"`
	// 车牌是否加密，true为加密，false为不加密，默认为false
	IsEncryptPlateNo *bool `json:"is_encrypt_plate_no,omitempty"`
	// 当前行程是否需要计费。true：需要，false：不需要。不传默认为true。
	NeedCharge *bool `json:"need_charge,omitempty"`
	// 蚂蚁会员统一ID对应的归属应用appid
	OpenAppid *string `json:"open_appid,omitempty"`
	// 蚂蚁会员统一ID
	OpenId *string `json:"open_id,omitempty"`
	// 外部停车流水号(用于串通进场与出场信息)
	OutSerialNo *string `json:"out_serial_no,omitempty"`
	// 支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create (录入停车场信息)接口获取。
	ParkingId *string `json:"parking_id,omitempty"`
	// 车牌颜色，车牌颜色，枚举支持： *BLUE：蓝。 *GREEN：绿。 *YELLOW：黄。 *WHITE：白。 *BLACK：黑。 *LIMEGREEN：黄绿色。
	PlateColor *string `json:"plate_color,omitempty"`
	// 车牌号（支持加密格式）
	PlateNo *string `json:"plate_no,omitempty"`
	// 停车服务页面地址。 1、服务商停车服务页面地址必须是支付宝小程序URL（无需转换https），详见：https://opendocs.alipay.com/support/01rb18#URL%20%E6%A0%BC%E5%BC%8F  2、若服务商没有服务链接，可传输支付宝停车官方小程序的服务链接：alipays://platformapi/startapp?appId=2021001102642986&page=pages%2Fparking-fee%2Findex  3、若此次对接的是无感支付，则服务链接传输为：alipays://platformapi/startapp?appId=2021001102642986&page=%2Fpages%2Fparking-bill%2Findex
	ServiceUrl *string `json:"service_url,omitempty"`
}

// NewAlipayCommerceTransportParkingEnterinfoSyncModel instantiates a new AlipayCommerceTransportParkingEnterinfoSyncModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceTransportParkingEnterinfoSyncModel() *AlipayCommerceTransportParkingEnterinfoSyncModel {
	this := AlipayCommerceTransportParkingEnterinfoSyncModel{}
	return &this
}

// NewAlipayCommerceTransportParkingEnterinfoSyncModelWithDefaults instantiates a new AlipayCommerceTransportParkingEnterinfoSyncModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceTransportParkingEnterinfoSyncModelWithDefaults() *AlipayCommerceTransportParkingEnterinfoSyncModel {
	this := AlipayCommerceTransportParkingEnterinfoSyncModel{}
	return &this
}

// GetAgreementQuery returns the AgreementQuery field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetAgreementQuery() bool {
	if o == nil || IsNil(o.AgreementQuery) {
		var ret bool
		return ret
	}
	return *o.AgreementQuery
}

// GetAgreementQueryOk returns a tuple with the AgreementQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetAgreementQueryOk() (*bool, bool) {
	if o == nil || IsNil(o.AgreementQuery) {
		return nil, false
	}
	return o.AgreementQuery, true
}

// HasAgreementQuery returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasAgreementQuery() bool {
	if o != nil && !IsNil(o.AgreementQuery) {
		return true
	}

	return false
}

// SetAgreementQuery gets a reference to the given bool and assigns it to the AgreementQuery field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetAgreementQuery(v bool) {
	o.AgreementQuery = &v
}

// GetFreeEnterMinutes returns the FreeEnterMinutes field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetFreeEnterMinutes() string {
	if o == nil || IsNil(o.FreeEnterMinutes) {
		var ret string
		return ret
	}
	return *o.FreeEnterMinutes
}

// GetFreeEnterMinutesOk returns a tuple with the FreeEnterMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetFreeEnterMinutesOk() (*string, bool) {
	if o == nil || IsNil(o.FreeEnterMinutes) {
		return nil, false
	}
	return o.FreeEnterMinutes, true
}

// HasFreeEnterMinutes returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasFreeEnterMinutes() bool {
	if o != nil && !IsNil(o.FreeEnterMinutes) {
		return true
	}

	return false
}

// SetFreeEnterMinutes gets a reference to the given string and assigns it to the FreeEnterMinutes field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetFreeEnterMinutes(v string) {
	o.FreeEnterMinutes = &v
}

// GetInTime returns the InTime field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetInTime() string {
	if o == nil || IsNil(o.InTime) {
		var ret string
		return ret
	}
	return *o.InTime
}

// GetInTimeOk returns a tuple with the InTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetInTimeOk() (*string, bool) {
	if o == nil || IsNil(o.InTime) {
		return nil, false
	}
	return o.InTime, true
}

// HasInTime returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasInTime() bool {
	if o != nil && !IsNil(o.InTime) {
		return true
	}

	return false
}

// SetInTime gets a reference to the given string and assigns it to the InTime field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetInTime(v string) {
	o.InTime = &v
}

// GetIsEncryptPlateNo returns the IsEncryptPlateNo field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetIsEncryptPlateNo() bool {
	if o == nil || IsNil(o.IsEncryptPlateNo) {
		var ret bool
		return ret
	}
	return *o.IsEncryptPlateNo
}

// GetIsEncryptPlateNoOk returns a tuple with the IsEncryptPlateNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetIsEncryptPlateNoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEncryptPlateNo) {
		return nil, false
	}
	return o.IsEncryptPlateNo, true
}

// HasIsEncryptPlateNo returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasIsEncryptPlateNo() bool {
	if o != nil && !IsNil(o.IsEncryptPlateNo) {
		return true
	}

	return false
}

// SetIsEncryptPlateNo gets a reference to the given bool and assigns it to the IsEncryptPlateNo field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetIsEncryptPlateNo(v bool) {
	o.IsEncryptPlateNo = &v
}

// GetNeedCharge returns the NeedCharge field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetNeedCharge() bool {
	if o == nil || IsNil(o.NeedCharge) {
		var ret bool
		return ret
	}
	return *o.NeedCharge
}

// GetNeedChargeOk returns a tuple with the NeedCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetNeedChargeOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedCharge) {
		return nil, false
	}
	return o.NeedCharge, true
}

// HasNeedCharge returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasNeedCharge() bool {
	if o != nil && !IsNil(o.NeedCharge) {
		return true
	}

	return false
}

// SetNeedCharge gets a reference to the given bool and assigns it to the NeedCharge field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetNeedCharge(v bool) {
	o.NeedCharge = &v
}

// GetOpenAppid returns the OpenAppid field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetOpenAppid() string {
	if o == nil || IsNil(o.OpenAppid) {
		var ret string
		return ret
	}
	return *o.OpenAppid
}

// GetOpenAppidOk returns a tuple with the OpenAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetOpenAppidOk() (*string, bool) {
	if o == nil || IsNil(o.OpenAppid) {
		return nil, false
	}
	return o.OpenAppid, true
}

// HasOpenAppid returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasOpenAppid() bool {
	if o != nil && !IsNil(o.OpenAppid) {
		return true
	}

	return false
}

// SetOpenAppid gets a reference to the given string and assigns it to the OpenAppid field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetOpenAppid(v string) {
	o.OpenAppid = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOutSerialNo returns the OutSerialNo field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetOutSerialNo() string {
	if o == nil || IsNil(o.OutSerialNo) {
		var ret string
		return ret
	}
	return *o.OutSerialNo
}

// GetOutSerialNoOk returns a tuple with the OutSerialNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetOutSerialNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutSerialNo) {
		return nil, false
	}
	return o.OutSerialNo, true
}

// HasOutSerialNo returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasOutSerialNo() bool {
	if o != nil && !IsNil(o.OutSerialNo) {
		return true
	}

	return false
}

// SetOutSerialNo gets a reference to the given string and assigns it to the OutSerialNo field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetOutSerialNo(v string) {
	o.OutSerialNo = &v
}

// GetParkingId returns the ParkingId field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetParkingId() string {
	if o == nil || IsNil(o.ParkingId) {
		var ret string
		return ret
	}
	return *o.ParkingId
}

// GetParkingIdOk returns a tuple with the ParkingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetParkingIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingId) {
		return nil, false
	}
	return o.ParkingId, true
}

// HasParkingId returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasParkingId() bool {
	if o != nil && !IsNil(o.ParkingId) {
		return true
	}

	return false
}

// SetParkingId gets a reference to the given string and assigns it to the ParkingId field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetParkingId(v string) {
	o.ParkingId = &v
}

// GetPlateColor returns the PlateColor field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetPlateColor() string {
	if o == nil || IsNil(o.PlateColor) {
		var ret string
		return ret
	}
	return *o.PlateColor
}

// GetPlateColorOk returns a tuple with the PlateColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetPlateColorOk() (*string, bool) {
	if o == nil || IsNil(o.PlateColor) {
		return nil, false
	}
	return o.PlateColor, true
}

// HasPlateColor returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasPlateColor() bool {
	if o != nil && !IsNil(o.PlateColor) {
		return true
	}

	return false
}

// SetPlateColor gets a reference to the given string and assigns it to the PlateColor field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetPlateColor(v string) {
	o.PlateColor = &v
}

// GetPlateNo returns the PlateNo field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetPlateNo() string {
	if o == nil || IsNil(o.PlateNo) {
		var ret string
		return ret
	}
	return *o.PlateNo
}

// GetPlateNoOk returns a tuple with the PlateNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetPlateNoOk() (*string, bool) {
	if o == nil || IsNil(o.PlateNo) {
		return nil, false
	}
	return o.PlateNo, true
}

// HasPlateNo returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasPlateNo() bool {
	if o != nil && !IsNil(o.PlateNo) {
		return true
	}

	return false
}

// SetPlateNo gets a reference to the given string and assigns it to the PlateNo field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetPlateNo(v string) {
	o.PlateNo = &v
}

// GetServiceUrl returns the ServiceUrl field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetServiceUrl() string {
	if o == nil || IsNil(o.ServiceUrl) {
		var ret string
		return ret
	}
	return *o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) GetServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUrl) {
		return nil, false
	}
	return o.ServiceUrl, true
}

// HasServiceUrl returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) HasServiceUrl() bool {
	if o != nil && !IsNil(o.ServiceUrl) {
		return true
	}

	return false
}

// SetServiceUrl gets a reference to the given string and assigns it to the ServiceUrl field.
func (o *AlipayCommerceTransportParkingEnterinfoSyncModel) SetServiceUrl(v string) {
	o.ServiceUrl = &v
}

func (o AlipayCommerceTransportParkingEnterinfoSyncModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceTransportParkingEnterinfoSyncModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementQuery) {
		toSerialize["agreement_query"] = o.AgreementQuery
	}
	if !IsNil(o.FreeEnterMinutes) {
		toSerialize["free_enter_minutes"] = o.FreeEnterMinutes
	}
	if !IsNil(o.InTime) {
		toSerialize["in_time"] = o.InTime
	}
	if !IsNil(o.IsEncryptPlateNo) {
		toSerialize["is_encrypt_plate_no"] = o.IsEncryptPlateNo
	}
	if !IsNil(o.NeedCharge) {
		toSerialize["need_charge"] = o.NeedCharge
	}
	if !IsNil(o.OpenAppid) {
		toSerialize["open_appid"] = o.OpenAppid
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OutSerialNo) {
		toSerialize["out_serial_no"] = o.OutSerialNo
	}
	if !IsNil(o.ParkingId) {
		toSerialize["parking_id"] = o.ParkingId
	}
	if !IsNil(o.PlateColor) {
		toSerialize["plate_color"] = o.PlateColor
	}
	if !IsNil(o.PlateNo) {
		toSerialize["plate_no"] = o.PlateNo
	}
	if !IsNil(o.ServiceUrl) {
		toSerialize["service_url"] = o.ServiceUrl
	}
	return toSerialize, nil
}

type NullableAlipayCommerceTransportParkingEnterinfoSyncModel struct {
	value *AlipayCommerceTransportParkingEnterinfoSyncModel
	isSet bool
}

func (v NullableAlipayCommerceTransportParkingEnterinfoSyncModel) Get() *AlipayCommerceTransportParkingEnterinfoSyncModel {
	return v.value
}

func (v *NullableAlipayCommerceTransportParkingEnterinfoSyncModel) Set(val *AlipayCommerceTransportParkingEnterinfoSyncModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceTransportParkingEnterinfoSyncModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceTransportParkingEnterinfoSyncModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceTransportParkingEnterinfoSyncModel(val *AlipayCommerceTransportParkingEnterinfoSyncModel) *NullableAlipayCommerceTransportParkingEnterinfoSyncModel {
	return &NullableAlipayCommerceTransportParkingEnterinfoSyncModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceTransportParkingEnterinfoSyncModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceTransportParkingEnterinfoSyncModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

