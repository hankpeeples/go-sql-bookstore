package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hankpeeples/sql-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3080", r))
}
