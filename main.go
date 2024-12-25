package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sermer07/panel/handlers"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Запуск http сервера")
	router := NewRouter()
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/lol", handlers.HomeHandler)
	http.ListenAndServe(":9090", nil)
}
