package controller

import (
	"html/template"
	"net/http"
	"path"
)

func DashboardView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(path.Join("view", "layout.html"), path.Join("view", "dashboard.html"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = tmpl.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
