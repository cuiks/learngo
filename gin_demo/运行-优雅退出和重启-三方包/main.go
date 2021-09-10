package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handler)
	// [...]
	endless.ListenAndServe(":4242", router)
}

func handler(context *gin.Context) {

}
