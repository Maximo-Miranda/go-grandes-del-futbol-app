package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/inertia"
	"grandesdelfutbol/app/models"
)

type DashboardController struct {
	inertia *inertia.Inertia
}

func NewDashboardController() *DashboardController {
	return &DashboardController{inertia: inertia.New()}
}

func (c *DashboardController) Index(ctx http.Context) http.Response {
	totalTournaments, _ := facades.Orm().Query().Model(&models.Tournament{}).Count()
	totalTeams, _ := facades.Orm().Query().Model(&models.Team{}).Count()
	totalMatches, _ := facades.Orm().Query().Model(&models.Match{}).Count()
	totalPlayers, _ := facades.Orm().Query().Model(&models.Player{}).Count()

	var upcomingMatches []models.Match
	facades.Orm().Query().
		With("HomeTeam").With("AwayTeam").With("Tournament").
		Where("status = ?", "scheduled").
		Order("match_date ASC").
		Limit(5).
		Find(&upcomingMatches)

	return c.inertia.Render(ctx, "Dashboard", map[string]any{
		"stats": map[string]any{
			"tournaments": totalTournaments,
			"teams":       totalTeams,
			"matches":     totalMatches,
			"players":     totalPlayers,
		},
		"upcomingMatches": upcomingMatches,
	})
}
