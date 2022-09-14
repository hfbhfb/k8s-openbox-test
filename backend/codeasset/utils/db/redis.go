package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

var (
	RClient *redis.Client
)

type RConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func InitRedis(rc *RConfig) (*redis.Client, error) {
	//初始化 redis
	redisHost := rc.Host
	redisPort := rc.Port
	redisPassword := rc.Password
	redisDB := rc.DB
	redisURL := fmt.Sprintf("%s:%s", redisHost, redisPort)
	fmt.Printf("redis connction redisURL:%v\n", redisURL)
	RClient = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPassword,
		DB:       redisDB,
	})

	_, err := RClient.Ping().Result()
	if err != nil {
		fmt.Printf("redis connction err:%v\n", err)
		return nil, err
	}

	return RClient, err
}

func CloseRedis() {
	RClient.Close()
}
