package service

import (
	"github.com/LuuDinhTheTai/tzone/internal/repository"
)

type BrandService struct {
	mongoDbRepo *repository.MongoDbRepository
}

func NewBrandService(mongoDbRepo *repository.MongoDbRepository) *BrandService {
	return &BrandService{
		mongoDbRepo: mongoDbRepo,
	}
}
