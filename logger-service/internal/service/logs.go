package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Dimoonevs/SportsApp/logger-service/internal/repository"
	"github.com/Dimoonevs/SportsApp/logger-service/pkg/data"
)

type LogsService struct {
	Repo repository.RepositoryLogs
}

func (s *LogsService) WriteLogs(ctx context.Context, req *http.Request) error {
	var requestPayload data.LogRequest
	err := json.NewDecoder(req.Body).Decode(&requestPayload)
	if err != nil {
		return err
	}

	return s.Repo.WriteLogs(ctx, &data.LogsPayload{
		Data: requestPayload.Data,
		Name: requestPayload.Name,
	}, requestPayload.Field)
}
