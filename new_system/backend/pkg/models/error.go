package models

type ErrorPayload struct {
	BaseModel
	ClientID string `json:"client_id"` // company.account: exness.10086
	ErrMsg   string `json:"err_msg"`   // error message
}

func (e *ErrorPayload) TableName() string {
	return "errors"
}
