package routes

import (
	"github.com/gorilla/mux"
	"github.com/scriptnsam/web-server/pkg/controllers"
)

var RegisterFoodieRoutes = func(router *mux.Router){
	// use controllers to handle functions
	router.HandleFunc("/",controllers.GetFoods).Methods("GET")
	router.HandleFunc("/{id}",controllers.GetFoodsById).Methods("GET")
	router.HandleFunc("/",controllers.CreateFood).Methods("POST")
	router.HandleFunc("/{id}",controllers.UpdateFood).Methods("PUT")
	router.HandleFunc("/{id}",controllers.DeleteFood).Methods("DELETE")
}