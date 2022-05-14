package db

import (
	"context"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
)

func (d *Db) GetPositionById(ctx context.Context, id uint64) (int64, error) {
	ok, err := d.Exist(ctx, id)
	if !ok {
		return 0, err
	}
	users := make([]models.User, 0)
	err = d.db.Order("COUNT desc").Find(&users).Error
	return Pos(users, id), err
}

func Pos(users []models.User, id uint64) int64 {
	for i, v := range users {
		if v.TelegramId == id {
			return int64(i + 1)
		}
	}
	return 0
}
