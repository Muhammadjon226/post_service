package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbFirst "github.com/Muhammadjon226/post_service/genproto/first_service"
	pbPost "github.com/Muhammadjon226/post_service/genproto/post_service"
	l "github.com/Muhammadjon226/post_service/pkg/logger"
	grpcclient "github.com/Muhammadjon226/post_service/service/grpc_client"
)

// PostService ...
type PostService struct {
	logger l.Logger
	client grpcclient.IGrpcClient
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpcclient.IGrpcClient) *PostService {
	return &PostService{
		logger: log,
		client: client,
	}
}
//CreatePost for create new post
func (s *PostService) CreatePost(ctx context.Context, req *pbPost.Post) (*pbPost.PostResponse, error) {

	post, err := s.client.FirstService().CreatePost(ctx, &pbFirst.Post{
		Id:     req.Id,
		UserId: req.UserId,
		Title:  req.Title,
		Body:   req.Body,
	})
	if err != nil {
		s.logger.Error("failed to create Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create Post")
	}

	return &pbPost.PostResponse{
		Id:        post.Id,
		UserId:    post.UserId,
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}
//GetPostByID for get post by id
func (s *PostService) GetPostByID(ctx context.Context, req *pbPost.ByIdReq) (*pbPost.PostResponse, error) {

	post, err := s.client.FirstService().GetPostByID(ctx, &pbFirst.ByIdReq{
		Id: req.Id,
	})
	if err != nil {
		s.logger.Error("failed to update Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update Post")
	}

	return &pbPost.PostResponse{
		Id:        post.Id,
		UserId:    post.UserId,
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

//ListPosts for listing posts
func (s *PostService) ListPosts(ctx context.Context, req *pbPost.ListReq) (*pbPost.ListResp, error) {
	posts, err := s.client.FirstService().ListPosts(ctx, &pbFirst.ListReq{
		Limit: req.Limit,
		Page:  req.Page,
	})
	if err != nil {
		s.logger.Error("failed to list Posts", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list Posts")
	}

	results := HelperFunction(posts.Posts)

	return &pbPost.ListResp{
		Posts: results,
		Count: posts.Count,
	}, nil
}

//UpdatePost for update post
func (s *PostService) UpdatePost(ctx context.Context, req *pbPost.Post) (*pbPost.PostResponse, error) {
	post, err := s.client.FirstService().UpdatePost(ctx, &pbFirst.Post{
		Id:     req.Id,
		UserId: req.UserId,
		Title:  req.Title,
		Body:   req.Body,
	})
	if err != nil {
		s.logger.Error("failed to update Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update Post")
	}
	return &pbPost.PostResponse{
		Id:        post.Id,
		UserId:    post.UserId,
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

//DeletePost for delete post by id
func (s *PostService) DeletePost(ctx context.Context, req *pbPost.ByIdReq) (*pbPost.EmptyResp, error) {
	_, err := s.client.FirstService().DeletePost(ctx, &pbFirst.ByIdReq{
		Id: req.Id,
	})
	if err != nil {
		s.logger.Error("failed to delete Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete Post")
	}

	return &pbPost.EmptyResp{}, nil
}
