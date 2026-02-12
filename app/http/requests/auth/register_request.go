package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type RegisterRequest struct {
	Name                 string `form:"name" json:"name"`
	Email                string `form:"email" json:"email"`
	Password             string `form:"password" json:"password"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation"`
}

func (r *RegisterRequest) Authorize(_ http.Context) error {
	return nil
}

func (r *RegisterRequest) Rules(_ http.Context) map[string]string {
	return map[string]string{
		"name":     "required|min_len:2",
		"email":    "required|email",
		"password": "required|min_len:8",
	}
}

func (r *RegisterRequest) Messages(_ http.Context) map[string]string {
	return map[string]string{
		"name.required":     "El nombre es obligatorio",
		"name.min_len":      "El nombre debe tener al menos 2 caracteres",
		"email.required":    "El correo electrónico es obligatorio",
		"email.email":       "Ingresa un correo electrónico válido",
		"password.required": "La contraseña es obligatoria",
		"password.min_len":  "La contraseña debe tener al menos 8 caracteres",
	}
}

func (r *RegisterRequest) Attributes(_ http.Context) map[string]string {
	return map[string]string{}
}

func (r *RegisterRequest) PrepareForValidation(_ http.Context, _ validation.Data) error {
	return nil
}
