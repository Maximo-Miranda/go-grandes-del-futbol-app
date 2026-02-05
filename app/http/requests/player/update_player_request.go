package player

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdatePlayerRequest struct {
	Name        string `form:"name" json:"name"`
	Nickname    string `form:"nickname" json:"nickname"`
	DocumentID  string `form:"document_id" json:"document_id"`
	Phone       string `form:"phone" json:"phone"`
	Position    string `form:"position" json:"position"`
	RemovePhoto string `form:"remove_photo" json:"remove_photo"`
}

func (r *UpdatePlayerRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdatePlayerRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":        "required|min_len:2|max_len:100",
		"nickname":    "max_len:50",
		"document_id": "max_len:20",
		"phone":       "max_len:20",
		"position":    "in:goalkeeper,defender,midfielder,forward,",
	}
}

func (r *UpdatePlayerRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":  "El nombre es obligatorio",
		"name.min_len":   "El nombre debe tener al menos 2 caracteres",
		"name.max_len":   "El nombre no puede exceder 100 caracteres",
		"nickname.max_len": "El apodo no puede exceder 50 caracteres",
		"document_id.max_len": "El documento no puede exceder 20 caracteres",
		"phone.max_len":  "El teléfono no puede exceder 20 caracteres",
		"position.in":    "La posición debe ser: portero, defensa, mediocampista o delantero",
	}
}

func (r *UpdatePlayerRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"name":        "nombre",
		"nickname":    "apodo",
		"document_id": "documento",
		"phone":       "teléfono",
		"position":    "posición",
	}
}

func (r *UpdatePlayerRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
