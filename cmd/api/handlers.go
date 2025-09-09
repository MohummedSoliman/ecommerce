package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MohummedSoliman/ecommerce/internal/cards"
	"github.com/go-chi/chi/v5"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
}

func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	var payload stripePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	card := cards.Card{
		Secret:   app.config.stripe.secret,
		Key:      app.config.stripe.key,
		Currency: payload.Currency,
	}

	okey := true
	payIntent, msg, err := card.Charge(payload.Currency, amount)
	if err != nil {
		okey = false
	}

	if okey {
		out, err := json.MarshalIndent(payIntent, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	} else {
		res := jsonResponse{
			OK:      false,
			Message: msg,
		}

		out, err := json.MarshalIndent(res, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}
}

func (app *application) GetWidgetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	widget, err := app.DB.GetWedgit(id)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	out, err := json.Marshal(widget)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Typ", "application/json")
	w.Write(out)
}
