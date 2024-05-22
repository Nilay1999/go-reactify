package services

import (
	"github.com/Nilay1999/gin-gonic-server/initializers"
	"github.com/Nilay1999/gin-gonic-server/models"
	"github.com/Nilay1999/gin-gonic-server/types"
)

type Post struct {
	models.Post
}
type Vote struct {
	models.Vote
}

func (p Post) Create(payload types.CreatePost) (*Post, error) {
	post := Post{
		Post: models.Post{
			Title:  payload.Title,
			Body:   payload.Body,
			UserID: payload.UserId,
		},
	}
	result := initializers.Repository.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}

func (v Vote) Upvote(postId uint, payload types.Upvote) (*Post, error) {
	var vote Vote
	result := initializers.Repository.Where("post_id = ? AND user_id = ?", vote.PostID, vote.UserID).First(&vote)

	if result.Error != nil {
		return nil, result.Error
	}
}
