package database

import (
	"gorm.io/gorm"
)

type database[T any] struct {
	db *gorm.DB
}

type Database[T any] interface {
	Get(params T, id string) (T, error)
	Post(params T) error
}

func NewDatabase[T any](db *gorm.DB) Database[T] {
	return &database[T]{
		db: db,
	}
}

func (d database[T]) Get(params T, id string) (T, error) {
	tx := d.db.Find(&params, "id = ?", id)
	err := tx.Error
	return params, err
}

func (d database[T]) Post(params T) error {
	tx := d.db.Create(&params)
	err := tx.Error
	return err
}

func (d database[T]) Update(params T, id string) error {
	tx := d.db.Where("id = ?", id).Updates(&params)
	err := tx.Error
	return err
}

func (d database[T]) Delete(params T, id string) error {
	tx := d.db.Where("id = ?", id).Delete(&params)
	err := tx.Error
	return err
}
