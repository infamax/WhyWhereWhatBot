package server

import (
	"context"
	"errors"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"gorm.io/gorm"
)

func (m MainServer) GetPositionUser(ctx context.Context, in *pb.TelegramId) (*pb.Position, error) {
	pos, err := m.storage.GetPositionById(ctx, in.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &pb.Position{
			Pos: 0,
		}, err
	}
	return &pb.Position{
		Pos: uint64(pos),
	}, nil
}
