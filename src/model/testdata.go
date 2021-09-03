package model

import "github.com/masa-4/jojo_api/src/db"

func InsertTestData() {
	db := db.ConnectDb()
	sub := Subtitle{
		SUBTITLE: "戦闘潮流",
		PART:     2,
	}
	db.Create(&sub)

	char := Character{
		Name:  "空条 承太郎",
		STAND: "スタープラチナ",
		PART:  "3,4,5",
	}
	db.Create(&char)
	defer db.Close()
}
