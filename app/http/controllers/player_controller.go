package controllers

import (
	"path/filepath"
	"strings"

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
	facades.Orm().Query().With("User").Order("nickname ASC").Find(&players)

	for i := range players {
		players[i].SetPhotoURL()
	}

	return c.inertia.Render(ctx, "players/Index", map[string]any{
		"players": players,
	})
}

func (c *PlayerController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var playerModel models.Player
	if err := facades.Orm().Query().With("User").Find(&playerModel, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/players")
	}

	playerModel.SetPhotoURL()

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
		if !lineup.Match.MatchDate.IsZero() {
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

// Photo serves player photos only to authenticated users
func (c *PlayerController) Photo(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var playerModel models.Player
	if err := facades.Orm().Query().Find(&playerModel, id); err != nil || playerModel.Photo == "" {
		return ctx.Response().Status(http.StatusNotFound).String("Not found")
	}

	data, err := facades.Storage().Disk("minio").GetBytes(playerModel.Photo)
	if err != nil {
		return ctx.Response().Status(http.StatusNotFound).String("Not found")
	}

	contentType := "application/octet-stream"
	ext := strings.ToLower(filepath.Ext(playerModel.Photo))
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
