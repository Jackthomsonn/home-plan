package connect

import (
	"net/http"
)

type Connect struct {
	addr string
}

func NewConnectionService(addr string) *Connect {
	return &Connect{
		addr: addr,
	}
}

func (svc *Connect) setupRoutes() {
	http.HandleFunc("/connect", HandleConnect)
}

func (svc *Connect) Start() {
	svc.setupRoutes()
	http.ListenAndServe(svc.addr, nil)
}
