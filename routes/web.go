package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"

	"grandesdelfutbol/app/http/controllers"
	authControllers "grandesdelfutbol/app/http/controllers/auth"
	"grandesdelfutbol/app/http/middleware"
	"grandesdelfutbol/app/inertia"
	authService "grandesdelfutbol/app/services/auth"
)

func Web() {
	i := inertia.New()

	facades.Route().Static("/build", "./public/build")

	facades.Route().Middleware(middleware.SecurityHeaders(), middleware.CSRF()).Group(func(router route.Router) {
		router.Get("/", func(ctx http.Context) http.Response {
			return i.Render(ctx, "Welcome", map[string]any{
				"version": support.Version,
			})
		})

		authController := authControllers.NewAuthController(authService.NewAuthService())

		router.Prefix("/auth").Middleware(middleware.Guest(), middleware.RateLimitAuth()).Group(func(router route.Router) {
			router.Get("/login", authController.ShowLogin)
			router.Post("/login", authController.Login)
			router.Get("/register", authController.ShowRegister)
			router.Post("/register", authController.Register)
		})

		router.Prefix("/auth").Middleware(middleware.Authenticate()).Group(func(router route.Router) {
			router.Post("/logout", authController.Logout)
		})

		// Profile management (authenticated but no profile required for create)
		profileController := controllers.NewProfileController()
		router.Middleware(middleware.Authenticate()).Group(func(router route.Router) {
			router.Get("/profile/create", profileController.Create)
			router.Post("/profile", profileController.Store)
			router.Get("/profile/edit", profileController.Edit)
			router.Put("/profile", profileController.Update)
			router.Post("/profile/update", profileController.Update) // Method spoofing for file uploads
		})

		// Authenticated routes (require profile)
		router.Middleware(middleware.Authenticate(), middleware.RequireProfile()).Group(func(router route.Router) {
			// Dashboard
			dashboardController := controllers.NewDashboardController()
			router.Get("/dashboard", dashboardController.Index)

			// Venues
			venueController := controllers.NewVenueController()
			router.Get("/venues", venueController.Index)
			router.Get("/venues/create", venueController.Create)
			router.Post("/venues", venueController.Store)
			router.Get("/venues/{uuid}/edit", venueController.Edit)
			router.Put("/venues/{uuid}", venueController.Update)
			router.Post("/venues/{uuid}", venueController.Update) // For method spoofing with file uploads
			router.Delete("/venues/{uuid}", venueController.Destroy)
			router.Delete("/venues/{uuid}/photo", venueController.DestroyPhoto)
			router.Get("/venues/{uuid}/photo", venueController.Photo)

			// Tournaments
			tournamentController := controllers.NewTournamentController()
			router.Get("/tournaments", tournamentController.Index)
			router.Get("/tournaments/create", tournamentController.Create)
			router.Post("/tournaments", tournamentController.Store)
			router.Get("/tournaments/{id}", tournamentController.Show)
			router.Get("/tournaments/{id}/edit", tournamentController.Edit)
			router.Put("/tournaments/{id}", tournamentController.Update)
			router.Delete("/tournaments/{id}", tournamentController.Destroy)
			router.Post("/tournaments/{id}/status", tournamentController.UpdateStatus)
			router.Post("/tournaments/{id}/teams", tournamentController.AddTeam)
			router.Delete("/tournaments/{id}/teams/{teamId}", tournamentController.RemoveTeam)
			router.Get("/tournaments/{id}/available-teams", tournamentController.AvailableTeams)
			router.Post("/tournaments/{id}/groups", tournamentController.CreateGroup)
			router.Put("/tournaments/{id}/groups/{groupId}", tournamentController.UpdateGroup)
			router.Delete("/tournaments/{id}/groups/{groupId}", tournamentController.DeleteGroup)
			router.Put("/tournaments/{id}/teams/{teamId}/group", tournamentController.AssignTeamGroup)

			// Teams
			teamController := controllers.NewTeamController()
			router.Get("/teams", teamController.Index)
			router.Get("/teams/create", teamController.Create)
			router.Post("/teams", teamController.Store)
			router.Get("/teams/{id}", teamController.Show)
			router.Get("/teams/{id}/edit", teamController.Edit)
			router.Put("/teams/{id}", teamController.Update)
			router.Post("/teams/{id}", teamController.Update) // For method spoofing with file uploads
			router.Delete("/teams/{id}", teamController.Destroy)
			router.Get("/teams/{id}/logo", teamController.Logo)
			router.Post("/teams/{id}/players", teamController.AddPlayer)
			router.Delete("/teams/{id}/players/{playerId}", teamController.RemovePlayer)
			router.Get("/teams/{id}/available-players", teamController.AvailablePlayers)

			// Players (read-only)
			playerController := controllers.NewPlayerController()
			router.Get("/players", playerController.Index)
			router.Get("/players/{id}", playerController.Show)
			router.Get("/players/{id}/photo", playerController.Photo)

			// Matches
			matchController := controllers.NewMatchController()
			router.Get("/matches", matchController.Index)
			router.Get("/matches/create", matchController.Create)
			router.Post("/matches", matchController.Store)
			router.Get("/matches/{id}", matchController.Show)
			router.Get("/matches/{id}/edit", matchController.Edit)
			router.Put("/matches/{id}", matchController.Update)
			router.Delete("/matches/{id}", matchController.Destroy)
			router.Post("/matches/{id}/result", matchController.RecordResult)
			router.Post("/matches/{id}/close", matchController.CloseMatch)
			router.Post("/matches/{id}/events", matchController.AddEvent)
			router.Delete("/matches/{id}/events/{eventId}", matchController.DeleteEvent)
			router.Post("/matches/{id}/lineups", matchController.AddLineup)
			router.Delete("/matches/{id}/lineups/{playerId}", matchController.RemoveLineup)

			// Standings
			standingController := controllers.NewStandingController()
			router.Get("/standings", standingController.Index)
		})
	})
}
