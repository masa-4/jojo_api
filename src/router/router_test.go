package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
	reqForSubtitle, _ := http.NewRequest("GET", "/s/1", nil)
	r.ServeHTTP(w, reqForSubtitle)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), "{\"subtitle\":\"Fantom Blood\"}")
}
