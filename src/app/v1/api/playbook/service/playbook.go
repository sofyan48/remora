package service

import (
	"time"

	"github.com/sofyan48/remora/src/app/v1/api/playbook/entity"
)

// PlaybookService ...
type PlaybookService struct{}

// PlaybookServiceHandler ...
func PlaybookServiceHandler() *PlaybookService {
	return &PlaybookService{}
}

// PlaybookServiceInterface ...
type PlaybookServiceInterface interface {
	PlaybookService() *entity.PlaybookResponse
}

// PlaybookService ...
func (service *PlaybookService) PlaybookService() *entity.PlaybookResponse {
	now := time.Now()
	result := &entity.PlaybookResponse{}
	result.Status = "OK"
	result.CreatedAt = &now
	return result
}
