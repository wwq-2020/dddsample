package facade

import (
	"github.com/google/wire"
	"github.com/wwq1988/dddsampple/pkg/application/user"
	"github.com/wwq1988/errors"
)

type Facade interface {
	Register(name, password string) error
}

type facade struct {
	userService user.Service
}

var SuperSet = wire.NewSet(New, user.SuperSet)

func New(userService user.Service) Facade {
	return &facade{userService: userService}
}

func (f *facade) Register(name, password string) error {
	if err := f.userService.Register(name, password); err != nil {
		return errors.Trace(err)
	}
	return nil
}
