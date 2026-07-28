package internal

import (
	"time"

	"github.com/redis/go-redis/v9"
)

const DefaultRedisTTL = 7 * 24 * time.Hour

type Config struct {
	Username       string
	Password       string
	Url            string
	FraudAPIConfig FraudAPIConfig
	RedisClient    *redis.Client
	RedisTTL       time.Duration
}

type FraudAPIConfig struct {
	Authorization string
	ForwardHost   string
	Url           string
}

var coreAPI *Config

func SetConfig(username, password, url, authorization, fraud_url, forward_host string, redisClient *redis.Client) *Config {
	coreAPI = &Config{
		Username: username,
		Password: password,
		Url:      url,
		FraudAPIConfig: FraudAPIConfig{
			Authorization: authorization,
			ForwardHost:   forward_host,
			Url:           fraud_url,
		},
		RedisClient: redisClient,
		RedisTTL:    DefaultRedisTTL,
	}
	return coreAPI
}

func GetConfig() *Config {
	if coreAPI == nil {
		panic("CBECoreAPI not initialized. Please call SetConfig() first.")
	}
	return coreAPI
}
