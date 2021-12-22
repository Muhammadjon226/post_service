package repo

import (
	"time"

	pb "github.com/Muhammadjon226/toDo-service/genproto"
)

// TaskStorageI ...
type TaskStorageI interface {
	Create(pb.Task) (pb.Task, error)
	Get(id string) (pb.Task, error)
	List(page, limit int64) ([]*pb.Task, int64, error)
	Update(pb.Task) (pb.Task, error)
	Delete(id string) error
	ListOverDue(t time.Time, page, limit int64) ([]*pb.Task, int64, error)
}
