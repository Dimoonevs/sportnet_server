package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Dimoonevs/SportsApp/athletes-service/internal/repository"
	"github.com/Dimoonevs/SportsApp/athletes-service/internal/service"
	proto "github.com/Dimoonevs/SportsApp/athletes-service/proto/athletes"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	gRpcPort    = "50004"
	grpcWebPort = "50014"
)

func main() {

	repositoryAthletes := &repository.AthletesRepositoryPostgres{
		Db:    repository.ConnectToPostgres(),
		Redis: repository.ConnectToRedis(),
	}

	service := &service.AthletesService{
		Repo: repositoryAthletes,
	}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		grpcListen(service)
	}()

	go func() {
		defer wg.Done()
		grpcWebListen(service)
	}()

	wg.Wait()
}

func grpcListen(service *service.AthletesService) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	s := grpc.NewServer()

	proto.RegisterAthleteServiceServer(s, service)
	log.Printf("grps server started on port %s", gRpcPort)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
func grpcWebListen(service *service.AthletesService) {
	s := grpc.NewServer()

	proto.RegisterAthleteServiceServer(s, service)

	wrappedGrpc := grpcweb.WrapServer(s)

	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%s", grpcWebPort),
		Handler: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			resp.Header().Set("Access-Control-Allow-Origin", "*")
			resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			resp.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-user-agent, x-grpc-web, token")
			if wrappedGrpc.IsGrpcWebRequest(req) {
				wrappedGrpc.ServeHTTP(resp, req)
			}
		}),
	}

	log.Printf("gRPC and gRPC-Web server started on port %s", gRpcPort)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
