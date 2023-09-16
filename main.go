package main

import (
	"test-build-env/config"
	"test-build-env/controller"

	"github.com/gin-gonic/gin"
)

func main() {
  config.NewDBConnection("host=localhost user=root password=root dbname=test-build-env port=5432 sslmode=disable")

    r := gin.Default()
    
    r.GET("/", func(c *gin.Context) {
      c.JSON(200, gin.H{
        "message": "Testing route",
      })
    })

    r.POST("/book", controller.SaveBook)

    r.Run(":8080")
}
