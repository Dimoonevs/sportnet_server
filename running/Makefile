BROKER_BINARY=brokerApp
SECURITY_BINARY=securityApp
LOGGER_BINARY=loggerApp
SUB_BINARY=subscriptionApp
GROUP_BINARY=groupsApp
ATHLETES_BINARY=athletesApp
MANAGEMENT_BINARY=managementApp

## up: starts all containers in the background without forcing build
up:
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_security build_logger build_subscription build_groups build_athletes build_management
	@echo Stopping docker images (if running...)
	docker-compose down
	@echo Building (when required) and starting docker images...
	docker-compose up --build -d
	@echo Docker images built and started!

## down: stop docker compose
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo Building broker binary...
	chdir ..\broker-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${BROKER_BINARY} ./cmd
	@echo Done!

## build_auth: builds the auth binary as a linux executable
build_security:
	@echo Building security binary...
	chdir ..\security-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${SECURITY_BINARY} ./cmd
	@echo Done!

## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo Building logger binary...
	chdir ..\logger-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${LOGGER_BINARY} ./cmd
	@echo Done!

## build_subscription: builds the subscription binary as a linux executable
build_subscription:
	@echo Building subscription binary...
	chdir ..\typesubscription-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${SUB_BINARY} ./cmd
	@echo Done!

## build_groups: builds the groups binary as a linux executable
build_groups:
	@echo Building groups binary...
	chdir ..\groups-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${GROUP_BINARY} ./cmd
	@echo Done!

## build_athletes: builds the athletes binary as a linux executable
build_athletes:
	@echo Building athletes binary...
	chdir ..\athletes-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${ATHLETES_BINARY} ./cmd
	@echo Done!

## build_management: builds the management binary as a linux executable
build_management:
	@echo Building management binary...
	chdir ..\management-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${MANAGEMENT_BINARY} ./cmd
	@echo Done!


## generate_proto: generates the proto files
generate_proto:
	@echo Generating proto...
	export PATH=$PATH:$(go env GOPATH)/bin
## 	Security
	protoc --proto_path=../protoRoot/security --go_out=../broker-service/proto/ --go-grpc_out=../broker-service/proto/ --go_out=../security-service/proto/ --go-grpc_out=../security-service/proto/ security.proto
	protoc --proto_path=../protoRoot/security --go_out=../typesubscription-service/proto/ --go-grpc_out=../typesubscription-service/proto/ --go_out=../athletes-service/proto/ --go-grpc_out=../athletes-service/proto/  security.proto
## 	Athletes
	protoc --proto_path=../protoRoot/athletes --go_out=../broker-service/proto/ --go-grpc_out=../broker-service/proto/ --go_out=../athletes-service/proto/ --go-grpc_out=../athletes-service/proto/ athletes.proto
	protoc --proto_path=../protoRoot/athletes --go_out=../management-service/proto/ --go-grpc_out=../management-service/proto/ athletes.proto
## 	Groups
	protoc --proto_path=../protoRoot/groups --go_out=../broker-service/proto/ --go-grpc_out=../broker-service/proto/ --go_out=../groups-service/proto/ --go-grpc_out=../groups-service/proto/ groups.proto
## 	Subscription
	protoc --proto_path=../protoRoot/subscription --go_out=../broker-service/proto/ --go-grpc_out=../broker-service/proto/ --go_out=../typesubscription-service/proto/ --go-grpc_out=../typesubscription-service/proto/ subscription.proto
## 	Management
	protoc --proto_path=../protoRoot/management --go_out=../typesubscription-service/proto/ --go-grpc_out=../typesubscription-service/proto/ management.proto
	protoc --proto_path=../protoRoot/management --go_out=../management-service/proto/ --go-grpc_out=../management-service/proto/ management.proto
	@echo Done!