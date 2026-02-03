package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/inertia"
	"grandesdelfutbol/app/models"
)

type StandingController struct {
	inertia *inertia.Inertia
}

func NewStandingController() *StandingController {
	return &StandingController{inertia: inertia.New()}
}

func (c *StandingController) Index(ctx http.Context) http.Response {
	tournamentID := ctx.Request().Query("tournament_id", "")

	query := facades.Orm().Query().With("Team").With("Group").With("Tournament")
	if tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}

	var standings []models.Standing
	query.Order("points DESC, goal_difference DESC, goals_for DESC").Find(&standings)

	var tournaments []models.Tournament
	facades.Orm().Query().Where("status != ?", "draft").Find(&tournaments)

	return c.inertia.Render(ctx, "standings/Index", map[string]any{
		"standings":    standings,
		"tournaments":  tournaments,
		"tournamentId": tournamentID,
	})
}
