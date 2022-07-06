package repo

import (
	pb "github.com/Muhammadjon226/post_service/genproto/post_service"
)

// PostStorageI ...
type PostStorageI interface {
	Create(*pb.Post) (*pb.Post, error)
	Get(*pb.ByIdReq) (*pb.Post, error)
	List(*pb.ListReq) (*pb.ListResp, error)
	Update(*pb.Post) (*pb.Post, error)
	Delete(*pb.ByIdReq) error
}
