package main

import (
	"log"
	"net/http"
	"os"
	"github.com/izzettinozbektas/userControlSystem/auth-service/internal/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := routes.SetupRouter()

	log.Println("Auth service is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
