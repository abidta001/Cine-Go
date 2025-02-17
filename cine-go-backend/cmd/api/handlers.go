package main

import (
	"fmt"
	"net/http"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there from %s", app.Domain)
}
