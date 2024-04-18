package controller

import (
	"bytes"
	"html/template"
	"net/http"
	"path"
)

func LoginView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles(
			path.Join("view", "layout.html"),
			path.Join("view", "login.html"),
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = templ.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		password := r.FormValue("password")
		if password != "rizalrizal" {
			tmpl, err := template.ParseFiles(
				path.Join("view", "components", "login", "error_message.html"),
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			b := new(bytes.Buffer)
			err = tmpl.Execute(b, nil)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			// should be 401 Unauthorized but make HTMX error
			w.WriteHeader(http.StatusOK)
			w.Write(b.Bytes())
			return
		}

		w.Header().Add("hx-redirect", "/dashboard")
		w.WriteHeader(http.StatusFound)
	}
}
