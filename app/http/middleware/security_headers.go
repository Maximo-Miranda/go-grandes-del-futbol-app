package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func SecurityHeaders() http.Middleware {
	return func(ctx http.Context) {
		ctx.Response().Header("X-Content-Type-Options", "nosniff")
		ctx.Response().Header("X-Frame-Options", "SAMEORIGIN")
		ctx.Response().Header("X-XSS-Protection", "1; mode=block")
		ctx.Response().Header("Referrer-Policy", "strict-origin-when-cross-origin")
		if facades.Config().GetString("app.env") != "local" {
			ctx.Response().Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		}
		ctx.Request().Next()
	}
}
