package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ccontreras/crispy-potato/internal/core/ports"
)

const (
	contentTypeHeader = "Content-Type"
	applicationJSON   = "application/json"
	imageJPEG         = "image/jpeg"
	invalidBodyMsg    = "Invalid request body"
	userIDRequiredMsg = "User ID is required"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userService ports.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register handles user registration
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, invalidBodyMsg, http.StatusBadRequest)
		return
	}

	user, err := h.userService.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(contentTypeHeader, applicationJSON)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login handles user login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, invalidBodyMsg, http.StatusBadRequest)
		return
	}

	loginResponse, err := h.userService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(contentTypeHeader, applicationJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginResponse)
}

// GetProfile handles getting user profile
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, userIDRequiredMsg, http.StatusBadRequest)
		return
	}

	profile, err := h.userService.GetProfile(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(contentTypeHeader, applicationJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

// UpdateProfile handles updating user profile
func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	var req struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Biographic string `json:"biographic"`
		Location   string `json:"location"`
		Website    string `json:"website"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, invalidBodyMsg, http.StatusBadRequest)
		return
	}

	err := h.userService.UpdateProfile(r.Context(), userID, req.Name, req.Surname, req.Biographic, req.Location, req.Website)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetUsers handles getting users list
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	pageStr := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")
	userType := r.URL.Query().Get("type")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	users, err := h.userService.GetUsers(r.Context(), userID, page, search, userType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(contentTypeHeader, applicationJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// UploadAvatar handles avatar upload
func (h *UserHandler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file data", http.StatusBadRequest)
		return
	}

	err = h.userService.UploadAvatar(r.Context(), userID, fileData, handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UploadBanner handles banner upload
func (h *UserHandler) UploadBanner(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file data", http.StatusBadRequest)
		return
	}

	err = h.userService.UploadBanner(r.Context(), userID, fileData, handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAvatar handles getting user avatar
func (h *UserHandler) GetAvatar(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, userIDRequiredMsg, http.StatusBadRequest)
		return
	}

	avatarData, err := h.userService.GetAvatar(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(contentTypeHeader, imageJPEG)
	w.Write(avatarData)
}

// GetBanner handles getting user banner
func (h *UserHandler) GetBanner(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, userIDRequiredMsg, http.StatusBadRequest)
		return
	}

	bannerData, err := h.userService.GetBanner(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(contentTypeHeader, imageJPEG)
	w.Write(bannerData)
}
