package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"	
)

func main()  {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	router.Use(cors.New(config))
	router.GET("/",HelloWorld)
	router.Run(":8080")
}

func HelloWorld(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"message": "Hello World!!",
	})
}