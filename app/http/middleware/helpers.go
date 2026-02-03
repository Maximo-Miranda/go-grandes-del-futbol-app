package middleware

import "github.com/goravel/framework/contracts/http"

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
