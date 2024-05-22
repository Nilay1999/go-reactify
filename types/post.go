package types

type CreatePost struct {
	Title  string `json:"title" binding:"required" required:"title is required"`
	Body   string `json:"body" binding:"required" required:"body is required"`
	UserId uint   `json:"userId" binding:"required" required:"userId is required"`
}

type Upvote struct {
	UserId uint `json:"userId" binding:"required" required:"userId is required"`
}
