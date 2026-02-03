package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/inertia"
	"grandesdelfutbol/app/models"
)

type PlayerController struct {
	inertia *inertia.Inertia
}

func NewPlayerController() *PlayerController {
	return &PlayerController{inertia: inertia.New()}
}

func (c *PlayerController) Index(ctx http.Context) http.Response {
	var players []models.Player
	facades.Orm().Query().Order("name ASC").Find(&players)

	return c.inertia.Render(ctx, "players/Index", map[string]any{
		"players": players,
	})
}

func (c *PlayerController) Create(ctx http.Context) http.Response {
	return c.inertia.Render(ctx, "players/Create", map[string]any{})
}

func (c *PlayerController) Store(ctx http.Context) http.Response {
	player := models.Player{
		Name:       ctx.Request().Input("name"),
		Nickname:   ctx.Request().Input("nickname"),
		DocumentID: ctx.Request().Input("document_id"),
		Phone:      ctx.Request().Input("phone"),
		Position:   ctx.Request().Input("position"),
	}

	if err := facades.Orm().Query().Create(&player); err != nil {
		return c.inertia.Render(ctx, "players/Create", map[string]any{
			"errors": map[string]string{"name": "Error al crear el jugador"},
		})
	}

	return ctx.Response().Redirect(http.StatusFound, "/players")
}

func (c *PlayerController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var player models.Player
	if err := facades.Orm().Query().Find(&player, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/players")
	}

	// Get player stats: goals, yellow cards, red cards
	goals, _ := facades.Orm().Query().Model(&models.MatchEvent{}).Where("player_id = ? AND event_type = ?", id, "goal").Count()
	yellowCards, _ := facades.Orm().Query().Model(&models.MatchEvent{}).Where("player_id = ? AND event_type = ?", id, "yellow_card").Count()
	redCards, _ := facades.Orm().Query().Model(&models.MatchEvent{}).Where("player_id = ? AND event_type = ?", id, "red_card").Count()

	return c.inertia.Render(ctx, "players/Show", map[string]any{
		"player": player,
		"stats": map[string]any{
			"goals":        goals,
			"yellow_cards": yellowCards,
			"red_cards":    redCards,
		},
	})
}

func (c *PlayerController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var player models.Player
	if err := facades.Orm().Query().Find(&player, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/players")
	}

	return c.inertia.Render(ctx, "players/Edit", map[string]any{
		"player": player,
	})
}

func (c *PlayerController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var player models.Player
	if err := facades.Orm().Query().Find(&player, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/players")
	}

	player.Name = ctx.Request().Input("name", player.Name)
	player.Nickname = ctx.Request().Input("nickname", player.Nickname)
	player.DocumentID = ctx.Request().Input("document_id", player.DocumentID)
	player.Phone = ctx.Request().Input("phone", player.Phone)
	player.Position = ctx.Request().Input("position", player.Position)

	facades.Orm().Query().Save(&player)
	return ctx.Response().Redirect(http.StatusFound, "/players/"+id)
}

func (c *PlayerController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var player models.Player
	if err := facades.Orm().Query().Find(&player, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/players")
	}
	facades.Orm().Query().Delete(&player)
	return ctx.Response().Redirect(http.StatusFound, "/players")
}
