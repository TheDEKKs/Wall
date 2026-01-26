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

type Wall struct {
	Comments        []comments.Comments
	UserName        string `json:"user_name"`
	NumberOfRecords int    `json:"number_of_records"`
}

type CommentsWallOut struct {
	Body Wall
}

func (h *WallHandler) GetWall(ctx context.Context, input *struct {
	Wall string `path:"wall"`
}) (*CommentsWallOut, error) {
	comment, index, err := h.commentService.GetCommentsWall(ctx, input.Wall)
	if err != nil {
		return nil, err
	}

	wall := Wall{Comments: *comment, UserName: input.Wall, NumberOfRecords: index}

	return &CommentsWallOut{Body: wall}, nil
}
