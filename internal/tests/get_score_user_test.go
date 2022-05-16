package tests

import (
	"context"
	"github.com/gojuno/minimock/v3"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"github.com/infamax/WhyWhereWhatBot/internal/app/server"
	"github.com/infamax/WhyWhereWhatBot/internal/models"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestGetScoreUserSimple(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.GetMock.Expect(ctx, 1123131).Return(&models.User{
		ID:         1,
		TelegramId: 1123131,
		Name:       "max_on",
		Count:      0,
	}, nil)
	user, err := svc.GetScoreUser(ctx, &pb.TelegramId{
		Id: 1123131,
	})
	assert.Nil(t, err)
	assert.Equal(t, user.Count, uint64(0))
}

func TestGetScoreUserWhereUserIsNotExist(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.GetMock.Expect(ctx, 1123131).Return(nil,
		status.Errorf(codes.NotFound, "not found"))
	user, err := svc.GetScoreUser(ctx, &pb.TelegramId{
		Id: 1123131,
	})
	assert.Nil(t, user)
	assert.Error(t, err, status.Errorf(codes.NotFound, "not found"))
}
