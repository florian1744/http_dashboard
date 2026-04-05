package webserver

import (
	"encoding/json"
	"html/template"
	"net/http"

	systemutil "example.com/http_dashboard/system_util"
)

func StartWebServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		systemResources, err := systemutil.BuildDashboardData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("dashboard.html"))
		tmpl.Execute(w, systemResources)
	})

	http.HandleFunc("/api/system", func(w http.ResponseWriter, r *http.Request) {
		systemResources, err := systemutil.BuildDashboardData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(systemResources)
	})
	http.ListenAndServe(":8080", nil)
}
