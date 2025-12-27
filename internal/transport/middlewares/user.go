package middlewares

import (

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

type CookieCtx struct {
	UserID uuid.UUID `json:"user_id"`
	Token  string    `json:"token"`
}

func MyMiddleware(ctx huma.Context, next func(huma.Context)) {
	var cookie CookieCtx

	token, err0 := huma.ReadCookie(ctx, "token")

	userID, err1 := huma.ReadCookie(ctx, "user_id")
	
	if err0 == nil && err1 == nil {
		cookie.Token = token.Value
		cookie.UserID, _ = uuid.Parse(userID.Value)

		ctx = huma.WithValue(ctx, "cookie", cookie)

		next(ctx)

	}

	next(ctx)
}
