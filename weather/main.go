package weather

import (
	"net/http"
)

type Weather struct {
	addr string
}

func NewWeatherService(addr string) *Weather {
	return &Weather{
		addr: addr,
	}
}

func (svc *Weather) setupRoutes() {
	http.HandleFunc("/optimise", HandleWeather)
}

func (svc *Weather) Start() {
	svc.setupRoutes()
	http.ListenAndServe(svc.addr, nil)
}
