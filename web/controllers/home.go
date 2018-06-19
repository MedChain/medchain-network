package controllers

import (
	"net/http"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryHello(app.CcStorage, "hello")
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
	helloProviderValue, err2 := app.Fabric.QueryHello(app.CcProvider, "helloprovider")
	if err2 != nil {
		http.Error(w, "Unable to query the blockchain for provider info", 500)
	}

	data := &struct {
		Hello         string
		HelloProvider string
	}{
		Hello:         helloValue,
		HelloProvider: helloProviderValue,
	}
	renderTemplate(w, r, "home.html", data)
}
