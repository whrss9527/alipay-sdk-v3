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

// checks if the PromiseDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromiseDetail{}

// PromiseDetail struct for PromiseDetail
type PromiseDetail struct {
	// 授权状态
	AuthStatus *bool `json:"auth_status,omitempty"`
	// 任务创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 任务结束时间
	EndTime *string `json:"end_time,omitempty"`
	// 任务到达终态的时间
	FinalTime *string `json:"final_time,omitempty"`
	// 任务完成期数
	FinishPeriods *int32 `json:"finish_periods,omitempty"`
	// 芝麻侧的商户id
	MerchantId *string `json:"merchant_id,omitempty"`
	// 商户logo
	MerchantLogo *string `json:"merchant_logo,omitempty"`
	// 商户名称
	MerchantName *string `json:"merchant_name,omitempty"`
	// 生活记录加入时的外部业务号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 周期类型
	PeriodType *string `json:"period_type,omitempty"`
	// 生活记录模板名称
	PromiseName *string `json:"promise_name,omitempty"`
	// 生活记录主记录id
	RecordId *string `json:"record_id,omitempty"`
	// 主任务状态
	RecordStatus *string `json:"record_status,omitempty"`
	// 任务开始时间
	StartTime *string `json:"start_time,omitempty"`
	// 子记录状态
	SubRecordStatus *string `json:"sub_record_status,omitempty"`
	// 副标题
	SubTitle *string `json:"sub_title,omitempty"`
	// 生活记录模板id
	TemplateId *string `json:"template_id,omitempty"`
	// 任务总期数
	TotalPeriods *int32 `json:"total_periods,omitempty"`
}

// NewPromiseDetail instantiates a new PromiseDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromiseDetail() *PromiseDetail {
	this := PromiseDetail{}
	return &this
}

// NewPromiseDetailWithDefaults instantiates a new PromiseDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromiseDetailWithDefaults() *PromiseDetail {
	this := PromiseDetail{}
	return &this
}

// GetAuthStatus returns the AuthStatus field value if set, zero value otherwise.
func (o *PromiseDetail) GetAuthStatus() bool {
	if o == nil || IsNil(o.AuthStatus) {
		var ret bool
		return ret
	}
	return *o.AuthStatus
}

// GetAuthStatusOk returns a tuple with the AuthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetAuthStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthStatus) {
		return nil, false
	}
	return o.AuthStatus, true
}

// HasAuthStatus returns a boolean if a field has been set.
func (o *PromiseDetail) HasAuthStatus() bool {
	if o != nil && !IsNil(o.AuthStatus) {
		return true
	}

	return false
}

// SetAuthStatus gets a reference to the given bool and assigns it to the AuthStatus field.
func (o *PromiseDetail) SetAuthStatus(v bool) {
	o.AuthStatus = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *PromiseDetail) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *PromiseDetail) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *PromiseDetail) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *PromiseDetail) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *PromiseDetail) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *PromiseDetail) SetEndTime(v string) {
	o.EndTime = &v
}

// GetFinalTime returns the FinalTime field value if set, zero value otherwise.
func (o *PromiseDetail) GetFinalTime() string {
	if o == nil || IsNil(o.FinalTime) {
		var ret string
		return ret
	}
	return *o.FinalTime
}

// GetFinalTimeOk returns a tuple with the FinalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetFinalTimeOk() (*string, bool) {
	if o == nil || IsNil(o.FinalTime) {
		return nil, false
	}
	return o.FinalTime, true
}

// HasFinalTime returns a boolean if a field has been set.
func (o *PromiseDetail) HasFinalTime() bool {
	if o != nil && !IsNil(o.FinalTime) {
		return true
	}

	return false
}

// SetFinalTime gets a reference to the given string and assigns it to the FinalTime field.
func (o *PromiseDetail) SetFinalTime(v string) {
	o.FinalTime = &v
}

// GetFinishPeriods returns the FinishPeriods field value if set, zero value otherwise.
func (o *PromiseDetail) GetFinishPeriods() int32 {
	if o == nil || IsNil(o.FinishPeriods) {
		var ret int32
		return ret
	}
	return *o.FinishPeriods
}

// GetFinishPeriodsOk returns a tuple with the FinishPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetFinishPeriodsOk() (*int32, bool) {
	if o == nil || IsNil(o.FinishPeriods) {
		return nil, false
	}
	return o.FinishPeriods, true
}

// HasFinishPeriods returns a boolean if a field has been set.
func (o *PromiseDetail) HasFinishPeriods() bool {
	if o != nil && !IsNil(o.FinishPeriods) {
		return true
	}

	return false
}

