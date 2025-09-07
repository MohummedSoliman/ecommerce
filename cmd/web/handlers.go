package main

import "net/http"

func (app *application) VertualTerminal(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Hit The Handler")
}
