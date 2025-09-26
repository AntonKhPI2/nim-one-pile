package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/AntonKhPI2/nim-one-pile/internal/database"
	"github.com/AntonKhPI2/nim-one-pile/internal/handlers"
	"github.com/AntonKhPI2/nim-one-pile/internal/repositories"
	"github.com/AntonKhPI2/nim-one-pile/internal/services"
)

func main() {
	_ = godotenv.Load(".env")

	ctx := context.Background()

	gormDB := database.MustOpenMySQLGorm(ctx)
	repo := repositories.NewGormGameRepo(gormDB)

	gameSvc := services.NewGameService(repo)

	mux := http.NewServeMux()
	handlers.RegisterGameRoutes(mux, gameSvc)

	addr := ":" + getenv("PORT", "8080")
	cors := parseCSV(getenv("CORS_ORIGINS", "http://localhost:5173,*"))

	srv := &http.Server{
		Addr:         addr,
		Handler:      withCORS(mux, cors),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Backend (GORM+MySQL) listening on", addr)
	log.Fatal(srv.ListenAndServe())
}

func withCORS(next http.Handler, origins []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if allowOrigin(origin, origins) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else if contains(origins, "*") {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
func parseCSV(s string) []string {
	if s == "" {
		return nil
	}
	a := strings.Split(s, ",")
	for i := range a {
		a[i] = strings.TrimSpace(a[i])
	}
	return a
}
func allowOrigin(o string, list []string) bool {
	if o == "" {
		return false
	}
	for _, v := range list {
		if v == o {
			return true
		}
	}
	return false
}
func contains(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}
