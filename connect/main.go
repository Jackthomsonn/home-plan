package connect

import (
	"net/http"

	"gorm.io/gorm"
)

type Connect struct {
	addr string
	db   *gorm.DB
}

func NewConnectionService(addr string, db *gorm.DB) *Connect {
	return &Connect{
		addr: addr,
		db:   db,
	}
}

func (svc *Connect) setupRoutes() {
	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		HandleConnect(w, r, *svc)
	})
}

func (svc *Connect) Start() {
	svc.setupRoutes()
	http.ListenAndServe(svc.addr, nil)
}
