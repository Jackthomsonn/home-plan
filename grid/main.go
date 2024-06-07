package grid

import (
	"net/http"
)

type Grid struct {
	addr string
}

func NewGridService(addr string) *Grid {
	return &Grid{
		addr: addr,
	}
}

func (svc *Grid) setupRoutes() {
	http.HandleFunc("/optimise", HandleGrid)
}

func (svc *Grid) Start() {
	svc.setupRoutes()
	http.ListenAndServe(svc.addr, nil)
}
