package verification

import (
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
)

var (
	height    = 96
	width     = 241
	base64Cap *base64Captcha.Captcha
)

func InitVerification(redisConn *redis.Client) {
	// 初始化验证码生成
	dm := base64Captcha.NewDriverMath(height, width, 5, base64Captcha.OptionShowSlimeLine, nil, nil, nil)
	base64Cap = base64Captcha.NewCaptcha(dm, &RedisStore{conn: redisConn})
}

// Generate 生成图片验证码
func Generate() (id, base64Images string, err error) {
	id, content, answer := base64Cap.Driver.GenerateIdQuestionAnswer()
	item, err := base64Cap.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", err
	}

	base64Images = item.EncodeB64string()

	err = base64Cap.Store.Set(id, answer)
	if err != nil {
		return "", "", err
	}

	return id, base64Images, err
}

// Verify 验证图片验证码
func Verify(id, answer string) bool {
	return base64Cap.Verify(id, answer, true)
}
