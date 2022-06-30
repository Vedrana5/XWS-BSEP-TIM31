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
	"github.com/go-playground/validator"
	"github.com/mikespook/gorbac"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPasswordUtil() *util.PasswordUtil {
	return &util.PasswordUtil{}
}

func initRegisterHandler(confirmationTokenService *service.ConfirmationTokenService, passwordUtil *util.PasswordUtil, LogInfo *logrus.Logger, LogError *logrus.Logger, userService *service.UserService) *handler.RegisterHandler {
	return &handler.RegisterHandler{
		confirmationTokenService,
		passwordUtil,
		userService,
		LogInfo,
		LogError,
	}

}

func initLoginHandler(validationCodeService *service.ValidationCodeService, userService *service.UserService, passwordUtil *util.PasswordUtil, LogInfo *logrus.Logger, LogError *logrus.Logger) *handler.LogInHandler {
	return &handler.LogInHandler{
		validationCodeService,
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


func Handle(registerHandler *handler.RegisterHandler, logInHandler *handler.LogInHandler, updateProfilHandler *handler.UpdateProfileHandler, userHandler *handler.UserHandler, confirmationTokenHandler *handler.ConfirmationTokenHandler) {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	router := mux.NewRouter()
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowCredentials(),
	)

	s := http.Server{
		Addr:         ":8089",           // configure the bind address
		Handler:      cors(router),      // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	//router.HandleFunc("/register", registerHandler.CreateUser).Methods("POST")
	router.HandleFunc("/login", logInHandler.LogIn).Methods("POST")
	//router.HandleFunc("/loginPasswordless", logInHandler.LogInPasswordless).Methods("POST")
	//router.HandleFunc("/updateProfil", updateProfilHandler.UpdateUserProfileInfo).Methods("POST")
	//router.HandleFunc("/findPublicUser", userHandler.FindPublicByUserName).Methods("GET")
	//router.HandleFunc("/findByUsername/{username}", userHandler.FindByUserName).Methods("GET")
	//router.HandleFunc("/findPublic/{username}", userHandler.FindPublic).Methods("GET")
	//router.HandleFunc("/findById/{id}", userHandler.FindById).Methods("GET")
	//router.HandleFunc("/resetPassword/{username}", userHandler.SendMailForResetPassword).Methods("POST")
	//router.HandleFunc("/resettPassword", userHandler.ResetPassword).Methods("POST")
	//router.HandleFunc("/confirmRegistration", confirmationTokenHandler.VerifyConfirmationToken).Methods("POST")
	//router.HandleFunc("/changePassword", userHandler.ChangePassword).Methods("POST")
	//router.HandleFunc("/findTokenByCode/{confirmationToken}", userHandler.FindTokenByCode).Methods("GET")
	//router.HandleFunc("/linkForPasswordless/{email}", userHandler.SendMailForPasswordlessLogin).Methods("POST")

	s.ListenAndServe()
}

func initUserHandler(validationCodeService *service.ValidationCodeService, userService *service.UserService, LogInfo *logrus.Logger, LogError *logrus.Logger) *handler.UserHandler {
	return &handler.UserHandler{
		ValidationCodeService: validationCodeService,
		UserService:           userService,
		LogInfo:               LogInfo,
		LogError:              LogError,
	}
}

func initUpdateProfileHandler(rbac *gorbac.RBAC, permissionFindAllUsers *gorbac.Permission, UserService *service.UserService, permissionUpdateUserInfo *gorbac.Permission, passwordUtil *util.PasswordUtil, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.UpdateProfileHandler {
	return &handler.UpdateProfileHandler{
		Rbac:                     rbac,
		PermissionFindAllUsers:   permissionFindAllUsers,
		UserService:              UserService,
		PermissionUpdateUserInfo: permissionUpdateUserInfo,
		PasswordUtil:             passwordUtil,
		LogInfo:                  LogInfo,
		LogError:                 LogError,
		Validator:                validator,
	}
}

func initConfirmationTokenService(repo *repository.ConfirmationTokenRepository) *service.ConfirmationTokenService {
	return &service.ConfirmationTokenService{Repo: repo}
}

func initConfirmationTokenRepo(database *gorm.DB) *repository.ConfirmationTokenRepository {
	return &repository.ConfirmationTokenRepository{Database: database}
}

func initValidationCodeRepo(database *gorm.DB) *repository.ValidationCodeRepo {
	return &repository.ValidationCodeRepo{Database: database}
}

func itValidationCodeService(repo *repository.ValidationCodeRepo) *service.ValidationCodeService {
	return &service.ValidationCodeService{Repo: repo}
}

func initConfirmationTokenHandler(LogInfo *logrus.Logger, LogError *logrus.Logger, confirmationTokenService *service.ConfirmationTokenService, userService *service.UserService) *handler.ConfirmationTokenHandler {
	return &handler.ConfirmationTokenHandler{
		ConfirmationTokenService: confirmationTokenService,
		UserService:              userService,
		LogInfo:                  LogInfo,
		LogError:                 LogError,
	}
}

const LOG_INFO_FILE = "logrusInfo.log"
const LOG_ERROR_FILE = "logrusError.log"

func main() {
	permissionFindAllUsers := gorbac.NewStdPermission("permission-find-all-users")
	permissionUpdateUserInfo := gorbac.NewStdPermission("permission-update-user-info")
	roleRegisteredUser := gorbac.NewStdRole("role-registered-user")
	//roleUnregisteredUser := gorbac.NewStdRole("role-unregistered-user")
	roleAdmin := gorbac.NewStdRole("role-admin")
	roleAdmin.Assign(permissionUpdateUserInfo)
	roleRegisteredUser.Assign(permissionUpdateUserInfo)
	rbac := gorbac.New()
	rbac.Add(roleAdmin)
	rbac.Add(roleRegisteredUser)

	db = SetupDatabase()
	passwordUtil := initPasswordUtil()

	logInfo := logrus.New()
	logInfo.Formatter = &logrus.JSONFormatter{}
	logInfo.SetOutput(os.Stdout)
	file, err := os.OpenFile(LOG_INFO_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logInfo.Fatal(err)
	}
	defer file.Close()
	logInfo.SetOutput(file)

	logError := logrus.New()
	logError.Formatter = &logrus.JSONFormatter{}
	logError.SetOutput(os.Stdout)
	file1, err1 := os.OpenFile(LOG_ERROR_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err1 != nil {
		logError.Fatal(err1)
	}
	defer file.Close()
	logError.SetOutput(file1)

	userRepo := initUserRepo(db)
	userService := initUserService(userRepo)
	confirmationTokenRepo := initConfirmationTokenRepo(db)
	confirmationTokenService := initConfirmationTokenService(confirmationTokenRepo)

	validationCodeRepo := initValidationCodeRepo(db)
	validationCodeService := itValidationCodeService(validationCodeRepo)

	loginHandler := initLoginHandler(validationCodeService, userService, passwordUtil, logInfo, logError)
	registerHandler := initRegisterHandler(confirmationTokenService, passwordUtil, logInfo, logError, userService)

	validator := validator.New()
	userHandler := initUserHandler(validationCodeService, userService, logInfo, logError)
	updateProfilHandler := initUpdateProfileHandler(rbac, &permissionFindAllUsers, userService, &permissionUpdateUserInfo, passwordUtil, logInfo, logError, validator)

	confirmationTokenHandler := initConfirmationTokenHandler(logInfo, logError, confirmationTokenService, userService)

	Handle(registerHandler, loginHandler, updateProfilHandler, userHandler, confirmationTokenHandler)
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

	db.AutoMigrate(&model.User{}, &model.ConfirmationToken{}, &model.ValidationCode{})

	return db
}
