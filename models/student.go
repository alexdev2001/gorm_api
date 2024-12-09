package models

type Student struct {
	name  string `gorm:"primary_key"`
	age   uint   `gorm:"unique"`
	grade int    `gorm:"size:5"`
}
