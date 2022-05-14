package server

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
)

func (m MainServer) UpdateUserScore(ctx context.Context, in *pb.UserTelegram) (*pb.Empty, error) {
	err := m.storage.UpdateUserScoreById(ctx, in.TelegramId, in.NewScore)
	return &pb.Empty{}, err
}
