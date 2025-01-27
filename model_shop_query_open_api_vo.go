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

// checks if the ShopQueryOpenApiVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShopQueryOpenApiVO{}

// ShopQueryOpenApiVO struct for ShopQueryOpenApiVO
type ShopQueryOpenApiVO struct {
	BusinessAddress *AddressInfo `json:"business_address,omitempty"`
	// 店铺经营时间
	BusinessTime []ShopBusinessTime `json:"business_time,omitempty"`
	// 店铺联系手机
	ContactMobile *string `json:"contact_mobile,omitempty"`
	// 店铺的联系固话
	ContactPhone *string `json:"contact_phone,omitempty"`
	// 新版门店类目标准二级类目code。类目标准及与原shop_category映射关系参见文档https://gw.alipayobjects.com/os/bmw-prod/4b3f82df-e53e-4b84-bc41-fe025101e726.xlsx
	NewShopCategory *string `json:"new_shop_category,omitempty"`
	// 店铺类目.取值参见文件<a href=\"https://mif-pub.alipayobjects.com/ShopCategory.xlsx\">文件</a>中的三级门店类目
	ShopCategory *string `json:"shop_category,omitempty"`
	// 蚂蚁店铺id
	ShopId *string `json:"shop_id,omitempty"`
	// 当前名称、地址、经纬度信息准确一致，可用于数字化经营场景消费(如商品、券、消费圈等场域的分发)，不影响门店支付结算
	ShopInfoStatus *string `json:"shop_info_status,omitempty"`
	// 店铺名称
	ShopName *string `json:"shop_name,omitempty"`
	// 门店状态
	ShopStatus *string `json:"shop_status,omitempty"`
	// 店铺经营类型
	ShopType *string `json:"shop_type,omitempty"`
	// 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号
	StoreId *string `json:"store_id,omitempty"`
}

// NewShopQueryOpenApiVO instantiates a new ShopQueryOpenApiVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopQueryOpenApiVO() *ShopQueryOpenApiVO {
	this := ShopQueryOpenApiVO{}
	return &this
}

// NewShopQueryOpenApiVOWithDefaults instantiates a new ShopQueryOpenApiVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopQueryOpenApiVOWithDefaults() *ShopQueryOpenApiVO {
	this := ShopQueryOpenApiVO{}
	return &this
}

// GetBusinessAddress returns the BusinessAddress field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetBusinessAddress() AddressInfo {
	if o == nil || IsNil(o.BusinessAddress) {
		var ret AddressInfo
		return ret
	}
	return *o.BusinessAddress
}

// GetBusinessAddressOk returns a tuple with the BusinessAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetBusinessAddressOk() (*AddressInfo, bool) {
	if o == nil || IsNil(o.BusinessAddress) {
		return nil, false
	}
	return o.BusinessAddress, true
}

// HasBusinessAddress returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasBusinessAddress() bool {
	if o != nil && !IsNil(o.BusinessAddress) {
		return true
	}

	return false
}

// SetBusinessAddress gets a reference to the given AddressInfo and assigns it to the BusinessAddress field.
func (o *ShopQueryOpenApiVO) SetBusinessAddress(v AddressInfo) {
	o.BusinessAddress = &v
}

// GetBusinessTime returns the BusinessTime field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetBusinessTime() []ShopBusinessTime {
	if o == nil || IsNil(o.BusinessTime) {
		var ret []ShopBusinessTime
		return ret
	}
	return o.BusinessTime
}

// GetBusinessTimeOk returns a tuple with the BusinessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetBusinessTimeOk() ([]ShopBusinessTime, bool) {
	if o == nil || IsNil(o.BusinessTime) {
		return nil, false
	}
	return o.BusinessTime, true
}

// HasBusinessTime returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasBusinessTime() bool {
	if o != nil && !IsNil(o.BusinessTime) {
		return true
	}

	return false
}

// SetBusinessTime gets a reference to the given []ShopBusinessTime and assigns it to the BusinessTime field.
func (o *ShopQueryOpenApiVO) SetBusinessTime(v []ShopBusinessTime) {
	o.BusinessTime = v
}

