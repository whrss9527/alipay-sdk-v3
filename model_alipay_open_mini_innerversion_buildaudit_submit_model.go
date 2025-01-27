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

// checks if the AlipayOpenMiniInnerversionBuildauditSubmitModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerversionBuildauditSubmitModel{}

// AlipayOpenMiniInnerversionBuildauditSubmitModel struct for AlipayOpenMiniInnerversionBuildauditSubmitModel
type AlipayOpenMiniInnerversionBuildauditSubmitModel struct {
	// 小程序类目，可不传，不传取基础信息中的小程序类目
	AppCategoryIds *string `json:"app_category_ids,omitempty"`
	// 小程序描述，可不传，不传取基础信息中的小程序描述
	AppDesc *string `json:"app_desc,omitempty"`
	// 小程序在从未上架过版本之前，英文名称是可以修改的，该字段用于在提交审核时候修改小程序英文名称。注意：小程序一旦有过上架版本之后就不可以修改英文名称。
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// 小程序logo，可不传，不传取基础信息中的小程序logo，请调用专用的logo上传接口上传logo文件，并将返回的地址作为本字段传入
	AppLogo *string `json:"app_logo,omitempty"`
	// 小程序在从未上架过版本之前，中文名称是可以修改的，该字段用于在提交审核时候修改小程序中文名称。注意：小程序一旦有过上架版本之后就不可以修改中文名称。
	AppName *string `json:"app_name,omitempty"`
	// 来源类型，新接入方需要向支付宝申请专用来源，否则不予接入，申请方式请参见接入手册。
	AppOrigin *string `json:"app_origin,omitempty"`
	// 小程序应用简介，一句话描述小程序功能，如果不填默认采用当前小程序应用简介，10~32个字符
	AppSlogan *string `json:"app_slogan,omitempty"`
	// 小程序版本号
	AppVersion *string `json:"app_version,omitempty"`
	// 构建扩展参数
	BuildExt *string `json:"build_ext,omitempty"`
	// 端信息
	BundleId *string `json:"bundle_id,omitempty"`
	// 三方应用ID
	IsvAppId    *string           `json:"isv_app_id,omitempty"`
	LicenseInfo *AuditLicenseInfo `json:"license_info,omitempty"`
	// 小程序ID，特殊场景专用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 https://docs.open.alipay.com/api_49/alipay.open.mini.category.query接口查询mini_category_list，如果不填默认采用当前小程序应用类目。使用默认应用类目后不需要再次上传营业执照号、营业执照名、营业执照截图、营业执照有效期。使用后不再读取app_category_ids值，老前台类目将废弃
	MiniCategoryIds *string `json:"mini_category_ids,omitempty"`
	// 小程序开发者ID，可不传，不传取基础信息中的小程序开发者ID
	Pid *string `json:"pid,omitempty"`
	// 服务区域类型,可不传，不传取基础信息中的小程序服务区域类型
	RegionType *string `json:"region_type,omitempty"`
	// 版本截图，最少2张，最多5张，必传
	ScreenShotList *string `json:"screen_shot_list,omitempty"`
	// 客服电话，可不传，不传取基础信息中的小程序客服电话
	ServicePhone *string `json:"service_phone,omitempty"`
	// 特殊资质图片地址列表，逗号分割,
	SpecialLicensePicList *string `json:"special_license_pic_list,omitempty"`
	// 小程序模板ID
	TemplateId *string `json:"template_id,omitempty"`
	// 模板的版本号，如果不填写，则默认用模板当前最新的在架版本
	TemplateVersion *string `json:"template_version,omitempty"`
	// 小程序版本描述，30-500个字符，一个中文占两个字符，一个英文或者数字占一个字符
	VersionDesc *string `json:"version_desc,omitempty"`
}

// NewAlipayOpenMiniInnerversionBuildauditSubmitModel instantiates a new AlipayOpenMiniInnerversionBuildauditSubmitModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerversionBuildauditSubmitModel() *AlipayOpenMiniInnerversionBuildauditSubmitModel {
	this := AlipayOpenMiniInnerversionBuildauditSubmitModel{}
	return &this
}

// NewAlipayOpenMiniInnerversionBuildauditSubmitModelWithDefaults instantiates a new AlipayOpenMiniInnerversionBuildauditSubmitModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerversionBuildauditSubmitModelWithDefaults() *AlipayOpenMiniInnerversionBuildauditSubmitModel {
	this := AlipayOpenMiniInnerversionBuildauditSubmitModel{}
	return &this
}

// GetAppCategoryIds returns the AppCategoryIds field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppCategoryIds() string {
	if o == nil || IsNil(o.AppCategoryIds) {
		var ret string
		return ret
	}
	return *o.AppCategoryIds
}

