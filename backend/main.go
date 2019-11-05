package main

import (
	"image/png"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lozy219/trustscience/backend/matching"
	"github.com/lozy219/trustscience/backend/record"
)

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func router() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://uygnim.com", "http://129.204.1.146"}

	r.Use(cors.New(config))

	r.POST("match", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("match")
		handleErr(err)

		src := matching.LoadImage(file)

		fname := "./screenshots/" + matching.HashImage(src) + ".PNG"
		fout, err := os.Create(fname)
		handleErr(err)
		defer fout.Close()

		encodeErr := png.Encode(fout, src)
		handleErr(encodeErr)

		c.JSON(200, matching.Match(src))
	})

	r.GET("result", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"current":  record.CurrentRecord(),
			"previous": record.PreviousRecord(),
			"result":   record.PreviousResult(),
		})
	})

	r.GET("report/:index", func(c *gin.Context) {
		index := c.Param("index")
		count, err := record.ReportResult(index)
		c.JSON(200, gin.H{
			"error": err,
			"count": count,
		})
	})

	r.GET("history/:y/:m/:d", func(c *gin.Context) {
		y, m, d := c.Param("y"), c.Param("m"), c.Param("d")
		c.JSON(200, gin.H{
			"results": record.History(y, m, d),
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
