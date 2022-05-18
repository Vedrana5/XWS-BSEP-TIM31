package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/handler"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/repository"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/service"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/util"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPasswordUtil() *util.PasswordUtil {
	return &util.PasswordUtil{}
}

func initUserHandler(passwordUtil *util.PasswordUtil, LogInfo *logrus.Logger, LogError *logrus.Logger, userService *service.UserService) *handler.RegisterHandler {
	return &handler.RegisterHandler{
		passwordUtil,
		userService,
		LogInfo,
		LogError,
	}

}

func initLoginHandler(userService *service.UserService, passwordUtil *util.PasswordUtil, LogInfo *logrus.Logger, LogError *logrus.Logger) *handler.LogInHandler {
	return &handler.LogInHandler{
		userService,
		passwordUtil,
		LogInfo,
		LogError,
	}
}

func initUserRepo(database *gorm.DB) *repository.UserRepo {
	return &repository.UserRepo{Database: database}
}

func initUserService(repo *repository.UserRepo) *service.UserService {
	return &service.UserService{Repo: repo}
}

func Pocetn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AAAAA")
}

func Handle(registerHandler *handler.RegisterHandler, logInHandler *handler.LogInHandler) {
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
	/*
		// start the server
		go func() {
			l.Println("Starting server on port 8082")

			err := s.ListenAndServe()
			if err != nil {
				l.Printf("Error starting server: %s\n", err)
				os.Exit(1)
			}
		}()*/
	router.HandleFunc("/register", registerHandler.CreateUser).Methods("POST")
	router.HandleFunc("/login", logInHandler.LogIn).Methods("POST")
	s.ListenAndServe()
}

func main() {
	db = SetupDatabase()
	passwordUtil := initPasswordUtil()
	logInfo := logrus.New()
	logError := logrus.New()
	userRepo := initUserRepo(db)
	userService := initUserService(userRepo)
	loginHandler := initLoginHandler(userService, passwordUtil, logInfo, logError)
	userHandler := initUserHandler(passwordUtil, logInfo, logError, userService)
	Handle(userHandler, loginHandler)
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

	db.AutoMigrate(&model.User{})

	return db
}
