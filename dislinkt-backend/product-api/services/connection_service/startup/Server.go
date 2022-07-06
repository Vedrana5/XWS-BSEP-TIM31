package startup

import (
	cproto "common/module/proto/connection_service"
	"connection/module/handler"
	"connection/module/repository"
	"connection/module/service"
	"connection/module/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
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

func (server *Server) InitConnectionService(repo repository.ConnectionRepo) *service.ConnectionService {
	return service.NewConnectionService(repo)
}

func (server *Server) InitConnectionHandler(connectionService *service.ConnectionService) *handler.ConnectionHandler {
	return handler.NewConnectionHandler(connectionService)
}

func (server *Server) InitConnectionRepo(client *mongo.Client) repository.ConnectionRepo {
	return repository.NewConnectionRepository(client)
}

func (server *Server) Start() {

	mongoClient := server.InitMongoClient()

	postRepo := server.InitConnectionRepo(mongoClient)
	postService := server.InitConnectionService(postRepo)
	postHandler := server.InitConnectionHandler(postService)

	server.StartGrpcServer(postHandler)

}
func (server *Server) StartGrpcServer(handler *handler.ConnectionHandler) {
	fmt.Printf("USLA SAM U CONNECTION SERVER")
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
	cproto.RegisterConnectionServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
