package providers

import "github.com/goravel/framework/contracts/foundation"

type ConsoleServiceProvider struct{}

func (receiver *ConsoleServiceProvider) Register(app foundation.Application) {}
func (receiver *ConsoleServiceProvider) Boot(app foundation.Application)     {}
