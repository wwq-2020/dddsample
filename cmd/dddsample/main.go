package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wwq1988/dddsampple/pkg/interfaces"
	userWeb "github.com/wwq1988/dddsampple/pkg/interfaces/user/web"
	"github.com/wwq1988/errors"
)

func main() {
	server := CreateServer()
	if err := server.Serve(); err != nil {
		logrus.WithFields(errors.Fields(err).KVs()).Error(err)
	}
}

func webs() []interfaces.Web {
	userWeb := userWeb.Create()
	return []interfaces.Web{userWeb}
}
