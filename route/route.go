package route

import (
	"database/sql"
	"net/http"

	"github.com/dickidarmawansaputra/go-crud/controller"
)

func MapRoute(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexEmployee(db))
	server.HandleFunc("/employee/create", controller.NewCreateEmployee(db))
	server.HandleFunc("/employee/update", controller.NewUpdateEmployee(db))
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployee(db))
}
