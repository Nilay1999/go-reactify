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

	// Start a new transaction
	tx := initializers.Repository.Begin()
	if tx.Error != nil {
		return nil, nil, tx.Error
	}

	// Check if the upvote already exists
	upvoteExists := tx.
		Where("post_id = ? AND user_id = ? AND vote_type = ?", newVote.PostID, newVote.UserID, newVote.VoteType).
		First(&models.Vote{})

	// Check if a downvote exists
	downvoteExists := tx.
		Where("post_id = ? AND user_id = ? AND vote_type = ?", newVote.PostID, newVote.UserID, "downvote").
		First(&existingDownvote)

	// Remove existing downvote if it exists
	if downvoteExists.Error == nil {
		if err := tx.Unscoped().Delete(&existingDownvote).Error; err != nil {
			tx.Rollback()
			return nil, nil, err
		}
	}

	// Return error if the upvote already exists
	if upvoteExists.Error == nil {
		tx.Rollback()
		return nil, nil, errors.New("User already upvoted")
	} else if !errors.Is(upvoteExists.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, nil, upvoteExists.Error
	}

	// Create new upvote
	if err := tx.Create(&newVote).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	// Count upvotes
	if err := tx.
		Model(&models.Vote{}).
		Where("post_id = ? AND vote_type = ?", newVote.PostID, "upvote").
		Count(&upvoteCount).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	// Count downvotes
	if err := tx.
		Model(&models.Vote{}).
		Where("post_id = ? AND vote_type = ?", newVote.PostID, "downvote").
		Count(&downvoteCount).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
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

	// Start a new transaction
	tx := initializers.Repository.Begin()
	if tx.Error != nil {
		return nil, nil, tx.Error
	}

	// Check if the downvote already exists
	downvoteExists := tx.
		Where("post_id = ? AND user_id = ? AND vote_type = ?", newVote.PostID, newVote.UserID, newVote.VoteType).
		First(&models.Vote{})

	// Check if an upvote exists
	upvoteExists := tx.
		Where("post_id = ? AND user_id = ? AND vote_type = ?", newVote.PostID, newVote.UserID, "upvote").
		First(&existingUpvote)

	// Remove existing upvote if it exists
	if upvoteExists.Error == nil {
		if err := tx.Unscoped().Delete(&existingUpvote).Error; err != nil {
			tx.Rollback()
			return nil, nil, err
		}
	}

	// Return error if the downvote already exists
	if downvoteExists.Error == nil {
		tx.Rollback()
		return nil, nil, errors.New("User already downvoted")
	} else if !errors.Is(downvoteExists.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, nil, downvoteExists.Error
	}

	// Create new downvote
	if err := tx.Create(&newVote).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	// Count upvotes
	if err := tx.
		Model(&models.Vote{}).
		Where("post_id = ? AND vote_type = ?", newVote.PostID, "upvote").
		Count(&upvoteCount).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	// Count downvotes
	if err := tx.
		Model(&models.Vote{}).
		Where("post_id = ? AND vote_type = ?", newVote.PostID, "downvote").
		Count(&downvoteCount).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, nil, err
	}

	return &upvoteCount, &downvoteCount, nil
}
