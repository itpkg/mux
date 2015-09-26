package mux

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

type Route struct {
	method    string
	pattern   string
	regexp    *regexp.Regexp
	functions []Handler
}

func (p *Route) String() string {
	names := make([]string, 0)
	show := func(fns []Handler) {
		for _, f := range fns {
			names = append(names, p.funcName(f))
		}
	}
	show(p.functions)
	return fmt.Sprintf("%s\t%s\n\t%s", p.method, p.pattern, strings.Join(names, "\t"))
}

func (p *Route) Match(req *http.Request) map[string]string {
	if p.method != req.Method {
		return nil
	}
	values := p.regexp.FindStringSubmatch(req.URL.Path)
	if values == nil {
		return nil
	}
	params := make(map[string]string, 0)
	for i, n := range p.regexp.SubexpNames() {
		params[n] = values[i]
	}
	return params
}

func (p *Route) funcName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
