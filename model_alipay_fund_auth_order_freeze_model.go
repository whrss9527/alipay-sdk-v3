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

// checks if the AlipayFundAuthOrderFreezeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundAuthOrderFreezeModel{}

// AlipayFundAuthOrderFreezeModel struct for AlipayFundAuthOrderFreezeModel
type AlipayFundAuthOrderFreezeModel struct {
	// 需要冻结的金额，单位为：元（人民币），精确到小数点后两位。 取值范围：[0.01,100000000.00]
	Amount *string `json:"amount,omitempty"`
	// 用户付款码。 1.条码场景：25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准； 2.刷脸场景： 1）fp开头的35位字符串； 2）300-700字符的随机字符串； 注：刷脸场景考虑到未来可能拓展更多格式，建议外围不必做规则拦截，由支付宝统一做有效性校验
	AuthCode *string `json:"auth_code,omitempty"`
	// 付款码类型。 1.条码场景：bar_code 2.刷脸场景：security_code
	AuthCodeType *string `json:"auth_code_type,omitempty"`
	// 业务参数，如风控参数outRiskInfo等。
	BusinessParams *string `json:"business_params,omitempty"`
	// 免押受理台模式，使用免押产品必传该字段。根据免押不同业务模式将开通受理台区分三种模式，商家可根据调用预授权冻结接口传入的参数决定该笔免押订单选择哪种受理台模式。不同受理台模式需要传入不同参数，其中：POSTPAY 表示后付金额已知，POSTPAY_UNCERTAIN 表示后付金额未知，DEPOSIT_ONLY 表示纯免押。 具体规则参考文档：<a href=\"https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545\">https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545</a>
	DepositProductMode *string `json:"deposit_product_mode,omitempty"`
	// 无特殊需要请勿传入；商户可用该参数禁用支付渠道。 传入后用户不可使用列表中的渠道进行支付，目前支持两种禁用渠道：信用卡快捷（OPTIMIZED_MOTO）、信用卡卡通（BIGAMOUNT_CREDIT_CARTOON）。与可用支付渠道不能同时传入
	DisablePayChannels *string `json:"disable_pay_channels,omitempty"`
	// 无特殊需要请勿传入；商户可用该参数指定支付渠道。 传入后用户仅能使用列表中的渠道进行支付，目前支持三种渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。与禁用支付渠道不可同时传入
	EnablePayChannels *string `json:"enable_pay_channels,omitempty"`
	// 业务扩展参数，用于特定业务信息的传递，json格式。 1、category，信用类目，信用预授权场景必传，具体类目信息见<a href=\"https://opendocs.alipay.com/open/10719\">https://opendocs.alipay.com/open/10719</a>； 2、serviceId，信用服务ID：信用预授权场景必传。需要商家在 <a href=\"https://b.alipay.com/page/zmgaia/pre-auth/index\">开放平台-芝麻免押-信用服务管理</a> 创建信用服务获取，详情可查看 <a href=\"https://opendocs.alipay.com/open/03w0a7?pathHash=51f6b4f2&ref=api#%E7%94%B3%E8%AF%B7%E4%BF%A1%E7%94%A8%E6%9C%8D%E5%8A%A1\">创建信用服务</a>。在创建过程中如果有其它疑问，可以咨询芝麻客服小二（0571-88158055 转 2）； 3、creditExtInfo，信用参数，可选，如有需要请与芝麻约定后传入，信用服务说明见<a href=\"https://opendocs.alipay.com/open/11157/qlsxya\">https://opendocs.alipay.com/open/11157/qlsxya</a>
	ExtraParam *string `json:"extra_param,omitempty"`
	// 用户实名信息参数，包含：姓名+身份证号的hash值、指定用户的uid。商户传入用户实名信息参数，支付宝会对比用户在支付宝端的实名信息。 姓名+身份证号hash值使用SHA256摘要方式与UTF8编码,返回十六进制的字符串。 identity_hash和alipay_user_id都是可选的，如果两个都传，则会先校验identity_hash，然后校验alipay_user_id。其中identity_hash的待加密字样如\"张三4566498798498498498498\"
	IdentityParams *string `json:"identity_params,omitempty"`
	// 通知地址
	NotifyUrl *string `json:"notify_url,omitempty"`
	// 订单标题。 业务订单的简单描述，如商品名称等
	OrderTitle *string `json:"order_title,omitempty"`
	// 商户授权资金订单号。 商家自定义需保证在商户端不重复。仅支持字母、数字、下划线。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// 商户本次资金操作的请求流水号，用于标示请求流水的唯一性。 可与out_order_no相同，仅支持字母、数字、下划线。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 预授权订单相对超时时间，从商户请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。
	PayTimeout *string `json:"pay_timeout,omitempty"`
	// 收款账户的支付宝登录号（email或手机号）。 如果传入则会校验该登录号对应的账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一
	PayeeLogonId *string `json:"payee_logon_id,omitempty"`
	// 收款账户的支付宝用户号。 以2088开头的16位纯数字，如果传入则会校验该账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一
	PayeeUserId *string `json:"payee_user_id,omitempty"`
	// 后付费项目， 有付费项目时需要传入该字段。不同受理台模式需要传入不同参数，后付费项目名称和计费说明需要通过校验规则，同时计费说明将展示在开通受理台上。当受理台模式（deposit_product_mode）传入POSTPAY 时，后付费项目名称（name）、金额（amount）必传，计费说明（description）选传；当传入 POSTPAY_UNCERTAIN 时，后付费项目名称（name）、计费说明（description）必传，金额（amount）不传。 具体规则参考文档： <a href=\"https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545\">https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545</a>
	PostPayments []PostPayment `json:"post_payments,omitempty"`
	// 销售产品码。 当面资金预授权固定为 PRE_AUTH
	ProductCode *string `json:"product_code,omitempty"`
	// 场景码。 刷脸场景下传入HOTEL，其他情况下无需传入
	SceneCode *string `json:"scene_code,omitempty"`
	// 商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP
	SettleCurrency *string `json:"settle_currency,omitempty"`
	// 机具管控sdk加签参数，参数示例 \"terminal_params\":\"{\"terminalType\":\"IOT\",\"signature\":\"QIIAX8DqbFbNf2oe97FI1RSLAycC/tU4GVjer3bN8K4qLtAB\",\"apdidToken\":\"xPA3ptuArwYc3F6Va_pjVwv7Qx7Tg5TJdrA_Jb_moYte9AqGZgEAAA==\",\"hardToken\":\"\",\"time\":\"1539847253\",\"bizCode\":\"11000200040004000121\",\"bizTid\":\"010100F01i1XyacMgpOinHerfdBw1xA9dNDocctlnqhLD8lfODr1A7Q\",\"signedKeys\":\"authCode,totalAmount,apdidToken,hardToken,time,bizCode,bizTid\"}\"
	TerminalParams *string `json:"terminal_params,omitempty"`
	// 预授权订单相对超时时间，从商户请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。
	TimeoutExpress *string `json:"timeout_express,omitempty"`
	// 标价币种,  amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP
	TransCurrency *string `json:"trans_currency,omitempty"`
}

