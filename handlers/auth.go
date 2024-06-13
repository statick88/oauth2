package handlers

import (
    "net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Auth handler"))
}