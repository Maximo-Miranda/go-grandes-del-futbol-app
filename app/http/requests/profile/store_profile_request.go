package profile

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type StoreProfileRequest struct {
	Nickname   string `form:"nickname" json:"nickname"`
	DocumentID string `form:"document_id" json:"document_id"`
	Phone      string `form:"phone" json:"phone"`
	Position   string `form:"position" json:"position"`
}

func (r *StoreProfileRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreProfileRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"nickname":    "required|min_len:2|max_len:50",
		"document_id": "required|max_len:20",
		"phone":       "required|max_len:20",
		"position":    "required|in:goalkeeper,defender,midfielder,forward",
	}
}

func (r *StoreProfileRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"nickname.required":    "El apodo es obligatorio",
		"nickname.min_len":     "El apodo debe tener al menos 2 caracteres",
		"nickname.max_len":     "El apodo no puede exceder 50 caracteres",
		"document_id.required": "El documento de identidad es obligatorio",
		"document_id.max_len":  "El documento no puede exceder 20 caracteres",
		"phone.required":       "El teléfono es obligatorio",
		"phone.max_len":        "El teléfono no puede exceder 20 caracteres",
		"position.required":    "La posición es obligatoria",
		"position.in":          "La posición debe ser: portero, defensa, mediocampista o delantero",
	}
}

func (r *StoreProfileRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"nickname":    "apodo",
		"document_id": "documento de identidad",
		"phone":       "teléfono",
		"position":    "posición",
	}
}

func (r *StoreProfileRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
