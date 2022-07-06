package service

import (
	"log"
	"os"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/Muhammadjon226/post_service/genproto/post_service"
)

var client pb.PostServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	client = pb.NewPostServiceClient(conn)

	os.Exit(m.Run())
}
