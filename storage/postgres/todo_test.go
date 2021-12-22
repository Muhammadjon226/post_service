package postgres

import (
	"reflect"
	"testing"
	"time"

	pb "github.com/Muhammadjon226/toDo-service/genproto"
)

func TestTaskRepo_Create(t *testing.T) {
	tests := map[string]struct {
		input pb.Task
		want  pb.Task
	}{
		"successful": {
			input: pb.Task{
				Id:       "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a1",
				Assignee: "Assignee",
				Title:    "Title",
				Summary:  "Summary",
				Deadline: "2021-12-20",
				Status:   "active",
			},
			want: pb.Task{
				Id:        "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a1",
				Assignee:  "Assignee",
				Title:     "Title",
				Summary:   "Summary",
				Deadline:  "2021-12-20T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-20",
				UpdatedAt: "2021-12-20",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := pgRepo.Create(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
			got.CreatedAt = "2021-12-20"
			got.UpdatedAt = "2021-12-20"
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Get(t *testing.T) {
	tests := map[string]struct {
		input string
		want  pb.Task
	}{
		"successful": {
			input: "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a1",
			want: pb.Task{
				Id:        "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a1",
				Assignee:  "Assignee",
				Title:     "Title",
				Summary:   "Summary",
				Deadline:  "2021-12-20T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-21T13:34:32.268858Z",
				UpdatedAt: "2021-12-21T08:34:32.268173Z",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := pgRepo.Get(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, \ngot: %v", name, tc.want, got)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, \n got: %v", name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_List(t *testing.T) {
	tests := map[string]struct {
		inputP int64
		inputL int64
		want   []pb.Task
	}{
		"successful": {
			inputP: 1,
			inputL: 2,
			want: []pb.Task{
				{
					Id:        "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a3",
					Assignee:  "Assignee",
					Title:     "Title",
					Summary:   "Summary",
					Deadline:  "2021-11-16T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-21 11:16:36.025121",
					UpdatedAt: "2021-12-21 11:16:36.025121",
				},
				{
					Id:        "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a1",
					Assignee:  "Assignee",
					Title:     "Title",
					Summary:   "Summary",
					Deadline:  "2021-12-20T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-20T03:55:43.124544Z",
					UpdatedAt: "2021-12-20T03:55:43.124544Z",
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _, err := pgRepo.List(tc.inputP, tc.inputL)
			for i := range got {
				if err != nil {
					t.Fatalf("%s: expected: %v, got: %v", name, tc.want[i], *got[i])
				}
				if !reflect.DeepEqual(tc.want[i], *got[i]) {
					t.Fatalf("%s: expected: %v, \ngot: %v", name, tc.want[i], *got[i])
				}
			}
		})
	}
}

func TestTaskRepo_Update(t *testing.T) {
	tests := map[string]struct {
		input pb.Task
		want  pb.Task
	}{
		"successful": {
			input: pb.Task{
				Id:       "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a1",
				Assignee: "Assignee",
				Title:    "Title",
				Summary:  "Summary something Update",
				Deadline: "2021-11-21",
				Status:   "active",
			},
			want: pb.Task{
				Id:        "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a1",
				Assignee:  "Assignee",
				Title:     "Title",
				Summary:   "Summary something Update",
				Deadline:  "2021-11-21T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-20",
				UpdatedAt: "2021-12-20",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := pgRepo.Update(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
			got.CreatedAt = "2021-12-20"
			got.UpdatedAt = "2021-12-20"
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, \n\t\t\t\t\t\t\t\tgot: %v", name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Delete(t *testing.T) {
	tests := map[string]struct {
		input string
		want  pb.Task
	}{
		"successful": {
			input: "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a1",
			want:  pb.Task{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := pgRepo.Delete(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, nil)
			}
		})
	}
}

func TestTaskRepo_ListOverDue(t *testing.T) {
	timer, _ := time.Parse("2006-01-02", "2021-11-16")
	tests := map[string]struct {
		inputP int64
		inputL int64
		timer  time.Time
		want   []pb.Task
	}{
		"successful": {
			inputP: 1,
			inputL: 2,
			timer:  timer,
			want: []pb.Task{
				{
					Id:        "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a3",
					Assignee:  "Assignee",
					Title:     "Title",
					Summary:   "Summary",
					Deadline:  "2021-11-16T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-21 11:16:36.025121",
					UpdatedAt: "2021-12-21 11:16:36.025121",
				},
				{
					Id:        "d41db3c4-84fe-4f39-9e6f-0c8a9bc5b3a1",
					Assignee:  "Assignee",
					Title:     "Title",
					Summary:   "Summary",
					Deadline:  "2021-12-20T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-20T03:55:43.124544Z",
					UpdatedAt: "2021-12-20T03:55:43.124544Z",
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _, err := pgRepo.ListOverDue(tc.timer, tc.inputP, tc.inputL)
			for i := range got {
				if err != nil {
					t.Fatalf("%s: expected: %v, got: %v", name, tc.want[i], *got[i])
				}
				if !reflect.DeepEqual(tc.want[i], *got[i]) {
					t.Fatalf("%s: expected: %v, \n\t\t\t\t\t\t\t\tgot: %v", name, tc.want[i], *got[i])
				}
			}
		})
	}
}
