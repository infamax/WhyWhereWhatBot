package server

import (
	"context"
	"errors"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (m MainServer) GetTop(ctx context.Context, req *pb.GetLeaderResponse) (*pb.Leader, error) {
	users, err := m.storage.GetTop(ctx, req.Limit)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}
	ans := new(pb.Leader)
	ans.Name = make([]string, len(users), len(users))
	ans.Score = make([]*pb.Score, len(users), len(users))
	for i, user := range users {
		ans.Name[i] = user.Name
		ans.Score[i] = new(pb.Score)
		ans.Score[i].Count = user.Count
	}
	return ans, err
}