// GetContactMobile returns the ContactMobile field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetContactMobile() string {
	if o == nil || IsNil(o.ContactMobile) {
		var ret string
		return ret
	}
	return *o.ContactMobile
}

// GetContactMobileOk returns a tuple with the ContactMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetContactMobileOk() (*string, bool) {
	if o == nil || IsNil(o.ContactMobile) {
		return nil, false
	}
	return o.ContactMobile, true
}

// HasContactMobile returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasContactMobile() bool {
	if o != nil && !IsNil(o.ContactMobile) {
		return true
	}

	return false
}

// SetContactMobile gets a reference to the given string and assigns it to the ContactMobile field.
func (o *ShopQueryOpenApiVO) SetContactMobile(v string) {
	o.ContactMobile = &v
}

// GetContactPhone returns the ContactPhone field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetContactPhone() string {
	if o == nil || IsNil(o.ContactPhone) {
		var ret string
		return ret
	}
	return *o.ContactPhone
}

// GetContactPhoneOk returns a tuple with the ContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetContactPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.ContactPhone) {
		return nil, false
	}
	return o.ContactPhone, true
}

// HasContactPhone returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasContactPhone() bool {
	if o != nil && !IsNil(o.ContactPhone) {
		return true
	}

	return false
}

// SetContactPhone gets a reference to the given string and assigns it to the ContactPhone field.
func (o *ShopQueryOpenApiVO) SetContactPhone(v string) {
	o.ContactPhone = &v
}

// GetNewShopCategory returns the NewShopCategory field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetNewShopCategory() string {
	if o == nil || IsNil(o.NewShopCategory) {
		var ret string
		return ret
	}
	return *o.NewShopCategory
}

// GetNewShopCategoryOk returns a tuple with the NewShopCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetNewShopCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.NewShopCategory) {
		return nil, false
	}
	return o.NewShopCategory, true
}

// HasNewShopCategory returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasNewShopCategory() bool {
	if o != nil && !IsNil(o.NewShopCategory) {
		return true
	}

	return false
}

// SetNewShopCategory gets a reference to the given string and assigns it to the NewShopCategory field.
func (o *ShopQueryOpenApiVO) SetNewShopCategory(v string) {
	o.NewShopCategory = &v
}

// GetShopCategory returns the ShopCategory field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetShopCategory() string {
	if o == nil || IsNil(o.ShopCategory) {
		var ret string
		return ret
	}
	return *o.ShopCategory
}

// GetShopCategoryOk returns a tuple with the ShopCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetShopCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ShopCategory) {
		return nil, false
	}
	return o.ShopCategory, true
}

// HasShopCategory returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasShopCategory() bool {
	if o != nil && !IsNil(o.ShopCategory) {
		return true
	}

	return false
}

// SetShopCategory gets a reference to the given string and assigns it to the ShopCategory field.
func (o *ShopQueryOpenApiVO) SetShopCategory(v string) {
	o.ShopCategory = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetShopId() string {
	if o == nil || IsNil(o.ShopId) {
		var ret string
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given string and assigns it to the ShopId field.
func (o *ShopQueryOpenApiVO) SetShopId(v string) {
	o.ShopId = &v
}

// GetShopInfoStatus returns the ShopInfoStatus field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetShopInfoStatus() string {
	if o == nil || IsNil(o.ShopInfoStatus) {
		var ret string
		return ret
	}
	return *o.ShopInfoStatus
}

// GetShopInfoStatusOk returns a tuple with the ShopInfoStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetShopInfoStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ShopInfoStatus) {
		return nil, false
	}
	return o.ShopInfoStatus, true
}

// HasShopInfoStatus returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasShopInfoStatus() bool {
	if o != nil && !IsNil(o.ShopInfoStatus) {
		return true
	}

	return false
}

// SetShopInfoStatus gets a reference to the given string and assigns it to the ShopInfoStatus field.
func (o *ShopQueryOpenApiVO) SetShopInfoStatus(v string) {
	o.ShopInfoStatus = &v
}

// GetShopName returns the ShopName field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetShopName() string {
	if o == nil || IsNil(o.ShopName) {
		var ret string
		return ret
	}
	return *o.ShopName
}

