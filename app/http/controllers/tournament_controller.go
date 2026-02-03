package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/inertia"
	"grandesdelfutbol/app/models"
)

type TournamentController struct {
	inertia *inertia.Inertia
}

func NewTournamentController() *TournamentController {
	return &TournamentController{inertia: inertia.New()}
}

func (c *TournamentController) Index(ctx http.Context) http.Response {
	var tournaments []models.Tournament
	facades.Orm().Query().With("Venue").Order("created_at DESC").Find(&tournaments)

	return c.inertia.Render(ctx, "tournaments/Index", map[string]any{
		"tournaments": tournaments,
	})
}

func (c *TournamentController) Create(ctx http.Context) http.Response {
	var venues []models.Venue
	facades.Orm().Query().Find(&venues)

	return c.inertia.Render(ctx, "tournaments/Create", map[string]any{
		"venues": venues,
	})
}

func (c *TournamentController) Store(ctx http.Context) http.Response {
	tournament := models.Tournament{
		Name:        ctx.Request().Input("name"),
		Description: ctx.Request().Input("description"),
		Format:      ctx.Request().Input("format", "round_robin"),
		GameType:    ctx.Request().Input("game_type", "5v5"),
		Status:      "draft",
	}

	if err := facades.Orm().Query().Create(&tournament); err != nil {
		return c.inertia.Render(ctx, "tournaments/Create", map[string]any{
			"errors": map[string]string{"name": "Error al crear el torneo"},
		})
	}

	return ctx.Response().Redirect(http.StatusFound, "/tournaments")
}

func (c *TournamentController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var tournament models.Tournament
	if err := facades.Orm().Query().With("Venue").With("Groups").Find(&tournament, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/tournaments")
	}

	var teams []models.TournamentTeam
	facades.Orm().Query().With("Team").With("Group").Where("tournament_id = ?", id).Find(&teams)

	var matches []models.Match
	facades.Orm().Query().With("HomeTeam").With("AwayTeam").Where("tournament_id = ?", id).Order("match_date ASC").Find(&matches)

	var standings []models.Standing
	facades.Orm().Query().With("Team").With("Group").Where("tournament_id = ?", id).Order("points DESC, goal_difference DESC").Find(&standings)

	return c.inertia.Render(ctx, "tournaments/Show", map[string]any{
		"tournament": tournament,
		"teams":      teams,
		"matches":    matches,
		"standings":  standings,
	})
}

func (c *TournamentController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var tournament models.Tournament
	if err := facades.Orm().Query().Find(&tournament, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/tournaments")
	}

	var venues []models.Venue
	facades.Orm().Query().Find(&venues)

	return c.inertia.Render(ctx, "tournaments/Edit", map[string]any{
		"tournament": tournament,
		"venues":     venues,
	})
}

func (c *TournamentController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var tournament models.Tournament
	if err := facades.Orm().Query().Find(&tournament, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/tournaments")
	}

	tournament.Name = ctx.Request().Input("name", tournament.Name)
	tournament.Description = ctx.Request().Input("description", tournament.Description)
	tournament.Format = ctx.Request().Input("format", tournament.Format)
	tournament.GameType = ctx.Request().Input("game_type", tournament.GameType)

	facades.Orm().Query().Save(&tournament)

	return ctx.Response().Redirect(http.StatusFound, "/tournaments/"+id)
}

func (c *TournamentController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var tournament models.Tournament
	if err := facades.Orm().Query().Find(&tournament, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/tournaments")
	}
	facades.Orm().Query().Delete(&tournament)
	return ctx.Response().Redirect(http.StatusFound, "/tournaments")
}

func (c *TournamentController) UpdateStatus(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var tournament models.Tournament
	if err := facades.Orm().Query().Find(&tournament, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/tournaments")
	}

	tournament.Status = ctx.Request().Input("status", tournament.Status)
	facades.Orm().Query().Save(&tournament)

	return ctx.Response().Redirect(http.StatusFound, "/tournaments/"+id)
}
