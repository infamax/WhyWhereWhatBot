package models

import "gorm.io/gorm"

const (
	Easy = iota
	Medium
	Hard
)

type User struct {
	gorm.Model
	TelegramId uint64
	Name       string
	Count      uint64
}
