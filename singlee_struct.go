package singlee_sdk

import (
	"strconv"
	"time"
)

type PayStatus int

type PayMethod int

const (
	PayStatusPaying PayStatus = iota
	PayStatusPayed
	PayStatusFailed
	PayMethodAli   PayMethod = 1 // 支付宝
	PayMethodWx    PayMethod = 2 // 微信
	PayMethodLong  PayMethod = 3 // 龙支付
	PayMethodUnion PayMethod = 4 // 银联
)

// PayResponse 交易响应
type PayResult interface {
	// 交易结果
	GetPayStatus() PayStatus
	// 交易失败原因
	GetFailedMsg() string
	// 接入方商户侧的单号
	GetPayNo() string
	// 银行通道侧的单号
	GetThirdNo() string
	// 微信支付官方侧单号
	GetFirstNo() string
	// 交易支付金额 单位分
	GetAmount() int
	// 交易用户的openid 实际取不到值
	GetOpenid() string
	// 交易时间
	GetPayTime() time.Time
	// 支付方式
	GetPayMethod() PayMethod
}

type PaymentRequest struct {
	ShopId        string `json:"shop_id"`
	PosId         string `json:"pos_id"`
	MchOrderNo    string `json:"mch_order_no"`
	PayQrcode     string `json:"pay_qrcode"`
	Amount        string `json:"amount"`
	MchInfoString string `json:"mch_info_string"`
	Timestamp     string `json:"timestamp"`
	Sign          string `json:"sign"`
}

type OrderQueryResponse struct {
	ResultCode     string `json:"result_code"`
	ErrCode        string `json:"err_code"`
	ErrMsg         string `json:"err_msg"`
	ShopId         string `json:"shop_id"`
	PosId          string `json:"pos_id"`
	MchOrderNo     string `json:"mch_order_no"`
	MchId          string `json:"mch_id"`  // 银行商户号
	TermId         string `json:"term_id"` // 银行实际终端id
	BankOrderNo    string `json:"bank_order_no"`
	ChannelType    string `json:"channel_type"` // 1支付宝 2微信 3 龙支付 4 银联支付
	Refno          string `json:"refno"`
	Posser         string `json:"posser"`
	Batser         string `json:"batser"`
	Amount         string `json:"amount"`
	Actpayamt      string `json:"actpayamt"`
	Couponamt      string `json:"couponamt"`
	Disamt         string `json:"disamt"`
	OriBankBizTime string `json:"ori_bank_biz_time"`
	MchName        string `json:"mch_name"`
	Timestamp      string `json:"timestamp"`
	Sign           string `json:"sign,omitempty"`
	Jylx           string `json:"jylx"`
}

type PaymentResponse struct {
	ResultCode         string `json:"result_code"`
	ErrCode            string `json:"err_code"`
	ErrMsg             string `json:"err_msg"`
	ShopId             string `json:"shop_id"`
	PosId              string `json:"pos_id"`
	MchId              string `json:"mch_id"`
	TermId             string `json:"term_id"`
	MchOrderNo         string `json:"mch_order_no"`
	BankOrderNo        string `json:"bank_order_no"`
	ChannelType        string `json:"channel_type"`
	Refno              string `json:"refno"`
	Posser             string `json:"posser"`
	Batser             string `json:"batser"`
	Amount             string `json:"amount"`
	Actpayamt          string `json:"actpayamt"`
	Disamt             string `json:"disamt"`
	PreferentialDetail string `json:"preferential_detail"`
	GoodsId            string `json:"goods_id"`
	DiscountAmount     string `json:"discount_amount"`
	BankBizTime        string `json:"bank_biz_time"`
	MchName            string `json:"mch_name"`
	Timestamp          string `json:"timestamp"`
	Sign               string `json:"sign"`
}

func (r *OrderQueryResponse) GetPayMethod() PayMethod {
	t, _ := strconv.Atoi(r.ChannelType)
	return PayMethod(t)
}

func (r *PaymentResponse) GetPayMethod() PayMethod {
	t, _ := strconv.Atoi(r.ChannelType)
	return PayMethod(t)
}

func (r *OrderQueryResponse) GetPayTime() time.Time {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("20060102150405", r.Timestamp, loc)
	return t
}

func (r *PaymentResponse) GetPayTime() time.Time {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("20060102150405", r.Timestamp, loc)
	return t
}

func (r *OrderQueryResponse) GetAmount() int {
	amount, _ := strconv.Atoi(r.Amount)
	return amount
}

func (r *PaymentResponse) GetAmount() int {
	amount, _ := strconv.Atoi(r.Amount)
	return amount
}

func (r *OrderQueryResponse) GetOpenid() string {
	return ""
}

func (r *PaymentResponse) GetOpenid() string {
	return ""
}

func (r *OrderQueryResponse) GetFirstNo() string {
	return r.BankOrderNo
}

func (r *PaymentResponse) GetFirstNo() string {
	return r.BankOrderNo
}

func (r *OrderQueryResponse) GetThirdNo() string {
	return r.Posser
}

func (r *PaymentResponse) GetThirdNo() string {
	return r.Posser
}

func (r *OrderQueryResponse) GetPayNo() string {
	return r.MchOrderNo
}

func (r *PaymentResponse) GetPayNo() string {
	return r.MchOrderNo
}

func (r *PaymentResponse) GetPayStatus() PayStatus {
	if r.ErrCode == "000000" && r.ResultCode == "0" {
		return PayStatusPayed
	} else if r.ErrCode == "D23" && r.ResultCode == "9997" {
		return PayStatusPaying
	}
	return PayStatusFailed
}

func (r *PaymentResponse) GetFailedMsg() string {
	if r.GetPayStatus() == PayStatusFailed {
		return r.ErrMsg
	}
	return ""
}

func (r *OrderQueryResponse) GetPayStatus() PayStatus {
	if r.ErrCode == "000000" && r.ResultCode == "0" {
		return PayStatusPayed
	} else if r.ErrCode == "C89005" && r.ResultCode == "9996" {
		return PayStatusPaying
	}
	return PayStatusFailed
}

func (r *OrderQueryResponse) GetFailedMsg() string {
	if r.GetPayStatus() == PayStatusFailed {
		return r.ErrMsg
	}
	return ""
}
