package main

import (
		"github.com/go-chi/chi"
		"github.com/go-chi/cors"
		"log"
		"net/http"
		"time"
)

func main() {
		r := chi.NewRouter()
		// Basic CORS
		// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
		cors := cors.New(cors.Options{
				// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
				AllowedOrigins:   []string{"*"},
				// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: true,
				MaxAge:           300, // Maximum value not ignored by any of major browsers
		})
		r.Use(cors.Handler)
		r.Route("/api", func(r chi.Router) {
				r.Get("/msg", func(w http.ResponseWriter, req *http.Request) {
						if _, err := w.Write([]byte("Hello world")); err != nil {
								log.Fatal(err)
						}
				})
		})

		srv := &http.Server{
				Handler:      r,
				Addr:         ":5000",
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
		}

		log.Fatal(srv.ListenAndServe())
}
