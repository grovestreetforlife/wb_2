package api

import (
	"dev11/internal/logger"
	"dev11/internal/service"
	"dev11/internal/storage/inmemory"
)

func InitializeServer() (*server, error) {
	storage := inmemory.NewStorage()
	calendarService := service.NewCalendarService(storage)
	logrusLogger := logger.NewLogger()
	apiServer := newServer(calendarService, storage, logrusLogger)
	return apiServer, nil
}
