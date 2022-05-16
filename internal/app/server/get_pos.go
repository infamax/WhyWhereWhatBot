package server

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
)

func (m MainServer) GetPositionUser(ctx context.Context, in *pb.TelegramId) (*pb.Position, error) {
	pos, err := m.storage.GetPositionById(ctx, in.Id)
	if err != nil {
		return &pb.Position{
			Pos: 0,
		}, err
	}
	return &pb.Position{
		Pos: uint64(pos),
	}, nil
}
