package main

import (
	"fmt"
	"log"
	"net/http"
	"vitabu/components/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterSiteRoutes(router)
	http.Handle("/", router)
	fmt.Println(">>>> Starting server for vitabu app at port 8000: ")
	log.Fatal(http.ListenAndServe(":8000", router))
}
