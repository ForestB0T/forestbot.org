package controllers

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
)

type Routes struct {
	DB *sql.DB
}

func (f *Routes) Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Testing route"}`))

	db := f.DB

	rows, err := db.Query("USE database1; SELECT * FROM users")
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
