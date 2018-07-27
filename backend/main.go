package main

import (
	"io"
	"os"
	"time"

	"./matching"
	"github.com/gin-gonic/gin"
)

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func router() *gin.Engine {
	r := gin.Default()

	r.POST("match", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("match")
		handleErr(err)

		fname := "./screenshots/" + time.Now().Format("1994032005") + ".PNG"
		fout, err := os.Create(fname)
		handleErr(err)

		_, copyErr := io.Copy(fout, file)
		handleErr(copyErr)

		fout.Close()

		fin, err := os.Open(fname)
		handleErr(err)
		defer fin.Close()

		c.JSON(200, matching.Match(fname))
	})

	return r
}

func main() {
	router().Run(":8734")
}
