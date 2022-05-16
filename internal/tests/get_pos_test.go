package tests

import (
	"context"
	"github.com/gojuno/minimock/v3"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"github.com/infamax/WhyWhereWhatBot/internal/app/server"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestGetPositionUser(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.GetPositionByIdMock.Expect(ctx, 218313).Return(12, nil)
	res, err := svc.GetPositionUser(ctx, &pb.TelegramId{Id: 218313})
	assert.Nil(t, err)
	assert.Equal(t, 12, int(res.Pos))
}

func TestGetPositionWhereUserNotExist(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.GetPositionByIdMock.Expect(ctx, 218313).Return(0, status.Error(codes.NotFound, "not found"))
	res, err := svc.GetPositionUser(ctx, &pb.TelegramId{Id: 218313})
	assert.Equal(t, 0, int(res.Pos))
	assert.Equal(t, err, status.Error(codes.NotFound, "not found"))
}
