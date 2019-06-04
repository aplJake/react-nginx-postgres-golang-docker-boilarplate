package main

import (
		"encoding/json"
		"fmt"
		"github.com/aplJake/react-nginx-docker-test-1/db"
		"github.com/go-chi/chi"
		"github.com/go-chi/cors"
		"github.com/joho/godotenv"
		"github.com/kelseyhightower/envconfig"
		"log"
		"net/http"
		"time"
)

type Config struct {
		PostgresDB       string `envconfig:"POSTGRES_DB"`
		PostgresUser     string `envconfig:"POSTGRES_USER"`
		PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
		PostgresHost     string `envconfig:"POSTGRES_HOST"`
		PostgresPort     string `envconfig:"POSTGRES_PORT"`
}

func ResponseOk(w http.ResponseWriter, body interface{}) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(body)
}

func main() {
		fmt.Println("Prints")
		if err := godotenv.Load("postgresql.env"); err != nil {
				log.Print("No .env file found")
		}


		//cfg := &Config{
		//		PostgresDB: os.Getenv("POSTGRES_DB"),
		//		PostgresUser: os.Getenv("POSTGRES_USER"),
		//		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		//		PostgresHost: os.Getenv("POSTGRES_HOST"),
		//		PostgresPort: os.Getenv("POSTGRES_PORT"),
		//}
		//fmt.Println("Host ", os.Getenv("POSTGRES_HOST"))

		var cfg Config
		err := envconfig.Process("", &cfg)
		if err != nil {
				log.Fatal(err.Error())
		}

		fmt.Println("User", cfg.PostgresUser)
		//err := envconfig.Process("", &cfg)
		//if err != nil {
		//		log.Fatal(err)
		//}

		conn := fmt.Sprintf("postgres://%s:%s@postgres_container:5432/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
		fmt.Println("DB connection name", conn)

		r := chi.NewRouter()
		// Basic CORS
		// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
		cors := cors.New(cors.Options{
				// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
				AllowedOrigins: []string{"*"},
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
						// connect to db
						//conn := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", "postgres", "jake23apple", "mydb1")
						//psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
						//		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
						//posgresDB, err := db.NewPostgresDB("postgresql://user:password@localhost:5432/?sslmode=disable")

						posgresDB, err := db.NewPostgresDB(conn)
						if err != nil {
								panic(err.Error())
						}

						var message string
						res := posgresDB.GetDB().QueryRow(`SELECT message FROM messages WHERE id=$1;`, 1)


						err = res.Scan(&message)
						if err != nil {
								panic(err.Error())
						}

						defer posgresDB.Close()

						ResponseOk(w, message)

						//if _, err := w.Write([]byte(message)); err != nil {
						//		panic(err)
						//}

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
