package handlers

import (
	"context"
	"thedekk/WWT/internal/domains/comments"
)

type CommentHandler struct {
	commentService *comments.CommentsService
}

func NewCommentHandler(commentService *comments.CommentsService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}


func (c *CommentHandler) NewComment(ctx context.Context, input *struct{}) (*struct{}, error) 
