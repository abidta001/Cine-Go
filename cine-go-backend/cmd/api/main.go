package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type Application struct {
	DSN    string
	Domain string
}

func main() {
	var app Application

	flag.StringVar(&app.DSN, "dsn", "hsot=localhost port=5432 user=admin password=Iwle@3674 dbname=movies sslmode=disable timezone=UTC conenct_timeout=5", "Postgres connection string")
	flag.Parse()

	app.Domain = "exapmle.go"
	log.Println("Starting application on port ", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {

	}

}
