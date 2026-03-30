package service

import (
	"github.com/LuuDinhTheTai/tzone/internal/repository"
	"github.com/google/uuid"
)

type PermissionService struct {
	repo *repository.PermissionRepository
}

func NewPermissionService(repo *repository.PermissionRepository) *PermissionService {
	return &PermissionService{repo: repo}
}

func (s *PermissionService) HasPermission(userID uuid.UUID, action string, resource string) (bool, error) {
	return s.repo.HasPermission(userID, action, resource)
}
