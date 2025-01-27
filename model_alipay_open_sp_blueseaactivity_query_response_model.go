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

// checks if the AlipayOpenSpBlueseaactivityQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSpBlueseaactivityQueryResponseModel{}

// AlipayOpenSpBlueseaactivityQueryResponseModel struct for AlipayOpenSpBlueseaactivityQueryResponseModel
type AlipayOpenSpBlueseaactivityQueryResponseModel struct {
	// 详细地址
	Address *string `json:"address,omitempty"`
	// 蓝海活动的场景，包括直连餐饮（BLUE_SEA_FOOD_APPLY）、直连快消（BLUE_SEA_FMCG_APPLY）、间连餐饮（BLUE_SEA_FOOD_INDIRECT_APPLY）、间连快消（BLUE_SEA_FMCG_INDIRECT_APPLY）场景
	BizScene *string `json:"biz_scene,omitempty"`
	// 营业执照，要求证件文本信息清晰可见。 资质信息请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	BusinessLic *string `json:"business_lic,omitempty"`
	// 城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）
	CityCode *string `json:"city_code,omitempty"`
	// 区县编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）
	DistrictCode *string `json:"district_code,omitempty"`
	// 食品经营许可证，要求证件文本信息清晰可见。资质信息请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	FoodBusinessLic *string `json:"food_business_lic,omitempty"`
	// 食品流通许可证，要求证件文本信息清晰可见。资质信息请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	FoodCirculateLic *string `json:"food_circulate_lic,omitempty"`
	// 食品卫生许可证，要求证件文本信息清晰可见。资质信息请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	FoodHealthLic *string `json:"food_health_lic,omitempty"`
	// 食品生产许可证，要求证件文本信息清晰可见。资质信息请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	FoodProductionLic *string `json:"food_production_lic,omitempty"`
	// 餐饮服务许可证，要求证件文本信息清晰可见。资质信息请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	FoodServiceLic *string `json:"food_service_lic,omitempty"`
	// 门头照，要求内景照片清晰可见。资质信息请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	IndoorPic *string `json:"indoor_pic,omitempty"`
	// 审核驳回原因，在订单状态为失败时有效
	Memo *string `json:"memo,omitempty"`
	// 参与蓝海活动的商户支付宝账号，只有当参与直连蓝海活动场景（BLUE_SEA_FOOD_APPLY/BLUE_SEA_FMCG_APPLY）时才返回
	MerchantLogon *string `json:"merchant_logon,omitempty"`
	// 省份编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）
	ProvinceCode *string `json:"province_code,omitempty"`
	// 门头照，要求店铺外观照片清晰可见。资质信息请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	ShopEntrancePic *string `json:"shop_entrance_pic,omitempty"`
	// 申请单状态，状态机参考（AUDITING:审核中，FAIL:报名失败，PASS:报名成功）
	Status *string `json:"status,omitempty"`
	// 参与蓝海活动的间连商户账号，只有当参与间连蓝海活动场景（BLUE_SEA_FOOD_INDIRECT_APPLY/BLUE_SEA_FMCG_INDIRECT_APPLY）时才返回
	SubMerchantId *string `json:"sub_merchant_id,omitempty"`
	// 烟草专卖零售许可证，要求证件文本信息清晰可见。资质信息请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	TobaccoLic *string `json:"tobacco_lic,omitempty"`
}

// NewAlipayOpenSpBlueseaactivityQueryResponseModel instantiates a new AlipayOpenSpBlueseaactivityQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSpBlueseaactivityQueryResponseModel() *AlipayOpenSpBlueseaactivityQueryResponseModel {
	this := AlipayOpenSpBlueseaactivityQueryResponseModel{}
	return &this
}

// NewAlipayOpenSpBlueseaactivityQueryResponseModelWithDefaults instantiates a new AlipayOpenSpBlueseaactivityQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSpBlueseaactivityQueryResponseModelWithDefaults() *AlipayOpenSpBlueseaactivityQueryResponseModel {
	this := AlipayOpenSpBlueseaactivityQueryResponseModel{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetAddress(v string) {
	o.Address = &v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetBusinessLic returns the BusinessLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetBusinessLic() string {
	if o == nil || IsNil(o.BusinessLic) {
		var ret string
		return ret
	}
	return *o.BusinessLic
}

// GetBusinessLicOk returns a tuple with the BusinessLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetBusinessLicOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessLic) {
		return nil, false
	}
	return o.BusinessLic, true
}

// HasBusinessLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasBusinessLic() bool {
	if o != nil && !IsNil(o.BusinessLic) {
		return true
	}

	return false
}

// SetBusinessLic gets a reference to the given string and assigns it to the BusinessLic field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetBusinessLic(v string) {
	o.BusinessLic = &v
}

// GetCityCode returns the CityCode field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetCityCode() string {
	if o == nil || IsNil(o.CityCode) {
		var ret string
		return ret
	}
	return *o.CityCode
}

// GetCityCodeOk returns a tuple with the CityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetCityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CityCode) {
		return nil, false
	}
	return o.CityCode, true
}

// HasCityCode returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasCityCode() bool {
	if o != nil && !IsNil(o.CityCode) {
		return true
	}

	return false
}

// SetCityCode gets a reference to the given string and assigns it to the CityCode field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetCityCode(v string) {
	o.CityCode = &v
}

// GetDistrictCode returns the DistrictCode field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetDistrictCode() string {
	if o == nil || IsNil(o.DistrictCode) {
		var ret string
		return ret
	}
	return *o.DistrictCode
}

// GetDistrictCodeOk returns a tuple with the DistrictCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetDistrictCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DistrictCode) {
		return nil, false
	}
	return o.DistrictCode, true
}

// HasDistrictCode returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasDistrictCode() bool {
	if o != nil && !IsNil(o.DistrictCode) {
		return true
	}

	return false
}

// SetDistrictCode gets a reference to the given string and assigns it to the DistrictCode field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetDistrictCode(v string) {
	o.DistrictCode = &v
}

// GetFoodBusinessLic returns the FoodBusinessLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetFoodBusinessLic() string {
	if o == nil || IsNil(o.FoodBusinessLic) {
		var ret string
		return ret
	}
	return *o.FoodBusinessLic
}

// GetFoodBusinessLicOk returns a tuple with the FoodBusinessLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetFoodBusinessLicOk() (*string, bool) {
	if o == nil || IsNil(o.FoodBusinessLic) {
		return nil, false
	}
	return o.FoodBusinessLic, true
}

// HasFoodBusinessLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasFoodBusinessLic() bool {
	if o != nil && !IsNil(o.FoodBusinessLic) {
		return true
	}

	return false
}

// SetFoodBusinessLic gets a reference to the given string and assigns it to the FoodBusinessLic field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetFoodBusinessLic(v string) {
	o.FoodBusinessLic = &v
}

// GetFoodCirculateLic returns the FoodCirculateLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetFoodCirculateLic() string {
	if o == nil || IsNil(o.FoodCirculateLic) {
		var ret string
		return ret
	}
	return *o.FoodCirculateLic
}

// GetFoodCirculateLicOk returns a tuple with the FoodCirculateLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetFoodCirculateLicOk() (*string, bool) {
	if o == nil || IsNil(o.FoodCirculateLic) {
		return nil, false
	}
	return o.FoodCirculateLic, true
}

// HasFoodCirculateLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasFoodCirculateLic() bool {
	if o != nil && !IsNil(o.FoodCirculateLic) {
		return true
	}

	return false
}

// SetFoodCirculateLic gets a reference to the given string and assigns it to the FoodCirculateLic field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetFoodCirculateLic(v string) {
	o.FoodCirculateLic = &v
}

// GetFoodHealthLic returns the FoodHealthLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetFoodHealthLic() string {
	if o == nil || IsNil(o.FoodHealthLic) {
		var ret string
		return ret
	}
	return *o.FoodHealthLic
}

// GetFoodHealthLicOk returns a tuple with the FoodHealthLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetFoodHealthLicOk() (*string, bool) {
	if o == nil || IsNil(o.FoodHealthLic) {
		return nil, false
	}
	return o.FoodHealthLic, true
}

// HasFoodHealthLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasFoodHealthLic() bool {
	if o != nil && !IsNil(o.FoodHealthLic) {
		return true
	}

	return false
}

// SetFoodHealthLic gets a reference to the given string and assigns it to the FoodHealthLic field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetFoodHealthLic(v string) {
	o.FoodHealthLic = &v
}

// GetFoodProductionLic returns the FoodProductionLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetFoodProductionLic() string {
	if o == nil || IsNil(o.FoodProductionLic) {
		var ret string
		return ret
	}
	return *o.FoodProductionLic
}

