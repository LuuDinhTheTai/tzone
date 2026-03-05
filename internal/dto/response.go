package dto

import (
	"time"
)

// BrandResponse represents the response structure for a brand
type BrandResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"brand_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BrandListResponse represents the response structure for a list of brands
type BrandListResponse struct {
	Brands []BrandResponse `json:"brands"`
	Total  int             `json:"total"`
}
