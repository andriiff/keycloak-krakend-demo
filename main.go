package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("handling %s %s\n", r.Method, r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"data": "42",
		})
	})

	http.HandleFunc("/private", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("handling %s %s user=%s\n", r.Method, r.URL.Path, r.Header.Get("x-user"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"secret": "42",
		})
	})

	port := os.Getenv("PORT")
	fmt.Printf("listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}
