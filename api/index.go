package handler

import (
	"fmt"
	"net/http"
	"os"
	"tc-service-otp/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
   setupRouter().ServeHTTP(w, r)
}

func setupRouter() http.Handler {
   r := chi.NewRouter()

   r.Use(cors.Handler(cors.Options{
       AllowedOrigins:   []string{"*"},
       AllowedHeaders:   []string{"*"},
       ExposedHeaders:   []string{"*"},
       AllowCredentials: false,
       MaxAge:           300,
       AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
   }))

   r.Use(setCORSHeaders)
   r.Use(middleware.Recoverer)
   r.Get("/health", apiToGetHealth)
   r.Get("/port", apiToGetPort)
   r.Route("/user", func(r chi.Router) {
       r.Post("/", controller.CreateUser)
       r.Put("/{userId}", controller.UpdateUser)
       r.Delete("/{userId}", controller.DeleteUser)
       r.Get("/{userId}", controller.GetUserById)
       r.Get("/username/{username}", controller.GetUserByUsername)
   })
   r.Get("/users", controller.GetAllUsers)

   return r
}

func setCORSHeaders(next http.Handler) http.Handler {
   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       if r.Method == "OPTIONS" {
           w.Header().Set("Access-Control-Allow-Origin", "*")
           w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
           w.Header().Set("Access-Control-Allow-Headers", "*")
           w.Header().Set("Access-Control-Max-Age", "300")
           w.Header().Set("Access-Control-Expose-Headers", "*")
           w.Header().Set("Access-Control-Allow-Credentials", "false")
           w.WriteHeader(http.StatusOK)
           return
       }
       next.ServeHTTP(w, r)
   })
}

func apiToGetHealth(w http.ResponseWriter, _ *http.Request) {
   w.WriteHeader(http.StatusOK)
   w.Write([]byte("OK"))
}

func apiToGetPort(w http.ResponseWriter, _ *http.Request) {
   port := os.Getenv("GO_PORT")
   addr := fmt.Sprintf("Port:%s", port)
   w.WriteHeader(http.StatusOK)
   w.Write([]byte(addr))
}
