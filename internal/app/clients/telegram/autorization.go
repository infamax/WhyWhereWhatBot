package telegram

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *Bot) initUser(chatId int64, username string) error {
	ok, _ := b.client.Exist(context.TODO(), &pb.TelegramId{
		Id: uint64(chatId),
	})
	if ok.Ans {
		return status.Errorf(codes.AlreadyExists, "this user already exist in db")
	}
	_, err := b.client.Add(context.TODO(), &pb.User{
		TelegramId: uint64(chatId),
		Name:       username,
		Score: &pb.Score{
			Count: 0,
		},
	})

	return err
}
