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

// checks if the DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel{}

// DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel struct for DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel
type DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel struct {
	// 人脸检测流水ID，请保留方便排查问题
	CertifyId *string `json:"certify_id,omitempty"`
	// 检测页面url，需要给到终端SDK(该参数后续不再维护，建议使用该参数的商户进行如下升级： 1.在App端活体检测场景中使用了该参数的商户，新版客户端SDK不再依赖该参数，请参考产品官方文档接入指南中的<a href=\"https://opendocs.alipay.com/open/04pxpy?pathHash=950007fa&ref=api\">App端活体检测</a>部分，进行客户端升级； 2.在H5活体检测场景中使用了该参数的商户，请替换为web_url进行认证。)
	// Deprecated
	PageUrl *string `json:"page_url,omitempty"`
	// H5活体检测地址
	WebUrl *string `json:"web_url,omitempty"`
}

// NewDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel instantiates a new DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel() *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel {
	this := DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel{}
	return &this
}

// NewDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModelWithDefaults instantiates a new DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModelWithDefaults() *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel {
	this := DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel{}
	return &this
}

// GetCertifyId returns the CertifyId field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) GetCertifyId() string {
	if o == nil || IsNil(o.CertifyId) {
		var ret string
		return ret
	}
	return *o.CertifyId
}

// GetCertifyIdOk returns a tuple with the CertifyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) GetCertifyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CertifyId) {
		return nil, false
	}
	return o.CertifyId, true
}

// HasCertifyId returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) HasCertifyId() bool {
	if o != nil && !IsNil(o.CertifyId) {
		return true
	}

	return false
}

// SetCertifyId gets a reference to the given string and assigns it to the CertifyId field.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) SetCertifyId(v string) {
	o.CertifyId = &v
}

// GetPageUrl returns the PageUrl field value if set, zero value otherwise.
// Deprecated
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) GetPageUrl() string {
	if o == nil || IsNil(o.PageUrl) {
		var ret string
		return ret
	}
	return *o.PageUrl
}

// GetPageUrlOk returns a tuple with the PageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) GetPageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PageUrl) {
		return nil, false
	}
	return o.PageUrl, true
}

// HasPageUrl returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) HasPageUrl() bool {
	if o != nil && !IsNil(o.PageUrl) {
		return true
	}

	return false
}

// SetPageUrl gets a reference to the given string and assigns it to the PageUrl field.
// Deprecated
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) SetPageUrl(v string) {
	o.PageUrl = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl) {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) GetWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebUrl) {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) HasWebUrl() bool {
	if o != nil && !IsNil(o.WebUrl) {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) SetWebUrl(v string) {
	o.WebUrl = &v
}

func (o DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertifyId) {
		toSerialize["certify_id"] = o.CertifyId
	}
	if !IsNil(o.PageUrl) {
		toSerialize["page_url"] = o.PageUrl
	}
	if !IsNil(o.WebUrl) {
		toSerialize["web_url"] = o.WebUrl
	}
	return toSerialize, nil
}

type NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel struct {
	value *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel
	isSet bool
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) Get() *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel {
	return v.value
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) Set(val *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel(val *DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) *NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel {
	return &NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel{value: val, isSet: true}
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

