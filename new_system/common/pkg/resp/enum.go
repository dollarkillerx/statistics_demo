package resp

// 常量
const (
	SuccessCode             = 200 //成功code
	FailureCode             = 500 //失败code
	UnAuthorized            = 401 //需要授权
	NotFoundCode            = 404 //未找到
	UnprocessableEntityCode = 422 //数据校验失败

	NoBindAddr = 501 //未绑定地址
	NoAuth     = 502 //未实名认证
	NoPayPwd   = 503 //未设置支付密码
	NoNickname = 504 //未设置昵称
)
