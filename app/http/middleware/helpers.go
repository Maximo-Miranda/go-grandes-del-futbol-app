package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

const authCookieMaxAge = 60 * 60 * 24 * 30 // 30 days

func extractBearerToken(ctx http.Context) string {
	token := ctx.Request().Cookie("token", "")
	if token == "" {
		authHeader := ctx.Request().Header("Authorization", "")
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		}
	}
	return token
}

func SetAuthCookie(ctx http.Context, token string) {
	ctx.Response().Cookie(http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   authCookieMaxAge,
		Path:     "/",
		HttpOnly: true,
		Secure:   facades.Config().GetString("app.env") != "local",
		SameSite: "Lax",
	})
}
