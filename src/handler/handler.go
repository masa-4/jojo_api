package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masa-4/jojo_api/src/model"
)

type StandRequest struct {
	Character string `json:"field_character`
}

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

func ReturnStand(c *gin.Context) {
	var reqJson StandRequest
	if err := c.BindJSON(&reqJson); err != nil {
		c.JSON(403, gin.H{
			"error": "リクエストボディーはjson形式にしてください",
		})
	}
	character := model.GetStandNameByCharacter(reqJson.Character)
	if character.STAND == "" {
		c.JSON(404, gin.H{
			"error": "指定された本体に対応するスタンドがDBに登録されていません",
		})
	} else {
		c.JSON(200, gin.H{
			"stand": character.STAND,
		})
	}

}
