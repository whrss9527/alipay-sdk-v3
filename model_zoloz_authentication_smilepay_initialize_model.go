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

// checks if the ZolozAuthenticationSmilepayInitializeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZolozAuthenticationSmilepayInitializeModel{}

// ZolozAuthenticationSmilepayInitializeModel struct for ZolozAuthenticationSmilepayInitializeModel
type ZolozAuthenticationSmilepayInitializeModel struct {
	// 设备指纹，用于唯一标识一台设备
	ApdidToken *string `json:"apdid_token,omitempty"`
	// 人脸识别应用名称
	AppName *string `json:"app_name,omitempty"`
	// 人脸识别应用版本号
	AppVersion *string `json:"app_version,omitempty"`
	// 基础包版本号
	BaseVer *string `json:"base_ver,omitempty"`
	// 生物识别元信息
	BioMetaInfo *string `json:"bio_meta_info,omitempty"`
	// 设备型号
	DeviceModel *string `json:"device_model,omitempty"`
	// 设备类型
	DeviceType   *string           `json:"device_type,omitempty"`
	ExtInfo      *FaceExtParams    `json:"ext_info,omitempty"`
	MachineInfo  *FaceMachineInfo  `json:"machine_info,omitempty"`
	MerchantInfo *FaceMerchantInfo `json:"merchant_info,omitempty"`
	// 操作系统版本号
	OsVersion *string `json:"os_version,omitempty"`
	// 业务ID
	RemoteLogId *string `json:"remote_log_id,omitempty"`
	// ZIM版本号
	ZimVer *string `json:"zim_ver,omitempty"`
}

// NewZolozAuthenticationSmilepayInitializeModel instantiates a new ZolozAuthenticationSmilepayInitializeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZolozAuthenticationSmilepayInitializeModel() *ZolozAuthenticationSmilepayInitializeModel {
	this := ZolozAuthenticationSmilepayInitializeModel{}
	return &this
}

// NewZolozAuthenticationSmilepayInitializeModelWithDefaults instantiates a new ZolozAuthenticationSmilepayInitializeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZolozAuthenticationSmilepayInitializeModelWithDefaults() *ZolozAuthenticationSmilepayInitializeModel {
	this := ZolozAuthenticationSmilepayInitializeModel{}
	return &this
}

// GetApdidToken returns the ApdidToken field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetApdidToken() string {
	if o == nil || IsNil(o.ApdidToken) {
		var ret string
		return ret
	}
	return *o.ApdidToken
}

// GetApdidTokenOk returns a tuple with the ApdidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetApdidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ApdidToken) {
		return nil, false
	}
	return o.ApdidToken, true
}

// HasApdidToken returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasApdidToken() bool {
	if o != nil && !IsNil(o.ApdidToken) {
		return true
	}

	return false
}

// SetApdidToken gets a reference to the given string and assigns it to the ApdidToken field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetApdidToken(v string) {
	o.ApdidToken = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetAppName(v string) {
	o.AppName = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBaseVer returns the BaseVer field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetBaseVer() string {
	if o == nil || IsNil(o.BaseVer) {
		var ret string
		return ret
	}
	return *o.BaseVer
}

// GetBaseVerOk returns a tuple with the BaseVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetBaseVerOk() (*string, bool) {
	if o == nil || IsNil(o.BaseVer) {
		return nil, false
	}
	return o.BaseVer, true
}

// HasBaseVer returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasBaseVer() bool {
	if o != nil && !IsNil(o.BaseVer) {
		return true
	}

	return false
}

// SetBaseVer gets a reference to the given string and assigns it to the BaseVer field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetBaseVer(v string) {
	o.BaseVer = &v
}

// GetBioMetaInfo returns the BioMetaInfo field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetBioMetaInfo() string {
	if o == nil || IsNil(o.BioMetaInfo) {
		var ret string
		return ret
	}
	return *o.BioMetaInfo
}

// GetBioMetaInfoOk returns a tuple with the BioMetaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetBioMetaInfoOk() (*string, bool) {
	if o == nil || IsNil(o.BioMetaInfo) {
		return nil, false
	}
	return o.BioMetaInfo, true
}

// HasBioMetaInfo returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasBioMetaInfo() bool {
	if o != nil && !IsNil(o.BioMetaInfo) {
		return true
	}

	return false
}

// SetBioMetaInfo gets a reference to the given string and assigns it to the BioMetaInfo field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetBioMetaInfo(v string) {
	o.BioMetaInfo = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetDeviceModel() string {
	if o == nil || IsNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetDeviceModelOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceModel) {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasDeviceModel() bool {
	if o != nil && !IsNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetExtInfo() FaceExtParams {
	if o == nil || IsNil(o.ExtInfo) {
		var ret FaceExtParams
		return ret
	}
	return *o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetExtInfoOk() (*FaceExtParams, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given FaceExtParams and assigns it to the ExtInfo field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetExtInfo(v FaceExtParams) {
	o.ExtInfo = &v
}

// GetMachineInfo returns the MachineInfo field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetMachineInfo() FaceMachineInfo {
	if o == nil || IsNil(o.MachineInfo) {
		var ret FaceMachineInfo
		return ret
	}
	return *o.MachineInfo
}

// GetMachineInfoOk returns a tuple with the MachineInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetMachineInfoOk() (*FaceMachineInfo, bool) {
	if o == nil || IsNil(o.MachineInfo) {
		return nil, false
	}
	return o.MachineInfo, true
}

// HasMachineInfo returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasMachineInfo() bool {
	if o != nil && !IsNil(o.MachineInfo) {
		return true
	}

	return false
}

// SetMachineInfo gets a reference to the given FaceMachineInfo and assigns it to the MachineInfo field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetMachineInfo(v FaceMachineInfo) {
	o.MachineInfo = &v
}

