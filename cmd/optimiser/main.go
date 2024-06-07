package main

import (
	"github.com/jackthomsonn/home-plan/grid"
	"github.com/jackthomsonn/home-plan/optimiser"
	"github.com/jackthomsonn/home-plan/weather"
)

func main() {
	gridSvc := grid.NewGridService()
	weatherSvc := weather.NewWeatherService()
	optimiserSvc := optimiser.NewOptimiserService(*gridSvc, *weatherSvc, ":8080")

	optimiserSvc.Start()
}