// NewAlipayFundAuthOrderFreezeModel instantiates a new AlipayFundAuthOrderFreezeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundAuthOrderFreezeModel() *AlipayFundAuthOrderFreezeModel {
	this := AlipayFundAuthOrderFreezeModel{}
	return &this
}

// NewAlipayFundAuthOrderFreezeModelWithDefaults instantiates a new AlipayFundAuthOrderFreezeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundAuthOrderFreezeModelWithDefaults() *AlipayFundAuthOrderFreezeModel {
	this := AlipayFundAuthOrderFreezeModel{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AlipayFundAuthOrderFreezeModel) SetAmount(v string) {
	o.Amount = &v
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetAuthCode() string {
	if o == nil || IsNil(o.AuthCode) {
		var ret string
		return ret
	}
	return *o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetAuthCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthCode) {
		return nil, false
	}
	return o.AuthCode, true
}

// HasAuthCode returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasAuthCode() bool {
	if o != nil && !IsNil(o.AuthCode) {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given string and assigns it to the AuthCode field.
func (o *AlipayFundAuthOrderFreezeModel) SetAuthCode(v string) {
	o.AuthCode = &v
}

// GetAuthCodeType returns the AuthCodeType field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetAuthCodeType() string {
	if o == nil || IsNil(o.AuthCodeType) {
		var ret string
		return ret
	}
	return *o.AuthCodeType
}

// GetAuthCodeTypeOk returns a tuple with the AuthCodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetAuthCodeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthCodeType) {
		return nil, false
	}
	return o.AuthCodeType, true
}

// HasAuthCodeType returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasAuthCodeType() bool {
	if o != nil && !IsNil(o.AuthCodeType) {
		return true
	}

	return false
}

