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

func TestGetScore(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.GetMock.Expect(ctx, 3113223).Return(
		&models.User{
			ID:         1,
			TelegramId: 3113223,
			Name:       "diana",
			Count:      10,
		}, nil)
	res, err := svc.GetScoreUser(ctx, &pb.TelegramId{
		Id: 3113223,
	})
	assert.Nil(t, err)
	assert.Equal(t, 10, int(res.Count))
}

func GetScoreUserWhereNotExist(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.GetMock.Expect(ctx, 3113223).Return(
		nil, status.Error(codes.NotFound, "not found"))
	res, err := svc.GetScoreUser(ctx, &pb.TelegramId{
		Id: 3113223,
	})
	assert.Equal(t, status.Error(codes.NotFound, "not found"), err)
	assert.Nil(t, res)
}
