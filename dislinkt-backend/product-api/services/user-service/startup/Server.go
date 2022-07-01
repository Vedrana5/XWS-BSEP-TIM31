package startup

import (
	userProto "common/module/proto/user_service"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"net"
	"os"
	"user/module/handler"
	"user/module/model"
	"user/module/repository"
	"user/module/service"
	"user/module/startup/config"
	"user/module/util"

	"gorm.io/gorm"
	"log"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "user-service"
)

func initPasswordUtil() *util.PasswordUtil {
	return &util.PasswordUtil{}
}

func (server *Server) initRegisterHandler(rbac *gorbac.RBAC, permissionUpdateUserInfo *gorbac.Permission, confirmationTokenService *service.ConfirmationTokenService, passwordUtil *util.PasswordUtil, LogInfo *logrus.Logger, LogError *logrus.Logger, userService *service.UserService) *handler.RegisterHandler {
	return handler.NewRegisterHandler(rbac, permissionUpdateUserInfo, confirmationTokenService, passwordUtil, LogInfo, LogError, userService)
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

func (server *Server) Start() {
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
	registerHandler := server.initRegisterHandler(rbac, &permissionUpdateUserInfo, confirmationTokenService, passwordUtil, logInfo, logError, userService)

	validator := validator.New()
	userHandler := initUserHandler(validationCodeService, userService, logInfo, logError)
	updateProfilHandler := initUpdateProfileHandler(rbac, &permissionFindAllUsers, userService, &permissionUpdateUserInfo, passwordUtil, logInfo, logError, validator)

	confirmationTokenHandler := initConfirmationTokenHandler(logInfo, logError, confirmationTokenService, userService)

	server.StartGrpcServer(userHandler, loginHandler, registerHandler, updateProfilHandler, confirmationTokenHandler)

}

func (server *Server) StartGrpcServer(handler *handler.UserHandler, loginHandler *handler.LogInHandler, registerHandler *handler.RegisterHandler, profilHandler *handler.UpdateProfileHandler, tokenHandler *handler.ConfirmationTokenHandler) {
	fmt.Printf("USLA SAM U STARTGRPC SERVER")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	/*
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
		if err != nil {
			log.Fatalf("failed to parse public key: %v", err)
		}
		interceptor := interceptor.NewAuthInterceptor(config.AccessibleRoles(), publicKey, logError)
	*/
	grpcServer := grpc.NewServer()
	userProto.RegisterUserServiceServer(grpcServer, registerHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

const LOG_INFO_FILE = "logrusInfo.log"
const LOG_ERROR_FILE = "logrusError.log"

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
