package main

import (
	"fmt"
	"log"
	"net/http"
)

const HOST = "0.0.0.0"
const PORT = 5050

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /jobs/list", listAPI)
	mux.HandleFunc("POST /jobs/insert", insertAPI)
	mux.HandleFunc("POST /jobs/update", updateAPI)

	log.Printf("Starting server on %s:%d\n", HOST, PORT)

	hostAndPort := fmt.Sprintf("%s:%d", HOST, PORT)
	err := http.ListenAndServe(hostAndPort, mux)
	log.Fatal(err)
}
