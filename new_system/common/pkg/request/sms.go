package request

type SmsPayload struct {
	// Type "1:注册 2:找回密码 3:提币 4:转账 5:修改支付密码 6:修改登录密码 7：增加钱包地址"
	Type   int64  `json:"type"`
	Mobile string `json:"mobile"`
	CapId  string `json:"cap_id"` // 验证码ID
	Answer string `json:"answer"` // 验证码答案
}
