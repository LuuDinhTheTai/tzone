package dto

// CreateBrandRequest represents the request body for creating a new brand
type CreateBrandRequest struct {
	Name string `json:"brand_name" binding:"required,min=1,max=100"`
}

// UpdateBrandRequest represents the request body for updating a brand
type UpdateBrandRequest struct {
	Name string `json:"brand_name" binding:"required,min=1,max=100"`
}
