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

func previousSlot() string {
	currentTime := time.Now()
	return fmt.Sprintf("%d-%d-%d", currentTime.Month(), currentTime.Day(), currentTime.Hour()/2-1)
}

func NewRecord(record []string) {
	redisClient.ZIncrBy(currentSlot(), 1, strings.Join(record, " "))
}

func getRecord(slot string) []string {
	result, error := redisClient.ZRevRangeByScore(slot, redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  1,
	}).Result()

	if error != nil {
		fmt.Print(error)
	}
	return result
}

func CurrentRecord() []string {
	return getRecord(currentSlot())
}

func PreviousRecord() []string {
	return getRecord(previousSlot())
}
