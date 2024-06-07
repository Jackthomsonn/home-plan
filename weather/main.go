package weather

import (
	"net/http"

	"gorm.io/gorm"
)

type Weather struct {
	addr string
	db   *gorm.DB
}

func NewWeatherService(addr string, db *gorm.DB) *Weather {
	return &Weather{
		addr: addr,
		db:   db,
	}
}

func (svc *Weather) setupRoutes() {
	http.HandleFunc("/weather", HandleWeather)
}

func (svc *Weather) Start() {
	svc.setupRoutes()
	http.ListenAndServe(svc.addr, nil)
}
