package routes

import (
	"github.com/pocketbase/pocketbase/core"
)

type Path = string
type ServerEventFunc = func(e *core.ServeEvent) error

// type Router struct {
// 	pb *pocketbase.PocketBase
// }

// func NewRouter(pb *pocketbase.PocketBase) *Router {
// 	return &Router{pb}
// }
