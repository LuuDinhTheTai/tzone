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

	client, ctx, cancel, errDB := database.Connect(cfg.Database.MongoDbAtlas.URL)
	if errDB != nil {
		slog.Error("Database connect error", "error", errDB)
		return
	}
	errPing := database.Ping(client, ctx)
	if errPing != nil {
		slog.Error("Database ping error", "error", errPing)
		return
	}
	defer database.Close(client, ctx, cancel)

	supaClient, errSupa := database.ConnectSupabase(cfg.Database.Supabase.URL, cfg.Database.Supabase.Key)
	if errSupa != nil {
		slog.Warn("Supabase connect warning", "error", errSupa)
	}

	r := gin.Default()

	server := server2.NewServer(r, cfg, nil, client, supaClient)

	err := server.Run()
	if err != nil {
		slog.Error("Server start error", "error", err)
		return
	}
}
