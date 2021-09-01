package model

import (
	"time"

	"github.com/masa-4/jojo_api/src/db"
)

type Character struct {
	ID        int    `gorm:"primarykey;autoIncrement;not null"`
	Name      string `gorm:"not null"`
	STAND     string `gorm:"not null"`
	PART      int    `gorm:"type:varchar(100)"`
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
	databse.Where("name=?", character).First(&chara_struct)
	defer databse.Close()
	return chara_struct
}
