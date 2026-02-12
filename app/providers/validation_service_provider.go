package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/rules"
)

type ValidationServiceProvider struct{}

func (receiver *ValidationServiceProvider) Register(app foundation.Application) {}
func (receiver *ValidationServiceProvider) Boot(app foundation.Application) {
	facades.Validation().AddRules([]validation.Rule{
		&rules.MaxFileSize{},
	})
}
