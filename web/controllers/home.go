package controllers

import (
	"net/http"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryHello()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
	helloProviderValue, err2 := app.Fabric.QueryHelloProvider()
	if err2 != nil {
		http.Error(w, "Unable to query the blockchain for provider info", 500)
	}
	patientData, err3 := app.Fabric.QueryPatientData("123456")
	if err3 != nil {
		http.Error(w, "Unable to query the blockchain for patientData", 500)
	}

	patientRecords, err4 := app.Fabric.QueryPatientRecords([]string{"1", "2", "3"})
	if err4 != nil {
		http.Error(w, "Unable to query the blockchain for patientRecords", 500)
	}

	data := &struct {
		Hello          string
		HelloProvider  string
		PatientData    string
		PatientRecords string
	}{
		Hello:          helloValue,
		HelloProvider:  helloProviderValue,
		PatientData:    patientData,
		PatientRecords: patientRecords,
	}
	renderTemplate(w, r, "home.html", data)
}
