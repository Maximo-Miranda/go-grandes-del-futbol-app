package controllers

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	profileRequest "grandesdelfutbol/app/http/requests/profile"
	"grandesdelfutbol/app/inertia"
	"grandesdelfutbol/app/models"
)

type ProfileController struct {
	inertia *inertia.Inertia
}

func NewProfileController() *ProfileController {
	return &ProfileController{inertia: inertia.New()}
}

func (c *ProfileController) Create(ctx http.Context) http.Response {
	user, _ := ctx.Value("user").(*models.User)

	// If user already has a profile, redirect to dashboard
	if user != nil {
		var existing models.Player
		if err := facades.Orm().Query().Where("user_id = ?", user.ID).First(&existing); err == nil && existing.ID != 0 {
			return ctx.Response().Redirect(http.StatusSeeOther, "/dashboard")
		}
	}

	return c.inertia.Render(ctx, "profile/Create", map[string]any{})
}

func (c *ProfileController) Store(ctx http.Context) http.Response {
	user, _ := ctx.Value("user").(*models.User)
	if user == nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/auth/login")
	}

	// Check if user already has a profile
	var existing models.Player
	if err := facades.Orm().Query().Where("user_id = ?", user.ID).First(&existing); err == nil && existing.ID != 0 {
		return ctx.Response().Redirect(http.StatusSeeOther, "/dashboard")
	}

	var request profileRequest.StoreProfileRequest
	errors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "profile/Create", map[string]any{
			"errors": map[string]string{"nickname": "Error de validación"},
		})
	}

	if errors != nil {
		return c.inertia.Render(ctx, "profile/Create", map[string]any{
			"errors": inertia.ValidationErrors(errors.All()),
		})
	}

	newPlayer := models.Player{
		UserID:     user.ID,
		Nickname:   request.Nickname,
		DocumentID: request.DocumentID,
		Phone:      request.Phone,
		Position:   request.Position,
	}

	if err := facades.Orm().Query().Create(&newPlayer); err != nil {
		return c.inertia.Render(ctx, "profile/Create", map[string]any{
			"errors": map[string]string{"nickname": "Error al crear el perfil deportivo"},
		})
	}

	// Upload photo using the generated ID
	if file, err := ctx.Request().File("photo"); err == nil && file != nil {
		photoPath, uploadErr := c.handlePhotoUpload(file, newPlayer.ID)
		if uploadErr == nil {
			newPlayer.Photo = photoPath
			facades.Orm().Query().Save(&newPlayer)
		}
	}

	inertia.Flash(ctx, "success", "Perfil deportivo creado exitosamente")
	return ctx.Response().Redirect(http.StatusSeeOther, "/dashboard")
}

func (c *ProfileController) Edit(ctx http.Context) http.Response {
	user, _ := ctx.Value("user").(*models.User)
	if user == nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/auth/login")
	}

	var player models.Player
	if err := facades.Orm().Query().Where("user_id = ?", user.ID).First(&player); err != nil || player.ID == 0 {
		return ctx.Response().Redirect(http.StatusSeeOther, "/profile/create")
	}

	player.SetPhotoURL()

	return c.inertia.Render(ctx, "profile/Edit", map[string]any{
		"player": player,
	})
}

func (c *ProfileController) Update(ctx http.Context) http.Response {
	user, _ := ctx.Value("user").(*models.User)
	if user == nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/auth/login")
	}

	var player models.Player
	if err := facades.Orm().Query().Where("user_id = ?", user.ID).First(&player); err != nil || player.ID == 0 {
		return ctx.Response().Redirect(http.StatusSeeOther, "/profile/create")
	}

	var request profileRequest.StoreProfileRequest
	errors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "profile/Edit", map[string]any{
			"player": player,
			"errors": map[string]string{"nickname": "Error de validación"},
		})
	}

	if errors != nil {
		return c.inertia.Render(ctx, "profile/Edit", map[string]any{
			"player": player,
			"errors": inertia.ValidationErrors(errors.All()),
		})
	}

	player.Nickname = request.Nickname
	player.DocumentID = request.DocumentID
	player.Phone = request.Phone
	player.Position = request.Position

	// Handle photo upload
	if file, err := ctx.Request().File("photo"); err == nil && file != nil {
		if player.Photo != "" {
			facades.Storage().Disk("minio").Delete(player.Photo)
		}
		photoPath, uploadErr := c.handlePhotoUpload(file, player.ID)
		if uploadErr == nil {
			player.Photo = photoPath
		}
	}

	// Handle photo removal
	if ctx.Request().Input("remove_photo") == "true" && player.Photo != "" {
		facades.Storage().Disk("minio").Delete(player.Photo)
		player.Photo = ""
	}

	facades.Orm().Query().Save(&player)
	inertia.Flash(ctx, "success", "Perfil actualizado exitosamente")
	return ctx.Response().Redirect(http.StatusSeeOther, fmt.Sprintf("/players/%d", player.ID))
}

func (c *ProfileController) handlePhotoUpload(file filesystem.File, playerID uint) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.GetClientOriginalName()))
	filename := fmt.Sprintf("player_profile_%d%s", time.Now().UnixNano(), ext)

	storedPath, err := file.Disk("minio").StoreAs(fmt.Sprintf("players/%d", playerID), filename)
	if err != nil {
		return "", err
	}

	return storedPath, nil
}
