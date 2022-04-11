package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/febzey/forestbot-api/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Routes struct {
	DB *sql.DB
}

//
// Playtime endpoint handler
//
func (f *Routes) GetPlaytime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.DB

	username := params["user"]
	mc_server := params["server"]

	rows, err := db.Query("SELECT playtime FROM users WHERE username = ? AND mc_server = ?", username, mc_server)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		playtime int
	)
	for rows.Next() {
		err = rows.Scan(&playtime)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	w.Write([]byte(`{"playtime": ` + strconv.Itoa(playtime) + `}`))

	return
}

//
// kills / deaths endpoint handler
//
func (f *Routes) GetKD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.DB

	username := params["user"]
	mc_server := params["server"]

	rows, err := db.Query("SELECT kills, deaths FROM users WHERE username = ? AND mc_server = ?", username, mc_server)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		kills  int
		deaths int
	)
	for rows.Next() {
		err = rows.Scan(&kills, &deaths)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	w.Write([]byte(`{"kills": ` + strconv.Itoa(kills) + `, "deaths": ` + strconv.Itoa(deaths) + `}`))

	return
}

//
// Join count endpoint handler
//
func (f *Routes) GetJoins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.DB

	username := params["user"]
	mc_server := params["server"]

	rows, err := db.Query("SELECT joins FROM users WHERE username = ? AND mc_server = ?", username, mc_server)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		joins int
	)
	for rows.Next() {
		err = rows.Scan(&joins)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	w.Write([]byte(`{"joins": ` + strconv.Itoa(joins) + `}`))

	return
}

//
// Lastseen endpoint handler
//
func (f *Routes) GetLastSeen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.DB

	username := params["user"]
	mc_server := params["server"]

	rows, err := db.Query("SELECT lastseen FROM users WHERE username = ? AND mc_server = ?", username, mc_server)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		lastseen string
	)
	for rows.Next() {
		err = rows.Scan(&lastseen)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	w.Write([]byte(`{"lastseen": "` + lastseen + `"}`))

	return
}

//
// Joindate endpoint handler
//
func (f *Routes) GetJoinDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.DB

	username := params["user"]
	mc_server := params["server"]

	rows, err := db.Query("SELECT joindate FROM users WHERE username = ? AND mc_server = ?", username, mc_server)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		joindate string
	)
	for rows.Next() {
		err = rows.Scan(&joindate)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	w.Write([]byte(`{"joindate": "` + joindate + `"}`))

	return
}

//
// Get all user stats endpoint
//
func (f *Routes) GetAllUserStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.DB

	user := params["user"]
	mc := params["server"]

	rows, err := db.Query("SELECT * FROM users WHERE username = ? AND mc_server = ?", user, mc)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		username        string
		playtime        int
		kills           int
		deaths          int
		joins           int
		leaves          int
		lastseen        string
		joindate        string
		uuid            string
		lastdeathString string
		lastdeathTime   int
		mc_server       string
	)
	for rows.Next() {
		err = rows.Scan(
			&username,
			&kills,
			&deaths,
			&joindate,
			&lastseen,
			&uuid,
			&playtime,
			&joins,
			&leaves,
			&lastdeathTime,
			&lastdeathString,
			&mc_server,
		)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	w.Write([]byte(`{"username": "` + username + `", "kills": ` + strconv.Itoa(kills) + `, "deaths": ` + strconv.Itoa(deaths) + `, "joins": ` + strconv.Itoa(joins) + `, "leaves": ` + strconv.Itoa(leaves) + `, "lastseen": "` + lastseen + `", "joindate": "` + joindate + `", "uuid": "` + uuid + `", "playtime": ` + strconv.Itoa(playtime) + `, "lastdeath": "` + lastdeathString + `", "mc_server": "` + mc_server + `"}`))

	return
}

//
// Total message count endpoint handler
//
func (f *Routes) GetMessageCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.DB

	user := params["user"]
	mc := params["server"]

	rows, err := db.Query("SELECT name,COUNT(name) AS cnt FROM messages WHERE name = ? AND mc_server = ?", user, mc)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		e   string
		cnt int
	)
	for rows.Next() {
		err = rows.Scan(&e, &cnt)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	w.Write([]byte(`{"messagecount": ` + strconv.Itoa(cnt) + `}`))

	return
}

//
// Random quote endpoint handler
//
func (f *Routes) GetQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.DB

	user := params["user"]
	mc := params["server"]

	rows, err := db.Query("SELECT name,message,date FROM messages WHERE name=? AND mc_server = ? AND LENGTH(message) > 30 ORDER BY RAND() LIMIT 1", user, mc)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		name    string
		message string
		date    string
	)
	for rows.Next() {
		err = rows.Scan(&name, &message, &date)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	w.Write([]byte(`{"name": "` + name + `", "message": "` + message + `", "date": "` + date + `"}`))

	return
}

//
// lastdeath endpoint handler
//
func (f *Routes) GetLastDeath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.DB

	user := params["user"]
	mc := params["server"]

	rows, err := db.Query("SELECT lastdeathTime, lastdeathString from users WHERE username = ? AND mc_server = ?", user, mc)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		lastdeathTime   int
		lastdeathString string
	)
	for rows.Next() {
		err = rows.Scan(&lastdeathTime, &lastdeathString)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	w.Write([]byte(`{"death": "` + lastdeathString + `", "time": ` + strconv.Itoa(lastdeathTime) + `}`))

	return
}

//create a struct that models a object with a playerListExtra that is an array of objects

type webSocketContents struct {
	PlayerList      []string     `json:"playerList"`
	PlayerListExtra []utils.User `json:"playerListExtra"`
	UniquePlayers   int          `json:"uniquePlayers"`
}

//
// Tablist endpoint handler
//
func (f *Routes) GetTablist(w http.ResponseWriter, r *http.Request) {

	var contents webSocketContents

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/playerlist"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		err = json.Unmarshal(message, &contents)
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.GenerateTablist(contents.PlayerListExtra, w)
		conn.Close()
		return
	}

}
