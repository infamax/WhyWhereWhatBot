package server

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m MainServer) Exist(ctx context.Context, req *pb.TelegramId) (*pb.ExistResponse, error) {
	ok, err := m.storage.Exist(ctx, req.Id)
	if !ok {
		return &pb.ExistResponse{
			Ans: false,
		}, status.Error(codes.NotFound, "not found")
	}
	return &pb.ExistResponse{
		Ans: true,
	}, err
}
