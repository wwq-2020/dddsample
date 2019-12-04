package web

import (
	"encoding/json"
	"net/http"

	"github.com/wwq1988/dddsampple/pkg/interfaces"
	"github.com/wwq1988/dddsampple/pkg/interfaces/user/facade"
)

type Web struct {
	facade facade.Facade
}

type registerReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type registerResp struct {
}

func New(facade facade.Facade) interfaces.Web {
	web := &Web{facade: facade}
	return web
}

func (web *Web) Register() {
	http.HandleFunc("/register", web.register)
}

func (web *Web) register(w http.ResponseWriter, r *http.Request) {
	req := &registerReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := web.facade.Register(req.Name, req.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func init() {
	interfaces.Register(Create())
}
