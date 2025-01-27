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

// checks if the AlipayIserviceCcmSwArticleGetResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmSwArticleGetResponseModel{}

// AlipayIserviceCcmSwArticleGetResponseModel struct for AlipayIserviceCcmSwArticleGetResponseModel
type AlipayIserviceCcmSwArticleGetResponseModel struct {
	// 文章对应附件集合
	Attachments []ArticleAttachmentInfo `json:"attachments,omitempty"`
	// 所属类目ID
	CategoryId *int32 `json:"category_id,omitempty"`
	// 类目名称路径
	CategoryNamePath *string `json:"category_name_path,omitempty"`
	// 类目路径
	CategoryPath []ArticleCategoryInfo `json:"category_path,omitempty"`
	// 内容
	Content *string `json:"content,omitempty"`
	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 创建人
	CreatorId *string `json:"creator_id,omitempty"`
	// 扩展标题
	ExtendTitles []string `json:"extend_titles,omitempty"`
	// 文章ID
	Id *int32 `json:"id,omitempty"`
	// 标签
	Keywords []string `json:"keywords,omitempty"`
	// 知识库ID
	LibraryId *int32 `json:"library_id,omitempty"`
	// 排序值
	OrderNo *int32 `json:"order_no,omitempty"`
	// 文章对应图片集合
	Pictures []ArticleAttachmentInfo `json:"pictures,omitempty"`
	// 有效期止
	PublishEnd *string `json:"publish_end,omitempty"`
	// 有效期起始
	PublishStart *string `json:"publish_start,omitempty"`
	// 场景ID。1（内部知识库）；2（机器人）;3（帮助中心）；4（无线帮助中心）
	SceneCodes []string `json:"scene_codes,omitempty"`
	// 来源
	Source *string `json:"source,omitempty"`
	// 文章状态
	Status *string `json:"status,omitempty"`
	// 文章状态码
	StatusCode *string `json:"status_code,omitempty"`
	// 标题
	Title *string `json:"title,omitempty"`
	// 修改时间
	UpdateTime *string `json:"update_time,omitempty"`
	// 修改人
	UpdaterId *string `json:"updater_id,omitempty"`
	// 修改人名称
	UpdaterName *string `json:"updater_name,omitempty"`
}

// NewAlipayIserviceCcmSwArticleGetResponseModel instantiates a new AlipayIserviceCcmSwArticleGetResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmSwArticleGetResponseModel() *AlipayIserviceCcmSwArticleGetResponseModel {
	this := AlipayIserviceCcmSwArticleGetResponseModel{}
	return &this
}

// NewAlipayIserviceCcmSwArticleGetResponseModelWithDefaults instantiates a new AlipayIserviceCcmSwArticleGetResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmSwArticleGetResponseModelWithDefaults() *AlipayIserviceCcmSwArticleGetResponseModel {
	this := AlipayIserviceCcmSwArticleGetResponseModel{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetAttachments() []ArticleAttachmentInfo {
	if o == nil || IsNil(o.Attachments) {
		var ret []ArticleAttachmentInfo
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetAttachmentsOk() ([]ArticleAttachmentInfo, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []ArticleAttachmentInfo and assigns it to the Attachments field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetAttachments(v []ArticleAttachmentInfo) {
	o.Attachments = v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetCategoryId() int32 {
	if o == nil || IsNil(o.CategoryId) {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetCategoryNamePath returns the CategoryNamePath field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetCategoryNamePath() string {
	if o == nil || IsNil(o.CategoryNamePath) {
		var ret string
		return ret
	}
	return *o.CategoryNamePath
}

// GetCategoryNamePathOk returns a tuple with the CategoryNamePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetCategoryNamePathOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryNamePath) {
		return nil, false
	}
	return o.CategoryNamePath, true
}

// HasCategoryNamePath returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasCategoryNamePath() bool {
	if o != nil && !IsNil(o.CategoryNamePath) {
		return true
	}

	return false
}

// SetCategoryNamePath gets a reference to the given string and assigns it to the CategoryNamePath field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetCategoryNamePath(v string) {
	o.CategoryNamePath = &v
}

// GetCategoryPath returns the CategoryPath field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetCategoryPath() []ArticleCategoryInfo {
	if o == nil || IsNil(o.CategoryPath) {
		var ret []ArticleCategoryInfo
		return ret
	}
	return o.CategoryPath
}

// GetCategoryPathOk returns a tuple with the CategoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetCategoryPathOk() ([]ArticleCategoryInfo, bool) {
	if o == nil || IsNil(o.CategoryPath) {
		return nil, false
	}
	return o.CategoryPath, true
}

// HasCategoryPath returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasCategoryPath() bool {
	if o != nil && !IsNil(o.CategoryPath) {
		return true
	}

	return false
}