// GetFoodProductionLicOk returns a tuple with the FoodProductionLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetFoodProductionLicOk() (*string, bool) {
	if o == nil || IsNil(o.FoodProductionLic) {
		return nil, false
	}
	return o.FoodProductionLic, true
}

// HasFoodProductionLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasFoodProductionLic() bool {
	if o != nil && !IsNil(o.FoodProductionLic) {
		return true
	}

	return false
}

// SetFoodProductionLic gets a reference to the given string and assigns it to the FoodProductionLic field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetFoodProductionLic(v string) {
	o.FoodProductionLic = &v
}

// GetFoodServiceLic returns the FoodServiceLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetFoodServiceLic() string {
	if o == nil || IsNil(o.FoodServiceLic) {
		var ret string
		return ret
	}
	return *o.FoodServiceLic
}

// GetFoodServiceLicOk returns a tuple with the FoodServiceLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetFoodServiceLicOk() (*string, bool) {
	if o == nil || IsNil(o.FoodServiceLic) {
		return nil, false
	}
	return o.FoodServiceLic, true
}

// HasFoodServiceLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasFoodServiceLic() bool {
	if o != nil && !IsNil(o.FoodServiceLic) {
		return true
	}

	return false
}

// SetFoodServiceLic gets a reference to the given string and assigns it to the FoodServiceLic field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetFoodServiceLic(v string) {
	o.FoodServiceLic = &v
}

// GetIndoorPic returns the IndoorPic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetIndoorPic() string {
	if o == nil || IsNil(o.IndoorPic) {
		var ret string
		return ret
	}
	return *o.IndoorPic
}

// GetIndoorPicOk returns a tuple with the IndoorPic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetIndoorPicOk() (*string, bool) {
	if o == nil || IsNil(o.IndoorPic) {
		return nil, false
	}
	return o.IndoorPic, true
}

// HasIndoorPic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasIndoorPic() bool {
	if o != nil && !IsNil(o.IndoorPic) {
		return true
	}

	return false
}

// SetIndoorPic gets a reference to the given string and assigns it to the IndoorPic field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetIndoorPic(v string) {
	o.IndoorPic = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetMemo(v string) {
	o.Memo = &v
}

// GetMerchantLogon returns the MerchantLogon field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetMerchantLogon() string {
	if o == nil || IsNil(o.MerchantLogon) {
		var ret string
		return ret
	}
	return *o.MerchantLogon
}

// GetMerchantLogonOk returns a tuple with the MerchantLogon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetMerchantLogonOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantLogon) {
		return nil, false
	}
	return o.MerchantLogon, true
}

// HasMerchantLogon returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasMerchantLogon() bool {
	if o != nil && !IsNil(o.MerchantLogon) {
		return true
	}

	return false
}

// SetMerchantLogon gets a reference to the given string and assigns it to the MerchantLogon field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetMerchantLogon(v string) {
	o.MerchantLogon = &v
}

// GetProvinceCode returns the ProvinceCode field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetProvinceCode() string {
	if o == nil || IsNil(o.ProvinceCode) {
		var ret string
		return ret
	}
	return *o.ProvinceCode
}

// GetProvinceCodeOk returns a tuple with the ProvinceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetProvinceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProvinceCode) {
		return nil, false
	}
	return o.ProvinceCode, true
}

// HasProvinceCode returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasProvinceCode() bool {
	if o != nil && !IsNil(o.ProvinceCode) {
		return true
	}

	return false
}

// SetProvinceCode gets a reference to the given string and assigns it to the ProvinceCode field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetProvinceCode(v string) {
	o.ProvinceCode = &v
}

// GetShopEntrancePic returns the ShopEntrancePic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetShopEntrancePic() string {
	if o == nil || IsNil(o.ShopEntrancePic) {
		var ret string
		return ret
	}
	return *o.ShopEntrancePic
}

// GetShopEntrancePicOk returns a tuple with the ShopEntrancePic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetShopEntrancePicOk() (*string, bool) {
	if o == nil || IsNil(o.ShopEntrancePic) {
		return nil, false
	}
	return o.ShopEntrancePic, true
}

// HasShopEntrancePic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasShopEntrancePic() bool {
	if o != nil && !IsNil(o.ShopEntrancePic) {
		return true
	}

	return false
}

