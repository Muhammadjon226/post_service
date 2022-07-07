package service

import (
	pbFirst "github.com/Muhammadjon226/post_service/genproto/first_service"
	pbPost "github.com/Muhammadjon226/post_service/genproto/post_service"
)

// HelperFunction ...
func HelperFunction(posts []*pbFirst.PostResponse) []*pbPost.PostResponse {

	results := make([]*pbPost.PostResponse, 0, len(posts))

	for _, post := range posts {
		restult := pbPost.PostResponse{}

		restult.Id = post.Id
		restult.UserId = post.UserId
		restult.Body = post.Body
		restult.Title = post.Title
		restult.CreatedAt = post.CreatedAt
		restult.UpdatedAt = post.UpdatedAt

		results = append(results, &restult)
	}

	return results
}
