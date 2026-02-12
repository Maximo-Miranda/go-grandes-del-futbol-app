package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/inertia"
	"grandesdelfutbol/app/models"
)

type TeamController struct {
	inertia *inertia.Inertia
}

func NewTeamController() *TeamController {
	return &TeamController{inertia: inertia.New()}
}

func (c *TeamController) Index(ctx http.Context) http.Response {
	var teams []models.Team
	facades.Orm().Query().Order("name ASC").Find(&teams)

	return c.inertia.Render(ctx, "teams/Index", map[string]any{
		"teams": teams,
	})
}

func (c *TeamController) Create(ctx http.Context) http.Response {
	return c.inertia.Render(ctx, "teams/Create", map[string]any{})
}

func (c *TeamController) Store(ctx http.Context) http.Response {
	team := models.Team{
		Name:         ctx.Request().Input("name"),
		Color:        ctx.Request().Input("color"),
		ContactPhone: ctx.Request().Input("contact_phone"),
	}

	if err := facades.Orm().Query().Create(&team); err != nil {
		return c.inertia.Render(ctx, "teams/Create", map[string]any{
			"errors": map[string]string{"name": "Error al crear el equipo"},
		})
	}

	return ctx.Response().Redirect(http.StatusSeeOther, "/teams")
}

func (c *TeamController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/teams")
	}

	var teamPlayers []models.TeamPlayer
	facades.Orm().Query().With("Player.User").Where("team_id = ?", id).Find(&teamPlayers)

	return c.inertia.Render(ctx, "teams/Show", map[string]any{
		"team":    team,
		"players": teamPlayers,
	})
}

func (c *TeamController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/teams")
	}

	return c.inertia.Render(ctx, "teams/Edit", map[string]any{
		"team": team,
	})
}

func (c *TeamController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/teams")
	}

	team.Name = ctx.Request().Input("name", team.Name)
	team.Color = ctx.Request().Input("color", team.Color)
	team.ContactPhone = ctx.Request().Input("contact_phone", team.ContactPhone)

	facades.Orm().Query().Save(&team)
	return ctx.Response().Redirect(http.StatusSeeOther, "/teams/"+id)
}

func (c *TeamController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/teams")
	}
	facades.Orm().Query().Delete(&team)
	return ctx.Response().Redirect(http.StatusSeeOther, "/teams")
}

func (c *TeamController) AddPlayer(ctx http.Context) http.Response {
	teamID := ctx.Request().Route("id")
	playerID := ctx.Request().Input("player_id")

	var team models.Team
	if err := facades.Orm().Query().Find(&team, teamID); err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{"error": "Equipo no encontrado"})
	}

	var playerIDUint uint
	for _, c := range playerID {
		if c >= '0' && c <= '9' {
			playerIDUint = playerIDUint*10 + uint(c-'0')
		}
	}

	teamPlayer := models.TeamPlayer{
		TeamID:   team.ID,
		PlayerID: playerIDUint,
	}

	if err := facades.Orm().Query().Create(&teamPlayer); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": "Error al agregar jugador"})
	}

	return ctx.Response().Redirect(http.StatusSeeOther, "/teams/"+teamID)
}

func (c *TeamController) RemovePlayer(ctx http.Context) http.Response {
	teamID := ctx.Request().Route("id")
	playerID := ctx.Request().Route("playerId")

	facades.Orm().Query().Where("team_id = ? AND player_id = ?", teamID, playerID).Delete(&models.TeamPlayer{})

	return ctx.Response().Redirect(http.StatusSeeOther, "/teams/"+teamID)
}

func (c *TeamController) AvailablePlayers(ctx http.Context) http.Response {
	teamID := ctx.Request().Route("id")

	// Get player IDs already in this team
	var teamPlayers []models.TeamPlayer
	facades.Orm().Query().Where("team_id = ?", teamID).Find(&teamPlayers)

	var existingPlayerIDs []any
	for _, tp := range teamPlayers {
		existingPlayerIDs = append(existingPlayerIDs, tp.PlayerID)
	}

	// Get players not in the list
	var players []models.Player
	if len(existingPlayerIDs) > 0 {
		facades.Orm().Query().With("User").WhereNotIn("id", existingPlayerIDs).Find(&players)
	} else {
		facades.Orm().Query().With("User").Find(&players)
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"players": players})
}
