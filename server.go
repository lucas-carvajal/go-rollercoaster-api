package main

import (
	"fmt"
	"go-rollercoaster-api/handlers"
	"net/http"
)

const port int = 8080

func main() {
	admin := handlers.NewAdminPortal()
	coasterHandlers := handlers.NewCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandlers.Coasters)
	http.HandleFunc("/coasters/", coasterHandlers.GetCoaster)
	http.HandleFunc("/addCoaster", coasterHandlers.AddCoaster)
	http.HandleFunc("/admin", admin.Handler)
	http.HandleFunc("/", handlers.LandingHandler)
	fmt.Printf("Listening on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
