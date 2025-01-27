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

// checks if the SearchAbilityOrderInfoOpenApi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchAbilityOrderInfoOpenApi{}

// SearchAbilityOrderInfoOpenApi struct for SearchAbilityOrderInfoOpenApi
type SearchAbilityOrderInfoOpenApi struct {
	// 申请单类型
	AccessType *string `json:"access_type,omitempty"`
	// 小程序名称
	AppName *string `json:"app_name,omitempty"`
	// 小程序状态
	AppStatus *string `json:"app_status,omitempty"`
	// 审核状态
	AuditStatus *string `json:"audit_status,omitempty"`
	// 运营申请单详情biz_id
	BizId *string `json:"biz_id,omitempty"`
	// 上架状态
	BoxStatus *string `json:"box_status,omitempty"`
	// 品牌模板id
	BrandTemplateId *string `json:"brand_template_id,omitempty"`
	// 二级服务唯一标识
	DataKey *string `json:"data_key,omitempty"`
	// 最近更新时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// 唯一id
	Id *string `json:"id,omitempty"`
	// 是否老工单
	IsOldData *bool `json:"is_old_data,omitempty"`
	// 服务主状态
	MajorStatus *string `json:"major_status,omitempty"`
	// 上架时间
	OnlineTime *string `json:"online_time,omitempty"`
	// 小程序可见性
	OpenStatus *bool `json:"open_status,omitempty"`
	// 下架操作者
	Operator *string `json:"operator,omitempty"`
	// 驳回原因
	RejectReason *string `json:"reject_reason,omitempty"`
	// 场景码
	SceneCode *string `json:"scene_code,omitempty"`
	// 场景名称
	SceneName *string `json:"scene_name,omitempty"`
	// 小程序id
	SearchAppId *string `json:"search_app_id,omitempty"`
	// 应用标识
	SepcCode *string `json:"sepc_code,omitempty"`
	// 服务码
	ServiceCode *string `json:"service_code,omitempty"`
	// 子功能描述
	SubServiceDesc *string `json:"sub_service_desc,omitempty"`
	// 子功能名称
	SubServiceName *string `json:"sub_service_name,omitempty"`
}

// NewSearchAbilityOrderInfoOpenApi instantiates a new SearchAbilityOrderInfoOpenApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchAbilityOrderInfoOpenApi() *SearchAbilityOrderInfoOpenApi {
	this := SearchAbilityOrderInfoOpenApi{}
	return &this
}

// NewSearchAbilityOrderInfoOpenApiWithDefaults instantiates a new SearchAbilityOrderInfoOpenApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchAbilityOrderInfoOpenApiWithDefaults() *SearchAbilityOrderInfoOpenApi {
	this := SearchAbilityOrderInfoOpenApi{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetAccessType() string {
	if o == nil || IsNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *SearchAbilityOrderInfoOpenApi) SetAccessType(v string) {
	o.AccessType = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *SearchAbilityOrderInfoOpenApi) SetAppName(v string) {
	o.AppName = &v
}

// GetAppStatus returns the AppStatus field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetAppStatus() string {
	if o == nil || IsNil(o.AppStatus) {
		var ret string
		return ret
	}
	return *o.AppStatus
}

// GetAppStatusOk returns a tuple with the AppStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetAppStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AppStatus) {
		return nil, false
	}
	return o.AppStatus, true
}

// HasAppStatus returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasAppStatus() bool {
	if o != nil && !IsNil(o.AppStatus) {
		return true
	}

	return false
}

// SetAppStatus gets a reference to the given string and assigns it to the AppStatus field.
func (o *SearchAbilityOrderInfoOpenApi) SetAppStatus(v string) {
	o.AppStatus = &v
}

// GetAuditStatus returns the AuditStatus field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetAuditStatus() string {
	if o == nil || IsNil(o.AuditStatus) {
		var ret string
		return ret
	}
	return *o.AuditStatus
}

// GetAuditStatusOk returns a tuple with the AuditStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetAuditStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AuditStatus) {
		return nil, false
	}
	return o.AuditStatus, true
}

