package service

import (
	"context"
	"log"
	"time"

	"github.com/Dimoonevs/SportsApp/typesubscription-service/internal/repository"
	"github.com/Dimoonevs/SportsApp/typesubscription-service/proto/management"
	"github.com/Dimoonevs/SportsApp/typesubscription-service/proto/security"
	proto "github.com/Dimoonevs/SportsApp/typesubscription-service/proto/subscription"
	"google.golang.org/grpc"
)

type SubscriptionService struct {
	Repo repository.SubscriptionRepository
	proto.UnimplementedTypeSubscriptionServiceServer
}

func (s *SubscriptionService) CreateSubscription(ctx context.Context, req *proto.SubscriptionRequest) (*proto.SubscriptionResponse, error) {
	id, err := s.Repo.CreateSubscription(ctx, req)
	if err != nil {
		return nil, err
	}

	if req.AutomaticallyManagement {
		timeZone, err := getTimeZoneFromSecurity(req.GetCoachId())
		if err != nil {
			return nil, err
		}
		err = createShodullerFromGroup(timeZone, req.DaysOfWeek, id, req.Time)
		if err != nil {
			return nil, err
		}
	}
	log.Println("Subscription created successfully")
	return &proto.SubscriptionResponse{
		Message: "Subscription created successfully",
		Id:      id,
	}, nil
}

func (s *SubscriptionService) GetAllSubscriptions(ctx context.Context, req *proto.GetSubscriptionRequest) (*proto.GetSubscriptionResponse, error) {
	subscriptions, err := s.Repo.GetAllSubscriptions(ctx, req)
	if err != nil {
		return nil, err
	}
	return &proto.GetSubscriptionResponse{
		Subscriptions: subscriptions,
	}, nil
}

func (s *SubscriptionService) EditSubscription(ctx context.Context, req *proto.SubscriptionEditRequest) (*proto.SubscriptionResponse, error) {
	err := s.Repo.EditSubscription(ctx, req)
	if err != nil {
		log.Println("Service Error")
		return nil, err
	}
	if req.AutomaticallyManagement && req.CronId == 0 {
		timeZone, err := getTimeZoneFromSecurity(req.CoachId)
		if err != nil {
			return nil, err
		}
		err = createShodullerFromGroup(timeZone, req.DaysOfWeek, req.Id, req.Time)
		if err != nil {
			return nil, err
		}
	}
	if req.CronId != 0 {
		if !req.AutomaticallyManagement {
			err = deleteScheduler(req.IdScheduler, req.CronId)
			if err != nil {
				return nil, err
			}
		} else {
			timeZone, err := getTimeZoneFromSecurity(req.CoachId)
			if err != nil {
				return nil, err
			}
			err = editScheduler(timeZone, req.DaysOfWeek, req.Id, req.Time, req.CronId, req.IdScheduler)
			if err != nil {
				return nil, err
			}
		}
	}

	// wait for gorutine
	return &proto.SubscriptionResponse{
		Id:      req.Id,
		Message: "Subscription edited successfully",
	}, nil
}

func createShodullerFromGroup(timeZone string, daysOfWeek []string, subscriptionId int32, timeHours []string) error {
	conn, err := grpc.Dial("management-service:50005", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := management.NewManagementServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err = client.CreateScheduler(ctx, &management.ManagementRequest{
		TimeZone:       timeZone,
		DaysOfWeek:     daysOfWeek,
		SubscriptionId: subscriptionId,
		Time:           timeHours,
	})

	if err != nil {
		return err
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
func editScheduler(timeZone string, daysOfWeek []string, subscriptionId int32, timeHours []string, cronId int32, idScheduler int32) error {
	conn, err := grpc.Dial("management-service:50005", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := management.NewManagementServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err = client.UpdateScheduler(ctx, &management.ManagementUpdateRequest{
		TimeZone:       timeZone,
		DaysOfWeek:     daysOfWeek,
		SubscriptionId: subscriptionId,
		Time:           timeHours,
		IdCron:         cronId,
		IdScheduler:    idScheduler,
	})
	if err != nil {
		return err
	}
	return nil
}
func deleteScheduler(idScheduler int32, idCron int32) error {
	conn, err := grpc.Dial("management-service:50005", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := management.NewManagementServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err = client.DeleteScheduler(ctx, &management.DeleteRequest{
		IdScheduler: idScheduler,
		IdCron:      idCron,
	})
	if err != nil {
		return err
	}
	return nil
}
