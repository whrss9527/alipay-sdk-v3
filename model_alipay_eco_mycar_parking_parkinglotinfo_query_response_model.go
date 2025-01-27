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

// checks if the AlipayEcoMycarParkingParkinglotinfoQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoMycarParkingParkinglotinfoQueryResponseModel{}

// AlipayEcoMycarParkingParkinglotinfoQueryResponseModel struct for AlipayEcoMycarParkingParkinglotinfoQueryResponseModel
type AlipayEcoMycarParkingParkinglotinfoQueryResponseModel struct {
	// 地区编码
	AddressId *string `json:"address_id,omitempty"`
	// 服务商ID（2088开头的16位纯数字），由服务商提供给ISV
	AgentId *string `json:"agent_id,omitempty"`
	// 车场业务归属列表
	BusinessIsv []BusinessItem `json:"business_isv,omitempty"`
	// 城市编码
	CityId *string `json:"city_id,omitempty"`
	// 高德检索地址
	MapPoiAddress *string `json:"map_poi_address,omitempty"`
	// 高德检索名称
	MapPoiName *string `json:"map_poi_name,omitempty"`
	// 收款方ID（2088开头的16位纯数字），由停车场收款的业主方提供给ISV，该字段暂用于机具和物料申领
	MchntId *string `json:"mchnt_id,omitempty"`
	// ISV停车场ID，由ISV提供，同一个isv或商户范围内唯一
	OutParkingId *string `json:"out_parking_id,omitempty"`
	// 停车场地址
	ParkingAddress *string `json:"parking_address,omitempty"`
	// 收费说明
	ParkingFeeDescription *string `json:"parking_fee_description,omitempty"`
	// 停车场价格明细图片
	ParkingFeeDescriptionImg *string `json:"parking_fee_description_img,omitempty"`
	// 支付宝返回停车场id，系统唯一
	ParkingId *string `json:"parking_id,omitempty"`
	// 停车场位置经度（优先高德坐标获取结果）
	ParkingLatitude *string `json:"parking_latitude,omitempty"`
	// 停车场位置经度（优先高德坐标获取结果）
	ParkingLongitude *string `json:"parking_longitude,omitempty"`
	// 停车场类型，1为居民小区、2为商圈停车场（购物中心商业广场商场等）、3为路侧停车、4为公园景点（景点乐园公园老街古镇等）、5为商务楼宇（酒店写字楼商务楼园区等）、6为其他、7为交通枢纽（机场火车站汽车站码头港口等）、8为市政设施（体育场博物图书馆医院学校等）
	ParkingLotType *string `json:"parking_lot_type,omitempty"`
	// 停车场客服电话
	ParkingMobile *string `json:"parking_mobile,omitempty"`
	// 停车场名称，由ISV定义，尽量与高德地图上的一致
	ParkingName *string `json:"parking_name,omitempty"`
	// 高德地图唯一标识
	ParkingPoiid *string `json:"parking_poiid,omitempty"`
	// 支付方式（1为支付宝在线缴费，2为支付宝代扣缴费，3当面付)，如支持多种方式以','进行间隔
	PayType *string `json:"pay_type,omitempty"`
	// 省份编码
	ProvinceId *string `json:"province_id,omitempty"`
	// 商圈id
	ShopingmallId *string `json:"shopingmall_id,omitempty"`
	// 用户支付未离场的超时时间(以分钟为单位)
	TimeOut *string `json:"time_out,omitempty"`
}

// NewAlipayEcoMycarParkingParkinglotinfoQueryResponseModel instantiates a new AlipayEcoMycarParkingParkinglotinfoQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoMycarParkingParkinglotinfoQueryResponseModel() *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel {
	this := AlipayEcoMycarParkingParkinglotinfoQueryResponseModel{}
	return &this
}