// HasAuditStatus returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasAuditStatus() bool {
	if o != nil && !IsNil(o.AuditStatus) {
		return true
	}

	return false
}

// SetAuditStatus gets a reference to the given string and assigns it to the AuditStatus field.
func (o *SearchAbilityOrderInfoOpenApi) SetAuditStatus(v string) {
	o.AuditStatus = &v
}

// GetBizId returns the BizId field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetBizId() string {
	if o == nil || IsNil(o.BizId) {
		var ret string
		return ret
	}
	return *o.BizId
}

// GetBizIdOk returns a tuple with the BizId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetBizIdOk() (*string, bool) {
	if o == nil || IsNil(o.BizId) {
		return nil, false
	}
	return o.BizId, true
}

// HasBizId returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasBizId() bool {
	if o != nil && !IsNil(o.BizId) {
		return true
	}

	return false
}

// SetBizId gets a reference to the given string and assigns it to the BizId field.
func (o *SearchAbilityOrderInfoOpenApi) SetBizId(v string) {
	o.BizId = &v
}

// GetBoxStatus returns the BoxStatus field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetBoxStatus() string {
	if o == nil || IsNil(o.BoxStatus) {
		var ret string
		return ret
	}
	return *o.BoxStatus
}

// GetBoxStatusOk returns a tuple with the BoxStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetBoxStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BoxStatus) {
		return nil, false
	}
	return o.BoxStatus, true
}

// HasBoxStatus returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasBoxStatus() bool {
	if o != nil && !IsNil(o.BoxStatus) {
		return true
	}

	return false
}

// SetBoxStatus gets a reference to the given string and assigns it to the BoxStatus field.
func (o *SearchAbilityOrderInfoOpenApi) SetBoxStatus(v string) {
	o.BoxStatus = &v
}

// GetBrandTemplateId returns the BrandTemplateId field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetBrandTemplateId() string {
	if o == nil || IsNil(o.BrandTemplateId) {
		var ret string
		return ret
	}
	return *o.BrandTemplateId
}

// GetBrandTemplateIdOk returns a tuple with the BrandTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetBrandTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandTemplateId) {
		return nil, false
	}
	return o.BrandTemplateId, true
}

// HasBrandTemplateId returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasBrandTemplateId() bool {
	if o != nil && !IsNil(o.BrandTemplateId) {
		return true
	}

	return false
}

// SetBrandTemplateId gets a reference to the given string and assigns it to the BrandTemplateId field.
func (o *SearchAbilityOrderInfoOpenApi) SetBrandTemplateId(v string) {
	o.BrandTemplateId = &v
}

// GetDataKey returns the DataKey field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetDataKey() string {
	if o == nil || IsNil(o.DataKey) {
		var ret string
		return ret
	}
	return *o.DataKey
}

// GetDataKeyOk returns a tuple with the DataKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetDataKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DataKey) {
		return nil, false
	}
	return o.DataKey, true
}

// HasDataKey returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasDataKey() bool {
	if o != nil && !IsNil(o.DataKey) {
		return true
	}

	return false
}

// SetDataKey gets a reference to the given string and assigns it to the DataKey field.
func (o *SearchAbilityOrderInfoOpenApi) SetDataKey(v string) {
	o.DataKey = &v
}

// GetGmtModified returns the GmtModified field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetGmtModified() string {
	if o == nil || IsNil(o.GmtModified) {
		var ret string
		return ret
	}
	return *o.GmtModified
}

// GetGmtModifiedOk returns a tuple with the GmtModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetGmtModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModified) {
		return nil, false
	}
	return o.GmtModified, true
}

// HasGmtModified returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasGmtModified() bool {
	if o != nil && !IsNil(o.GmtModified) {
		return true
	}

	return false
}

// SetGmtModified gets a reference to the given string and assigns it to the GmtModified field.
func (o *SearchAbilityOrderInfoOpenApi) SetGmtModified(v string) {
	o.GmtModified = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchAbilityOrderInfoOpenApi) SetId(v string) {
	o.Id = &v
}

