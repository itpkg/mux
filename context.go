package mux

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/itpkg/log"
)

const (
	MT_XML  = "application/xml"
	MT_JSON = "application/json"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Logger   log.Logger
}

func (p *Context) BYTES(contentType string, body []byte) {
	wrt := p.Response
	wrt.Header().Set("Content-Type", contentType)
	wrt.Write(body)
}

func (p *Context) JSON(obj interface{}) error {
	wrt := p.Response
	wrt.Header().Set("Content-Type", MT_JSON)
	return json.NewEncoder(wrt).Encode(obj)
}

func (p *Context) XML(obj interface{}) error {
	wrt := p.Response
	wrt.Header().Set("Content-Type", MT_XML)
	return xml.NewEncoder(wrt).Encode(obj)
}

func (p *Context) Abort(code int, err error) {
	p.Logger.Error(err.Error())

	wrt := p.Response
	wrt.WriteHeader(code)
	fmt.Fprintf(wrt, err.Error())
}

func (p *Context) Unauthorized(err error) {
	p.Abort(http.StatusUnauthorized, err)
}

func (p *Context) BadRequest(err error) {
	p.Abort(http.StatusBadRequest, err)
}

func (p *Context) InternalServerError(err error) {
	p.Abort(http.StatusInternalServerError, err)
}

func (p *Context) Forbidden(err error) {
	p.Abort(http.StatusForbidden, err)
}

func (p *Context) NotFound(err error) {
	p.Abort(http.StatusNotFound, err)
}
