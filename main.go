package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
)

type Item struct {
	Topic string `json:"topic"`
	Data  int    `json:"data"`
}

type datastore struct {
	sync.RWMutex
	m map[string]Item
}

func main() {
	//TODO db init
	mux := http.NewServeMux()
	store := &datastore{
		m:       map[string]Item{},
		RWMutex: sync.RWMutex{},
	}

	mux.Handle("/temperature", store)
	mux.Handle("/temperature/", store)
	log.Println(http.ListenAndServe(":8080", mux))
}

func (store *datastore) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		store.getTemperature(w, r)
	case http.MethodPost:
		store.addTemperature(w, r)
	default:
		// 現状その他は想定していない
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (store *datastore) getTemperature(w http.ResponseWriter, r *http.Request) {
	topic := strings.TrimPrefix(r.URL.Path, "/temperature/")
	store.RLock()
	item, ok := store.m[topic]
	store.RUnlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	jsonBytes, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (store *datastore) addTemperature(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store.Lock()
	store.m[item.Topic] = item
	store.Unlock()

	jsonBytes, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
