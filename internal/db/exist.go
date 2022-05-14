package db

import (
	"context"
	"errors"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (d *Db) Exist(ctx context.Context, id uint64) (bool, error) {
	var user models.User
	err := d.db.Where("telegram_id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, status.Error(codes.NotFound, "not found")
	}
	return true, err
}
