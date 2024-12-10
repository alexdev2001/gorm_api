package models

type Student struct {
	Name  string `gorm:"primary_key"`
	Age   uint   `gorm:"unique"`
	Grade int    `gorm:"size:5"`
}
