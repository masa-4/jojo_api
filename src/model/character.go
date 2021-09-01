package model

import (
	"time"

	"github.com/masa-4/jojo_api/src/db"
)

type Character struct {
	ID        int    `gorm:"primarykey;autoIncrement;not null"`
	Name      string `gorm:"not null"`
	STAND     string `gorm:"not null"`
	PART      []int  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func MigrateCharacterTable() {
	database := db.ConnectDb()
	database.AutoMigrate(&Character{})
	defer database.Close()
}

func GetStandNameByCharacter(character string) Character {
	databse := db.ConnectDb()
	var chara_struct Character
	databse.First(&chara_struct, character)
	defer databse.Close()
	return chara_struct
}
