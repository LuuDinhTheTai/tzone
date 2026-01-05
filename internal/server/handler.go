package server

import (
	"github.com/LuuDinhTheTai/tzone/internal/delivery/handler"
	"github.com/LuuDinhTheTai/tzone/internal/delivery/route"
	"github.com/LuuDinhTheTai/tzone/internal/repository"
	"github.com/LuuDinhTheTai/tzone/internal/service"
)

func (s *Server) MapHandlers() error {
	// Init repository
	mongoDBRepo := repository.NewMongoDbRepository()
	//postgreRepo := repository.NewPostgreRepository()

	// Init service
	brandService := service.NewBrandService(mongoDBRepo)
	deviceService := service.NewDeviceService(mongoDBRepo)

	// Init handler
	commonHandler := handler.NewCommonHandler()
	brandHandler := handler.NewBrandHandler(brandService)
	deviceHandler := handler.NewDeviceHandler(deviceService)

	// Init route
	route.MapCommonRoutes(s.r, commonHandler)
	route.MapBrandRoutes(s.r, brandHandler)
	route.MapDeviceRoutes(s.r, deviceHandler)

	return nil
}
