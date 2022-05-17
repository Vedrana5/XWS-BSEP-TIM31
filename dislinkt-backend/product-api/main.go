package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	router := mux.NewRouter()

	s := http.Server{
		Addr:         ":8082",           // configure the bind address
		Handler:      router,            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 8082")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	SetupDatabase()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

var db *gorm.DB
var err error

func SetupDatabase() *gorm.DB {
	const DIALECT = "postgres"
	const HOST = "localhost"
	const PG_DBPORT = "8080"
	const PG_USER = "postgres"
	const XML_DB_NAME = "xws_project"
	const PG_PASSWORD = "postgres" //fakultet

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode= disable ", HOST, PG_DBPORT, PG_USER, PG_PASSWORD, XML_DB_NAME)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully connected to database")
	}

	//db.AutoMigrate(&data.User{})

	return db
}
