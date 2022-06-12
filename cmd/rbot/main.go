package main

import (
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"github.com/infamax/WhyWhereWhatBot/config"
	"github.com/infamax/WhyWhereWhatBot/internal/app/clients/telegram"
	"github.com/infamax/WhyWhereWhatBot/internal/app/server"
	"github.com/infamax/WhyWhereWhatBot/internal/db"
	"google.golang.org/grpc"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Чтение конфига
	b, err := os.ReadFile("config/config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}

	// Подключение к базе данных
	adp, err := db.New(cfg.Dsn)

	if err != nil {
		log.Fatal(err)
	}

	newServer := server.New(adp)
	//lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterWhyWhereWhatServerServer(grpcServer, newServer)
	//err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

	// Запускаем бота
	bot, err := tgbotapi.NewBotAPI(cfg.ApiKey.Telegram)
	if err != nil {
		log.Panic(err)
	}
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	client := pb.NewWhyWhereWhatServerClient(conn)

	bot.Debug = true
	telegramBot := telegram.New(bot, client)
	err = telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
