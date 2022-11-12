package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Routes struct {
	DB *sql.DB
}

func (f *Routes) GetTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}
