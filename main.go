package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	Version  = "1.0.0"
	Priority = 1

	jwsIndicator = "jws"

	db  *sql.DB
	err error
)

const (
	driver   = "postgres"
	host     = "postgres"
	port     = 5432
	user     = "admin"
	password = "secret"
	dbname   = "api_log"
)

func conection() {
	log.Println("Accessing %s ... ", dbname)

	var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open(driver, DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Connected!")
	}

	defer db.Close()
}

func main() {
	conection()
}
