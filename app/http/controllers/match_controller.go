package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/inertia"
	"grandesdelfutbol/app/models"
)

type MatchController struct {
	inertia *inertia.Inertia
}

func NewMatchController() *MatchController {
	return &MatchController{inertia: inertia.New()}
}

func (c *MatchController) Index(ctx http.Context) http.Response {
	tournamentID := ctx.Request().Query("tournament_id", "")

	query := facades.Orm().Query().With("HomeTeam").With("AwayTeam").With("Tournament").With("Venue")
	if tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}

	var matches []models.Match
	query.Order("match_date ASC").Find(&matches)

	var tournaments []models.Tournament
	facades.Orm().Query().Find(&tournaments)

	return c.inertia.Render(ctx, "matches/Index", map[string]any{
		"matches":      matches,
		"tournaments":  tournaments,
		"tournamentId": tournamentID,
	})
}

func (c *MatchController) Create(ctx http.Context) http.Response {
	tournamentID := ctx.Request().Query("tournament_id", "")

	var tournaments []models.Tournament
	facades.Orm().Query().Find(&tournaments)

	var teams []models.Team
	facades.Orm().Query().Find(&teams)

	var venues []models.Venue
	facades.Orm().Query().Find(&venues)

	return c.inertia.Render(ctx, "matches/Create", map[string]any{
		"tournaments":  tournaments,
		"teams":        teams,
		"venues":       venues,
		"tournamentId": tournamentID,
	})
}

func (c *MatchController) Store(ctx http.Context) http.Response {
	var tournaments []models.Tournament
	facades.Orm().Query().Find(&tournaments)
	var teams []models.Team
	facades.Orm().Query().Find(&teams)
	var venues []models.Venue
	facades.Orm().Query().Find(&venues)

	// Validation
	errors := make(map[string]string)
	tournamentID := ctx.Request().Input("tournament_id")
	homeTeamID := ctx.Request().Input("home_team_id")
	awayTeamID := ctx.Request().Input("away_team_id")

	if tournamentID == "" {
		errors["tournament_id"] = "El torneo es obligatorio"
	}
	if homeTeamID == "" {
		errors["home_team_id"] = "El equipo local es obligatorio"
	}
	if awayTeamID == "" {
		errors["away_team_id"] = "El equipo visitante es obligatorio"
	}
	if homeTeamID == awayTeamID && homeTeamID != "" {
		errors["away_team_id"] = "El equipo visitante debe ser diferente al local"
	}

	if len(errors) > 0 {
		return c.inertia.Render(ctx, "matches/Create", map[string]any{
			"tournaments": tournaments,
			"teams":       teams,
			"venues":      venues,
			"errors":      errors,
		})
	}

	match := models.Match{
		Status: "scheduled",
	}

	// Parse IDs
	var tID, hID, aID uint
	fmt.Sscanf(tournamentID, "%d", &tID)
	fmt.Sscanf(homeTeamID, "%d", &hID)
	fmt.Sscanf(awayTeamID, "%d", &aID)

	match.TournamentID = tID
	match.HomeTeamID = hID
	match.AwayTeamID = aID

	// Optional venue
	if venueID := ctx.Request().Input("venue_id"); venueID != "" {
		var vID uint
		fmt.Sscanf(venueID, "%d", &vID)
		match.VenueID = &vID
	}

	// Match date
	if matchDate := ctx.Request().Input("match_date"); matchDate != "" {
		if t, err := time.Parse("2006-01-02T15:04", matchDate); err == nil {
			match.MatchDate = &t
		} else if t, err := time.Parse("2006-01-02", matchDate); err == nil {
			match.MatchDate = &t
		}
	}

	// Round/Matchday
	if matchday := ctx.Request().Input("matchday"); matchday != "" {
		md, _ := strconv.Atoi(matchday)
		match.Round = md
	}

	if err := facades.Orm().Query().Create(&match); err != nil {
		return c.inertia.Render(ctx, "matches/Create", map[string]any{
			"tournaments": tournaments,
			"teams":       teams,
			"venues":      venues,
			"errors":      map[string]string{"tournament_id": "Error al crear el partido"},
		})
	}

	return ctx.Response().Redirect(http.StatusFound, "/matches")
}

func (c *MatchController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/matches")
	}

	var tournaments []models.Tournament
	facades.Orm().Query().Find(&tournaments)
	var teams []models.Team
	facades.Orm().Query().Find(&teams)
	var venues []models.Venue
	facades.Orm().Query().Find(&venues)

	return c.inertia.Render(ctx, "matches/Edit", map[string]any{
		"match":       match,
		"tournaments": tournaments,
		"teams":       teams,
		"venues":      venues,
	})
}

