package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewRandHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("server is listening on 8081")
	server.ListenAndServe()
}
