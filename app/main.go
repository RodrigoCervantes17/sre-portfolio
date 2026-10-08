package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
func readyz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ready")
}

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/readyz", readyz)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}
	fmt.Println("listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
