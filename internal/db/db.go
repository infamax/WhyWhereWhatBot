package db

import (
	"github.com/infamax/WhyWhereWhatBot/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	db *gorm.DB
}

func New(dsn string) (*Db, error) {
	adp, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Миграция схем
	err = adp.AutoMigrate(&models.User{})
	return &Db{
		db: adp,
	}, err
}
