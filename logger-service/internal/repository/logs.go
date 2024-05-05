package repository

import (
	"context"
	"time"

	"github.com/Dimoonevs/SportsApp/logger-service/pkg/data"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryLogs interface {
	WriteLogs(ctx context.Context, req *data.LogsPayload, field string) error
}

type RepositoryLogsMongo struct {
	Mongo *mongo.Client
}

func (r *RepositoryLogsMongo) WriteLogs(ctx context.Context, req *data.LogsPayload, field string) error {
	req.Time = time.Now().Format("2006-01-02 15:04:05")
	_, err := r.Mongo.Database("logs").Collection(field).InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
