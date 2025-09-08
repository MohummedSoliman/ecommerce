package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", app.VertualTerminal)
	mux.Post("/payment-successed", app.PaymentSucceeded)

	return mux
}
