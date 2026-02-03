package providers

import "github.com/goravel/framework/contracts/foundation"

type QueueServiceProvider struct{}

func (receiver *QueueServiceProvider) Register(app foundation.Application) {}
func (receiver *QueueServiceProvider) Boot(app foundation.Application)     {}