// SetAuthCodeType gets a reference to the given string and assigns it to the AuthCodeType field.
func (o *AlipayFundAuthOrderFreezeModel) SetAuthCodeType(v string) {
	o.AuthCodeType = &v
}

// GetBusinessParams returns the BusinessParams field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetBusinessParams() string {
	if o == nil || IsNil(o.BusinessParams) {
		var ret string
		return ret
	}
	return *o.BusinessParams
}

// GetBusinessParamsOk returns a tuple with the BusinessParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetBusinessParamsOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessParams) {
		return nil, false
	}
	return o.BusinessParams, true
}

// HasBusinessParams returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasBusinessParams() bool {
	if o != nil && !IsNil(o.BusinessParams) {
		return true
	}

	return false
}

// SetBusinessParams gets a reference to the given string and assigns it to the BusinessParams field.
func (o *AlipayFundAuthOrderFreezeModel) SetBusinessParams(v string) {
	o.BusinessParams = &v
}

// GetDepositProductMode returns the DepositProductMode field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetDepositProductMode() string {
	if o == nil || IsNil(o.DepositProductMode) {
		var ret string
		return ret
	}
	return *o.DepositProductMode
}

// GetDepositProductModeOk returns a tuple with the DepositProductMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetDepositProductModeOk() (*string, bool) {
	if o == nil || IsNil(o.DepositProductMode) {
		return nil, false
	}
	return o.DepositProductMode, true
}

// HasDepositProductMode returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasDepositProductMode() bool {
	if o != nil && !IsNil(o.DepositProductMode) {
		return true
	}

	return false
}

// SetDepositProductMode gets a reference to the given string and assigns it to the DepositProductMode field.
func (o *AlipayFundAuthOrderFreezeModel) SetDepositProductMode(v string) {
	o.DepositProductMode = &v
}

// GetDisablePayChannels returns the DisablePayChannels field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetDisablePayChannels() string {
	if o == nil || IsNil(o.DisablePayChannels) {
		var ret string
		return ret
	}
	return *o.DisablePayChannels
}

// GetDisablePayChannelsOk returns a tuple with the DisablePayChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetDisablePayChannelsOk() (*string, bool) {
	if o == nil || IsNil(o.DisablePayChannels) {
		return nil, false
	}
	return o.DisablePayChannels, true
}

// HasDisablePayChannels returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasDisablePayChannels() bool {
	if o != nil && !IsNil(o.DisablePayChannels) {
		return true
	}

	return false
}

// SetDisablePayChannels gets a reference to the given string and assigns it to the DisablePayChannels field.
func (o *AlipayFundAuthOrderFreezeModel) SetDisablePayChannels(v string) {
	o.DisablePayChannels = &v
}

// GetEnablePayChannels returns the EnablePayChannels field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetEnablePayChannels() string {
	if o == nil || IsNil(o.EnablePayChannels) {
		var ret string
		return ret
	}
	return *o.EnablePayChannels
}

// GetEnablePayChannelsOk returns a tuple with the EnablePayChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetEnablePayChannelsOk() (*string, bool) {
	if o == nil || IsNil(o.EnablePayChannels) {
		return nil, false
	}
	return o.EnablePayChannels, true
}

// HasEnablePayChannels returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasEnablePayChannels() bool {
	if o != nil && !IsNil(o.EnablePayChannels) {
		return true
	}

	return false
}

// SetEnablePayChannels gets a reference to the given string and assigns it to the EnablePayChannels field.
func (o *AlipayFundAuthOrderFreezeModel) SetEnablePayChannels(v string) {
	o.EnablePayChannels = &v
}

// GetExtraParam returns the ExtraParam field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetExtraParam() string {
	if o == nil || IsNil(o.ExtraParam) {
		var ret string
		return ret
	}
	return *o.ExtraParam
}

// GetExtraParamOk returns a tuple with the ExtraParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetExtraParamOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraParam) {
		return nil, false
	}
	return o.ExtraParam, true
}

// HasExtraParam returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasExtraParam() bool {
	if o != nil && !IsNil(o.ExtraParam) {
		return true
	}

	return false
}

// SetExtraParam gets a reference to the given string and assigns it to the ExtraParam field.
func (o *AlipayFundAuthOrderFreezeModel) SetExtraParam(v string) {
	o.ExtraParam = &v
}

