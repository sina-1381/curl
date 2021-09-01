package main

import (
	"curl/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/curl" , controller.Dextools)
	err := router.Run(":8080")
	if err != nil{
		panic(err)
	}
}