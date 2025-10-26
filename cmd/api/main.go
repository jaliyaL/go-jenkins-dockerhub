package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaliyaL/go-jenkins-dockerhub/internal/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	fmt.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
