package services

import (
	"github.com/Nilay1999/gin-gonic-server/initializers"
	"github.com/Nilay1999/gin-gonic-server/models"
	"github.com/Nilay1999/gin-gonic-server/types"
)

type Post struct {
	models.Post
}

func (p Post) Create(data types.CreatePost) (*Post, error) {
	post := Post{
		Post: models.Post{
			Title:  data.Title,
			Body:   data.Body,
			UserID: data.UserId,
		},
	}
	result := initializers.Repository.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	if err := initializers.Repository.Preload("User").First(&post, post.ID).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (p Post) Get(offset int, limit int) ([]Post, error) {
	var posts []Post
	result := initializers.Repository.Preload("User").Limit(limit).Offset(offset).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (p Post) GetById(id string) (*Post, error) {
	var post Post
	result := initializers.Repository.Preload("User").First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}
