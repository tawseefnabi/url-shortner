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
	controller "url-shortner/Controller"
	repository "url-shortner/Repository"
	service "url-shortner/Service"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Database = "tinyUrl.db"
)
var db *gorm.DB

// var rdb *redis.Client
var err error

func init() {
	db, err = gorm.Open(sqlite.Open(Database), &gorm.Config{})
	if err != nil {
		log.Panicln("Database Not initilized!!!!!")
		log.Panic("Error: ", err.Error())
		return
	}
	fmt.Println("Database Initialized")
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	router := mux.NewRouter()
	router.HandleFunc("/generateTinyUrl/", controller.GenerateTinyUrl).Methods("POST")

	log.Println("Application Started")
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
