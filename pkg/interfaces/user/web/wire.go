//+build wireinject

package web

import (
	"github.com/google/wire"
	"github.com/wwq1988/dddsampple/pkg/conf"
	"github.com/wwq1988/dddsampple/pkg/interfaces"
	"github.com/wwq1988/dddsampple/pkg/interfaces/user/facade"
)

func Create() interfaces.Web {
	panic(wire.Build(conf.Parse, New, facade.SuperSet))
}
