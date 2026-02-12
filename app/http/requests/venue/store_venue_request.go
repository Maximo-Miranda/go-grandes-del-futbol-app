package venue

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type StoreVenueRequest struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func (r *StoreVenueRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreVenueRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":    "required|min_len:2|max_len:100",
		"address": "required|min_len:2|max_len:255",
		"photo":   "file|image|max_file_size:1024",
	}
}

func (r *StoreVenueRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":       "El nombre es obligatorio",
		"name.min_len":        "El nombre debe tener al menos 2 caracteres",
		"name.max_len":        "El nombre no puede exceder 100 caracteres",
		"address.required":    "La dirección es obligatoria",
		"address.min_len":     "La dirección debe tener al menos 2 caracteres",
		"address.max_len":     "La dirección no puede exceder 255 caracteres",
		"photo.image":         "La foto debe ser una imagen válida",
		"photo.max_file_size": "La foto no puede exceder 1 MB",
	}
}

func (r *StoreVenueRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"name":    "nombre",
		"address": "dirección",
		"photo":   "foto",
	}
}

func (r *StoreVenueRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
