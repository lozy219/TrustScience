package record

import (
	"fmt"
	"strconv"
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

// NewRecord increments the counter for a record in current slot
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

// CurrentRecord returns the record of current slot
func CurrentRecord() []string {
	return getRecord(currentSlot())
}

// PreviousRecord returns the record of previous slot
func PreviousRecord() []string {
	return getRecord(previousSlot())
}

// ReportResult increments the coresponding result for previous slot
func ReportResult(index string) bool {
	if index != "1" && index != "0" {
		return true
	}
	redisClient.Incr(fmt.Sprintf("result.%s.%s", index, previousSlot()))
	return false
}

// PreviousResult returns the reported results of previous slot
func PreviousResult() [2]int {
	slot := previousSlot()
	result0, err0 := strconv.Atoi(redisClient.Get(fmt.Sprintf("result.%d.%s", 0, slot)).Val())
	if err0 != nil {
		result0 = 0
	}
	result1, err1 := strconv.Atoi(redisClient.Get(fmt.Sprintf("result.%d.%s", 1, slot)).Val())
	if err1 != nil {
		result1 = 0
	}
	return [2]int{result0, result1}
}
