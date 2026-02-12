package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/models"
)

func RequireProfile() http.Middleware {
	return func(ctx http.Context) {
		user, ok := ctx.Value("user").(*models.User)
		if !ok || user == nil {
			redirectToLogin(ctx)
			return
		}

		var player models.Player
		if err := facades.Orm().Query().Where("user_id = ?", user.ID).First(&player); err != nil || player.ID == 0 {
			redirectToProfileCreate(ctx)
			return
		}

		ctx.WithValue("player", &player)
		ctx.Request().Next()
	}
}

func redirectToProfileCreate(ctx http.Context) {
	if ctx.Request().Header("X-Inertia", "") == "true" {
		ctx.Response().Header("X-Inertia-Location", "/profile/create")
		ctx.Response().String(http.StatusConflict, "").Abort()
		return
	}
	ctx.Response().Redirect(http.StatusSeeOther, "/profile/create").Abort()
}
