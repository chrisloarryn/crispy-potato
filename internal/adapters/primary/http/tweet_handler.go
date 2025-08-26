package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ccontreras/crispy-potato/internal/core/ports"
)

const (
	tweetIDRequiredMsg = "Tweet ID is required"
	pageRequiredMsg    = "Page parameter is required"
)

// TweetHandler handles HTTP requests for tweet operations
type TweetHandler struct {
	tweetService ports.TweetService
}

// NewTweetHandler creates a new TweetHandler instance
func NewTweetHandler(tweetService ports.TweetService) *TweetHandler {
	return &TweetHandler{
		tweetService: tweetService,
	}
}

// CreateTweet handles tweet creation
func (h *TweetHandler) CreateTweet(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	var req struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, invalidBodyMsg, http.StatusBadRequest)
		return
	}

	err := h.tweetService.CreateTweet(r.Context(), userID, req.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetTweets handles getting tweets by user
func (h *TweetHandler) GetTweets(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, userIDRequiredMsg, http.StatusBadRequest)
		return
	}

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		http.Error(w, pageRequiredMsg, http.StatusBadRequest)
		return
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	tweets, err := h.tweetService.GetTweetsByUser(r.Context(), userID, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(contentTypeHeader, applicationJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}

// DeleteTweet handles tweet deletion
func (h *TweetHandler) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	tweetID := r.URL.Query().Get("id")

	if tweetID == "" {
		http.Error(w, tweetIDRequiredMsg, http.StatusBadRequest)
		return
	}

	err := h.tweetService.DeleteTweet(r.Context(), tweetID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetTweetsFromFollowers handles getting tweets from followed users
func (h *TweetHandler) GetTweetsFromFollowers(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		http.Error(w, pageRequiredMsg, http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	tweets, err := h.tweetService.GetTweetsFromFollowers(r.Context(), userID, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(contentTypeHeader, applicationJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}
