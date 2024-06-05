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

func (v Vote) Upvote(postID uint, data types.VotePost) (*int64, *int64, error) {
	newVote := models.Vote{
		UserID:   data.UserId,
		PostID:   postID,
		VoteType: "upvote",
	}
	var upvoteCount, downvoteCount int64
	var existingDownvote models.Vote

	// Check if the upvote already exists
	upvoteExists := initializers.Repository.
		Where("post_id = ? AND user_id = ? AND vote_type = ?", newVote.PostID, newVote.UserID, newVote.VoteType).
		First(&models.Vote{})

	// Check if a downvote exists
	downvoteExists := initializers.Repository.
		Where("post_id = ? AND user_id = ? AND vote_type = ?", newVote.PostID, newVote.UserID, "downvote").
		First(&existingDownvote)

	// Remove existing downvote if it exists
	if downvoteExists.Error == nil {
		initializers.Repository.Unscoped().Delete(&existingDownvote)
	}

	// Return error if the upvote already exists
	if upvoteExists.Error == nil {
		return nil, nil, errors.New("User already upvoted")
	} else if !errors.Is(upvoteExists.Error, gorm.ErrRecordNotFound) {
		return nil, nil, upvoteExists.Error
	}

	// Create new upvote
	if err := initializers.Repository.Create(&newVote).Error; err != nil {
		return nil, nil, err
	}

	// Count upvotes
	if err := initializers.Repository.
		Model(&models.Vote{}).
		Where("post_id = ? AND vote_type = ?", newVote.PostID, "upvote").
		Count(&upvoteCount).Error; err != nil {
		return nil, nil, err
	}

	// Count downvotes
	if err := initializers.Repository.
		Model(&models.Vote{}).
		Where("post_id = ? AND vote_type = ?", newVote.PostID, "downvote").
		Count(&downvoteCount).Error; err != nil {
		return nil, nil, err
	}

	return &upvoteCount, &downvoteCount, nil
}

func (v Vote) Downvote(postID uint, data types.VotePost) (*int64, *int64, error) {
	newVote := models.Vote{
		UserID:   data.UserId,
		PostID:   postID,
		VoteType: "downvote",
	}
	var upvoteCount, downvoteCount int64
	var existingUpvote models.Vote

	// Check if the downvote already exists
	downvoteExists := initializers.Repository.
		Where("post_id = ? AND user_id = ? AND vote_type = ?", newVote.PostID, newVote.UserID, newVote.VoteType).
		First(&models.Vote{})

	// Check if an upvote exists
	upvoteExists := initializers.Repository.
		Where("post_id = ? AND user_id = ? AND vote_type = ?", newVote.PostID, newVote.UserID, "upvote").
		First(&existingUpvote)

	// Remove existing upvote if it exists
	if upvoteExists.Error == nil {
		initializers.Repository.Unscoped().Delete(&existingUpvote)
	}

	// Return error if the downvote already exists
	if downvoteExists.Error == nil {
		return nil, nil, errors.New("User already downvoted")
	} else if !errors.Is(downvoteExists.Error, gorm.ErrRecordNotFound) {
		return nil, nil, downvoteExists.Error
	}

	// Create new downvote
	if err := initializers.Repository.Create(&newVote).Error; err != nil {
		return nil, nil, err
	}

	// Count upvotes
	if err := initializers.Repository.
		Model(&models.Vote{}).
		Where("post_id = ? AND vote_type = ?", newVote.PostID, "upvote").
		Count(&upvoteCount).Error; err != nil {
		return nil, nil, err
	}

	// Count downvotes
	if err := initializers.Repository.
		Model(&models.Vote{}).
		Where("post_id = ? AND vote_type = ?", newVote.PostID, "downvote").
		Count(&downvoteCount).Error; err != nil {
		return nil, nil, err
	}

	return &upvoteCount, &downvoteCount, nil
}
