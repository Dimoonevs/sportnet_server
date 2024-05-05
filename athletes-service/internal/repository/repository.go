package repository

import (
	"context"
	"database/sql"

	proto "github.com/Dimoonevs/SportsApp/athletes-service/proto/athletes"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AthletesRepository interface {
	MinusDay(ctx context.Context, id int32) error
	CreateAthlete(ctx context.Context, req *proto.AthleteRequest) error
	GetAthletes(id int32) ([]*proto.AthleteRequest, error)
	DeleteAthletes(ctx context.Context, req *proto.DeleteAthletesRequest) error
	EditAthlete(ctx context.Context, req *proto.AthleteRequest) error
	AddTraining(ctx context.Context, req *proto.TrainingData) error
}

type AthletesRepositoryPostgres struct {
	Db    *pgxpool.Pool
	Redis *redis.Client
}

func (a *AthletesRepositoryPostgres) MinusDay(ctx context.Context, id int32) error {
	conn, err := a.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "UPDATE athletes SET days_left = days_left - 1 WHERE subscription_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (a *AthletesRepositoryPostgres) CreateAthlete(ctx context.Context, req *proto.AthleteRequest) error {
	conn, err := a.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	subscriptionId := sql.NullInt64{Int64: int64(req.SubscriptionId), Valid: req.SubscriptionId != 0}
	_, err = conn.Query(ctx, "INSERT INTO athletes (first_name, last_name, subscription_id, group_id, days_left, date_last, status) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		req.FirstName, req.LastName, subscriptionId, req.GroupId, req.DaysLeft, req.DateLast, req.Status)
	if err != nil {
		return err
	}
	return nil
}
func (a *AthletesRepositoryPostgres) GetAthletes(id int32) ([]*proto.AthleteRequest, error) {
	var athletes []*proto.AthleteRequest

	conn, err := a.Db.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), `SELECT id, first_name, last_name, subscription_id, group_id, days_left, date_last, status 
	FROM athletes 
	WHERE group_id = $1 
	ORDER BY (status = 'Active') DESC;
	`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		athlete := &proto.AthleteRequest{}
		if err := rows.Scan(&athlete.Id, &athlete.FirstName, &athlete.LastName, &athlete.SubscriptionId, &athlete.GroupId, &athlete.DaysLeft, &athlete.DateLast, &athlete.Status); err != nil {
			return nil, err
		}
		athletes = append(athletes, athlete)
	}
	return athletes, nil
}

func (a *AthletesRepositoryPostgres) DeleteAthletes(ctx context.Context, req *proto.DeleteAthletesRequest) error {

	conn, err := a.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Query(ctx, "DELETE FROM athletes WHERE id = ANY($1)", req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (a *AthletesRepositoryPostgres) EditAthlete(ctx context.Context, req *proto.AthleteRequest) error {

	conn, err := a.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Query(ctx, "UPDATE athletes SET first_name = $1, last_name = $2, subscription_id = $3, group_id = $4, days_left = $5, date_last = $6, status = $7 WHERE id = $8",
		req.FirstName, req.LastName, req.SubscriptionId, req.GroupId, req.DaysLeft, req.DateLast, req.Status, req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (a *AthletesRepositoryPostgres) AddTraining(ctx context.Context, req *proto.TrainingData) error {

	conn, err := a.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	if req.DaysLeft != 0 {
		_, err = conn.Query(ctx, "UPDATE athletes SET days_left = days_left + $1 WHERE id = $2",
			req.DaysLeft, req.Id)
		if err != nil {
			return err
		}
	} else {
		_, err = conn.Query(ctx, "UPDATE athletes SET date_last = $1 WHERE id = $2",
			req.DateLast, req.Id)
		if err != nil {
			return err
		}
	}
	return nil
}
