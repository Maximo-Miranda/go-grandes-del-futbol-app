package controllers

import (
	"strconv"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	matchRequest "grandesdelfutbol/app/http/requests/match"
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

	renderData := map[string]any{
		"tournaments": tournaments,
		"teams":       teams,
		"venues":      venues,
	}

	var request matchRequest.StoreMatchRequest
	validationErrors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		renderData["errors"] = map[string]string{"tournament_id": "Error de validación"}
		return c.inertia.Render(ctx, "matches/Create", renderData)
	}

	if validationErrors != nil {
		renderData["errors"] = inertia.ValidationErrors(validationErrors.All())
		return c.inertia.Render(ctx, "matches/Create", renderData)
	}

	// Custom validation: teams must be different
	if request.HomeTeamID == request.AwayTeamID {
		renderData["errors"] = map[string]string{"away_team_id": "El equipo visitante debe ser diferente al local"}
		return c.inertia.Render(ctx, "matches/Create", renderData)
	}

	// Parse IDs safely
	tournamentID, err2 := strconv.ParseUint(request.TournamentID, 10, 64)
	homeTeamID, err3 := strconv.ParseUint(request.HomeTeamID, 10, 64)
	awayTeamID, err4 := strconv.ParseUint(request.AwayTeamID, 10, 64)
	venueID, err5 := strconv.ParseUint(request.VenueID, 10, 64)
	matchday, err6 := strconv.Atoi(request.Matchday)

	if err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		renderData["errors"] = map[string]string{"tournament_id": "Datos inválidos, verifica los campos"}
		return c.inertia.Render(ctx, "matches/Create", renderData)
	}

	matchDate, errDate := time.Parse("2006-01-02T15:04", request.MatchDate)
	if errDate != nil {
		matchDate, errDate = time.Parse("2006-01-02", request.MatchDate)
	}
	if errDate != nil {
		renderData["errors"] = map[string]string{"match_date": "Formato de fecha inválido"}
		return c.inertia.Render(ctx, "matches/Create", renderData)
	}

	match := models.Match{
		TournamentID: uint(tournamentID),
		HomeTeamID:   uint(homeTeamID),
		AwayTeamID:   uint(awayTeamID),
		VenueID:      uint(venueID),
		MatchDate:    matchDate,
		Round:        matchday,
		Status:       "scheduled",
	}

	if err := facades.Orm().Query().Create(&match); err != nil {
		renderData["errors"] = map[string]string{"tournament_id": "Error al crear el partido"}
		return c.inertia.Render(ctx, "matches/Create", renderData)
	}

	inertia.Flash(ctx, "success", "Partido creado exitosamente")
	return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
}

func (c *MatchController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
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
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
	}

	var tournaments []models.Tournament
	facades.Orm().Query().Find(&tournaments)
	var teams []models.Team
	facades.Orm().Query().Find(&teams)
	var venues []models.Venue
	facades.Orm().Query().Find(&venues)

	renderData := map[string]any{
		"match":       match,
		"tournaments": tournaments,
		"teams":       teams,
		"venues":      venues,
	}

	var request matchRequest.UpdateMatchRequest
	validationErrors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		renderData["errors"] = map[string]string{"tournament_id": "Error de validación"}
		return c.inertia.Render(ctx, "matches/Edit", renderData)
	}

	if validationErrors != nil {
		renderData["errors"] = inertia.ValidationErrors(validationErrors.All())
		return c.inertia.Render(ctx, "matches/Edit", renderData)
	}

	// Custom validation: teams must be different
	if request.HomeTeamID == request.AwayTeamID {
		renderData["errors"] = map[string]string{"away_team_id": "El equipo visitante debe ser diferente al local"}
		return c.inertia.Render(ctx, "matches/Edit", renderData)
	}

	// Parse IDs safely
	tournamentID, err2 := strconv.ParseUint(request.TournamentID, 10, 64)
	homeTeamID, err3 := strconv.ParseUint(request.HomeTeamID, 10, 64)
	awayTeamID, err4 := strconv.ParseUint(request.AwayTeamID, 10, 64)
	venueID, err5 := strconv.ParseUint(request.VenueID, 10, 64)
	matchday, err6 := strconv.Atoi(request.Matchday)

	if err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		renderData["errors"] = map[string]string{"tournament_id": "Datos inválidos, verifica los campos"}
		return c.inertia.Render(ctx, "matches/Edit", renderData)
	}

	matchDate, errDate := time.Parse("2006-01-02T15:04", request.MatchDate)
	if errDate != nil {
		matchDate, errDate = time.Parse("2006-01-02", request.MatchDate)
	}
	if errDate != nil {
		renderData["errors"] = map[string]string{"match_date": "Formato de fecha inválido"}
		return c.inertia.Render(ctx, "matches/Edit", renderData)
	}

	match.TournamentID = uint(tournamentID)
	match.HomeTeamID = uint(homeTeamID)
	match.AwayTeamID = uint(awayTeamID)
	match.VenueID = uint(venueID)
	match.MatchDate = matchDate
	match.Round = matchday

	facades.Orm().Query().Save(&match)
	inertia.Flash(ctx, "success", "Partido actualizado exitosamente")
	return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+id)
}

