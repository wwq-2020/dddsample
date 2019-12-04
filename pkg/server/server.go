package server

import (
	"net/http"

	"github.com/google/wire"
	"github.com/wwq1988/dddsampple/pkg/conf"
	"github.com/wwq1988/dddsampple/pkg/interfaces"
	"github.com/wwq1988/errors"
)

type Server struct {
	conf    *conf.Config
	handler http.Handler
}

var SuperSet = wire.NewSet(New, interfaces.Webs)

func New(conf *conf.Config, webs ...interfaces.Web) *Server {
	for _, each := range webs {
		each.Register()
	}
	server := &Server{conf: conf}
	return server
}

func (s *Server) Serve() error {
	if err := http.ListenAndServe(s.conf.Addr, nil); err != nil {
		return errors.Trace(err)
	}
	return nil
}
