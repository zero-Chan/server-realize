package router

import (
	"fmt"
	"net/http"

	//	"entry"
)

var (
	DefaultHTTPRouter HTTPRouter
)

func init() {
	//  实现 DefaultHTTPRouter 服务发现功能
}

type HTTPRouter struct {
	address string
	mux     *http.ServeMux
}

func NewHTTPRouter(addr string) *HTTPRouter {
	hr := &HTTPRouter{
		mux:     http.NewServeMux(),
		address: addr,
	}

	return hr
}

func (r *HTTPRouter) Register(e Entry) {

	hdl := func(w http.ResponseWriter, req *http.Request) {
		if e.RouterRule(req) == false {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		e.ServeHTTP(w, req)
	}

	r.mux.HandleFunc(e.Path(), hdl)
}

func (r *HTTPRouter) ResetAddr(addr string) {
	r.address = addr
}

func (r *HTTPRouter) Run() error {
	fmt.Printf("Listen: [%s]\n", r.address)

	return http.ListenAndServe(r.address, r.mux)
}
