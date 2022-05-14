package server

import (
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"github.com/infamax/WhyWhereWhatBot/internal/storage"
)

type MainServer struct {
	storage storage.Storage
	pb.UnimplementedWhyWhereWhatServerServer
}

func New(s storage.Storage) *MainServer {
	return &MainServer{
		storage: s,
	}
}