// SetShopEntrancePic gets a reference to the given string and assigns it to the ShopEntrancePic field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetShopEntrancePic(v string) {
	o.ShopEntrancePic = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetStatus(v string) {
	o.Status = &v
}

// GetSubMerchantId returns the SubMerchantId field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetSubMerchantId() string {
	if o == nil || IsNil(o.SubMerchantId) {
		var ret string
		return ret
	}
	return *o.SubMerchantId
}

// GetSubMerchantIdOk returns a tuple with the SubMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetSubMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubMerchantId) {
		return nil, false
	}
	return o.SubMerchantId, true
}

// HasSubMerchantId returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasSubMerchantId() bool {
	if o != nil && !IsNil(o.SubMerchantId) {
		return true
	}

	return false
}

// SetSubMerchantId gets a reference to the given string and assigns it to the SubMerchantId field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetSubMerchantId(v string) {
	o.SubMerchantId = &v
}

// GetTobaccoLic returns the TobaccoLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetTobaccoLic() string {
	if o == nil || IsNil(o.TobaccoLic) {
		var ret string
		return ret
	}
	return *o.TobaccoLic
}

// GetTobaccoLicOk returns a tuple with the TobaccoLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) GetTobaccoLicOk() (*string, bool) {
	if o == nil || IsNil(o.TobaccoLic) {
		return nil, false
	}
	return o.TobaccoLic, true
}

// HasTobaccoLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) HasTobaccoLic() bool {
	if o != nil && !IsNil(o.TobaccoLic) {
		return true
	}

	return false
}

// SetTobaccoLic gets a reference to the given string and assigns it to the TobaccoLic field.
func (o *AlipayOpenSpBlueseaactivityQueryResponseModel) SetTobaccoLic(v string) {
	o.TobaccoLic = &v
}

func (o AlipayOpenSpBlueseaactivityQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSpBlueseaactivityQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.BusinessLic) {
		toSerialize["business_lic"] = o.BusinessLic
	}
	if !IsNil(o.CityCode) {
		toSerialize["city_code"] = o.CityCode
	}
	if !IsNil(o.DistrictCode) {
		toSerialize["district_code"] = o.DistrictCode
	}
	if !IsNil(o.FoodBusinessLic) {
		toSerialize["food_business_lic"] = o.FoodBusinessLic
	}
	if !IsNil(o.FoodCirculateLic) {
		toSerialize["food_circulate_lic"] = o.FoodCirculateLic
	}
	if !IsNil(o.FoodHealthLic) {
		toSerialize["food_health_lic"] = o.FoodHealthLic
	}
	if !IsNil(o.FoodProductionLic) {
		toSerialize["food_production_lic"] = o.FoodProductionLic
	}
	if !IsNil(o.FoodServiceLic) {
		toSerialize["food_service_lic"] = o.FoodServiceLic
	}
	if !IsNil(o.IndoorPic) {
		toSerialize["indoor_pic"] = o.IndoorPic
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.MerchantLogon) {
		toSerialize["merchant_logon"] = o.MerchantLogon
	}
	if !IsNil(o.ProvinceCode) {
		toSerialize["province_code"] = o.ProvinceCode
	}
	if !IsNil(o.ShopEntrancePic) {
		toSerialize["shop_entrance_pic"] = o.ShopEntrancePic
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubMerchantId) {
		toSerialize["sub_merchant_id"] = o.SubMerchantId
	}
	if !IsNil(o.TobaccoLic) {
		toSerialize["tobacco_lic"] = o.TobaccoLic
	}
	return toSerialize, nil
}

type NullableAlipayOpenSpBlueseaactivityQueryResponseModel struct {
	value *AlipayOpenSpBlueseaactivityQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenSpBlueseaactivityQueryResponseModel) Get() *AlipayOpenSpBlueseaactivityQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSpBlueseaactivityQueryResponseModel) Set(val *AlipayOpenSpBlueseaactivityQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpBlueseaactivityQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpBlueseaactivityQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpBlueseaactivityQueryResponseModel(val *AlipayOpenSpBlueseaactivityQueryResponseModel) *NullableAlipayOpenSpBlueseaactivityQueryResponseModel {
	return &NullableAlipayOpenSpBlueseaactivityQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSpBlueseaactivityQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpBlueseaactivityQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
