package route

import (
	"github.com/LuuDinhTheTai/tzone/internal/delivery/handler"
	"github.com/LuuDinhTheTai/tzone/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
)

func MapReviewRoutes(r *gin.Engine, reviewHandler *handler.ReviewHandler) {
	reviewGroup := r.Group("/api/v1/reviews")
	reviewGroup.Use(middleware.APIRateLimit())
	{
		reviewGroup.GET("/device/:deviceId", reviewHandler.GetByDeviceID)

		authReviewGroup := reviewGroup.Group("")
		authReviewGroup.Use(middleware.JWTAuth())
		{
			authReviewGroup.POST("/device/:deviceId/rating", reviewHandler.SetRating)
			authReviewGroup.POST("/device/:deviceId/comment", reviewHandler.SetComment)
			authReviewGroup.PUT("/:id/comment", reviewHandler.UpdateComment)
			authReviewGroup.DELETE("/:id", reviewHandler.Delete)
		}
	}
}