// NewAlipayEcoMycarParkingParkinglotinfoQueryResponseModelWithDefaults instantiates a new AlipayEcoMycarParkingParkinglotinfoQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoMycarParkingParkinglotinfoQueryResponseModelWithDefaults() *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel {
	this := AlipayEcoMycarParkingParkinglotinfoQueryResponseModel{}
	return &this
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetAddressId() string {
	if o == nil || IsNil(o.AddressId) {
		var ret string
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.AddressId) {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasAddressId() bool {
	if o != nil && !IsNil(o.AddressId) {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given string and assigns it to the AddressId field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetAddressId(v string) {
	o.AddressId = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetAgentId() string {
	if o == nil || IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetAgentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasAgentId() bool {
	if o != nil && !IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetAgentId(v string) {
	o.AgentId = &v
}

// GetBusinessIsv returns the BusinessIsv field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetBusinessIsv() []BusinessItem {
	if o == nil || IsNil(o.BusinessIsv) {
		var ret []BusinessItem
		return ret
	}
	return o.BusinessIsv
}

// GetBusinessIsvOk returns a tuple with the BusinessIsv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetBusinessIsvOk() ([]BusinessItem, bool) {
	if o == nil || IsNil(o.BusinessIsv) {
		return nil, false
	}
	return o.BusinessIsv, true
}

// HasBusinessIsv returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasBusinessIsv() bool {
	if o != nil && !IsNil(o.BusinessIsv) {
		return true
	}

	return false
}

// SetBusinessIsv gets a reference to the given []BusinessItem and assigns it to the BusinessIsv field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetBusinessIsv(v []BusinessItem) {
	o.BusinessIsv = v
}

// GetCityId returns the CityId field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetCityId() string {
	if o == nil || IsNil(o.CityId) {
		var ret string
		return ret
	}
	return *o.CityId
}

// GetCityIdOk returns a tuple with the CityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetCityIdOk() (*string, bool) {
	if o == nil || IsNil(o.CityId) {
		return nil, false
	}
	return o.CityId, true
}

// HasCityId returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasCityId() bool {
	if o != nil && !IsNil(o.CityId) {
		return true
	}

	return false
}

// SetCityId gets a reference to the given string and assigns it to the CityId field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetCityId(v string) {
	o.CityId = &v
}

// GetMapPoiAddress returns the MapPoiAddress field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetMapPoiAddress() string {
	if o == nil || IsNil(o.MapPoiAddress) {
		var ret string
		return ret
	}
	return *o.MapPoiAddress
}

// GetMapPoiAddressOk returns a tuple with the MapPoiAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetMapPoiAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MapPoiAddress) {
		return nil, false
	}
	return o.MapPoiAddress, true
}

// HasMapPoiAddress returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasMapPoiAddress() bool {
	if o != nil && !IsNil(o.MapPoiAddress) {
		return true
	}

	return false
}

// SetMapPoiAddress gets a reference to the given string and assigns it to the MapPoiAddress field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetMapPoiAddress(v string) {
	o.MapPoiAddress = &v
}

// GetMapPoiName returns the MapPoiName field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetMapPoiName() string {
	if o == nil || IsNil(o.MapPoiName) {
		var ret string
		return ret
	}
	return *o.MapPoiName
}

// GetMapPoiNameOk returns a tuple with the MapPoiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetMapPoiNameOk() (*string, bool) {
	if o == nil || IsNil(o.MapPoiName) {
		return nil, false
	}
	return o.MapPoiName, true
}

// HasMapPoiName returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasMapPoiName() bool {
	if o != nil && !IsNil(o.MapPoiName) {
		return true
	}

	return false
}

// SetMapPoiName gets a reference to the given string and assigns it to the MapPoiName field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetMapPoiName(v string) {
	o.MapPoiName = &v
}

// GetMchntId returns the MchntId field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetMchntId() string {
	if o == nil || IsNil(o.MchntId) {
		var ret string
		return ret
	}
	return *o.MchntId
}

// GetMchntIdOk returns a tuple with the MchntId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetMchntIdOk() (*string, bool) {
	if o == nil || IsNil(o.MchntId) {
		return nil, false
	}
	return o.MchntId, true
}

// HasMchntId returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasMchntId() bool {
	if o != nil && !IsNil(o.MchntId) {
		return true
	}

	return false
}

// SetMchntId gets a reference to the given string and assigns it to the MchntId field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetMchntId(v string) {
	o.MchntId = &v
}

// GetOutParkingId returns the OutParkingId field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetOutParkingId() string {
	if o == nil || IsNil(o.OutParkingId) {
		var ret string
		return ret
	}
	return *o.OutParkingId
}

// GetOutParkingIdOk returns a tuple with the OutParkingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetOutParkingIdOk() (*string, bool) {
	if o == nil || IsNil(o.OutParkingId) {
		return nil, false
	}
	return o.OutParkingId, true
}

// HasOutParkingId returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasOutParkingId() bool {
	if o != nil && !IsNil(o.OutParkingId) {
		return true
	}

	return false
}

