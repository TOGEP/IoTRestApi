package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Item struct {
	Topic string `json:"topic"`
	Data  int    `json:"data"`
}

func main() {
	//TODO db init
	mux := http.NewServeMux()
	mux.Handle("/temperature", http.HandlerFunc(temperature))
	log.Println(http.ListenAndServe(":8080", mux))
}

func temperature(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTemperature(w, r)
	case http.MethodPost:
		addTemperature(w, r)
	default:
		// 現状その他は想定していない
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func getTemperature(w http.ResponseWriter, r *http.Request) {
	item := Item{http.MethodGet, http.StatusOK}
	res, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func addTemperature(w http.ResponseWriter, r *http.Request) {
	item := Item{http.MethodPost, http.StatusOK}
	res, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
