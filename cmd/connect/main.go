package main

import (
	"github.com/jackthomsonn/home-plan/connect"
)

func main() {
	connectionSvc := connect.NewConnectionService(":8080")
	connectionSvc.Start()
}
