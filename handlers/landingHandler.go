package handlers

import (
	"net/http"
	"strings"
)

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")

	if parts[1] == "" {
		w.Header().Add("location","/coasters")
		w.WriteHeader(http.StatusFound)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("<html><h1>Page not found</h1></html>"))
	}
}
