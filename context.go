package mux

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/itpkg/log"
)

type Response struct {
	Header map[string]string
	Body []byte
}

const (
	MT_XML  = "application/xml"
	MT_JSON = "application/json"
)

type Context struct {
	Request  *http.Request
	Writer http.ResponseWriter
	Params map[string]interface {}
	Response *Response
	Logger   log.Logger
}

func (p *Context) Write(){
	for k, v := range p.Response.Header{
		p.Writer.Header().Set(k,v)
	}
	p.Request.Write(p.Response.Body)
}

func (p *Context) BYTES(contentType string, body []byte) {
	wrt := p.Writer
	wrt.Header().Set("Content-Type", contentType)
	wrt.Write(body)
}

func (p *Context) JSON(obj interface{}) error {
	wrt := p.Writer
	wrt.Header().Set("Content-Type", MT_JSON)
	return json.NewEncoder(wrt).Encode(obj)
}

func (p *Context) XML(obj interface{}) error {
	wrt := p.Writer
	wrt.Header().Set("Content-Type", MT_XML)
	return xml.NewEncoder(wrt).Encode(obj)
}

func (p *Context) Abort(code int, err error) {
	p.Logger.Error(err.Error())

	wrt := p.Writer
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
