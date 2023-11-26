package web

import (
	"fmt"
	"net/http"
	"ttd/config"
	"ttd/middleware"
	"ttd/routes"
)


func App () {
	logging := &middleware.Logging{
		Handler: routes.Routes(),
	}
	server := http.Server{
		Addr: config.Address,
		Handler: logging,
	}
	fmt.Printf("Server Running http://%v 🚀🚀🚀\n",config.Address)
	server.ListenAndServe()
}
