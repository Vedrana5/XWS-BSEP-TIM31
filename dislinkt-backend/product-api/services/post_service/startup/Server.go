package startup

import (
	postProto "common/module/proto/post_service"
	"fmt"
	_ "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

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

func (server *Server) InitMongoClient() *mongo.Client {
	client, err := repository.GetClient(server.config.PostDBHost, server.config.PostDBPort)


	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connected to mongo database!")
	}

	return client
}

func (server *Server) InitPostService(repo repository.PostRepo) *service.PostService {
	return service.NewPostService(repo)
}

func (server *Server) InitPostHandler(postService *service.PostService) *handler.PostHandler {
	return handler.NewPostHandler(postService)
}

func (server *Server) InitPostRepo(client *mongo.Client) repository.PostRepo {
	return repository.NewPostRepository(client)
}

func (server *Server) Start() {

	mongoClient := server.InitMongoClient()

	postRepo := server.InitPostRepo(mongoClient)
	postService := server.InitPostService(postRepo)
	postHandler := server.InitPostHandler(postService)

	server.StartGrpcServer(postHandler)

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
