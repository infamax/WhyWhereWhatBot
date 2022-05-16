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

func TestSimpleUpdateUser(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.ExistMock.Expect(ctx, 1213311).Return(false, status.Errorf(codes.NotFound, "not found"))
	mockRepo.AddMock.Expect(ctx, models.User{
		TelegramId: 1213311,
		Name:       "max_on",
		Count:      0,
	}).Return(1, nil)
	id, _ := svc.Add(ctx, &pb.User{
		TelegramId: 1213311,
		Name:       "max_on",
		Score: &pb.Score{
			Count: 0,
		},
	})
	mockRepo.UpdateMock.Expect(ctx, models.User{
		ID:         uint(id.Id),
		TelegramId: 1213311,
		Name:       "max_on",
		Count:      10,
	}).Return(nil)

	_, err := svc.Update(ctx, &pb.User{
		Id:         id.Id,
		TelegramId: 1213311,
		Name:       "max_on",
		Score: &pb.Score{
			Count: 10,
		},
	})

	assert.Nil(t, err)
	mockRepo.GetMock.Expect(ctx, 1213311).Return(
		&models.User{
			TelegramId: 1213311,
			Name:       "max_on",
			Count:      10,
		}, nil)
	user, _ := svc.GetScoreUser(ctx, &pb.TelegramId{
		Id: 1213311,
	})
	assert.Equal(t, int(user.Count), 10)
}
