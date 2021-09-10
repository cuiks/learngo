package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/local/file", func(c *gin.Context) {
		c.File("/Users/cks/go_home/src/learngo/gin_demo/demo/local/file.go")
	})

	//var fs http.FileSystem = // ...
	//r.GET("/fs/file", func(c *gin.Context) {
	//	c.FileFromFS("fs/file.go", fs)
	//})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")

	// client
	// curl http://127.0.0.1:8080/JSONP?callback=x
}
