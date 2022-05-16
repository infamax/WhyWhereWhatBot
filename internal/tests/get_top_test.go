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
	"gorm.io/gorm"
	"testing"
)

func TestSimpleGetTop(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.GetTopMock.Expect(ctx, 3).Return([]models.User{
		models.User{
			ID:         1,
			TelegramId: 13131,
			Name:       "nedinai",
			Count:      121,
		},
		models.User{
			ID:         2,
			TelegramId: 19131,
			Name:       "darrrk",
			Count:      102,
		},
		models.User{
			ID:         3,
			TelegramId: 9313412,
			Name:       "diana",
			Count:      71,
		},
	}, nil)

	res, err := svc.GetTop(ctx, &pb.GetLeaderResponse{
		Limit: 3,
	})

	assert.Nil(t, err)
	assert.Equal(t, len(res.Name), 3)
	expectedName := []string{"nedinai", "darrrk", "diana"}
	for i := 0; i < 3; i++ {
		assert.Equal(t, res.Name[i], expectedName[i])
	}
	expectedScore := []uint64{121, 102, 71}
	for i := 0; i < 3; i++ {
		assert.Equal(t, int(res.Score[i].Count), int(expectedScore[i]))
	}
}

func TestGetTopWhereUsersNotExist(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.GetTopMock.Expect(ctx,
		3).Return(nil, gorm.ErrRecordNotFound)
	res, err := svc.GetTop(ctx, &pb.GetLeaderResponse{
		Limit: 3,
	})
	assert.Nil(t, res)
	assert.Equal(t, err, status.Error(codes.NotFound, "not found"))
}
