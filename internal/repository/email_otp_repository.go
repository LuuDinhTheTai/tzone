package repository

import (
	"time"

	"github.com/LuuDinhTheTai/tzone/internal/model"
	"gorm.io/gorm"
)

type EmailOTPRepository struct {
	db *gorm.DB
}

func NewEmailOTPRepository(db *gorm.DB) *EmailOTPRepository {
	return &EmailOTPRepository{db: db}
}

func (r *EmailOTPRepository) Create(otp *model.EmailOTP) error {
	return r.db.Create(otp).Error
}

func (r *EmailOTPRepository) GetLatestByEmailPurpose(email string, purpose string) (*model.EmailOTP, error) {
	var otp model.EmailOTP
	err := r.db.
		Where("email = ? AND purpose = ?", email, purpose).
		Order("created_at DESC").
		First(&otp).Error
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *EmailOTPRepository) GetLatestActiveByEmailPurpose(email string, purpose string, now time.Time) (*model.EmailOTP, error) {
	var otp model.EmailOTP
	err := r.db.
		Where("email = ? AND purpose = ? AND is_used = ? AND expires_at > ?", email, purpose, false, now).
		Order("created_at DESC").
		First(&otp).Error
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *EmailOTPRepository) InvalidateActiveByEmailPurpose(email string, purpose string, now time.Time) error {
	return r.db.Model(&model.EmailOTP{}).
		Where("email = ? AND purpose = ? AND is_used = ? AND expires_at > ?", email, purpose, false, now).
		Updates(map[string]interface{}{"is_used": true}).Error
}

func (r *EmailOTPRepository) IncrementAttempt(id uint) error {
	return r.db.Model(&model.EmailOTP{}).
		Where("id = ?", id).
		UpdateColumn("attempt_count", gorm.Expr("attempt_count + ?", 1)).Error
}

func (r *EmailOTPRepository) MarkUsed(id uint) error {
	return r.db.Model(&model.EmailOTP{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"is_used": true}).Error
}
