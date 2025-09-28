package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"coursegolang/internal/database"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// 🔹 Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, usando variables del sistema")
	}

	// 🔹 URL de la base de datos
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL no encontrada en el entorno")
	}

	// 🔹 Puerto donde va a escuchar el servidor HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // puerto por defecto
	}

	// 🔹 Conexión a la DB
	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	// 🔹 Verificar conexión activa
	if err := connection.Ping(); err != nil {
		log.Fatal("No se pudo hacer ping a la base de datos:", err)
	}

	// 🔹 Configuración del API con sqlc
	apiCfg := apiConfig{
		DB: database.New(connection),
	}

	// 🔹 Router principal
	router := chi.NewRouter()

	// 🔹 Middleware CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://example.com", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // 5 minutos
	}))

	// 🔹 Subrouter v1
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIKey))
	router.Mount("/v1", v1Router)

	// 🔹 Configuración del servidor HTTP
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port, // 🔹 Cambio: usar puerto, no DB_URL
	}

	log.Printf("Server starting on port %v", port)

	// 🔹 Iniciar servidor
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", port)
}
