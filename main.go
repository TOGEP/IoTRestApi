package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/temperature", http.HandlerFunc(temperature))
	log.Println(http.ListenAndServe(":8080", mux))
}

func temperature(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "24.00")
}
