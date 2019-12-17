package cmd

import (
	"fmt"
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

	// init client layers
	cStore := store.NewClientStore(db)
	cService := service.NewClientService(cStore)
	cHandler := handler.NewClientHandler(cService)

	// init user layers
	uStore := store.NewUserStore(db)
	uService := service.NewUserService(uStore)
	uHandler := handler.NewUserHandler(uService)

	// init auth layers
	aServ := service.NewAuthorizeService()
	ahandler := handler.NewAuthorizeHandler(aServ)

	// TODO: WRAP ALL HANDLERS WITH ServeHTTP
	// TODO: handler interface
	router.HandleFunc("/client/register", cHandler.Register).Methods("POST")
	router.HandleFunc("/user/register", uHandler.Register).Methods("POST")
	router.HandleFunc("/user/login", uHandler.Login).Methods("POST")

	router.HandleFunc("/authorize", ahandler.Authorize).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Running server on port: 8080 .....")
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Error: %+v", err)
	}
	return nil
}
