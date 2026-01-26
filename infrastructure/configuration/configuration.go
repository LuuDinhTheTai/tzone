package configuration

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	MongoDbAtlas MongoDBConfig
	Supabase     SupabaseConfig
}

type MongoDBConfig struct {
	URL string
}

type SupabaseConfig struct {
	URL string
}

func LoadEnv() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("SERVER_PORT")

	uri := os.Getenv("MONGODB_ATLAS_URL")

	supabaseURL := os.Getenv("SUPABASE_URL")

	return Config{
		Server: ServerConfig{
			Port: port,
		},
		Database: DatabaseConfig{
			MongoDbAtlas: MongoDBConfig{
				URL: uri,
			},
			Supabase: SupabaseConfig{
				URL: supabaseURL,
			},
		},
	}
}
