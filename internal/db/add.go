package db

import (
	"context"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
)

func (d *Db) Add(ctx context.Context, user models.User) (int, error) {
	d.db.Create(&user)
	return int(user.ID), nil
}