func (c *MatchController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
	}
	facades.Orm().Query().Delete(&match)
	inertia.Flash(ctx, "success", "Partido eliminado exitosamente")
	return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
}

func (c *MatchController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().
		With("HomeTeam").With("AwayTeam").With("Tournament").With("Venue").
		Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
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

	// Get available players for both teams
	var homeTeamPlayers []models.Player
	var awayTeamPlayers []models.Player

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
		"isEditable":      match.IsEditable(),
	})
}

func (c *MatchController) RecordResult(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
	}

	if !match.IsEditable() {
		inertia.Flash(ctx, "error", "El partido no se puede editar hasta 24 horas antes de su inicio")
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+id)
	}

	homeScore, _ := strconv.Atoi(ctx.Request().Input("home_score", "0"))
	awayScore, _ := strconv.Atoi(ctx.Request().Input("away_score", "0"))

	match.HomeScore = homeScore
	match.AwayScore = awayScore
	match.Status = "completed"

	facades.Orm().Query().Save(&match)

	c.updateStandings(&match)

	inertia.Flash(ctx, "success", "Resultado registrado exitosamente")
	return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+id)
}

func (c *MatchController) CloseMatch(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
	}

	if !match.IsEditable() {
		inertia.Flash(ctx, "error", "El partido no se puede cerrar hasta 24 horas antes de su inicio")
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+id)
	}

	match.Status = "completed"
	facades.Orm().Query().Save(&match)

	c.updateStandings(&match)

	inertia.Flash(ctx, "success", "Partido cerrado exitosamente")
	return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+id)
}

func (c *MatchController) AddEvent(ctx http.Context) http.Response {
	matchID := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, matchID); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
	}

	if !match.IsEditable() {
		inertia.Flash(ctx, "error", "No se pueden agregar eventos hasta 24 horas antes del partido")
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
	}

	playerID, _ := strconv.ParseUint(ctx.Request().Input("player_id"), 10, 64)
	teamID, _ := strconv.ParseUint(ctx.Request().Input("team_id"), 10, 64)
	minute, _ := strconv.Atoi(ctx.Request().Input("minute", "0"))
	eventType := ctx.Request().Input("event_type")

	if playerID == 0 || teamID == 0 || eventType == "" {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
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

	return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
}

func (c *MatchController) DeleteEvent(ctx http.Context) http.Response {
	matchID := ctx.Request().Route("id")
	eventID := ctx.Request().Route("eventId")

	var match models.Match
	if err := facades.Orm().Query().Find(&match, matchID); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
	}

	if !match.IsEditable() {
		inertia.Flash(ctx, "error", "No se pueden eliminar eventos hasta 24 horas antes del partido")
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
	}

	var event models.MatchEvent
	if err := facades.Orm().Query().Where("id = ? AND match_id = ?", eventID, matchID).Find(&event); err != nil || event.ID == 0 {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
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

	return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
}

func (c *MatchController) AddLineup(ctx http.Context) http.Response {
	matchID := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().Find(&match, matchID); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
	}

	if !match.IsEditable() {
		inertia.Flash(ctx, "error", "No se pueden modificar alineaciones hasta 24 horas antes del partido")
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
	}

	playerID, _ := strconv.ParseUint(ctx.Request().Input("player_id"), 10, 64)
	teamID, _ := strconv.ParseUint(ctx.Request().Input("team_id"), 10, 64)
	minutesPlayed, _ := strconv.Atoi(ctx.Request().Input("minutes_played", "90"))
	isStarter := ctx.Request().Input("is_starter") == "true"

	if playerID == 0 || teamID == 0 {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
	}

	// Check if player is already in lineup
	var existing models.MatchLineup
	facades.Orm().Query().Where("match_id = ? AND player_id = ?", matchID, playerID).First(&existing)
	if existing.ID > 0 {
		existing.MinutesPlayed = minutesPlayed
		existing.IsStarter = isStarter
		facades.Orm().Query().Save(&existing)
	} else {
		lineup := models.MatchLineup{
			MatchID:       match.ID,
			PlayerID:      uint(playerID),
			TeamID:        uint(teamID),
			IsStarter:     isStarter,
			MinutesPlayed: minutesPlayed,
		}
		facades.Orm().Query().Create(&lineup)
	}

	return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
}

func (c *MatchController) RemoveLineup(ctx http.Context) http.Response {
	matchID := ctx.Request().Route("id")
	playerID := ctx.Request().Route("playerId")

	var match models.Match
	if err := facades.Orm().Query().Find(&match, matchID); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches")
	}

	if !match.IsEditable() {
		inertia.Flash(ctx, "error", "No se pueden modificar alineaciones hasta 24 horas antes del partido")
		return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
	}

	facades.Orm().Query().Where("match_id = ? AND player_id = ?", matchID, playerID).Delete(&models.MatchLineup{})

	return ctx.Response().Redirect(http.StatusSeeOther, "/matches/"+matchID)
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
