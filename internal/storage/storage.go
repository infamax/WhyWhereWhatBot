package storage

import (
	"context"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
)

type Storage interface {
	Add(context.Context, models.User) (int, error)
	Get(context.Context, uint64) (*models.User, error)
	Update(context.Context, models.User) error
	Delete(context.Context, uint64) error
	GetTop(context.Context, uint64) ([]models.User, error)
	Exist(context.Context, uint64) (bool, error)
	GetPositionById(context.Context, uint64) (int64, error)
	UpdateUserScoreById(context.Context, uint64, uint64) error
}
