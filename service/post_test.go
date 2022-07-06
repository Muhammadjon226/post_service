package service

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/Muhammadjon226/post_service/genproto/post_service"
)

func TestPostService_CreatePost(t *testing.T) {
	tests := map[string]struct {
		input pb.Post
		want  pb.Post
	}{
		"successful": {
			input: pb.Post{
				Id:     1,
				UserId: 2,
				Title:  "Title",
				Body:   "body",
			},
			want: pb.Post{
				Id:     1,
				UserId: 2,
				Title:  "Title",
				Body:   "body",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.CreatePost(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create task", err)
			}
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v,\n got: %v", name, tc.want, *got)
			}
		})
	}
}

func TestPostService_GetPostById(t *testing.T) {
	tests := map[string]struct {
		input pb.ByIdReq
		want  pb.Post
	}{
		"successful": {
			input: pb.ByIdReq{Id: 1},
			want: pb.Post{
				Id:    1,
				Title: "Title",
				Body:  "Body",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.GetPostById(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to Get task", err)
			}

			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v,\n got: %v", name, tc.want, *got)
			}
		})
	}
}

func TestPostService_ListPosts(t *testing.T) {
	tests := map[string]struct {
		input pb.ListReq
		want  []pb.Post
	}{
		"successful": {
			input: pb.ListReq{Page: 1, Limit: 2},
			want: []pb.Post{
				{
					Id:    1,
					Title: "Title",
					Body:  "Tester body 1",
				},
				{
					Id:    2,
					Title: "Title 2",
					Body:  "Tester body 2",
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.ListPosts(context.Background(), &tc.input)
			for i := range got.Posts {
				if err != nil {
					t.Error("failed to Get Lists of task", err)
				}
				if !reflect.DeepEqual(tc.want[i], *got.Posts[i]) {
					t.Fatalf("%s: expected: %v, \ngot: %v", name, tc.want[i], *got.Posts[i])
				}
			}
		})
	}
}

func TestPostService_UpdatePosts(t *testing.T) {
	tests := map[string]struct {
		input pb.Post
		want  pb.Post
	}{
		"successful": {
			input: pb.Post{
				Id:    2,
				Title: "Title after update",
				Body:  "Body after update",
			},
			want: pb.Post{
				Id:    2,
				Title: "Title after update",
				Body:  "Body after update",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.UpdatePost(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to update task", err)
			}
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, \n\t\t\t\t\t\t\t\tgot: %v", name, tc.want, *got)
			}
		})
	}
}

func TestPostService_DeletePost(t *testing.T) {
	tests := map[string]struct {
		input pb.ByIdReq
		want  pb.EmptyResp
	}{
		"successful": {
			input: pb.ByIdReq{Id: 2},
			want:  pb.EmptyResp{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.DeletePost(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to Delete task", err)
			}
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v,\n got: %v", name, tc.want, *got)
			}
		})
	}
}
