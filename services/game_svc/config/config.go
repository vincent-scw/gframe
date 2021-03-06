package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	port         = "PORT"
	jwtKey       = "JWT_KEY"
	kafkaBrokers = "KAFKA_BROKERS"
	brokerRPC    = "BROKER_RPC"
	redisServer  = "REDIS_SERVER"
)

func init() {
	viper.SetDefault(port, 8441)
	viper.SetDefault(jwtKey, "00000000")
	viper.SetDefault(kafkaBrokers, []string{"localhost:9092"})
	viper.SetDefault(brokerRPC, "localhost:8543")
	viper.SetDefault(redisServer, "localhost:6379")

	viper.AutomaticEnv()
}

// GetRedisServer returns RedisServer
func GetRedisServer() string {
	return viper.GetString(redisServer)
}

// GetPort returns port
func GetPort() string {
	return fmt.Sprintf(":%d", viper.GetInt(port))
}

// GetJwtKey returns JwtKey
func GetJwtKey() []byte {
	return []byte(viper.GetString(jwtKey))
}

// GetKafkaBrokers returns KafkaBrokers
func GetKafkaBrokers() []string {
	return viper.GetStringSlice(kafkaBrokers)
}

// GetBrokerRPC returns BrokerRPC
func GetBrokerRPC() string {
	return viper.GetString(brokerRPC)
}
