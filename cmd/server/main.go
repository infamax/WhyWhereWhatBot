package main

import (
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"github.com/infamax/WhyWhereWhatBot/config"
	"github.com/infamax/WhyWhereWhatBot/internal/app/server"
	"github.com/infamax/WhyWhereWhatBot/internal/db"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	b, err := os.ReadFile("./config/config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}

	adp, err := db.New(cfg.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	newServer := server.New(adp)
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterWhyWhereWhatServerServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
