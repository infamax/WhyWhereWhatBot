package db

import (
	"context"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
)

func (d *Db) UpdateUserScoreById(ctx context.Context, id uint64, score uint64) error {
	ok, err := d.Exist(ctx, id)
	if !ok {
		return err
	}
	var user models.User
	_ = d.db.Where("telegram_id = ?", id).First(&user).Error
	user.Count = score
	return d.Update(ctx, user)
}
