package providers

import "github.com/goravel/framework/contracts/foundation"

type EventServiceProvider struct{}

func (receiver *EventServiceProvider) Register(app foundation.Application) {}
func (receiver *EventServiceProvider) Boot(app foundation.Application)     {}
