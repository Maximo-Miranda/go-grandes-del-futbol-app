package providers

import "github.com/goravel/framework/contracts/foundation"

type GrpcServiceProvider struct{}

func (receiver *GrpcServiceProvider) Register(app foundation.Application) {}
func (receiver *GrpcServiceProvider) Boot(app foundation.Application)     {}
