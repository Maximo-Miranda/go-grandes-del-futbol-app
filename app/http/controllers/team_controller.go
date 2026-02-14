package controllers

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	teamRequest "grandesdelfutbol/app/http/requests/team"
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
	facades.Orm().Query().With("Players").Order("name ASC").Find(&teams)

	for i := range teams {
		teams[i].SetLogoURL()
	}

	return c.inertia.Render(ctx, "teams/Index", map[string]any{
		"teams": teams,
	})
}

func (c *TeamController) Create(ctx http.Context) http.Response {
	return c.inertia.Render(ctx, "teams/Create", map[string]any{})
}

func (c *TeamController) Store(ctx http.Context) http.Response {
	var request teamRequest.StoreTeamRequest
	validationErrors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "teams/Create", map[string]any{
			"errors": map[string]string{"name": "Error de validación"},
		})
	}

	if validationErrors != nil {
		return c.inertia.Render(ctx, "teams/Create", map[string]any{
			"errors": inertia.ValidationErrors(validationErrors.All()),
		})
	}

	team := models.Team{
		Name:         request.Name,
		Color:        request.Color,
		ContactPhone: request.ContactPhone,
	}

	if err := facades.Orm().Query().Create(&team); err != nil {
		return c.inertia.Render(ctx, "teams/Create", map[string]any{
			"errors": map[string]string{"name": "Error al crear el equipo"},
		})
	}

	// Upload logo using the generated ID
	if file, err := ctx.Request().File("logo"); err == nil && file != nil {
		logoPath, uploadErr := c.handleLogoUpload(file, team.ID)
		if uploadErr == nil {
			team.Logo = logoPath
			facades.Orm().Query().Save(&team)
		}
	}

	inertia.Flash(ctx, "success", "Equipo creado exitosamente")
	return ctx.Response().Redirect(http.StatusSeeOther, "/teams")
}

func (c *TeamController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/teams")
	}

	team.SetLogoURL()

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

	team.SetLogoURL()

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

	var request teamRequest.UpdateTeamRequest
	validationErrors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		team.SetLogoURL()
		return c.inertia.Render(ctx, "teams/Edit", map[string]any{
			"team":   team,
			"errors": map[string]string{"name": "Error de validación"},
		})
	}

	if validationErrors != nil {
		team.SetLogoURL()
		return c.inertia.Render(ctx, "teams/Edit", map[string]any{
			"team":   team,
			"errors": inertia.ValidationErrors(validationErrors.All()),
		})
	}

	team.Name = request.Name
	team.Color = request.Color
	team.ContactPhone = request.ContactPhone

	// Handle logo upload
	if file, err := ctx.Request().File("logo"); err == nil && file != nil {
		if team.Logo != "" {
			facades.Storage().Disk("minio").Delete(team.Logo)
		}
		logoPath, uploadErr := c.handleLogoUpload(file, team.ID)
		if uploadErr == nil {
			team.Logo = logoPath
		}
	}

	// Handle logo removal
	if request.RemoveLogo == "true" && team.Logo != "" {
		facades.Storage().Disk("minio").Delete(team.Logo)
		team.Logo = ""
	}

	facades.Orm().Query().Save(&team)
	inertia.Flash(ctx, "success", "Equipo actualizado exitosamente")
	return ctx.Response().Redirect(http.StatusSeeOther, fmt.Sprintf("/teams/%d", team.ID))
}

func (c *TeamController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/teams")
	}

	if team.Logo != "" {
		facades.Storage().Disk("minio").Delete(team.Logo)
	}

	facades.Orm().Query().Delete(&team)
	inertia.Flash(ctx, "success", "Equipo eliminado exitosamente")
	return ctx.Response().Redirect(http.StatusSeeOther, "/teams")
}

func (c *TeamController) Logo(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var team models.Team
	if err := facades.Orm().Query().Find(&team, id); err != nil || team.Logo == "" {
		return ctx.Response().Status(http.StatusNotFound).String("Not found")
	}

	data, err := facades.Storage().Disk("minio").GetBytes(team.Logo)
	if err != nil {
		return ctx.Response().Status(http.StatusNotFound).String("Not found")
	}

	contentType := "application/octet-stream"
	ext := strings.ToLower(filepath.Ext(team.Logo))
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".webp":
		contentType = "image/webp"
	}

	return ctx.Response().Header("Content-Type", contentType).Data(http.StatusOK, contentType, data)
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

	var teamPlayers []models.TeamPlayer
	facades.Orm().Query().Where("team_id = ?", teamID).Find(&teamPlayers)

	var existingPlayerIDs []any
	for _, tp := range teamPlayers {
		existingPlayerIDs = append(existingPlayerIDs, tp.PlayerID)
	}

	var players []models.Player
	if len(existingPlayerIDs) > 0 {
		facades.Orm().Query().With("User").WhereNotIn("id", existingPlayerIDs).Find(&players)
	} else {
		facades.Orm().Query().With("User").Find(&players)
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"players": players})
}

func (c *TeamController) handleLogoUpload(file filesystem.File, teamID uint) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.GetClientOriginalName()))
	filename := fmt.Sprintf("team_logo_%d%s", time.Now().UnixNano(), ext)

	storedPath, err := file.Disk("minio").StoreAs(fmt.Sprintf("teams/%d", teamID), filename)
	if err != nil {
		return "", err
	}

	return storedPath, nil
}
