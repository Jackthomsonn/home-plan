package main

import (
	"github.com/jackthomsonn/home-plan/internal"
	"github.com/jackthomsonn/home-plan/weather"
)

func main() {
	db, err := internal.CreateConnection()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&internal.Weather{})

	connectionSvc := weather.NewWeatherService(":8080", db)
	connectionSvc.Start()
}