// SetOutParkingId gets a reference to the given string and assigns it to the OutParkingId field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetOutParkingId(v string) {
	o.OutParkingId = &v
}

// GetParkingAddress returns the ParkingAddress field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingAddress() string {
	if o == nil || IsNil(o.ParkingAddress) {
		var ret string
		return ret
	}
	return *o.ParkingAddress
}

// GetParkingAddressOk returns a tuple with the ParkingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingAddress) {
		return nil, false
	}
	return o.ParkingAddress, true
}

// HasParkingAddress returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasParkingAddress() bool {
	if o != nil && !IsNil(o.ParkingAddress) {
		return true
	}

	return false
}

// SetParkingAddress gets a reference to the given string and assigns it to the ParkingAddress field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetParkingAddress(v string) {
	o.ParkingAddress = &v
}

// GetParkingFeeDescription returns the ParkingFeeDescription field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingFeeDescription() string {
	if o == nil || IsNil(o.ParkingFeeDescription) {
		var ret string
		return ret
	}
	return *o.ParkingFeeDescription
}

// GetParkingFeeDescriptionOk returns a tuple with the ParkingFeeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingFeeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingFeeDescription) {
		return nil, false
	}
	return o.ParkingFeeDescription, true
}

// HasParkingFeeDescription returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasParkingFeeDescription() bool {
	if o != nil && !IsNil(o.ParkingFeeDescription) {
		return true
	}

	return false
}

// SetParkingFeeDescription gets a reference to the given string and assigns it to the ParkingFeeDescription field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetParkingFeeDescription(v string) {
	o.ParkingFeeDescription = &v
}

// GetParkingFeeDescriptionImg returns the ParkingFeeDescriptionImg field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingFeeDescriptionImg() string {
	if o == nil || IsNil(o.ParkingFeeDescriptionImg) {
		var ret string
		return ret
	}
	return *o.ParkingFeeDescriptionImg
}

// GetParkingFeeDescriptionImgOk returns a tuple with the ParkingFeeDescriptionImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingFeeDescriptionImgOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingFeeDescriptionImg) {
		return nil, false
	}
	return o.ParkingFeeDescriptionImg, true
}

// HasParkingFeeDescriptionImg returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasParkingFeeDescriptionImg() bool {
	if o != nil && !IsNil(o.ParkingFeeDescriptionImg) {
		return true
	}

	return false
}

// SetParkingFeeDescriptionImg gets a reference to the given string and assigns it to the ParkingFeeDescriptionImg field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetParkingFeeDescriptionImg(v string) {
	o.ParkingFeeDescriptionImg = &v
}

// GetParkingId returns the ParkingId field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingId() string {
	if o == nil || IsNil(o.ParkingId) {
		var ret string
		return ret
	}
	return *o.ParkingId
}

// GetParkingIdOk returns a tuple with the ParkingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingId) {
		return nil, false
	}
	return o.ParkingId, true
}

// HasParkingId returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasParkingId() bool {
	if o != nil && !IsNil(o.ParkingId) {
		return true
	}

	return false
}

// SetParkingId gets a reference to the given string and assigns it to the ParkingId field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetParkingId(v string) {
	o.ParkingId = &v
}

// GetParkingLatitude returns the ParkingLatitude field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingLatitude() string {
	if o == nil || IsNil(o.ParkingLatitude) {
		var ret string
		return ret
	}
	return *o.ParkingLatitude
}

// GetParkingLatitudeOk returns a tuple with the ParkingLatitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingLatitude) {
		return nil, false
	}
	return o.ParkingLatitude, true
}

// HasParkingLatitude returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasParkingLatitude() bool {
	if o != nil && !IsNil(o.ParkingLatitude) {
		return true
	}

	return false
}

// SetParkingLatitude gets a reference to the given string and assigns it to the ParkingLatitude field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetParkingLatitude(v string) {
	o.ParkingLatitude = &v
}

// GetParkingLongitude returns the ParkingLongitude field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingLongitude() string {
	if o == nil || IsNil(o.ParkingLongitude) {
		var ret string
		return ret
	}
	return *o.ParkingLongitude
}

// GetParkingLongitudeOk returns a tuple with the ParkingLongitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingLongitudeOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingLongitude) {
		return nil, false
	}
	return o.ParkingLongitude, true
}

// HasParkingLongitude returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasParkingLongitude() bool {
	if o != nil && !IsNil(o.ParkingLongitude) {
		return true
	}

	return false
}

