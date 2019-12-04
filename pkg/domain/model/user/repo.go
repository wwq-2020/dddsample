package user

import (
	"github.com/wwq1988/dddsampple/pkg/conf"
)

type Repo interface {
	Create(*User) error
}

type repo struct{}

func New(conf *conf.Config) Repo {
	return &repo{}
}

func (rp *repo) Create(*User) error {
	return nil
}
