package handlers

import (
	"context"
	"net/http"
	"thedekk/WWT/internal/domains/users"

	"github.com/google/uuid"
)

type UserHandler struct {
	userService *users.UserService
}

func NewUserHandler(userService *users.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type RegistrationUserInput struct {
	Body struct {
		UserName   string    `json:"user_name"`
		Password   string    `json:"password"`
		TelegramID uuid.UUID `json:"telegram_id"`
	}
}

type UserCookieOut struct {
	SetCookie []http.Cookie `header:"Set-Cookie"`
}

func (h *UserHandler) RegistrationUser(ctx context.Context, input *RegistrationUserInput) (*UserCookieOut, error) {
	userCookie, err := h.userService.RegistrationUser(ctx, input.Body.UserName, input.Body.Password, input.Body.TelegramID)
	if err != nil {
		return nil, err
	}

	var userCookieOut UserCookieOut

	userCookieOut.SetCookie = append(userCookieOut.SetCookie, http.Cookie{
		Name:   "token",
		Value:  (*userCookie)["Token"],
		MaxAge: 14 * 24 * 60 * 60,
	}, http.Cookie{
		Name:   "user_id",
		Value:  (*userCookie)["UserID"],
		MaxAge: 14 * 24 * 60 * 60,
	})

	return &userCookieOut, nil
}

type LoginUserInput struct {
	Body struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
}

func (h *UserHandler) LoginUser(ctx context.Context, input *LoginUserInput) (*UserCookieOut, error) {
	userCookie, err := h.userService.LoginUser(ctx, input.Body.UserName, input.Body.Password)
	if err != nil {
		return nil, err
	}

	var userCookieOut UserCookieOut

	userCookieOut.SetCookie = append(userCookieOut.SetCookie, http.Cookie{
		Name:   "token",
		Value:  (*userCookie)["Token"],
		MaxAge: 14 * 24 * 60 * 60,
	}, http.Cookie{
		Name:   "user_id",
		Value:  (*userCookie)["UserID"],
		MaxAge: 14 * 24 * 60 * 60,
	})

	return &userCookieOut, nil

}
