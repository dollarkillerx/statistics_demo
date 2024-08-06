package models

type Error struct {
	BaseModel
	ClientID string `json:"client_id"` // company.account: exness.10086
	ErrMsg   string `json:"err_msg"`   // error message
}

func (e *Error) TableName() string {
	return "errors"
}
