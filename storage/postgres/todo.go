package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "github.com/Muhammadjon226/post_service/genproto/post_service"
)

type postRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (pr *postRepo) Create(*pb.Post) (*pb.Post, error) {

	return nil, nil
}

func (pr *postRepo) Get(*pb.ByIdReq) (*pb.Post, error) {

	return nil, nil
}

func (pr *postRepo) Delete(*pb.ByIdReq) error {

	return nil
}

func (pr *postRepo) Update(*pb.Post) (*pb.Post, error) {

	return nil, nil
}
func (pr *postRepo) List(*pb.ListReq) (*pb.ListResp, error) {

	return nil, nil
}