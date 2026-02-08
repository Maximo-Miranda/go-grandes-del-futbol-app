package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"

	"grandesdelfutbol/app/http/requests/player"
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
	var request player.StorePlayerRequest
	errors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "players/Create", map[string]any{
			"errors": map[string]string{"name": "Error de validación"},
		})
	}

	if errors != nil {
		return c.inertia.Render(ctx, "players/Create", map[string]any{
			"errors": inertia.ValidationErrors(errors.All()),
		})
	}

	newPlayer := models.Player{
		Name:       request.Name,
		Nickname:   request.Nickname,
		DocumentID: request.DocumentID,
		Phone:      request.Phone,
		Position:   request.Position,
	}

	// Handle photo upload
	if file, err := ctx.Request().File("photo"); err == nil && file != nil {
		photoPath, uploadErr := c.handlePhotoUpload(file)
		if uploadErr == nil {
			newPlayer.Photo = photoPath
		}
	}

	if err := facades.Orm().Query().Create(&newPlayer); err != nil {
		return c.inertia.Render(ctx, "players/Create", map[string]any{
			"errors": map[string]string{"name": "Error al crear el jugador"},
		})
	}

	return ctx.Response().Redirect(http.StatusFound, "/players")
}

func (c *PlayerController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var playerModel models.Player
	if err := facades.Orm().Query().Find(&playerModel, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/players")
	}

	// Get player stats from match_events
	goals, _ := facades.Orm().Query().Model(&models.MatchEvent{}).Where("player_id = ? AND event_type = ?", id, "goal").Count()
	assists, _ := facades.Orm().Query().Model(&models.MatchEvent{}).Where("player_id = ? AND event_type = ?", id, "assist").Count()
	yellowCards, _ := facades.Orm().Query().Model(&models.MatchEvent{}).Where("player_id = ? AND event_type = ?", id, "yellow_card").Count()
	redCards, _ := facades.Orm().Query().Model(&models.MatchEvent{}).Where("player_id = ? AND event_type = ?", id, "red_card").Count()

	// Get matches played via lineups
	matchesPlayed, _ := facades.Orm().Query().Model(&models.MatchLineup{}).Where("player_id = ?", id).Count()

	// Get recent matches with player events
	var recentMatches []map[string]any
	var lineups []models.MatchLineup
	facades.Orm().Query().
		With("Match.HomeTeam").With("Match.AwayTeam").With("Team").
		Where("player_id = ?", id).
		Order("created_at DESC").
		Limit(5).
		Find(&lineups)

	for _, lineup := range lineups {
		if lineup.Match == nil {
			continue
		}

		// Get events for this player in this match
		matchGoals, _ := facades.Orm().Query().Model(&models.MatchEvent{}).
			Where("match_id = ? AND player_id = ? AND event_type = ?", lineup.MatchID, id, "goal").
			Count()
		matchAssists, _ := facades.Orm().Query().Model(&models.MatchEvent{}).
			Where("match_id = ? AND player_id = ? AND event_type = ?", lineup.MatchID, id, "assist").
			Count()
		yellowEvent, _ := facades.Orm().Query().Model(&models.MatchEvent{}).
			Where("match_id = ? AND player_id = ? AND event_type = ?", lineup.MatchID, id, "yellow_card").
			Count()
		redEvent, _ := facades.Orm().Query().Model(&models.MatchEvent{}).
			Where("match_id = ? AND player_id = ? AND event_type = ?", lineup.MatchID, id, "red_card").
			Count()

		var matchCard string
		if redEvent > 0 {
			matchCard = "red"
		} else if yellowEvent > 0 {
			matchCard = "yellow"
		}

		// Determine opponent and result
		var opponent string
		var result string
		isHome := lineup.TeamID == lineup.Match.HomeTeamID
		if isHome {
			if lineup.Match.AwayTeam != nil {
				opponent = lineup.Match.AwayTeam.Name
			}
			if lineup.Match.Status == "completed" {
				if lineup.Match.HomeScore > lineup.Match.AwayScore {
					result = "W"
				} else if lineup.Match.HomeScore < lineup.Match.AwayScore {
					result = "L"
				} else {
					result = "D"
				}
			}
		} else {
			if lineup.Match.HomeTeam != nil {
				opponent = lineup.Match.HomeTeam.Name
			}
			if lineup.Match.Status == "completed" {
				if lineup.Match.AwayScore > lineup.Match.HomeScore {
					result = "W"
				} else if lineup.Match.AwayScore < lineup.Match.HomeScore {
					result = "L"
				} else {
					result = "D"
				}
			}
		}

		dateStr := ""
		if lineup.Match.MatchDate != nil {
			dateStr = lineup.Match.MatchDate.Format("02/01/2006")
		}

		recentMatches = append(recentMatches, map[string]any{
			"date":     dateStr,
			"opponent": opponent,
			"result":   result,
			"goals":    matchGoals,
			"assists":  matchAssists,
			"cards":    matchCard,
			"isMvp":    false,
		})
	}

	return c.inertia.Render(ctx, "players/Show", map[string]any{
		"player": playerModel,
		"stats": map[string]any{
			"goals":         goals,
			"assists":       assists,
			"yellow_cards":  yellowCards,
			"red_cards":     redCards,
			"matchesPlayed": matchesPlayed,
		},
		"recentMatches": recentMatches,
	})
}

