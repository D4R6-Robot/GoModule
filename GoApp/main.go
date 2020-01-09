package main

import (
	"malwaresystem/httplib"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/list", httplib.ListHandler)

	http.ListenAndServe(":8001", r)
}