// SetFinishPeriods gets a reference to the given int32 and assigns it to the FinishPeriods field.
func (o *PromiseDetail) SetFinishPeriods(v int32) {
	o.FinishPeriods = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *PromiseDetail) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *PromiseDetail) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *PromiseDetail) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetMerchantLogo returns the MerchantLogo field value if set, zero value otherwise.
func (o *PromiseDetail) GetMerchantLogo() string {
	if o == nil || IsNil(o.MerchantLogo) {
		var ret string
		return ret
	}
	return *o.MerchantLogo
}

// GetMerchantLogoOk returns a tuple with the MerchantLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetMerchantLogoOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantLogo) {
		return nil, false
	}
	return o.MerchantLogo, true
}

// HasMerchantLogo returns a boolean if a field has been set.
func (o *PromiseDetail) HasMerchantLogo() bool {
	if o != nil && !IsNil(o.MerchantLogo) {
		return true
	}

	return false
}

// SetMerchantLogo gets a reference to the given string and assigns it to the MerchantLogo field.
func (o *PromiseDetail) SetMerchantLogo(v string) {
	o.MerchantLogo = &v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *PromiseDetail) GetMerchantName() string {
	if o == nil || IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *PromiseDetail) HasMerchantName() bool {
	if o != nil && !IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *PromiseDetail) SetMerchantName(v string) {
	o.MerchantName = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *PromiseDetail) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *PromiseDetail) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *PromiseDetail) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPeriodType returns the PeriodType field value if set, zero value otherwise.
func (o *PromiseDetail) GetPeriodType() string {
	if o == nil || IsNil(o.PeriodType) {
		var ret string
		return ret
	}
	return *o.PeriodType
}

// GetPeriodTypeOk returns a tuple with the PeriodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetPeriodTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodType) {
		return nil, false
	}
	return o.PeriodType, true
}

// HasPeriodType returns a boolean if a field has been set.
func (o *PromiseDetail) HasPeriodType() bool {
	if o != nil && !IsNil(o.PeriodType) {
		return true
	}

	return false
}

// SetPeriodType gets a reference to the given string and assigns it to the PeriodType field.
func (o *PromiseDetail) SetPeriodType(v string) {
	o.PeriodType = &v
}

// GetPromiseName returns the PromiseName field value if set, zero value otherwise.
func (o *PromiseDetail) GetPromiseName() string {
	if o == nil || IsNil(o.PromiseName) {
		var ret string
		return ret
	}
	return *o.PromiseName
}

// GetPromiseNameOk returns a tuple with the PromiseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetPromiseNameOk() (*string, bool) {
	if o == nil || IsNil(o.PromiseName) {
		return nil, false
	}
	return o.PromiseName, true
}

// HasPromiseName returns a boolean if a field has been set.
func (o *PromiseDetail) HasPromiseName() bool {
	if o != nil && !IsNil(o.PromiseName) {
		return true
	}

	return false
}

// SetPromiseName gets a reference to the given string and assigns it to the PromiseName field.
func (o *PromiseDetail) SetPromiseName(v string) {
	o.PromiseName = &v
}

// GetRecordId returns the RecordId field value if set, zero value otherwise.
func (o *PromiseDetail) GetRecordId() string {
	if o == nil || IsNil(o.RecordId) {
		var ret string
		return ret
	}
	return *o.RecordId
}

// GetRecordIdOk returns a tuple with the RecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecordId) {
		return nil, false
	}
	return o.RecordId, true
}

// HasRecordId returns a boolean if a field has been set.
func (o *PromiseDetail) HasRecordId() bool {
	if o != nil && !IsNil(o.RecordId) {
		return true
	}

	return false
}

// SetRecordId gets a reference to the given string and assigns it to the RecordId field.
func (o *PromiseDetail) SetRecordId(v string) {
	o.RecordId = &v
}

// GetRecordStatus returns the RecordStatus field value if set, zero value otherwise.
func (o *PromiseDetail) GetRecordStatus() string {
	if o == nil || IsNil(o.RecordStatus) {
		var ret string
		return ret
	}
	return *o.RecordStatus
}

// GetRecordStatusOk returns a tuple with the RecordStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetRecordStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RecordStatus) {
		return nil, false
	}
	return o.RecordStatus, true
}

// HasRecordStatus returns a boolean if a field has been set.
func (o *PromiseDetail) HasRecordStatus() bool {
	if o != nil && !IsNil(o.RecordStatus) {
		return true
	}

	return false
}

