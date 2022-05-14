package server

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
	"gorm.io/gorm"
)

func (m MainServer) Update(ctx context.Context, req *pb.User) (*pb.Empty, error) {
	var user = models.User{
		Model:      gorm.Model{ID: uint(req.Id)},
		TelegramId: req.TelegramId,
		Name:       req.Name,
		Count:      req.Score.Count,
	}

	err := m.storage.Update(ctx, user)
	return &pb.Empty{}, err
}
