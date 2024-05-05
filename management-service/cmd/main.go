package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Dimoonevs/SportsApp/management-service/internal/repository"
	"github.com/Dimoonevs/SportsApp/management-service/internal/service"
	"github.com/Dimoonevs/SportsApp/management-service/proto/management"
	"google.golang.org/grpc"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	gRpcPort = "50005"
)

func main() {
	var repo repository.Repository

	repo = &repository.RepositoryPostgres{
		PostgresDb: repository.ConnectToPostgres(),
		MongoDB:    repository.ConnectToMongo(),
	}

	service := &service.ManagementService{
		Repo: repo,
	}
	go func() {
		err := service.LoadCron(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()

	grpcListen(service)

}

func grpcListen(service *service.ManagementService) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	s := grpc.NewServer()

	management.RegisterManagementServiceServer(s, service)
	log.Printf("grps server started on port %s", gRpcPort)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