// GetShopNameOk returns a tuple with the ShopName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetShopNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShopName) {
		return nil, false
	}
	return o.ShopName, true
}

// HasShopName returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasShopName() bool {
	if o != nil && !IsNil(o.ShopName) {
		return true
	}

	return false
}

// SetShopName gets a reference to the given string and assigns it to the ShopName field.
func (o *ShopQueryOpenApiVO) SetShopName(v string) {
	o.ShopName = &v
}

// GetShopStatus returns the ShopStatus field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetShopStatus() string {
	if o == nil || IsNil(o.ShopStatus) {
		var ret string
		return ret
	}
	return *o.ShopStatus
}

// GetShopStatusOk returns a tuple with the ShopStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetShopStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ShopStatus) {
		return nil, false
	}
	return o.ShopStatus, true
}

// HasShopStatus returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasShopStatus() bool {
	if o != nil && !IsNil(o.ShopStatus) {
		return true
	}

	return false
}

// SetShopStatus gets a reference to the given string and assigns it to the ShopStatus field.
func (o *ShopQueryOpenApiVO) SetShopStatus(v string) {
	o.ShopStatus = &v
}

// GetShopType returns the ShopType field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetShopType() string {
	if o == nil || IsNil(o.ShopType) {
		var ret string
		return ret
	}
	return *o.ShopType
}

// GetShopTypeOk returns a tuple with the ShopType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetShopTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ShopType) {
		return nil, false
	}
	return o.ShopType, true
}

// HasShopType returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasShopType() bool {
	if o != nil && !IsNil(o.ShopType) {
		return true
	}

	return false
}

// SetShopType gets a reference to the given string and assigns it to the ShopType field.
func (o *ShopQueryOpenApiVO) SetShopType(v string) {
	o.ShopType = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *ShopQueryOpenApiVO) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopQueryOpenApiVO) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *ShopQueryOpenApiVO) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *ShopQueryOpenApiVO) SetStoreId(v string) {
	o.StoreId = &v
}

func (o ShopQueryOpenApiVO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShopQueryOpenApiVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessAddress) {
		toSerialize["business_address"] = o.BusinessAddress
	}
	if !IsNil(o.BusinessTime) {
		toSerialize["business_time"] = o.BusinessTime
	}
	if !IsNil(o.ContactMobile) {
		toSerialize["contact_mobile"] = o.ContactMobile
	}
	if !IsNil(o.ContactPhone) {
		toSerialize["contact_phone"] = o.ContactPhone
	}
	if !IsNil(o.NewShopCategory) {
		toSerialize["new_shop_category"] = o.NewShopCategory
	}
	if !IsNil(o.ShopCategory) {
		toSerialize["shop_category"] = o.ShopCategory
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.ShopInfoStatus) {
		toSerialize["shop_info_status"] = o.ShopInfoStatus
	}
	if !IsNil(o.ShopName) {
		toSerialize["shop_name"] = o.ShopName
	}
	if !IsNil(o.ShopStatus) {
		toSerialize["shop_status"] = o.ShopStatus
	}
	if !IsNil(o.ShopType) {
		toSerialize["shop_type"] = o.ShopType
	}
	if !IsNil(o.StoreId) {
		toSerialize["store_id"] = o.StoreId
	}
	return toSerialize, nil
}

type NullableShopQueryOpenApiVO struct {
	value *ShopQueryOpenApiVO
	isSet bool
}

func (v NullableShopQueryOpenApiVO) Get() *ShopQueryOpenApiVO {
	return v.value
}

func (v *NullableShopQueryOpenApiVO) Set(val *ShopQueryOpenApiVO) {
	v.value = val
	v.isSet = true
}

func (v NullableShopQueryOpenApiVO) IsSet() bool {
	return v.isSet
}

func (v *NullableShopQueryOpenApiVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopQueryOpenApiVO(val *ShopQueryOpenApiVO) *NullableShopQueryOpenApiVO {
	return &NullableShopQueryOpenApiVO{value: val, isSet: true}
}

func (v NullableShopQueryOpenApiVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShopQueryOpenApiVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
