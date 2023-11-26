package middleware

import (
	"log"
	"net/http"
)

type Logging struct {
	Handler http.Handler
}

func (logging *Logging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v \n",r.Method,r.URL.Path)
	logging.Handler.ServeHTTP(w,r)
}