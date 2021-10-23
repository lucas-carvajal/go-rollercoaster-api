package main

import (
	"go-rollercoaster-api/handlers"
	"net/http"
)

func main() {
	admin := handlers.NewAdminPortal()
	coasterHandlers := handlers.NewCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandlers.Coasters)
	http.HandleFunc("/coasters/", coasterHandlers.GetCoaster)
	http.HandleFunc("/addCoaster", coasterHandlers.AddCoaster)
	http.HandleFunc("/admin", admin.Handler)
	http.HandleFunc("/", handlers.LandingHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
