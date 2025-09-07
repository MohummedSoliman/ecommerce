package main

import (
	"net/http"
)

func (app *application) VertualTerminal(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(w, r, "terminal", &templateData{})
	if err != nil {
		app.errorLog.Println(err)
	}
}
