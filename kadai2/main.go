package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"	
)


func main() {
    port, _ := strconv.Atoi(os.Args[1])
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	router.Use(cors.New(config))
	router.GET("/",HelloWorld)
	router.Run(fmt.Sprintf(":%d", port))
}
func HelloWorld(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"message": "Hello World!!",
	})
}