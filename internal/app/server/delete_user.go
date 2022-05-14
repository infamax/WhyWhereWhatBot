package server

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
)

func (m MainServer) Delete(ctx context.Context, req *pb.UserId) (*pb.Empty, error) {
	err := m.storage.Delete(ctx, req.Id)
	return &pb.Empty{}, err
}
