package nipper

import (
	"github.com/gin-gonic/gin"
)

type HttpRoute struct {
	router *HttpRouter
	route  string
}

func (route *HttpRoute) Route() string {
	return route.route
}

func (route *HttpRoute) Get(handler interface{}) Route {
	injectInto(route.router.ginrouter.GET, route, handler)
	return route
}

func (route *HttpRoute) Post(handler interface{}) Route {
	injectInto(route.router.ginrouter.POST, route, handler)
	return route
}

func (route *HttpRoute) Put(handler interface{}) Route {
	injectInto(route.router.ginrouter.PUT, route, handler)
	return route
}

func (route *HttpRoute) Delete(handler interface{}) Route {
	injectInto(route.router.ginrouter.DELETE, route, handler)
	return route
}

func (route *HttpRoute) Patch(handler interface{}) Route {
	injectInto(route.router.ginrouter.PATCH, route, handler)
	return route
}

func (route *HttpRoute) Head(handler interface{}) Route {
	injectInto(route.router.ginrouter.HEAD, route, handler)
	return route
}

func (route *HttpRoute) Options(handler interface{}) Route {
	injectInto(route.router.ginrouter.OPTIONS, route, handler)
	return route
}

type HttpRouter struct {
	ginrouter gin.IRouter
}

func NewHttpRouter(ginrouter gin.IRouter) Router {
	return &HttpRouter{
		ginrouter: ginrouter,
	}
}

func (router *HttpRouter) Route(route string) Route {
	return &HttpRoute{router, route}
}
