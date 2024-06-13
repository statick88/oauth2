// main.go

package main

import (
    "fmt"
    "log"

    "github.com/statick88/oauth2/db"
    "github.com/statick88/oauth2/handlers"
)

func main() {
    // Initialize database connection
    _, err := db.InitDB()
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }

    // Example usage
    email := "test@example.com"
    password := "password123"

    // Create a new user
    err = handlers.CreateUser(email, password)
    if err != nil {
        log.Fatalf("Error creating user: %v", err)
    }

    // Get user by ID
    user, err := handlers.GetUserByID(1)
    if err != nil {
        log.Fatalf("Error getting user: %v", err)
    }
    fmt.Printf("User found: %+v\n", user)

    // Delete user by ID
    err = handlers.DeleteUserByID(1)
    if err != nil {
        log.Fatalf("Error deleting user: %v", err)
    }
}
