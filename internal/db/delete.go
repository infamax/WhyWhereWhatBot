package db

import (
	"context"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
)

func (d *Db) Delete(ctx context.Context, id uint64) error {
	err := d.db.Delete(&models.User{}, id).Error
	return err
}
