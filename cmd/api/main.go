package main

import (
	"log"

	"github.com/ccontreras/crispy-potato/config"
	"github.com/ccontreras/crispy-potato/internal/adapters/primary/http"
	"github.com/ccontreras/crispy-potato/internal/adapters/secondary/auth"
	"github.com/ccontreras/crispy-potato/internal/adapters/secondary/database/mongodb"
	"github.com/ccontreras/crispy-potato/internal/adapters/secondary/storage"
	"github.com/ccontreras/crispy-potato/internal/core/services"
	"github.com/ccontreras/crispy-potato/pkg/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database connection
	conn, err := mongodb.NewConnection(cfg.Database.URI)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repositories
	userRepo := mongodb.NewUserRepository(conn)
	tweetRepo := mongodb.NewTweetRepository(conn)
	relationRepo := mongodb.NewRelationRepository(conn)

	// Initialize secondary adapters
	hasher := auth.NewPasswordHasher()
	tokenGen := auth.NewJWTTokenGenerator(cfg.JWT.SecretKey)
	fileStorage := storage.NewLocalFileStorage(cfg.Storage.BasePath)

	// Initialize services
	userService := services.NewUserService(userRepo, hasher, tokenGen, fileStorage)
	tweetService := services.NewTweetService(tweetRepo)
	relationService := services.NewRelationService(relationRepo)

	// Initialize middlewares
	authMiddleware := middleware.NewAuthMiddleware(tokenGen, userRepo)
	dbMiddleware := middleware.NewDatabaseMiddleware(conn)

	// Initialize and start router
	router := http.NewRouter(
		userService,
		tweetService,
		relationService,
		authMiddleware,
		dbMiddleware,
	)

	router.Start(cfg.Server.Port)
}
