package service

import "github.com/LuuDinhTheTai/tzone/internal/repository"

type DeviceService struct {
	mongoDbRepo *repository.BrandRepository
}

func NewDeviceService(mongoDbRepo *repository.BrandRepository) *DeviceService {
	return &DeviceService{
		mongoDbRepo: mongoDbRepo,
	}
}
