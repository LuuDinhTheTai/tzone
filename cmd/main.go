package main

import (
	"log/slog"

	"github.com/LuuDinhTheTai/tzone/infrastructure/configuration"
	"github.com/LuuDinhTheTai/tzone/infrastructure/database"
	server2 "github.com/LuuDinhTheTai/tzone/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configuration.LoadEnv()

	client, ctx, cancel, errDB := database.Connect(cfg.Database.MongoDBConfig.URI)
	if errDB != nil {
		slog.Error("Database connect err:", errDB)
		return
	}
	errPing := database.Ping(client, ctx)
	if errPing != nil {
		slog.Error("Database ping err:", errPing)
		return
	}
	defer database.Close(client, ctx, cancel)

	r := gin.Default()

	server := server2.NewServer(r, cfg, nil, client)

	err := server.Run()
	if err != nil {
		slog.Error("Server start err:", err)
		return
	}
}
