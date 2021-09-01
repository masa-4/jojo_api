package router

import (
	"github.com/gin-gonic/gin"
	"github.com/masa-4/jojo_api/src/handler"
	"github.com/masa-4/jojo_api/src/model"
)

func Router() *gin.Engine {
	router := gin.Default()
	model.MigrateSubtitleTable()
	model.MigrateCharacterTable()

	// ヘルスチェック
	router.GET("/", handler.Health_check)

	// サブタイトル
	router.GET("/s/:part", handler.ReturnSubtitle)

	// スタンド
	router.POST("/p", handler.ReturnStand)
	return router
}
