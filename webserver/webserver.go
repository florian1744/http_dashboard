package webserver

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	systemutil "example.com/http_dashboard/system_util"
)

func StartWebServer() {
	tmpl := template.Must(template.ParseFiles("dashboard.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		systemResources, err := systemutil.BuildDashboardData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, systemResources); err != nil {
			log.Println("template error:", err)
		}
	})

	http.HandleFunc("/api/system", func(w http.ResponseWriter, r *http.Request) {
		systemResources, err := systemutil.BuildDashboardData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(systemResources); err != nil {
			log.Println("json error", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
