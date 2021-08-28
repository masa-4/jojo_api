package router

import "github.com/gin-gonic/gin"

func Router() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	router.Run()
}
