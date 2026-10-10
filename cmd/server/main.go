package main

import (
	"fmt"
	"log"
	"net/http"

	"go-auth/internal/handler"
	"go-auth/internal/repository"
	"go-auth/internal/service"
)

func main() {
	// Inisialisasi Repository, Service, dan Handler
	userRepo := repository.NewUserRepository()
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// set routing memakai bawaan go(serverMux)
	mux := http.NewServeMux()

	// Daftarkan EndPoint (Wajib menggunakan HandleFunc)
	mux.HandleFunc("/api/register", authHandler.Register)
	mux.HandleFunc("/api/login", authHandler.Login)

	port := ":8080"
	fmt.Println("Server starting on", port)
	
	// Jalankan server (cukup panggil satu kali saja)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server gagal dijalankan: %v", err)
	}
}