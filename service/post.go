package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbPost "github.com/Muhammadjon226/post_service/genproto/post_service"
	l "github.com/Muhammadjon226/post_service/pkg/logger"
	"github.com/Muhammadjon226/post_service/storage"
)

// PostService ...
type PostService struct {
	logger  l.Logger
	storage storage.IStorage
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger) *PostService {
	return &PostService{
		logger: log,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *pbPost.Post) (*pbPost.Post, error) {

	Post, err := s.storage.Post().Create(req)
	if err != nil {
		s.logger.Error("failed to create Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create Post")
	}

	return Post, nil
}

func (s *PostService) GetPostById(ctx context.Context, req *pbPost.ByIdReq) (*pbPost.Post, error) {

	// if err != nil {
	// 	s.logger.Error("failed to get Post", l.Error(err))
	// 	return nil, status.Error(codes.Internal, "failed to get Post")
	// }

	return nil, nil
}

func (s *PostService) ListPosts(ctx context.Context, req *pbPost.ListReq) (*pbPost.ListResp, error) {
	Posts, err := s.storage.Post().List(req)
	if err != nil {
		s.logger.Error("failed to list Posts", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list Posts")
	}

	return &pbPost.ListResp{
		Posts: Posts.Posts,
		Count: Posts.Count,
	}, nil
}

func (s *PostService) UpdatePost(ctx context.Context, req *pbPost.Post) (*pbPost.Post, error) {
	Post, err := s.storage.Post().Update(req)
	if err != nil {
		s.logger.Error("failed to update Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update Post")
	}

	return Post, nil
}

func (s *PostService) DeletePost(ctx context.Context, req *pbPost.ByIdReq) (*pbPost.EmptyResp, error) {
	err := s.storage.Post().Delete(req)
	if err != nil {
		s.logger.Error("failed to delete Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete Post")
	}

	return &pbPost.EmptyResp{}, nil
}
