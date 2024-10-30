// main.go
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	mymiddlware "github.com/lucassantoss1701/kuririn/kuririn-backend/middleware"
)

func main() {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Altere para o domínio do seu frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Adiciona o middleware CORS ao router
	r.Use(cors.Handler)

	// Middleware global, se necessário
	r.Use(middleware.Logger) // Exemplo: log de requisições

	// Rota pública
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the public page!"))
	})

	// Rota privada protegida
	r.Route("/dashboard", func(r chi.Router) {
		r.Use(mymiddlware.ValidateJWT) // Aplicando middleware de validação JWT
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to the dashboard!"))
		})
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	// Inicializa o servidor e mantém o processo em execução
	log.Println("Servidor rodando na porta 8000...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
