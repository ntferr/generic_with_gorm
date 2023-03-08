package models

type Motorbike struct {
	ID     string `gorm:"primaryKey"`
	Name   string `gorm:"colunm"`
	Wheels int    `gorm:"colunm"`
}
