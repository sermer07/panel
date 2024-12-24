package panel

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com./sermer07/panel/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ViewData struct {
	Title   string
	Message string
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Запуск http сервера")
	// logger.Warn("Warning message")
	// logger.Error("Error message")

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/lol", handlers.HomeHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title:   "World Cup",
			Message: "FIFA will never regret it",
		}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":9090", nil)
}
