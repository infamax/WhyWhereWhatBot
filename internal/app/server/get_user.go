package server

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
)

func (m MainServer) GetScoreUser(ctx context.Context, req *pb.TelegramId) (*pb.Score, error) {
	user, err := m.storage.Get(ctx, req.Id)
	if user == nil {
		return nil, err
	}
	return &pb.Score{
		Count: user.Count,
	}, err
}
