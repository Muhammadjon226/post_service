package service

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Muhammadjon226/toDo-service/genproto"
	l "github.com/Muhammadjon226/toDo-service/pkg/logger"
	"github.com/Muhammadjon226/toDo-service/storage"
)

// TaskService ...
type TaskService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewTaskService ...
func NewTaskService(db *sqlx.DB, log l.Logger) *TaskService {
	return &TaskService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *TaskService) Create(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	Task, err := s.storage.Task().Create(*req)
	if err != nil {
		s.logger.Error("failed to create Task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create Task")
	}

	return &Task, nil
}

func (s *TaskService) Get(ctx context.Context, req *pb.ByIdReq) (*pb.Task, error) {
	Task, err := s.storage.Task().Get(req.GetId())
	if err != nil {
		s.logger.Error("failed to get Task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get Task")
	}

	return &Task, nil
}

func (s *TaskService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
	Tasks, count, err := s.storage.Task().List(req.Page, req.Limit)
	if err != nil {
		s.logger.Error("failed to list Tasks", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list Tasks")
	}

	return &pb.ListResp{
		Tasks: Tasks,
		Count: count,
	}, nil
}

func (s *TaskService) Update(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	Task, err := s.storage.Task().Update(*req)
	if err != nil {
		s.logger.Error("failed to update Task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update Task")
	}

	return &Task, nil
}

func (s *TaskService) Delete(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := s.storage.Task().Delete(req.Id)
	if err != nil {
		s.logger.Error("failed to delete Task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete Task")
	}

	return &pb.EmptyResp{}, nil

	
}

func (s *TaskService) ListOverDue(ctx context.Context, req *pb.ListOverReq) (*pb.ListOverResp, error) {
	duration, err := time.Parse("2006-01-02", req.Time)
	if err != nil {
		s.logger.Error("failed to parse Time", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to parse Time")
	}
	Tasks, count, err := s.storage.Task().ListOverDue(duration, req.Page, req.Limit)
	if err != nil {
		s.logger.Error("failed to list Tasks", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list Tasks")
	}

	return &pb.ListOverResp{
		Tasks: Tasks,
		Count: count,
	}, nil
}
