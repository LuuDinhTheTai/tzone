package service

import "github.com/LuuDinhTheTai/tzone/internal/repository"

type DeviceService struct {
	mongoDbRepo *repository.MongoDbRepository
}

func NewDeviceService(mongoDbRepo *repository.MongoDbRepository) *DeviceService {
	return &DeviceService{
		mongoDbRepo: mongoDbRepo,
	}
}
