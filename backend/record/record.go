package record

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "0.0.0.0:6379",
})

func currentSlot() string {
	currentTime := time.Now()
	return fmt.Sprintf("%d-%d-%d", currentTime.Month(), currentTime.Day(), currentTime.Hour()/2)
}

func NewRecord(record []string) {
	redisClient.ZIncrBy(currentSlot(), 1, strings.Join(record, " "))
}

func CurrentRecord() []string {
	result, error := redisClient.ZRevRangeByScore(currentSlot(), redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  1,
	}).Result()

	fmt.Print(error)
	return result
}
