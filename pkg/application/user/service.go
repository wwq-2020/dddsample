package user

import (
	"github.com/google/wire"
	"github.com/wwq1988/dddsampple/pkg/domain/model/user"
	"github.com/wwq1988/errors"
)

type Service interface {
	Register(name, password string) error
}

type service struct {
	repo user.Repo
}

var SuperSet = wire.NewSet(New, user.New)

func New(repo user.Repo) Service {
	return &service{repo: repo}
}

func (s *service) Register(name, password string) error {
	user := &user.User{Name: name, Password: password}
	if err := s.repo.Create(user); err != nil {
		return errors.Trace(err)
	}
	return nil
}
