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
	MongoDBConfig MongoDBConfig
}

type MongoDBConfig struct {
	URI string
}

func LoadEnv() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("SERVER_PORT")

	uri := os.Getenv("MONGODB_DATABASE_URI")

	return Config{
		Server: ServerConfig{
			Port: port,
		},
		Database: DatabaseConfig{
			MongoDBConfig: MongoDBConfig{
				URI: uri,
			},
		},
	}
}
