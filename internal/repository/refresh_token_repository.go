package repository

import (
	"github.com/LuuDinhTheTai/tzone/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Create(token *model.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *RefreshTokenRepository) FindByID(id uuid.UUID) (*model.RefreshToken, error) {
	var token model.RefreshToken
	err := r.db.Where("id = ?", id).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *RefreshTokenRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.RefreshToken{}).Error
}

func (r *RefreshTokenRepository) DeleteAllByUserID(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.RefreshToken{}).Error
}
