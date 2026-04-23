package dto

import (
	"time"

	"github.com/LuuDinhTheTai/tzone/internal/model"
)

type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}

// BrandResponse represents the response structure for a brand
type BrandResponse struct {
	Id   string `json:"id"`
	Name string `json:"brand_name"`
}

// BrandListResponse represents the response structure for a list of brands
type BrandListResponse struct {
	Brands     []BrandResponse `json:"brands"`
	Total      int             `json:"total"`
	Pagination PaginationMeta  `json:"pagination"`
}

// DeviceResponse represents the response structure for a device
type DeviceResponse struct {
	ID             string               `json:"id,omitempty"`
	BrandID        string               `json:"brand_id,omitempty"`
	ModelName      string               `json:"model_name,omitempty"`
	ImageUrl       string               `json:"imageUrl,omitempty"`
	Specifications model.Specifications `json:"specifications,omitempty"`
}

// DeviceListResponse represents the response structure for a list of devices
type DeviceListResponse struct {
	Devices    []DeviceResponse `json:"devices"`
	Total      int              `json:"total"`
	Pagination PaginationMeta   `json:"pagination"`
}

type FavoriteListResponse struct {
	DeviceIDs []string `json:"device_ids"`
}

type RecommendedDeviceCard struct {
	ID        string `json:"id"`
	BrandName string `json:"brand_name"`
	ModelName string `json:"model_name"`
	ImageURL  string `json:"imageUrl"`
	DetailURL string `json:"detail_url"`
	OS        string `json:"os,omitempty"`
	Chipset   string `json:"chipset,omitempty"`
	Memory    string `json:"memory,omitempty"`
	Battery   string `json:"battery,omitempty"`
	Price     string `json:"price,omitempty"`
}

type AIChatRecommendResponse struct {
	Reply   string                  `json:"reply"`
	Devices []RecommendedDeviceCard `json:"devices"`
}

type AIVideoReview struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type AIVideoReviewResponse struct {
	Reply  string          `json:"reply"`
	Videos []AIVideoReview `json:"videos"`
}

type ReviewResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	UserEmail string    `json:"user_email"`
	DeviceID  string    `json:"device_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReviewListResponse struct {
	Reviews       []ReviewResponse `json:"reviews"`
	Total         int              `json:"total"`
	Pagination    PaginationMeta   `json:"pagination"`
	RatingSummary RatingSummary    `json:"rating_summary"`
}

type RatingSummary struct {
	Average float64 `json:"average"`
	Count   int64   `json:"count"`
}