// GetAppCategoryIdsOk returns a tuple with the AppCategoryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppCategoryIdsOk() (*string, bool) {
	if o == nil || IsNil(o.AppCategoryIds) {
		return nil, false
	}
	return o.AppCategoryIds, true
}

// HasAppCategoryIds returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasAppCategoryIds() bool {
	if o != nil && !IsNil(o.AppCategoryIds) {
		return true
	}

	return false
}

// SetAppCategoryIds gets a reference to the given string and assigns it to the AppCategoryIds field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetAppCategoryIds(v string) {
	o.AppCategoryIds = &v
}

// GetAppDesc returns the AppDesc field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppDesc() string {
	if o == nil || IsNil(o.AppDesc) {
		var ret string
		return ret
	}
	return *o.AppDesc
}

// GetAppDescOk returns a tuple with the AppDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppDescOk() (*string, bool) {
	if o == nil || IsNil(o.AppDesc) {
		return nil, false
	}
	return o.AppDesc, true
}

// HasAppDesc returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasAppDesc() bool {
	if o != nil && !IsNil(o.AppDesc) {
		return true
	}

	return false
}

// SetAppDesc gets a reference to the given string and assigns it to the AppDesc field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetAppDesc(v string) {
	o.AppDesc = &v
}

// GetAppEnglishName returns the AppEnglishName field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppEnglishName() string {
	if o == nil || IsNil(o.AppEnglishName) {
		var ret string
		return ret
	}
	return *o.AppEnglishName
}

// GetAppEnglishNameOk returns a tuple with the AppEnglishName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppEnglishNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppEnglishName) {
		return nil, false
	}
	return o.AppEnglishName, true
}

// HasAppEnglishName returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasAppEnglishName() bool {
	if o != nil && !IsNil(o.AppEnglishName) {
		return true
	}

	return false
}

// SetAppEnglishName gets a reference to the given string and assigns it to the AppEnglishName field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetAppEnglishName(v string) {
	o.AppEnglishName = &v
}

// GetAppLogo returns the AppLogo field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppLogo() string {
	if o == nil || IsNil(o.AppLogo) {
		var ret string
		return ret
	}
	return *o.AppLogo
}

// GetAppLogoOk returns a tuple with the AppLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppLogoOk() (*string, bool) {
	if o == nil || IsNil(o.AppLogo) {
		return nil, false
	}
	return o.AppLogo, true
}

// HasAppLogo returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasAppLogo() bool {
	if o != nil && !IsNil(o.AppLogo) {
		return true
	}

	return false
}

// SetAppLogo gets a reference to the given string and assigns it to the AppLogo field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetAppLogo(v string) {
	o.AppLogo = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetAppName(v string) {
	o.AppName = &v
}

// GetAppOrigin returns the AppOrigin field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppOrigin() string {
	if o == nil || IsNil(o.AppOrigin) {
		var ret string
		return ret
	}
	return *o.AppOrigin
}

// GetAppOriginOk returns a tuple with the AppOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AppOrigin) {
		return nil, false
	}
	return o.AppOrigin, true
}

// HasAppOrigin returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasAppOrigin() bool {
	if o != nil && !IsNil(o.AppOrigin) {
		return true
	}

	return false
}

// SetAppOrigin gets a reference to the given string and assigns it to the AppOrigin field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetAppOrigin(v string) {
	o.AppOrigin = &v
}

// GetAppSlogan returns the AppSlogan field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppSlogan() string {
	if o == nil || IsNil(o.AppSlogan) {
		var ret string
		return ret
	}
	return *o.AppSlogan
}

// GetAppSloganOk returns a tuple with the AppSlogan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppSloganOk() (*string, bool) {
	if o == nil || IsNil(o.AppSlogan) {
		return nil, false
	}
	return o.AppSlogan, true
}

// HasAppSlogan returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasAppSlogan() bool {
	if o != nil && !IsNil(o.AppSlogan) {
		return true
	}

	return false
}

// SetAppSlogan gets a reference to the given string and assigns it to the AppSlogan field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetAppSlogan(v string) {
	o.AppSlogan = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBuildExt returns the BuildExt field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetBuildExt() string {
	if o == nil || IsNil(o.BuildExt) {
		var ret string
		return ret
	}
	return *o.BuildExt
}

// GetBuildExtOk returns a tuple with the BuildExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetBuildExtOk() (*string, bool) {
	if o == nil || IsNil(o.BuildExt) {
		return nil, false
	}
	return o.BuildExt, true
}

// HasBuildExt returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasBuildExt() bool {
	if o != nil && !IsNil(o.BuildExt) {
		return true
	}

	return false
}

// SetBuildExt gets a reference to the given string and assigns it to the BuildExt field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetBuildExt(v string) {
	o.BuildExt = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetBundleId(v string) {
	o.BundleId = &v
}

