package main

import (
	"fmt"
	"goapi/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go API Server...")

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}
}
