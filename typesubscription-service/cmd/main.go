package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Dimoonevs/SportsApp/typesubscription-service/internal/repository"
	"github.com/Dimoonevs/SportsApp/typesubscription-service/internal/service"
	proto "github.com/Dimoonevs/SportsApp/typesubscription-service/proto/subscription"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	gRpcPort = "50002"
)

func main() {
	var repo repository.SubscriptionRepository

	repo = &repository.SubscriptionRepositoryPostgres{
		Db: repository.ConnectToPostgres(),
	}

	service := &service.SubscriptionService{
		Repo: repo,
	}
	grpcListen(service)

}

func grpcListen(service *service.SubscriptionService) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	s := grpc.NewServer()

	proto.RegisterTypeSubscriptionServiceServer(s, service)
	log.Printf("grps server started on port %s", gRpcPort)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
