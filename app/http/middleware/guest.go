package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Guest() http.Middleware {
	return func(ctx http.Context) {
		token := extractBearerToken(ctx)
		if token == "" {
			ctx.Request().Next()
			return
		}

		_, err := facades.Auth(ctx).Parse(token)
		if err != nil {
			ctx.Request().Next()
			return
		}

		if ctx.Request().Header("X-Inertia", "") == "true" {
			ctx.Response().Header("X-Inertia-Location", "/dashboard")
			ctx.Response().String(http.StatusConflict, "").Abort()
			return
		}

		ctx.Response().Redirect(http.StatusFound, "/dashboard").Abort()
	}
}
