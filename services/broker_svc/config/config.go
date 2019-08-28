package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	webPort      = "WEB_PORT"
	rpcPort 	 = "RPC_PORT"
	redisServer  = "REDIS_SERVER"
	kafkaBrokers = "KAFKA_BROKERS"
)

func init() {
	// Set default configurations
	viper.SetDefault(redisServer, "localhost:6379")
	viper.SetDefault(webPort, 8443)
	viper.SetDefault(rpcPort, 8543)
	viper.SetDefault(kafkaBrokers, []string{"localhost:9092"})

	viper.AutomaticEnv() // automatically bind env
}

// GetRedisServer returns RedisServer
func GetRedisServer() string {
	return viper.GetString(redisServer)
}

// GetWebPort returns port
func GetWebPort() string {
	return fmt.Sprintf(":%d", viper.GetInt(webPort))
}

// GetRPCPort returns port
func GetRPCPort() string {
	return fmt.Sprintf(":%d", viper.GetInt(rpcPort))
}

// GetKafkaBrokers returns kafka brokers
func GetKafkaBrokers() []string {
	return viper.GetStringSlice(kafkaBrokers)
}