package server

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
)

func (m MainServer) Add(ctx context.Context, req *pb.User) (*pb.UserId, error) {
	var user = models.User{
		TelegramId: req.TelegramId,
		Name:       req.Name,
		Count:      req.Score.Count,
	}
	id, err := m.storage.Add(ctx, user)
	return &pb.UserId{
		Id: uint64(id),
	}, err
}
