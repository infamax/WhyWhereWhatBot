package db

import (
	"context"
	"errors"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (d *Db) Get(ctx context.Context, id uint64) (*models.User, error) {
	var user models.User
	err := d.db.First(&user, int(id)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return &user, err
}