// GetIdentityParams returns the IdentityParams field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetIdentityParams() string {
	if o == nil || IsNil(o.IdentityParams) {
		var ret string
		return ret
	}
	return *o.IdentityParams
}

// GetIdentityParamsOk returns a tuple with the IdentityParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetIdentityParamsOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityParams) {
		return nil, false
	}
	return o.IdentityParams, true
}

// HasIdentityParams returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasIdentityParams() bool {
	if o != nil && !IsNil(o.IdentityParams) {
		return true
	}

	return false
}

// SetIdentityParams gets a reference to the given string and assigns it to the IdentityParams field.
func (o *AlipayFundAuthOrderFreezeModel) SetIdentityParams(v string) {
	o.IdentityParams = &v
}

// GetNotifyUrl returns the NotifyUrl field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetNotifyUrl() string {
	if o == nil || IsNil(o.NotifyUrl) {
		var ret string
		return ret
	}
	return *o.NotifyUrl
}

// GetNotifyUrlOk returns a tuple with the NotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetNotifyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyUrl) {
		return nil, false
	}
	return o.NotifyUrl, true
}

// HasNotifyUrl returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasNotifyUrl() bool {
	if o != nil && !IsNil(o.NotifyUrl) {
		return true
	}

	return false
}

// SetNotifyUrl gets a reference to the given string and assigns it to the NotifyUrl field.
func (o *AlipayFundAuthOrderFreezeModel) SetNotifyUrl(v string) {
	o.NotifyUrl = &v
}

// GetOrderTitle returns the OrderTitle field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetOrderTitle() string {
	if o == nil || IsNil(o.OrderTitle) {
		var ret string
		return ret
	}
	return *o.OrderTitle
}

// GetOrderTitleOk returns a tuple with the OrderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetOrderTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OrderTitle) {
		return nil, false
	}
	return o.OrderTitle, true
}

// HasOrderTitle returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasOrderTitle() bool {
	if o != nil && !IsNil(o.OrderTitle) {
		return true
	}

	return false
}

// SetOrderTitle gets a reference to the given string and assigns it to the OrderTitle field.
func (o *AlipayFundAuthOrderFreezeModel) SetOrderTitle(v string) {
	o.OrderTitle = &v
}

// GetOutOrderNo returns the OutOrderNo field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetOutOrderNo() string {
	if o == nil || IsNil(o.OutOrderNo) {
		var ret string
		return ret
	}
	return *o.OutOrderNo
}

// GetOutOrderNoOk returns a tuple with the OutOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetOutOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutOrderNo) {
		return nil, false
	}
	return o.OutOrderNo, true
}

// HasOutOrderNo returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasOutOrderNo() bool {
	if o != nil && !IsNil(o.OutOrderNo) {
		return true
	}

	return false
}

// SetOutOrderNo gets a reference to the given string and assigns it to the OutOrderNo field.
func (o *AlipayFundAuthOrderFreezeModel) SetOutOrderNo(v string) {
	o.OutOrderNo = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayFundAuthOrderFreezeModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetPayTimeout returns the PayTimeout field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetPayTimeout() string {
	if o == nil || IsNil(o.PayTimeout) {
		var ret string
		return ret
	}
	return *o.PayTimeout
}

// GetPayTimeoutOk returns a tuple with the PayTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetPayTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.PayTimeout) {
		return nil, false
	}
	return o.PayTimeout, true
}

// HasPayTimeout returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasPayTimeout() bool {
	if o != nil && !IsNil(o.PayTimeout) {
		return true
	}

	return false
}

// SetPayTimeout gets a reference to the given string and assigns it to the PayTimeout field.
func (o *AlipayFundAuthOrderFreezeModel) SetPayTimeout(v string) {
	o.PayTimeout = &v
}

// GetPayeeLogonId returns the PayeeLogonId field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetPayeeLogonId() string {
	if o == nil || IsNil(o.PayeeLogonId) {
		var ret string
		return ret
	}
	return *o.PayeeLogonId
}

// GetPayeeLogonIdOk returns a tuple with the PayeeLogonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetPayeeLogonIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeLogonId) {
		return nil, false
	}
	return o.PayeeLogonId, true
}

// HasPayeeLogonId returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasPayeeLogonId() bool {
	if o != nil && !IsNil(o.PayeeLogonId) {
		return true
	}

	return false
}

// SetPayeeLogonId gets a reference to the given string and assigns it to the PayeeLogonId field.
func (o *AlipayFundAuthOrderFreezeModel) SetPayeeLogonId(v string) {
	o.PayeeLogonId = &v
}

