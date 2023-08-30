package handlers

import (
	"fmt"
	"net/http"

	"abc.com/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Router() http.Handler {
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	//check health router
	router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "1.1.0")
	})
	userSub := router.PathPrefix("/users").Subrouter()
	userSub.HandleFunc("", middlewares.AuthMiddleware(GetUserHandler)).Methods("GET")
	userSub.HandleFunc("", PostUserHandler).Methods("POST")
	userSub.HandleFunc("/{id}", middlewares.AuthMiddleware(GetUserHandler)).Methods("GET")
	userSub.HandleFunc("/{id}", middlewares.AuthMiddleware(DeleteUserHandler)).Methods("DELETE")

	authSub := router.PathPrefix("/auth").Subrouter()
	authSub.HandleFunc("/login", LoginHandler).Methods("POST")

	handler := c.Handler(router)
	return handler
}
