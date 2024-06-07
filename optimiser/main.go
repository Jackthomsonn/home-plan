package optimiser

import (
	"net/http"

	"github.com/jackthomsonn/home-plan/connections"
	"gorm.io/gorm"
)

type Optimiser struct {
	connection connections.Connection
	addr       string
	db         *gorm.DB
}

type OptimiserRequest struct {
	Type string `json:"type"`
}

func NewOptimiserService(addr string, db *gorm.DB) *Optimiser {
	return &Optimiser{
		addr: addr,
		db:   db,
	}
}

func (svc *Optimiser) setupRoutes() {
	http.HandleFunc("/optimise", func(w http.ResponseWriter, r *http.Request) {
		HandleOptimisation(w, r, *svc)
	})
}

func (svc *Optimiser) Start() {
	svc.setupRoutes()
	http.ListenAndServe(svc.addr, nil)
}
