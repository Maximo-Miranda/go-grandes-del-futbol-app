package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/inertia"
	"grandesdelfutbol/app/models"
)

type VenueController struct {
	inertia *inertia.Inertia
}

func NewVenueController() *VenueController {
	return &VenueController{inertia: inertia.New()}
}

func (c *VenueController) Index(ctx http.Context) http.Response {
	var venues []models.Venue
	facades.Orm().Query().Order("name ASC").Find(&venues)

	return c.inertia.Render(ctx, "venues/Index", map[string]any{
		"venues": venues,
	})
}

func (c *VenueController) Create(ctx http.Context) http.Response {
	return c.inertia.Render(ctx, "venues/Create", map[string]any{})
}

func (c *VenueController) Store(ctx http.Context) http.Response {
	venue := models.Venue{
		Name:    ctx.Request().Input("name"),
		Address: ctx.Request().Input("address"),
	}

	if err := facades.Orm().Query().Create(&venue); err != nil {
		return c.inertia.Render(ctx, "venues/Create", map[string]any{
			"errors": map[string]string{"name": "Error al crear la sede"},
		})
	}

	return ctx.Response().Redirect(http.StatusFound, "/venues")
}

func (c *VenueController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var venue models.Venue
	if err := facades.Orm().Query().Find(&venue, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/venues")
	}

	return c.inertia.Render(ctx, "venues/Edit", map[string]any{
		"venue": venue,
	})
}

func (c *VenueController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var venue models.Venue
	if err := facades.Orm().Query().Find(&venue, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/venues")
	}

	venue.Name = ctx.Request().Input("name", venue.Name)
	venue.Address = ctx.Request().Input("address", venue.Address)

	facades.Orm().Query().Save(&venue)
	return ctx.Response().Redirect(http.StatusFound, "/venues")
}

func (c *VenueController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var venue models.Venue
	if err := facades.Orm().Query().Find(&venue, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/venues")
	}
	facades.Orm().Query().Delete(&venue)
	return ctx.Response().Redirect(http.StatusFound, "/venues")
}
