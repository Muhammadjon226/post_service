package grpcclient

import (
	"fmt"

	"github.com/Muhammadjon226/post_service/config"
	pbFirst "github.com/Muhammadjon226/post_service/genproto/first_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// IGrpcClient ...
type IGrpcClient interface {
	FirstService() pbFirst.FirstServiceClient
}

// grpcClient ...
type grpcClient struct {
	cfg          config.Config
	firstService pbFirst.FirstServiceClient
}

// New ...
func New(cfg config.Config) (IGrpcClient, error) {

	connFirst, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.FirstServiceHost, cfg.FirstServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("first service dial host: %s port: %d",
			cfg.FirstServiceHost, cfg.FirstServicePort)
	}
	client := &grpcClient{
		cfg:          cfg,
		firstService: pbFirst.NewFirstServiceClient(connFirst),
	}

	return client, nil
}

// FirstService ...
func (s *grpcClient) FirstService() pbFirst.FirstServiceClient {
	return s.firstService
}
