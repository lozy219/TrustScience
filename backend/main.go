package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"./matching"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("test", func(c *gin.Context) {
		fmt.Println("abc")
		file, _, err := c.Request.FormFile("match")
		handleErr(err)

		fmt.Println(file)

		fname := "./screenshots/" + time.Now().Format("1994032005") + ".PNG"
		fout, err := os.Create(fname)
		handleErr(err)

		_, copyErr := io.Copy(fout, file)
		handleErr(copyErr)

		fout.Close()

		fin, err := os.Open(fname)
		handleErr(err)
		defer fin.Close()

		matching.MatchV3(fname)
		c.String(200, "test")
	})

	return r
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
