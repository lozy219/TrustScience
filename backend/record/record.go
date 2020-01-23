package record

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "0.0.0.0:6379",
})

type Record string
type Result [2]int
type Match struct {
	Date   string
	Record Record
	Result Result
}

func timeToSlotNew(c *gin.Context, t time.Time) string {
	slot := fmt.Sprintf("%d-%d-%d-%d", t.Year(), t.Month(), t.Day(), t.Hour()/2)
	flag, _ := c.Cookie("ryf")
	if flag == "1" {
		slot += "-ryf"
	}
	return slot
}

func currentSlot(c *gin.Context) string {
	currentTime := time.Now()
	return timeToSlotNew(c, currentTime)
}

func previousSlot(c *gin.Context) string {
	prevTime := time.Now().Add(time.Duration(-2) * time.Hour)
	return timeToSlotNew(c, prevTime)
}

// NewRecord increments the counter for a record in current slot
func NewRecord(c *gin.Context, record []string) {
	redisClient.ZIncrBy(currentSlot(c), 1, strings.Join(record, " "))
}

func getRecord(slot string) Record {
	result, error := redisClient.ZRevRangeByScore(slot, redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  1,
	}).Result()

	if error != nil {
		fmt.Print(error)
	}
	if len(result) > 0 {
		return Record(result[0])
	}
	return ""
}

func getResult(slot string) Result {
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

// CurrentRecord returns the record of current slot
func CurrentRecord(c *gin.Context) Record {
	return getRecord(currentSlot(c))
}

// PreviousRecord returns the record of previous slot
func PreviousRecord(c *gin.Context) Record {
	return getRecord(previousSlot(c))
}

// ReportResult increments the coresponding result for previous slot
func ReportResult(c *gin.Context, index string) (int64, bool) {
	if index != "1" && index != "0" {
		return 0, true
	}
	count := redisClient.Incr(fmt.Sprintf("result.%s.%s", index, previousSlot(c))).Val()
	return count, false
}

// PreviousResult returns the reported results of previous slot
func PreviousResult(c *gin.Context) Result {
	slot := previousSlot(c)
	return getResult(slot)
}

// History returns history at a specific date
func History(c *gin.Context, y, m, d string) []Match {
	var matches []Match
	for i := 0; i <= 11; i++ {
		slot := fmt.Sprintf("%v-%v-%v", m, d, i)
		if y != "old" {
			slot = fmt.Sprintf("%v-%v-%v-%v", y, m, d, i)
		}
		record := getRecord(slot)
		result := getResult(slot)
		if len(record) > 0 {
			matches = append(matches, Match{
				fmt.Sprintf("%v/%v/%v %v:00", d, m, y, i*2),
				record,
				result,
			})
		}
	}
	return matches
}
