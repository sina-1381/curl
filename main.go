package main

import (
	"curl/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	router := gin.Default()
	err := godotenv.Load(".env")
	if err != nil{
		panic(err)
	}
	router.POST("/curl" , controller.Dextools)
	err = router.Run(os.Getenv("PORT"))
}