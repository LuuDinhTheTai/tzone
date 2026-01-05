package handler

import (
	"github.com/LuuDinhTheTai/tzone/util/response"
	"github.com/LuuDinhTheTai/tzone/web/page"
	"github.com/gin-gonic/gin"
)

type CommonHandler struct {
}

func NewCommonHandler() *CommonHandler {
	return &CommonHandler{}
}

func (h *CommonHandler) IndexHandler(ctx *gin.Context) {
	response.HTML(ctx, page.HomePage())
}
