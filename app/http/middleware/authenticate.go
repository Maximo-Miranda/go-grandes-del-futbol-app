package middleware

import (
	"errors"

	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/models"
)

func Authenticate() http.Middleware {
	return func(ctx http.Context) {
		token := extractBearerToken(ctx)
		if token == "" {
			redirectToLogin(ctx)
			return
		}

		payload, err := facades.Auth(ctx).Parse(token)
		if err != nil {
			if !errors.Is(err, auth.ErrorTokenExpired) {
				redirectToLogin(ctx)
				return
			}

			// Token expired — attempt silent refresh
			newToken, refreshErr := facades.Auth(ctx).Refresh()
			if refreshErr != nil || newToken == "" {
				redirectToLogin(ctx)
				return
			}

			// Refresh succeeded — set new cookie and re-parse for payload
			SetAuthCookie(ctx, newToken)
			payload, err = facades.Auth(ctx).Parse(newToken)
			if err != nil {
				redirectToLogin(ctx)
				return
			}
		}

		userID := payload.Key
		if userID == "" {
			redirectToLogin(ctx)
			return
		}

		var user models.User
		if err := facades.Orm().Query().Find(&user, userID); err != nil {
			redirectToLogin(ctx)
			return
		}

		if user.IsLocked() {
			redirectToLogin(ctx)
			return
		}

		ctx.WithValue("user", &user)
		ctx.WithValue("user_id", user.ID)
		ctx.Request().Next()
	}
}

func redirectToLogin(ctx http.Context) {
	clearAuthCookie(ctx)

	if ctx.Request().Header("X-Inertia", "") == "true" {
		ctx.Response().Header("X-Inertia-Location", "/auth/login")
		ctx.Response().String(http.StatusConflict, "").Abort()
		return
	}
	ctx.Response().Redirect(http.StatusFound, "/auth/login").Abort()
}

func clearAuthCookie(ctx http.Context) {
	ctx.Response().Cookie(http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})
}
