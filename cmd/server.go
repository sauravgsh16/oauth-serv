package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sauravgsh16/oauth-serv/db"
	"github.com/sauravgsh16/oauth-serv/handler"
	"github.com/sauravgsh16/oauth-serv/service"
	"github.com/sauravgsh16/oauth-serv/store"
)

// RunServer runs the auth server
func RunServer() error {
	db, err := db.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	router := mux.NewRouter()

	// init layers
	store := store.NewClientStore(db)
	service := service.NewClientService(store)
	handler := handler.NewClientHandler(service)

	// TODO: WRAP WITH HANDLERS WITH ServeHTTP
	// TODO: handler interface
	router.HandleFunc("/register", handler.Register).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Error: %+v", err)
	}
	return nil
}
