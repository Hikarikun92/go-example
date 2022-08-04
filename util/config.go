package util

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseServer   string
	DatabasePort     string
	DatabaseName     string
	ServerAddress    string
	ServerPort       string
	JwtSecret        []byte
	JwtExpirationMs  time.Duration
}

func LoadConfigFromEnvironment() *Config {
	config := &Config{
		DatabaseUser:     getOptionalEnv("DATABASE_USERNAME", "blog_backend_user"),
		DatabasePassword: getOptionalEnv("DATABASE_PASSWORD", "blog123"),
		DatabaseServer:   getOptionalEnv("DATABASE_SERVER", "localhost"),
		DatabasePort:     getOptionalEnv("DATABASE_PORT", "3306"),
		DatabaseName:     getOptionalEnv("DATABASE_NAME", "blog_backend_go"),
		ServerAddress:    getOptionalEnv("SERVER_ADDRESS", "localhost"),
		ServerPort:       getOptionalEnv("PORT", "8080"),
		JwtSecret:        []byte(getOptionalEnv("JWT_SECRET", "jwtsecret")),
	}

	expiration, err := strconv.ParseInt(getOptionalEnv("JWT_EXPIRATION_MS", "7200000"), 10, 64)
	if err != nil {
		log.Panicln("Error parsing JWT expiration", err)
	}
	config.JwtExpirationMs = time.Duration(expiration)

	return config
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
