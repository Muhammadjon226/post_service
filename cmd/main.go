package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Muhammadjon226/post_service/config"
	pbPost "github.com/Muhammadjon226/post_service/genproto/post_service"
	"github.com/Muhammadjon226/post_service/pkg/db"
	"github.com/Muhammadjon226/post_service/pkg/logger"
	"github.com/Muhammadjon226/post_service/service"
	grpcclient "github.com/Muhammadjon226/post_service/service/grpc_client"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "post_service")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	client, err := grpcclient.New(cfg)
	if err != nil {
		log.Error("error while connecting other services")
		return
	}

	postService := service.NewPostService(connDB, log, client)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pbPost.RegisterPostServiceServer(s, postService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
