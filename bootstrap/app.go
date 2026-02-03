package bootstrap

import (
	"github.com/goravel/framework/foundation"

	"grandesdelfutbol/config"
)

func Boot() {
	app := foundation.NewApplication()
	app.Boot()
	config.Boot()
}
