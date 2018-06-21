package controllers

import (
	"net/http"
)

func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionId         string
		Success               bool
		Response              bool
		TransactionIdProvider string
		SuccessProvider       bool
		ResponseProvider      bool
	}{
		TransactionId:         "",
		Success:               false,
		Response:              false,
		TransactionIdProvider: "",
		SuccessProvider:       false,
		ResponseProvider:      false,
	}
	if r.FormValue("submitted") == "true" {
		helloValue := r.FormValue("hello")
		helloProviderValue := r.FormValue("helloprovider")
		txid, err := app.Fabric.InvokeHello(app.CcStorage, "hello", helloValue)
		if err != nil {
			http.Error(w, "Unable to invoke hello in the blockchain", 500)
		}
		txidprovider, err2 := app.Fabric.InvokeHello(app.CcProvider, "helloprovider", helloProviderValue)
		if err2 != nil {
			http.Error(w, "Unable to invoke provider hello in the blockchain", 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
		data.TransactionId = txidprovider
		data.SuccessProvider = true
		data.ResponseProvider = true
	}
	renderTemplate(w, r, "request.html", data)
}
