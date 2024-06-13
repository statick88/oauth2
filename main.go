package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"               // Importa el enrutador mux de Gorilla
	"github.com/statick88/oauth2/db"       // Importa el paquete db desde oauth2
	"github.com/statick88/oauth2/handlers" // Importa los manejadores desde el paquete oauth2/handlers
)

func main() {
	// Initialize database connection
	_, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Create a new router
	r := NewRouter()

	// Start the server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/auth", handlers.Auth).Methods("GET")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/logout", handlers.Logout).Methods("GET")
	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := handlers.GetUserByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}).Methods("GET")

	return r
}