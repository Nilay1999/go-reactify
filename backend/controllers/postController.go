package controllers

import (
	"net/http"
	"strconv"

	"github.com/Nilay1999/gin-gonic-server/services"
	"github.com/Nilay1999/gin-gonic-server/types"
	"github.com/gin-gonic/gin"
)

type PostController struct{}

var postService = new(services.Post)
var voteService = new(services.Vote)

func (p PostController) GetAllPost(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset := (page - 1) * limit

	users, err := postService.Get(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (p PostController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := postService.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (p PostController) CreatePost(ctx *gin.Context) {
	var payload types.CreatePost
	if validationError := ctx.BindJSON(&payload); validationError != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": customizer.DecryptErrors(validationError)})
		return
	}
	response, serviceError := postService.Create(payload)
	if serviceError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": serviceError.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (p PostController) Upvote(ctx *gin.Context) {
	var payload types.VotePost
	postID, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if validationError := ctx.BindJSON(&payload); validationError != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": customizer.DecryptErrors(validationError)})
		return
	}

	upvote, downvote, serviceError := voteService.Upvote(uint(postID), payload)
	if serviceError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": serviceError.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"upvote": upvote, "downvote": downvote})
}

func (p PostController) Downvote(ctx *gin.Context) {
	var payload types.VotePost
	postID, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if validationError := ctx.BindJSON(&payload); validationError != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": customizer.DecryptErrors(validationError)})
		return
	}

	upvote, downvote, serviceError := voteService.Downvote(uint(postID), payload)
	if serviceError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": serviceError.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"upvote": upvote, "downvote": downvote})
}
