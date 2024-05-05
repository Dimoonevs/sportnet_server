package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Dimoonevs/SportsApp/auth-service/internal/repository"
	"github.com/Dimoonevs/SportsApp/auth-service/internal/service"
	"github.com/Dimoonevs/SportsApp/auth-service/pkg/utils"
	proto "github.com/Dimoonevs/SportsApp/auth-service/proto/security"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
)

const (
	gRpcPort  = "50001"
	secretKey = "r43t18sc"
)

func main() {
	var repo repository.Repository

	repo = &repository.RepositoryPostgres{
		DB: repository.ConnectToPostgres(),
	}

	jwt := utils.JwtWrapper{
		SecretKey:       secretKey,
		Issuer:          "go-secrit-service",
		ExpirationHours: 24 * 100,
	}

	service := &service.SecurityService{
		Repo:       repo,
		JwtWrapper: jwt,
	}
	grpcListen(service)
}

func grpcListen(service *service.SecurityService) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	s := grpc.NewServer()

	proto.RegisterSecurityServiceServer(s, service)
	log.Printf("grps server started on port %s", gRpcPort)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
