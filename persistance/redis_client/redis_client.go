package redis_client

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func ReadContraintValue(key string, limit int64) bool {
	currentMinute := time.Now().Minute()

	redisKey := fmt.Sprintf("%s:%d", key, currentMinute)
	reqCheck, err := client.Incr(ctx, redisKey).Result()

	if err != nil || reqCheck > limit {
		return false
	}

	return true
}
