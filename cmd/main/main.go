package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scriptnsam/web-server/pkg/routes"
)

// The main function sets up a router, registers routes for a foodie application, and starts a server
// listening on localhost:4000.
func main(){
	r:=mux.NewRouter()
	routes.RegisterFoodieRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:4000", r))
}