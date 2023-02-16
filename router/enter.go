package router

import (
	"go-generate/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
