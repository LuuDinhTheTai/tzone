package route

import (
	"github.com/LuuDinhTheTai/tzone/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func MapCommonRoutes(r *gin.Engine, commonHandler *handler.CommonHandler) {
	r.GET("/", commonHandler.IndexHandler)
}
