package router

import (
	"github.com/gin-gonic/gin"
	"github.com/masa-4/jojo_api/src/db"
	"github.com/masa-4/jojo_api/src/handler"
	"github.com/masa-4/jojo_api/src/model"
)

func Router() *gin.Engine {
	router := gin.Default()
	database := db.ConnectDb()
	database.AutoMigrate(&model.Subtitle{})
	defer database.Close()

	// ヘルスチェック
	router.GET("/", handler.Health_check)

	// サブタイトル
	router.GET("/s/:part", handler.ReturnSubtitle)
	return router
}
