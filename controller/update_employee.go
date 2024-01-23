package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewUpdateEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form["name"][0]
			npwp := r.Form["npwp"][0]
			address := r.Form["address"][0]

			_, err := db.Exec("UPDATE employee SET name = ?, npwp = ?, address = ? WHERE id = ?", name, npwp, address, id)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
			return
		}

		row := db.QueryRow("SELECT * FROM employee WHERE id = ?", id)
		if row.Err() != nil {
			w.Write([]byte(row.Err().Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var employee Employee
		err := row.Scan(
			&employee.Id,
			&employee.Name,
			&employee.NPWP,
			&employee.Address,
		)
		employee.Id = id
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fp := filepath.Join("views", "update.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["employee"] = employee

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
