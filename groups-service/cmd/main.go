package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Dimoonevs/SportAspp/groups-service/internal/repository"
	"github.com/Dimoonevs/SportAspp/groups-service/internal/service"
	proto "github.com/Dimoonevs/SportAspp/groups-service/proto/groups"
	"google.golang.org/grpc"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	gRpcPort = "50003"
)

func main() {
	var repo repository.GroupRepository

	repo = &repository.GroupRepositoryPortgres{
		DB: repository.ConnectToPostgres(),
	}

	service := &service.GroupService{
		Repo: repo,
	}

	grpcListen(service)
}
func grpcListen(service *service.GroupService) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	s := grpc.NewServer()

	proto.RegisterGroupServiceServer(s, service)
	log.Printf("grps server started on port %s", gRpcPort)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
