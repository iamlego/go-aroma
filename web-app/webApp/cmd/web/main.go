package main

import (
	"fmt"
	handler "github/iamlego/go-web/pkg/handlers"
	"log"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/About", handler.About)

	log.Println(fmt.Sprintf("The server is running on Port %s", port))
	_ = http.ListenAndServe(port, nil)
}