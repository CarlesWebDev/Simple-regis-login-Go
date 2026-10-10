package handler

import (
	"go-auth/internal/dto"
	"encoding/json"
	"net/http"

	"go-auth/internal/service"
)

// Struct Handler
type AuthHandler struct {
	service service.AuthService
}

// Constructor
func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

// Helper function Json
func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// handler Register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithJSON(w, http.StatusMethodNotAllowed, dto.AuthResponse{Message: "Method Not Allowed"})
		return
	}

	// tangkap Body JSon dari User lalu Masuk Ke DTO
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, dto.AuthResponse{Message: "Format JSON Tidak Valid"})
		return
	}

	//  Kirim DTO ke Layer Service untuk di proses
	if err := h.service.Register(req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, dto.AuthResponse{Message: err.Error()})
		return
	}
	respondWithJSON(w, http.StatusCreated, dto.AuthResponse{Message: "User registered successfully"})
}

// handler Login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithJSON(w, http.StatusMethodNotAllowed, dto.AuthResponse{Message: "Method Not Allowed"})
		return
	}

	// tangkap Body JSon dari User lalu Masuk Ke DTO
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, dto.AuthResponse{Message: "Format JSON Tidak Valid"})
		return
	}

	//  Kirim DTO ke Layer Service untuk di proses
	if err := h.service.Login(req); err != nil {
		respondWithJSON(w, http.StatusUnauthorized, dto.AuthResponse{Message: err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, dto.AuthResponse{Message: "Login successful"})
}