// SetCategoryPath gets a reference to the given []ArticleCategoryInfo and assigns it to the CategoryPath field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetCategoryPath(v []ArticleCategoryInfo) {
	o.CategoryPath = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetContent(v string) {
	o.Content = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetExtendTitles returns the ExtendTitles field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetExtendTitles() []string {
	if o == nil || IsNil(o.ExtendTitles) {
		var ret []string
		return ret
	}
	return o.ExtendTitles
}

// GetExtendTitlesOk returns a tuple with the ExtendTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetExtendTitlesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtendTitles) {
		return nil, false
	}
	return o.ExtendTitles, true
}

// HasExtendTitles returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasExtendTitles() bool {
	if o != nil && !IsNil(o.ExtendTitles) {
		return true
	}

	return false
}

// SetExtendTitles gets a reference to the given []string and assigns it to the ExtendTitles field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetExtendTitles(v []string) {
	o.ExtendTitles = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetId(v int32) {
	o.Id = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetKeywords() []string {
	if o == nil || IsNil(o.Keywords) {
		var ret []string
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []string and assigns it to the Keywords field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetKeywords(v []string) {
	o.Keywords = v
}

// GetLibraryId returns the LibraryId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetLibraryId() int32 {
	if o == nil || IsNil(o.LibraryId) {
		var ret int32
		return ret
	}
	return *o.LibraryId
}

// GetLibraryIdOk returns a tuple with the LibraryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetLibraryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LibraryId) {
		return nil, false
	}
	return o.LibraryId, true
}

// HasLibraryId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasLibraryId() bool {
	if o != nil && !IsNil(o.LibraryId) {
		return true
	}

	return false
}

// SetLibraryId gets a reference to the given int32 and assigns it to the LibraryId field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetLibraryId(v int32) {
	o.LibraryId = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetOrderNo() int32 {
	if o == nil || IsNil(o.OrderNo) {
		var ret int32
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetOrderNoOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given int32 and assigns it to the OrderNo field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetOrderNo(v int32) {
	o.OrderNo = &v
}

// GetPictures returns the Pictures field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetPictures() []ArticleAttachmentInfo {
	if o == nil || IsNil(o.Pictures) {
		var ret []ArticleAttachmentInfo
		return ret
	}
	return o.Pictures
}

// GetPicturesOk returns a tuple with the Pictures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetPicturesOk() ([]ArticleAttachmentInfo, bool) {
	if o == nil || IsNil(o.Pictures) {
		return nil, false
	}
	return o.Pictures, true
}

// HasPictures returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasPictures() bool {
	if o != nil && !IsNil(o.Pictures) {
		return true
	}

	return false
}

// SetPictures gets a reference to the given []ArticleAttachmentInfo and assigns it to the Pictures field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetPictures(v []ArticleAttachmentInfo) {
	o.Pictures = v
}

// GetPublishEnd returns the PublishEnd field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetPublishEnd() string {
	if o == nil || IsNil(o.PublishEnd) {
		var ret string
		return ret
	}
	return *o.PublishEnd
}

// GetPublishEndOk returns a tuple with the PublishEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetPublishEndOk() (*string, bool) {
	if o == nil || IsNil(o.PublishEnd) {
		return nil, false
	}
	return o.PublishEnd, true
}

// HasPublishEnd returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasPublishEnd() bool {
	if o != nil && !IsNil(o.PublishEnd) {
		return true
	}

	return false
}

// SetPublishEnd gets a reference to the given string and assigns it to the PublishEnd field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetPublishEnd(v string) {
	o.PublishEnd = &v
}

// GetPublishStart returns the PublishStart field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetPublishStart() string {
	if o == nil || IsNil(o.PublishStart) {
		var ret string
		return ret
	}
	return *o.PublishStart
}

// GetPublishStartOk returns a tuple with the PublishStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetPublishStartOk() (*string, bool) {
	if o == nil || IsNil(o.PublishStart) {
		return nil, false
	}
	return o.PublishStart, true
}

