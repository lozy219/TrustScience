package main

import (
	"TrustScience/backend/matching"
	"image/png"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://uygnim.com"}

	r.Use(cors.New(config))

	// 预留数据库连接池
	// db, err := sql.Open("mysql", "user:pwd@tcp(127.0.0.1:3306)/yys?parseTime=true")
	// defer db.Close()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// db.SetMaxIdleConns(20)
	// db.SetMaxOpenConns(20)

	// if err := db.Ping(); err != nil {
	// 	log.Fatalln(err)
	// }
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "It works")
	// })

	r.POST("match", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("match")
		HandleErr(err)

		src := matching.LoadImage(file)

		fname := "./screenshots/" + matching.HashImage(src) + ".PNG"
		fout, err := os.Create(fname)
		HandleErr(err)
		defer fout.Close()

		encodeErr := png.Encode(fout, src)
		HandleErr(encodeErr)

		c.JSON(200, matching.Match(src))
	})

	return r
}

func main() {
	router().Run(":8734")
}
