package handlers

import (
	"context"
	"thedekk/WWT/internal/domains/comments"
	"thedekk/WWT/internal/domains/walls"
)

type WallHandler struct {
	wallService    *walls.WallService
	commentService *comments.CommentsService
}

func NewWallHandler(wallService *walls.WallService, commentService *comments.CommentsService) *WallHandler {
	return &WallHandler{
		commentService: commentService,
		wallService:    wallService,
	}
}

type CommentsWallOut struct {
	Body []comments.Comments
}

func (h *WallHandler) GetCommentsWall(ctx context.Context, input *struct {
	Wall string `path:"wall"`
}) (*CommentsWallOut, error) {
	comment, err := h.commentService.GetCommentsWall(ctx, input.Wall)
	if err != nil {
		return nil, err
	}

	return &CommentsWallOut{Body: *comment}, nil
}
