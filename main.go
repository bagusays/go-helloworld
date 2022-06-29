package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	})
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
