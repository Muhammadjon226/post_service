package service

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/Muhammadjon226/toDo-service/genproto"
)

func TestTaskService_Create(t *testing.T) {
	tests := map[string]struct {
		input pb.Task
		want  pb.Task
	}{
		"successful": {
			input: pb.Task{
				Assignee: "Assignee",
				Title:    "Title",
				Summary:  "Summary",
				Deadline: "2021-12-20",
				Status:   "active",
			},
			want: pb.Task{
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
			got, err := client.Create(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create task", err)
			}
			got.Id = ""
			got.CreatedAt = "2021-12-20"
			got.UpdatedAt = "2021-12-20"
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v,\n got: %v", name, tc.want, *got)
			}
		})
	}
}

func TestTaskService_Get(t *testing.T) {
	tests := map[string]struct {
		input pb.ByIdReq
		want  pb.Task
	}{
		"successful": {
			input: pb.ByIdReq{Id: "19af8b83-8605-4414-85ca-381c1012f5a2"},
			want: pb.Task{
				Id:        "19af8b83-8605-4414-85ca-381c1012f5a2",
				Assignee:  "Assignee",
				Title:     "Title",
				Summary:   "Summary",
				Deadline:  "2021-12-20T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-20T17:05:04.915623Z",
				UpdatedAt: "2021-12-20T17:05:04.915623Z",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.Get(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to Get task", err)
			}

			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v,\n got: %v", name, tc.want, *got)
			}
		})
	}
}

func TestTaskService_List(t *testing.T) {
	tests := map[string]struct {
		input pb.ListReq
		want  []pb.Task
	}{
		"successful": {
			input: pb.ListReq{Page: 1, Limit: 2},
			want: []pb.Task{
				{
					Id:        "19af8b83-8605-4414-85ca-381c1012f5a2",
					Assignee:  "Assignee",
					Title:     "Title",
					Summary:   "Summary",
					Deadline:  "2021-12-20T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-20T17:05:04.915623Z",
					UpdatedAt: "2021-12-20T17:05:04.915623Z",
				},
				{
					Id:        "905ece3d-0fae-4982-a27d-99c73dba41eb",
					Assignee:  "Assignee",
					Title:     "Title",
					Summary:   "Summary",
					Deadline:  "2021-12-20T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-20T17:19:13.260022Z",
					UpdatedAt: "2021-12-20T17:19:13.260022Z",
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.List(context.Background(), &tc.input)
			for i := range got.Tasks {
				if err != nil {
					t.Error("failed to Get Lists of task", err)
				}
				if !reflect.DeepEqual(tc.want[i], *got.Tasks[i]) {
					t.Fatalf("%s: expected: %v, \ngot: %v", name, tc.want[i], *got.Tasks[i])
				}
			}
		})
	}
}

func TestTaskService_Update(t *testing.T) {
	tests := map[string]struct {
		input pb.Task
		want  pb.Task
	}{
		"successful": {
			input: pb.Task{
				Id:       "905ece3d-0fae-4982-a27d-99c73dba41eb",
				Assignee: "Assignee11",
				Title:    "Title",
				Summary:  "Summary",
				Deadline: "2021-12-20",
				Status:   "active",
			},
			want: pb.Task{
				Id:        "905ece3d-0fae-4982-a27d-99c73dba41eb",
				Assignee:  "Assignee11",
				Title:     "Title",
				Summary:   "Summary",
				Deadline:  "2021-12-20T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-20T17:19:13.260022Z",
				UpdatedAt: "2021-12-20",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.Update(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to update task", err)
			}
			got.UpdatedAt = "2021-12-20"
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, \n\t\t\t\t\t\t\t\tgot: %v", name, tc.want, *got)
			}
		})
	}
}

func TestTaskService_Delete(t *testing.T) {
	tests := map[string]struct {
		input pb.ByIdReq
		want  pb.EmptyResp
	}{
		"successful": {
			input: pb.ByIdReq{Id: "905ece3d-0fae-4982-a27d-99c73dba41eb"},
			want:  pb.EmptyResp{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.Delete(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to Delete task", err)
			}
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v,\n got: %v", name, tc.want, *got)
			}
		})
	}
}

func TestTaskService_ListOverdue(t *testing.T) {
	tests := map[string]struct {
		input pb.ListOverReq
		want  []pb.Task
	}{
		"successful": {
			input: pb.ListOverReq{Time: "2021-12-20", Page: 1, Limit: 2},
			want: []pb.Task{
				{
					Id:        "19af8b83-8605-4414-85ca-381c1012f5a2",
					Assignee:  "Assignee",
					Title:     "Title",
					Summary:   "Summary",
					Deadline:  "2021-12-20T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-20T17:05:04.915623Z",
					UpdatedAt: "2021-12-20T17:05:04.915623Z",
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.ListOverDue(context.Background(), &tc.input)
			for i := range got.Tasks {
				if err != nil {
					t.Error("failed to Get Lists of task", err)
				}
				if !reflect.DeepEqual(tc.want[i], *got.Tasks[i]) {
					t.Fatalf("%s: expected: %v, \ngot: %v", name, tc.want[i], *got.Tasks[i])
				}
			}
		})
	}
}
