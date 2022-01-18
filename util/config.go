package util

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseServer   string
	DatabasePort     string
	DatabaseName     string
	ServerAddress    string
	ServerPort       string
}

func LoadConfigFromEnvironment() *Config {
	return &Config{
		DatabaseUser:     getOptionalEnv("DATABASE_USERNAME", "blog_backend_user"),
		DatabasePassword: getOptionalEnv("DATABASE_PASSWORD", "blog123"),
		DatabaseServer:   getOptionalEnv("DATABASE_SERVER", "localhost"),
		DatabasePort:     getOptionalEnv("DATABASE_PORT", "3306"),
		DatabaseName:     getOptionalEnv("DATABASE_NAME", "blog_backend_go"),
		ServerAddress:    getOptionalEnv("SERVER_ADDRESS", "localhost"),
		ServerPort:       getOptionalEnv("SERVER_PORT", "8080"),
	}
}

func getOptionalEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	} else {
		return defaultValue
	}
}

func (c *Config) GetDataSourceName() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", c.DatabaseUser, c.DatabasePassword, c.DatabaseServer, c.DatabasePort, c.DatabaseName)
}

func (c *Config) GetServerAddress() string {
	return c.ServerAddress + ":" + c.ServerPort
}
