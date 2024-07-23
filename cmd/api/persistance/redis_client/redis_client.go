package redis_client

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	addr   string
	port   string
	passwd string
	db     int

	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(addr string, port string, passwd string, db int) *RedisClient {
	return &RedisClient{
		addr:   addr,
		port:   port,
		passwd: passwd,
		db:     db,
		ctx:    context.Background(),
		client: redis.NewClient(&redis.Options{
			Addr:     addr + ":" + port,
			Password: passwd,
			DB:       0,
		}),
	}
}

func (redisSession *RedisClient) ReadContraintValue(key string, limit int64) bool {
	currentMinute := time.Now().Minute()
	println("Trying to get key")

	println(key)

	redisKey := fmt.Sprintf("%s:%d", key, currentMinute)
	reqCheck, err := redisSession.client.Incr(redisSession.ctx, redisKey).Result()

	if err != nil || reqCheck > limit {
		return false
	}

	return true
}
