package main

import (
	"cinego/internal/repository"
	"cinego/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type Application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	var app Application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=admin password=Iwle@3674 dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.Domain = "exapmle.go"
	log.Println("Starting application on port ", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {

	}

}
