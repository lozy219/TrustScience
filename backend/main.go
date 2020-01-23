package main

import (
	"fmt"
	"image/png"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lozy219/trustscience/backend/matching"
	"github.com/lozy219/trustscience/backend/record"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_count",
		},
		[]string{
			"endpoint",
		},
	)
)

func handleErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func router() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://uygnim.com", "http://129.204.1.146"}
	r.Use(cors.New(config))

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.POST("match", func(c *gin.Context) {
		requestCounter.WithLabelValues("match").Inc()
		file, _, err := c.Request.FormFile("match")
		handleErr(err)

		src := matching.LoadImage(file)

		fname := "./screenshots/" + matching.HashImage(src) + ".PNG"
		fout, err := os.Create(fname)
		handleErr(err)
		defer fout.Close()

		encodeErr := png.Encode(fout, src)
		handleErr(encodeErr)

		c.JSON(200, matching.Match(c, src))
	})

	r.GET("result", func(c *gin.Context) {
		requestCounter.WithLabelValues("result").Inc()
		c.JSON(200, gin.H{
			"current":  record.CurrentRecord(c),
			"previous": record.PreviousRecord(c),
			"result":   record.PreviousResult(c),
		})
	})

	r.GET("report/:index", func(c *gin.Context) {
		requestCounter.WithLabelValues("report").Inc()
		index := c.Param("index")
		count, err := record.ReportResult(c, index)
		c.JSON(200, gin.H{
			"error": err,
			"count": count,
		})
	})

	r.GET("history/:y/:m/:d", func(c *gin.Context) {
		requestCounter.WithLabelValues("history").Inc()
		y, m, d := c.Param("y"), c.Param("m"), c.Param("d")
		c.JSON(200, gin.H{
			"results": record.History(c, y, m, d),
		})
	})

	return r
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	f, _ := os.Create("/var/log/yys/yys-" + strconv.Itoa(int(time.Now().Unix())) + ".log")
	gin.DefaultWriter = io.MultiWriter(f)

	router().Run(":8734")
}
