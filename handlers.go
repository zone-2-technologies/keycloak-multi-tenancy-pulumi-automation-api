package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// CreateTenantRequest POST request payload
type CreateTenantRequest struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// CreateTenantResponse POST request response
type CreateTenantResponse struct {
	RealmId string `json:"realmId"`
}

func createTenantHandler(w http.ResponseWriter, req *http.Request) {
	var payload CreateTenantRequest
	err := json.NewDecoder(req.Body).Decode(&payload)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)

		return
	}

	// TODO: implement creation of the Keycloak realm

	response := &CreateTenantResponse{
		RealmId: "TODO",
	}

	err = json.NewEncoder(w).Encode(&response)

	if err != nil {
		http.Error(w, "Could not write response json", http.StatusInternalServerError)
	}
}

func deleteTenantHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	if len(params["slug"]) == 0 {
		http.Error(w, "slug is required", http.StatusBadRequest)

		return
	}

	// TODO: implement deletion of the Keycloak realm
}
