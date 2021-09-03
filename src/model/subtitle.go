package model

import (
	"time"

	"github.com/masa-4/jojo_api/src/db"
)

type Subtitle struct {
	ID        int    `gorm:"primarykey;autoIncrement;not null"`
	SUBTITLE  string `gorm:"not null"`
	PART      int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func MigrateSubtitleTable() {
	database := db.ConnectDb()
	database.AutoMigrate(&Subtitle{})
	defer database.Close()
}

func GetSubtitleByPart(part int) Subtitle {
	database := db.ConnectDb()
	defer database.Close()

	var subtile Subtitle
	database.Where("part=?", part).First(&subtile)

	return subtile
}
