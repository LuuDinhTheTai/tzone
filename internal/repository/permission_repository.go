package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

// HasPermission performs a 5-table JOIN to verify if a user has access to a resource via a specific action (method)
func (r *PermissionRepository) HasPermission(userID uuid.UUID, action string, resource string) (bool, error) {
	var count int64
	err := r.db.Table("user_roles").
		Joins("JOIN role_permissions ON role_permissions.role_id = user_roles.role_id").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN actions ON actions.id = permissions.action_id").
		Joins("JOIN resources ON resources.id = permissions.resource_id").
		Where("user_roles.user_id = ? AND actions.name = ? AND resources.endpoint = ?", userID, action, resource).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
