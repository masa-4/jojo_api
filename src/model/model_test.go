package model_test

import (
	"fmt"
	"testing"

	"github.com/masa-4/jojo_api/src/model"
	"github.com/stretchr/testify/assert"
)

// サブタイトルを部数を指定してDBから取得するテスト
func TestGetSubtitleByPart(t *testing.T) {
	subtitle := model.GetSubtitleByPart(1)

	pass_or_fail := assert.Equal(t, subtitle.SUBTITLE, "戦闘潮流")
	if pass_or_fail == false {
		fmt.Println("部数からサブタイトルを取得する処理のテスト失敗")
	} else {
		fmt.Println("部数からサブタイトルを取得する処理テスト成功")
	}
}

// スタンドを本体を指定してDBから取得するテスト
func TestGetStandNameByCharacter(t *testing.T) {
	charactername := model.GetStandNameByCharacter("空条 承太郎")

	pass_or_fail := assert.Equal(t, charactername.STAND, "スタープラチナ")
	if pass_or_fail == false {
		fmt.Println("本体からスタンドを取得する処理のテスト失敗")
	} else {
		fmt.Println("本体からスタンドを取得する処理のテスト成功")
	}
}

// 本体をスタンドを指定してDBから取得するテスト
func TestGetCharacterNameByStand(t *testing.T) {
	standname := model.GetCharacterNameByStand("スタープラチナ")

	pass_or_fail := assert.Equal(t, standname.Name, "空条 承太郎")
	if pass_or_fail == false {
		fmt.Println("スタンドから本体を取得する処理のテスト失敗")
	} else {
		fmt.Println("スタンドから本体を取得する処理のテスト成功")
	}
}
