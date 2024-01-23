package main

import (
	"net/http"

	"github.com/dickidarmawansaputra/go-crud/database"
	"github.com/dickidarmawansaputra/go-crud/route"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	route.MapRoute(server, db)

	http.ListenAndServe(":8080", server)
}
