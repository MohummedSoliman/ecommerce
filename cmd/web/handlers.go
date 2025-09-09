package main

import (
	"net/http"

	"github.com/MohummedSoliman/ecommerce/internal/models"
)

func (app *application) VertualTerminal(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(w, r, "terminal", &templateData{}, "stripe-js")
	if err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	cardHolder := r.Form.Get("cardholder_name")
	email := r.Form.Get("email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	data := make(map[string]any)
	data["cardholder"] = cardHolder
	data["email"] = email
	data["payment_intent"] = paymentIntent
	data["payment_method"] = paymentMethod
	data["payment_amount"] = paymentAmount
	data["currency"] = paymentCurrency

	if err = app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) ChargeOne(w http.ResponseWriter, r *http.Request) {
	widget := models.Widget{
		ID:             1,
		Name:           "Custom Widget",
		Description:    "Very Nice Widget",
		InventroyLevel: 10,
		Price:          10.00,
	}

	data := make(map[string]any)
	data["widget"] = widget

	err := app.renderTemplate(w, r, "buy-once", &templateData{
		Data: data,
	}, "stripe-js")
	if err != nil {
		app.errorLog.Println(err)
		return
	}
}
