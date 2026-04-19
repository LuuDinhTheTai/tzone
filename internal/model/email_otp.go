package model

import "time"

const (
	OTPPurposeRegister      = "register"
	OTPPurposeResetPassword = "reset_password"
)

type EmailOTP struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Email        string    `gorm:"index:idx_email_purpose_status,priority:1;not null;column:email"`
	Purpose      string    `gorm:"index:idx_email_purpose_status,priority:2;not null;column:purpose"`
	CodeHash     string    `gorm:"not null;column:code_hash"`
	ExpiresAt    time.Time `gorm:"index;not null;column:expires_at"`
	AttemptCount int       `gorm:"default:0;column:attempt_count"`
	IsUsed       bool      `gorm:"index:idx_email_purpose_status,priority:3;default:false;column:is_used"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at"`
}

func (EmailOTP) TableName() string {
	return "email_otps"
}
