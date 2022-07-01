package startup

import (
	postProto "common/module/proto/post_service"
	"fmt"

	_ "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"

	"post/module/handler"
	"post/module/repository"
	"post/module/service"
	"post/module/startup/config"
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
	QueueGroup = "post_service"
)

func (server *Server) initPostHandler(postService *service.PostService) *handler.PostHandler {
	return handler.NewPostHandler(postService)
}

func initPostService(repo *repository.PostRepo) *service.PostService {
	return &service.PostService{Repo: repo}
}

func initUserRepo(database *gorm.DB) *repository.PostRepo {
	return &repository.PostRepo{Database: database}
}

func (server *Server) Start() {

	db = SetupDatabase()

	postRepo := initUserRepo(db)
	postService := initPostService(postRepo)
	postHandler := server.initPostHandler(postService)

	server.StartGrpcServer(postHandler)

}

var db *gorm.DB
var err error

func SetupDatabase() *gorm.DB {
	return nil
}

func (server *Server) StartGrpcServer(handler *handler.PostHandler) {
	fmt.Printf("USLA SAM U POST SERVER")
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
	postProto.RegisterPostServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
