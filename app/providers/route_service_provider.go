package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/http"
	"grandesdelfutbol/routes"
)

type RouteServiceProvider struct{}

func (receiver *RouteServiceProvider) Register(app foundation.Application) {}

func (receiver *RouteServiceProvider) Boot(app foundation.Application) {
	facades.Route().GlobalMiddleware(http.Kernel{}.Middleware()...)
	routes.Web()
}
