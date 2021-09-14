package router_test

import (
	"bytes"
	"fmt"
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
	c, _ := gin.CreateTestContext(w)
	testjson := bytes.NewBufferString("{\"part\":\"1\"}")

	// サブタイトル用のパスのテスト
	c.Request, _ = http.NewRequest("POST", "/t", testjson)
	r.ServeHTTP(w, c.Request)
	statuscodetest := assert.Equal(t, 200, w.Code)
	if statuscodetest == false {
		fmt.Println("サブタイトル取得用APIのテストで200以外のステータスコードが返却されました")
	} else {
		fmt.Println("サブタイトル取得用APIで200が返却されました")
	}
	subtitletest := assert.Equal(t, "{\"subtitle\":\"ファントムブラッド\"}", w.Body.String())
	if subtitletest == false {
		fmt.Println("サブタイトル取得用APIで期待しないサブタイトルが返却されました")
	} else {
		fmt.Println("サブタイトル取得用APIで期待値が返却されましたs")
	}
}

func TestGetStand(t *testing.T) {
	r := router.Router()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	testjson := bytes.NewBufferString("{\"character\":\"空条 承太郎\"}")

	// スタンド取得用のパス
	c.Request, _ = http.NewRequest("POST", "/p", testjson)
	r.ServeHTTP(w, c.Request)
	teststand := assert.JSONEq(t, "{\"stand\":\"スタープラチナ\"}", w.Body.String())
	if teststand == false {
		fmt.Println("本体からスタンドを取得するルーティングのテストで失敗")
	} else {
		fmt.Println("本体からスタンドを取得するルーティングのテスト成功")
	}
	statuscode := assert.Equal(t, 200, w.Code)
	if statuscode == false {
		fmt.Println("スタンド取得APIで200以外のステータスコードが返却されました")
	} else {
		fmt.Println("スタンド取得APIで正常なステータスコードが返却されました。")
	}
}

func TestGetCharacter(t *testing.T) {
	r := router.Router()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	testjson := bytes.NewBufferString("{\"stand\":\"スタープラチナ\"}")

	// 本体取得用のパス
	c.Request, _ = http.NewRequest("POST", "/s", testjson)
	r.ServeHTTP(w, c.Request)
	charactertest := assert.JSONEq(t, "{\"character\":\"空条 承太郎\"}", w.Body.String())
	if charactertest == false {
		fmt.Println("本体取得APIで期待値と異なる本体名が返却されました")
	} else {
		fmt.Println("本体取得APIで正常に本体名が返却されました")
	}
	statuscode := assert.Equal(t, 200, w.Code)
	if statuscode == false {
		fmt.Println("本体取得APIで200以外のステータスコードが返却されました")
	} else {
		fmt.Println("本体取得APIで正常なステータスコードが返却されました")
	}
}
