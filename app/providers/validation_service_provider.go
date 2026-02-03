package providers

import "github.com/goravel/framework/contracts/foundation"

type ValidationServiceProvider struct{}

func (receiver *ValidationServiceProvider) Register(app foundation.Application) {}
func (receiver *ValidationServiceProvider) Boot(app foundation.Application)     {}