// SetParkingLongitude gets a reference to the given string and assigns it to the ParkingLongitude field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetParkingLongitude(v string) {
	o.ParkingLongitude = &v
}

// GetParkingLotType returns the ParkingLotType field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingLotType() string {
	if o == nil || IsNil(o.ParkingLotType) {
		var ret string
		return ret
	}
	return *o.ParkingLotType
}

// GetParkingLotTypeOk returns a tuple with the ParkingLotType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingLotTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingLotType) {
		return nil, false
	}
	return o.ParkingLotType, true
}

// HasParkingLotType returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasParkingLotType() bool {
	if o != nil && !IsNil(o.ParkingLotType) {
		return true
	}

	return false
}

// SetParkingLotType gets a reference to the given string and assigns it to the ParkingLotType field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetParkingLotType(v string) {
	o.ParkingLotType = &v
}

// GetParkingMobile returns the ParkingMobile field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingMobile() string {
	if o == nil || IsNil(o.ParkingMobile) {
		var ret string
		return ret
	}
	return *o.ParkingMobile
}

// GetParkingMobileOk returns a tuple with the ParkingMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingMobileOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingMobile) {
		return nil, false
	}
	return o.ParkingMobile, true
}

// HasParkingMobile returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasParkingMobile() bool {
	if o != nil && !IsNil(o.ParkingMobile) {
		return true
	}

	return false
}

// SetParkingMobile gets a reference to the given string and assigns it to the ParkingMobile field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetParkingMobile(v string) {
	o.ParkingMobile = &v
}

// GetParkingName returns the ParkingName field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingName() string {
	if o == nil || IsNil(o.ParkingName) {
		var ret string
		return ret
	}
	return *o.ParkingName
}

// GetParkingNameOk returns a tuple with the ParkingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingName) {
		return nil, false
	}
	return o.ParkingName, true
}

// HasParkingName returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasParkingName() bool {
	if o != nil && !IsNil(o.ParkingName) {
		return true
	}

	return false
}

// SetParkingName gets a reference to the given string and assigns it to the ParkingName field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetParkingName(v string) {
	o.ParkingName = &v
}

// GetParkingPoiid returns the ParkingPoiid field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingPoiid() string {
	if o == nil || IsNil(o.ParkingPoiid) {
		var ret string
		return ret
	}
	return *o.ParkingPoiid
}

// GetParkingPoiidOk returns a tuple with the ParkingPoiid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetParkingPoiidOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingPoiid) {
		return nil, false
	}
	return o.ParkingPoiid, true
}

// HasParkingPoiid returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasParkingPoiid() bool {
	if o != nil && !IsNil(o.ParkingPoiid) {
		return true
	}

	return false
}

// SetParkingPoiid gets a reference to the given string and assigns it to the ParkingPoiid field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetParkingPoiid(v string) {
	o.ParkingPoiid = &v
}

// GetPayType returns the PayType field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetPayType() string {
	if o == nil || IsNil(o.PayType) {
		var ret string
		return ret
	}
	return *o.PayType
}

// GetPayTypeOk returns a tuple with the PayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetPayTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PayType) {
		return nil, false
	}
	return o.PayType, true
}

// HasPayType returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasPayType() bool {
	if o != nil && !IsNil(o.PayType) {
		return true
	}

	return false
}

// SetPayType gets a reference to the given string and assigns it to the PayType field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetPayType(v string) {
	o.PayType = &v
}

// GetProvinceId returns the ProvinceId field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetProvinceId() string {
	if o == nil || IsNil(o.ProvinceId) {
		var ret string
		return ret
	}
	return *o.ProvinceId
}

// GetProvinceIdOk returns a tuple with the ProvinceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetProvinceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProvinceId) {
		return nil, false
	}
	return o.ProvinceId, true
}

// HasProvinceId returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasProvinceId() bool {
	if o != nil && !IsNil(o.ProvinceId) {
		return true
	}

	return false
}

// SetProvinceId gets a reference to the given string and assigns it to the ProvinceId field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetProvinceId(v string) {
	o.ProvinceId = &v
}

// GetShopingmallId returns the ShopingmallId field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetShopingmallId() string {
	if o == nil || IsNil(o.ShopingmallId) {
		var ret string
		return ret
	}
	return *o.ShopingmallId
}

