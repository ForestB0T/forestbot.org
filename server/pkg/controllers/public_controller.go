package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	DB *sql.DB
}

func (f *Routes) Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.DB

	username := params["user"]
	datab := params["db"]

	_, err := db.Exec("use ?", datab)
	if err != nil {
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		fmt.Println(err)
		return
	}

	rows, err := db.Query("SELECT playtime FROM users WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer fmt.Println("Hello, defered.")
	defer rows.Close()

	var playtime string
	for rows.Next() {
		err = rows.Scan(&playtime)
		if err != nil {
			fmt.Println(err)
		}
	}

	w.Write([]byte(`{"playtime": "` + playtime + `"}`))
}

func (f *Routes) GetPlaytime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)

	w.Write([]byte(`{"message": "Get playtime route", "id": "` + id + `"}`))

}
