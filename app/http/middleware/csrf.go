package middleware

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

const (
	csrfTokenLength = 32
	csrfCookieName  = "XSRF-TOKEN"
	csrfHeaderName  = "X-XSRF-TOKEN"
)

var safeMethods = map[string]bool{
	"GET": true, "HEAD": true, "OPTIONS": true,
}

func CSRF() http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Cookie(csrfCookieName, "")
		if token == "" {
			token = generateCSRFToken()
		}

		if !safeMethods[ctx.Request().Method()] {
			headerToken := ctx.Request().Header(csrfHeaderName, "")
			formToken := ctx.Request().Input("_token", "")
			providedToken := headerToken
			if providedToken == "" {
				providedToken = formToken
			}
			if !validateCSRFToken(token, providedToken) {
				if ctx.Request().Header("X-Inertia", "") == "true" {
					ctx.Response().Header("X-Inertia-Location", ctx.Request().Path())
					ctx.Response().String(http.StatusConflict, "CSRF token mismatch").Abort()
					return
				}
				ctx.Response().Json(http.StatusForbidden, http.Json{
					"error":   "CSRF token mismatch",
					"message": "La sesión ha expirado. Por favor recarga la página.",
				}).Abort()
				return
			}

			// Rotate token after every successful mutation (POST/PUT/DELETE)
			token = generateCSRFToken()
		}

		setCSRFCookie(ctx, token)
		ctx.WithValue("csrf_token", token)
		ctx.Request().Next()
	}
}

func generateCSRFToken() string {
	bytes := make([]byte, csrfTokenLength)
	if _, err := rand.Read(bytes); err != nil {
		facades.Log().Error("Failed to generate CSRF token: " + err.Error())
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

func setCSRFCookie(ctx http.Context, token string) {
	ctx.Response().Cookie(http.Cookie{
		Name:     csrfCookieName,
		Value:    token,
		Path:     "/",
		MaxAge:   60 * 60 * 2,
		HttpOnly: false,
		Secure:   facades.Config().GetString("app.env") != "local",
		SameSite: "Lax",
	})
}

func validateCSRFToken(expected, provided string) bool {
	if expected == "" || provided == "" {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(expected), []byte(provided)) == 1
}
