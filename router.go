package mux

import (
	"fmt"
	"regexp"
)

type Rest struct {
	Index  []Handler
	Create []Handler
	Show   []Handler
	Update []Handler
	Delete []Handler
}

type Router struct {
	routes []*Route
}

func (p *Router) GET(pat string, fns ...Handler) {
	p.add("GET", pat, fns...)
}
func (p *Router) POST(pat string, fns ...Handler) {
	p.add("POST", pat, fns...)
}
func (p *Router) PATCH(pat string, fns ...Handler) {
	p.add("PATCH", pat, fns...)
}
func (p *Router) PUT(pat string, fns ...Handler) {
	p.add("PUT", pat, fns...)
}
func (p *Router) DELETE(pat string, fns ...Handler) {
	p.add("DELETE", pat, fns...)
}

func (p *Router) REST(res string, rest Rest) {
	p.GET(fmt.Sprintf("%s/", res), rest.Index...)
	p.POST(fmt.Sprintf("%s/", res), rest.Create...)
	p.GET(fmt.Sprintf(`%s/(?P<id>\d+)`, res), rest.Show...)
	p.PUT(fmt.Sprintf(`%s/(?P<id>\d+)`, res), rest.Update...)
	p.DELETE(fmt.Sprintf(`%s/(?P<id>\d+)`, res), rest.Delete...)
}

func (p *Router) add(act string, pat string, fns ...Handler) error {
	if reg, err := regexp.Compile(pat); err == nil {
		p.routes = append(p.routes, &Route{method: act, pattern: pat, regexp: reg, functions: fns})
		return nil
	} else {
		return err
	}
}

//-----------------------------------------------------------------------------

func NewRouter() *Router {
	return &Router{
		routes: make([]*Route, 0),
	}
}
