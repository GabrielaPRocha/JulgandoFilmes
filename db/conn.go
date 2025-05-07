package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "102030"
	dbname   = "filmes_criticas"
	apiKey   = "371c84afe53f4ebb7026156c896b1732"
)

func ConnectDB() (*sql.DB, error) {
	/*	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)*/
	psqlinfo := "root:102030@tcp(127.0.0.1:3306)/filmes_criticas"
	db, err := sql.Open("mysql", psqlinfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to " + dbname)
	return db, nil
}
