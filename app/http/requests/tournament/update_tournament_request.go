package tournament

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateTournamentRequest struct {
	Name        string `form:"name" json:"name"`
	Format      string `form:"format" json:"format"`
	GameType    string `form:"game_type" json:"game_type"`
	Status      string `form:"status" json:"status"`
	StartDate   string `form:"start_date" json:"start_date"`
	EndDate     string `form:"end_date" json:"end_date"`
	VenueID     string `form:"venue_id" json:"venue_id"`
	Description string `form:"description" json:"description"`
}

func (r *UpdateTournamentRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateTournamentRequest) Rules(ctx http.Context) map[string]string {
	rules := map[string]string{
		"name":       "required|min_len:2|max_len:100",
		"format":     "required|in:round_robin,knockout,group_knockout",
		"game_type":  "required|in:5v5,7v7,11v11",
		"status":     "required|in:draft,active,completed,cancelled",
		"venue_id":   "required",
		"start_date": "required|date",
	}

	if r.StartDate != "" {
		rules["end_date"] = "required|date|gte_date:" + r.StartDate
	} else {
		rules["end_date"] = "required|date"
	}

	return rules
}

func (r *UpdateTournamentRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":       "El nombre del torneo es obligatorio",
		"name.min_len":        "El nombre debe tener al menos 2 caracteres",
		"name.max_len":        "El nombre no puede exceder 100 caracteres",
		"format.required":     "El formato es obligatorio",
		"format.in":           "El formato debe ser: todos contra todos, eliminación directa o grupos + eliminación",
		"game_type.required":  "El tipo de juego es obligatorio",
		"game_type.in":        "El tipo de juego debe ser: 5v5, 7v7 o 11v11",
		"status.required":     "El estado es obligatorio",
		"status.in":           "El estado debe ser: borrador, activo, finalizado o cancelado",
		"venue_id.required":   "La sede es obligatoria",
		"start_date.required": "La fecha de inicio es obligatoria",
		"start_date.date":     "La fecha de inicio debe ser una fecha válida",
		"end_date.required":   "La fecha de fin es obligatoria",
		"end_date.date":       "La fecha de fin debe ser una fecha válida",
		"end_date.gte_date":   "La fecha de fin no puede ser anterior a la fecha de inicio",
	}
}

func (r *UpdateTournamentRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"name":       "nombre",
		"format":     "formato",
		"game_type":  "tipo de juego",
		"status":     "estado",
		"start_date": "fecha de inicio",
		"end_date":   "fecha de fin",
		"venue_id":   "sede",
	}
}

func (r *UpdateTournamentRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
