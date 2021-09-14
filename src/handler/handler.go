package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masa-4/jojo_api/src/model"
)

type PartRequest struct {
	Part string `form:part`
}

type StandRequest struct {
	Character string `form:"character"`
}

type CharacterRequest struct {
	Stand string `form:"stand"`
}

func Health_check(c *gin.Context) {
	c.JSON(200, gin.H{
		"health_check": "OK",
	})
}

func ReturnSubtitle(c *gin.Context) {
	var reqJson PartRequest
	if err := c.ShouldBindJSON(&reqJson); err != nil {
		c.JSON(403, gin.H{
			"error": "リクエストボディーはjson形式にしてください",
		})
	}
	partInt, _ := strconv.Atoi(reqJson.Part)
	subtitle := model.GetSubtitleByPart(partInt)
	if subtitle.SUBTITLE == "" {
		c.JSON(404, gin.H{
			"error": "指定された部のサブタイトルがDBに登録されていません",
		})
	} else {
		c.JSON(200, gin.H{
			"subtitle": subtitle.SUBTITLE,
		})
	}
}

func ReturnStand(c *gin.Context) {
	var reqJson StandRequest
	if err := c.ShouldBindJSON(&reqJson); err != nil {
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

func ReturnCharacter(c *gin.Context) {
	var reqJson CharacterRequest
	if err := c.ShouldBindJSON(&reqJson); err != nil {
		c.JSON(403, gin.H{
			"error": "リクエストボディーはjson形式にしてください",
		})
	}
	fmt.Println(reqJson.Stand)
	character := model.GetCharacterNameByStand(reqJson.Stand)
	if character.Name == "" {
		c.JSON(404, gin.H{
			"error": "指定されたスタンドを持つ本体がDBに登録されていません",
		})
	} else {
		c.JSON(200, gin.H{
			"character": character.Name,
		})
	}
}
