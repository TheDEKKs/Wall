package handlers

import (
	"context"
	"fmt"
	"net/http"
	"thedekk/WWT/internal/domains/users"
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
		UserName       string `json:"user_name"`
		Password       string `json:"password"`
		RepeatPassword string `json:"repeat_password"`
	}
}

type UserCookieOut struct {
	SetCookie http.Cookie `header:"Set-Cookie"`
}

func (h *UserHandler) RegistrationUser(ctx context.Context, input *RegistrationUserInput) (*UserCookieOut, error) {
	if input.Body.Password != input.Body.RepeatPassword {
		return nil, fmt.Errorf("Password no repeat")
	}

	token, err := h.userService.RegistrationUser(ctx, input.Body.UserName, input.Body.Password)
	if err != nil {
		return nil, err
	}

	return &UserCookieOut{SetCookie: http.Cookie{
		Name:  "token",
		Value: *token,
		MaxAge:   14 * 24 * 60 * 60, 
	}}, nil
}

type LoginUserInput struct {
	Body struct {
		UserName       string `json:"user_name"`
		Password       string `json:"password"`
	}
}

func (h *UserHandler) LoginUser(ctx context.Context, input *LoginUserInput) (*UserCookieOut, error) {
	token, err := h.userService.LoginUser(ctx, input.Body.UserName, input.Body.Password)
	if err != nil {
		return nil, err
	}

	return &UserCookieOut{SetCookie: http.Cookie{
		Name:  "token",
		Value: *token,
		MaxAge:   14 * 24 * 60 * 60, 
	}}, nil
}
