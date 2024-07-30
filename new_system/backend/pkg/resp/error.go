package resp

type ErrorPayload struct {
	ClientID string `json:"client_id"` // company.account: exness.10086
	ErrMsg   string `json:"err_msg"`   // error message
}
