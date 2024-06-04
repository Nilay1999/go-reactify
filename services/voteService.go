package services

import (
	"errors"

	"github.com/Nilay1999/gin-gonic-server/initializers"
	"github.com/Nilay1999/gin-gonic-server/models"
	"github.com/Nilay1999/gin-gonic-server/types"
	"gorm.io/gorm"
)

type Vote struct {
	models.Vote
}

func (v Vote) Upvote(postId uint, data types.VotePost) (*int64, *int64, error) {
	vote := Vote{
		Vote: models.Vote{
			UserID:   data.UserId,
			PostID:   postId,
			VoteType: "upvote",
		},
	}
	var upvoteCount int64
	var downvoteCount int64

	upvoteExists := initializers.Repository.
		Where("post_id = ? AND user_id = ? AND vote_type = ?", vote.PostID, vote.UserID, vote.VoteType).
		First(&vote)

	if upvoteExists.Error == nil {
		return nil, nil, errors.New("User already upvoted")
	} else if !errors.Is(upvoteExists.Error, gorm.ErrRecordNotFound) {
		return nil, nil, upvoteExists.Error
	}

	upvote := initializers.Repository.Create(&vote)
	if upvote.Error != nil {
		return nil, nil, upvote.Error
	}

	upvoteCountQuery := initializers.Repository.
		Model(&v.Vote).
		Where("post_id = ? AND vote_type = ?", vote.PostID, "upvote").
		Count(&upvoteCount)

	downvoteCountQuery := initializers.Repository.
		Model(&v.Vote).
		Where("post_id = ? AND vote_type = ?", vote.PostID, "downvote").
		Count(&downvoteCount)

	if upvoteCountQuery.Error != nil {
		return nil, nil, upvoteCountQuery.Error
	}
	if downvoteCountQuery.Error != nil {
		return nil, nil, downvoteCountQuery.Error
	}

	return &upvoteCount, &downvoteCount, nil
}

func (v Vote) Downvote(postId uint, data types.VotePost) (*int64, *int64, error) {
	vote := Vote{
		Vote: models.Vote{
			UserID:   data.UserId,
			PostID:   postId,
			VoteType: "downvote",
		},
	}
	var upvoteCount int64
	var downvoteCount int64

	upvoteExists := initializers.Repository.
		Where("post_id = ? AND user_id = ? AND vote_type = ?", vote.PostID, vote.UserID, vote.VoteType).
		First(&vote)

	if upvoteExists.Error == nil {
		return nil, nil, errors.New("User already downvoted")
	} else if !errors.Is(upvoteExists.Error, gorm.ErrRecordNotFound) {
		return nil, nil, upvoteExists.Error
	}

	upvote := initializers.Repository.Create(&vote)
	if upvote.Error != nil {
		return nil, nil, upvote.Error
	}

	upvoteCountQuery := initializers.Repository.
		Model(&v.Vote).
		Where("post_id = ? AND vote_type = ?", vote.PostID, "upvote").
		Count(&upvoteCount)

	downvoteCountQuery := initializers.Repository.
		Model(&v.Vote).
		Where("post_id = ? AND vote_type = ?", vote.PostID, "downvote").
		Count(&downvoteCount)

	if upvoteCountQuery.Error != nil {
		return nil, nil, upvoteCountQuery.Error
	}
	if downvoteCountQuery.Error != nil {
		return nil, nil, downvoteCountQuery.Error
	}

	return &upvoteCount, &downvoteCount, nil
}
