package repository

import (
	"context"
	"database/sql"

	proto "github.com/Dimoonevs/SportsApp/typesubscription-service/proto/subscription"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SubscriptionRepository interface {
	CreateSubscription(ctx context.Context, req *proto.SubscriptionRequest) (int32, error)
	GetAllSubscriptions(ctx context.Context, req *proto.GetSubscriptionRequest) ([]*proto.SubscriptionData, error)
	EditSubscription(ctx context.Context, req *proto.SubscriptionEditRequest) error
}

type SubscriptionRepositoryPostgres struct {
	Db *pgxpool.Pool
}

func (s *SubscriptionRepositoryPostgres) CreateSubscription(ctx context.Context, req *proto.SubscriptionRequest) (int32, error) {
	// open transaction
	conn, err := s.Db.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()
	var id int32

	err = conn.QueryRow(ctx, "INSERT INTO subscription (name, description, type_sub, time_limited, custom_time_limited, price, currency, days_of_week, coach_id, automatically_management) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		req.Name, req.Description, req.StatusSubscription.TypeSub, req.StatusSubscription.TimeLimited, req.StatusSubscription.CustomTimeLimited, req.Price.Price, req.Price.Currency, req.DaysOfWeek, req.CoachId, req.AutomaticallyManagement).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SubscriptionRepositoryPostgres) GetAllSubscriptions(ctx context.Context, req *proto.GetSubscriptionRequest) ([]*proto.SubscriptionData, error) {
	conn, err := s.Db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var timeLimited int32
	var typeSub int32
	var subs []*proto.SubscriptionData
	var schedulerId sql.NullInt32
	var cronId sql.NullInt32

	rows, err := conn.Query(ctx, "SELECT subscription.*, scheduler_data.id AS scheduler_id, scheduler_data.cron_id, scheduler_data.time_training FROM subscription LEFT JOIN subscriptions_scheduler ON subscription.id = subscriptions_scheduler.subscription_id LEFT JOIN scheduler_data ON subscriptions_scheduler.scheduler_id = scheduler_data.id;")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		sub := &proto.SubscriptionData{}
		err = rows.Scan(&sub.Id, &sub.Name, &sub.Description, &typeSub, &timeLimited, &sub.CustomTimeLimited, &sub.Price, &sub.Currency, &sub.DaysOfWeek, &sub.AutomaticallyManagement, &sub.CoachId, &schedulerId, &cronId, &sub.Time)
		if err != nil {
			return nil, err
		}

		if schedulerId.Valid {
			sub.IdScheduler = schedulerId.Int32
		}
		if cronId.Valid {
			sub.CronId = cronId.Int32
		}
		sub.TypeSub = proto.TypeSub(typeSub).String()
		sub.TimeLimited = proto.TimeLimited(timeLimited).String()
		subs = append(subs, sub)
	}

	return subs, nil
}

func (s *SubscriptionRepositoryPostgres) EditSubscription(ctx context.Context, req *proto.SubscriptionEditRequest) error {
	conn, err := s.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "UPDATE subscription SET name = $1, description = $2, type_sub = $3, time_limited = $4, custom_time_limited = $5, price = $6, currency = $7, days_of_week = $8, automatically_management = $9 WHERE id = $10", req.Name, req.Description, req.StatusSubscription.TypeSub, req.StatusSubscription.TimeLimited, req.StatusSubscription.CustomTimeLimited, req.Price.Price, req.Price.Currency, req.DaysOfWeek, req.AutomaticallyManagement, req.Id)
	if err != nil {
		return err
	}
	return nil
}
