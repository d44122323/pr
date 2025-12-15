package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := LoadConfig()
	mux := SetupRouter(cfg)

	fmt.Println("SERVER RUNNING: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
