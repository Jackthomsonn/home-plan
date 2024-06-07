package main

import (
	"github.com/jackthomsonn/home-plan/internal"
	"github.com/jackthomsonn/home-plan/optimiser"
)

func main() {
	db, err := internal.CreateConnection()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&internal.Plan{})

	optimiserSvc := optimiser.NewOptimiserService(":8080", db)
	optimiserSvc.Start()
}
