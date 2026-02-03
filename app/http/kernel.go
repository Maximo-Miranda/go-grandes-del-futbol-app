package http

import "github.com/goravel/framework/contracts/http"

type Kernel struct{}

func (kernel Kernel) Middleware() []http.Middleware {
	return []http.Middleware{}
}