// GetShopingmallIdOk returns a tuple with the ShopingmallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetShopingmallIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShopingmallId) {
		return nil, false
	}
	return o.ShopingmallId, true
}

// HasShopingmallId returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasShopingmallId() bool {
	if o != nil && !IsNil(o.ShopingmallId) {
		return true
	}

	return false
}

// SetShopingmallId gets a reference to the given string and assigns it to the ShopingmallId field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetShopingmallId(v string) {
	o.ShopingmallId = &v
}

// GetTimeOut returns the TimeOut field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetTimeOut() string {
	if o == nil || IsNil(o.TimeOut) {
		var ret string
		return ret
	}
	return *o.TimeOut
}

// GetTimeOutOk returns a tuple with the TimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) GetTimeOutOk() (*string, bool) {
	if o == nil || IsNil(o.TimeOut) {
		return nil, false
	}
	return o.TimeOut, true
}

// HasTimeOut returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) HasTimeOut() bool {
	if o != nil && !IsNil(o.TimeOut) {
		return true
	}

	return false
}

// SetTimeOut gets a reference to the given string and assigns it to the TimeOut field.
func (o *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) SetTimeOut(v string) {
	o.TimeOut = &v
}

func (o AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressId) {
		toSerialize["address_id"] = o.AddressId
	}
	if !IsNil(o.AgentId) {
		toSerialize["agent_id"] = o.AgentId
	}
	if !IsNil(o.BusinessIsv) {
		toSerialize["business_isv"] = o.BusinessIsv
	}
	if !IsNil(o.CityId) {
		toSerialize["city_id"] = o.CityId
	}
	if !IsNil(o.MapPoiAddress) {
		toSerialize["map_poi_address"] = o.MapPoiAddress
	}
	if !IsNil(o.MapPoiName) {
		toSerialize["map_poi_name"] = o.MapPoiName
	}
	if !IsNil(o.MchntId) {
		toSerialize["mchnt_id"] = o.MchntId
	}
	if !IsNil(o.OutParkingId) {
		toSerialize["out_parking_id"] = o.OutParkingId
	}
	if !IsNil(o.ParkingAddress) {
		toSerialize["parking_address"] = o.ParkingAddress
	}
	if !IsNil(o.ParkingFeeDescription) {
		toSerialize["parking_fee_description"] = o.ParkingFeeDescription
	}
	if !IsNil(o.ParkingFeeDescriptionImg) {
		toSerialize["parking_fee_description_img"] = o.ParkingFeeDescriptionImg
	}
	if !IsNil(o.ParkingId) {
		toSerialize["parking_id"] = o.ParkingId
	}
	if !IsNil(o.ParkingLatitude) {
		toSerialize["parking_latitude"] = o.ParkingLatitude
	}
	if !IsNil(o.ParkingLongitude) {
		toSerialize["parking_longitude"] = o.ParkingLongitude
	}
	if !IsNil(o.ParkingLotType) {
		toSerialize["parking_lot_type"] = o.ParkingLotType
	}
	if !IsNil(o.ParkingMobile) {
		toSerialize["parking_mobile"] = o.ParkingMobile
	}
	if !IsNil(o.ParkingName) {
		toSerialize["parking_name"] = o.ParkingName
	}
	if !IsNil(o.ParkingPoiid) {
		toSerialize["parking_poiid"] = o.ParkingPoiid
	}
	if !IsNil(o.PayType) {
		toSerialize["pay_type"] = o.PayType
	}
	if !IsNil(o.ProvinceId) {
		toSerialize["province_id"] = o.ProvinceId
	}
	if !IsNil(o.ShopingmallId) {
		toSerialize["shopingmall_id"] = o.ShopingmallId
	}
	if !IsNil(o.TimeOut) {
		toSerialize["time_out"] = o.TimeOut
	}
	return toSerialize, nil
}

type NullableAlipayEcoMycarParkingParkinglotinfoQueryResponseModel struct {
	value *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel
	isSet bool
}

func (v NullableAlipayEcoMycarParkingParkinglotinfoQueryResponseModel) Get() *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingParkinglotinfoQueryResponseModel) Set(val *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingParkinglotinfoQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingParkinglotinfoQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingParkinglotinfoQueryResponseModel(val *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel) *NullableAlipayEcoMycarParkingParkinglotinfoQueryResponseModel {
	return &NullableAlipayEcoMycarParkingParkinglotinfoQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingParkinglotinfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingParkinglotinfoQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
