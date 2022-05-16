package models

const (
	Easy = iota
	Medium
	Hard
)

type User struct {
	ID         uint `gorm:"primarykey"`
	TelegramId uint64
	Name       string
	Count      uint64
}