// GetIsOldData returns the IsOldData field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetIsOldData() bool {
	if o == nil || IsNil(o.IsOldData) {
		var ret bool
		return ret
	}
	return *o.IsOldData
}

// GetIsOldDataOk returns a tuple with the IsOldData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetIsOldDataOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOldData) {
		return nil, false
	}
	return o.IsOldData, true
}

// HasIsOldData returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasIsOldData() bool {
	if o != nil && !IsNil(o.IsOldData) {
		return true
	}

	return false
}

// SetIsOldData gets a reference to the given bool and assigns it to the IsOldData field.
func (o *SearchAbilityOrderInfoOpenApi) SetIsOldData(v bool) {
	o.IsOldData = &v
}

// GetMajorStatus returns the MajorStatus field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetMajorStatus() string {
	if o == nil || IsNil(o.MajorStatus) {
		var ret string
		return ret
	}
	return *o.MajorStatus
}

// GetMajorStatusOk returns a tuple with the MajorStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetMajorStatusOk() (*string, bool) {
	if o == nil || IsNil(o.MajorStatus) {
		return nil, false
	}
	return o.MajorStatus, true
}

// HasMajorStatus returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasMajorStatus() bool {
	if o != nil && !IsNil(o.MajorStatus) {
		return true
	}

	return false
}

// SetMajorStatus gets a reference to the given string and assigns it to the MajorStatus field.
func (o *SearchAbilityOrderInfoOpenApi) SetMajorStatus(v string) {
	o.MajorStatus = &v
}

// GetOnlineTime returns the OnlineTime field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetOnlineTime() string {
	if o == nil || IsNil(o.OnlineTime) {
		var ret string
		return ret
	}
	return *o.OnlineTime
}

// GetOnlineTimeOk returns a tuple with the OnlineTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetOnlineTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OnlineTime) {
		return nil, false
	}
	return o.OnlineTime, true
}

// HasOnlineTime returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasOnlineTime() bool {
	if o != nil && !IsNil(o.OnlineTime) {
		return true
	}

	return false
}

// SetOnlineTime gets a reference to the given string and assigns it to the OnlineTime field.
func (o *SearchAbilityOrderInfoOpenApi) SetOnlineTime(v string) {
	o.OnlineTime = &v
}

// GetOpenStatus returns the OpenStatus field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetOpenStatus() bool {
	if o == nil || IsNil(o.OpenStatus) {
		var ret bool
		return ret
	}
	return *o.OpenStatus
}

// GetOpenStatusOk returns a tuple with the OpenStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetOpenStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.OpenStatus) {
		return nil, false
	}
	return o.OpenStatus, true
}

// HasOpenStatus returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasOpenStatus() bool {
	if o != nil && !IsNil(o.OpenStatus) {
		return true
	}

	return false
}

// SetOpenStatus gets a reference to the given bool and assigns it to the OpenStatus field.
func (o *SearchAbilityOrderInfoOpenApi) SetOpenStatus(v bool) {
	o.OpenStatus = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *SearchAbilityOrderInfoOpenApi) SetOperator(v string) {
	o.Operator = &v
}

// GetRejectReason returns the RejectReason field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetRejectReason() string {
	if o == nil || IsNil(o.RejectReason) {
		var ret string
		return ret
	}
	return *o.RejectReason
}

// GetRejectReasonOk returns a tuple with the RejectReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetRejectReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RejectReason) {
		return nil, false
	}
	return o.RejectReason, true
}

// HasRejectReason returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasRejectReason() bool {
	if o != nil && !IsNil(o.RejectReason) {
		return true
	}

	return false
}

// SetRejectReason gets a reference to the given string and assigns it to the RejectReason field.
func (o *SearchAbilityOrderInfoOpenApi) SetRejectReason(v string) {
	o.RejectReason = &v
}

// GetSceneCode returns the SceneCode field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetSceneCode() string {
	if o == nil || IsNil(o.SceneCode) {
		var ret string
		return ret
	}
	return *o.SceneCode
}

// GetSceneCodeOk returns a tuple with the SceneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetSceneCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SceneCode) {
		return nil, false
	}
	return o.SceneCode, true
}