func (c *MatchController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/matches")
	}

	// Parse IDs
	if tournamentID := ctx.Request().Input("tournament_id"); tournamentID != "" {
		var tID uint
		fmt.Sscanf(tournamentID, "%d", &tID)
		match.TournamentID = tID
	}
	if homeTeamID := ctx.Request().Input("home_team_id"); homeTeamID != "" {
		var hID uint
		fmt.Sscanf(homeTeamID, "%d", &hID)
		match.HomeTeamID = hID
	}
	if awayTeamID := ctx.Request().Input("away_team_id"); awayTeamID != "" {
		var aID uint
		fmt.Sscanf(awayTeamID, "%d", &aID)
		match.AwayTeamID = aID
	}
	if venueID := ctx.Request().Input("venue_id"); venueID != "" {
		var vID uint
		fmt.Sscanf(venueID, "%d", &vID)
		match.VenueID = &vID
	}
	if matchDate := ctx.Request().Input("match_date"); matchDate != "" {
		if t, err := time.Parse("2006-01-02T15:04", matchDate); err == nil {
			match.MatchDate = &t
		}
	}
	if matchday := ctx.Request().Input("matchday"); matchday != "" {
		md, _ := strconv.Atoi(matchday)
		match.Round = md
	}

	facades.Orm().Query().Save(&match)
	return ctx.Response().Redirect(http.StatusFound, "/matches/"+id)
}

func (c *MatchController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/matches")
	}
	facades.Orm().Query().Delete(&match)
	return ctx.Response().Redirect(http.StatusFound, "/matches")
}

func (c *MatchController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().
		With("HomeTeam").With("AwayTeam").With("Tournament").With("Venue").
		Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/matches")
	}

	// Get events with player info
	var events []models.MatchEvent
	facades.Orm().Query().
		With("Player").With("Team").
		Where("match_id = ?", id).
		Order("minute ASC").
		Find(&events)

	// Get lineups with player info
	var lineups []models.MatchLineup
	facades.Orm().Query().
		With("Player").With("Team").
		Where("match_id = ?", id).
		Find(&lineups)

	// Get available players for both teams (for adding events)
	var homeTeamPlayers []models.Player
	var awayTeamPlayers []models.Player

	// Players via team_players table
	var homeTeamPlayerLinks []models.TeamPlayer
	var awayTeamPlayerLinks []models.TeamPlayer
	facades.Orm().Query().With("Player").Where("team_id = ?", match.HomeTeamID).Find(&homeTeamPlayerLinks)
	facades.Orm().Query().With("Player").Where("team_id = ?", match.AwayTeamID).Find(&awayTeamPlayerLinks)

	for _, tp := range homeTeamPlayerLinks {
		if tp.Player != nil {
			homeTeamPlayers = append(homeTeamPlayers, *tp.Player)
		}
	}
	for _, tp := range awayTeamPlayerLinks {
		if tp.Player != nil {
			awayTeamPlayers = append(awayTeamPlayers, *tp.Player)
		}
	}

	return c.inertia.Render(ctx, "matches/Show", map[string]any{
		"match":           match,
		"events":          events,
		"lineups":         lineups,
		"homeTeamPlayers": homeTeamPlayers,
		"awayTeamPlayers": awayTeamPlayers,
	})
}

func (c *MatchController) RecordResult(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/matches")
	}

	homeScore, _ := strconv.Atoi(ctx.Request().Input("home_score", "0"))
	awayScore, _ := strconv.Atoi(ctx.Request().Input("away_score", "0"))

	match.HomeScore = homeScore
	match.AwayScore = awayScore
	match.Status = "completed"

	facades.Orm().Query().Save(&match)

	// Update standings
	c.updateStandings(&match)

	return ctx.Response().Redirect(http.StatusFound, "/matches/"+id)
}

// CloseMatch marks a match as finished without changing the score
func (c *MatchController) CloseMatch(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/matches")
	}

	match.Status = "completed"
	facades.Orm().Query().Save(&match)

	// Update standings if not already done
	c.updateStandings(&match)

	return ctx.Response().Redirect(http.StatusFound, "/matches/"+id)
}

// AddEvent adds a match event (goal, assist, card, etc.)
func (c *MatchController) AddEvent(ctx http.Context) http.Response {
	matchID := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, matchID); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/matches")
	}

	playerID, _ := strconv.ParseUint(ctx.Request().Input("player_id"), 10, 64)
	teamID, _ := strconv.ParseUint(ctx.Request().Input("team_id"), 10, 64)
	minute, _ := strconv.Atoi(ctx.Request().Input("minute", "0"))
	eventType := ctx.Request().Input("event_type")

	if playerID == 0 || teamID == 0 || eventType == "" {
		return ctx.Response().Redirect(http.StatusFound, "/matches/"+matchID)
	}

	event := models.MatchEvent{
		MatchID:   match.ID,
		PlayerID:  uint(playerID),
		TeamID:    uint(teamID),
		EventType: eventType,
		Minute:    minute,
	}

	facades.Orm().Query().Create(&event)

	// If it's a goal, update the match score
	if eventType == "goal" {
		if uint(teamID) == match.HomeTeamID {
			match.HomeScore++
		} else if uint(teamID) == match.AwayTeamID {
			match.AwayScore++
		}
		facades.Orm().Query().Save(&match)
	}

	return ctx.Response().Redirect(http.StatusFound, "/matches/"+matchID)
}

