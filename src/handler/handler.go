package handler

import (
	"github.com/gin-gonic/gin"
)

func Health_check(c *gin.Context) {
	c.JSON(200, gin.H{
		"health_check": "OK",
	})
}

func ReturnSubtitle(c *gin.Context) {

}