// HasPublishStart returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasPublishStart() bool {
	if o != nil && !IsNil(o.PublishStart) {
		return true
	}

	return false
}

// SetPublishStart gets a reference to the given string and assigns it to the PublishStart field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetPublishStart(v string) {
	o.PublishStart = &v
}

// GetSceneCodes returns the SceneCodes field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetSceneCodes() []string {
	if o == nil || IsNil(o.SceneCodes) {
		var ret []string
		return ret
	}
	return o.SceneCodes
}

// GetSceneCodesOk returns a tuple with the SceneCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetSceneCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.SceneCodes) {
		return nil, false
	}
	return o.SceneCodes, true
}

// HasSceneCodes returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasSceneCodes() bool {
	if o != nil && !IsNil(o.SceneCodes) {
		return true
	}

	return false
}

// SetSceneCodes gets a reference to the given []string and assigns it to the SceneCodes field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetSceneCodes(v []string) {
	o.SceneCodes = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetStatus(v string) {
	o.Status = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetStatusCode() string {
	if o == nil || IsNil(o.StatusCode) {
		var ret string
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given string and assigns it to the StatusCode field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetStatusCode(v string) {
	o.StatusCode = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetTitle(v string) {
	o.Title = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetUpdaterId returns the UpdaterId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetUpdaterId() string {
	if o == nil || IsNil(o.UpdaterId) {
		var ret string
		return ret
	}
	return *o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetUpdaterIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpdaterId) {
		return nil, false
	}
	return o.UpdaterId, true
}

// HasUpdaterId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasUpdaterId() bool {
	if o != nil && !IsNil(o.UpdaterId) {
		return true
	}

	return false
}

// SetUpdaterId gets a reference to the given string and assigns it to the UpdaterId field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetUpdaterId(v string) {
	o.UpdaterId = &v
}

// GetUpdaterName returns the UpdaterName field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetUpdaterName() string {
	if o == nil || IsNil(o.UpdaterName) {
		var ret string
		return ret
	}
	return *o.UpdaterName
}

// GetUpdaterNameOk returns a tuple with the UpdaterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) GetUpdaterNameOk() (*string, bool) {
	if o == nil || IsNil(o.UpdaterName) {
		return nil, false
	}
	return o.UpdaterName, true
}

// HasUpdaterName returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) HasUpdaterName() bool {
	if o != nil && !IsNil(o.UpdaterName) {
		return true
	}

	return false
}

// SetUpdaterName gets a reference to the given string and assigns it to the UpdaterName field.
func (o *AlipayIserviceCcmSwArticleGetResponseModel) SetUpdaterName(v string) {
	o.UpdaterName = &v
}

func (o AlipayIserviceCcmSwArticleGetResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmSwArticleGetResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.CategoryNamePath) {
		toSerialize["category_name_path"] = o.CategoryNamePath
	}
	if !IsNil(o.CategoryPath) {
		toSerialize["category_path"] = o.CategoryPath
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creator_id"] = o.CreatorId
	}
	if !IsNil(o.ExtendTitles) {
		toSerialize["extend_titles"] = o.ExtendTitles
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.LibraryId) {
		toSerialize["library_id"] = o.LibraryId
	}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	if !IsNil(o.Pictures) {
		toSerialize["pictures"] = o.Pictures
	}
	if !IsNil(o.PublishEnd) {
		toSerialize["publish_end"] = o.PublishEnd
	}
	if !IsNil(o.PublishStart) {
		toSerialize["publish_start"] = o.PublishStart
	}
	if !IsNil(o.SceneCodes) {
		toSerialize["scene_codes"] = o.SceneCodes
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	if !IsNil(o.UpdaterId) {
		toSerialize["updater_id"] = o.UpdaterId
	}
	if !IsNil(o.UpdaterName) {
		toSerialize["updater_name"] = o.UpdaterName
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmSwArticleGetResponseModel struct {
	value *AlipayIserviceCcmSwArticleGetResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmSwArticleGetResponseModel) Get() *AlipayIserviceCcmSwArticleGetResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwArticleGetResponseModel) Set(val *AlipayIserviceCcmSwArticleGetResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwArticleGetResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwArticleGetResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwArticleGetResponseModel(val *AlipayIserviceCcmSwArticleGetResponseModel) *NullableAlipayIserviceCcmSwArticleGetResponseModel {
	return &NullableAlipayIserviceCcmSwArticleGetResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwArticleGetResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwArticleGetResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
