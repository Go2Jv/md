package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

const PORT = ":8080"

func main() {
	mux := chi.NewMux()

	fmt.Printf("the service is running on %v", PORT)
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		panic(err.Error())
	}
}
