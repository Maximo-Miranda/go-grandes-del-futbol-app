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

	return ctx.Response().Redirect(http.StatusFound, "/teams")
}

func (c *TeamController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/teams")
	}

	var teamPlayers []models.TeamPlayer
	facades.Orm().Query().With("Player").Where("team_id = ?", id).Find(&teamPlayers)

	return c.inertia.Render(ctx, "teams/Show", map[string]any{
		"team":    team,
		"players": teamPlayers,
	})
}

func (c *TeamController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/teams")
	}

	return c.inertia.Render(ctx, "teams/Edit", map[string]any{
		"team": team,
	})
}

func (c *TeamController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/teams")
	}

	team.Name = ctx.Request().Input("name", team.Name)
	team.Color = ctx.Request().Input("color", team.Color)
	team.ContactPhone = ctx.Request().Input("contact_phone", team.ContactPhone)

	facades.Orm().Query().Save(&team)
	return ctx.Response().Redirect(http.StatusFound, "/teams/"+id)
}

func (c *TeamController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/teams")
	}
	facades.Orm().Query().Delete(&team)
	return ctx.Response().Redirect(http.StatusFound, "/teams")
}
