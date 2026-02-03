package controllers

import (
	"strconv"

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

	return c.inertia.Render(ctx, "matches/Index", map[string]any{
		"matches":      matches,
		"tournamentId": tournamentID,
	})
}

func (c *MatchController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var match models.Match
	if err := facades.Orm().Query().
		With("HomeTeam").With("AwayTeam").With("Tournament").With("Venue").
		With("Events").With("Lineups").
		Find(&match, id); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/matches")
	}

	return c.inertia.Render(ctx, "matches/Show", map[string]any{
		"match": match,
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
