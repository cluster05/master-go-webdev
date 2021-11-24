package main

import (
	"github.com/gin-gonic/gin"
)

const PORT string = ":3000"

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(PORT)

}
