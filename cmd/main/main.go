package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaipalsinghkhangarot/bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
