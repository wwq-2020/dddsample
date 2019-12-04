//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/wwq1988/dddsampple/pkg/conf"
	"github.com/wwq1988/dddsampple/pkg/server"
)

func CreateServer() *server.Server {
	panic(wire.Build(conf.Parse, server.SuperSet))
}
