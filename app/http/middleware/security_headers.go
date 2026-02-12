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

		if facades.Config().GetString("app.env") == "local" {
			ctx.Response().Header("Content-Security-Policy",
				"default-src 'self'; "+
					"script-src 'self' 'unsafe-inline' 'unsafe-eval' http://localhost:5173; "+
					"style-src 'self' 'unsafe-inline' http://localhost:5173; "+
					"font-src 'self' data: http://localhost:5173; "+
					"img-src 'self' data: blob:; "+
					"connect-src 'self' ws://localhost:5173 http://localhost:5173")
		} else {
			ctx.Response().Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
			ctx.Response().Header("Content-Security-Policy",
				"default-src 'self'; "+
					"script-src 'self'; "+
					"style-src 'self' 'unsafe-inline'; "+
					"font-src 'self' data:; "+
					"img-src 'self' data: blob:; "+
					"connect-src 'self'")
		}

		ctx.Request().Next()
	}
}
