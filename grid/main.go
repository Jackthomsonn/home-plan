package grid

import (
	"net/http"

	"gorm.io/gorm"
)

type Grid struct {
	addr string
	db   *gorm.DB
}

func NewGridService(addr string, db *gorm.DB) *Grid {
	return &Grid{
		addr: addr,
		db:   db,
	}
}

func (svc *Grid) setupRoutes() {
	http.HandleFunc("/grid", HandleGrid)
}

func (svc *Grid) Start() {
	svc.setupRoutes()
	http.ListenAndServe(svc.addr, nil)
}
