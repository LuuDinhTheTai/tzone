package handler

import (
	"net/http"
	"strings"

	"github.com/LuuDinhTheTai/tzone/internal/dto"
	"github.com/LuuDinhTheTai/tzone/internal/service"
	"github.com/LuuDinhTheTai/tzone/util/response"
	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewService *service.ReviewService
}

func NewReviewHandler(reviewService *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{reviewService: reviewService}
}

func (h *ReviewHandler) GetByDeviceID(c *gin.Context) {
	deviceID := strings.TrimSpace(c.Param("deviceId"))
	if deviceID == "" {
		response.Error(c, http.StatusBadRequest, "deviceId is required", nil)
		return
	}

	var query dto.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid pagination query", []response.ErrorResponse{{Field: "query", Error: err.Error()}})
		return
	}
	query.Normalize()

	result, err := h.reviewService.GetByDeviceID(deviceID, query.Page, query.Limit)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "reviews retrieved successfully", result)
}

func (h *ReviewHandler) SetRating(c *gin.Context) {
	userID, ok := getAuthUserID(c)
	if !ok {
		return
	}

	deviceID := strings.TrimSpace(c.Param("deviceId"))
	if deviceID == "" {
		response.Error(c, http.StatusBadRequest, "deviceId is required", nil)
		return
	}

	var req dto.SetRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid rating request", []response.ErrorResponse{{Field: "request", Error: err.Error()}})
		return
	}

	result, err := h.reviewService.SetRating(userID, deviceID, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "rating saved successfully", result)
}

func (h *ReviewHandler) SetComment(c *gin.Context) {
	userID, ok := getAuthUserID(c)
	if !ok {
		return
	}

	deviceID := strings.TrimSpace(c.Param("deviceId"))
	if deviceID == "" {
		response.Error(c, http.StatusBadRequest, "deviceId is required", nil)
		return
	}

	var req dto.SetCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid comment request", []response.ErrorResponse{{Field: "request", Error: err.Error()}})
		return
	}

	result, err := h.reviewService.SetComment(userID, deviceID, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "comment saved successfully", result)
}

func (h *ReviewHandler) UpdateComment(c *gin.Context) {
	userID, ok := getAuthUserID(c)
	if !ok {
		return
	}

	reviewID := strings.TrimSpace(c.Param("id"))
	if reviewID == "" {
		response.Error(c, http.StatusBadRequest, "review id is required", nil)
		return
	}

	var req dto.UpdateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid comment update request", []response.ErrorResponse{{Field: "request", Error: err.Error()}})
		return
	}

	result, err := h.reviewService.UpdateComment(userID, reviewID, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "comment updated successfully", result)
}

func (h *ReviewHandler) Create(c *gin.Context) {
	userID, ok := getAuthUserID(c)
	if !ok {
		return
	}

	deviceID := strings.TrimSpace(c.Param("deviceId"))
	if deviceID == "" {
		response.Error(c, http.StatusBadRequest, "deviceId is required", nil)
		return
	}

	var req dto.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid create review request", []response.ErrorResponse{{Field: "request", Error: err.Error()}})
		return
	}

	result, err := h.reviewService.Create(userID, deviceID, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "review created successfully", result)
}

func (h *ReviewHandler) Update(c *gin.Context) {
	userID, ok := getAuthUserID(c)
	if !ok {
		return
	}

	reviewID := strings.TrimSpace(c.Param("id"))
	if reviewID == "" {
		response.Error(c, http.StatusBadRequest, "review id is required", nil)
		return
	}

	var req dto.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid update review request", []response.ErrorResponse{{Field: "request", Error: err.Error()}})
		return
	}

	result, err := h.reviewService.Update(userID, reviewID, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "review updated successfully", result)
}

func (h *ReviewHandler) Delete(c *gin.Context) {
	userID, ok := getAuthUserID(c)
	if !ok {
		return
	}

	reviewID := strings.TrimSpace(c.Param("id"))
	if reviewID == "" {
		response.Error(c, http.StatusBadRequest, "review id is required", nil)
		return
	}

	if err := h.reviewService.Delete(userID, reviewID); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "review deleted successfully", nil)
}
