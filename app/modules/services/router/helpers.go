package router

import (
	"github.com/peyman-abdi/avalanche/app/interfaces/services"
	"github.com/peyman-abdi/fasthttp-routing"
)

func handleNotFound(c *routing.Context) error {
	c.Error("Page not found!", 404)
	return nil
}

func (r *routerImpl) handleRoute(route *services.Route) func(context *routing.Context) error {
	return func(context *routing.Context) error {
		if route.Verify != nil {
			if err := route.Verify(r.requestFromContext(context)); err != nil {
				return err
			}
		}

		if route.Handle != nil {
			return route.Handle(r.requestFromContext(context), r.responseFromContext(context))
		}

		return nil
	}
}

func (r *routerImpl) handleGroups(wares []*services.RouteGroup) []routing.Handler {
	if len(wares) > 0 {
		var methods = make([]routing.Handler, len(wares))
		for index, ware := range wares {
			methods[index] = r.handleMiddleWare(ware.Handler)
		}
		return methods
	}

	return []routing.Handler{r.handleEmpty}
}
func (r *routerImpl) handleMiddleWares(wares []*services.MiddleWare) []routing.Handler {
	if len(wares) > 0 {
		var methods = make([]routing.Handler, len(wares))
		for index, ware := range wares {
			methods[index] = r.handleMiddleWare(ware.Handler)
		}
		return methods
	}

	return []routing.Handler{r.handleEmpty}
}
func (r *routerImpl) handleMiddleWare(callback services.RequestHandler) routing.Handler {
	if callback != nil {
		return func(context *routing.Context) error {
			return callback(r.requestFromContext(context), r.responseFromContext(context))
		}
	}

	return r.handleEmpty
}

func (r *routerImpl) handleEmpty(context *routing.Context) error {
	return nil
}

func (r *routerImpl) methodsFromInt(method int) (methods []string) {
	if method&services.GET != 0 {
		methods = append(methods, "GET")
	}
	if method&services.POST != 0 {
		methods = append(methods, "POST")
	}
	if method&services.PUT != 0 {
		methods = append(methods, "PUT")
	}
	if method&services.DELETE != 0 {
		methods = append(methods, "DELETE")
	}
	return
}
func (r *routerImpl) responseFromContext(ctx *routing.Context) services.Response {
	return &responseImpl{context: ctx, log: r.log, engine: r.engine}
}
