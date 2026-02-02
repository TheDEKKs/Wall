package handlers

import (
	"context"
	"thedekk/WWT/internal/domains/comments"
	"thedekk/WWT/internal/domains/telegram"
	"thedekk/WWT/internal/domains/users"
	"thedekk/WWT/internal/domains/walls"
)

type WallHandler struct {
	wallService     *walls.WallService
	commentService  *comments.CommentsService
	telegramService *telegram.TelegramService
	userService *users.UserService

}

func NewWallHandler(wallService *walls.WallService, commentService *comments.CommentsService, telegramService *telegram.TelegramService, userService *users.UserService) *WallHandler {
	return &WallHandler{
		commentService:  commentService,
		wallService:     wallService,
		telegramService: telegramService,
		userService: userService,
	}
}

type Wall struct {
	Comments        []comments.Comments
	UserName        string `json:"user_name"`
	NumberOfRecords int    `json:"number_of_records"`

	UserNameTelegram  string `json:"user_name_telegram"`
	FirstNameTelegram string `json:"first_name_telegram"`
	LastNAmeTelegram  string `json:"last_name_telegram"`
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

	user, err := h.userService.GetTelegramIDUserByUserName(ctx, input.Wall)
	if err != nil {
		return nil, err
	}

	userTelegram, err := h.telegramService.GetUserByID(ctx, user.TelegramRecordID)
	if err != nil {
		return nil, err
	}

	wall := Wall{Comments: *comment, UserName: input.Wall, NumberOfRecords: index, UserNameTelegram: userTelegram.Username, FirstNameTelegram: userTelegram.FirstName, LastNAmeTelegram: userTelegram.LastName}

	return &CommentsWallOut{Body: wall}, nil
}
