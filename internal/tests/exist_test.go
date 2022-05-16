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

func TestGetUser(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.ExistMock.Expect(ctx, 10113).Return(true, nil)
	res, err := svc.Exist(ctx, &pb.TelegramId{
		Id: 10113,
	})
	assert.Nil(t, err)
	assert.Equal(t, true, res.Ans)
	mockRepo.ExistMock.Expect(ctx, 31313).Return(false, status.Error(codes.NotFound, "not found"))
	res, err = svc.Exist(ctx, &pb.TelegramId{
		Id: 31313,
	})
	assert.Equal(t, false, res.Ans)
	assert.Equal(t, err, status.Error(codes.NotFound, "not found"))
}
