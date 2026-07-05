package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok"}`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	fmt.Println("Listening on 127.0.0.1:9000")

	err := http.ListenAndServe("127.0.0.1:9000", mux)
	if err != nil {
		panic(err)
	}
}