// GetMerchantInfo returns the MerchantInfo field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetMerchantInfo() FaceMerchantInfo {
	if o == nil || IsNil(o.MerchantInfo) {
		var ret FaceMerchantInfo
		return ret
	}
	return *o.MerchantInfo
}

// GetMerchantInfoOk returns a tuple with the MerchantInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetMerchantInfoOk() (*FaceMerchantInfo, bool) {
	if o == nil || IsNil(o.MerchantInfo) {
		return nil, false
	}
	return o.MerchantInfo, true
}

// HasMerchantInfo returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasMerchantInfo() bool {
	if o != nil && !IsNil(o.MerchantInfo) {
		return true
	}

	return false
}

// SetMerchantInfo gets a reference to the given FaceMerchantInfo and assigns it to the MerchantInfo field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetMerchantInfo(v FaceMerchantInfo) {
	o.MerchantInfo = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetOsVersion() string {
	if o == nil || IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetOsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetRemoteLogId returns the RemoteLogId field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetRemoteLogId() string {
	if o == nil || IsNil(o.RemoteLogId) {
		var ret string
		return ret
	}
	return *o.RemoteLogId
}

// GetRemoteLogIdOk returns a tuple with the RemoteLogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetRemoteLogIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteLogId) {
		return nil, false
	}
	return o.RemoteLogId, true
}

// HasRemoteLogId returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasRemoteLogId() bool {
	if o != nil && !IsNil(o.RemoteLogId) {
		return true
	}

	return false
}

// SetRemoteLogId gets a reference to the given string and assigns it to the RemoteLogId field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetRemoteLogId(v string) {
	o.RemoteLogId = &v
}

// GetZimVer returns the ZimVer field value if set, zero value otherwise.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetZimVer() string {
	if o == nil || IsNil(o.ZimVer) {
		var ret string
		return ret
	}
	return *o.ZimVer
}

// GetZimVerOk returns a tuple with the ZimVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) GetZimVerOk() (*string, bool) {
	if o == nil || IsNil(o.ZimVer) {
		return nil, false
	}
	return o.ZimVer, true
}

// HasZimVer returns a boolean if a field has been set.
func (o *ZolozAuthenticationSmilepayInitializeModel) HasZimVer() bool {
	if o != nil && !IsNil(o.ZimVer) {
		return true
	}

	return false
}

// SetZimVer gets a reference to the given string and assigns it to the ZimVer field.
func (o *ZolozAuthenticationSmilepayInitializeModel) SetZimVer(v string) {
	o.ZimVer = &v
}

func (o ZolozAuthenticationSmilepayInitializeModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZolozAuthenticationSmilepayInitializeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApdidToken) {
		toSerialize["apdid_token"] = o.ApdidToken
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AppVersion) {
		toSerialize["app_version"] = o.AppVersion
	}
	if !IsNil(o.BaseVer) {
		toSerialize["base_ver"] = o.BaseVer
	}
	if !IsNil(o.BioMetaInfo) {
		toSerialize["bio_meta_info"] = o.BioMetaInfo
	}
	if !IsNil(o.DeviceModel) {
		toSerialize["device_model"] = o.DeviceModel
	}
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.MachineInfo) {
		toSerialize["machine_info"] = o.MachineInfo
	}
	if !IsNil(o.MerchantInfo) {
		toSerialize["merchant_info"] = o.MerchantInfo
	}
	if !IsNil(o.OsVersion) {
		toSerialize["os_version"] = o.OsVersion
	}
	if !IsNil(o.RemoteLogId) {
		toSerialize["remote_log_id"] = o.RemoteLogId
	}
	if !IsNil(o.ZimVer) {
		toSerialize["zim_ver"] = o.ZimVer
	}
	return toSerialize, nil
}

type NullableZolozAuthenticationSmilepayInitializeModel struct {
	value *ZolozAuthenticationSmilepayInitializeModel
	isSet bool
}

func (v NullableZolozAuthenticationSmilepayInitializeModel) Get() *ZolozAuthenticationSmilepayInitializeModel {
	return v.value
}

func (v *NullableZolozAuthenticationSmilepayInitializeModel) Set(val *ZolozAuthenticationSmilepayInitializeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZolozAuthenticationSmilepayInitializeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZolozAuthenticationSmilepayInitializeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZolozAuthenticationSmilepayInitializeModel(val *ZolozAuthenticationSmilepayInitializeModel) *NullableZolozAuthenticationSmilepayInitializeModel {
	return &NullableZolozAuthenticationSmilepayInitializeModel{value: val, isSet: true}
}

func (v NullableZolozAuthenticationSmilepayInitializeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZolozAuthenticationSmilepayInitializeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
