package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/handler"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/repository"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/service"
	"github.com/go-playground/validator"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*	func initActivityRepo(database *gorm.DB) *repository.ActivityRepository{
	return &repository.ActivityRepository{ Database: database }
}*/

func initCommentRepo(database *gorm.DB) *repository.CommentRepository {
	return &repository.CommentRepository{Database: database}
}

func initPostRepo(database *gorm.DB) *repository.PostRepository {
	return &repository.PostRepository{Database: database}
}

/*
	func initActivityService(repo *repository.ActivityRepository) *service.ActivityService{
		return &service.ActivityService{ Repo: repo }
	}*/

func initCommentService(repo *repository.CommentRepository) *service.CommentService {
	return &service.CommentService{Repo: repo}
}

func initPostService(repo *repository.PostRepository) *service.PostService {
	return &service.PostService{Repo: repo}
}

/*	func initActivityHandler(SinglePostService *service.SinglePostService,logInfo *logrus.Logger, logError *logrus.Logger,service *service.ActivityService) *handler.ActivityHandler{
		return &handler.ActivityHandler{
			Service:           service,
			SinglePostService: SinglePostService,
			LogInfo:           logInfo,
			LogError:          logError,
		}
	}
*/
func initCommentHandler(PostService *service.PostService, logInfo *logrus.Logger, logError *logrus.Logger, service *service.CommentService, validate *validator.Validate) *handler.CommentHandler {
	return &handler.CommentHandler{
		Service:     service,
		PostService: PostService,
	}
}

func initPostHandler(logInfo *logrus.Logger, logError *logrus.Logger, postService *service.PostService) *handler.PostHandler {
	return &handler.PostHandler{LogInfo: logInfo, LogError: logError, PostService: postService}
}

func Handle(postHandler *handler.PostHandler, commentHandler *handler.CommentHandler) {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	router := mux.NewRouter()

	s := http.Server{
		Addr:         ":8080",           // configure the bind address
		Handler:      router,            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 8080")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	/*	router.HandleFunc("/register", registerHandler.CreateUser).Methods("POST")
		router.HandleFunc("/login", logInHandler.LogIn).Methods("POST")
		router.HandleFunc("/updateProfil", updateProfilHandler.UpdateUserProfileInfo).Methods("POST")
		router.HandleFunc("/findPublicUser", userHandler.FindByUserName).Methods("GET")
	*/
	s.ListenAndServe()
}

func main() {
	db = SetupDatabase()
	/*	rbac := gorbac.New()*/

	logInfo := logrus.New()
	logError := logrus.New()

	/*	repoActivity := initActivityRepo(db)*/
	repoComment := initCommentRepo(db)
	repoPost := initPostRepo(db)

	/*serviceActivity := initActivityService(repoActivity)*/
	serviceComment := initCommentService(repoComment)
	servicePost := initPostService(repoPost)

	validator := validator.New()

	/*	ActivityHandler := initActivityHandler(servicePost,logInfo, logError, serviceActivity)*/
	commentHandler := initCommentHandler(servicePost, logInfo, logError, serviceComment, validator)
	postHandler := initPostHandler(logInfo, logError, servicePost)

	Handle(postHandler, commentHandler)
}

var db *gorm.DB
var err error

func SetupDatabase() *gorm.DB {
	const DIALECT = "postgres"
	const HOST = "localhost"
	const PG_DBPORT = "5432"
	const PG_USER = "postgres"
	const XML_DB_NAME = "xws_project"
	const PG_PASSWORD = "fakultet" //fakultet

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode= disable ", HOST, PG_DBPORT, PG_USER, PG_PASSWORD, XML_DB_NAME)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully connected to database")
	}

	db.AutoMigrate(&model.Post{})

	return db
}
