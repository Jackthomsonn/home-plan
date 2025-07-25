package main

import (
	"github.com/jackthomsonn/home-plan/connect"
	"github.com/jackthomsonn/home-plan/internal"
)

func main() {
	db, err := internal.CreateConnection()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&internal.Device{})

	connectionSvc := connect.NewConnectionService(":8080", db)
	connectionSvc.Start()
}
