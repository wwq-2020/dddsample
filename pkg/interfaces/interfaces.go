package interfaces

type Web interface {
	Register()
}

var webs []Web

func Webs() []Web {
	return webs
}

func Register(web Web) {
	webs = append(webs, web)
}
