package telegram

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
)

func (b *Bot) getTop(limit int) (*pb.Leader, error) {
	leaders, err := b.client.GetTop(context.TODO(), &pb.GetLeaderResponse{
		Limit: uint64(limit),
	})
	return leaders, err
}
