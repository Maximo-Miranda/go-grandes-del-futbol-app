package match

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateMatchRequest struct {
	TournamentID string `form:"tournament_id" json:"tournament_id"`
	HomeTeamID   string `form:"home_team_id" json:"home_team_id"`
	AwayTeamID   string `form:"away_team_id" json:"away_team_id"`
	VenueID      string `form:"venue_id" json:"venue_id"`
	MatchDate    string `form:"match_date" json:"match_date"`
	Matchday     string `form:"matchday" json:"matchday"`
}

func (r *UpdateMatchRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateMatchRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"tournament_id": "required",
		"home_team_id":  "required",
		"away_team_id":  "required",
		"venue_id":      "required",
		"match_date":    "required",
		"matchday":      "required",
	}
}

func (r *UpdateMatchRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"tournament_id.required": "El torneo es obligatorio",
		"home_team_id.required":  "El equipo local es obligatorio",
		"away_team_id.required":  "El equipo visitante es obligatorio",
		"venue_id.required":      "La sede es obligatoria",
		"match_date.required":    "La fecha y hora del partido es obligatoria",
		"matchday.required":      "La jornada es obligatoria",
	}
}

func (r *UpdateMatchRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"tournament_id": "torneo",
		"home_team_id":  "equipo local",
		"away_team_id":  "equipo visitante",
		"venue_id":      "sede",
		"match_date":    "fecha del partido",
		"matchday":      "jornada",
	}
}

func (r *UpdateMatchRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