// SetRecordStatus gets a reference to the given string and assigns it to the RecordStatus field.
func (o *PromiseDetail) SetRecordStatus(v string) {
	o.RecordStatus = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *PromiseDetail) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *PromiseDetail) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *PromiseDetail) SetStartTime(v string) {
	o.StartTime = &v
}

// GetSubRecordStatus returns the SubRecordStatus field value if set, zero value otherwise.
func (o *PromiseDetail) GetSubRecordStatus() string {
	if o == nil || IsNil(o.SubRecordStatus) {
		var ret string
		return ret
	}
	return *o.SubRecordStatus
}

// GetSubRecordStatusOk returns a tuple with the SubRecordStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetSubRecordStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SubRecordStatus) {
		return nil, false
	}
	return o.SubRecordStatus, true
}

// HasSubRecordStatus returns a boolean if a field has been set.
func (o *PromiseDetail) HasSubRecordStatus() bool {
	if o != nil && !IsNil(o.SubRecordStatus) {
		return true
	}

	return false
}

// SetSubRecordStatus gets a reference to the given string and assigns it to the SubRecordStatus field.
func (o *PromiseDetail) SetSubRecordStatus(v string) {
	o.SubRecordStatus = &v
}

// GetSubTitle returns the SubTitle field value if set, zero value otherwise.
func (o *PromiseDetail) GetSubTitle() string {
	if o == nil || IsNil(o.SubTitle) {
		var ret string
		return ret
	}
	return *o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetSubTitleOk() (*string, bool) {
	if o == nil || IsNil(o.SubTitle) {
		return nil, false
	}
	return o.SubTitle, true
}

// HasSubTitle returns a boolean if a field has been set.
func (o *PromiseDetail) HasSubTitle() bool {
	if o != nil && !IsNil(o.SubTitle) {
		return true
	}

	return false
}

// SetSubTitle gets a reference to the given string and assigns it to the SubTitle field.
func (o *PromiseDetail) SetSubTitle(v string) {
	o.SubTitle = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *PromiseDetail) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *PromiseDetail) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *PromiseDetail) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTotalPeriods returns the TotalPeriods field value if set, zero value otherwise.
func (o *PromiseDetail) GetTotalPeriods() int32 {
	if o == nil || IsNil(o.TotalPeriods) {
		var ret int32
		return ret
	}
	return *o.TotalPeriods
}

// GetTotalPeriodsOk returns a tuple with the TotalPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromiseDetail) GetTotalPeriodsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPeriods) {
		return nil, false
	}
	return o.TotalPeriods, true
}

// HasTotalPeriods returns a boolean if a field has been set.
func (o *PromiseDetail) HasTotalPeriods() bool {
	if o != nil && !IsNil(o.TotalPeriods) {
		return true
	}

	return false
}

// SetTotalPeriods gets a reference to the given int32 and assigns it to the TotalPeriods field.
func (o *PromiseDetail) SetTotalPeriods(v int32) {
	o.TotalPeriods = &v
}

func (o PromiseDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromiseDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthStatus) {
		toSerialize["auth_status"] = o.AuthStatus
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.FinalTime) {
		toSerialize["final_time"] = o.FinalTime
	}
	if !IsNil(o.FinishPeriods) {
		toSerialize["finish_periods"] = o.FinishPeriods
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchant_id"] = o.MerchantId
	}
	if !IsNil(o.MerchantLogo) {
		toSerialize["merchant_logo"] = o.MerchantLogo
	}
	if !IsNil(o.MerchantName) {
		toSerialize["merchant_name"] = o.MerchantName
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PeriodType) {
		toSerialize["period_type"] = o.PeriodType
	}
	if !IsNil(o.PromiseName) {
		toSerialize["promise_name"] = o.PromiseName
	}
	if !IsNil(o.RecordId) {
		toSerialize["record_id"] = o.RecordId
	}
	if !IsNil(o.RecordStatus) {
		toSerialize["record_status"] = o.RecordStatus
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.SubRecordStatus) {
		toSerialize["sub_record_status"] = o.SubRecordStatus
	}
	if !IsNil(o.SubTitle) {
		toSerialize["sub_title"] = o.SubTitle
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.TotalPeriods) {
		toSerialize["total_periods"] = o.TotalPeriods
	}
	return toSerialize, nil
}

type NullablePromiseDetail struct {
	value *PromiseDetail
	isSet bool
}

func (v NullablePromiseDetail) Get() *PromiseDetail {
	return v.value
}

func (v *NullablePromiseDetail) Set(val *PromiseDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePromiseDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePromiseDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromiseDetail(val *PromiseDetail) *NullablePromiseDetail {
	return &NullablePromiseDetail{value: val, isSet: true}
}

func (v NullablePromiseDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromiseDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
