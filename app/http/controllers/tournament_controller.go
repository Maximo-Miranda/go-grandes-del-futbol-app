package controllers

import (
	"fmt"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/http/requests/tournament"
	"grandesdelfutbol/app/inertia"
	"grandesdelfutbol/app/models"
)

type TournamentController struct {
	inertia *inertia.Inertia
}

func NewTournamentController() *TournamentController {
	return &TournamentController{inertia: inertia.New()}
}

func (c *TournamentController) Index(ctx http.Context) http.Response {
	var tournaments []models.Tournament
	facades.Orm().Query().With("Venue").Order("created_at DESC").Find(&tournaments)

	return c.inertia.Render(ctx, "tournaments/Index", map[string]any{
		"tournaments": tournaments,
	})
}

func (c *TournamentController) Create(ctx http.Context) http.Response {
	var venues []models.Venue
	facades.Orm().Query().Find(&venues)

	return c.inertia.Render(ctx, "tournaments/Create", map[string]any{
		"venues": venues,
	})
}

func (c *TournamentController) Store(ctx http.Context) http.Response {
	var venues []models.Venue
	facades.Orm().Query().Find(&venues)

	var request tournament.StoreTournamentRequest
	validationErrors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "tournaments/Create", map[string]any{
			"venues": venues,
			"errors": map[string]string{"name": "Error de validación"},
		})
	}

	if validationErrors != nil {
		return c.inertia.Render(ctx, "tournaments/Create", map[string]any{
			"venues": venues,
			"errors": inertia.ValidationErrors(validationErrors.All()),
		})
	}

	startDate, _ := time.Parse("2006-01-02", request.StartDate)
	endDate, _ := time.Parse("2006-01-02", request.EndDate)

	var venueID uint
	fmt.Sscanf(request.VenueID, "%d", &venueID)

	newTournament := models.Tournament{
		Name:        request.Name,
		Description: request.Description,
		Format:      request.Format,
		GameType:    request.GameType,
		Status:      "draft",
		StartDate:   startDate,
		EndDate:     endDate,
		VenueID:     venueID,
	}

	if err := facades.Orm().Query().Create(&newTournament); err != nil {
		return c.inertia.Render(ctx, "tournaments/Create", map[string]any{
			"venues": venues,
			"errors": map[string]string{"name": "Error al crear el torneo"},
		})
	}

	return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments")
}

func (c *TournamentController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var tournamentModel models.Tournament
	if err := facades.Orm().Query().With("Venue").With("Groups").Find(&tournamentModel, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments")
	}

	var teams []models.TournamentTeam
	facades.Orm().Query().With("Team").With("Group").Where("tournament_id = ?", id).Find(&teams)

	var matches []models.Match
	facades.Orm().Query().With("HomeTeam").With("AwayTeam").Where("tournament_id = ?", id).Order("match_date ASC").Find(&matches)

	var standings []models.Standing
	facades.Orm().Query().With("Team").With("Group").Where("tournament_id = ?", id).Order("points DESC, goal_difference DESC").Find(&standings)

	return c.inertia.Render(ctx, "tournaments/Show", map[string]any{
		"tournament": tournamentModel,
		"teams":      teams,
		"matches":    matches,
		"standings":  standings,
	})
}

func (c *TournamentController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var tournamentModel models.Tournament
	if err := facades.Orm().Query().Find(&tournamentModel, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments")
	}

	var venues []models.Venue
	facades.Orm().Query().Find(&venues)

	return c.inertia.Render(ctx, "tournaments/Edit", map[string]any{
		"tournament": tournamentModel,
		"venues":     venues,
	})
}

func (c *TournamentController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var tournamentModel models.Tournament
	if err := facades.Orm().Query().Find(&tournamentModel, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments")
	}

	var venues []models.Venue
	facades.Orm().Query().Find(&venues)

	var request tournament.UpdateTournamentRequest
	validationErrors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return c.inertia.Render(ctx, "tournaments/Edit", map[string]any{
			"tournament": tournamentModel,
			"venues":     venues,
			"errors":     map[string]string{"name": "Error de validación"},
		})
	}

	if validationErrors != nil {
		return c.inertia.Render(ctx, "tournaments/Edit", map[string]any{
			"tournament": tournamentModel,
			"venues":     venues,
			"errors":     inertia.ValidationErrors(validationErrors.All()),
		})
	}

	tournamentModel.Name = request.Name
	tournamentModel.Description = request.Description
	tournamentModel.Format = request.Format
	tournamentModel.GameType = request.GameType

	startDate, _ := time.Parse("2006-01-02", request.StartDate)
	endDate, _ := time.Parse("2006-01-02", request.EndDate)
	tournamentModel.StartDate = startDate
	tournamentModel.EndDate = endDate

	var venueID uint
	fmt.Sscanf(request.VenueID, "%d", &venueID)
	tournamentModel.VenueID = venueID

	facades.Orm().Query().Save(&tournamentModel)
	return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments/"+id)
}

func (c *TournamentController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var tournamentModel models.Tournament
	if err := facades.Orm().Query().Find(&tournamentModel, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments")
	}
	facades.Orm().Query().Delete(&tournamentModel)
	return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments")
}

func (c *TournamentController) UpdateStatus(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var tournamentModel models.Tournament
	if err := facades.Orm().Query().Find(&tournamentModel, id); err != nil {
		return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments")
	}

	tournamentModel.Status = ctx.Request().Input("status", tournamentModel.Status)
	facades.Orm().Query().Save(&tournamentModel)

	return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments/"+id)
}

func (c *TournamentController) AddTeam(ctx http.Context) http.Response {
	tournamentID := ctx.Request().Route("id")
	teamID := ctx.Request().Input("team_id")

	var tournamentModel models.Tournament
	if err := facades.Orm().Query().Find(&tournamentModel, tournamentID); err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{"error": "Torneo no encontrado"})
	}

	var teamIDUint uint
	fmt.Sscanf(teamID, "%d", &teamIDUint)

	tournamentTeam := models.TournamentTeam{
		TournamentID: tournamentModel.ID,
		TeamID:       teamIDUint,
	}

	if err := facades.Orm().Query().Create(&tournamentTeam); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": "Error al agregar equipo"})
	}

	return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments/"+tournamentID)
}

func (c *TournamentController) RemoveTeam(ctx http.Context) http.Response {
	tournamentID := ctx.Request().Route("id")
	teamID := ctx.Request().Route("teamId")

	facades.Orm().Query().Where("tournament_id = ? AND team_id = ?", tournamentID, teamID).Delete(&models.TournamentTeam{})

	return ctx.Response().Redirect(http.StatusSeeOther, "/tournaments/"+tournamentID)
}

func (c *TournamentController) AvailableTeams(ctx http.Context) http.Response {
	tournamentID := ctx.Request().Route("id")

	var tournamentTeams []models.TournamentTeam
	facades.Orm().Query().Where("tournament_id = ?", tournamentID).Find(&tournamentTeams)

	var existingTeamIDs []any
	for _, tt := range tournamentTeams {
		existingTeamIDs = append(existingTeamIDs, tt.TeamID)
	}

	var teams []models.Team
	if len(existingTeamIDs) > 0 {
		facades.Orm().Query().WhereNotIn("id", existingTeamIDs).Find(&teams)
	} else {
		facades.Orm().Query().Find(&teams)
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"teams": teams})
}