// GetIsvAppId returns the IsvAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetIsvAppId() string {
	if o == nil || IsNil(o.IsvAppId) {
		var ret string
		return ret
	}
	return *o.IsvAppId
}

// GetIsvAppIdOk returns a tuple with the IsvAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetIsvAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.IsvAppId) {
		return nil, false
	}
	return o.IsvAppId, true
}

// HasIsvAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasIsvAppId() bool {
	if o != nil && !IsNil(o.IsvAppId) {
		return true
	}

	return false
}

// SetIsvAppId gets a reference to the given string and assigns it to the IsvAppId field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetIsvAppId(v string) {
	o.IsvAppId = &v
}

// GetLicenseInfo returns the LicenseInfo field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetLicenseInfo() AuditLicenseInfo {
	if o == nil || IsNil(o.LicenseInfo) {
		var ret AuditLicenseInfo
		return ret
	}
	return *o.LicenseInfo
}

// GetLicenseInfoOk returns a tuple with the LicenseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetLicenseInfoOk() (*AuditLicenseInfo, bool) {
	if o == nil || IsNil(o.LicenseInfo) {
		return nil, false
	}
	return o.LicenseInfo, true
}

// HasLicenseInfo returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasLicenseInfo() bool {
	if o != nil && !IsNil(o.LicenseInfo) {
		return true
	}

	return false
}

// SetLicenseInfo gets a reference to the given AuditLicenseInfo and assigns it to the LicenseInfo field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetLicenseInfo(v AuditLicenseInfo) {
	o.LicenseInfo = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetMiniCategoryIds returns the MiniCategoryIds field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetMiniCategoryIds() string {
	if o == nil || IsNil(o.MiniCategoryIds) {
		var ret string
		return ret
	}
	return *o.MiniCategoryIds
}

// GetMiniCategoryIdsOk returns a tuple with the MiniCategoryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetMiniCategoryIdsOk() (*string, bool) {
	if o == nil || IsNil(o.MiniCategoryIds) {
		return nil, false
	}
	return o.MiniCategoryIds, true
}

// HasMiniCategoryIds returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasMiniCategoryIds() bool {
	if o != nil && !IsNil(o.MiniCategoryIds) {
		return true
	}

	return false
}

// SetMiniCategoryIds gets a reference to the given string and assigns it to the MiniCategoryIds field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetMiniCategoryIds(v string) {
	o.MiniCategoryIds = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetPid(v string) {
	o.Pid = &v
}

// GetRegionType returns the RegionType field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetRegionType() string {
	if o == nil || IsNil(o.RegionType) {
		var ret string
		return ret
	}
	return *o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetRegionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionType) {
		return nil, false
	}
	return o.RegionType, true
}

// HasRegionType returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasRegionType() bool {
	if o != nil && !IsNil(o.RegionType) {
		return true
	}

	return false
}

// SetRegionType gets a reference to the given string and assigns it to the RegionType field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetRegionType(v string) {
	o.RegionType = &v
}

// GetScreenShotList returns the ScreenShotList field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetScreenShotList() string {
	if o == nil || IsNil(o.ScreenShotList) {
		var ret string
		return ret
	}
	return *o.ScreenShotList
}

// GetScreenShotListOk returns a tuple with the ScreenShotList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetScreenShotListOk() (*string, bool) {
	if o == nil || IsNil(o.ScreenShotList) {
		return nil, false
	}
	return o.ScreenShotList, true
}

// HasScreenShotList returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasScreenShotList() bool {
	if o != nil && !IsNil(o.ScreenShotList) {
		return true
	}

	return false
}

// SetScreenShotList gets a reference to the given string and assigns it to the ScreenShotList field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetScreenShotList(v string) {
	o.ScreenShotList = &v
}

// GetServicePhone returns the ServicePhone field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetServicePhone() string {
	if o == nil || IsNil(o.ServicePhone) {
		var ret string
		return ret
	}
	return *o.ServicePhone
}

// GetServicePhoneOk returns a tuple with the ServicePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetServicePhoneOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePhone) {
		return nil, false
	}
	return o.ServicePhone, true
}

// HasServicePhone returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasServicePhone() bool {
	if o != nil && !IsNil(o.ServicePhone) {
		return true
	}

	return false
}

// SetServicePhone gets a reference to the given string and assigns it to the ServicePhone field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetServicePhone(v string) {
	o.ServicePhone = &v
}

// GetSpecialLicensePicList returns the SpecialLicensePicList field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetSpecialLicensePicList() string {
	if o == nil || IsNil(o.SpecialLicensePicList) {
		var ret string
		return ret
	}
	return *o.SpecialLicensePicList
}

