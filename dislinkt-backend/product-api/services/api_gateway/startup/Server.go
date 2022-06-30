package startup

import (
	userGw "common/module/proto/user_service"
	"context"
	"fmt"
	cfg "gateway/module/startup/config"
	gorilla_handlers "github.com/gorilla/handlers"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux // Part of grpcGateway library
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	//	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	//Povezuje sa grpc generisanim fajlovima
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(20*1024*1024),
			grpc.MaxCallSendMsgSize(20*1024*1024)),
	}

	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)

	//postsEndpoint := fmt.Sprintf("%s:%s", server.config.PostsHost, server.config.PostsPort)

	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("fdfddf" + userEndpoint)
	}

	/*	err = postsGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postsEndpoint, opts)
		if err != nil {
			panic(err)
		}*/

}

func (server *Server) Start() {
	cors := gorilla_handlers.CORS(
		gorilla_handlers.AllowedOrigins([]string{"https://localhost:4200", "https://localhost:4200/**", "http://localhost:4200", "http://localhost:4200/**", "http://localhost:8080/**"}),
		gorilla_handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorilla_handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Authorization", "Access-Control-Allow-*", "Access-Control-Allow-Origin", "*"}),
		gorilla_handlers.AllowCredentials(),
	)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "9090"), cors(muxMiddleware(server))))

}

func muxMiddleware(server *Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.mux.ServeHTTP(w, r)
	})
}
