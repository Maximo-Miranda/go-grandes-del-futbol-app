package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/http/requests/auth"
	"grandesdelfutbol/app/inertia"
	"grandesdelfutbol/app/models"
	authService "grandesdelfutbol/app/services/auth"
)

type AuthController struct {
	inertia     *inertia.Inertia
	authService authService.AuthServiceInterface
}

func NewAuthController(service authService.AuthServiceInterface) *AuthController {
	return &AuthController{
		inertia:     inertia.New(),
		authService: service,
	}
}

func (c *AuthController) ShowLogin(ctx http.Context) http.Response {
	return c.inertia.Render(ctx, "auth/Login", map[string]any{})
}

func (c *AuthController) Login(ctx http.Context) http.Response {
	var request auth.LoginRequest
	validationErrors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "auth/Login", map[string]any{
			"errors": map[string]string{"email": "Error de validación"},
		})
	}
	if validationErrors != nil {
		return c.inertia.Render(ctx, "auth/Login", map[string]any{
			"errors": inertia.ValidationErrors(validationErrors.All()),
		})
	}

	result := c.authService.Login(ctx, request.Email, request.Password, ctx.Request().Ip())
	if !result.Success {
		return c.inertia.Render(ctx, "auth/Login", map[string]any{
			"errors": map[string]string{"email": result.Error},
		})
	}

	ctx.Response().Cookie(http.Cookie{
		Name:     "token",
		Value:    result.Token,
		MaxAge:   60 * 60 * 12,
		Path:     "/",
		HttpOnly: true,
		Secure:   facades.Config().GetString("app.env") != "local",
		SameSite: "Lax",
	})

	return ctx.Response().Redirect(http.StatusFound, "/dashboard")
}

func (c *AuthController) ShowRegister(ctx http.Context) http.Response {
	return c.inertia.Render(ctx, "auth/Register", map[string]any{})
}

func (c *AuthController) Register(ctx http.Context) http.Response {
	var request auth.RegisterRequest
	validationErrors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "auth/Register", map[string]any{
			"errors": map[string]string{"email": "Error de validación"},
		})
	}
	if validationErrors != nil {
		return c.inertia.Render(ctx, "auth/Register", map[string]any{
			"errors": inertia.ValidationErrors(validationErrors.All()),
		})
	}

	result := c.authService.Register(request.Name, request.Email, request.Password)
	if !result.Success {
		return c.inertia.Render(ctx, "auth/Register", map[string]any{
			"errors": map[string]string{"email": result.Error},
		})
	}

	return c.inertia.Render(ctx, "auth/Login", map[string]any{
		"success": "Cuenta creada exitosamente. Ya puedes iniciar sesión.",
	})
}

func (c *AuthController) Logout(ctx http.Context) http.Response {
	user, ok := ctx.Value("user").(*models.User)
	if ok && user != nil {
		_ = c.authService.Logout(ctx, user.ID)
	}

	ctx.Response().Cookie(http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})

	return ctx.Response().Redirect(http.StatusFound, "/auth/login")
}