// GetPayeeUserId returns the PayeeUserId field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetPayeeUserId() string {
	if o == nil || IsNil(o.PayeeUserId) {
		var ret string
		return ret
	}
	return *o.PayeeUserId
}

// GetPayeeUserIdOk returns a tuple with the PayeeUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetPayeeUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeUserId) {
		return nil, false
	}
	return o.PayeeUserId, true
}

// HasPayeeUserId returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasPayeeUserId() bool {
	if o != nil && !IsNil(o.PayeeUserId) {
		return true
	}

	return false
}

// SetPayeeUserId gets a reference to the given string and assigns it to the PayeeUserId field.
func (o *AlipayFundAuthOrderFreezeModel) SetPayeeUserId(v string) {
	o.PayeeUserId = &v
}

// GetPostPayments returns the PostPayments field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetPostPayments() []PostPayment {
	if o == nil || IsNil(o.PostPayments) {
		var ret []PostPayment
		return ret
	}
	return o.PostPayments
}

// GetPostPaymentsOk returns a tuple with the PostPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetPostPaymentsOk() ([]PostPayment, bool) {
	if o == nil || IsNil(o.PostPayments) {
		return nil, false
	}
	return o.PostPayments, true
}

// HasPostPayments returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasPostPayments() bool {
	if o != nil && !IsNil(o.PostPayments) {
		return true
	}

	return false
}

// SetPostPayments gets a reference to the given []PostPayment and assigns it to the PostPayments field.
func (o *AlipayFundAuthOrderFreezeModel) SetPostPayments(v []PostPayment) {
	o.PostPayments = v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundAuthOrderFreezeModel) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetSceneCode returns the SceneCode field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetSceneCode() string {
	if o == nil || IsNil(o.SceneCode) {
		var ret string
		return ret
	}
	return *o.SceneCode
}

// GetSceneCodeOk returns a tuple with the SceneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetSceneCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SceneCode) {
		return nil, false
	}
	return o.SceneCode, true
}

// HasSceneCode returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasSceneCode() bool {
	if o != nil && !IsNil(o.SceneCode) {
		return true
	}

	return false
}

// SetSceneCode gets a reference to the given string and assigns it to the SceneCode field.
func (o *AlipayFundAuthOrderFreezeModel) SetSceneCode(v string) {
	o.SceneCode = &v
}

// GetSettleCurrency returns the SettleCurrency field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetSettleCurrency() string {
	if o == nil || IsNil(o.SettleCurrency) {
		var ret string
		return ret
	}
	return *o.SettleCurrency
}

// GetSettleCurrencyOk returns a tuple with the SettleCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetSettleCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.SettleCurrency) {
		return nil, false
	}
	return o.SettleCurrency, true
}

// HasSettleCurrency returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasSettleCurrency() bool {
	if o != nil && !IsNil(o.SettleCurrency) {
		return true
	}

	return false
}

// SetSettleCurrency gets a reference to the given string and assigns it to the SettleCurrency field.
func (o *AlipayFundAuthOrderFreezeModel) SetSettleCurrency(v string) {
	o.SettleCurrency = &v
}

// GetTerminalParams returns the TerminalParams field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetTerminalParams() string {
	if o == nil || IsNil(o.TerminalParams) {
		var ret string
		return ret
	}
	return *o.TerminalParams
}

// GetTerminalParamsOk returns a tuple with the TerminalParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetTerminalParamsOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalParams) {
		return nil, false
	}
	return o.TerminalParams, true
}

// HasTerminalParams returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasTerminalParams() bool {
	if o != nil && !IsNil(o.TerminalParams) {
		return true
	}

	return false
}

// SetTerminalParams gets a reference to the given string and assigns it to the TerminalParams field.
func (o *AlipayFundAuthOrderFreezeModel) SetTerminalParams(v string) {
	o.TerminalParams = &v
}

// GetTimeoutExpress returns the TimeoutExpress field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetTimeoutExpress() string {
	if o == nil || IsNil(o.TimeoutExpress) {
		var ret string
		return ret
	}
	return *o.TimeoutExpress
}

// GetTimeoutExpressOk returns a tuple with the TimeoutExpress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetTimeoutExpressOk() (*string, bool) {
	if o == nil || IsNil(o.TimeoutExpress) {
		return nil, false
	}
	return o.TimeoutExpress, true
}