// HasSceneCode returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasSceneCode() bool {
	if o != nil && !IsNil(o.SceneCode) {
		return true
	}

	return false
}

// SetSceneCode gets a reference to the given string and assigns it to the SceneCode field.
func (o *SearchAbilityOrderInfoOpenApi) SetSceneCode(v string) {
	o.SceneCode = &v
}

// GetSceneName returns the SceneName field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetSceneName() string {
	if o == nil || IsNil(o.SceneName) {
		var ret string
		return ret
	}
	return *o.SceneName
}

// GetSceneNameOk returns a tuple with the SceneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetSceneNameOk() (*string, bool) {
	if o == nil || IsNil(o.SceneName) {
		return nil, false
	}
	return o.SceneName, true
}

// HasSceneName returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasSceneName() bool {
	if o != nil && !IsNil(o.SceneName) {
		return true
	}

	return false
}

// SetSceneName gets a reference to the given string and assigns it to the SceneName field.
func (o *SearchAbilityOrderInfoOpenApi) SetSceneName(v string) {
	o.SceneName = &v
}

// GetSearchAppId returns the SearchAppId field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetSearchAppId() string {
	if o == nil || IsNil(o.SearchAppId) {
		var ret string
		return ret
	}
	return *o.SearchAppId
}

// GetSearchAppIdOk returns a tuple with the SearchAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetSearchAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.SearchAppId) {
		return nil, false
	}
	return o.SearchAppId, true
}

// HasSearchAppId returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasSearchAppId() bool {
	if o != nil && !IsNil(o.SearchAppId) {
		return true
	}

	return false
}

// SetSearchAppId gets a reference to the given string and assigns it to the SearchAppId field.
func (o *SearchAbilityOrderInfoOpenApi) SetSearchAppId(v string) {
	o.SearchAppId = &v
}

// GetSepcCode returns the SepcCode field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetSepcCode() string {
	if o == nil || IsNil(o.SepcCode) {
		var ret string
		return ret
	}
	return *o.SepcCode
}

// GetSepcCodeOk returns a tuple with the SepcCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetSepcCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SepcCode) {
		return nil, false
	}
	return o.SepcCode, true
}

// HasSepcCode returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasSepcCode() bool {
	if o != nil && !IsNil(o.SepcCode) {
		return true
	}

	return false
}

// SetSepcCode gets a reference to the given string and assigns it to the SepcCode field.
func (o *SearchAbilityOrderInfoOpenApi) SetSepcCode(v string) {
	o.SepcCode = &v
}

// GetServiceCode returns the ServiceCode field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetServiceCode() string {
	if o == nil || IsNil(o.ServiceCode) {
		var ret string
		return ret
	}
	return *o.ServiceCode
}

// GetServiceCodeOk returns a tuple with the ServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceCode) {
		return nil, false
	}
	return o.ServiceCode, true
}

// HasServiceCode returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasServiceCode() bool {
	if o != nil && !IsNil(o.ServiceCode) {
		return true
	}

	return false
}

// SetServiceCode gets a reference to the given string and assigns it to the ServiceCode field.
func (o *SearchAbilityOrderInfoOpenApi) SetServiceCode(v string) {
	o.ServiceCode = &v
}

// GetSubServiceDesc returns the SubServiceDesc field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetSubServiceDesc() string {
	if o == nil || IsNil(o.SubServiceDesc) {
		var ret string
		return ret
	}
	return *o.SubServiceDesc
}

// GetSubServiceDescOk returns a tuple with the SubServiceDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetSubServiceDescOk() (*string, bool) {
	if o == nil || IsNil(o.SubServiceDesc) {
		return nil, false
	}
	return o.SubServiceDesc, true
}

// HasSubServiceDesc returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasSubServiceDesc() bool {
	if o != nil && !IsNil(o.SubServiceDesc) {
		return true
	}

	return false
}

// SetSubServiceDesc gets a reference to the given string and assigns it to the SubServiceDesc field.
func (o *SearchAbilityOrderInfoOpenApi) SetSubServiceDesc(v string) {
	o.SubServiceDesc = &v
}

