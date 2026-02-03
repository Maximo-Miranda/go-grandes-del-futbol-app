package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type LoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *LoginRequest) Authorize(_ http.Context) error {
	return nil
}

func (r *LoginRequest) Rules(_ http.Context) map[string]string {
	return map[string]string{
		"email":    "required|email",
		"password": "required|min_len:6",
	}
}

func (r *LoginRequest) Messages(_ http.Context) map[string]string {
	return map[string]string{
		"email.required":    "El correo electrónico es obligatorio",
		"email.email":       "Ingresa un correo electrónico válido",
		"password.required": "La contraseña es obligatoria",
		"password.min_len":  "La contraseña debe tener al menos 6 caracteres",
	}
}

func (r *LoginRequest) Attributes(_ http.Context) map[string]string {
	return map[string]string{}
}

func (r *LoginRequest) PrepareForValidation(_ http.Context, _ validation.Data) error {
	return nil
}
