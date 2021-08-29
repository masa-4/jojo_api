package model

type Subtitle struct {
	ID       int    `gorm:"primarykey;autoIncrement;not null"`
	SUBTITLE string `gorm:"not null"`
	PART     int    `gorm:"not null"`
}
