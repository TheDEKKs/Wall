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
	SetCookie []http.Cookie `header:"Set-Cookie"`
}

func (h *UserHandler) RegistrationUser(ctx context.Context, input *RegistrationUserInput) (*UserCookieOut, error) {
	if input.Body.Password != input.Body.RepeatPassword {
		return nil, fmt.Errorf("Password no repeat")
	}

	userCookie, err := h.userService.RegistrationUser(ctx, input.Body.UserName, input.Body.Password)
	if err != nil {
		return nil, err
	}

	var userCookieOut UserCookieOut

	userCookieOut.SetCookie = append(userCookieOut.SetCookie, http.Cookie{
		Name:  "token",
		Value: (*userCookie)["Token"],
		MaxAge:   14 * 24 * 60 * 60, 
	}, http.Cookie{
		Name:  "user_id",
		Value: (*userCookie)["UserID"],
		MaxAge:   14 * 24 * 60 * 60,
	}, )

	return &userCookieOut, nil
}

type LoginUserInput struct {
	Body struct {
		UserName       string `json:"user_name"`
		Password       string `json:"password"`
	}
}

func (h *UserHandler) LoginUser(ctx context.Context, input *LoginUserInput) (*UserCookieOut, error) {
	userCookie, err := h.userService.LoginUser(ctx, input.Body.UserName, input.Body.Password)
	if err != nil {
		return nil, err
	}

	var userCookieOut UserCookieOut

	userCookieOut.SetCookie = append(userCookieOut.SetCookie, http.Cookie{
		Name:  "token",
		Value: (*userCookie)["Token"],
		MaxAge:   14 * 24 * 60 * 60, 
	}, http.Cookie{
		Name:  "user_id",
		Value: (*userCookie)["UserID"],
		MaxAge:   14 * 24 * 60 * 60, 
	}, )

	return &userCookieOut, nil

}
