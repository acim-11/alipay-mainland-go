package alipay

const (
	NotifyTypeTradeStatusSync = "trade_status_sync"
)

// TradeNotification
// Deprecated: use Notification instead.
type TradeNotification Notification

// Notification 通知响应参数 https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.8AmJwg&treeId=203&articleId=105286&docType=1
type Notification struct {
	AuthAppId           string      `json:"auth_app_id"`           // App Id
	NotifyTime          string      `json:"notify_time"`           // 通知时间
	NotifyType          string      `json:"notify_type"`           // 通知类型
	NotifyId            string      `json:"notify_id"`             // 通知校验ID
	AppId               string      `json:"app_id"`                // 开发者的app_id
	Charset             string      `json:"charset"`               // 编码格式
	Version             string      `json:"version"`               // 接口版本
	SignType            string      `json:"sign_type"`             // 签名类型
	Sign                string      `json:"sign"`                  // 签名
	TradeNo             string      `json:"trade_no"`              // 支付宝交易号
	OutTradeNo          string      `json:"out_trade_no"`          // 商户订单号
	OutRequestNo        string      `json:"out_request_no"`        // 退款请求编号
	OutBizNo            string      `json:"out_biz_no"`            // 商户业务号
	BuyerId             string      `json:"buyer_id"`              // 买家支付宝用户号
	BuyerLogonId        string      `json:"buyer_logon_id"`        // 买家支付宝账号
	SellerId            string      `json:"seller_id"`             // 卖家支付宝用户号
	SellerEmail         string      `json:"seller_email"`          // 卖家支付宝账号
	TradeStatus         TradeStatus `json:"trade_status"`          // 交易状态
	RefundStatus        string      `json:"refund_status"`         // 退款状态
	RefundReason        string      `json:"refund_reason"`         // 退款原因
	RefundAmount        string      `json:"refund_amount"`         // 退款金额
	TotalAmount         string      `json:"total_amount"`          // 订单金额
	ReceiptAmount       string      `json:"receipt_amount"`        // 实收金额
	InvoiceAmount       string      `json:"invoice_amount"`        // 开票金额
	BuyerPayAmount      string      `json:"buyer_pay_amount"`      // 付款金额
	PointAmount         string      `json:"point_amount"`          // 集分宝金额
	RefundFee           string      `json:"refund_fee"`            // 总退款金额
	Subject             string      `json:"subject"`               // 商品的标题/交易标题/订单标题/订单关键字等，是请求时对应的参数，原样通知回来。
	Body                string      `json:"body"`                  // 商品描述
	GmtCreate           string      `json:"gmt_create"`            // 交易创建时间
	GmtPayment          string      `json:"gmt_payment"`           // 交易付款时间
	GmtRefund           string      `json:"gmt_refund"`            // 交易退款时间
	GmtClose            string      `json:"gmt_close"`             // 交易结束时间
	FundBillList        string      `json:"fund_bill_list"`        // 支付金额信息
	PassbackParams      string      `json:"passback_params"`       // 回传参数
	VoucherDetailList   string      `json:"voucher_detail_list"`   // 优惠券信息
	AgreementNo         string      `json:"agreement_no"`          // 支付宝签约号
	ExternalAgreementNo string      `json:"external_agreement_no"` // 商户自定义签约号
	DBackStatus         string      `json:"dback_status"`          // 银行卡冲退状态
	DBackAmount         string      `json:"dback_amount"`          // 银行卡冲退金额
	BankAckTime         string      `json:"bank_ack_time"`         // 银行响应时间
}

// FundAuthNotification 预授权通知明细
type FundAuthNotification struct {
	AuthNo                    string `json:"auth_no"`                                // 支付宝资金授权订单号
	NotifyType                string `json:"notify_type"`                            // 通知类型
	OutOrderNo                string `json:"out_order_no"`                           // 商家的资金授权订单号
	OperationID               string `json:"operation_id"`                           // 支付宝的资金操作流水号
	OutRequestNo              string `json:"out_request_no"`                         // 商家资金操作流水号
	OperationType             string `json:"operation_type"`                         // 资金操作类型
	Amount                    string `json:"amount"`                                 // 本次操作冻结的金额
	Status                    string `json:"status"`                                 // 资金预授权明细的状态
	GmtCreate                 string `json:"gmt_create"`                             // 明细创建时间
	GmtTrans                  string `json:"gmt_trans"`                              // 明细处理完成时间
	PayerLogonID              string `json:"payer_logon_id"`                         // 付款方支付宝账号登录号，脱敏
	PayerUserID               string `json:"payer_user_id"`                          // 付款方支付宝账号 UID
	PayeeLogonID              string `json:"payee_logon_id,omitempty"`               // 收款方支付宝账号，脱敏
	PayeeUserID               string `json:"payee_user_id,omitempty"`                // 收款方支付宝账号 UID
	TotalFreezeAmount         string `json:"total_freeze_amount"`                    // 累计冻结金额
	TotalUnfreezeAmount       string `json:"total_unfreeze_amount"`                  // 累计解冻金额
	TotalPayAmount            string `json:"total_pay_amount"`                       // 累计支付金额
	RestAmount                string `json:"rest_amount"`                            // 剩余冻结金额
	CreditAmount              string `json:"credit_amount,omitempty"`                // 本次操作中信用金额
	FundAmount                string `json:"fund_amount,omitempty"`                  // 本次操作中自有资金金额
	TotalFreezeCreditAmount   string `json:"total_freeze_credit_amount,omitempty"`   // 累计冻结信用金额
	TotalFreezeFundAmount     string `json:"total_freeze_fund_amount,omitempty"`     // 累计冻结自有资金金额
	TotalUnfreezeCreditAmount string `json:"total_unfreeze_credit_amount,omitempty"` // 累计解冻信用金额
	TotalUnfreezeFundAmount   string `json:"total_unfreeze_fund_amount,omitempty"`   // 累计解冻自有资金金额
	TotalPayCreditAmount      string `json:"total_pay_credit_amount,omitempty"`      // 累计支付信用金额
	TotalPayFundAmount        string `json:"total_pay_fund_amount,omitempty"`        // 累计支付自有资金金额
	RestCreditAmount          string `json:"rest_credit_amount,omitempty"`           // 剩余冻结信用金额
	RestFundAmount            string `json:"rest_fund_amount,omitempty"`             // 剩余冻结自有资金金额
	PreAuthType               string `json:"pre_auth_type,omitempty"`                // 预授权类型
	CreditMerchantExt         string `json:"credit_merchant_ext,omitempty"`          // 芝麻透出给商家的信息
}
