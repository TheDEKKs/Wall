package handlers

import (
	"context"
	"fmt"
	"thedekk/WWT/internal/domains/comments"
	"thedekk/WWT/internal/transport/middlewares"
)

type CommentHandler struct {
	commentService *comments.CommentsService
}

func NewCommentHandler(commentService *comments.CommentsService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}


func (c *CommentHandler) NewComment(ctx context.Context, input *struct{}) (*struct{}, error) {
	token := ctx.Value("cookie").(middlewares.CookieCtx)

	fmt.Println(token.Token, token.UserID)

	return nil, nil
}
