package server

import (
	"github.com/LuuDinhTheTai/tzone/infrastructure/configuration"
	"github.com/gin-gonic/gin"
	supabase "github.com/supabase-community/supabase-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

type Server struct {
	r              *gin.Engine
	cfg            configuration.Config
	db             *gorm.DB
	mongoClient    *mongo.Client
	supabaseClient *supabase.Client
}

func NewServer(
	r *gin.Engine,
	cfg configuration.Config,
	db *gorm.DB,
	mongoClient *mongo.Client,
	supabaseClient *supabase.Client,
) *Server {
	return &Server{
		r:              r,
		cfg:            cfg,
		db:             db,
		mongoClient:    mongoClient,
		supabaseClient: supabaseClient,
	}
}

func (s *Server) Run() error {
	s.r = gin.Default()

	if err := s.MapHandlers(); err != nil {
		return err
	}

	if err := s.r.Run(":" + s.cfg.Server.Port); err != nil {
		return err
	}

	return nil
}