func (c *PlayerController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var playerModel models.Player
	if err := facades.Orm().Query().Find(&playerModel, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/players")
	}

	return c.inertia.Render(ctx, "players/Edit", map[string]any{
		"player": playerModel,
	})
}

func (c *PlayerController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var playerModel models.Player
	if err := facades.Orm().Query().Find(&playerModel, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/players")
	}

	var request player.UpdatePlayerRequest
	errors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "players/Edit", map[string]any{
			"player": playerModel,
			"errors": map[string]string{"name": "Error de validación"},
		})
	}

	if errors != nil {
		return c.inertia.Render(ctx, "players/Edit", map[string]any{
			"player": playerModel,
			"errors": inertia.ValidationErrors(errors.All()),
		})
	}

	playerModel.Name = request.Name
	playerModel.Nickname = request.Nickname
	playerModel.DocumentID = request.DocumentID
	playerModel.Phone = request.Phone
	playerModel.Position = request.Position

	// Handle photo upload
	if file, err := ctx.Request().File("photo"); err == nil && file != nil {
		// Delete old photo if exists
		if playerModel.Photo != "" {
			oldPath := path.Storage("app/public/" + playerModel.Photo)
			os.Remove(oldPath)
		}

		photoPath, uploadErr := c.handlePhotoUpload(file)
		if uploadErr == nil {
			playerModel.Photo = photoPath
		}
	}

	// Handle photo removal
	if request.RemovePhoto == "true" && playerModel.Photo != "" {
		oldPath := path.Storage("app/public/" + playerModel.Photo)
		os.Remove(oldPath)
		playerModel.Photo = ""
	}

	facades.Orm().Query().Save(&playerModel)
	return ctx.Response().Redirect(http.StatusFound, "/players/"+id)
}

func (c *PlayerController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var playerModel models.Player
	if err := facades.Orm().Query().Find(&playerModel, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/players")
	}

	// Delete photo file if exists
	if playerModel.Photo != "" {
		photoPath := path.Storage("app/public/" + playerModel.Photo)
		os.Remove(photoPath)
	}

	facades.Orm().Query().Delete(&playerModel)
	return ctx.Response().Redirect(http.StatusFound, "/players")
}

// Photo serves player photos only to authenticated users
func (c *PlayerController) Photo(ctx http.Context) http.Response {
	filename := ctx.Request().Route("filename")
	photoPath := path.Storage("app/public/players/" + filename)

	// Check if file exists
	if _, err := os.Stat(photoPath); os.IsNotExist(err) {
		return ctx.Response().Status(http.StatusNotFound).String("Not found")
	}

	// Serve file
	return ctx.Response().File(photoPath)
}

func (c *PlayerController) handlePhotoUpload(file filesystem.File) (string, error) {
	// Validate file type
	ext := strings.ToLower(filepath.Ext(file.GetClientOriginalName()))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		return "", fmt.Errorf("invalid file type")
	}

	// Generate unique filename
	filename := fmt.Sprintf("player_%d%s", time.Now().UnixNano(), ext)

	// Ensure directory exists (absolute path for mkdir)
	uploadDir := path.Storage("app/public/players")
	os.MkdirAll(uploadDir, 0755)

	// Save file using relative path to disk root (storage/app/)
	// StoreAs expects path relative to the disk's root, not absolute
	if _, err := file.StoreAs("public/players", filename); err != nil {
		return "", err
	}

	return "players/" + filename, nil
}
