package grpcClient

import (
	"github.com/Muhammadjon226/post_service/config"
)

// IGrpcClient ...
type IGrpcClient interface{}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	return &GrpcClient{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}
