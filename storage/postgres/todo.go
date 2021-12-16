package postgres

import (
	"time"
	"database/sql"

	"github.com/jmoiron/sqlx"
	pb "github.com/Muhammadjon226/toDo-service/genproto"
)

type taskRepo struct {
	db *sqlx.DB
}

// NewTaskRepo ...
func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task pb.Task) (pb.Task, error) {
	var id int64
	err := r.db.QueryRow(`
        INSERT INTO tasks(assignee, title, summary, deadline, status)
        VALUES ($1,$2, $3, $4, $5) returning id`, task.Assignee, task.Title, task.Summary, task.Deadline, task.Status).Scan(&id)
	if err != nil {
		return pb.Task{}, err
	}

	task, err = r.Get(id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Get(id int64) (pb.Task, error) {
	var task pb.Task
	err := r.db.QueryRow(`
        SELECT id, assignee, title, summary, deadline,status,created_at,updated_at FROM tasks
        WHERE id=$1 AND deleted_at IS NOT NULL AND updated_at IS NOT NULL`, id).Scan(
			&task.Id, 
			&task.Assignee, 
			&task.Title, 
			&task.Summary, 
			&task.Deadline, 
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) List(page, limit int64) ([]*pb.Task, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(`
		SELECT id, assignee, title, summary, deadline,status,created_at,updated_at FROM tasks 
		WHERE deleted_at IS NOT NULL
		LIMIT $1 OFFSET $2`,
		limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		tasks []*pb.Task
		task  pb.Task
		count int64
	)
	for rows.Next() {
		err = rows.Scan(
			&task.Id, 
			&task.Assignee, 
			&task.Title, 
			&task.Summary, 
			&task.Deadline, 
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM tasks`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (r *taskRepo) Update(task pb.Task) (pb.Task, error) {
	result, err := r.db.Exec(`
	UPDATE tasks 
	SET
		
		assignee = case when $1 = '' then title else $1 end,
		title = case when $2 = '' then title else $2 end,
		summary = case when $3 = '' then summary else $3 end,
		deadline = case when $4 = '' then deadline else $4 end,
		status = case when $5 = '' then status else $5 end,
		updated_at = current_timestamp
	 WHERE id=$6`,
		task.Assignee, task.Title, task.Summary, task.Deadline, task.Status, task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Task{}, sql.ErrNoRows
	}

	task, err = r.Get(task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Delete(id int64) error {
	result, err := r.db.Exec(`UPDATE tasks 
	SET
		deleted_at = current_timestamp
	WHERE id=$1`, id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *taskRepo) ListOverDue(t time.Time, page,limit int64) ([]*pb.Task, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(`
		SELECT id, assignee, title, summary, deadline,status from tasks WHERE deadline > $1 
		WHERE deleted_at IS NOT NULL
		LIMIT $2 OFFSET $3`,
		t,limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		tasks []*pb.Task
		task  pb.Task
		count int64
	)
	for rows.Next() {
		err = rows.Scan(&task.Id, &task.Assignee, &task.Title, &task.Summary, &task.Deadline, &task.Status)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM tasks`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}