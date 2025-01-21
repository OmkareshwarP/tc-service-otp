package main

import (
	"log"
	"net/http"
	api "tc-service-otp/api"
)

func main() {
   port := "8080"
   log.Printf("Server running on http://localhost:%s", port)
   log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(api.Handler)))
}
