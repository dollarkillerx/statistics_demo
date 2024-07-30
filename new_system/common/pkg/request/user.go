package request

import (
	"github.com/pkg/errors"
	"strings"
)

type RegisterPayload struct {
	Mobile     string `json:"mobile"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	VerifyId   string `json:"verify_id"`
	VerifyCode string `json:"verify_code"`
	InviteCode string `json:"invite_code"`
}

func (r *RegisterPayload) Validate() error {
	r.Mobile = strings.TrimSpace(r.Mobile)
	r.Password = strings.TrimSpace(r.Password)
	r.RePassword = strings.TrimSpace(r.RePassword)
	r.VerifyCode = strings.TrimSpace(r.VerifyCode)
	r.InviteCode = strings.TrimSpace(r.InviteCode)

	if r.Mobile == "" || r.Password == "" || r.RePassword == "" || r.VerifyCode == "" {
		return errors.New("invalid_input")
	}

	if len(r.Password) < 8 || len(r.Password) > 22 {
		return errors.New("user_pwd_range")
	}

	if r.Password != r.RePassword {
		return errors.New("user_pwd_nts")
	}

	return nil
}
