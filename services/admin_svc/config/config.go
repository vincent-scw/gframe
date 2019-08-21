package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	port        = "PORT"
	redisServer = "REDIS_SERVER"
	gameURL     = "GAME_URL"
	oauthURL    = "OAUTH_URL"
)

func init() {
	// Set default configurations
	viper.SetDefault(redisServer, "localhost:6379")
	viper.SetDefault(port, 8451)
	viper.SetDefault(gameURL, "http://localhost:8441")
	viper.SetDefault(oauthURL, "http://localhost:8440")

	viper.AutomaticEnv() // automatically bind env
}

// GetRedisServer returns RedisServer
func GetRedisServer() string {
	return viper.GetString(redisServer)
}

// GetPort returns port
func GetPort() string {
	return fmt.Sprintf(":%d", viper.GetInt(port))
}

// GetGameURL returns game url
func GetGameURL() string {
	return viper.GetString(gameURL)
}

// GetOAuthURL returns oauth url
func GetOAuthURL() string {
	return viper.GetString(oauthURL)
}
