package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Dimoonevs/SportsApp/management-service/internal/repository"
	schedulerElem "github.com/Dimoonevs/SportsApp/management-service/internal/scheduler"
	"github.com/Dimoonevs/SportsApp/management-service/proto/athletes"
	"github.com/Dimoonevs/SportsApp/management-service/proto/management"
	empty "github.com/golang/protobuf/ptypes/empty"
	cronLib "github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

type ManagementService struct {
	Repo repository.Repository
	management.UnimplementedManagementServiceServer
}

var daysMap = map[string]int{
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
	"Sunday":    0,
}

func (m *ManagementService) CreateScheduler(ctx context.Context, req *management.ManagementRequest) (*empty.Empty, error) {
	cron := schedulerElem.GetScheduler().GetCron()

	sliceDay, err := daysOfWeekToSliceInt(req.DaysOfWeek)
	if err != nil {
		return nil, err
	}
	cronId, err := createCron(ctx, sliceDay, req.TimeZone, cron, m.Repo, req.SubscriptionId, req.Time)
	if err != nil {
		return nil, err
	}
	schedulerId, err := m.Repo.SaveSchedulerData(ctx, cronId, sliceDay, req.TimeZone, req.Time)
	if err != nil {
		return nil, err
	}
	err = m.Repo.SaveSubscriptionsSchedulerId(ctx, req.SubscriptionId, schedulerId)
	if err != nil {
		return nil, err
	}
	cron.Start()
	return &empty.Empty{}, nil
}

func (m *ManagementService) LoadCron(ctx context.Context) error {
	cron := schedulerElem.GetScheduler().GetCron()
	schedulerData, err := m.Repo.GetSchedulerData(ctx)
	if err != nil {
		return err
	}
	for _, cronData := range schedulerData {
		IdSubscription, err := m.Repo.GetIdSubscriptionBySchedulerId(ctx, cronData.Id)
		if err != nil {
			return err
		}
		cronId, err := createCron(ctx, cronData.DaysOfWeek, cronData.TimeZone, cron, m.Repo, IdSubscription, cronData.TimeTraining)
		if err != nil {
			return err
		}
		m.Repo.UpdateShedulerData(ctx, cronData.Id, cronId, cronData.DaysOfWeek, cronData.TimeZone, cronData.TimeTraining)
	}
	cron.Start()
	return nil
}

func (m *ManagementService) DeleteScheduler(ctx context.Context, req *management.DeleteRequest) (*empty.Empty, error) {
	cron := schedulerElem.GetScheduler().GetCron()
	cron.Remove(cronLib.EntryID(int(req.IdCron)))
	err := m.Repo.DeleteSchedulerData(ctx, req.IdScheduler)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (m *ManagementService) UpdateScheduler(ctx context.Context, req *management.ManagementUpdateRequest) (*empty.Empty, error) {
	cron := schedulerElem.GetScheduler().GetCron()
	cron.Remove(cronLib.EntryID(int(req.IdCron)))

	sliceDay, err := daysOfWeekToSliceInt(req.DaysOfWeek)
	if err != nil {
		return nil, err
	}
	cronId, err := createCron(ctx, sliceDay, req.TimeZone, cron, m.Repo, req.SubscriptionId, req.Time)
	if err != nil {
		return nil, err
	}
	m.Repo.UpdateShedulerData(ctx, req.IdScheduler, cronId, sliceDay, req.TimeZone, req.Time)
	return &empty.Empty{}, nil
}

func createCron(ctx context.Context, sliceDay []int, timezone string, cron *cronLib.Cron, repo repository.Repository, subscriptionId int32, timeHours []string) (int32, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}
	cronExpression := fmt.Sprintf("CRON_TZ=%s %s %s * * %s ", location, timeHours[1], timeHours[0], transformationsSliceIntInString(sliceDay))
	log.Println(transformationsSliceIntInString(sliceDay))
	id, err := cron.AddFunc(cronExpression, func() {
		err := minusDay(ctx, subscriptionId)
		if err != nil {
			log.Println(err)
		}
		log.Println("OK")
	})
	if err != nil {
		return 0, err
	}
	return int32(id), nil
}

func minusDay(ctx context.Context, id int32) error {
	conn, err := grpc.Dial("athletes-service:50004", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		return err
	}
	defer conn.Close()

	client := athletes.NewAthleteServiceClient(conn)
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	_, err = client.MinusTrainingOfSubscription(ctx, &athletes.MinusTrainingOfSubscriptionRequest{
		SubcriptionId: id,
	})

	if err != nil {
		return err
	}
	return nil

}

func daysOfWeekToSliceInt(days []string) ([]int, error) {
	sliceDay := make([]int, 0)
	for _, day := range days {
		if num, ok := daysMap[day]; ok {
			sliceDay = append(sliceDay, num)
		} else {
			return nil, fmt.Errorf("Wrong day of the week: %s", day)
		}
	}
	return sliceDay, nil
}

func transformationsSliceIntInString(daysOfWeek []int) string {

	daysOfWeekStr := make([]string, len(daysOfWeek))
	for i, day := range daysOfWeek {
		daysOfWeekStr[i] = fmt.Sprint(day)
	}
	daysOfWeekCron := strings.Join(daysOfWeekStr, ",")
	return daysOfWeekCron
}
