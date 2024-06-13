package db

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", "oauth2", "oauth2", "oauth2")
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }

    // Check if the connection is established
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Connected to MySQL database!")

    DB = db
    return DB, nil
}
