package handlers

import (
	"context"
	"net/http"
	"thedekk/WWT/internal/domains/users"
	"thedekk/WWT/internal/transport/web/middlewares"

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

type UserResponse struct {
	Body string `json:"user_name"`
}

func (h *UserHandler) Check(ctx context.Context, input *struct{}) (*UserResponse, error) {
	cookie := ctx.Value("cookie").(middlewares.CookieCtx)

	user, err := h.userService.GetUserByUserID(ctx, cookie.UserID)
	if err != nil {
		return nil, err	
	}

	return  &UserResponse{Body: user.UserName}, nil

}
type ResetPasswordInput struct {
	Body struct {
		Password    string `json:"password"`
		NewPassword string `json:"new_password"`
	}
}

func (h *UserHandler) ResetPassword(ctx context.Context, input *ResetPasswordInput) (*UserCookieOut, error) {
	cookie := ctx.Value("cookie").(middlewares.CookieCtx)

	userCookie, err := h.userService.ResetPassword(ctx, cookie.UserID, input.Body.Password, input.Body.NewPassword)
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


type ResetUserNameInput struct {
	Body struct {
		UserName string `json:"new_user_name"`
	}
}

func (h *UserHandler) ResetUserName(ctx context.Context, input *ResetUserNameInput) (*struct{},error) {
	cookie := ctx.Value("cookie").(middlewares.CookieCtx)

	if err := h.userService.ResetUserName(ctx, cookie.UserID, input.Body.UserName); err != nil {
		return nil, err
	}

	return nil, nil 
}


type ResetTelegramInput struct {
	Body struct {
		TelegramID uuid.UUID `json:"telegram_id"`
	}
}

func (h *UserHandler) ResetTelegram(ctx context.Context, input *ResetTelegramInput) (*struct{},error)  {
	cookie := ctx.Value("cookie").(middlewares.CookieCtx)

	if err := h.userService.ResetTelegram(ctx, cookie.UserID, input.Body.TelegramID); err != nil {
		return nil, err
	}

	return nil, nil 
}
