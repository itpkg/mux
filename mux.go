package mux

import (
	"errors"
	"net/http"

	"github.com/itpkg/log"
)

type Handler func(*Context) error

type Mux struct {
	routers []*Router

	Logger log.Logger `inject:""`
}

func (p *Mux) ServeHTTP(wrt http.ResponseWriter, req *http.Request) {
	ctx := &Context{
		Request:  req,
		Response: wrt,
		Params: make(map[string]interface {}, 0),
		Logger:   p.Logger,
	}

	p.Logger.Info("%s %s", req.Method, req.URL.Path)
	for _, router := range p.routers {
		for _, route := range router.routes {
			if params := route.Match(req); params != nil {
				for k, v := range params{
					ctx.Params[k] = v
				}
				p.Logger.Debug("Match %s", route.pattern)
				for _, fn := range route.functions {
					if err := fn(ctx); err != nil {
						p.Logger.Error("http error %v", err)
						break
					}
				}
				return
			}
		}
	}

	ctx.BadRequest(errors.New("Bad Http Request."))
}

func (p *Mux) AddRouter(router *Router) {
	p.routers = append(p.routers, router)
}
