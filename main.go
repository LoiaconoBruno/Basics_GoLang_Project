package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not found in the enviroment")
	} else {
		fmt.Println("Port: ", portString)
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// 🔹 Orígenes permitidos (podes ajustar según tu frontend)
		AllowedOrigins: []string{"https://example.com", "http://localhost:3000"},

		// 🔹 Métodos que acepta tu API
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},

		// 🔹 Headers que acepta
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},

		// 🔹 Headers expuestos al cliente (ej: tokens en headers)
		ExposedHeaders: []string{"Link"},

		// 🔹 Permitir cookies/autenticación (true si usas sesiones o JWT en cookies)
		AllowCredentials: true,

		// 🔹 Cache del preflight (OPTIONS)
		MaxAge: 300, // 5 minutos
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port: ", portString)
}
