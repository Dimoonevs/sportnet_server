package repository

import (
	"context"

	"github.com/Dimoonevs/SportsApp/management-service/pkg/data"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	SaveSubscriptionsSchedulerId(ctx context.Context, idSubscription int32, idScheduler int32) error
	SaveSchedulerData(ctx context.Context, cronId int32, sliceDay []int, timezone string, timeHours []string) (int32, error)
	GetSchedulerData(ctx context.Context) ([]*data.CronData, error)
	UpdateShedulerData(ctx context.Context, id int32, cronId int32, sliceDay []int, timezone string, timeHours []string) error
	GetIdSubscriptionBySchedulerId(ctx context.Context, idScheduler int32) (int32, error)
	DeleteSchedulerData(ctx context.Context, idScheduler int32) error
}
type RepositoryPostgres struct {
	MongoDB    *mongo.Client
	PostgresDb *pgxpool.Pool
}

type ReqPayloadMongoCronReference struct {
	IDSubscription int32
	IDCron         int32
}

func (r *RepositoryPostgres) SaveSubscriptionsSchedulerId(ctx context.Context, idSubscription int32, idScheduler int32) error {

	conn, err := r.PostgresDb.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "INSERT INTO subscriptions_scheduler (subscription_id, scheduler_id) VALUES ($1, $2)", idSubscription, idScheduler)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryPostgres) SaveSchedulerData(ctx context.Context, cronId int32, sliceDay []int, timezone string, timeHours []string) (int32, error) {
	conn, err := r.PostgresDb.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()
	var id int32
	err = conn.QueryRow(ctx, "INSERT INTO scheduler_data (cron_id, days_of_week, time_zone, time_training) VALUES ($1, $2, $3, $4) RETURNING id", cronId, sliceDay, timezone, timeHours).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (r *RepositoryPostgres) GetSchedulerData(ctx context.Context) ([]*data.CronData, error) {
	conn, err := r.PostgresDb.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, "SELECT id, cron_id, days_of_week, time_zone, time_training FROM scheduler_data")
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	var cronDataList []*data.CronData
	for rows.Next() {
		var cronData data.CronData
		err = rows.Scan(&cronData.Id, &cronData.CronId, &cronData.DaysOfWeek, &cronData.TimeZone, &cronData.TimeTraining)
		if err != nil {
			return nil, err
		}
		cronDataList = append(cronDataList, &cronData)
	}
	return cronDataList, nil
}

func (r *RepositoryPostgres) UpdateShedulerData(ctx context.Context, id int32, cronId int32, sliceDay []int, timezone string, timeHours []string) error {
	conn, err := r.PostgresDb.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "UPDATE scheduler_data SET cron_id = $1, days_of_week = $2, time_zone = $3, time_training = $4 WHERE id = $5", cronId, sliceDay, timezone, timeHours, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *RepositoryPostgres) DeleteSchedulerData(ctx context.Context, idScheduler int32) error {
	conn, err := r.PostgresDb.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "DELETE FROM subscriptions_scheduler WHERE scheduler_id = $1", idScheduler)
	if err != nil {
		return err
	}
	_, err = conn.Exec(ctx, "DELETE FROM scheduler_data WHERE id = $1", idScheduler)
	if err != nil {
		return err
	}
	return nil
}
func (r *RepositoryPostgres) GetIdSubscriptionBySchedulerId(ctx context.Context, idScheduler int32) (int32, error) {

	conn, err := r.PostgresDb.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()
	var id int32
	err = conn.QueryRow(ctx, "SELECT subscription_id FROM subscriptions_scheduler WHERE scheduler_id = $1", idScheduler).Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return id, nil
}
