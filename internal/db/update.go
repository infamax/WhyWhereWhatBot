package db

import (
	"context"
	"errors"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (d *Db) Update(ctx context.Context, user models.User) error {
	err := d.db.Model(&user).Updates(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return status.Error(codes.NotFound, "not found")
	}
	return err
}
