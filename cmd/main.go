package main

import (
	"fmt"
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
	fmt.Print("Server running on port :3080\n")
	err := http.ListenAndServe("localhost:3080", r)
	if err != nil {
		log.Fatal("Server crashed: \n", err.Error())
	}
}
