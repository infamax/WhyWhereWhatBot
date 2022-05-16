package server

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m MainServer) Add(ctx context.Context, req *pb.User) (*pb.UserId, error) {
	ok, err := m.storage.Exist(ctx, req.TelegramId)
	if ok {
		return nil, status.Errorf(codes.AlreadyExists, "This is user already exist in db")
	}
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
