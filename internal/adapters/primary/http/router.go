package http

import (
	"log"
	"net/http"

	"github.com/ccontreras/crispy-potato/internal/core/ports"
	"github.com/ccontreras/crispy-potato/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	tweetsEndpoint    = "/tweets"
	relationsEndpoint = "/relations"
	avatarsEndpoint   = "/avatars"
	bannersEndpoint   = "/banners"
)

// Router handles HTTP routing
type Router struct {
	userHandler     *UserHandler
	tweetHandler    *TweetHandler
	relationHandler *RelationHandler
	authMiddleware  *middleware.AuthMiddleware
	dbMiddleware    *middleware.DatabaseMiddleware
}

// NewRouter creates a new Router instance
func NewRouter(
	userService ports.UserService,
	tweetService ports.TweetService,
	relationService ports.RelationService,
	authMiddleware *middleware.AuthMiddleware,
	dbMiddleware *middleware.DatabaseMiddleware,
) *Router {
	return &Router{
		userHandler:     NewUserHandler(userService),
		tweetHandler:    NewTweetHandler(tweetService),
		relationHandler: NewRelationHandler(relationService),
		authMiddleware:  authMiddleware,
		dbMiddleware:    dbMiddleware,
	}
}

// SetupRoutes configures all routes
func (r *Router) SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/signUp", r.dbMiddleware.CheckConnection(r.userHandler.Register)).Methods("POST")
	router.HandleFunc("/signIn", r.dbMiddleware.CheckConnection(r.userHandler.Login)).Methods("POST")

	// Protected routes
	router.HandleFunc("/{id}/{idBox}/me", r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.userHandler.GetProfile))).Methods("GET")
	router.HandleFunc("/me", r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.userHandler.UpdateProfile))).Methods("PUT")

	// Tweet routes
	router.HandleFunc(tweetsEndpoint, r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.tweetHandler.CreateTweet))).Methods("POST")
	router.HandleFunc(tweetsEndpoint, r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.tweetHandler.GetTweets))).Methods("GET")
	router.HandleFunc(tweetsEndpoint, r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.tweetHandler.DeleteTweet))).Methods("DELETE")

	// File upload routes
	router.HandleFunc(avatarsEndpoint, r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.userHandler.GetAvatar))).Methods("GET")
	router.HandleFunc(avatarsEndpoint, r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.userHandler.UploadAvatar))).Methods("POST")
	router.HandleFunc(bannersEndpoint, r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.userHandler.GetBanner))).Methods("GET")
	router.HandleFunc(bannersEndpoint, r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.userHandler.UploadBanner))).Methods("POST")

	// Relation routes
	router.HandleFunc(relationsEndpoint, r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.relationHandler.Follow))).Methods("POST")
	router.HandleFunc(relationsEndpoint, r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.relationHandler.Unfollow))).Methods("DELETE")
	router.HandleFunc(relationsEndpoint, r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.relationHandler.GetRelationStatus))).Methods("GET")

	// User list routes
	router.HandleFunc("/usersFollow", r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.userHandler.GetUsers))).Methods("GET")
	router.HandleFunc("/tweetsFollowers", r.dbMiddleware.CheckConnection(r.authMiddleware.ValidateJWT(r.tweetHandler.GetTweetsFromFollowers))).Methods("GET")

	return router
}

// Start starts the HTTP server
func (r *Router) Start(port string) {
	router := r.SetupRoutes()
	handler := cors.AllowAll().Handler(router)

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
