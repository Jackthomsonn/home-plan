package optimiser

import (
	"net/http"

	"github.com/jackthomsonn/home-plan/connections"
	"github.com/jackthomsonn/home-plan/grid"
	"github.com/jackthomsonn/home-plan/weather"
)

type Optimiser struct {
	connection connections.Connection
	addr       string
	gridSvc    grid.Grid
	weatherSvc weather.Weather
}

type OptimiserRequest struct {
	Type string `json:"type"`
}

func NewOptimiserService(grid grid.Grid, weather weather.Weather, addr string) *Optimiser {
	return &Optimiser{
		addr:       addr,
		gridSvc:    grid,
		weatherSvc: weather,
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
