package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KushagraMehta/blog/JWT-with-React+Go/Code/backend/controller"
	"github.com/KushagraMehta/blog/JWT-with-React+Go/Code/backend/middleware"
	"github.com/gorilla/mux"
)

var handler controller.Handler

var port = os.Getenv("PORT")

func main() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	handler.Init()

	router := mux.NewRouter()

	router.HandleFunc("/signup", handler.PostUser).Methods(http.MethodPost)
	router.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/logout", handler.Logout).Methods(http.MethodGet)

	secure := router.PathPrefix("/auth").Subrouter()
	secure.HandleFunc("/", handler.Auth).Methods(http.MethodGet)
	secure.HandleFunc("/user", handler.GetUser).Methods(http.MethodGet)
	secure.Use(middleware.Auth)

	// Hosting React build files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("build")))

	// ---------Code Under this is not important just for good programming practice--------------
	var addr string
	if port == "" {
		addr = "0.0.0.0:8090"
	} else {
		addr = fmt.Sprintf("0.0.0.0:%v", port)
	}
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	go func() {
		log.Printf("Starting server at port http://%v\n", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
