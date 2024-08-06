package api

import (
	"github.com/dollarkillerx/backend/pkg/resp"
	"github.com/gin-gonic/gin"
)

func (a *ApiServer) errors(ctx *gin.Context) {
	var input resp.ErrorPayload
	if err := ctx.ShouldBindJSON(&input); err != nil {
		resp.Return(ctx, 400, err.Error(), nil)
		return
	}

	a.storage.SetError(input)

	resp.Return(ctx, 200, "ok", nil)
}
