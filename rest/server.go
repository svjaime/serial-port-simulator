package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"ubiwhere.com/serial-port-simulator/config"
	"ubiwhere.com/serial-port-simulator/inverter"
)

//getActiveConfig function will handle GET requests to /config
func getActiveConfig(w http.ResponseWriter, r *http.Request) {

	fmt.Println("++REST++ Endpoint Hit: getActiveConfig")
	w.Header().Set("Content-Type", "application/json")

	jsonBody, err := json.Marshal(config.Active)
	if err != nil {
		http.Error(w, "Error converting results to json", http.StatusInternalServerError)
	}

	w.Write(jsonBody)
}

//getInverter function will handle GET requests to /inverter
func getInverter(w http.ResponseWriter, r *http.Request) {

	fmt.Println("++REST++ Endpoint Hit: getInverter")
	w.Header().Set("Content-Type", "application/json")

	jsonBody, err := json.Marshal(inverter.Active)
	if err != nil {
		http.Error(w, "Error converting results to json", http.StatusInternalServerError)
	}

	w.Write(jsonBody)
}

//getInverterInfo function will handle GET requests to /inverter/info
func getInverterInfo(w http.ResponseWriter, r *http.Request) {

	fmt.Println("++REST++ Endpoint Hit: getInverterInfo")
	w.Header().Set("Content-Type", "application/json")

	jsonBody, err := json.Marshal(inverter.Active.Info)
	if err != nil {
		http.Error(w, "Error converting results to json", http.StatusInternalServerError)
	}

	w.Write(jsonBody)
}

//getBattery function will handle GET requests to /inverter/battery
func getBattery(w http.ResponseWriter, r *http.Request) {

	fmt.Println("++REST++ Endpoint Hit: getBattery")
	w.Header().Set("Content-Type", "application/json")

	jsonBody, err := json.Marshal(inverter.Active.Bms)
	if err != nil {
		http.Error(w, "Error converting results to json", http.StatusInternalServerError)
	}

	w.Write(jsonBody)
}

//getBatteryInfo function will handle GET requests to /inverter/battery/info
func getBatteryInfo(w http.ResponseWriter, r *http.Request) {

	fmt.Println("++REST++ Endpoint Hit: getBatteryInfo")
	w.Header().Set("Content-Type", "application/json")

	jsonBody, err := json.Marshal(inverter.Active.Bms.Info)
	if err != nil {
		http.Error(w, "Error converting results to json", http.StatusInternalServerError)
	}

	w.Write(jsonBody)
}

//getBatteryMetrics function will handle GET requests to /inverter/battery/metrics
func getBatteryMetrics(w http.ResponseWriter, r *http.Request) {

	fmt.Println("++REST++ Endpoint Hit: getBatteryMetrics")
	w.Header().Set("Content-Type", "application/json")

	jsonBody, err := json.Marshal(inverter.Active.Bms.Data)
	if err != nil {
		http.Error(w, "Error converting results to json", http.StatusInternalServerError)
	}

	w.Write(jsonBody)
}

//updateInverter function will handle POST requests to /inverter
func updateInverter(w http.ResponseWriter, r *http.Request) {

	fmt.Println("++REST++ Endpoint Hit: updateInverterInfo")
	w.Header().Set("Content-Type", "application/json")

	var inv inverter.Inverter

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &inv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = inv.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	inverter.Active = inv
	fmt.Printf("%v\n", inverter.Active)

	jsonBody, _ := json.Marshal(inverter.Active)

	err = ioutil.WriteFile(config.DataFile.Name(), jsonBody, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	config.DataFile.Sync()

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBody)
}

//HandleRequests function will match the URL path hit with a defined function
func HandleRequests() {
	fmt.Println("++REST++ api running on localhost:10000")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/config", getActiveConfig).Methods(http.MethodGet)

	router.HandleFunc("/inverter", getInverter).Methods(http.MethodGet)
	router.HandleFunc("/inverter", updateInverter).Methods(http.MethodPost)

	router.HandleFunc("/inverter/info", getInverterInfo).Methods(http.MethodGet)

	router.HandleFunc("/inverter/battery", getBattery).Methods(http.MethodGet)
	router.HandleFunc("/inverter/battery/info", getBatteryInfo).Methods(http.MethodGet)
	router.HandleFunc("/inverter/battery/metrics", getBatteryMetrics).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":10000", router))
}
