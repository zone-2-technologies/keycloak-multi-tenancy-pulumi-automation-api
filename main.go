package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tenants", createTenantHandler).Methods("POST")
	router.HandleFunc("/tenants/{slug}", deleteTenantHandler).Methods("DELETE")

	s := &http.Server{Addr: ":8000", Handler: router}

	log.Println("Listening on :8000")

	err := s.ListenAndServe()

	if err != nil {
		log.Fatal("Unable to start server", err)
	}

}
