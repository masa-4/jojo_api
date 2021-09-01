package router_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/masa-4/jojo_api/src/router"
	"github.com/stretchr/testify/assert"
)

func TestPingRouter(t *testing.T) {
	r := router.Router()
	w := httptest.NewRecorder()

	// ヘルスチェック用のパスのテスト
	reqForHealth, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, reqForHealth)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), "{\"health_check\":\"OK\"}")

}

func TestGetSubtile(t *testing.T) {
	r := router.Router()
	w := httptest.NewRecorder()

	// サブタイトル用のパスのテスト
	reqForSubtitle, _ := http.NewRequest("GET", "/t/1", nil)
	r.ServeHTTP(w, reqForSubtitle)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), "{\"subtitle\":\"Fantom Blood\"}")
}

func TestGetCharacter(t *testing.T) {
	r := router.Router()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	testjson := bytes.NewBufferString("{\"character\":\"空 承太郎\"}")

	// 本体取得用のパス
	c.Request, _ = http.NewRequest("POST", "/p", testjson)
	r.ServeHTTP(w, c.Request)
	assert.JSONEq(t, w.Body.String(), "{\"stand\":\"スタープラチナ\"}")
	assert.Equal(t, w.Code, 200)
}
