package http

import (
	"encoding/json"
	"net/http"

	"github.com/ccontreras/crispy-potato/internal/core/ports"
)

// RelationHandler handles HTTP requests for relation operations
type RelationHandler struct {
	relationService ports.RelationService
}

// NewRelationHandler creates a new RelationHandler instance
func NewRelationHandler(relationService ports.RelationService) *RelationHandler {
	return &RelationHandler{
		relationService: relationService,
	}
}

// Follow handles following a user
func (h *RelationHandler) Follow(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	targetUserID := r.URL.Query().Get("id")

	if targetUserID == "" {
		http.Error(w, userIDRequiredMsg, http.StatusBadRequest)
		return
	}

	err := h.relationService.Follow(r.Context(), userID, targetUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Unfollow handles unfollowing a user
func (h *RelationHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	targetUserID := r.URL.Query().Get("id")

	if targetUserID == "" {
		http.Error(w, userIDRequiredMsg, http.StatusBadRequest)
		return
	}

	err := h.relationService.Unfollow(r.Context(), userID, targetUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetRelationStatus handles checking relation status
func (h *RelationHandler) GetRelationStatus(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	targetUserID := r.URL.Query().Get("id")

	if targetUserID == "" {
		http.Error(w, userIDRequiredMsg, http.StatusBadRequest)
		return
	}

	response, err := h.relationService.IsFollowing(r.Context(), userID, targetUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(contentTypeHeader, applicationJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
