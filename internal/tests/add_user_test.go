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
	"log"
	"testing"
)

func TestAddUser(t *testing.T) {

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
	_, err := svc.Add(ctx, &pb.User{
		TelegramId: 1213311,
		Name:       "max_on",
		Score: &pb.Score{
			Count: 0,
		},
	})
	log.Println(err)
	assert.Nil(t, err)
}

func TestAddUserThisAlreadyExist(t *testing.T) {
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
	_, err := svc.Add(ctx, &pb.User{
		TelegramId: 1213311,
		Name:       "max_on",
		Score: &pb.Score{
			Count: 0,
		},
	})
	mockRepo.ExistMock.Expect(ctx, 1213311).Return(true, nil)
	_, err = svc.Add(ctx, &pb.User{
		TelegramId: 1213311,
		Name:       "max_on",
		Score: &pb.Score{
			Count: 0,
		},
	})
	//assert.Nil(t, err)
	assert.Error(t, err, codes.AlreadyExists)
}
