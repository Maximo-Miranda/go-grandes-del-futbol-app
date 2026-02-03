package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"

	authControllers "grandesdelfutbol/app/http/controllers/auth"
	"grandesdelfutbol/app/http/controllers"
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

		// Authenticated routes
		router.Middleware(middleware.Authenticate()).Group(func(router route.Router) {
			// Dashboard
			dashboardController := controllers.NewDashboardController()
			router.Get("/dashboard", dashboardController.Index)

			// Venues
			venueController := controllers.NewVenueController()
			router.Get("/venues", venueController.Index)
			router.Get("/venues/create", venueController.Create)
			router.Post("/venues", venueController.Store)
			router.Get("/venues/{id}/edit", venueController.Edit)
			router.Put("/venues/{id}", venueController.Update)
			router.Delete("/venues/{id}", venueController.Destroy)

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

			// Teams
			teamController := controllers.NewTeamController()
			router.Get("/teams", teamController.Index)
			router.Get("/teams/create", teamController.Create)
			router.Post("/teams", teamController.Store)
			router.Get("/teams/{id}", teamController.Show)
			router.Get("/teams/{id}/edit", teamController.Edit)
			router.Put("/teams/{id}", teamController.Update)
			router.Delete("/teams/{id}", teamController.Destroy)
			router.Post("/teams/{id}/players", teamController.AddPlayer)
			router.Delete("/teams/{id}/players/{playerId}", teamController.RemovePlayer)
			router.Get("/teams/{id}/available-players", teamController.AvailablePlayers)

			// Players
			playerController := controllers.NewPlayerController()
			router.Get("/players", playerController.Index)
			router.Get("/players/create", playerController.Create)
			router.Post("/players", playerController.Store)
			router.Get("/players/{id}", playerController.Show)
			router.Get("/players/{id}/edit", playerController.Edit)
			router.Put("/players/{id}", playerController.Update)
			router.Delete("/players/{id}", playerController.Destroy)

			// Matches
			matchController := controllers.NewMatchController()
			router.Get("/matches", matchController.Index)
			router.Get("/matches/{id}", matchController.Show)
			router.Post("/matches/{id}/result", matchController.RecordResult)

			// Standings
			standingController := controllers.NewStandingController()
			router.Get("/standings", standingController.Index)
		})
	})
}
