// handlers/user.go

package handlers

import (
    "database/sql"
    "fmt"

    "github.com/username/oauth2/db"
)

type User struct {
    ID       int
    Email    string
    Password string
}

// Create a new user
func CreateUser(email, password string) error {
    _, err := db.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, password)
    if err != nil {
        fmt.Println("Error creating user:", err)
        return err
    }

    fmt.Println("User created successfully")
    return nil
}

// Get a user by ID
func GetUserByID(id int) (*User, error) {
    var user User
    err := db.DB.QueryRow("SELECT id, email, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email, &user.Password)
    switch {
    case err == sql.ErrNoRows:
        fmt.Println("No user with that ID.")
        return nil, nil
    case err != nil:
        fmt.Println("Error querying user:", err)
        return nil, err
    }

    return &user, nil
}

// Delete a user by ID
func DeleteUserByID(id int) error {
    _, err := db.DB.Exec("DELETE FROM users WHERE id = ?", id)
    if err != nil {
        fmt.Println("Error deleting user:", err)
        return err
    }

    fmt.Println("User deleted successfully")
    return nil
}
