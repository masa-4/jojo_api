package model

import "github.com/masa-4/jojo_api/src/db"

func InsertTestData() {
	db := db.ConnectDb()
	sub := Subtitle{
		SUBTITLE: "ファントムブラッド",
		PART:     1,
	}
	db.Create(&sub)

	char := Character{
		Name:  "空条 承太郎",
		STAND: "スタープラチナ",
		PART:  345,
	}
	db.Create(&char)
	defer db.Close()
}
