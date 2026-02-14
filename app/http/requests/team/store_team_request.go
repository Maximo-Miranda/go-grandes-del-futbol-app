package team

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type StoreTeamRequest struct {
	Name         string `form:"name" json:"name"`
	Color        string `form:"color" json:"color"`
	ContactPhone string `form:"contact_phone" json:"contact_phone"`
}

func (r *StoreTeamRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreTeamRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":          "required|min_len:2|max_len:100",
		"color":         "required|max_len:20",
		"contact_phone": "required|max_len:20",
		"logo":          "file|image|max_file_size:1024",
	}
}

func (r *StoreTeamRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":          "El nombre del equipo es obligatorio",
		"name.min_len":           "El nombre debe tener al menos 2 caracteres",
		"name.max_len":           "El nombre no puede exceder 100 caracteres",
		"color.required":         "El color del equipo es obligatorio",
		"color.max_len":          "El color no puede exceder 20 caracteres",
		"contact_phone.required": "El teléfono de contacto es obligatorio",
		"contact_phone.max_len":  "El teléfono no puede exceder 20 caracteres",
		"logo.image":             "El logo debe ser una imagen válida",
		"logo.max_file_size":     "El logo no puede exceder 1 MB",
	}
}

func (r *StoreTeamRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"name":          "nombre",
		"color":         "color",
		"contact_phone": "teléfono de contacto",
		"logo":          "logo",
	}
}

func (r *StoreTeamRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
