package main

import (
	"github.com/jackthomsonn/home-plan/grid"
	"github.com/jackthomsonn/home-plan/internal"
)

func main() {
	db, err := internal.CreateConnection()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&internal.Grid{})

	connectionSvc := grid.NewGridService(":8080", db)
	connectionSvc.Start()
}
