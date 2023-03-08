package models

type Car struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"colunm"`
}
