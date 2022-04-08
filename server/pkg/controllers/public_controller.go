package controllers

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/febzey/forestbot-api/pkg/database"
	"github.com/gorilla/mux"
)

// func Root(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	//send a response that says hello world
// 	w.Write([]byte(`{"message": "Root route"}`))
// }

type Routes struct {
	DB *sql.DB
}

func (f *Routes) Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Testing route"}`))

	db := f.DB

	//run the query

	//TODO! Fix query input error. sanatize;
	rows, err := database.RunQuery("SELECT * FROM users", db)
	if err != nil {
		fmt.Println("Error running query:", err)
	}

	//print rows

	fmt.Printf("%+v", rows)

}

func (f *Routes) GetPlaytime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)

	w.Write([]byte(`{"message": "Get playtime route", "id": "` + id + `"}`))

}