// GetSubServiceName returns the SubServiceName field value if set, zero value otherwise.
func (o *SearchAbilityOrderInfoOpenApi) GetSubServiceName() string {
	if o == nil || IsNil(o.SubServiceName) {
		var ret string
		return ret
	}
	return *o.SubServiceName
}

// GetSubServiceNameOk returns a tuple with the SubServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderInfoOpenApi) GetSubServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubServiceName) {
		return nil, false
	}
	return o.SubServiceName, true
}

// HasSubServiceName returns a boolean if a field has been set.
func (o *SearchAbilityOrderInfoOpenApi) HasSubServiceName() bool {
	if o != nil && !IsNil(o.SubServiceName) {
		return true
	}

	return false
}

// SetSubServiceName gets a reference to the given string and assigns it to the SubServiceName field.
func (o *SearchAbilityOrderInfoOpenApi) SetSubServiceName(v string) {
	o.SubServiceName = &v
}

func (o SearchAbilityOrderInfoOpenApi) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchAbilityOrderInfoOpenApi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessType) {
		toSerialize["access_type"] = o.AccessType
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AppStatus) {
		toSerialize["app_status"] = o.AppStatus
	}
	if !IsNil(o.AuditStatus) {
		toSerialize["audit_status"] = o.AuditStatus
	}
	if !IsNil(o.BizId) {
		toSerialize["biz_id"] = o.BizId
	}
	if !IsNil(o.BoxStatus) {
		toSerialize["box_status"] = o.BoxStatus
	}
	if !IsNil(o.BrandTemplateId) {
		toSerialize["brand_template_id"] = o.BrandTemplateId
	}
	if !IsNil(o.DataKey) {
		toSerialize["data_key"] = o.DataKey
	}
	if !IsNil(o.GmtModified) {
		toSerialize["gmt_modified"] = o.GmtModified
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsOldData) {
		toSerialize["is_old_data"] = o.IsOldData
	}
	if !IsNil(o.MajorStatus) {
		toSerialize["major_status"] = o.MajorStatus
	}
	if !IsNil(o.OnlineTime) {
		toSerialize["online_time"] = o.OnlineTime
	}
	if !IsNil(o.OpenStatus) {
		toSerialize["open_status"] = o.OpenStatus
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.RejectReason) {
		toSerialize["reject_reason"] = o.RejectReason
	}
	if !IsNil(o.SceneCode) {
		toSerialize["scene_code"] = o.SceneCode
	}
	if !IsNil(o.SceneName) {
		toSerialize["scene_name"] = o.SceneName
	}
	if !IsNil(o.SearchAppId) {
		toSerialize["search_app_id"] = o.SearchAppId
	}
	if !IsNil(o.SepcCode) {
		toSerialize["sepc_code"] = o.SepcCode
	}
	if !IsNil(o.ServiceCode) {
		toSerialize["service_code"] = o.ServiceCode
	}
	if !IsNil(o.SubServiceDesc) {
		toSerialize["sub_service_desc"] = o.SubServiceDesc
	}
	if !IsNil(o.SubServiceName) {
		toSerialize["sub_service_name"] = o.SubServiceName
	}
	return toSerialize, nil
}

type NullableSearchAbilityOrderInfoOpenApi struct {
	value *SearchAbilityOrderInfoOpenApi
	isSet bool
}

func (v NullableSearchAbilityOrderInfoOpenApi) Get() *SearchAbilityOrderInfoOpenApi {
	return v.value
}

func (v *NullableSearchAbilityOrderInfoOpenApi) Set(val *SearchAbilityOrderInfoOpenApi) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAbilityOrderInfoOpenApi) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAbilityOrderInfoOpenApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAbilityOrderInfoOpenApi(val *SearchAbilityOrderInfoOpenApi) *NullableSearchAbilityOrderInfoOpenApi {
	return &NullableSearchAbilityOrderInfoOpenApi{value: val, isSet: true}
}

func (v NullableSearchAbilityOrderInfoOpenApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAbilityOrderInfoOpenApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
