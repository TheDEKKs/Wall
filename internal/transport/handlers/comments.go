package handlers

import (
	"context"
	"fmt"
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

type NewCommentInput struct {
	Wall string `path:"wall"`
	Body struct {
		Comment string `json:"comment"`
	}
}

func (h *CommentHandler) NewComment(ctx context.Context, input *NewCommentInput) (*struct{}, error) {
	if len(input.Body.Comment) > 500 {
		return nil, fmt.Errorf("Error comment very big")
	} else if len(input.Body.Comment) <= 0 {
		return nil, fmt.Errorf("Comment is null")
	}

	if err := h.commentService.NewComment(ctx, input.Wall, input.Body.Comment); err != nil {
		return nil, err
	}

	return nil, nil
}
