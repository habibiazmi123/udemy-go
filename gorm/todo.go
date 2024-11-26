package belajar_golang_gorm

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID          int64  `gorm:"primary_key;column:id;autoIncrement"`
	UserID      int64  `gorm:"column:user_id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
}

func (t *Todo) TableName() string {
	return "todos"
}
