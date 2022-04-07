package controllers

import (
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//send a response that says hello world
	w.Write([]byte(`{"message": "Root route"}`))
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//send a response that says hello world
	w.Write([]byte(`{"message": "Testing route"}`))
}
