package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "forest:treehouse@tcp(170.199.19.152:3306)/discord")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database")

	return db, nil
}

func EndConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection to database closed")
}

func RunQuery(query string, db *sql.DB, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
