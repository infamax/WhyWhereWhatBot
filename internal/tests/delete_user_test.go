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

func TestSimpleDeleteUser(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewStorageMock(mc)
	svc := server.New(mockRepo)
	ctx := context.Background()
	mockRepo.DeleteMock.Expect(ctx, 1).Return(nil)
	_, err := svc.Delete(ctx, &pb.UserId{Id: 1})
	mockRepo.DeleteMock.Expect(ctx, 3).Return(status.Errorf(codes.NotFound, "not found"))
	assert.Nil(t, err)
	_, err = svc.Delete(ctx, &pb.UserId{Id: 3})
	if err == nil {
		t.Errorf("Error")
	}
}
