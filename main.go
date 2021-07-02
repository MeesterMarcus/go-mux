package main

import (
	"github.com/MeesterMarcus/go-mux/config"
	"github.com/MeesterMarcus/go-mux/controllers"
)

func main() {
	db, _ := config.ConnectToDB()
	controllers.HandleRequests(db)
}
