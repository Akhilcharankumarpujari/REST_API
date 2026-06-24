package main

import (
	"log"
	"net/http"

	"restapi/config"
	"restapi/database"
	"restapi/routes"
)

func main() {
	config.Load()
	database.Connect()

	mux := http.NewServeMux()
	routes.Setup(mux)

	addr := ":" + config.AppPort
	log.Println("REST API running at http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
