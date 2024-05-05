package service

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/Dimoonevs/SportsApp/athletes-service/internal/repository"
	proto "github.com/Dimoonevs/SportsApp/athletes-service/proto/athletes"
	"github.com/Dimoonevs/SportsApp/athletes-service/proto/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AthletesService struct {
	Repo repository.AthletesRepository
	proto.UnimplementedAthleteServiceServer
}

func (a *AthletesService) CreateAthlete(ctx context.Context, req *proto.AthleteRequest) (*proto.AthleteResponse, error) {
	req.Status = "Active"
	if err := a.Repo.CreateAthlete(ctx, req); err != nil {
		log.Println(err)
		return &proto.AthleteResponse{}, err
	}
	return &proto.AthleteResponse{
		Message: "Athlete created successfully",
	}, nil
}

func (a *AthletesService) GetAthletes(req *proto.GetAthletesRequest, stream proto.AthleteService_GetAthletesServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("failed to get metadata")
	}
	token := md["token"][0]
	if err := validateToken(token); err != nil {
		log.Println(err)
		return err
	}
	for {
		athletes, err := a.Repo.GetAthletes(req.GroupId)
		if err != nil {
			log.Println(err)
			return err
		}

		for _, athlete := range athletes {
			err := a.checkAthleteStatus(athlete)
			if err != nil {
				return err
			}
			if err := stream.Send(athlete); err != nil {
				return err
			}
		}

		time.Sleep(5 * time.Second)
	}

}

func (a *AthletesService) DeleteAthletes(ctx context.Context, req *proto.DeleteAthletesRequest) (*proto.AthleteResponse, error) {
	err := a.Repo.DeleteAthletes(ctx, req)
	if err != nil {
		log.Println(err)
		return &proto.AthleteResponse{}, err
	}
	return &proto.AthleteResponse{
		Message: "Athletes deleted successfully",
	}, nil
}

func (a *AthletesService) EditAthlete(ctx context.Context, req *proto.EditAthletesRequest) (*proto.AthleteResponse, error) {

	errs := make(chan error, 2)
	var wg sync.WaitGroup

	wg.Add(len(req.Athlete))
	for _, athlete := range req.Athlete {
		go func(athlete *proto.AthleteRequest) {
			defer wg.Done()
			if err := a.Repo.EditAthlete(ctx, athlete); err != nil {
				errs <- err
			}
		}(athlete)
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return nil, err
		}
	}
	return &proto.AthleteResponse{
		Message: "Athletes edited successfully",
	}, nil
}

func (a *AthletesService) AddTraining(ctx context.Context, req *proto.AddTrainingRequest) (*proto.AthleteResponse, error) {
	errs := make(chan error, 2)
	var wg sync.WaitGroup

	wg.Add(len(req.Training))
	for _, training := range req.Training {
		go func(training *proto.TrainingData) {
			defer wg.Done()
			if err := a.Repo.AddTraining(ctx, training); err != nil {
				errs <- err
			}
		}(training)
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return nil, err
		}
	}
	return &proto.AthleteResponse{
		Message: "Training added successfully",
	}, nil
}

func validateToken(token string) error {
	conn, err := grpc.Dial("security-service:50001", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		return err
	}
	defer conn.Close()

	client := security.NewSecurityServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err = client.Validate(ctx, &security.ValidateRequest{
		Token: token,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *AthletesService) checkAthleteStatus(athlete *proto.AthleteRequest) error {
	timeZone, err := getTimeZoneFromSecurity(athlete.Id)
	if err != nil {
		return err
	}
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		return err
	}
	currentTime := time.Now().In(location)

	dateLastAthlete, err := time.Parse("02.01.2006", athlete.DateLast)
	if err != nil {
		return err
	}
	if currentTime.After(dateLastAthlete) {
		athlete.Status = "Expired"
		athlete.DaysLeft = 0
		err := a.Repo.EditAthlete(context.Background(), athlete)
		if err != nil {
			return err
		}
	}
	return nil
}

func getTimeZoneFromSecurity(id int32) (string, error) {
	conn, err := grpc.Dial("security-service:50001", grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := security.NewSecurityServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := client.GetCoachTimeZone(ctx, &security.GetCoachTimeZoneRequest{
		Id: id,
	})
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}
