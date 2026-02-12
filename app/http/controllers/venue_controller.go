package controllers

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/http/requests/venue"
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

	for i := range venues {
		venues[i].SetPhotoURL()
	}

	return c.inertia.Render(ctx, "venues/Index", map[string]any{
		"venues": venues,
	})
}

func (c *VenueController) Create(ctx http.Context) http.Response {
	return c.inertia.Render(ctx, "venues/Create", map[string]any{})
}

func (c *VenueController) Store(ctx http.Context) http.Response {
	var request venue.StoreVenueRequest
	validationErrors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "venues/Create", map[string]any{
			"errors": map[string]string{"name": "Error de validación"},
		})
	}

	if validationErrors != nil {
		return c.inertia.Render(ctx, "venues/Create", map[string]any{
			"errors": inertia.ValidationErrors(validationErrors.All()),
		})
	}

	newVenue := models.Venue{
		Name:    request.Name,
		Address: request.Address,
	}

	// Create venue first to generate UUID via DispatchesEvents
	if err := facades.Orm().Query().Create(&newVenue); err != nil {
		return c.inertia.Render(ctx, "venues/Create", map[string]any{
			"errors": map[string]string{"name": "Error al crear la sede"},
		})
	}

	// Upload photo using the generated UUID
	if file, err := ctx.Request().File("photo"); err == nil && file != nil {
		photoPath, uploadErr := c.handlePhotoUpload(file, newVenue.UUID)
		if uploadErr == nil {
			newVenue.Photo = photoPath
			facades.Orm().Query().Save(&newVenue)
		}
	}

	return ctx.Response().Redirect(http.StatusSeeOther, "/venues")
}

func (c *VenueController) Edit(ctx http.Context) http.Response {
	venueUUID := ctx.Request().Route("uuid")
	var venueModel models.Venue
	if err := facades.Orm().Query().Where("uuid", venueUUID).First(&venueModel); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/venues")
	}

	venueModel.SetPhotoURL()

	return c.inertia.Render(ctx, "venues/Edit", map[string]any{
		"venue": venueModel,
	})
}

func (c *VenueController) Update(ctx http.Context) http.Response {
	venueUUID := ctx.Request().Route("uuid")
	var venueModel models.Venue
	if err := facades.Orm().Query().Where("uuid", venueUUID).First(&venueModel); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/venues")
	}

	var request venue.UpdateVenueRequest
	validationErrors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "venues/Edit", map[string]any{
			"venue":  venueModel,
			"errors": map[string]string{"name": "Error de validación"},
		})
	}

	if validationErrors != nil {
		return c.inertia.Render(ctx, "venues/Edit", map[string]any{
			"venue":  venueModel,
			"errors": inertia.ValidationErrors(validationErrors.All()),
		})
	}

	venueModel.Name = request.Name
	venueModel.Address = request.Address

	if file, err := ctx.Request().File("photo"); err == nil && file != nil {
		if venueModel.Photo != "" {
			facades.Storage().Disk("minio").Delete(venueModel.Photo)
		}

		photoPath, uploadErr := c.handlePhotoUpload(file, venueModel.UUID)
		if uploadErr == nil {
			venueModel.Photo = photoPath
		}
	}

	if request.RemovePhoto == "true" && venueModel.Photo != "" {
		facades.Storage().Disk("minio").Delete(venueModel.Photo)
		venueModel.Photo = ""
	}

	facades.Orm().Query().Save(&venueModel)
	return ctx.Response().Redirect(http.StatusSeeOther, "/venues")
}

func (c *VenueController) Destroy(ctx http.Context) http.Response {
	venueUUID := ctx.Request().Route("uuid")
	var venueModel models.Venue
	if err := facades.Orm().Query().Where("uuid", venueUUID).First(&venueModel); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/venues")
	}

	if venueModel.Photo != "" {
		facades.Storage().Disk("minio").Delete(venueModel.Photo)
	}

	facades.Orm().Query().Delete(&venueModel)
	return ctx.Response().Redirect(http.StatusSeeOther, "/venues")
}

func (c *VenueController) DestroyPhoto(ctx http.Context) http.Response {
	venueUUID := ctx.Request().Route("uuid")
	var venueModel models.Venue
	if err := facades.Orm().Query().Where("uuid", venueUUID).First(&venueModel); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/venues")
	}

	if venueModel.Photo != "" {
		facades.Storage().Disk("minio").Delete(venueModel.Photo)
		venueModel.Photo = ""
		facades.Orm().Query().Save(&venueModel)
	}

	return ctx.Response().Redirect(http.StatusSeeOther, fmt.Sprintf("/venues/%s/edit", venueUUID))
}

func (c *VenueController) Photo(ctx http.Context) http.Response {
	venueUUID := ctx.Request().Route("uuid")
	var venueModel models.Venue
	if err := facades.Orm().Query().Where("uuid", venueUUID).First(&venueModel); err != nil || venueModel.Photo == "" {
		return ctx.Response().Status(http.StatusNotFound).String("Not found")
	}

	data, err := facades.Storage().Disk("minio").GetBytes(venueModel.Photo)
	if err != nil {
		return ctx.Response().Status(http.StatusNotFound).String("Not found")
	}

	contentType := "application/octet-stream"
	ext := strings.ToLower(filepath.Ext(venueModel.Photo))
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".webp":
		contentType = "image/webp"
	case ".gif":
		contentType = "image/gif"
	case ".bmp":
		contentType = "image/bmp"
	case ".tiff", ".tif":
		contentType = "image/tiff"
	case ".avif":
		contentType = "image/avif"
	case ".heif", ".heic":
		contentType = "image/heif"
	}

	return ctx.Response().Header("Content-Type", contentType).Data(http.StatusOK, contentType, data)
}

func (c *VenueController) handlePhotoUpload(file filesystem.File, venueUUID string) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.GetClientOriginalName()))
	filename := fmt.Sprintf("venue_profile_%d%s", time.Now().UnixNano(), ext)

	storedPath, err := file.Disk("minio").StoreAs("venues/"+venueUUID, filename)
	if err != nil {
		return "", err
	}

	return storedPath, nil
}
