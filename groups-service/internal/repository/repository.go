package repository

import (
	"context"
	"database/sql"
	"log"

	proto "github.com/Dimoonevs/SportAspp/groups-service/proto/groups"
	"github.com/jackc/pgx/v4/pgxpool"
)

type GroupRepository interface {
	CreateGroup(ctx context.Context, req *proto.GroupRequest) (int32, error)
	GetGroups(ctx context.Context, req *proto.GetGroupRequest) ([]*proto.GroupData, error)
	EditGroup(ctx context.Context, req *proto.GroupEditRequest) error
	ChangeGroupsSubscriptions(ctx context.Context, id int32, subscriptionId int32) error
}

type GroupRepositoryPortgres struct {
	DB *pgxpool.Pool
}

func (g *GroupRepositoryPortgres) CreateGroup(ctx context.Context, req *proto.GroupRequest) (int32, error) {
	// open transaction
	conn, err := g.DB.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()
	var id int32
	// subscriptionId := sql.NullInt64{Int64: int64(req.SubscriptionId), Valid: req.SubscriptionId != 0}
	err = conn.QueryRow(ctx, "INSERT INTO groups (name, coach_id) VALUES ($1, $2) RETURNING id", req.Name, req.CoachId).Scan(&id)
	if err != nil {
		return 0, err
	}
	if req.SubscriptionId != 0 {
		_, err = conn.Query(ctx, "INSERT INTO groups_subscriptions (group_id, subscription_id) VALUES ($1, $2)", id, req.SubscriptionId)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (g *GroupRepositoryPortgres) GetGroups(ctx context.Context, req *proto.GetGroupRequest) ([]*proto.GroupData, error) {
	conn, err := g.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `
		SELECT 
		    groups.id AS group_id,
		    groups.name AS group_name,
		    groups.coach_id,
		    subscription.id AS subscription_id,
		    subscription.name AS subscription_name,
		    subscription.time_limited,
		    subscription.custom_time_limited
		FROM 
		    groups
		LEFT JOIN 
		    groups_subscriptions ON groups.id = groups_subscriptions.group_id
		LEFT JOIN 
		    subscription ON groups_subscriptions.subscription_id = subscription.id
		WHERE 
		    groups.coach_id = $1`, req.CoachId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*proto.GroupData
	var subscription struct {
		Id                sql.NullInt32
		Name              sql.NullString
		TimeLimited       sql.NullInt32
		CustomTimeLimited sql.NullInt32
	}
	for rows.Next() {
		group := &proto.GroupData{}
		err = rows.Scan(&group.Id, &group.Name, &group.CoachId, &subscription.Id, &subscription.Name, &subscription.TimeLimited, &subscription.CustomTimeLimited)
		if err != nil {
			return nil, err
		}
		if subscription.Id.Valid {
			log.Println("subscription valid")
			group.SubscriptionId = int32(subscription.Id.Int32)
			group.SubscriptionName = subscription.Name.String
			group.TimeLimited = int32(subscription.TimeLimited.Int32)
			group.CustomTimeLimited = int32(subscription.CustomTimeLimited.Int32)
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (g *GroupRepositoryPortgres) EditGroup(ctx context.Context, req *proto.GroupEditRequest) error {
	conn, err := g.DB.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Query(ctx, "UPDATE groups SET name = $1 WHERE id = $2", req.Name, req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (g *GroupRepositoryPortgres) ChangeGroupsSubscriptions(ctx context.Context, id int32, subscriptionId int32) error {
	conn, err := g.DB.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	if subscriptionId != 0 {
		_, err = conn.Query(ctx, `INSERT INTO groups_subscriptions (group_id, subscription_id)
		VALUES ($1, $2)
		ON CONFLICT (group_id) DO UPDATE SET subscription_id = $2`, id, subscriptionId)
		if err != nil {
			return err
		}
	} else {
		_, err = conn.Query(ctx, "DELETE FROM groups_subscriptions WHERE group_id = $1", id)
		if err != nil {
			return err
		}
	}
	return nil
}