// GetSpecialLicensePicListOk returns a tuple with the SpecialLicensePicList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetSpecialLicensePicListOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialLicensePicList) {
		return nil, false
	}
	return o.SpecialLicensePicList, true
}

// HasSpecialLicensePicList returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasSpecialLicensePicList() bool {
	if o != nil && !IsNil(o.SpecialLicensePicList) {
		return true
	}

	return false
}

// SetSpecialLicensePicList gets a reference to the given string and assigns it to the SpecialLicensePicList field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetSpecialLicensePicList(v string) {
	o.SpecialLicensePicList = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateVersion returns the TemplateVersion field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetTemplateVersion() string {
	if o == nil || IsNil(o.TemplateVersion) {
		var ret string
		return ret
	}
	return *o.TemplateVersion
}

// GetTemplateVersionOk returns a tuple with the TemplateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetTemplateVersionOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateVersion) {
		return nil, false
	}
	return o.TemplateVersion, true
}

// HasTemplateVersion returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasTemplateVersion() bool {
	if o != nil && !IsNil(o.TemplateVersion) {
		return true
	}

	return false
}

// SetTemplateVersion gets a reference to the given string and assigns it to the TemplateVersion field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetTemplateVersion(v string) {
	o.TemplateVersion = &v
}

// GetVersionDesc returns the VersionDesc field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetVersionDesc() string {
	if o == nil || IsNil(o.VersionDesc) {
		var ret string
		return ret
	}
	return *o.VersionDesc
}

// GetVersionDescOk returns a tuple with the VersionDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) GetVersionDescOk() (*string, bool) {
	if o == nil || IsNil(o.VersionDesc) {
		return nil, false
	}
	return o.VersionDesc, true
}

// HasVersionDesc returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) HasVersionDesc() bool {
	if o != nil && !IsNil(o.VersionDesc) {
		return true
	}

	return false
}

// SetVersionDesc gets a reference to the given string and assigns it to the VersionDesc field.
func (o *AlipayOpenMiniInnerversionBuildauditSubmitModel) SetVersionDesc(v string) {
	o.VersionDesc = &v
}

func (o AlipayOpenMiniInnerversionBuildauditSubmitModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerversionBuildauditSubmitModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppCategoryIds) {
		toSerialize["app_category_ids"] = o.AppCategoryIds
	}
	if !IsNil(o.AppDesc) {
		toSerialize["app_desc"] = o.AppDesc
	}
	if !IsNil(o.AppEnglishName) {
		toSerialize["app_english_name"] = o.AppEnglishName
	}
	if !IsNil(o.AppLogo) {
		toSerialize["app_logo"] = o.AppLogo
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AppOrigin) {
		toSerialize["app_origin"] = o.AppOrigin
	}
	if !IsNil(o.AppSlogan) {
		toSerialize["app_slogan"] = o.AppSlogan
	}
	if !IsNil(o.AppVersion) {
		toSerialize["app_version"] = o.AppVersion
	}
	if !IsNil(o.BuildExt) {
		toSerialize["build_ext"] = o.BuildExt
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.IsvAppId) {
		toSerialize["isv_app_id"] = o.IsvAppId
	}
	if !IsNil(o.LicenseInfo) {
		toSerialize["license_info"] = o.LicenseInfo
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.MiniCategoryIds) {
		toSerialize["mini_category_ids"] = o.MiniCategoryIds
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.RegionType) {
		toSerialize["region_type"] = o.RegionType
	}
	if !IsNil(o.ScreenShotList) {
		toSerialize["screen_shot_list"] = o.ScreenShotList
	}
	if !IsNil(o.ServicePhone) {
		toSerialize["service_phone"] = o.ServicePhone
	}
	if !IsNil(o.SpecialLicensePicList) {
		toSerialize["special_license_pic_list"] = o.SpecialLicensePicList
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.TemplateVersion) {
		toSerialize["template_version"] = o.TemplateVersion
	}
	if !IsNil(o.VersionDesc) {
		toSerialize["version_desc"] = o.VersionDesc
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerversionBuildauditSubmitModel struct {
	value *AlipayOpenMiniInnerversionBuildauditSubmitModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionBuildauditSubmitModel) Get() *AlipayOpenMiniInnerversionBuildauditSubmitModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionBuildauditSubmitModel) Set(val *AlipayOpenMiniInnerversionBuildauditSubmitModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionBuildauditSubmitModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionBuildauditSubmitModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionBuildauditSubmitModel(val *AlipayOpenMiniInnerversionBuildauditSubmitModel) *NullableAlipayOpenMiniInnerversionBuildauditSubmitModel {
	return &NullableAlipayOpenMiniInnerversionBuildauditSubmitModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionBuildauditSubmitModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionBuildauditSubmitModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
