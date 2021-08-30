package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masa-4/jojo_api/src/model"
)

func Health_check(c *gin.Context) {
	c.JSON(200, gin.H{
		"health_check": "OK",
	})
}

func ReturnSubtitle(c *gin.Context) {
	partString := c.Param("part")
	partInt, _ := strconv.Atoi(partString)
	subetitle := model.GetSubtitleByPart(partInt)
	c.JSON(200, gin.H{
		"subtitle": subetitle.SUBTITLE,
	})
}
