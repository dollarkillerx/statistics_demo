package api

import (
	"github.com/dollarkillerx/backend/pkg/resp"
	"github.com/gin-gonic/gin"

	"strings"
)

func (a *ApiServer) accounts(ctx *gin.Context) {
	positions := a.storage.GetAccounts()
	resp.Return(ctx, 200, "ok", positions)
}

func (a *ApiServer) account(ctx *gin.Context) {
	param := strings.TrimSpace(ctx.Param("account"))

	if param == "" {
		resp.Return(ctx, 200, "key is null", nil)
		return
	}

}
