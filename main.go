package main

import (
	"database/sql"
	"github.com/gtrodrigues04/desafio-go/api"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(db:3306)/db_routes")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()

	routesHandler := api.RoutesAPI(db)
	r.Mount("/api", routesHandler)

	http.ListenAndServe(":8081", r)

}