// DeleteEvent removes a match event
func (c *MatchController) DeleteEvent(ctx http.Context) http.Response {
	matchID := ctx.Request().Route("id")
	eventID := ctx.Request().Route("eventId")

	var match models.Match
	if err := facades.Orm().Query().Find(&match, matchID); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/matches")
	}

	var event models.MatchEvent
	if err := facades.Orm().Query().Where("id = ? AND match_id = ?", eventID, matchID).Find(&event); err != nil || event.ID == 0 {
		return ctx.Response().Redirect(http.StatusFound, "/matches/"+matchID)
	}

	// If it's a goal, decrement the score
	if event.EventType == "goal" {
		if event.TeamID == match.HomeTeamID {
			if match.HomeScore > 0 {
				match.HomeScore--
			}
		} else if event.TeamID == match.AwayTeamID {
			if match.AwayScore > 0 {
				match.AwayScore--
			}
		}
		facades.Orm().Query().Save(&match)
	}

	facades.Orm().Query().Delete(&event)

	return ctx.Response().Redirect(http.StatusFound, "/matches/"+matchID)
}

// AddLineup adds a player to the match lineup
func (c *MatchController) AddLineup(ctx http.Context) http.Response {
	matchID := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, matchID); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/matches")
	}

	playerID, _ := strconv.ParseUint(ctx.Request().Input("player_id"), 10, 64)
	teamID, _ := strconv.ParseUint(ctx.Request().Input("team_id"), 10, 64)
	minutesPlayed, _ := strconv.Atoi(ctx.Request().Input("minutes_played", "90"))
	isStarter := ctx.Request().Input("is_starter") == "true"

	if playerID == 0 || teamID == 0 {
		return ctx.Response().Redirect(http.StatusFound, "/matches/"+matchID)
	}

	// Check if player is already in lineup
	var existing models.MatchLineup
	facades.Orm().Query().Where("match_id = ? AND player_id = ?", matchID, playerID).First(&existing)
	if existing.ID > 0 {
		// Update existing
		existing.MinutesPlayed = minutesPlayed
		existing.IsStarter = isStarter
		facades.Orm().Query().Save(&existing)
	} else {
		// Create new
		lineup := models.MatchLineup{
			MatchID:       match.ID,
			PlayerID:      uint(playerID),
			TeamID:        uint(teamID),
			IsStarter:     isStarter,
			MinutesPlayed: minutesPlayed,
		}
		facades.Orm().Query().Create(&lineup)
	}

	return ctx.Response().Redirect(http.StatusFound, "/matches/"+matchID)
}

// RemoveLineup removes a player from the match lineup
func (c *MatchController) RemoveLineup(ctx http.Context) http.Response {
	matchID := ctx.Request().Route("id")
	playerID := ctx.Request().Route("playerId")

	facades.Orm().Query().Where("match_id = ? AND player_id = ?", matchID, playerID).Delete(&models.MatchLineup{})

	return ctx.Response().Redirect(http.StatusFound, "/matches/"+matchID)
}

func (c *MatchController) updateStandings(match *models.Match) {
	updateTeamStanding := func(teamID uint, goalsFor, goalsAgainst int, won, drawn, lost bool) {
		var standing models.Standing
		err := facades.Orm().Query().
			Where("tournament_id = ? AND team_id = ?", match.TournamentID, teamID).
			First(&standing)

		if err != nil {
			standing = models.Standing{
				TournamentID: match.TournamentID,
				GroupID:      match.GroupID,
				TeamID:       teamID,
			}
		}

		standing.Played++
		standing.GoalsFor += goalsFor
		standing.GoalsAgainst += goalsAgainst
		standing.GoalDifference = standing.GoalsFor - standing.GoalsAgainst

		if won {
			standing.Won++
			standing.Points += 3
		} else if drawn {
			standing.Drawn++
			standing.Points += 1
		} else if lost {
			standing.Lost++
		}

		if standing.ID == 0 {
			facades.Orm().Query().Create(&standing)
		} else {
			facades.Orm().Query().Save(&standing)
		}
	}

	homeWon := match.HomeScore > match.AwayScore
	awayWon := match.AwayScore > match.HomeScore
	isDraw := match.HomeScore == match.AwayScore

	updateTeamStanding(match.HomeTeamID, match.HomeScore, match.AwayScore, homeWon, isDraw, awayWon)
	updateTeamStanding(match.AwayTeamID, match.AwayScore, match.HomeScore, awayWon, isDraw, homeWon)
}
