package middlewares

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

type CookieCtx struct {
	UserID uuid.UUID `json:"user_id"`
	Token  string    `json:"token"`
}

func MyMiddleware(ctx huma.Context, next func(huma.Context)) {
	var cookie CookieCtx

	token, err := huma.ReadCookie(ctx, "token")
	if err != nil {
		fmt.Println(err)
	}
	userID, err := huma.ReadCookie(ctx, "user_id")
	if err != nil {
		fmt.Println(err)
	}

	cookie.Token = token.Value
	cookie.UserID, err = uuid.Parse(userID.Value)
	if err != nil {
		fmt.Println(err)
	}

	ctx = huma.WithValue(ctx, "cookie", cookie)

	next(ctx)
}
