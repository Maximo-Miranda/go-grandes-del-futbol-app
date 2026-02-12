package providers

import (
	"fmt"
	"os"

	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

type AppServiceProvider struct{}

func (receiver *AppServiceProvider) Register(app foundation.Application) {}

func (receiver *AppServiceProvider) Boot(app foundation.Application) {
	jwtSecret := facades.Config().GetString("jwt.secret")
	if jwtSecret == "" {
		if facades.Config().GetString("app.env") != "local" {
			fmt.Fprintln(os.Stderr, "FATAL: JWT_SECRET environment variable must be set in non-local environments")
			os.Exit(1)
		}
		facades.Log().Warning("JWT_SECRET is not set. Authentication will not work securely.")
	}
}
