package routes

import (
	"net/http"
	"ttd/controller"

	"github.com/gorilla/mux"
)

func Routes () *mux.Router {
	route := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	//Static File 
	route.PathPrefix("/static").Handler(http.StripPrefix("/static",fs))


	route.HandleFunc("/students/{category}",controller.Home).Methods("GET")
	route.HandleFunc("/add",controller.AddStudent).Methods("POST")
	route.HandleFunc("/search",controller.Search).Methods("GET")
	route.HandleFunc("/delete/{NIS}",controller.DeleteStudent).Methods("POST")
	route.HandleFunc("/update/{NIS}",controller.UpdateStudent).Methods("POST")
	return route
}