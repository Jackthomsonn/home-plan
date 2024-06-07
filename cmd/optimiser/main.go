package main

import (
	"github.com/jackthomsonn/home-plan/connections"
	"github.com/jackthomsonn/home-plan/grid"
	"github.com/jackthomsonn/home-plan/optimiser"
	"github.com/jackthomsonn/home-plan/weather"
)

func main() {
	nest := connections.NewNestConnection()
	gridSvc := grid.NewGridService(":8081")
	weatherSvc := weather.NewWeatherService(":8082")
	optimiserSvc := optimiser.NewOptimiserService(nest, *gridSvc, *weatherSvc, ":8080")
	optimiserSvc.Start()
}
