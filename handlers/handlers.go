package handlers

import (
	"github.com/ccontreras/crispy-potato/middlew"
	"github.com/ccontreras/crispy-potato/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// Handlers for manage some routes, sets the port and start the server
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signUp", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/signIn", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/me", middlew.CheckDB(middlew.ValidateJWT(routers.Me))).Methods("GET")
	router.HandleFunc("/me", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweets", middlew.CheckDB(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/tweets", middlew.CheckDB(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/tweets", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/avatars", middlew.CheckDB(middlew.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/avatars", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/banners", middlew.CheckDB(middlew.ValidateJWT(routers.GetBanner))).Methods("GET")
	router.HandleFunc("/banners", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")

	router.HandleFunc("/relations", middlew.CheckDB(middlew.ValidateJWT(routers.InsertRelation))).Methods("POST")
	router.HandleFunc("/relations", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/relations", middlew.CheckDB(middlew.ValidateJWT(routers.ReadRelation))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