// HasTimeoutExpress returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasTimeoutExpress() bool {
	if o != nil && !IsNil(o.TimeoutExpress) {
		return true
	}

	return false
}

// SetTimeoutExpress gets a reference to the given string and assigns it to the TimeoutExpress field.
func (o *AlipayFundAuthOrderFreezeModel) SetTimeoutExpress(v string) {
	o.TimeoutExpress = &v
}

// GetTransCurrency returns the TransCurrency field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderFreezeModel) GetTransCurrency() string {
	if o == nil || IsNil(o.TransCurrency) {
		var ret string
		return ret
	}
	return *o.TransCurrency
}

// GetTransCurrencyOk returns a tuple with the TransCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderFreezeModel) GetTransCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.TransCurrency) {
		return nil, false
	}
	return o.TransCurrency, true
}

// HasTransCurrency returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderFreezeModel) HasTransCurrency() bool {
	if o != nil && !IsNil(o.TransCurrency) {
		return true
	}

	return false
}

// SetTransCurrency gets a reference to the given string and assigns it to the TransCurrency field.
func (o *AlipayFundAuthOrderFreezeModel) SetTransCurrency(v string) {
	o.TransCurrency = &v
}

func (o AlipayFundAuthOrderFreezeModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundAuthOrderFreezeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AuthCode) {
		toSerialize["auth_code"] = o.AuthCode
	}
	if !IsNil(o.AuthCodeType) {
		toSerialize["auth_code_type"] = o.AuthCodeType
	}
	if !IsNil(o.BusinessParams) {
		toSerialize["business_params"] = o.BusinessParams
	}
	if !IsNil(o.DepositProductMode) {
		toSerialize["deposit_product_mode"] = o.DepositProductMode
	}
	if !IsNil(o.DisablePayChannels) {
		toSerialize["disable_pay_channels"] = o.DisablePayChannels
	}
	if !IsNil(o.EnablePayChannels) {
		toSerialize["enable_pay_channels"] = o.EnablePayChannels
	}
	if !IsNil(o.ExtraParam) {
		toSerialize["extra_param"] = o.ExtraParam
	}
	if !IsNil(o.IdentityParams) {
		toSerialize["identity_params"] = o.IdentityParams
	}
	if !IsNil(o.NotifyUrl) {
		toSerialize["notify_url"] = o.NotifyUrl
	}
	if !IsNil(o.OrderTitle) {
		toSerialize["order_title"] = o.OrderTitle
	}
	if !IsNil(o.OutOrderNo) {
		toSerialize["out_order_no"] = o.OutOrderNo
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.PayTimeout) {
		toSerialize["pay_timeout"] = o.PayTimeout
	}
	if !IsNil(o.PayeeLogonId) {
		toSerialize["payee_logon_id"] = o.PayeeLogonId
	}
	if !IsNil(o.PayeeUserId) {
		toSerialize["payee_user_id"] = o.PayeeUserId
	}
	if !IsNil(o.PostPayments) {
		toSerialize["post_payments"] = o.PostPayments
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	if !IsNil(o.SceneCode) {
		toSerialize["scene_code"] = o.SceneCode
	}
	if !IsNil(o.SettleCurrency) {
		toSerialize["settle_currency"] = o.SettleCurrency
	}
	if !IsNil(o.TerminalParams) {
		toSerialize["terminal_params"] = o.TerminalParams
	}
	if !IsNil(o.TimeoutExpress) {
		toSerialize["timeout_express"] = o.TimeoutExpress
	}
	if !IsNil(o.TransCurrency) {
		toSerialize["trans_currency"] = o.TransCurrency
	}
	return toSerialize, nil
}

type NullableAlipayFundAuthOrderFreezeModel struct {
	value *AlipayFundAuthOrderFreezeModel
	isSet bool
}

func (v NullableAlipayFundAuthOrderFreezeModel) Get() *AlipayFundAuthOrderFreezeModel {
	return v.value
}

func (v *NullableAlipayFundAuthOrderFreezeModel) Set(val *AlipayFundAuthOrderFreezeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAuthOrderFreezeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAuthOrderFreezeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAuthOrderFreezeModel(val *AlipayFundAuthOrderFreezeModel) *NullableAlipayFundAuthOrderFreezeModel {
	return &NullableAlipayFundAuthOrderFreezeModel{value: val, isSet: true}
}

func (v NullableAlipayFundAuthOrderFreezeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAuthOrderFreezeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
