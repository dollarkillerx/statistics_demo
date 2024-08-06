package storage

import (
	"github.com/dollarkillerx/backend/pkg/models"
	"github.com/dollarkillerx/backend/pkg/resp"
	"github.com/rs/xid"
)

func (s *Storage) SetError(payload resp.ErrorPayload) {
	s.db.Model(&models.Error{}).Create(&models.Error{
		BaseModel: models.BaseModel{
			ID: xid.New().String(),
		},
		ClientID: payload.ClientID,
		ErrMsg:   payload.ErrMsg,
	})
}
