package db

import (
	"context"
	"errors"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
	"gorm.io/gorm"
)

func (d *Db) GetTop(ctx context.Context, limit uint64) ([]models.User, error) {
	users := make([]models.User, 0, limit)
	err := d.db.Order("COUNT desc").Limit(int(limit)).Find(&users).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return users, nil
}
